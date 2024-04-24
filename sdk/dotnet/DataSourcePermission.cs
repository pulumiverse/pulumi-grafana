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
    [GrafanaResourceType("grafana:index/dataSourcePermission:DataSourcePermission")]
    public partial class DataSourcePermission : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Deprecated: Use `datasource_uid` instead.
        /// </summary>
        [Output("datasourceId")]
        public Output<string> DatasourceId { get; private set; } = null!;

        /// <summary>
        /// UID of the datasource to apply permissions to.
        /// </summary>
        [Output("datasourceUid")]
        public Output<string> DatasourceUid { get; private set; } = null!;

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        /// <summary>
        /// The permission items to add/update. Items that are omitted from the list will be removed.
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableArray<Outputs.DataSourcePermissionPermission>> Permissions { get; private set; } = null!;


        /// <summary>
        /// Create a DataSourcePermission resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataSourcePermission(string name, DataSourcePermissionArgs? args = null, CustomResourceOptions? options = null)
            : base("grafana:index/dataSourcePermission:DataSourcePermission", name, args ?? new DataSourcePermissionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataSourcePermission(string name, Input<string> id, DataSourcePermissionState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/dataSourcePermission:DataSourcePermission", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DataSourcePermission resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataSourcePermission Get(string name, Input<string> id, DataSourcePermissionState? state = null, CustomResourceOptions? options = null)
        {
            return new DataSourcePermission(name, id, state, options);
        }
    }

    public sealed class DataSourcePermissionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Deprecated: Use `datasource_uid` instead.
        /// </summary>
        [Input("datasourceId")]
        public Input<string>? DatasourceId { get; set; }

        /// <summary>
        /// UID of the datasource to apply permissions to.
        /// </summary>
        [Input("datasourceUid")]
        public Input<string>? DatasourceUid { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("permissions")]
        private InputList<Inputs.DataSourcePermissionPermissionArgs>? _permissions;

        /// <summary>
        /// The permission items to add/update. Items that are omitted from the list will be removed.
        /// </summary>
        public InputList<Inputs.DataSourcePermissionPermissionArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.DataSourcePermissionPermissionArgs>());
            set => _permissions = value;
        }

        public DataSourcePermissionArgs()
        {
        }
        public static new DataSourcePermissionArgs Empty => new DataSourcePermissionArgs();
    }

    public sealed class DataSourcePermissionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Deprecated: Use `datasource_uid` instead.
        /// </summary>
        [Input("datasourceId")]
        public Input<string>? DatasourceId { get; set; }

        /// <summary>
        /// UID of the datasource to apply permissions to.
        /// </summary>
        [Input("datasourceUid")]
        public Input<string>? DatasourceUid { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("permissions")]
        private InputList<Inputs.DataSourcePermissionPermissionGetArgs>? _permissions;

        /// <summary>
        /// The permission items to add/update. Items that are omitted from the list will be removed.
        /// </summary>
        public InputList<Inputs.DataSourcePermissionPermissionGetArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.DataSourcePermissionPermissionGetArgs>());
            set => _permissions = value;
        }

        public DataSourcePermissionState()
        {
        }
        public static new DataSourcePermissionState Empty => new DataSourcePermissionState();
    }
}
