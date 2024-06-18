// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.SyntheticMonitoring.Inputs
{

    public sealed class CheckSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Settings for DNS check. The target must be a valid hostname (or IP address for `PTR` records).
        /// </summary>
        [Input("dns")]
        public Input<Inputs.CheckSettingsDnsGetArgs>? Dns { get; set; }

        /// <summary>
        /// Settings for HTTP check. The target must be a URL (http or https).
        /// </summary>
        [Input("http")]
        public Input<Inputs.CheckSettingsHttpGetArgs>? Http { get; set; }

        /// <summary>
        /// Settings for MultiHTTP check. The target must be a URL (http or https)
        /// </summary>
        [Input("multihttp")]
        public Input<Inputs.CheckSettingsMultihttpGetArgs>? Multihttp { get; set; }

        /// <summary>
        /// Settings for ping (ICMP) check. The target must be a valid hostname or IP address.
        /// </summary>
        [Input("ping")]
        public Input<Inputs.CheckSettingsPingGetArgs>? Ping { get; set; }

        /// <summary>
        /// Settings for scripted check. See https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/create-checks/checks/k6/.
        /// </summary>
        [Input("scripted")]
        public Input<Inputs.CheckSettingsScriptedGetArgs>? Scripted { get; set; }

        /// <summary>
        /// Settings for TCP check. The target must be of the form `&lt;host&gt;:&lt;port&gt;`, where the host portion must be a valid hostname or IP address.
        /// </summary>
        [Input("tcp")]
        public Input<Inputs.CheckSettingsTcpGetArgs>? Tcp { get; set; }

        /// <summary>
        /// Settings for traceroute check. The target must be a valid hostname or IP address
        /// </summary>
        [Input("traceroute")]
        public Input<Inputs.CheckSettingsTracerouteGetArgs>? Traceroute { get; set; }

        public CheckSettingsGetArgs()
        {
        }
        public static new CheckSettingsGetArgs Empty => new CheckSettingsGetArgs();
    }
}
