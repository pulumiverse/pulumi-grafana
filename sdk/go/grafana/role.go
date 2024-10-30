// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// **Note:** This resource is available only with Grafana Enterprise 8.+.
//
// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/roles-and-permissions/access-control/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/access_control/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/enterprise"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := enterprise.NewRole(ctx, "super_user", &enterprise.RoleArgs{
//				Name:        pulumi.String("Super User"),
//				Description: pulumi.String("My Super User description"),
//				Uid:         pulumi.String("superuseruid"),
//				Version:     pulumi.Int(1),
//				Global:      pulumi.Bool(true),
//				Permissions: enterprise.RolePermissionArray{
//					&enterprise.RolePermissionArgs{
//						Action: pulumi.String("org.users:add"),
//						Scope:  pulumi.String("users:*"),
//					},
//					&enterprise.RolePermissionArgs{
//						Action: pulumi.String("org.users:write"),
//						Scope:  pulumi.String("users:*"),
//					},
//					&enterprise.RolePermissionArgs{
//						Action: pulumi.String("org.users:read"),
//						Scope:  pulumi.String("users:*"),
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
// $ pulumi import grafana:index/role:Role name "{{ uid }}"
// ```
//
// ```sh
// $ pulumi import grafana:index/role:Role name "{{ orgID }}:{{ uid }}"
// ```
//
// Deprecated: grafana.index/role.Role has been deprecated in favor of grafana.enterprise/role.Role
type Role struct {
	pulumi.CustomResourceState

	// Whether the role version should be incremented automatically on updates (and set to 1 on creation). This field or `version` should be set.
	AutoIncrementVersion pulumi.BoolPtrOutput `pulumi:"autoIncrementVersion"`
	// Description of the role.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Display name of the role. Available with Grafana 8.5+.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Boolean to state whether the role is available across all organizations or not. Defaults to `false`.
	Global pulumi.BoolPtrOutput `pulumi:"global"`
	// Group of the role. Available with Grafana 8.5+.
	Group pulumi.StringPtrOutput `pulumi:"group"`
	// Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+. Defaults to `false`.
	Hidden pulumi.BoolPtrOutput `pulumi:"hidden"`
	// Name of the role
	Name pulumi.StringOutput `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// Specific set of actions granted by the role.
	Permissions RolePermissionArrayOutput `pulumi:"permissions"`
	// Unique identifier of the role. Used for assignments.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Version of the role. A role is updated only on version increase. This field or `autoIncrementVersion` should be set.
	Version pulumi.IntPtrOutput `pulumi:"version"`
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil {
		args = &RoleArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/role:Role"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Role
	err := ctx.RegisterResource("grafana:index/role:Role", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleState, opts ...pulumi.ResourceOption) (*Role, error) {
	var resource Role
	err := ctx.ReadResource("grafana:index/role:Role", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Role resources.
type roleState struct {
	// Whether the role version should be incremented automatically on updates (and set to 1 on creation). This field or `version` should be set.
	AutoIncrementVersion *bool `pulumi:"autoIncrementVersion"`
	// Description of the role.
	Description *string `pulumi:"description"`
	// Display name of the role. Available with Grafana 8.5+.
	DisplayName *string `pulumi:"displayName"`
	// Boolean to state whether the role is available across all organizations or not. Defaults to `false`.
	Global *bool `pulumi:"global"`
	// Group of the role. Available with Grafana 8.5+.
	Group *string `pulumi:"group"`
	// Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+. Defaults to `false`.
	Hidden *bool `pulumi:"hidden"`
	// Name of the role
	Name *string `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// Specific set of actions granted by the role.
	Permissions []RolePermission `pulumi:"permissions"`
	// Unique identifier of the role. Used for assignments.
	Uid *string `pulumi:"uid"`
	// Version of the role. A role is updated only on version increase. This field or `autoIncrementVersion` should be set.
	Version *int `pulumi:"version"`
}

