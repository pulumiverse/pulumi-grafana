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
    /// Manages the entire set of permissions for a dashboard. Permissions that aren't specified when applying this resource will be removed.
    /// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/roles-and-permissions/access-control/)
    /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/dashboard_permissions/)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Grafana = Pulumiverse.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var team = new Grafana.Oss.Team("team", new()
    ///     {
    ///         Name = "Team Name",
    ///     });
    /// 
    ///     var user = new Grafana.Oss.User("user", new()
    ///     {
    ///         Email = "user.name@example.com",
    ///         Password = "my-password",
    ///         Login = "user.name",
    ///     });
    /// 
    ///     var metrics = new Grafana.Oss.Dashboard("metrics", new()
    ///     {
    ///         ConfigJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["title"] = "My Dashboard",
    ///             ["uid"] = "my-dashboard-uid",
    ///         }),
    ///     });
    /// 
    ///     var collectionPermission = new Grafana.Oss.DashboardPermission("collectionPermission", new()
    ///     {
    ///         DashboardUid = metrics.Uid,
    ///         Permissions = new[]
    ///         {
    ///             new Grafana.Oss.Inputs.DashboardPermissionPermissionArgs
    ///             {
    ///                 Role = "Editor",
    ///                 Permission = "Edit",
    ///             },
    ///             new Grafana.Oss.Inputs.DashboardPermissionPermissionArgs
    ///             {
    ///                 TeamId = team.Id,
    ///                 Permission = "View",
    ///             },
    ///             new Grafana.Oss.Inputs.DashboardPermissionPermissionArgs
    ///             {
    ///                 UserId = user.Id,
    ///                 Permission = "Admin",
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
    /// $ pulumi import grafana:index/dashboardPermission:DashboardPermission name "{{ dashboardUID }}"
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import grafana:index/dashboardPermission:DashboardPermission name "{{ orgID }}:{{ dashboardUID }}"
    /// ```
    /// </summary>
    [Obsolete(@"grafana.index/dashboardpermission.DashboardPermission has been deprecated in favor of grafana.oss/dashboardpermission.DashboardPermission")]
    [GrafanaResourceType("grafana:index/dashboardPermission:DashboardPermission")]
    public partial class DashboardPermission : global::Pulumi.CustomResource
    {
        /// <summary>
        /// UID of the dashboard to apply permissions to.
        /// </summary>
        [Output("dashboardUid")]
        public Output<string> DashboardUid { get; private set; } = null!;

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        /// <summary>
        /// The permission items to add/update. Items that are omitted from the list will be removed.
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableArray<Outputs.DashboardPermissionPermission>> Permissions { get; private set; } = null!;


        /// <summary>
        /// Create a DashboardPermission resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DashboardPermission(string name, DashboardPermissionArgs? args = null, CustomResourceOptions? options = null)
            : base("grafana:index/dashboardPermission:DashboardPermission", name, args ?? new DashboardPermissionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DashboardPermission(string name, Input<string> id, DashboardPermissionState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/dashboardPermission:DashboardPermission", name, state, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "grafana:index/dashboardPermission:DashboardPermission" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DashboardPermission resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DashboardPermission Get(string name, Input<string> id, DashboardPermissionState? state = null, CustomResourceOptions? options = null)
        {
            return new DashboardPermission(name, id, state, options);
        }
    }

    public sealed class DashboardPermissionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// UID of the dashboard to apply permissions to.
        /// </summary>
        [Input("dashboardUid")]
        public Input<string>? DashboardUid { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("permissions")]
        private InputList<Inputs.DashboardPermissionPermissionArgs>? _permissions;

        /// <summary>
        /// The permission items to add/update. Items that are omitted from the list will be removed.
        /// </summary>
        public InputList<Inputs.DashboardPermissionPermissionArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.DashboardPermissionPermissionArgs>());
            set => _permissions = value;
        }

        public DashboardPermissionArgs()
        {
        }
        public static new DashboardPermissionArgs Empty => new DashboardPermissionArgs();
    }

    public sealed class DashboardPermissionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// UID of the dashboard to apply permissions to.
        /// </summary>
        [Input("dashboardUid")]
        public Input<string>? DashboardUid { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("permissions")]
        private InputList<Inputs.DashboardPermissionPermissionGetArgs>? _permissions;

        /// <summary>
        /// The permission items to add/update. Items that are omitted from the list will be removed.
        /// </summary>
        public InputList<Inputs.DashboardPermissionPermissionGetArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.DashboardPermissionPermissionGetArgs>());
            set => _permissions = value;
        }

        public DashboardPermissionState()
        {
        }
        public static new DashboardPermissionState Empty => new DashboardPermissionState();
    }
}
