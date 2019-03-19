package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LibertyAppSpec defines the desired state of LibertyApp
// +k8s:openapi-gen=true
type LibertyAppSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              HelmAppSpec   `json:"spec"`
	Status            HelmAppStatus `json:"status,omitempty"`
}

type HelmAppSpec map[string]interface{}

// LibertyAppStatus defines the observed state of LibertyApp
// +k8s:openapi-gen=true
type LibertyAppStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LibertyApp is the Schema for the libertyapps API
// +k8s:openapi-gen=true
type LibertyApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LibertyAppSpec   `json:"spec,omitempty"`
	Status LibertyAppStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LibertyAppList contains a list of LibertyApp
type LibertyAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LibertyApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LibertyApp{}, &LibertyAppList{})
}
