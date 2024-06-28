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
    /// **Note:** This resource is available only with Grafana Enterprise 8.+.
    /// 
    /// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/roles-and-permissions/access-control/)
    /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/access_control/)
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
    ///     var superUser = new Grafana.Enterprise.Role("superUser", new()
    ///     {
    ///         Description = "My Super User description",
    ///         Global = true,
    ///         Permissions = new[]
    ///         {
    ///             new Grafana.Enterprise.Inputs.RolePermissionArgs
    ///             {
    ///                 Action = "org.users:add",
    ///                 Scope = "users:*",
    ///             },
    ///             new Grafana.Enterprise.Inputs.RolePermissionArgs
    ///             {
    ///                 Action = "org.users:write",
    ///                 Scope = "users:*",
    ///             },
    ///             new Grafana.Enterprise.Inputs.RolePermissionArgs
    ///             {
    ///                 Action = "org.users:read",
    ///                 Scope = "users:*",
    ///             },
    ///         },
    ///         Uid = "superuseruid",
    ///         Version = 1,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:index/role:Role name "{{ uid }}"
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import grafana:index/role:Role name "{{ orgID }}:{{ uid }}"
    /// ```
    /// </summary>
    [Obsolete(@"grafana.index/role.Role has been deprecated in favor of grafana.enterprise/role.Role")]
    [GrafanaResourceType("grafana:index/role:Role")]
    public partial class Role : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the role version should be incremented automatically on updates (and set to 1 on creation). This field or `version` should be set.
        /// </summary>
        [Output("autoIncrementVersion")]
        public Output<bool?> AutoIncrementVersion { get; private set; } = null!;

        /// <summary>
        /// Description of the role.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Display name of the role. Available with Grafana 8.5+.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Boolean to state whether the role is available across all organizations or not. Defaults to `false`.
        /// </summary>
        [Output("global")]
        public Output<bool?> Global { get; private set; } = null!;

        /// <summary>
        /// Group of the role. Available with Grafana 8.5+.
        /// </summary>
        [Output("group")]
        public Output<string?> Group { get; private set; } = null!;

        /// <summary>
        /// Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+. Defaults to `false`.
        /// </summary>
        [Output("hidden")]
        public Output<bool?> Hidden { get; private set; } = null!;

        /// <summary>
        /// Name of the role
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        /// <summary>
        /// Specific set of actions granted by the role.
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableArray<Outputs.RolePermission>> Permissions { get; private set; } = null!;

        /// <summary>
        /// Unique identifier of the role. Used for assignments.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// Version of the role. A role is updated only on version increase. This field or `auto_increment_version` should be set.
        /// </summary>
        [Output("version")]
        public Output<int?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Role resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Role(string name, RoleArgs? args = null, CustomResourceOptions? options = null)
            : base("grafana:index/role:Role", name, args ?? new RoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Role(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/role:Role", name, state, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "grafana:index/role:Role" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Role resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Role Get(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
        {
            return new Role(name, id, state, options);
        }
    }

    public sealed class RoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the role version should be incremented automatically on updates (and set to 1 on creation). This field or `version` should be set.
        /// </summary>
        [Input("autoIncrementVersion")]
        public Input<bool>? AutoIncrementVersion { get; set; }

        /// <summary>
        /// Description of the role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Display name of the role. Available with Grafana 8.5+.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Boolean to state whether the role is available across all organizations or not. Defaults to `false`.
        /// </summary>
        [Input("global")]
        public Input<bool>? Global { get; set; }

        /// <summary>
        /// Group of the role. Available with Grafana 8.5+.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+. Defaults to `false`.
        /// </summary>
        [Input("hidden")]
        public Input<bool>? Hidden { get; set; }

        /// <summary>
        /// Name of the role
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("permissions")]
        private InputList<Inputs.RolePermissionArgs>? _permissions;

        /// <summary>
        /// Specific set of actions granted by the role.
        /// </summary>
        public InputList<Inputs.RolePermissionArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.RolePermissionArgs>());
            set => _permissions = value;
        }

        /// <summary>
        /// Unique identifier of the role. Used for assignments.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// Version of the role. A role is updated only on version increase. This field or `auto_increment_version` should be set.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public RoleArgs()
        {
        }
        public static new RoleArgs Empty => new RoleArgs();
    }

    public sealed class RoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the role version should be incremented automatically on updates (and set to 1 on creation). This field or `version` should be set.
        /// </summary>
        [Input("autoIncrementVersion")]
        public Input<bool>? AutoIncrementVersion { get; set; }

        /// <summary>
        /// Description of the role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Display name of the role. Available with Grafana 8.5+.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Boolean to state whether the role is available across all organizations or not. Defaults to `false`.
        /// </summary>
        [Input("global")]
        public Input<bool>? Global { get; set; }

        /// <summary>
        /// Group of the role. Available with Grafana 8.5+.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+. Defaults to `false`.
        /// </summary>
        [Input("hidden")]
        public Input<bool>? Hidden { get; set; }

        /// <summary>
        /// Name of the role
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("permissions")]
        private InputList<Inputs.RolePermissionGetArgs>? _permissions;

        /// <summary>
        /// Specific set of actions granted by the role.
        /// </summary>
        public InputList<Inputs.RolePermissionGetArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.RolePermissionGetArgs>());
            set => _permissions = value;
        }

        /// <summary>
        /// Unique identifier of the role. Used for assignments.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// Version of the role. A role is updated only on version increase. This field or `auto_increment_version` should be set.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public RoleState()
        {
        }
        public static new RoleState Empty => new RoleState();
    }
}
