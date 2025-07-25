// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/manage-dashboards/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/folder/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := oss.NewFolder(ctx, "test", &oss.FolderArgs{
//				Title: pulumi.String("test-folder"),
//				Uid:   pulumi.String("test-ds-folder-uid"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = oss.LookupFolderOutput(ctx, oss.GetFolderOutputArgs{
//				Title: test.Title,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: grafana.index/getfolder.getFolder has been deprecated in favor of grafana.oss/getfolder.getFolder
func LookupFolder(ctx *pulumi.Context, args *LookupFolderArgs, opts ...pulumi.InvokeOption) (*LookupFolderResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFolderResult
	err := ctx.Invoke("grafana:index/getFolder:getFolder", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFolder.
type LookupFolderArgs struct {
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// The title of the folder.
	Title string `pulumi:"title"`
}

// A collection of values returned by getFolder.
type LookupFolderResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// The uid of the parent folder. If set, the folder will be nested. If not set, the folder will be created in the root folder. Note: This requires the nestedFolders feature flag to be enabled on your Grafana instance.
	ParentFolderUid string `pulumi:"parentFolderUid"`
	// The title of the folder.
	Title string `pulumi:"title"`
	// Unique identifier.
	Uid string `pulumi:"uid"`
	// The full URL of the folder.
	Url string `pulumi:"url"`
}

func LookupFolderOutput(ctx *pulumi.Context, args LookupFolderOutputArgs, opts ...pulumi.InvokeOption) LookupFolderResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFolderResultOutput, error) {
			args := v.(LookupFolderArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("grafana:index/getFolder:getFolder", args, LookupFolderResultOutput{}, options).(LookupFolderResultOutput), nil
		}).(LookupFolderResultOutput)
}

// A collection of arguments for invoking getFolder.
type LookupFolderOutputArgs struct {
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
	// The title of the folder.
	Title pulumi.StringInput `pulumi:"title"`
}

func (LookupFolderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderArgs)(nil)).Elem()
}

// A collection of values returned by getFolder.
type LookupFolderResultOutput struct{ *pulumi.OutputState }

func (LookupFolderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderResult)(nil)).Elem()
}

func (o LookupFolderResultOutput) ToLookupFolderResultOutput() LookupFolderResultOutput {
	return o
}

func (o LookupFolderResultOutput) ToLookupFolderResultOutputWithContext(ctx context.Context) LookupFolderResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFolderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o LookupFolderResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFolderResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// The uid of the parent folder. If set, the folder will be nested. If not set, the folder will be created in the root folder. Note: This requires the nestedFolders feature flag to be enabled on your Grafana instance.
func (o LookupFolderResultOutput) ParentFolderUid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.ParentFolderUid }).(pulumi.StringOutput)
}

// The title of the folder.
func (o LookupFolderResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Title }).(pulumi.StringOutput)
}

// Unique identifier.
func (o LookupFolderResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Uid }).(pulumi.StringOutput)
}

// The full URL of the folder.
func (o LookupFolderResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFolderResultOutput{})
}
