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

// * [Official documentation](https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#stacks/)
//
// Required access policy scopes:
//
// * stacks:read
// * stacks:write
// * stacks:delete
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
//			_, err := cloud.NewStack(ctx, "test", &cloud.StackArgs{
//				Name:        pulumi.String("gcloudstacktest"),
//				Slug:        pulumi.String("gcloudstacktest"),
//				RegionSlug:  pulumi.String("eu"),
//				Description: pulumi.String("Test Grafana Cloud Stack"),
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
// $ pulumi import grafana:cloud/stack:Stack name "{{ stackSlugOrID }}"
// ```
type Stack struct {
	pulumi.CustomResourceState

	// Name of the Alertmanager instance configured for this stack.
	AlertmanagerName pulumi.StringOutput `pulumi:"alertmanagerName"`
	// Status of the Alertmanager instance configured for this stack.
	AlertmanagerStatus pulumi.StringOutput `pulumi:"alertmanagerStatus"`
	// Base URL of the Alertmanager instance configured for this stack.
	AlertmanagerUrl pulumi.StringOutput `pulumi:"alertmanagerUrl"`
	// User ID of the Alertmanager instance configured for this stack.
	AlertmanagerUserId pulumi.IntOutput `pulumi:"alertmanagerUserId"`
	// Slug of the cluster where this stack resides.
	ClusterSlug pulumi.StringOutput `pulumi:"clusterSlug"`
	// Description of stack.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the Fleet Management instance configured for this stack.
	FleetManagementName pulumi.StringOutput `pulumi:"fleetManagementName"`
	// Status of the Fleet Management instance configured for this stack.
	FleetManagementStatus pulumi.StringOutput `pulumi:"fleetManagementStatus"`
	// Base URL of the Fleet Management instance configured for this stack.
	FleetManagementUrl pulumi.StringOutput `pulumi:"fleetManagementUrl"`
	// User ID of the Fleet Management instance configured for this stack.
	FleetManagementUserId pulumi.IntOutput    `pulumi:"fleetManagementUserId"`
	GraphiteName          pulumi.StringOutput `pulumi:"graphiteName"`
	GraphiteStatus        pulumi.StringOutput `pulumi:"graphiteStatus"`
	GraphiteUrl           pulumi.StringOutput `pulumi:"graphiteUrl"`
	GraphiteUserId        pulumi.IntOutput    `pulumi:"graphiteUserId"`
	// Base URL of the InfluxDB instance configured for this stack. The username is the same as the metrics' (`prometheusUserId` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/metrics/metrics-influxdb/push-from-telegraf/ for docs on how to use this.
	InfluxUrl pulumi.StringOutput `pulumi:"influxUrl"`
	// A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\-.]+$" and stacks cannot have more than 10 labels.
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
	// Base URL of the OTLP instance configured for this stack. The username is the stack's ID (`id` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/ for docs on how to use this.
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
	// Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired region. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	RegionSlug pulumi.StringPtrOutput `pulumi:"regionSlug"`
	// Subdomain that the Grafana instance will be available at. Setting slug to `<stack_slug>` will make the instance available at `https://<stack_slug>.grafana.net`.
	Slug pulumi.StringOutput `pulumi:"slug"`
	// Status of the stack.
	Status       pulumi.StringOutput `pulumi:"status"`
	TracesName   pulumi.StringOutput `pulumi:"tracesName"`
	TracesStatus pulumi.StringOutput `pulumi:"tracesStatus"`
	// Base URL of the Traces instance configured for this stack. To use this in the Tempo data source in Grafana, append `/tempo` to the URL.
	TracesUrl    pulumi.StringOutput `pulumi:"tracesUrl"`
	TracesUserId pulumi.IntOutput    `pulumi:"tracesUserId"`
	// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
	Url pulumi.StringPtrOutput `pulumi:"url"`
	// Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana instance). Defaults to `true`.
	WaitForReadiness pulumi.BoolPtrOutput `pulumi:"waitForReadiness"`
	// How long to wait for readiness (if enabled). Defaults to `5m0s`.
	WaitForReadinessTimeout pulumi.StringPtrOutput `pulumi:"waitForReadinessTimeout"`
}

