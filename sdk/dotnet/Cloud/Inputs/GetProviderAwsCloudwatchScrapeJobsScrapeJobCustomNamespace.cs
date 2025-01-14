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

    public sealed class GetProviderAwsCloudwatchScrapeJobsScrapeJobCustomNamespaceArgs : global::Pulumi.InvokeArgs
    {
        [Input("metrics")]
        private List<Inputs.GetProviderAwsCloudwatchScrapeJobsScrapeJobCustomNamespaceMetricArgs>? _metrics;

        /// <summary>
        /// One or more configuration blocks to configure metrics and their statistics to scrape. Each block must represent a distinct metric name. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public List<Inputs.GetProviderAwsCloudwatchScrapeJobsScrapeJobCustomNamespaceMetricArgs> Metrics
        {
            get => _metrics ?? (_metrics = new List<Inputs.GetProviderAwsCloudwatchScrapeJobsScrapeJobCustomNamespaceMetricArgs>());
            set => _metrics = value;
        }

        /// <summary>
        /// The name of the custom namespace to scrape.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The interval in seconds to scrape the custom namespace.
        /// </summary>
        [Input("scrapeIntervalSeconds", required: true)]
        public int ScrapeIntervalSeconds { get; set; }

        public GetProviderAwsCloudwatchScrapeJobsScrapeJobCustomNamespaceArgs()
        {
        }
        public static new GetProviderAwsCloudwatchScrapeJobsScrapeJobCustomNamespaceArgs Empty => new GetProviderAwsCloudwatchScrapeJobsScrapeJobCustomNamespaceArgs();
    }
}