// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.CloudProvider.Inputs
{

    public sealed class GetAwsCloudwatchScrapeJobsScrapeJobServiceInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("metrics")]
        private InputList<Inputs.GetAwsCloudwatchScrapeJobsScrapeJobServiceMetricInputArgs>? _metrics;

        /// <summary>
        /// One or more configuration blocks to configure metrics and their statistics to scrape. Each block must represent a distinct metric name. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public InputList<Inputs.GetAwsCloudwatchScrapeJobsScrapeJobServiceMetricInputArgs> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<Inputs.GetAwsCloudwatchScrapeJobsScrapeJobServiceMetricInputArgs>());
            set => _metrics = value;
        }

        /// <summary>
        /// The name of the service to scrape. See https://grafana.com/docs/grafana-cloud/monitor-infrastructure/aws/cloudwatch-metrics/services/ for supported services, metrics, and their statistics.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("resourceDiscoveryTagFilters")]
        private InputList<Inputs.GetAwsCloudwatchScrapeJobsScrapeJobServiceResourceDiscoveryTagFilterInputArgs>? _resourceDiscoveryTagFilters;

        /// <summary>
        /// One or more configuration blocks to configure tag filters applied to discovery of resource entities in the associated AWS account. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public InputList<Inputs.GetAwsCloudwatchScrapeJobsScrapeJobServiceResourceDiscoveryTagFilterInputArgs> ResourceDiscoveryTagFilters
        {
            get => _resourceDiscoveryTagFilters ?? (_resourceDiscoveryTagFilters = new InputList<Inputs.GetAwsCloudwatchScrapeJobsScrapeJobServiceResourceDiscoveryTagFilterInputArgs>());
            set => _resourceDiscoveryTagFilters = value;
        }

        /// <summary>
        /// The interval in seconds to scrape the service. See https://grafana.com/docs/grafana-cloud/monitor-infrastructure/aws/cloudwatch-metrics/services/ for supported scrape intervals.
        /// </summary>
        [Input("scrapeIntervalSeconds", required: true)]
        public Input<int> ScrapeIntervalSeconds { get; set; } = null!;

        [Input("tagsToAddToMetrics", required: true)]
        private InputList<string>? _tagsToAddToMetrics;

        /// <summary>
        /// A set of tags to add to all metrics exported by this scrape job, for use in PromQL queries.
        /// </summary>
        public InputList<string> TagsToAddToMetrics
        {
            get => _tagsToAddToMetrics ?? (_tagsToAddToMetrics = new InputList<string>());
            set => _tagsToAddToMetrics = value;
        }

        public GetAwsCloudwatchScrapeJobsScrapeJobServiceInputArgs()
        {
        }
        public static new GetAwsCloudwatchScrapeJobsScrapeJobServiceInputArgs Empty => new GetAwsCloudwatchScrapeJobsScrapeJobServiceInputArgs();
    }
}
