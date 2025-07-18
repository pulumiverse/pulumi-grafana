// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// Data source for retrieving sets of cloud IPs. See https://grafana.com/docs/grafana-cloud/reference/allow-list/ for more info
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
//			_, err := cloud.GetIps(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: grafana.index/getcloudips.getCloudIps has been deprecated in favor of grafana.cloud/getips.getIps
func GetCloudIps(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetCloudIpsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCloudIpsResult
	err := ctx.Invoke("grafana:index/getCloudIps:getCloudIps", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getCloudIps.
type GetCloudIpsResult struct {
	// Set of IP addresses that are used for hosted alerts.
	HostedAlerts []string `pulumi:"hostedAlerts"`
	// Set of IP addresses that are used for hosted Grafana.
	HostedGrafanas []string `pulumi:"hostedGrafanas"`
	// Set of IP addresses that are used for hosted logs.
	HostedLogs []string `pulumi:"hostedLogs"`
	// Set of IP addresses that are used for hosted metrics.
	HostedMetrics []string `pulumi:"hostedMetrics"`
	// Set of IP addresses that are used for hosted traces.
	HostedTraces []string `pulumi:"hostedTraces"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func GetCloudIpsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetCloudIpsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetCloudIpsResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("grafana:index/getCloudIps:getCloudIps", nil, GetCloudIpsResultOutput{}, options).(GetCloudIpsResultOutput), nil
	}).(GetCloudIpsResultOutput)
}

// A collection of values returned by getCloudIps.
type GetCloudIpsResultOutput struct{ *pulumi.OutputState }

func (GetCloudIpsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudIpsResult)(nil)).Elem()
}

func (o GetCloudIpsResultOutput) ToGetCloudIpsResultOutput() GetCloudIpsResultOutput {
	return o
}

func (o GetCloudIpsResultOutput) ToGetCloudIpsResultOutputWithContext(ctx context.Context) GetCloudIpsResultOutput {
	return o
}

// Set of IP addresses that are used for hosted alerts.
func (o GetCloudIpsResultOutput) HostedAlerts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCloudIpsResult) []string { return v.HostedAlerts }).(pulumi.StringArrayOutput)
}

// Set of IP addresses that are used for hosted Grafana.
func (o GetCloudIpsResultOutput) HostedGrafanas() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCloudIpsResult) []string { return v.HostedGrafanas }).(pulumi.StringArrayOutput)
}

// Set of IP addresses that are used for hosted logs.
func (o GetCloudIpsResultOutput) HostedLogs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCloudIpsResult) []string { return v.HostedLogs }).(pulumi.StringArrayOutput)
}

// Set of IP addresses that are used for hosted metrics.
func (o GetCloudIpsResultOutput) HostedMetrics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCloudIpsResult) []string { return v.HostedMetrics }).(pulumi.StringArrayOutput)
}

// Set of IP addresses that are used for hosted traces.
func (o GetCloudIpsResultOutput) HostedTraces() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCloudIpsResult) []string { return v.HostedTraces }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCloudIpsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudIpsResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudIpsResultOutput{})
}
