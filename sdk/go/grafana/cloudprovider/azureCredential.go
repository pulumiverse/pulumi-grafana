// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprovider

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
// $ pulumi import grafana:cloudProvider/azureCredential:AzureCredential name "{{ stack_id }}:{{ resource_id }}"
// ```
type AzureCredential struct {
	pulumi.CustomResourceState

	// The list of auto discovery configurations.
	AutoDiscoveryConfigurations AzureCredentialAutoDiscoveryConfigurationArrayOutput `pulumi:"autoDiscoveryConfigurations"`
	// The client ID of the Azure Credential.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The client secret of the Azure Credential.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// The name of the Azure Credential.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilters AzureCredentialResourceDiscoveryTagFilterArrayOutput `pulumi:"resourceDiscoveryTagFilters"`
	// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// A set of regions that this AWS Account resource applies to.
	ResourceTagsToAddToMetrics pulumi.StringArrayOutput `pulumi:"resourceTagsToAddToMetrics"`
	StackId                    pulumi.StringOutput      `pulumi:"stackId"`
	// The tenant ID of the Azure Credential.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewAzureCredential registers a new resource with the given unique name, arguments, and options.
func NewAzureCredential(ctx *pulumi.Context,
	name string, args *AzureCredentialArgs, opts ...pulumi.ResourceOption) (*AzureCredential, error) {
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:cloud/providerAzureCredential:ProviderAzureCredential"),
		},
	})
	opts = append(opts, aliases)
	if args.ClientSecret != nil {
		args.ClientSecret = pulumi.ToSecret(args.ClientSecret).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clientSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AzureCredential
	err := ctx.RegisterResource("grafana:cloudProvider/azureCredential:AzureCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureCredential gets an existing AzureCredential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureCredentialState, opts ...pulumi.ResourceOption) (*AzureCredential, error) {
	var resource AzureCredential
	err := ctx.ReadResource("grafana:cloudProvider/azureCredential:AzureCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureCredential resources.
type azureCredentialState struct {
	// The list of auto discovery configurations.
	AutoDiscoveryConfigurations []AzureCredentialAutoDiscoveryConfiguration `pulumi:"autoDiscoveryConfigurations"`
	// The client ID of the Azure Credential.
	ClientId *string `pulumi:"clientId"`
	// The client secret of the Azure Credential.
	ClientSecret *string `pulumi:"clientSecret"`
	// The name of the Azure Credential.
	Name *string `pulumi:"name"`
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilters []AzureCredentialResourceDiscoveryTagFilter `pulumi:"resourceDiscoveryTagFilters"`
	// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
	ResourceId *string `pulumi:"resourceId"`
	// A set of regions that this AWS Account resource applies to.
	ResourceTagsToAddToMetrics []string `pulumi:"resourceTagsToAddToMetrics"`
	StackId                    *string  `pulumi:"stackId"`
	// The tenant ID of the Azure Credential.
	TenantId *string `pulumi:"tenantId"`
}

type AzureCredentialState struct {
	// The list of auto discovery configurations.
	AutoDiscoveryConfigurations AzureCredentialAutoDiscoveryConfigurationArrayInput
	// The client ID of the Azure Credential.
	ClientId pulumi.StringPtrInput
	// The client secret of the Azure Credential.
	ClientSecret pulumi.StringPtrInput
	// The name of the Azure Credential.
	Name pulumi.StringPtrInput
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilters AzureCredentialResourceDiscoveryTagFilterArrayInput
	// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
	ResourceId pulumi.StringPtrInput
	// A set of regions that this AWS Account resource applies to.
	ResourceTagsToAddToMetrics pulumi.StringArrayInput
	StackId                    pulumi.StringPtrInput
	// The tenant ID of the Azure Credential.
	TenantId pulumi.StringPtrInput
}

func (AzureCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureCredentialState)(nil)).Elem()
}

type azureCredentialArgs struct {
	// The list of auto discovery configurations.
	AutoDiscoveryConfigurations []AzureCredentialAutoDiscoveryConfiguration `pulumi:"autoDiscoveryConfigurations"`
	// The client ID of the Azure Credential.
	ClientId string `pulumi:"clientId"`
	// The client secret of the Azure Credential.
	ClientSecret string `pulumi:"clientSecret"`
	// The name of the Azure Credential.
	Name *string `pulumi:"name"`
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilters []AzureCredentialResourceDiscoveryTagFilter `pulumi:"resourceDiscoveryTagFilters"`
	// A set of regions that this AWS Account resource applies to.
	ResourceTagsToAddToMetrics []string `pulumi:"resourceTagsToAddToMetrics"`
	StackId                    string   `pulumi:"stackId"`
	// The tenant ID of the Azure Credential.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a AzureCredential resource.
type AzureCredentialArgs struct {
	// The list of auto discovery configurations.
	AutoDiscoveryConfigurations AzureCredentialAutoDiscoveryConfigurationArrayInput
	// The client ID of the Azure Credential.
	ClientId pulumi.StringInput
	// The client secret of the Azure Credential.
	ClientSecret pulumi.StringInput
	// The name of the Azure Credential.
	Name pulumi.StringPtrInput
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilters AzureCredentialResourceDiscoveryTagFilterArrayInput
	// A set of regions that this AWS Account resource applies to.
	ResourceTagsToAddToMetrics pulumi.StringArrayInput
	StackId                    pulumi.StringInput
	// The tenant ID of the Azure Credential.
	TenantId pulumi.StringInput
}

func (AzureCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureCredentialArgs)(nil)).Elem()
}

