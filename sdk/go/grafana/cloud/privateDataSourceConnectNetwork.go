// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [Official documentation](https://grafana.com/docs/grafana-cloud/connect-externally-hosted/private-data-source-connect/)
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
//			current, err := cloud.LookupStack(ctx, &cloud.LookupStackArgs{
//				Slug: "<your slug>",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			test, err := cloud.NewPrivateDataSourceConnectNetwork(ctx, "test", &cloud.PrivateDataSourceConnectNetworkArgs{
//				Region:          pulumi.String("prod-us-east-0"),
//				Name:            pulumi.String("my-pdc"),
//				DisplayName:     pulumi.String("My PDC"),
//				StackIdentifier: pulumi.String(current.Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloud.NewPrivateDataSourceConnectNetworkToken(ctx, "test", &cloud.PrivateDataSourceConnectNetworkTokenArgs{
//				PdcNetworkId: test.PdcNetworkId,
//				Region:       test.Region,
//				Name:         pulumi.String("my-pdc-token"),
//				DisplayName:  pulumi.String("My PDC Token"),
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
// $ pulumi import grafana:cloud/privateDataSourceConnectNetwork:PrivateDataSourceConnectNetwork name "{{ region }}:{{ policyId }}"
// ```
type PrivateDataSourceConnectNetwork struct {
	pulumi.CustomResourceState

	// Creation date of the private data source connect network.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Display name of the PDC network. Defaults to the name.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Name of the PDC network.**Note:** The name must be lowercase and can contain hyphens or underscores. See full requirements here: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#request-body
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the private data source connect network.
	PdcNetworkId pulumi.StringOutput `pulumi:"pdcNetworkId"`
	// The region where your stack is deployed. Use the instances list API to get the region for your instance - use the regionSlug property: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-stacks
	Region pulumi.StringOutput `pulumi:"region"`
	// The identifier of the stack.
	StackIdentifier pulumi.StringOutput `pulumi:"stackIdentifier"`
	// Last update date of the private data source connect network.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewPrivateDataSourceConnectNetwork registers a new resource with the given unique name, arguments, and options.
func NewPrivateDataSourceConnectNetwork(ctx *pulumi.Context,
	name string, args *PrivateDataSourceConnectNetworkArgs, opts ...pulumi.ResourceOption) (*PrivateDataSourceConnectNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.StackIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'StackIdentifier'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PrivateDataSourceConnectNetwork
	err := ctx.RegisterResource("grafana:cloud/privateDataSourceConnectNetwork:PrivateDataSourceConnectNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateDataSourceConnectNetwork gets an existing PrivateDataSourceConnectNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateDataSourceConnectNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateDataSourceConnectNetworkState, opts ...pulumi.ResourceOption) (*PrivateDataSourceConnectNetwork, error) {
	var resource PrivateDataSourceConnectNetwork
	err := ctx.ReadResource("grafana:cloud/privateDataSourceConnectNetwork:PrivateDataSourceConnectNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateDataSourceConnectNetwork resources.
type privateDataSourceConnectNetworkState struct {
	// Creation date of the private data source connect network.
	CreatedAt *string `pulumi:"createdAt"`
	// Display name of the PDC network. Defaults to the name.
	DisplayName *string `pulumi:"displayName"`
	// Name of the PDC network.**Note:** The name must be lowercase and can contain hyphens or underscores. See full requirements here: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#request-body
	Name *string `pulumi:"name"`
	// ID of the private data source connect network.
	PdcNetworkId *string `pulumi:"pdcNetworkId"`
	// The region where your stack is deployed. Use the instances list API to get the region for your instance - use the regionSlug property: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-stacks
	Region *string `pulumi:"region"`
	// The identifier of the stack.
	StackIdentifier *string `pulumi:"stackIdentifier"`
	// Last update date of the private data source connect network.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type PrivateDataSourceConnectNetworkState struct {
	// Creation date of the private data source connect network.
	CreatedAt pulumi.StringPtrInput
	// Display name of the PDC network. Defaults to the name.
	DisplayName pulumi.StringPtrInput
	// Name of the PDC network.**Note:** The name must be lowercase and can contain hyphens or underscores. See full requirements here: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#request-body
	Name pulumi.StringPtrInput
	// ID of the private data source connect network.
	PdcNetworkId pulumi.StringPtrInput
	// The region where your stack is deployed. Use the instances list API to get the region for your instance - use the regionSlug property: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-stacks
	Region pulumi.StringPtrInput
	// The identifier of the stack.
	StackIdentifier pulumi.StringPtrInput
	// Last update date of the private data source connect network.
	UpdatedAt pulumi.StringPtrInput
}

func (PrivateDataSourceConnectNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDataSourceConnectNetworkState)(nil)).Elem()
}

type privateDataSourceConnectNetworkArgs struct {
	// Display name of the PDC network. Defaults to the name.
	DisplayName *string `pulumi:"displayName"`
	// Name of the PDC network.**Note:** The name must be lowercase and can contain hyphens or underscores. See full requirements here: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#request-body
	Name *string `pulumi:"name"`
	// The region where your stack is deployed. Use the instances list API to get the region for your instance - use the regionSlug property: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-stacks
	Region string `pulumi:"region"`
	// The identifier of the stack.
	StackIdentifier string `pulumi:"stackIdentifier"`
}

// The set of arguments for constructing a PrivateDataSourceConnectNetwork resource.
type PrivateDataSourceConnectNetworkArgs struct {
	// Display name of the PDC network. Defaults to the name.
	DisplayName pulumi.StringPtrInput
	// Name of the PDC network.**Note:** The name must be lowercase and can contain hyphens or underscores. See full requirements here: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#request-body
	Name pulumi.StringPtrInput
	// The region where your stack is deployed. Use the instances list API to get the region for your instance - use the regionSlug property: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-stacks
	Region pulumi.StringInput
	// The identifier of the stack.
	StackIdentifier pulumi.StringInput
}

func (PrivateDataSourceConnectNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDataSourceConnectNetworkArgs)(nil)).Elem()
}

