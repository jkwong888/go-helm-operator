package libertyapp

import (
	"context"
	"fmt"
	"strings"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/helm/pkg/kube"

	libertyv1alpha1 "github.com/jkwong888/websphere-liberty-operator/pkg/apis/liberty/v1alpha1"
	"github.com/jkwong888/websphere-liberty-operator/pkg/image"
	"github.com/jkwong888/websphere-liberty-operator/pkg/internal/util/diffutil"
	"github.com/jkwong888/websphere-liberty-operator/pkg/release"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	rpb "k8s.io/helm/pkg/proto/hapi/release"
)

const (
	finalizer = "uninstall-helm-release"
)

var _ reconcile.Reconciler = &ReconcileLibertyApp{}

// ReleaseHookFunc defines a function signature for release hooks.
type ReleaseHookFunc func(*rpb.Release) error

// ReconcileLibertyApp reconciles a LibertyApp object
type ReconcileLibertyApp struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	KubeClient      *kube.Client
	Client          client.Client
	scheme          *runtime.Scheme
	ManagerFactory  release.ManagerFactory
	ReconcilePeriod time.Duration
	releaseHook     ReleaseHookFunc
}

// Reconcile reads that state of the cluster for a LibertyApp object and makes changes based on the state read
// and what is in the LibertyApp.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileLibertyApp) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling LibertyApp")

	// Fetch the LibertyApp instance
	instance := &libertyv1alpha1.LibertyApp{}
	err := r.Client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	oMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	o := &unstructured.Unstructured{
		Object: oMap,
	}

	log := log.WithValues(
		"namespace", o.GetNamespace(),
		"name", o.GetName(),
		"apiVersion", o.GetAPIVersion(),
		"kind", o.GetKind(),
	)

	manager := r.ManagerFactory.NewManager(o)
	if err != nil {
		log.Error(err, "Failed to get release manager")
		return reconcile.Result{}, err
	}

	status := release.StatusFor(o)
	log = log.WithValues("release", manager.ReleaseName())

	deleted := o.GetDeletionTimestamp() != nil
	pendingFinalizers := o.GetFinalizers()
	if !deleted && !contains(pendingFinalizers, finalizer) {
		log.V(1).Info("Adding finalizer", "finalizer", finalizer)
		finalizers := append(pendingFinalizers, finalizer)
		o.SetFinalizers(finalizers)
		err = r.updateResource(o)

		// Need to requeue because finalizer update does not change metadata.generation
		return reconcile.Result{Requeue: true}, err
	}

	status.SetCondition(release.AppCondition{
		Type:   release.ConditionInitialized,
		Status: release.StatusTrue,
	})

	if err := manager.Sync(context.TODO(), r.transformRelease); err != nil {
		log.Error(err, "Failed to sync release")
		status.SetCondition(release.AppCondition{
			Type:    release.ConditionIrreconcilable,
			Status:  release.StatusTrue,
			Reason:  release.ReasonReconcileError,
			Message: err.Error(),
		})
		_ = r.updateResourceStatus(o, status)
		return reconcile.Result{}, err
	}
	status.RemoveCondition(release.ConditionIrreconcilable)

	if deleted {
		if !contains(pendingFinalizers, finalizer) {
			log.Info("Resource is terminated, skipping reconciliation")
			return reconcile.Result{}, nil
		}

		uninstalledRelease, err := manager.UninstallRelease(context.TODO())
		if err != nil && err != release.ErrNotFound {
			log.Error(err, "Failed to uninstall release")
			status.SetCondition(release.AppCondition{
				Type:    release.ConditionReleaseFailed,
				Status:  release.StatusTrue,
				Reason:  release.ReasonUninstallError,
				Message: err.Error(),
			})
			_ = r.updateResourceStatus(o, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(release.ConditionReleaseFailed)

		if err == release.ErrNotFound {
			log.Info("Release not found, removing finalizer")
		} else {
			log.Info("Uninstalled release")
			if log.Enabled() {
				fmt.Println(diffutil.Diff(uninstalledRelease.GetManifest(), ""))
			}
			status.SetCondition(release.AppCondition{
				Type:   release.ConditionDeployed,
				Status: release.StatusFalse,
				Reason: release.ReasonUninstallSuccessful,
			})
		}
		if err := r.updateResourceStatus(o, status); err != nil {
			return reconcile.Result{}, err
		}

		finalizers := []string{}
		for _, pendingFinalizer := range pendingFinalizers {
			if pendingFinalizer != finalizer {
				finalizers = append(finalizers, pendingFinalizer)
			}
		}
		o.SetFinalizers(finalizers)
		err = r.updateResource(o)

		// Need to requeue because finalizer update does not change metadata.generation
		return reconcile.Result{Requeue: true}, err
	}

	if !manager.IsInstalled() {
		installedRelease, err := manager.InstallRelease(context.TODO())
		if err != nil {
			log.Error(err, "Failed to install release")
			status.SetCondition(release.AppCondition{
				Type:    release.ConditionReleaseFailed,
				Status:  release.StatusTrue,
				Reason:  release.ReasonInstallError,
				Message: err.Error(),
				Release: installedRelease,
			})
			_ = r.updateResourceStatus(o, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(release.ConditionReleaseFailed)

		if r.releaseHook != nil {
			if err := r.releaseHook(installedRelease); err != nil {
				log.Error(err, "Failed to run release hook")
				return reconcile.Result{}, err
			}
		}

		log.Info("Installed release")

		if log.Enabled() {
			fmt.Println(diffutil.Diff("", installedRelease.GetManifest()))
		}
		log.V(1).Info("Config values", "values", installedRelease.GetConfig())
		status.SetCondition(release.AppCondition{
			Type:    release.ConditionDeployed,
			Status:  release.StatusTrue,
			Reason:  release.ReasonInstallSuccessful,
			Message: installedRelease.GetInfo().GetStatus().GetNotes(),
			Release: installedRelease,
		})

		err = r.updateResourceStatus(o, status)
		return reconcile.Result{RequeueAfter: r.ReconcilePeriod}, err

	}

	if manager.IsUpdateRequired() {
		previousRelease, updatedRelease, err := manager.UpdateRelease(context.TODO())
		if err != nil {
			log.Error(err, "Failed to update release")
			status.SetCondition(release.AppCondition{
				Type:    release.ConditionReleaseFailed,
				Status:  release.StatusTrue,
				Reason:  release.ReasonUpdateError,
				Message: err.Error(),
				Release: updatedRelease,
			})
			_ = r.updateResourceStatus(o, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(release.ConditionReleaseFailed)

		if r.releaseHook != nil {
			if err := r.releaseHook(updatedRelease); err != nil {
				log.Error(err, "Failed to run release hook")
				return reconcile.Result{}, err
			}
		}

		log.Info("Updated release")
		if log.Enabled() {
			fmt.Println(diffutil.Diff(previousRelease.GetManifest(), updatedRelease.GetManifest()))
		}
		log.V(1).Info("Config values", "values", updatedRelease.GetConfig())
		status.SetCondition(release.AppCondition{
			Type:    release.ConditionDeployed,
			Status:  release.StatusTrue,
			Reason:  release.ReasonUpdateSuccessful,
			Message: updatedRelease.GetInfo().GetStatus().GetNotes(),
			Release: updatedRelease,
		})
		err = r.updateResourceStatus(o, status)
		return reconcile.Result{RequeueAfter: r.ReconcilePeriod}, err
	}

	expectedRelease, err := manager.ReconcileRelease(context.TODO())
	if err != nil {
		log.Error(err, "Failed to reconcile release")
		status.SetCondition(release.AppCondition{
			Type:    release.ConditionIrreconcilable,
			Status:  release.StatusTrue,
			Reason:  release.ReasonReconcileError,
			Message: err.Error(),
		})
		_ = r.updateResourceStatus(o, status)
		return reconcile.Result{}, err
	}
	status.RemoveCondition(release.ConditionIrreconcilable)

	if r.releaseHook != nil {
		if err := r.releaseHook(expectedRelease); err != nil {
			log.Error(err, "Failed to run release hook")
			return reconcile.Result{}, err
		}
	}

	log.Info("Reconciled release")
	err = r.updateResourceStatus(o, status)

	return reconcile.Result{RequeueAfter: r.ReconcilePeriod}, err
}

func (r *ReconcileLibertyApp) transformRelease(objectMap map[string]runtime.Object) (map[string]runtime.Object, error) {
	toReturn := make(map[string]runtime.Object)

	clientset, err := r.KubeClient.KubernetesClientSet()
	if err != nil {
		log.Error(err, "Failed to get kubernetes clientset")
	}

	for k, v := range objectMap {
		// find the Deployment
		if !strings.HasPrefix(k, "Deployment") {
			toReturn[k] = v
			continue
		}

		// convert the runtime.Object to unstructured.Unstructured
		unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(v)
		if err != nil {
			return nil, err
		}

		// convert to Deployment
		deployment := &appsv1.Deployment{}
		err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredObj, deployment)
		if err != nil {
			return nil, err
		}

		namespace := deployment.ObjectMeta.GetNamespace()

		// find the images
		containers := deployment.Spec.Template.Spec.Containers
		var newContainers []corev1.Container
		for _, container := range containers {
			newContainer := container.DeepCopy()
			// get the image repo and tag
			img, err := image.NewLibertyAppImage(clientset, container.Image, namespace, nil)
			if err != nil {
				return nil, err
			}

			digest, err := img.GetDigest(context.TODO())
			if err != nil {
				return nil, err
			}

			// this is a small hack to get the image to be split the way we like, so the new
			// repository is repo@sha256 and the tag is abcdefg123456
			imageRepo := strings.Split(container.Image, ":")

			// to grab an image by digest from the image repository, use the repo@sha256:abcdefg123456
			newImageName := fmt.Sprintf("%s@%s", imageRepo[0], *digest)
			newContainer.Image = newImageName
			newContainers = append(newContainers, *newContainer)
		}

		deployment.Spec.Template.Spec.Containers = newContainers
		toReturn[k] = deployment.DeepCopyObject()
	}

	return toReturn, nil
}

func (r ReconcileLibertyApp) updateResource(o *unstructured.Unstructured) error {
	return r.Client.Update(context.TODO(), o)
}

func (r ReconcileLibertyApp) updateResourceStatus(o *unstructured.Unstructured, status *release.AppStatus) error {
	o.Object["status"] = status
	return r.Client.Status().Update(context.TODO(), o)
}

func contains(l []string, s string) bool {
	for _, elem := range l {
		if elem == s {
			return true
		}
	}
	return false
}
