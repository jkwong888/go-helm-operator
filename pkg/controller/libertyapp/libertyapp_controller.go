package libertyapp

import (
	libertyv1alpha1 "github.com/jkwong888/websphere-liberty-operator/pkg/apis/liberty/v1alpha1"

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
	"github.com/operator-framework/operator-sdk/pkg/helm/release"
)

var log = logf.Log.WithName("controller_libertyapp")

// Add creates a new LibertyApp Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	// Create Tiller's storage backend and kubernetes client
	storageBackend := storage.Init(driver.NewMemory())
	tillerKubeClient, err := client.NewFromManager(mgr)
	if err != nil {
		log.Error(err, "Failed to create new Tiller client.")
		return err
	}

	return add(mgr, newReconciler(mgr, storageBackend, tillerKubeClient))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager, storageBackend *storage.Storage, tillerKubeClient *kube.Client) reconcile.Reconciler {
	return &ReconcileLibertyApp{
		Client:         mgr.GetClient(),
		scheme:         mgr.GetScheme(),
		ManagerFactory: release.NewManagerFactory(storageBackend, tillerKubeClient, "/opt/helm/charts"),
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
