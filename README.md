# websphere-liberty-operator

Operator for Websphere Liberty apps in Kubernetes

# Operator functions

- deploy the [ibm-websphere-liberty helm chart](https://github.com/IBM/charts/tree/master/stable/ibm-websphere-liberty) when `LibertyApp` CR is created
  - The CR `spec` field contains `values.yaml` properties that can be specified as if installing the helm chart
- patch `default` service account in a namespace with an image pull secret specified by user
  - must be pre-created before installation
  - specify the secret name in `image.pullSecret` in the `spec` section of the CR.
- poll the container image registry for the app image and roll out the new deployment when the digest changes

# Building the Operator

There are various scripts in `hack` directory to generate code.

- `hack/libertyappspec_gen.sh` will generate go structs from the liberty chart's `values.yaml`
  - Note that `image.pullSecret` is not originally part of the chart and I added this by hand, so the resulting `pkg/apis/liberty/v1alpha1/libertyappspec.go` was hand modified to include this field after this step
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
- use a separate service account for the running deployment instead of `default`
- I borrowed heavily from the helm operator code in [operator-sdk](https://github.com/operator-framework/operator-sdk/tree/master/pkg/helm), which uses an embedded tiller.  I would like to reduce the dependency on embedded Tiller, just render the templates and manage them manually afterward.
- We could also remove the embedded helm release in the `status` section of the CR which is very large and doesn't allow easy edits in text editor.
- Monitor configmaps for changes and roll the deployments to update
- Dynamically build and deploy app images from a war file stored in a Maven repo
- Monitor git for configmap changes and pull them into the namespace when they change (and roll the deployment)