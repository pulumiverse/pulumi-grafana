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

    public sealed class SyntheticMonitoringCheckSettingsMultihttpEntryRequestQueryFieldArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the query field to send
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Value of the query field to send
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public SyntheticMonitoringCheckSettingsMultihttpEntryRequestQueryFieldArgs()
        {
        }
        public static new SyntheticMonitoringCheckSettingsMultihttpEntryRequestQueryFieldArgs Empty => new SyntheticMonitoringCheckSettingsMultihttpEntryRequestQueryFieldArgs();
    }
}
