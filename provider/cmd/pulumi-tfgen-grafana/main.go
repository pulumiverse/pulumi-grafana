package main

import (
	"github.com/pulumi/pulumi-terraform-bridge/pf/tfgen"

	grafana "github.com/pulumiverse/pulumi-grafana/provider"
)

func main() {
	// Modify the path to point to the new provider
	tfgen.MainWithMuxer("grafana", grafana.Provider())
}
