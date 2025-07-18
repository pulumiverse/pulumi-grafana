// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages limits for a k6 project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const testProjectLimits = new grafana.k6.Project("test_project_limits", {name: "Terraform Project Test Limits"});
 * const testLimits = new grafana.k6.ProjectLimits("test_limits", {
 *     projectId: testProjectLimits.id,
 *     vuhMaxPerMonth: 1000,
 *     vuMaxPerTest: 100,
 *     vuBrowserMaxPerTest: 10,
 *     durationMaxPerTest: 3600,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:k6/projectLimits:ProjectLimits name "{{ project_id }}"
 * ```
 */
export class ProjectLimits extends pulumi.CustomResource {
    /**
     * Get an existing ProjectLimits resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectLimitsState, opts?: pulumi.CustomResourceOptions): ProjectLimits {
        return new ProjectLimits(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:k6/projectLimits:ProjectLimits';

    /**
     * Returns true if the given object is an instance of ProjectLimits.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectLimits {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectLimits.__pulumiType;
    }

    /**
     * Maximum duration of a test in seconds.
     */
    public readonly durationMaxPerTest!: pulumi.Output<number | undefined>;
    /**
     * The identifier of the project to manage limits for.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Maximum number of concurrent browser virtual users (VUs) used in one test.
     */
    public readonly vuBrowserMaxPerTest!: pulumi.Output<number | undefined>;
    /**
     * Maximum number of concurrent virtual users (VUs) used in one test.
     */
    public readonly vuMaxPerTest!: pulumi.Output<number | undefined>;
    /**
     * Maximum amount of virtual user hours (VU/h) used per one calendar month.
     */
    public readonly vuhMaxPerMonth!: pulumi.Output<number | undefined>;

    /**
     * Create a ProjectLimits resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectLimitsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectLimitsArgs | ProjectLimitsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectLimitsState | undefined;
            resourceInputs["durationMaxPerTest"] = state ? state.durationMaxPerTest : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["vuBrowserMaxPerTest"] = state ? state.vuBrowserMaxPerTest : undefined;
            resourceInputs["vuMaxPerTest"] = state ? state.vuMaxPerTest : undefined;
            resourceInputs["vuhMaxPerMonth"] = state ? state.vuhMaxPerMonth : undefined;
        } else {
            const args = argsOrState as ProjectLimitsArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["durationMaxPerTest"] = args ? args.durationMaxPerTest : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["vuBrowserMaxPerTest"] = args ? args.vuBrowserMaxPerTest : undefined;
            resourceInputs["vuMaxPerTest"] = args ? args.vuMaxPerTest : undefined;
            resourceInputs["vuhMaxPerMonth"] = args ? args.vuhMaxPerMonth : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectLimits.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectLimits resources.
 */
export interface ProjectLimitsState {
    /**
     * Maximum duration of a test in seconds.
     */
    durationMaxPerTest?: pulumi.Input<number>;
    /**
     * The identifier of the project to manage limits for.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Maximum number of concurrent browser virtual users (VUs) used in one test.
     */
    vuBrowserMaxPerTest?: pulumi.Input<number>;
    /**
     * Maximum number of concurrent virtual users (VUs) used in one test.
     */
    vuMaxPerTest?: pulumi.Input<number>;
    /**
     * Maximum amount of virtual user hours (VU/h) used per one calendar month.
     */
    vuhMaxPerMonth?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ProjectLimits resource.
 */
export interface ProjectLimitsArgs {
    /**
     * Maximum duration of a test in seconds.
     */
    durationMaxPerTest?: pulumi.Input<number>;
    /**
     * The identifier of the project to manage limits for.
     */
    projectId: pulumi.Input<string>;
    /**
     * Maximum number of concurrent browser virtual users (VUs) used in one test.
     */
    vuBrowserMaxPerTest?: pulumi.Input<number>;
    /**
     * Maximum number of concurrent virtual users (VUs) used in one test.
     */
    vuMaxPerTest?: pulumi.Input<number>;
    /**
     * Maximum amount of virtual user hours (VU/h) used per one calendar month.
     */
    vuhMaxPerMonth?: pulumi.Input<number>;
}
