// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class MuteTimingIntervalTime {
    /**
     * @return The time, in hh:mm format, of when the interval should end exclusively.
     * 
     */
    private String end;
    /**
     * @return The time, in hh:mm format, of when the interval should begin inclusively.
     * 
     */
    private String start;

    private MuteTimingIntervalTime() {}
    /**
     * @return The time, in hh:mm format, of when the interval should end exclusively.
     * 
     */
    public String end() {
        return this.end;
    }
    /**
     * @return The time, in hh:mm format, of when the interval should begin inclusively.
     * 
     */
    public String start() {
        return this.start;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(MuteTimingIntervalTime defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String end;
        private String start;
        public Builder() {}
        public Builder(MuteTimingIntervalTime defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.end = defaults.end;
    	      this.start = defaults.start;
        }

        @CustomType.Setter
        public Builder end(String end) {
            this.end = Objects.requireNonNull(end);
            return this;
        }
        @CustomType.Setter
        public Builder start(String start) {
            this.start = Objects.requireNonNull(start);
            return this;
        }
        public MuteTimingIntervalTime build() {
            final var o = new MuteTimingIntervalTime();
            o.end = end;
            o.start = start;
            return o;
        }
    }
}
