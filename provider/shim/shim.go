package shim

import (
	"github.com/grafana/terraform-provider-grafana/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pulumiverse/pulumi-grafana/provider/pkg/version"
)

func NewProvider() *schema.Provider {
	prov := provider.Provider(version.Version)()
	return prov
}
