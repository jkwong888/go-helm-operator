package libertyapp

import (
	"os"

	"github.com/operator-framework/operator-sdk/pkg/k8sutil"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"k8s.io/helm/pkg/storage"
  "k8s.io/helm/pkg/storage/driver"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"

	"github.com/operator-framework/operator-sdk/pkg/helm/client"
	helmController "github.com/operator-framework/operator-sdk/pkg/helm/controller"
	release "github.com/operator-framework/operator-sdk/pkg/helm/release"

)

var log = logf.Log.WithName("controller_libertyapp")

// Add creates a new LibertyApp Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
// TODO: for now, liberty is just a helm release so just use the helm operator.
// if we need to do other fancy stuff for liberty we can add in a second reconciler


namespace, err := k8sutil.GetWatchNamespace()
if err != nil {
	log.Error(err, "Failed to get watch namespace")
	os.Exit(1)
}

// Create Tiller's storage backend and kubernetes client
storageBackend := storage.Init(driver.NewMemory())
tillerKubeClient, err := client.NewFromManager(mgr)
if err != nil {
	log.Error(err, "Failed to create new Tiller client.")
	return err
}

//	err := add(mgr, newReconciler(mgr))
//
//	if (err != nil) {
//		return err
//	}
	helmWatchOptions := helmController.WatchOptions{
		Namespace: namespace,
		GVK: schema.GroupVersionKind{
			Group: "liberty.ibm.com",
			Version: "v1alpha1",
			Kind: "LibertyApp",
		},
		ManagerFactory: release.NewManagerFactory(storageBackend, tillerKubeClient, "/opt/helm/charts"),
		ReconcilePeriod: 5, //TODO
		WatchDependentResources: true,
	}

  // add helm controller watcher
	return helmController.Add(mgr, helmWatchOptions)
}
