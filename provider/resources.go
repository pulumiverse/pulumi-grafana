package grafana

import (
	"context"
	"fmt"
	"path/filepath"
	"unicode"

	// The linter requires unnamed imports to have a doc comment
	_ "embed"

	grafana "github.com/grafana/terraform-provider-grafana/v4/pkg/provider"

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tks "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumiverse/pulumi-grafana/provider/v2/pkg/version"
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
	alertingMod              = "alerting"
	appsMod                  = "apps"
	assertsMod               = "assert"
	cloudMod                 = "cloud"
	cloudProviderMod         = "cloudProvider"
	connectionsMod           = "connections"
	enterpriseMod            = "enterprise"
	fleetManagementMod       = "fleetManagement"
	frontendObservabilityMod = "frontendObservability"
	k6Mod                    = "k6"
	mlMod                    = "machineLearning"
	oncallMod                = "onCall"
	ossMod                   = "oss"
	sloMod                   = "slo"
	syntheticMonitoringMod   = "syntheticMonitoring"
	// Experimental resources
	experimentalMod = "experimental"
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

// grafanaVersionedResource manufactures a versioned resource token given a module and resource name.
// It verifies the
func grafanaVersionedResource(mod string, version string, res string) tokens.Type {
	return grafanaResource(mod+"/"+version, res)
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
		TFProviderModuleVersion: "v4",
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
			"http_headers": {},
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
			"grafana_apps_alertenrichment_alertenrichment_v1beta1": {
				Tok: grafanaVersionedResource(alertingMod, "v1beta1", "AlertEnrichment"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(alertingMod, "AlertEnrichment"),
					},
				},
			},
			"grafana_apps_notifications_inhibitionrule_v1beta1": {
				Tok: grafanaVersionedResource(alertingMod, "v1beta1", "NotificationsInhibitionRule"),
			},
			"grafana_apps_rules_alertrule_v0alpha1": {
				Tok: grafanaVersionedResource(alertingMod, "v0alpha1", "AlertRule"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(alertingMod, "AlertRuleV0Alpha1"),
					},
				},
			},
			"grafana_apps_rules_recordingrule_v0alpha1": {
				Tok: grafanaVersionedResource(alertingMod, "v0alpha1", "RecordingRule"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(alertingMod, "RecordingRuleV0Alpha1"),
					},
				},
			},
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

			// Apps
			"grafana_apps_dashboard_dashboard_v1beta1": {
				Tok: grafanaVersionedResource(appsMod, "v1beta1", "Dashboard"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(experimentalMod, "AppsDashboard"),
					},
				},
			},
			"grafana_apps_dashboard_dashboard_v2beta1": {
				Tok: grafanaVersionedResource(appsMod, "v2beta1", "Dashboard"),
			},
			"grafana_apps_dashboard_dashboard_v2": {
				Tok: grafanaVersionedResource(appsMod, "v2", "Dashboard"),
			},
			"grafana_apps_generic_resource": {
				Tok: grafanaResource(appsMod, "GenericResource"),
			},
			"grafana_apps_playlist_playlist_v0alpha1": {
				Tok: grafanaVersionedResource(appsMod, "v0alpha1", "Playlist"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(experimentalMod, "AppsPlaylistV0Alpha1"),
					},
				},
			},
			"grafana_apps_playlist_playlist_v1": {
				Tok: grafanaVersionedResource(appsMod, "v1", "Playlist"),
			},
			"grafana_apps_provisioning_connection_v0alpha1": {
				Tok: grafanaVersionedResource(appsMod, "v0alpha1", "ProvisioningConnection"),
			},
			"grafana_apps_provisioning_repository_v0alpha1": {
				Tok: grafanaVersionedResource(appsMod, "v0alpha1", "ProvisioningRepository"),
			},

			// Cloud
			"grafana_apps_productactivation_appo11yconfig_v1alpha1": {
				Tok: grafanaVersionedResource(cloudMod, "v1alpha1", "ProductActivationAppO11yConfig"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(cloudMod, "ProductActivationAppO11yConfigV1Alpha1"),
					},
				},
			},
			"grafana_apps_productactivation_dbo11yconfig_v1alpha1": {
				Tok: grafanaVersionedResource(cloudMod, "v1alpha1", "ProductActivationDbO11yConfig"),
			},
			"grafana_apps_productactivation_k8so11yconfig_v1alpha1": {
				Tok: grafanaVersionedResource(cloudMod, "v1alpha1", "ProductActivationK8sO11yConfig"),
				Aliases: []tfbridge.AliasInfo{
					{
						Type: grafanaResourceAlias(cloudMod, "ProductActivationK8sO11yConfigV1Alpha1"),
					},
				},
			},

			// Enterprise
			"grafana_apps_secret_keeper_activation_v1beta1": {
				Tok: grafanaVersionedResource(enterpriseMod, "v1beta1", "SecretKeeperActivation"),
			},
			"grafana_apps_secret_keeper_v1beta1": {
				Tok: grafanaVersionedResource(enterpriseMod, "v1beta1", "SecretKeeper"),
			},
			"grafana_apps_secret_securevalue_v1beta1": {
				Tok: grafanaVersionedResource(enterpriseMod, "v1beta1", "SecretSecureValue"),
			},
			"grafana_data_source_config_lbac_rules": {
				Tok: grafanaResource(enterpriseMod, "DataSourceConfigLbacRules"),
				Docs: &tfbridge.DocInfo{
					Source: "data_source_config_lbac_rules.md",
				},
			},
			"grafana_data_source_cache_config": {
				Tok: grafanaResource(enterpriseMod, "DataSourceCacheConfig"),
			},
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
			"grafana_scim_config": {
				Tok: grafanaResource(enterpriseMod, "ScimConfig"),
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

			// Fleet Management
			"grafana_fleet_management_collector": {
				ComputeID: tfbridge.DelegateIDField("id", "grafana", "https://github.com/pulumiverse/pulumi-grafana"),
			},

			// Frontend Observability
			"grafana_frontend_o11y_app": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {Type: "string"},
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
			"grafana_service_account_rotating_token": {
				Tok: grafanaResource(ossMod, "ServiceAccountRotatingToken"),
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
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
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
			"grafana_library_panels": {
				Tok: grafanaDataSource(ossMod, "getLibraryPanels"),
				Docs: &tfbridge.DocInfo{
					Source: "library_panels.md",
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
			"grafana_organization_user": {
				Tok: grafanaDataSource(ossMod, "getOrganizationUser"),
				Docs: &tfbridge.DocInfo{
					Source: "organization_user.md",
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

			// SLO
			"grafana_slos": {
				Tok: grafanaDataSource(sloMod, "getSlos"),
				Docs: &tfbridge.DocInfo{
					Source: "slos.md",
				},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@pulumiverse/grafana",
			// List any npm dependencies and their versions
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
			RespectSchemaVersion: true,
			ModuleToPackage: map[string]string{
				"alerting/v0alpha1":  "alerting/v0alpha1",
				"alerting/v1beta1":   "alerting/v1beta1",
				"apps/v0alpha1":      "apps/v0alpha1",
				"apps/v1beta1":       "apps/v1beta1",
				"apps/v2beta1":       "apps/v2beta1",
				"cloud/v1alpha1":     "cloud/v1alpha1",
				"enterprise/v1beta1": "enterprise/v1beta1",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			PackageName:          "pulumiverse_grafana",
			PyProject:            struct{ Enabled bool }{true},
			RespectSchemaVersion: true,
			ModuleNameOverrides: map[string]string{
				"alerting/v0alpha1":  "alerting/v0alpha1",
				"alerting/v1beta1":   "alerting/v1beta1",
				"apps/v0alpha1":      "apps/v0alpha1",
				"apps/v1beta1":       "apps/v1beta1",
				"apps/v2beta1":       "apps/v2beta1",
				"cloud/v1alpha1":     "cloud/v1alpha1",
				"enterprise/v1beta1": "enterprise/v1beta1",
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
			RespectSchemaVersion:           true,
			ModuleToPackage: map[string]string{
				"alerting/v0alpha1":  "alerting/v0alpha1",
				"alerting/v1beta1":   "alerting/v1beta1",
				"apps/v0alpha1":      "apps/v0alpha1",
				"apps/v1beta1":       "apps/v1beta1",
				"apps/v2beta1":       "apps/v2beta1",
				"cloud/v1alpha1":     "cloud/v1alpha1",
				"enterprise/v1beta1": "enterprise/v1beta1",
			},
			PackageImportAliases: map[string]string{
				"github.com/pulumiverse/pulumi-grafana/sdk/v2/go/grafana/alerting/v0alpha1":  "alertingv0alpha1",
				"github.com/pulumiverse/pulumi-grafana/sdk/v2/go/grafana/alerting/v1beta1":   "alertingv1beta1",
				"github.com/pulumiverse/pulumi-grafana/sdk/v2/go/grafana/apps/v0alpha1":      "appsv0alpha1",
				"github.com/pulumiverse/pulumi-grafana/sdk/v2/go/grafana/apps/v1beta1":       "appsv1beta1",
				"github.com/pulumiverse/pulumi-grafana/sdk/v2/go/grafana/apps/v2beta1":       "appsv2beta1",
				"github.com/pulumiverse/pulumi-grafana/sdk/v2/go/grafana/cloud/v1alpha1":     "cloudv1alpha1",
				"github.com/pulumiverse/pulumi-grafana/sdk/v2/go/grafana/enterprise/v1beta1": "enterprisev1beta1",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace:        "Pulumiverse",
			RespectSchemaVersion: true,
			Namespaces: map[string]string{
				"alerting/v0alpha1":  "Alerting.V0Alpha1",
				"alerting/v1beta1":   "Alerting.V1Beta1",
				"apps/v0alpha1":      "Apps.V0Alpha1",
				"apps/v1beta1":       "Apps.V1Beta1",
				"apps/v2beta1":       "Apps.V2Beta1",
				"cloud/v1alpha1":     "Cloud.V1Alpha1",
				"enterprise/v1beta1": "Enterprise.V1Beta1",
			},
		},
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
	}

	prov.MustComputeTokens(
		tks.MappedModules(
			"grafana_",
			"index",
			map[string]string{
				"alerting":             alertingMod,
				"asserts":              assertsMod,
				"cloud":                cloudMod,
				"cloud_provider":       cloudProviderMod,
				"connections":          connectionsMod,
				"enterprise":           enterpriseMod,
				"fleet_management":     fleetManagementMod,
				"frontend_o11y":        frontendObservabilityMod,
				"k6":                   k6Mod,
				"machine_learning":     mlMod,
				"oncall":               oncallMod,
				"oss":                  ossMod,
				"slo":                  sloMod,
				"synthetic_monitoring": syntheticMonitoringMod,
			},
			tks.MakeStandard(grafanaPkg),
		),
	)
	prov.SetAutonaming(255, "-")
	prov.MustApplyAutoAliases()

	return prov
}
