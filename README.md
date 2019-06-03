# helm-operator-liberty

Operator for Websphere Liberty apps in Kubernetes, using the operator-sdk, generated from helm chart

# Operator functions

- deploy the [ibm-websphere-liberty helm chart](https://github.com/IBM/charts/tree/master/stable/ibm-websphere-liberty) when `LibertyApp` CR is created
  - The CR `spec` field contains `values.yaml` properties that can be specified as if installing the helm chart
- patch `default` service account in a namespace with an image pull secret specified by user
  - must be pre-created before installation
  - specify the secret name in `image.pullSecret` in the `spec` section of the CR.
- poll the container image registry for the app image and roll out the new deployment when the digest changes

# Building the Operator

We used the [Operator SDK](https://github.com/operator-framework/operator-sdk) to generate the skeleton

```
# in $GOPATH, create the liberty operator
~/go/bin/operator-sdk new websphere-liberty-operator

# add the CRD
~/go/bin/operator-sdk add api --api-version=liberty.ibm.com/v1alpha1 --kind=LibertyApp

# add the controller:
~/go/bin/operator-sdk add controller --api-version liberty.ibm.com/v1alpha1 --kind=LibertyApp

```

The `operator-sdk` will generate an `interface{}` struct called `LibertyAppSpec`.  There are various scripts in `hack` directory to generate code for this out of `values.yaml`, which builds the necessary nested structs in order to build go code that can reconcile the CR.

- `hack/libertyappspec_gen.sh` will generate go structs from the liberty chart's `values.yaml`
  - Specifically in our operator, `image.pullSecret` is not originally part of the chart and I added this by hand, so the resulting `pkg/apis/liberty/v1alpha1/libertyappspec.go` was hand modified to include this field after this step
- `hack/deepcopy_gen.sh` will generate deep copy functions for the go structs
- `hack/gen_from_liberty_chart.sh` wraps the above functions; it downloads the latest liberty chart, calls `hack/libertyappspec_gen.sh` to generate go structs, and uses `hack/deepcopy_gen.sh` to generate the deep copy functions needed for the operator sdk
- `hack/run_local.sh` will use build the operator and run it (assuming you've unzipped the chart in `/opt/helm/charts`) reusing the current kube context
- `hack/deploy_operator.sh` will use a golang container to build and deploy the operator, roles, and CRDs using the current kube context

# Sample CR

Here is a sample CR I used for testing which uses the sample plantsbywebsphere.

```
apiVersion: liberty.ibm.com/v1alpha1
kind: LibertyApp
metadata:
  name: plants
spec:
  image:
    repository: registry.jkwong.cloudns.cx/jkwong/plantsbywebsphere
    tag: "latest"
    pullSecret: "my-pull-secret"
  ssl:
    enabled: false 
  service:
    port: 9080
    targetPort: 9080
```

# TODO

The following ideas to enhance the operator:

- fix the operator RBAC so that the permissions are streamlined
- I borrowed heavily from the helm operator code in [operator-sdk](https://github.com/operator-framework/operator-sdk/tree/master/pkg/helm), which uses an embedded tiller.  I would like to reduce the dependency on embedded Tiller, just render the templates and manage them manually afterward.
- We could also remove the embedded helm release in the `status` section of the CR which is very large and doesn't allow easy edits in text editor. (`operator-sdk` v0.8.1 addresses this but it seems the full manifest is still in status section)
- Monitor configmaps for changes and roll the deployments to update
  - configmaps with a certain annotation can be pulled into the deployment and mounted into `/config/configDropins`
  - for GitOps use cases, we can monitor configmaps directly from git as well
- Dynamically build and deploy app images from a war file stored in a Maven repo
