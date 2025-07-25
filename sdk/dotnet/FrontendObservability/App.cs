// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.FrontendObservability
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:frontendObservability/app:App name "{{ stack_id }}:{{ name }}"
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:frontendObservability/app:App")]
    public partial class App : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of allowed origins for CORS.
        /// </summary>
        [Output("allowedOrigins")]
        public Output<ImmutableArray<string>> AllowedOrigins { get; private set; } = null!;

        /// <summary>
        /// The collector URL Grafana Cloud Frontend Observability. Use this endpoint to send your Telemetry.
        /// </summary>
        [Output("collectorEndpoint")]
        public Output<string> CollectorEndpoint { get; private set; } = null!;

        /// <summary>
        /// The extra attributes to append in each signal.
        /// </summary>
        [Output("extraLogAttributes")]
        public Output<ImmutableDictionary<string, string>> ExtraLogAttributes { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The key-value settings of the Frontend Observability app. Available Settings: `{combineLabData=(0|1),geolocation.level=(0|1),geolocation.level=0-4,geolocation.country_denylist=&lt;comma-separated-list-of-country-codes&gt;}`
        /// </summary>
        [Output("settings")]
        public Output<ImmutableDictionary<string, string>> Settings { get; private set; } = null!;

        [Output("stackId")]
        public Output<int> StackId { get; private set; } = null!;


        /// <summary>
        /// Create a App resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public App(string name, AppArgs args, CustomResourceOptions? options = null)
            : base("grafana:frontendObservability/app:App", name, args ?? new AppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private App(string name, Input<string> id, AppState? state = null, CustomResourceOptions? options = null)
            : base("grafana:frontendObservability/app:App", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing App resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static App Get(string name, Input<string> id, AppState? state = null, CustomResourceOptions? options = null)
        {
            return new App(name, id, state, options);
        }
    }

    public sealed class AppArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedOrigins", required: true)]
        private InputList<string>? _allowedOrigins;

        /// <summary>
        /// A list of allowed origins for CORS.
        /// </summary>
        public InputList<string> AllowedOrigins
        {
            get => _allowedOrigins ?? (_allowedOrigins = new InputList<string>());
            set => _allowedOrigins = value;
        }

        [Input("extraLogAttributes", required: true)]
        private InputMap<string>? _extraLogAttributes;

        /// <summary>
        /// The extra attributes to append in each signal.
        /// </summary>
        public InputMap<string> ExtraLogAttributes
        {
            get => _extraLogAttributes ?? (_extraLogAttributes = new InputMap<string>());
            set => _extraLogAttributes = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("settings", required: true)]
        private InputMap<string>? _settings;

        /// <summary>
        /// The key-value settings of the Frontend Observability app. Available Settings: `{combineLabData=(0|1),geolocation.level=(0|1),geolocation.level=0-4,geolocation.country_denylist=&lt;comma-separated-list-of-country-codes&gt;}`
        /// </summary>
        public InputMap<string> Settings
        {
            get => _settings ?? (_settings = new InputMap<string>());
            set => _settings = value;
        }

        [Input("stackId", required: true)]
        public Input<int> StackId { get; set; } = null!;

        public AppArgs()
        {
        }
        public static new AppArgs Empty => new AppArgs();
    }

    public sealed class AppState : global::Pulumi.ResourceArgs
    {
        [Input("allowedOrigins")]
        private InputList<string>? _allowedOrigins;

        /// <summary>
        /// A list of allowed origins for CORS.
        /// </summary>
        public InputList<string> AllowedOrigins
        {
            get => _allowedOrigins ?? (_allowedOrigins = new InputList<string>());
            set => _allowedOrigins = value;
        }

        /// <summary>
        /// The collector URL Grafana Cloud Frontend Observability. Use this endpoint to send your Telemetry.
        /// </summary>
        [Input("collectorEndpoint")]
        public Input<string>? CollectorEndpoint { get; set; }

        [Input("extraLogAttributes")]
        private InputMap<string>? _extraLogAttributes;

        /// <summary>
        /// The extra attributes to append in each signal.
        /// </summary>
        public InputMap<string> ExtraLogAttributes
        {
            get => _extraLogAttributes ?? (_extraLogAttributes = new InputMap<string>());
            set => _extraLogAttributes = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("settings")]
        private InputMap<string>? _settings;

        /// <summary>
        /// The key-value settings of the Frontend Observability app. Available Settings: `{combineLabData=(0|1),geolocation.level=(0|1),geolocation.level=0-4,geolocation.country_denylist=&lt;comma-separated-list-of-country-codes&gt;}`
        /// </summary>
        public InputMap<string> Settings
        {
            get => _settings ?? (_settings = new InputMap<string>());
            set => _settings = value;
        }

        [Input("stackId")]
        public Input<int>? StackId { get; set; }

        public AppState()
        {
        }
        public static new AppState Empty => new AppState();
    }
}
