// Copyright 2018 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// TODO shameless cut and paste from operator-sdk, for the most part

package release

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"

	yaml "gopkg.in/yaml.v2"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	apitypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/cli-runtime/pkg/genericclioptions/resource"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/kube"
	cpb "k8s.io/helm/pkg/proto/hapi/chart"
	rpb "k8s.io/helm/pkg/proto/hapi/release"
	"k8s.io/helm/pkg/proto/hapi/services"
	"k8s.io/helm/pkg/storage"
	"k8s.io/helm/pkg/tiller"

	"github.com/jkwong888/websphere-liberty-operator/pkg/internal/util/yamlutil"
	"github.com/mattbaird/jsonpatch"
)

var (
	// ErrNotFound indicates the release was not found.
	ErrNotFound = errors.New("release not found")
)

type transformFunc func(objectMap map[string]runtime.Object) (map[string]runtime.Object, error)

// Manager manages a Helm release. It can install, update, reconcile,
// and uninstall a release.
type Manager interface {
	ReleaseName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context, transformFunc) error
	InstallRelease(context.Context) (*rpb.Release, error)
	UpdateRelease(context.Context) (*rpb.Release, *rpb.Release, error)
	ReconcileRelease(context.Context) (*rpb.Release, error)
	UninstallRelease(context.Context) (*rpb.Release, error)
}

type manager struct {
	storageBackend   *storage.Storage
	tillerKubeClient *kube.Client
	chartDir         string

	//TODO hang on to the groupResources of the cluster so we don't get them every loop

	tiller      *tiller.ReleaseServer
	releaseName string
	namespace   string

	spec   interface{}
	status *AppStatus

	isInstalled      bool
	isUpdateRequired bool
	deployedRelease  *rpb.Release
	chart            *cpb.Chart
	config           *cpb.Config

	pendingRelease *rpb.Release
	objects        map[string]runtime.Object
}

// ReleaseName returns the name of the release.
func (m manager) ReleaseName() string {
	return m.releaseName
}

func (m manager) IsInstalled() bool {
	return m.isInstalled
}

func (m manager) IsUpdateRequired() bool {
	return m.isUpdateRequired
}

// Sync ensures the Helm storage backend is in sync with the status of the
// custom resource.
func (m *manager) Sync(ctx context.Context, transform transformFunc) error {
	if err := m.syncReleaseStatus(*m.status); err != nil {
		return fmt.Errorf("failed to sync release status to storage backend: %s", err)
	}

	// Get release history for this release name
	releases, err := m.storageBackend.History(m.releaseName)
	if err != nil && !notFoundErr(err) {
		return fmt.Errorf("failed to retrieve release history: %s", err)
	}

	// Cleanup non-deployed release versions. If all release versions are
	// non-deployed, this will ensure that failed installations are correctly
	// retried.
	for _, rel := range releases {
		if rel.GetInfo().GetStatus().GetCode() != rpb.Status_DEPLOYED {
			_, err := m.storageBackend.Delete(rel.GetName(), rel.GetVersion())
			if err != nil && !notFoundErr(err) {
				return fmt.Errorf("failed to delete stale release version: %s", err)
			}
		}
	}

	// Load the chart and config based on the current state of the custom resource.
	chart, config, err := m.loadChartAndConfig()
	if err != nil {
		return fmt.Errorf("failed to load chart and config: %s", err)
	}
	m.chart = chart
	m.config = config

	// Load the most recently deployed release from the storage backend.
	deployedRelease, err := m.getDeployedRelease()
	if err == ErrNotFound {
		// render the chart here
		m.pendingRelease, err = renderRelease(ctx, m.tiller, m.tillerKubeClient, m.namespace, m.releaseName, chart, config)
		if err != nil {
			return fmt.Errorf("failed to get candidate release: %s", err)
		}

		// get the objects that will be created as part of the chart
		objects, err := getReleaseObjects(m.pendingRelease)
		if err != nil {
			return fmt.Errorf("failed to render candidate release: %s", err)
		}

		// allow the caller to transform generated objects
		if transform != nil {
			m.objects, err = transform(objects)
			if err != nil {
				return fmt.Errorf("failed to call transform function: %s", err)
			}
		}

		return nil
	}

	if err != nil {
		return fmt.Errorf("failed to get deployed release: %s", err)
	}
	m.deployedRelease = deployedRelease
	m.isInstalled = true

	// Get the next candidate release to determine if an update is necessary.
	m.pendingRelease, err = getCandidateRelease(ctx, m.tiller, m.releaseName, chart, config)
	if err != nil {
		return fmt.Errorf("failed to get candidate release: %s", err)
	}

	// get the objects that will be created as part of the chart
	objects, err := getReleaseObjects(m.pendingRelease)
	if err != nil {
		return fmt.Errorf("failed to render candidate release: %s", err)
	}

	// allow the caller to transform generated objects
	if transform != nil {
		m.objects, err = transform(objects)
		if err != nil {
			return fmt.Errorf("failed to call transform function: %s", err)
		}
	}

	// get the deployed objects from last release
	deployedObjects, err := getReleaseObjects(deployedRelease)
	if err != nil {
		return fmt.Errorf("failed to render deployed release: %s", err)
	}

	// compare the two object maps; we'll do this with maps
	/*
		if !reflect.DeepEqual(m.objects, deployedObjects) {
			m.isUpdateRequired = true
		}
	*/

	for k, v := range deployedObjects {
		pendingObj := m.objects[k]

		if pendingObj == nil {
			// not in the pending release, object was deleted
			m.isUpdateRequired = true
			break
		}

		// compare the two objects, if the objects aren't equal, update required
		if !reflect.DeepEqual(pendingObj, v) {
			m.isUpdateRequired = true
			break
		}
	}

	// check for new objects
	for k := range m.objects {
		if deployedObjects[k] != nil {
			// existing objects were handled in the above loop
			continue
		}

		// at this point, the object is new in this release and must be created
		m.isUpdateRequired = true
		break
	}

	return nil
}

