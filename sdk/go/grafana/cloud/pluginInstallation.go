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

// Manages Grafana Cloud Plugin Installations.
//
// * [Plugin Catalog](https://grafana.com/grafana/plugins/)
//
// Required access policy scopes:
//
// * stack-plugins:read
// * stack-plugins:write
// * stack-plugins:delete
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
//			_, err := cloud.NewPluginInstallation(ctx, "test", &cloud.PluginInstallationArgs{
//				StackSlug: pulumi.String("stackname"),
//				Slug:      pulumi.String("some-plugin"),
//				Version:   pulumi.String("1.2.3"),
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
// $ pulumi import grafana:cloud/pluginInstallation:PluginInstallation name "{{ stackSlug }}:{{ pluginSlug }}"
// ```
type PluginInstallation struct {
	pulumi.CustomResourceState

	// Slug of the plugin to be installed.
	Slug pulumi.StringOutput `pulumi:"slug"`
	// The stack id to which the plugin should be installed.
	StackSlug pulumi.StringOutput `pulumi:"stackSlug"`
	// Version of the plugin to be installed.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewPluginInstallation registers a new resource with the given unique name, arguments, and options.
func NewPluginInstallation(ctx *pulumi.Context,
	name string, args *PluginInstallationArgs, opts ...pulumi.ResourceOption) (*PluginInstallation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Slug == nil {
		return nil, errors.New("invalid value for required argument 'Slug'")
	}
	if args.StackSlug == nil {
		return nil, errors.New("invalid value for required argument 'StackSlug'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PluginInstallation
	err := ctx.RegisterResource("grafana:cloud/pluginInstallation:PluginInstallation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPluginInstallation gets an existing PluginInstallation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPluginInstallation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PluginInstallationState, opts ...pulumi.ResourceOption) (*PluginInstallation, error) {
	var resource PluginInstallation
	err := ctx.ReadResource("grafana:cloud/pluginInstallation:PluginInstallation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PluginInstallation resources.
type pluginInstallationState struct {
	// Slug of the plugin to be installed.
	Slug *string `pulumi:"slug"`
	// The stack id to which the plugin should be installed.
	StackSlug *string `pulumi:"stackSlug"`
	// Version of the plugin to be installed.
	Version *string `pulumi:"version"`
}

type PluginInstallationState struct {
	// Slug of the plugin to be installed.
	Slug pulumi.StringPtrInput
	// The stack id to which the plugin should be installed.
	StackSlug pulumi.StringPtrInput
	// Version of the plugin to be installed.
	Version pulumi.StringPtrInput
}

func (PluginInstallationState) ElementType() reflect.Type {
	return reflect.TypeOf((*pluginInstallationState)(nil)).Elem()
}

type pluginInstallationArgs struct {
	// Slug of the plugin to be installed.
	Slug string `pulumi:"slug"`
	// The stack id to which the plugin should be installed.
	StackSlug string `pulumi:"stackSlug"`
	// Version of the plugin to be installed.
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a PluginInstallation resource.
type PluginInstallationArgs struct {
	// Slug of the plugin to be installed.
	Slug pulumi.StringInput
	// The stack id to which the plugin should be installed.
	StackSlug pulumi.StringInput
	// Version of the plugin to be installed.
	Version pulumi.StringInput
}

func (PluginInstallationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pluginInstallationArgs)(nil)).Elem()
}

type PluginInstallationInput interface {
	pulumi.Input

	ToPluginInstallationOutput() PluginInstallationOutput
	ToPluginInstallationOutputWithContext(ctx context.Context) PluginInstallationOutput
}

func (*PluginInstallation) ElementType() reflect.Type {
	return reflect.TypeOf((**PluginInstallation)(nil)).Elem()
}

func (i *PluginInstallation) ToPluginInstallationOutput() PluginInstallationOutput {
	return i.ToPluginInstallationOutputWithContext(context.Background())
}

func (i *PluginInstallation) ToPluginInstallationOutputWithContext(ctx context.Context) PluginInstallationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluginInstallationOutput)
}

// PluginInstallationArrayInput is an input type that accepts PluginInstallationArray and PluginInstallationArrayOutput values.
// You can construct a concrete instance of `PluginInstallationArrayInput` via:
//
//	PluginInstallationArray{ PluginInstallationArgs{...} }
type PluginInstallationArrayInput interface {
	pulumi.Input

	ToPluginInstallationArrayOutput() PluginInstallationArrayOutput
	ToPluginInstallationArrayOutputWithContext(context.Context) PluginInstallationArrayOutput
}

type PluginInstallationArray []PluginInstallationInput

func (PluginInstallationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PluginInstallation)(nil)).Elem()
}

func (i PluginInstallationArray) ToPluginInstallationArrayOutput() PluginInstallationArrayOutput {
	return i.ToPluginInstallationArrayOutputWithContext(context.Background())
}

func (i PluginInstallationArray) ToPluginInstallationArrayOutputWithContext(ctx context.Context) PluginInstallationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluginInstallationArrayOutput)
}

