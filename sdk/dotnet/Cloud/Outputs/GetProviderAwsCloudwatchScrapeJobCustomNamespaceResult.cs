// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Cloud.Outputs
{

    [OutputType]
    public sealed class GetProviderAwsCloudwatchScrapeJobCustomNamespaceResult
    {
        /// <summary>
        /// One or more configuration blocks to configure metrics and their statistics to scrape. Each block must represent a distinct metric name. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProviderAwsCloudwatchScrapeJobCustomNamespaceMetricResult> Metrics;
        /// <summary>
        /// The name of the custom namespace to scrape.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The interval in seconds to scrape the custom namespace.
        /// </summary>
        public readonly int ScrapeIntervalSeconds;

        [OutputConstructor]
        private GetProviderAwsCloudwatchScrapeJobCustomNamespaceResult(
            ImmutableArray<Outputs.GetProviderAwsCloudwatchScrapeJobCustomNamespaceMetricResult> metrics,

            string name,

            int scrapeIntervalSeconds)
        {
            Metrics = metrics;
            Name = name;
            ScrapeIntervalSeconds = scrapeIntervalSeconds;
        }
    }
}
