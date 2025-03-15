// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages Grafana Fleet Management pipelines.
 *
 * * [Official documentation](https://grafana.com/docs/grafana-cloud/send-data/fleet-management/)
 * * [API documentation](https://grafana.com/docs/grafana-cloud/send-data/fleet-management/api-reference/pipeline-api/)
 *
 * Required access policy scopes:
 *
 * * fleet-management:read
 * * fleet-management:write
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:fleetManagement/pipeline:Pipeline name "{{ name }}"
 * ```
 */
export class Pipeline extends pulumi.CustomResource {
    /**
     * Get an existing Pipeline resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PipelineState, opts?: pulumi.CustomResourceOptions): Pipeline {
        return new Pipeline(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:fleetManagement/pipeline:Pipeline';

    /**
     * Returns true if the given object is an instance of Pipeline.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pipeline {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pipeline.__pulumiType;
    }

    /**
     * Configuration contents of the pipeline to be used by collectors
     */
    public readonly contents!: pulumi.Output<string>;
    /**
     * Whether the pipeline is enabled for collectors
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * Used to match against collectors and assign pipelines to them; follows the syntax of Prometheus Alertmanager matchers
     */
    public readonly matchers!: pulumi.Output<string[]>;
    /**
     * Name of the pipeline which is the unique identifier for the pipeline
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Pipeline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PipelineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PipelineArgs | PipelineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PipelineState | undefined;
            resourceInputs["contents"] = state ? state.contents : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["matchers"] = state ? state.matchers : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as PipelineArgs | undefined;
            if ((!args || args.contents === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contents'");
            }
            resourceInputs["contents"] = args ? args.contents : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["matchers"] = args ? args.matchers : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Pipeline.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pipeline resources.
 */
export interface PipelineState {
    /**
     * Configuration contents of the pipeline to be used by collectors
     */
    contents?: pulumi.Input<string>;
    /**
     * Whether the pipeline is enabled for collectors
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Used to match against collectors and assign pipelines to them; follows the syntax of Prometheus Alertmanager matchers
     */
    matchers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the pipeline which is the unique identifier for the pipeline
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Pipeline resource.
 */
export interface PipelineArgs {
    /**
     * Configuration contents of the pipeline to be used by collectors
     */
    contents: pulumi.Input<string>;
    /**
     * Whether the pipeline is enabled for collectors
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Used to match against collectors and assign pipelines to them; follows the syntax of Prometheus Alertmanager matchers
     */
    matchers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the pipeline which is the unique identifier for the pipeline
     */
    name?: pulumi.Input<string>;
}
