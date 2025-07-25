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

    public sealed class RuleGroupRuleDataGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UID of the datasource being queried, or "-100" if this stage is an expression stage.
        /// </summary>
        [Input("datasourceUid", required: true)]
        public Input<string> DatasourceUid { get; set; } = null!;

        /// <summary>
        /// Custom JSON data to send to the specified datasource when querying.
        /// </summary>
        [Input("model", required: true)]
        public Input<string> Model { get; set; } = null!;

        /// <summary>
        /// An optional identifier for the type of query being executed. Defaults to ``.
        /// </summary>
        [Input("queryType")]
        public Input<string>? QueryType { get; set; }

        /// <summary>
        /// A unique string to identify this query stage within a rule.
        /// </summary>
        [Input("refId", required: true)]
        public Input<string> RefId { get; set; } = null!;

        /// <summary>
        /// The time range, relative to when the query is executed, across which to query.
        /// </summary>
        [Input("relativeTimeRange", required: true)]
        public Input<Inputs.RuleGroupRuleDataRelativeTimeRangeGetArgs> RelativeTimeRange { get; set; } = null!;

        public RuleGroupRuleDataGetArgs()
        {
        }
        public static new RuleGroupRuleDataGetArgs Empty => new RuleGroupRuleDataGetArgs();
    }
}
