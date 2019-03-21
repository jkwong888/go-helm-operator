package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/helm/pkg/proto/hapi/release"
)

// LibertyAppSpec defines the desired state of LibertyApp
// +k8s:openapi-gen=true
type LibertyApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              LibertyAppSpec   `json:"spec"`
	Status            LibertyAppStatus `json:"status,omitempty"`
}

type LibertyAppConditionType string
type ConditionStatus string
type LibertyAppConditionReason string

type LibertyAppCondition struct {
	Type    LibertyAppConditionType   `json:"type"`
	Status  ConditionStatus        `json:"status"`
	Reason  LibertyAppConditionReason `json:"reason,omitempty"`
	Message string                 `json:"message,omitempty"`
	Release *release.Release       `json:"release,omitempty"`

	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

const (
	ConditionInitialized    LibertyAppConditionType = "Initialized"
	ConditionDeployed       LibertyAppConditionType = "Deployed"
	ConditionReleaseFailed  LibertyAppConditionType = "ReleaseFailed"
	ConditionIrreconcilable LibertyAppConditionType = "Irreconcilable"

	StatusTrue    ConditionStatus = "True"
	StatusFalse   ConditionStatus = "False"
	StatusUnknown ConditionStatus = "Unknown"

	ReasonInstallSuccessful   LibertyAppConditionReason = "InstallSuccessful"
	ReasonUpdateSuccessful    LibertyAppConditionReason = "UpdateSuccessful"
	ReasonUninstallSuccessful LibertyAppConditionReason = "UninstallSuccessful"
	ReasonInstallError        LibertyAppConditionReason = "InstallError"
	ReasonUpdateError         LibertyAppConditionReason = "UpdateError"
	ReasonReconcileError      LibertyAppConditionReason = "ReconcileError"
	ReasonUninstallError      LibertyAppConditionReason = "UninstallError"
)

type LibertyAppStatus struct {
	Conditions []LibertyAppCondition `json:"conditions"`
}

func (s *LibertyAppStatus) ToMap() (map[string]interface{}, error) {
	var out map[string]interface{}
	jsonObj, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsonObj, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SetCondition sets a condition on the status object. If the condition already
// exists, it will be replaced. SetCondition does not update the resource in
// the cluster.
func (s *LibertyAppStatus) SetCondition(condition LibertyAppCondition) *LibertyAppStatus {
	now := metav1.Now()
	for i := range s.Conditions {
		if s.Conditions[i].Type == condition.Type {
			if s.Conditions[i].Status != condition.Status {
				condition.LastTransitionTime = now
			} else {
				condition.LastTransitionTime = s.Conditions[i].LastTransitionTime
			}
			s.Conditions[i] = condition
			return s
		}
	}

	// If the condition does not exist,
	// initialize the lastTransitionTime
	condition.LastTransitionTime = now
	s.Conditions = append(s.Conditions, condition)
	return s
}

// RemoveCondition removes the condition with the passed condition type from
// the status object. If the condition is not already present, the returned
// status object is returned unchanged. RemoveCondition does not update the
// resource in the cluster.
func (s *LibertyAppStatus) RemoveCondition(conditionType LibertyAppConditionType) *LibertyAppStatus {
	for i := range s.Conditions {
		if s.Conditions[i].Type == conditionType {
			s.Conditions = append(s.Conditions[:i], s.Conditions[i+1:]...)
			return s
		}
	}
	return s
}

// StatusFor safely returns a typed status block from a custom resource.
func StatusFor(cr *unstructured.Unstructured) *LibertyAppStatus {
	switch s := cr.Object["status"].(type) {
	case *LibertyAppStatus:
		return s
	case map[string]interface{}:
		var status *LibertyAppStatus
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(s, &status); err != nil {
			return &LibertyAppStatus{}
		}
		return status
	default:
		return &LibertyAppStatus{}
	}
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