// NewStack registers a new resource with the given unique name, arguments, and options.
func NewStack(ctx *pulumi.Context,
	name string, args *StackArgs, opts ...pulumi.ResourceOption) (*Stack, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Slug == nil {
		return nil, errors.New("invalid value for required argument 'Slug'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Stack
	err := ctx.RegisterResource("grafana:cloud/stack:Stack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStack gets an existing Stack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StackState, opts ...pulumi.ResourceOption) (*Stack, error) {
	var resource Stack
	err := ctx.ReadResource("grafana:cloud/stack:Stack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stack resources.
type stackState struct {
	// Name of the Alertmanager instance configured for this stack.
	AlertmanagerName *string `pulumi:"alertmanagerName"`
	// Status of the Alertmanager instance configured for this stack.
	AlertmanagerStatus *string `pulumi:"alertmanagerStatus"`
	// Base URL of the Alertmanager instance configured for this stack.
	AlertmanagerUrl *string `pulumi:"alertmanagerUrl"`
	// User ID of the Alertmanager instance configured for this stack.
	AlertmanagerUserId *int `pulumi:"alertmanagerUserId"`
	// Slug of the cluster where this stack resides.
	ClusterSlug *string `pulumi:"clusterSlug"`
	// Description of stack.
	Description *string `pulumi:"description"`
	// Name of the Fleet Management instance configured for this stack.
	FleetManagementName *string `pulumi:"fleetManagementName"`
	// Status of the Fleet Management instance configured for this stack.
	FleetManagementStatus *string `pulumi:"fleetManagementStatus"`
	// Base URL of the Fleet Management instance configured for this stack.
	FleetManagementUrl *string `pulumi:"fleetManagementUrl"`
	// User ID of the Fleet Management instance configured for this stack.
	FleetManagementUserId *int    `pulumi:"fleetManagementUserId"`
	GraphiteName          *string `pulumi:"graphiteName"`
	GraphiteStatus        *string `pulumi:"graphiteStatus"`
	GraphiteUrl           *string `pulumi:"graphiteUrl"`
	GraphiteUserId        *int    `pulumi:"graphiteUserId"`
	// Base URL of the InfluxDB instance configured for this stack. The username is the same as the metrics' (`prometheusUserId` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/metrics/metrics-influxdb/push-from-telegraf/ for docs on how to use this.
	InfluxUrl *string `pulumi:"influxUrl"`
	// A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\-.]+$" and stacks cannot have more than 10 labels.
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
	// Base URL of the OTLP instance configured for this stack. The username is the stack's ID (`id` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/ for docs on how to use this.
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
	// Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired region. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	RegionSlug *string `pulumi:"regionSlug"`
	// Subdomain that the Grafana instance will be available at. Setting slug to `<stack_slug>` will make the instance available at `https://<stack_slug>.grafana.net`.
	Slug *string `pulumi:"slug"`
	// Status of the stack.
	Status       *string `pulumi:"status"`
	TracesName   *string `pulumi:"tracesName"`
	TracesStatus *string `pulumi:"tracesStatus"`
	// Base URL of the Traces instance configured for this stack. To use this in the Tempo data source in Grafana, append `/tempo` to the URL.
	TracesUrl    *string `pulumi:"tracesUrl"`
	TracesUserId *int    `pulumi:"tracesUserId"`
	// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
	Url *string `pulumi:"url"`
	// Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana instance). Defaults to `true`.
	WaitForReadiness *bool `pulumi:"waitForReadiness"`
	// How long to wait for readiness (if enabled). Defaults to `5m0s`.
	WaitForReadinessTimeout *string `pulumi:"waitForReadinessTimeout"`
}

type StackState struct {
	// Name of the Alertmanager instance configured for this stack.
	AlertmanagerName pulumi.StringPtrInput
	// Status of the Alertmanager instance configured for this stack.
	AlertmanagerStatus pulumi.StringPtrInput
	// Base URL of the Alertmanager instance configured for this stack.
	AlertmanagerUrl pulumi.StringPtrInput
	// User ID of the Alertmanager instance configured for this stack.
	AlertmanagerUserId pulumi.IntPtrInput
	// Slug of the cluster where this stack resides.
	ClusterSlug pulumi.StringPtrInput
	// Description of stack.
	Description pulumi.StringPtrInput
	// Name of the Fleet Management instance configured for this stack.
	FleetManagementName pulumi.StringPtrInput
	// Status of the Fleet Management instance configured for this stack.
	FleetManagementStatus pulumi.StringPtrInput
	// Base URL of the Fleet Management instance configured for this stack.
	FleetManagementUrl pulumi.StringPtrInput
	// User ID of the Fleet Management instance configured for this stack.
	FleetManagementUserId pulumi.IntPtrInput
	GraphiteName          pulumi.StringPtrInput
	GraphiteStatus        pulumi.StringPtrInput
	GraphiteUrl           pulumi.StringPtrInput
	GraphiteUserId        pulumi.IntPtrInput
	// Base URL of the InfluxDB instance configured for this stack. The username is the same as the metrics' (`prometheusUserId` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/metrics/metrics-influxdb/push-from-telegraf/ for docs on how to use this.
	InfluxUrl pulumi.StringPtrInput
	// A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\-.]+$" and stacks cannot have more than 10 labels.
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
	// Base URL of the OTLP instance configured for this stack. The username is the stack's ID (`id` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/ for docs on how to use this.
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
	// Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired region. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	RegionSlug pulumi.StringPtrInput
	// Subdomain that the Grafana instance will be available at. Setting slug to `<stack_slug>` will make the instance available at `https://<stack_slug>.grafana.net`.
	Slug pulumi.StringPtrInput
	// Status of the stack.
	Status       pulumi.StringPtrInput
	TracesName   pulumi.StringPtrInput
	TracesStatus pulumi.StringPtrInput
	// Base URL of the Traces instance configured for this stack. To use this in the Tempo data source in Grafana, append `/tempo` to the URL.
	TracesUrl    pulumi.StringPtrInput
	TracesUserId pulumi.IntPtrInput
	// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
	Url pulumi.StringPtrInput
	// Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana instance). Defaults to `true`.
	WaitForReadiness pulumi.BoolPtrInput
	// How long to wait for readiness (if enabled). Defaults to `5m0s`.
	WaitForReadinessTimeout pulumi.StringPtrInput
}

func (StackState) ElementType() reflect.Type {
	return reflect.TypeOf((*stackState)(nil)).Elem()
}

type stackArgs struct {
	// Description of stack.
	Description *string `pulumi:"description"`
	// A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\-.]+$" and stacks cannot have more than 10 labels.
	Labels map[string]string `pulumi:"labels"`
	// Name of stack. Conventionally matches the url of the instance (e.g. `<stack_slug>.grafana.net`).
	Name *string `pulumi:"name"`
	// Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired region. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	RegionSlug *string `pulumi:"regionSlug"`
	// Subdomain that the Grafana instance will be available at. Setting slug to `<stack_slug>` will make the instance available at `https://<stack_slug>.grafana.net`.
	Slug string `pulumi:"slug"`
	// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
	Url *string `pulumi:"url"`
	// Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana instance). Defaults to `true`.
	WaitForReadiness *bool `pulumi:"waitForReadiness"`
	// How long to wait for readiness (if enabled). Defaults to `5m0s`.
	WaitForReadinessTimeout *string `pulumi:"waitForReadinessTimeout"`
}

// The set of arguments for constructing a Stack resource.
type StackArgs struct {
	// Description of stack.
	Description pulumi.StringPtrInput
	// A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\-.]+$" and stacks cannot have more than 10 labels.
	Labels pulumi.StringMapInput
	// Name of stack. Conventionally matches the url of the instance (e.g. `<stack_slug>.grafana.net`).
	Name pulumi.StringPtrInput
	// Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired region. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	RegionSlug pulumi.StringPtrInput
	// Subdomain that the Grafana instance will be available at. Setting slug to `<stack_slug>` will make the instance available at `https://<stack_slug>.grafana.net`.
	Slug pulumi.StringInput
	// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
	Url pulumi.StringPtrInput
	// Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana instance). Defaults to `true`.
	WaitForReadiness pulumi.BoolPtrInput
	// How long to wait for readiness (if enabled). Defaults to `5m0s`.
	WaitForReadinessTimeout pulumi.StringPtrInput
}

func (StackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stackArgs)(nil)).Elem()
}

type StackInput interface {
	pulumi.Input

	ToStackOutput() StackOutput
	ToStackOutputWithContext(ctx context.Context) StackOutput
}

func (*Stack) ElementType() reflect.Type {
	return reflect.TypeOf((**Stack)(nil)).Elem()
}

func (i *Stack) ToStackOutput() StackOutput {
	return i.ToStackOutputWithContext(context.Background())
}

func (i *Stack) ToStackOutputWithContext(ctx context.Context) StackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackOutput)
}

// StackArrayInput is an input type that accepts StackArray and StackArrayOutput values.
// You can construct a concrete instance of `StackArrayInput` via:
//
//	StackArray{ StackArgs{...} }
type StackArrayInput interface {
	pulumi.Input

	ToStackArrayOutput() StackArrayOutput
	ToStackArrayOutputWithContext(context.Context) StackArrayOutput
}

type StackArray []StackInput

func (StackArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Stack)(nil)).Elem()
}

