package v1alpha1

import (
	"k8s.io/client-go/transport"
	identityapi "k8s.io/identity/pkg/apis/identity/v1alpha1"
)

type IdentityDocumentExpansion interface {
	Create(sar *identityapi.IdentityDocument, user string) (result *identityapi.IdentityDocument, err error)
}

func (c *identityDocuments) Create(sar *identityapi.IdentityDocument, user string) (result *identityapi.IdentityDocument, err error) {
	result = &identityapi.IdentityDocument{}
	err = c.client.Post().
		SetHeader(transport.ImpersonateUserHeader, user).
		Resource("identitydocuments").
		Body(sar).
		Do().
		Into(result)
	return
}
