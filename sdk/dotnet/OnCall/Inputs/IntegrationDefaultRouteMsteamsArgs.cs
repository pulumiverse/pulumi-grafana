// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.OnCall.Inputs
{

    public sealed class IntegrationDefaultRouteMsteamsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable notification in MS teams. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// MS teams channel id. Alerts will be directed to this channel in Microsoft teams.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public IntegrationDefaultRouteMsteamsArgs()
        {
        }
        public static new IntegrationDefaultRouteMsteamsArgs Empty => new IntegrationDefaultRouteMsteamsArgs();
    }
}
