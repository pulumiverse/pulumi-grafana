// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprovider

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/cloud"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/cloudprovider"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := cloud.LookupStack(ctx, &cloud.LookupStackArgs{
//				Slug: "gcloudstacktest",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			testAwsAccount, err := cloudprovider.NewAwsAccount(ctx, "test", &cloudprovider.AwsAccountArgs{
//				StackId: pulumi.String(test.Id),
//				RoleArn: pulumi.Any(testAwsIamRole.Arn),
//				Regions: pulumi.StringArray{
//					pulumi.String("us-east-2"),
//					pulumi.String("eu-west-3"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = testAwsAccount.ResourceId.ApplyT(func(resourceId string) (cloudprovider.GetAwsAccountResult, error) {
//				return cloudprovider.GetAwsAccountResult(interface{}(cloudprovider.GetAwsAccountOutput(ctx, cloudprovider.GetAwsAccountOutputArgs{
//					StackId:    test.Id,
//					ResourceId: resourceId,
//				}, nil))), nil
//			}).(cloudprovider.GetAwsAccountResultOutput)
//			return nil
//		})
//	}
//
// ```
func LookupAwsAccount(ctx *pulumi.Context, args *LookupAwsAccountArgs, opts ...pulumi.InvokeOption) (*LookupAwsAccountResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAwsAccountResult
	err := ctx.Invoke("grafana:cloudProvider/getAwsAccount:getAwsAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAwsAccount.
type LookupAwsAccountArgs struct {
	// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
	ResourceId string `pulumi:"resourceId"`
	StackId    string `pulumi:"stackId"`
}

// A collection of values returned by getAwsAccount.
type LookupAwsAccountResult struct {
	Id string `pulumi:"id"`
	// An optional human-readable name for this AWS Account resource.
	Name string `pulumi:"name"`
	// A set of regions that this AWS Account resource applies to.
	Regions []string `pulumi:"regions"`
	// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
	ResourceId string `pulumi:"resourceId"`
	// An IAM Role ARN string to represent with this AWS Account resource.
	RoleArn string `pulumi:"roleArn"`
	StackId string `pulumi:"stackId"`
}

func LookupAwsAccountOutput(ctx *pulumi.Context, args LookupAwsAccountOutputArgs, opts ...pulumi.InvokeOption) LookupAwsAccountResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAwsAccountResultOutput, error) {
			args := v.(LookupAwsAccountArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("grafana:cloudProvider/getAwsAccount:getAwsAccount", args, LookupAwsAccountResultOutput{}, options).(LookupAwsAccountResultOutput), nil
		}).(LookupAwsAccountResultOutput)
}

// A collection of arguments for invoking getAwsAccount.
type LookupAwsAccountOutputArgs struct {
	// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	StackId    pulumi.StringInput `pulumi:"stackId"`
}

func (LookupAwsAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAwsAccountArgs)(nil)).Elem()
}

// A collection of values returned by getAwsAccount.
type LookupAwsAccountResultOutput struct{ *pulumi.OutputState }

func (LookupAwsAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAwsAccountResult)(nil)).Elem()
}

func (o LookupAwsAccountResultOutput) ToLookupAwsAccountResultOutput() LookupAwsAccountResultOutput {
	return o
}

func (o LookupAwsAccountResultOutput) ToLookupAwsAccountResultOutputWithContext(ctx context.Context) LookupAwsAccountResultOutput {
	return o
}

func (o LookupAwsAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// An optional human-readable name for this AWS Account resource.
func (o LookupAwsAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// A set of regions that this AWS Account resource applies to.
func (o LookupAwsAccountResultOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAwsAccountResult) []string { return v.Regions }).(pulumi.StringArrayOutput)
}

// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
func (o LookupAwsAccountResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsAccountResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

// An IAM Role ARN string to represent with this AWS Account resource.
func (o LookupAwsAccountResultOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsAccountResult) string { return v.RoleArn }).(pulumi.StringOutput)
}

func (o LookupAwsAccountResultOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsAccountResult) string { return v.StackId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAwsAccountResultOutput{})
}
