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

// Deprecated: grafana.index/cloudstack.CloudStack has been deprecated in favor of grafana.cloud/stack.Stack
type CloudStack struct {
	pulumi.CustomResourceState

	// Name of the Alertmanager instance configured for this stack.
	AlertmanagerName pulumi.StringOutput `pulumi:"alertmanagerName"`
	// Status of the Alertmanager instance configured for this stack.
	AlertmanagerStatus pulumi.StringOutput `pulumi:"alertmanagerStatus"`
	// Base URL of the Alertmanager instance configured for this stack.
	AlertmanagerUrl pulumi.StringOutput `pulumi:"alertmanagerUrl"`
	// User ID of the Alertmanager instance configured for this stack.
	AlertmanagerUserId pulumi.IntOutput `pulumi:"alertmanagerUserId"`
	// Description of stack.
	Description    pulumi.StringPtrOutput `pulumi:"description"`
	GraphiteName   pulumi.StringOutput    `pulumi:"graphiteName"`
	GraphiteStatus pulumi.StringOutput    `pulumi:"graphiteStatus"`
	GraphiteUrl    pulumi.StringOutput    `pulumi:"graphiteUrl"`
	GraphiteUserId pulumi.IntOutput       `pulumi:"graphiteUserId"`
	// Base URL of the InfluxDB instance configured for this stack. The username is the same as the metrics'
	// (`prometheusUserId` attribute of this resource). See
	// https://grafana.com/docs/grafana-cloud/send-data/metrics/metrics-influxdb/push-from-telegraf/ for docs on how to use
	// this.
	InfluxUrl pulumi.StringOutput `pulumi:"influxUrl"`
	// A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\\-.]+$" and
	// stacks cannot have more than 10 labels.
	Labels     pulumi.StringMapOutput `pulumi:"labels"`
	LogsName   pulumi.StringOutput    `pulumi:"logsName"`
	LogsStatus pulumi.StringOutput    `pulumi:"logsStatus"`
	LogsUrl    pulumi.StringOutput    `pulumi:"logsUrl"`
	LogsUserId pulumi.IntOutput       `pulumi:"logsUserId"`
	// Name of stack. Conventionally matches the url of the instance (e.g. `<stack_slug>.grafana.net`).
	Name pulumi.StringOutput `pulumi:"name"`
	// Organization id to assign to this stack.
	OrgId pulumi.IntOutput `pulumi:"orgId"`
	// Organization name to assign to this stack.
	OrgName pulumi.StringOutput `pulumi:"orgName"`
	// Organization slug to assign to this stack.
	OrgSlug pulumi.StringOutput `pulumi:"orgSlug"`
	// Base URL of the OTLP instance configured for this stack. The username is the stack's ID (`id` attribute of this
	// resource). See https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/ for docs on how to use this.
	OtlpUrl        pulumi.StringOutput `pulumi:"otlpUrl"`
	ProfilesName   pulumi.StringOutput `pulumi:"profilesName"`
	ProfilesStatus pulumi.StringOutput `pulumi:"profilesStatus"`
	ProfilesUrl    pulumi.StringOutput `pulumi:"profilesUrl"`
	ProfilesUserId pulumi.IntOutput    `pulumi:"profilesUserId"`
	// Prometheus name for this instance.
	PrometheusName pulumi.StringOutput `pulumi:"prometheusName"`
	// Use this URL to query hosted metrics data e.g. Prometheus data source in Grafana
	PrometheusRemoteEndpoint pulumi.StringOutput `pulumi:"prometheusRemoteEndpoint"`
	// Use this URL to send prometheus metrics to Grafana cloud
	PrometheusRemoteWriteEndpoint pulumi.StringOutput `pulumi:"prometheusRemoteWriteEndpoint"`
	// Prometheus status for this instance.
	PrometheusStatus pulumi.StringOutput `pulumi:"prometheusStatus"`
	// Prometheus url for this instance.
	PrometheusUrl pulumi.StringOutput `pulumi:"prometheusUrl"`
	// Prometheus user ID. Used for e.g. remote_write.
	PrometheusUserId pulumi.IntOutput `pulumi:"prometheusUserId"`
	// Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired
	// region. Use the region list API to get the list of available regions:
	// https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	RegionSlug pulumi.StringPtrOutput `pulumi:"regionSlug"`
	// Subdomain that the Grafana instance will be available at. Setting slug to `<stack_slug>` will make the instance
	// available at `https://<stack_slug>.grafana.net`.
	Slug pulumi.StringOutput `pulumi:"slug"`
	// Status of the stack.
	Status       pulumi.StringOutput `pulumi:"status"`
	TracesName   pulumi.StringOutput `pulumi:"tracesName"`
	TracesStatus pulumi.StringOutput `pulumi:"tracesStatus"`
	// Base URL of the Traces instance configured for this stack. To use this in the Tempo data source in Grafana, append
	// `/tempo` to the URL.
	TracesUrl    pulumi.StringOutput `pulumi:"tracesUrl"`
	TracesUserId pulumi.IntOutput    `pulumi:"tracesUserId"`
	// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
	Url pulumi.StringPtrOutput `pulumi:"url"`
	// Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana
	// instance).
	WaitForReadiness pulumi.BoolPtrOutput `pulumi:"waitForReadiness"`
	// How long to wait for readiness (if enabled).
	WaitForReadinessTimeout pulumi.StringPtrOutput `pulumi:"waitForReadinessTimeout"`
}

