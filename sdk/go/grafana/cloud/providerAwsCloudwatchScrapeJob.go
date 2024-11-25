// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud

import (
	"context"
	"reflect"

	"errors"
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
//			testProviderAwsAccount, err := cloud.NewProviderAwsAccount(ctx, "test", &cloud.ProviderAwsAccountArgs{
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
//			_, err = cloud.NewProviderAwsCloudwatchScrapeJob(ctx, "test", &cloud.ProviderAwsCloudwatchScrapeJobArgs{
//				StackId:              pulumi.String(test.Id),
//				Name:                 pulumi.String("my-cloudwatch-scrape-job"),
//				AwsAccountResourceId: testProviderAwsAccount.ResourceId,
//				ExportTags:           pulumi.Bool(true),
//				Services: cloud.ProviderAwsCloudwatchScrapeJobServiceArray{
//					&cloud.ProviderAwsCloudwatchScrapeJobServiceArgs{
//						Name: pulumi.String("AWS/EC2"),
//						Metrics: cloud.ProviderAwsCloudwatchScrapeJobServiceMetricArray{
//							&cloud.ProviderAwsCloudwatchScrapeJobServiceMetricArgs{
//								Name: pulumi.String("CPUUtilization"),
//								Statistics: pulumi.StringArray{
//									pulumi.String("Average"),
//								},
//							},
//							&cloud.ProviderAwsCloudwatchScrapeJobServiceMetricArgs{
//								Name: pulumi.String("StatusCheckFailed"),
//								Statistics: pulumi.StringArray{
//									pulumi.String("Maximum"),
//								},
//							},
//						},
//						ScrapeIntervalSeconds: pulumi.Int(300),
//						ResourceDiscoveryTagFilters: cloud.ProviderAwsCloudwatchScrapeJobServiceResourceDiscoveryTagFilterArray{
//							&cloud.ProviderAwsCloudwatchScrapeJobServiceResourceDiscoveryTagFilterArgs{
//								Key:   pulumi.String("k8s.io/cluster-autoscaler/enabled"),
//								Value: pulumi.String("true"),
//							},
//						},
//						TagsToAddToMetrics: pulumi.StringArray{
//							pulumi.String("eks:cluster-name"),
//						},
//					},
//				},
//				CustomNamespaces: cloud.ProviderAwsCloudwatchScrapeJobCustomNamespaceArray{
//					&cloud.ProviderAwsCloudwatchScrapeJobCustomNamespaceArgs{
//						Name: pulumi.String("CoolApp"),
//						Metrics: cloud.ProviderAwsCloudwatchScrapeJobCustomNamespaceMetricArray{
//							&cloud.ProviderAwsCloudwatchScrapeJobCustomNamespaceMetricArgs{
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
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
// $ pulumi import grafana:cloud/providerAwsCloudwatchScrapeJob:ProviderAwsCloudwatchScrapeJob name "{{ stack_id }}:{{ name }}"
// ```
type ProviderAwsCloudwatchScrapeJob struct {
	pulumi.CustomResourceState

	// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resourceId` attribute of the `cloud.ProviderAwsAccount` resource.
	AwsAccountResourceId pulumi.StringOutput `pulumi:"awsAccountResourceId"`
	// Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	CustomNamespaces ProviderAwsCloudwatchScrapeJobCustomNamespaceArrayOutput `pulumi:"customNamespaces"`
	// When the CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
	DisabledReason pulumi.StringOutput `pulumi:"disabledReason"`
	// Whether the CloudWatch Scrape Job is enabled or not.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`.
	ExportTags pulumi.BoolOutput   `pulumi:"exportTags"`
	Name       pulumi.StringOutput `pulumi:"name"`
	// A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
	RegionsSubsetOverrides pulumi.StringArrayOutput `pulumi:"regionsSubsetOverrides"`
	// One or more configuration blocks to configure AWS services for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	Services ProviderAwsCloudwatchScrapeJobServiceArrayOutput `pulumi:"services"`
	StackId  pulumi.StringOutput                              `pulumi:"stackId"`
}

// NewProviderAwsCloudwatchScrapeJob registers a new resource with the given unique name, arguments, and options.
func NewProviderAwsCloudwatchScrapeJob(ctx *pulumi.Context,
	name string, args *ProviderAwsCloudwatchScrapeJobArgs, opts ...pulumi.ResourceOption) (*ProviderAwsCloudwatchScrapeJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AwsAccountResourceId == nil {
		return nil, errors.New("invalid value for required argument 'AwsAccountResourceId'")
	}
	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProviderAwsCloudwatchScrapeJob
	err := ctx.RegisterResource("grafana:cloud/providerAwsCloudwatchScrapeJob:ProviderAwsCloudwatchScrapeJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProviderAwsCloudwatchScrapeJob gets an existing ProviderAwsCloudwatchScrapeJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProviderAwsCloudwatchScrapeJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderAwsCloudwatchScrapeJobState, opts ...pulumi.ResourceOption) (*ProviderAwsCloudwatchScrapeJob, error) {
	var resource ProviderAwsCloudwatchScrapeJob
	err := ctx.ReadResource("grafana:cloud/providerAwsCloudwatchScrapeJob:ProviderAwsCloudwatchScrapeJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProviderAwsCloudwatchScrapeJob resources.
type providerAwsCloudwatchScrapeJobState struct {
	// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resourceId` attribute of the `cloud.ProviderAwsAccount` resource.
	AwsAccountResourceId *string `pulumi:"awsAccountResourceId"`
	// Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	CustomNamespaces []ProviderAwsCloudwatchScrapeJobCustomNamespace `pulumi:"customNamespaces"`
	// When the CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
	DisabledReason *string `pulumi:"disabledReason"`
	// Whether the CloudWatch Scrape Job is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`.
	ExportTags *bool   `pulumi:"exportTags"`
	Name       *string `pulumi:"name"`
	// A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
	RegionsSubsetOverrides []string `pulumi:"regionsSubsetOverrides"`
	// One or more configuration blocks to configure AWS services for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	Services []ProviderAwsCloudwatchScrapeJobService `pulumi:"services"`
	StackId  *string                                 `pulumi:"stackId"`
}

type ProviderAwsCloudwatchScrapeJobState struct {
	// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resourceId` attribute of the `cloud.ProviderAwsAccount` resource.
	AwsAccountResourceId pulumi.StringPtrInput
	// Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	CustomNamespaces ProviderAwsCloudwatchScrapeJobCustomNamespaceArrayInput
	// When the CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
	DisabledReason pulumi.StringPtrInput
	// Whether the CloudWatch Scrape Job is enabled or not.
	Enabled pulumi.BoolPtrInput
	// When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`.
	ExportTags pulumi.BoolPtrInput
	Name       pulumi.StringPtrInput
	// A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
	RegionsSubsetOverrides pulumi.StringArrayInput
	// One or more configuration blocks to configure AWS services for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	Services ProviderAwsCloudwatchScrapeJobServiceArrayInput
	StackId  pulumi.StringPtrInput
}

func (ProviderAwsCloudwatchScrapeJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerAwsCloudwatchScrapeJobState)(nil)).Elem()
}

type providerAwsCloudwatchScrapeJobArgs struct {
	// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resourceId` attribute of the `cloud.ProviderAwsAccount` resource.
	AwsAccountResourceId string `pulumi:"awsAccountResourceId"`
	// Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	CustomNamespaces []ProviderAwsCloudwatchScrapeJobCustomNamespace `pulumi:"customNamespaces"`
	// Whether the CloudWatch Scrape Job is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`.
	ExportTags *bool   `pulumi:"exportTags"`
	Name       *string `pulumi:"name"`
	// A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
	RegionsSubsetOverrides []string `pulumi:"regionsSubsetOverrides"`
	// One or more configuration blocks to configure AWS services for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	Services []ProviderAwsCloudwatchScrapeJobService `pulumi:"services"`
	StackId  string                                  `pulumi:"stackId"`
}

// The set of arguments for constructing a ProviderAwsCloudwatchScrapeJob resource.
type ProviderAwsCloudwatchScrapeJobArgs struct {
	// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resourceId` attribute of the `cloud.ProviderAwsAccount` resource.
	AwsAccountResourceId pulumi.StringInput
	// Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	CustomNamespaces ProviderAwsCloudwatchScrapeJobCustomNamespaceArrayInput
	// Whether the CloudWatch Scrape Job is enabled or not.
	Enabled pulumi.BoolPtrInput
	// When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`.
	ExportTags pulumi.BoolPtrInput
	Name       pulumi.StringPtrInput
	// A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
	RegionsSubsetOverrides pulumi.StringArrayInput
	// One or more configuration blocks to configure AWS services for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	Services ProviderAwsCloudwatchScrapeJobServiceArrayInput
	StackId  pulumi.StringInput
}

func (ProviderAwsCloudwatchScrapeJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerAwsCloudwatchScrapeJobArgs)(nil)).Elem()
}

