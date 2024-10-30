// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/integrations/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/onCall"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := onCall.GetIntegration(ctx, &oncall.GetIntegrationArgs{
//				Id: "CEXAMPLEID123",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: grafana.index/getoncallintegration.getOncallIntegration has been deprecated in favor of grafana.oncall/getintegration.getIntegration
func LookupOncallIntegration(ctx *pulumi.Context, args *LookupOncallIntegrationArgs, opts ...pulumi.InvokeOption) (*LookupOncallIntegrationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOncallIntegrationResult
	err := ctx.Invoke("grafana:index/getOncallIntegration:getOncallIntegration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOncallIntegration.
type LookupOncallIntegrationArgs struct {
	// The integration ID.
	Id string `pulumi:"id"`
}

// A collection of values returned by getOncallIntegration.
type LookupOncallIntegrationResult struct {
	// The integration ID.
	Id string `pulumi:"id"`
	// The integration name.
	Name string `pulumi:"name"`
}

func LookupOncallIntegrationOutput(ctx *pulumi.Context, args LookupOncallIntegrationOutputArgs, opts ...pulumi.InvokeOption) LookupOncallIntegrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOncallIntegrationResultOutput, error) {
			args := v.(LookupOncallIntegrationArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupOncallIntegrationResult
			secret, err := ctx.InvokePackageRaw("grafana:index/getOncallIntegration:getOncallIntegration", args, &rv, "", opts...)
			if err != nil {
				return LookupOncallIntegrationResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupOncallIntegrationResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupOncallIntegrationResultOutput), nil
			}
			return output, nil
		}).(LookupOncallIntegrationResultOutput)
}

// A collection of arguments for invoking getOncallIntegration.
type LookupOncallIntegrationOutputArgs struct {
	// The integration ID.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupOncallIntegrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOncallIntegrationArgs)(nil)).Elem()
}

// A collection of values returned by getOncallIntegration.
type LookupOncallIntegrationResultOutput struct{ *pulumi.OutputState }

func (LookupOncallIntegrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOncallIntegrationResult)(nil)).Elem()
}

func (o LookupOncallIntegrationResultOutput) ToLookupOncallIntegrationResultOutput() LookupOncallIntegrationResultOutput {
	return o
}

func (o LookupOncallIntegrationResultOutput) ToLookupOncallIntegrationResultOutputWithContext(ctx context.Context) LookupOncallIntegrationResultOutput {
	return o
}

// The integration ID.
func (o LookupOncallIntegrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOncallIntegrationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The integration name.
func (o LookupOncallIntegrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOncallIntegrationResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOncallIntegrationResultOutput{})
}
