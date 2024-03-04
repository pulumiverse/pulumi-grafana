//go:generate go run ./generate.go

package main

import (
	_ "embed"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	grafana "github.com/pulumiverse/pulumi-grafana/provider"
	"github.com/pulumiverse/pulumi-grafana/provider/pkg/version"
)

//go:embed schema-embed.json
var pulumiSchema []byte

func main() {
	// Modify the path to point to the new provider
	tfbridge.Main("grafana", version.Version, grafana.Provider(), pulumiSchema)
}
