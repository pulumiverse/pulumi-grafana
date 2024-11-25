// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud

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
//			testProviderAwsAccount, err := cloud.NewProviderAwsAccount(ctx, "test", &cloud.ProviderAwsAccountArgs{
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
//			_ = testProviderAwsAccount.ResourceId.ApplyT(func(resourceId string) (cloud.GetProviderAwsAccountResult, error) {
//				return cloud.GetProviderAwsAccountResult(interface{}(cloud.LookupProviderAwsAccountOutput(ctx, cloud.GetProviderAwsAccountOutputArgs{
//					StackId:    test.Id,
//					ResourceId: resourceId,
//				}, nil))), nil
//			}).(cloud.GetProviderAwsAccountResultOutput)
//			return nil
//		})
//	}
//
// ```
func LookupProviderAwsAccount(ctx *pulumi.Context, args *LookupProviderAwsAccountArgs, opts ...pulumi.InvokeOption) (*LookupProviderAwsAccountResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProviderAwsAccountResult
	err := ctx.Invoke("grafana:cloud/getProviderAwsAccount:getProviderAwsAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProviderAwsAccount.
type LookupProviderAwsAccountArgs struct {
	// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
	ResourceId string `pulumi:"resourceId"`
	StackId    string `pulumi:"stackId"`
}

// A collection of values returned by getProviderAwsAccount.
type LookupProviderAwsAccountResult struct {
	Id string `pulumi:"id"`
	// A set of regions that this AWS Account resource applies to.
	Regions []string `pulumi:"regions"`
	// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
	ResourceId string `pulumi:"resourceId"`
	// An IAM Role ARN string to represent with this AWS Account resource.
	RoleArn string `pulumi:"roleArn"`
	StackId string `pulumi:"stackId"`
}

func LookupProviderAwsAccountOutput(ctx *pulumi.Context, args LookupProviderAwsAccountOutputArgs, opts ...pulumi.InvokeOption) LookupProviderAwsAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProviderAwsAccountResultOutput, error) {
			args := v.(LookupProviderAwsAccountArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupProviderAwsAccountResult
			secret, err := ctx.InvokePackageRaw("grafana:cloud/getProviderAwsAccount:getProviderAwsAccount", args, &rv, "", opts...)
			if err != nil {
				return LookupProviderAwsAccountResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupProviderAwsAccountResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupProviderAwsAccountResultOutput), nil
			}
			return output, nil
		}).(LookupProviderAwsAccountResultOutput)
}

// A collection of arguments for invoking getProviderAwsAccount.
type LookupProviderAwsAccountOutputArgs struct {
	// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	StackId    pulumi.StringInput `pulumi:"stackId"`
}

func (LookupProviderAwsAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProviderAwsAccountArgs)(nil)).Elem()
}

// A collection of values returned by getProviderAwsAccount.
type LookupProviderAwsAccountResultOutput struct{ *pulumi.OutputState }

func (LookupProviderAwsAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProviderAwsAccountResult)(nil)).Elem()
}

func (o LookupProviderAwsAccountResultOutput) ToLookupProviderAwsAccountResultOutput() LookupProviderAwsAccountResultOutput {
	return o
}

func (o LookupProviderAwsAccountResultOutput) ToLookupProviderAwsAccountResultOutputWithContext(ctx context.Context) LookupProviderAwsAccountResultOutput {
	return o
}

func (o LookupProviderAwsAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderAwsAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// A set of regions that this AWS Account resource applies to.
func (o LookupProviderAwsAccountResultOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProviderAwsAccountResult) []string { return v.Regions }).(pulumi.StringArrayOutput)
}

// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
func (o LookupProviderAwsAccountResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderAwsAccountResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

// An IAM Role ARN string to represent with this AWS Account resource.
func (o LookupProviderAwsAccountResultOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderAwsAccountResult) string { return v.RoleArn }).(pulumi.StringOutput)
}

func (o LookupProviderAwsAccountResultOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderAwsAccountResult) string { return v.StackId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProviderAwsAccountResultOutput{})
}