type AzureCredentialInput interface {
	pulumi.Input

	ToAzureCredentialOutput() AzureCredentialOutput
	ToAzureCredentialOutputWithContext(ctx context.Context) AzureCredentialOutput
}

func (*AzureCredential) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureCredential)(nil)).Elem()
}

func (i *AzureCredential) ToAzureCredentialOutput() AzureCredentialOutput {
	return i.ToAzureCredentialOutputWithContext(context.Background())
}

func (i *AzureCredential) ToAzureCredentialOutputWithContext(ctx context.Context) AzureCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureCredentialOutput)
}

// AzureCredentialArrayInput is an input type that accepts AzureCredentialArray and AzureCredentialArrayOutput values.
// You can construct a concrete instance of `AzureCredentialArrayInput` via:
//
//	AzureCredentialArray{ AzureCredentialArgs{...} }
type AzureCredentialArrayInput interface {
	pulumi.Input

	ToAzureCredentialArrayOutput() AzureCredentialArrayOutput
	ToAzureCredentialArrayOutputWithContext(context.Context) AzureCredentialArrayOutput
}

type AzureCredentialArray []AzureCredentialInput

func (AzureCredentialArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureCredential)(nil)).Elem()
}

func (i AzureCredentialArray) ToAzureCredentialArrayOutput() AzureCredentialArrayOutput {
	return i.ToAzureCredentialArrayOutputWithContext(context.Background())
}

func (i AzureCredentialArray) ToAzureCredentialArrayOutputWithContext(ctx context.Context) AzureCredentialArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureCredentialArrayOutput)
}

// AzureCredentialMapInput is an input type that accepts AzureCredentialMap and AzureCredentialMapOutput values.
// You can construct a concrete instance of `AzureCredentialMapInput` via:
//
//	AzureCredentialMap{ "key": AzureCredentialArgs{...} }
type AzureCredentialMapInput interface {
	pulumi.Input

	ToAzureCredentialMapOutput() AzureCredentialMapOutput
	ToAzureCredentialMapOutputWithContext(context.Context) AzureCredentialMapOutput
}

type AzureCredentialMap map[string]AzureCredentialInput

func (AzureCredentialMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureCredential)(nil)).Elem()
}

func (i AzureCredentialMap) ToAzureCredentialMapOutput() AzureCredentialMapOutput {
	return i.ToAzureCredentialMapOutputWithContext(context.Background())
}

