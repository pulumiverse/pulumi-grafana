// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package syntheticmonitoring

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// Besides the public probes run by Grafana Labs, you can also install your
// own private probes. These are only accessible to you and only write data to
// your Grafana Cloud account. Private probes are instances of the open source
// Grafana Synthetic Monitoring Agent.
//
// * [Official documentation](https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/set-up-private-probes/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/syntheticmonitoring"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := syntheticmonitoring.NewProbe(ctx, "main", &syntheticmonitoring.ProbeArgs{
//				Name:      pulumi.String("Mount Everest"),
//				Latitude:  pulumi.Float64(27.98606),
//				Longitude: pulumi.Float64(86.92262),
//				Region:    pulumi.String("APAC"),
//				Labels: pulumi.StringMap{
//					"type": pulumi.String("mountain"),
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
// $ pulumi import grafana:syntheticMonitoring/probe:Probe name "{{ id }}"
// ```
//
// ```sh
// $ pulumi import grafana:syntheticMonitoring/probe:Probe name "{{ id }}:{{ authToken }}"
// ```
type Probe struct {
	pulumi.CustomResourceState

	// The probe authentication token. Your probe must use this to authenticate with Grafana Cloud.
	AuthToken pulumi.StringOutput `pulumi:"authToken"`
	// Disables scripted checks for this probe. Defaults to `false`.
	DisableScriptedChecks pulumi.BoolPtrOutput `pulumi:"disableScriptedChecks"`
	// Custom labels to be included with collected metrics and logs.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Latitude coordinates.
	Latitude pulumi.Float64Output `pulumi:"latitude"`
	// Longitude coordinates.
	Longitude pulumi.Float64Output `pulumi:"longitude"`
	// Name of the probe.
	Name pulumi.StringOutput `pulumi:"name"`
	// Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
	Public pulumi.BoolPtrOutput `pulumi:"public"`
	// Region of the probe.
	Region pulumi.StringOutput `pulumi:"region"`
	// The tenant ID of the probe.
	TenantId pulumi.IntOutput `pulumi:"tenantId"`
}

