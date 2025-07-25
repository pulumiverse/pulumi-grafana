// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Inputs
{

    public sealed class OncallIntegrationDefaultRouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the escalation chain.
        /// </summary>
        [Input("escalationChainId")]
        public Input<string>? EscalationChainId { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// MS teams-specific settings for a route.
        /// </summary>
        [Input("msteams")]
        public Input<Inputs.OncallIntegrationDefaultRouteMsteamsArgs>? Msteams { get; set; }

        /// <summary>
        /// Slack-specific settings for a route.
        /// </summary>
        [Input("slack")]
        public Input<Inputs.OncallIntegrationDefaultRouteSlackArgs>? Slack { get; set; }

        /// <summary>
        /// Telegram-specific settings for a route.
        /// </summary>
        [Input("telegram")]
        public Input<Inputs.OncallIntegrationDefaultRouteTelegramArgs>? Telegram { get; set; }

        public OncallIntegrationDefaultRouteArgs()
        {
        }
        public static new OncallIntegrationDefaultRouteArgs Empty => new OncallIntegrationDefaultRouteArgs();
    }
}
