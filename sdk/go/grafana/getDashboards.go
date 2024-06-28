// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// Datasource for retrieving all dashboards. Specify list of folder IDs to search in for dashboards.
//
// * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/)
// * [Folder/Dashboard Search HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/folder_dashboard_search/)
// * [Dashboard HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/dashboard/)
//
// Deprecated: grafana.index/getdashboards.getDashboards has been deprecated in favor of grafana.oss/getdashboards.getDashboards
func GetDashboards(ctx *pulumi.Context, args *GetDashboardsArgs, opts ...pulumi.InvokeOption) (*GetDashboardsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDashboardsResult
	err := ctx.Invoke("grafana:index/getDashboards:getDashboards", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDashboards.
type GetDashboardsArgs struct {
	// UIDs of Grafana folders containing dashboards. Specify to filter for dashboards by folder (eg. `["General"]` for General folder), or leave blank to get all dashboards in all folders.
	FolderUids []string `pulumi:"folderUids"`
	// Maximum number of dashboard search results to return. Defaults to `5000`.
	Limit *int `pulumi:"limit"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// List of string Grafana dashboard tags to search for, eg. `["prod"]`. Used only as search input, i.e., attribute value will remain unchanged.
	Tags []string `pulumi:"tags"`
}

// A collection of values returned by getDashboards.
type GetDashboardsResult struct {
	Dashboards []GetDashboardsDashboard `pulumi:"dashboards"`
	// UIDs of Grafana folders containing dashboards. Specify to filter for dashboards by folder (eg. `["General"]` for General folder), or leave blank to get all dashboards in all folders.
	FolderUids []string `pulumi:"folderUids"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Maximum number of dashboard search results to return. Defaults to `5000`.
	Limit *int `pulumi:"limit"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// List of string Grafana dashboard tags to search for, eg. `["prod"]`. Used only as search input, i.e., attribute value will remain unchanged.
	Tags []string `pulumi:"tags"`
}

func GetDashboardsOutput(ctx *pulumi.Context, args GetDashboardsOutputArgs, opts ...pulumi.InvokeOption) GetDashboardsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDashboardsResult, error) {
			args := v.(GetDashboardsArgs)
			r, err := GetDashboards(ctx, &args, opts...)
			var s GetDashboardsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDashboardsResultOutput)
}

// A collection of arguments for invoking getDashboards.
type GetDashboardsOutputArgs struct {
	// UIDs of Grafana folders containing dashboards. Specify to filter for dashboards by folder (eg. `["General"]` for General folder), or leave blank to get all dashboards in all folders.
	FolderUids pulumi.StringArrayInput `pulumi:"folderUids"`
	// Maximum number of dashboard search results to return. Defaults to `5000`.
	Limit pulumi.IntPtrInput `pulumi:"limit"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
	// List of string Grafana dashboard tags to search for, eg. `["prod"]`. Used only as search input, i.e., attribute value will remain unchanged.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
}

func (GetDashboardsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDashboardsArgs)(nil)).Elem()
}

// A collection of values returned by getDashboards.
type GetDashboardsResultOutput struct{ *pulumi.OutputState }

func (GetDashboardsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDashboardsResult)(nil)).Elem()
}

func (o GetDashboardsResultOutput) ToGetDashboardsResultOutput() GetDashboardsResultOutput {
	return o
}

func (o GetDashboardsResultOutput) ToGetDashboardsResultOutputWithContext(ctx context.Context) GetDashboardsResultOutput {
	return o
}

func (o GetDashboardsResultOutput) Dashboards() GetDashboardsDashboardArrayOutput {
	return o.ApplyT(func(v GetDashboardsResult) []GetDashboardsDashboard { return v.Dashboards }).(GetDashboardsDashboardArrayOutput)
}

// UIDs of Grafana folders containing dashboards. Specify to filter for dashboards by folder (eg. `["General"]` for General folder), or leave blank to get all dashboards in all folders.
func (o GetDashboardsResultOutput) FolderUids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDashboardsResult) []string { return v.FolderUids }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDashboardsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDashboardsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Maximum number of dashboard search results to return. Defaults to `5000`.
func (o GetDashboardsResultOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDashboardsResult) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o GetDashboardsResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDashboardsResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// List of string Grafana dashboard tags to search for, eg. `["prod"]`. Used only as search input, i.e., attribute value will remain unchanged.
func (o GetDashboardsResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDashboardsResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDashboardsResultOutput{})
}
