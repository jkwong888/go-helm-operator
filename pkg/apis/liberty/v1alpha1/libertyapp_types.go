package v1alpha1

import (
	"reflect"

	"github.com/fatih/structtag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LibertyApp defines the desired state of LibertyApp
// +k8s:openapi-gen=true
type LibertyApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   LibertyAppSpec   `json:"spec"`
	Status LibertyAppStatus `json:"status,omitempty"`
}

// LibertyAppStatus defines the current state of LibertyApp
type LibertyAppStatus struct {
	AppStatus
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

// StructToJSONTagMap generates a map from an struct using jsontags as the keys
func StructToJSONTagMap(intf interface{}) map[string]interface{} {
	s := reflect.TypeOf(intf)

	out := make(map[string]interface{})
	numMetadataFields := s.NumField()

	for i := 0; i < numMetadataFields; i++ {
		field := s.Field(i)
		fieldName := field.Name
		tag := field.Tag

		// get the jsonTag associated with the field
		tags, err := structtag.Parse(string(tag))
		if err != nil {
			continue
		}

		jsonTag, err := tags.Get("json")
		if err != nil {
			continue
		}

		v := reflect.ValueOf(intf).FieldByName(fieldName).Interface()
		switch v.(type) {
		case int:
			out[jsonTag.Name] = v
		case string:
			out[jsonTag.Name] = v
		case bool:
			out[jsonTag.Name] = v
		case map[string]string:
			out[jsonTag.Name] = v
		case []map[string]string:
			out[jsonTag.Name] = v
		case []interface{}:
			var arr []interface{}

			for _, elem := range v.([]interface{}) {
				arr = append(arr, StructToJSONTagMap(elem))
			}
			out[jsonTag.Name] = arr
		default:
			if reflect.TypeOf(v).Kind() == reflect.Ptr {
				// if null value, just omit it (TODO according to the jsontag options)
				if reflect.ValueOf(v).IsNil() {
					continue
				}
			}
			if reflect.TypeOf(v).Kind() != reflect.Struct {
				// all other unknown types
				out[jsonTag.Name] = v
				continue
			}

			out[jsonTag.Name] = StructToJSONTagMap(v)
		}
	}

	return out
}

// ToUnstructured return an Unstructured with contents of libertyapp mapped by the jsonTag name
func (app *LibertyApp) ToUnstructured() *unstructured.Unstructured {
	o := &unstructured.Unstructured{}

	o.SetGroupVersionKind(app.TypeMeta.GroupVersionKind())

	o.SetName(app.ObjectMeta.GetName())
	o.SetNamespace(app.ObjectMeta.GetNamespace())
	o.SetAnnotations(app.ObjectMeta.GetAnnotations())
	o.SetFinalizers(app.ObjectMeta.GetFinalizers())
	o.SetResourceVersion(app.ObjectMeta.GetResourceVersion())
	o.SetCreationTimestamp(app.ObjectMeta.GetCreationTimestamp())
	o.SetGeneration(app.ObjectMeta.GetGeneration())
	o.SetDeletionTimestamp(app.ObjectMeta.GetDeletionTimestamp())
	o.SetDeletionGracePeriodSeconds(app.ObjectMeta.GetDeletionGracePeriodSeconds())
	o.SetClusterName(app.ObjectMeta.GetClusterName())
	o.SetLabels(app.ObjectMeta.GetLabels())
	o.SetUID(app.ObjectMeta.GetUID())

	// get the map so far
	unstructuredMap := o.UnstructuredContent()

	// add the spec and status using the json tags as keys
	unstructuredMap["spec"] = StructToJSONTagMap(app.Spec)
	unstructuredMap["status"] = StructToJSONTagMap(app.Status)

	// set it back into the unstructured
	o.SetUnstructuredContent(unstructuredMap)

	return o
}
