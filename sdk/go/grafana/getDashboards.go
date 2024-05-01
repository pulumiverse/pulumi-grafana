// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

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
	// Deprecated: Use `folderUids` instead.
	FolderIds  []int    `pulumi:"folderIds"`
	FolderUids []string `pulumi:"folderUids"`
	Limit      *int     `pulumi:"limit"`
	OrgId      *string  `pulumi:"orgId"`
	Tags       []string `pulumi:"tags"`
}

// A collection of values returned by getDashboards.
type GetDashboardsResult struct {
	Dashboards []GetDashboardsDashboard `pulumi:"dashboards"`
	// Deprecated: Use `folderUids` instead.
	FolderIds  []int    `pulumi:"folderIds"`
	FolderUids []string `pulumi:"folderUids"`
	// The provider-assigned unique ID for this managed resource.
	Id    string   `pulumi:"id"`
	Limit *int     `pulumi:"limit"`
	OrgId *string  `pulumi:"orgId"`
	Tags  []string `pulumi:"tags"`
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
	// Deprecated: Use `folderUids` instead.
	FolderIds  pulumi.IntArrayInput    `pulumi:"folderIds"`
	FolderUids pulumi.StringArrayInput `pulumi:"folderUids"`
	Limit      pulumi.IntPtrInput      `pulumi:"limit"`
	OrgId      pulumi.StringPtrInput   `pulumi:"orgId"`
	Tags       pulumi.StringArrayInput `pulumi:"tags"`
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

// Deprecated: Use `folderUids` instead.
func (o GetDashboardsResultOutput) FolderIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetDashboardsResult) []int { return v.FolderIds }).(pulumi.IntArrayOutput)
}

func (o GetDashboardsResultOutput) FolderUids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDashboardsResult) []string { return v.FolderUids }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDashboardsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDashboardsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDashboardsResultOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDashboardsResult) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

func (o GetDashboardsResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDashboardsResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

func (o GetDashboardsResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDashboardsResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDashboardsResultOutput{})
}
