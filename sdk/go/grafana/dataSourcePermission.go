// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// Manages the entire set of permissions for a datasource. Permissions that aren't specified when applying this resource will be removed.
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/datasource_permissions/)
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
//			_, err = enterprise.NewDataSourcePermission(ctx, "fooPermissions", &enterprise.DataSourcePermissionArgs{
//				DatasourceUid: foo.Uid,
//				Permissions: enterprise.DataSourcePermissionPermissionArray{
//					&enterprise.DataSourcePermissionPermissionArgs{
//						TeamId:     team.ID(),
//						Permission: pulumi.String("Edit"),
//					},
//					&enterprise.DataSourcePermissionPermissionArgs{
//						UserId:     user.ID(),
//						Permission: pulumi.String("Edit"),
//					},
//					&enterprise.DataSourcePermissionPermissionArgs{
//						BuiltInRole: pulumi.String("Viewer"),
//						Permission:  pulumi.String("Query"),
//					},
//					&enterprise.DataSourcePermissionPermissionArgs{
//						UserId:     sa.ID(),
//						Permission: pulumi.String("Query"),
//					},
//				},
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
// $ pulumi import grafana:index/dataSourcePermission:DataSourcePermission name "{{ datasourceID }}"
// ```
//
// ```sh
// $ pulumi import grafana:index/dataSourcePermission:DataSourcePermission name "{{ orgID }}:{{ datasourceID }}"
// ```
//
// Deprecated: grafana.index/datasourcepermission.DataSourcePermission has been deprecated in favor of grafana.enterprise/datasourcepermission.DataSourcePermission
type DataSourcePermission struct {
	pulumi.CustomResourceState

	// UID of the datasource to apply permissions to.
	DatasourceUid pulumi.StringOutput `pulumi:"datasourceUid"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// The permission items to add/update. Items that are omitted from the list will be removed.
	Permissions DataSourcePermissionPermissionArrayOutput `pulumi:"permissions"`
}

// NewDataSourcePermission registers a new resource with the given unique name, arguments, and options.
func NewDataSourcePermission(ctx *pulumi.Context,
	name string, args *DataSourcePermissionArgs, opts ...pulumi.ResourceOption) (*DataSourcePermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasourceUid == nil {
		return nil, errors.New("invalid value for required argument 'DatasourceUid'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/dataSourcePermission:DataSourcePermission"),
		},
	})
	opts = append(opts, aliases)
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
	// UID of the datasource to apply permissions to.
	DatasourceUid *string `pulumi:"datasourceUid"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// The permission items to add/update. Items that are omitted from the list will be removed.
	Permissions []DataSourcePermissionPermission `pulumi:"permissions"`
}

type DataSourcePermissionState struct {
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
	// UID of the datasource to apply permissions to.
	DatasourceUid string `pulumi:"datasourceUid"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// The permission items to add/update. Items that are omitted from the list will be removed.
	Permissions []DataSourcePermissionPermission `pulumi:"permissions"`
}

// The set of arguments for constructing a DataSourcePermission resource.
type DataSourcePermissionArgs struct {
	// UID of the datasource to apply permissions to.
	DatasourceUid pulumi.StringInput
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

// UID of the datasource to apply permissions to.
func (o DataSourcePermissionOutput) DatasourceUid() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSourcePermission) pulumi.StringOutput { return v.DatasourceUid }).(pulumi.StringOutput)
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