func (i StackArray) ToStackArrayOutput() StackArrayOutput {
	return i.ToStackArrayOutputWithContext(context.Background())
}

func (i StackArray) ToStackArrayOutputWithContext(ctx context.Context) StackArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackArrayOutput)
}

// StackMapInput is an input type that accepts StackMap and StackMapOutput values.
// You can construct a concrete instance of `StackMapInput` via:
//
//	StackMap{ "key": StackArgs{...} }
type StackMapInput interface {
	pulumi.Input

	ToStackMapOutput() StackMapOutput
	ToStackMapOutputWithContext(context.Context) StackMapOutput
}

type StackMap map[string]StackInput

func (StackMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Stack)(nil)).Elem()
}

func (i StackMap) ToStackMapOutput() StackMapOutput {
	return i.ToStackMapOutputWithContext(context.Background())
}

func (i StackMap) ToStackMapOutputWithContext(ctx context.Context) StackMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackMapOutput)
}

type StackOutput struct{ *pulumi.OutputState }

func (StackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Stack)(nil)).Elem()
}

func (o StackOutput) ToStackOutput() StackOutput {
	return o
}

func (o StackOutput) ToStackOutputWithContext(ctx context.Context) StackOutput {
	return o
}

// Name of the Alertmanager instance configured for this stack.
func (o StackOutput) AlertmanagerName() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.AlertmanagerName }).(pulumi.StringOutput)
}

