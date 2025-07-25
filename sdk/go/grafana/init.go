// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

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
	case "grafana:index/annotation:Annotation":
		r = &Annotation{}
	case "grafana:index/cloudAccessPolicy:CloudAccessPolicy":
		r = &CloudAccessPolicy{}
	case "grafana:index/cloudAccessPolicyToken:CloudAccessPolicyToken":
		r = &CloudAccessPolicyToken{}
	case "grafana:index/cloudOrgMember:CloudOrgMember":
		r = &CloudOrgMember{}
	case "grafana:index/cloudPluginInstallation:CloudPluginInstallation":
		r = &CloudPluginInstallation{}
	case "grafana:index/cloudStack:CloudStack":
		r = &CloudStack{}
	case "grafana:index/cloudStackServiceAccount:CloudStackServiceAccount":
		r = &CloudStackServiceAccount{}
	case "grafana:index/cloudStackServiceAccountToken:CloudStackServiceAccountToken":
		r = &CloudStackServiceAccountToken{}
	case "grafana:index/connectionsMetricsEndpointScrapeJob:ConnectionsMetricsEndpointScrapeJob":
		r = &ConnectionsMetricsEndpointScrapeJob{}
	case "grafana:index/contactPoint:ContactPoint":
		r = &ContactPoint{}
	case "grafana:index/dashboard:Dashboard":
		r = &Dashboard{}
	case "grafana:index/dashboardPermission:DashboardPermission":
		r = &DashboardPermission{}
	case "grafana:index/dashboardPermissionItem:DashboardPermissionItem":
		r = &DashboardPermissionItem{}
	case "grafana:index/dashboardPublic:DashboardPublic":
		r = &DashboardPublic{}
	case "grafana:index/dataSource:DataSource":
		r = &DataSource{}
	case "grafana:index/dataSourceConfig:DataSourceConfig":
		r = &DataSourceConfig{}
	case "grafana:index/dataSourcePermission:DataSourcePermission":
		r = &DataSourcePermission{}
	case "grafana:index/dataSourcePermissionItem:DataSourcePermissionItem":
		r = &DataSourcePermissionItem{}
	case "grafana:index/folder:Folder":
		r = &Folder{}
	case "grafana:index/folderPermission:FolderPermission":
		r = &FolderPermission{}
	case "grafana:index/folderPermissionItem:FolderPermissionItem":
		r = &FolderPermissionItem{}
	case "grafana:index/libraryPanel:LibraryPanel":
		r = &LibraryPanel{}
	case "grafana:index/machineLearningHoliday:MachineLearningHoliday":
		r = &MachineLearningHoliday{}
	case "grafana:index/machineLearningJob:MachineLearningJob":
		r = &MachineLearningJob{}
	case "grafana:index/machineLearningOutlierDetector:MachineLearningOutlierDetector":
		r = &MachineLearningOutlierDetector{}
	case "grafana:index/messageTemplate:MessageTemplate":
		r = &MessageTemplate{}
	case "grafana:index/muteTiming:MuteTiming":
		r = &MuteTiming{}
	case "grafana:index/notificationPolicy:NotificationPolicy":
		r = &NotificationPolicy{}
	case "grafana:index/oncallEscalation:OncallEscalation":
		r = &OncallEscalation{}
	case "grafana:index/oncallEscalationChain:OncallEscalationChain":
		r = &OncallEscalationChain{}
	case "grafana:index/oncallIntegration:OncallIntegration":
		r = &OncallIntegration{}
	case "grafana:index/oncallOnCallShift:OncallOnCallShift":
		r = &OncallOnCallShift{}
	case "grafana:index/oncallOutgoingWebhook:OncallOutgoingWebhook":
		r = &OncallOutgoingWebhook{}
	case "grafana:index/oncallRoute:OncallRoute":
		r = &OncallRoute{}
	case "grafana:index/oncallSchedule:OncallSchedule":
		r = &OncallSchedule{}
	case "grafana:index/oncallUserNotificationRule:OncallUserNotificationRule":
		r = &OncallUserNotificationRule{}
	case "grafana:index/organization:Organization":
		r = &Organization{}
	case "grafana:index/organizationPreferences:OrganizationPreferences":
		r = &OrganizationPreferences{}
	case "grafana:index/playlist:Playlist":
		r = &Playlist{}
	case "grafana:index/report:Report":
		r = &Report{}
	case "grafana:index/role:Role":
		r = &Role{}
	case "grafana:index/roleAssignment:RoleAssignment":
		r = &RoleAssignment{}
	case "grafana:index/roleAssignmentItem:RoleAssignmentItem":
		r = &RoleAssignmentItem{}
	case "grafana:index/ruleGroup:RuleGroup":
		r = &RuleGroup{}
	case "grafana:index/sLO:SLO":
		r = &SLO{}
	case "grafana:index/serviceAccount:ServiceAccount":
		r = &ServiceAccount{}
	case "grafana:index/serviceAccountPermission:ServiceAccountPermission":
		r = &ServiceAccountPermission{}
	case "grafana:index/serviceAccountPermissionItem:ServiceAccountPermissionItem":
		r = &ServiceAccountPermissionItem{}
	case "grafana:index/serviceAccountToken:ServiceAccountToken":
		r = &ServiceAccountToken{}
	case "grafana:index/ssoSettings:SsoSettings":
		r = &SsoSettings{}
	case "grafana:index/syntheticMonitoringCheck:SyntheticMonitoringCheck":
		r = &SyntheticMonitoringCheck{}
	case "grafana:index/syntheticMonitoringInstallation:SyntheticMonitoringInstallation":
		r = &SyntheticMonitoringInstallation{}
	case "grafana:index/syntheticMonitoringProbe:SyntheticMonitoringProbe":
		r = &SyntheticMonitoringProbe{}
	case "grafana:index/team:Team":
		r = &Team{}
	case "grafana:index/teamExternalGroup:TeamExternalGroup":
		r = &TeamExternalGroup{}
	case "grafana:index/user:User":
		r = &User{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:grafana" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"grafana",
		"index/annotation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/cloudAccessPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/cloudAccessPolicyToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/cloudOrgMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/cloudPluginInstallation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/cloudStack",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/cloudStackServiceAccount",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/cloudStackServiceAccountToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/connectionsMetricsEndpointScrapeJob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/contactPoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/dashboard",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/dashboardPermission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/dashboardPermissionItem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/dashboardPublic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/dataSource",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/dataSourceConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/dataSourcePermission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/dataSourcePermissionItem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/folder",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/folderPermission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/folderPermissionItem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/libraryPanel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/machineLearningHoliday",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/machineLearningJob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/machineLearningOutlierDetector",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/messageTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/muteTiming",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/notificationPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/oncallEscalation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/oncallEscalationChain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/oncallIntegration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/oncallOnCallShift",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/oncallOutgoingWebhook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/oncallRoute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/oncallSchedule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/oncallUserNotificationRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/organization",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/organizationPreferences",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/playlist",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/report",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/role",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/roleAssignment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/roleAssignmentItem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/ruleGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/sLO",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/serviceAccount",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/serviceAccountPermission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/serviceAccountPermissionItem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/serviceAccountToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/ssoSettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/syntheticMonitoringCheck",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/syntheticMonitoringInstallation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/syntheticMonitoringProbe",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/team",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/teamExternalGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"index/user",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"grafana",
		&pkg{version},
	)
}
