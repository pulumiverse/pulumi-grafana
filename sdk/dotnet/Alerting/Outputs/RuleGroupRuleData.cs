// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Alerting.Outputs
{

    [OutputType]
    public sealed class RuleGroupRuleData
    {
        /// <summary>
        /// The UID of the datasource being queried, or "-100" if this stage is an expression stage.
        /// </summary>
        public readonly string DatasourceUid;
        /// <summary>
        /// Custom JSON data to send to the specified datasource when querying.
        /// </summary>
        public readonly string Model;
        /// <summary>
        /// An optional identifier for the type of query being executed. Defaults to ``.
        /// </summary>
        public readonly string? QueryType;
        /// <summary>
        /// A unique string to identify this query stage within a rule.
        /// </summary>
        public readonly string RefId;
        /// <summary>
        /// The time range, relative to when the query is executed, across which to query.
        /// </summary>
        public readonly Outputs.RuleGroupRuleDataRelativeTimeRange RelativeTimeRange;

        [OutputConstructor]
        private RuleGroupRuleData(
            string datasourceUid,

            string model,

            string? queryType,

            string refId,

            Outputs.RuleGroupRuleDataRelativeTimeRange relativeTimeRange)
        {
            DatasourceUid = datasourceUid;
            Model = model;
            QueryType = queryType;
            RefId = refId;
            RelativeTimeRange = relativeTimeRange;
        }
    }
}
