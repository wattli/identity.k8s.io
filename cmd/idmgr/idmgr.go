package main

import (
	"flag"
	"fmt"
	"os"

	"k8s.io/client-go/tools/clientcmd"
	api "k8s.io/identity/pkg/apis/idmgr"
	idclient "k8s.io/identity/pkg/client/clientset/typed/identity/v1alpha1"
	"k8s.io/identity/pkg/management"
	"k8s.io/identity/pkg/uds"
	"k8s.io/identity/pkg/util"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	restclient "k8s.io/client-go/rest"
)

func main() {
	flag.Set("alsologtostderr", "true")
	cmd := &cobra.Command{
		Short: "Identity manager",
		RunE: func(c *cobra.Command, args []string) error {
			return run()
		},
	}
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if err := setupPluginDir(); err != nil {
		return fmt.Errorf("failed to setup volume plugin dir: %v", err)
	}

	cc, err := client()
	if err != nil {
		return fmt.Errorf("failed to create kubernetes client: %v", err)
	}
	c, err := idclient.NewForConfig(cc)
	if err != nil {
		return err
	}

	s, err := uds.New(
		"/tmp/idmgr.sock",
		func(s *grpc.Server) {
			api.RegisterManagementServer(s, management.NewServer(c.IdentityDocuments()))
		},
		uds.LoggingInterceptor,
	)
	if err != nil {
		return err
	}
	glog.Fatalf("Server exited unexpectedly: %v", s.Serve())
	return nil
}

func setupPluginDir() error {
	if err := os.MkdirAll("/volumeplugin/k8s.io~identity", 0777); err != nil {
		return err
	}
	return util.CopyFile("/usr/local/bin/idmgr-driver", "/volumeplugin/k8s.io~identity/identity")
}

func client() (*restclient.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: "/var/lib/kubelet/kubeconfig"},
		&clientcmd.ConfigOverrides{}).ClientConfig()
}
