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
//			_, err = cloud.NewProviderAwsAccount(ctx, "test", &cloud.ProviderAwsAccountArgs{
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
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
// $ pulumi import grafana:cloud/providerAwsAccount:ProviderAwsAccount name "{{ stack_id }}:{{ resource_id }}"
// ```
type ProviderAwsAccount struct {
	pulumi.CustomResourceState

	// An optional human-readable name for this AWS Account resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// A set of regions that this AWS Account resource applies to.
	Regions pulumi.StringArrayOutput `pulumi:"regions"`
	// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// An IAM Role ARN string to represent with this AWS Account resource.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	StackId pulumi.StringOutput `pulumi:"stackId"`
}

// NewProviderAwsAccount registers a new resource with the given unique name, arguments, and options.
func NewProviderAwsAccount(ctx *pulumi.Context,
	name string, args *ProviderAwsAccountArgs, opts ...pulumi.ResourceOption) (*ProviderAwsAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Regions == nil {
		return nil, errors.New("invalid value for required argument 'Regions'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProviderAwsAccount
	err := ctx.RegisterResource("grafana:cloud/providerAwsAccount:ProviderAwsAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProviderAwsAccount gets an existing ProviderAwsAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProviderAwsAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderAwsAccountState, opts ...pulumi.ResourceOption) (*ProviderAwsAccount, error) {
	var resource ProviderAwsAccount
	err := ctx.ReadResource("grafana:cloud/providerAwsAccount:ProviderAwsAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProviderAwsAccount resources.
type providerAwsAccountState struct {
	// An optional human-readable name for this AWS Account resource.
	Name *string `pulumi:"name"`
	// A set of regions that this AWS Account resource applies to.
	Regions []string `pulumi:"regions"`
	// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
	ResourceId *string `pulumi:"resourceId"`
	// An IAM Role ARN string to represent with this AWS Account resource.
	RoleArn *string `pulumi:"roleArn"`
	StackId *string `pulumi:"stackId"`
}

type ProviderAwsAccountState struct {
	// An optional human-readable name for this AWS Account resource.
	Name pulumi.StringPtrInput
	// A set of regions that this AWS Account resource applies to.
	Regions pulumi.StringArrayInput
	// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
	ResourceId pulumi.StringPtrInput
	// An IAM Role ARN string to represent with this AWS Account resource.
	RoleArn pulumi.StringPtrInput
	StackId pulumi.StringPtrInput
}

func (ProviderAwsAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerAwsAccountState)(nil)).Elem()
}

type providerAwsAccountArgs struct {
	// An optional human-readable name for this AWS Account resource.
	Name *string `pulumi:"name"`
	// A set of regions that this AWS Account resource applies to.
	Regions []string `pulumi:"regions"`
	// An IAM Role ARN string to represent with this AWS Account resource.
	RoleArn string `pulumi:"roleArn"`
	StackId string `pulumi:"stackId"`
}

// The set of arguments for constructing a ProviderAwsAccount resource.
type ProviderAwsAccountArgs struct {
	// An optional human-readable name for this AWS Account resource.
	Name pulumi.StringPtrInput
	// A set of regions that this AWS Account resource applies to.
	Regions pulumi.StringArrayInput
	// An IAM Role ARN string to represent with this AWS Account resource.
	RoleArn pulumi.StringInput
	StackId pulumi.StringInput
}

func (ProviderAwsAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerAwsAccountArgs)(nil)).Elem()
}

type ProviderAwsAccountInput interface {
	pulumi.Input

	ToProviderAwsAccountOutput() ProviderAwsAccountOutput
	ToProviderAwsAccountOutputWithContext(ctx context.Context) ProviderAwsAccountOutput
}

func (*ProviderAwsAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderAwsAccount)(nil)).Elem()
}

func (i *ProviderAwsAccount) ToProviderAwsAccountOutput() ProviderAwsAccountOutput {
	return i.ToProviderAwsAccountOutputWithContext(context.Background())
}

func (i *ProviderAwsAccount) ToProviderAwsAccountOutputWithContext(ctx context.Context) ProviderAwsAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAwsAccountOutput)
}

// ProviderAwsAccountArrayInput is an input type that accepts ProviderAwsAccountArray and ProviderAwsAccountArrayOutput values.
// You can construct a concrete instance of `ProviderAwsAccountArrayInput` via:
//
//	ProviderAwsAccountArray{ ProviderAwsAccountArgs{...} }
type ProviderAwsAccountArrayInput interface {
	pulumi.Input

	ToProviderAwsAccountArrayOutput() ProviderAwsAccountArrayOutput
	ToProviderAwsAccountArrayOutputWithContext(context.Context) ProviderAwsAccountArrayOutput
}

