// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// Manages service account tokens of a Grafana Cloud stack using the Cloud API
// This can be used to bootstrap a management service account token for a new stack
//
// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api)
//
// Required access policy scopes:
//
// * stack-service-accounts:write
//
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
//			cloudSa, err := cloud.NewStackServiceAccount(ctx, "cloud_sa", &cloud.StackServiceAccountArgs{
//				StackSlug:  pulumi.String("<your stack slug>"),
//				Name:       pulumi.String("cloud service account"),
//				Role:       pulumi.String("Admin"),
//				IsDisabled: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			foo, err := cloud.NewStackServiceAccountToken(ctx, "foo", &cloud.StackServiceAccountTokenArgs{
//				Name:             pulumi.String("key_foo"),
//				ServiceAccountId: cloudSa.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("serviceAccountTokenFooKey", foo.Key)
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: grafana.index/cloudstackserviceaccounttoken.CloudStackServiceAccountToken has been deprecated in favor of grafana.cloud/stackserviceaccounttoken.StackServiceAccountToken
type CloudStackServiceAccountToken struct {
	pulumi.CustomResourceState

	Expiration       pulumi.StringOutput `pulumi:"expiration"`
	HasExpired       pulumi.BoolOutput   `pulumi:"hasExpired"`
	Key              pulumi.StringOutput `pulumi:"key"`
	Name             pulumi.StringOutput `pulumi:"name"`
	SecondsToLive    pulumi.IntPtrOutput `pulumi:"secondsToLive"`
	ServiceAccountId pulumi.StringOutput `pulumi:"serviceAccountId"`
	StackSlug        pulumi.StringOutput `pulumi:"stackSlug"`
}

// NewCloudStackServiceAccountToken registers a new resource with the given unique name, arguments, and options.
func NewCloudStackServiceAccountToken(ctx *pulumi.Context,
	name string, args *CloudStackServiceAccountTokenArgs, opts ...pulumi.ResourceOption) (*CloudStackServiceAccountToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceAccountId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountId'")
	}
	if args.StackSlug == nil {
		return nil, errors.New("invalid value for required argument 'StackSlug'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/cloudStackServiceAccountToken:CloudStackServiceAccountToken"),
		},
	})
	opts = append(opts, aliases)
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"key",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CloudStackServiceAccountToken
	err := ctx.RegisterResource("grafana:index/cloudStackServiceAccountToken:CloudStackServiceAccountToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudStackServiceAccountToken gets an existing CloudStackServiceAccountToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudStackServiceAccountToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudStackServiceAccountTokenState, opts ...pulumi.ResourceOption) (*CloudStackServiceAccountToken, error) {
	var resource CloudStackServiceAccountToken
	err := ctx.ReadResource("grafana:index/cloudStackServiceAccountToken:CloudStackServiceAccountToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudStackServiceAccountToken resources.
type cloudStackServiceAccountTokenState struct {
	Expiration       *string `pulumi:"expiration"`
	HasExpired       *bool   `pulumi:"hasExpired"`
	Key              *string `pulumi:"key"`
	Name             *string `pulumi:"name"`
	SecondsToLive    *int    `pulumi:"secondsToLive"`
	ServiceAccountId *string `pulumi:"serviceAccountId"`
	StackSlug        *string `pulumi:"stackSlug"`
}

type CloudStackServiceAccountTokenState struct {
	Expiration       pulumi.StringPtrInput
	HasExpired       pulumi.BoolPtrInput
	Key              pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	SecondsToLive    pulumi.IntPtrInput
	ServiceAccountId pulumi.StringPtrInput
	StackSlug        pulumi.StringPtrInput
}

func (CloudStackServiceAccountTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudStackServiceAccountTokenState)(nil)).Elem()
}

type cloudStackServiceAccountTokenArgs struct {
	Name             *string `pulumi:"name"`
	SecondsToLive    *int    `pulumi:"secondsToLive"`
	ServiceAccountId string  `pulumi:"serviceAccountId"`
	StackSlug        string  `pulumi:"stackSlug"`
}

