package main

import (
	"context"
	"os"

	"github.com/jkwong888/websphere-liberty-operator/pkg/image"
	"github.com/spf13/pflag"

	"github.com/operator-framework/operator-sdk/pkg/log/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var (
	//  Leave blank for the default context in your kube config.
	kubecontext = ""
)

var log = logf.Log.WithName("cmd")

func main() {
	// Add the zap logger flag set to the CLI. The flag set must
	// be added before calling pflag.Parse().
	pflag.CommandLine.AddFlagSet(zap.FlagSet())

	// Use a zap logr.Logger implementation. If none of the zap
	// flags are configured (or if the zap flag set is not being
	// used), this defaults to a production zap logger.
	//
	// The logger instantiated here can be changed to any logger
	// implementing the logr.Logger interface. This logger will
	// be propagated through the whole operator, generating
	// uniform and structured logs.
	logf.SetLogger(zap.Logger())

	ctx := context.Background()

	imageName := os.Args[1]

	clientconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{CurrentContext: kubecontext})
	config, err := clientconfig.ClientConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	rawConfig, err := clientconfig.RawConfig()
	if err != nil {
		panic(err)
	}
	namespace := rawConfig.Contexts[rawConfig.CurrentContext].Namespace

	log.Info("connecting to ", "host", config.Host)

	secret := "my-pull-secret"
	img, err := image.NewLibertyAppImage(clientset, imageName, namespace, &secret)
	if err != nil {
		log.Error(err, "Error")
		panic(err)
	}

	digest, err := img.GetDigest(ctx)
	if err != nil {
		panic(err)
	}

	log.Info("Digest is ", "digest", *digest)

}
