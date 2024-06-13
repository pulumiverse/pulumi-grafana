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

// Resource manages Grafana SLOs.
//
// * [Official documentation](https://grafana.com/docs/grafana-cloud/alerting-and-irm/slo/)
// * [API documentation](https://grafana.com/docs/grafana-cloud/alerting-and-irm/slo/api/)
// * [Additional Information On Alerting Rule Annotations and Labels](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/#templating/)
//
// ## Example Usage
//
// ## Import
//
// ```sh
// $ pulumi import grafana:index/sLO:SLO name "{{ uuid }}"
// ```
type SLO struct {
	pulumi.CustomResourceState

	// Configures the alerting rules that will be generated for each
	// 			time window associated with the SLO. Grafana SLOs can generate
	// 			alerts when the short-term error budget burn is very high, the
	// 			long-term error budget burn rate is high, or when the remaining
	// 			error budget is below a certain threshold. Annotations and Labels support templating.
	Alertings SLOAlertingArrayOutput `pulumi:"alertings"`
	// Description is a free-text field that can provide more context to an SLO.
	Description pulumi.StringOutput `pulumi:"description"`
	// Destination Datasource sets the datasource defined for an SLO
	DestinationDatasource SLODestinationDatasourcePtrOutput `pulumi:"destinationDatasource"`
	// Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
	Labels SLOLabelArrayOutput `pulumi:"labels"`
	// Name should be a short description of your indicator. Consider names like "API Availability"
	Name pulumi.StringOutput `pulumi:"name"`
	// Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
	Objectives SLOObjectiveArrayOutput `pulumi:"objectives"`
	// Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
	Queries SLOQueryArrayOutput `pulumi:"queries"`
}

