// Code generated by pulumi-language-go DO NOT EDIT.
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
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
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
//			testGetRole, err := iam.LookupRole(ctx, &iam.LookupRoleArgs{
//				Name: "my-role",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			testAwsAccount, err := cloudprovider.NewAwsAccount(ctx, "test", &cloudprovider.AwsAccountArgs{
//				StackId: pulumi.String(test.Id),
//				RoleArn: pulumi.String(testGetRole.Arn),
//				Regions: pulumi.StringArray{
//					pulumi.String("us-east-1"),
//					pulumi.String("us-east-2"),
//					pulumi.String("us-west-1"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			testAwsCloudwatchScrapeJob, err := cloudprovider.NewAwsCloudwatchScrapeJob(ctx, "test", &cloudprovider.AwsCloudwatchScrapeJobArgs{
//				StackId:              pulumi.String(test.Id),
//				Name:                 pulumi.String("my-cloudwatch-scrape-job"),
//				AwsAccountResourceId: testAwsAccount.ResourceId,
//				ExportTags:           pulumi.Bool(true),
//				Services: cloudprovider.AwsCloudwatchScrapeJobServiceArray{
//					&cloudprovider.AwsCloudwatchScrapeJobServiceArgs{
//						Name: pulumi.String("AWS/EC2"),
//						Metrics: cloudprovider.AwsCloudwatchScrapeJobServiceMetricArray{
//							&cloudprovider.AwsCloudwatchScrapeJobServiceMetricArgs{
//								Name: pulumi.String("CPUUtilization"),
//								Statistics: pulumi.StringArray{
//									pulumi.String("Average"),
//								},
//							},
//							&cloudprovider.AwsCloudwatchScrapeJobServiceMetricArgs{
//								Name: pulumi.String("StatusCheckFailed"),
//								Statistics: pulumi.StringArray{
//									pulumi.String("Maximum"),
//								},
//							},
//						},
//						ScrapeIntervalSeconds: pulumi.Int(300),
//						ResourceDiscoveryTagFilters: cloudprovider.AwsCloudwatchScrapeJobServiceResourceDiscoveryTagFilterArray{
//							&cloudprovider.AwsCloudwatchScrapeJobServiceResourceDiscoveryTagFilterArgs{
//								Key:   pulumi.String("k8s.io/cluster-autoscaler/enabled"),
//								Value: pulumi.String("true"),
//							},
//						},
//						TagsToAddToMetrics: pulumi.StringArray{
//							pulumi.String("eks:cluster-name"),
//						},
//					},
//				},
//				CustomNamespaces: cloudprovider.AwsCloudwatchScrapeJobCustomNamespaceArray{
//					&cloudprovider.AwsCloudwatchScrapeJobCustomNamespaceArgs{
//						Name: pulumi.String("CoolApp"),
//						Metrics: cloudprovider.AwsCloudwatchScrapeJobCustomNamespaceMetricArray{
//							&cloudprovider.AwsCloudwatchScrapeJobCustomNamespaceMetricArgs{
//								Name: pulumi.String("CoolMetric"),
//								Statistics: pulumi.StringArray{
//									pulumi.String("Maximum"),
//									pulumi.String("Sum"),
//								},
//							},
//						},
//						ScrapeIntervalSeconds: pulumi.Int(300),
//					},
//				},
//				StaticLabels: pulumi.StringMap{
//					"label1": pulumi.String("value1"),
//					"label2": pulumi.String("value2"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = testAwsCloudwatchScrapeJob.Name.ApplyT(func(name string) (cloudprovider.GetAwsCloudwatchScrapeJobResult, error) {
//				return cloudprovider.GetAwsCloudwatchScrapeJobResult(cloudprovider.GetAwsCloudwatchScrapeJob(ctx, &cloudprovider.GetAwsCloudwatchScrapeJobArgs{
//					StackId: test.Id,
//					Name:    name,
//				}, nil)), nil
//			}).(cloudprovider.GetAwsCloudwatchScrapeJobResultOutput)
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: grafana.cloud/getproviderawscloudwatchscrapejob.getProviderAwsCloudwatchScrapeJob has been deprecated in favor of grafana.cloudprovider/getawscloudwatchscrapejob.getAwsCloudwatchScrapeJob
func LookupProviderAwsCloudwatchScrapeJob(ctx *pulumi.Context, args *LookupProviderAwsCloudwatchScrapeJobArgs, opts ...pulumi.InvokeOption) (*LookupProviderAwsCloudwatchScrapeJobResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProviderAwsCloudwatchScrapeJobResult
	err := ctx.Invoke("grafana:cloud/getProviderAwsCloudwatchScrapeJob:getProviderAwsCloudwatchScrapeJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProviderAwsCloudwatchScrapeJob.
type LookupProviderAwsCloudwatchScrapeJobArgs struct {
	// Zero or more configuration blocks to configure custom namespaces for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	CustomNamespaces []GetProviderAwsCloudwatchScrapeJobCustomNamespace `pulumi:"customNamespaces"`
	Name             string                                             `pulumi:"name"`
	// One or more configuration blocks to dictate what this AWS CloudWatch Scrape Job should scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	Services []GetProviderAwsCloudwatchScrapeJobService `pulumi:"services"`
	StackId  string                                     `pulumi:"stackId"`
}

// A collection of values returned by getProviderAwsCloudwatchScrapeJob.
type LookupProviderAwsCloudwatchScrapeJobResult struct {
	// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this AWS CloudWatch Scrape Job. This can be provided by the `resourceId` attribute of the `cloudProvider.AwsAccount` resource.
	AwsAccountResourceId string `pulumi:"awsAccountResourceId"`
	// Zero or more configuration blocks to configure custom namespaces for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	CustomNamespaces []GetProviderAwsCloudwatchScrapeJobCustomNamespace `pulumi:"customNamespaces"`
	// When the AWS CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
	DisabledReason string `pulumi:"disabledReason"`
	// Whether the AWS CloudWatch Scrape Job is enabled or not.
	Enabled bool `pulumi:"enabled"`
	// When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`.
	ExportTags bool   `pulumi:"exportTags"`
	Id         string `pulumi:"id"`
	Name       string `pulumi:"name"`
	// The set of AWS region names that this AWS CloudWatch Scrape Job is configured to scrape.
	Regions []string `pulumi:"regions"`
	// When true, the `regions` attribute will be the set of regions configured in the override. When false, the `regions` attribute will be the set of regions belonging to the AWS Account resource that is associated with this AWS CloudWatch Scrape Job.
	RegionsSubsetOverrideUsed bool `pulumi:"regionsSubsetOverrideUsed"`
	// The AWS ARN of the IAM role associated with the AWS Account resource that is being used by this AWS CloudWatch Scrape Job.
	RoleArn string `pulumi:"roleArn"`
	// One or more configuration blocks to dictate what this AWS CloudWatch Scrape Job should scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	Services []GetProviderAwsCloudwatchScrapeJobService `pulumi:"services"`
	StackId  string                                     `pulumi:"stackId"`
	// A set of static labels to add to all metrics exported by this scrape job.
	StaticLabels map[string]string `pulumi:"staticLabels"`
}

func LookupProviderAwsCloudwatchScrapeJobOutput(ctx *pulumi.Context, args LookupProviderAwsCloudwatchScrapeJobOutputArgs, opts ...pulumi.InvokeOption) LookupProviderAwsCloudwatchScrapeJobResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProviderAwsCloudwatchScrapeJobResultOutput, error) {
			args := v.(LookupProviderAwsCloudwatchScrapeJobArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("grafana:cloud/getProviderAwsCloudwatchScrapeJob:getProviderAwsCloudwatchScrapeJob", args, LookupProviderAwsCloudwatchScrapeJobResultOutput{}, options).(LookupProviderAwsCloudwatchScrapeJobResultOutput), nil
		}).(LookupProviderAwsCloudwatchScrapeJobResultOutput)
}

