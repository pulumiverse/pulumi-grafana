// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Alerting.Outputs
{

    [OutputType]
    public sealed class ContactPointEmail
    {
        /// <summary>
        /// The addresses to send emails to.
        /// </summary>
        public readonly ImmutableArray<string> Addresses;
        /// <summary>
        /// Whether to disable sending resolve messages. Defaults to `false`.
        /// </summary>
        public readonly bool? DisableResolveMessage;
        /// <summary>
        /// The templated content of the email. Defaults to ``.
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// Additional custom properties to attach to the notifier. Defaults to `map[]`.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Settings;
        /// <summary>
        /// Whether to send a single email CC'ing all addresses, rather than a separate email to each address. Defaults to `false`.
        /// </summary>
        public readonly bool? SingleEmail;
        /// <summary>
        /// The templated subject line of the email. Defaults to ``.
        /// </summary>
        public readonly string? Subject;
        /// <summary>
        /// The UID of the contact point.
        /// </summary>
        public readonly string? Uid;

        [OutputConstructor]
        private ContactPointEmail(
            ImmutableArray<string> addresses,

            bool? disableResolveMessage,

            string? message,

            ImmutableDictionary<string, string>? settings,

            bool? singleEmail,

            string? subject,

            string? uid)
        {
            Addresses = addresses;
            DisableResolveMessage = disableResolveMessage;
            Message = message;
            Settings = settings;
            SingleEmail = singleEmail;
            Subject = subject;
            Uid = uid;
        }
    }
}
