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

    public sealed class OncallScheduleSlackGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Slack channel id. Reminder about schedule shifts will be directed to this channel in Slack.
        /// </summary>
        [Input("channelId")]
        public Input<string>? ChannelId { get; set; }

        /// <summary>
        /// Slack user group id. Members of user group will be updated when on-call users change.
        /// </summary>
        [Input("userGroupId")]
        public Input<string>? UserGroupId { get; set; }

        public OncallScheduleSlackGetArgs()
        {
        }
        public static new OncallScheduleSlackGetArgs Empty => new OncallScheduleSlackGetArgs();
    }
}