func (i AzureCredentialMap) ToAzureCredentialMapOutputWithContext(ctx context.Context) AzureCredentialMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureCredentialMapOutput)
}

type AzureCredentialOutput struct{ *pulumi.OutputState }

func (AzureCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureCredential)(nil)).Elem()
}

func (o AzureCredentialOutput) ToAzureCredentialOutput() AzureCredentialOutput {
	return o
}

func (o AzureCredentialOutput) ToAzureCredentialOutputWithContext(ctx context.Context) AzureCredentialOutput {
	return o
}

// The list of auto discovery configurations.
func (o AzureCredentialOutput) AutoDiscoveryConfigurations() AzureCredentialAutoDiscoveryConfigurationArrayOutput {
	return o.ApplyT(func(v *AzureCredential) AzureCredentialAutoDiscoveryConfigurationArrayOutput {
		return v.AutoDiscoveryConfigurations
	}).(AzureCredentialAutoDiscoveryConfigurationArrayOutput)
}

// The client ID of the Azure Credential.
func (o AzureCredentialOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCredential) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// The client secret of the Azure Credential.
func (o AzureCredentialOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCredential) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// The name of the Azure Credential.
func (o AzureCredentialOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCredential) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of tag filters to apply to resources.
func (o AzureCredentialOutput) ResourceDiscoveryTagFilters() AzureCredentialResourceDiscoveryTagFilterArrayOutput {
	return o.ApplyT(func(v *AzureCredential) AzureCredentialResourceDiscoveryTagFilterArrayOutput {
		return v.ResourceDiscoveryTagFilters
	}).(AzureCredentialResourceDiscoveryTagFilterArrayOutput)
}

// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
func (o AzureCredentialOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCredential) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// A set of regions that this AWS Account resource applies to.
func (o AzureCredentialOutput) ResourceTagsToAddToMetrics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureCredential) pulumi.StringArrayOutput { return v.ResourceTagsToAddToMetrics }).(pulumi.StringArrayOutput)
}

func (o AzureCredentialOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCredential) pulumi.StringOutput { return v.StackId }).(pulumi.StringOutput)
}

// The tenant ID of the Azure Credential.
func (o AzureCredentialOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCredential) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

type AzureCredentialArrayOutput struct{ *pulumi.OutputState }

func (AzureCredentialArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureCredential)(nil)).Elem()
}

func (o AzureCredentialArrayOutput) ToAzureCredentialArrayOutput() AzureCredentialArrayOutput {
	return o
}

func (o AzureCredentialArrayOutput) ToAzureCredentialArrayOutputWithContext(ctx context.Context) AzureCredentialArrayOutput {
	return o
}

func (o AzureCredentialArrayOutput) Index(i pulumi.IntInput) AzureCredentialOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AzureCredential {
		return vs[0].([]*AzureCredential)[vs[1].(int)]
	}).(AzureCredentialOutput)
}

type AzureCredentialMapOutput struct{ *pulumi.OutputState }

func (AzureCredentialMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureCredential)(nil)).Elem()
}

func (o AzureCredentialMapOutput) ToAzureCredentialMapOutput() AzureCredentialMapOutput {
	return o
}

func (o AzureCredentialMapOutput) ToAzureCredentialMapOutputWithContext(ctx context.Context) AzureCredentialMapOutput {
	return o
}

func (o AzureCredentialMapOutput) MapIndex(k pulumi.StringInput) AzureCredentialOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AzureCredential {
		return vs[0].(map[string]*AzureCredential)[vs[1].(string)]
	}).(AzureCredentialOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AzureCredentialInput)(nil)).Elem(), &AzureCredential{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureCredentialArrayInput)(nil)).Elem(), AzureCredentialArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureCredentialMapInput)(nil)).Elem(), AzureCredentialMap{})
	pulumi.RegisterOutputType(AzureCredentialOutput{})
	pulumi.RegisterOutputType(AzureCredentialArrayOutput{})
	pulumi.RegisterOutputType(AzureCredentialMapOutput{})
}
