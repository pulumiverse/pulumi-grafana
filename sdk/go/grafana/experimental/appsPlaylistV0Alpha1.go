// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package experimental

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// Manages Grafana playlists using the new Grafana APIs.
//
// * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/create-manage-playlists/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/apis/)
type AppsPlaylistV0Alpha1 struct {
	pulumi.CustomResourceState

	// The metadata of the resource.
	Metadata AppsPlaylistV0Alpha1MetadataPtrOutput `pulumi:"metadata"`
	// Options for applying the resource.
	Options AppsPlaylistV0Alpha1OptionsPtrOutput `pulumi:"options"`
	// The spec of the resource.
	Spec AppsPlaylistV0Alpha1SpecPtrOutput `pulumi:"spec"`
}

// NewAppsPlaylistV0Alpha1 registers a new resource with the given unique name, arguments, and options.
func NewAppsPlaylistV0Alpha1(ctx *pulumi.Context,
	name string, args *AppsPlaylistV0Alpha1Args, opts ...pulumi.ResourceOption) (*AppsPlaylistV0Alpha1, error) {
	if args == nil {
		args = &AppsPlaylistV0Alpha1Args{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AppsPlaylistV0Alpha1
	err := ctx.RegisterResource("grafana:experimental/appsPlaylistV0Alpha1:AppsPlaylistV0Alpha1", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppsPlaylistV0Alpha1 gets an existing AppsPlaylistV0Alpha1 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppsPlaylistV0Alpha1(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppsPlaylistV0Alpha1State, opts ...pulumi.ResourceOption) (*AppsPlaylistV0Alpha1, error) {
	var resource AppsPlaylistV0Alpha1
	err := ctx.ReadResource("grafana:experimental/appsPlaylistV0Alpha1:AppsPlaylistV0Alpha1", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppsPlaylistV0Alpha1 resources.
type appsPlaylistV0Alpha1State struct {
	// The metadata of the resource.
	Metadata *AppsPlaylistV0Alpha1Metadata `pulumi:"metadata"`
	// Options for applying the resource.
	Options *AppsPlaylistV0Alpha1Options `pulumi:"options"`
	// The spec of the resource.
	Spec *AppsPlaylistV0Alpha1Spec `pulumi:"spec"`
}

type AppsPlaylistV0Alpha1State struct {
	// The metadata of the resource.
	Metadata AppsPlaylistV0Alpha1MetadataPtrInput
	// Options for applying the resource.
	Options AppsPlaylistV0Alpha1OptionsPtrInput
	// The spec of the resource.
	Spec AppsPlaylistV0Alpha1SpecPtrInput
}

func (AppsPlaylistV0Alpha1State) ElementType() reflect.Type {
	return reflect.TypeOf((*appsPlaylistV0Alpha1State)(nil)).Elem()
}

type appsPlaylistV0Alpha1Args struct {
	// The metadata of the resource.
	Metadata *AppsPlaylistV0Alpha1Metadata `pulumi:"metadata"`
	// Options for applying the resource.
	Options *AppsPlaylistV0Alpha1Options `pulumi:"options"`
	// The spec of the resource.
	Spec *AppsPlaylistV0Alpha1Spec `pulumi:"spec"`
}

// The set of arguments for constructing a AppsPlaylistV0Alpha1 resource.
type AppsPlaylistV0Alpha1Args struct {
	// The metadata of the resource.
	Metadata AppsPlaylistV0Alpha1MetadataPtrInput
	// Options for applying the resource.
	Options AppsPlaylistV0Alpha1OptionsPtrInput
	// The spec of the resource.
	Spec AppsPlaylistV0Alpha1SpecPtrInput
}

func (AppsPlaylistV0Alpha1Args) ElementType() reflect.Type {
	return reflect.TypeOf((*appsPlaylistV0Alpha1Args)(nil)).Elem()
}

type AppsPlaylistV0Alpha1Input interface {
	pulumi.Input

	ToAppsPlaylistV0Alpha1Output() AppsPlaylistV0Alpha1Output
	ToAppsPlaylistV0Alpha1OutputWithContext(ctx context.Context) AppsPlaylistV0Alpha1Output
}

func (*AppsPlaylistV0Alpha1) ElementType() reflect.Type {
	return reflect.TypeOf((**AppsPlaylistV0Alpha1)(nil)).Elem()
}

func (i *AppsPlaylistV0Alpha1) ToAppsPlaylistV0Alpha1Output() AppsPlaylistV0Alpha1Output {
	return i.ToAppsPlaylistV0Alpha1OutputWithContext(context.Background())
}

func (i *AppsPlaylistV0Alpha1) ToAppsPlaylistV0Alpha1OutputWithContext(ctx context.Context) AppsPlaylistV0Alpha1Output {
	return pulumi.ToOutputWithContext(ctx, i).(AppsPlaylistV0Alpha1Output)
}

// AppsPlaylistV0Alpha1ArrayInput is an input type that accepts AppsPlaylistV0Alpha1Array and AppsPlaylistV0Alpha1ArrayOutput values.
// You can construct a concrete instance of `AppsPlaylistV0Alpha1ArrayInput` via:
//
//	AppsPlaylistV0Alpha1Array{ AppsPlaylistV0Alpha1Args{...} }
type AppsPlaylistV0Alpha1ArrayInput interface {
	pulumi.Input

	ToAppsPlaylistV0Alpha1ArrayOutput() AppsPlaylistV0Alpha1ArrayOutput
	ToAppsPlaylistV0Alpha1ArrayOutputWithContext(context.Context) AppsPlaylistV0Alpha1ArrayOutput
}

type AppsPlaylistV0Alpha1Array []AppsPlaylistV0Alpha1Input

func (AppsPlaylistV0Alpha1Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppsPlaylistV0Alpha1)(nil)).Elem()
}

func (i AppsPlaylistV0Alpha1Array) ToAppsPlaylistV0Alpha1ArrayOutput() AppsPlaylistV0Alpha1ArrayOutput {
	return i.ToAppsPlaylistV0Alpha1ArrayOutputWithContext(context.Background())
}

func (i AppsPlaylistV0Alpha1Array) ToAppsPlaylistV0Alpha1ArrayOutputWithContext(ctx context.Context) AppsPlaylistV0Alpha1ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppsPlaylistV0Alpha1ArrayOutput)
}

// AppsPlaylistV0Alpha1MapInput is an input type that accepts AppsPlaylistV0Alpha1Map and AppsPlaylistV0Alpha1MapOutput values.
// You can construct a concrete instance of `AppsPlaylistV0Alpha1MapInput` via:
//
//	AppsPlaylistV0Alpha1Map{ "key": AppsPlaylistV0Alpha1Args{...} }
type AppsPlaylistV0Alpha1MapInput interface {
	pulumi.Input

	ToAppsPlaylistV0Alpha1MapOutput() AppsPlaylistV0Alpha1MapOutput
	ToAppsPlaylistV0Alpha1MapOutputWithContext(context.Context) AppsPlaylistV0Alpha1MapOutput
}

type AppsPlaylistV0Alpha1Map map[string]AppsPlaylistV0Alpha1Input

func (AppsPlaylistV0Alpha1Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppsPlaylistV0Alpha1)(nil)).Elem()
}

