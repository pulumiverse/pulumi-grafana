// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * A job defines the queries and model parameters for a machine learning task.
 *
 * See [the Grafana Cloud docs](https://grafana.com/docs/grafana-cloud/alerting-and-irm/machine-learning/dynamic-alerting/forecasting/config/) for more information
 * on available hyperparameters for use in the `hyperParams` field.
 *
 * ## Example Usage
 *
 * ### Basic Forecast
 *
 * This forecast uses a Prometheus datasource, where the source query is defined in the `expr` field of the `queryParams` attribute.
 *
 * Other datasources are supported, but the structure `queryParams` may differ.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const foo = new grafana.oss.DataSource("foo", {
 *     type: "prometheus",
 *     name: "prometheus-ds-test",
 *     uid: "prometheus-ds-test-uid",
 *     url: "https://my-instance.com",
 *     basicAuthEnabled: true,
 *     basicAuthUsername: "username",
 *     jsonDataEncoded: JSON.stringify({
 *         httpMethod: "POST",
 *         prometheusType: "Mimir",
 *         prometheusVersion: "2.4.0",
 *     }),
 *     secureJsonDataEncoded: JSON.stringify({
 *         basicAuthPassword: "password",
 *     }),
 * });
 * const testJob = new grafana.machinelearning.Job("test_job", {
 *     name: "Test Job",
 *     metric: "tf_test_job",
 *     datasourceType: "prometheus",
 *     datasourceUid: foo.uid,
 *     queryParams: {
 *         expr: "grafanacloud_grafana_instance_active_user_count",
 *     },
 * });
 * ```
 *
 * ### Tuned Forecast
 *
 * This forecast has tuned hyperparameters to improve the accuracy of the model.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const foo = new grafana.oss.DataSource("foo", {
 *     type: "prometheus",
 *     name: "prometheus-ds-test",
 *     uid: "prometheus-ds-test-uid",
 *     url: "https://my-instance.com",
 *     basicAuthEnabled: true,
 *     basicAuthUsername: "username",
 *     jsonDataEncoded: JSON.stringify({
 *         httpMethod: "POST",
 *         prometheusType: "Mimir",
 *         prometheusVersion: "2.4.0",
 *     }),
 *     secureJsonDataEncoded: JSON.stringify({
 *         basicAuthPassword: "password",
 *     }),
 * });
 * const testJob = new grafana.machinelearning.Job("test_job", {
 *     name: "Test Job",
 *     metric: "tf_test_job",
 *     datasourceType: "prometheus",
 *     datasourceUid: foo.uid,
 *     queryParams: {
 *         expr: "grafanacloud_grafana_instance_active_user_count",
 *     },
 *     hyperParams: {
 *         daily_seasonality: "15",
 *         weekly_seasonality: "10",
 *     },
 *     customLabels: {
 *         example_label: "example_value",
 *     },
 * });
 * ```
 *
 * ### Rescaled Forecast
 *
 * This forecast has had the data transformed using a power transformation in order to avoid negative lower predictions.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const foo = new grafana.oss.DataSource("foo", {
 *     type: "prometheus",
 *     name: "prometheus-ds-test",
 *     uid: "prometheus-ds-test-uid",
 *     url: "https://my-instance.com",
 *     basicAuthEnabled: true,
 *     basicAuthUsername: "username",
 *     jsonDataEncoded: JSON.stringify({
 *         httpMethod: "POST",
 *         prometheusType: "Mimir",
 *         prometheusVersion: "2.4.0",
 *     }),
 *     secureJsonDataEncoded: JSON.stringify({
 *         basicAuthPassword: "password",
 *     }),
 * });
 * const testJob = new grafana.machinelearning.Job("test_job", {
 *     name: "Test Job",
 *     metric: "tf_test_job",
 *     datasourceType: "prometheus",
 *     datasourceUid: foo.uid,
 *     queryParams: {
 *         expr: "grafanacloud_grafana_instance_active_user_count",
 *     },
 *     hyperParams: {
 *         transformation_id: "power",
 *     },
 * });
 * ```
 *
 * ### Forecast with Holidays
 *
 * This forecast has holidays which will be taken into account when training the model.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const foo = new grafana.oss.DataSource("foo", {
 *     type: "prometheus",
 *     name: "prometheus-ds-test",
 *     uid: "prometheus-ds-test-uid",
 *     url: "https://my-instance.com",
 *     basicAuthEnabled: true,
 *     basicAuthUsername: "username",
 *     jsonDataEncoded: JSON.stringify({
 *         httpMethod: "POST",
 *         prometheusType: "Mimir",
 *         prometheusVersion: "2.4.0",
 *     }),
 *     secureJsonDataEncoded: JSON.stringify({
 *         basicAuthPassword: "password",
 *     }),
 * });
 * const testHoliday = new grafana.machinelearning.Holiday("test_holiday", {
 *     name: "Test Holiday",
 *     customPeriods: [{
 *         name: "First of January",
 *         startTime: "2023-01-01T00:00:00Z",
 *         endTime: "2023-01-02T00:00:00Z",
 *     }],
 * });
 * const testJob = new grafana.machinelearning.Job("test_job", {
 *     name: "Test Job",
 *     metric: "tf_test_job",
 *     datasourceType: "prometheus",
 *     datasourceUid: foo.uid,
 *     queryParams: {
 *         expr: "grafanacloud_grafana_instance_active_user_count",
 *     },
 *     holidays: [testHoliday.id],
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:index/machineLearningJob:MachineLearningJob name "{{ id }}"
 * ```
 *
 * @deprecated grafana.index/machinelearningjob.MachineLearningJob has been deprecated in favor of grafana.machinelearning/job.Job
 */
export class MachineLearningJob extends pulumi.CustomResource {
    /**
     * Get an existing MachineLearningJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MachineLearningJobState, opts?: pulumi.CustomResourceOptions): MachineLearningJob {
        pulumi.log.warn("MachineLearningJob is deprecated: grafana.index/machinelearningjob.MachineLearningJob has been deprecated in favor of grafana.machinelearning/job.Job")
        return new MachineLearningJob(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/machineLearningJob:MachineLearningJob';

    /**
     * Returns true if the given object is an instance of MachineLearningJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MachineLearningJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MachineLearningJob.__pulumiType;
    }

    /**
     * An object representing the custom labels added on the forecast.
     */
    public readonly customLabels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
     */
    public readonly datasourceType!: pulumi.Output<string>;
    /**
     * The uid of the datasource to query.
     */
    public readonly datasourceUid!: pulumi.Output<string>;
    /**
     * A description of the job.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A list of holiday IDs or names to take into account when training the model.
     */
    public readonly holidays!: pulumi.Output<string[] | undefined>;
    /**
     * The hyperparameters used to fine tune the algorithm. See
     * https://grafana.com/docs/grafana-cloud/alerting-and-irm/machine-learning/forecasts/models/ for the full list of
     * available hyperparameters.
     */
    public readonly hyperParams!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The data interval in seconds to train the data on.
     */
    public readonly interval!: pulumi.Output<number | undefined>;
    /**
     * The metric used to query the job results.
     */
    public readonly metric!: pulumi.Output<string>;
    /**
     * The name of the job.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An object representing the query params to query Grafana with.
     */
    public readonly queryParams!: pulumi.Output<{[key: string]: string}>;
    /**
     * The data interval in seconds to train the data on.
     */
    public readonly trainingWindow!: pulumi.Output<number | undefined>;

    /**
     * Create a MachineLearningJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated grafana.index/machinelearningjob.MachineLearningJob has been deprecated in favor of grafana.machinelearning/job.Job */
    constructor(name: string, args: MachineLearningJobArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated grafana.index/machinelearningjob.MachineLearningJob has been deprecated in favor of grafana.machinelearning/job.Job */
    constructor(name: string, argsOrState?: MachineLearningJobArgs | MachineLearningJobState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("MachineLearningJob is deprecated: grafana.index/machinelearningjob.MachineLearningJob has been deprecated in favor of grafana.machinelearning/job.Job")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MachineLearningJobState | undefined;
            resourceInputs["customLabels"] = state ? state.customLabels : undefined;
            resourceInputs["datasourceType"] = state ? state.datasourceType : undefined;
            resourceInputs["datasourceUid"] = state ? state.datasourceUid : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["holidays"] = state ? state.holidays : undefined;
            resourceInputs["hyperParams"] = state ? state.hyperParams : undefined;
            resourceInputs["interval"] = state ? state.interval : undefined;
            resourceInputs["metric"] = state ? state.metric : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["queryParams"] = state ? state.queryParams : undefined;
            resourceInputs["trainingWindow"] = state ? state.trainingWindow : undefined;
        } else {
            const args = argsOrState as MachineLearningJobArgs | undefined;
            if ((!args || args.datasourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasourceType'");
            }
            if ((!args || args.datasourceUid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasourceUid'");
            }
            if ((!args || args.metric === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metric'");
            }
            if ((!args || args.queryParams === undefined) && !opts.urn) {
                throw new Error("Missing required property 'queryParams'");
            }
            resourceInputs["customLabels"] = args ? args.customLabels : undefined;
            resourceInputs["datasourceType"] = args ? args.datasourceType : undefined;
            resourceInputs["datasourceUid"] = args ? args.datasourceUid : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["holidays"] = args ? args.holidays : undefined;
            resourceInputs["hyperParams"] = args ? args.hyperParams : undefined;
            resourceInputs["interval"] = args ? args.interval : undefined;
            resourceInputs["metric"] = args ? args.metric : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["queryParams"] = args ? args.queryParams : undefined;
            resourceInputs["trainingWindow"] = args ? args.trainingWindow : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MachineLearningJob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MachineLearningJob resources.
 */
export interface MachineLearningJobState {
    /**
     * An object representing the custom labels added on the forecast.
     */
    customLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
     */
    datasourceType?: pulumi.Input<string>;
    /**
     * The uid of the datasource to query.
     */
    datasourceUid?: pulumi.Input<string>;
    /**
     * A description of the job.
     */
    description?: pulumi.Input<string>;
    /**
     * A list of holiday IDs or names to take into account when training the model.
     */
    holidays?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The hyperparameters used to fine tune the algorithm. See
     * https://grafana.com/docs/grafana-cloud/alerting-and-irm/machine-learning/forecasts/models/ for the full list of
     * available hyperparameters.
     */
    hyperParams?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The data interval in seconds to train the data on.
     */
    interval?: pulumi.Input<number>;
    /**
     * The metric used to query the job results.
     */
    metric?: pulumi.Input<string>;
    /**
     * The name of the job.
     */
    name?: pulumi.Input<string>;
    /**
     * An object representing the query params to query Grafana with.
     */
    queryParams?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The data interval in seconds to train the data on.
     */
    trainingWindow?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a MachineLearningJob resource.
 */
export interface MachineLearningJobArgs {
    /**
     * An object representing the custom labels added on the forecast.
     */
    customLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
     */
    datasourceType: pulumi.Input<string>;
    /**
     * The uid of the datasource to query.
     */
    datasourceUid: pulumi.Input<string>;
    /**
     * A description of the job.
     */
    description?: pulumi.Input<string>;
    /**
     * A list of holiday IDs or names to take into account when training the model.
     */
    holidays?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The hyperparameters used to fine tune the algorithm. See
     * https://grafana.com/docs/grafana-cloud/alerting-and-irm/machine-learning/forecasts/models/ for the full list of
     * available hyperparameters.
     */
    hyperParams?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The data interval in seconds to train the data on.
     */
    interval?: pulumi.Input<number>;
    /**
     * The metric used to query the job results.
     */
    metric: pulumi.Input<string>;
    /**
     * The name of the job.
     */
    name?: pulumi.Input<string>;
    /**
     * An object representing the query params to query Grafana with.
     */
    queryParams: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The data interval in seconds to train the data on.
     */
    trainingWindow?: pulumi.Input<number>;
}
