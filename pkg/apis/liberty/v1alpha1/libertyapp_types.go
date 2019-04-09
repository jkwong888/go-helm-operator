package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LibertyApp defines the desired state of LibertyApp
// +k8s:openapi-gen=true
type LibertyApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   LibertyAppSpec `json:"spec"`
	Status AppStatus      `json:"status,omitempty"`
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
