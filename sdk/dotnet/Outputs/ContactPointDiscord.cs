// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    public sealed class ContactPointDiscord
    {
        /// <summary>
        /// The URL of a custom avatar image to use. Defaults to ``.
        /// </summary>
        public readonly string? AvatarUrl;
        /// <summary>
        /// Whether to disable sending resolve messages. Defaults to `false`.
        /// </summary>
        public readonly bool? DisableResolveMessage;
        /// <summary>
        /// The templated content of the message. Defaults to ``.
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// Additional custom properties to attach to the notifier. Defaults to `map[]`.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Settings;
        /// <summary>
        /// The templated content of the title.
        /// </summary>
        public readonly string? Title;
        /// <summary>
        /// The UID of the contact point.
        /// </summary>
        public readonly string? Uid;
        /// <summary>
        /// The discord webhook URL.
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// Whether to use the bot account's plain username instead of "Grafana." Defaults to `false`.
        /// </summary>
        public readonly bool? UseDiscordUsername;

        [OutputConstructor]
        private ContactPointDiscord(
            string? avatarUrl,

            bool? disableResolveMessage,

            string? message,

            ImmutableDictionary<string, string>? settings,

            string? title,

            string? uid,

            string url,

            bool? useDiscordUsername)
        {
            AvatarUrl = avatarUrl;
            DisableResolveMessage = disableResolveMessage;
            Message = message;
            Settings = settings;
            Title = title;
            Uid = uid;
            Url = url;
            UseDiscordUsername = useDiscordUsername;
        }
    }
}
