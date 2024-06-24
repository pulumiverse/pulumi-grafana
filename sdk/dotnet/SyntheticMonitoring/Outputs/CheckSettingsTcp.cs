// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.SyntheticMonitoring.Outputs
{

    [OutputType]
    public sealed class CheckSettingsTcp
    {
        /// <summary>
        /// Options are `V4`, `V6`, `Any`. Specifies whether the corresponding check will be performed using IPv4 or IPv6. The `Any` value indicates that IPv6 should be used, falling back to IPv4 if that's not available. Defaults to `V4`.
        /// </summary>
        public readonly string? IpVersion;
        /// <summary>
        /// The query sent in the TCP probe and the expected associated response.
        /// </summary>
        public readonly ImmutableArray<Outputs.CheckSettingsTcpQueryResponse> QueryResponses;
        /// <summary>
        /// Source IP address.
        /// </summary>
        public readonly string? SourceIpAddress;
        /// <summary>
        /// Whether or not TLS is used when the connection is initiated. Defaults to `false`.
        /// </summary>
        public readonly bool? Tls;
        /// <summary>
        /// TLS config.
        /// </summary>
        public readonly Outputs.CheckSettingsTcpTlsConfig? TlsConfig;

        [OutputConstructor]
        private CheckSettingsTcp(
            string? ipVersion,

            ImmutableArray<Outputs.CheckSettingsTcpQueryResponse> queryResponses,

            string? sourceIpAddress,

            bool? tls,

            Outputs.CheckSettingsTcpTlsConfig? tlsConfig)
        {
            IpVersion = ipVersion;
            QueryResponses = queryResponses;
            SourceIpAddress = sourceIpAddress;
            Tls = tls;
            TlsConfig = tlsConfig;
        }
    }
}