type PrivateDataSourceConnectNetworkInput interface {
	pulumi.Input

	ToPrivateDataSourceConnectNetworkOutput() PrivateDataSourceConnectNetworkOutput
	ToPrivateDataSourceConnectNetworkOutputWithContext(ctx context.Context) PrivateDataSourceConnectNetworkOutput
}

func (*PrivateDataSourceConnectNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateDataSourceConnectNetwork)(nil)).Elem()
}

func (i *PrivateDataSourceConnectNetwork) ToPrivateDataSourceConnectNetworkOutput() PrivateDataSourceConnectNetworkOutput {
	return i.ToPrivateDataSourceConnectNetworkOutputWithContext(context.Background())
}

func (i *PrivateDataSourceConnectNetwork) ToPrivateDataSourceConnectNetworkOutputWithContext(ctx context.Context) PrivateDataSourceConnectNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDataSourceConnectNetworkOutput)
}

// PrivateDataSourceConnectNetworkArrayInput is an input type that accepts PrivateDataSourceConnectNetworkArray and PrivateDataSourceConnectNetworkArrayOutput values.
// You can construct a concrete instance of `PrivateDataSourceConnectNetworkArrayInput` via:
//
//	PrivateDataSourceConnectNetworkArray{ PrivateDataSourceConnectNetworkArgs{...} }
type PrivateDataSourceConnectNetworkArrayInput interface {
	pulumi.Input

	ToPrivateDataSourceConnectNetworkArrayOutput() PrivateDataSourceConnectNetworkArrayOutput
	ToPrivateDataSourceConnectNetworkArrayOutputWithContext(context.Context) PrivateDataSourceConnectNetworkArrayOutput
}

type PrivateDataSourceConnectNetworkArray []PrivateDataSourceConnectNetworkInput

func (PrivateDataSourceConnectNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateDataSourceConnectNetwork)(nil)).Elem()
}

func (i PrivateDataSourceConnectNetworkArray) ToPrivateDataSourceConnectNetworkArrayOutput() PrivateDataSourceConnectNetworkArrayOutput {
	return i.ToPrivateDataSourceConnectNetworkArrayOutputWithContext(context.Background())
}

func (i PrivateDataSourceConnectNetworkArray) ToPrivateDataSourceConnectNetworkArrayOutputWithContext(ctx context.Context) PrivateDataSourceConnectNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDataSourceConnectNetworkArrayOutput)
}

// PrivateDataSourceConnectNetworkMapInput is an input type that accepts PrivateDataSourceConnectNetworkMap and PrivateDataSourceConnectNetworkMapOutput values.
// You can construct a concrete instance of `PrivateDataSourceConnectNetworkMapInput` via:
//
//	PrivateDataSourceConnectNetworkMap{ "key": PrivateDataSourceConnectNetworkArgs{...} }
type PrivateDataSourceConnectNetworkMapInput interface {
	pulumi.Input

	ToPrivateDataSourceConnectNetworkMapOutput() PrivateDataSourceConnectNetworkMapOutput
	ToPrivateDataSourceConnectNetworkMapOutputWithContext(context.Context) PrivateDataSourceConnectNetworkMapOutput
}

type PrivateDataSourceConnectNetworkMap map[string]PrivateDataSourceConnectNetworkInput

func (PrivateDataSourceConnectNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateDataSourceConnectNetwork)(nil)).Elem()
}

