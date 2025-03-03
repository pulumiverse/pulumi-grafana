package grafana

import (
	"context"
	"fmt"
	"path/filepath"
	"unicode"

	// The linter requires unnamed imports to have a doc comment
	_ "embed"

	grafana "github.com/grafana/terraform-provider-grafana/v3/pkg/provider"

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
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
	// grafanaMod = "index" // the toplevel module
	// further modules follow the grouping of the upstream TF provider
	// https://registry.terraform.io/providers/grafana/grafana/latest/docs
	alertingMod            = "alerting"
	cloudMod               = "cloud"
	cloudProviderMod       = "cloudProvider"
	connectionsMod         = "connections"
	enterpriseMod          = "enterprise"
	fleetManagementMod     = "fleetManagement"
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
			},
			"grafana_message_template": {
				Tok: grafanaResource(alertingMod, "MessageTemplate"),
			},
			"grafana_mute_timing": {
				Tok: grafanaResource(alertingMod, "MuteTiming"),
			},
			"grafana_notification_policy": {
				Tok: grafanaResource(alertingMod, "NotificationPolicy"),
			},
			"grafana_rule_group": {
				Tok: grafanaResource(alertingMod, "RuleGroup"),
			},

			// Enterprise
			"grafana_data_source_config_lbac_rules": {
				Tok: grafanaResource(enterpriseMod, "DataSourceConfigLbacRules"),
			},
			"grafana_data_source_permission": {
				Tok: grafanaResource(enterpriseMod, "DataSourcePermission"),
			},
			"grafana_data_source_permission_item": {
				Tok: grafanaResource(enterpriseMod, "DataSourcePermissionItem"),
			},
			"grafana_report": {
				Tok: grafanaResource(enterpriseMod, "Report"),
			},
			"grafana_role": {
				Tok: grafanaResource(enterpriseMod, "Role"),
			},
			"grafana_role_assignment": {
				Tok: grafanaResource(enterpriseMod, "RoleAssignment"),
			},
			"grafana_role_assignment_item": {
				Tok: grafanaResource(enterpriseMod, "RoleAssignmentItem"),
			},
			"grafana_team_external_group": {
				Tok: grafanaResource(enterpriseMod, "TeamExternalGroup"),
			},

			// Fleet Management
			"grafana_fleet_management_collector": {
				ComputeID: tfbridge.DelegateIDField("id", "grafana", "https://github.com/pulumiveerse/pulumi-grafana"),
			},

			// OSS
			"grafana_annotation": {
				Tok: grafanaResource(ossMod, "Annotation"),
			},
			"grafana_dashboard": {
				Tok: grafanaResource(ossMod, "Dashboard"),
			},
			"grafana_dashboard_permission": {
				Tok: grafanaResource(ossMod, "DashboardPermission"),
			},
			"grafana_dashboard_permission_item": {
				Tok: grafanaResource(ossMod, "DashboardPermissionItem"),
			},
			"grafana_dashboard_public": {
				Tok: grafanaResource(ossMod, "DashboardPublic"),
			},
			"grafana_data_source": {
				Tok: grafanaResource(ossMod, "DataSource"),
			},
			"grafana_data_source_config": {
				Tok: grafanaResource(ossMod, "DataSourceConfig"),
			},
			"grafana_folder": {
				Tok: grafanaResource(ossMod, "Folder"),
			},
			"grafana_folder_permission": {
				Tok: grafanaResource(ossMod, "FolderPermission"),
			},
			"grafana_folder_permission_item": {
				Tok: grafanaResource(ossMod, "FolderPermissionItem"),
			},
			"grafana_library_panel": {
				Tok: grafanaResource(ossMod, "LibraryPanel"),
			},
			"grafana_organization": {
				Tok: grafanaResource(ossMod, "Organization"),
			},
			"grafana_organization_preferences": {
				Tok: grafanaResource(ossMod, "OrganizationPreferences"),
			},
			"grafana_playlist": {
				Tok: grafanaResource(ossMod, "Playlist"),
			},
			"grafana_service_account": {
				Tok: grafanaResource(ossMod, "ServiceAccount"),
			},
			"grafana_service_account_permission": {
				Tok: grafanaResource(ossMod, "ServiceAccountPermission"),
			},
			"grafana_service_account_permission_item": {
				Tok: grafanaResource(ossMod, "ServiceAccountPermissionItem"),
			},
			"grafana_service_account_token": {
				Tok: grafanaResource(ossMod, "ServiceAccountToken"),
			},
			"grafana_sso_settings": {
				Tok: grafanaResource(ossMod, "SsoSettings"),
			},
			"grafana_team": {
				Tok: grafanaResource(ossMod, "Team"),
			},
			"grafana_user": {
				Tok: grafanaResource(ossMod, "User"),
			},

			// SLO
			"grafana_slo": {
				Tok: grafanaResource(sloMod, "SLO"),
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Enterprise
			"grafana_role": {
				Tok: grafanaDataSource(enterpriseMod, "getRole"),
			},

			// OSS
			"grafana_dashboard": {
				Tok: grafanaDataSource(ossMod, "getDashboard"),
			},
			"grafana_dashboards": {
				Tok: grafanaDataSource(ossMod, "getDashboards"),
			},
			"grafana_data_source": {
				Tok: grafanaDataSource(ossMod, "getDataSource"),
			},
			"grafana_folder": {
				Tok: grafanaDataSource(ossMod, "getFolder"),
			},
			"grafana_folders": {
				Tok: grafanaDataSource(ossMod, "getFolders"),
			},
			"grafana_library_panel": {
				Tok: grafanaDataSource(ossMod, "getLibraryPanel"),
			},
			"grafana_organization": {
				Tok: grafanaDataSource(ossMod, "getOrganization"),
			},
			"grafana_organization_preferences": {
				Tok: grafanaDataSource(ossMod, "getOrganizationPreferences"),
			},
			"grafana_service_account": {
				Tok: grafanaDataSource(ossMod, "getServiceAccount"),
			},
			"grafana_team": {
				Tok: grafanaDataSource(ossMod, "getTeam"),
			},
			"grafana_user": {
				Tok: grafanaDataSource(ossMod, "getUser"),
			},
			"grafana_users": {
				Tok: grafanaDataSource(ossMod, "getUsers"),
			},

			// SLO
			"grafana_slos": {
				Tok: grafanaDataSource(sloMod, "getSlos"),
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
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			PackageName:          "pulumiverse_grafana",
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
			RootNamespace:        "Pulumiverse",
			RespectSchemaVersion: true,
		},
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
	}

	prov.MustComputeTokens(
		tks.MappedModules(
			"grafana_",
			"index",
			map[string]string{
				"alerting":             alertingMod,
				"cloud":                cloudMod,
				"cloud_provider":       cloudProviderMod,
				"connections":          connectionsMod,
				"enterprise":           enterpriseMod,
				"fleet_management":     fleetManagementMod,
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
