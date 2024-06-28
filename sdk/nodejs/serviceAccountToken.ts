// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Note:** This resource is available only with Grafana 9.1+.
 *
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const admin = new grafana.oss.ServiceAccount("admin", {
 *     isDisabled: false,
 *     role: "Admin",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:index/serviceAccountToken:ServiceAccountToken name "{{ id }}"
 * ```
 *
 * ```sh
 * $ pulumi import grafana:index/serviceAccountToken:ServiceAccountToken name "{{ orgID }}:{{ id }}"
 * ```
 *
 * @deprecated grafana.index/serviceaccounttoken.ServiceAccountToken has been deprecated in favor of grafana.oss/serviceaccounttoken.ServiceAccountToken
 */
export class ServiceAccountToken extends pulumi.CustomResource {
    /**
     * Get an existing ServiceAccountToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceAccountTokenState, opts?: pulumi.CustomResourceOptions): ServiceAccountToken {
        pulumi.log.warn("ServiceAccountToken is deprecated: grafana.index/serviceaccounttoken.ServiceAccountToken has been deprecated in favor of grafana.oss/serviceaccounttoken.ServiceAccountToken")
        return new ServiceAccountToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/serviceAccountToken:ServiceAccountToken';

    /**
     * Returns true if the given object is an instance of ServiceAccountToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceAccountToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceAccountToken.__pulumiType;
    }

    /**
     * The expiration date of the service account token.
     */
    public /*out*/ readonly expiration!: pulumi.Output<string>;
    /**
     * The status of the service account token.
     */
    public /*out*/ readonly hasExpired!: pulumi.Output<boolean>;
    /**
     * The key of the service account token.
     */
    public /*out*/ readonly key!: pulumi.Output<string>;
    /**
     * The name of the service account.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The key expiration in seconds. It is optional. If it is a positive number an expiration date for the key is set. If it
     * is null, zero or is omitted completely (unless `apiKeyMaxSecondsToLive` configuration option is set) the key will never
     * expire.
     */
    public readonly secondsToLive!: pulumi.Output<number | undefined>;
    /**
     * The ID of the service account to which the token belongs.
     */
    public readonly serviceAccountId!: pulumi.Output<string>;

    /**
     * Create a ServiceAccountToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated grafana.index/serviceaccounttoken.ServiceAccountToken has been deprecated in favor of grafana.oss/serviceaccounttoken.ServiceAccountToken */
    constructor(name: string, args: ServiceAccountTokenArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated grafana.index/serviceaccounttoken.ServiceAccountToken has been deprecated in favor of grafana.oss/serviceaccounttoken.ServiceAccountToken */
    constructor(name: string, argsOrState?: ServiceAccountTokenArgs | ServiceAccountTokenState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ServiceAccountToken is deprecated: grafana.index/serviceaccounttoken.ServiceAccountToken has been deprecated in favor of grafana.oss/serviceaccounttoken.ServiceAccountToken")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceAccountTokenState | undefined;
            resourceInputs["expiration"] = state ? state.expiration : undefined;
            resourceInputs["hasExpired"] = state ? state.hasExpired : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["secondsToLive"] = state ? state.secondsToLive : undefined;
            resourceInputs["serviceAccountId"] = state ? state.serviceAccountId : undefined;
        } else {
            const args = argsOrState as ServiceAccountTokenArgs | undefined;
            if ((!args || args.serviceAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceAccountId'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["secondsToLive"] = args ? args.secondsToLive : undefined;
            resourceInputs["serviceAccountId"] = args ? args.serviceAccountId : undefined;
            resourceInputs["expiration"] = undefined /*out*/;
            resourceInputs["hasExpired"] = undefined /*out*/;
            resourceInputs["key"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "grafana:index/serviceAccountToken:ServiceAccountToken" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        const secretOpts = { additionalSecretOutputs: ["key"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServiceAccountToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceAccountToken resources.
 */
export interface ServiceAccountTokenState {
    /**
     * The expiration date of the service account token.
     */
    expiration?: pulumi.Input<string>;
    /**
     * The status of the service account token.
     */
    hasExpired?: pulumi.Input<boolean>;
    /**
     * The key of the service account token.
     */
    key?: pulumi.Input<string>;
    /**
     * The name of the service account.
     */
    name?: pulumi.Input<string>;
    /**
     * The key expiration in seconds. It is optional. If it is a positive number an expiration date for the key is set. If it
     * is null, zero or is omitted completely (unless `apiKeyMaxSecondsToLive` configuration option is set) the key will never
     * expire.
     */
    secondsToLive?: pulumi.Input<number>;
    /**
     * The ID of the service account to which the token belongs.
     */
    serviceAccountId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceAccountToken resource.
 */
export interface ServiceAccountTokenArgs {
    /**
     * The name of the service account.
     */
    name?: pulumi.Input<string>;
    /**
     * The key expiration in seconds. It is optional. If it is a positive number an expiration date for the key is set. If it
     * is null, zero or is omitted completely (unless `apiKeyMaxSecondsToLive` configuration option is set) the key will never
     * expire.
     */
    secondsToLive?: pulumi.Input<number>;
    /**
     * The ID of the service account to which the token belongs.
     */
    serviceAccountId: pulumi.Input<string>;
}