// A collection of arguments for invoking getProviderAwsCloudwatchScrapeJob.
type LookupProviderAwsCloudwatchScrapeJobOutputArgs struct {
	// Zero or more configuration blocks to configure custom namespaces for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	CustomNamespaces GetProviderAwsCloudwatchScrapeJobCustomNamespaceArrayInput `pulumi:"customNamespaces"`
	Name             pulumi.StringInput                                         `pulumi:"name"`
	// One or more configuration blocks to dictate what this AWS CloudWatch Scrape Job should scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	Services GetProviderAwsCloudwatchScrapeJobServiceArrayInput `pulumi:"services"`
	StackId  pulumi.StringInput                                 `pulumi:"stackId"`
}

func (LookupProviderAwsCloudwatchScrapeJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProviderAwsCloudwatchScrapeJobArgs)(nil)).Elem()
}

// A collection of values returned by getProviderAwsCloudwatchScrapeJob.
type LookupProviderAwsCloudwatchScrapeJobResultOutput struct{ *pulumi.OutputState }

func (LookupProviderAwsCloudwatchScrapeJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProviderAwsCloudwatchScrapeJobResult)(nil)).Elem()
}

func (o LookupProviderAwsCloudwatchScrapeJobResultOutput) ToLookupProviderAwsCloudwatchScrapeJobResultOutput() LookupProviderAwsCloudwatchScrapeJobResultOutput {
	return o
}

