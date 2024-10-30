// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud

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
func GetIps(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetIpsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIpsResult
	err := ctx.Invoke("grafana:cloud/getIps:getIps", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getIps.
type GetIpsResult struct {
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

func GetIpsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetIpsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetIpsResultOutput, error) {
		opts = internal.PkgInvokeDefaultOpts(opts)
		var rv GetIpsResult
		secret, err := ctx.InvokePackageRaw("grafana:cloud/getIps:getIps", nil, &rv, "", opts...)
		if err != nil {
			return GetIpsResultOutput{}, err
		}

		output := pulumi.ToOutput(rv).(GetIpsResultOutput)
		if secret {
			return pulumi.ToSecret(output).(GetIpsResultOutput), nil
		}
		return output, nil
	}).(GetIpsResultOutput)
}

// A collection of values returned by getIps.
type GetIpsResultOutput struct{ *pulumi.OutputState }

func (GetIpsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpsResult)(nil)).Elem()
}

func (o GetIpsResultOutput) ToGetIpsResultOutput() GetIpsResultOutput {
	return o
}

func (o GetIpsResultOutput) ToGetIpsResultOutputWithContext(ctx context.Context) GetIpsResultOutput {
	return o
}

// Set of IP addresses that are used for hosted alerts.
func (o GetIpsResultOutput) HostedAlerts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpsResult) []string { return v.HostedAlerts }).(pulumi.StringArrayOutput)
}

// Set of IP addresses that are used for hosted Grafana.
func (o GetIpsResultOutput) HostedGrafanas() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpsResult) []string { return v.HostedGrafanas }).(pulumi.StringArrayOutput)
}

// Set of IP addresses that are used for hosted logs.
func (o GetIpsResultOutput) HostedLogs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpsResult) []string { return v.HostedLogs }).(pulumi.StringArrayOutput)
}

// Set of IP addresses that are used for hosted metrics.
func (o GetIpsResultOutput) HostedMetrics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpsResult) []string { return v.HostedMetrics }).(pulumi.StringArrayOutput)
}

// Set of IP addresses that are used for hosted traces.
func (o GetIpsResultOutput) HostedTraces() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpsResult) []string { return v.HostedTraces }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIpsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpsResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIpsResultOutput{})
}
