// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// **Note:** This resource is available only with Grafana Enterprise 7.+.
//
// * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/create-reports/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/reporting/)
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
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := oss.NewDashboard(ctx, "test", &oss.DashboardArgs{
//				ConfigJson: pulumi.String("{\n  \"uid\": \"report-dashboard\",\n  \"title\": \"report-dashboard\"\n}\n"),
//				Message:    pulumi.String("inital commit."),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = enterprise.NewReport(ctx, "test", &enterprise.ReportArgs{
//				Name: pulumi.String("my report"),
//				Recipients: pulumi.StringArray{
//					pulumi.String("some@email.com"),
//				},
//				Dashboards: enterprise.ReportDashboardArray{
//					&enterprise.ReportDashboardArgs{
//						Uid: test.Uid,
//					},
//				},
//				Schedule: &enterprise.ReportScheduleArgs{
//					Frequency: pulumi.String("hourly"),
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
// $ pulumi import grafana:index/report:Report name "{{ id }}"
// ```
//
// ```sh
// $ pulumi import grafana:index/report:Report name "{{ orgID }}:{{ id }}"
// ```
//
// Deprecated: grafana.index/report.Report has been deprecated in favor of grafana.enterprise/report.Report
type Report struct {
	pulumi.CustomResourceState

	// List of dashboards to render into the report
	Dashboards ReportDashboardArrayOutput `pulumi:"dashboards"`
	// Specifies what kind of attachment to generate for the report. Allowed values: `pdf`, `csv`, `image`.
	Formats pulumi.StringArrayOutput `pulumi:"formats"`
	// Whether to include a link to the dashboard in the report. Defaults to `true`.
	IncludeDashboardLink pulumi.BoolPtrOutput `pulumi:"includeDashboardLink"`
	// Whether to include a CSV file of table panel data. Defaults to `false`.
	IncludeTableCsv pulumi.BoolPtrOutput `pulumi:"includeTableCsv"`
	// Layout of the report. Allowed values: `simple`, `grid`. Defaults to `grid`.
	Layout pulumi.StringPtrOutput `pulumi:"layout"`
	// Message to be sent in the report.
	Message pulumi.StringPtrOutput `pulumi:"message"`
	// Name of the report.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// Orientation of the report. Allowed values: `landscape`, `portrait`. Defaults to `landscape`.
	Orientation pulumi.StringPtrOutput `pulumi:"orientation"`
	// List of recipients of the report.
	Recipients pulumi.StringArrayOutput `pulumi:"recipients"`
	// Reply-to email address of the report.
	ReplyTo pulumi.StringPtrOutput `pulumi:"replyTo"`
	// Schedule of the report.
	Schedule ReportScheduleOutput `pulumi:"schedule"`
}

// NewReport registers a new resource with the given unique name, arguments, and options.
func NewReport(ctx *pulumi.Context,
	name string, args *ReportArgs, opts ...pulumi.ResourceOption) (*Report, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Recipients == nil {
		return nil, errors.New("invalid value for required argument 'Recipients'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/report:Report"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Report
	err := ctx.RegisterResource("grafana:index/report:Report", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReport gets an existing Report resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReportState, opts ...pulumi.ResourceOption) (*Report, error) {
	var resource Report
	err := ctx.ReadResource("grafana:index/report:Report", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Report resources.
type reportState struct {
	// List of dashboards to render into the report
	Dashboards []ReportDashboard `pulumi:"dashboards"`
	// Specifies what kind of attachment to generate for the report. Allowed values: `pdf`, `csv`, `image`.
	Formats []string `pulumi:"formats"`
	// Whether to include a link to the dashboard in the report. Defaults to `true`.
	IncludeDashboardLink *bool `pulumi:"includeDashboardLink"`
	// Whether to include a CSV file of table panel data. Defaults to `false`.
	IncludeTableCsv *bool `pulumi:"includeTableCsv"`
	// Layout of the report. Allowed values: `simple`, `grid`. Defaults to `grid`.
	Layout *string `pulumi:"layout"`
	// Message to be sent in the report.
	Message *string `pulumi:"message"`
	// Name of the report.
	Name *string `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// Orientation of the report. Allowed values: `landscape`, `portrait`. Defaults to `landscape`.
	Orientation *string `pulumi:"orientation"`
	// List of recipients of the report.
	Recipients []string `pulumi:"recipients"`
	// Reply-to email address of the report.
	ReplyTo *string `pulumi:"replyTo"`
	// Schedule of the report.
	Schedule *ReportSchedule `pulumi:"schedule"`
}

type ReportState struct {
	// List of dashboards to render into the report
	Dashboards ReportDashboardArrayInput
	// Specifies what kind of attachment to generate for the report. Allowed values: `pdf`, `csv`, `image`.
	Formats pulumi.StringArrayInput
	// Whether to include a link to the dashboard in the report. Defaults to `true`.
	IncludeDashboardLink pulumi.BoolPtrInput
	// Whether to include a CSV file of table panel data. Defaults to `false`.
	IncludeTableCsv pulumi.BoolPtrInput
	// Layout of the report. Allowed values: `simple`, `grid`. Defaults to `grid`.
	Layout pulumi.StringPtrInput
	// Message to be sent in the report.
	Message pulumi.StringPtrInput
	// Name of the report.
	Name pulumi.StringPtrInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// Orientation of the report. Allowed values: `landscape`, `portrait`. Defaults to `landscape`.
	Orientation pulumi.StringPtrInput
	// List of recipients of the report.
	Recipients pulumi.StringArrayInput
	// Reply-to email address of the report.
	ReplyTo pulumi.StringPtrInput
	// Schedule of the report.
	Schedule ReportSchedulePtrInput
}

func (ReportState) ElementType() reflect.Type {
	return reflect.TypeOf((*reportState)(nil)).Elem()
}

type reportArgs struct {
	// List of dashboards to render into the report
	Dashboards []ReportDashboard `pulumi:"dashboards"`
	// Specifies what kind of attachment to generate for the report. Allowed values: `pdf`, `csv`, `image`.
	Formats []string `pulumi:"formats"`
	// Whether to include a link to the dashboard in the report. Defaults to `true`.
	IncludeDashboardLink *bool `pulumi:"includeDashboardLink"`
	// Whether to include a CSV file of table panel data. Defaults to `false`.
	IncludeTableCsv *bool `pulumi:"includeTableCsv"`
	// Layout of the report. Allowed values: `simple`, `grid`. Defaults to `grid`.
	Layout *string `pulumi:"layout"`
	// Message to be sent in the report.
	Message *string `pulumi:"message"`
	// Name of the report.
	Name *string `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// Orientation of the report. Allowed values: `landscape`, `portrait`. Defaults to `landscape`.
	Orientation *string `pulumi:"orientation"`
	// List of recipients of the report.
	Recipients []string `pulumi:"recipients"`
	// Reply-to email address of the report.
	ReplyTo *string `pulumi:"replyTo"`
	// Schedule of the report.
	Schedule ReportSchedule `pulumi:"schedule"`
}

// The set of arguments for constructing a Report resource.
type ReportArgs struct {
	// List of dashboards to render into the report
	Dashboards ReportDashboardArrayInput
	// Specifies what kind of attachment to generate for the report. Allowed values: `pdf`, `csv`, `image`.
	Formats pulumi.StringArrayInput
	// Whether to include a link to the dashboard in the report. Defaults to `true`.
	IncludeDashboardLink pulumi.BoolPtrInput
	// Whether to include a CSV file of table panel data. Defaults to `false`.
	IncludeTableCsv pulumi.BoolPtrInput
	// Layout of the report. Allowed values: `simple`, `grid`. Defaults to `grid`.
	Layout pulumi.StringPtrInput
	// Message to be sent in the report.
	Message pulumi.StringPtrInput
	// Name of the report.
	Name pulumi.StringPtrInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// Orientation of the report. Allowed values: `landscape`, `portrait`. Defaults to `landscape`.
	Orientation pulumi.StringPtrInput
	// List of recipients of the report.
	Recipients pulumi.StringArrayInput
	// Reply-to email address of the report.
	ReplyTo pulumi.StringPtrInput
	// Schedule of the report.
	Schedule ReportScheduleInput
}

func (ReportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reportArgs)(nil)).Elem()
}

type ReportInput interface {
	pulumi.Input

	ToReportOutput() ReportOutput
	ToReportOutputWithContext(ctx context.Context) ReportOutput
}

func (*Report) ElementType() reflect.Type {
	return reflect.TypeOf((**Report)(nil)).Elem()
}

func (i *Report) ToReportOutput() ReportOutput {
	return i.ToReportOutputWithContext(context.Background())
}

func (i *Report) ToReportOutputWithContext(ctx context.Context) ReportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportOutput)
}

// ReportArrayInput is an input type that accepts ReportArray and ReportArrayOutput values.
// You can construct a concrete instance of `ReportArrayInput` via:
//
//	ReportArray{ ReportArgs{...} }
type ReportArrayInput interface {
	pulumi.Input

	ToReportArrayOutput() ReportArrayOutput
	ToReportArrayOutputWithContext(context.Context) ReportArrayOutput
}

type ReportArray []ReportInput

func (ReportArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Report)(nil)).Elem()
}

func (i ReportArray) ToReportArrayOutput() ReportArrayOutput {
	return i.ToReportArrayOutputWithContext(context.Background())
}

func (i ReportArray) ToReportArrayOutputWithContext(ctx context.Context) ReportArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportArrayOutput)
}