// NewProbe registers a new resource with the given unique name, arguments, and options.
func NewProbe(ctx *pulumi.Context,
	name string, args *ProbeArgs, opts ...pulumi.ResourceOption) (*Probe, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Latitude == nil {
		return nil, errors.New("invalid value for required argument 'Latitude'")
	}
	if args.Longitude == nil {
		return nil, errors.New("invalid value for required argument 'Longitude'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/syntheticMonitoringProbe:SyntheticMonitoringProbe"),
		},
	})
	opts = append(opts, aliases)
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"authToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Probe
	err := ctx.RegisterResource("grafana:syntheticMonitoring/probe:Probe", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProbe gets an existing Probe resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProbe(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProbeState, opts ...pulumi.ResourceOption) (*Probe, error) {
	var resource Probe
	err := ctx.ReadResource("grafana:syntheticMonitoring/probe:Probe", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Probe resources.
type probeState struct {
	// The probe authentication token. Your probe must use this to authenticate with Grafana Cloud.
	AuthToken *string `pulumi:"authToken"`
	// Disables scripted checks for this probe. Defaults to `false`.
	DisableScriptedChecks *bool `pulumi:"disableScriptedChecks"`
	// Custom labels to be included with collected metrics and logs.
	Labels map[string]string `pulumi:"labels"`
	// Latitude coordinates.
	Latitude *float64 `pulumi:"latitude"`
	// Longitude coordinates.
	Longitude *float64 `pulumi:"longitude"`
	// Name of the probe.
	Name *string `pulumi:"name"`
	// Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
	Public *bool `pulumi:"public"`
	// Region of the probe.
	Region *string `pulumi:"region"`
	// The tenant ID of the probe.
	TenantId *int `pulumi:"tenantId"`
}

type ProbeState struct {
	// The probe authentication token. Your probe must use this to authenticate with Grafana Cloud.
	AuthToken pulumi.StringPtrInput
	// Disables scripted checks for this probe. Defaults to `false`.
	DisableScriptedChecks pulumi.BoolPtrInput
	// Custom labels to be included with collected metrics and logs.
	Labels pulumi.StringMapInput
	// Latitude coordinates.
	Latitude pulumi.Float64PtrInput
	// Longitude coordinates.
	Longitude pulumi.Float64PtrInput
	// Name of the probe.
	Name pulumi.StringPtrInput
	// Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
	Public pulumi.BoolPtrInput
	// Region of the probe.
	Region pulumi.StringPtrInput
	// The tenant ID of the probe.
	TenantId pulumi.IntPtrInput
}

func (ProbeState) ElementType() reflect.Type {
	return reflect.TypeOf((*probeState)(nil)).Elem()
}

type probeArgs struct {
	// Disables scripted checks for this probe. Defaults to `false`.
	DisableScriptedChecks *bool `pulumi:"disableScriptedChecks"`
	// Custom labels to be included with collected metrics and logs.
	Labels map[string]string `pulumi:"labels"`
	// Latitude coordinates.
	Latitude float64 `pulumi:"latitude"`
	// Longitude coordinates.
	Longitude float64 `pulumi:"longitude"`
	// Name of the probe.
	Name *string `pulumi:"name"`
	// Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
	Public *bool `pulumi:"public"`
	// Region of the probe.
	Region string `pulumi:"region"`
}

// The set of arguments for constructing a Probe resource.
type ProbeArgs struct {
	// Disables scripted checks for this probe. Defaults to `false`.
	DisableScriptedChecks pulumi.BoolPtrInput
	// Custom labels to be included with collected metrics and logs.
	Labels pulumi.StringMapInput
	// Latitude coordinates.
	Latitude pulumi.Float64Input
	// Longitude coordinates.
	Longitude pulumi.Float64Input
	// Name of the probe.
	Name pulumi.StringPtrInput
	// Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
	Public pulumi.BoolPtrInput
	// Region of the probe.
	Region pulumi.StringInput
}

func (ProbeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*probeArgs)(nil)).Elem()
}

type ProbeInput interface {
	pulumi.Input

	ToProbeOutput() ProbeOutput
	ToProbeOutputWithContext(ctx context.Context) ProbeOutput
}

func (*Probe) ElementType() reflect.Type {
	return reflect.TypeOf((**Probe)(nil)).Elem()
}

func (i *Probe) ToProbeOutput() ProbeOutput {
	return i.ToProbeOutputWithContext(context.Background())
}

func (i *Probe) ToProbeOutputWithContext(ctx context.Context) ProbeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProbeOutput)
}

// ProbeArrayInput is an input type that accepts ProbeArray and ProbeArrayOutput values.
// You can construct a concrete instance of `ProbeArrayInput` via:
//
//	ProbeArray{ ProbeArgs{...} }
type ProbeArrayInput interface {
	pulumi.Input

	ToProbeArrayOutput() ProbeArrayOutput
	ToProbeArrayOutputWithContext(context.Context) ProbeArrayOutput
}

type ProbeArray []ProbeInput

func (ProbeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Probe)(nil)).Elem()
}

func (i ProbeArray) ToProbeArrayOutput() ProbeArrayOutput {
	return i.ToProbeArrayOutputWithContext(context.Background())
}

func (i ProbeArray) ToProbeArrayOutputWithContext(ctx context.Context) ProbeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProbeArrayOutput)
}

// ProbeMapInput is an input type that accepts ProbeMap and ProbeMapOutput values.
// You can construct a concrete instance of `ProbeMapInput` via:
//
//	ProbeMap{ "key": ProbeArgs{...} }
type ProbeMapInput interface {
	pulumi.Input

	ToProbeMapOutput() ProbeMapOutput
	ToProbeMapOutputWithContext(context.Context) ProbeMapOutput
}

