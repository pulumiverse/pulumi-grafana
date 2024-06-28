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
    /// * [Official documentation](https://grafana.com/docs/grafana/latest/datasources/)
    /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/data_source/)
    /// 
    /// The required arguments for this resource vary depending on the type of data
    /// source selected (via the 'type' argument).
    /// 
    /// Use this resource for configuring multiple datasources, when that configuration (`json_data_encoded` field) requires circular references like in the example below.
    /// 
    /// &gt; When using the `grafana.oss.DataSourceConfig` resource, the corresponding `grafana.oss.DataSource` resources must have the `json_data_encoded` and `http_headers` fields ignored. Otherwise, an infinite update loop will occur. See the example below.
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
    ///     var lokiDataSource = new Grafana.Oss.DataSource("lokiDataSource", new()
    ///     {
    ///         Type = "loki",
    ///         Url = "http://localhost:3100",
    ///     });
    /// 
    ///     var tempoDataSource = new Grafana.Oss.DataSource("tempoDataSource", new()
    ///     {
    ///         Type = "tempo",
    ///         Url = "http://localhost:3200",
    ///     });
    /// 
    ///     var lokiDataSourceConfig = new Grafana.Oss.DataSourceConfig("lokiDataSourceConfig", new()
    ///     {
    ///         Uid = lokiDataSource.Uid,
    ///         JsonDataEncoded = Output.JsonSerialize(Output.Create(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["derivedFields"] = new[]
    ///             {
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["datasourceUid"] = tempoDataSource.Uid,
    ///                     ["matcherRegex"] = "[tT]race_?[iI][dD]\"?[:=]\"?(\\w+)",
    ///                     ["matcherType"] = "regex",
    ///                     ["name"] = "traceID",
    ///                     ["url"] = "${__value.raw}",
    ///                 },
    ///             },
    ///         })),
    ///     });
    /// 
    ///     var tempoDataSourceConfig = new Grafana.Oss.DataSourceConfig("tempoDataSourceConfig", new()
    ///     {
    ///         Uid = tempoDataSource.Uid,
    ///         JsonDataEncoded = Output.JsonSerialize(Output.Create(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["tracesToLogsV2"] = new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 ["customQuery"] = true,
    ///                 ["datasourceUid"] = lokiDataSource.Uid,
    ///                 ["filterBySpanID"] = false,
    ///                 ["filterByTraceID"] = false,
    ///                 ["query"] = "|=\"${__trace.traceId}\" | json",
    ///             },
    ///         })),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:index/dataSourceConfig:DataSourceConfig name "{{ uid }}"
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import grafana:index/dataSourceConfig:DataSourceConfig name "{{ orgID }}:{{ uid }}"
    /// ```
    /// </summary>
    [Obsolete(@"grafana.index/datasourceconfig.DataSourceConfig has been deprecated in favor of grafana.oss/datasourceconfig.DataSourceConfig")]
    [GrafanaResourceType("grafana:index/dataSourceConfig:DataSourceConfig")]
    public partial class DataSourceConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Custom HTTP headers
        /// </summary>
        [Output("httpHeaders")]
        public Output<ImmutableDictionary<string, string>?> HttpHeaders { get; private set; } = null!;

        /// <summary>
        /// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        /// </summary>
        [Output("jsonDataEncoded")]
        public Output<string?> JsonDataEncoded { get; private set; } = null!;

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        /// <summary>
        /// Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        /// </summary>
        [Output("secureJsonDataEncoded")]
        public Output<string?> SecureJsonDataEncoded { get; private set; } = null!;

        /// <summary>
        /// Unique identifier. If unset, this will be automatically generated.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;


        /// <summary>
        /// Create a DataSourceConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataSourceConfig(string name, DataSourceConfigArgs? args = null, CustomResourceOptions? options = null)
            : base("grafana:index/dataSourceConfig:DataSourceConfig", name, args ?? new DataSourceConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataSourceConfig(string name, Input<string> id, DataSourceConfigState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/dataSourceConfig:DataSourceConfig", name, state, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "grafana:index/dataSourceConfig:DataSourceConfig" },
                },
                AdditionalSecretOutputs =
                {
                    "httpHeaders",
                    "secureJsonDataEncoded",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DataSourceConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataSourceConfig Get(string name, Input<string> id, DataSourceConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new DataSourceConfig(name, id, state, options);
        }
    }

    public sealed class DataSourceConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("httpHeaders")]
        private InputMap<string>? _httpHeaders;

        /// <summary>
        /// Custom HTTP headers
        /// </summary>
        public InputMap<string> HttpHeaders
        {
            get => _httpHeaders ?? (_httpHeaders = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _httpHeaders = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        /// </summary>
        [Input("jsonDataEncoded")]
        public Input<string>? JsonDataEncoded { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("secureJsonDataEncoded")]
        private Input<string>? _secureJsonDataEncoded;

        /// <summary>
        /// Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        /// </summary>
        public Input<string>? SecureJsonDataEncoded
        {
            get => _secureJsonDataEncoded;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secureJsonDataEncoded = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Unique identifier. If unset, this will be automatically generated.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        public DataSourceConfigArgs()
        {
        }
        public static new DataSourceConfigArgs Empty => new DataSourceConfigArgs();
    }

    public sealed class DataSourceConfigState : global::Pulumi.ResourceArgs
    {
        [Input("httpHeaders")]
        private InputMap<string>? _httpHeaders;

        /// <summary>
        /// Custom HTTP headers
        /// </summary>
        public InputMap<string> HttpHeaders
        {
            get => _httpHeaders ?? (_httpHeaders = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _httpHeaders = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        /// </summary>
        [Input("jsonDataEncoded")]
        public Input<string>? JsonDataEncoded { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("secureJsonDataEncoded")]
        private Input<string>? _secureJsonDataEncoded;

        /// <summary>
        /// Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        /// </summary>
        public Input<string>? SecureJsonDataEncoded
        {
            get => _secureJsonDataEncoded;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secureJsonDataEncoded = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Unique identifier. If unset, this will be automatically generated.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        public DataSourceConfigState()
        {
        }
        public static new DataSourceConfigState Empty => new DataSourceConfigState();
    }
}
