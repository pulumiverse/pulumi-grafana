// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Slo.Outputs
{

    [OutputType]
    public sealed class SLOAlerting
    {
        /// <summary>
        /// Advanced Options for Alert Rules
        /// </summary>
        public readonly Outputs.SLOAlertingAdvancedOptions? AdvancedOptions;
        /// <summary>
        /// Annotations will be attached to all alerts generated by any of these rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.SLOAlertingAnnotation> Annotations;
        /// <summary>
        /// Alerting Rules generated for Fast Burn alerts
        /// </summary>
        public readonly ImmutableArray<Outputs.SLOAlertingFastburn> Fastburns;
        /// <summary>
        /// Labels will be attached to all alerts generated by any of these rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.SLOAlertingLabel> Labels;
        /// <summary>
        /// Alerting Rules generated for Slow Burn alerts
        /// </summary>
        public readonly ImmutableArray<Outputs.SLOAlertingSlowburn> Slowburns;

        [OutputConstructor]
        private SLOAlerting(
            Outputs.SLOAlertingAdvancedOptions? advancedOptions,

            ImmutableArray<Outputs.SLOAlertingAnnotation> annotations,

            ImmutableArray<Outputs.SLOAlertingFastburn> fastburns,

            ImmutableArray<Outputs.SLOAlertingLabel> labels,

            ImmutableArray<Outputs.SLOAlertingSlowburn> slowburns)
        {
            AdvancedOptions = advancedOptions;
            Annotations = annotations;
            Fastburns = fastburns;
            Labels = labels;
            Slowburns = slowburns;
        }
    }
}