type RoleState struct {
	// Whether the role version should be incremented automatically on updates (and set to 1 on creation). This field or `version` should be set.
	AutoIncrementVersion pulumi.BoolPtrInput
	// Description of the role.
	Description pulumi.StringPtrInput
	// Display name of the role. Available with Grafana 8.5+.
	DisplayName pulumi.StringPtrInput
	// Boolean to state whether the role is available across all organizations or not. Defaults to `false`.
	Global pulumi.BoolPtrInput
	// Group of the role. Available with Grafana 8.5+.
	Group pulumi.StringPtrInput
	// Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+. Defaults to `false`.
	Hidden pulumi.BoolPtrInput
	// Name of the role
	Name pulumi.StringPtrInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// Specific set of actions granted by the role.
	Permissions RolePermissionArrayInput
	// Unique identifier of the role. Used for assignments.
	Uid pulumi.StringPtrInput
	// Version of the role. A role is updated only on version increase. This field or `autoIncrementVersion` should be set.
	Version pulumi.IntPtrInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	// Whether the role version should be incremented automatically on updates (and set to 1 on creation). This field or `version` should be set.
	AutoIncrementVersion *bool `pulumi:"autoIncrementVersion"`
	// Description of the role.
	Description *string `pulumi:"description"`
	// Display name of the role. Available with Grafana 8.5+.
	DisplayName *string `pulumi:"displayName"`
	// Boolean to state whether the role is available across all organizations or not. Defaults to `false`.
	Global *bool `pulumi:"global"`
	// Group of the role. Available with Grafana 8.5+.
	Group *string `pulumi:"group"`
	// Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+. Defaults to `false`.
	Hidden *bool `pulumi:"hidden"`
	// Name of the role
	Name *string `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// Specific set of actions granted by the role.
	Permissions []RolePermission `pulumi:"permissions"`
	// Unique identifier of the role. Used for assignments.
	Uid *string `pulumi:"uid"`
	// Version of the role. A role is updated only on version increase. This field or `autoIncrementVersion` should be set.
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	// Whether the role version should be incremented automatically on updates (and set to 1 on creation). This field or `version` should be set.
	AutoIncrementVersion pulumi.BoolPtrInput
	// Description of the role.
	Description pulumi.StringPtrInput
	// Display name of the role. Available with Grafana 8.5+.
	DisplayName pulumi.StringPtrInput
	// Boolean to state whether the role is available across all organizations or not. Defaults to `false`.
	Global pulumi.BoolPtrInput
	// Group of the role. Available with Grafana 8.5+.
	Group pulumi.StringPtrInput
	// Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+. Defaults to `false`.
	Hidden pulumi.BoolPtrInput
	// Name of the role
	Name pulumi.StringPtrInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// Specific set of actions granted by the role.
	Permissions RolePermissionArrayInput
	// Unique identifier of the role. Used for assignments.
	Uid pulumi.StringPtrInput
	// Version of the role. A role is updated only on version increase. This field or `autoIncrementVersion` should be set.
	Version pulumi.IntPtrInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}

type RoleInput interface {
	pulumi.Input

	ToRoleOutput() RoleOutput
	ToRoleOutputWithContext(ctx context.Context) RoleOutput
}

func (*Role) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (i *Role) ToRoleOutput() RoleOutput {
	return i.ToRoleOutputWithContext(context.Background())
}

func (i *Role) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleOutput)
}

// RoleArrayInput is an input type that accepts RoleArray and RoleArrayOutput values.
// You can construct a concrete instance of `RoleArrayInput` via:
//
//	RoleArray{ RoleArgs{...} }
type RoleArrayInput interface {
	pulumi.Input

	ToRoleArrayOutput() RoleArrayOutput
	ToRoleArrayOutputWithContext(context.Context) RoleArrayOutput
}

type RoleArray []RoleInput

func (RoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (i RoleArray) ToRoleArrayOutput() RoleArrayOutput {
	return i.ToRoleArrayOutputWithContext(context.Background())
}

func (i RoleArray) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleArrayOutput)
}

// RoleMapInput is an input type that accepts RoleMap and RoleMapOutput values.
// You can construct a concrete instance of `RoleMapInput` via:
//
//	RoleMap{ "key": RoleArgs{...} }
type RoleMapInput interface {
	pulumi.Input

	ToRoleMapOutput() RoleMapOutput
	ToRoleMapOutputWithContext(context.Context) RoleMapOutput
}

type RoleMap map[string]RoleInput

func (RoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (i RoleMap) ToRoleMapOutput() RoleMapOutput {
	return i.ToRoleMapOutputWithContext(context.Background())
}

func (i RoleMap) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleMapOutput)
}

type RoleOutput struct{ *pulumi.OutputState }

func (RoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (o RoleOutput) ToRoleOutput() RoleOutput {
	return o
}

func (o RoleOutput) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return o
}

// Whether the role version should be incremented automatically on updates (and set to 1 on creation). This field or `version` should be set.
func (o RoleOutput) AutoIncrementVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.BoolPtrOutput { return v.AutoIncrementVersion }).(pulumi.BoolPtrOutput)
}

// Description of the role.
func (o RoleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Display name of the role. Available with Grafana 8.5+.
func (o RoleOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Boolean to state whether the role is available across all organizations or not. Defaults to `false`.
func (o RoleOutput) Global() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.BoolPtrOutput { return v.Global }).(pulumi.BoolPtrOutput)
}

// Group of the role. Available with Grafana 8.5+.
func (o RoleOutput) Group() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.StringPtrOutput { return v.Group }).(pulumi.StringPtrOutput)
}

// Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+. Defaults to `false`.
func (o RoleOutput) Hidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.BoolPtrOutput { return v.Hidden }).(pulumi.BoolPtrOutput)
}

// Name of the role
func (o RoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o RoleOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Specific set of actions granted by the role.
func (o RoleOutput) Permissions() RolePermissionArrayOutput {
	return o.ApplyT(func(v *Role) RolePermissionArrayOutput { return v.Permissions }).(RolePermissionArrayOutput)
}

// Unique identifier of the role. Used for assignments.
func (o RoleOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Version of the role. A role is updated only on version increase. This field or `autoIncrementVersion` should be set.
func (o RoleOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.IntPtrOutput { return v.Version }).(pulumi.IntPtrOutput)
}

type RoleArrayOutput struct{ *pulumi.OutputState }

func (RoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (o RoleArrayOutput) ToRoleArrayOutput() RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) Index(i pulumi.IntInput) RoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Role {
		return vs[0].([]*Role)[vs[1].(int)]
	}).(RoleOutput)
}

type RoleMapOutput struct{ *pulumi.OutputState }

func (RoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (o RoleMapOutput) ToRoleMapOutput() RoleMapOutput {
	return o
}

func (o RoleMapOutput) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return o
}

func (o RoleMapOutput) MapIndex(k pulumi.StringInput) RoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Role {
		return vs[0].(map[string]*Role)[vs[1].(string)]
	}).(RoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleInput)(nil)).Elem(), &Role{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleArrayInput)(nil)).Elem(), RoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleMapInput)(nil)).Elem(), RoleMap{})
	pulumi.RegisterOutputType(RoleOutput{})
	pulumi.RegisterOutputType(RoleArrayOutput{})
	pulumi.RegisterOutputType(RoleMapOutput{})
}
