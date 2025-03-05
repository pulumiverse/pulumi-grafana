// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Cloud
{
    [Obsolete(@"grafana.cloud/getproviderazurecredential.getProviderAzureCredential has been deprecated in favor of grafana.cloudprovider/getazurecredential.getAzureCredential")]
    public static class GetProviderAzureCredential
    {
        /// <summary>
        /// ## Example Usage
        /// </summary>
        public static Task<GetProviderAzureCredentialResult> InvokeAsync(GetProviderAzureCredentialArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProviderAzureCredentialResult>("grafana:cloud/getProviderAzureCredential:getProviderAzureCredential", args ?? new GetProviderAzureCredentialArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// </summary>
        public static Output<GetProviderAzureCredentialResult> Invoke(GetProviderAzureCredentialInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProviderAzureCredentialResult>("grafana:cloud/getProviderAzureCredential:getProviderAzureCredential", args ?? new GetProviderAzureCredentialInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// </summary>
        public static Output<GetProviderAzureCredentialResult> Invoke(GetProviderAzureCredentialInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetProviderAzureCredentialResult>("grafana:cloud/getProviderAzureCredential:getProviderAzureCredential", args ?? new GetProviderAzureCredentialInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProviderAzureCredentialArgs : global::Pulumi.InvokeArgs
    {
        [Input("autoDiscoveryConfigurations")]
        private List<Inputs.GetProviderAzureCredentialAutoDiscoveryConfigurationArgs>? _autoDiscoveryConfigurations;

        /// <summary>
        /// The list of auto discovery configurations.
        /// </summary>
        public List<Inputs.GetProviderAzureCredentialAutoDiscoveryConfigurationArgs> AutoDiscoveryConfigurations
        {
            get => _autoDiscoveryConfigurations ?? (_autoDiscoveryConfigurations = new List<Inputs.GetProviderAzureCredentialAutoDiscoveryConfigurationArgs>());
            set => _autoDiscoveryConfigurations = value;
        }

        [Input("resourceDiscoveryTagFilters")]
        private List<Inputs.GetProviderAzureCredentialResourceDiscoveryTagFilterArgs>? _resourceDiscoveryTagFilters;

        /// <summary>
        /// The list of tag filters to apply to resources.
        /// </summary>
        public List<Inputs.GetProviderAzureCredentialResourceDiscoveryTagFilterArgs> ResourceDiscoveryTagFilters
        {
            get => _resourceDiscoveryTagFilters ?? (_resourceDiscoveryTagFilters = new List<Inputs.GetProviderAzureCredentialResourceDiscoveryTagFilterArgs>());
            set => _resourceDiscoveryTagFilters = value;
        }

        /// <summary>
        /// The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
        /// </summary>
        [Input("resourceId", required: true)]
        public string ResourceId { get; set; } = null!;

        [Input("stackId", required: true)]
        public string StackId { get; set; } = null!;

        public GetProviderAzureCredentialArgs()
        {
        }
        public static new GetProviderAzureCredentialArgs Empty => new GetProviderAzureCredentialArgs();
    }

    public sealed class GetProviderAzureCredentialInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("autoDiscoveryConfigurations")]
        private InputList<Inputs.GetProviderAzureCredentialAutoDiscoveryConfigurationInputArgs>? _autoDiscoveryConfigurations;

        /// <summary>
        /// The list of auto discovery configurations.
        /// </summary>
        public InputList<Inputs.GetProviderAzureCredentialAutoDiscoveryConfigurationInputArgs> AutoDiscoveryConfigurations
        {
            get => _autoDiscoveryConfigurations ?? (_autoDiscoveryConfigurations = new InputList<Inputs.GetProviderAzureCredentialAutoDiscoveryConfigurationInputArgs>());
            set => _autoDiscoveryConfigurations = value;
        }

        [Input("resourceDiscoveryTagFilters")]
        private InputList<Inputs.GetProviderAzureCredentialResourceDiscoveryTagFilterInputArgs>? _resourceDiscoveryTagFilters;

        /// <summary>
        /// The list of tag filters to apply to resources.
        /// </summary>
        public InputList<Inputs.GetProviderAzureCredentialResourceDiscoveryTagFilterInputArgs> ResourceDiscoveryTagFilters
        {
            get => _resourceDiscoveryTagFilters ?? (_resourceDiscoveryTagFilters = new InputList<Inputs.GetProviderAzureCredentialResourceDiscoveryTagFilterInputArgs>());
            set => _resourceDiscoveryTagFilters = value;
        }

        /// <summary>
        /// The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        [Input("stackId", required: true)]
        public Input<string> StackId { get; set; } = null!;

        public GetProviderAzureCredentialInvokeArgs()
        {
        }
        public static new GetProviderAzureCredentialInvokeArgs Empty => new GetProviderAzureCredentialInvokeArgs();
    }


    [OutputType]
    public sealed class GetProviderAzureCredentialResult
    {
        /// <summary>
        /// The list of auto discovery configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProviderAzureCredentialAutoDiscoveryConfigurationResult> AutoDiscoveryConfigurations;
        /// <summary>
        /// The client ID of the Azure Credential.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The client secret of the Azure Credential.
        /// </summary>
        public readonly string ClientSecret;
        public readonly string Id;
        /// <summary>
        /// The name of the Azure Credential.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The list of tag filters to apply to resources.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProviderAzureCredentialResourceDiscoveryTagFilterResult> ResourceDiscoveryTagFilters;
        /// <summary>
        /// The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
        /// </summary>
        public readonly string ResourceId;
        /// <summary>
        /// The list of resource tags to add to metrics.
        /// </summary>
        public readonly ImmutableArray<string> ResourceTagsToAddToMetrics;
        public readonly string StackId;
        /// <summary>
        /// The tenant ID of the Azure Credential.
        /// </summary>
        public readonly string TenantId;

        [OutputConstructor]
        private GetProviderAzureCredentialResult(
            ImmutableArray<Outputs.GetProviderAzureCredentialAutoDiscoveryConfigurationResult> autoDiscoveryConfigurations,

            string clientId,

            string clientSecret,

            string id,

            string name,

            ImmutableArray<Outputs.GetProviderAzureCredentialResourceDiscoveryTagFilterResult> resourceDiscoveryTagFilters,

            string resourceId,

            ImmutableArray<string> resourceTagsToAddToMetrics,

            string stackId,

            string tenantId)
        {
            AutoDiscoveryConfigurations = autoDiscoveryConfigurations;
            ClientId = clientId;
            ClientSecret = clientSecret;
            Id = id;
            Name = name;
            ResourceDiscoveryTagFilters = resourceDiscoveryTagFilters;
            ResourceId = resourceId;
            ResourceTagsToAddToMetrics = resourceTagsToAddToMetrics;
            StackId = stackId;
            TenantId = tenantId;
        }
    }
}
