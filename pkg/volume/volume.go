package volume

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

func NewCommand(out, errOut io.Writer) *cobra.Command {
	d := &Driver{}
	cmd := &cobra.Command{
		Short: "Identifier volume driver",
		Long:  "Inject identifiers for pods",
	}
	cmd.AddCommand(
		&cobra.Command{
			Use:   "init",
			Short: "Initialize the volume driver",
			RunE: func(c *cobra.Command, args []string) error {
				if len(args) != 0 {
					return fmt.Errorf("no arguments are supported")
				}
				status, err := d.Init()
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
				status, err := d.Mount(mountDir, mountOptions)
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
				status, err := d.Unmount(mountDir)
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
	switch {
	case err != nil:
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

type Driver struct{}

func (d *Driver) Init() (*Status, error) {
	return &Status{
		Status:       Success,
		Message:      "identifier volume plugin ready",
		Capabilities: Capabilities{Attach: false},
	}, nil
}

func (d *Driver) Mount(dir string, options MountOptions) (*Status, error) {
	return &Status{
		Status:  Success,
		Message: fmt.Sprintf("mounted %s", dir),
	}, nil
}

func (d *Driver) Unmount(dir string) (*Status, error) {
	return &Status{
		Status:  Success,
		Message: fmt.Sprintf("mounted %s", dir),
	}, nil
}
