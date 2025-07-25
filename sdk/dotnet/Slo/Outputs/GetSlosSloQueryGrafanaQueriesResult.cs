// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Slo.Outputs
{

    [OutputType]
    public sealed class GetSlosSloQueryGrafanaQueriesResult
    {
        /// <summary>
        /// Query Object - Array of Grafana Query JSON objects
        /// </summary>
        public readonly string GrafanaQueries;

        [OutputConstructor]
        private GetSlosSloQueryGrafanaQueriesResult(string grafanaQueries)
        {
            GrafanaQueries = grafanaQueries;
        }
    }
}
