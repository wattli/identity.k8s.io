package workload

import (
	"context"
	"fmt"

	identityapi "k8s.io/identity/pkg/apis/identity/v1alpha1"
	mapi "k8s.io/identity/pkg/apis/idmgr"
	api "k8s.io/identity/pkg/apis/workload"
	idclient "k8s.io/identity/pkg/client/clientset/typed/identity/v1alpha1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Server struct {
	IdCli idclient.IdentityDocumentInterface
}

func (w *Server) GetToken(ctx context.Context, req *api.GetTokenRequest) (*api.GetTokenResponse, error) {
	info, ok := GetInfo(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Internal, "unable to retrieve authentication from channel")
	}
	resp, err := w.IdCli.Create(
		&identityapi.IdentityDocument{
			Audience: req.Audience,
		},
		fmt.Sprintf("system:serviceaccount:%s:%s", info.Namespace, info.Name),
	)
	if err != nil {
		return nil, err
	}
	return &api.GetTokenResponse{
		Token: []byte(resp.JWT),
	}, nil
}

func (w *Server) ValidateToken(ctx context.Context, req *api.ValidateTokenRequest) (*api.ValidateTokenResponse, error) {
	return &api.ValidateTokenResponse{}, nil
}

func GetInfo(ctx context.Context) (*mapi.PodInfo, bool) {
	info, ok := ctx.Value("pod-info").(*mapi.PodInfo)
	return info, ok
}
