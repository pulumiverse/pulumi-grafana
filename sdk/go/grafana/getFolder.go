// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
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
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := grafana.NewFolder(ctx, "test", &grafana.FolderArgs{
//				Title: pulumi.String("test-folder"),
//				Uid:   pulumi.String("test-ds-folder-uid"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = grafana.LookupFolderOutput(ctx, grafana.GetFolderOutputArgs{
//				Title: test.Title,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
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
	// The name of the Grafana folder.
	Title string `pulumi:"title"`
}

// A collection of values returned by getFolder.
type LookupFolderResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId           *string `pulumi:"orgId"`
	ParentFolderUid string  `pulumi:"parentFolderUid"`
	// The name of the Grafana folder.
	Title string `pulumi:"title"`
	// The uid of the Grafana folder.
	Uid string `pulumi:"uid"`
	// The full URL of the folder.
	Url string `pulumi:"url"`
}

func LookupFolderOutput(ctx *pulumi.Context, args LookupFolderOutputArgs, opts ...pulumi.InvokeOption) LookupFolderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFolderResult, error) {
			args := v.(LookupFolderArgs)
			r, err := LookupFolder(ctx, &args, opts...)
			var s LookupFolderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFolderResultOutput)
}

// A collection of arguments for invoking getFolder.
type LookupFolderOutputArgs struct {
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
	// The name of the Grafana folder.
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

func (o LookupFolderResultOutput) ParentFolderUid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.ParentFolderUid }).(pulumi.StringOutput)
}

// The name of the Grafana folder.
func (o LookupFolderResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Title }).(pulumi.StringOutput)
}

// The uid of the Grafana folder.
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