// NewSLO registers a new resource with the given unique name, arguments, and options.
func NewSLO(ctx *pulumi.Context,
	name string, args *SLOArgs, opts ...pulumi.ResourceOption) (*SLO, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Objectives == nil {
		return nil, errors.New("invalid value for required argument 'Objectives'")
	}
	if args.Queries == nil {
		return nil, errors.New("invalid value for required argument 'Queries'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SLO
	err := ctx.RegisterResource("grafana:index/sLO:SLO", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSLO gets an existing SLO resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSLO(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SLOState, opts ...pulumi.ResourceOption) (*SLO, error) {
	var resource SLO
	err := ctx.ReadResource("grafana:index/sLO:SLO", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SLO resources.
type sloState struct {
	// Configures the alerting rules that will be generated for each
	// 			time window associated with the SLO. Grafana SLOs can generate
	// 			alerts when the short-term error budget burn is very high, the
	// 			long-term error budget burn rate is high, or when the remaining
	// 			error budget is below a certain threshold. Annotations and Labels support templating.
	Alertings []SLOAlerting `pulumi:"alertings"`
	// Description is a free-text field that can provide more context to an SLO.
	Description *string `pulumi:"description"`
	// Destination Datasource sets the datasource defined for an SLO
	DestinationDatasource *SLODestinationDatasource `pulumi:"destinationDatasource"`
	// Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
	Labels []SLOLabel `pulumi:"labels"`
	// Name should be a short description of your indicator. Consider names like "API Availability"
	Name *string `pulumi:"name"`
	// Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
	Objectives []SLOObjective `pulumi:"objectives"`
	// Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
	Queries []SLOQuery `pulumi:"queries"`
}

type SLOState struct {
	// Configures the alerting rules that will be generated for each
	// 			time window associated with the SLO. Grafana SLOs can generate
	// 			alerts when the short-term error budget burn is very high, the
	// 			long-term error budget burn rate is high, or when the remaining
	// 			error budget is below a certain threshold. Annotations and Labels support templating.
	Alertings SLOAlertingArrayInput
	// Description is a free-text field that can provide more context to an SLO.
	Description pulumi.StringPtrInput
	// Destination Datasource sets the datasource defined for an SLO
	DestinationDatasource SLODestinationDatasourcePtrInput
	// Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
	Labels SLOLabelArrayInput
	// Name should be a short description of your indicator. Consider names like "API Availability"
	Name pulumi.StringPtrInput
	// Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
	Objectives SLOObjectiveArrayInput
	// Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
	Queries SLOQueryArrayInput
}

func (SLOState) ElementType() reflect.Type {
	return reflect.TypeOf((*sloState)(nil)).Elem()
}

type sloArgs struct {
	// Configures the alerting rules that will be generated for each
	// 			time window associated with the SLO. Grafana SLOs can generate
	// 			alerts when the short-term error budget burn is very high, the
	// 			long-term error budget burn rate is high, or when the remaining
	// 			error budget is below a certain threshold. Annotations and Labels support templating.
	Alertings []SLOAlerting `pulumi:"alertings"`
	// Description is a free-text field that can provide more context to an SLO.
	Description string `pulumi:"description"`
	// Destination Datasource sets the datasource defined for an SLO
	DestinationDatasource *SLODestinationDatasource `pulumi:"destinationDatasource"`
	// Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
	Labels []SLOLabel `pulumi:"labels"`
	// Name should be a short description of your indicator. Consider names like "API Availability"
	Name *string `pulumi:"name"`
	// Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
	Objectives []SLOObjective `pulumi:"objectives"`
	// Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
	Queries []SLOQuery `pulumi:"queries"`
}

// The set of arguments for constructing a SLO resource.
type SLOArgs struct {
	// Configures the alerting rules that will be generated for each
	// 			time window associated with the SLO. Grafana SLOs can generate
	// 			alerts when the short-term error budget burn is very high, the
	// 			long-term error budget burn rate is high, or when the remaining
	// 			error budget is below a certain threshold. Annotations and Labels support templating.
	Alertings SLOAlertingArrayInput
	// Description is a free-text field that can provide more context to an SLO.
	Description pulumi.StringInput
	// Destination Datasource sets the datasource defined for an SLO
	DestinationDatasource SLODestinationDatasourcePtrInput
	// Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
	Labels SLOLabelArrayInput
	// Name should be a short description of your indicator. Consider names like "API Availability"
	Name pulumi.StringPtrInput
	// Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
	Objectives SLOObjectiveArrayInput
	// Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
	Queries SLOQueryArrayInput
}

func (SLOArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sloArgs)(nil)).Elem()
}

type SLOInput interface {
	pulumi.Input

	ToSLOOutput() SLOOutput
	ToSLOOutputWithContext(ctx context.Context) SLOOutput
}

func (*SLO) ElementType() reflect.Type {
	return reflect.TypeOf((**SLO)(nil)).Elem()
}

func (i *SLO) ToSLOOutput() SLOOutput {
	return i.ToSLOOutputWithContext(context.Background())
}

func (i *SLO) ToSLOOutputWithContext(ctx context.Context) SLOOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SLOOutput)
}

// SLOArrayInput is an input type that accepts SLOArray and SLOArrayOutput values.
// You can construct a concrete instance of `SLOArrayInput` via:
//
//	SLOArray{ SLOArgs{...} }
type SLOArrayInput interface {
	pulumi.Input

	ToSLOArrayOutput() SLOArrayOutput
	ToSLOArrayOutputWithContext(context.Context) SLOArrayOutput
}

type SLOArray []SLOInput

func (SLOArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SLO)(nil)).Elem()
}

func (i SLOArray) ToSLOArrayOutput() SLOArrayOutput {
	return i.ToSLOArrayOutputWithContext(context.Background())
}

func (i SLOArray) ToSLOArrayOutputWithContext(ctx context.Context) SLOArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SLOArrayOutput)
}

// SLOMapInput is an input type that accepts SLOMap and SLOMapOutput values.
// You can construct a concrete instance of `SLOMapInput` via:
//
//	SLOMap{ "key": SLOArgs{...} }
type SLOMapInput interface {
	pulumi.Input

	ToSLOMapOutput() SLOMapOutput
	ToSLOMapOutputWithContext(context.Context) SLOMapOutput
}

