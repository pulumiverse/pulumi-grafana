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
    public sealed class GetAwsCloudwatchScrapeJobServiceResourceDiscoveryTagFilterResult
    {
        /// <summary>
        /// The key of the tag filter.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The value of the tag filter.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetAwsCloudwatchScrapeJobServiceResourceDiscoveryTagFilterResult(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