// Status of the Alertmanager instance configured for this stack.
func (o StackOutput) AlertmanagerStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.AlertmanagerStatus }).(pulumi.StringOutput)
}

// Base URL of the Alertmanager instance configured for this stack.
func (o StackOutput) AlertmanagerUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.AlertmanagerUrl }).(pulumi.StringOutput)
}

// User ID of the Alertmanager instance configured for this stack.
func (o StackOutput) AlertmanagerUserId() pulumi.IntOutput {
	return o.ApplyT(func(v *Stack) pulumi.IntOutput { return v.AlertmanagerUserId }).(pulumi.IntOutput)
}

// Slug of the cluster where this stack resides.
func (o StackOutput) ClusterSlug() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.ClusterSlug }).(pulumi.StringOutput)
}

// Description of stack.
func (o StackOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the Fleet Management instance configured for this stack.
func (o StackOutput) FleetManagementName() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.FleetManagementName }).(pulumi.StringOutput)
}

// Status of the Fleet Management instance configured for this stack.
func (o StackOutput) FleetManagementStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.FleetManagementStatus }).(pulumi.StringOutput)
}

// Base URL of the Fleet Management instance configured for this stack.
func (o StackOutput) FleetManagementUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.FleetManagementUrl }).(pulumi.StringOutput)
}