// NewCloudStack registers a new resource with the given unique name, arguments, and options.
func NewCloudStack(ctx *pulumi.Context,
	name string, args *CloudStackArgs, opts ...pulumi.ResourceOption) (*CloudStack, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Slug == nil {
		return nil, errors.New("invalid value for required argument 'Slug'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/cloudStack:CloudStack"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CloudStack
	err := ctx.RegisterResource("grafana:index/cloudStack:CloudStack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudStack gets an existing CloudStack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudStack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudStackState, opts ...pulumi.ResourceOption) (*CloudStack, error) {
	var resource CloudStack
	err := ctx.ReadResource("grafana:index/cloudStack:CloudStack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudStack resources.
type cloudStackState struct {
	// Name of the Alertmanager instance configured for this stack.
	AlertmanagerName *string `pulumi:"alertmanagerName"`
	// Status of the Alertmanager instance configured for this stack.
	AlertmanagerStatus *string `pulumi:"alertmanagerStatus"`
	// Base URL of the Alertmanager instance configured for this stack.
	AlertmanagerUrl *string `pulumi:"alertmanagerUrl"`
	// User ID of the Alertmanager instance configured for this stack.
	AlertmanagerUserId *int `pulumi:"alertmanagerUserId"`
	// Description of stack.
	Description    *string `pulumi:"description"`
	GraphiteName   *string `pulumi:"graphiteName"`
	GraphiteStatus *string `pulumi:"graphiteStatus"`
	GraphiteUrl    *string `pulumi:"graphiteUrl"`
	GraphiteUserId *int    `pulumi:"graphiteUserId"`
	// Base URL of the InfluxDB instance configured for this stack. The username is the same as the metrics'
	// (`prometheusUserId` attribute of this resource). See
	// https://grafana.com/docs/grafana-cloud/send-data/metrics/metrics-influxdb/push-from-telegraf/ for docs on how to use
	// this.
	InfluxUrl *string `pulumi:"influxUrl"`
	// A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\\-.]+$" and
	// stacks cannot have more than 10 labels.
	Labels     map[string]string `pulumi:"labels"`
	LogsName   *string           `pulumi:"logsName"`
	LogsStatus *string           `pulumi:"logsStatus"`
	LogsUrl    *string           `pulumi:"logsUrl"`
	LogsUserId *int              `pulumi:"logsUserId"`
	// Name of stack. Conventionally matches the url of the instance (e.g. `<stack_slug>.grafana.net`).
	Name *string `pulumi:"name"`
	// Organization id to assign to this stack.
	OrgId *int `pulumi:"orgId"`
	// Organization name to assign to this stack.
	OrgName *string `pulumi:"orgName"`
	// Organization slug to assign to this stack.
	OrgSlug *string `pulumi:"orgSlug"`
	// Base URL of the OTLP instance configured for this stack. The username is the stack's ID (`id` attribute of this
	// resource). See https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/ for docs on how to use this.
	OtlpUrl        *string `pulumi:"otlpUrl"`
	ProfilesName   *string `pulumi:"profilesName"`
	ProfilesStatus *string `pulumi:"profilesStatus"`
	ProfilesUrl    *string `pulumi:"profilesUrl"`
	ProfilesUserId *int    `pulumi:"profilesUserId"`
	// Prometheus name for this instance.
	PrometheusName *string `pulumi:"prometheusName"`
	// Use this URL to query hosted metrics data e.g. Prometheus data source in Grafana
	PrometheusRemoteEndpoint *string `pulumi:"prometheusRemoteEndpoint"`
	// Use this URL to send prometheus metrics to Grafana cloud
	PrometheusRemoteWriteEndpoint *string `pulumi:"prometheusRemoteWriteEndpoint"`
	// Prometheus status for this instance.
	PrometheusStatus *string `pulumi:"prometheusStatus"`
	// Prometheus url for this instance.
	PrometheusUrl *string `pulumi:"prometheusUrl"`
	// Prometheus user ID. Used for e.g. remote_write.
	PrometheusUserId *int `pulumi:"prometheusUserId"`
	// Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired
	// region. Use the region list API to get the list of available regions:
	// https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	RegionSlug *string `pulumi:"regionSlug"`
	// Subdomain that the Grafana instance will be available at. Setting slug to `<stack_slug>` will make the instance
	// available at `https://<stack_slug>.grafana.net`.
	Slug *string `pulumi:"slug"`
	// Status of the stack.
	Status       *string `pulumi:"status"`
	TracesName   *string `pulumi:"tracesName"`
	TracesStatus *string `pulumi:"tracesStatus"`
	// Base URL of the Traces instance configured for this stack. To use this in the Tempo data source in Grafana, append
	// `/tempo` to the URL.
	TracesUrl    *string `pulumi:"tracesUrl"`
	TracesUserId *int    `pulumi:"tracesUserId"`
	// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
	Url *string `pulumi:"url"`
	// Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana
	// instance).
	WaitForReadiness *bool `pulumi:"waitForReadiness"`
	// How long to wait for readiness (if enabled).
	WaitForReadinessTimeout *string `pulumi:"waitForReadinessTimeout"`
}

type CloudStackState struct {
	// Name of the Alertmanager instance configured for this stack.
	AlertmanagerName pulumi.StringPtrInput
	// Status of the Alertmanager instance configured for this stack.
	AlertmanagerStatus pulumi.StringPtrInput
	// Base URL of the Alertmanager instance configured for this stack.
	AlertmanagerUrl pulumi.StringPtrInput
	// User ID of the Alertmanager instance configured for this stack.
	AlertmanagerUserId pulumi.IntPtrInput
	// Description of stack.
	Description    pulumi.StringPtrInput
	GraphiteName   pulumi.StringPtrInput
	GraphiteStatus pulumi.StringPtrInput
	GraphiteUrl    pulumi.StringPtrInput
	GraphiteUserId pulumi.IntPtrInput
	// Base URL of the InfluxDB instance configured for this stack. The username is the same as the metrics'
	// (`prometheusUserId` attribute of this resource). See
	// https://grafana.com/docs/grafana-cloud/send-data/metrics/metrics-influxdb/push-from-telegraf/ for docs on how to use
	// this.
	InfluxUrl pulumi.StringPtrInput
	// A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\\-.]+$" and
	// stacks cannot have more than 10 labels.
	Labels     pulumi.StringMapInput
	LogsName   pulumi.StringPtrInput
	LogsStatus pulumi.StringPtrInput
	LogsUrl    pulumi.StringPtrInput
	LogsUserId pulumi.IntPtrInput
	// Name of stack. Conventionally matches the url of the instance (e.g. `<stack_slug>.grafana.net`).
	Name pulumi.StringPtrInput
	// Organization id to assign to this stack.
	OrgId pulumi.IntPtrInput
	// Organization name to assign to this stack.
	OrgName pulumi.StringPtrInput
	// Organization slug to assign to this stack.
	OrgSlug pulumi.StringPtrInput
	// Base URL of the OTLP instance configured for this stack. The username is the stack's ID (`id` attribute of this
	// resource). See https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/ for docs on how to use this.
	OtlpUrl        pulumi.StringPtrInput
	ProfilesName   pulumi.StringPtrInput
	ProfilesStatus pulumi.StringPtrInput
	ProfilesUrl    pulumi.StringPtrInput
	ProfilesUserId pulumi.IntPtrInput
	// Prometheus name for this instance.
	PrometheusName pulumi.StringPtrInput
	// Use this URL to query hosted metrics data e.g. Prometheus data source in Grafana
	PrometheusRemoteEndpoint pulumi.StringPtrInput
	// Use this URL to send prometheus metrics to Grafana cloud
	PrometheusRemoteWriteEndpoint pulumi.StringPtrInput
	// Prometheus status for this instance.
	PrometheusStatus pulumi.StringPtrInput
	// Prometheus url for this instance.
	PrometheusUrl pulumi.StringPtrInput
	// Prometheus user ID. Used for e.g. remote_write.
	PrometheusUserId pulumi.IntPtrInput
	// Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired
	// region. Use the region list API to get the list of available regions:
	// https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	RegionSlug pulumi.StringPtrInput
	// Subdomain that the Grafana instance will be available at. Setting slug to `<stack_slug>` will make the instance
	// available at `https://<stack_slug>.grafana.net`.
	Slug pulumi.StringPtrInput
	// Status of the stack.
	Status       pulumi.StringPtrInput
	TracesName   pulumi.StringPtrInput
	TracesStatus pulumi.StringPtrInput
	// Base URL of the Traces instance configured for this stack. To use this in the Tempo data source in Grafana, append
	// `/tempo` to the URL.
	TracesUrl    pulumi.StringPtrInput
	TracesUserId pulumi.IntPtrInput
	// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
	Url pulumi.StringPtrInput
	// Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana
	// instance).
	WaitForReadiness pulumi.BoolPtrInput
	// How long to wait for readiness (if enabled).
	WaitForReadinessTimeout pulumi.StringPtrInput
}

func (CloudStackState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudStackState)(nil)).Elem()
}

type cloudStackArgs struct {
	// Description of stack.
	Description *string `pulumi:"description"`
	// A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\\-.]+$" and
	// stacks cannot have more than 10 labels.
	Labels map[string]string `pulumi:"labels"`
	// Name of stack. Conventionally matches the url of the instance (e.g. `<stack_slug>.grafana.net`).
	Name *string `pulumi:"name"`
	// Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired
	// region. Use the region list API to get the list of available regions:
	// https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	RegionSlug *string `pulumi:"regionSlug"`
	// Subdomain that the Grafana instance will be available at. Setting slug to `<stack_slug>` will make the instance
	// available at `https://<stack_slug>.grafana.net`.
	Slug string `pulumi:"slug"`
	// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
	Url *string `pulumi:"url"`
	// Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana
	// instance).
	WaitForReadiness *bool `pulumi:"waitForReadiness"`
	// How long to wait for readiness (if enabled).
	WaitForReadinessTimeout *string `pulumi:"waitForReadinessTimeout"`
}

// The set of arguments for constructing a CloudStack resource.
type CloudStackArgs struct {
	// Description of stack.
	Description pulumi.StringPtrInput
	// A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\\-.]+$" and
	// stacks cannot have more than 10 labels.
	Labels pulumi.StringMapInput
	// Name of stack. Conventionally matches the url of the instance (e.g. `<stack_slug>.grafana.net`).
	Name pulumi.StringPtrInput
	// Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired
	// region. Use the region list API to get the list of available regions:
	// https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	RegionSlug pulumi.StringPtrInput
	// Subdomain that the Grafana instance will be available at. Setting slug to `<stack_slug>` will make the instance
	// available at `https://<stack_slug>.grafana.net`.
	Slug pulumi.StringInput
	// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
	Url pulumi.StringPtrInput
	// Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana
	// instance).
	WaitForReadiness pulumi.BoolPtrInput
	// How long to wait for readiness (if enabled).
	WaitForReadinessTimeout pulumi.StringPtrInput
}

func (CloudStackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudStackArgs)(nil)).Elem()
}

type CloudStackInput interface {
	pulumi.Input

	ToCloudStackOutput() CloudStackOutput
	ToCloudStackOutputWithContext(ctx context.Context) CloudStackOutput
}

func (*CloudStack) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudStack)(nil)).Elem()
}

