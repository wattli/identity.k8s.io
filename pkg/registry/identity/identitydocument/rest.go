package identitydocument

import (
	"k8s.io/apimachinery/pkg/runtime"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	identityapi "k8s.io/identity/pkg/apis/identity"
	"k8s.io/identity/pkg/jwt"
)

type REST struct {
	signer *jwt.Signer
}

func NewREST(s *jwt.Signer) *REST {
	return &REST{signer: s}
}

func (r *REST) New() runtime.Object {
	return &identityapi.IdentityDocument{}
}

func (r *REST) Create(ctx genericapirequest.Context, obj runtime.Object, includeUninitialized bool) (runtime.Object, error) {
	doc := obj.(*identityapi.IdentityDocument)
	var rawJwt string

	if user, ok := genericapirequest.UserFrom(ctx); ok {
		rawJwt = r.signer.Sign(
			jwt.PublicClaims{
				Subject:  user.GetName(),
				Audience: doc.Audience,
			},
			jwt.PrivateClaims{
				Kubernetes: jwt.KubernetesClaim{
					Groups: user.GetGroups(),
				},
			},
		)
	} else {
		rawJwt = r.signer.Sign(
			jwt.PublicClaims{
				Subject:  "system:unauthenticated",
				Audience: doc.Audience,
			},
			jwt.PrivateClaims{
				Kubernetes: jwt.KubernetesClaim{
					Groups: []string{"system:unauthenticated"},
				},
			},
		)
	}

	doc.JWT = rawJwt
	return doc, nil
}