// User ID of the Fleet Management instance configured for this stack.
func (o StackOutput) FleetManagementUserId() pulumi.IntOutput {
	return o.ApplyT(func(v *Stack) pulumi.IntOutput { return v.FleetManagementUserId }).(pulumi.IntOutput)
}

func (o StackOutput) GraphiteName() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.GraphiteName }).(pulumi.StringOutput)
}

func (o StackOutput) GraphiteStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.GraphiteStatus }).(pulumi.StringOutput)
}

func (o StackOutput) GraphiteUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.GraphiteUrl }).(pulumi.StringOutput)
}

func (o StackOutput) GraphiteUserId() pulumi.IntOutput {
	return o.ApplyT(func(v *Stack) pulumi.IntOutput { return v.GraphiteUserId }).(pulumi.IntOutput)
}

// Base URL of the InfluxDB instance configured for this stack. The username is the same as the metrics' (`prometheusUserId` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/metrics/metrics-influxdb/push-from-telegraf/ for docs on how to use this.
func (o StackOutput) InfluxUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.InfluxUrl }).(pulumi.StringOutput)
}

// A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\-.]+$" and stacks cannot have more than 10 labels.
func (o StackOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o StackOutput) LogsName() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.LogsName }).(pulumi.StringOutput)
}

func (o StackOutput) LogsStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.LogsStatus }).(pulumi.StringOutput)
}

func (o StackOutput) LogsUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.LogsUrl }).(pulumi.StringOutput)
}

func (o StackOutput) LogsUserId() pulumi.IntOutput {
	return o.ApplyT(func(v *Stack) pulumi.IntOutput { return v.LogsUserId }).(pulumi.IntOutput)
}

// Name of stack. Conventionally matches the url of the instance (e.g. `<stack_slug>.grafana.net`).
func (o StackOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Organization id to assign to this stack.
func (o StackOutput) OrgId() pulumi.IntOutput {
	return o.ApplyT(func(v *Stack) pulumi.IntOutput { return v.OrgId }).(pulumi.IntOutput)
}

// Organization name to assign to this stack.
func (o StackOutput) OrgName() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.OrgName }).(pulumi.StringOutput)
}

// Organization slug to assign to this stack.
func (o StackOutput) OrgSlug() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.OrgSlug }).(pulumi.StringOutput)
}

// Base URL of the OTLP instance configured for this stack. The username is the stack's ID (`id` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/ for docs on how to use this.
func (o StackOutput) OtlpUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.OtlpUrl }).(pulumi.StringOutput)
}

func (o StackOutput) ProfilesName() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.ProfilesName }).(pulumi.StringOutput)
}

func (o StackOutput) ProfilesStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.ProfilesStatus }).(pulumi.StringOutput)
}

func (o StackOutput) ProfilesUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.ProfilesUrl }).(pulumi.StringOutput)
}

func (o StackOutput) ProfilesUserId() pulumi.IntOutput {
	return o.ApplyT(func(v *Stack) pulumi.IntOutput { return v.ProfilesUserId }).(pulumi.IntOutput)
}

