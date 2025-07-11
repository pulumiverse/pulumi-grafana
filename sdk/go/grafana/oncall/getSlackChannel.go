// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oncall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/slack_channels/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/oncall"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oncall.GetSlackChannel(ctx, &oncall.GetSlackChannelArgs{
//				Name: "example_slack_channel",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSlackChannel(ctx *pulumi.Context, args *GetSlackChannelArgs, opts ...pulumi.InvokeOption) (*GetSlackChannelResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSlackChannelResult
	err := ctx.Invoke("grafana:onCall/getSlackChannel:getSlackChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSlackChannel.
type GetSlackChannelArgs struct {
	// The Slack channel name.
	Name string `pulumi:"name"`
}

// A collection of values returned by getSlackChannel.
type GetSlackChannelResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Slack channel name.
	Name string `pulumi:"name"`
	// The Slack ID of the channel.
	SlackId string `pulumi:"slackId"`
}

func GetSlackChannelOutput(ctx *pulumi.Context, args GetSlackChannelOutputArgs, opts ...pulumi.InvokeOption) GetSlackChannelResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetSlackChannelResultOutput, error) {
			args := v.(GetSlackChannelArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("grafana:onCall/getSlackChannel:getSlackChannel", args, GetSlackChannelResultOutput{}, options).(GetSlackChannelResultOutput), nil
		}).(GetSlackChannelResultOutput)
}

// A collection of arguments for invoking getSlackChannel.
type GetSlackChannelOutputArgs struct {
	// The Slack channel name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetSlackChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlackChannelArgs)(nil)).Elem()
}

// A collection of values returned by getSlackChannel.
type GetSlackChannelResultOutput struct{ *pulumi.OutputState }

func (GetSlackChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlackChannelResult)(nil)).Elem()
}

func (o GetSlackChannelResultOutput) ToGetSlackChannelResultOutput() GetSlackChannelResultOutput {
	return o
}

func (o GetSlackChannelResultOutput) ToGetSlackChannelResultOutputWithContext(ctx context.Context) GetSlackChannelResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSlackChannelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlackChannelResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Slack channel name.
func (o GetSlackChannelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlackChannelResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Slack ID of the channel.
func (o GetSlackChannelResultOutput) SlackId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlackChannelResult) string { return v.SlackId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSlackChannelResultOutput{})
}
