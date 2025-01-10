// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = new grafana.cloud.ProviderAzureCredential("test", {
 *     stackId: "1",
 *     name: "test-name",
 *     clientId: "my-client-id",
 *     clientSecret: "my-client-secret",
 *     tenantId: "my-tenant-id",
 *     resourceDiscoveryTagFilters: [
 *         {
 *             key: "key-1",
 *             value: "value-1",
 *         },
 *         {
 *             key: "key-2",
 *             value: "value-2",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:cloud/providerAzureCredential:ProviderAzureCredential name "{{ stack_id }}:{{ resource_id }}"
 * ```
 */
export class ProviderAzureCredential extends pulumi.CustomResource {
    /**
     * Get an existing ProviderAzureCredential resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProviderAzureCredentialState, opts?: pulumi.CustomResourceOptions): ProviderAzureCredential {
        return new ProviderAzureCredential(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:cloud/providerAzureCredential:ProviderAzureCredential';

    /**
     * Returns true if the given object is an instance of ProviderAzureCredential.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProviderAzureCredential {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProviderAzureCredential.__pulumiType;
    }

    /**
     * The client ID of the Azure Credential.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * The client secret of the Azure Credential.
     */
    public readonly clientSecret!: pulumi.Output<string>;
    /**
     * The name of the Azure Credential.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The list of tag filters to apply to resources.
     */
    public readonly resourceDiscoveryTagFilters!: pulumi.Output<outputs.cloud.ProviderAzureCredentialResourceDiscoveryTagFilter[] | undefined>;
    /**
     * The ID given by the Grafana Cloud Provider API to this AWS Account resource.
     */
    public /*out*/ readonly resourceId!: pulumi.Output<string>;
    public readonly stackId!: pulumi.Output<string>;
    /**
     * The tenant ID of the Azure Credential.
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a ProviderAzureCredential resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderAzureCredentialArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProviderAzureCredentialArgs | ProviderAzureCredentialState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProviderAzureCredentialState | undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["resourceDiscoveryTagFilters"] = state ? state.resourceDiscoveryTagFilters : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["stackId"] = state ? state.stackId : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as ProviderAzureCredentialArgs | undefined;
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.clientSecret === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientSecret'");
            }
            if ((!args || args.stackId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stackId'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSecret"] = args?.clientSecret ? pulumi.secret(args.clientSecret) : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceDiscoveryTagFilters"] = args ? args.resourceDiscoveryTagFilters : undefined;
            resourceInputs["stackId"] = args ? args.stackId : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["resourceId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["clientSecret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ProviderAzureCredential.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProviderAzureCredential resources.
 */
export interface ProviderAzureCredentialState {
    /**
     * The client ID of the Azure Credential.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The client secret of the Azure Credential.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * The name of the Azure Credential.
     */
    name?: pulumi.Input<string>;
    /**
     * The list of tag filters to apply to resources.
     */
    resourceDiscoveryTagFilters?: pulumi.Input<pulumi.Input<inputs.cloud.ProviderAzureCredentialResourceDiscoveryTagFilter>[]>;
    /**
     * The ID given by the Grafana Cloud Provider API to this AWS Account resource.
     */
    resourceId?: pulumi.Input<string>;
    stackId?: pulumi.Input<string>;
    /**
     * The tenant ID of the Azure Credential.
     */
    tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProviderAzureCredential resource.
 */
export interface ProviderAzureCredentialArgs {
    /**
     * The client ID of the Azure Credential.
     */
    clientId: pulumi.Input<string>;
    /**
     * The client secret of the Azure Credential.
     */
    clientSecret: pulumi.Input<string>;
    /**
     * The name of the Azure Credential.
     */
    name?: pulumi.Input<string>;
    /**
     * The list of tag filters to apply to resources.
     */
    resourceDiscoveryTagFilters?: pulumi.Input<pulumi.Input<inputs.cloud.ProviderAzureCredentialResourceDiscoveryTagFilter>[]>;
    stackId: pulumi.Input<string>;
    /**
     * The tenant ID of the Azure Credential.
     */
    tenantId: pulumi.Input<string>;
}