// Prometheus name for this instance.
func (o StackOutput) PrometheusName() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.PrometheusName }).(pulumi.StringOutput)
}

// Use this URL to query hosted metrics data e.g. Prometheus data source in Grafana
func (o StackOutput) PrometheusRemoteEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.PrometheusRemoteEndpoint }).(pulumi.StringOutput)
}

// Use this URL to send prometheus metrics to Grafana cloud
func (o StackOutput) PrometheusRemoteWriteEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.PrometheusRemoteWriteEndpoint }).(pulumi.StringOutput)
}

// Prometheus status for this instance.
func (o StackOutput) PrometheusStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.PrometheusStatus }).(pulumi.StringOutput)
}

// Prometheus url for this instance.
func (o StackOutput) PrometheusUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.PrometheusUrl }).(pulumi.StringOutput)
}

// Prometheus user ID. Used for e.g. remote_write.
func (o StackOutput) PrometheusUserId() pulumi.IntOutput {
	return o.ApplyT(func(v *Stack) pulumi.IntOutput { return v.PrometheusUserId }).(pulumi.IntOutput)
}

// Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired region. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
func (o StackOutput) RegionSlug() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringPtrOutput { return v.RegionSlug }).(pulumi.StringPtrOutput)
}

// Subdomain that the Grafana instance will be available at. Setting slug to `<stack_slug>` will make the instance available at `https://<stack_slug>.grafana.net`.
func (o StackOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

// Status of the stack.
func (o StackOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o StackOutput) TracesName() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.TracesName }).(pulumi.StringOutput)
}

func (o StackOutput) TracesStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.TracesStatus }).(pulumi.StringOutput)
}

// Base URL of the Traces instance configured for this stack. To use this in the Tempo data source in Grafana, append `/tempo` to the URL.
func (o StackOutput) TracesUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.TracesUrl }).(pulumi.StringOutput)
}

func (o StackOutput) TracesUserId() pulumi.IntOutput {
	return o.ApplyT(func(v *Stack) pulumi.IntOutput { return v.TracesUserId }).(pulumi.IntOutput)
}

// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
func (o StackOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

// Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana instance). Defaults to `true`.
func (o StackOutput) WaitForReadiness() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.BoolPtrOutput { return v.WaitForReadiness }).(pulumi.BoolPtrOutput)
}

// How long to wait for readiness (if enabled). Defaults to `5m0s`.
func (o StackOutput) WaitForReadinessTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringPtrOutput { return v.WaitForReadinessTimeout }).(pulumi.StringPtrOutput)
}

type StackArrayOutput struct{ *pulumi.OutputState }

func (StackArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Stack)(nil)).Elem()
}

func (o StackArrayOutput) ToStackArrayOutput() StackArrayOutput {
	return o
}

func (o StackArrayOutput) ToStackArrayOutputWithContext(ctx context.Context) StackArrayOutput {
	return o
}

func (o StackArrayOutput) Index(i pulumi.IntInput) StackOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Stack {
		return vs[0].([]*Stack)[vs[1].(int)]
	}).(StackOutput)
}

type StackMapOutput struct{ *pulumi.OutputState }

func (StackMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Stack)(nil)).Elem()
}

func (o StackMapOutput) ToStackMapOutput() StackMapOutput {
	return o
}

func (o StackMapOutput) ToStackMapOutputWithContext(ctx context.Context) StackMapOutput {
	return o
}

func (o StackMapOutput) MapIndex(k pulumi.StringInput) StackOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Stack {
		return vs[0].(map[string]*Stack)[vs[1].(string)]
	}).(StackOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StackInput)(nil)).Elem(), &Stack{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackArrayInput)(nil)).Elem(), StackArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackMapInput)(nil)).Elem(), StackMap{})
	pulumi.RegisterOutputType(StackOutput{})
	pulumi.RegisterOutputType(StackArrayOutput{})
	pulumi.RegisterOutputType(StackMapOutput{})
}
