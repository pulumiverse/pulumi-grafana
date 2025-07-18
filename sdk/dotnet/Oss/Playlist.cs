// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Oss
{
    /// <summary>
    /// * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/create-manage-playlists/)
    /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/playlist/)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Grafana = Pulumiverse.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new Grafana.Oss.Playlist("test", new()
    ///     {
    ///         Name = "My Playlist!",
    ///         Interval = "5m",
    ///         Items = new[]
    ///         {
    ///             new Grafana.Oss.Inputs.PlaylistItemArgs
    ///             {
    ///                 Order = 2,
    ///                 Title = "Terraform Dashboard By Tag",
    ///                 Type = "dashboard_by_tag",
    ///                 Value = "terraform",
    ///             },
    ///             new Grafana.Oss.Inputs.PlaylistItemArgs
    ///             {
    ///                 Order = 1,
    ///                 Title = "Terraform Dashboard By UID",
    ///                 Type = "dashboard_by_uid",
    ///                 Value = "cIBgcSjkk",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:oss/playlist:Playlist name "{{ uid }}"
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import grafana:oss/playlist:Playlist name "{{ orgID }}:{{ uid }}"
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:oss/playlist:Playlist")]
    public partial class Playlist : global::Pulumi.CustomResource
    {
        [Output("interval")]
        public Output<string> Interval { get; private set; } = null!;

        [Output("items")]
        public Output<ImmutableArray<Outputs.PlaylistItem>> Items { get; private set; } = null!;

        /// <summary>
        /// The name of the playlist.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;


        /// <summary>
        /// Create a Playlist resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Playlist(string name, PlaylistArgs args, CustomResourceOptions? options = null)
            : base("grafana:oss/playlist:Playlist", name, args ?? new PlaylistArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Playlist(string name, Input<string> id, PlaylistState? state = null, CustomResourceOptions? options = null)
            : base("grafana:oss/playlist:Playlist", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "grafana:index/playlist:Playlist" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Playlist resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Playlist Get(string name, Input<string> id, PlaylistState? state = null, CustomResourceOptions? options = null)
        {
            return new Playlist(name, id, state, options);
        }
    }

    public sealed class PlaylistArgs : global::Pulumi.ResourceArgs
    {
        [Input("interval", required: true)]
        public Input<string> Interval { get; set; } = null!;

        [Input("items", required: true)]
        private InputList<Inputs.PlaylistItemArgs>? _items;
        public InputList<Inputs.PlaylistItemArgs> Items
        {
            get => _items ?? (_items = new InputList<Inputs.PlaylistItemArgs>());
            set => _items = value;
        }

        /// <summary>
        /// The name of the playlist.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        public PlaylistArgs()
        {
        }
        public static new PlaylistArgs Empty => new PlaylistArgs();
    }

    public sealed class PlaylistState : global::Pulumi.ResourceArgs
    {
        [Input("interval")]
        public Input<string>? Interval { get; set; }

        [Input("items")]
        private InputList<Inputs.PlaylistItemGetArgs>? _items;
        public InputList<Inputs.PlaylistItemGetArgs> Items
        {
            get => _items ?? (_items = new InputList<Inputs.PlaylistItemGetArgs>());
            set => _items = value;
        }

        /// <summary>
        /// The name of the playlist.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        public PlaylistState()
        {
        }
        public static new PlaylistState Empty => new PlaylistState();
    }
}
