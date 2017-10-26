package workload

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/golang/glog"

	mapi "k8s.io/identity/pkg/apis/idmgr"
	api "k8s.io/identity/pkg/apis/workload"
)

type Server struct {
}

func (w *Server) GetToken(ctx context.Context, req *api.GetTokenRequest) (*api.GetTokenResponse, error) {
	info, ok := GetInfo(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Internal, "unable to retrieve authentication from channel")
	}
	glog.Infof("info: %#v", info)
	return &api.GetTokenResponse{}, nil
}

func (w *Server) ValidateToken(ctx context.Context, req *api.ValidateTokenRequest) (*api.ValidateTokenResponse, error) {
	return &api.ValidateTokenResponse{}, nil
}

func GetInfo(ctx context.Context) (*mapi.PodInfo, bool) {
	info, ok := ctx.Value("pod-info").(*mapi.PodInfo)
	return info, ok
}
