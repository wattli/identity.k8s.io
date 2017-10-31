package identity

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:nonNamespaced
// +genclient:noVerbs
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IdentityDocument struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Audience []string
	JWT      string
}