func (i PrivateDataSourceConnectNetworkMap) ToPrivateDataSourceConnectNetworkMapOutput() PrivateDataSourceConnectNetworkMapOutput {
	return i.ToPrivateDataSourceConnectNetworkMapOutputWithContext(context.Background())
}

func (i PrivateDataSourceConnectNetworkMap) ToPrivateDataSourceConnectNetworkMapOutputWithContext(ctx context.Context) PrivateDataSourceConnectNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDataSourceConnectNetworkMapOutput)
}

type PrivateDataSourceConnectNetworkOutput struct{ *pulumi.OutputState }

func (PrivateDataSourceConnectNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateDataSourceConnectNetwork)(nil)).Elem()
}

func (o PrivateDataSourceConnectNetworkOutput) ToPrivateDataSourceConnectNetworkOutput() PrivateDataSourceConnectNetworkOutput {
	return o
}

func (o PrivateDataSourceConnectNetworkOutput) ToPrivateDataSourceConnectNetworkOutputWithContext(ctx context.Context) PrivateDataSourceConnectNetworkOutput {
	return o
}

// Creation date of the private data source connect network.
func (o PrivateDataSourceConnectNetworkOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDataSourceConnectNetwork) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Display name of the PDC network. Defaults to the name.
func (o PrivateDataSourceConnectNetworkOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateDataSourceConnectNetwork) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Name of the PDC network.**Note:** The name must be lowercase and can contain hyphens or underscores. See full requirements here: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#request-body
func (o PrivateDataSourceConnectNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDataSourceConnectNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the private data source connect network.
func (o PrivateDataSourceConnectNetworkOutput) PdcNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDataSourceConnectNetwork) pulumi.StringOutput { return v.PdcNetworkId }).(pulumi.StringOutput)
}

// The region where your stack is deployed. Use the instances list API to get the region for your instance - use the regionSlug property: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-stacks
func (o PrivateDataSourceConnectNetworkOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDataSourceConnectNetwork) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The identifier of the stack.
func (o PrivateDataSourceConnectNetworkOutput) StackIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDataSourceConnectNetwork) pulumi.StringOutput { return v.StackIdentifier }).(pulumi.StringOutput)
}

// Last update date of the private data source connect network.
func (o PrivateDataSourceConnectNetworkOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDataSourceConnectNetwork) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type PrivateDataSourceConnectNetworkArrayOutput struct{ *pulumi.OutputState }

func (PrivateDataSourceConnectNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateDataSourceConnectNetwork)(nil)).Elem()
}

func (o PrivateDataSourceConnectNetworkArrayOutput) ToPrivateDataSourceConnectNetworkArrayOutput() PrivateDataSourceConnectNetworkArrayOutput {
	return o
}

func (o PrivateDataSourceConnectNetworkArrayOutput) ToPrivateDataSourceConnectNetworkArrayOutputWithContext(ctx context.Context) PrivateDataSourceConnectNetworkArrayOutput {
	return o
}

func (o PrivateDataSourceConnectNetworkArrayOutput) Index(i pulumi.IntInput) PrivateDataSourceConnectNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrivateDataSourceConnectNetwork {
		return vs[0].([]*PrivateDataSourceConnectNetwork)[vs[1].(int)]
	}).(PrivateDataSourceConnectNetworkOutput)
}

type PrivateDataSourceConnectNetworkMapOutput struct{ *pulumi.OutputState }

func (PrivateDataSourceConnectNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateDataSourceConnectNetwork)(nil)).Elem()
}

func (o PrivateDataSourceConnectNetworkMapOutput) ToPrivateDataSourceConnectNetworkMapOutput() PrivateDataSourceConnectNetworkMapOutput {
	return o
}

func (o PrivateDataSourceConnectNetworkMapOutput) ToPrivateDataSourceConnectNetworkMapOutputWithContext(ctx context.Context) PrivateDataSourceConnectNetworkMapOutput {
	return o
}

func (o PrivateDataSourceConnectNetworkMapOutput) MapIndex(k pulumi.StringInput) PrivateDataSourceConnectNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrivateDataSourceConnectNetwork {
		return vs[0].(map[string]*PrivateDataSourceConnectNetwork)[vs[1].(string)]
	}).(PrivateDataSourceConnectNetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDataSourceConnectNetworkInput)(nil)).Elem(), &PrivateDataSourceConnectNetwork{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDataSourceConnectNetworkArrayInput)(nil)).Elem(), PrivateDataSourceConnectNetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDataSourceConnectNetworkMapInput)(nil)).Elem(), PrivateDataSourceConnectNetworkMap{})
	pulumi.RegisterOutputType(PrivateDataSourceConnectNetworkOutput{})
	pulumi.RegisterOutputType(PrivateDataSourceConnectNetworkArrayOutput{})
	pulumi.RegisterOutputType(PrivateDataSourceConnectNetworkMapOutput{})
}