// PluginInstallationMapInput is an input type that accepts PluginInstallationMap and PluginInstallationMapOutput values.
// You can construct a concrete instance of `PluginInstallationMapInput` via:
//
//	PluginInstallationMap{ "key": PluginInstallationArgs{...} }
type PluginInstallationMapInput interface {
	pulumi.Input

	ToPluginInstallationMapOutput() PluginInstallationMapOutput
	ToPluginInstallationMapOutputWithContext(context.Context) PluginInstallationMapOutput
}

type PluginInstallationMap map[string]PluginInstallationInput

func (PluginInstallationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PluginInstallation)(nil)).Elem()
}

func (i PluginInstallationMap) ToPluginInstallationMapOutput() PluginInstallationMapOutput {
	return i.ToPluginInstallationMapOutputWithContext(context.Background())
}

func (i PluginInstallationMap) ToPluginInstallationMapOutputWithContext(ctx context.Context) PluginInstallationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluginInstallationMapOutput)
}

type PluginInstallationOutput struct{ *pulumi.OutputState }

func (PluginInstallationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PluginInstallation)(nil)).Elem()
}

func (o PluginInstallationOutput) ToPluginInstallationOutput() PluginInstallationOutput {
	return o
}

func (o PluginInstallationOutput) ToPluginInstallationOutputWithContext(ctx context.Context) PluginInstallationOutput {
	return o
}

// Slug of the plugin to be installed.
func (o PluginInstallationOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *PluginInstallation) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

// The stack id to which the plugin should be installed.
func (o PluginInstallationOutput) StackSlug() pulumi.StringOutput {
	return o.ApplyT(func(v *PluginInstallation) pulumi.StringOutput { return v.StackSlug }).(pulumi.StringOutput)
}

// Version of the plugin to be installed.
func (o PluginInstallationOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *PluginInstallation) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type PluginInstallationArrayOutput struct{ *pulumi.OutputState }

func (PluginInstallationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PluginInstallation)(nil)).Elem()
}

func (o PluginInstallationArrayOutput) ToPluginInstallationArrayOutput() PluginInstallationArrayOutput {
	return o
}

func (o PluginInstallationArrayOutput) ToPluginInstallationArrayOutputWithContext(ctx context.Context) PluginInstallationArrayOutput {
	return o
}

func (o PluginInstallationArrayOutput) Index(i pulumi.IntInput) PluginInstallationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PluginInstallation {
		return vs[0].([]*PluginInstallation)[vs[1].(int)]
	}).(PluginInstallationOutput)
}

type PluginInstallationMapOutput struct{ *pulumi.OutputState }

func (PluginInstallationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PluginInstallation)(nil)).Elem()
}

func (o PluginInstallationMapOutput) ToPluginInstallationMapOutput() PluginInstallationMapOutput {
	return o
}

func (o PluginInstallationMapOutput) ToPluginInstallationMapOutputWithContext(ctx context.Context) PluginInstallationMapOutput {
	return o
}

func (o PluginInstallationMapOutput) MapIndex(k pulumi.StringInput) PluginInstallationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PluginInstallation {
		return vs[0].(map[string]*PluginInstallation)[vs[1].(string)]
	}).(PluginInstallationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PluginInstallationInput)(nil)).Elem(), &PluginInstallation{})
	pulumi.RegisterInputType(reflect.TypeOf((*PluginInstallationArrayInput)(nil)).Elem(), PluginInstallationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PluginInstallationMapInput)(nil)).Elem(), PluginInstallationMap{})
	pulumi.RegisterOutputType(PluginInstallationOutput{})
	pulumi.RegisterOutputType(PluginInstallationArrayOutput{})
	pulumi.RegisterOutputType(PluginInstallationMapOutput{})
}
