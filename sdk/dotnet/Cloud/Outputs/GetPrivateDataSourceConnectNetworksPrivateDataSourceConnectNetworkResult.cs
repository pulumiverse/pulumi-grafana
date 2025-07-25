// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Cloud.Outputs
{

    [OutputType]
    public sealed class GetPrivateDataSourceConnectNetworksPrivateDataSourceConnectNetworkResult
    {
        public readonly string DisplayName;
        public readonly string Id;
        public readonly string Name;
        public readonly string Region;
        public readonly string Status;

        [OutputConstructor]
        private GetPrivateDataSourceConnectNetworksPrivateDataSourceConnectNetworkResult(
            string displayName,

            string id,

            string name,

            string region,

            string status)
        {
            DisplayName = displayName;
            Id = id;
            Name = name;
            Region = region;
            Status = status;
        }
    }
}
