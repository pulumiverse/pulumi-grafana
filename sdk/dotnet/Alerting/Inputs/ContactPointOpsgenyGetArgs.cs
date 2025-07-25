// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Alerting.Inputs
{

    public sealed class ContactPointOpsgenyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiKey", required: true)]
        private Input<string>? _apiKey;

        /// <summary>
        /// The OpsGenie API key to use.
        /// </summary>
        public Input<string>? ApiKey
        {
            get => _apiKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Whether to auto-close alerts in OpsGenie when they resolve in the Alertmanager.
        /// </summary>
        [Input("autoClose")]
        public Input<bool>? AutoClose { get; set; }

        /// <summary>
        /// A templated high-level description to use for the alert.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to disable sending resolve messages. Defaults to `false`.
        /// </summary>
        [Input("disableResolveMessage")]
        public Input<bool>? DisableResolveMessage { get; set; }

        /// <summary>
        /// The templated content of the message.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// Whether to allow the alert priority to be configured via the value of the `og_priority` annotation on the alert.
        /// </summary>
        [Input("overridePriority")]
        public Input<bool>? OverridePriority { get; set; }

        [Input("responders")]
        private InputList<Inputs.ContactPointOpsgenyResponderGetArgs>? _responders;

        /// <summary>
        /// Teams, users, escalations and schedules that the alert will be routed to send notifications. If the API Key belongs to a team integration, this field will be overwritten with the owner team. This feature is available from Grafana 10.3+.
        /// </summary>
        public InputList<Inputs.ContactPointOpsgenyResponderGetArgs> Responders
        {
            get => _responders ?? (_responders = new InputList<Inputs.ContactPointOpsgenyResponderGetArgs>());
            set => _responders = value;
        }

        /// <summary>
        /// Whether to send annotations to OpsGenie as Tags, Details, or both. Supported values are `tags`, `details`, `both`, or empty to use the default behavior of Tags.
        /// </summary>
        [Input("sendTagsAs")]
        public Input<string>? SendTagsAs { get; set; }

        [Input("settings")]
        private InputMap<string>? _settings;

        /// <summary>
        /// Additional custom properties to attach to the notifier. Defaults to `map[]`.
        /// </summary>
        public InputMap<string> Settings
        {
            get => _settings ?? (_settings = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _settings = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// The UID of the contact point.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// Allows customization of the OpsGenie API URL.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ContactPointOpsgenyGetArgs()
        {
        }
        public static new ContactPointOpsgenyGetArgs Empty => new ContactPointOpsgenyGetArgs();
    }
}
