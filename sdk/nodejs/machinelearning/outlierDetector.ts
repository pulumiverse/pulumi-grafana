// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * An outlier detector monitors the results of a query and reports when its values are outside normal bands.
 *
 * The normal band is configured by choice of algorithm, its sensitivity and other configuration.
 *
 * Visit https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for more details.
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:machineLearning/outlierDetector:OutlierDetector name "{{ id }}"
 * ```
 */
export class OutlierDetector extends pulumi.CustomResource {
    /**
     * Get an existing OutlierDetector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OutlierDetectorState, opts?: pulumi.CustomResourceOptions): OutlierDetector {
        return new OutlierDetector(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:machineLearning/outlierDetector:OutlierDetector';

    /**
     * Returns true if the given object is an instance of OutlierDetector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OutlierDetector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OutlierDetector.__pulumiType;
    }

    /**
     * The algorithm to use and its configuration. See https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
     */
    public readonly algorithm!: pulumi.Output<outputs.machineLearning.OutlierDetectorAlgorithm>;
    /**
     * The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
     */
    public readonly datasourceType!: pulumi.Output<string>;
    /**
     * The uid of the datasource to query.
     */
    public readonly datasourceUid!: pulumi.Output<string>;
    /**
     * A description of the outlier detector.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The data interval in seconds to monitor. Defaults to `300`.
     */
    public readonly interval!: pulumi.Output<number | undefined>;
    /**
     * The metric used to query the outlier detector results.
     */
    public readonly metric!: pulumi.Output<string>;
    /**
     * The name of the outlier detector.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An object representing the query params to query Grafana with.
     */
    public readonly queryParams!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a OutlierDetector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OutlierDetectorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OutlierDetectorArgs | OutlierDetectorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OutlierDetectorState | undefined;
            resourceInputs["algorithm"] = state ? state.algorithm : undefined;
            resourceInputs["datasourceType"] = state ? state.datasourceType : undefined;
            resourceInputs["datasourceUid"] = state ? state.datasourceUid : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["interval"] = state ? state.interval : undefined;
            resourceInputs["metric"] = state ? state.metric : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["queryParams"] = state ? state.queryParams : undefined;
        } else {
            const args = argsOrState as OutlierDetectorArgs | undefined;
            if ((!args || args.algorithm === undefined) && !opts.urn) {
                throw new Error("Missing required property 'algorithm'");
            }
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
            resourceInputs["algorithm"] = args ? args.algorithm : undefined;
            resourceInputs["datasourceType"] = args ? args.datasourceType : undefined;
            resourceInputs["datasourceUid"] = args ? args.datasourceUid : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["interval"] = args ? args.interval : undefined;
            resourceInputs["metric"] = args ? args.metric : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["queryParams"] = args ? args.queryParams : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "grafana:index/machineLearningOutlierDetector:MachineLearningOutlierDetector" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(OutlierDetector.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OutlierDetector resources.
 */
export interface OutlierDetectorState {
    /**
     * The algorithm to use and its configuration. See https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
     */
    algorithm?: pulumi.Input<inputs.machineLearning.OutlierDetectorAlgorithm>;
    /**
     * The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
     */
    datasourceType?: pulumi.Input<string>;
    /**
     * The uid of the datasource to query.
     */
    datasourceUid?: pulumi.Input<string>;
    /**
     * A description of the outlier detector.
     */
    description?: pulumi.Input<string>;
    /**
     * The data interval in seconds to monitor. Defaults to `300`.
     */
    interval?: pulumi.Input<number>;
    /**
     * The metric used to query the outlier detector results.
     */
    metric?: pulumi.Input<string>;
    /**
     * The name of the outlier detector.
     */
    name?: pulumi.Input<string>;
    /**
     * An object representing the query params to query Grafana with.
     */
    queryParams?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a OutlierDetector resource.
 */
export interface OutlierDetectorArgs {
    /**
     * The algorithm to use and its configuration. See https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
     */
    algorithm: pulumi.Input<inputs.machineLearning.OutlierDetectorAlgorithm>;
    /**
     * The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
     */
    datasourceType: pulumi.Input<string>;
    /**
     * The uid of the datasource to query.
     */
    datasourceUid: pulumi.Input<string>;
    /**
     * A description of the outlier detector.
     */
    description?: pulumi.Input<string>;
    /**
     * The data interval in seconds to monitor. Defaults to `300`.
     */
    interval?: pulumi.Input<number>;
    /**
     * The metric used to query the outlier detector results.
     */
    metric: pulumi.Input<string>;
    /**
     * The name of the outlier detector.
     */
    name?: pulumi.Input<string>;
    /**
     * An object representing the query params to query Grafana with.
     */
    queryParams: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
