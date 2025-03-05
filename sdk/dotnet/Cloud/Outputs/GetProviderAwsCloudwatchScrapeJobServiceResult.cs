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
    public sealed class GetProviderAwsCloudwatchScrapeJobServiceResult
    {
        /// <summary>
        /// One or more configuration blocks to configure metrics and their statistics to scrape. Each block must represent a distinct metric name. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProviderAwsCloudwatchScrapeJobServiceMetricResult> Metrics;
        /// <summary>
        /// The name of the service to scrape. See https://grafana.com/docs/grafana-cloud/monitor-infrastructure/monitor-cloud-provider/aws/cloudwatch-metrics/services/ for supported services, metrics, and their statistics.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// One or more configuration blocks to configure tag filters applied to discovery of resource entities in the associated AWS account. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProviderAwsCloudwatchScrapeJobServiceResourceDiscoveryTagFilterResult> ResourceDiscoveryTagFilters;
        /// <summary>
        /// The interval in seconds to scrape the service. See https://grafana.com/docs/grafana-cloud/monitor-infrastructure/monitor-cloud-provider/aws/cloudwatch-metrics/services/ for supported scrape intervals.
        /// </summary>
        public readonly int ScrapeIntervalSeconds;
        /// <summary>
        /// A set of tags to add to all metrics exported by this scrape job, for use in PromQL queries.
        /// </summary>
        public readonly ImmutableArray<string> TagsToAddToMetrics;

        [OutputConstructor]
        private GetProviderAwsCloudwatchScrapeJobServiceResult(
            ImmutableArray<Outputs.GetProviderAwsCloudwatchScrapeJobServiceMetricResult> metrics,

            string name,

            ImmutableArray<Outputs.GetProviderAwsCloudwatchScrapeJobServiceResourceDiscoveryTagFilterResult> resourceDiscoveryTagFilters,

            int scrapeIntervalSeconds,

            ImmutableArray<string> tagsToAddToMetrics)
        {
            Metrics = metrics;
            Name = name;
            ResourceDiscoveryTagFilters = resourceDiscoveryTagFilters;
            ScrapeIntervalSeconds = scrapeIntervalSeconds;
            TagsToAddToMetrics = tagsToAddToMetrics;
        }
    }
}
