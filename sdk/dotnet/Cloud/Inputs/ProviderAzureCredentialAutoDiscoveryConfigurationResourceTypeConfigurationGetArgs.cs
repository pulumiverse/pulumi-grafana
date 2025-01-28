// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Cloud.Inputs
{

    public sealed class ProviderAzureCredentialAutoDiscoveryConfigurationResourceTypeConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("metricConfigurations", required: true)]
        private InputList<Inputs.ProviderAzureCredentialAutoDiscoveryConfigurationResourceTypeConfigurationMetricConfigurationGetArgs>? _metricConfigurations;
        public InputList<Inputs.ProviderAzureCredentialAutoDiscoveryConfigurationResourceTypeConfigurationMetricConfigurationGetArgs> MetricConfigurations
        {
            get => _metricConfigurations ?? (_metricConfigurations = new InputList<Inputs.ProviderAzureCredentialAutoDiscoveryConfigurationResourceTypeConfigurationMetricConfigurationGetArgs>());
            set => _metricConfigurations = value;
        }

        [Input("resourceTypeName", required: true)]
        public Input<string> ResourceTypeName { get; set; } = null!;

        public ProviderAzureCredentialAutoDiscoveryConfigurationResourceTypeConfigurationGetArgs()
        {
        }
        public static new ProviderAzureCredentialAutoDiscoveryConfigurationResourceTypeConfigurationGetArgs Empty => new ProviderAzureCredentialAutoDiscoveryConfigurationResourceTypeConfigurationGetArgs();
    }
}
