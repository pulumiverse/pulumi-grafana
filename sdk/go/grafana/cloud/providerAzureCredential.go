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
// ## Import
//
// ```sh
// $ pulumi import grafana:cloud/providerAzureCredential:ProviderAzureCredential name "{{ stack_id }}:{{ resource_id }}"
// ```
//
// Deprecated: grafana.cloud/providerazurecredential.ProviderAzureCredential has been deprecated in favor of grafana.cloudprovider/azurecredential.AzureCredential
type ProviderAzureCredential struct {
	pulumi.CustomResourceState

	// The list of auto discovery configurations.
	AutoDiscoveryConfigurations ProviderAzureCredentialAutoDiscoveryConfigurationArrayOutput `pulumi:"autoDiscoveryConfigurations"`
	// The client ID of the Azure Credential.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The client secret of the Azure Credential.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// The name of the Azure Credential.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilters ProviderAzureCredentialResourceDiscoveryTagFilterArrayOutput `pulumi:"resourceDiscoveryTagFilters"`
	// The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The list of resource tags to add to metrics.
	ResourceTagsToAddToMetrics pulumi.StringArrayOutput `pulumi:"resourceTagsToAddToMetrics"`
	StackId                    pulumi.StringOutput      `pulumi:"stackId"`
	// The tenant ID of the Azure Credential.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewProviderAzureCredential registers a new resource with the given unique name, arguments, and options.
func NewProviderAzureCredential(ctx *pulumi.Context,
	name string, args *ProviderAzureCredentialArgs, opts ...pulumi.ResourceOption) (*ProviderAzureCredential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'ClientSecret'")
	}
	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	if args.ClientSecret != nil {
		args.ClientSecret = pulumi.ToSecret(args.ClientSecret).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clientSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProviderAzureCredential
	err := ctx.RegisterResource("grafana:cloud/providerAzureCredential:ProviderAzureCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProviderAzureCredential gets an existing ProviderAzureCredential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProviderAzureCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderAzureCredentialState, opts ...pulumi.ResourceOption) (*ProviderAzureCredential, error) {
	var resource ProviderAzureCredential
	err := ctx.ReadResource("grafana:cloud/providerAzureCredential:ProviderAzureCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProviderAzureCredential resources.
type providerAzureCredentialState struct {
	// The list of auto discovery configurations.
	AutoDiscoveryConfigurations []ProviderAzureCredentialAutoDiscoveryConfiguration `pulumi:"autoDiscoveryConfigurations"`
	// The client ID of the Azure Credential.
	ClientId *string `pulumi:"clientId"`
	// The client secret of the Azure Credential.
	ClientSecret *string `pulumi:"clientSecret"`
	// The name of the Azure Credential.
	Name *string `pulumi:"name"`
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilters []ProviderAzureCredentialResourceDiscoveryTagFilter `pulumi:"resourceDiscoveryTagFilters"`
	// The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
	ResourceId *string `pulumi:"resourceId"`
	// The list of resource tags to add to metrics.
	ResourceTagsToAddToMetrics []string `pulumi:"resourceTagsToAddToMetrics"`
	StackId                    *string  `pulumi:"stackId"`
	// The tenant ID of the Azure Credential.
	TenantId *string `pulumi:"tenantId"`
}

type ProviderAzureCredentialState struct {
	// The list of auto discovery configurations.
	AutoDiscoveryConfigurations ProviderAzureCredentialAutoDiscoveryConfigurationArrayInput
	// The client ID of the Azure Credential.
	ClientId pulumi.StringPtrInput
	// The client secret of the Azure Credential.
	ClientSecret pulumi.StringPtrInput
	// The name of the Azure Credential.
	Name pulumi.StringPtrInput
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilters ProviderAzureCredentialResourceDiscoveryTagFilterArrayInput
	// The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
	ResourceId pulumi.StringPtrInput
	// The list of resource tags to add to metrics.
	ResourceTagsToAddToMetrics pulumi.StringArrayInput
	StackId                    pulumi.StringPtrInput
	// The tenant ID of the Azure Credential.
	TenantId pulumi.StringPtrInput
}

func (ProviderAzureCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerAzureCredentialState)(nil)).Elem()
}

type providerAzureCredentialArgs struct {
	// The list of auto discovery configurations.
	AutoDiscoveryConfigurations []ProviderAzureCredentialAutoDiscoveryConfiguration `pulumi:"autoDiscoveryConfigurations"`
	// The client ID of the Azure Credential.
	ClientId string `pulumi:"clientId"`
	// The client secret of the Azure Credential.
	ClientSecret string `pulumi:"clientSecret"`
	// The name of the Azure Credential.
	Name *string `pulumi:"name"`
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilters []ProviderAzureCredentialResourceDiscoveryTagFilter `pulumi:"resourceDiscoveryTagFilters"`
	// The list of resource tags to add to metrics.
	ResourceTagsToAddToMetrics []string `pulumi:"resourceTagsToAddToMetrics"`
	StackId                    string   `pulumi:"stackId"`
	// The tenant ID of the Azure Credential.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a ProviderAzureCredential resource.
type ProviderAzureCredentialArgs struct {
	// The list of auto discovery configurations.
	AutoDiscoveryConfigurations ProviderAzureCredentialAutoDiscoveryConfigurationArrayInput
	// The client ID of the Azure Credential.
	ClientId pulumi.StringInput
	// The client secret of the Azure Credential.
	ClientSecret pulumi.StringInput
	// The name of the Azure Credential.
	Name pulumi.StringPtrInput
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilters ProviderAzureCredentialResourceDiscoveryTagFilterArrayInput
	// The list of resource tags to add to metrics.
	ResourceTagsToAddToMetrics pulumi.StringArrayInput
	StackId                    pulumi.StringInput
	// The tenant ID of the Azure Credential.
	TenantId pulumi.StringInput
}

func (ProviderAzureCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerAzureCredentialArgs)(nil)).Elem()
}

type ProviderAzureCredentialInput interface {
	pulumi.Input

	ToProviderAzureCredentialOutput() ProviderAzureCredentialOutput
	ToProviderAzureCredentialOutputWithContext(ctx context.Context) ProviderAzureCredentialOutput
}

func (*ProviderAzureCredential) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderAzureCredential)(nil)).Elem()
}