type SLOMap map[string]SLOInput

func (SLOMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SLO)(nil)).Elem()
}

func (i SLOMap) ToSLOMapOutput() SLOMapOutput {
	return i.ToSLOMapOutputWithContext(context.Background())
}

func (i SLOMap) ToSLOMapOutputWithContext(ctx context.Context) SLOMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SLOMapOutput)
}

type SLOOutput struct{ *pulumi.OutputState }

func (SLOOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SLO)(nil)).Elem()
}

func (o SLOOutput) ToSLOOutput() SLOOutput {
	return o
}

func (o SLOOutput) ToSLOOutputWithContext(ctx context.Context) SLOOutput {
	return o
}

// Configures the alerting rules that will be generated for each
//
//	time window associated with the SLO. Grafana SLOs can generate
//	alerts when the short-term error budget burn is very high, the
//	long-term error budget burn rate is high, or when the remaining
//	error budget is below a certain threshold. Annotations and Labels support templating.
func (o SLOOutput) Alertings() SLOAlertingArrayOutput {
	return o.ApplyT(func(v *SLO) SLOAlertingArrayOutput { return v.Alertings }).(SLOAlertingArrayOutput)
}

// Description is a free-text field that can provide more context to an SLO.
func (o SLOOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SLO) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Destination Datasource sets the datasource defined for an SLO
func (o SLOOutput) DestinationDatasource() SLODestinationDatasourcePtrOutput {
	return o.ApplyT(func(v *SLO) SLODestinationDatasourcePtrOutput { return v.DestinationDatasource }).(SLODestinationDatasourcePtrOutput)
}

// Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
func (o SLOOutput) Labels() SLOLabelArrayOutput {
	return o.ApplyT(func(v *SLO) SLOLabelArrayOutput { return v.Labels }).(SLOLabelArrayOutput)
}

// Name should be a short description of your indicator. Consider names like "API Availability"
func (o SLOOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SLO) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
func (o SLOOutput) Objectives() SLOObjectiveArrayOutput {
	return o.ApplyT(func(v *SLO) SLOObjectiveArrayOutput { return v.Objectives }).(SLOObjectiveArrayOutput)
}

// Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
func (o SLOOutput) Queries() SLOQueryArrayOutput {
	return o.ApplyT(func(v *SLO) SLOQueryArrayOutput { return v.Queries }).(SLOQueryArrayOutput)
}

type SLOArrayOutput struct{ *pulumi.OutputState }

func (SLOArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SLO)(nil)).Elem()
}

func (o SLOArrayOutput) ToSLOArrayOutput() SLOArrayOutput {
	return o
}

func (o SLOArrayOutput) ToSLOArrayOutputWithContext(ctx context.Context) SLOArrayOutput {
	return o
}

func (o SLOArrayOutput) Index(i pulumi.IntInput) SLOOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SLO {
		return vs[0].([]*SLO)[vs[1].(int)]
	}).(SLOOutput)
}

type SLOMapOutput struct{ *pulumi.OutputState }

func (SLOMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SLO)(nil)).Elem()
}

func (o SLOMapOutput) ToSLOMapOutput() SLOMapOutput {
	return o
}

func (o SLOMapOutput) ToSLOMapOutputWithContext(ctx context.Context) SLOMapOutput {
	return o
}

func (o SLOMapOutput) MapIndex(k pulumi.StringInput) SLOOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SLO {
		return vs[0].(map[string]*SLO)[vs[1].(string)]
	}).(SLOOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SLOInput)(nil)).Elem(), &SLO{})
	pulumi.RegisterInputType(reflect.TypeOf((*SLOArrayInput)(nil)).Elem(), SLOArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SLOMapInput)(nil)).Elem(), SLOMap{})
	pulumi.RegisterOutputType(SLOOutput{})
	pulumi.RegisterOutputType(SLOArrayOutput{})
	pulumi.RegisterOutputType(SLOMapOutput{})
}
