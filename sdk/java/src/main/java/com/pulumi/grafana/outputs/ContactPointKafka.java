// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ContactPointKafka {
    /**
     * @return The API version to use when contacting the Kafka REST Server. Supported: v2 (default) and v3. Defaults to `v2`.
     * 
     */
    private @Nullable String apiVersion;
    /**
     * @return The Id of cluster to use when contacting the Kafka REST Server. Required api_version to be &#39;v3&#39;
     * 
     */
    private @Nullable String clusterId;
    /**
     * @return The templated description of the Kafka message.
     * 
     */
    private @Nullable String description;
    /**
     * @return The templated details to include with the message.
     * 
     */
    private @Nullable String details;
    /**
     * @return Whether to disable sending resolve messages. Defaults to `false`.
     * 
     */
    private @Nullable Boolean disableResolveMessage;
    /**
     * @return The password to use when making a call to the Kafka REST Proxy
     * 
     */
    private @Nullable String password;
    /**
     * @return The URL of the Kafka REST proxy to send requests to.
     * 
     */
    private String restProxyUrl;
    /**
     * @return Additional custom properties to attach to the notifier. Defaults to `map[]`.
     * 
     */
    private @Nullable Map<String,String> settings;
    /**
     * @return The name of the Kafka topic to publish to.
     * 
     */
    private String topic;
    /**
     * @return The UID of the contact point.
     * 
     */
    private @Nullable String uid;
    /**
     * @return The user name to use when making a call to the Kafka REST Proxy
     * 
     */
    private @Nullable String username;

    private ContactPointKafka() {}
    /**
     * @return The API version to use when contacting the Kafka REST Server. Supported: v2 (default) and v3. Defaults to `v2`.
     * 
     */
    public Optional<String> apiVersion() {
        return Optional.ofNullable(this.apiVersion);
    }
    /**
     * @return The Id of cluster to use when contacting the Kafka REST Server. Required api_version to be &#39;v3&#39;
     * 
     */
    public Optional<String> clusterId() {
        return Optional.ofNullable(this.clusterId);
    }
    /**
     * @return The templated description of the Kafka message.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return The templated details to include with the message.
     * 
     */
    public Optional<String> details() {
        return Optional.ofNullable(this.details);
    }
    /**
     * @return Whether to disable sending resolve messages. Defaults to `false`.
     * 
     */
    public Optional<Boolean> disableResolveMessage() {
        return Optional.ofNullable(this.disableResolveMessage);
    }
    /**
     * @return The password to use when making a call to the Kafka REST Proxy
     * 
     */
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }
    /**
     * @return The URL of the Kafka REST proxy to send requests to.
     * 
     */
    public String restProxyUrl() {
        return this.restProxyUrl;
    }
    /**
     * @return Additional custom properties to attach to the notifier. Defaults to `map[]`.
     * 
     */
    public Map<String,String> settings() {
        return this.settings == null ? Map.of() : this.settings;
    }
    /**
     * @return The name of the Kafka topic to publish to.
     * 
     */
    public String topic() {
        return this.topic;
    }
    /**
     * @return The UID of the contact point.
     * 
     */
    public Optional<String> uid() {
        return Optional.ofNullable(this.uid);
    }
    /**
     * @return The user name to use when making a call to the Kafka REST Proxy
     * 
     */
    public Optional<String> username() {
        return Optional.ofNullable(this.username);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContactPointKafka defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String apiVersion;
        private @Nullable String clusterId;
        private @Nullable String description;
        private @Nullable String details;
        private @Nullable Boolean disableResolveMessage;
        private @Nullable String password;
        private String restProxyUrl;
        private @Nullable Map<String,String> settings;
        private String topic;
        private @Nullable String uid;
        private @Nullable String username;
        public Builder() {}
        public Builder(ContactPointKafka defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.apiVersion = defaults.apiVersion;
    	      this.clusterId = defaults.clusterId;
    	      this.description = defaults.description;
    	      this.details = defaults.details;
    	      this.disableResolveMessage = defaults.disableResolveMessage;
    	      this.password = defaults.password;
    	      this.restProxyUrl = defaults.restProxyUrl;
    	      this.settings = defaults.settings;
    	      this.topic = defaults.topic;
    	      this.uid = defaults.uid;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder apiVersion(@Nullable String apiVersion) {
            this.apiVersion = apiVersion;
            return this;
        }
        @CustomType.Setter
        public Builder clusterId(@Nullable String clusterId) {
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder details(@Nullable String details) {
            this.details = details;
            return this;
        }
        @CustomType.Setter
        public Builder disableResolveMessage(@Nullable Boolean disableResolveMessage) {
            this.disableResolveMessage = disableResolveMessage;
            return this;
        }
        @CustomType.Setter
        public Builder password(@Nullable String password) {
            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder restProxyUrl(String restProxyUrl) {
            this.restProxyUrl = Objects.requireNonNull(restProxyUrl);
            return this;
        }
        @CustomType.Setter
        public Builder settings(@Nullable Map<String,String> settings) {
            this.settings = settings;
            return this;
        }
        @CustomType.Setter
        public Builder topic(String topic) {
            this.topic = Objects.requireNonNull(topic);
            return this;
        }
        @CustomType.Setter
        public Builder uid(@Nullable String uid) {
            this.uid = uid;
            return this;
        }
        @CustomType.Setter
        public Builder username(@Nullable String username) {
            this.username = username;
            return this;
        }
        public ContactPointKafka build() {
            final var o = new ContactPointKafka();
            o.apiVersion = apiVersion;
            o.clusterId = clusterId;
            o.description = description;
            o.details = details;
            o.disableResolveMessage = disableResolveMessage;
            o.password = password;
            o.restProxyUrl = restProxyUrl;
            o.settings = settings;
            o.topic = topic;
            o.uid = uid;
            o.username = username;
            return o;
        }
    }
}
