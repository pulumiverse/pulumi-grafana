// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Inputs
{

    public sealed class CloudAccessPolicyRealmArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identifier of the org or stack. For orgs, this is the slug, for stacks, this is the stack ID.
        /// </summary>
        [Input("identifier", required: true)]
        public Input<string> Identifier { get; set; } = null!;

        [Input("labelPolicies")]
        private InputList<Inputs.CloudAccessPolicyRealmLabelPolicyArgs>? _labelPolicies;
        public InputList<Inputs.CloudAccessPolicyRealmLabelPolicyArgs> LabelPolicies
        {
            get => _labelPolicies ?? (_labelPolicies = new InputList<Inputs.CloudAccessPolicyRealmLabelPolicyArgs>());
            set => _labelPolicies = value;
        }

        /// <summary>
        /// Whether a policy applies to a Cloud org or a specific stack. Should be one of `org` or `stack`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public CloudAccessPolicyRealmArgs()
        {
        }
        public static new CloudAccessPolicyRealmArgs Empty => new CloudAccessPolicyRealmArgs();
    }
}
