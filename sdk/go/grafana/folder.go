// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"errors"
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
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testFolderFolder, err := grafana.NewFolder(ctx, "testFolderFolder", &grafana.FolderArgs{
//				Title: pulumi.String("Terraform Test Folder"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewDashboard(ctx, "testFolderDashboard", &grafana.DashboardArgs{
//				Folder:     testFolderFolder.ID(),
//				ConfigJson: pulumi.String("{\n  \"title\": \"Dashboard in folder\",\n  \"uid\": \"dashboard-in-folder\"\n}\n"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewFolder(ctx, "testFolderWithUid", &grafana.FolderArgs{
//				Uid:   pulumi.String("test-folder-uid"),
//				Title: pulumi.String("Terraform Test Folder With UID"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
// $ pulumi import grafana:index/folder:Folder name "{{ uid }}"
// ```
//
// ```sh
// $ pulumi import grafana:index/folder:Folder name "{{ orgID }}:{{ uid }}"
// ```
type Folder struct {
	pulumi.CustomResourceState

	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// The uid of the parent folder. If set, the folder will be nested. If not set, the folder will be created in the root folder. Note: This requires the nestedFolders feature flag to be enabled on your Grafana instance.
	ParentFolderUid pulumi.StringPtrOutput `pulumi:"parentFolderUid"`
	// Prevent deletion of the folder if it is not empty (contains dashboards or alert rules). Defaults to `false`.
	PreventDestroyIfNotEmpty pulumi.BoolPtrOutput `pulumi:"preventDestroyIfNotEmpty"`
	// The title of the folder.
	Title pulumi.StringOutput `pulumi:"title"`
	// Unique identifier.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The full URL of the folder.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewFolder registers a new resource with the given unique name, arguments, and options.
func NewFolder(ctx *pulumi.Context,
	name string, args *FolderArgs, opts ...pulumi.ResourceOption) (*Folder, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Folder
	err := ctx.RegisterResource("grafana:index/folder:Folder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolder gets an existing Folder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderState, opts ...pulumi.ResourceOption) (*Folder, error) {
	var resource Folder
	err := ctx.ReadResource("grafana:index/folder:Folder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Folder resources.
type folderState struct {
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// The uid of the parent folder. If set, the folder will be nested. If not set, the folder will be created in the root folder. Note: This requires the nestedFolders feature flag to be enabled on your Grafana instance.
	ParentFolderUid *string `pulumi:"parentFolderUid"`
	// Prevent deletion of the folder if it is not empty (contains dashboards or alert rules). Defaults to `false`.
	PreventDestroyIfNotEmpty *bool `pulumi:"preventDestroyIfNotEmpty"`
	// The title of the folder.
	Title *string `pulumi:"title"`
	// Unique identifier.
	Uid *string `pulumi:"uid"`
	// The full URL of the folder.
	Url *string `pulumi:"url"`
}

type FolderState struct {
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// The uid of the parent folder. If set, the folder will be nested. If not set, the folder will be created in the root folder. Note: This requires the nestedFolders feature flag to be enabled on your Grafana instance.
	ParentFolderUid pulumi.StringPtrInput
	// Prevent deletion of the folder if it is not empty (contains dashboards or alert rules). Defaults to `false`.
	PreventDestroyIfNotEmpty pulumi.BoolPtrInput
	// The title of the folder.
	Title pulumi.StringPtrInput
	// Unique identifier.
	Uid pulumi.StringPtrInput
	// The full URL of the folder.
	Url pulumi.StringPtrInput
}

func (FolderState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderState)(nil)).Elem()
}

type folderArgs struct {
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// The uid of the parent folder. If set, the folder will be nested. If not set, the folder will be created in the root folder. Note: This requires the nestedFolders feature flag to be enabled on your Grafana instance.
	ParentFolderUid *string `pulumi:"parentFolderUid"`
	// Prevent deletion of the folder if it is not empty (contains dashboards or alert rules). Defaults to `false`.
	PreventDestroyIfNotEmpty *bool `pulumi:"preventDestroyIfNotEmpty"`
	// The title of the folder.
	Title string `pulumi:"title"`
	// Unique identifier.
	Uid *string `pulumi:"uid"`
}

// The set of arguments for constructing a Folder resource.
type FolderArgs struct {
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// The uid of the parent folder. If set, the folder will be nested. If not set, the folder will be created in the root folder. Note: This requires the nestedFolders feature flag to be enabled on your Grafana instance.
	ParentFolderUid pulumi.StringPtrInput
	// Prevent deletion of the folder if it is not empty (contains dashboards or alert rules). Defaults to `false`.
	PreventDestroyIfNotEmpty pulumi.BoolPtrInput
	// The title of the folder.
	Title pulumi.StringInput
	// Unique identifier.
	Uid pulumi.StringPtrInput
}

func (FolderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderArgs)(nil)).Elem()
}

type FolderInput interface {
	pulumi.Input

	ToFolderOutput() FolderOutput
	ToFolderOutputWithContext(ctx context.Context) FolderOutput
}

func (*Folder) ElementType() reflect.Type {
	return reflect.TypeOf((**Folder)(nil)).Elem()
}

func (i *Folder) ToFolderOutput() FolderOutput {
	return i.ToFolderOutputWithContext(context.Background())
}

func (i *Folder) ToFolderOutputWithContext(ctx context.Context) FolderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderOutput)
}

