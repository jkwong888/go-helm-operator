package image

import (
	"context"
	"fmt"

	"github.com/containers/image/manifest"
	"github.com/containers/image/transports/alltransports"
	"github.com/containers/image/types"

	"k8s.io/client-go/kubernetes"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("libertyappimage")

// LibertyAppImage represents a liberty app image on a remote registry
type LibertyAppImage struct {
	imageName      string
	systemContext  *types.SystemContext
	ImageReference types.ImageReference
}

// newSystemContext returns a *types.SystemContext corresponding to opts.
// It is guaranteed to return a fresh instance, so it is safe to make additional updates to it.
func newSystemContext() *types.SystemContext {
	ctx := &types.SystemContext{}
	return ctx
}

// NewLibertyAppImage Create a new LibertyAppImage
func NewLibertyAppImage(clientset *kubernetes.Clientset, imageName string, namespace string, secretName *string) (*LibertyAppImage, error) {
	log.WithValues("image", imageName)

	img := &LibertyAppImage{
		imageName:     imageName,
		systemContext: newSystemContext(),
	}

	if secretName != nil {
		log.WithValues("secret", secretName)

		secret, err := getDockerSecret(clientset, namespace, *secretName)
		if err != nil {
			return nil, err
		}

		dockerAuth, err := getDockerAuth(imageName, secret)
		if err != nil {
			return nil, err
		}

		img.systemContext.DockerAuthConfig = dockerAuth
	}

	ref, err := alltransports.ParseImageName(fmt.Sprintf("docker://%s", img.imageName))
	if err != nil {
		return nil, err
	}

	img.ImageReference = ref

	return img, nil
}

// GetDigest get the repo digest of the image
func (img *LibertyAppImage) GetDigest(ctx context.Context) (*string, error) {
	imgCloser, err := img.ImageReference.NewImage(ctx, img.systemContext)
	if err != nil {
		return nil, err
	}
	defer imgCloser.Close()

	//rawManifest, _, err := img.Manifest(ctx)
	rawManifest, _, err := imgCloser.Manifest(ctx)
	if err != nil {
		panic(err)
	}

	digest, err := manifest.Digest(rawManifest)
	if err != nil {
		panic(err)
	}

	str := digest.String()

	return &str, nil
}
