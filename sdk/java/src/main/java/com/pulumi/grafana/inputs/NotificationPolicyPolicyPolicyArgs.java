// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.grafana.inputs.NotificationPolicyPolicyPolicyMatcherArgs;
import com.pulumi.grafana.inputs.NotificationPolicyPolicyPolicyPolicyArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NotificationPolicyPolicyPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final NotificationPolicyPolicyPolicyArgs Empty = new NotificationPolicyPolicyPolicyArgs();

    /**
     * The contact point to route notifications that match this rule to.
     * 
     */
    @Import(name="contactPoint", required=true)
    private Output<String> contactPoint;

    /**
     * @return The contact point to route notifications that match this rule to.
     * 
     */
    public Output<String> contactPoint() {
        return this.contactPoint;
    }

    /**
     * Whether to continue matching subsequent rules if an alert matches the current rule. Otherwise, the rule will be &#39;consumed&#39; by the first policy to match it.
     * 
     */
    @Import(name="continue")
    private @Nullable Output<Boolean> continue_;

    /**
     * @return Whether to continue matching subsequent rules if an alert matches the current rule. Otherwise, the rule will be &#39;consumed&#39; by the first policy to match it.
     * 
     */
    public Optional<Output<Boolean>> continue_() {
        return Optional.ofNullable(this.continue_);
    }

    /**
     * A list of alert labels to group alerts into notifications by. Use the special label `...` to group alerts by all labels, effectively disabling grouping. Required for root policy only. If empty, the parent grouping is used.
     * 
     */
    @Import(name="groupBies")
    private @Nullable Output<List<String>> groupBies;

    /**
     * @return A list of alert labels to group alerts into notifications by. Use the special label `...` to group alerts by all labels, effectively disabling grouping. Required for root policy only. If empty, the parent grouping is used.
     * 
     */
    public Optional<Output<List<String>>> groupBies() {
        return Optional.ofNullable(this.groupBies);
    }

    /**
     * Minimum time interval between two notifications for the same group. Default is 5 minutes.
     * 
     */
    @Import(name="groupInterval")
    private @Nullable Output<String> groupInterval;

    /**
     * @return Minimum time interval between two notifications for the same group. Default is 5 minutes.
     * 
     */
    public Optional<Output<String>> groupInterval() {
        return Optional.ofNullable(this.groupInterval);
    }

    /**
     * Time to wait to buffer alerts of the same group before sending a notification. Default is 30 seconds.
     * 
     */
    @Import(name="groupWait")
    private @Nullable Output<String> groupWait;

    /**
     * @return Time to wait to buffer alerts of the same group before sending a notification. Default is 30 seconds.
     * 
     */
    public Optional<Output<String>> groupWait() {
        return Optional.ofNullable(this.groupWait);
    }

    /**
     * Describes which labels this rule should match. When multiple matchers are supplied, an alert must match ALL matchers to be accepted by this policy. When no matchers are supplied, the rule will match all alert instances.
     * 
     */
    @Import(name="matchers")
    private @Nullable Output<List<NotificationPolicyPolicyPolicyMatcherArgs>> matchers;

    /**
     * @return Describes which labels this rule should match. When multiple matchers are supplied, an alert must match ALL matchers to be accepted by this policy. When no matchers are supplied, the rule will match all alert instances.
     * 
     */
    public Optional<Output<List<NotificationPolicyPolicyPolicyMatcherArgs>>> matchers() {
        return Optional.ofNullable(this.matchers);
    }

    /**
     * A list of mute timing names to apply to alerts that match this policy.
     * 
     */
    @Import(name="muteTimings")
    private @Nullable Output<List<String>> muteTimings;

    /**
     * @return A list of mute timing names to apply to alerts that match this policy.
     * 
     */
    public Optional<Output<List<String>>> muteTimings() {
        return Optional.ofNullable(this.muteTimings);
    }

    /**
     * Routing rules for specific label sets.
     * 
     */
    @Import(name="policies")
    private @Nullable Output<List<NotificationPolicyPolicyPolicyPolicyArgs>> policies;

    /**
     * @return Routing rules for specific label sets.
     * 
     */
    public Optional<Output<List<NotificationPolicyPolicyPolicyPolicyArgs>>> policies() {
        return Optional.ofNullable(this.policies);
    }

    /**
     * Minimum time interval for re-sending a notification if an alert is still firing. Default is 4 hours.
     * 
     */
    @Import(name="repeatInterval")
    private @Nullable Output<String> repeatInterval;

    /**
     * @return Minimum time interval for re-sending a notification if an alert is still firing. Default is 4 hours.
     * 
     */
    public Optional<Output<String>> repeatInterval() {
        return Optional.ofNullable(this.repeatInterval);
    }

    private NotificationPolicyPolicyPolicyArgs() {}

    private NotificationPolicyPolicyPolicyArgs(NotificationPolicyPolicyPolicyArgs $) {
        this.contactPoint = $.contactPoint;
        this.continue_ = $.continue_;
        this.groupBies = $.groupBies;
        this.groupInterval = $.groupInterval;
        this.groupWait = $.groupWait;
        this.matchers = $.matchers;
        this.muteTimings = $.muteTimings;
        this.policies = $.policies;
        this.repeatInterval = $.repeatInterval;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NotificationPolicyPolicyPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NotificationPolicyPolicyPolicyArgs $;

        public Builder() {
            $ = new NotificationPolicyPolicyPolicyArgs();
        }

        public Builder(NotificationPolicyPolicyPolicyArgs defaults) {
            $ = new NotificationPolicyPolicyPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param contactPoint The contact point to route notifications that match this rule to.
         * 
         * @return builder
         * 
         */
        public Builder contactPoint(Output<String> contactPoint) {
            $.contactPoint = contactPoint;
            return this;
        }

        /**
         * @param contactPoint The contact point to route notifications that match this rule to.
         * 
         * @return builder
         * 
         */
        public Builder contactPoint(String contactPoint) {
            return contactPoint(Output.of(contactPoint));
        }

        /**
         * @param continue_ Whether to continue matching subsequent rules if an alert matches the current rule. Otherwise, the rule will be &#39;consumed&#39; by the first policy to match it.
         * 
         * @return builder
         * 
         */
        public Builder continue_(@Nullable Output<Boolean> continue_) {
            $.continue_ = continue_;
            return this;
        }

        /**
         * @param continue_ Whether to continue matching subsequent rules if an alert matches the current rule. Otherwise, the rule will be &#39;consumed&#39; by the first policy to match it.
         * 
         * @return builder
         * 
         */
        public Builder continue_(Boolean continue_) {
            return continue_(Output.of(continue_));
        }

        /**
         * @param groupBies A list of alert labels to group alerts into notifications by. Use the special label `...` to group alerts by all labels, effectively disabling grouping. Required for root policy only. If empty, the parent grouping is used.
         * 
         * @return builder
         * 
         */
        public Builder groupBies(@Nullable Output<List<String>> groupBies) {
            $.groupBies = groupBies;
            return this;
        }

        /**
         * @param groupBies A list of alert labels to group alerts into notifications by. Use the special label `...` to group alerts by all labels, effectively disabling grouping. Required for root policy only. If empty, the parent grouping is used.
         * 
         * @return builder
         * 
         */
        public Builder groupBies(List<String> groupBies) {
            return groupBies(Output.of(groupBies));
        }

        /**
         * @param groupBies A list of alert labels to group alerts into notifications by. Use the special label `...` to group alerts by all labels, effectively disabling grouping. Required for root policy only. If empty, the parent grouping is used.
         * 
         * @return builder
         * 
         */
        public Builder groupBies(String... groupBies) {
            return groupBies(List.of(groupBies));
        }

        /**
         * @param groupInterval Minimum time interval between two notifications for the same group. Default is 5 minutes.
         * 
         * @return builder
         * 
         */
        public Builder groupInterval(@Nullable Output<String> groupInterval) {
            $.groupInterval = groupInterval;
            return this;
        }

        /**
         * @param groupInterval Minimum time interval between two notifications for the same group. Default is 5 minutes.
         * 
         * @return builder
         * 
         */
        public Builder groupInterval(String groupInterval) {
            return groupInterval(Output.of(groupInterval));
        }

        /**
         * @param groupWait Time to wait to buffer alerts of the same group before sending a notification. Default is 30 seconds.
         * 
         * @return builder
         * 
         */
        public Builder groupWait(@Nullable Output<String> groupWait) {
            $.groupWait = groupWait;
            return this;
        }

        /**
         * @param groupWait Time to wait to buffer alerts of the same group before sending a notification. Default is 30 seconds.
         * 
         * @return builder
         * 
         */
        public Builder groupWait(String groupWait) {
            return groupWait(Output.of(groupWait));
        }

        /**
         * @param matchers Describes which labels this rule should match. When multiple matchers are supplied, an alert must match ALL matchers to be accepted by this policy. When no matchers are supplied, the rule will match all alert instances.
         * 
         * @return builder
         * 
         */
        public Builder matchers(@Nullable Output<List<NotificationPolicyPolicyPolicyMatcherArgs>> matchers) {
            $.matchers = matchers;
            return this;
        }

        /**
         * @param matchers Describes which labels this rule should match. When multiple matchers are supplied, an alert must match ALL matchers to be accepted by this policy. When no matchers are supplied, the rule will match all alert instances.
         * 
         * @return builder
         * 
         */
        public Builder matchers(List<NotificationPolicyPolicyPolicyMatcherArgs> matchers) {
            return matchers(Output.of(matchers));
        }

        /**
         * @param matchers Describes which labels this rule should match. When multiple matchers are supplied, an alert must match ALL matchers to be accepted by this policy. When no matchers are supplied, the rule will match all alert instances.
         * 
         * @return builder
         * 
         */
        public Builder matchers(NotificationPolicyPolicyPolicyMatcherArgs... matchers) {
            return matchers(List.of(matchers));
        }

        /**
         * @param muteTimings A list of mute timing names to apply to alerts that match this policy.
         * 
         * @return builder
         * 
         */
        public Builder muteTimings(@Nullable Output<List<String>> muteTimings) {
            $.muteTimings = muteTimings;
            return this;
        }

        /**
         * @param muteTimings A list of mute timing names to apply to alerts that match this policy.
         * 
         * @return builder
         * 
         */
        public Builder muteTimings(List<String> muteTimings) {
            return muteTimings(Output.of(muteTimings));
        }

        /**
         * @param muteTimings A list of mute timing names to apply to alerts that match this policy.
         * 
         * @return builder
         * 
         */
        public Builder muteTimings(String... muteTimings) {
            return muteTimings(List.of(muteTimings));
        }

        /**
         * @param policies Routing rules for specific label sets.
         * 
         * @return builder
         * 
         */
        public Builder policies(@Nullable Output<List<NotificationPolicyPolicyPolicyPolicyArgs>> policies) {
            $.policies = policies;
            return this;
        }

        /**
         * @param policies Routing rules for specific label sets.
         * 
         * @return builder
         * 
         */
        public Builder policies(List<NotificationPolicyPolicyPolicyPolicyArgs> policies) {
            return policies(Output.of(policies));
        }

        /**
         * @param policies Routing rules for specific label sets.
         * 
         * @return builder
         * 
         */
        public Builder policies(NotificationPolicyPolicyPolicyPolicyArgs... policies) {
            return policies(List.of(policies));
        }

        /**
         * @param repeatInterval Minimum time interval for re-sending a notification if an alert is still firing. Default is 4 hours.
         * 
         * @return builder
         * 
         */
        public Builder repeatInterval(@Nullable Output<String> repeatInterval) {
            $.repeatInterval = repeatInterval;
            return this;
        }

        /**
         * @param repeatInterval Minimum time interval for re-sending a notification if an alert is still firing. Default is 4 hours.
         * 
         * @return builder
         * 
         */
        public Builder repeatInterval(String repeatInterval) {
            return repeatInterval(Output.of(repeatInterval));
        }

        public NotificationPolicyPolicyPolicyArgs build() {
            $.contactPoint = Objects.requireNonNull($.contactPoint, "expected parameter 'contactPoint' to be non-null");
            return $;
        }
    }

}
