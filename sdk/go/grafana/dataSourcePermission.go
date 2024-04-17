// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

type DataSourcePermission struct {
	pulumi.CustomResourceState

	// Deprecated: Use `datasource_uid` instead.
	//
	// Deprecated: Use `datasourceUid` instead
	DatasourceId pulumi.StringPtrOutput `pulumi:"datasourceId"`
	// UID of the datasource to apply permissions to.
	DatasourceUid pulumi.StringPtrOutput `pulumi:"datasourceUid"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// The permission items to add/update. Items that are omitted from the list will be removed.
	Permissions DataSourcePermissionPermissionArrayOutput `pulumi:"permissions"`
}

// NewDataSourcePermission registers a new resource with the given unique name, arguments, and options.
func NewDataSourcePermission(ctx *pulumi.Context,
	name string, args *DataSourcePermissionArgs, opts ...pulumi.ResourceOption) (*DataSourcePermission, error) {
	if args == nil {
		args = &DataSourcePermissionArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataSourcePermission
	err := ctx.RegisterResource("grafana:index/dataSourcePermission:DataSourcePermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataSourcePermission gets an existing DataSourcePermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataSourcePermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSourcePermissionState, opts ...pulumi.ResourceOption) (*DataSourcePermission, error) {
	var resource DataSourcePermission
	err := ctx.ReadResource("grafana:index/dataSourcePermission:DataSourcePermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataSourcePermission resources.
type dataSourcePermissionState struct {
	// Deprecated: Use `datasource_uid` instead.
	//
	// Deprecated: Use `datasourceUid` instead
	DatasourceId *string `pulumi:"datasourceId"`
	// UID of the datasource to apply permissions to.
	DatasourceUid *string `pulumi:"datasourceUid"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// The permission items to add/update. Items that are omitted from the list will be removed.
	Permissions []DataSourcePermissionPermission `pulumi:"permissions"`
}

type DataSourcePermissionState struct {
	// Deprecated: Use `datasource_uid` instead.
	//
	// Deprecated: Use `datasourceUid` instead
	DatasourceId pulumi.StringPtrInput
	// UID of the datasource to apply permissions to.
	DatasourceUid pulumi.StringPtrInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// The permission items to add/update. Items that are omitted from the list will be removed.
	Permissions DataSourcePermissionPermissionArrayInput
}

func (DataSourcePermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourcePermissionState)(nil)).Elem()
}

