package grafana

import (
	"context"
	"fmt"
	"path/filepath"
	"unicode"

	// The linter requires unnamed imports to have a doc comment
	_ "embed"

	grafana "github.com/grafana/terraform-provider-grafana/v3/pkg/provider"

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
	// further modules follow the grouping of the upstream TF provider
	// https://registry.terraform.io/providers/grafana/grafana/latest/docs
	alertingMod            = "alerting"
	cloudMod               = "cloud"
	enterpriseMod          = "enterprise"
	mlMod                  = "machineLearning"
	oncallMod              = "onCall"
	ossMod                 = "oss"
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

func preConfigureCallback(_ resource.PropertyMap, _ shim.ResourceConfig) error {
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
		P:           p,
		Name:        "grafana",
		DisplayName: "Grafana",
		Publisher:   "pulumiverse",
		//nolint:lll
		LogoURL:                 "https://raw.githubusercontent.com/pulumiverse/pulumi-grafana/main/assets/grafana.png",
		PluginDownloadURL:       "github://api.github.com/pulumiverse",
		Version:                 version.Version,
		Description:             "A Pulumi package for creating and managing grafana.",
		Keywords:                []string{"pulumi", "grafana", "pulumiverse"},
		License:                 "Apache-2.0",
		Homepage:                "https://grafana.com",
		Repository:              "https://github.com/pulumiverse/pulumi-grafana",
		GitHubOrg:               "grafana",
		TFProviderModuleVersion: "v3",
		Config: map[string]*tfbridge.SchemaInfo{
			"auth": {
				Type:   "string",
				Secret: boolRef(true),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_AUTH"},
				},
			},
			"ca_cert": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_CA_CERT"},
				},
			},
			"cloud_access_policy_token": {
				Type:   "string",
				Secret: boolRef(true),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_CLOUD_ACCESS_POLICY_TOKEN"},
				},
			},
			"cloud_api_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_CLOUD_API_URL"},
				},
			},
			"http_headers": {
				SuppressEmptyMapElements: boolRef(true),
				Omit:                     true,
			},
			"insecure_skip_verify": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_INSECURE_SKIP_VERIFY"},
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
			"retries": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_RETRIES"},
				},
			},
			// "retry_status_codes": {
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"GRAFANA_RETRY_STATUS_CODES"},
			// 	},
			// },
			"retry_wait": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_RETRY_WAIT"},
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
			"tls_cert": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_TLS_CERT"},
				},
			},
			"tls_key": {
				Secret: boolRef(true),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_TLS_KEY"},
				},
			},
			"url": {
				Type: "string",
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_URL"},
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
				Docs: &tfbridge.DocInfo{
					Source: "contact_point.md",
				},
			},
			"grafana_message_template": {
				Tok: grafanaResource(alertingMod, "MessageTemplate"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "MessageTemplate"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "message_template.md",
				},
			},
			"grafana_mute_timing": {
				Tok: grafanaResource(alertingMod, "MuteTiming"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "MuteTiming"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "mute_timing.md",
				},
			},
			"grafana_notification_policy": {
				Tok: grafanaResource(alertingMod, "NotificationPolicy"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "NotificationPolicy"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "notification_policy.md",
				},
			},
			"grafana_rule_group": {
				Tok: grafanaResource(alertingMod, "RuleGroup"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "RuleGroup"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "rule_group.md",
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
				Docs: &tfbridge.DocInfo{
					Source: "cloud_access_policy.md",
				},
			},
			"grafana_cloud_access_policy_token": {
				Tok: grafanaResource(cloudMod, "AccessPolicyToken"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "CloudAccessPolicyToken"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "cloud_access_policy_token.md",
				},
			},
			"grafana_cloud_org_member": {
				Tok: grafanaResource(cloudMod, "OrgMember"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "CloudOrgMember"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "cloud_org_member.md",
				},
			},
			"grafana_cloud_plugin_installation": {
				Tok: grafanaResource(cloudMod, "PluginInstallation"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "CloudPluginInstallation"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "cloud_plugin_installation.md",
				},
			},
			"grafana_cloud_stack": {
				Tok: grafanaResource(cloudMod, "Stack"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "CloudStack"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "cloud_stack.md",
				},
			},
			"grafana_cloud_stack_service_account": {
				Tok: grafanaResource(cloudMod, "StackServiceAccount"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "CloudStackServiceAccount"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "cloud_stack_service_account.md",
				},
			},
			"grafana_cloud_stack_service_account_token": {
				Tok: grafanaResource(cloudMod, "StackServiceAccountToken"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "CloudStackServiceAccountToken"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "cloud_stack_service_account_token.md",
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
				Docs: &tfbridge.DocInfo{
					Source: "data_source_permission.md",
				},
			},
			"grafana_data_source_permission_item": {
				Tok: grafanaResource(enterpriseMod, "DataSourcePermissionItem"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "DataSourcePermissionItem"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "data_source_permission_item.md",
				},
			},
			"grafana_report": {
				Tok: grafanaResource(enterpriseMod, "Report"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "Report"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "report.md",
				},
			},
			"grafana_role": {
				Tok: grafanaResource(enterpriseMod, "Role"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "Role"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "role.md",
				},
			},
			"grafana_role_assignment": {
				Tok: grafanaResource(enterpriseMod, "RoleAssignment"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "RoleAssignment"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "role_assignment.md",
				},
			},
			"grafana_role_assignment_item": {
				Tok: grafanaResource(enterpriseMod, "RoleAssignmentItem"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "RoleAssignmentItem"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "role_assignment_item.md",
				},
			},
			"grafana_team_external_group": {
				Tok: grafanaResource(enterpriseMod, "TeamExternalGroup"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "TeamExternalGroup"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "team_external_group.md",
				},
			},

			// OSS
			"grafana_annotation": {
				Tok: grafanaResource(ossMod, "Annotation"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "Annotation"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "annotation.md",
				},
			},
			"grafana_dashboard": {
				Tok: grafanaResource(ossMod, "Dashboard"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "Dashboard"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "dashboard.md",
				},
			},
			"grafana_dashboard_permission": {
				Tok: grafanaResource(ossMod, "DashboardPermission"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "DashboardPermission"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "dashboard_permission.md",
				},
			},
			"grafana_dashboard_permission_item": {
				Tok: grafanaResource(ossMod, "DashboardPermissionItem"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "DashboardPermissionItem"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "dashboard_permission_item.md",
				},
			},
			"grafana_dashboard_public": {
				Tok: grafanaResource(ossMod, "DashboardPublic"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "DashboardPublic"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "dashboard_public.md",
				},
			},
			"grafana_data_source": {
				Tok: grafanaResource(ossMod, "DataSource"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "DataSource"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "data_source.md",
				},
			},
			"grafana_data_source_config": {
				Tok: grafanaResource(ossMod, "DataSourceConfig"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "DataSourceConfig"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "data_source_config.md",
				},
			},
			"grafana_folder": {
				Tok: grafanaResource(ossMod, "Folder"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "Folder"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "folder.md",
				},
			},
			"grafana_folder_permission": {
				Tok: grafanaResource(ossMod, "FolderPermission"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "FolderPermission"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "folder_permission.md",
				},
			},
			"grafana_folder_permission_item": {
				Tok: grafanaResource(ossMod, "FolderPermissionItem"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "FolderPermissionItem"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "folder_permission_item.md",
				},
			},
			"grafana_library_panel": {
				Tok: grafanaResource(ossMod, "LibraryPanel"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "LibraryPanel"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "library_panel.md",
				},
			},
			"grafana_organization": {
				Tok: grafanaResource(ossMod, "Organization"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "Organization"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "organization.md",
				},
			},
			"grafana_organization_preferences": {
				Tok: grafanaResource(ossMod, "OrganizationPreferences"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "OrganizationPreferences"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "organization_preferences.md",
				},
			},
			"grafana_playlist": {
				Tok: grafanaResource(ossMod, "Playlist"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "Playlist"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "playlist.md",
				},
			},
			"grafana_service_account": {
				Tok: grafanaResource(ossMod, "ServiceAccount"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "ServiceAccount"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "service_account.md",
				},
			},
			"grafana_service_account_permission": {
				Tok: grafanaResource(ossMod, "ServiceAccountPermission"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "ServiceAccountPermission"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "service_account_permission.md",
				},
			},
			"grafana_service_account_permission_item": {
				Tok: grafanaResource(ossMod, "ServiceAccountPermissionItem"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "ServiceAccountPermissionItem"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "service_account_permission_item.md",
				},
			},
			"grafana_service_account_token": {
				Tok: grafanaResource(ossMod, "ServiceAccountToken"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "ServiceAccountToken"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "service_account.md",
				},
			},
			"grafana_sso_settings": {
				Tok: grafanaResource(ossMod, "SsoSettings"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "SsoSettings"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "sso_settings.md",
				},
			},
			"grafana_team": {
				Tok: grafanaResource(ossMod, "Team"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "Team"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "team.md",
				},
			},
			"grafana_user": {
				Tok: grafanaResource(ossMod, "User"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "User"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "user.md",
				},
			},

			// Machine Learning
			"grafana_machine_learning_holiday": {
				Tok: grafanaResource(mlMod, "Holiday"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "MachineLearningHoliday"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "machine_learning_holiday.md",
				},
			},
			"grafana_machine_learning_job": {
				Tok: grafanaResource(mlMod, "Job"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "MachineLearningJob"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "machine_learning_job.md",
				},
			},
			"grafana_machine_learning_outlier_detector": {
				Tok: grafanaResource(mlMod, "OutlierDetector"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "MachineLearningOutlierDetector"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "machine_learning_outlier_detector.md",
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
				Docs: &tfbridge.DocInfo{
					Source: "oncall_escalation.md",
				},
			},
			"grafana_oncall_escalation_chain": {
				Tok: grafanaResource(oncallMod, "EscalationChain"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "OncallEscalationChain"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "oncall_escalation_chain.md",
				},
			},
			"grafana_oncall_integration": {
				Tok: grafanaResource(oncallMod, "Integration"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "OncallIntegration"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "oncall_integration.md",
				},
			},
			"grafana_oncall_on_call_shift": {
				Tok: grafanaResource(oncallMod, "OnCallShift"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "OncallOnCallShift"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "oncall_on_call_shift.md",
				},
			},
			"grafana_oncall_outgoing_webhook": {
				Tok: grafanaResource(oncallMod, "OutgoingWebhook"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "OncallOutgoingWebhook"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "oncall_outgoing_webhook.md",
				},
			},
			"grafana_oncall_route": {
				Tok: grafanaResource(oncallMod, "Route"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "OncallRoute"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "oncall_route.md",
				},
			},
			"grafana_oncall_schedule": {
				Tok: grafanaResource(oncallMod, "Schedule"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "OncallSchedule"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "oncall_schedule.md",
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
				Docs: &tfbridge.DocInfo{
					Source: "slo.md",
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
				Docs: &tfbridge.DocInfo{
					Source: "synthetic_monitoring_check.md",
				},
			},
			"grafana_synthetic_monitoring_installation": {
				Tok: grafanaResource(syntheticMonitoringMod, "Installation"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "SyntheticMonitoringInstallation"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "synthetic_monitoring_installation.md",
				},
			},
			"grafana_synthetic_monitoring_probe": {
				Tok: grafanaResource(syntheticMonitoringMod, "Probe"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(grafanaMod, "SyntheticMonitoringProbe"),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "synthetic_monitoring_probe.md",
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Alerting

			// Cloud
			"grafana_cloud_ips": {
				Tok: grafanaDataSource(cloudMod, "getIps"),
				Docs: &tfbridge.DocInfo{
					Source: "cloud_ips.md",
				},
			},
			"grafana_cloud_organization": {
				Tok: grafanaDataSource(cloudMod, "getOrganization"),
				Docs: &tfbridge.DocInfo{
					Source: "cloud_organization.md",
				},
			},
			"grafana_cloud_stack": {
				Tok: grafanaDataSource(cloudMod, "getStack"),
				Docs: &tfbridge.DocInfo{
					Source: "cloud_stack.md",
				},
			},

			// Enterprise
			"grafana_role": {
				Tok: grafanaDataSource(enterpriseMod, "getRole"),
				Docs: &tfbridge.DocInfo{
					Source: "role.md",
				},
			},

			// OSS
			"grafana_dashboard": {
				Tok: grafanaDataSource(ossMod, "getDashboard"),
				Docs: &tfbridge.DocInfo{
					Source: "dashboard.md",
				},
			},
			"grafana_dashboards": {
				Tok: grafanaDataSource(ossMod, "getDashboards"),
				Docs: &tfbridge.DocInfo{
					Source: "dashboards.md",
				},
			},
			"grafana_data_source": {
				Tok: grafanaDataSource(ossMod, "getDataSource"),
				Docs: &tfbridge.DocInfo{
					Source: "data_source.md",
				},
			},
			"grafana_folder": {
				Tok: grafanaDataSource(ossMod, "getFolder"),
				Docs: &tfbridge.DocInfo{
					Source: "folder.md",
				},
			},
			"grafana_folders": {
				Tok: grafanaDataSource(ossMod, "getFolders"),
				Docs: &tfbridge.DocInfo{
					Source: "folders.md",
				},
			},
			"grafana_library_panel": {
				Tok: grafanaDataSource(ossMod, "getLibraryPanel"),
				Docs: &tfbridge.DocInfo{
					Source: "library_panel.md",
				},
			},
			"grafana_organization": {
				Tok: grafanaDataSource(ossMod, "getOrganization"),
				Docs: &tfbridge.DocInfo{
					Source: "organization.md",
				},
			},
			"grafana_organization_preferences": {
				Tok: grafanaDataSource(ossMod, "getOrganizationPreferences"),
				Docs: &tfbridge.DocInfo{
					Source: "organization_preferences.md",
				},
			},
			"grafana_service_account": {
				Tok: grafanaDataSource(ossMod, "getServiceAccount"),
				Docs: &tfbridge.DocInfo{
					Source: "service_account.md",
				},
			},
			"grafana_team": {
				Tok: grafanaDataSource(ossMod, "getTeam"),
				Docs: &tfbridge.DocInfo{
					Source: "team.md",
				},
			},
			"grafana_user": {
				Tok: grafanaDataSource(ossMod, "getUser"),
				Docs: &tfbridge.DocInfo{
					Source: "user.md",
				},
			},
			"grafana_users": {
				Tok: grafanaDataSource(ossMod, "getUsers"),
				Docs: &tfbridge.DocInfo{
					Source: "users.md",
				},
			},

			// Machine Learning

			// OnCall
			"grafana_oncall_escalation_chain": {
				Tok: grafanaDataSource(oncallMod, "getEscalationChain"),
				Docs: &tfbridge.DocInfo{
					Source: "oncall_escalation_chain.md",
				},
			},
			"grafana_oncall_integration": {
				Tok: grafanaDataSource(oncallMod, "getIntegration"),
				Docs: &tfbridge.DocInfo{
					Source: "oncall_integration.md",
				},
			},
			"grafana_oncall_outgoing_webhook": {
				Tok: grafanaDataSource(oncallMod, "getOutgoingWebhook"),
				Docs: &tfbridge.DocInfo{
					Source: "oncall_outgoing_webhook.md",
				},
			},
			"grafana_oncall_schedule": {
				Tok: grafanaDataSource(oncallMod, "getSchedule"),
				Docs: &tfbridge.DocInfo{
					Source: "oncall_schedule.md",
				},
			},
			"grafana_oncall_slack_channel": {
				Tok: grafanaDataSource(oncallMod, "getSlackChannel"),
				Docs: &tfbridge.DocInfo{
					Source: "oncall_slack_channel.md",
				},
			},
			"grafana_oncall_team": {
				Tok: grafanaDataSource(oncallMod, "getTeam"),
				Docs: &tfbridge.DocInfo{
					Source: "oncall_team.md",
				},
			},
			"grafana_oncall_user": {
				Tok: grafanaDataSource(oncallMod, "getUser"),
				Docs: &tfbridge.DocInfo{
					Source: "oncall_user.md",
				},
			},
			"grafana_oncall_user_group": {
				Tok: grafanaDataSource(oncallMod, "getUserGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "oncall_user_group.md",
				},
			},

			// SLO
			"grafana_slos": {
				Tok: grafanaDataSource(sloMod, "getSlos"),
				Docs: &tfbridge.DocInfo{
					Source: "slos.md",
				},
			},

			// Synthetic Monitoring
			"grafana_synthetic_monitoring_probe": {
				Tok: grafanaDataSource(syntheticMonitoringMod, "getProbe"),
				Docs: &tfbridge.DocInfo{
					Source: "synthetic_monitoring_probe.md",
				},
			},
			"grafana_synthetic_monitoring_probes": {
				Tok: grafanaDataSource(syntheticMonitoringMod, "getProbes"),
				Docs: &tfbridge.DocInfo{
					Source: "synthetic_monitoring_probes.md",
				},
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
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			PackageName: "pulumiverse_grafana",
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
			PyProject:            struct{ Enabled bool }{true},
			RespectSchemaVersion: true,
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumiverse/pulumi-%[1]s/sdk/", grafanaPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				grafanaPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Pulumiverse",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			RespectSchemaVersion: true,
		},
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
	}

	prov.MustComputeTokens(tks.SingleModule("grafana_", grafanaMod, tks.MakeStandard(grafanaPkg)))

	prov.SetAutonaming(255, "-")

	prov.MustApplyAutoAliases()

	return prov
}
