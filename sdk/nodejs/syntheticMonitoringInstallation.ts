// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SyntheticMonitoringInstallation extends pulumi.CustomResource {
    /**
     * Get an existing SyntheticMonitoringInstallation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SyntheticMonitoringInstallationState, opts?: pulumi.CustomResourceOptions): SyntheticMonitoringInstallation {
        return new SyntheticMonitoringInstallation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/syntheticMonitoringInstallation:SyntheticMonitoringInstallation';

    /**
     * Returns true if the given object is an instance of SyntheticMonitoringInstallation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SyntheticMonitoringInstallation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SyntheticMonitoringInstallation.__pulumiType;
    }

    /**
     * The [Grafana Cloud access
     * policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with
     * the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and
     * logs to Grafana Cloud stack.
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
     * The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here:
     * https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/set-up/set-up-private-probes/#probe-api-server-url. A
     * static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region,
     * this field is computed automatically and readable.
     */
    public readonly stackSmApiUrl!: pulumi.Output<string>;

    /**
     * Create a SyntheticMonitoringInstallation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SyntheticMonitoringInstallationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SyntheticMonitoringInstallationArgs | SyntheticMonitoringInstallationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SyntheticMonitoringInstallationState | undefined;
            resourceInputs["metricsPublisherKey"] = state ? state.metricsPublisherKey : undefined;
            resourceInputs["smAccessToken"] = state ? state.smAccessToken : undefined;
            resourceInputs["stackId"] = state ? state.stackId : undefined;
            resourceInputs["stackSmApiUrl"] = state ? state.stackSmApiUrl : undefined;
        } else {
            const args = argsOrState as SyntheticMonitoringInstallationArgs | undefined;
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
        super(SyntheticMonitoringInstallation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SyntheticMonitoringInstallation resources.
 */
export interface SyntheticMonitoringInstallationState {
    /**
     * The [Grafana Cloud access
     * policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with
     * the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and
     * logs to Grafana Cloud stack.
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
     * The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here:
     * https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/set-up/set-up-private-probes/#probe-api-server-url. A
     * static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region,
     * this field is computed automatically and readable.
     */
    stackSmApiUrl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SyntheticMonitoringInstallation resource.
 */
export interface SyntheticMonitoringInstallationArgs {
    /**
     * The [Grafana Cloud access
     * policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with
     * the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and
     * logs to Grafana Cloud stack.
     */
    metricsPublisherKey: pulumi.Input<string>;
    /**
     * The ID or slug of the stack to install SM on.
     */
    stackId: pulumi.Input<string>;
    /**
     * The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here:
     * https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/set-up/set-up-private-probes/#probe-api-server-url. A
     * static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region,
     * this field is computed automatically and readable.
     */
    stackSmApiUrl?: pulumi.Input<string>;
}
