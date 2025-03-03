// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// Manages service accounts of a Grafana Cloud stack using the Cloud API
// This can be used to bootstrap a management service account for a new stack
//
// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api)
//
// Required access policy scopes:
//
// * stacks:read
// * stack-service-accounts:write
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/cloud"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloud.NewStackServiceAccount(ctx, "cloud_sa", &cloud.StackServiceAccountArgs{
//				StackSlug:  pulumi.String("<your stack slug>"),
//				Name:       pulumi.String("cloud service account"),
//				Role:       pulumi.String("Admin"),
//				IsDisabled: pulumi.Bool(false),
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
// $ pulumi import grafana:cloud/stackServiceAccount:StackServiceAccount name "{{ stackSlug }}:{{ serviceAccountID }}"
// ```
type StackServiceAccount struct {
	pulumi.CustomResourceState

	// The disabled status for the service account. Defaults to `false`.
	IsDisabled pulumi.BoolPtrOutput `pulumi:"isDisabled"`
	// The name of the service account.
	Name pulumi.StringOutput `pulumi:"name"`
	// The basic role of the service account in the organization.
	Role      pulumi.StringOutput `pulumi:"role"`
	StackSlug pulumi.StringOutput `pulumi:"stackSlug"`
}

// NewStackServiceAccount registers a new resource with the given unique name, arguments, and options.
func NewStackServiceAccount(ctx *pulumi.Context,
	name string, args *StackServiceAccountArgs, opts ...pulumi.ResourceOption) (*StackServiceAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.StackSlug == nil {
		return nil, errors.New("invalid value for required argument 'StackSlug'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StackServiceAccount
	err := ctx.RegisterResource("grafana:cloud/stackServiceAccount:StackServiceAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStackServiceAccount gets an existing StackServiceAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStackServiceAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StackServiceAccountState, opts ...pulumi.ResourceOption) (*StackServiceAccount, error) {
	var resource StackServiceAccount
	err := ctx.ReadResource("grafana:cloud/stackServiceAccount:StackServiceAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StackServiceAccount resources.
type stackServiceAccountState struct {
	// The disabled status for the service account. Defaults to `false`.
	IsDisabled *bool `pulumi:"isDisabled"`
	// The name of the service account.
	Name *string `pulumi:"name"`
	// The basic role of the service account in the organization.
	Role      *string `pulumi:"role"`
	StackSlug *string `pulumi:"stackSlug"`
}

type StackServiceAccountState struct {
	// The disabled status for the service account. Defaults to `false`.
	IsDisabled pulumi.BoolPtrInput
	// The name of the service account.
	Name pulumi.StringPtrInput
	// The basic role of the service account in the organization.
	Role      pulumi.StringPtrInput
	StackSlug pulumi.StringPtrInput
}

func (StackServiceAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*stackServiceAccountState)(nil)).Elem()
}

type stackServiceAccountArgs struct {
	// The disabled status for the service account. Defaults to `false`.
	IsDisabled *bool `pulumi:"isDisabled"`
	// The name of the service account.
	Name *string `pulumi:"name"`
	// The basic role of the service account in the organization.
	Role      string `pulumi:"role"`
	StackSlug string `pulumi:"stackSlug"`
}

// The set of arguments for constructing a StackServiceAccount resource.
type StackServiceAccountArgs struct {
	// The disabled status for the service account. Defaults to `false`.
	IsDisabled pulumi.BoolPtrInput
	// The name of the service account.
	Name pulumi.StringPtrInput
	// The basic role of the service account in the organization.
	Role      pulumi.StringInput
	StackSlug pulumi.StringInput
}

func (StackServiceAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stackServiceAccountArgs)(nil)).Elem()
}

type StackServiceAccountInput interface {
	pulumi.Input

	ToStackServiceAccountOutput() StackServiceAccountOutput
	ToStackServiceAccountOutputWithContext(ctx context.Context) StackServiceAccountOutput
}

func (*StackServiceAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**StackServiceAccount)(nil)).Elem()
}

func (i *StackServiceAccount) ToStackServiceAccountOutput() StackServiceAccountOutput {
	return i.ToStackServiceAccountOutputWithContext(context.Background())
}

func (i *StackServiceAccount) ToStackServiceAccountOutputWithContext(ctx context.Context) StackServiceAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackServiceAccountOutput)
}

// StackServiceAccountArrayInput is an input type that accepts StackServiceAccountArray and StackServiceAccountArrayOutput values.
// You can construct a concrete instance of `StackServiceAccountArrayInput` via:
//
//	StackServiceAccountArray{ StackServiceAccountArgs{...} }
type StackServiceAccountArrayInput interface {
	pulumi.Input

	ToStackServiceAccountArrayOutput() StackServiceAccountArrayOutput
	ToStackServiceAccountArrayOutputWithContext(context.Context) StackServiceAccountArrayOutput
}

