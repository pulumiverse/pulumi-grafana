// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// Data source for Grafana Stack
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
//			testStack, err := cloud.NewStack(ctx, "test", &cloud.StackArgs{
//				Name:        pulumi.String("gcloudstacktest"),
//				Slug:        pulumi.String("gcloudstacktest"),
//				RegionSlug:  pulumi.String("eu"),
//				Description: pulumi.String("Test Grafana Cloud Stack"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = cloud.LookupStackOutput(ctx, cloud.GetStackOutputArgs{
//				Slug: testStack.Slug,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupStack(ctx *pulumi.Context, args *LookupStackArgs, opts ...pulumi.InvokeOption) (*LookupStackResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStackResult
	err := ctx.Invoke("grafana:cloud/getStack:getStack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStack.
type LookupStackArgs struct {
	// Subdomain that the Grafana instance will be available at (i.e. setting slug to “\n\n” will make the instance
	// available at “https://\n\n.grafana.net".
	Slug string `pulumi:"slug"`
}

// A collection of values returned by getStack.
type LookupStackResult struct {
	// Name of the Alertmanager instance configured for this stack.
	AlertmanagerName string `pulumi:"alertmanagerName"`
	// Status of the Alertmanager instance configured for this stack.
	AlertmanagerStatus string `pulumi:"alertmanagerStatus"`
	// Base URL of the Alertmanager instance configured for this stack.
	AlertmanagerUrl string `pulumi:"alertmanagerUrl"`
	// User ID of the Alertmanager instance configured for this stack.
	AlertmanagerUserId int `pulumi:"alertmanagerUserId"`
	// Description of stack.
	Description    string `pulumi:"description"`
	GraphiteName   string `pulumi:"graphiteName"`
	GraphiteStatus string `pulumi:"graphiteStatus"`
	GraphiteUrl    string `pulumi:"graphiteUrl"`
	GraphiteUserId int    `pulumi:"graphiteUserId"`
	// The stack id assigned to this stack by Grafana.
	Id string `pulumi:"id"`
	// Base URL of the InfluxDB instance configured for this stack. The username is the same as the metrics' (`prometheusUserId` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/metrics/metrics-influxdb/push-from-telegraf/ for docs on how to use this.
	InfluxUrl string `pulumi:"influxUrl"`
	// A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\-.]+$" and stacks cannot have more than 10 labels.
	Labels     map[string]string `pulumi:"labels"`
	LogsName   string            `pulumi:"logsName"`
	LogsStatus string            `pulumi:"logsStatus"`
	LogsUrl    string            `pulumi:"logsUrl"`
	LogsUserId int               `pulumi:"logsUserId"`
	// Name of stack. Conventionally matches the url of the instance (e.g. `<stack_slug>.grafana.net`).
	Name string `pulumi:"name"`
	// Organization id to assign to this stack.
	OrgId int `pulumi:"orgId"`
	// Organization name to assign to this stack.
	OrgName string `pulumi:"orgName"`
	// Organization slug to assign to this stack.
	OrgSlug string `pulumi:"orgSlug"`
	// Base URL of the OTLP instance configured for this stack. The username is the stack's ID (`id` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/ for docs on how to use this.
	OtlpUrl        string `pulumi:"otlpUrl"`
	ProfilesName   string `pulumi:"profilesName"`
	ProfilesStatus string `pulumi:"profilesStatus"`
	ProfilesUrl    string `pulumi:"profilesUrl"`
	ProfilesUserId int    `pulumi:"profilesUserId"`
	// Prometheus name for this instance.
	PrometheusName string `pulumi:"prometheusName"`
	// Use this URL to query hosted metrics data e.g. Prometheus data source in Grafana
	PrometheusRemoteEndpoint string `pulumi:"prometheusRemoteEndpoint"`
	// Use this URL to send prometheus metrics to Grafana cloud
	PrometheusRemoteWriteEndpoint string `pulumi:"prometheusRemoteWriteEndpoint"`
	// Prometheus status for this instance.
	PrometheusStatus string `pulumi:"prometheusStatus"`
	// Prometheus url for this instance.
	PrometheusUrl string `pulumi:"prometheusUrl"`
	// Prometheus user ID. Used for e.g. remote_write.
	PrometheusUserId int `pulumi:"prometheusUserId"`
	// The region this stack is deployed to.
	RegionSlug string `pulumi:"regionSlug"`
	// Subdomain that the Grafana instance will be available at (i.e. setting slug to “\n\n” will make the instance
	// available at “https://\n\n.grafana.net".
	Slug string `pulumi:"slug"`
	// Status of the stack.
	Status       string `pulumi:"status"`
	TracesName   string `pulumi:"tracesName"`
	TracesStatus string `pulumi:"tracesStatus"`
	// Base URL of the Traces instance configured for this stack. To use this in the Tempo data source in Grafana, append `/tempo` to the URL.
	TracesUrl    string `pulumi:"tracesUrl"`
	TracesUserId int    `pulumi:"tracesUserId"`
	// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
	Url string `pulumi:"url"`
}

func LookupStackOutput(ctx *pulumi.Context, args LookupStackOutputArgs, opts ...pulumi.InvokeOption) LookupStackResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStackResult, error) {
			args := v.(LookupStackArgs)
			r, err := LookupStack(ctx, &args, opts...)
			var s LookupStackResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStackResultOutput)
}

