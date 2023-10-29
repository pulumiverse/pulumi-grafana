// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class OncallIntegrationDefaultRouteSlack {
    /**
     * @return Slack channel id. Alerts will be directed to this channel in Slack.
     * 
     */
    private @Nullable String channelId;
    /**
     * @return Enable notification in MS teams. Defaults to `true`.
     * 
     */
    private @Nullable Boolean enabled;

    private OncallIntegrationDefaultRouteSlack() {}
    /**
     * @return Slack channel id. Alerts will be directed to this channel in Slack.
     * 
     */
    public Optional<String> channelId() {
        return Optional.ofNullable(this.channelId);
    }
    /**
     * @return Enable notification in MS teams. Defaults to `true`.
     * 
     */
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OncallIntegrationDefaultRouteSlack defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String channelId;
        private @Nullable Boolean enabled;
        public Builder() {}
        public Builder(OncallIntegrationDefaultRouteSlack defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.channelId = defaults.channelId;
    	      this.enabled = defaults.enabled;
        }

        @CustomType.Setter
        public Builder channelId(@Nullable String channelId) {
            this.channelId = channelId;
            return this;
        }
        @CustomType.Setter
        public Builder enabled(@Nullable Boolean enabled) {
            this.enabled = enabled;
            return this;
        }
        public OncallIntegrationDefaultRouteSlack build() {
            final var o = new OncallIntegrationDefaultRouteSlack();
            o.channelId = channelId;
            o.enabled = enabled;
            return o;
        }
    }
}
