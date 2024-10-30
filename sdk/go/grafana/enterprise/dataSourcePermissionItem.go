// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package enterprise

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// Manages a single permission item for a datasource. Conflicts with the "enterprise.DataSourcePermission" resource which manages the entire set of permissions for a datasource.
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
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/enterprise"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			team, err := oss.NewTeam(ctx, "team", &oss.TeamArgs{
//				Name: pulumi.String("Team Name"),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"defaultRegion": "us-east-1",
//				"authType":      "keys",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			tmpJSON1, err := json.Marshal(map[string]interface{}{
//				"accessKey": "123",
//				"secretKey": "456",
//			})
//			if err != nil {
//				return err
//			}
//			json1 := string(tmpJSON1)
//			foo, err := oss.NewDataSource(ctx, "foo", &oss.DataSourceArgs{
//				Type:                  pulumi.String("cloudwatch"),
//				Name:                  pulumi.String("cw-example"),
//				JsonDataEncoded:       pulumi.String(json0),
//				SecureJsonDataEncoded: pulumi.String(json1),
//			})
//			if err != nil {
//				return err
//			}
//			user, err := oss.NewUser(ctx, "user", &oss.UserArgs{
//				Name:     pulumi.String("test-ds-permissions"),
//				Email:    pulumi.String("test-ds-permissions@example.com"),
//				Login:    pulumi.String("test-ds-permissions"),
//				Password: pulumi.String("hunter2"),
//			})
//			if err != nil {
//				return err
//			}
//			sa, err := oss.NewServiceAccount(ctx, "sa", &oss.ServiceAccountArgs{
//				Name: pulumi.String("test-ds-permissions"),
//				Role: pulumi.String("Viewer"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = enterprise.NewDataSourcePermissionItem(ctx, "team", &enterprise.DataSourcePermissionItemArgs{
//				DatasourceUid: foo.Uid,
//				Team:          team.ID(),
//				Permission:    pulumi.String("Edit"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = enterprise.NewDataSourcePermissionItem(ctx, "user", &enterprise.DataSourcePermissionItemArgs{
//				DatasourceUid: foo.Uid,
//				User:          user.ID(),
//				Permission:    pulumi.String("Edit"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = enterprise.NewDataSourcePermissionItem(ctx, "role", &enterprise.DataSourcePermissionItemArgs{
//				DatasourceUid: foo.Uid,
//				Role:          pulumi.String("Viewer"),
//				Permission:    pulumi.String("Query"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = enterprise.NewDataSourcePermissionItem(ctx, "service_account", &enterprise.DataSourcePermissionItemArgs{
//				DatasourceUid: foo.Uid,
//				User:          sa.ID(),
//				Permission:    pulumi.String("Query"),
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
// $ pulumi import grafana:enterprise/dataSourcePermissionItem:DataSourcePermissionItem name "{{ datasourceUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
// ```
//
// ```sh
// $ pulumi import grafana:enterprise/dataSourcePermissionItem:DataSourcePermissionItem name "{{ orgID }}:{{ datasourceUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
// ```
type DataSourcePermissionItem struct {
	pulumi.CustomResourceState

	// The UID of the datasource.
	DatasourceUid pulumi.StringOutput `pulumi:"datasourceUid"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// the permission to be assigned
	Permission pulumi.StringOutput `pulumi:"permission"`
	// the role onto which the permission is to be assigned
	Role pulumi.StringPtrOutput `pulumi:"role"`
	// the team onto which the permission is to be assigned
	Team pulumi.StringPtrOutput `pulumi:"team"`
	// the user or service account onto which the permission is to be assigned
	User pulumi.StringPtrOutput `pulumi:"user"`
}

// NewDataSourcePermissionItem registers a new resource with the given unique name, arguments, and options.
func NewDataSourcePermissionItem(ctx *pulumi.Context,
	name string, args *DataSourcePermissionItemArgs, opts ...pulumi.ResourceOption) (*DataSourcePermissionItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasourceUid == nil {
		return nil, errors.New("invalid value for required argument 'DatasourceUid'")
	}
	if args.Permission == nil {
		return nil, errors.New("invalid value for required argument 'Permission'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/dataSourcePermissionItem:DataSourcePermissionItem"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataSourcePermissionItem
	err := ctx.RegisterResource("grafana:enterprise/dataSourcePermissionItem:DataSourcePermissionItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataSourcePermissionItem gets an existing DataSourcePermissionItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataSourcePermissionItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSourcePermissionItemState, opts ...pulumi.ResourceOption) (*DataSourcePermissionItem, error) {
	var resource DataSourcePermissionItem
	err := ctx.ReadResource("grafana:enterprise/dataSourcePermissionItem:DataSourcePermissionItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataSourcePermissionItem resources.
type dataSourcePermissionItemState struct {
	// The UID of the datasource.
	DatasourceUid *string `pulumi:"datasourceUid"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// the permission to be assigned
	Permission *string `pulumi:"permission"`
	// the role onto which the permission is to be assigned
	Role *string `pulumi:"role"`
	// the team onto which the permission is to be assigned
	Team *string `pulumi:"team"`
	// the user or service account onto which the permission is to be assigned
	User *string `pulumi:"user"`
}