func (o LookupProviderAwsCloudwatchScrapeJobResultOutput) ToLookupProviderAwsCloudwatchScrapeJobResultOutputWithContext(ctx context.Context) LookupProviderAwsCloudwatchScrapeJobResultOutput {
	return o
}

// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this AWS CloudWatch Scrape Job. This can be provided by the `resourceId` attribute of the `cloudProvider.AwsAccount` resource.
func (o LookupProviderAwsCloudwatchScrapeJobResultOutput) AwsAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderAwsCloudwatchScrapeJobResult) string { return v.AwsAccountResourceId }).(pulumi.StringOutput)
}

// Zero or more configuration blocks to configure custom namespaces for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
func (o LookupProviderAwsCloudwatchScrapeJobResultOutput) CustomNamespaces() GetProviderAwsCloudwatchScrapeJobCustomNamespaceArrayOutput {
	return o.ApplyT(func(v LookupProviderAwsCloudwatchScrapeJobResult) []GetProviderAwsCloudwatchScrapeJobCustomNamespace {
		return v.CustomNamespaces
	}).(GetProviderAwsCloudwatchScrapeJobCustomNamespaceArrayOutput)
}

// When the AWS CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
func (o LookupProviderAwsCloudwatchScrapeJobResultOutput) DisabledReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderAwsCloudwatchScrapeJobResult) string { return v.DisabledReason }).(pulumi.StringOutput)
}

// Whether the AWS CloudWatch Scrape Job is enabled or not.
func (o LookupProviderAwsCloudwatchScrapeJobResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProviderAwsCloudwatchScrapeJobResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`.
func (o LookupProviderAwsCloudwatchScrapeJobResultOutput) ExportTags() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProviderAwsCloudwatchScrapeJobResult) bool { return v.ExportTags }).(pulumi.BoolOutput)
}

func (o LookupProviderAwsCloudwatchScrapeJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderAwsCloudwatchScrapeJobResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProviderAwsCloudwatchScrapeJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderAwsCloudwatchScrapeJobResult) string { return v.Name }).(pulumi.StringOutput)
}

// The set of AWS region names that this AWS CloudWatch Scrape Job is configured to scrape.
func (o LookupProviderAwsCloudwatchScrapeJobResultOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProviderAwsCloudwatchScrapeJobResult) []string { return v.Regions }).(pulumi.StringArrayOutput)
}

// When true, the `regions` attribute will be the set of regions configured in the override. When false, the `regions` attribute will be the set of regions belonging to the AWS Account resource that is associated with this AWS CloudWatch Scrape Job.
func (o LookupProviderAwsCloudwatchScrapeJobResultOutput) RegionsSubsetOverrideUsed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProviderAwsCloudwatchScrapeJobResult) bool { return v.RegionsSubsetOverrideUsed }).(pulumi.BoolOutput)
}

// The AWS ARN of the IAM role associated with the AWS Account resource that is being used by this AWS CloudWatch Scrape Job.
func (o LookupProviderAwsCloudwatchScrapeJobResultOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderAwsCloudwatchScrapeJobResult) string { return v.RoleArn }).(pulumi.StringOutput)
}

// One or more configuration blocks to dictate what this AWS CloudWatch Scrape Job should scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
func (o LookupProviderAwsCloudwatchScrapeJobResultOutput) Services() GetProviderAwsCloudwatchScrapeJobServiceArrayOutput {
	return o.ApplyT(func(v LookupProviderAwsCloudwatchScrapeJobResult) []GetProviderAwsCloudwatchScrapeJobService {
		return v.Services
	}).(GetProviderAwsCloudwatchScrapeJobServiceArrayOutput)
}

func (o LookupProviderAwsCloudwatchScrapeJobResultOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderAwsCloudwatchScrapeJobResult) string { return v.StackId }).(pulumi.StringOutput)
}

// A set of static labels to add to all metrics exported by this scrape job.
func (o LookupProviderAwsCloudwatchScrapeJobResultOutput) StaticLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProviderAwsCloudwatchScrapeJobResult) map[string]string { return v.StaticLabels }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProviderAwsCloudwatchScrapeJobResultOutput{})
}