func (i AppsPlaylistV0Alpha1Map) ToAppsPlaylistV0Alpha1MapOutput() AppsPlaylistV0Alpha1MapOutput {
	return i.ToAppsPlaylistV0Alpha1MapOutputWithContext(context.Background())
}

func (i AppsPlaylistV0Alpha1Map) ToAppsPlaylistV0Alpha1MapOutputWithContext(ctx context.Context) AppsPlaylistV0Alpha1MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppsPlaylistV0Alpha1MapOutput)
}

type AppsPlaylistV0Alpha1Output struct{ *pulumi.OutputState }

func (AppsPlaylistV0Alpha1Output) ElementType() reflect.Type {
	return reflect.TypeOf((**AppsPlaylistV0Alpha1)(nil)).Elem()
}

func (o AppsPlaylistV0Alpha1Output) ToAppsPlaylistV0Alpha1Output() AppsPlaylistV0Alpha1Output {
	return o
}

func (o AppsPlaylistV0Alpha1Output) ToAppsPlaylistV0Alpha1OutputWithContext(ctx context.Context) AppsPlaylistV0Alpha1Output {
	return o
}

// The metadata of the resource.
func (o AppsPlaylistV0Alpha1Output) Metadata() AppsPlaylistV0Alpha1MetadataPtrOutput {
	return o.ApplyT(func(v *AppsPlaylistV0Alpha1) AppsPlaylistV0Alpha1MetadataPtrOutput { return v.Metadata }).(AppsPlaylistV0Alpha1MetadataPtrOutput)
}

