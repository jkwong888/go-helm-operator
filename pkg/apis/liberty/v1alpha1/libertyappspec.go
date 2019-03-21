package v1alpha1

type LibertyAppSpec struct {
	// TODO: fill in the liberty chart values.
	// better if we can generate this based on chart version
  image string `json:"image"`
}
