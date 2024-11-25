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

    public sealed class CheckSettingsMultihttpEntryAssertionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The condition of the assertion: NOT*CONTAINS, EQUALS, STARTS*WITH, ENDS*WITH, TYPE*OF, CONTAINS
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        /// <summary>
        /// The expression of the assertion. Should start with $.
        /// </summary>
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        /// <summary>
        /// The subject of the assertion: RESPONSE*HEADERS, HTTP*STATUS*CODE, RESPONSE*BODY
        /// </summary>
        [Input("subject")]
        public Input<string>? Subject { get; set; }

        /// <summary>
        /// The type of assertion to make: TEXT, JSON*PATH*VALUE, JSON*PATH*ASSERTION, REGEX_ASSERTION
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The value of the assertion
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public CheckSettingsMultihttpEntryAssertionArgs()
        {
        }
        public static new CheckSettingsMultihttpEntryAssertionArgs Empty => new CheckSettingsMultihttpEntryAssertionArgs();
    }
}