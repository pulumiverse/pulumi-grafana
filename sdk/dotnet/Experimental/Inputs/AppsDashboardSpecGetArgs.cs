// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Experimental.Inputs
{

    public sealed class AppsDashboardSpecGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The JSON representation of the dashboard spec.
        /// </summary>
        [Input("json", required: true)]
        public Input<string> Json { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags of the dashboard. If not set, the tags will be derived from the JSON spec.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The title of the dashboard. If not set, the title will be derived from the JSON spec.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public AppsDashboardSpecGetArgs()
        {
        }
        public static new AppsDashboardSpecGetArgs Empty => new AppsDashboardSpecGetArgs();
    }
}