func (i *CloudStack) ToCloudStackOutput() CloudStackOutput {
	return i.ToCloudStackOutputWithContext(context.Background())
}

func (i *CloudStack) ToCloudStackOutputWithContext(ctx context.Context) CloudStackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudStackOutput)
}

// CloudStackArrayInput is an input type that accepts CloudStackArray and CloudStackArrayOutput values.
// You can construct a concrete instance of `CloudStackArrayInput` via:
//
//	CloudStackArray{ CloudStackArgs{...} }
type CloudStackArrayInput interface {
	pulumi.Input

	ToCloudStackArrayOutput() CloudStackArrayOutput
	ToCloudStackArrayOutputWithContext(context.Context) CloudStackArrayOutput
}

type CloudStackArray []CloudStackInput

func (CloudStackArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudStack)(nil)).Elem()
}

func (i CloudStackArray) ToCloudStackArrayOutput() CloudStackArrayOutput {
	return i.ToCloudStackArrayOutputWithContext(context.Background())
}

func (i CloudStackArray) ToCloudStackArrayOutputWithContext(ctx context.Context) CloudStackArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudStackArrayOutput)
}

// CloudStackMapInput is an input type that accepts CloudStackMap and CloudStackMapOutput values.
// You can construct a concrete instance of `CloudStackMapInput` via:
//
//	CloudStackMap{ "key": CloudStackArgs{...} }
type CloudStackMapInput interface {
	pulumi.Input

	ToCloudStackMapOutput() CloudStackMapOutput
	ToCloudStackMapOutputWithContext(context.Context) CloudStackMapOutput
}

