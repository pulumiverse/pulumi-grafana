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

    public sealed class OncallIntegrationTemplatesPhoneCallArgs : global::Pulumi.ResourceArgs
    {
        [Input("title")]
        public Input<string>? Title { get; set; }

        public OncallIntegrationTemplatesPhoneCallArgs()
        {
        }
        public static new OncallIntegrationTemplatesPhoneCallArgs Empty => new OncallIntegrationTemplatesPhoneCallArgs();
    }
}