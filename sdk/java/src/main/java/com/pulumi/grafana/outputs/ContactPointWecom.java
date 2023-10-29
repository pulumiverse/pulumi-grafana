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
public final class ContactPointWecom {
    /**
     * @return Agent ID added to the request payload when using APIAPP.
     * 
     */
    private @Nullable String agentId;
    /**
     * @return Corp ID used to get token when using APIAPP.
     * 
     */
    private @Nullable String corpId;
    /**
     * @return Whether to disable sending resolve messages. Defaults to `false`.
     * 
     */
    private @Nullable Boolean disableResolveMessage;
    /**
     * @return The templated content of the message to send.
     * 
     */
    private @Nullable String message;
    /**
     * @return The type of them message. Supported: markdown, text. Default: text.
     * 
     */
    private @Nullable String msgType;
    /**
     * @return The secret key required to obtain access token when using APIAPP. See https://work.weixin.qq.com/wework_admin/frame#apps to create APIAPP.
     * 
     */
    private @Nullable String secret;
    /**
     * @return Additional custom properties to attach to the notifier. Defaults to `map[]`.
     * 
     */
    private @Nullable Map<String,String> settings;
    /**
     * @return The templated title of the message to send.
     * 
     */
    private @Nullable String title;
    /**
     * @return The ID of user that should receive the message. Multiple entries should be separated by &#39;|&#39;. Default: @all.
     * 
     */
    private @Nullable String toUser;
    /**
     * @return The UID of the contact point.
     * 
     */
    private @Nullable String uid;
    /**
     * @return The WeCom webhook URL. Required if using GroupRobot.
     * 
     */
    private @Nullable String url;

    private ContactPointWecom() {}
    /**
     * @return Agent ID added to the request payload when using APIAPP.
     * 
     */
    public Optional<String> agentId() {
        return Optional.ofNullable(this.agentId);
    }
    /**
     * @return Corp ID used to get token when using APIAPP.
     * 
     */
    public Optional<String> corpId() {
        return Optional.ofNullable(this.corpId);
    }
    /**
     * @return Whether to disable sending resolve messages. Defaults to `false`.
     * 
     */
    public Optional<Boolean> disableResolveMessage() {
        return Optional.ofNullable(this.disableResolveMessage);
    }
    /**
     * @return The templated content of the message to send.
     * 
     */
    public Optional<String> message() {
        return Optional.ofNullable(this.message);
    }
    /**
     * @return The type of them message. Supported: markdown, text. Default: text.
     * 
     */
    public Optional<String> msgType() {
        return Optional.ofNullable(this.msgType);
    }
    /**
     * @return The secret key required to obtain access token when using APIAPP. See https://work.weixin.qq.com/wework_admin/frame#apps to create APIAPP.
     * 
     */
    public Optional<String> secret() {
        return Optional.ofNullable(this.secret);
    }
    /**
     * @return Additional custom properties to attach to the notifier. Defaults to `map[]`.
     * 
     */
    public Map<String,String> settings() {
        return this.settings == null ? Map.of() : this.settings;
    }
    /**
     * @return The templated title of the message to send.
     * 
     */
    public Optional<String> title() {
        return Optional.ofNullable(this.title);
    }
    /**
     * @return The ID of user that should receive the message. Multiple entries should be separated by &#39;|&#39;. Default: @all.
     * 
     */
    public Optional<String> toUser() {
        return Optional.ofNullable(this.toUser);
    }
    /**
     * @return The UID of the contact point.
     * 
     */
    public Optional<String> uid() {
        return Optional.ofNullable(this.uid);
    }
    /**
     * @return The WeCom webhook URL. Required if using GroupRobot.
     * 
     */
    public Optional<String> url() {
        return Optional.ofNullable(this.url);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContactPointWecom defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String agentId;
        private @Nullable String corpId;
        private @Nullable Boolean disableResolveMessage;
        private @Nullable String message;
        private @Nullable String msgType;
        private @Nullable String secret;
        private @Nullable Map<String,String> settings;
        private @Nullable String title;
        private @Nullable String toUser;
        private @Nullable String uid;
        private @Nullable String url;
        public Builder() {}
        public Builder(ContactPointWecom defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.agentId = defaults.agentId;
    	      this.corpId = defaults.corpId;
    	      this.disableResolveMessage = defaults.disableResolveMessage;
    	      this.message = defaults.message;
    	      this.msgType = defaults.msgType;
    	      this.secret = defaults.secret;
    	      this.settings = defaults.settings;
    	      this.title = defaults.title;
    	      this.toUser = defaults.toUser;
    	      this.uid = defaults.uid;
    	      this.url = defaults.url;
        }

        @CustomType.Setter
        public Builder agentId(@Nullable String agentId) {
            this.agentId = agentId;
            return this;
        }
        @CustomType.Setter
        public Builder corpId(@Nullable String corpId) {
            this.corpId = corpId;
            return this;
        }
        @CustomType.Setter
        public Builder disableResolveMessage(@Nullable Boolean disableResolveMessage) {
            this.disableResolveMessage = disableResolveMessage;
            return this;
        }
        @CustomType.Setter
        public Builder message(@Nullable String message) {
            this.message = message;
            return this;
        }
        @CustomType.Setter
        public Builder msgType(@Nullable String msgType) {
            this.msgType = msgType;
            return this;
        }
        @CustomType.Setter
        public Builder secret(@Nullable String secret) {
            this.secret = secret;
            return this;
        }
        @CustomType.Setter
        public Builder settings(@Nullable Map<String,String> settings) {
            this.settings = settings;
            return this;
        }
        @CustomType.Setter
        public Builder title(@Nullable String title) {
            this.title = title;
            return this;
        }
        @CustomType.Setter
        public Builder toUser(@Nullable String toUser) {
            this.toUser = toUser;
            return this;
        }
        @CustomType.Setter
        public Builder uid(@Nullable String uid) {
            this.uid = uid;
            return this;
        }
        @CustomType.Setter
        public Builder url(@Nullable String url) {
            this.url = url;
            return this;
        }
        public ContactPointWecom build() {
            final var o = new ContactPointWecom();
            o.agentId = agentId;
            o.corpId = corpId;
            o.disableResolveMessage = disableResolveMessage;
            o.message = message;
            o.msgType = msgType;
            o.secret = secret;
            o.settings = settings;
            o.title = title;
            o.toUser = toUser;
            o.uid = uid;
            o.url = url;
            return o;
        }
    }
}