type CloudStackMap map[string]CloudStackInput

func (CloudStackMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudStack)(nil)).Elem()
}

func (i CloudStackMap) ToCloudStackMapOutput() CloudStackMapOutput {
	return i.ToCloudStackMapOutputWithContext(context.Background())
}

func (i CloudStackMap) ToCloudStackMapOutputWithContext(ctx context.Context) CloudStackMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudStackMapOutput)
}

type CloudStackOutput struct{ *pulumi.OutputState }

func (CloudStackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudStack)(nil)).Elem()
}

func (o CloudStackOutput) ToCloudStackOutput() CloudStackOutput {
	return o
}

func (o CloudStackOutput) ToCloudStackOutputWithContext(ctx context.Context) CloudStackOutput {
	return o
}

// Name of the Alertmanager instance configured for this stack.
func (o CloudStackOutput) AlertmanagerName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.AlertmanagerName }).(pulumi.StringOutput)
}

// Status of the Alertmanager instance configured for this stack.
func (o CloudStackOutput) AlertmanagerStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.AlertmanagerStatus }).(pulumi.StringOutput)
}

// Base URL of the Alertmanager instance configured for this stack.
func (o CloudStackOutput) AlertmanagerUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.AlertmanagerUrl }).(pulumi.StringOutput)
}

