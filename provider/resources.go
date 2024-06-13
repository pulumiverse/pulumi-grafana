package grafana

import (
	"context"
	"fmt"
	"path/filepath"
	"unicode"

	// The linter requires unnamed imports to have a doc comment
	_ "embed"

	grafana "github.com/grafana/terraform-provider-grafana/v2/pkg/provider"
	pfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tks "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumiverse/pulumi-grafana/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	grafanaPkg = "grafana"
	// modules:
	grafanaMod = "index" // the toplevel module
	// further modules follow the grouping of the upstream TF provider, except for OSS
	// https://registry.terraform.io/providers/grafana/grafana/latest/docs
	alertingMod            = "alerting"
	cloudMod               = "cloud"
	enterpriseMod          = "enterprise"
	mlMod                  = "machineLearning"
	oncallMod              = "onCall"
	sloMod                 = "slo"
	syntheticMonitoringMod = "syntheticMonitoring"
)

// grafanaMember manufactures a type token for the grafana package and the given module and type.
func grafanaMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(grafanaPkg + ":" + mod + ":" + mem)
}

// grafanaType manufactures a type token for the grafana package and the given module and type.
func grafanaType(mod string, typ string) tokens.Type {
	return tokens.Type(grafanaMember(mod, typ))
}

// grafanaDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the grafana package and names the file by simply lower casing the data
// source's first character.
func grafanaDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return grafanaMember(mod+"/"+fn, res)
}

// grafanaResource manufactures a standard resource token given a module and resource name.
// It automatically uses the grafana package and names the file by simply lower casing the resource's
// first character.
func grafanaResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return grafanaType(mod+"/"+fn, res)
}

func grafanaResourceAlias(mod string, res string) *string {
	alias := grafanaResource(mod, res).String()
	return &alias
}

func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

func boolRef(b bool) *bool {
	return &b
}

