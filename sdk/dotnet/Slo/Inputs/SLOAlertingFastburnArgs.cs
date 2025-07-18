// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Slo.Inputs
{

    public sealed class SLOAlertingFastburnArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputList<Inputs.SLOAlertingFastburnAnnotationArgs>? _annotations;

        /// <summary>
        /// Annotations to attach only to Fast Burn alerts.
        /// </summary>
        public InputList<Inputs.SLOAlertingFastburnAnnotationArgs> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<Inputs.SLOAlertingFastburnAnnotationArgs>());
            set => _annotations = value;
        }

        [Input("labels")]
        private InputList<Inputs.SLOAlertingFastburnLabelArgs>? _labels;

        /// <summary>
        /// Labels to attach only to Fast Burn alerts.
        /// </summary>
        public InputList<Inputs.SLOAlertingFastburnLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.SLOAlertingFastburnLabelArgs>());
            set => _labels = value;
        }

        public SLOAlertingFastburnArgs()
        {
        }
        public static new SLOAlertingFastburnArgs Empty => new SLOAlertingFastburnArgs();
    }
}
