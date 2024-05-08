// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Outputs
{

    [OutputType]
    public sealed class NotificationPolicyPolicyPolicyPolicy
    {
        /// <summary>
        /// The contact point to route notifications that match this rule to.
        /// </summary>
        public readonly string? ContactPoint;
        /// <summary>
        /// Whether to continue matching subsequent rules if an alert matches the current rule. Otherwise, the rule will be 'consumed' by the first policy to match it.
        /// </summary>
        public readonly bool? Continue;
        /// <summary>
        /// A list of alert labels to group alerts into notifications by. Use the special label `...` to group alerts by all labels, effectively disabling grouping. Required for root policy only. If empty, the parent grouping is used.
        /// </summary>
        public readonly ImmutableArray<string> GroupBies;
        /// <summary>
        /// Minimum time interval between two notifications for the same group. Default is 5 minutes.
        /// </summary>
        public readonly string? GroupInterval;
        /// <summary>
        /// Time to wait to buffer alerts of the same group before sending a notification. Default is 30 seconds.
        /// </summary>
        public readonly string? GroupWait;
        /// <summary>
        /// Describes which labels this rule should match. When multiple matchers are supplied, an alert must match ALL matchers to be accepted by this policy. When no matchers are supplied, the rule will match all alert instances.
        /// </summary>
        public readonly ImmutableArray<Outputs.NotificationPolicyPolicyPolicyPolicyMatcher> Matchers;
        /// <summary>
        /// A list of mute timing names to apply to alerts that match this policy.
        /// </summary>
        public readonly ImmutableArray<string> MuteTimings;
        /// <summary>
        /// Routing rules for specific label sets.
        /// </summary>
        public readonly ImmutableArray<Outputs.NotificationPolicyPolicyPolicyPolicyPolicy> Policies;
        /// <summary>
        /// Minimum time interval for re-sending a notification if an alert is still firing. Default is 4 hours.
        /// </summary>
        public readonly string? RepeatInterval;

        [OutputConstructor]
        private NotificationPolicyPolicyPolicyPolicy(
            string? contactPoint,

            bool? @continue,

            ImmutableArray<string> groupBies,

            string? groupInterval,

            string? groupWait,

            ImmutableArray<Outputs.NotificationPolicyPolicyPolicyPolicyMatcher> matchers,

            ImmutableArray<string> muteTimings,

            ImmutableArray<Outputs.NotificationPolicyPolicyPolicyPolicyPolicy> policies,

            string? repeatInterval)
        {
            ContactPoint = contactPoint;
            Continue = @continue;
            GroupBies = groupBies;
            GroupInterval = groupInterval;
            GroupWait = groupWait;
            Matchers = matchers;
            MuteTimings = muteTimings;
            Policies = policies;
            RepeatInterval = repeatInterval;
        }
    }
}
