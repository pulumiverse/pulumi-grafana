// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearning

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

var _ = internal.GetEnvOrDefault

type HolidayCustomPeriod struct {
	EndTime string `pulumi:"endTime"`
	// The name of the custom period.
	Name      *string `pulumi:"name"`
	StartTime string  `pulumi:"startTime"`
}

// HolidayCustomPeriodInput is an input type that accepts HolidayCustomPeriodArgs and HolidayCustomPeriodOutput values.
// You can construct a concrete instance of `HolidayCustomPeriodInput` via:
//
//	HolidayCustomPeriodArgs{...}
type HolidayCustomPeriodInput interface {
	pulumi.Input

	ToHolidayCustomPeriodOutput() HolidayCustomPeriodOutput
	ToHolidayCustomPeriodOutputWithContext(context.Context) HolidayCustomPeriodOutput
}

type HolidayCustomPeriodArgs struct {
	EndTime pulumi.StringInput `pulumi:"endTime"`
	// The name of the custom period.
	Name      pulumi.StringPtrInput `pulumi:"name"`
	StartTime pulumi.StringInput    `pulumi:"startTime"`
}

func (HolidayCustomPeriodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HolidayCustomPeriod)(nil)).Elem()
}

func (i HolidayCustomPeriodArgs) ToHolidayCustomPeriodOutput() HolidayCustomPeriodOutput {
	return i.ToHolidayCustomPeriodOutputWithContext(context.Background())
}

func (i HolidayCustomPeriodArgs) ToHolidayCustomPeriodOutputWithContext(ctx context.Context) HolidayCustomPeriodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HolidayCustomPeriodOutput)
}

// HolidayCustomPeriodArrayInput is an input type that accepts HolidayCustomPeriodArray and HolidayCustomPeriodArrayOutput values.
// You can construct a concrete instance of `HolidayCustomPeriodArrayInput` via:
//
//	HolidayCustomPeriodArray{ HolidayCustomPeriodArgs{...} }
type HolidayCustomPeriodArrayInput interface {
	pulumi.Input

	ToHolidayCustomPeriodArrayOutput() HolidayCustomPeriodArrayOutput
	ToHolidayCustomPeriodArrayOutputWithContext(context.Context) HolidayCustomPeriodArrayOutput
}

type HolidayCustomPeriodArray []HolidayCustomPeriodInput

func (HolidayCustomPeriodArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HolidayCustomPeriod)(nil)).Elem()
}

func (i HolidayCustomPeriodArray) ToHolidayCustomPeriodArrayOutput() HolidayCustomPeriodArrayOutput {
	return i.ToHolidayCustomPeriodArrayOutputWithContext(context.Background())
}

func (i HolidayCustomPeriodArray) ToHolidayCustomPeriodArrayOutputWithContext(ctx context.Context) HolidayCustomPeriodArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HolidayCustomPeriodArrayOutput)
}

type HolidayCustomPeriodOutput struct{ *pulumi.OutputState }

func (HolidayCustomPeriodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HolidayCustomPeriod)(nil)).Elem()
}

func (o HolidayCustomPeriodOutput) ToHolidayCustomPeriodOutput() HolidayCustomPeriodOutput {
	return o
}

func (o HolidayCustomPeriodOutput) ToHolidayCustomPeriodOutputWithContext(ctx context.Context) HolidayCustomPeriodOutput {
	return o
}

func (o HolidayCustomPeriodOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v HolidayCustomPeriod) string { return v.EndTime }).(pulumi.StringOutput)
}

// The name of the custom period.
func (o HolidayCustomPeriodOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HolidayCustomPeriod) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HolidayCustomPeriodOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v HolidayCustomPeriod) string { return v.StartTime }).(pulumi.StringOutput)
}

type HolidayCustomPeriodArrayOutput struct{ *pulumi.OutputState }

func (HolidayCustomPeriodArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HolidayCustomPeriod)(nil)).Elem()
}

func (o HolidayCustomPeriodArrayOutput) ToHolidayCustomPeriodArrayOutput() HolidayCustomPeriodArrayOutput {
	return o
}

func (o HolidayCustomPeriodArrayOutput) ToHolidayCustomPeriodArrayOutputWithContext(ctx context.Context) HolidayCustomPeriodArrayOutput {
	return o
}

