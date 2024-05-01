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
    [GrafanaResourceType("grafana:index/libraryPanel:LibraryPanel")]
    public partial class LibraryPanel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Timestamp when the library panel was created.
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// Numerical IDs of Grafana dashboards containing the library panel.
        /// </summary>
        [Output("dashboardIds")]
        public Output<ImmutableArray<int>> DashboardIds { get; private set; } = null!;

        /// <summary>
        /// Description of the library panel.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Deprecated. Use `folder_uid` instead
        /// </summary>
        [Output("folderId")]
        public Output<string?> FolderId { get; private set; } = null!;

        /// <summary>
        /// Name of the folder containing the library panel.
        /// </summary>
        [Output("folderName")]
        public Output<string> FolderName { get; private set; } = null!;

        /// <summary>
        /// Unique ID (UID) of the folder containing the library panel.
        /// </summary>
        [Output("folderUid")]
        public Output<string?> FolderUid { get; private set; } = null!;

        /// <summary>
        /// The JSON model for the library panel.
        /// </summary>
        [Output("modelJson")]
        public Output<string> ModelJson { get; private set; } = null!;

        /// <summary>
        /// Name of the library panel.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        /// <summary>
        /// The numeric ID of the library panel computed by Grafana.
        /// </summary>
        [Output("panelId")]
        public Output<int> PanelId { get; private set; } = null!;

        /// <summary>
        /// Type of the library panel (eg. text).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The unique identifier (UID) of a library panel uniquely identifies library panels between multiple Grafana installs.
        /// It’s automatically generated unless you specify it during library panel creation.The UID provides consistent URLs for
        /// accessing library panels and when syncing library panels between multiple Grafana installs.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// Timestamp when the library panel was last modified.
        /// </summary>
        [Output("updated")]
        public Output<string> Updated { get; private set; } = null!;

        /// <summary>
        /// Version of the library panel.
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a LibraryPanel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LibraryPanel(string name, LibraryPanelArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/libraryPanel:LibraryPanel", name, args ?? new LibraryPanelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LibraryPanel(string name, Input<string> id, LibraryPanelState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/libraryPanel:LibraryPanel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LibraryPanel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LibraryPanel Get(string name, Input<string> id, LibraryPanelState? state = null, CustomResourceOptions? options = null)
        {
            return new LibraryPanel(name, id, state, options);
        }
    }

    public sealed class LibraryPanelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Deprecated. Use `folder_uid` instead
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// Unique ID (UID) of the folder containing the library panel.
        /// </summary>
        [Input("folderUid")]
        public Input<string>? FolderUid { get; set; }

        /// <summary>
        /// The JSON model for the library panel.
        /// </summary>
        [Input("modelJson", required: true)]
        public Input<string> ModelJson { get; set; } = null!;

        /// <summary>
        /// Name of the library panel.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// The unique identifier (UID) of a library panel uniquely identifies library panels between multiple Grafana installs.
        /// It’s automatically generated unless you specify it during library panel creation.The UID provides consistent URLs for
        /// accessing library panels and when syncing library panels between multiple Grafana installs.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        public LibraryPanelArgs()
        {
        }
        public static new LibraryPanelArgs Empty => new LibraryPanelArgs();
    }

    public sealed class LibraryPanelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Timestamp when the library panel was created.
        /// </summary>
        [Input("created")]
        public Input<string>? Created { get; set; }

        [Input("dashboardIds")]
        private InputList<int>? _dashboardIds;

        /// <summary>
        /// Numerical IDs of Grafana dashboards containing the library panel.
        /// </summary>
        public InputList<int> DashboardIds
        {
            get => _dashboardIds ?? (_dashboardIds = new InputList<int>());
            set => _dashboardIds = value;
        }

        /// <summary>
        /// Description of the library panel.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Deprecated. Use `folder_uid` instead
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// Name of the folder containing the library panel.
        /// </summary>
        [Input("folderName")]
        public Input<string>? FolderName { get; set; }

        /// <summary>
        /// Unique ID (UID) of the folder containing the library panel.
        /// </summary>
        [Input("folderUid")]
        public Input<string>? FolderUid { get; set; }

        /// <summary>
        /// The JSON model for the library panel.
        /// </summary>
        [Input("modelJson")]
        public Input<string>? ModelJson { get; set; }

        /// <summary>
        /// Name of the library panel.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// The numeric ID of the library panel computed by Grafana.
        /// </summary>
        [Input("panelId")]
        public Input<int>? PanelId { get; set; }

        /// <summary>
        /// Type of the library panel (eg. text).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The unique identifier (UID) of a library panel uniquely identifies library panels between multiple Grafana installs.
        /// It’s automatically generated unless you specify it during library panel creation.The UID provides consistent URLs for
        /// accessing library panels and when syncing library panels between multiple Grafana installs.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// Timestamp when the library panel was last modified.
        /// </summary>
        [Input("updated")]
        public Input<string>? Updated { get; set; }

        /// <summary>
        /// Version of the library panel.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public LibraryPanelState()
        {
        }
        public static new LibraryPanelState Empty => new LibraryPanelState();
    }
}
