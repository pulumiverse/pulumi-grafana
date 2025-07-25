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
    public sealed class ContactPointTelegram
    {
        /// <summary>
        /// The chat ID to send messages to.
        /// </summary>
        public readonly string ChatId;
        /// <summary>
        /// When set users will receive a notification with no sound.
        /// </summary>
        public readonly bool? DisableNotifications;
        /// <summary>
        /// Whether to disable sending resolve messages. Defaults to `false`.
        /// </summary>
        public readonly bool? DisableResolveMessage;
        /// <summary>
        /// When set it disables link previews for links in the message.
        /// </summary>
        public readonly bool? DisableWebPagePreview;
        /// <summary>
        /// The templated content of the message.
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// The ID of the message thread to send the message to.
        /// </summary>
        public readonly string? MessageThreadId;
        /// <summary>
        /// Mode for parsing entities in the message text. Supported: None, Markdown, MarkdownV2, and HTML. HTML is the default.
        /// </summary>
        public readonly string? ParseMode;
        /// <summary>
        /// When set it protects the contents of the message from forwarding and saving.
        /// </summary>
        public readonly bool? ProtectContent;
        /// <summary>
        /// Additional custom properties to attach to the notifier. Defaults to `map[]`.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Settings;
        /// <summary>
        /// The Telegram bot token.
        /// </summary>
        public readonly string Token;
        /// <summary>
        /// The UID of the contact point.
        /// </summary>
        public readonly string? Uid;

        [OutputConstructor]
        private ContactPointTelegram(
            string chatId,

            bool? disableNotifications,

            bool? disableResolveMessage,

            bool? disableWebPagePreview,

            string? message,

            string? messageThreadId,

            string? parseMode,

            bool? protectContent,

            ImmutableDictionary<string, string>? settings,

            string token,

            string? uid)
        {
            ChatId = chatId;
            DisableNotifications = disableNotifications;
            DisableResolveMessage = disableResolveMessage;
            DisableWebPagePreview = disableWebPagePreview;
            Message = message;
            MessageThreadId = messageThreadId;
            ParseMode = parseMode;
            ProtectContent = protectContent;
            Settings = settings;
            Token = token;
            Uid = uid;
        }
    }
}
