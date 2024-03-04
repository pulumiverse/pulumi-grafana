package main

import (
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen"
	grafana "github.com/pulumiverse/pulumi-grafana/provider"
	"github.com/pulumiverse/pulumi-grafana/provider/pkg/version"
)

func main() {
	// Modify the path to point to the new provider
	tfgen.Main("grafana", version.Version, grafana.Provider())
}
