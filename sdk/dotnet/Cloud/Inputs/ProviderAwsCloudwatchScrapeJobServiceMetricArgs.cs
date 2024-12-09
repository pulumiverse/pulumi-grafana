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

    public sealed class ProviderAwsCloudwatchScrapeJobServiceMetricArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the metric to scrape.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("statistics", required: true)]
        private InputList<string>? _statistics;

        /// <summary>
        /// A set of statistics to scrape.
        /// </summary>
        public InputList<string> Statistics
        {
            get => _statistics ?? (_statistics = new InputList<string>());
            set => _statistics = value;
        }

        public ProviderAwsCloudwatchScrapeJobServiceMetricArgs()
        {
        }
        public static new ProviderAwsCloudwatchScrapeJobServiceMetricArgs Empty => new ProviderAwsCloudwatchScrapeJobServiceMetricArgs();
    }
}