// A collection of arguments for invoking getStack.
type LookupStackOutputArgs struct {
	// Subdomain that the Grafana instance will be available at (i.e. setting slug to “\n\n” will make the instance
	// available at “https://\n\n.grafana.net".
	Slug pulumi.StringInput `pulumi:"slug"`
}

func (LookupStackOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackArgs)(nil)).Elem()
}

// A collection of values returned by getStack.
type LookupStackResultOutput struct{ *pulumi.OutputState }

func (LookupStackResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackResult)(nil)).Elem()
}

func (o LookupStackResultOutput) ToLookupStackResultOutput() LookupStackResultOutput {
	return o
}

func (o LookupStackResultOutput) ToLookupStackResultOutputWithContext(ctx context.Context) LookupStackResultOutput {
	return o
}

// Name of the Alertmanager instance configured for this stack.
func (o LookupStackResultOutput) AlertmanagerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.AlertmanagerName }).(pulumi.StringOutput)
}

// Status of the Alertmanager instance configured for this stack.
func (o LookupStackResultOutput) AlertmanagerStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.AlertmanagerStatus }).(pulumi.StringOutput)
}

// Base URL of the Alertmanager instance configured for this stack.
func (o LookupStackResultOutput) AlertmanagerUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.AlertmanagerUrl }).(pulumi.StringOutput)
}

// User ID of the Alertmanager instance configured for this stack.
func (o LookupStackResultOutput) AlertmanagerUserId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStackResult) int { return v.AlertmanagerUserId }).(pulumi.IntOutput)
}

// Description of stack.
func (o LookupStackResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupStackResultOutput) GraphiteName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.GraphiteName }).(pulumi.StringOutput)
}

func (o LookupStackResultOutput) GraphiteStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.GraphiteStatus }).(pulumi.StringOutput)
}

func (o LookupStackResultOutput) GraphiteUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.GraphiteUrl }).(pulumi.StringOutput)
}

func (o LookupStackResultOutput) GraphiteUserId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStackResult) int { return v.GraphiteUserId }).(pulumi.IntOutput)
}

// The stack id assigned to this stack by Grafana.
func (o LookupStackResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.Id }).(pulumi.StringOutput)
}

// Base URL of the InfluxDB instance configured for this stack. The username is the same as the metrics' (`prometheusUserId` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/metrics/metrics-influxdb/push-from-telegraf/ for docs on how to use this.
func (o LookupStackResultOutput) InfluxUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.InfluxUrl }).(pulumi.StringOutput)
}

