// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Slo.Inputs
{

    public sealed class SLOAlertingAnnotationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key for filtering and identification
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Templatable value
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public SLOAlertingAnnotationGetArgs()
        {
        }
        public static new SLOAlertingAnnotationGetArgs Empty => new SLOAlertingAnnotationGetArgs();
    }
}
