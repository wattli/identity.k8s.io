package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	api "k8s.io/identity/pkg/apis/idmgr"
	"k8s.io/identity/pkg/management"
	"k8s.io/identity/pkg/uds"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func main() {
	flag.Set("alsologtostderr", "true")
	cmd := &cobra.Command{
		Short: "Identity manager",
		RunE: func(c *cobra.Command, args []string) error {

			if err := setupPluginDir(); err != nil {
				return fmt.Errorf("failed to setup volume plugin dir: %v", err)
			}

			s, err := uds.New(
				"/tmp/idmgr.sock",
				func(s *grpc.Server) {
					api.RegisterManagementServer(s, management.NewServer())
				},
				uds.LoggingInterceptor,
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

func setupPluginDir() error {
	if err := os.MkdirAll("/volumeplugin/k8s.io~identity", 0777); err != nil {
		return err
	}
	fsrc, err := os.Open("/usr/local/bin/idmgr-driver")
	if err != nil {
		return err
	}
	defer fsrc.Close()

	stat, err := fsrc.Stat()
	if err != nil {
		return err
	}

	ftrgt, err := os.OpenFile("/volumeplugin/k8s.io~identity/identity", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, stat.Mode())
	if err != nil {
		return err
	}
	defer ftrgt.Close()

	_, err = io.Copy(ftrgt, fsrc)
	return err
}