// The set of arguments for constructing a CloudStackServiceAccountToken resource.
type CloudStackServiceAccountTokenArgs struct {
	Name             pulumi.StringPtrInput
	SecondsToLive    pulumi.IntPtrInput
	ServiceAccountId pulumi.StringInput
	StackSlug        pulumi.StringInput
}

func (CloudStackServiceAccountTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudStackServiceAccountTokenArgs)(nil)).Elem()
}

type CloudStackServiceAccountTokenInput interface {
	pulumi.Input

	ToCloudStackServiceAccountTokenOutput() CloudStackServiceAccountTokenOutput
	ToCloudStackServiceAccountTokenOutputWithContext(ctx context.Context) CloudStackServiceAccountTokenOutput
}

func (*CloudStackServiceAccountToken) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudStackServiceAccountToken)(nil)).Elem()
}

func (i *CloudStackServiceAccountToken) ToCloudStackServiceAccountTokenOutput() CloudStackServiceAccountTokenOutput {
	return i.ToCloudStackServiceAccountTokenOutputWithContext(context.Background())
}

func (i *CloudStackServiceAccountToken) ToCloudStackServiceAccountTokenOutputWithContext(ctx context.Context) CloudStackServiceAccountTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudStackServiceAccountTokenOutput)
}

// CloudStackServiceAccountTokenArrayInput is an input type that accepts CloudStackServiceAccountTokenArray and CloudStackServiceAccountTokenArrayOutput values.
// You can construct a concrete instance of `CloudStackServiceAccountTokenArrayInput` via:
//
//	CloudStackServiceAccountTokenArray{ CloudStackServiceAccountTokenArgs{...} }
type CloudStackServiceAccountTokenArrayInput interface {
	pulumi.Input

	ToCloudStackServiceAccountTokenArrayOutput() CloudStackServiceAccountTokenArrayOutput
	ToCloudStackServiceAccountTokenArrayOutputWithContext(context.Context) CloudStackServiceAccountTokenArrayOutput
}

type CloudStackServiceAccountTokenArray []CloudStackServiceAccountTokenInput

func (CloudStackServiceAccountTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudStackServiceAccountToken)(nil)).Elem()
}

func (i CloudStackServiceAccountTokenArray) ToCloudStackServiceAccountTokenArrayOutput() CloudStackServiceAccountTokenArrayOutput {
	return i.ToCloudStackServiceAccountTokenArrayOutputWithContext(context.Background())
}

func (i CloudStackServiceAccountTokenArray) ToCloudStackServiceAccountTokenArrayOutputWithContext(ctx context.Context) CloudStackServiceAccountTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudStackServiceAccountTokenArrayOutput)
}

// CloudStackServiceAccountTokenMapInput is an input type that accepts CloudStackServiceAccountTokenMap and CloudStackServiceAccountTokenMapOutput values.
// You can construct a concrete instance of `CloudStackServiceAccountTokenMapInput` via:
//
//	CloudStackServiceAccountTokenMap{ "key": CloudStackServiceAccountTokenArgs{...} }
type CloudStackServiceAccountTokenMapInput interface {
	pulumi.Input

	ToCloudStackServiceAccountTokenMapOutput() CloudStackServiceAccountTokenMapOutput
	ToCloudStackServiceAccountTokenMapOutputWithContext(context.Context) CloudStackServiceAccountTokenMapOutput
}

type CloudStackServiceAccountTokenMap map[string]CloudStackServiceAccountTokenInput

func (CloudStackServiceAccountTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudStackServiceAccountToken)(nil)).Elem()
}

func (i CloudStackServiceAccountTokenMap) ToCloudStackServiceAccountTokenMapOutput() CloudStackServiceAccountTokenMapOutput {
	return i.ToCloudStackServiceAccountTokenMapOutputWithContext(context.Background())
}

func (i CloudStackServiceAccountTokenMap) ToCloudStackServiceAccountTokenMapOutputWithContext(ctx context.Context) CloudStackServiceAccountTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudStackServiceAccountTokenMapOutput)
}

type CloudStackServiceAccountTokenOutput struct{ *pulumi.OutputState }