type DataSourcePermissionItemState struct {
	// The UID of the datasource.
	DatasourceUid pulumi.StringPtrInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// the permission to be assigned
	Permission pulumi.StringPtrInput
	// the role onto which the permission is to be assigned
	Role pulumi.StringPtrInput
	// the team onto which the permission is to be assigned
	Team pulumi.StringPtrInput
	// the user or service account onto which the permission is to be assigned
	User pulumi.StringPtrInput
}

func (DataSourcePermissionItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourcePermissionItemState)(nil)).Elem()
}

type dataSourcePermissionItemArgs struct {
	// The UID of the datasource.
	DatasourceUid string `pulumi:"datasourceUid"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// the permission to be assigned
	Permission string `pulumi:"permission"`
	// the role onto which the permission is to be assigned
	Role *string `pulumi:"role"`
	// the team onto which the permission is to be assigned
	Team *string `pulumi:"team"`
	// the user or service account onto which the permission is to be assigned
	User *string `pulumi:"user"`
}

// The set of arguments for constructing a DataSourcePermissionItem resource.
type DataSourcePermissionItemArgs struct {
	// The UID of the datasource.
	DatasourceUid pulumi.StringInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// the permission to be assigned
	Permission pulumi.StringInput
	// the role onto which the permission is to be assigned
	Role pulumi.StringPtrInput
	// the team onto which the permission is to be assigned
	Team pulumi.StringPtrInput
	// the user or service account onto which the permission is to be assigned
	User pulumi.StringPtrInput
}

func (DataSourcePermissionItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourcePermissionItemArgs)(nil)).Elem()
}

type DataSourcePermissionItemInput interface {
	pulumi.Input

	ToDataSourcePermissionItemOutput() DataSourcePermissionItemOutput
	ToDataSourcePermissionItemOutputWithContext(ctx context.Context) DataSourcePermissionItemOutput
}

func (*DataSourcePermissionItem) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSourcePermissionItem)(nil)).Elem()
}

func (i *DataSourcePermissionItem) ToDataSourcePermissionItemOutput() DataSourcePermissionItemOutput {
	return i.ToDataSourcePermissionItemOutputWithContext(context.Background())
}

func (i *DataSourcePermissionItem) ToDataSourcePermissionItemOutputWithContext(ctx context.Context) DataSourcePermissionItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourcePermissionItemOutput)
}

// DataSourcePermissionItemArrayInput is an input type that accepts DataSourcePermissionItemArray and DataSourcePermissionItemArrayOutput values.
// You can construct a concrete instance of `DataSourcePermissionItemArrayInput` via:
//
//	DataSourcePermissionItemArray{ DataSourcePermissionItemArgs{...} }
type DataSourcePermissionItemArrayInput interface {
	pulumi.Input

	ToDataSourcePermissionItemArrayOutput() DataSourcePermissionItemArrayOutput
	ToDataSourcePermissionItemArrayOutputWithContext(context.Context) DataSourcePermissionItemArrayOutput
}

type DataSourcePermissionItemArray []DataSourcePermissionItemInput

func (DataSourcePermissionItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataSourcePermissionItem)(nil)).Elem()
}

func (i DataSourcePermissionItemArray) ToDataSourcePermissionItemArrayOutput() DataSourcePermissionItemArrayOutput {
	return i.ToDataSourcePermissionItemArrayOutputWithContext(context.Background())
}

func (i DataSourcePermissionItemArray) ToDataSourcePermissionItemArrayOutputWithContext(ctx context.Context) DataSourcePermissionItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourcePermissionItemArrayOutput)
}

// DataSourcePermissionItemMapInput is an input type that accepts DataSourcePermissionItemMap and DataSourcePermissionItemMapOutput values.
// You can construct a concrete instance of `DataSourcePermissionItemMapInput` via:
//
//	DataSourcePermissionItemMap{ "key": DataSourcePermissionItemArgs{...} }
type DataSourcePermissionItemMapInput interface {
	pulumi.Input

	ToDataSourcePermissionItemMapOutput() DataSourcePermissionItemMapOutput
	ToDataSourcePermissionItemMapOutputWithContext(context.Context) DataSourcePermissionItemMapOutput
}

type DataSourcePermissionItemMap map[string]DataSourcePermissionItemInput

func (DataSourcePermissionItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataSourcePermissionItem)(nil)).Elem()
}

func (i DataSourcePermissionItemMap) ToDataSourcePermissionItemMapOutput() DataSourcePermissionItemMapOutput {
	return i.ToDataSourcePermissionItemMapOutputWithContext(context.Background())
}

func (i DataSourcePermissionItemMap) ToDataSourcePermissionItemMapOutputWithContext(ctx context.Context) DataSourcePermissionItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourcePermissionItemMapOutput)
}

type DataSourcePermissionItemOutput struct{ *pulumi.OutputState }

func (DataSourcePermissionItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSourcePermissionItem)(nil)).Elem()
}

func (o DataSourcePermissionItemOutput) ToDataSourcePermissionItemOutput() DataSourcePermissionItemOutput {
	return o
}

func (o DataSourcePermissionItemOutput) ToDataSourcePermissionItemOutputWithContext(ctx context.Context) DataSourcePermissionItemOutput {
	return o
}

// The UID of the datasource.
func (o DataSourcePermissionItemOutput) DatasourceUid() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSourcePermissionItem) pulumi.StringOutput { return v.DatasourceUid }).(pulumi.StringOutput)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o DataSourcePermissionItemOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSourcePermissionItem) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// the permission to be assigned
func (o DataSourcePermissionItemOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSourcePermissionItem) pulumi.StringOutput { return v.Permission }).(pulumi.StringOutput)
}

// the role onto which the permission is to be assigned
func (o DataSourcePermissionItemOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSourcePermissionItem) pulumi.StringPtrOutput { return v.Role }).(pulumi.StringPtrOutput)
}

// the team onto which the permission is to be assigned
func (o DataSourcePermissionItemOutput) Team() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSourcePermissionItem) pulumi.StringPtrOutput { return v.Team }).(pulumi.StringPtrOutput)
}

// the user or service account onto which the permission is to be assigned
func (o DataSourcePermissionItemOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSourcePermissionItem) pulumi.StringPtrOutput { return v.User }).(pulumi.StringPtrOutput)
}

type DataSourcePermissionItemArrayOutput struct{ *pulumi.OutputState }

func (DataSourcePermissionItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataSourcePermissionItem)(nil)).Elem()
}

func (o DataSourcePermissionItemArrayOutput) ToDataSourcePermissionItemArrayOutput() DataSourcePermissionItemArrayOutput {
	return o
}

func (o DataSourcePermissionItemArrayOutput) ToDataSourcePermissionItemArrayOutputWithContext(ctx context.Context) DataSourcePermissionItemArrayOutput {
	return o
}

func (o DataSourcePermissionItemArrayOutput) Index(i pulumi.IntInput) DataSourcePermissionItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataSourcePermissionItem {
		return vs[0].([]*DataSourcePermissionItem)[vs[1].(int)]
	}).(DataSourcePermissionItemOutput)
}

type DataSourcePermissionItemMapOutput struct{ *pulumi.OutputState }

func (DataSourcePermissionItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataSourcePermissionItem)(nil)).Elem()
}

func (o DataSourcePermissionItemMapOutput) ToDataSourcePermissionItemMapOutput() DataSourcePermissionItemMapOutput {
	return o
}

func (o DataSourcePermissionItemMapOutput) ToDataSourcePermissionItemMapOutputWithContext(ctx context.Context) DataSourcePermissionItemMapOutput {
	return o
}

func (o DataSourcePermissionItemMapOutput) MapIndex(k pulumi.StringInput) DataSourcePermissionItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataSourcePermissionItem {
		return vs[0].(map[string]*DataSourcePermissionItem)[vs[1].(string)]
	}).(DataSourcePermissionItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourcePermissionItemInput)(nil)).Elem(), &DataSourcePermissionItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourcePermissionItemArrayInput)(nil)).Elem(), DataSourcePermissionItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourcePermissionItemMapInput)(nil)).Elem(), DataSourcePermissionItemMap{})
	pulumi.RegisterOutputType(DataSourcePermissionItemOutput{})
	pulumi.RegisterOutputType(DataSourcePermissionItemArrayOutput{})
	pulumi.RegisterOutputType(DataSourcePermissionItemMapOutput{})
}
