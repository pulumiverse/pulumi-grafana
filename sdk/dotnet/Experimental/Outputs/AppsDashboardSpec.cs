// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Experimental.Outputs
{

    [OutputType]
    public sealed class AppsDashboardSpec
    {
        /// <summary>
        /// The JSON representation of the dashboard spec.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// The tags of the dashboard. If not set, the tags will be derived from the JSON spec.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The title of the dashboard. If not set, the title will be derived from the JSON spec.
        /// </summary>
        public readonly string? Title;

        [OutputConstructor]
        private AppsDashboardSpec(
            string json,

            ImmutableArray<string> tags,

            string? title)
        {
            Json = json;
            Tags = tags;
            Title = title;
        }
    }
}
