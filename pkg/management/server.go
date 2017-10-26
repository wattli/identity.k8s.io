package management

import (
	"context"
	"fmt"
	"sync"

	api "k8s.io/identity/pkg/apis/idmgr"
)

func NewServer() *Server {
	return &Server{
		state: make(map[string]*Manager),
	}
}

type Server struct {
	sync.Mutex
	state map[string]*Manager
}

func (ms *Server) CreateIdentityVolume(ctx context.Context, in *api.CreateIdentityVolumeRequest) (*api.CreateIdentityVolumeResponse, error) {
	vmgr := &Manager{
		Dir:     in.GetMountPath(),
		PodInfo: in.GetPodInfo(),
	}
	ms.Lock()
	defer ms.Unlock()
	// todo make sure it doesn't exist already
	ms.state[in.MountPath] = vmgr
	vmgr.Start()
	return &api.CreateIdentityVolumeResponse{}, vmgr.Start()
}

func (ms *Server) DestroyIdentityVolume(ctx context.Context, in *api.DestroyIdentityVolumeRequest) (*api.DestroyIdentityVolumeResponse, error) {
	ms.Lock()
	defer ms.Unlock()
	vmgr, ok := ms.state[in.MountPath]
	if !ok {
		return nil, fmt.Errorf("couldn't find manager for %q", in.MountPath)
	}
	if err := vmgr.Stop(); err != nil {
		return nil, fmt.Errorf("couldn't stop workload api %q: %v", in.MountPath, err)
	}
	delete(ms.state, in.MountPath)
	return &api.DestroyIdentityVolumeResponse{}, nil
}
