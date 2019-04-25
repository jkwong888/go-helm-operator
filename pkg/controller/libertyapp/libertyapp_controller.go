package libertyapp

import (
	"fmt"
	"time"

	"k8s.io/client-go/tools/clientcmd"

	libertyv1alpha1 "github.com/jkwong888/websphere-liberty-operator/pkg/apis/liberty/v1alpha1"
	"github.com/jkwong888/websphere-liberty-operator/pkg/controller/options"
	"github.com/jkwong888/websphere-liberty-operator/pkg/release"

	corev1 "k8s.io/api/core/v1"
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
)

var log = logf.Log.WithName("controller_libertyapp")

// namespace returns the namespace of tiller
func namespace() (string, error) {
	// Get a config to talk to the apiserver
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	namespace, _, err := kubeconfig.Namespace()
	if err != nil {
		log.Error(err, "Unable to retrieve current namespace.")
		return "", err
	}

	return namespace, nil

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

	namespace, err := namespace()
	if err != nil {
		log.Error(err, "Failed to get storage client context.")
		return err
	}

	storageBackend := storage.Init(driver.NewSecrets(clientset.CoreV1().Secrets(namespace)))

	return add(mgr, newReconciler(mgr, tillerKubeClient, storageBackend))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager, kubeClient *kube.Client, storageBackend *storage.Storage) reconcile.Reconciler {
	time, err := time.ParseDuration(options.ReconcilePeriod)
	if err != nil {
		log.Error(err, fmt.Sprintf("Failed to parse duration: %s.", options.ReconcilePeriod))
		return nil
	}

	return &ReconcileLibertyApp{
		KubeClient:      kubeClient,
		Client:          mgr.GetClient(),
		scheme:          mgr.GetScheme(),
		ManagerFactory:  release.NewManagerFactory(storageBackend, kubeClient, "/opt/helm/charts"),
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