type ProviderAwsCloudwatchScrapeJobInput interface {
	pulumi.Input

	ToProviderAwsCloudwatchScrapeJobOutput() ProviderAwsCloudwatchScrapeJobOutput
	ToProviderAwsCloudwatchScrapeJobOutputWithContext(ctx context.Context) ProviderAwsCloudwatchScrapeJobOutput
}

func (*ProviderAwsCloudwatchScrapeJob) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderAwsCloudwatchScrapeJob)(nil)).Elem()
}

func (i *ProviderAwsCloudwatchScrapeJob) ToProviderAwsCloudwatchScrapeJobOutput() ProviderAwsCloudwatchScrapeJobOutput {
	return i.ToProviderAwsCloudwatchScrapeJobOutputWithContext(context.Background())
}

func (i *ProviderAwsCloudwatchScrapeJob) ToProviderAwsCloudwatchScrapeJobOutputWithContext(ctx context.Context) ProviderAwsCloudwatchScrapeJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAwsCloudwatchScrapeJobOutput)
}

// ProviderAwsCloudwatchScrapeJobArrayInput is an input type that accepts ProviderAwsCloudwatchScrapeJobArray and ProviderAwsCloudwatchScrapeJobArrayOutput values.
// You can construct a concrete instance of `ProviderAwsCloudwatchScrapeJobArrayInput` via:
//
//	ProviderAwsCloudwatchScrapeJobArray{ ProviderAwsCloudwatchScrapeJobArgs{...} }
type ProviderAwsCloudwatchScrapeJobArrayInput interface {
	pulumi.Input

	ToProviderAwsCloudwatchScrapeJobArrayOutput() ProviderAwsCloudwatchScrapeJobArrayOutput
	ToProviderAwsCloudwatchScrapeJobArrayOutputWithContext(context.Context) ProviderAwsCloudwatchScrapeJobArrayOutput
}

