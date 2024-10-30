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

// * [Official documentation](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/)
// * [API documentation](https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#create-an-access-policy)
//
// Required access policy scopes:
//
// * accesspolicies:read
// * accesspolicies:write
// * accesspolicies:delete
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
//			current, err := cloud.GetOrganization(ctx, &cloud.GetOrganizationArgs{
//				Slug: pulumi.StringRef("<your org slug>"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			test, err := cloud.NewAccessPolicy(ctx, "test", &cloud.AccessPolicyArgs{
//				Region:      pulumi.String("us"),
//				Name:        pulumi.String("my-policy"),
//				DisplayName: pulumi.String("My Policy"),
//				Scopes: pulumi.StringArray{
//					pulumi.String("metrics:read"),
//					pulumi.String("logs:read"),
//				},
//				Realms: cloud.AccessPolicyRealmArray{
//					&cloud.AccessPolicyRealmArgs{
//						Type:       pulumi.String("org"),
//						Identifier: pulumi.String(current.Id),
//						LabelPolicies: cloud.AccessPolicyRealmLabelPolicyArray{
//							&cloud.AccessPolicyRealmLabelPolicyArgs{
//								Selector: pulumi.String("{namespace=\"default\"}"),
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloud.NewAccessPolicyToken(ctx, "test", &cloud.AccessPolicyTokenArgs{
//				Region:         pulumi.String("us"),
//				AccessPolicyId: test.PolicyId,
//				Name:           pulumi.String("my-policy-token"),
//				DisplayName:    pulumi.String("My Policy Token"),
//				ExpiresAt:      pulumi.String("2023-01-01T00:00:00Z"),
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
// $ pulumi import grafana:index/cloudAccessPolicy:CloudAccessPolicy name "{{ region }}:{{ policyId }}"
// ```
//
// Deprecated: grafana.index/cloudaccesspolicy.CloudAccessPolicy has been deprecated in favor of grafana.cloud/accesspolicy.AccessPolicy
type CloudAccessPolicy struct {
	pulumi.CustomResourceState

	// Creation date of the access policy.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Display name of the access policy. Defaults to the name.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Name of the access policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the access policy.
	PolicyId pulumi.StringOutput               `pulumi:"policyId"`
	Realms   CloudAccessPolicyRealmArrayOutput `pulumi:"realms"`
	// Region where the API is deployed. Generally where the stack is deployed. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	Region pulumi.StringOutput `pulumi:"region"`
	// Scopes of the access policy. See https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/#scopes for possible values.
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// Last update date of the access policy.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewCloudAccessPolicy registers a new resource with the given unique name, arguments, and options.
func NewCloudAccessPolicy(ctx *pulumi.Context,
	name string, args *CloudAccessPolicyArgs, opts ...pulumi.ResourceOption) (*CloudAccessPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Realms == nil {
		return nil, errors.New("invalid value for required argument 'Realms'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/cloudAccessPolicy:CloudAccessPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CloudAccessPolicy
	err := ctx.RegisterResource("grafana:index/cloudAccessPolicy:CloudAccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudAccessPolicy gets an existing CloudAccessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudAccessPolicyState, opts ...pulumi.ResourceOption) (*CloudAccessPolicy, error) {
	var resource CloudAccessPolicy
	err := ctx.ReadResource("grafana:index/cloudAccessPolicy:CloudAccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudAccessPolicy resources.
type cloudAccessPolicyState struct {
	// Creation date of the access policy.
	CreatedAt *string `pulumi:"createdAt"`
	// Display name of the access policy. Defaults to the name.
	DisplayName *string `pulumi:"displayName"`
	// Name of the access policy.
	Name *string `pulumi:"name"`
	// ID of the access policy.
	PolicyId *string                  `pulumi:"policyId"`
	Realms   []CloudAccessPolicyRealm `pulumi:"realms"`
	// Region where the API is deployed. Generally where the stack is deployed. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	Region *string `pulumi:"region"`
	// Scopes of the access policy. See https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/#scopes for possible values.
	Scopes []string `pulumi:"scopes"`
	// Last update date of the access policy.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type CloudAccessPolicyState struct {
	// Creation date of the access policy.
	CreatedAt pulumi.StringPtrInput
	// Display name of the access policy. Defaults to the name.
	DisplayName pulumi.StringPtrInput
	// Name of the access policy.
	Name pulumi.StringPtrInput
	// ID of the access policy.
	PolicyId pulumi.StringPtrInput
	Realms   CloudAccessPolicyRealmArrayInput
	// Region where the API is deployed. Generally where the stack is deployed. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	Region pulumi.StringPtrInput
	// Scopes of the access policy. See https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/#scopes for possible values.
	Scopes pulumi.StringArrayInput
	// Last update date of the access policy.
	UpdatedAt pulumi.StringPtrInput
}

func (CloudAccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudAccessPolicyState)(nil)).Elem()
}

type cloudAccessPolicyArgs struct {
	// Display name of the access policy. Defaults to the name.
	DisplayName *string `pulumi:"displayName"`
	// Name of the access policy.
	Name   *string                  `pulumi:"name"`
	Realms []CloudAccessPolicyRealm `pulumi:"realms"`
	// Region where the API is deployed. Generally where the stack is deployed. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	Region string `pulumi:"region"`
	// Scopes of the access policy. See https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/#scopes for possible values.
	Scopes []string `pulumi:"scopes"`
}

// The set of arguments for constructing a CloudAccessPolicy resource.
type CloudAccessPolicyArgs struct {
	// Display name of the access policy. Defaults to the name.
	DisplayName pulumi.StringPtrInput
	// Name of the access policy.
	Name   pulumi.StringPtrInput
	Realms CloudAccessPolicyRealmArrayInput
	// Region where the API is deployed. Generally where the stack is deployed. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	Region pulumi.StringInput
	// Scopes of the access policy. See https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/#scopes for possible values.
	Scopes pulumi.StringArrayInput
}

func (CloudAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudAccessPolicyArgs)(nil)).Elem()
}

type CloudAccessPolicyInput interface {
	pulumi.Input

	ToCloudAccessPolicyOutput() CloudAccessPolicyOutput
	ToCloudAccessPolicyOutputWithContext(ctx context.Context) CloudAccessPolicyOutput
}

func (*CloudAccessPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudAccessPolicy)(nil)).Elem()
}

func (i *CloudAccessPolicy) ToCloudAccessPolicyOutput() CloudAccessPolicyOutput {
	return i.ToCloudAccessPolicyOutputWithContext(context.Background())
}

func (i *CloudAccessPolicy) ToCloudAccessPolicyOutputWithContext(ctx context.Context) CloudAccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudAccessPolicyOutput)
}

