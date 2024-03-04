package shim

import (
	"github.com/grafana/terraform-provider-grafana/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewProvider(version string) *schema.Provider {
	prov := provider.Provider(version)()
	return prov
}
