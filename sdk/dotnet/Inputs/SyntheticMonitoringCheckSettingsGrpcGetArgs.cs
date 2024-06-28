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

    public sealed class SyntheticMonitoringCheckSettingsGrpcGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Options are `V4`, `V6`, `Any`. Specifies whether the corresponding check will be performed using IPv4 or IPv6. The `Any` value indicates that IPv6 should be used, falling back to IPv4 if that's not available. Defaults to `V4`.
        /// </summary>
        [Input("ipVersion")]
        public Input<string>? IpVersion { get; set; }

        /// <summary>
        /// gRPC service.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// Whether or not TLS is used when the connection is initiated. Defaults to `false`.
        /// </summary>
        [Input("tls")]
        public Input<bool>? Tls { get; set; }

        /// <summary>
        /// TLS config.
        /// </summary>
        [Input("tlsConfig")]
        public Input<Inputs.SyntheticMonitoringCheckSettingsGrpcTlsConfigGetArgs>? TlsConfig { get; set; }

        public SyntheticMonitoringCheckSettingsGrpcGetArgs()
        {
        }
        public static new SyntheticMonitoringCheckSettingsGrpcGetArgs Empty => new SyntheticMonitoringCheckSettingsGrpcGetArgs();
    }
}
