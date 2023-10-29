// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Grafana.Inputs
{

    public sealed class OncallIntegrationTemplatesEmailArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Template for Alert message.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// Template for Alert title.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public OncallIntegrationTemplatesEmailArgs()
        {
        }
        public static new OncallIntegrationTemplatesEmailArgs Empty => new OncallIntegrationTemplatesEmailArgs();
    }
}
