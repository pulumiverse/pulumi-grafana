// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearning

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// A holiday describes time periods where a time series is expected to behave differently to normal.
//
// To use a holiday in a job, use its id in the `holidays` attribute of a `machineLearning.Job`:
//
// ### iCal Holiday
//
// This holiday uses an iCal file to define the holidays.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/machinelearning"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := machinelearning.NewHoliday(ctx, "ical", &machinelearning.HolidayArgs{
//				Name:         pulumi.String("My iCal holiday"),
//				Description:  pulumi.String("My Holiday"),
//				IcalUrl:      pulumi.String("https://calendar.google.com/calendar/ical/en.uk%23holiday%40group.v.calendar.google.com/public/basic.ics"),
//				IcalTimezone: pulumi.String("Europe/London"),
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
// ### Custom Periods Holiday
//
// This holiday uses custom periods to define the holidays.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/machinelearning"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := machinelearning.NewHoliday(ctx, "custom_periods", &machinelearning.HolidayArgs{
//				Name:        pulumi.String("My custom periods holiday"),
//				Description: pulumi.String("My Holiday"),
//				CustomPeriods: machinelearning.HolidayCustomPeriodArray{
//					&machinelearning.HolidayCustomPeriodArgs{
//						Name:      pulumi.String("First of January"),
//						StartTime: pulumi.String("2023-01-01T00:00:00Z"),
//						EndTime:   pulumi.String("2023-01-02T00:00:00Z"),
//					},
//					&machinelearning.HolidayCustomPeriodArgs{
//						Name:      pulumi.String("First of Feburary"),
//						StartTime: pulumi.String("2023-02-01T00:00:00Z"),
//						EndTime:   pulumi.String("2023-02-02T00:00:00Z"),
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
// $ pulumi import grafana:machineLearning/holiday:Holiday name "{{ id }}"
// ```
type Holiday struct {
	pulumi.CustomResourceState

	// A list of custom periods for the holiday.
	CustomPeriods HolidayCustomPeriodArrayOutput `pulumi:"customPeriods"`
	// A description of the holiday.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The timezone to use for events in the iCal file pointed to by ical_url.
	IcalTimezone pulumi.StringPtrOutput `pulumi:"icalTimezone"`
	// A URL to an iCal file containing all occurrences of the holiday.
	IcalUrl pulumi.StringPtrOutput `pulumi:"icalUrl"`
	// The name of the holiday.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewHoliday registers a new resource with the given unique name, arguments, and options.
func NewHoliday(ctx *pulumi.Context,
	name string, args *HolidayArgs, opts ...pulumi.ResourceOption) (*Holiday, error) {
	if args == nil {
		args = &HolidayArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Holiday
	err := ctx.RegisterResource("grafana:machineLearning/holiday:Holiday", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHoliday gets an existing Holiday resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHoliday(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HolidayState, opts ...pulumi.ResourceOption) (*Holiday, error) {
	var resource Holiday
	err := ctx.ReadResource("grafana:machineLearning/holiday:Holiday", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Holiday resources.
type holidayState struct {
	// A list of custom periods for the holiday.
	CustomPeriods []HolidayCustomPeriod `pulumi:"customPeriods"`
	// A description of the holiday.
	Description *string `pulumi:"description"`
	// The timezone to use for events in the iCal file pointed to by ical_url.
	IcalTimezone *string `pulumi:"icalTimezone"`
	// A URL to an iCal file containing all occurrences of the holiday.
	IcalUrl *string `pulumi:"icalUrl"`
	// The name of the holiday.
	Name *string `pulumi:"name"`
}

type HolidayState struct {
	// A list of custom periods for the holiday.
	CustomPeriods HolidayCustomPeriodArrayInput
	// A description of the holiday.
	Description pulumi.StringPtrInput
	// The timezone to use for events in the iCal file pointed to by ical_url.
	IcalTimezone pulumi.StringPtrInput
	// A URL to an iCal file containing all occurrences of the holiday.
	IcalUrl pulumi.StringPtrInput
	// The name of the holiday.
	Name pulumi.StringPtrInput
}

func (HolidayState) ElementType() reflect.Type {
	return reflect.TypeOf((*holidayState)(nil)).Elem()
}

type holidayArgs struct {
	// A list of custom periods for the holiday.
	CustomPeriods []HolidayCustomPeriod `pulumi:"customPeriods"`
	// A description of the holiday.
	Description *string `pulumi:"description"`
	// The timezone to use for events in the iCal file pointed to by ical_url.
	IcalTimezone *string `pulumi:"icalTimezone"`
	// A URL to an iCal file containing all occurrences of the holiday.
	IcalUrl *string `pulumi:"icalUrl"`
	// The name of the holiday.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Holiday resource.
type HolidayArgs struct {
	// A list of custom periods for the holiday.
	CustomPeriods HolidayCustomPeriodArrayInput
	// A description of the holiday.
	Description pulumi.StringPtrInput
	// The timezone to use for events in the iCal file pointed to by ical_url.
	IcalTimezone pulumi.StringPtrInput
	// A URL to an iCal file containing all occurrences of the holiday.
	IcalUrl pulumi.StringPtrInput
	// The name of the holiday.
	Name pulumi.StringPtrInput
}

func (HolidayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*holidayArgs)(nil)).Elem()
}

type HolidayInput interface {
	pulumi.Input

	ToHolidayOutput() HolidayOutput
	ToHolidayOutputWithContext(ctx context.Context) HolidayOutput
}

func (*Holiday) ElementType() reflect.Type {
	return reflect.TypeOf((**Holiday)(nil)).Elem()
}

func (i *Holiday) ToHolidayOutput() HolidayOutput {
	return i.ToHolidayOutputWithContext(context.Background())
}

func (i *Holiday) ToHolidayOutputWithContext(ctx context.Context) HolidayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HolidayOutput)
}

// HolidayArrayInput is an input type that accepts HolidayArray and HolidayArrayOutput values.
// You can construct a concrete instance of `HolidayArrayInput` via:
//
//	HolidayArray{ HolidayArgs{...} }
type HolidayArrayInput interface {
	pulumi.Input

	ToHolidayArrayOutput() HolidayArrayOutput
	ToHolidayArrayOutputWithContext(context.Context) HolidayArrayOutput
}

type HolidayArray []HolidayInput

func (HolidayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Holiday)(nil)).Elem()
}

func (i HolidayArray) ToHolidayArrayOutput() HolidayArrayOutput {
	return i.ToHolidayArrayOutputWithContext(context.Background())
}

func (i HolidayArray) ToHolidayArrayOutputWithContext(ctx context.Context) HolidayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HolidayArrayOutput)
}

