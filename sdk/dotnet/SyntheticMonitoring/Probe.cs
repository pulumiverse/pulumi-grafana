// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.SyntheticMonitoring
{
    /// <summary>
    /// Besides the public probes run by Grafana Labs, you can also install your
    /// own private probes. These are only accessible to you and only write data to
    /// your Grafana Cloud account. Private probes are instances of the open source
    /// Grafana Synthetic Monitoring Agent.
    /// 
    /// * [Official documentation](https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/set-up-private-probes/)
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
    ///     var main = new Grafana.SyntheticMonitoring.Probe("main", new()
    ///     {
    ///         Name = "Mount Everest",
    ///         Latitude = 27.98606,
    ///         Longitude = 86.92262,
    ///         Region = "APAC",
    ///         Labels = 
    ///         {
    ///             { "type", "mountain" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:syntheticMonitoring/probe:Probe name "{{ id }}"
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import grafana:syntheticMonitoring/probe:Probe name "{{ id }}:{{ authToken }}"
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:syntheticMonitoring/probe:Probe")]
    public partial class Probe : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The probe authentication token. Your probe must use this to authenticate with Grafana Cloud.
        /// </summary>
        [Output("authToken")]
        public Output<string> AuthToken { get; private set; } = null!;

        /// <summary>
        /// Disables browser checks for this probe. Defaults to `false`.
        /// </summary>
        [Output("disableBrowserChecks")]
        public Output<bool?> DisableBrowserChecks { get; private set; } = null!;

        /// <summary>
        /// Disables scripted checks for this probe. Defaults to `false`.
        /// </summary>
        [Output("disableScriptedChecks")]
        public Output<bool?> DisableScriptedChecks { get; private set; } = null!;

        /// <summary>
        /// Custom labels to be included with collected metrics and logs.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Latitude coordinates.
        /// </summary>
        [Output("latitude")]
        public Output<double> Latitude { get; private set; } = null!;

        /// <summary>
        /// Longitude coordinates.
        /// </summary>
        [Output("longitude")]
        public Output<double> Longitude { get; private set; } = null!;

        /// <summary>
        /// Name of the probe.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
        /// </summary>
        [Output("public")]
        public Output<bool?> Public { get; private set; } = null!;

        /// <summary>
        /// Region of the probe.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The tenant ID of the probe.
        /// </summary>
        [Output("tenantId")]
        public Output<int> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a Probe resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Probe(string name, ProbeArgs args, CustomResourceOptions? options = null)
            : base("grafana:syntheticMonitoring/probe:Probe", name, args ?? new ProbeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Probe(string name, Input<string> id, ProbeState? state = null, CustomResourceOptions? options = null)
            : base("grafana:syntheticMonitoring/probe:Probe", name, state, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "grafana:index/syntheticMonitoringProbe:SyntheticMonitoringProbe" },
                },
                AdditionalSecretOutputs =
                {
                    "authToken",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Probe resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Probe Get(string name, Input<string> id, ProbeState? state = null, CustomResourceOptions? options = null)
        {
            return new Probe(name, id, state, options);
        }
    }

    public sealed class ProbeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disables browser checks for this probe. Defaults to `false`.
        /// </summary>
        [Input("disableBrowserChecks")]
        public Input<bool>? DisableBrowserChecks { get; set; }

        /// <summary>
        /// Disables scripted checks for this probe. Defaults to `false`.
        /// </summary>
        [Input("disableScriptedChecks")]
        public Input<bool>? DisableScriptedChecks { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Custom labels to be included with collected metrics and logs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Latitude coordinates.
        /// </summary>
        [Input("latitude", required: true)]
        public Input<double> Latitude { get; set; } = null!;

        /// <summary>
        /// Longitude coordinates.
        /// </summary>
        [Input("longitude", required: true)]
        public Input<double> Longitude { get; set; } = null!;

        /// <summary>
        /// Name of the probe.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
        /// </summary>
        [Input("public")]
        public Input<bool>? Public { get; set; }

        /// <summary>
        /// Region of the probe.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public ProbeArgs()
        {
        }
        public static new ProbeArgs Empty => new ProbeArgs();
    }

    public sealed class ProbeState : global::Pulumi.ResourceArgs
    {
        [Input("authToken")]
        private Input<string>? _authToken;

        /// <summary>
        /// The probe authentication token. Your probe must use this to authenticate with Grafana Cloud.
        /// </summary>
        public Input<string>? AuthToken
        {
            get => _authToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _authToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Disables browser checks for this probe. Defaults to `false`.
        /// </summary>
        [Input("disableBrowserChecks")]
        public Input<bool>? DisableBrowserChecks { get; set; }

        /// <summary>
        /// Disables scripted checks for this probe. Defaults to `false`.
        /// </summary>
        [Input("disableScriptedChecks")]
        public Input<bool>? DisableScriptedChecks { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Custom labels to be included with collected metrics and logs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Latitude coordinates.
        /// </summary>
        [Input("latitude")]
        public Input<double>? Latitude { get; set; }

        /// <summary>
        /// Longitude coordinates.
        /// </summary>
        [Input("longitude")]
        public Input<double>? Longitude { get; set; }

        /// <summary>
        /// Name of the probe.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
        /// </summary>
        [Input("public")]
        public Input<bool>? Public { get; set; }

        /// <summary>
        /// Region of the probe.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The tenant ID of the probe.
        /// </summary>
        [Input("tenantId")]
        public Input<int>? TenantId { get; set; }

        public ProbeState()
        {
        }
        public static new ProbeState Empty => new ProbeState();
    }
}
