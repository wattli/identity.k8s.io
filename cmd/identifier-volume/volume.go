package main

import (
	"flag"
	"fmt"
	"os"

	"k8s.io/apiserver/pkg/util/logs"

	"github.com/openshift/identifier/pkg/volume"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	cmd := volume.NewCommand(os.Stdout, os.Stderr)
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