type ProviderAwsAccountArray []ProviderAwsAccountInput

func (ProviderAwsAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderAwsAccount)(nil)).Elem()
}

func (i ProviderAwsAccountArray) ToProviderAwsAccountArrayOutput() ProviderAwsAccountArrayOutput {
	return i.ToProviderAwsAccountArrayOutputWithContext(context.Background())
}

func (i ProviderAwsAccountArray) ToProviderAwsAccountArrayOutputWithContext(ctx context.Context) ProviderAwsAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAwsAccountArrayOutput)
}

// ProviderAwsAccountMapInput is an input type that accepts ProviderAwsAccountMap and ProviderAwsAccountMapOutput values.
// You can construct a concrete instance of `ProviderAwsAccountMapInput` via:
//
//	ProviderAwsAccountMap{ "key": ProviderAwsAccountArgs{...} }
type ProviderAwsAccountMapInput interface {
	pulumi.Input

	ToProviderAwsAccountMapOutput() ProviderAwsAccountMapOutput
	ToProviderAwsAccountMapOutputWithContext(context.Context) ProviderAwsAccountMapOutput
}

type ProviderAwsAccountMap map[string]ProviderAwsAccountInput

func (ProviderAwsAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderAwsAccount)(nil)).Elem()
}

func (i ProviderAwsAccountMap) ToProviderAwsAccountMapOutput() ProviderAwsAccountMapOutput {
	return i.ToProviderAwsAccountMapOutputWithContext(context.Background())
}

func (i ProviderAwsAccountMap) ToProviderAwsAccountMapOutputWithContext(ctx context.Context) ProviderAwsAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAwsAccountMapOutput)
}

type ProviderAwsAccountOutput struct{ *pulumi.OutputState }

func (ProviderAwsAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderAwsAccount)(nil)).Elem()
}

func (o ProviderAwsAccountOutput) ToProviderAwsAccountOutput() ProviderAwsAccountOutput {
	return o
}

func (o ProviderAwsAccountOutput) ToProviderAwsAccountOutputWithContext(ctx context.Context) ProviderAwsAccountOutput {
	return o
}

// An optional human-readable name for this AWS Account resource.
func (o ProviderAwsAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderAwsAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A set of regions that this AWS Account resource applies to.
func (o ProviderAwsAccountOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderAwsAccount) pulumi.StringArrayOutput { return v.Regions }).(pulumi.StringArrayOutput)
}

// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
func (o ProviderAwsAccountOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderAwsAccount) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// An IAM Role ARN string to represent with this AWS Account resource.
func (o ProviderAwsAccountOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderAwsAccount) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

func (o ProviderAwsAccountOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderAwsAccount) pulumi.StringOutput { return v.StackId }).(pulumi.StringOutput)
}

type ProviderAwsAccountArrayOutput struct{ *pulumi.OutputState }

func (ProviderAwsAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderAwsAccount)(nil)).Elem()
}

func (o ProviderAwsAccountArrayOutput) ToProviderAwsAccountArrayOutput() ProviderAwsAccountArrayOutput {
	return o
}

func (o ProviderAwsAccountArrayOutput) ToProviderAwsAccountArrayOutputWithContext(ctx context.Context) ProviderAwsAccountArrayOutput {
	return o
}

func (o ProviderAwsAccountArrayOutput) Index(i pulumi.IntInput) ProviderAwsAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProviderAwsAccount {
		return vs[0].([]*ProviderAwsAccount)[vs[1].(int)]
	}).(ProviderAwsAccountOutput)
}

type ProviderAwsAccountMapOutput struct{ *pulumi.OutputState }

func (ProviderAwsAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderAwsAccount)(nil)).Elem()
}

func (o ProviderAwsAccountMapOutput) ToProviderAwsAccountMapOutput() ProviderAwsAccountMapOutput {
	return o
}

func (o ProviderAwsAccountMapOutput) ToProviderAwsAccountMapOutputWithContext(ctx context.Context) ProviderAwsAccountMapOutput {
	return o
}

func (o ProviderAwsAccountMapOutput) MapIndex(k pulumi.StringInput) ProviderAwsAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProviderAwsAccount {
		return vs[0].(map[string]*ProviderAwsAccount)[vs[1].(string)]
	}).(ProviderAwsAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderAwsAccountInput)(nil)).Elem(), &ProviderAwsAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderAwsAccountArrayInput)(nil)).Elem(), ProviderAwsAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderAwsAccountMapInput)(nil)).Elem(), ProviderAwsAccountMap{})
	pulumi.RegisterOutputType(ProviderAwsAccountOutput{})
	pulumi.RegisterOutputType(ProviderAwsAccountArrayOutput{})
	pulumi.RegisterOutputType(ProviderAwsAccountMapOutput{})
}
