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

    public sealed class SyntheticMonitoringCheckSettingsMultihttpEntryVariableGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The attribute to use when finding the variable value. Only used when type is CSS_SELECTOR
        /// </summary>
        [Input("attribute")]
        public Input<string>? Attribute { get; set; }

        /// <summary>
        /// The expression of the assertion. Should start with $.
        /// </summary>
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        /// <summary>
        /// Name of the header to send
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of assertion to make: TEXT, JSON*PATH*VALUE, JSON*PATH*ASSERTION, REGEX_ASSERTION
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public SyntheticMonitoringCheckSettingsMultihttpEntryVariableGetArgs()
        {
        }
        public static new SyntheticMonitoringCheckSettingsMultihttpEntryVariableGetArgs Empty => new SyntheticMonitoringCheckSettingsMultihttpEntryVariableGetArgs();
    }
}
