// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana
{
    /// <summary>
    /// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/escalation_chains/)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Grafana = Pulumi.Grafana;
    /// using Grafana = Pulumiverse.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myTeam = Grafana.Oss.GetTeam.Invoke(new()
    ///     {
    ///         Name = "my team",
    ///     });
    /// 
    ///     var myTeamGetTeam = Grafana.OnCall.GetTeam.Invoke(new()
    ///     {
    ///         Name = myTeam.Apply(getTeamResult =&gt; getTeamResult.Name),
    ///     });
    /// 
    ///     var @default = new Grafana.OnCall.EscalationChain("default", new()
    ///     {
    ///         Name = "default",
    ///         TeamId = myTeamGetTeam.Apply(getTeamResult =&gt; getTeamResult.Id),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:index/oncallEscalationChain:OncallEscalationChain name "{{ id }}"
    /// ```
    /// </summary>
    [Obsolete(@"grafana.index/oncallescalationchain.OncallEscalationChain has been deprecated in favor of grafana.oncall/escalationchain.EscalationChain")]
    [GrafanaResourceType("grafana:index/oncallEscalationChain:OncallEscalationChain")]
    public partial class OncallEscalationChain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the escalation chain.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the OnCall team (using the `grafana.onCall.getTeam` datasource).
        /// </summary>
        [Output("teamId")]
        public Output<string?> TeamId { get; private set; } = null!;


        /// <summary>
        /// Create a OncallEscalationChain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OncallEscalationChain(string name, OncallEscalationChainArgs? args = null, CustomResourceOptions? options = null)
            : base("grafana:index/oncallEscalationChain:OncallEscalationChain", name, args ?? new OncallEscalationChainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OncallEscalationChain(string name, Input<string> id, OncallEscalationChainState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/oncallEscalationChain:OncallEscalationChain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OncallEscalationChain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OncallEscalationChain Get(string name, Input<string> id, OncallEscalationChainState? state = null, CustomResourceOptions? options = null)
        {
            return new OncallEscalationChain(name, id, state, options);
        }
    }

    public sealed class OncallEscalationChainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the escalation chain.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the OnCall team (using the `grafana.onCall.getTeam` datasource).
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public OncallEscalationChainArgs()
        {
        }
        public static new OncallEscalationChainArgs Empty => new OncallEscalationChainArgs();
    }

    public sealed class OncallEscalationChainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the escalation chain.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the OnCall team (using the `grafana.onCall.getTeam` datasource).
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public OncallEscalationChainState()
        {
        }
        public static new OncallEscalationChainState Empty => new OncallEscalationChainState();
    }
}
