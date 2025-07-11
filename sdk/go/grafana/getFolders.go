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
//			_, err := oss.NewFolder(ctx, "test_a", &oss.FolderArgs{
//				Title: pulumi.String("test-folder-a"),
//				Uid:   pulumi.String("test-ds-folder-uid-a"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewFolder(ctx, "test_b", &oss.FolderArgs{
//				Title: pulumi.String("test-folder-b"),
//				Uid:   pulumi.String("test-ds-folder-uid-b"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.GetFolders(ctx, &oss.GetFoldersArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: grafana.index/getfolders.getFolders has been deprecated in favor of grafana.oss/getfolders.getFolders
func GetFolders(ctx *pulumi.Context, args *GetFoldersArgs, opts ...pulumi.InvokeOption) (*GetFoldersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetFoldersResult
	err := ctx.Invoke("grafana:index/getFolders:getFolders", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFolders.
type GetFoldersArgs struct {
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
}

// A collection of values returned by getFolders.
type GetFoldersResult struct {
	// The Grafana instance's folders.
	Folders []GetFoldersFolder `pulumi:"folders"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
}

func GetFoldersOutput(ctx *pulumi.Context, args GetFoldersOutputArgs, opts ...pulumi.InvokeOption) GetFoldersResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetFoldersResultOutput, error) {
			args := v.(GetFoldersArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("grafana:index/getFolders:getFolders", args, GetFoldersResultOutput{}, options).(GetFoldersResultOutput), nil
		}).(GetFoldersResultOutput)
}

// A collection of arguments for invoking getFolders.
type GetFoldersOutputArgs struct {
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
}

func (GetFoldersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFoldersArgs)(nil)).Elem()
}

// A collection of values returned by getFolders.
type GetFoldersResultOutput struct{ *pulumi.OutputState }

func (GetFoldersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFoldersResult)(nil)).Elem()
}

func (o GetFoldersResultOutput) ToGetFoldersResultOutput() GetFoldersResultOutput {
	return o
}

func (o GetFoldersResultOutput) ToGetFoldersResultOutputWithContext(ctx context.Context) GetFoldersResultOutput {
	return o
}

// The Grafana instance's folders.
func (o GetFoldersResultOutput) Folders() GetFoldersFolderArrayOutput {
	return o.ApplyT(func(v GetFoldersResult) []GetFoldersFolder { return v.Folders }).(GetFoldersFolderArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFoldersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFoldersResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o GetFoldersResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFoldersResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFoldersResultOutput{})
}
