// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:frontendObservability/app:App name "{{ stack_id }}:{{ name }}"
 * ```
 */
export class App extends pulumi.CustomResource {
    /**
     * Get an existing App resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppState, opts?: pulumi.CustomResourceOptions): App {
        return new App(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:frontendObservability/app:App';

    /**
     * Returns true if the given object is an instance of App.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is App {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === App.__pulumiType;
    }

    /**
     * A list of allowed origins for CORS.
     */
    public readonly allowedOrigins!: pulumi.Output<string[]>;
    /**
     * The extra attributes to append in each signal.
     */
    public readonly extraLogAttributes!: pulumi.Output<{[key: string]: string}>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The key-value settings of the Frontend Observability app. Available Settings: `{combineLabData=(0|1)}`
     */
    public readonly settings!: pulumi.Output<{[key: string]: string}>;
    public readonly stackId!: pulumi.Output<number>;

    /**
     * Create a App resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppArgs | AppState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppState | undefined;
            resourceInputs["allowedOrigins"] = state ? state.allowedOrigins : undefined;
            resourceInputs["extraLogAttributes"] = state ? state.extraLogAttributes : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["settings"] = state ? state.settings : undefined;
            resourceInputs["stackId"] = state ? state.stackId : undefined;
        } else {
            const args = argsOrState as AppArgs | undefined;
            if ((!args || args.allowedOrigins === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowedOrigins'");
            }
            if ((!args || args.extraLogAttributes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'extraLogAttributes'");
            }
            if ((!args || args.settings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'settings'");
            }
            if ((!args || args.stackId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stackId'");
            }
            resourceInputs["allowedOrigins"] = args ? args.allowedOrigins : undefined;
            resourceInputs["extraLogAttributes"] = args ? args.extraLogAttributes : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["settings"] = args ? args.settings : undefined;
            resourceInputs["stackId"] = args ? args.stackId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(App.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering App resources.
 */
export interface AppState {
    /**
     * A list of allowed origins for CORS.
     */
    allowedOrigins?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The extra attributes to append in each signal.
     */
    extraLogAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    name?: pulumi.Input<string>;
    /**
     * The key-value settings of the Frontend Observability app. Available Settings: `{combineLabData=(0|1)}`
     */
    settings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    stackId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a App resource.
 */
export interface AppArgs {
    /**
     * A list of allowed origins for CORS.
     */
    allowedOrigins: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The extra attributes to append in each signal.
     */
    extraLogAttributes: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    name?: pulumi.Input<string>;
    /**
     * The key-value settings of the Frontend Observability app. Available Settings: `{combineLabData=(0|1)}`
     */
    settings: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    stackId: pulumi.Input<number>;
}
