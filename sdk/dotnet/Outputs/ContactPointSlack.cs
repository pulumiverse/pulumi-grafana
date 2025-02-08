// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Outputs
{

    [OutputType]
    public sealed class ContactPointSlack
    {
        /// <summary>
        /// Templated color of the slack message.
        /// </summary>
        public readonly string? Color;
        /// <summary>
        /// Whether to disable sending resolve messages. Defaults to `false`.
        /// </summary>
        public readonly bool? DisableResolveMessage;
        /// <summary>
        /// Use this to override the Slack API endpoint URL to send requests to.
        /// </summary>
        public readonly string? EndpointUrl;
        /// <summary>
        /// The name of a Slack workspace emoji to use as the bot icon.
        /// </summary>
        public readonly string? IconEmoji;
        /// <summary>
        /// A URL of an image to use as the bot icon.
        /// </summary>
        public readonly string? IconUrl;
        /// <summary>
        /// Describes how to ping the slack channel that messages are being sent to. Options are `here` for an @here ping, `channel` for @channel, or empty for no ping.
        /// </summary>
        public readonly string? MentionChannel;
        /// <summary>
        /// Comma-separated list of groups to mention in the message.
        /// </summary>
        public readonly string? MentionGroups;
        /// <summary>
        /// Comma-separated list of users to mention in the message.
        /// </summary>
        public readonly string? MentionUsers;
        /// <summary>
        /// Channel, private group, or IM channel (can be an encoded ID or a name) to send messages to.
        /// </summary>
        public readonly string? Recipient;
        /// <summary>
        /// Additional custom properties to attach to the notifier. Defaults to `map[]`.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Settings;
        /// <summary>
        /// Templated content of the message.
        /// </summary>
        public readonly string? Text;
        /// <summary>
        /// Templated title of the message.
        /// </summary>
        public readonly string? Title;
        /// <summary>
        /// A Slack API token,for sending messages directly without the webhook method.
        /// </summary>
        public readonly string? Token;
        /// <summary>
        /// The UID of the contact point.
        /// </summary>
        public readonly string? Uid;
        /// <summary>
        /// A Slack webhook URL,for sending messages via the webhook method.
        /// </summary>
        public readonly string? Url;
        /// <summary>
        /// Username for the bot to use.
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private ContactPointSlack(
            string? color,

            bool? disableResolveMessage,

            string? endpointUrl,

            string? iconEmoji,

            string? iconUrl,

            string? mentionChannel,

            string? mentionGroups,

            string? mentionUsers,

            string? recipient,

            ImmutableDictionary<string, string>? settings,

            string? text,

            string? title,

            string? token,

            string? uid,

            string? url,

            string? username)
        {
            Color = color;
            DisableResolveMessage = disableResolveMessage;
            EndpointUrl = endpointUrl;
            IconEmoji = iconEmoji;
            IconUrl = iconUrl;
            MentionChannel = mentionChannel;
            MentionGroups = mentionGroups;
            MentionUsers = mentionUsers;
            Recipient = recipient;
            Settings = settings;
            Text = text;
            Title = title;
            Token = token;
            Uid = uid;
            Url = url;
            Username = username;
        }
    }
}
