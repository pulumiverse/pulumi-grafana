package shim

import (
	"github.com/grafana/terraform-provider-grafana/internal/provider"
	"github.com/lbrlabs/pulumi-grafana/provider/pkg/version"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewProvider() *schema.Provider {
	prov := provider.Provider(version.Version)()
	return prov
}