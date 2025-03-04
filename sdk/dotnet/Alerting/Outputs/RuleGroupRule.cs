// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
    public sealed class RuleGroupRule
    {
        /// <summary>
        /// Key-value pairs of metadata to attach to the alert rule. They add additional information, such as a `summary` or `runbook_url`, to help identify and investigate alerts. The `dashboardUId` and `panelId` annotations, which link alerts to a panel, must be set together. Defaults to `map[]`.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Annotations;
        /// <summary>
        /// The `ref_id` of the query node in the `data` field to use as the alert condition.
        /// </summary>
        public readonly string? Condition;
        /// <summary>
        /// A sequence of stages that describe the contents of the rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleGroupRuleData> Datas;
        /// <summary>
        /// Describes what state to enter when the rule's query is invalid and the rule cannot be executed. Options are OK, Error, KeepLast, and Alerting.  Defaults to Alerting if not set.
        /// </summary>
        public readonly string? ExecErrState;
        /// <summary>
        /// The amount of time for which the rule must be breached for the rule to be considered to be Firing. Before this time has elapsed, the rule is only considered to be Pending. Defaults to `0`.
        /// </summary>
        public readonly string? For;
        /// <summary>
        /// Sets whether the alert should be paused or not. Defaults to `false`.
        /// </summary>
        public readonly bool? IsPaused;
        /// <summary>
        /// Key-value pairs to attach to the alert rule that can be used in matching, grouping, and routing. Defaults to `map[]`.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;
        /// <summary>
        /// The name of the alert rule.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Describes what state to enter when the rule's query returns No Data. Options are OK, NoData, KeepLast, and Alerting. Defaults to NoData if not set.
        /// </summary>
        public readonly string? NoDataState;
        /// <summary>
        /// Notification settings for the rule. If specified, it overrides the notification policies. Available since Grafana 10.4, requires feature flag 'alertingSimplifiedRouting' to be enabled.
        /// </summary>
        public readonly Outputs.RuleGroupRuleNotificationSettings? NotificationSettings;
        /// <summary>
        /// Settings for a recording rule. Available since Grafana 11.2, requires feature flag 'grafanaManagedRecordingRules' to be enabled.
        /// </summary>
        public readonly Outputs.RuleGroupRuleRecord? Record;
        /// <summary>
        /// The unique identifier of the alert rule.
        /// </summary>
        public readonly string? Uid;

        [OutputConstructor]
        private RuleGroupRule(
            ImmutableDictionary<string, string>? annotations,

            string? condition,

            ImmutableArray<Outputs.RuleGroupRuleData> datas,

            string? execErrState,

            string? @for,

            bool? isPaused,

            ImmutableDictionary<string, string>? labels,

            string name,

            string? noDataState,

            Outputs.RuleGroupRuleNotificationSettings? notificationSettings,

            Outputs.RuleGroupRuleRecord? record,

            string? uid)
        {
            Annotations = annotations;
            Condition = condition;
            Datas = datas;
            ExecErrState = execErrState;
            For = @for;
            IsPaused = isPaused;
            Labels = labels;
            Name = name;
            NoDataState = noDataState;
            NotificationSettings = notificationSettings;
            Record = record;
            Uid = uid;
        }
    }
}
