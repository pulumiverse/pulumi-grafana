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
//			_, err = cloud.GetProviderAwsCloudwatchScrapeJobs(ctx, &cloud.GetProviderAwsCloudwatchScrapeJobsArgs{
//				StackId: test.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetProviderAwsCloudwatchScrapeJobs(ctx *pulumi.Context, args *GetProviderAwsCloudwatchScrapeJobsArgs, opts ...pulumi.InvokeOption) (*GetProviderAwsCloudwatchScrapeJobsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProviderAwsCloudwatchScrapeJobsResult
	err := ctx.Invoke("grafana:cloud/getProviderAwsCloudwatchScrapeJobs:getProviderAwsCloudwatchScrapeJobs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProviderAwsCloudwatchScrapeJobs.
type GetProviderAwsCloudwatchScrapeJobsArgs struct {
	// A list of AWS CloudWatch Scrape Job objects associated with the given StackID.
	ScrapeJobs []GetProviderAwsCloudwatchScrapeJobsScrapeJob `pulumi:"scrapeJobs"`
	StackId    string                                        `pulumi:"stackId"`
}

// A collection of values returned by getProviderAwsCloudwatchScrapeJobs.
type GetProviderAwsCloudwatchScrapeJobsResult struct {
	Id string `pulumi:"id"`
	// A list of AWS CloudWatch Scrape Job objects associated with the given StackID.
	ScrapeJobs []GetProviderAwsCloudwatchScrapeJobsScrapeJob `pulumi:"scrapeJobs"`
	StackId    string                                        `pulumi:"stackId"`
}

func GetProviderAwsCloudwatchScrapeJobsOutput(ctx *pulumi.Context, args GetProviderAwsCloudwatchScrapeJobsOutputArgs, opts ...pulumi.InvokeOption) GetProviderAwsCloudwatchScrapeJobsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetProviderAwsCloudwatchScrapeJobsResultOutput, error) {
			args := v.(GetProviderAwsCloudwatchScrapeJobsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("grafana:cloud/getProviderAwsCloudwatchScrapeJobs:getProviderAwsCloudwatchScrapeJobs", args, GetProviderAwsCloudwatchScrapeJobsResultOutput{}, options).(GetProviderAwsCloudwatchScrapeJobsResultOutput), nil
		}).(GetProviderAwsCloudwatchScrapeJobsResultOutput)
}

// A collection of arguments for invoking getProviderAwsCloudwatchScrapeJobs.
type GetProviderAwsCloudwatchScrapeJobsOutputArgs struct {
	// A list of AWS CloudWatch Scrape Job objects associated with the given StackID.
	ScrapeJobs GetProviderAwsCloudwatchScrapeJobsScrapeJobArrayInput `pulumi:"scrapeJobs"`
	StackId    pulumi.StringInput                                    `pulumi:"stackId"`
}

func (GetProviderAwsCloudwatchScrapeJobsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProviderAwsCloudwatchScrapeJobsArgs)(nil)).Elem()
}

// A collection of values returned by getProviderAwsCloudwatchScrapeJobs.
type GetProviderAwsCloudwatchScrapeJobsResultOutput struct{ *pulumi.OutputState }

func (GetProviderAwsCloudwatchScrapeJobsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProviderAwsCloudwatchScrapeJobsResult)(nil)).Elem()
}

func (o GetProviderAwsCloudwatchScrapeJobsResultOutput) ToGetProviderAwsCloudwatchScrapeJobsResultOutput() GetProviderAwsCloudwatchScrapeJobsResultOutput {
	return o
}

func (o GetProviderAwsCloudwatchScrapeJobsResultOutput) ToGetProviderAwsCloudwatchScrapeJobsResultOutputWithContext(ctx context.Context) GetProviderAwsCloudwatchScrapeJobsResultOutput {
	return o
}

func (o GetProviderAwsCloudwatchScrapeJobsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProviderAwsCloudwatchScrapeJobsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of AWS CloudWatch Scrape Job objects associated with the given StackID.
func (o GetProviderAwsCloudwatchScrapeJobsResultOutput) ScrapeJobs() GetProviderAwsCloudwatchScrapeJobsScrapeJobArrayOutput {
	return o.ApplyT(func(v GetProviderAwsCloudwatchScrapeJobsResult) []GetProviderAwsCloudwatchScrapeJobsScrapeJob {
		return v.ScrapeJobs
	}).(GetProviderAwsCloudwatchScrapeJobsScrapeJobArrayOutput)
}

func (o GetProviderAwsCloudwatchScrapeJobsResultOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProviderAwsCloudwatchScrapeJobsResult) string { return v.StackId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProviderAwsCloudwatchScrapeJobsResultOutput{})
}
