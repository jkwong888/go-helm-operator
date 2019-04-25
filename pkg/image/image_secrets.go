package image

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiTypes "k8s.io/apimachinery/pkg/types"

	"github.com/containers/image/types"
	"k8s.io/client-go/kubernetes"
)

func parseCreds(creds64 string) (string, string, error) {
	if creds64 == "" {
		return "", "", errors.New("credentials can't be empty")
	}

	// base64 decode the creds
	creds, err := base64.StdEncoding.DecodeString(creds64)
	if err != nil {
		return "", "", fmt.Errorf("Error decoding creds: %s", err.Error())
	}

	up := strings.SplitN(string(creds), ":", 2)
	if len(up) == 1 {
		return up[0], "", nil
	}
	if up[0] == "" {
		return "", "", errors.New("username can't be empty")
	}
	return up[0], up[1], nil
}

func getDockerSecret(clientset *kubernetes.Clientset, namespace string, secretname string) (*corev1.Secret, error) {
	secret, err := clientset.CoreV1().Secrets(namespace).Get(secretname, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	if secret == nil {
		return nil, fmt.Errorf("Could not get secret, %s does it exist?", secretname)
	}

	return secret, nil
}

func getDockerAuth(imageName string, secret *corev1.Secret) (*types.DockerAuthConfig, error) {
	b := secret.Data[".dockerconfigjson"]

	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}

	registryURL := strings.SplitN(imageName, "/", 2)[0]
	if m["auths"] == nil {
		log.V(1).Info("Did not find auths key in imagePullSecret", "secret", secret.ObjectMeta.GetName())
		return &types.DockerAuthConfig{}, nil
	}
	authMap := m["auths"].(map[string]interface{})

	if authMap[registryURL] == nil {
		log.V(1).Info("Did not find auth for registry in imagePullSecret", "secret", secret.ObjectMeta.GetName(), "registry", registryURL)
		return &types.DockerAuthConfig{}, nil
	}

	val := authMap[registryURL].(map[string]interface{})
	if val == nil {
		log.V(1).Info("Did not find auth for registry in imagePullSecret", "secret", secret.ObjectMeta.GetName(), "registry", registryURL)
		return &types.DockerAuthConfig{}, nil
	}

	// get the "auths" key
	username, password, err := parseCreds(val["auth"].(string))
	if err != nil {
		return nil, err
	}

	return &types.DockerAuthConfig{
		Username: username,
		Password: password,
	}, nil
}

// AddPullSecretToServiceAccount add pull secret to service account
func AddPullSecretToServiceAccount(clientset *kubernetes.Clientset, namespace string, serviceaccount string, secretName *string) error {
	if secretName == nil {
		// nothing to do
		return nil
	}

	log.V(1).Info("Adding pull secret to service account",
		"serviceaccount", serviceaccount,
		"imagePullSecret", *secretName,
		"namespace", namespace)

	sa, err := clientset.CoreV1().ServiceAccounts(namespace).Get(serviceaccount, metav1.GetOptions{})
	if err != nil {
		log.Error(err,
			"Unable to find serviceaccount",
			"serviceaccount", serviceaccount,
			"namespace", namespace)

		return err
	}

	_, err = clientset.CoreV1().Secrets(namespace).Get(*secretName, metav1.GetOptions{})
	if err != nil {
		log.Error(err,
			"Unable to find imagePullSecret",
			"imagePullSecret", *secretName,
			"namespace", namespace)

		return err
	}

	existingPullSecrets := sa.ImagePullSecrets

	existingSecretStr := ""
	for _, existingSecret := range existingPullSecrets {
		if existingSecret.Name == *secretName {
			log.V(1).Info("Service account already has pull secret",
				"serviceaccount", serviceaccount,
				"imagePullSecret", secretName,
				"namespace", namespace)
			return nil
		}

		existingSecretStr += fmt.Sprintf(`{"name": %q}, `, existingSecret.Name)
	}

	newSecrets := fmt.Sprintf(`{"imagePullSecrets": [ %s{"name": %q } ] }`, existingSecretStr, *secretName)
	payloadBytes := []byte(newSecrets)
	newSA, err := clientset.CoreV1().
		ServiceAccounts(namespace).
		Patch(serviceaccount, apiTypes.MergePatchType, payloadBytes)

	if err != nil {
		log.Error(err,
			"Unable to patch service account with new pull secret",
			"serviceaccount", serviceaccount,
			"imagePullSecret", *secretName,
			"namespace", namespace)

		return err
	}

	log.V(1).Info("Patched service account",
		"serviceaccount", newSA)

	return nil
}

// AddPullSecretToDeployment add pull secret to deployment
func AddPullSecretToDeployment(clientset *kubernetes.Clientset, namespace string, deploymentName string, secretName *string) error {
	if secretName == nil {
		// nothing to do
		return nil
	}

	log.V(1).Info("Adding pull secret to deployment",
		"deployment", deploymentName,
		"imagePullSecret", *secretName,
		"namespace", namespace)

	deploy, err := clientset.AppsV1().Deployments(namespace).Get(deploymentName, metav1.GetOptions{})
	if err != nil {
		log.Error(err,
			"Unable to find deployment",
			"deployment", deploymentName,
			"namespace", namespace)

		return err
	}

	_, err = clientset.CoreV1().Secrets(namespace).Get(*secretName, metav1.GetOptions{})
	if err != nil {
		log.Error(err,
			"Unable to find imagePullSecret",
			"imagePullSecret", *secretName,
			"namespace", namespace)

		return err
	}

	existingPullSecrets := deploy.Spec.Template.Spec.ImagePullSecrets
	for _, existingSecret := range existingPullSecrets {
		if existingSecret.Name == *secretName {
			log.V(1).Info("Deployment already has pull secret",
				"deployment", deploymentName,
				"imagePullSecret", secretName,
				"namespace", namespace)
			return nil
		}
	}

	newPullSecrets := append(existingPullSecrets,
		corev1.LocalObjectReference{
			Name: *secretName,
		})

	// patch the service account with the pull secret
	payload := map[string]interface{}{
		"op":    "replace",
		"path":  "/spec/template/spec/imagePullSecrets",
		"value": newPullSecrets,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Error(err,
			"Unable to marshal payload for patch",
			"deployment", deploymentName,
			"imagePullSecret", *secretName,
			"namespace", namespace)

		return err
	}

	newSA, err := clientset.CoreV1().
		ServiceAccounts(namespace).
		Patch(deploymentName, apiTypes.JSONPatchType, payloadBytes)

	if err != nil {
		log.Error(err,
			"Unable to patch service account with new pull secret",
			"deployment", deploymentName,
			"imagePullSecret", *secretName,
			"namespace", namespace)

		return err
	}

	log.V(1).Info("Patched deployment",
		"deployment", newSA)

	return nil
}