func (i *ProviderAzureCredential) ToProviderAzureCredentialOutput() ProviderAzureCredentialOutput {
	return i.ToProviderAzureCredentialOutputWithContext(context.Background())
}

func (i *ProviderAzureCredential) ToProviderAzureCredentialOutputWithContext(ctx context.Context) ProviderAzureCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAzureCredentialOutput)
}

// ProviderAzureCredentialArrayInput is an input type that accepts ProviderAzureCredentialArray and ProviderAzureCredentialArrayOutput values.
// You can construct a concrete instance of `ProviderAzureCredentialArrayInput` via:
//
//	ProviderAzureCredentialArray{ ProviderAzureCredentialArgs{...} }
type ProviderAzureCredentialArrayInput interface {
	pulumi.Input

	ToProviderAzureCredentialArrayOutput() ProviderAzureCredentialArrayOutput
	ToProviderAzureCredentialArrayOutputWithContext(context.Context) ProviderAzureCredentialArrayOutput
}

type ProviderAzureCredentialArray []ProviderAzureCredentialInput

func (ProviderAzureCredentialArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderAzureCredential)(nil)).Elem()
}

func (i ProviderAzureCredentialArray) ToProviderAzureCredentialArrayOutput() ProviderAzureCredentialArrayOutput {
	return i.ToProviderAzureCredentialArrayOutputWithContext(context.Background())
}

func (i ProviderAzureCredentialArray) ToProviderAzureCredentialArrayOutputWithContext(ctx context.Context) ProviderAzureCredentialArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAzureCredentialArrayOutput)
}

// ProviderAzureCredentialMapInput is an input type that accepts ProviderAzureCredentialMap and ProviderAzureCredentialMapOutput values.
// You can construct a concrete instance of `ProviderAzureCredentialMapInput` via:
//
//	ProviderAzureCredentialMap{ "key": ProviderAzureCredentialArgs{...} }
type ProviderAzureCredentialMapInput interface {
	pulumi.Input

	ToProviderAzureCredentialMapOutput() ProviderAzureCredentialMapOutput
	ToProviderAzureCredentialMapOutputWithContext(context.Context) ProviderAzureCredentialMapOutput
}

type ProviderAzureCredentialMap map[string]ProviderAzureCredentialInput

func (ProviderAzureCredentialMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderAzureCredential)(nil)).Elem()
}

func (i ProviderAzureCredentialMap) ToProviderAzureCredentialMapOutput() ProviderAzureCredentialMapOutput {
	return i.ToProviderAzureCredentialMapOutputWithContext(context.Background())
}

func (i ProviderAzureCredentialMap) ToProviderAzureCredentialMapOutputWithContext(ctx context.Context) ProviderAzureCredentialMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAzureCredentialMapOutput)
}

