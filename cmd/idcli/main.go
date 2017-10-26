package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	api "k8s.io/identity/pkg/apis/workload"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"k8s.io/apiserver/pkg/util/logs"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	cmd := NewCommand(os.Stdout, os.Stderr)
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func NewCommand(out, errOut io.Writer) *cobra.Command {
	ctx := context.Background()

	path := "/data/id.sock"

	conn, err := grpc.DialContext(ctx, path,
		grpc.WithInsecure(),
		grpc.WithTimeout(100*time.Second),
		grpc.WithDialer(func(address string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("unix", address, timeout)
		}),
	)
	if err != nil {
		glog.Fatalf("fail to dial: %v", err)
	}

	c := api.NewWorkloadClient(conn)

	cmd := &cobra.Command{
		Short: "Workload API cli",
	}
	cmd.AddCommand(
		&cobra.Command{
			Use: "get_token",
			RunE: func(cmd *cobra.Command, args []string) error {
				c.GetToken(ctx, &api.GetTokenRequest{Audience: []string{"foo"}})
				return nil
			},
		},
	)
	return cmd
}
