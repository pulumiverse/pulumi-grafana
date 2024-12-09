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

    public sealed class GetProviderAwsCloudwatchScrapeJobCustomNamespaceInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("metrics")]
        private InputList<Inputs.GetProviderAwsCloudwatchScrapeJobCustomNamespaceMetricInputArgs>? _metrics;

        /// <summary>
        /// One or more configuration blocks to configure metrics and their statistics to scrape. Each block must represent a distinct metric name. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public InputList<Inputs.GetProviderAwsCloudwatchScrapeJobCustomNamespaceMetricInputArgs> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<Inputs.GetProviderAwsCloudwatchScrapeJobCustomNamespaceMetricInputArgs>());
            set => _metrics = value;
        }

        /// <summary>
        /// The name of the custom namespace to scrape.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The interval in seconds to scrape the custom namespace.
        /// </summary>
        [Input("scrapeIntervalSeconds", required: true)]
        public Input<int> ScrapeIntervalSeconds { get; set; } = null!;

        public GetProviderAwsCloudwatchScrapeJobCustomNamespaceInputArgs()
        {
        }
        public static new GetProviderAwsCloudwatchScrapeJobCustomNamespaceInputArgs Empty => new GetProviderAwsCloudwatchScrapeJobCustomNamespaceInputArgs();
    }
}