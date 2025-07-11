// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.CloudProvider.Inputs
{

    public sealed class GetAzureCredentialAutoDiscoveryConfigurationArgs : global::Pulumi.InvokeArgs
    {
        [Input("resourceTypeConfigurations", required: true)]
        private List<Inputs.GetAzureCredentialAutoDiscoveryConfigurationResourceTypeConfigurationArgs>? _resourceTypeConfigurations;

        /// <summary>
        /// The list of resource type configurations.
        /// </summary>
        public List<Inputs.GetAzureCredentialAutoDiscoveryConfigurationResourceTypeConfigurationArgs> ResourceTypeConfigurations
        {
            get => _resourceTypeConfigurations ?? (_resourceTypeConfigurations = new List<Inputs.GetAzureCredentialAutoDiscoveryConfigurationResourceTypeConfigurationArgs>());
            set => _resourceTypeConfigurations = value;
        }

        /// <summary>
        /// The subscription ID of the Azure account.
        /// </summary>
        [Input("subscriptionId", required: true)]
        public string SubscriptionId { get; set; } = null!;

        public GetAzureCredentialAutoDiscoveryConfigurationArgs()
        {
        }
        public static new GetAzureCredentialAutoDiscoveryConfigurationArgs Empty => new GetAzureCredentialAutoDiscoveryConfigurationArgs();
    }
}
