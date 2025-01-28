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

    public sealed class ProviderAzureCredentialAutoDiscoveryConfigurationResourceTypeConfigurationMetricConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("aggregations", required: true)]
        private InputList<string>? _aggregations;
        public InputList<string> Aggregations
        {
            get => _aggregations ?? (_aggregations = new InputList<string>());
            set => _aggregations = value;
        }

        [Input("dimensions", required: true)]
        private InputList<string>? _dimensions;
        public InputList<string> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<string>());
            set => _dimensions = value;
        }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ProviderAzureCredentialAutoDiscoveryConfigurationResourceTypeConfigurationMetricConfigurationGetArgs()
        {
        }
        public static new ProviderAzureCredentialAutoDiscoveryConfigurationResourceTypeConfigurationMetricConfigurationGetArgs Empty => new ProviderAzureCredentialAutoDiscoveryConfigurationResourceTypeConfigurationMetricConfigurationGetArgs();
    }
}
