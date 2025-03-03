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

    public sealed class SLOAlertingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Advanced Options for Alert Rules
        /// </summary>
        [Input("advancedOptions")]
        public Input<Inputs.SLOAlertingAdvancedOptionsArgs>? AdvancedOptions { get; set; }

        [Input("annotations")]
        private InputList<Inputs.SLOAlertingAnnotationArgs>? _annotations;

        /// <summary>
        /// Annotations will be attached to all alerts generated by any of these rules.
        /// </summary>
        public InputList<Inputs.SLOAlertingAnnotationArgs> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<Inputs.SLOAlertingAnnotationArgs>());
            set => _annotations = value;
        }

        /// <summary>
        /// Alerting Rules generated for Fast Burn alerts
        /// </summary>
        [Input("fastburn")]
        public Input<Inputs.SLOAlertingFastburnArgs>? Fastburn { get; set; }

        [Input("labels")]
        private InputList<Inputs.SLOAlertingLabelArgs>? _labels;

        /// <summary>
        /// Labels will be attached to all alerts generated by any of these rules.
        /// </summary>
        public InputList<Inputs.SLOAlertingLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.SLOAlertingLabelArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// Alerting Rules generated for Slow Burn alerts
        /// </summary>
        [Input("slowburn")]
        public Input<Inputs.SLOAlertingSlowburnArgs>? Slowburn { get; set; }

        public SLOAlertingArgs()
        {
        }
        public static new SLOAlertingArgs Empty => new SLOAlertingArgs();
    }
}
