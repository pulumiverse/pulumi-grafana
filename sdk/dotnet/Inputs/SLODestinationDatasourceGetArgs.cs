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

    public sealed class SLODestinationDatasourceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// UID for the Datasource
        /// </summary>
        [Input("uid", required: true)]
        public Input<string> Uid { get; set; } = null!;

        public SLODestinationDatasourceGetArgs()
        {
        }
        public static new SLODestinationDatasourceGetArgs Empty => new SLODestinationDatasourceGetArgs();
    }
}
