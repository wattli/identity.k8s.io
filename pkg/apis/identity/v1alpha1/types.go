package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type IdentityDocument struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Audience []string `json:"audience,omitempty"`
	JWT      string   `json:"jwt,omitempty"`
}