// CloudAccessPolicyArrayInput is an input type that accepts CloudAccessPolicyArray and CloudAccessPolicyArrayOutput values.
// You can construct a concrete instance of `CloudAccessPolicyArrayInput` via:
//
//	CloudAccessPolicyArray{ CloudAccessPolicyArgs{...} }
type CloudAccessPolicyArrayInput interface {
	pulumi.Input

	ToCloudAccessPolicyArrayOutput() CloudAccessPolicyArrayOutput
	ToCloudAccessPolicyArrayOutputWithContext(context.Context) CloudAccessPolicyArrayOutput
}

type CloudAccessPolicyArray []CloudAccessPolicyInput

func (CloudAccessPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudAccessPolicy)(nil)).Elem()
}

func (i CloudAccessPolicyArray) ToCloudAccessPolicyArrayOutput() CloudAccessPolicyArrayOutput {
	return i.ToCloudAccessPolicyArrayOutputWithContext(context.Background())
}

func (i CloudAccessPolicyArray) ToCloudAccessPolicyArrayOutputWithContext(ctx context.Context) CloudAccessPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudAccessPolicyArrayOutput)
}

// CloudAccessPolicyMapInput is an input type that accepts CloudAccessPolicyMap and CloudAccessPolicyMapOutput values.
// You can construct a concrete instance of `CloudAccessPolicyMapInput` via:
//
//	CloudAccessPolicyMap{ "key": CloudAccessPolicyArgs{...} }
type CloudAccessPolicyMapInput interface {
	pulumi.Input

	ToCloudAccessPolicyMapOutput() CloudAccessPolicyMapOutput
	ToCloudAccessPolicyMapOutputWithContext(context.Context) CloudAccessPolicyMapOutput
}