//go:embed cmd/pulumi-resource-grafana/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := pfbridge.MuxShimWithPF(context.Background(),
		shimv2.NewProvider(grafana.Provider(version.Version)),
		grafana.FrameworkProvider(version.Version))

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                       p,
		Name:                    "grafana",
		DisplayName:             "Grafana",
		Publisher:               "pulumiverse",
		LogoURL:                 "https://raw.githubusercontent.com/pulumiverse/pulumi-grafana/main/assets/grafana.png", // nolint[:lll]
		PluginDownloadURL:       "github://api.github.com/pulumiverse",
		Version:                 version.Version,
		Description:             "A Pulumi package for creating and managing grafana.",
		Keywords:                []string{"pulumi", "grafana", "pulumiverse"},
		License:                 "Apache-2.0",
		Homepage:                "https://grafana.com",
		Repository:              "https://github.com/pulumiverse/pulumi-grafana",
		GitHubOrg:               "grafana",
		TFProviderModuleVersion: "v2",
		Config: map[string]*tfbridge.SchemaInfo{
			"url": {
				Type: "string",
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_URL"},
				},
			},
			"auth": {
				Type:   "string",
				Secret: boolRef(true),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_AUTH"},
				},
			},
			"http_headers": {
				SuppressEmptyMapElements: boolRef(true),
				Omit:                     true,
			},
			"retries": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_RETRIES"},
				},
			},
			"org_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_ORG_ID"},
				},
			},
			"tls_key": {
				Secret: boolRef(true),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_TLS_KEY"},
				},
			},
			"tls_cert": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_TLS_CERT"},
				},
			},
			"ca_cert": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_CA_CERT"},
				},
			},
			"insecure_skip_verify": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_INSECURE_SKIP_VERIFY"},
				},
			},
			"cloud_api_key": {
				Secret: boolRef(true),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_CLOUD_API_KEY"},
				},
			},
			"cloud_api_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_CLOUD_API_URL"},
				},
			},
			"sm_access_token": {
				Secret: boolRef(true),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_SM_ACCESS_TOKEN"},
				},
			},
			"sm_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_SM_URL"},
				},
			},
			"store_dashboard_sha256": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_STORE_DASHBOARD_SHA256"},
				},
			},
			"oncall_access_token": {
				Secret: boolRef(true),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_ONCALL_ACCESS_TOKEN"},
				},
			},
			"oncall_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_ONCALL_URL"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Alerting
			"grafana_contact_point": {
				Tok: grafanaResource(alertingMod, "ContactPoint"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "ContactPoint"),
					},
				},
			},
			"grafana_message_template": {
				Tok: grafanaResource(alertingMod, "MessageTemplate"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "MessageTemplate"),
					},
				},
			},
			"grafana_mute_timing": {
				Tok: grafanaResource(alertingMod, "MuteTiming"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "MuteTiming"),
					},
				},
			},
			"grafana_notification_policy": {
				Tok: grafanaResource(alertingMod, "NotificationPolicy"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "NotificationPolicy"),
					},
				},
			},
			"grafana_rule_group": {
				Tok: grafanaResource(alertingMod, "RuleGroup"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "RuleGroup"),
					},
				},
			},

			// Cloud
			"grafana_cloud_access_policy": {
				Tok: grafanaResource(cloudMod, "AccessPolicy"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "CloudAccessPolicy"),
					},
				},
			},
			"grafana_cloud_access_policy_token": {
				Tok: grafanaResource(cloudMod, "AccessPolicyToken"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "CloudAccessPolicyToken"),
					},
				},
			},
			"grafana_cloud_api_key": {
				Tok: grafanaResource(cloudMod, "ApiKey"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "CloudApiKey"),
					},
				},
			},
			// "grafana_cloud_org_member": {
			// 	Tok: grafanaResource(cloudMod, "OrgMember"),
			// 	Aliases: []tfbridge.AliasInfo{
			// 		{
			// 			Type: grafanaResourceAlias(grafanaMod, "CloudOrgMember"),
			// 		},
			// 	},
			// },
			"grafana_cloud_plugin_installation": {
				Tok: grafanaResource(cloudMod, "PluginInstallation"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "CloudPluginInstallation"),
					},
				},
			},
			"grafana_cloud_stack": {
				Tok: grafanaResource(cloudMod, "Stack"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "CloudStack"),
					},
				},
			},
			"grafana_cloud_stack_api_key": {
				Tok: grafanaResource(cloudMod, "StackApiKey"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "CloudStackApiKey"),
					},
				},
			},
			"grafana_cloud_stack_service_account": {
				Tok: grafanaResource(cloudMod, "StackServiceAccount"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "CloudStackServiceAccount"),
					},
				},
			},
			"grafana_cloud_stack_service_account_token": {
				Tok: grafanaResource(cloudMod, "StackServiceAccountToken"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "CloudStackServiceAccountToken"),
					},
				},
			},

			// Enterprise
			"grafana_data_source_permission": {
				Tok: grafanaResource(enterpriseMod, "DataSourcePermission"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "DataSourcePermission"),
					},
				},
			},
			// "grafana_data_source_permission_item": {
			// 	Tok: grafanaResource(enterpriseMod, "DataSourcePermissionItem"),
			// 	Aliases: []tfbridge.AliasInfo{
			// 		{
			// 			Type: grafanaResourceAlias(grafanaMod, "DataSourcePermissionItem"),
			// 		},
			// 	},
			// },
			"grafana_report": {
				Tok: grafanaResource(enterpriseMod, "Report"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "Report"),
					},
				},
			},
			"grafana_role": {
				Tok: grafanaResource(enterpriseMod, "Role"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "Role"),
					},
				},
			},
			"grafana_role_assignment": {
				Tok: grafanaResource(enterpriseMod, "RoleAssignment"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "RoleAssignment"),
					},
				},
			},
			// "grafana_role_assignment_item": {
			// 	Tok: grafanaResource(enterpriseMod, "RoleAssignmentItem"),
			// 	Aliases: []tfbridge.AliasInfo{
			// 		{
			// 			Type: grafanaResourceAlias(grafanaMod, "RoleAssignmentItem"),
			// 		},
			// 	},
			// },
			"grafana_team_external_group": {
				Tok: grafanaResource(enterpriseMod, "TeamExternalGroup"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "TeamExternalGroup"),
					},
				},
			},

			// OSS is not remapped to a submodule. Mapping handled by auto-mapping below.

			// Machine Learning
			"grafana_machine_learning_holiday": {
				Tok: grafanaResource(mlMod, "Holiday"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "MachineLearningHoliday"),
					},
				},
			},
			"grafana_machine_learning_job": {
				Tok: grafanaResource(mlMod, "Job"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "MachineLearningJob"),
					},
				},
			},
			"grafana_machine_learning_outlier_detector": {
				Tok: grafanaResource(mlMod, "OutlierDetector"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "MachineLearningOutlierDetector"),
					},
				},
			},

			// OnCall
			"grafana_oncall_escalation": {
				Tok: grafanaResource(oncallMod, "Escalation"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "OncallEscalation"),
					},
				},
			},
			"grafana_oncall_escalation_chain": {
				Tok: grafanaResource(oncallMod, "EscalationChain"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "OncallEscalationChain"),
					},
				},
			},
			"grafana_oncall_integration": {
				Tok: grafanaResource(oncallMod, "Integration"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "OncallIntegration"),
					},
				},
			},
			"grafana_oncall_on_call_shift": {
				Tok: grafanaResource(oncallMod, "OnCallShift"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "OncallOnCallShift"),
					},
				},
			},
			"grafana_oncall_outgoing_webhook": {
				Tok: grafanaResource(oncallMod, "OutgoingWebhook"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "OncallOutgoingWebhook"),
					},
				},
			},
			"grafana_oncall_route": {
				Tok: grafanaResource(oncallMod, "Route"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "OncallRoute"),
					},
				},
			},
			"grafana_oncall_schedule": {
				Tok: grafanaResource(oncallMod, "Schedule"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "OncallSchedule"),
					},
				},
			},

			// SLO
			"grafana_slo": {
				Tok: grafanaResource(sloMod, "SLO"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "SLO"),
					},
				},
			},

			// Synthetic Monitoring
			"grafana_synthetic_monitoring_check": {
				Tok: grafanaResource(syntheticMonitoringMod, "Check"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "SyntheticMonitoringCheck"),
					},
				},
			},
			"grafana_synthetic_monitoring_installation": {
				Tok: grafanaResource(syntheticMonitoringMod, "Installation"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "SyntheticMonitoringInstallation"),
					},
				},
			},
			"grafana_synthetic_monitoring_probe": {
				Tok: grafanaResource(syntheticMonitoringMod, "Probe"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "SyntheticMonitoringProbe"),
					},
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Alerting

			// Cloud
			"grafana_cloud_ips": {
				Tok: grafanaDataSource(cloudMod, "getIps"),
			},
			"grafana_cloud_organization": {
				Tok: grafanaDataSource(cloudMod, "getOrganization"),
			},
			"grafana_cloud_stack": {
				Tok: grafanaDataSource(cloudMod, "getStack"),
			},

			// Enterprise
			"grafana_role": {
				Tok: grafanaDataSource(enterpriseMod, "getRole"),
			},

			// OSS is not remapped to a submodule. Mapping handled by auto-mapping below.

			// Machine Learning

			// OnCall
			"grafana_oncall_action": {
				Tok: grafanaDataSource(oncallMod, "getAction"),
			},
			"grafana_oncall_escalation_chain": {
				Tok: grafanaDataSource(oncallMod, "getEscalationChain"),
			},
			"grafana_oncall_outgoing_webhook": {
				Tok: grafanaDataSource(oncallMod, "getOutgoingWebhook"),
			},
			"grafana_oncall_schedule": {
				Tok: grafanaDataSource(oncallMod, "getSchedule"),
			},
			"grafana_oncall_slack_channel": {
				Tok: grafanaDataSource(oncallMod, "getSlackChannel"),
			},
			"grafana_oncall_team": {
				Tok: grafanaDataSource(oncallMod, "getTeam"),
			},
			"grafana_oncall_user": {
				Tok: grafanaDataSource(oncallMod, "getUser"),
			},
			"grafana_oncall_user_group": {
				Tok: grafanaDataSource(oncallMod, "getUserGroup"),
			},

			// SLO
			"grafana_slos": {
				Tok: grafanaDataSource(sloMod, "getSlos"),
			},

			// Synthetic Monitoring
			"grafana_synthetic_monitoring_probe": {
				Tok: grafanaDataSource(syntheticMonitoringMod, "getProbe"),
			},
			"grafana_synthetic_monitoring_probes": {
				Tok: grafanaDataSource(syntheticMonitoringMod, "getProbes"),
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@pulumiverse/grafana",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			PackageName: "pulumiverse_grafana",
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumiverse/pulumi-%[1]s/sdk/", grafanaPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				grafanaPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Pulumiverse",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
	}

	prov.MustComputeTokens(tks.SingleModule("grafana_", grafanaMod, tks.MakeStandard(grafanaPkg)))

	prov.SetAutonaming(255, "-")

	prov.MustApplyAutoAliases()

	return prov
}