func (m manager) syncReleaseStatus(status AppStatus) error {
	var release *rpb.Release
	for _, condition := range status.Conditions {
		if condition.Type == ConditionDeployed && condition.Status == StatusTrue {
			release = condition.Release
			break
		}
	}
	if release == nil {
		return nil
	}

	name := release.GetName()
	version := release.GetVersion()
	_, err := m.storageBackend.Get(name, version)
	if err == nil {
		return nil
	}

	if !notFoundErr(err) {
		return err
	}
	return m.storageBackend.Create(release)
}

func notFoundErr(err error) bool {
	return strings.Contains(err.Error(), "not found")
}

func (m manager) loadChartAndConfig() (*cpb.Chart, *cpb.Config, error) {
	// chart is mutated by the call to processRequirements,
	// so we need to reload it from disk every time.
	chart, err := chartutil.LoadDir(m.chartDir)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load chart: %s", err)
	}

	cr, err := yaml.Marshal(m.spec)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse values: %s", err)
	}
	config := &cpb.Config{Raw: string(cr)}

	err = processRequirements(chart, config)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to process chart requirements: %s", err)
	}
	return chart, config, nil
}

// processRequirements will process the requirements file
// It will disable/enable the charts based on condition in requirements file
// Also imports the specified chart values from child to parent.
func processRequirements(chart *cpb.Chart, values *cpb.Config) error {
	err := chartutil.ProcessRequirementsEnabled(chart, values)
	if err != nil {
		return err
	}
	err = chartutil.ProcessRequirementsImportValues(chart)
	if err != nil {
		return err
	}
	return nil
}

