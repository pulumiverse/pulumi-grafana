// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Inputs
{

    public sealed class SyntheticMonitoringCheckSettingsBrowserArgs : global::Pulumi.ResourceArgs
    {
        [Input("script", required: true)]
        public Input<string> Script { get; set; } = null!;

        public SyntheticMonitoringCheckSettingsBrowserArgs()
        {
        }
        public static new SyntheticMonitoringCheckSettingsBrowserArgs Empty => new SyntheticMonitoringCheckSettingsBrowserArgs();
    }
}
