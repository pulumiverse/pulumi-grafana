// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana
{
    [GrafanaResourceType("grafana:index/cloudOrgMember:CloudOrgMember")]
    public partial class CloudOrgMember : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The slug or ID of the organization.
        /// </summary>
        [Output("org")]
        public Output<string> Org { get; private set; } = null!;

        /// <summary>
        /// Whether the user should receive billing emails.
        /// </summary>
        [Output("receiveBillingEmails")]
        public Output<bool> ReceiveBillingEmails { get; private set; } = null!;

        /// <summary>
        /// The role to assign to the user in the organization.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// Username or ID of the user to add to the org's members.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;


        /// <summary>
        /// Create a CloudOrgMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudOrgMember(string name, CloudOrgMemberArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/cloudOrgMember:CloudOrgMember", name, args ?? new CloudOrgMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudOrgMember(string name, Input<string> id, CloudOrgMemberState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/cloudOrgMember:CloudOrgMember", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CloudOrgMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudOrgMember Get(string name, Input<string> id, CloudOrgMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new CloudOrgMember(name, id, state, options);
        }
    }

    public sealed class CloudOrgMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The slug or ID of the organization.
        /// </summary>
        [Input("org", required: true)]
        public Input<string> Org { get; set; } = null!;

        /// <summary>
        /// Whether the user should receive billing emails.
        /// </summary>
        [Input("receiveBillingEmails")]
        public Input<bool>? ReceiveBillingEmails { get; set; }

        /// <summary>
        /// The role to assign to the user in the organization.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// Username or ID of the user to add to the org's members.
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        public CloudOrgMemberArgs()
        {
        }
        public static new CloudOrgMemberArgs Empty => new CloudOrgMemberArgs();
    }

    public sealed class CloudOrgMemberState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The slug or ID of the organization.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// Whether the user should receive billing emails.
        /// </summary>
        [Input("receiveBillingEmails")]
        public Input<bool>? ReceiveBillingEmails { get; set; }

        /// <summary>
        /// The role to assign to the user in the organization.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Username or ID of the user to add to the org's members.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public CloudOrgMemberState()
        {
        }
        public static new CloudOrgMemberState Empty => new CloudOrgMemberState();
    }
}