// ReportMapInput is an input type that accepts ReportMap and ReportMapOutput values.
// You can construct a concrete instance of `ReportMapInput` via:
//
//	ReportMap{ "key": ReportArgs{...} }
type ReportMapInput interface {
	pulumi.Input

	ToReportMapOutput() ReportMapOutput
	ToReportMapOutputWithContext(context.Context) ReportMapOutput
}

type ReportMap map[string]ReportInput

func (ReportMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Report)(nil)).Elem()
}

func (i ReportMap) ToReportMapOutput() ReportMapOutput {
	return i.ToReportMapOutputWithContext(context.Background())
}

func (i ReportMap) ToReportMapOutputWithContext(ctx context.Context) ReportMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportMapOutput)
}

type ReportOutput struct{ *pulumi.OutputState }

func (ReportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Report)(nil)).Elem()
}

func (o ReportOutput) ToReportOutput() ReportOutput {
	return o
}

func (o ReportOutput) ToReportOutputWithContext(ctx context.Context) ReportOutput {
	return o
}

// List of dashboards to render into the report
func (o ReportOutput) Dashboards() ReportDashboardArrayOutput {
	return o.ApplyT(func(v *Report) ReportDashboardArrayOutput { return v.Dashboards }).(ReportDashboardArrayOutput)
}

