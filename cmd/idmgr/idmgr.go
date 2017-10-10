package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sync"

	api "k8s.io/identity/pkg/apis/idmgr"
	"k8s.io/identity/pkg/volumemgr"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func main() {
	flag.Set("alsologtostderr", "true")
	cmd := &cobra.Command{
		Short: "Identity manager",
		RunE: func(c *cobra.Command, args []string) error {
			s, err := newUnixServer(
				"/tmp/idmgr.sock",
				func(s *grpc.Server) {
					api.RegisterManagementServer(s, &managementServer{})
				},
				loggingInterceptor,
			)
			if err != nil {
				return err
			}
			glog.Fatalf("Server exited unexpectedly: %v", s.Serve())
			return nil
		},
	}
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

type managementServer struct {
	sync.Mutex
	mgrs map[string]volumemgr.Manager
}

func (ms *managementServer) CreateIdentityVolume(ctx context.Context, in *api.CreateIdentityVolumeRequest) (*api.CreateIdentityVolumeResponse, error) {
	return &api.CreateIdentityVolumeResponse{}, nil
}

func (s *managementServer) DestroyIdentityVolume(ctx context.Context, in *api.DestroyIdentityVolumeRequest) (*api.DestroyIdentityVolumeResponse, error) {
	return &api.DestroyIdentityVolumeResponse{}, nil
}
