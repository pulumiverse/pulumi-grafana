// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Grafana.Outputs
{

    [OutputType]
    public sealed class ContactPointGooglechat
    {
        /// <summary>
        /// Whether to disable sending resolve messages. Defaults to `false`.
        /// </summary>
        public readonly bool? DisableResolveMessage;
        /// <summary>
        /// The templated content of the message.
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
        /// The Google Chat webhook URL.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private ContactPointGooglechat(
            bool? disableResolveMessage,

            string? message,

            ImmutableDictionary<string, string>? settings,

            string? title,

            string? uid,

            string url)
        {
            DisableResolveMessage = disableResolveMessage;
            Message = message;
            Settings = settings;
            Title = title;
            Uid = uid;
            Url = url;
        }
    }
}
