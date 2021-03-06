package management

import (
	"context"
	"os"
	"path/filepath"

	mapi "k8s.io/identity/pkg/apis/idmgr"
	api "k8s.io/identity/pkg/apis/workload"
	idclient "k8s.io/identity/pkg/client/clientset/typed/identity/v1alpha1"
	"k8s.io/identity/pkg/uds"
	"k8s.io/identity/pkg/util"
	"k8s.io/identity/pkg/workload"

	"google.golang.org/grpc"
)

type Manager struct {
	Dir     string
	PodInfo *mapi.PodInfo
	IdCli   idclient.IdentityDocumentInterface
	stop    func() error
}

func (m *Manager) Start() error {
	os.MkdirAll(m.Dir, 0744)
	s, err := uds.New(filepath.Join(m.Dir, "id.sock"), func(s *grpc.Server) {
		api.RegisterWorkloadServer(s, &workload.Server{
			IdCli: m.IdCli,
		})
	},
		uds.LoggingInterceptor,
		podInfoInterceptor(m.PodInfo),
	)
	if err != nil {
		return err
	}

	if err := util.CopyFile("/usr/local/bin/idcli", filepath.Join(m.Dir, "idcli")); err != nil {
		return err
	}

	m.stop = func() error {
		s.Stop()
		os.RemoveAll(m.Dir)
		return nil
	}
	go s.Serve()
	return nil
}

func (m *Manager) Stop() error {
	return m.stop()
}

func podInfoInterceptor(pod *mapi.PodInfo) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, "pod-info", pod), req)
	}
}
