package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"

	api "k8s.io/identity/pkg/apis/idmgr"

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
	d := &Driver{}
	cmd := &cobra.Command{
		Short: "Identity volume driver",
	}
	cmd.AddCommand(
		&cobra.Command{
			Use:   "init",
			Short: "Initialize the volume driver",
			RunE: func(c *cobra.Command, args []string) error {
				if len(args) != 0 {
					return fmt.Errorf("no arguments are supported")
				}
				status, err := d.Init(ctx)
				return writeResponse(out, status, err)
			},
		},
		&cobra.Command{
			Use:   "mount",
			Short: "Mount a pod-specific volume",
			RunE: func(c *cobra.Command, args []string) error {
				if len(args) != 2 {
					return fmt.Errorf("expects exactly 2 arguments")
				}
				mountDir := args[0]
				mountOptions := make(MountOptions)
				if err := json.Unmarshal([]byte(args[1]), &mountOptions); err != nil {
					return writeResponse(out, nil, err)
				}
				status, err := d.Mount(ctx, mountDir, mountOptions)
				return writeResponse(out, status, err)
			},
		},
		&cobra.Command{
			Use:   "unmount",
			Short: "Unmount a pod-specific volume",
			RunE: func(c *cobra.Command, args []string) error {
				if len(args) != 1 {
					return fmt.Errorf("expects exactly 1 argument")
				}
				mountDir := args[0]
				status, err := d.Unmount(ctx, mountDir)
				return writeResponse(out, status, err)
			},
		},
		&cobra.Command{Use: "mountdevice", Hidden: true, RunE: func(c *cobra.Command, args []string) error { return unsupported(out, c.Name()) }},
		&cobra.Command{Use: "unmountdevice", Hidden: true, RunE: func(c *cobra.Command, args []string) error { return unsupported(out, c.Name()) }},
		&cobra.Command{Use: "attach", Hidden: true, RunE: func(c *cobra.Command, args []string) error { return unsupported(out, c.Name()) }},
		&cobra.Command{Use: "detach", Hidden: true, RunE: func(c *cobra.Command, args []string) error { return unsupported(out, c.Name()) }},
		&cobra.Command{Use: "waitforattach", Hidden: true, RunE: func(c *cobra.Command, args []string) error { return unsupported(out, c.Name()) }},
		&cobra.Command{Use: "isattached", Hidden: true, RunE: func(c *cobra.Command, args []string) error { return unsupported(out, c.Name()) }},
	)
	return cmd
}

func unsupported(out io.Writer, op string) error {
	status := &Status{Status: NotSupported, Message: fmt.Sprintf("this operation %q is not supported", op)}
	if err := writeResponse(out, status, nil); err != nil {
		return err
	}
	return errUnsupported
}

func writeResponse(out io.Writer, status *Status, err error) error {
	if err != nil {
		status = &Status{Status: Failure, Message: err.Error()}
	}
	body, writeErr := json.Marshal(status)
	if err != nil {
		fmt.Fprintln(out, `{"status":"Failure"}`)
		return fmt.Errorf("unable to write response: %v", writeErr)
	}
	out.Write(body)
	fmt.Fprintln(out)
	return err
}

var errUnsupported = fmt.Errorf("operation is unsupported")

const (
	Success      = "Success"
	Failure      = "Failure"
	NotSupported = "Not supported"
)

type MountOptions map[string]string

type Status struct {
	Status       string       `json:"status"`
	Message      string       `json:"message"`
	Device       string       `json:"device"`
	VolumeName   string       `json:"volumeName"`
	Attached     bool         `json:"attached"`
	Capabilities Capabilities `json:"capabilities,omitempty"`
}

type Capabilities struct {
	Attach bool `json:"attach"`
}

type Driver struct {
	initClientOnce sync.Once
	client         api.ManagementClient
}

func (d *Driver) Init(ctx context.Context) (*Status, error) {
	return &Status{
		Status:       Success,
		Message:      "identifier volume plugin ready",
		Capabilities: Capabilities{Attach: false},
	}, nil
}

func (d *Driver) Mount(ctx context.Context, dir string, options MountOptions) (*Status, error) {
	d.initClient(ctx)
	d.client.CreateIdentityVolume(ctx, &api.CreateIdentityVolumeRequest{
		MountPath: dir,
		PodInfo: &api.PodInfo{
			Name:           options["kubernetes.io/pod.name"],
			Namespace:      options["kubernetes.io/pod.namespace"],
			Uid:            options["kubernetes.io/pod.uid"],
			ServiceAccount: options["kubernetes.io/serviceAccount.name"],
		},
	})
	return &Status{
		Status:  Success,
		Message: fmt.Sprintf("mounted %s", dir),
	}, nil
}

func (d *Driver) Unmount(ctx context.Context, dir string) (*Status, error) {
	d.initClient(ctx)
	d.client.DestroyIdentityVolume(ctx, &api.DestroyIdentityVolumeRequest{
		MountPath: dir,
	})
	return &Status{
		Status:  Success,
		Message: fmt.Sprintf("mounted %s", dir),
	}, nil
}

func (d *Driver) initClient(ctx context.Context) {
	d.initClientOnce.Do(func() {
		path := "/tmp/idmgr.sock"

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

		d.client = api.NewManagementClient(conn)
	})
}
