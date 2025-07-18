// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.SyntheticMonitoring.Inputs
{

    public sealed class CheckSettingsTcpQueryResponseGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Response to expect.
        /// </summary>
        [Input("expect", required: true)]
        public Input<string> Expect { get; set; } = null!;

        /// <summary>
        /// Data to send.
        /// </summary>
        [Input("send", required: true)]
        public Input<string> Send { get; set; } = null!;

        /// <summary>
        /// Upgrade TCP connection to TLS. Defaults to `false`.
        /// </summary>
        [Input("startTls")]
        public Input<bool>? StartTls { get; set; }

        public CheckSettingsTcpQueryResponseGetArgs()
        {
        }
        public static new CheckSettingsTcpQueryResponseGetArgs Empty => new CheckSettingsTcpQueryResponseGetArgs();
    }
}