// User ID of the Alertmanager instance configured for this stack.
func (o CloudStackOutput) AlertmanagerUserId() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.IntOutput { return v.AlertmanagerUserId }).(pulumi.IntOutput)
}

// Description of stack.
func (o CloudStackOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CloudStackOutput) GraphiteName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.GraphiteName }).(pulumi.StringOutput)
}

func (o CloudStackOutput) GraphiteStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.GraphiteStatus }).(pulumi.StringOutput)
}

func (o CloudStackOutput) GraphiteUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.GraphiteUrl }).(pulumi.StringOutput)
}

func (o CloudStackOutput) GraphiteUserId() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.IntOutput { return v.GraphiteUserId }).(pulumi.IntOutput)
}

// Base URL of the InfluxDB instance configured for this stack. The username is the same as the metrics'
// (`prometheusUserId` attribute of this resource). See
// https://grafana.com/docs/grafana-cloud/send-data/metrics/metrics-influxdb/push-from-telegraf/ for docs on how to use
// this.
func (o CloudStackOutput) InfluxUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.InfluxUrl }).(pulumi.StringOutput)
}

// A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\\-.]+$" and
// stacks cannot have more than 10 labels.
func (o CloudStackOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o CloudStackOutput) LogsName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.LogsName }).(pulumi.StringOutput)
}

func (o CloudStackOutput) LogsStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.LogsStatus }).(pulumi.StringOutput)
}

func (o CloudStackOutput) LogsUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.LogsUrl }).(pulumi.StringOutput)
}

func (o CloudStackOutput) LogsUserId() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.IntOutput { return v.LogsUserId }).(pulumi.IntOutput)
}

// Name of stack. Conventionally matches the url of the instance (e.g. `<stack_slug>.grafana.net`).
func (o CloudStackOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Organization id to assign to this stack.
func (o CloudStackOutput) OrgId() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.IntOutput { return v.OrgId }).(pulumi.IntOutput)
}

// Organization name to assign to this stack.
func (o CloudStackOutput) OrgName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.OrgName }).(pulumi.StringOutput)
}

// Organization slug to assign to this stack.
func (o CloudStackOutput) OrgSlug() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.OrgSlug }).(pulumi.StringOutput)
}

// Base URL of the OTLP instance configured for this stack. The username is the stack's ID (`id` attribute of this
// resource). See https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/ for docs on how to use this.
func (o CloudStackOutput) OtlpUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.OtlpUrl }).(pulumi.StringOutput)
}

func (o CloudStackOutput) ProfilesName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.ProfilesName }).(pulumi.StringOutput)
}