type ProviderAwsCloudwatchScrapeJobArray []ProviderAwsCloudwatchScrapeJobInput

func (ProviderAwsCloudwatchScrapeJobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderAwsCloudwatchScrapeJob)(nil)).Elem()
}

func (i ProviderAwsCloudwatchScrapeJobArray) ToProviderAwsCloudwatchScrapeJobArrayOutput() ProviderAwsCloudwatchScrapeJobArrayOutput {
	return i.ToProviderAwsCloudwatchScrapeJobArrayOutputWithContext(context.Background())
}

func (i ProviderAwsCloudwatchScrapeJobArray) ToProviderAwsCloudwatchScrapeJobArrayOutputWithContext(ctx context.Context) ProviderAwsCloudwatchScrapeJobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAwsCloudwatchScrapeJobArrayOutput)
}

// ProviderAwsCloudwatchScrapeJobMapInput is an input type that accepts ProviderAwsCloudwatchScrapeJobMap and ProviderAwsCloudwatchScrapeJobMapOutput values.
// You can construct a concrete instance of `ProviderAwsCloudwatchScrapeJobMapInput` via:
//
//	ProviderAwsCloudwatchScrapeJobMap{ "key": ProviderAwsCloudwatchScrapeJobArgs{...} }
type ProviderAwsCloudwatchScrapeJobMapInput interface {
	pulumi.Input

	ToProviderAwsCloudwatchScrapeJobMapOutput() ProviderAwsCloudwatchScrapeJobMapOutput
	ToProviderAwsCloudwatchScrapeJobMapOutputWithContext(context.Context) ProviderAwsCloudwatchScrapeJobMapOutput
}

