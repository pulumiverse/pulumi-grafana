// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Grafana.Inputs
{

    public sealed class ContactPointKafkaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API version to use when contacting the Kafka REST Server. Supported: v2 (default) and v3. Defaults to `v2`.
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// The Id of cluster to use when contacting the Kafka REST Server. Required api_version to be 'v3'
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The templated description of the Kafka message.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The templated details to include with the message.
        /// </summary>
        [Input("details")]
        public Input<string>? Details { get; set; }

        /// <summary>
        /// Whether to disable sending resolve messages. Defaults to `false`.
        /// </summary>
        [Input("disableResolveMessage")]
        public Input<bool>? DisableResolveMessage { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The password to use when making a call to the Kafka REST Proxy
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("restProxyUrl", required: true)]
        private Input<string>? _restProxyUrl;

        /// <summary>
        /// The URL of the Kafka REST proxy to send requests to.
        /// </summary>
        public Input<string>? RestProxyUrl
        {
            get => _restProxyUrl;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _restProxyUrl = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

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
        /// The name of the Kafka topic to publish to.
        /// </summary>
        [Input("topic", required: true)]
        public Input<string> Topic { get; set; } = null!;

        /// <summary>
        /// The UID of the contact point.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// The user name to use when making a call to the Kafka REST Proxy
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ContactPointKafkaArgs()
        {
        }
        public static new ContactPointKafkaArgs Empty => new ContactPointKafkaArgs();
    }
}
