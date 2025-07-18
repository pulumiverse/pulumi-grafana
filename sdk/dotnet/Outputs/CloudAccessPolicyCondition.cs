// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Outputs
{

    [OutputType]
    public sealed class CloudAccessPolicyCondition
    {
        /// <summary>
        /// Conditions that apply to the access policy,such as IP Allow lists.
        /// </summary>
        public readonly ImmutableArray<string> AllowedSubnets;

        [OutputConstructor]
        private CloudAccessPolicyCondition(ImmutableArray<string> allowedSubnets)
        {
            AllowedSubnets = allowedSubnets;
        }
    }
}
