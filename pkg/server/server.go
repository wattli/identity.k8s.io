package server

import (
	"fmt"
	"io"
	"net"

	"github.com/spf13/cobra"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
)

type ServerOptions struct {
	RecommendedOptions *genericoptions.RecommendedOptions

	StdOut io.Writer
	StdErr io.Writer
}

func NewServerOptions(out, errOut io.Writer) *ServerOptions {
	o := &ServerOptions{
		RecommendedOptions: &genericoptions.RecommendedOptions{
			SecureServing:  genericoptions.NewSecureServingOptions(),
			Authentication: genericoptions.NewDelegatingAuthenticationOptions(),
			Authorization:  genericoptions.NewDelegatingAuthorizationOptions(),
			Audit:          genericoptions.NewAuditOptions(),
			Features:       genericoptions.NewFeatureOptions(),
			CoreAPI:        genericoptions.NewCoreAPIOptions(),
		},

		StdOut: out,
		StdErr: errOut,
	}

	return o
}

// NewCommandStartServer starts the API server.
func NewCommandStartServer(out, errOut io.Writer, stopCh <-chan struct{}) *cobra.Command {
	o := NewServerOptions(out, errOut)

	cmd := &cobra.Command{
		Short: "Launch API server",
		Long:  "Launch API server",
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(args); err != nil {
				return err
			}
			if err := o.RunServer(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	flags := cmd.Flags()
	o.RecommendedOptions.AddFlags(flags)

	return cmd
}

func (o ServerOptions) Validate(args []string) error {
	errors := []error{}
	errors = append(errors, o.RecommendedOptions.Validate()...)
	return utilerrors.NewAggregate(errors)
}

func (o *ServerOptions) Complete() error {
	return nil
}

func (o ServerOptions) Config() (*Config, error) {
	// TODO have a "real" external address
	if err := o.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	serverConfig := genericapiserver.NewRecommendedConfig(Codecs)
	if err := o.RecommendedOptions.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	config := &Config{
		GenericConfig: serverConfig,
		ExtraConfig:   ExtraConfig{},
	}
	return config, nil
}

func (o ServerOptions) RunServer(stopCh <-chan struct{}) error {
	config, err := o.Config()
	if err != nil {
		return err
	}

	server, err := config.Complete().New()
	if err != nil {
		return err
	}

	return server.GenericAPIServer.PrepareRun().Run(stopCh)
}
