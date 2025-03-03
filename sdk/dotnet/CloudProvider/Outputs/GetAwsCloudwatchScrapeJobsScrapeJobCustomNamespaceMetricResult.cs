// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.CloudProvider.Outputs
{

    [OutputType]
    public sealed class GetAwsCloudwatchScrapeJobsScrapeJobCustomNamespaceMetricResult
    {
        /// <summary>
        /// The name of the metric to scrape.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A set of statistics to scrape.
        /// </summary>
        public readonly ImmutableArray<string> Statistics;

        [OutputConstructor]
        private GetAwsCloudwatchScrapeJobsScrapeJobCustomNamespaceMetricResult(
            string name,

            ImmutableArray<string> statistics)
        {
            Name = name;
            Statistics = statistics;
        }
    }
}
