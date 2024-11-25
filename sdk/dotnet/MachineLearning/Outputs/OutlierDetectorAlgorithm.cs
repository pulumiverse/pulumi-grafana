// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.MachineLearning.Outputs
{

    [OutputType]
    public sealed class OutlierDetectorAlgorithm
    {
        /// <summary>
        /// For DBSCAN only, specify the configuration map
        /// </summary>
        public readonly Outputs.OutlierDetectorAlgorithmConfig? Config;
        /// <summary>
        /// The name of the algorithm to use ('mad' or 'dbscan').
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specify the sensitivity of the detector (in range [0,1]).
        /// </summary>
        public readonly double Sensitivity;

        [OutputConstructor]
        private OutlierDetectorAlgorithm(
            Outputs.OutlierDetectorAlgorithmConfig? config,

            string name,

            double sensitivity)
        {
            Config = config;
            Name = name;
            Sensitivity = sensitivity;
        }
    }
}