type ProviderAwsCloudwatchScrapeJobMap map[string]ProviderAwsCloudwatchScrapeJobInput

func (ProviderAwsCloudwatchScrapeJobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderAwsCloudwatchScrapeJob)(nil)).Elem()
}

func (i ProviderAwsCloudwatchScrapeJobMap) ToProviderAwsCloudwatchScrapeJobMapOutput() ProviderAwsCloudwatchScrapeJobMapOutput {
	return i.ToProviderAwsCloudwatchScrapeJobMapOutputWithContext(context.Background())
}

func (i ProviderAwsCloudwatchScrapeJobMap) ToProviderAwsCloudwatchScrapeJobMapOutputWithContext(ctx context.Context) ProviderAwsCloudwatchScrapeJobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAwsCloudwatchScrapeJobMapOutput)
}

type ProviderAwsCloudwatchScrapeJobOutput struct{ *pulumi.OutputState }

func (ProviderAwsCloudwatchScrapeJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderAwsCloudwatchScrapeJob)(nil)).Elem()
}

func (o ProviderAwsCloudwatchScrapeJobOutput) ToProviderAwsCloudwatchScrapeJobOutput() ProviderAwsCloudwatchScrapeJobOutput {
	return o
}

func (o ProviderAwsCloudwatchScrapeJobOutput) ToProviderAwsCloudwatchScrapeJobOutputWithContext(ctx context.Context) ProviderAwsCloudwatchScrapeJobOutput {
	return o
}

// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resourceId` attribute of the `cloud.ProviderAwsAccount` resource.
func (o ProviderAwsCloudwatchScrapeJobOutput) AwsAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderAwsCloudwatchScrapeJob) pulumi.StringOutput { return v.AwsAccountResourceId }).(pulumi.StringOutput)
}

// Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
func (o ProviderAwsCloudwatchScrapeJobOutput) CustomNamespaces() ProviderAwsCloudwatchScrapeJobCustomNamespaceArrayOutput {
	return o.ApplyT(func(v *ProviderAwsCloudwatchScrapeJob) ProviderAwsCloudwatchScrapeJobCustomNamespaceArrayOutput {
		return v.CustomNamespaces
	}).(ProviderAwsCloudwatchScrapeJobCustomNamespaceArrayOutput)
}

// When the CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
func (o ProviderAwsCloudwatchScrapeJobOutput) DisabledReason() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderAwsCloudwatchScrapeJob) pulumi.StringOutput { return v.DisabledReason }).(pulumi.StringOutput)
}

