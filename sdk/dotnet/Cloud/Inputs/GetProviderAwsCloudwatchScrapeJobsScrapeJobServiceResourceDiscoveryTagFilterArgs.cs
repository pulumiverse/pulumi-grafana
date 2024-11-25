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

    public sealed class GetProviderAwsCloudwatchScrapeJobsScrapeJobServiceResourceDiscoveryTagFilterInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key of the tag filter.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The value of the tag filter.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public GetProviderAwsCloudwatchScrapeJobsScrapeJobServiceResourceDiscoveryTagFilterInputArgs()
        {
        }
        public static new GetProviderAwsCloudwatchScrapeJobsScrapeJobServiceResourceDiscoveryTagFilterInputArgs Empty => new GetProviderAwsCloudwatchScrapeJobsScrapeJobServiceResourceDiscoveryTagFilterInputArgs();
    }
}
