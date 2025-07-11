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
    /// Manages a single permission item for a datasource. Conflicts with the "grafana.enterprise.DataSourcePermission" resource which manages the entire set of permissions for a datasource.
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
    ///     var foo = new Grafana.Oss.DataSource("foo", new()
    ///     {
    ///         Type = "cloudwatch",
    ///         Name = "cw-example",
    ///         JsonDataEncoded = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["defaultRegion"] = "us-east-1",
    ///             ["authType"] = "keys",
    ///         }),
    ///         SecureJsonDataEncoded = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["accessKey"] = "123",
    ///             ["secretKey"] = "456",
    ///         }),
    ///     });
    /// 
    ///     var user = new Grafana.Oss.User("user", new()
    ///     {
    ///         Name = "test-ds-permissions",
    ///         Email = "test-ds-permissions@example.com",
    ///         Login = "test-ds-permissions",
    ///         Password = "hunter2",
    ///     });
    /// 
    ///     var sa = new Grafana.Oss.ServiceAccount("sa", new()
    ///     {
    ///         Name = "test-ds-permissions",
    ///         Role = "Viewer",
    ///     });
    /// 
    ///     var teamDataSourcePermissionItem = new Grafana.Enterprise.DataSourcePermissionItem("team", new()
    ///     {
    ///         DatasourceUid = foo.Uid,
    ///         Team = team.Id,
    ///         Permission = "Edit",
    ///     });
    /// 
    ///     var userDataSourcePermissionItem = new Grafana.Enterprise.DataSourcePermissionItem("user", new()
    ///     {
    ///         DatasourceUid = foo.Uid,
    ///         User = user.Id,
    ///         Permission = "Edit",
    ///     });
    /// 
    ///     var role = new Grafana.Enterprise.DataSourcePermissionItem("role", new()
    ///     {
    ///         DatasourceUid = foo.Uid,
    ///         Role = "Viewer",
    ///         Permission = "Query",
    ///     });
    /// 
    ///     var serviceAccount = new Grafana.Enterprise.DataSourcePermissionItem("service_account", new()
    ///     {
    ///         DatasourceUid = foo.Uid,
    ///         User = sa.Id,
    ///         Permission = "Query",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:index/dataSourcePermissionItem:DataSourcePermissionItem name "{{ datasourceUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import grafana:index/dataSourcePermissionItem:DataSourcePermissionItem name "{{ orgID }}:{{ datasourceUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
    /// ```
    /// </summary>
    [Obsolete(@"grafana.index/datasourcepermissionitem.DataSourcePermissionItem has been deprecated in favor of grafana.enterprise/datasourcepermissionitem.DataSourcePermissionItem")]
    [GrafanaResourceType("grafana:index/dataSourcePermissionItem:DataSourcePermissionItem")]
    public partial class DataSourcePermissionItem : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The UID of the datasource.
        /// </summary>
        [Output("datasourceUid")]
        public Output<string> DatasourceUid { get; private set; } = null!;

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
        /// Create a DataSourcePermissionItem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataSourcePermissionItem(string name, DataSourcePermissionItemArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/dataSourcePermissionItem:DataSourcePermissionItem", name, args ?? new DataSourcePermissionItemArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataSourcePermissionItem(string name, Input<string> id, DataSourcePermissionItemState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/dataSourcePermissionItem:DataSourcePermissionItem", name, state, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "grafana:index/dataSourcePermissionItem:DataSourcePermissionItem" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DataSourcePermissionItem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataSourcePermissionItem Get(string name, Input<string> id, DataSourcePermissionItemState? state = null, CustomResourceOptions? options = null)
        {
            return new DataSourcePermissionItem(name, id, state, options);
        }
    }

    public sealed class DataSourcePermissionItemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UID of the datasource.
        /// </summary>
        [Input("datasourceUid", required: true)]
        public Input<string> DatasourceUid { get; set; } = null!;

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

        public DataSourcePermissionItemArgs()
        {
        }
        public static new DataSourcePermissionItemArgs Empty => new DataSourcePermissionItemArgs();
    }

    public sealed class DataSourcePermissionItemState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UID of the datasource.
        /// </summary>
        [Input("datasourceUid")]
        public Input<string>? DatasourceUid { get; set; }

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

        public DataSourcePermissionItemState()
        {
        }
        public static new DataSourcePermissionItemState Empty => new DataSourcePermissionItemState();
    }
}