func (o HolidayCustomPeriodArrayOutput) Index(i pulumi.IntInput) HolidayCustomPeriodOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HolidayCustomPeriod {
		return vs[0].([]HolidayCustomPeriod)[vs[1].(int)]
	}).(HolidayCustomPeriodOutput)
}

type OutlierDetectorAlgorithm struct {
	// For DBSCAN only, specify the configuration map
	Config *OutlierDetectorAlgorithmConfig `pulumi:"config"`
	// The name of the algorithm to use ('mad' or 'dbscan').
	Name string `pulumi:"name"`
	// Specify the sensitivity of the detector (in range [0,1]).
	Sensitivity float64 `pulumi:"sensitivity"`
}

// OutlierDetectorAlgorithmInput is an input type that accepts OutlierDetectorAlgorithmArgs and OutlierDetectorAlgorithmOutput values.
// You can construct a concrete instance of `OutlierDetectorAlgorithmInput` via:
//
//	OutlierDetectorAlgorithmArgs{...}
type OutlierDetectorAlgorithmInput interface {
	pulumi.Input

	ToOutlierDetectorAlgorithmOutput() OutlierDetectorAlgorithmOutput
	ToOutlierDetectorAlgorithmOutputWithContext(context.Context) OutlierDetectorAlgorithmOutput
}

type OutlierDetectorAlgorithmArgs struct {
	// For DBSCAN only, specify the configuration map
	Config OutlierDetectorAlgorithmConfigPtrInput `pulumi:"config"`
	// The name of the algorithm to use ('mad' or 'dbscan').
	Name pulumi.StringInput `pulumi:"name"`
	// Specify the sensitivity of the detector (in range [0,1]).
	Sensitivity pulumi.Float64Input `pulumi:"sensitivity"`
}

func (OutlierDetectorAlgorithmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OutlierDetectorAlgorithm)(nil)).Elem()
}

func (i OutlierDetectorAlgorithmArgs) ToOutlierDetectorAlgorithmOutput() OutlierDetectorAlgorithmOutput {
	return i.ToOutlierDetectorAlgorithmOutputWithContext(context.Background())
}

func (i OutlierDetectorAlgorithmArgs) ToOutlierDetectorAlgorithmOutputWithContext(ctx context.Context) OutlierDetectorAlgorithmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutlierDetectorAlgorithmOutput)
}

func (i OutlierDetectorAlgorithmArgs) ToOutlierDetectorAlgorithmPtrOutput() OutlierDetectorAlgorithmPtrOutput {
	return i.ToOutlierDetectorAlgorithmPtrOutputWithContext(context.Background())
}

func (i OutlierDetectorAlgorithmArgs) ToOutlierDetectorAlgorithmPtrOutputWithContext(ctx context.Context) OutlierDetectorAlgorithmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutlierDetectorAlgorithmOutput).ToOutlierDetectorAlgorithmPtrOutputWithContext(ctx)
}

// OutlierDetectorAlgorithmPtrInput is an input type that accepts OutlierDetectorAlgorithmArgs, OutlierDetectorAlgorithmPtr and OutlierDetectorAlgorithmPtrOutput values.
// You can construct a concrete instance of `OutlierDetectorAlgorithmPtrInput` via:
//
//	        OutlierDetectorAlgorithmArgs{...}
//
//	or:
//
//	        nil
type OutlierDetectorAlgorithmPtrInput interface {
	pulumi.Input

	ToOutlierDetectorAlgorithmPtrOutput() OutlierDetectorAlgorithmPtrOutput
	ToOutlierDetectorAlgorithmPtrOutputWithContext(context.Context) OutlierDetectorAlgorithmPtrOutput
}

type outlierDetectorAlgorithmPtrType OutlierDetectorAlgorithmArgs

func OutlierDetectorAlgorithmPtr(v *OutlierDetectorAlgorithmArgs) OutlierDetectorAlgorithmPtrInput {
	return (*outlierDetectorAlgorithmPtrType)(v)
}

func (*outlierDetectorAlgorithmPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OutlierDetectorAlgorithm)(nil)).Elem()
}