func (CloudStackServiceAccountTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudStackServiceAccountToken)(nil)).Elem()
}

func (o CloudStackServiceAccountTokenOutput) ToCloudStackServiceAccountTokenOutput() CloudStackServiceAccountTokenOutput {
	return o
}

func (o CloudStackServiceAccountTokenOutput) ToCloudStackServiceAccountTokenOutputWithContext(ctx context.Context) CloudStackServiceAccountTokenOutput {
	return o
}

func (o CloudStackServiceAccountTokenOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStackServiceAccountToken) pulumi.StringOutput { return v.Expiration }).(pulumi.StringOutput)
}

func (o CloudStackServiceAccountTokenOutput) HasExpired() pulumi.BoolOutput {
	return o.ApplyT(func(v *CloudStackServiceAccountToken) pulumi.BoolOutput { return v.HasExpired }).(pulumi.BoolOutput)
}

func (o CloudStackServiceAccountTokenOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStackServiceAccountToken) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

func (o CloudStackServiceAccountTokenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStackServiceAccountToken) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CloudStackServiceAccountTokenOutput) SecondsToLive() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CloudStackServiceAccountToken) pulumi.IntPtrOutput { return v.SecondsToLive }).(pulumi.IntPtrOutput)
}

func (o CloudStackServiceAccountTokenOutput) ServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStackServiceAccountToken) pulumi.StringOutput { return v.ServiceAccountId }).(pulumi.StringOutput)
}

func (o CloudStackServiceAccountTokenOutput) StackSlug() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudStackServiceAccountToken) pulumi.StringOutput { return v.StackSlug }).(pulumi.StringOutput)
}

type CloudStackServiceAccountTokenArrayOutput struct{ *pulumi.OutputState }

func (CloudStackServiceAccountTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudStackServiceAccountToken)(nil)).Elem()
}

func (o CloudStackServiceAccountTokenArrayOutput) ToCloudStackServiceAccountTokenArrayOutput() CloudStackServiceAccountTokenArrayOutput {
	return o
}

func (o CloudStackServiceAccountTokenArrayOutput) ToCloudStackServiceAccountTokenArrayOutputWithContext(ctx context.Context) CloudStackServiceAccountTokenArrayOutput {
	return o
}

func (o CloudStackServiceAccountTokenArrayOutput) Index(i pulumi.IntInput) CloudStackServiceAccountTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudStackServiceAccountToken {
		return vs[0].([]*CloudStackServiceAccountToken)[vs[1].(int)]
	}).(CloudStackServiceAccountTokenOutput)
}

type CloudStackServiceAccountTokenMapOutput struct{ *pulumi.OutputState }

func (CloudStackServiceAccountTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudStackServiceAccountToken)(nil)).Elem()
}

func (o CloudStackServiceAccountTokenMapOutput) ToCloudStackServiceAccountTokenMapOutput() CloudStackServiceAccountTokenMapOutput {
	return o
}

func (o CloudStackServiceAccountTokenMapOutput) ToCloudStackServiceAccountTokenMapOutputWithContext(ctx context.Context) CloudStackServiceAccountTokenMapOutput {
	return o
}

func (o CloudStackServiceAccountTokenMapOutput) MapIndex(k pulumi.StringInput) CloudStackServiceAccountTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudStackServiceAccountToken {
		return vs[0].(map[string]*CloudStackServiceAccountToken)[vs[1].(string)]
	}).(CloudStackServiceAccountTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudStackServiceAccountTokenInput)(nil)).Elem(), &CloudStackServiceAccountToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudStackServiceAccountTokenArrayInput)(nil)).Elem(), CloudStackServiceAccountTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudStackServiceAccountTokenMapInput)(nil)).Elem(), CloudStackServiceAccountTokenMap{})
	pulumi.RegisterOutputType(CloudStackServiceAccountTokenOutput{})
	pulumi.RegisterOutputType(CloudStackServiceAccountTokenArrayOutput{})
	pulumi.RegisterOutputType(CloudStackServiceAccountTokenMapOutput{})
}
