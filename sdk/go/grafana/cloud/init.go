// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "grafana:cloud/accessPolicy:AccessPolicy":
		r = &AccessPolicy{}
	case "grafana:cloud/accessPolicyToken:AccessPolicyToken":
		r = &AccessPolicyToken{}
	case "grafana:cloud/orgMember:OrgMember":
		r = &OrgMember{}
	case "grafana:cloud/pluginInstallation:PluginInstallation":
		r = &PluginInstallation{}
	case "grafana:cloud/privateDataSourceConnectNetwork:PrivateDataSourceConnectNetwork":
		r = &PrivateDataSourceConnectNetwork{}
	case "grafana:cloud/privateDataSourceConnectNetworkToken:PrivateDataSourceConnectNetworkToken":
		r = &PrivateDataSourceConnectNetworkToken{}
	case "grafana:cloud/providerAwsAccount:ProviderAwsAccount":
		r = &ProviderAwsAccount{}
	case "grafana:cloud/providerAwsCloudwatchScrapeJob:ProviderAwsCloudwatchScrapeJob":
		r = &ProviderAwsCloudwatchScrapeJob{}
	case "grafana:cloud/providerAzureCredential:ProviderAzureCredential":
		r = &ProviderAzureCredential{}
	case "grafana:cloud/stack:Stack":
		r = &Stack{}
	case "grafana:cloud/stackServiceAccount:StackServiceAccount":
		r = &StackServiceAccount{}
	case "grafana:cloud/stackServiceAccountToken:StackServiceAccountToken":
		r = &StackServiceAccountToken{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"grafana",
		"cloud/accessPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"cloud/accessPolicyToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"cloud/orgMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"cloud/pluginInstallation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"cloud/privateDataSourceConnectNetwork",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"cloud/privateDataSourceConnectNetworkToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"cloud/providerAwsAccount",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"cloud/providerAwsCloudwatchScrapeJob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"cloud/providerAzureCredential",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"cloud/stack",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"cloud/stackServiceAccount",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"cloud/stackServiceAccountToken",
		&module{version},
	)
}