type dataSourcePermissionArgs struct {
	// Deprecated: Use `datasource_uid` instead.
	//
	// Deprecated: Use `datasourceUid` instead
	DatasourceId *string `pulumi:"datasourceId"`
	// UID of the datasource to apply permissions to.
	DatasourceUid *string `pulumi:"datasourceUid"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// The permission items to add/update. Items that are omitted from the list will be removed.
	Permissions []DataSourcePermissionPermission `pulumi:"permissions"`
}

// The set of arguments for constructing a DataSourcePermission resource.
type DataSourcePermissionArgs struct {
	// Deprecated: Use `datasource_uid` instead.
	//
	// Deprecated: Use `datasourceUid` instead
	DatasourceId pulumi.StringPtrInput
	// UID of the datasource to apply permissions to.
	DatasourceUid pulumi.StringPtrInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// The permission items to add/update. Items that are omitted from the list will be removed.
	Permissions DataSourcePermissionPermissionArrayInput
}

func (DataSourcePermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourcePermissionArgs)(nil)).Elem()
}

type DataSourcePermissionInput interface {
	pulumi.Input

	ToDataSourcePermissionOutput() DataSourcePermissionOutput
	ToDataSourcePermissionOutputWithContext(ctx context.Context) DataSourcePermissionOutput
}

func (*DataSourcePermission) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSourcePermission)(nil)).Elem()
}

func (i *DataSourcePermission) ToDataSourcePermissionOutput() DataSourcePermissionOutput {
	return i.ToDataSourcePermissionOutputWithContext(context.Background())
}

func (i *DataSourcePermission) ToDataSourcePermissionOutputWithContext(ctx context.Context) DataSourcePermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourcePermissionOutput)
}

// DataSourcePermissionArrayInput is an input type that accepts DataSourcePermissionArray and DataSourcePermissionArrayOutput values.
// You can construct a concrete instance of `DataSourcePermissionArrayInput` via:
//
//	DataSourcePermissionArray{ DataSourcePermissionArgs{...} }
type DataSourcePermissionArrayInput interface {
	pulumi.Input

	ToDataSourcePermissionArrayOutput() DataSourcePermissionArrayOutput
	ToDataSourcePermissionArrayOutputWithContext(context.Context) DataSourcePermissionArrayOutput
}

type DataSourcePermissionArray []DataSourcePermissionInput

func (DataSourcePermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataSourcePermission)(nil)).Elem()
}

func (i DataSourcePermissionArray) ToDataSourcePermissionArrayOutput() DataSourcePermissionArrayOutput {
	return i.ToDataSourcePermissionArrayOutputWithContext(context.Background())
}

func (i DataSourcePermissionArray) ToDataSourcePermissionArrayOutputWithContext(ctx context.Context) DataSourcePermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourcePermissionArrayOutput)
}

// DataSourcePermissionMapInput is an input type that accepts DataSourcePermissionMap and DataSourcePermissionMapOutput values.
// You can construct a concrete instance of `DataSourcePermissionMapInput` via:
//
//	DataSourcePermissionMap{ "key": DataSourcePermissionArgs{...} }
type DataSourcePermissionMapInput interface {
	pulumi.Input

	ToDataSourcePermissionMapOutput() DataSourcePermissionMapOutput
	ToDataSourcePermissionMapOutputWithContext(context.Context) DataSourcePermissionMapOutput
}

type DataSourcePermissionMap map[string]DataSourcePermissionInput

func (DataSourcePermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataSourcePermission)(nil)).Elem()
}

func (i DataSourcePermissionMap) ToDataSourcePermissionMapOutput() DataSourcePermissionMapOutput {
	return i.ToDataSourcePermissionMapOutputWithContext(context.Background())
}

func (i DataSourcePermissionMap) ToDataSourcePermissionMapOutputWithContext(ctx context.Context) DataSourcePermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourcePermissionMapOutput)
}

type DataSourcePermissionOutput struct{ *pulumi.OutputState }

func (DataSourcePermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSourcePermission)(nil)).Elem()
}

func (o DataSourcePermissionOutput) ToDataSourcePermissionOutput() DataSourcePermissionOutput {
	return o
}

func (o DataSourcePermissionOutput) ToDataSourcePermissionOutputWithContext(ctx context.Context) DataSourcePermissionOutput {
	return o
}

// Deprecated: Use `datasource_uid` instead.
//
// Deprecated: Use `datasourceUid` instead
func (o DataSourcePermissionOutput) DatasourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSourcePermission) pulumi.StringPtrOutput { return v.DatasourceId }).(pulumi.StringPtrOutput)
}

// UID of the datasource to apply permissions to.
func (o DataSourcePermissionOutput) DatasourceUid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSourcePermission) pulumi.StringPtrOutput { return v.DatasourceUid }).(pulumi.StringPtrOutput)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o DataSourcePermissionOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSourcePermission) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// The permission items to add/update. Items that are omitted from the list will be removed.
func (o DataSourcePermissionOutput) Permissions() DataSourcePermissionPermissionArrayOutput {
	return o.ApplyT(func(v *DataSourcePermission) DataSourcePermissionPermissionArrayOutput { return v.Permissions }).(DataSourcePermissionPermissionArrayOutput)
}

type DataSourcePermissionArrayOutput struct{ *pulumi.OutputState }

func (DataSourcePermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataSourcePermission)(nil)).Elem()
}

func (o DataSourcePermissionArrayOutput) ToDataSourcePermissionArrayOutput() DataSourcePermissionArrayOutput {
	return o
}

func (o DataSourcePermissionArrayOutput) ToDataSourcePermissionArrayOutputWithContext(ctx context.Context) DataSourcePermissionArrayOutput {
	return o
}

func (o DataSourcePermissionArrayOutput) Index(i pulumi.IntInput) DataSourcePermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataSourcePermission {
		return vs[0].([]*DataSourcePermission)[vs[1].(int)]
	}).(DataSourcePermissionOutput)
}

type DataSourcePermissionMapOutput struct{ *pulumi.OutputState }

func (DataSourcePermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataSourcePermission)(nil)).Elem()
}

func (o DataSourcePermissionMapOutput) ToDataSourcePermissionMapOutput() DataSourcePermissionMapOutput {
	return o
}

func (o DataSourcePermissionMapOutput) ToDataSourcePermissionMapOutputWithContext(ctx context.Context) DataSourcePermissionMapOutput {
	return o
}

func (o DataSourcePermissionMapOutput) MapIndex(k pulumi.StringInput) DataSourcePermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataSourcePermission {
		return vs[0].(map[string]*DataSourcePermission)[vs[1].(string)]
	}).(DataSourcePermissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourcePermissionInput)(nil)).Elem(), &DataSourcePermission{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourcePermissionArrayInput)(nil)).Elem(), DataSourcePermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourcePermissionMapInput)(nil)).Elem(), DataSourcePermissionMap{})
	pulumi.RegisterOutputType(DataSourcePermissionOutput{})
	pulumi.RegisterOutputType(DataSourcePermissionArrayOutput{})
	pulumi.RegisterOutputType(DataSourcePermissionMapOutput{})
}
