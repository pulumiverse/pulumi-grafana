// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages Grafana Cloud Plugin Installations.
 *
 * * [Plugin Catalog](https://grafana.com/grafana/plugins/)
 *
 * Required access policy scopes:
 *
 * * stack-plugins:read
 * * stack-plugins:write
 * * stack-plugins:delete
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = new grafana.cloud.PluginInstallation("test", {
 *     stackSlug: "stackname",
 *     slug: "some-plugin",
 *     version: "1.2.3",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:cloud/pluginInstallation:PluginInstallation name "{{ stackSlug }}:{{ pluginSlug }}"
 * ```
 */
export class PluginInstallation extends pulumi.CustomResource {
    /**
     * Get an existing PluginInstallation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PluginInstallationState, opts?: pulumi.CustomResourceOptions): PluginInstallation {
        return new PluginInstallation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:cloud/pluginInstallation:PluginInstallation';

    /**
     * Returns true if the given object is an instance of PluginInstallation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PluginInstallation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PluginInstallation.__pulumiType;
    }

    /**
     * Slug of the plugin to be installed.
     */
    public readonly slug!: pulumi.Output<string>;
    /**
     * The stack id to which the plugin should be installed.
     */
    public readonly stackSlug!: pulumi.Output<string>;
    /**
     * Version of the plugin to be installed.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a PluginInstallation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PluginInstallationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PluginInstallationArgs | PluginInstallationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PluginInstallationState | undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["stackSlug"] = state ? state.stackSlug : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as PluginInstallationArgs | undefined;
            if ((!args || args.slug === undefined) && !opts.urn) {
                throw new Error("Missing required property 'slug'");
            }
            if ((!args || args.stackSlug === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stackSlug'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            resourceInputs["slug"] = args ? args.slug : undefined;
            resourceInputs["stackSlug"] = args ? args.stackSlug : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "grafana:index/cloudPluginInstallation:CloudPluginInstallation" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(PluginInstallation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PluginInstallation resources.
 */
export interface PluginInstallationState {
    /**
     * Slug of the plugin to be installed.
     */
    slug?: pulumi.Input<string>;
    /**
     * The stack id to which the plugin should be installed.
     */
    stackSlug?: pulumi.Input<string>;
    /**
     * Version of the plugin to be installed.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PluginInstallation resource.
 */
export interface PluginInstallationArgs {
    /**
     * Slug of the plugin to be installed.
     */
    slug: pulumi.Input<string>;
    /**
     * The stack id to which the plugin should be installed.
     */
    stackSlug: pulumi.Input<string>;
    /**
     * Version of the plugin to be installed.
     */
    version: pulumi.Input<string>;
}