// Options for applying the resource.
func (o AppsPlaylistV0Alpha1Output) Options() AppsPlaylistV0Alpha1OptionsPtrOutput {
	return o.ApplyT(func(v *AppsPlaylistV0Alpha1) AppsPlaylistV0Alpha1OptionsPtrOutput { return v.Options }).(AppsPlaylistV0Alpha1OptionsPtrOutput)
}

// The spec of the resource.
func (o AppsPlaylistV0Alpha1Output) Spec() AppsPlaylistV0Alpha1SpecPtrOutput {
	return o.ApplyT(func(v *AppsPlaylistV0Alpha1) AppsPlaylistV0Alpha1SpecPtrOutput { return v.Spec }).(AppsPlaylistV0Alpha1SpecPtrOutput)
}

type AppsPlaylistV0Alpha1ArrayOutput struct{ *pulumi.OutputState }

func (AppsPlaylistV0Alpha1ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppsPlaylistV0Alpha1)(nil)).Elem()
}

func (o AppsPlaylistV0Alpha1ArrayOutput) ToAppsPlaylistV0Alpha1ArrayOutput() AppsPlaylistV0Alpha1ArrayOutput {
	return o
}

func (o AppsPlaylistV0Alpha1ArrayOutput) ToAppsPlaylistV0Alpha1ArrayOutputWithContext(ctx context.Context) AppsPlaylistV0Alpha1ArrayOutput {
	return o
}

func (o AppsPlaylistV0Alpha1ArrayOutput) Index(i pulumi.IntInput) AppsPlaylistV0Alpha1Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppsPlaylistV0Alpha1 {
		return vs[0].([]*AppsPlaylistV0Alpha1)[vs[1].(int)]
	}).(AppsPlaylistV0Alpha1Output)
}

type AppsPlaylistV0Alpha1MapOutput struct{ *pulumi.OutputState }

func (AppsPlaylistV0Alpha1MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppsPlaylistV0Alpha1)(nil)).Elem()
}

func (o AppsPlaylistV0Alpha1MapOutput) ToAppsPlaylistV0Alpha1MapOutput() AppsPlaylistV0Alpha1MapOutput {
	return o
}

func (o AppsPlaylistV0Alpha1MapOutput) ToAppsPlaylistV0Alpha1MapOutputWithContext(ctx context.Context) AppsPlaylistV0Alpha1MapOutput {
	return o
}

func (o AppsPlaylistV0Alpha1MapOutput) MapIndex(k pulumi.StringInput) AppsPlaylistV0Alpha1Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppsPlaylistV0Alpha1 {
		return vs[0].(map[string]*AppsPlaylistV0Alpha1)[vs[1].(string)]
	}).(AppsPlaylistV0Alpha1Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppsPlaylistV0Alpha1Input)(nil)).Elem(), &AppsPlaylistV0Alpha1{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppsPlaylistV0Alpha1ArrayInput)(nil)).Elem(), AppsPlaylistV0Alpha1Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppsPlaylistV0Alpha1MapInput)(nil)).Elem(), AppsPlaylistV0Alpha1Map{})
	pulumi.RegisterOutputType(AppsPlaylistV0Alpha1Output{})
	pulumi.RegisterOutputType(AppsPlaylistV0Alpha1ArrayOutput{})
	pulumi.RegisterOutputType(AppsPlaylistV0Alpha1MapOutput{})
}
