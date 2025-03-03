// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.FleetManagement
{
    /// <summary>
    /// Manages Grafana Fleet Management collectors.
    /// 
    /// * [Official documentation](https://grafana.com/docs/grafana-cloud/send-data/fleet-management/)
    /// * [API documentation](https://grafana.com/docs/grafana-cloud/send-data/fleet-management/api-reference/collector-api/)
    /// 
    /// **Note:** Fleet Management is in [public preview](https://grafana.com/docs/release-life-cycle/#public-preview) and this resource is experimental. Grafana Labs offers limited support, and breaking changes might occur.
    /// 
    /// Required access policy scopes:
    /// 
    /// * fleet-management:read
    /// * fleet-management:write
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:fleetManagement/collector:Collector name "{{ id }}"
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:fleetManagement/collector:Collector")]
    public partial class Collector : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the collector is enabled or not
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// Remote attributes for the collector
        /// </summary>
        [Output("remoteAttributes")]
        public Output<ImmutableDictionary<string, string>> RemoteAttributes { get; private set; } = null!;


        /// <summary>
        /// Create a Collector resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Collector(string name, CollectorArgs? args = null, CustomResourceOptions? options = null)
            : base("grafana:fleetManagement/collector:Collector", name, args ?? new CollectorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Collector(string name, Input<string> id, CollectorState? state = null, CustomResourceOptions? options = null)
            : base("grafana:fleetManagement/collector:Collector", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Collector resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Collector Get(string name, Input<string> id, CollectorState? state = null, CustomResourceOptions? options = null)
        {
            return new Collector(name, id, state, options);
        }
    }

    public sealed class CollectorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the collector is enabled or not
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("remoteAttributes")]
        private InputMap<string>? _remoteAttributes;

        /// <summary>
        /// Remote attributes for the collector
        /// </summary>
        public InputMap<string> RemoteAttributes
        {
            get => _remoteAttributes ?? (_remoteAttributes = new InputMap<string>());
            set => _remoteAttributes = value;
        }

        public CollectorArgs()
        {
        }
        public static new CollectorArgs Empty => new CollectorArgs();
    }

    public sealed class CollectorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the collector is enabled or not
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("remoteAttributes")]
        private InputMap<string>? _remoteAttributes;

        /// <summary>
        /// Remote attributes for the collector
        /// </summary>
        public InputMap<string> RemoteAttributes
        {
            get => _remoteAttributes ?? (_remoteAttributes = new InputMap<string>());
            set => _remoteAttributes = value;
        }

        public CollectorState()
        {
        }
        public static new CollectorState Empty => new CollectorState();
    }
}
