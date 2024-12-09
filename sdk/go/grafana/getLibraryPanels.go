// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

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
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"title":       "test name",
//				"type":        "text",
//				"version":     0,
//				"description": "test description",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = oss.NewLibraryPanel(ctx, "test", &oss.LibraryPanelArgs{
//				Name:      pulumi.String("panelname"),
//				ModelJson: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			testFolder, err := oss.NewFolder(ctx, "test", &oss.FolderArgs{
//				Title: pulumi.String("Panel Folder"),
//				Uid:   pulumi.String("panelname-folder"),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON1, err := json.Marshal(map[string]interface{}{
//				"gridPos": map[string]interface{}{
//					"x": 0,
//					"y": 0,
//					"h": 10,
//					"w": 10,
//				},
//				"title":   "panel",
//				"type":    "text",
//				"version": 0,
//			})
//			if err != nil {
//				return err
//			}
//			json1 := string(tmpJSON1)
//			_, err = oss.NewLibraryPanel(ctx, "folder", &oss.LibraryPanelArgs{
//				Name:      pulumi.String("panelname In Folder"),
//				FolderUid: testFolder.Uid,
//				ModelJson: pulumi.String(json1),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = grafana.GetLibraryPanels(ctx, &grafana.GetLibraryPanelsArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetLibraryPanels(ctx *pulumi.Context, args *GetLibraryPanelsArgs, opts ...pulumi.InvokeOption) (*GetLibraryPanelsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLibraryPanelsResult
	err := ctx.Invoke("grafana:index/getLibraryPanels:getLibraryPanels", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLibraryPanels.
type GetLibraryPanelsArgs struct {
	// The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
	OrgId *string `pulumi:"orgId"`
}

// A collection of values returned by getLibraryPanels.
type GetLibraryPanelsResult struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
	// The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
	OrgId  string                  `pulumi:"orgId"`
	Panels []GetLibraryPanelsPanel `pulumi:"panels"`
}

func GetLibraryPanelsOutput(ctx *pulumi.Context, args GetLibraryPanelsOutputArgs, opts ...pulumi.InvokeOption) GetLibraryPanelsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLibraryPanelsResultOutput, error) {
			args := v.(GetLibraryPanelsArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetLibraryPanelsResult
			secret, err := ctx.InvokePackageRaw("grafana:index/getLibraryPanels:getLibraryPanels", args, &rv, "", opts...)
			if err != nil {
				return GetLibraryPanelsResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetLibraryPanelsResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetLibraryPanelsResultOutput), nil
			}
			return output, nil
		}).(GetLibraryPanelsResultOutput)
}

// A collection of arguments for invoking getLibraryPanels.
type GetLibraryPanelsOutputArgs struct {
	// The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
}

func (GetLibraryPanelsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLibraryPanelsArgs)(nil)).Elem()
}

// A collection of values returned by getLibraryPanels.
type GetLibraryPanelsResultOutput struct{ *pulumi.OutputState }

func (GetLibraryPanelsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLibraryPanelsResult)(nil)).Elem()
}

func (o GetLibraryPanelsResultOutput) ToGetLibraryPanelsResultOutput() GetLibraryPanelsResultOutput {
	return o
}

func (o GetLibraryPanelsResultOutput) ToGetLibraryPanelsResultOutputWithContext(ctx context.Context) GetLibraryPanelsResultOutput {
	return o
}

// The ID of this resource.
func (o GetLibraryPanelsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLibraryPanelsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
func (o GetLibraryPanelsResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLibraryPanelsResult) string { return v.OrgId }).(pulumi.StringOutput)
}

func (o GetLibraryPanelsResultOutput) Panels() GetLibraryPanelsPanelArrayOutput {
	return o.ApplyT(func(v GetLibraryPanelsResult) []GetLibraryPanelsPanel { return v.Panels }).(GetLibraryPanelsPanelArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLibraryPanelsResultOutput{})
}