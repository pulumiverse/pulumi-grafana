// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oncall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/outgoing_webhooks/)
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
//			_, err := onCall.GetOutgoingWebhook(ctx, &oncall.GetOutgoingWebhookArgs{
//				Name: "example_outgoing_webhook",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupOutgoingWebhook(ctx *pulumi.Context, args *LookupOutgoingWebhookArgs, opts ...pulumi.InvokeOption) (*LookupOutgoingWebhookResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOutgoingWebhookResult
	err := ctx.Invoke("grafana:onCall/getOutgoingWebhook:getOutgoingWebhook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOutgoingWebhook.
type LookupOutgoingWebhookArgs struct {
	// The outgoing webhook name.
	Name string `pulumi:"name"`
}

// A collection of values returned by getOutgoingWebhook.
type LookupOutgoingWebhookResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The outgoing webhook name.
	Name string `pulumi:"name"`
}

func LookupOutgoingWebhookOutput(ctx *pulumi.Context, args LookupOutgoingWebhookOutputArgs, opts ...pulumi.InvokeOption) LookupOutgoingWebhookResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupOutgoingWebhookResultOutput, error) {
			args := v.(LookupOutgoingWebhookArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("grafana:onCall/getOutgoingWebhook:getOutgoingWebhook", args, LookupOutgoingWebhookResultOutput{}, options).(LookupOutgoingWebhookResultOutput), nil
		}).(LookupOutgoingWebhookResultOutput)
}

// A collection of arguments for invoking getOutgoingWebhook.
type LookupOutgoingWebhookOutputArgs struct {
	// The outgoing webhook name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupOutgoingWebhookOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOutgoingWebhookArgs)(nil)).Elem()
}

// A collection of values returned by getOutgoingWebhook.
type LookupOutgoingWebhookResultOutput struct{ *pulumi.OutputState }

func (LookupOutgoingWebhookResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOutgoingWebhookResult)(nil)).Elem()
}

func (o LookupOutgoingWebhookResultOutput) ToLookupOutgoingWebhookResultOutput() LookupOutgoingWebhookResultOutput {
	return o
}

func (o LookupOutgoingWebhookResultOutput) ToLookupOutgoingWebhookResultOutputWithContext(ctx context.Context) LookupOutgoingWebhookResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOutgoingWebhookResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutgoingWebhookResult) string { return v.Id }).(pulumi.StringOutput)
}

// The outgoing webhook name.
func (o LookupOutgoingWebhookResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutgoingWebhookResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOutgoingWebhookResultOutput{})
}
