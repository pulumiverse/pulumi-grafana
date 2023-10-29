// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrs {
    /**
     * @return Fail if value matches regex.
     * 
     */
    private @Nullable List<String> failIfMatchesRegexps;
    /**
     * @return Fail if value does not match regex.
     * 
     */
    private @Nullable List<String> failIfNotMatchesRegexps;

    private SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrs() {}
    /**
     * @return Fail if value matches regex.
     * 
     */
    public List<String> failIfMatchesRegexps() {
        return this.failIfMatchesRegexps == null ? List.of() : this.failIfMatchesRegexps;
    }
    /**
     * @return Fail if value does not match regex.
     * 
     */
    public List<String> failIfNotMatchesRegexps() {
        return this.failIfNotMatchesRegexps == null ? List.of() : this.failIfNotMatchesRegexps;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrs defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> failIfMatchesRegexps;
        private @Nullable List<String> failIfNotMatchesRegexps;
        public Builder() {}
        public Builder(SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrs defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.failIfMatchesRegexps = defaults.failIfMatchesRegexps;
    	      this.failIfNotMatchesRegexps = defaults.failIfNotMatchesRegexps;
        }

        @CustomType.Setter
        public Builder failIfMatchesRegexps(@Nullable List<String> failIfMatchesRegexps) {
            this.failIfMatchesRegexps = failIfMatchesRegexps;
            return this;
        }
        public Builder failIfMatchesRegexps(String... failIfMatchesRegexps) {
            return failIfMatchesRegexps(List.of(failIfMatchesRegexps));
        }
        @CustomType.Setter
        public Builder failIfNotMatchesRegexps(@Nullable List<String> failIfNotMatchesRegexps) {
            this.failIfNotMatchesRegexps = failIfNotMatchesRegexps;
            return this;
        }
        public Builder failIfNotMatchesRegexps(String... failIfNotMatchesRegexps) {
            return failIfNotMatchesRegexps(List.of(failIfNotMatchesRegexps));
        }
        public SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrs build() {
            final var o = new SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrs();
            o.failIfMatchesRegexps = failIfMatchesRegexps;
            o.failIfNotMatchesRegexps = failIfNotMatchesRegexps;
            return o;
        }
    }
}
