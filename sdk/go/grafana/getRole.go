// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
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
//	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := grafana.NewRole(ctx, "test", &grafana.RoleArgs{
//				Description: pulumi.String("test-role description"),
//				Uid:         pulumi.String("test-ds-role-uid"),
//				Version:     pulumi.Int(1),
//				Global:      pulumi.Bool(true),
//				Permissions: grafana.RolePermissionArray{
//					&grafana.RolePermissionArgs{
//						Action: pulumi.String("org.users:add"),
//						Scope:  pulumi.String("users:*"),
//					},
//					&grafana.RolePermissionArgs{
//						Action: pulumi.String("org.users:write"),
//						Scope:  pulumi.String("users:*"),
//					},
//					&grafana.RolePermissionArgs{
//						Action: pulumi.String("org.users:read"),
//						Scope:  pulumi.String("users:*"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = grafana.LookupRoleOutput(ctx, grafana.GetRoleOutputArgs{
//				Name: test.Name,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupRole(ctx *pulumi.Context, args *LookupRoleArgs, opts ...pulumi.InvokeOption) (*LookupRoleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRoleResult
	err := ctx.Invoke("grafana:index/getRole:getRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRole.
type LookupRoleArgs struct {
	// Name of the role
	Name string `pulumi:"name"`
}

// A collection of values returned by getRole.
type LookupRoleResult struct {
	// Description of the role.
	Description string `pulumi:"description"`
	// Display name of the role. Available with Grafana 8.5+.
	DisplayName string `pulumi:"displayName"`
	// Boolean to state whether the role is available across all organizations or not.
	Global bool `pulumi:"global"`
	// Group of the role. Available with Grafana 8.5+.
	Group string `pulumi:"group"`
	// Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+.
	Hidden bool `pulumi:"hidden"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the role
	Name string `pulumi:"name"`
	// Specific set of actions granted by the role.
	Permissions []GetRolePermission `pulumi:"permissions"`
	// Unique identifier of the role. Used for assignments.
	Uid string `pulumi:"uid"`
	// Version of the role. A role is updated only on version increase.
	Version int `pulumi:"version"`
}

func LookupRoleOutput(ctx *pulumi.Context, args LookupRoleOutputArgs, opts ...pulumi.InvokeOption) LookupRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoleResult, error) {
			args := v.(LookupRoleArgs)
			r, err := LookupRole(ctx, &args, opts...)
			var s LookupRoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoleResultOutput)
}

// A collection of arguments for invoking getRole.
type LookupRoleOutputArgs struct {
	// Name of the role
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleArgs)(nil)).Elem()
}

// A collection of values returned by getRole.
type LookupRoleResultOutput struct{ *pulumi.OutputState }

func (LookupRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleResult)(nil)).Elem()
}

func (o LookupRoleResultOutput) ToLookupRoleResultOutput() LookupRoleResultOutput {
	return o
}

func (o LookupRoleResultOutput) ToLookupRoleResultOutputWithContext(ctx context.Context) LookupRoleResultOutput {
	return o
}

func (o LookupRoleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRoleResult] {
	return pulumix.Output[LookupRoleResult]{
		OutputState: o.OutputState,
	}
}

// Description of the role.
func (o LookupRoleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.Description }).(pulumi.StringOutput)
}

// Display name of the role. Available with Grafana 8.5+.
func (o LookupRoleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Boolean to state whether the role is available across all organizations or not.
func (o LookupRoleResultOutput) Global() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRoleResult) bool { return v.Global }).(pulumi.BoolOutput)
}

// Group of the role. Available with Grafana 8.5+.
func (o LookupRoleResultOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.Group }).(pulumi.StringOutput)
}

// Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+.
func (o LookupRoleResultOutput) Hidden() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRoleResult) bool { return v.Hidden }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the role
func (o LookupRoleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specific set of actions granted by the role.
func (o LookupRoleResultOutput) Permissions() GetRolePermissionArrayOutput {
	return o.ApplyT(func(v LookupRoleResult) []GetRolePermission { return v.Permissions }).(GetRolePermissionArrayOutput)
}

// Unique identifier of the role. Used for assignments.
func (o LookupRoleResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.Uid }).(pulumi.StringOutput)
}

// Version of the role. A role is updated only on version increase.
func (o LookupRoleResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRoleResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoleResultOutput{})
}