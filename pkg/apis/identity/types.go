package identity

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type IdentityDocument struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Audience []string
	JWT      string
}

type JSONWebKey struct {
	Use string `json:"use,omitempty"`
	Kty string `json:"kty,omitempty"`
	Kid string `json:"kid,omitempty"`
	Crv string `json:"crv,omitempty"`
	Alg string `json:"alg,omitempty"`
	K   []byte `json:"k,omitempty"`
	X   []byte `json:"x,omitempty"`
	Y   []byte `json:"y,omitempty"`
	N   []byte `json:"n,omitempty"`
	E   []byte `json:"e,omitempty"`
	// we only care about a subset of fields
}
