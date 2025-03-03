// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Sets up Synthetic Monitoring on a Grafana cloud stack and generates a token.
 * Once a Grafana Cloud stack is created, a user can either use this resource or go into the UI to install synthetic monitoring.
 * This resource cannot be imported but it can be used on an existing Synthetic Monitoring installation without issues.
 *
 * **Note that this resource must be used on a provider configured with Grafana Cloud credentials.**
 *
 * * [Official documentation](https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/)
 * * [API documentation](https://github.com/grafana/synthetic-monitoring-api-go-client/blob/main/docs/API.md#apiv1registerinstall)
 *
 * Required access policy scopes:
 *
 * * stacks:read
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const config = new pulumi.Config();
 * // Cloud Access Policy token for Grafana Cloud with the following scopes: accesspolicies:read|write|delete, stacks:read|write|delete
 * const cloudAccessPolicyToken = config.requireObject("cloudAccessPolicyToken");
 * const stackSlug = config.requireObject("stackSlug");
 * const cloudRegion = config.get("cloudRegion") || "us";
 * const smStack = new grafana.cloud.Stack("sm_stack", {
 *     name: stackSlug,
 *     slug: stackSlug,
 *     regionSlug: cloudRegion,
 * });
 * // Step 2: Install Synthetic Monitoring on the stack
 * const smMetricsPublish = new grafana.cloud.AccessPolicy("sm_metrics_publish", {
 *     region: cloudRegion,
 *     name: "metric-publisher-for-sm",
 *     scopes: [
 *         "metrics:write",
 *         "stacks:read",
 *         "logs:write",
 *         "traces:write",
 *     ],
 *     realms: [{
 *         type: "stack",
 *         identifier: smStack.id,
 *     }],
 * });
 * const smMetricsPublishAccessPolicyToken = new grafana.cloud.AccessPolicyToken("sm_metrics_publish", {
 *     region: cloudRegion,
 *     accessPolicyId: smMetricsPublish.policyId,
 *     name: "metric-publisher-for-sm",
 * });
 * const smStackInstallation = new grafana.syntheticmonitoring.Installation("sm_stack", {
 *     stackId: smStack.id,
 *     metricsPublisherKey: smMetricsPublishAccessPolicyToken.token,
 * });
 * const main = grafana.syntheticMonitoring.getProbes({});
 * ```
 */
export class Installation extends pulumi.CustomResource {
    /**
     * Get an existing Installation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstallationState, opts?: pulumi.CustomResourceOptions): Installation {
        return new Installation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:syntheticMonitoring/installation:Installation';

    /**
     * Returns true if the given object is an instance of Installation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Installation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Installation.__pulumiType;
    }

    /**
     * The [Grafana Cloud access policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and logs to Grafana Cloud stack.
     */
    public readonly metricsPublisherKey!: pulumi.Output<string>;
    /**
     * Generated token to access the SM API.
     */
    public /*out*/ readonly smAccessToken!: pulumi.Output<string>;
    /**
     * The ID or slug of the stack to install SM on.
     */
    public readonly stackId!: pulumi.Output<string>;
    /**
     * The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/set-up-private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
     */
    public readonly stackSmApiUrl!: pulumi.Output<string>;

    /**
     * Create a Installation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstallationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstallationArgs | InstallationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstallationState | undefined;
            resourceInputs["metricsPublisherKey"] = state ? state.metricsPublisherKey : undefined;
            resourceInputs["smAccessToken"] = state ? state.smAccessToken : undefined;
            resourceInputs["stackId"] = state ? state.stackId : undefined;
            resourceInputs["stackSmApiUrl"] = state ? state.stackSmApiUrl : undefined;
        } else {
            const args = argsOrState as InstallationArgs | undefined;
            if ((!args || args.metricsPublisherKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metricsPublisherKey'");
            }
            if ((!args || args.stackId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stackId'");
            }
            resourceInputs["metricsPublisherKey"] = args?.metricsPublisherKey ? pulumi.secret(args.metricsPublisherKey) : undefined;
            resourceInputs["stackId"] = args ? args.stackId : undefined;
            resourceInputs["stackSmApiUrl"] = args ? args.stackSmApiUrl : undefined;
            resourceInputs["smAccessToken"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["metricsPublisherKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Installation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Installation resources.
 */
export interface InstallationState {
    /**
     * The [Grafana Cloud access policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and logs to Grafana Cloud stack.
     */
    metricsPublisherKey?: pulumi.Input<string>;
    /**
     * Generated token to access the SM API.
     */
    smAccessToken?: pulumi.Input<string>;
    /**
     * The ID or slug of the stack to install SM on.
     */
    stackId?: pulumi.Input<string>;
    /**
     * The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/set-up-private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
     */
    stackSmApiUrl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Installation resource.
 */
export interface InstallationArgs {
    /**
     * The [Grafana Cloud access policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and logs to Grafana Cloud stack.
     */
    metricsPublisherKey: pulumi.Input<string>;
    /**
     * The ID or slug of the stack to install SM on.
     */
    stackId: pulumi.Input<string>;
    /**
     * The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/set-up-private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
     */
    stackSmApiUrl?: pulumi.Input<string>;
}