// A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\-.]+$" and stacks cannot have more than 10 labels.
func (o LookupStackResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStackResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupStackResultOutput) LogsName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.LogsName }).(pulumi.StringOutput)
}

func (o LookupStackResultOutput) LogsStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.LogsStatus }).(pulumi.StringOutput)
}

func (o LookupStackResultOutput) LogsUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.LogsUrl }).(pulumi.StringOutput)
}

func (o LookupStackResultOutput) LogsUserId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStackResult) int { return v.LogsUserId }).(pulumi.IntOutput)
}

// Name of stack. Conventionally matches the url of the instance (e.g. `<stack_slug>.grafana.net`).
func (o LookupStackResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.Name }).(pulumi.StringOutput)
}

// Organization id to assign to this stack.
func (o LookupStackResultOutput) OrgId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStackResult) int { return v.OrgId }).(pulumi.IntOutput)
}

// Organization name to assign to this stack.
func (o LookupStackResultOutput) OrgName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.OrgName }).(pulumi.StringOutput)
}

// Organization slug to assign to this stack.
func (o LookupStackResultOutput) OrgSlug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.OrgSlug }).(pulumi.StringOutput)
}

// Base URL of the OTLP instance configured for this stack. The username is the stack's ID (`id` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/ for docs on how to use this.
func (o LookupStackResultOutput) OtlpUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.OtlpUrl }).(pulumi.StringOutput)
}

func (o LookupStackResultOutput) ProfilesName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.ProfilesName }).(pulumi.StringOutput)
}

func (o LookupStackResultOutput) ProfilesStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.ProfilesStatus }).(pulumi.StringOutput)
}

func (o LookupStackResultOutput) ProfilesUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.ProfilesUrl }).(pulumi.StringOutput)
}

func (o LookupStackResultOutput) ProfilesUserId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStackResult) int { return v.ProfilesUserId }).(pulumi.IntOutput)
}

// Prometheus name for this instance.
func (o LookupStackResultOutput) PrometheusName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.PrometheusName }).(pulumi.StringOutput)
}

// Use this URL to query hosted metrics data e.g. Prometheus data source in Grafana
func (o LookupStackResultOutput) PrometheusRemoteEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.PrometheusRemoteEndpoint }).(pulumi.StringOutput)
}

// Use this URL to send prometheus metrics to Grafana cloud
func (o LookupStackResultOutput) PrometheusRemoteWriteEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.PrometheusRemoteWriteEndpoint }).(pulumi.StringOutput)
}

// Prometheus status for this instance.
func (o LookupStackResultOutput) PrometheusStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.PrometheusStatus }).(pulumi.StringOutput)
}

// Prometheus url for this instance.
func (o LookupStackResultOutput) PrometheusUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.PrometheusUrl }).(pulumi.StringOutput)
}

// Prometheus user ID. Used for e.g. remote_write.
func (o LookupStackResultOutput) PrometheusUserId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStackResult) int { return v.PrometheusUserId }).(pulumi.IntOutput)
}

// The region this stack is deployed to.
func (o LookupStackResultOutput) RegionSlug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.RegionSlug }).(pulumi.StringOutput)
}

// Subdomain that the Grafana instance will be available at (i.e. setting slug to “\n\n” will make the instance
// available at “https://\n\n.grafana.net".
func (o LookupStackResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.Slug }).(pulumi.StringOutput)
}

// Status of the stack.
func (o LookupStackResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupStackResultOutput) TracesName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.TracesName }).(pulumi.StringOutput)
}

func (o LookupStackResultOutput) TracesStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.TracesStatus }).(pulumi.StringOutput)
}

// Base URL of the Traces instance configured for this stack. To use this in the Tempo data source in Grafana, append `/tempo` to the URL.
func (o LookupStackResultOutput) TracesUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.TracesUrl }).(pulumi.StringOutput)
}

func (o LookupStackResultOutput) TracesUserId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStackResult) int { return v.TracesUserId }).(pulumi.IntOutput)
}

// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
func (o LookupStackResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStackResultOutput{})
}
