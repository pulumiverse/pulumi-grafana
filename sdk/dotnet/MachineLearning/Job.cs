// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.MachineLearning
{
    /// <summary>
    /// A job defines the queries and model parameters for a machine learning task.
    /// 
    /// See [the Grafana Cloud docs](https://grafana.com/docs/grafana-cloud/alerting-and-irm/machine-learning/dynamic-alerting/forecasting/config/) for more information
    /// on available hyperparameters for use in the `hyper_params` field.
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic Forecast
    /// 
    /// This forecast uses a Prometheus datasource, where the source query is defined in the `expr` field of the `query_params` attribute.
    /// 
    /// Other datasources are supported, but the structure `query_params` may differ.
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
    ///     var foo = new Grafana.Oss.DataSource("foo", new()
    ///     {
    ///         Type = "prometheus",
    ///         Name = "prometheus-ds-test",
    ///         Uid = "prometheus-ds-test-uid",
    ///         Url = "https://my-instance.com",
    ///         BasicAuthEnabled = true,
    ///         BasicAuthUsername = "username",
    ///         JsonDataEncoded = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["httpMethod"] = "POST",
    ///             ["prometheusType"] = "Mimir",
    ///             ["prometheusVersion"] = "2.4.0",
    ///         }),
    ///         SecureJsonDataEncoded = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["basicAuthPassword"] = "password",
    ///         }),
    ///     });
    /// 
    ///     var testJob = new Grafana.MachineLearning.Job("test_job", new()
    ///     {
    ///         Name = "Test Job",
    ///         Metric = "tf_test_job",
    ///         DatasourceType = "prometheus",
    ///         DatasourceUid = foo.Uid,
    ///         QueryParams = 
    ///         {
    ///             { "expr", "grafanacloud_grafana_instance_active_user_count" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Tuned Forecast
    /// 
    /// This forecast has tuned hyperparameters to improve the accuracy of the model.
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
    ///     var foo = new Grafana.Oss.DataSource("foo", new()
    ///     {
    ///         Type = "prometheus",
    ///         Name = "prometheus-ds-test",
    ///         Uid = "prometheus-ds-test-uid",
    ///         Url = "https://my-instance.com",
    ///         BasicAuthEnabled = true,
    ///         BasicAuthUsername = "username",
    ///         JsonDataEncoded = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["httpMethod"] = "POST",
    ///             ["prometheusType"] = "Mimir",
    ///             ["prometheusVersion"] = "2.4.0",
    ///         }),
    ///         SecureJsonDataEncoded = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["basicAuthPassword"] = "password",
    ///         }),
    ///     });
    /// 
    ///     var testJob = new Grafana.MachineLearning.Job("test_job", new()
    ///     {
    ///         Name = "Test Job",
    ///         Metric = "tf_test_job",
    ///         DatasourceType = "prometheus",
    ///         DatasourceUid = foo.Uid,
    ///         QueryParams = 
    ///         {
    ///             { "expr", "grafanacloud_grafana_instance_active_user_count" },
    ///         },
    ///         HyperParams = 
    ///         {
    ///             { "daily_seasonality", "15" },
    ///             { "weekly_seasonality", "10" },
    ///         },
    ///         CustomLabels = 
    ///         {
    ///             { "example_label", "example_value" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Rescaled Forecast
    /// 
    /// This forecast has had the data transformed using a power transformation in order to avoid negative lower predictions.
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
    ///     var foo = new Grafana.Oss.DataSource("foo", new()
    ///     {
    ///         Type = "prometheus",
    ///         Name = "prometheus-ds-test",
    ///         Uid = "prometheus-ds-test-uid",
    ///         Url = "https://my-instance.com",
    ///         BasicAuthEnabled = true,
    ///         BasicAuthUsername = "username",
    ///         JsonDataEncoded = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["httpMethod"] = "POST",
    ///             ["prometheusType"] = "Mimir",
    ///             ["prometheusVersion"] = "2.4.0",
    ///         }),
    ///         SecureJsonDataEncoded = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["basicAuthPassword"] = "password",
    ///         }),
    ///     });
    /// 
    ///     var testJob = new Grafana.MachineLearning.Job("test_job", new()
    ///     {
    ///         Name = "Test Job",
    ///         Metric = "tf_test_job",
    ///         DatasourceType = "prometheus",
    ///         DatasourceUid = foo.Uid,
    ///         QueryParams = 
    ///         {
    ///             { "expr", "grafanacloud_grafana_instance_active_user_count" },
    ///         },
    ///         HyperParams = 
    ///         {
    ///             { "transformation_id", "power" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Forecast with Holidays
    /// 
    /// This forecast has holidays which will be taken into account when training the model.
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
    ///     var foo = new Grafana.Oss.DataSource("foo", new()
    ///     {
    ///         Type = "prometheus",
    ///         Name = "prometheus-ds-test",
    ///         Uid = "prometheus-ds-test-uid",
    ///         Url = "https://my-instance.com",
    ///         BasicAuthEnabled = true,
    ///         BasicAuthUsername = "username",
    ///         JsonDataEncoded = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["httpMethod"] = "POST",
    ///             ["prometheusType"] = "Mimir",
    ///             ["prometheusVersion"] = "2.4.0",
    ///         }),
    ///         SecureJsonDataEncoded = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["basicAuthPassword"] = "password",
    ///         }),
    ///     });
    /// 
    ///     var testHoliday = new Grafana.MachineLearning.Holiday("test_holiday", new()
    ///     {
    ///         Name = "Test Holiday",
    ///         CustomPeriods = new[]
    ///         {
    ///             new Grafana.MachineLearning.Inputs.HolidayCustomPeriodArgs
    ///             {
    ///                 Name = "First of January",
    ///                 StartTime = "2023-01-01T00:00:00Z",
    ///                 EndTime = "2023-01-02T00:00:00Z",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var testJob = new Grafana.MachineLearning.Job("test_job", new()
    ///     {
    ///         Name = "Test Job",
    ///         Metric = "tf_test_job",
    ///         DatasourceType = "prometheus",
    ///         DatasourceUid = foo.Uid,
    ///         QueryParams = 
    ///         {
    ///             { "expr", "grafanacloud_grafana_instance_active_user_count" },
    ///         },
    ///         Holidays = new[]
    ///         {
    ///             testHoliday.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:machineLearning/job:Job name "{{ id }}"
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:machineLearning/job:Job")]
    public partial class Job : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An object representing the custom labels added on the forecast.
        /// </summary>
        [Output("customLabels")]
        public Output<ImmutableDictionary<string, string>?> CustomLabels { get; private set; } = null!;

        /// <summary>
        /// The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        /// </summary>
        [Output("datasourceType")]
        public Output<string> DatasourceType { get; private set; } = null!;

        /// <summary>
        /// The uid of the datasource to query.
        /// </summary>
        [Output("datasourceUid")]
        public Output<string> DatasourceUid { get; private set; } = null!;

        /// <summary>
        /// A description of the job.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A list of holiday IDs or names to take into account when training the model.
        /// </summary>
        [Output("holidays")]
        public Output<ImmutableArray<string>> Holidays { get; private set; } = null!;

        /// <summary>
        /// The hyperparameters used to fine tune the algorithm. See
        /// https://grafana.com/docs/grafana-cloud/alerting-and-irm/machine-learning/forecasts/models/ for the full list of
        /// available hyperparameters.
        /// </summary>
        [Output("hyperParams")]
        public Output<ImmutableDictionary<string, string>?> HyperParams { get; private set; } = null!;

        /// <summary>
        /// The data interval in seconds to train the data on.
        /// </summary>
        [Output("interval")]
        public Output<int?> Interval { get; private set; } = null!;

        /// <summary>
        /// The metric used to query the job results.
        /// </summary>
        [Output("metric")]
        public Output<string> Metric { get; private set; } = null!;

        /// <summary>
        /// The name of the job.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An object representing the query params to query Grafana with.
        /// </summary>
        [Output("queryParams")]
        public Output<ImmutableDictionary<string, string>> QueryParams { get; private set; } = null!;

        /// <summary>
        /// The data interval in seconds to train the data on.
        /// </summary>
        [Output("trainingWindow")]
        public Output<int?> TrainingWindow { get; private set; } = null!;


        /// <summary>
        /// Create a Job resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Job(string name, JobArgs args, CustomResourceOptions? options = null)
            : base("grafana:machineLearning/job:Job", name, args ?? new JobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Job(string name, Input<string> id, JobState? state = null, CustomResourceOptions? options = null)
            : base("grafana:machineLearning/job:Job", name, state, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "grafana:index/machineLearningJob:MachineLearningJob" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Job resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Job Get(string name, Input<string> id, JobState? state = null, CustomResourceOptions? options = null)
        {
            return new Job(name, id, state, options);
        }
    }

    public sealed class JobArgs : global::Pulumi.ResourceArgs
    {
        [Input("customLabels")]
        private InputMap<string>? _customLabels;

        /// <summary>
        /// An object representing the custom labels added on the forecast.
        /// </summary>
        public InputMap<string> CustomLabels
        {
            get => _customLabels ?? (_customLabels = new InputMap<string>());
            set => _customLabels = value;
        }

        /// <summary>
        /// The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        /// </summary>
        [Input("datasourceType", required: true)]
        public Input<string> DatasourceType { get; set; } = null!;

        /// <summary>
        /// The uid of the datasource to query.
        /// </summary>
        [Input("datasourceUid", required: true)]
        public Input<string> DatasourceUid { get; set; } = null!;

        /// <summary>
        /// A description of the job.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("holidays")]
        private InputList<string>? _holidays;

        /// <summary>
        /// A list of holiday IDs or names to take into account when training the model.
        /// </summary>
        public InputList<string> Holidays
        {
            get => _holidays ?? (_holidays = new InputList<string>());
            set => _holidays = value;
        }

        [Input("hyperParams")]
        private InputMap<string>? _hyperParams;

        /// <summary>
        /// The hyperparameters used to fine tune the algorithm. See
        /// https://grafana.com/docs/grafana-cloud/alerting-and-irm/machine-learning/forecasts/models/ for the full list of
        /// available hyperparameters.
        /// </summary>
        public InputMap<string> HyperParams
        {
            get => _hyperParams ?? (_hyperParams = new InputMap<string>());
            set => _hyperParams = value;
        }

        /// <summary>
        /// The data interval in seconds to train the data on.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// The metric used to query the job results.
        /// </summary>
        [Input("metric", required: true)]
        public Input<string> Metric { get; set; } = null!;

        /// <summary>
        /// The name of the job.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("queryParams", required: true)]
        private InputMap<string>? _queryParams;

        /// <summary>
        /// An object representing the query params to query Grafana with.
        /// </summary>
        public InputMap<string> QueryParams
        {
            get => _queryParams ?? (_queryParams = new InputMap<string>());
            set => _queryParams = value;
        }

        /// <summary>
        /// The data interval in seconds to train the data on.
        /// </summary>
        [Input("trainingWindow")]
        public Input<int>? TrainingWindow { get; set; }

        public JobArgs()
        {
        }
        public static new JobArgs Empty => new JobArgs();
    }

    public sealed class JobState : global::Pulumi.ResourceArgs
    {
        [Input("customLabels")]
        private InputMap<string>? _customLabels;

        /// <summary>
        /// An object representing the custom labels added on the forecast.
        /// </summary>
        public InputMap<string> CustomLabels
        {
            get => _customLabels ?? (_customLabels = new InputMap<string>());
            set => _customLabels = value;
        }

        /// <summary>
        /// The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        /// </summary>
        [Input("datasourceType")]
        public Input<string>? DatasourceType { get; set; }

        /// <summary>
        /// The uid of the datasource to query.
        /// </summary>
        [Input("datasourceUid")]
        public Input<string>? DatasourceUid { get; set; }

        /// <summary>
        /// A description of the job.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("holidays")]
        private InputList<string>? _holidays;

        /// <summary>
        /// A list of holiday IDs or names to take into account when training the model.
        /// </summary>
        public InputList<string> Holidays
        {
            get => _holidays ?? (_holidays = new InputList<string>());
            set => _holidays = value;
        }

        [Input("hyperParams")]
        private InputMap<string>? _hyperParams;

        /// <summary>
        /// The hyperparameters used to fine tune the algorithm. See
        /// https://grafana.com/docs/grafana-cloud/alerting-and-irm/machine-learning/forecasts/models/ for the full list of
        /// available hyperparameters.
        /// </summary>
        public InputMap<string> HyperParams
        {
            get => _hyperParams ?? (_hyperParams = new InputMap<string>());
            set => _hyperParams = value;
        }

        /// <summary>
        /// The data interval in seconds to train the data on.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// The metric used to query the job results.
        /// </summary>
        [Input("metric")]
        public Input<string>? Metric { get; set; }

        /// <summary>
        /// The name of the job.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("queryParams")]
        private InputMap<string>? _queryParams;

        /// <summary>
        /// An object representing the query params to query Grafana with.
        /// </summary>
        public InputMap<string> QueryParams
        {
            get => _queryParams ?? (_queryParams = new InputMap<string>());
            set => _queryParams = value;
        }

        /// <summary>
        /// The data interval in seconds to train the data on.
        /// </summary>
        [Input("trainingWindow")]
        public Input<int>? TrainingWindow { get; set; }

        public JobState()
        {
        }
        public static new JobState Empty => new JobState();
    }
}
