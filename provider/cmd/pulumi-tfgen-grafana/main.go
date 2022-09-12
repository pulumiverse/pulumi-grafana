package main

import (
	grafana "github.com/lbrlabs/pulumi-grafana/provider"
	"github.com/lbrlabs/pulumi-grafana/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen"
)

func main() {
	// Modify the path to point to the new provider
	tfgen.Main("grafana", version.Version, grafana.Provider())
}
