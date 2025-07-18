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

    public sealed class SyntheticMonitoringCheckSettingsMultihttpEntryRequestHeaderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the header to send
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Value of the header to send
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public SyntheticMonitoringCheckSettingsMultihttpEntryRequestHeaderArgs()
        {
        }
        public static new SyntheticMonitoringCheckSettingsMultihttpEntryRequestHeaderArgs Empty => new SyntheticMonitoringCheckSettingsMultihttpEntryRequestHeaderArgs();
    }
}
