// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprovider

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// ## Example Usage
func LookupAzureCredential(ctx *pulumi.Context, args *LookupAzureCredentialArgs, opts ...pulumi.InvokeOption) (*LookupAzureCredentialResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAzureCredentialResult
	err := ctx.Invoke("grafana:cloudProvider/getAzureCredential:getAzureCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAzureCredential.
type LookupAzureCredentialArgs struct {
	// The list of auto discovery configurations.
	AutoDiscoveryConfigurations []GetAzureCredentialAutoDiscoveryConfiguration `pulumi:"autoDiscoveryConfigurations"`
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilters []GetAzureCredentialResourceDiscoveryTagFilter `pulumi:"resourceDiscoveryTagFilters"`
	// The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
	ResourceId string `pulumi:"resourceId"`
	StackId    string `pulumi:"stackId"`
}

// A collection of values returned by getAzureCredential.
type LookupAzureCredentialResult struct {
	// The list of auto discovery configurations.
	AutoDiscoveryConfigurations []GetAzureCredentialAutoDiscoveryConfiguration `pulumi:"autoDiscoveryConfigurations"`
	// The client ID of the Azure Credential.
	ClientId string `pulumi:"clientId"`
	// The client secret of the Azure Credential.
	ClientSecret string `pulumi:"clientSecret"`
	Id           string `pulumi:"id"`
	// The name of the Azure Credential.
	Name string `pulumi:"name"`
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilters []GetAzureCredentialResourceDiscoveryTagFilter `pulumi:"resourceDiscoveryTagFilters"`
	// The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
	ResourceId string `pulumi:"resourceId"`
	// The list of resource tags to add to metrics.
	ResourceTagsToAddToMetrics []string `pulumi:"resourceTagsToAddToMetrics"`
	StackId                    string   `pulumi:"stackId"`
	// The tenant ID of the Azure Credential.
	TenantId string `pulumi:"tenantId"`
}

func LookupAzureCredentialOutput(ctx *pulumi.Context, args LookupAzureCredentialOutputArgs, opts ...pulumi.InvokeOption) LookupAzureCredentialResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAzureCredentialResultOutput, error) {
			args := v.(LookupAzureCredentialArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("grafana:cloudProvider/getAzureCredential:getAzureCredential", args, LookupAzureCredentialResultOutput{}, options).(LookupAzureCredentialResultOutput), nil
		}).(LookupAzureCredentialResultOutput)
}

// A collection of arguments for invoking getAzureCredential.
type LookupAzureCredentialOutputArgs struct {
	// The list of auto discovery configurations.
	AutoDiscoveryConfigurations GetAzureCredentialAutoDiscoveryConfigurationArrayInput `pulumi:"autoDiscoveryConfigurations"`
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilters GetAzureCredentialResourceDiscoveryTagFilterArrayInput `pulumi:"resourceDiscoveryTagFilters"`
	// The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	StackId    pulumi.StringInput `pulumi:"stackId"`
}

func (LookupAzureCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureCredentialArgs)(nil)).Elem()
}

// A collection of values returned by getAzureCredential.
type LookupAzureCredentialResultOutput struct{ *pulumi.OutputState }

func (LookupAzureCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureCredentialResult)(nil)).Elem()
}

func (o LookupAzureCredentialResultOutput) ToLookupAzureCredentialResultOutput() LookupAzureCredentialResultOutput {
	return o
}

func (o LookupAzureCredentialResultOutput) ToLookupAzureCredentialResultOutputWithContext(ctx context.Context) LookupAzureCredentialResultOutput {
	return o
}

// The list of auto discovery configurations.
func (o LookupAzureCredentialResultOutput) AutoDiscoveryConfigurations() GetAzureCredentialAutoDiscoveryConfigurationArrayOutput {
	return o.ApplyT(func(v LookupAzureCredentialResult) []GetAzureCredentialAutoDiscoveryConfiguration {
		return v.AutoDiscoveryConfigurations
	}).(GetAzureCredentialAutoDiscoveryConfigurationArrayOutput)
}

// The client ID of the Azure Credential.
func (o LookupAzureCredentialResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCredentialResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// The client secret of the Azure Credential.
func (o LookupAzureCredentialResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCredentialResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

func (o LookupAzureCredentialResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCredentialResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the Azure Credential.
func (o LookupAzureCredentialResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCredentialResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of tag filters to apply to resources.
func (o LookupAzureCredentialResultOutput) ResourceDiscoveryTagFilters() GetAzureCredentialResourceDiscoveryTagFilterArrayOutput {
	return o.ApplyT(func(v LookupAzureCredentialResult) []GetAzureCredentialResourceDiscoveryTagFilter {
		return v.ResourceDiscoveryTagFilters
	}).(GetAzureCredentialResourceDiscoveryTagFilterArrayOutput)
}

// The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
func (o LookupAzureCredentialResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCredentialResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

// The list of resource tags to add to metrics.
func (o LookupAzureCredentialResultOutput) ResourceTagsToAddToMetrics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAzureCredentialResult) []string { return v.ResourceTagsToAddToMetrics }).(pulumi.StringArrayOutput)
}

func (o LookupAzureCredentialResultOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCredentialResult) string { return v.StackId }).(pulumi.StringOutput)
}

// The tenant ID of the Azure Credential.
func (o LookupAzureCredentialResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCredentialResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAzureCredentialResultOutput{})
}