type CloudAccessPolicyMap map[string]CloudAccessPolicyInput

func (CloudAccessPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudAccessPolicy)(nil)).Elem()
}

func (i CloudAccessPolicyMap) ToCloudAccessPolicyMapOutput() CloudAccessPolicyMapOutput {
	return i.ToCloudAccessPolicyMapOutputWithContext(context.Background())
}

func (i CloudAccessPolicyMap) ToCloudAccessPolicyMapOutputWithContext(ctx context.Context) CloudAccessPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudAccessPolicyMapOutput)
}

type CloudAccessPolicyOutput struct{ *pulumi.OutputState }

func (CloudAccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudAccessPolicy)(nil)).Elem()
}

func (o CloudAccessPolicyOutput) ToCloudAccessPolicyOutput() CloudAccessPolicyOutput {
	return o
}

func (o CloudAccessPolicyOutput) ToCloudAccessPolicyOutputWithContext(ctx context.Context) CloudAccessPolicyOutput {
	return o
}

// Creation date of the access policy.
func (o CloudAccessPolicyOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccessPolicy) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Display name of the access policy. Defaults to the name.
func (o CloudAccessPolicyOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudAccessPolicy) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Name of the access policy.
func (o CloudAccessPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccessPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the access policy.
func (o CloudAccessPolicyOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccessPolicy) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

func (o CloudAccessPolicyOutput) Realms() CloudAccessPolicyRealmArrayOutput {
	return o.ApplyT(func(v *CloudAccessPolicy) CloudAccessPolicyRealmArrayOutput { return v.Realms }).(CloudAccessPolicyRealmArrayOutput)
}

// Region where the API is deployed. Generally where the stack is deployed. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
func (o CloudAccessPolicyOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccessPolicy) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Scopes of the access policy. See https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/#scopes for possible values.
func (o CloudAccessPolicyOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudAccessPolicy) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

// Last update date of the access policy.
func (o CloudAccessPolicyOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccessPolicy) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type CloudAccessPolicyArrayOutput struct{ *pulumi.OutputState }

func (CloudAccessPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudAccessPolicy)(nil)).Elem()
}

func (o CloudAccessPolicyArrayOutput) ToCloudAccessPolicyArrayOutput() CloudAccessPolicyArrayOutput {
	return o
}

func (o CloudAccessPolicyArrayOutput) ToCloudAccessPolicyArrayOutputWithContext(ctx context.Context) CloudAccessPolicyArrayOutput {
	return o
}

func (o CloudAccessPolicyArrayOutput) Index(i pulumi.IntInput) CloudAccessPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudAccessPolicy {
		return vs[0].([]*CloudAccessPolicy)[vs[1].(int)]
	}).(CloudAccessPolicyOutput)
}

type CloudAccessPolicyMapOutput struct{ *pulumi.OutputState }

func (CloudAccessPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudAccessPolicy)(nil)).Elem()
}

func (o CloudAccessPolicyMapOutput) ToCloudAccessPolicyMapOutput() CloudAccessPolicyMapOutput {
	return o
}

func (o CloudAccessPolicyMapOutput) ToCloudAccessPolicyMapOutputWithContext(ctx context.Context) CloudAccessPolicyMapOutput {
	return o
}

func (o CloudAccessPolicyMapOutput) MapIndex(k pulumi.StringInput) CloudAccessPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudAccessPolicy {
		return vs[0].(map[string]*CloudAccessPolicy)[vs[1].(string)]
	}).(CloudAccessPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudAccessPolicyInput)(nil)).Elem(), &CloudAccessPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudAccessPolicyArrayInput)(nil)).Elem(), CloudAccessPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudAccessPolicyMapInput)(nil)).Elem(), CloudAccessPolicyMap{})
	pulumi.RegisterOutputType(CloudAccessPolicyOutput{})
	pulumi.RegisterOutputType(CloudAccessPolicyArrayOutput{})
	pulumi.RegisterOutputType(CloudAccessPolicyMapOutput{})
}
