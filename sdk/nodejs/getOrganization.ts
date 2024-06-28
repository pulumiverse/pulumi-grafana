// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/organization-management/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/org/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = new grafana.oss.Organization("test", {
 *     adminUser: "admin",
 *     createUsers: true,
 *     viewers: [
 *         "viewer-01@example.com",
 *         "viewer-02@example.com",
 *     ],
 * });
 * const fromName = grafana.oss.getOrganizationOutput({
 *     name: test.name,
 * });
 * ```
 */
/** @deprecated grafana.index/getorganization.getOrganization has been deprecated in favor of grafana.oss/getorganization.getOrganization */
export function getOrganization(args: GetOrganizationArgs, opts?: pulumi.InvokeOptions): Promise<GetOrganizationResult> {
    pulumi.log.warn("getOrganization is deprecated: grafana.index/getorganization.getOrganization has been deprecated in favor of grafana.oss/getorganization.getOrganization")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:index/getOrganization:getOrganization", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrganization.
 */
export interface GetOrganizationArgs {
    /**
     * The name of the Organization.
     */
    name: string;
}

/**
 * A collection of values returned by getOrganization.
 */
export interface GetOrganizationResult {
    /**
     * A list of email addresses corresponding to users given admin access to the organization.
     */
    readonly admins: string[];
    /**
     * A list of email addresses corresponding to users given editor access to the organization.
     */
    readonly editors: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the Organization.
     */
    readonly name: string;
    /**
     * A list of email addresses corresponding to users given viewer access to the organization.
     */
    readonly viewers: string[];
}
/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/organization-management/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/org/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = new grafana.oss.Organization("test", {
 *     adminUser: "admin",
 *     createUsers: true,
 *     viewers: [
 *         "viewer-01@example.com",
 *         "viewer-02@example.com",
 *     ],
 * });
 * const fromName = grafana.oss.getOrganizationOutput({
 *     name: test.name,
 * });
 * ```
 */
/** @deprecated grafana.index/getorganization.getOrganization has been deprecated in favor of grafana.oss/getorganization.getOrganization */
export function getOrganizationOutput(args: GetOrganizationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrganizationResult> {
    return pulumi.output(args).apply((a: any) => getOrganization(a, opts))
}

/**
 * A collection of arguments for invoking getOrganization.
 */
export interface GetOrganizationOutputArgs {
    /**
     * The name of the Organization.
     */
    name: pulumi.Input<string>;
}
