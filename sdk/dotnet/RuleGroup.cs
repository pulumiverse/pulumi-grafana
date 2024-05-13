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
    [GrafanaResourceType("grafana:index/ruleGroup:RuleGroup")]
    public partial class RuleGroup : global::Pulumi.CustomResource
    {
        [Output("disableProvenance")]
        public Output<bool?> DisableProvenance { get; private set; } = null!;

        /// <summary>
        /// The UID of the folder that the group belongs to.
        /// </summary>
        [Output("folderUid")]
        public Output<string> FolderUid { get; private set; } = null!;

        /// <summary>
        /// The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are
        /// evaluated sequentially.
        /// </summary>
        [Output("intervalSeconds")]
        public Output<int> IntervalSeconds { get; private set; } = null!;

        /// <summary>
        /// The name of the rule group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        /// <summary>
        /// The rules within the group.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.RuleGroupRule>> Rules { get; private set; } = null!;


        /// <summary>
        /// Create a RuleGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RuleGroup(string name, RuleGroupArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/ruleGroup:RuleGroup", name, args ?? new RuleGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RuleGroup(string name, Input<string> id, RuleGroupState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/ruleGroup:RuleGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RuleGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RuleGroup Get(string name, Input<string> id, RuleGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new RuleGroup(name, id, state, options);
        }
    }

    public sealed class RuleGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("disableProvenance")]
        public Input<bool>? DisableProvenance { get; set; }

        /// <summary>
        /// The UID of the folder that the group belongs to.
        /// </summary>
        [Input("folderUid", required: true)]
        public Input<string> FolderUid { get; set; } = null!;

        /// <summary>
        /// The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are
        /// evaluated sequentially.
        /// </summary>
        [Input("intervalSeconds", required: true)]
        public Input<int> IntervalSeconds { get; set; } = null!;

        /// <summary>
        /// The name of the rule group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("rules", required: true)]
        private InputList<Inputs.RuleGroupRuleArgs>? _rules;

        /// <summary>
        /// The rules within the group.
        /// </summary>
        public InputList<Inputs.RuleGroupRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.RuleGroupRuleArgs>());
            set => _rules = value;
        }

        public RuleGroupArgs()
        {
        }
        public static new RuleGroupArgs Empty => new RuleGroupArgs();
    }

    public sealed class RuleGroupState : global::Pulumi.ResourceArgs
    {
        [Input("disableProvenance")]
        public Input<bool>? DisableProvenance { get; set; }

        /// <summary>
        /// The UID of the folder that the group belongs to.
        /// </summary>
        [Input("folderUid")]
        public Input<string>? FolderUid { get; set; }

        /// <summary>
        /// The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are
        /// evaluated sequentially.
        /// </summary>
        [Input("intervalSeconds")]
        public Input<int>? IntervalSeconds { get; set; }

        /// <summary>
        /// The name of the rule group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("rules")]
        private InputList<Inputs.RuleGroupRuleGetArgs>? _rules;

        /// <summary>
        /// The rules within the group.
        /// </summary>
        public InputList<Inputs.RuleGroupRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.RuleGroupRuleGetArgs>());
            set => _rules = value;
        }

        public RuleGroupState()
        {
        }
        public static new RuleGroupState Empty => new RuleGroupState();
    }
}
