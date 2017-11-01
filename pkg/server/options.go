package server

import (
	"io"

	"github.com/spf13/pflag"
	genericoptions "k8s.io/apiserver/pkg/server/options"
)

type ServerOptions struct {
	RecommendedOptions *genericoptions.RecommendedOptions
	ExtraOptions       ExtraOptions

	StdOut io.Writer
	StdErr io.Writer
}

func (o *ServerOptions) AddFlags(fs *pflag.FlagSet) {
	o.RecommendedOptions.AddFlags(fs)
	o.ExtraOptions.AddFlags(fs)
}

type ExtraOptions struct {
	Issuer string
}

func (o *ExtraOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Issuer, "oidc-issuer", o.Issuer, "")
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