// FolderArrayInput is an input type that accepts FolderArray and FolderArrayOutput values.
// You can construct a concrete instance of `FolderArrayInput` via:
//
//	FolderArray{ FolderArgs{...} }
type FolderArrayInput interface {
	pulumi.Input

	ToFolderArrayOutput() FolderArrayOutput
	ToFolderArrayOutputWithContext(context.Context) FolderArrayOutput
}

type FolderArray []FolderInput

func (FolderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Folder)(nil)).Elem()
}

func (i FolderArray) ToFolderArrayOutput() FolderArrayOutput {
	return i.ToFolderArrayOutputWithContext(context.Background())
}

func (i FolderArray) ToFolderArrayOutputWithContext(ctx context.Context) FolderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderArrayOutput)
}

// FolderMapInput is an input type that accepts FolderMap and FolderMapOutput values.
// You can construct a concrete instance of `FolderMapInput` via:
//
//	FolderMap{ "key": FolderArgs{...} }
type FolderMapInput interface {
	pulumi.Input

	ToFolderMapOutput() FolderMapOutput
	ToFolderMapOutputWithContext(context.Context) FolderMapOutput
}

type FolderMap map[string]FolderInput

func (FolderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Folder)(nil)).Elem()
}

func (i FolderMap) ToFolderMapOutput() FolderMapOutput {
	return i.ToFolderMapOutputWithContext(context.Background())
}

func (i FolderMap) ToFolderMapOutputWithContext(ctx context.Context) FolderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderMapOutput)
}

type FolderOutput struct{ *pulumi.OutputState }

func (FolderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Folder)(nil)).Elem()
}

func (o FolderOutput) ToFolderOutput() FolderOutput {
	return o
}

func (o FolderOutput) ToFolderOutputWithContext(ctx context.Context) FolderOutput {
	return o
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o FolderOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// The uid of the parent folder. If set, the folder will be nested. If not set, the folder will be created in the root folder. Note: This requires the nestedFolders feature flag to be enabled on your Grafana instance.
func (o FolderOutput) ParentFolderUid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringPtrOutput { return v.ParentFolderUid }).(pulumi.StringPtrOutput)
}

// Prevent deletion of the folder if it is not empty (contains dashboards or alert rules). Defaults to `false`.
func (o FolderOutput) PreventDestroyIfNotEmpty() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Folder) pulumi.BoolPtrOutput { return v.PreventDestroyIfNotEmpty }).(pulumi.BoolPtrOutput)
}

// The title of the folder.
func (o FolderOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// Unique identifier.
func (o FolderOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The full URL of the folder.
func (o FolderOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type FolderArrayOutput struct{ *pulumi.OutputState }

func (FolderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Folder)(nil)).Elem()
}

func (o FolderArrayOutput) ToFolderArrayOutput() FolderArrayOutput {
	return o
}

func (o FolderArrayOutput) ToFolderArrayOutputWithContext(ctx context.Context) FolderArrayOutput {
	return o
}

func (o FolderArrayOutput) Index(i pulumi.IntInput) FolderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Folder {
		return vs[0].([]*Folder)[vs[1].(int)]
	}).(FolderOutput)
}

type FolderMapOutput struct{ *pulumi.OutputState }

func (FolderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Folder)(nil)).Elem()
}

func (o FolderMapOutput) ToFolderMapOutput() FolderMapOutput {
	return o
}

func (o FolderMapOutput) ToFolderMapOutputWithContext(ctx context.Context) FolderMapOutput {
	return o
}

func (o FolderMapOutput) MapIndex(k pulumi.StringInput) FolderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Folder {
		return vs[0].(map[string]*Folder)[vs[1].(string)]
	}).(FolderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FolderInput)(nil)).Elem(), &Folder{})
	pulumi.RegisterInputType(reflect.TypeOf((*FolderArrayInput)(nil)).Elem(), FolderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FolderMapInput)(nil)).Elem(), FolderMap{})
	pulumi.RegisterOutputType(FolderOutput{})
	pulumi.RegisterOutputType(FolderArrayOutput{})
	pulumi.RegisterOutputType(FolderMapOutput{})
}