// Specifies what kind of attachment to generate for the report. Allowed values: `pdf`, `csv`, `image`.
func (o ReportOutput) Formats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Report) pulumi.StringArrayOutput { return v.Formats }).(pulumi.StringArrayOutput)
}

// Whether to include a link to the dashboard in the report. Defaults to `true`.
func (o ReportOutput) IncludeDashboardLink() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Report) pulumi.BoolPtrOutput { return v.IncludeDashboardLink }).(pulumi.BoolPtrOutput)
}

// Whether to include a CSV file of table panel data. Defaults to `false`.
func (o ReportOutput) IncludeTableCsv() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Report) pulumi.BoolPtrOutput { return v.IncludeTableCsv }).(pulumi.BoolPtrOutput)
}

// Layout of the report. Allowed values: `simple`, `grid`. Defaults to `grid`.
func (o ReportOutput) Layout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Report) pulumi.StringPtrOutput { return v.Layout }).(pulumi.StringPtrOutput)
}

// Message to be sent in the report.
func (o ReportOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Report) pulumi.StringPtrOutput { return v.Message }).(pulumi.StringPtrOutput)
}

// Name of the report.
func (o ReportOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Report) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o ReportOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Report) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Orientation of the report. Allowed values: `landscape`, `portrait`. Defaults to `landscape`.
func (o ReportOutput) Orientation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Report) pulumi.StringPtrOutput { return v.Orientation }).(pulumi.StringPtrOutput)
}

// List of recipients of the report.
func (o ReportOutput) Recipients() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Report) pulumi.StringArrayOutput { return v.Recipients }).(pulumi.StringArrayOutput)
}

// Reply-to email address of the report.
func (o ReportOutput) ReplyTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Report) pulumi.StringPtrOutput { return v.ReplyTo }).(pulumi.StringPtrOutput)
}

// Schedule of the report.
func (o ReportOutput) Schedule() ReportScheduleOutput {
	return o.ApplyT(func(v *Report) ReportScheduleOutput { return v.Schedule }).(ReportScheduleOutput)
}

type ReportArrayOutput struct{ *pulumi.OutputState }

func (ReportArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Report)(nil)).Elem()
}

func (o ReportArrayOutput) ToReportArrayOutput() ReportArrayOutput {
	return o
}

func (o ReportArrayOutput) ToReportArrayOutputWithContext(ctx context.Context) ReportArrayOutput {
	return o
}

func (o ReportArrayOutput) Index(i pulumi.IntInput) ReportOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Report {
		return vs[0].([]*Report)[vs[1].(int)]
	}).(ReportOutput)
}

type ReportMapOutput struct{ *pulumi.OutputState }

func (ReportMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Report)(nil)).Elem()
}

func (o ReportMapOutput) ToReportMapOutput() ReportMapOutput {
	return o
}

func (o ReportMapOutput) ToReportMapOutputWithContext(ctx context.Context) ReportMapOutput {
	return o
}

func (o ReportMapOutput) MapIndex(k pulumi.StringInput) ReportOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Report {
		return vs[0].(map[string]*Report)[vs[1].(string)]
	}).(ReportOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReportInput)(nil)).Elem(), &Report{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReportArrayInput)(nil)).Elem(), ReportArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReportMapInput)(nil)).Elem(), ReportMap{})
	pulumi.RegisterOutputType(ReportOutput{})
	pulumi.RegisterOutputType(ReportArrayOutput{})
	pulumi.RegisterOutputType(ReportMapOutput{})
}
