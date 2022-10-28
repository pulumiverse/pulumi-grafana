// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Grafana
{
    /// <summary>
    /// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/manage-organizations/)
    /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/preferences/#get-current-org-prefs)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Grafana = Lbrlabs.PulumiPackage.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new Grafana.OrganizationPreference("test", new()
    ///     {
    ///         Theme = "light",
    ///         Timezone = "utc",
    ///         WeekStart = "Tuesday",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:index/organizationPreference:OrganizationPreference")]
    public partial class OrganizationPreference : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Organization home dashboard ID.
        /// </summary>
        [Output("homeDashboardId")]
        public Output<int?> HomeDashboardId { get; private set; } = null!;

        /// <summary>
        /// The Organization home dashboard UID.
        /// </summary>
        [Output("homeDashboardUid")]
        public Output<string?> HomeDashboardUid { get; private set; } = null!;

        /// <summary>
        /// The Organization theme. Available values are `light`, `dark`, or an empty string for the default.
        /// </summary>
        [Output("theme")]
        public Output<string?> Theme { get; private set; } = null!;

        /// <summary>
        /// The Organization timezone. Available values are `utc`, `browser`, or an empty string for the default.
        /// </summary>
        [Output("timezone")]
        public Output<string?> Timezone { get; private set; } = null!;

        /// <summary>
        /// The Organization week start.
        /// </summary>
        [Output("weekStart")]
        public Output<string?> WeekStart { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationPreference resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationPreference(string name, OrganizationPreferenceArgs? args = null, CustomResourceOptions? options = null)
            : base("grafana:index/organizationPreference:OrganizationPreference", name, args ?? new OrganizationPreferenceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationPreference(string name, Input<string> id, OrganizationPreferenceState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/organizationPreference:OrganizationPreference", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lbrlabs",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OrganizationPreference resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationPreference Get(string name, Input<string> id, OrganizationPreferenceState? state = null, CustomResourceOptions? options = null)
        {
            return new OrganizationPreference(name, id, state, options);
        }
    }

    public sealed class OrganizationPreferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Organization home dashboard ID.
        /// </summary>
        [Input("homeDashboardId")]
        public Input<int>? HomeDashboardId { get; set; }

        /// <summary>
        /// The Organization home dashboard UID.
        /// </summary>
        [Input("homeDashboardUid")]
        public Input<string>? HomeDashboardUid { get; set; }

        /// <summary>
        /// The Organization theme. Available values are `light`, `dark`, or an empty string for the default.
        /// </summary>
        [Input("theme")]
        public Input<string>? Theme { get; set; }

        /// <summary>
        /// The Organization timezone. Available values are `utc`, `browser`, or an empty string for the default.
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        /// <summary>
        /// The Organization week start.
        /// </summary>
        [Input("weekStart")]
        public Input<string>? WeekStart { get; set; }

        public OrganizationPreferenceArgs()
        {
        }
        public static new OrganizationPreferenceArgs Empty => new OrganizationPreferenceArgs();
    }

    public sealed class OrganizationPreferenceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Organization home dashboard ID.
        /// </summary>
        [Input("homeDashboardId")]
        public Input<int>? HomeDashboardId { get; set; }

        /// <summary>
        /// The Organization home dashboard UID.
        /// </summary>
        [Input("homeDashboardUid")]
        public Input<string>? HomeDashboardUid { get; set; }

        /// <summary>
        /// The Organization theme. Available values are `light`, `dark`, or an empty string for the default.
        /// </summary>
        [Input("theme")]
        public Input<string>? Theme { get; set; }

        /// <summary>
        /// The Organization timezone. Available values are `utc`, `browser`, or an empty string for the default.
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        /// <summary>
        /// The Organization week start.
        /// </summary>
        [Input("weekStart")]
        public Input<string>? WeekStart { get; set; }

        public OrganizationPreferenceState()
        {
        }
        public static new OrganizationPreferenceState Empty => new OrganizationPreferenceState();
    }
}
