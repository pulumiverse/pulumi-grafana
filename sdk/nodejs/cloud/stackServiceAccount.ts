// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages service accounts of a Grafana Cloud stack using the Cloud API
 * This can be used to bootstrap a management service account for a new stack
 *
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api)
 *
 * Required access policy scopes:
 *
 * * stacks:read
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
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:cloud/stackServiceAccount:StackServiceAccount name "{{ stackSlug }}:{{ serviceAccountID }}"
 * ```
 */
export class StackServiceAccount extends pulumi.CustomResource {
    /**
     * Get an existing StackServiceAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StackServiceAccountState, opts?: pulumi.CustomResourceOptions): StackServiceAccount {
        return new StackServiceAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:cloud/stackServiceAccount:StackServiceAccount';

    /**
     * Returns true if the given object is an instance of StackServiceAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StackServiceAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StackServiceAccount.__pulumiType;
    }

    /**
     * The disabled status for the service account. Defaults to `false`.
     */
    public readonly isDisabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the service account.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The basic role of the service account in the organization.
     */
    public readonly role!: pulumi.Output<string>;
    public readonly stackSlug!: pulumi.Output<string>;

    /**
     * Create a StackServiceAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StackServiceAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StackServiceAccountArgs | StackServiceAccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StackServiceAccountState | undefined;
            resourceInputs["isDisabled"] = state ? state.isDisabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["stackSlug"] = state ? state.stackSlug : undefined;
        } else {
            const args = argsOrState as StackServiceAccountArgs | undefined;
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            if ((!args || args.stackSlug === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stackSlug'");
            }
            resourceInputs["isDisabled"] = args ? args.isDisabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["stackSlug"] = args ? args.stackSlug : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "grafana:index/cloudStackServiceAccount:CloudStackServiceAccount" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(StackServiceAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StackServiceAccount resources.
 */
export interface StackServiceAccountState {
    /**
     * The disabled status for the service account. Defaults to `false`.
     */
    isDisabled?: pulumi.Input<boolean>;
    /**
     * The name of the service account.
     */
    name?: pulumi.Input<string>;
    /**
     * The basic role of the service account in the organization.
     */
    role?: pulumi.Input<string>;
    stackSlug?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StackServiceAccount resource.
 */
export interface StackServiceAccountArgs {
    /**
     * The disabled status for the service account. Defaults to `false`.
     */
    isDisabled?: pulumi.Input<boolean>;
    /**
     * The name of the service account.
     */
    name?: pulumi.Input<string>;
    /**
     * The basic role of the service account in the organization.
     */
    role: pulumi.Input<string>;
    stackSlug: pulumi.Input<string>;
}
