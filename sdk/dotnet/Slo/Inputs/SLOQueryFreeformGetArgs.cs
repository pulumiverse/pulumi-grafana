// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Slo.Inputs
{

    public sealed class SLOQueryFreeformGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Freeform Query Field - valid promQl
        /// </summary>
        [Input("query", required: true)]
        public Input<string> Query { get; set; } = null!;

        public SLOQueryFreeformGetArgs()
        {
        }
        public static new SLOQueryFreeformGetArgs Empty => new SLOQueryFreeformGetArgs();
    }
}
