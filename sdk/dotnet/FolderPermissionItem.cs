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
    /// <summary>
    /// Manages a single permission item for a folder. Conflicts with the "grafana.oss.FolderPermission" resource which manages the entire set of permissions for a folder.
    /// 		* [Official documentation](https://grafana.com/docs/grafana/latest/administration/roles-and-permissions/access-control/)
    /// 		* [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/folder_permissions/)
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
    ///     var team = new Grafana.Oss.Team("team");
    /// 
    ///     var user = new Grafana.Oss.User("user", new()
    ///     {
    ///         Email = "user.name@example.com",
    ///         Login = "user.name",
    ///         Password = "my-password",
    ///     });
    /// 
    ///     var collection = new Grafana.Oss.Folder("collection", new()
    ///     {
    ///         Title = "Folder Title",
    ///     });
    /// 
    ///     var onRole = new Grafana.Oss.FolderPermissionItem("onRole", new()
    ///     {
    ///         FolderUid = collection.Uid,
    ///         Role = "Viewer",
    ///         Permission = "Edit",
    ///     });
    /// 
    ///     var onTeam = new Grafana.Oss.FolderPermissionItem("onTeam", new()
    ///     {
    ///         FolderUid = collection.Uid,
    ///         Team = team.Id,
    ///         Permission = "View",
    ///     });
    /// 
    ///     var onUser = new Grafana.Oss.FolderPermissionItem("onUser", new()
    ///     {
    ///         FolderUid = collection.Uid,
    ///         User = user.Id,
    ///         Permission = "Admin",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:index/folderPermissionItem:FolderPermissionItem name "{{ folderUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import grafana:index/folderPermissionItem:FolderPermissionItem name "{{ orgID }}:{{ folderUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
    /// ```
    /// </summary>
    [Obsolete(@"grafana.index/folderpermissionitem.FolderPermissionItem has been deprecated in favor of grafana.oss/folderpermissionitem.FolderPermissionItem")]
    [GrafanaResourceType("grafana:index/folderPermissionItem:FolderPermissionItem")]
    public partial class FolderPermissionItem : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The UID of the folder.
        /// </summary>
        [Output("folderUid")]
        public Output<string> FolderUid { get; private set; } = null!;

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// the permission to be assigned
        /// </summary>
        [Output("permission")]
        public Output<string> Permission { get; private set; } = null!;

        /// <summary>
        /// the role onto which the permission is to be assigned
        /// </summary>
        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;

        /// <summary>
        /// the team onto which the permission is to be assigned
        /// </summary>
        [Output("team")]
        public Output<string?> Team { get; private set; } = null!;

        /// <summary>
        /// the user or service account onto which the permission is to be assigned
        /// </summary>
        [Output("user")]
        public Output<string?> User { get; private set; } = null!;


        /// <summary>
        /// Create a FolderPermissionItem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FolderPermissionItem(string name, FolderPermissionItemArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/folderPermissionItem:FolderPermissionItem", name, args ?? new FolderPermissionItemArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FolderPermissionItem(string name, Input<string> id, FolderPermissionItemState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/folderPermissionItem:FolderPermissionItem", name, state, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "grafana:index/folderPermissionItem:FolderPermissionItem" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FolderPermissionItem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FolderPermissionItem Get(string name, Input<string> id, FolderPermissionItemState? state = null, CustomResourceOptions? options = null)
        {
            return new FolderPermissionItem(name, id, state, options);
        }
    }

    public sealed class FolderPermissionItemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UID of the folder.
        /// </summary>
        [Input("folderUid", required: true)]
        public Input<string> FolderUid { get; set; } = null!;

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// the permission to be assigned
        /// </summary>
        [Input("permission", required: true)]
        public Input<string> Permission { get; set; } = null!;

        /// <summary>
        /// the role onto which the permission is to be assigned
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// the team onto which the permission is to be assigned
        /// </summary>
        [Input("team")]
        public Input<string>? Team { get; set; }

        /// <summary>
        /// the user or service account onto which the permission is to be assigned
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public FolderPermissionItemArgs()
        {
        }
        public static new FolderPermissionItemArgs Empty => new FolderPermissionItemArgs();
    }

    public sealed class FolderPermissionItemState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UID of the folder.
        /// </summary>
        [Input("folderUid")]
        public Input<string>? FolderUid { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// the permission to be assigned
        /// </summary>
        [Input("permission")]
        public Input<string>? Permission { get; set; }

        /// <summary>
        /// the role onto which the permission is to be assigned
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// the team onto which the permission is to be assigned
        /// </summary>
        [Input("team")]
        public Input<string>? Team { get; set; }

        /// <summary>
        /// the user or service account onto which the permission is to be assigned
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public FolderPermissionItemState()
        {
        }
        public static new FolderPermissionItemState Empty => new FolderPermissionItemState();
    }
}
