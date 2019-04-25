package libertyapp

import (
	"context"
	"fmt"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

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

	cpb "k8s.io/helm/pkg/proto/hapi/chart"
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

	// TODO: here we're taking advantage of

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

	status := libertyv1alpha1.StatusFor(o)
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

	secretName := instance.Spec.Image.PullSecret
	clientset, err := r.KubeClient.KubernetesClientSet()
	if err != nil {
		log.Error(err, "Failed to get kube client")
	}

	// TODO: would prefer the liberty chart creates a new service account instead of using "default"
	// for now we'll patch the "default" service account with the pull secret here
	if secretName != nil {
		err = image.AddPullSecretToServiceAccount(clientset, request.Namespace, "default", secretName)
		if err != nil {
			log.Error(err, "Failed to patch service account")
			return reconcile.Result{}, err
		}
	}

	status.SetCondition(libertyv1alpha1.AppCondition{
		Type:   libertyv1alpha1.ConditionInitialized,
		Status: libertyv1alpha1.StatusTrue,
	})

	if err := manager.Sync(context.TODO(), r.transformRelease); err != nil {
		log.Error(err, "Failed to sync release")
		status.SetCondition(libertyv1alpha1.AppCondition{
			Type:    libertyv1alpha1.ConditionIrreconcilable,
			Status:  libertyv1alpha1.StatusTrue,
			Reason:  libertyv1alpha1.ReasonReconcileError,
			Message: err.Error(),
		})
		_ = r.updateResourceStatus(o, status)
		return reconcile.Result{}, err
	}
	status.RemoveCondition(libertyv1alpha1.ConditionIrreconcilable)

	if deleted {
		if !contains(pendingFinalizers, finalizer) {
			log.Info("Resource is terminated, skipping reconciliation")
			return reconcile.Result{}, nil
		}

		uninstalledRelease, err := manager.UninstallRelease(context.TODO())
		if err != nil && err != release.ErrNotFound {
			log.Error(err, "Failed to uninstall release")
			status.SetCondition(libertyv1alpha1.AppCondition{
				Type:    libertyv1alpha1.ConditionReleaseFailed,
				Status:  libertyv1alpha1.StatusTrue,
				Reason:  libertyv1alpha1.ReasonUninstallError,
				Message: err.Error(),
			})
			_ = r.updateResourceStatus(o, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(libertyv1alpha1.ConditionReleaseFailed)

		if err == release.ErrNotFound {
			log.Info("Release not found, removing finalizer")
		} else {
			log.Info("Uninstalled release")
			if log.Enabled() {
				fmt.Println(diffutil.Diff(uninstalledRelease.GetManifest(), ""))
			}
			status.SetCondition(libertyv1alpha1.AppCondition{
				Type:   libertyv1alpha1.ConditionDeployed,
				Status: libertyv1alpha1.StatusFalse,
				Reason: libertyv1alpha1.ReasonUninstallSuccessful,
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
			status.SetCondition(libertyv1alpha1.AppCondition{
				Type:    libertyv1alpha1.ConditionReleaseFailed,
				Status:  libertyv1alpha1.StatusTrue,
				Reason:  libertyv1alpha1.ReasonInstallError,
				Message: err.Error(),
				Release: installedRelease,
			})
			_ = r.updateResourceStatus(o, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(libertyv1alpha1.ConditionReleaseFailed)

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
		status.SetCondition(libertyv1alpha1.AppCondition{
			Type:    libertyv1alpha1.ConditionDeployed,
			Status:  libertyv1alpha1.StatusTrue,
			Reason:  libertyv1alpha1.ReasonInstallSuccessful,
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
			status.SetCondition(libertyv1alpha1.AppCondition{
				Type:    libertyv1alpha1.ConditionReleaseFailed,
				Status:  libertyv1alpha1.StatusTrue,
				Reason:  libertyv1alpha1.ReasonUpdateError,
				Message: err.Error(),
				Release: updatedRelease,
			})
			_ = r.updateResourceStatus(o, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(libertyv1alpha1.ConditionReleaseFailed)

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
		status.SetCondition(libertyv1alpha1.AppCondition{
			Type:    libertyv1alpha1.ConditionDeployed,
			Status:  libertyv1alpha1.StatusTrue,
			Reason:  libertyv1alpha1.ReasonUpdateSuccessful,
			Message: updatedRelease.GetInfo().GetStatus().GetNotes(),
			Release: updatedRelease,
		})
		err = r.updateResourceStatus(o, status)
		return reconcile.Result{RequeueAfter: r.ReconcilePeriod}, err
	}

	expectedRelease, err := manager.ReconcileRelease(context.TODO())
	if err != nil {
		log.Error(err, "Failed to reconcile release")
		status.SetCondition(libertyv1alpha1.AppCondition{
			Type:    libertyv1alpha1.ConditionIrreconcilable,
			Status:  libertyv1alpha1.StatusTrue,
			Reason:  libertyv1alpha1.ReasonReconcileError,
			Message: err.Error(),
		})
		_ = r.updateResourceStatus(o, status)
		return reconcile.Result{}, err
	}
	status.RemoveCondition(libertyv1alpha1.ConditionIrreconcilable)

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

func (r *ReconcileLibertyApp) transformRelease(namespace string, chart *cpb.Chart, config *cpb.Config) error {
	clientset, err := r.KubeClient.KubernetesClientSet()
	if err != nil {
		log.Error(err, "Failed to get kube client")
	}

	defaultMap := make(map[string]interface{})
	err = yaml.Unmarshal([]byte(chart.Values.Raw), &defaultMap)
	if err != nil {
		log.Error(err, "Failed to unmarshal raw values")
	}

	defaultImageSpec := defaultMap["image"].(map[interface{}]interface{})
	imageRepo := defaultImageSpec["repository"]
	imageTag := defaultImageSpec["tag"]
	var secretName *string

	/* parse config.Raw to get the overridden image repo and tag */
	valueMap := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(config.Raw), &valueMap)
	if err != nil {
		log.Error(err, "Failed to unmarshal raw values")
	}

	imageSpec := valueMap["image"].(map[interface{}]interface{})
	if imageSpec != nil {
		imageRepo = imageSpec["repository"]
		imageTag = imageSpec["tag"]
		secretNameTmp := imageSpec["pullSecret"].(string)
		secretName = &secretNameTmp
	}

	imageName := fmt.Sprintf("%s:%s", imageRepo, imageTag)

	img, err := image.NewLibertyAppImage(clientset, imageName, namespace, secretName)
	if err != nil {
		log.Error(err, "Failed to get image", "image", imageName)
	}

	// get the image digest, which may look like sha256:abcdefg12356....
	digest, err := img.GetDigest(context.TODO())
	if err != nil {
		log.Error(err, "Failed to get image digest", "image", imageName)
	}

	// to grab an image by digest from the image repository, use the repo@sha256:abcdefg123456
	newImageName := fmt.Sprintf("%s@%s", imageRepo, *digest)

	// this is a small hack to get the image to be split the way we like, so the new
	// repository is repo@sha256 and the tag is abcdefg123456
	newImgSplit := strings.Split(newImageName, ":")

	if imageSpec == nil {
		imageSpec = make(map[interface{}]interface{})
	}

	imageSpec["repository"] = newImgSplit[0]
	imageSpec["tag"] = newImgSplit[1]
	valueMap["image"] = imageSpec

	out, err := yaml.Marshal(valueMap)
	if err != nil {
		log.Error(err, "Failed to Marshal updated values")
	}

	config.Raw = string(out)

	log.V(1).Info("Updating chart parameters with liberty image digest",
		"LibertyAppImage", imageName,
		"Digest", digest)

	return nil
}

func (r ReconcileLibertyApp) updateResource(o *unstructured.Unstructured) error {
	return r.Client.Update(context.TODO(), o)
}

func (r ReconcileLibertyApp) updateResourceStatus(o *unstructured.Unstructured, status *libertyv1alpha1.AppStatus) error {
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