func (o CloudStackOutput) ProfilesStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.ProfilesStatus }).(pulumi.StringOutput)
}

func (o CloudStackOutput) ProfilesUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.ProfilesUrl }).(pulumi.StringOutput)
}

func (o CloudStackOutput) ProfilesUserId() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.IntOutput { return v.ProfilesUserId }).(pulumi.IntOutput)
}

// Prometheus name for this instance.
func (o CloudStackOutput) PrometheusName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.PrometheusName }).(pulumi.StringOutput)
}

// Use this URL to query hosted metrics data e.g. Prometheus data source in Grafana
func (o CloudStackOutput) PrometheusRemoteEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.PrometheusRemoteEndpoint }).(pulumi.StringOutput)
}

// Use this URL to send prometheus metrics to Grafana cloud
func (o CloudStackOutput) PrometheusRemoteWriteEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.PrometheusRemoteWriteEndpoint }).(pulumi.StringOutput)
}

// Prometheus status for this instance.
func (o CloudStackOutput) PrometheusStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.PrometheusStatus }).(pulumi.StringOutput)
}

// Prometheus url for this instance.
func (o CloudStackOutput) PrometheusUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.PrometheusUrl }).(pulumi.StringOutput)
}

// Prometheus user ID. Used for e.g. remote_write.
func (o CloudStackOutput) PrometheusUserId() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.IntOutput { return v.PrometheusUserId }).(pulumi.IntOutput)
}

// Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired
// region. Use the region list API to get the list of available regions:
// https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
func (o CloudStackOutput) RegionSlug() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringPtrOutput { return v.RegionSlug }).(pulumi.StringPtrOutput)
}

// Subdomain that the Grafana instance will be available at. Setting slug to `<stack_slug>` will make the instance
// available at `https://<stack_slug>.grafana.net`.
func (o CloudStackOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

// Status of the stack.
func (o CloudStackOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o CloudStackOutput) TracesName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.TracesName }).(pulumi.StringOutput)
}

func (o CloudStackOutput) TracesStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.TracesStatus }).(pulumi.StringOutput)
}

// Base URL of the Traces instance configured for this stack. To use this in the Tempo data source in Grafana, append
// `/tempo` to the URL.
func (o CloudStackOutput) TracesUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringOutput { return v.TracesUrl }).(pulumi.StringOutput)
}

func (o CloudStackOutput) TracesUserId() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.IntOutput { return v.TracesUserId }).(pulumi.IntOutput)
}

// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
func (o CloudStackOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

// Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana
// instance).
func (o CloudStackOutput) WaitForReadiness() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.BoolPtrOutput { return v.WaitForReadiness }).(pulumi.BoolPtrOutput)
}

// How long to wait for readiness (if enabled).
func (o CloudStackOutput) WaitForReadinessTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudStack) pulumi.StringPtrOutput { return v.WaitForReadinessTimeout }).(pulumi.StringPtrOutput)
}

type CloudStackArrayOutput struct{ *pulumi.OutputState }

func (CloudStackArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudStack)(nil)).Elem()
}

func (o CloudStackArrayOutput) ToCloudStackArrayOutput() CloudStackArrayOutput {
	return o
}

func (o CloudStackArrayOutput) ToCloudStackArrayOutputWithContext(ctx context.Context) CloudStackArrayOutput {
	return o
}

func (o CloudStackArrayOutput) Index(i pulumi.IntInput) CloudStackOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudStack {
		return vs[0].([]*CloudStack)[vs[1].(int)]
	}).(CloudStackOutput)
}

type CloudStackMapOutput struct{ *pulumi.OutputState }

func (CloudStackMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudStack)(nil)).Elem()
}

func (o CloudStackMapOutput) ToCloudStackMapOutput() CloudStackMapOutput {
	return o
}

func (o CloudStackMapOutput) ToCloudStackMapOutputWithContext(ctx context.Context) CloudStackMapOutput {
	return o
}

func (o CloudStackMapOutput) MapIndex(k pulumi.StringInput) CloudStackOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudStack {
		return vs[0].(map[string]*CloudStack)[vs[1].(string)]
	}).(CloudStackOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudStackInput)(nil)).Elem(), &CloudStack{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudStackArrayInput)(nil)).Elem(), CloudStackArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudStackMapInput)(nil)).Elem(), CloudStackMap{})
	pulumi.RegisterOutputType(CloudStackOutput{})
	pulumi.RegisterOutputType(CloudStackArrayOutput{})
	pulumi.RegisterOutputType(CloudStackMapOutput{})
}