func (i *outlierDetectorAlgorithmPtrType) ToOutlierDetectorAlgorithmPtrOutput() OutlierDetectorAlgorithmPtrOutput {
	return i.ToOutlierDetectorAlgorithmPtrOutputWithContext(context.Background())
}

func (i *outlierDetectorAlgorithmPtrType) ToOutlierDetectorAlgorithmPtrOutputWithContext(ctx context.Context) OutlierDetectorAlgorithmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutlierDetectorAlgorithmPtrOutput)
}

type OutlierDetectorAlgorithmOutput struct{ *pulumi.OutputState }

func (OutlierDetectorAlgorithmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutlierDetectorAlgorithm)(nil)).Elem()
}

func (o OutlierDetectorAlgorithmOutput) ToOutlierDetectorAlgorithmOutput() OutlierDetectorAlgorithmOutput {
	return o
}

func (o OutlierDetectorAlgorithmOutput) ToOutlierDetectorAlgorithmOutputWithContext(ctx context.Context) OutlierDetectorAlgorithmOutput {
	return o
}

func (o OutlierDetectorAlgorithmOutput) ToOutlierDetectorAlgorithmPtrOutput() OutlierDetectorAlgorithmPtrOutput {
	return o.ToOutlierDetectorAlgorithmPtrOutputWithContext(context.Background())
}

func (o OutlierDetectorAlgorithmOutput) ToOutlierDetectorAlgorithmPtrOutputWithContext(ctx context.Context) OutlierDetectorAlgorithmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OutlierDetectorAlgorithm) *OutlierDetectorAlgorithm {
		return &v
	}).(OutlierDetectorAlgorithmPtrOutput)
}

// For DBSCAN only, specify the configuration map
func (o OutlierDetectorAlgorithmOutput) Config() OutlierDetectorAlgorithmConfigPtrOutput {
	return o.ApplyT(func(v OutlierDetectorAlgorithm) *OutlierDetectorAlgorithmConfig { return v.Config }).(OutlierDetectorAlgorithmConfigPtrOutput)
}

// The name of the algorithm to use ('mad' or 'dbscan').
func (o OutlierDetectorAlgorithmOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v OutlierDetectorAlgorithm) string { return v.Name }).(pulumi.StringOutput)
}

// Specify the sensitivity of the detector (in range [0,1]).
func (o OutlierDetectorAlgorithmOutput) Sensitivity() pulumi.Float64Output {
	return o.ApplyT(func(v OutlierDetectorAlgorithm) float64 { return v.Sensitivity }).(pulumi.Float64Output)
}

type OutlierDetectorAlgorithmPtrOutput struct{ *pulumi.OutputState }

func (OutlierDetectorAlgorithmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutlierDetectorAlgorithm)(nil)).Elem()
}

func (o OutlierDetectorAlgorithmPtrOutput) ToOutlierDetectorAlgorithmPtrOutput() OutlierDetectorAlgorithmPtrOutput {
	return o
}

func (o OutlierDetectorAlgorithmPtrOutput) ToOutlierDetectorAlgorithmPtrOutputWithContext(ctx context.Context) OutlierDetectorAlgorithmPtrOutput {
	return o
}

func (o OutlierDetectorAlgorithmPtrOutput) Elem() OutlierDetectorAlgorithmOutput {
	return o.ApplyT(func(v *OutlierDetectorAlgorithm) OutlierDetectorAlgorithm {
		if v != nil {
			return *v
		}
		var ret OutlierDetectorAlgorithm
		return ret
	}).(OutlierDetectorAlgorithmOutput)
}

// For DBSCAN only, specify the configuration map
func (o OutlierDetectorAlgorithmPtrOutput) Config() OutlierDetectorAlgorithmConfigPtrOutput {
	return o.ApplyT(func(v *OutlierDetectorAlgorithm) *OutlierDetectorAlgorithmConfig {
		if v == nil {
			return nil
		}
		return v.Config
	}).(OutlierDetectorAlgorithmConfigPtrOutput)
}

// The name of the algorithm to use ('mad' or 'dbscan').
func (o OutlierDetectorAlgorithmPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OutlierDetectorAlgorithm) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// Specify the sensitivity of the detector (in range [0,1]).
func (o OutlierDetectorAlgorithmPtrOutput) Sensitivity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *OutlierDetectorAlgorithm) *float64 {
		if v == nil {
			return nil
		}
		return &v.Sensitivity
	}).(pulumi.Float64PtrOutput)
}

