// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    public sealed class CheckSettingsDns
    {
        /// <summary>
        /// Options are `V4`, `V6`, `Any`. Specifies whether the corresponding check will be performed using IPv4 or IPv6. The `Any` value indicates that IPv6 should be used, falling back to IPv4 if that's not available. Defaults to `V4`.
        /// </summary>
        public readonly string? IpVersion;
        /// <summary>
        /// Port to target. Defaults to `53`.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// `TCP` or `UDP`. Defaults to `UDP`.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// One of `ANY`, `A`, `AAAA`, `CNAME`, `MX`, `NS`, `PTR`, `SOA`, `SRV`, `TXT`. Defaults to `A`.
        /// </summary>
        public readonly string? RecordType;
        /// <summary>
        /// DNS server address to target. Defaults to `8.8.8.8`.
        /// </summary>
        public readonly string? Server;
        /// <summary>
        /// Source IP address.
        /// </summary>
        public readonly string? SourceIpAddress;
        /// <summary>
        /// List of valid response codes. Options include `NOERROR`, `BADALG`, `BADMODE`, `BADKEY`, `BADCOOKIE`, `BADNAME`, `BADSIG`, `BADTIME`, `BADTRUNC`, `BADVERS`, `FORMERR`, `NOTIMP`, `NOTAUTH`, `NOTZONE`, `NXDOMAIN`, `NXRRSET`, `REFUSED`, `SERVFAIL`, `YXDOMAIN`, `YXRRSET`.
        /// </summary>
        public readonly ImmutableArray<string> ValidRCodes;
        /// <summary>
        /// Validate additional matches.
        /// </summary>
        public readonly ImmutableArray<Outputs.CheckSettingsDnsValidateAdditionalRr> ValidateAdditionalRrs;
        /// <summary>
        /// Validate response answer.
        /// </summary>
        public readonly Outputs.CheckSettingsDnsValidateAnswerRrs? ValidateAnswerRrs;
        /// <summary>
        /// Validate response authority.
        /// </summary>
        public readonly Outputs.CheckSettingsDnsValidateAuthorityRrs? ValidateAuthorityRrs;

        [OutputConstructor]
        private CheckSettingsDns(
            string? ipVersion,

            int? port,

            string? protocol,

            string? recordType,

            string? server,

            string? sourceIpAddress,

            ImmutableArray<string> validRCodes,

            ImmutableArray<Outputs.CheckSettingsDnsValidateAdditionalRr> validateAdditionalRrs,

            Outputs.CheckSettingsDnsValidateAnswerRrs? validateAnswerRrs,

            Outputs.CheckSettingsDnsValidateAuthorityRrs? validateAuthorityRrs)
        {
            IpVersion = ipVersion;
            Port = port;
            Protocol = protocol;
            RecordType = recordType;
            Server = server;
            SourceIpAddress = sourceIpAddress;
            ValidRCodes = validRCodes;
            ValidateAdditionalRrs = validateAdditionalRrs;
            ValidateAnswerRrs = validateAnswerRrs;
            ValidateAuthorityRrs = validateAuthorityRrs;
        }
    }
}