// HolidayMapInput is an input type that accepts HolidayMap and HolidayMapOutput values.
// You can construct a concrete instance of `HolidayMapInput` via:
//
//	HolidayMap{ "key": HolidayArgs{...} }
type HolidayMapInput interface {
	pulumi.Input

	ToHolidayMapOutput() HolidayMapOutput
	ToHolidayMapOutputWithContext(context.Context) HolidayMapOutput
}

type HolidayMap map[string]HolidayInput

func (HolidayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Holiday)(nil)).Elem()
}

func (i HolidayMap) ToHolidayMapOutput() HolidayMapOutput {
	return i.ToHolidayMapOutputWithContext(context.Background())
}

func (i HolidayMap) ToHolidayMapOutputWithContext(ctx context.Context) HolidayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HolidayMapOutput)
}

type HolidayOutput struct{ *pulumi.OutputState }

func (HolidayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Holiday)(nil)).Elem()
}

func (o HolidayOutput) ToHolidayOutput() HolidayOutput {
	return o
}

func (o HolidayOutput) ToHolidayOutputWithContext(ctx context.Context) HolidayOutput {
	return o
}

// A list of custom periods for the holiday.
func (o HolidayOutput) CustomPeriods() HolidayCustomPeriodArrayOutput {
	return o.ApplyT(func(v *Holiday) HolidayCustomPeriodArrayOutput { return v.CustomPeriods }).(HolidayCustomPeriodArrayOutput)
}

// A description of the holiday.
func (o HolidayOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Holiday) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The timezone to use for events in the iCal file pointed to by ical_url.
func (o HolidayOutput) IcalTimezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Holiday) pulumi.StringPtrOutput { return v.IcalTimezone }).(pulumi.StringPtrOutput)
}

// A URL to an iCal file containing all occurrences of the holiday.
func (o HolidayOutput) IcalUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Holiday) pulumi.StringPtrOutput { return v.IcalUrl }).(pulumi.StringPtrOutput)
}

// The name of the holiday.
func (o HolidayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Holiday) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type HolidayArrayOutput struct{ *pulumi.OutputState }

func (HolidayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Holiday)(nil)).Elem()
}

func (o HolidayArrayOutput) ToHolidayArrayOutput() HolidayArrayOutput {
	return o
}

func (o HolidayArrayOutput) ToHolidayArrayOutputWithContext(ctx context.Context) HolidayArrayOutput {
	return o
}

func (o HolidayArrayOutput) Index(i pulumi.IntInput) HolidayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Holiday {
		return vs[0].([]*Holiday)[vs[1].(int)]
	}).(HolidayOutput)
}

type HolidayMapOutput struct{ *pulumi.OutputState }

func (HolidayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Holiday)(nil)).Elem()
}

func (o HolidayMapOutput) ToHolidayMapOutput() HolidayMapOutput {
	return o
}

func (o HolidayMapOutput) ToHolidayMapOutputWithContext(ctx context.Context) HolidayMapOutput {
	return o
}

func (o HolidayMapOutput) MapIndex(k pulumi.StringInput) HolidayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Holiday {
		return vs[0].(map[string]*Holiday)[vs[1].(string)]
	}).(HolidayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HolidayInput)(nil)).Elem(), &Holiday{})
	pulumi.RegisterInputType(reflect.TypeOf((*HolidayArrayInput)(nil)).Elem(), HolidayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HolidayMapInput)(nil)).Elem(), HolidayMap{})
	pulumi.RegisterOutputType(HolidayOutput{})
	pulumi.RegisterOutputType(HolidayArrayOutput{})
	pulumi.RegisterOutputType(HolidayMapOutput{})
}
