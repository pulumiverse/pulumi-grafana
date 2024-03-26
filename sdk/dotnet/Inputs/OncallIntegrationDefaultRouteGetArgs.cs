// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Inputs
{

    public sealed class OncallIntegrationDefaultRouteGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the escalation chain.
        /// </summary>
        [Input("escalationChainId")]
        public Input<string>? EscalationChainId { get; set; }

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// MS teams-specific settings for a route.
        /// </summary>
        [Input("msteams")]
        public Input<Inputs.OncallIntegrationDefaultRouteMsteamsGetArgs>? Msteams { get; set; }

        /// <summary>
        /// Slack-specific settings for a route.
        /// </summary>
        [Input("slack")]
        public Input<Inputs.OncallIntegrationDefaultRouteSlackGetArgs>? Slack { get; set; }

        /// <summary>
        /// Telegram-specific settings for a route.
        /// </summary>
        [Input("telegram")]
        public Input<Inputs.OncallIntegrationDefaultRouteTelegramGetArgs>? Telegram { get; set; }

        public OncallIntegrationDefaultRouteGetArgs()
        {
        }
        public static new OncallIntegrationDefaultRouteGetArgs Empty => new OncallIntegrationDefaultRouteGetArgs();
    }
}