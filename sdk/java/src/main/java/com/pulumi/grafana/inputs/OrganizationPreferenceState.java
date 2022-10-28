// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OrganizationPreferenceState extends com.pulumi.resources.ResourceArgs {

    public static final OrganizationPreferenceState Empty = new OrganizationPreferenceState();

    /**
     * The Organization home dashboard ID.
     * 
     */
    @Import(name="homeDashboardId")
    private @Nullable Output<Integer> homeDashboardId;

    /**
     * @return The Organization home dashboard ID.
     * 
     */
    public Optional<Output<Integer>> homeDashboardId() {
        return Optional.ofNullable(this.homeDashboardId);
    }

    /**
     * The Organization home dashboard UID.
     * 
     */
    @Import(name="homeDashboardUid")
    private @Nullable Output<String> homeDashboardUid;

    /**
     * @return The Organization home dashboard UID.
     * 
     */
    public Optional<Output<String>> homeDashboardUid() {
        return Optional.ofNullable(this.homeDashboardUid);
    }

    /**
     * The Organization theme. Available values are `light`, `dark`, or an empty string for the default.
     * 
     */
    @Import(name="theme")
    private @Nullable Output<String> theme;

    /**
     * @return The Organization theme. Available values are `light`, `dark`, or an empty string for the default.
     * 
     */
    public Optional<Output<String>> theme() {
        return Optional.ofNullable(this.theme);
    }

    /**
     * The Organization timezone. Available values are `utc`, `browser`, or an empty string for the default.
     * 
     */
    @Import(name="timezone")
    private @Nullable Output<String> timezone;

    /**
     * @return The Organization timezone. Available values are `utc`, `browser`, or an empty string for the default.
     * 
     */
    public Optional<Output<String>> timezone() {
        return Optional.ofNullable(this.timezone);
    }

    /**
     * The Organization week start.
     * 
     */
    @Import(name="weekStart")
    private @Nullable Output<String> weekStart;

    /**
     * @return The Organization week start.
     * 
     */
    public Optional<Output<String>> weekStart() {
        return Optional.ofNullable(this.weekStart);
    }

    private OrganizationPreferenceState() {}

    private OrganizationPreferenceState(OrganizationPreferenceState $) {
        this.homeDashboardId = $.homeDashboardId;
        this.homeDashboardUid = $.homeDashboardUid;
        this.theme = $.theme;
        this.timezone = $.timezone;
        this.weekStart = $.weekStart;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OrganizationPreferenceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OrganizationPreferenceState $;

        public Builder() {
            $ = new OrganizationPreferenceState();
        }

        public Builder(OrganizationPreferenceState defaults) {
            $ = new OrganizationPreferenceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param homeDashboardId The Organization home dashboard ID.
         * 
         * @return builder
         * 
         */
        public Builder homeDashboardId(@Nullable Output<Integer> homeDashboardId) {
            $.homeDashboardId = homeDashboardId;
            return this;
        }

        /**
         * @param homeDashboardId The Organization home dashboard ID.
         * 
         * @return builder
         * 
         */
        public Builder homeDashboardId(Integer homeDashboardId) {
            return homeDashboardId(Output.of(homeDashboardId));
        }

        /**
         * @param homeDashboardUid The Organization home dashboard UID.
         * 
         * @return builder
         * 
         */
        public Builder homeDashboardUid(@Nullable Output<String> homeDashboardUid) {
            $.homeDashboardUid = homeDashboardUid;
            return this;
        }

        /**
         * @param homeDashboardUid The Organization home dashboard UID.
         * 
         * @return builder
         * 
         */
        public Builder homeDashboardUid(String homeDashboardUid) {
            return homeDashboardUid(Output.of(homeDashboardUid));
        }

        /**
         * @param theme The Organization theme. Available values are `light`, `dark`, or an empty string for the default.
         * 
         * @return builder
         * 
         */
        public Builder theme(@Nullable Output<String> theme) {
            $.theme = theme;
            return this;
        }

        /**
         * @param theme The Organization theme. Available values are `light`, `dark`, or an empty string for the default.
         * 
         * @return builder
         * 
         */
        public Builder theme(String theme) {
            return theme(Output.of(theme));
        }

        /**
         * @param timezone The Organization timezone. Available values are `utc`, `browser`, or an empty string for the default.
         * 
         * @return builder
         * 
         */
        public Builder timezone(@Nullable Output<String> timezone) {
            $.timezone = timezone;
            return this;
        }

        /**
         * @param timezone The Organization timezone. Available values are `utc`, `browser`, or an empty string for the default.
         * 
         * @return builder
         * 
         */
        public Builder timezone(String timezone) {
            return timezone(Output.of(timezone));
        }

        /**
         * @param weekStart The Organization week start.
         * 
         * @return builder
         * 
         */
        public Builder weekStart(@Nullable Output<String> weekStart) {
            $.weekStart = weekStart;
            return this;
        }

        /**
         * @param weekStart The Organization week start.
         * 
         * @return builder
         * 
         */
        public Builder weekStart(String weekStart) {
            return weekStart(Output.of(weekStart));
        }

        public OrganizationPreferenceState build() {
            return $;
        }
    }

}
