// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Alerting.Inputs
{

    public sealed class MuteTimingIntervalTimeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time, in hh:mm format, of when the interval should end exclusively.
        /// </summary>
        [Input("end", required: true)]
        public Input<string> End { get; set; } = null!;

        /// <summary>
        /// The time, in hh:mm format, of when the interval should begin inclusively.
        /// </summary>
        [Input("start", required: true)]
        public Input<string> Start { get; set; } = null!;

        public MuteTimingIntervalTimeArgs()
        {
        }
        public static new MuteTimingIntervalTimeArgs Empty => new MuteTimingIntervalTimeArgs();
    }
}
