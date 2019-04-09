package libertyapp

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/operator-framework/operator-sdk/pkg/k8sutil"

	libertyv1alpha1 "github.com/jkwong888/websphere-liberty-operator/pkg/apis/liberty/v1alpha1"
	"github.com/jkwong888/websphere-liberty-operator/pkg/controller/options"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"k8s.io/helm/pkg/kube"
	"k8s.io/helm/pkg/storage"
	"k8s.io/helm/pkg/storage/driver"

	"github.com/operator-framework/operator-sdk/pkg/helm/client"
	"github.com/operator-framework/operator-sdk/pkg/helm/release"
)

var log = logf.Log.WithName("controller_libertyapp")

// namespace returns the namespace of tiller
func namespace() string {
	namespace, found := os.LookupEnv(k8sutil.WatchNamespaceEnvVar)
	if found {
		return namespace
	}

	// Fall back to the namespace associated with the service account token, if available
	if data, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace"); err == nil {
		if ns := strings.TrimSpace(string(data)); len(ns) > 0 {
			return ns
		}
	}

	return metav1.NamespaceAll
}

// Add creates a new LibertyApp Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	// Create Tiller's storage backend and kubernetes client
	tillerKubeClient, err := client.NewFromManager(mgr)
	if err != nil {
		log.Error(err, "Failed to create new Tiller client.")
		return err
	}

	clientset, err := tillerKubeClient.KubernetesClientSet()
	if err != nil {
		log.Error(err, "Failed to create storage client.")
		return err
	}
	storageBackend := storage.Init(driver.NewSecrets(clientset.CoreV1().Secrets(namespace())))

	return add(mgr, newReconciler(mgr, storageBackend, tillerKubeClient))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager, storageBackend *storage.Storage, tillerKubeClient *kube.Client) reconcile.Reconciler {
	time, err := time.ParseDuration(options.ReconcilePeriod)
	if err != nil {
		log.Error(err, fmt.Sprintf("Failed to parse duration: %s.", options.ReconcilePeriod))
		return nil
	}

	return &ReconcileLibertyApp{
		Client:          mgr.GetClient(),
		scheme:          mgr.GetScheme(),
		ManagerFactory:  release.NewManagerFactory(storageBackend, tillerKubeClient, "/opt/helm/charts"),
		ReconcilePeriod: time,
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("libertyapp-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource LibertyApp
	err = c.Watch(&source.Kind{Type: &libertyv1alpha1.LibertyApp{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner LibertyApp
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &libertyv1alpha1.LibertyApp{},
	})
	if err != nil {
		return err
	}

	return nil
}
