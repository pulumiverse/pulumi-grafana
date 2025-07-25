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

    public sealed class OncallIntegrationTemplatesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Template for sending a signal to acknowledge the Incident.
        /// </summary>
        [Input("acknowledgeSignal")]
        public Input<string>? AcknowledgeSignal { get; set; }

        /// <summary>
        /// Templates for Email.
        /// </summary>
        [Input("email")]
        public Input<Inputs.OncallIntegrationTemplatesEmailArgs>? Email { get; set; }

        /// <summary>
        /// Template for the key by which alerts are grouped.
        /// </summary>
        [Input("groupingKey")]
        public Input<string>? GroupingKey { get; set; }

        /// <summary>
        /// Templates for Microsoft Teams. **NOTE**: Microsoft Teams templates are only available on Grafana Cloud.
        /// </summary>
        [Input("microsoftTeams")]
        public Input<Inputs.OncallIntegrationTemplatesMicrosoftTeamsArgs>? MicrosoftTeams { get; set; }

        /// <summary>
        /// Templates for Mobile app push notifications.
        /// </summary>
        [Input("mobileApp")]
        public Input<Inputs.OncallIntegrationTemplatesMobileAppArgs>? MobileApp { get; set; }

        /// <summary>
        /// Templates for Phone Call.
        /// </summary>
        [Input("phoneCall")]
        public Input<Inputs.OncallIntegrationTemplatesPhoneCallArgs>? PhoneCall { get; set; }

        /// <summary>
        /// Template for sending a signal to resolve the Incident.
        /// </summary>
        [Input("resolveSignal")]
        public Input<string>? ResolveSignal { get; set; }

        /// <summary>
        /// Templates for Slack.
        /// </summary>
        [Input("slack")]
        public Input<Inputs.OncallIntegrationTemplatesSlackArgs>? Slack { get; set; }

        /// <summary>
        /// Templates for SMS.
        /// </summary>
        [Input("sms")]
        public Input<Inputs.OncallIntegrationTemplatesSmsArgs>? Sms { get; set; }

        /// <summary>
        /// Template for a source link.
        /// </summary>
        [Input("sourceLink")]
        public Input<string>? SourceLink { get; set; }

        /// <summary>
        /// Templates for Telegram.
        /// </summary>
        [Input("telegram")]
        public Input<Inputs.OncallIntegrationTemplatesTelegramArgs>? Telegram { get; set; }

        /// <summary>
        /// Templates for Web.
        /// </summary>
        [Input("web")]
        public Input<Inputs.OncallIntegrationTemplatesWebArgs>? Web { get; set; }

        public OncallIntegrationTemplatesArgs()
        {
        }
        public static new OncallIntegrationTemplatesArgs Empty => new OncallIntegrationTemplatesArgs();
    }
}