type OutlierDetectorAlgorithmConfig struct {
	// Specify the epsilon parameter (positive float)
	Epsilon float64 `pulumi:"epsilon"`
}

// OutlierDetectorAlgorithmConfigInput is an input type that accepts OutlierDetectorAlgorithmConfigArgs and OutlierDetectorAlgorithmConfigOutput values.
// You can construct a concrete instance of `OutlierDetectorAlgorithmConfigInput` via:
//
//	OutlierDetectorAlgorithmConfigArgs{...}
type OutlierDetectorAlgorithmConfigInput interface {
	pulumi.Input

	ToOutlierDetectorAlgorithmConfigOutput() OutlierDetectorAlgorithmConfigOutput
	ToOutlierDetectorAlgorithmConfigOutputWithContext(context.Context) OutlierDetectorAlgorithmConfigOutput
}

type OutlierDetectorAlgorithmConfigArgs struct {
	// Specify the epsilon parameter (positive float)
	Epsilon pulumi.Float64Input `pulumi:"epsilon"`
}

func (OutlierDetectorAlgorithmConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OutlierDetectorAlgorithmConfig)(nil)).Elem()
}

func (i OutlierDetectorAlgorithmConfigArgs) ToOutlierDetectorAlgorithmConfigOutput() OutlierDetectorAlgorithmConfigOutput {
	return i.ToOutlierDetectorAlgorithmConfigOutputWithContext(context.Background())
}

func (i OutlierDetectorAlgorithmConfigArgs) ToOutlierDetectorAlgorithmConfigOutputWithContext(ctx context.Context) OutlierDetectorAlgorithmConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutlierDetectorAlgorithmConfigOutput)
}

func (i OutlierDetectorAlgorithmConfigArgs) ToOutlierDetectorAlgorithmConfigPtrOutput() OutlierDetectorAlgorithmConfigPtrOutput {
	return i.ToOutlierDetectorAlgorithmConfigPtrOutputWithContext(context.Background())
}

func (i OutlierDetectorAlgorithmConfigArgs) ToOutlierDetectorAlgorithmConfigPtrOutputWithContext(ctx context.Context) OutlierDetectorAlgorithmConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutlierDetectorAlgorithmConfigOutput).ToOutlierDetectorAlgorithmConfigPtrOutputWithContext(ctx)
}

// OutlierDetectorAlgorithmConfigPtrInput is an input type that accepts OutlierDetectorAlgorithmConfigArgs, OutlierDetectorAlgorithmConfigPtr and OutlierDetectorAlgorithmConfigPtrOutput values.
// You can construct a concrete instance of `OutlierDetectorAlgorithmConfigPtrInput` via:
//
//	        OutlierDetectorAlgorithmConfigArgs{...}
//
//	or:
//
//	        nil
type OutlierDetectorAlgorithmConfigPtrInput interface {
	pulumi.Input

	ToOutlierDetectorAlgorithmConfigPtrOutput() OutlierDetectorAlgorithmConfigPtrOutput
	ToOutlierDetectorAlgorithmConfigPtrOutputWithContext(context.Context) OutlierDetectorAlgorithmConfigPtrOutput
}

type outlierDetectorAlgorithmConfigPtrType OutlierDetectorAlgorithmConfigArgs

func OutlierDetectorAlgorithmConfigPtr(v *OutlierDetectorAlgorithmConfigArgs) OutlierDetectorAlgorithmConfigPtrInput {
	return (*outlierDetectorAlgorithmConfigPtrType)(v)
}

func (*outlierDetectorAlgorithmConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OutlierDetectorAlgorithmConfig)(nil)).Elem()
}

func (i *outlierDetectorAlgorithmConfigPtrType) ToOutlierDetectorAlgorithmConfigPtrOutput() OutlierDetectorAlgorithmConfigPtrOutput {
	return i.ToOutlierDetectorAlgorithmConfigPtrOutputWithContext(context.Background())
}

func (i *outlierDetectorAlgorithmConfigPtrType) ToOutlierDetectorAlgorithmConfigPtrOutputWithContext(ctx context.Context) OutlierDetectorAlgorithmConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutlierDetectorAlgorithmConfigPtrOutput)
}