type StackServiceAccountArray []StackServiceAccountInput

func (StackServiceAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StackServiceAccount)(nil)).Elem()
}

func (i StackServiceAccountArray) ToStackServiceAccountArrayOutput() StackServiceAccountArrayOutput {
	return i.ToStackServiceAccountArrayOutputWithContext(context.Background())
}

func (i StackServiceAccountArray) ToStackServiceAccountArrayOutputWithContext(ctx context.Context) StackServiceAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackServiceAccountArrayOutput)
}

// StackServiceAccountMapInput is an input type that accepts StackServiceAccountMap and StackServiceAccountMapOutput values.
// You can construct a concrete instance of `StackServiceAccountMapInput` via:
//
//	StackServiceAccountMap{ "key": StackServiceAccountArgs{...} }
type StackServiceAccountMapInput interface {
	pulumi.Input

	ToStackServiceAccountMapOutput() StackServiceAccountMapOutput
	ToStackServiceAccountMapOutputWithContext(context.Context) StackServiceAccountMapOutput
}

type StackServiceAccountMap map[string]StackServiceAccountInput

func (StackServiceAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StackServiceAccount)(nil)).Elem()
}

func (i StackServiceAccountMap) ToStackServiceAccountMapOutput() StackServiceAccountMapOutput {
	return i.ToStackServiceAccountMapOutputWithContext(context.Background())
}

func (i StackServiceAccountMap) ToStackServiceAccountMapOutputWithContext(ctx context.Context) StackServiceAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackServiceAccountMapOutput)
}

type StackServiceAccountOutput struct{ *pulumi.OutputState }

func (StackServiceAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StackServiceAccount)(nil)).Elem()
}

func (o StackServiceAccountOutput) ToStackServiceAccountOutput() StackServiceAccountOutput {
	return o
}

func (o StackServiceAccountOutput) ToStackServiceAccountOutputWithContext(ctx context.Context) StackServiceAccountOutput {
	return o
}

// The disabled status for the service account. Defaults to `false`.
func (o StackServiceAccountOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StackServiceAccount) pulumi.BoolPtrOutput { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

// The name of the service account.
func (o StackServiceAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StackServiceAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The basic role of the service account in the organization.
func (o StackServiceAccountOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *StackServiceAccount) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o StackServiceAccountOutput) StackSlug() pulumi.StringOutput {
	return o.ApplyT(func(v *StackServiceAccount) pulumi.StringOutput { return v.StackSlug }).(pulumi.StringOutput)
}

type StackServiceAccountArrayOutput struct{ *pulumi.OutputState }

func (StackServiceAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StackServiceAccount)(nil)).Elem()
}

func (o StackServiceAccountArrayOutput) ToStackServiceAccountArrayOutput() StackServiceAccountArrayOutput {
	return o
}

func (o StackServiceAccountArrayOutput) ToStackServiceAccountArrayOutputWithContext(ctx context.Context) StackServiceAccountArrayOutput {
	return o
}

func (o StackServiceAccountArrayOutput) Index(i pulumi.IntInput) StackServiceAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StackServiceAccount {
		return vs[0].([]*StackServiceAccount)[vs[1].(int)]
	}).(StackServiceAccountOutput)
}

type StackServiceAccountMapOutput struct{ *pulumi.OutputState }

func (StackServiceAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StackServiceAccount)(nil)).Elem()
}

func (o StackServiceAccountMapOutput) ToStackServiceAccountMapOutput() StackServiceAccountMapOutput {
	return o
}

func (o StackServiceAccountMapOutput) ToStackServiceAccountMapOutputWithContext(ctx context.Context) StackServiceAccountMapOutput {
	return o
}

func (o StackServiceAccountMapOutput) MapIndex(k pulumi.StringInput) StackServiceAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StackServiceAccount {
		return vs[0].(map[string]*StackServiceAccount)[vs[1].(string)]
	}).(StackServiceAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StackServiceAccountInput)(nil)).Elem(), &StackServiceAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackServiceAccountArrayInput)(nil)).Elem(), StackServiceAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackServiceAccountMapInput)(nil)).Elem(), StackServiceAccountMap{})
	pulumi.RegisterOutputType(StackServiceAccountOutput{})
	pulumi.RegisterOutputType(StackServiceAccountArrayOutput{})
	pulumi.RegisterOutputType(StackServiceAccountMapOutput{})
}