// Whether the CloudWatch Scrape Job is enabled or not.
func (o ProviderAwsCloudwatchScrapeJobOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProviderAwsCloudwatchScrapeJob) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`.
func (o ProviderAwsCloudwatchScrapeJobOutput) ExportTags() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProviderAwsCloudwatchScrapeJob) pulumi.BoolOutput { return v.ExportTags }).(pulumi.BoolOutput)
}

func (o ProviderAwsCloudwatchScrapeJobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderAwsCloudwatchScrapeJob) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
func (o ProviderAwsCloudwatchScrapeJobOutput) RegionsSubsetOverrides() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderAwsCloudwatchScrapeJob) pulumi.StringArrayOutput { return v.RegionsSubsetOverrides }).(pulumi.StringArrayOutput)
}

// One or more configuration blocks to configure AWS services for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
func (o ProviderAwsCloudwatchScrapeJobOutput) Services() ProviderAwsCloudwatchScrapeJobServiceArrayOutput {
	return o.ApplyT(func(v *ProviderAwsCloudwatchScrapeJob) ProviderAwsCloudwatchScrapeJobServiceArrayOutput {
		return v.Services
	}).(ProviderAwsCloudwatchScrapeJobServiceArrayOutput)
}

func (o ProviderAwsCloudwatchScrapeJobOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderAwsCloudwatchScrapeJob) pulumi.StringOutput { return v.StackId }).(pulumi.StringOutput)
}

type ProviderAwsCloudwatchScrapeJobArrayOutput struct{ *pulumi.OutputState }

func (ProviderAwsCloudwatchScrapeJobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderAwsCloudwatchScrapeJob)(nil)).Elem()
}

func (o ProviderAwsCloudwatchScrapeJobArrayOutput) ToProviderAwsCloudwatchScrapeJobArrayOutput() ProviderAwsCloudwatchScrapeJobArrayOutput {
	return o
}

func (o ProviderAwsCloudwatchScrapeJobArrayOutput) ToProviderAwsCloudwatchScrapeJobArrayOutputWithContext(ctx context.Context) ProviderAwsCloudwatchScrapeJobArrayOutput {
	return o
}

func (o ProviderAwsCloudwatchScrapeJobArrayOutput) Index(i pulumi.IntInput) ProviderAwsCloudwatchScrapeJobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProviderAwsCloudwatchScrapeJob {
		return vs[0].([]*ProviderAwsCloudwatchScrapeJob)[vs[1].(int)]
	}).(ProviderAwsCloudwatchScrapeJobOutput)
}

type ProviderAwsCloudwatchScrapeJobMapOutput struct{ *pulumi.OutputState }

func (ProviderAwsCloudwatchScrapeJobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderAwsCloudwatchScrapeJob)(nil)).Elem()
}

func (o ProviderAwsCloudwatchScrapeJobMapOutput) ToProviderAwsCloudwatchScrapeJobMapOutput() ProviderAwsCloudwatchScrapeJobMapOutput {
	return o
}

func (o ProviderAwsCloudwatchScrapeJobMapOutput) ToProviderAwsCloudwatchScrapeJobMapOutputWithContext(ctx context.Context) ProviderAwsCloudwatchScrapeJobMapOutput {
	return o
}

func (o ProviderAwsCloudwatchScrapeJobMapOutput) MapIndex(k pulumi.StringInput) ProviderAwsCloudwatchScrapeJobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProviderAwsCloudwatchScrapeJob {
		return vs[0].(map[string]*ProviderAwsCloudwatchScrapeJob)[vs[1].(string)]
	}).(ProviderAwsCloudwatchScrapeJobOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderAwsCloudwatchScrapeJobInput)(nil)).Elem(), &ProviderAwsCloudwatchScrapeJob{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderAwsCloudwatchScrapeJobArrayInput)(nil)).Elem(), ProviderAwsCloudwatchScrapeJobArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderAwsCloudwatchScrapeJobMapInput)(nil)).Elem(), ProviderAwsCloudwatchScrapeJobMap{})
	pulumi.RegisterOutputType(ProviderAwsCloudwatchScrapeJobOutput{})
	pulumi.RegisterOutputType(ProviderAwsCloudwatchScrapeJobArrayOutput{})
	pulumi.RegisterOutputType(ProviderAwsCloudwatchScrapeJobMapOutput{})
}
