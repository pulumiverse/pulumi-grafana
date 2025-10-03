//go:generate go run ./generate.go

package main

import (
	"context"

	_ "embed"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"

	grafana "github.com/pulumiverse/pulumi-grafana/provider/v2"
)

//go:embed schema-embed.json
var pulumiSchema []byte

func main() {
	// Modify the path to point to the new provider
	tfbridge.MainWithMuxer(context.Background(), "grafana", grafana.Provider(), pulumiSchema)
}
