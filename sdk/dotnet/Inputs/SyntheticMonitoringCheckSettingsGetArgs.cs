// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Inputs
{

    public sealed class SyntheticMonitoringCheckSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Settings for browser check. See https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/create-checks/checks/k6-browser/.
        /// </summary>
        [Input("browser")]
        public Input<Inputs.SyntheticMonitoringCheckSettingsBrowserGetArgs>? Browser { get; set; }

        /// <summary>
        /// Settings for DNS check. The target must be a valid hostname (or IP address for `PTR` records).
        /// </summary>
        [Input("dns")]
        public Input<Inputs.SyntheticMonitoringCheckSettingsDnsGetArgs>? Dns { get; set; }

        /// <summary>
        /// Settings for gRPC Health check. The target must be of the form `&lt;host&gt;:&lt;port&gt;`, where the host portion must be a valid hostname or IP address.
        /// </summary>
        [Input("grpc")]
        public Input<Inputs.SyntheticMonitoringCheckSettingsGrpcGetArgs>? Grpc { get; set; }

        /// <summary>
        /// Settings for HTTP check. The target must be a URL (http or https).
        /// </summary>
        [Input("http")]
        public Input<Inputs.SyntheticMonitoringCheckSettingsHttpGetArgs>? Http { get; set; }

        /// <summary>
        /// Settings for MultiHTTP check. The target must be a URL (http or https)
        /// </summary>
        [Input("multihttp")]
        public Input<Inputs.SyntheticMonitoringCheckSettingsMultihttpGetArgs>? Multihttp { get; set; }

        /// <summary>
        /// Settings for ping (ICMP) check. The target must be a valid hostname or IP address.
        /// </summary>
        [Input("ping")]
        public Input<Inputs.SyntheticMonitoringCheckSettingsPingGetArgs>? Ping { get; set; }

        /// <summary>
        /// Settings for scripted check. See https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/create-checks/checks/k6/.
        /// </summary>
        [Input("scripted")]
        public Input<Inputs.SyntheticMonitoringCheckSettingsScriptedGetArgs>? Scripted { get; set; }

        /// <summary>
        /// Settings for TCP check. The target must be of the form `&lt;host&gt;:&lt;port&gt;`, where the host portion must be a valid hostname or IP address.
        /// </summary>
        [Input("tcp")]
        public Input<Inputs.SyntheticMonitoringCheckSettingsTcpGetArgs>? Tcp { get; set; }

        /// <summary>
        /// Settings for traceroute check. The target must be a valid hostname or IP address
        /// </summary>
        [Input("traceroute")]
        public Input<Inputs.SyntheticMonitoringCheckSettingsTracerouteGetArgs>? Traceroute { get; set; }

        public SyntheticMonitoringCheckSettingsGetArgs()
        {
        }
        public static new SyntheticMonitoringCheckSettingsGetArgs Empty => new SyntheticMonitoringCheckSettingsGetArgs();
    }
}