type ProviderAzureCredentialOutput struct{ *pulumi.OutputState }

func (ProviderAzureCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderAzureCredential)(nil)).Elem()
}

func (o ProviderAzureCredentialOutput) ToProviderAzureCredentialOutput() ProviderAzureCredentialOutput {
	return o
}

func (o ProviderAzureCredentialOutput) ToProviderAzureCredentialOutputWithContext(ctx context.Context) ProviderAzureCredentialOutput {
	return o
}

// The list of auto discovery configurations.
func (o ProviderAzureCredentialOutput) AutoDiscoveryConfigurations() ProviderAzureCredentialAutoDiscoveryConfigurationArrayOutput {
	return o.ApplyT(func(v *ProviderAzureCredential) ProviderAzureCredentialAutoDiscoveryConfigurationArrayOutput {
		return v.AutoDiscoveryConfigurations
	}).(ProviderAzureCredentialAutoDiscoveryConfigurationArrayOutput)
}

// The client ID of the Azure Credential.
func (o ProviderAzureCredentialOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderAzureCredential) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// The client secret of the Azure Credential.
func (o ProviderAzureCredentialOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderAzureCredential) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// The name of the Azure Credential.
func (o ProviderAzureCredentialOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderAzureCredential) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of tag filters to apply to resources.
func (o ProviderAzureCredentialOutput) ResourceDiscoveryTagFilters() ProviderAzureCredentialResourceDiscoveryTagFilterArrayOutput {
	return o.ApplyT(func(v *ProviderAzureCredential) ProviderAzureCredentialResourceDiscoveryTagFilterArrayOutput {
		return v.ResourceDiscoveryTagFilters
	}).(ProviderAzureCredentialResourceDiscoveryTagFilterArrayOutput)
}

// The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
func (o ProviderAzureCredentialOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderAzureCredential) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// The list of resource tags to add to metrics.
func (o ProviderAzureCredentialOutput) ResourceTagsToAddToMetrics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderAzureCredential) pulumi.StringArrayOutput { return v.ResourceTagsToAddToMetrics }).(pulumi.StringArrayOutput)
}

func (o ProviderAzureCredentialOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderAzureCredential) pulumi.StringOutput { return v.StackId }).(pulumi.StringOutput)
}

// The tenant ID of the Azure Credential.
func (o ProviderAzureCredentialOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderAzureCredential) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

type ProviderAzureCredentialArrayOutput struct{ *pulumi.OutputState }

func (ProviderAzureCredentialArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderAzureCredential)(nil)).Elem()
}

func (o ProviderAzureCredentialArrayOutput) ToProviderAzureCredentialArrayOutput() ProviderAzureCredentialArrayOutput {
	return o
}

func (o ProviderAzureCredentialArrayOutput) ToProviderAzureCredentialArrayOutputWithContext(ctx context.Context) ProviderAzureCredentialArrayOutput {
	return o
}

func (o ProviderAzureCredentialArrayOutput) Index(i pulumi.IntInput) ProviderAzureCredentialOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProviderAzureCredential {
		return vs[0].([]*ProviderAzureCredential)[vs[1].(int)]
	}).(ProviderAzureCredentialOutput)
}

type ProviderAzureCredentialMapOutput struct{ *pulumi.OutputState }

func (ProviderAzureCredentialMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderAzureCredential)(nil)).Elem()
}

func (o ProviderAzureCredentialMapOutput) ToProviderAzureCredentialMapOutput() ProviderAzureCredentialMapOutput {
	return o
}

func (o ProviderAzureCredentialMapOutput) ToProviderAzureCredentialMapOutputWithContext(ctx context.Context) ProviderAzureCredentialMapOutput {
	return o
}

func (o ProviderAzureCredentialMapOutput) MapIndex(k pulumi.StringInput) ProviderAzureCredentialOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProviderAzureCredential {
		return vs[0].(map[string]*ProviderAzureCredential)[vs[1].(string)]
	}).(ProviderAzureCredentialOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderAzureCredentialInput)(nil)).Elem(), &ProviderAzureCredential{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderAzureCredentialArrayInput)(nil)).Elem(), ProviderAzureCredentialArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderAzureCredentialMapInput)(nil)).Elem(), ProviderAzureCredentialMap{})
	pulumi.RegisterOutputType(ProviderAzureCredentialOutput{})
	pulumi.RegisterOutputType(ProviderAzureCredentialArrayOutput{})
	pulumi.RegisterOutputType(ProviderAzureCredentialMapOutput{})
}