type OutlierDetectorAlgorithmConfigOutput struct{ *pulumi.OutputState }

func (OutlierDetectorAlgorithmConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutlierDetectorAlgorithmConfig)(nil)).Elem()
}

func (o OutlierDetectorAlgorithmConfigOutput) ToOutlierDetectorAlgorithmConfigOutput() OutlierDetectorAlgorithmConfigOutput {
	return o
}

func (o OutlierDetectorAlgorithmConfigOutput) ToOutlierDetectorAlgorithmConfigOutputWithContext(ctx context.Context) OutlierDetectorAlgorithmConfigOutput {
	return o
}

func (o OutlierDetectorAlgorithmConfigOutput) ToOutlierDetectorAlgorithmConfigPtrOutput() OutlierDetectorAlgorithmConfigPtrOutput {
	return o.ToOutlierDetectorAlgorithmConfigPtrOutputWithContext(context.Background())
}

func (o OutlierDetectorAlgorithmConfigOutput) ToOutlierDetectorAlgorithmConfigPtrOutputWithContext(ctx context.Context) OutlierDetectorAlgorithmConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OutlierDetectorAlgorithmConfig) *OutlierDetectorAlgorithmConfig {
		return &v
	}).(OutlierDetectorAlgorithmConfigPtrOutput)
}

// Specify the epsilon parameter (positive float)
func (o OutlierDetectorAlgorithmConfigOutput) Epsilon() pulumi.Float64Output {
	return o.ApplyT(func(v OutlierDetectorAlgorithmConfig) float64 { return v.Epsilon }).(pulumi.Float64Output)
}

type OutlierDetectorAlgorithmConfigPtrOutput struct{ *pulumi.OutputState }

func (OutlierDetectorAlgorithmConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutlierDetectorAlgorithmConfig)(nil)).Elem()
}

func (o OutlierDetectorAlgorithmConfigPtrOutput) ToOutlierDetectorAlgorithmConfigPtrOutput() OutlierDetectorAlgorithmConfigPtrOutput {
	return o
}

func (o OutlierDetectorAlgorithmConfigPtrOutput) ToOutlierDetectorAlgorithmConfigPtrOutputWithContext(ctx context.Context) OutlierDetectorAlgorithmConfigPtrOutput {
	return o
}

func (o OutlierDetectorAlgorithmConfigPtrOutput) Elem() OutlierDetectorAlgorithmConfigOutput {
	return o.ApplyT(func(v *OutlierDetectorAlgorithmConfig) OutlierDetectorAlgorithmConfig {
		if v != nil {
			return *v
		}
		var ret OutlierDetectorAlgorithmConfig
		return ret
	}).(OutlierDetectorAlgorithmConfigOutput)
}

// Specify the epsilon parameter (positive float)
func (o OutlierDetectorAlgorithmConfigPtrOutput) Epsilon() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *OutlierDetectorAlgorithmConfig) *float64 {
		if v == nil {
			return nil
		}
		return &v.Epsilon
	}).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HolidayCustomPeriodInput)(nil)).Elem(), HolidayCustomPeriodArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HolidayCustomPeriodArrayInput)(nil)).Elem(), HolidayCustomPeriodArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutlierDetectorAlgorithmInput)(nil)).Elem(), OutlierDetectorAlgorithmArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutlierDetectorAlgorithmPtrInput)(nil)).Elem(), OutlierDetectorAlgorithmArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutlierDetectorAlgorithmConfigInput)(nil)).Elem(), OutlierDetectorAlgorithmConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutlierDetectorAlgorithmConfigPtrInput)(nil)).Elem(), OutlierDetectorAlgorithmConfigArgs{})
	pulumi.RegisterOutputType(HolidayCustomPeriodOutput{})
	pulumi.RegisterOutputType(HolidayCustomPeriodArrayOutput{})
	pulumi.RegisterOutputType(OutlierDetectorAlgorithmOutput{})
	pulumi.RegisterOutputType(OutlierDetectorAlgorithmPtrOutput{})
	pulumi.RegisterOutputType(OutlierDetectorAlgorithmConfigOutput{})
	pulumi.RegisterOutputType(OutlierDetectorAlgorithmConfigPtrOutput{})
}
