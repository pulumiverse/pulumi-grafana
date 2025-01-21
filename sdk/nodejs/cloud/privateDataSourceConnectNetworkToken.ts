// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * * [Official documentation](https://grafana.com/docs/grafana-cloud/connect-externally-hosted/private-data-source-connect/)
 * * [API documentation](https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#create-a-token)
 *
 * Required access policy scopes:
 *
 * * accesspolicies:read
 * * accesspolicies:write
 * * accesspolicies:delete
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const current = grafana.cloud.getStack({
 *     slug: "<your slug>",
 * });
 * const test = new grafana.cloud.PrivateDataSourceConnectNetwork("test", {
 *     region: "us",
 *     name: "my-pdc",
 *     displayName: "My PDC",
 *     stackIdentifier: current.then(current => current.id),
 * });
 * const testPrivateDataSourceConnectNetworkToken = new grafana.cloud.PrivateDataSourceConnectNetworkToken("test", {
 *     pdcNetworkId: test.pdcNetworkId,
 *     region: test.region,
 *     name: "my-pdc-token",
 *     displayName: "My PDC Token",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:cloud/privateDataSourceConnectNetworkToken:PrivateDataSourceConnectNetworkToken name "{{ region }}:{{ tokenId }}"
 * ```
 */
export class PrivateDataSourceConnectNetworkToken extends pulumi.CustomResource {
    /**
     * Get an existing PrivateDataSourceConnectNetworkToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivateDataSourceConnectNetworkTokenState, opts?: pulumi.CustomResourceOptions): PrivateDataSourceConnectNetworkToken {
        return new PrivateDataSourceConnectNetworkToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:cloud/privateDataSourceConnectNetworkToken:PrivateDataSourceConnectNetworkToken';

    /**
     * Returns true if the given object is an instance of PrivateDataSourceConnectNetworkToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateDataSourceConnectNetworkToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateDataSourceConnectNetworkToken.__pulumiType;
    }

    /**
     * Creation date of the private data source network token.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Display name of the private data source network token. Defaults to the name.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Expiration date of the private data source network token. Does not expire by default.
     */
    public readonly expiresAt!: pulumi.Output<string | undefined>;
    /**
     * Name of the private data source network token.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the private data source network for which to create a token.
     */
    public readonly pdcNetworkId!: pulumi.Output<string>;
    /**
     * Region of the private data source network. Should be set to the same region as the private data source network. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
     */
    public readonly region!: pulumi.Output<string>;
    public /*out*/ readonly token!: pulumi.Output<string>;
    /**
     * Last update date of the private data source network token.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a PrivateDataSourceConnectNetworkToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateDataSourceConnectNetworkTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivateDataSourceConnectNetworkTokenArgs | PrivateDataSourceConnectNetworkTokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrivateDataSourceConnectNetworkTokenState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pdcNetworkId"] = state ? state.pdcNetworkId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as PrivateDataSourceConnectNetworkTokenArgs | undefined;
            if ((!args || args.pdcNetworkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pdcNetworkId'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["expiresAt"] = args ? args.expiresAt : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pdcNetworkId"] = args ? args.pdcNetworkId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(PrivateDataSourceConnectNetworkToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivateDataSourceConnectNetworkToken resources.
 */
export interface PrivateDataSourceConnectNetworkTokenState {
    /**
     * Creation date of the private data source network token.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Display name of the private data source network token. Defaults to the name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Expiration date of the private data source network token. Does not expire by default.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * Name of the private data source network token.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the private data source network for which to create a token.
     */
    pdcNetworkId?: pulumi.Input<string>;
    /**
     * Region of the private data source network. Should be set to the same region as the private data source network. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
     */
    region?: pulumi.Input<string>;
    token?: pulumi.Input<string>;
    /**
     * Last update date of the private data source network token.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PrivateDataSourceConnectNetworkToken resource.
 */
export interface PrivateDataSourceConnectNetworkTokenArgs {
    /**
     * Display name of the private data source network token. Defaults to the name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Expiration date of the private data source network token. Does not expire by default.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * Name of the private data source network token.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the private data source network for which to create a token.
     */
    pdcNetworkId: pulumi.Input<string>;
    /**
     * Region of the private data source network. Should be set to the same region as the private data source network. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
     */
    region: pulumi.Input<string>;
}
