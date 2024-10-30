// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Slo.Inputs
{

    public sealed class SLOAlertingAdvancedOptionsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Minimum number of failed events to trigger an alert
        /// </summary>
        [Input("minFailures")]
        public Input<int>? MinFailures { get; set; }

        public SLOAlertingAdvancedOptionsGetArgs()
        {
        }
        public static new SLOAlertingAdvancedOptionsGetArgs Empty => new SLOAlertingAdvancedOptionsGetArgs();
    }
}