type ProbeMap map[string]ProbeInput

func (ProbeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Probe)(nil)).Elem()
}

func (i ProbeMap) ToProbeMapOutput() ProbeMapOutput {
	return i.ToProbeMapOutputWithContext(context.Background())
}

func (i ProbeMap) ToProbeMapOutputWithContext(ctx context.Context) ProbeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProbeMapOutput)
}

type ProbeOutput struct{ *pulumi.OutputState }

func (ProbeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Probe)(nil)).Elem()
}

func (o ProbeOutput) ToProbeOutput() ProbeOutput {
	return o
}

func (o ProbeOutput) ToProbeOutputWithContext(ctx context.Context) ProbeOutput {
	return o
}

// The probe authentication token. Your probe must use this to authenticate with Grafana Cloud.
func (o ProbeOutput) AuthToken() pulumi.StringOutput {
	return o.ApplyT(func(v *Probe) pulumi.StringOutput { return v.AuthToken }).(pulumi.StringOutput)
}

// Disables scripted checks for this probe. Defaults to `false`.
func (o ProbeOutput) DisableScriptedChecks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Probe) pulumi.BoolPtrOutput { return v.DisableScriptedChecks }).(pulumi.BoolPtrOutput)
}

// Custom labels to be included with collected metrics and logs.
func (o ProbeOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Probe) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Latitude coordinates.
func (o ProbeOutput) Latitude() pulumi.Float64Output {
	return o.ApplyT(func(v *Probe) pulumi.Float64Output { return v.Latitude }).(pulumi.Float64Output)
}

// Longitude coordinates.
func (o ProbeOutput) Longitude() pulumi.Float64Output {
	return o.ApplyT(func(v *Probe) pulumi.Float64Output { return v.Longitude }).(pulumi.Float64Output)
}

// Name of the probe.
func (o ProbeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Probe) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
func (o ProbeOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Probe) pulumi.BoolPtrOutput { return v.Public }).(pulumi.BoolPtrOutput)
}

// Region of the probe.
func (o ProbeOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Probe) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The tenant ID of the probe.
func (o ProbeOutput) TenantId() pulumi.IntOutput {
	return o.ApplyT(func(v *Probe) pulumi.IntOutput { return v.TenantId }).(pulumi.IntOutput)
}

type ProbeArrayOutput struct{ *pulumi.OutputState }

func (ProbeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Probe)(nil)).Elem()
}

func (o ProbeArrayOutput) ToProbeArrayOutput() ProbeArrayOutput {
	return o
}

func (o ProbeArrayOutput) ToProbeArrayOutputWithContext(ctx context.Context) ProbeArrayOutput {
	return o
}

func (o ProbeArrayOutput) Index(i pulumi.IntInput) ProbeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Probe {
		return vs[0].([]*Probe)[vs[1].(int)]
	}).(ProbeOutput)
}

type ProbeMapOutput struct{ *pulumi.OutputState }

func (ProbeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Probe)(nil)).Elem()
}

func (o ProbeMapOutput) ToProbeMapOutput() ProbeMapOutput {
	return o
}

func (o ProbeMapOutput) ToProbeMapOutputWithContext(ctx context.Context) ProbeMapOutput {
	return o
}

func (o ProbeMapOutput) MapIndex(k pulumi.StringInput) ProbeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Probe {
		return vs[0].(map[string]*Probe)[vs[1].(string)]
	}).(ProbeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProbeInput)(nil)).Elem(), &Probe{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProbeArrayInput)(nil)).Elem(), ProbeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProbeMapInput)(nil)).Elem(), ProbeMap{})
	pulumi.RegisterOutputType(ProbeOutput{})
	pulumi.RegisterOutputType(ProbeArrayOutput{})
	pulumi.RegisterOutputType(ProbeMapOutput{})
}
