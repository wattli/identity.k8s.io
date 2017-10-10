package main

import (
	"context"
	"flag"
	"os"

	"github.com/golang/glog"

	genericoptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/identity/pkg/server"
)

func main() {
	flag.Set("alsologtostderr", "true")
	flag.Parse()

	ctx := context.Background()
	o := &server.ServerOptions{
		RecommendedOptions: &genericoptions.RecommendedOptions{
			SecureServing: genericoptions.NewSecureServingOptions(),
			Features:      genericoptions.NewFeatureOptions(),
		},

		StdOut: os.Stdout,
		StdErr: os.Stderr,
	}
	o.RecommendedOptions.SecureServing.BindPort = 9443
	glog.Fatalf("err: %v", o.RunServer(ctx.Done()))
}
