// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * **Note:** This resource is available only with Grafana Enterprise 8.+.
 *
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/roles-and-permissions/access-control/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/access_control/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = new grafana.enterprise.Role("test", {
 *     description: "test-role description",
 *     uid: "test-ds-role-uid",
 *     version: 1,
 *     global: true,
 *     hidden: false,
 *     permissions: [
 *         {
 *             action: "org.users:add",
 *             scope: "users:*",
 *         },
 *         {
 *             action: "org.users:write",
 *             scope: "users:*",
 *         },
 *         {
 *             action: "org.users:read",
 *             scope: "users:*",
 *         },
 *     ],
 * });
 * const fromName = grafana.enterprise.getRoleOutput({
 *     name: test.name,
 * });
 * ```
 */
/** @deprecated grafana.index/getrole.getRole has been deprecated in favor of grafana.enterprise/getrole.getRole */
export function getRole(args: GetRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetRoleResult> {
    pulumi.log.warn("getRole is deprecated: grafana.index/getrole.getRole has been deprecated in favor of grafana.enterprise/getrole.getRole")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:index/getRole:getRole", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRole.
 */
export interface GetRoleArgs {
    /**
     * Name of the role
     */
    name: string;
}

/**
 * A collection of values returned by getRole.
 */
export interface GetRoleResult {
    /**
     * Description of the role.
     */
    readonly description: string;
    /**
     * Display name of the role. Available with Grafana 8.5+.
     */
    readonly displayName: string;
    /**
     * Boolean to state whether the role is available across all organizations or not.
     */
    readonly global: boolean;
    /**
     * Group of the role. Available with Grafana 8.5+.
     */
    readonly group: string;
    /**
     * Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+.
     */
    readonly hidden: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name of the role
     */
    readonly name: string;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    readonly orgId: string;
    /**
     * Specific set of actions granted by the role.
     */
    readonly permissions: outputs.GetRolePermission[];
    /**
     * Unique identifier of the role. Used for assignments.
     */
    readonly uid: string;
    /**
     * Version of the role. A role is updated only on version increase. This field or `autoIncrementVersion` should be set.
     */
    readonly version: number;
}
/**
 * **Note:** This resource is available only with Grafana Enterprise 8.+.
 *
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/roles-and-permissions/access-control/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/access_control/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = new grafana.enterprise.Role("test", {
 *     description: "test-role description",
 *     uid: "test-ds-role-uid",
 *     version: 1,
 *     global: true,
 *     hidden: false,
 *     permissions: [
 *         {
 *             action: "org.users:add",
 *             scope: "users:*",
 *         },
 *         {
 *             action: "org.users:write",
 *             scope: "users:*",
 *         },
 *         {
 *             action: "org.users:read",
 *             scope: "users:*",
 *         },
 *     ],
 * });
 * const fromName = grafana.enterprise.getRoleOutput({
 *     name: test.name,
 * });
 * ```
 */
/** @deprecated grafana.index/getrole.getRole has been deprecated in favor of grafana.enterprise/getrole.getRole */
export function getRoleOutput(args: GetRoleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRoleResult> {
    return pulumi.output(args).apply((a: any) => getRole(a, opts))
}

/**
 * A collection of arguments for invoking getRole.
 */
export interface GetRoleOutputArgs {
    /**
     * Name of the role
     */
    name: pulumi.Input<string>;
}
