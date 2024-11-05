// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages service account tokens of a Grafana Cloud stack using the Cloud API
 * This can be used to bootstrap a management service account token for a new stack
 *
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api)
 *
 * Required access policy scopes:
 *
 * * stack-service-accounts:write
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const cloudSa = new grafana.cloud.StackServiceAccount("cloud_sa", {
 *     stackSlug: "<your stack slug>",
 *     name: "cloud service account",
 *     role: "Admin",
 *     isDisabled: false,
 * });
 * const foo = new grafana.cloud.StackServiceAccountToken("foo", {
 *     name: "key_foo",
 *     serviceAccountId: cloudSa.id,
 * });
 * export const serviceAccountTokenFooKey = foo.key;
 * ```
 *
 * @deprecated grafana.index/cloudstackserviceaccounttoken.CloudStackServiceAccountToken has been deprecated in favor of grafana.cloud/stackserviceaccounttoken.StackServiceAccountToken
 */
export class CloudStackServiceAccountToken extends pulumi.CustomResource {
    /**
     * Get an existing CloudStackServiceAccountToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CloudStackServiceAccountTokenState, opts?: pulumi.CustomResourceOptions): CloudStackServiceAccountToken {
        pulumi.log.warn("CloudStackServiceAccountToken is deprecated: grafana.index/cloudstackserviceaccounttoken.CloudStackServiceAccountToken has been deprecated in favor of grafana.cloud/stackserviceaccounttoken.StackServiceAccountToken")
        return new CloudStackServiceAccountToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/cloudStackServiceAccountToken:CloudStackServiceAccountToken';

    /**
     * Returns true if the given object is an instance of CloudStackServiceAccountToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudStackServiceAccountToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudStackServiceAccountToken.__pulumiType;
    }

    public /*out*/ readonly expiration!: pulumi.Output<string>;
    public /*out*/ readonly hasExpired!: pulumi.Output<boolean>;
    public /*out*/ readonly key!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly secondsToLive!: pulumi.Output<number | undefined>;
    public readonly serviceAccountId!: pulumi.Output<string>;
    public readonly stackSlug!: pulumi.Output<string>;

    /**
     * Create a CloudStackServiceAccountToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated grafana.index/cloudstackserviceaccounttoken.CloudStackServiceAccountToken has been deprecated in favor of grafana.cloud/stackserviceaccounttoken.StackServiceAccountToken */
    constructor(name: string, args: CloudStackServiceAccountTokenArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated grafana.index/cloudstackserviceaccounttoken.CloudStackServiceAccountToken has been deprecated in favor of grafana.cloud/stackserviceaccounttoken.StackServiceAccountToken */
    constructor(name: string, argsOrState?: CloudStackServiceAccountTokenArgs | CloudStackServiceAccountTokenState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("CloudStackServiceAccountToken is deprecated: grafana.index/cloudstackserviceaccounttoken.CloudStackServiceAccountToken has been deprecated in favor of grafana.cloud/stackserviceaccounttoken.StackServiceAccountToken")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CloudStackServiceAccountTokenState | undefined;
            resourceInputs["expiration"] = state ? state.expiration : undefined;
            resourceInputs["hasExpired"] = state ? state.hasExpired : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["secondsToLive"] = state ? state.secondsToLive : undefined;
            resourceInputs["serviceAccountId"] = state ? state.serviceAccountId : undefined;
            resourceInputs["stackSlug"] = state ? state.stackSlug : undefined;
        } else {
            const args = argsOrState as CloudStackServiceAccountTokenArgs | undefined;
            if ((!args || args.serviceAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceAccountId'");
            }
            if ((!args || args.stackSlug === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stackSlug'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["secondsToLive"] = args ? args.secondsToLive : undefined;
            resourceInputs["serviceAccountId"] = args ? args.serviceAccountId : undefined;
            resourceInputs["stackSlug"] = args ? args.stackSlug : undefined;
            resourceInputs["expiration"] = undefined /*out*/;
            resourceInputs["hasExpired"] = undefined /*out*/;
            resourceInputs["key"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["key"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(CloudStackServiceAccountToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CloudStackServiceAccountToken resources.
 */
export interface CloudStackServiceAccountTokenState {
    expiration?: pulumi.Input<string>;
    hasExpired?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    secondsToLive?: pulumi.Input<number>;
    serviceAccountId?: pulumi.Input<string>;
    stackSlug?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CloudStackServiceAccountToken resource.
 */
export interface CloudStackServiceAccountTokenArgs {
    name?: pulumi.Input<string>;
    secondsToLive?: pulumi.Input<number>;
    serviceAccountId: pulumi.Input<string>;
    stackSlug: pulumi.Input<string>;
}
