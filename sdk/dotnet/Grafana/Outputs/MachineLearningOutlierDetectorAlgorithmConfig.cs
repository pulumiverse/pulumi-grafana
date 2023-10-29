// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Grafana.Outputs
{

    [OutputType]
    public sealed class MachineLearningOutlierDetectorAlgorithmConfig
    {
        /// <summary>
        /// Specify the epsilon parameter (positive float)
        /// </summary>
        public readonly double Epsilon;

        [OutputConstructor]
        private MachineLearningOutlierDetectorAlgorithmConfig(double epsilon)
        {
            Epsilon = epsilon;
        }
    }
}
