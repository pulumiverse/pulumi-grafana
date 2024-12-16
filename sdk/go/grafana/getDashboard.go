// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/)
// * [Folder/Dashboard Search HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/folder_dashboard_search/)
// * [Dashboard HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/dashboard/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"id":    12345,
//				"uid":   "test-ds-dashboard-uid",
//				"title": "Production Overview",
//				"tags": []string{
//					"templated",
//				},
//				"timezone":      "browser",
//				"schemaVersion": 16,
//				"version":       0,
//				"refresh":       "25s",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = oss.NewDashboard(ctx, "test", &oss.DashboardArgs{
//				ConfigJson: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.LookupDashboard(ctx, &oss.LookupDashboardArgs{
//				Uid: pulumi.StringRef("test-ds-dashboard-uid"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: grafana.index/getdashboard.getDashboard has been deprecated in favor of grafana.oss/getdashboard.getDashboard
func LookupDashboard(ctx *pulumi.Context, args *LookupDashboardArgs, opts ...pulumi.InvokeOption) (*LookupDashboardResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDashboardResult
	err := ctx.Invoke("grafana:index/getDashboard:getDashboard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDashboard.
type LookupDashboardArgs struct {
	// The numerical ID of the Grafana dashboard. Specify either this or `uid`. Defaults to `-1`.
	DashboardId *int `pulumi:"dashboardId"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// The uid of the Grafana dashboard. Specify either this or `dashboardId`. Defaults to ``.
	Uid *string `pulumi:"uid"`
}

// A collection of values returned by getDashboard.
type LookupDashboardResult struct {
	// The complete dashboard model JSON.
	ConfigJson string `pulumi:"configJson"`
	// The numerical ID of the Grafana dashboard. Specify either this or `uid`. Defaults to `-1`.
	DashboardId *int `pulumi:"dashboardId"`
	// The UID of the folder where the Grafana dashboard is found.
	FolderUid string `pulumi:"folderUid"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Whether or not the Grafana dashboard is starred. Starred Dashboards will show up on your own Home Dashboard by default, and are a convenient way to mark Dashboards that you’re interested in.
	IsStarred bool `pulumi:"isStarred"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// URL slug of the dashboard (deprecated).
	Slug string `pulumi:"slug"`
	// The title of the Grafana dashboard.
	Title string `pulumi:"title"`
	// The uid of the Grafana dashboard. Specify either this or `dashboardId`. Defaults to ``.
	Uid *string `pulumi:"uid"`
	// The full URL of the dashboard.
	Url string `pulumi:"url"`
	// The numerical version of the Grafana dashboard.
	Version int `pulumi:"version"`
}

func LookupDashboardOutput(ctx *pulumi.Context, args LookupDashboardOutputArgs, opts ...pulumi.InvokeOption) LookupDashboardResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDashboardResultOutput, error) {
			args := v.(LookupDashboardArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("grafana:index/getDashboard:getDashboard", args, LookupDashboardResultOutput{}, options).(LookupDashboardResultOutput), nil
		}).(LookupDashboardResultOutput)
}

// A collection of arguments for invoking getDashboard.
type LookupDashboardOutputArgs struct {
	// The numerical ID of the Grafana dashboard. Specify either this or `uid`. Defaults to `-1`.
	DashboardId pulumi.IntPtrInput `pulumi:"dashboardId"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
	// The uid of the Grafana dashboard. Specify either this or `dashboardId`. Defaults to ``.
	Uid pulumi.StringPtrInput `pulumi:"uid"`
}

func (LookupDashboardOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDashboardArgs)(nil)).Elem()
}

// A collection of values returned by getDashboard.
type LookupDashboardResultOutput struct{ *pulumi.OutputState }

func (LookupDashboardResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDashboardResult)(nil)).Elem()
}

func (o LookupDashboardResultOutput) ToLookupDashboardResultOutput() LookupDashboardResultOutput {
	return o
}

func (o LookupDashboardResultOutput) ToLookupDashboardResultOutputWithContext(ctx context.Context) LookupDashboardResultOutput {
	return o
}

// The complete dashboard model JSON.
func (o LookupDashboardResultOutput) ConfigJson() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDashboardResult) string { return v.ConfigJson }).(pulumi.StringOutput)
}

// The numerical ID of the Grafana dashboard. Specify either this or `uid`. Defaults to `-1`.
func (o LookupDashboardResultOutput) DashboardId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDashboardResult) *int { return v.DashboardId }).(pulumi.IntPtrOutput)
}

// The UID of the folder where the Grafana dashboard is found.
func (o LookupDashboardResultOutput) FolderUid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDashboardResult) string { return v.FolderUid }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDashboardResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDashboardResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether or not the Grafana dashboard is starred. Starred Dashboards will show up on your own Home Dashboard by default, and are a convenient way to mark Dashboards that you’re interested in.
func (o LookupDashboardResultOutput) IsStarred() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDashboardResult) bool { return v.IsStarred }).(pulumi.BoolOutput)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o LookupDashboardResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDashboardResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// URL slug of the dashboard (deprecated).
func (o LookupDashboardResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDashboardResult) string { return v.Slug }).(pulumi.StringOutput)
}

// The title of the Grafana dashboard.
func (o LookupDashboardResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDashboardResult) string { return v.Title }).(pulumi.StringOutput)
}

// The uid of the Grafana dashboard. Specify either this or `dashboardId`. Defaults to “.
func (o LookupDashboardResultOutput) Uid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDashboardResult) *string { return v.Uid }).(pulumi.StringPtrOutput)
}

// The full URL of the dashboard.
func (o LookupDashboardResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDashboardResult) string { return v.Url }).(pulumi.StringOutput)
}

// The numerical version of the Grafana dashboard.
func (o LookupDashboardResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDashboardResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDashboardResultOutput{})
}
