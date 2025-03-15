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

    public sealed class SLOQueryArgs : global::Pulumi.ResourceArgs
    {
        [Input("freeform")]
        public Input<Inputs.SLOQueryFreeformArgs>? Freeform { get; set; }

        /// <summary>
        /// Array for holding a set of grafana queries
        /// </summary>
        [Input("grafanaQueries")]
        public Input<Inputs.SLOQueryGrafanaQueriesArgs>? GrafanaQueries { get; set; }

        [Input("ratio")]
        public Input<Inputs.SLOQueryRatioArgs>? Ratio { get; set; }

        /// <summary>
        /// Query type must be one of: "freeform", "query", "ratio", "grafana_queries" or "threshold"
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public SLOQueryArgs()
        {
        }
        public static new SLOQueryArgs Empty => new SLOQueryArgs();
    }
}
