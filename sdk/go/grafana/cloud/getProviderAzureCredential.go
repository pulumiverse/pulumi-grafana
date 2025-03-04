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
// Deprecated: grafana.cloud/getproviderazurecredential.getProviderAzureCredential has been deprecated in favor of grafana.cloudprovider/getazurecredential.getAzureCredential
func LookupProviderAzureCredential(ctx *pulumi.Context, args *LookupProviderAzureCredentialArgs, opts ...pulumi.InvokeOption) (*LookupProviderAzureCredentialResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProviderAzureCredentialResult
	err := ctx.Invoke("grafana:cloud/getProviderAzureCredential:getProviderAzureCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProviderAzureCredential.
type LookupProviderAzureCredentialArgs struct {
	// The list of auto discovery configurations.
	AutoDiscoveryConfigurations []GetProviderAzureCredentialAutoDiscoveryConfiguration `pulumi:"autoDiscoveryConfigurations"`
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilters []GetProviderAzureCredentialResourceDiscoveryTagFilter `pulumi:"resourceDiscoveryTagFilters"`
	// The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
	ResourceId string `pulumi:"resourceId"`
	StackId    string `pulumi:"stackId"`
}

// A collection of values returned by getProviderAzureCredential.
type LookupProviderAzureCredentialResult struct {
	// The list of auto discovery configurations.
	AutoDiscoveryConfigurations []GetProviderAzureCredentialAutoDiscoveryConfiguration `pulumi:"autoDiscoveryConfigurations"`
	// The client ID of the Azure Credential.
	ClientId string `pulumi:"clientId"`
	// The client secret of the Azure Credential.
	ClientSecret string `pulumi:"clientSecret"`
	Id           string `pulumi:"id"`
	// The name of the Azure Credential.
	Name string `pulumi:"name"`
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilters []GetProviderAzureCredentialResourceDiscoveryTagFilter `pulumi:"resourceDiscoveryTagFilters"`
	// The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
	ResourceId string `pulumi:"resourceId"`
	// A set of regions that this AWS Account resource applies to.
	ResourceTagsToAddToMetrics []string `pulumi:"resourceTagsToAddToMetrics"`
	StackId                    string   `pulumi:"stackId"`
	// The tenant ID of the Azure Credential.
	TenantId string `pulumi:"tenantId"`
}

func LookupProviderAzureCredentialOutput(ctx *pulumi.Context, args LookupProviderAzureCredentialOutputArgs, opts ...pulumi.InvokeOption) LookupProviderAzureCredentialResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProviderAzureCredentialResultOutput, error) {
			args := v.(LookupProviderAzureCredentialArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("grafana:cloud/getProviderAzureCredential:getProviderAzureCredential", args, LookupProviderAzureCredentialResultOutput{}, options).(LookupProviderAzureCredentialResultOutput), nil
		}).(LookupProviderAzureCredentialResultOutput)
}

// A collection of arguments for invoking getProviderAzureCredential.
type LookupProviderAzureCredentialOutputArgs struct {
	// The list of auto discovery configurations.
	AutoDiscoveryConfigurations GetProviderAzureCredentialAutoDiscoveryConfigurationArrayInput `pulumi:"autoDiscoveryConfigurations"`
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilters GetProviderAzureCredentialResourceDiscoveryTagFilterArrayInput `pulumi:"resourceDiscoveryTagFilters"`
	// The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	StackId    pulumi.StringInput `pulumi:"stackId"`
}

func (LookupProviderAzureCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProviderAzureCredentialArgs)(nil)).Elem()
}

// A collection of values returned by getProviderAzureCredential.
type LookupProviderAzureCredentialResultOutput struct{ *pulumi.OutputState }

func (LookupProviderAzureCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProviderAzureCredentialResult)(nil)).Elem()
}

func (o LookupProviderAzureCredentialResultOutput) ToLookupProviderAzureCredentialResultOutput() LookupProviderAzureCredentialResultOutput {
	return o
}

func (o LookupProviderAzureCredentialResultOutput) ToLookupProviderAzureCredentialResultOutputWithContext(ctx context.Context) LookupProviderAzureCredentialResultOutput {
	return o
}

// The list of auto discovery configurations.
func (o LookupProviderAzureCredentialResultOutput) AutoDiscoveryConfigurations() GetProviderAzureCredentialAutoDiscoveryConfigurationArrayOutput {
	return o.ApplyT(func(v LookupProviderAzureCredentialResult) []GetProviderAzureCredentialAutoDiscoveryConfiguration {
		return v.AutoDiscoveryConfigurations
	}).(GetProviderAzureCredentialAutoDiscoveryConfigurationArrayOutput)
}

// The client ID of the Azure Credential.
func (o LookupProviderAzureCredentialResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderAzureCredentialResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// The client secret of the Azure Credential.
func (o LookupProviderAzureCredentialResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderAzureCredentialResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

func (o LookupProviderAzureCredentialResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderAzureCredentialResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the Azure Credential.
func (o LookupProviderAzureCredentialResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderAzureCredentialResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of tag filters to apply to resources.
func (o LookupProviderAzureCredentialResultOutput) ResourceDiscoveryTagFilters() GetProviderAzureCredentialResourceDiscoveryTagFilterArrayOutput {
	return o.ApplyT(func(v LookupProviderAzureCredentialResult) []GetProviderAzureCredentialResourceDiscoveryTagFilter {
		return v.ResourceDiscoveryTagFilters
	}).(GetProviderAzureCredentialResourceDiscoveryTagFilterArrayOutput)
}

// The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
func (o LookupProviderAzureCredentialResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderAzureCredentialResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

// A set of regions that this AWS Account resource applies to.
func (o LookupProviderAzureCredentialResultOutput) ResourceTagsToAddToMetrics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProviderAzureCredentialResult) []string { return v.ResourceTagsToAddToMetrics }).(pulumi.StringArrayOutput)
}

func (o LookupProviderAzureCredentialResultOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderAzureCredentialResult) string { return v.StackId }).(pulumi.StringOutput)
}

// The tenant ID of the Azure Credential.
func (o LookupProviderAzureCredentialResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderAzureCredentialResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProviderAzureCredentialResultOutput{})
}
