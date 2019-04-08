package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/helm/pkg/proto/hapi/release"
)

type AppConditionType string
type AppConditionStatus string
type AppConditionReason string

type AppCondition struct {
	Type    AppConditionType   `json:"type"`
	Status  AppConditionStatus `json:"status"`
	Reason  AppConditionReason `json:"reason,omitempty"`
	Message string             `json:"message,omitempty"`
	Release *release.Release   `json:"release,omitempty"`

	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

const (
	ConditionInitialized    AppConditionType = "Initialized"
	ConditionDeployed       AppConditionType = "Deployed"
	ConditionReleaseFailed  AppConditionType = "ReleaseFailed"
	ConditionIrreconcilable AppConditionType = "Irreconcilable"

	StatusTrue    AppConditionStatus = "True"
	StatusFalse   AppConditionStatus = "False"
	StatusUnknown AppConditionStatus = "Unknown"

	ReasonInstallSuccessful   AppConditionReason = "InstallSuccessful"
	ReasonUpdateSuccessful    AppConditionReason = "UpdateSuccessful"
	ReasonUninstallSuccessful AppConditionReason = "UninstallSuccessful"
	ReasonInstallError        AppConditionReason = "InstallError"
	ReasonUpdateError         AppConditionReason = "UpdateError"
	ReasonReconcileError      AppConditionReason = "ReconcileError"
	ReasonUninstallError      AppConditionReason = "UninstallError"
)

type AppStatus struct {
	Conditions []AppCondition `json:"conditions"`
}

func (s *AppStatus) ToMap() (map[string]interface{}, error) {
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
func (s *AppStatus) SetCondition(condition AppCondition) *AppStatus {
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
func (s *AppStatus) RemoveCondition(conditionType AppConditionType) *AppStatus {
	for i := range s.Conditions {
		if s.Conditions[i].Type == conditionType {
			s.Conditions = append(s.Conditions[:i], s.Conditions[i+1:]...)
			return s
		}
	}
	return s
}

// StatusFor safely returns a typed status block from a custom resource.
func StatusFor(cr *unstructured.Unstructured) *AppStatus {
	switch s := cr.Object["status"].(type) {
	case *AppStatus:
		return s
	case map[string]interface{}:
		var status *AppStatus
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(s, &status); err != nil {
			return &AppStatus{}
		}
		return status
	default:
		return &AppStatus{}
	}
}
