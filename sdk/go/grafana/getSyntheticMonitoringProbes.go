// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// Data source for retrieving all probes.
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
//			_, err := syntheticmonitoring.GetProbes(ctx, &syntheticmonitoring.GetProbesArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: grafana.index/getsyntheticmonitoringprobes.getSyntheticMonitoringProbes has been deprecated in favor of grafana.syntheticmonitoring/getprobes.getProbes
func GetSyntheticMonitoringProbes(ctx *pulumi.Context, args *GetSyntheticMonitoringProbesArgs, opts ...pulumi.InvokeOption) (*GetSyntheticMonitoringProbesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSyntheticMonitoringProbesResult
	err := ctx.Invoke("grafana:index/getSyntheticMonitoringProbes:getSyntheticMonitoringProbes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSyntheticMonitoringProbes.
type GetSyntheticMonitoringProbesArgs struct {
	// If true, only probes that are not deprecated will be returned. Defaults to `true`.
	FilterDeprecated *bool `pulumi:"filterDeprecated"`
}

// A collection of values returned by getSyntheticMonitoringProbes.
type GetSyntheticMonitoringProbesResult struct {
	// If true, only probes that are not deprecated will be returned. Defaults to `true`.
	FilterDeprecated *bool `pulumi:"filterDeprecated"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Map of probes with their names as keys and IDs as values.
	Probes map[string]int `pulumi:"probes"`
}

func GetSyntheticMonitoringProbesOutput(ctx *pulumi.Context, args GetSyntheticMonitoringProbesOutputArgs, opts ...pulumi.InvokeOption) GetSyntheticMonitoringProbesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetSyntheticMonitoringProbesResultOutput, error) {
			args := v.(GetSyntheticMonitoringProbesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("grafana:index/getSyntheticMonitoringProbes:getSyntheticMonitoringProbes", args, GetSyntheticMonitoringProbesResultOutput{}, options).(GetSyntheticMonitoringProbesResultOutput), nil
		}).(GetSyntheticMonitoringProbesResultOutput)
}

// A collection of arguments for invoking getSyntheticMonitoringProbes.
type GetSyntheticMonitoringProbesOutputArgs struct {
	// If true, only probes that are not deprecated will be returned. Defaults to `true`.
	FilterDeprecated pulumi.BoolPtrInput `pulumi:"filterDeprecated"`
}

func (GetSyntheticMonitoringProbesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSyntheticMonitoringProbesArgs)(nil)).Elem()
}

// A collection of values returned by getSyntheticMonitoringProbes.
type GetSyntheticMonitoringProbesResultOutput struct{ *pulumi.OutputState }

func (GetSyntheticMonitoringProbesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSyntheticMonitoringProbesResult)(nil)).Elem()
}

func (o GetSyntheticMonitoringProbesResultOutput) ToGetSyntheticMonitoringProbesResultOutput() GetSyntheticMonitoringProbesResultOutput {
	return o
}

func (o GetSyntheticMonitoringProbesResultOutput) ToGetSyntheticMonitoringProbesResultOutputWithContext(ctx context.Context) GetSyntheticMonitoringProbesResultOutput {
	return o
}

// If true, only probes that are not deprecated will be returned. Defaults to `true`.
func (o GetSyntheticMonitoringProbesResultOutput) FilterDeprecated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetSyntheticMonitoringProbesResult) *bool { return v.FilterDeprecated }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSyntheticMonitoringProbesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSyntheticMonitoringProbesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Map of probes with their names as keys and IDs as values.
func (o GetSyntheticMonitoringProbesResultOutput) Probes() pulumi.IntMapOutput {
	return o.ApplyT(func(v GetSyntheticMonitoringProbesResult) map[string]int { return v.Probes }).(pulumi.IntMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSyntheticMonitoringProbesResultOutput{})
}
