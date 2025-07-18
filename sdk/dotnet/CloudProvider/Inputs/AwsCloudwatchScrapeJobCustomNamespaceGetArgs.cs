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

    public sealed class AwsCloudwatchScrapeJobCustomNamespaceGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("metrics")]
        private InputList<Inputs.AwsCloudwatchScrapeJobCustomNamespaceMetricGetArgs>? _metrics;

        /// <summary>
        /// One or more configuration blocks to configure metrics and their statistics to scrape. Each block must represent a distinct metric name. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public InputList<Inputs.AwsCloudwatchScrapeJobCustomNamespaceMetricGetArgs> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<Inputs.AwsCloudwatchScrapeJobCustomNamespaceMetricGetArgs>());
            set => _metrics = value;
        }

        /// <summary>
        /// The name of the custom namespace to scrape.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The interval in seconds to scrape the custom namespace. Defaults to `300`.
        /// </summary>
        [Input("scrapeIntervalSeconds")]
        public Input<int>? ScrapeIntervalSeconds { get; set; }

        public AwsCloudwatchScrapeJobCustomNamespaceGetArgs()
        {
        }
        public static new AwsCloudwatchScrapeJobCustomNamespaceGetArgs Empty => new AwsCloudwatchScrapeJobCustomNamespaceGetArgs();
    }
}
