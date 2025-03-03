// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
    /// Manages a single permission item for a dashboard. Conflicts with the "grafana.oss.DashboardPermission" resource which manages the entire set of permissions for a dashboard.
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
    ///     var dashboard = new Grafana.Oss.Dashboard("dashboard", new()
    ///     {
    ///         ConfigJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["title"] = "My Dashboard",
    ///             ["uid"] = "my-dashboard-uid",
    ///         }),
    ///     });
    /// 
    ///     var role = new Grafana.Oss.DashboardPermissionItem("role", new()
    ///     {
    ///         DashboardUid = dashboard.Uid,
    ///         Role = "Viewer",
    ///         Permission = "View",
    ///     });
    /// 
    ///     var userDashboardPermissionItem = new Grafana.Oss.DashboardPermissionItem("user", new()
    ///     {
    ///         DashboardUid = dashboard.Uid,
    ///         User = user.Id,
    ///         Permission = "Admin",
    ///     });
    /// 
    ///     var teamDashboardPermissionItem = new Grafana.Oss.DashboardPermissionItem("team", new()
    ///     {
    ///         DashboardUid = dashboard.Uid,
    ///         Team = team.Id,
    ///         Permission = "Edit",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:oss/dashboardPermissionItem:DashboardPermissionItem name "{{ dashboardUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import grafana:oss/dashboardPermissionItem:DashboardPermissionItem name "{{ orgID }}:{{ dashboardUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:oss/dashboardPermissionItem:DashboardPermissionItem")]
    public partial class DashboardPermissionItem : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The UID of the dashboard.
        /// </summary>
        [Output("dashboardUid")]
        public Output<string> DashboardUid { get; private set; } = null!;

        /// <summary>
        /// The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
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
        /// Create a DashboardPermissionItem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DashboardPermissionItem(string name, DashboardPermissionItemArgs args, CustomResourceOptions? options = null)
            : base("grafana:oss/dashboardPermissionItem:DashboardPermissionItem", name, args ?? new DashboardPermissionItemArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DashboardPermissionItem(string name, Input<string> id, DashboardPermissionItemState? state = null, CustomResourceOptions? options = null)
            : base("grafana:oss/dashboardPermissionItem:DashboardPermissionItem", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DashboardPermissionItem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DashboardPermissionItem Get(string name, Input<string> id, DashboardPermissionItemState? state = null, CustomResourceOptions? options = null)
        {
            return new DashboardPermissionItem(name, id, state, options);
        }
    }

    public sealed class DashboardPermissionItemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UID of the dashboard.
        /// </summary>
        [Input("dashboardUid", required: true)]
        public Input<string> DashboardUid { get; set; } = null!;

        /// <summary>
        /// The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
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

        public DashboardPermissionItemArgs()
        {
        }
        public static new DashboardPermissionItemArgs Empty => new DashboardPermissionItemArgs();
    }

    public sealed class DashboardPermissionItemState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UID of the dashboard.
        /// </summary>
        [Input("dashboardUid")]
        public Input<string>? DashboardUid { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
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

        public DashboardPermissionItemState()
        {
        }
        public static new DashboardPermissionItemState Empty => new DashboardPermissionItemState();
    }
}