func (m manager) getDeployedRelease() (*rpb.Release, error) {
	deployedRelease, err := m.storageBackend.Deployed(m.releaseName)
	if err != nil {
		if strings.Contains(err.Error(), "has no deployed releases") {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return deployedRelease, nil
}

func getCandidateRelease(ctx context.Context, tiller *tiller.ReleaseServer, name string, chart *cpb.Chart, config *cpb.Config) (*rpb.Release, error) {
	dryRunReq := &services.UpdateReleaseRequest{
		Name:   name,
		Chart:  chart,
		Values: config,
		DryRun: true,
	}
	dryRunResponse, err := tiller.UpdateRelease(ctx, dryRunReq)
	if err != nil {
		return nil, err
	}
	return dryRunResponse.GetRelease(), nil
}

func renderRelease(ctx context.Context,
	tiller *tiller.ReleaseServer,
	tillerKubeClient *kube.Client,
	namespace string,
	name string,
	chart *cpb.Chart,
	config *cpb.Config) (*rpb.Release, error) {
	installReq := &services.InstallReleaseRequest{
		Namespace: namespace,
		Name:      name,
		Chart:     chart,
		Values:    config,
		DryRun:    true,
	}

	releaseResponse, err := tiller.InstallRelease(ctx, installReq)
	if err != nil {
		// Workaround for helm/helm#3338
		if releaseResponse.GetRelease() != nil {
			uninstallReq := &services.UninstallReleaseRequest{
				Name:  releaseResponse.GetRelease().GetName(),
				Purge: true,
			}
			_, uninstallErr := tiller.UninstallRelease(ctx, uninstallReq)
			if uninstallErr != nil {
				return nil, fmt.Errorf("failed to roll back failed installation: %s: %s", uninstallErr, err)
			}
		}
		return nil, err
	}

	return releaseResponse.Release, nil
}

// InstallRelease performs a Helm release install.
func (m manager) InstallRelease(ctx context.Context) (*rpb.Release, error) {
	return installRelease(ctx, m.tillerKubeClient, m.storageBackend, m.pendingRelease, m.objects)
}

func installRelease(ctx context.Context,
	kubeClient *kube.Client,
	storageBackend *storage.Storage,
	release *rpb.Release,
	objects map[string]runtime.Object) (*rpb.Release, error) {

	// write the resulting objects to the API server one by one
	// TODO: here we may wish to allow the controller to sort or prune the objects
	// so it can handle dependencies or other such special cases
	dynamicClient, err := kubeClient.DynamicClient()
	if err != nil {
		return nil, err
	}

	clientset, err := kubeClient.KubernetesClientSet()
	if err != nil {
		return nil, err
	}

	// TODO this shouldn't change too often so maybe it's ok to cache this
	groupResources, err := restmapper.GetAPIGroupResources(clientset.Discovery())
	if err != nil {
		return nil, err
	}

	rm := restmapper.NewDiscoveryRESTMapper(groupResources)

	// track all the installed objects
	var installed []unstructured.Unstructured

	for _, v := range objects {
		gvk := v.GetObjectKind().GroupVersionKind()
		gk := schema.GroupKind{Group: gvk.Group, Kind: gvk.Kind}
		mapping, err := rm.RESTMapping(gk, gvk.Version)

		// convert the runtime.Object to unstructured.Unstructured
		oMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(v)
		if err != nil {
			return nil, err
		}
		unstructuredObj := &unstructured.Unstructured{
			Object: oMap,
		}

		// create this object
		_, err = dynamicClient.Resource(mapping.Resource).Namespace(release.GetNamespace()).Create(unstructuredObj, metav1.CreateOptions{})
		if err != nil {
			if apierrors.IsAlreadyExists(err) {
				log.Info("Error creating resource, already exists", "namespace", release.GetNamespace(), "gvk", gvk, "release", release.GetName(), "name", unstructuredObj.GetName())
				installed = append(installed, *unstructuredObj)
				continue
			} else {
				log.Error(err, "Error creating resource", "namespace", release.GetNamespace(), "gvk", gvk, "release", release.GetName(), "name", unstructuredObj.GetName())
				return nil, err
			}
		}

		// track the installed object and add it to the manifest later
		installed = append(installed, *unstructuredObj)

		log.Info("Created resource", "namespace", release.GetNamespace(), "gvk", gvk, "release", release.GetName(), "name", unstructuredObj.GetName())
	}

	// hack, just set the release as deployed so the sync code doesn't try to remove it
	release.Info.Status.Code = rpb.Status_DEPLOYED

	// take the installed objects and write them to the release
	release.Manifest, err = toManifest(installed)
	if err != nil {
		return nil, err
	}

	err = storageBackend.Create(release)
	if err != nil {
		log.Error(err, "Error writing release to storage backend", "namespace", release.GetNamespace(), "name", release.GetName())
	}

	return release, nil
}

// UpdateRelease performs a Helm release update.
func (m manager) UpdateRelease(ctx context.Context) (*rpb.Release, *rpb.Release, error) {
	updatedRelease, err := updateRelease(ctx, m.tillerKubeClient, m.storageBackend, m.deployedRelease, m.pendingRelease, m.objects)
	return m.deployedRelease, updatedRelease, err
}

func updateRelease(ctx context.Context,
	kubeClient *kube.Client,
	storageBackend *storage.Storage,
	deployedRelease *rpb.Release,
	pendingRelease *rpb.Release,
	pendingObjects map[string]runtime.Object) (*rpb.Release, error) {
	// track all the updated objects
	var updated []unstructured.Unstructured

	// TODO: here we may wish to allow the controller to sort or prune the objects
	// so it can handle dependencies or other such special cases
	dynamicClient, err := kubeClient.DynamicClient()
	if err != nil {
		return nil, err
	}

	clientset, err := kubeClient.KubernetesClientSet()
	if err != nil {
		return nil, err
	}

	// TODO this shouldn't change too often so maybe it's ok to cache this
	groupResources, err := restmapper.GetAPIGroupResources(clientset.Discovery())
	if err != nil {
		return nil, err
	}

	rm := restmapper.NewDiscoveryRESTMapper(groupResources)

	// get the deployed objects from last release
	deployedObjects, err := getReleaseObjects(deployedRelease)
	if err != nil {
		return nil, fmt.Errorf("failed to render deployed release: %s", err)
	}

	for k, v := range deployedObjects {
		gvk := v.GetObjectKind().GroupVersionKind()
		gk := schema.GroupKind{Group: gvk.Group, Kind: gvk.Kind}
		mapping, err := rm.RESTMapping(gk, gvk.Version)
		pendingObj := pendingObjects[k]
		objectNameArr := strings.Split(k, "/")

		if pendingObj == nil {
			// not in the pending release, object was deleted
			err = dynamicClient.Resource(mapping.Resource).Namespace(deployedRelease.GetNamespace()).Delete(objectNameArr[1], &metav1.DeleteOptions{})
			if err != nil {
				return nil, fmt.Errorf("failed to deleted resource: %s", err)
			}

			log.Info("Deleted resource", "namespace", deployedRelease.GetNamespace(), "gvk", gvk, "release", deployedRelease.GetName(), "name", objectNameArr[1])
			continue
		}

		// convert the runtime.Object to unstructured.Unstructured
		oMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(v)
		if err != nil {
			return nil, err
		}

		unstructuredObj := &unstructured.Unstructured{
			Object: oMap,
		}

		// compare the two objects, if the objects aren't equal, update it
		if !reflect.DeepEqual(pendingObj, v) {
			_, err = dynamicClient.Resource(mapping.Resource).Namespace(deployedRelease.GetNamespace()).Update(unstructuredObj, metav1.UpdateOptions{})
			if err != nil {
				return nil, fmt.Errorf("failed to update resource: %s", err)
			}

			// add this object to the release
			updated = append(updated, *unstructuredObj)

			continue
		}

		// this object didn't change, but is still part of our release
		updated = append(updated, *unstructuredObj)
	}

	// check for new objects
	for k, v := range pendingObjects {
		if deployedObjects[k] != nil {
			// existing objects were handled in the above loop
			continue
		}

		// at this point, the object is new in this release and must be created

		gvk := v.GetObjectKind().GroupVersionKind()
		gk := schema.GroupKind{Group: gvk.Group, Kind: gvk.Kind}
		mapping, err := rm.RESTMapping(gk, gvk.Version)

		oMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(v)
		if err != nil {
			return nil, err
		}

		unstructuredObj := &unstructured.Unstructured{
			Object: oMap,
		}

		// new object
		_, err = dynamicClient.Resource(mapping.Resource).Namespace(deployedRelease.GetNamespace()).Create(unstructuredObj, metav1.CreateOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to create resource: %s", err)
		}

		updated = append(updated, *unstructuredObj)
	}

	releaseVersion := deployedRelease.GetVersion() + 1
	pendingRelease.Version = releaseVersion

	// set the release as deployed
	pendingRelease.Info.Status.Code = rpb.Status_DEPLOYED
	// take the installed objects and write them to the release
	pendingRelease.Manifest, err = toManifest(updated)
	if err != nil {
		return nil, err
	}

	err = storageBackend.Create(pendingRelease)
	if err != nil {
		log.Error(err, "Error writing release to storage backend", "namespace", pendingRelease.GetNamespace(), "name", pendingRelease.GetName())
		return nil, err
	}

	// set the old release to superseded
	deployedRelease.Info.Status.Code = rpb.Status_SUPERSEDED
	err = storageBackend.Update(deployedRelease)
	if err != nil {
		log.Error(err, "Error updating release to storage backend", "namespace", deployedRelease.GetNamespace(), "name", deployedRelease.GetName())
		return nil, err
	}

	return pendingRelease, nil
}

// ReconcileRelease creates or patches resources as necessary to match the
// deployed release's manifest.
func (m manager) ReconcileRelease(ctx context.Context) (*rpb.Release, error) {
	err := reconcileRelease(ctx, m.tillerKubeClient, m.namespace, m.deployedRelease.GetManifest())
	return m.deployedRelease, err
}

func reconcileRelease(ctx context.Context, tillerKubeClient *kube.Client, namespace string, expectedManifest string) error {
	expectedInfos, err := tillerKubeClient.BuildUnstructured(namespace, bytes.NewBufferString(expectedManifest))
	if err != nil {
		return err
	}
	return expectedInfos.Visit(func(expected *resource.Info, err error) error {
		if err != nil {
			return err
		}

		expectedClient := resource.NewClientWithOptions(expected.Client, func(r *rest.Request) {
			*r = *r.Context(ctx)
		})
		helper := resource.NewHelper(expectedClient, expected.Mapping)

		existing, err := helper.Get(expected.Namespace, expected.Name, false)
		if apierrors.IsNotFound(err) {
			if _, err := helper.Create(expected.Namespace, true, expected.Object, &metav1.CreateOptions{}); err != nil {
				return fmt.Errorf("create error: %s", err)
			}
			return nil
		} else if err != nil {
			return err
		}

		patch, err := generatePatch(existing, expected.Object)
		if err != nil {
			return fmt.Errorf("failed to marshal JSON patch: %s", err)
		}

		if patch == nil {
			return nil
		}

		_, err = helper.Patch(expected.Namespace, expected.Name, apitypes.JSONPatchType, patch, &metav1.UpdateOptions{})
		if err != nil {
			return fmt.Errorf("patch error: %s", err)
		}
		return nil
	})
}

func generatePatch(existing, expected runtime.Object) ([]byte, error) {
	existingJSON, err := json.Marshal(existing)
	if err != nil {
		return nil, err
	}
	expectedJSON, err := json.Marshal(expected)
	if err != nil {
		return nil, err
	}

	ops, err := jsonpatch.CreatePatch(existingJSON, expectedJSON)
	if err != nil {
		return nil, err
	}

	// We ignore the "remove" operations from the full patch because they are
	// fields added by Kubernetes or by the user after the existing release
	// resource has been applied. The goal for this patch is to make sure that
	// the fields managed by the Helm chart are applied.
	patchOps := make([]jsonpatch.JsonPatchOperation, 0)
	for _, op := range ops {
		if op.Operation != "remove" {
			patchOps = append(patchOps, op)
		}
	}

	// If there are no patch operations, return nil. Callers are expected
	// to check for a nil response and skip the patch operation to avoid
	// unnecessary chatter with the API server.
	if len(patchOps) == 0 {
		return nil, nil
	}

	return json.Marshal(patchOps)
}

// UninstallRelease performs a Helm release uninstall.
func (m manager) UninstallRelease(ctx context.Context) (*rpb.Release, error) {
	return uninstallRelease(ctx, m.storageBackend, m.tiller, m.releaseName)
}

func uninstallRelease(ctx context.Context, storageBackend *storage.Storage, tiller *tiller.ReleaseServer, releaseName string) (*rpb.Release, error) {
	// Get history of this release
	h, err := storageBackend.History(releaseName)
	if err != nil {
		return nil, fmt.Errorf("failed to get release history: %s", err)
	}

	// If there is no history, the release has already been uninstalled,
	// so return ErrNotFound.
	if len(h) == 0 {
		return nil, ErrNotFound
	}

	uninstallResponse, err := tiller.UninstallRelease(ctx, &services.UninstallReleaseRequest{
		Name:  releaseName,
		Purge: true,
	})
	return uninstallResponse.GetRelease(), err
}

func getReleaseObjects(release *rpb.Release) (map[string]runtime.Object, error) {
	objMap := make(map[string]runtime.Object)

	yamls, err := yamlutil.SplitYaml(release.Manifest)
	if err != nil {
		return nil, err
	}

	for _, y := range yamls {
		// parse the manifest into individual objects
		decode := scheme.Codecs.UniversalDeserializer().Decode
		obj, _, err := decode([]byte(y), nil, nil)
		if err != nil {
			return nil, err
		}

		// convert the runtime.Object to unstructured.Unstructured
		oMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			return nil, err
		}

		unstructuredObj := &unstructured.Unstructured{
			Object: oMap,
		}

		mapKey := unstructuredObj.GetKind() + "/" + unstructuredObj.GetName()
		objMap[mapKey] = obj
	}

	return objMap, nil
}

func toManifest(objects []unstructured.Unstructured) (string, error) {
	intfArr := make([]interface{}, len(objects))
	for i, v := range objects {
		intfArr[i] = v.Object
	}
	return yamlutil.ToYaml(intfArr)
}
