// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Slo
{
    /// <summary>
    /// Resource manages Grafana SLOs.
    /// 
    /// * [Official documentation](https://grafana.com/docs/grafana-cloud/alerting-and-irm/slo/)
    /// * [API documentation](https://grafana.com/docs/grafana-cloud/alerting-and-irm/slo/api/)
    /// * [Additional Information On Alerting Rule Annotations and Labels](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/#templating/)
    /// 
    /// ## Example Usage
    /// 
    /// ### Ratio
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Grafana = Pulumiverse.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var ratio = new Grafana.Slo.SLO("ratio", new()
    ///     {
    ///         Name = "Terraform Testing - Ratio Query",
    ///         Description = "Terraform Description - Ratio Query",
    ///         Queries = new[]
    ///         {
    ///             new Grafana.Slo.Inputs.SLOQueryArgs
    ///             {
    ///                 Ratio = new Grafana.Slo.Inputs.SLOQueryRatioArgs
    ///                 {
    ///                     SuccessMetric = "kubelet_http_requests_total{status!~\"5..\"}",
    ///                     TotalMetric = "kubelet_http_requests_total",
    ///                     GroupByLabels = new[]
    ///                     {
    ///                         "job",
    ///                         "instance",
    ///                     },
    ///                 },
    ///                 Type = "ratio",
    ///             },
    ///         },
    ///         Objectives = new[]
    ///         {
    ///             new Grafana.Slo.Inputs.SLOObjectiveArgs
    ///             {
    ///                 Value = 0.995,
    ///                 Window = "30d",
    ///             },
    ///         },
    ///         DestinationDatasource = new Grafana.Slo.Inputs.SLODestinationDatasourceArgs
    ///         {
    ///             Uid = "grafanacloud-prom",
    ///         },
    ///         Labels = new[]
    ///         {
    ///             new Grafana.Slo.Inputs.SLOLabelArgs
    ///             {
    ///                 Key = "slo",
    ///                 Value = "terraform",
    ///             },
    ///         },
    ///         Alertings = new[]
    ///         {
    ///             new Grafana.Slo.Inputs.SLOAlertingArgs
    ///             {
    ///                 Fastburns = new[]
    ///                 {
    ///                     new Grafana.Slo.Inputs.SLOAlertingFastburnArgs
    ///                     {
    ///                         Annotations = new[]
    ///                         {
    ///                             new Grafana.Slo.Inputs.SLOAlertingFastburnAnnotationArgs
    ///                             {
    ///                                 Key = "name",
    ///                                 Value = "SLO Burn Rate Very High",
    ///                             },
    ///                             new Grafana.Slo.Inputs.SLOAlertingFastburnAnnotationArgs
    ///                             {
    ///                                 Key = "description",
    ///                                 Value = "Error budget is burning too fast",
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///                 Slowburns = new[]
    ///                 {
    ///                     new Grafana.Slo.Inputs.SLOAlertingSlowburnArgs
    ///                     {
    ///                         Annotations = new[]
    ///                         {
    ///                             new Grafana.Slo.Inputs.SLOAlertingSlowburnAnnotationArgs
    ///                             {
    ///                                 Key = "name",
    ///                                 Value = "SLO Burn Rate High",
    ///                             },
    ///                             new Grafana.Slo.Inputs.SLOAlertingSlowburnAnnotationArgs
    ///                             {
    ///                                 Key = "description",
    ///                                 Value = "Error budget is burning too fast",
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Advanced
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Grafana = Pulumiverse.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new Grafana.Slo.SLO("test", new()
    ///     {
    ///         Name = "Terraform Testing",
    ///         Description = "Terraform Description",
    ///         Queries = new[]
    ///         {
    ///             new Grafana.Slo.Inputs.SLOQueryArgs
    ///             {
    ///                 Freeform = new Grafana.Slo.Inputs.SLOQueryFreeformArgs
    ///                 {
    ///                     Query = "sum(rate(apiserver_request_total{code!=\"500\"}[$__rate_interval])) / sum(rate(apiserver_request_total[$__rate_interval]))",
    ///                 },
    ///                 Type = "freeform",
    ///             },
    ///         },
    ///         Objectives = new[]
    ///         {
    ///             new Grafana.Slo.Inputs.SLOObjectiveArgs
    ///             {
    ///                 Value = 0.995,
    ///                 Window = "30d",
    ///             },
    ///         },
    ///         DestinationDatasource = new Grafana.Slo.Inputs.SLODestinationDatasourceArgs
    ///         {
    ///             Uid = "grafanacloud-prom",
    ///         },
    ///         Labels = new[]
    ///         {
    ///             new Grafana.Slo.Inputs.SLOLabelArgs
    ///             {
    ///                 Key = "slo",
    ///                 Value = "terraform",
    ///             },
    ///         },
    ///         Alertings = new[]
    ///         {
    ///             new Grafana.Slo.Inputs.SLOAlertingArgs
    ///             {
    ///                 Fastburns = new[]
    ///                 {
    ///                     new Grafana.Slo.Inputs.SLOAlertingFastburnArgs
    ///                     {
    ///                         Annotations = new[]
    ///                         {
    ///                             new Grafana.Slo.Inputs.SLOAlertingFastburnAnnotationArgs
    ///                             {
    ///                                 Key = "name",
    ///                                 Value = "SLO Burn Rate Very High",
    ///                             },
    ///                             new Grafana.Slo.Inputs.SLOAlertingFastburnAnnotationArgs
    ///                             {
    ///                                 Key = "description",
    ///                                 Value = "Error budget is burning too fast",
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///                 Slowburns = new[]
    ///                 {
    ///                     new Grafana.Slo.Inputs.SLOAlertingSlowburnArgs
    ///                     {
    ///                         Annotations = new[]
    ///                         {
    ///                             new Grafana.Slo.Inputs.SLOAlertingSlowburnAnnotationArgs
    ///                             {
    ///                                 Key = "name",
    ///                                 Value = "SLO Burn Rate High",
    ///                             },
    ///                             new Grafana.Slo.Inputs.SLOAlertingSlowburnAnnotationArgs
    ///                             {
    ///                                 Key = "description",
    ///                                 Value = "Error budget is burning too fast",
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Grafana Queries - Any supported datasource
    /// 
    /// Grafana Queries use the grafana_queries field. It expects a JSON string list of valid grafana query JSON objects, the same as you'll find assigned to a Grafana Dashboard panel `targets` field.
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
    ///     var test = new Grafana.Slo.SLO("test", new()
    ///     {
    ///         Name = "Terraform Testing",
    ///         Description = "Terraform Description",
    ///         Queries = new[]
    ///         {
    ///             new Grafana.Slo.Inputs.SLOQueryArgs
    ///             {
    ///                 GrafanaQueries = new Grafana.Slo.Inputs.SLOQueryGrafanaQueriesArgs
    ///                 {
    ///                     GrafanaQueries = JsonSerializer.Serialize(new[]
    ///                     {
    ///                         new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             ["datasource"] = new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 ["type"] = "graphite",
    ///                                 ["uid"] = "datasource-uid",
    ///                             },
    ///                             ["refId"] = "Success",
    ///                             ["target"] = "groupByNode(perSecond(web.*.http.2xx_success.*.*), 3, 'avg')",
    ///                         },
    ///                         new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             ["datasource"] = new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 ["type"] = "graphite",
    ///                                 ["uid"] = "datasource-uid",
    ///                             },
    ///                             ["refId"] = "Total",
    ///                             ["target"] = "groupByNode(perSecond(web.*.http.5xx_errors.*.*), 3, 'avg')",
    ///                         },
    ///                         new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             ["datasource"] = new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 ["type"] = "__expr__",
    ///                                 ["uid"] = "__expr__",
    ///                             },
    ///                             ["expression"] = "$Success / $Total",
    ///                             ["refId"] = "Expression",
    ///                             ["type"] = "math",
    ///                         },
    ///                     }),
    ///                 },
    ///                 Type = "grafana_queries",
    ///             },
    ///         },
    ///         DestinationDatasource = new Grafana.Slo.Inputs.SLODestinationDatasourceArgs
    ///         {
    ///             Uid = "grafanacloud-prom",
    ///         },
    ///         Objectives = new[]
    ///         {
    ///             new Grafana.Slo.Inputs.SLOObjectiveArgs
    ///             {
    ///                 Value = 0.995,
    ///                 Window = "30d",
    ///             },
    ///         },
    ///         Labels = new[]
    ///         {
    ///             new Grafana.Slo.Inputs.SLOLabelArgs
    ///             {
    ///                 Key = "slo",
    ///                 Value = "terraform",
    ///             },
    ///         },
    ///         Alertings = new[]
    ///         {
    ///             new Grafana.Slo.Inputs.SLOAlertingArgs
    ///             {
    ///                 Fastburns = new[]
    ///                 {
    ///                     new Grafana.Slo.Inputs.SLOAlertingFastburnArgs
    ///                     {
    ///                         Annotations = new[]
    ///                         {
    ///                             new Grafana.Slo.Inputs.SLOAlertingFastburnAnnotationArgs
    ///                             {
    ///                                 Key = "name",
    ///                                 Value = "SLO Burn Rate Very High",
    ///                             },
    ///                             new Grafana.Slo.Inputs.SLOAlertingFastburnAnnotationArgs
    ///                             {
    ///                                 Key = "description",
    ///                                 Value = "Error budget is burning too fast",
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///                 Slowburns = new[]
    ///                 {
    ///                     new Grafana.Slo.Inputs.SLOAlertingSlowburnArgs
    ///                     {
    ///                         Annotations = new[]
    ///                         {
    ///                             new Grafana.Slo.Inputs.SLOAlertingSlowburnAnnotationArgs
    ///                             {
    ///                                 Key = "name",
    ///                                 Value = "SLO Burn Rate High",
    ///                             },
    ///                             new Grafana.Slo.Inputs.SLOAlertingSlowburnAnnotationArgs
    ///                             {
    ///                                 Key = "description",
    ///                                 Value = "Error budget is burning too fast",
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// For a complete list, see [supported data sources](https://grafana.com/docs/grafana-cloud/alerting-and-irm/slo/set-up/additionaldatasources/#supported-data-sources).
    /// 
    /// For additional help with SLOs, view our [documentation](https://grafana.com/docs/grafana-cloud/alerting-and-irm/slo/).
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:slo/sLO:SLO name "{{ uuid }}"
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:slo/sLO:SLO")]
    public partial class SLO : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Configures the alerting rules that will be generated for each
        /// 			time window associated with the SLO. Grafana SLOs can generate
        /// 			alerts when the short-term error budget burn is very high, the
        /// 			long-term error budget burn rate is high, or when the remaining
        /// 			error budget is below a certain threshold. Annotations and Labels support templating.
        /// </summary>
        [Output("alertings")]
        public Output<ImmutableArray<Outputs.SLOAlerting>> Alertings { get; private set; } = null!;

        /// <summary>
        /// Description is a free-text field that can provide more context to an SLO.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Destination Datasource sets the datasource defined for an SLO
        /// </summary>
        [Output("destinationDatasource")]
        public Output<Outputs.SLODestinationDatasource> DestinationDatasource { get; private set; } = null!;

        /// <summary>
        /// UID for the SLO folder
        /// </summary>
        [Output("folderUid")]
        public Output<string?> FolderUid { get; private set; } = null!;

        /// <summary>
        /// Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<Outputs.SLOLabel>> Labels { get; private set; } = null!;

        /// <summary>
        /// Name should be a short description of your indicator. Consider names like "API Availability"
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
        /// </summary>
        [Output("objectives")]
        public Output<ImmutableArray<Outputs.SLOObjective>> Objectives { get; private set; } = null!;

        /// <summary>
        /// Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
        /// </summary>
        [Output("queries")]
        public Output<ImmutableArray<Outputs.SLOQuery>> Queries { get; private set; } = null!;

        /// <summary>
        /// The name of a search expression in Grafana Asserts. This is used in the SLO UI to open the Asserts RCA workbench and in alerts to link to the RCA workbench.
        /// </summary>
        [Output("searchExpression")]
        public Output<string?> SearchExpression { get; private set; } = null!;


        /// <summary>
        /// Create a SLO resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SLO(string name, SLOArgs args, CustomResourceOptions? options = null)
            : base("grafana:slo/sLO:SLO", name, args ?? new SLOArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SLO(string name, Input<string> id, SLOState? state = null, CustomResourceOptions? options = null)
            : base("grafana:slo/sLO:SLO", name, state, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "grafana:index/sLO:SLO" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SLO resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SLO Get(string name, Input<string> id, SLOState? state = null, CustomResourceOptions? options = null)
        {
            return new SLO(name, id, state, options);
        }
    }

    public sealed class SLOArgs : global::Pulumi.ResourceArgs
    {
        [Input("alertings")]
        private InputList<Inputs.SLOAlertingArgs>? _alertings;

        /// <summary>
        /// Configures the alerting rules that will be generated for each
        /// 			time window associated with the SLO. Grafana SLOs can generate
        /// 			alerts when the short-term error budget burn is very high, the
        /// 			long-term error budget burn rate is high, or when the remaining
        /// 			error budget is below a certain threshold. Annotations and Labels support templating.
        /// </summary>
        public InputList<Inputs.SLOAlertingArgs> Alertings
        {
            get => _alertings ?? (_alertings = new InputList<Inputs.SLOAlertingArgs>());
            set => _alertings = value;
        }

        /// <summary>
        /// Description is a free-text field that can provide more context to an SLO.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// Destination Datasource sets the datasource defined for an SLO
        /// </summary>
        [Input("destinationDatasource", required: true)]
        public Input<Inputs.SLODestinationDatasourceArgs> DestinationDatasource { get; set; } = null!;

        /// <summary>
        /// UID for the SLO folder
        /// </summary>
        [Input("folderUid")]
        public Input<string>? FolderUid { get; set; }

        [Input("labels")]
        private InputList<Inputs.SLOLabelArgs>? _labels;

        /// <summary>
        /// Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
        /// </summary>
        public InputList<Inputs.SLOLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.SLOLabelArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// Name should be a short description of your indicator. Consider names like "API Availability"
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("objectives", required: true)]
        private InputList<Inputs.SLOObjectiveArgs>? _objectives;

        /// <summary>
        /// Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
        /// </summary>
        public InputList<Inputs.SLOObjectiveArgs> Objectives
        {
            get => _objectives ?? (_objectives = new InputList<Inputs.SLOObjectiveArgs>());
            set => _objectives = value;
        }

        [Input("queries", required: true)]
        private InputList<Inputs.SLOQueryArgs>? _queries;

        /// <summary>
        /// Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
        /// </summary>
        public InputList<Inputs.SLOQueryArgs> Queries
        {
            get => _queries ?? (_queries = new InputList<Inputs.SLOQueryArgs>());
            set => _queries = value;
        }

        /// <summary>
        /// The name of a search expression in Grafana Asserts. This is used in the SLO UI to open the Asserts RCA workbench and in alerts to link to the RCA workbench.
        /// </summary>
        [Input("searchExpression")]
        public Input<string>? SearchExpression { get; set; }

        public SLOArgs()
        {
        }
        public static new SLOArgs Empty => new SLOArgs();
    }

    public sealed class SLOState : global::Pulumi.ResourceArgs
    {
        [Input("alertings")]
        private InputList<Inputs.SLOAlertingGetArgs>? _alertings;

        /// <summary>
        /// Configures the alerting rules that will be generated for each
        /// 			time window associated with the SLO. Grafana SLOs can generate
        /// 			alerts when the short-term error budget burn is very high, the
        /// 			long-term error budget burn rate is high, or when the remaining
        /// 			error budget is below a certain threshold. Annotations and Labels support templating.
        /// </summary>
        public InputList<Inputs.SLOAlertingGetArgs> Alertings
        {
            get => _alertings ?? (_alertings = new InputList<Inputs.SLOAlertingGetArgs>());
            set => _alertings = value;
        }

        /// <summary>
        /// Description is a free-text field that can provide more context to an SLO.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Destination Datasource sets the datasource defined for an SLO
        /// </summary>
        [Input("destinationDatasource")]
        public Input<Inputs.SLODestinationDatasourceGetArgs>? DestinationDatasource { get; set; }

        /// <summary>
        /// UID for the SLO folder
        /// </summary>
        [Input("folderUid")]
        public Input<string>? FolderUid { get; set; }

        [Input("labels")]
        private InputList<Inputs.SLOLabelGetArgs>? _labels;

        /// <summary>
        /// Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
        /// </summary>
        public InputList<Inputs.SLOLabelGetArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.SLOLabelGetArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// Name should be a short description of your indicator. Consider names like "API Availability"
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("objectives")]
        private InputList<Inputs.SLOObjectiveGetArgs>? _objectives;

        /// <summary>
        /// Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
        /// </summary>
        public InputList<Inputs.SLOObjectiveGetArgs> Objectives
        {
            get => _objectives ?? (_objectives = new InputList<Inputs.SLOObjectiveGetArgs>());
            set => _objectives = value;
        }

        [Input("queries")]
        private InputList<Inputs.SLOQueryGetArgs>? _queries;

        /// <summary>
        /// Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
        /// </summary>
        public InputList<Inputs.SLOQueryGetArgs> Queries
        {
            get => _queries ?? (_queries = new InputList<Inputs.SLOQueryGetArgs>());
            set => _queries = value;
        }

        /// <summary>
        /// The name of a search expression in Grafana Asserts. This is used in the SLO UI to open the Asserts RCA workbench and in alerts to link to the RCA workbench.
        /// </summary>
        [Input("searchExpression")]
        public Input<string>? SearchExpression { get; set; }

        public SLOState()
        {
        }
        public static new SLOState Empty => new SLOState();
    }
}
