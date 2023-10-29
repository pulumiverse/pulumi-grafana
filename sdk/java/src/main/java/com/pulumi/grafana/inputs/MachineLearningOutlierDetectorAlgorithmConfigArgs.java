// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.util.Objects;


public final class MachineLearningOutlierDetectorAlgorithmConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final MachineLearningOutlierDetectorAlgorithmConfigArgs Empty = new MachineLearningOutlierDetectorAlgorithmConfigArgs();

    /**
     * Specify the epsilon parameter (positive float)
     * 
     */
    @Import(name="epsilon", required=true)
    private Output<Double> epsilon;

    /**
     * @return Specify the epsilon parameter (positive float)
     * 
     */
    public Output<Double> epsilon() {
        return this.epsilon;
    }

    private MachineLearningOutlierDetectorAlgorithmConfigArgs() {}

    private MachineLearningOutlierDetectorAlgorithmConfigArgs(MachineLearningOutlierDetectorAlgorithmConfigArgs $) {
        this.epsilon = $.epsilon;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MachineLearningOutlierDetectorAlgorithmConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MachineLearningOutlierDetectorAlgorithmConfigArgs $;

        public Builder() {
            $ = new MachineLearningOutlierDetectorAlgorithmConfigArgs();
        }

        public Builder(MachineLearningOutlierDetectorAlgorithmConfigArgs defaults) {
            $ = new MachineLearningOutlierDetectorAlgorithmConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param epsilon Specify the epsilon parameter (positive float)
         * 
         * @return builder
         * 
         */
        public Builder epsilon(Output<Double> epsilon) {
            $.epsilon = epsilon;
            return this;
        }

        /**
         * @param epsilon Specify the epsilon parameter (positive float)
         * 
         * @return builder
         * 
         */
        public Builder epsilon(Double epsilon) {
            return epsilon(Output.of(epsilon));
        }

        public MachineLearningOutlierDetectorAlgorithmConfigArgs build() {
            $.epsilon = Objects.requireNonNull($.epsilon, "expected parameter 'epsilon' to be non-null");
            return $;
        }
    }

}
