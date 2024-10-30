// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manages the entire set of permissions for a dashboard. Permissions that aren't specified when applying this resource will be removed.
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/roles-and-permissions/access-control/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/dashboard_permissions/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const team = new grafana.oss.Team("team", {name: "Team Name"});
 * const user = new grafana.oss.User("user", {
 *     email: "user.name@example.com",
 *     password: "my-password",
 *     login: "user.name",
 * });
 * const metrics = new grafana.oss.Dashboard("metrics", {configJson: JSON.stringify({
 *     title: "My Dashboard",
 *     uid: "my-dashboard-uid",
 * })});
 * const collectionPermission = new grafana.oss.DashboardPermission("collectionPermission", {
 *     dashboardUid: metrics.uid,
 *     permissions: [
 *         {
 *             role: "Editor",
 *             permission: "Edit",
 *         },
 *         {
 *             teamId: team.id,
 *             permission: "View",
 *         },
 *         {
 *             userId: user.id,
 *             permission: "Admin",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:index/dashboardPermission:DashboardPermission name "{{ dashboardUID }}"
 * ```
 *
 * ```sh
 * $ pulumi import grafana:index/dashboardPermission:DashboardPermission name "{{ orgID }}:{{ dashboardUID }}"
 * ```
 *
 * @deprecated grafana.index/dashboardpermission.DashboardPermission has been deprecated in favor of grafana.oss/dashboardpermission.DashboardPermission
 */
export class DashboardPermission extends pulumi.CustomResource {
    /**
     * Get an existing DashboardPermission resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DashboardPermissionState, opts?: pulumi.CustomResourceOptions): DashboardPermission {
        pulumi.log.warn("DashboardPermission is deprecated: grafana.index/dashboardpermission.DashboardPermission has been deprecated in favor of grafana.oss/dashboardpermission.DashboardPermission")
        return new DashboardPermission(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/dashboardPermission:DashboardPermission';

    /**
     * Returns true if the given object is an instance of DashboardPermission.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DashboardPermission {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DashboardPermission.__pulumiType;
    }

    /**
     * UID of the dashboard to apply permissions to.
     */
    public readonly dashboardUid!: pulumi.Output<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * The permission items to add/update. Items that are omitted from the list will be removed.
     */
    public readonly permissions!: pulumi.Output<outputs.DashboardPermissionPermission[] | undefined>;

    /**
     * Create a DashboardPermission resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated grafana.index/dashboardpermission.DashboardPermission has been deprecated in favor of grafana.oss/dashboardpermission.DashboardPermission */
    constructor(name: string, args?: DashboardPermissionArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated grafana.index/dashboardpermission.DashboardPermission has been deprecated in favor of grafana.oss/dashboardpermission.DashboardPermission */
    constructor(name: string, argsOrState?: DashboardPermissionArgs | DashboardPermissionState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DashboardPermission is deprecated: grafana.index/dashboardpermission.DashboardPermission has been deprecated in favor of grafana.oss/dashboardpermission.DashboardPermission")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DashboardPermissionState | undefined;
            resourceInputs["dashboardUid"] = state ? state.dashboardUid : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
        } else {
            const args = argsOrState as DashboardPermissionArgs | undefined;
            resourceInputs["dashboardUid"] = args ? args.dashboardUid : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "grafana:index/dashboardPermission:DashboardPermission" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(DashboardPermission.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DashboardPermission resources.
 */
export interface DashboardPermissionState {
    /**
     * UID of the dashboard to apply permissions to.
     */
    dashboardUid?: pulumi.Input<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: pulumi.Input<string>;
    /**
     * The permission items to add/update. Items that are omitted from the list will be removed.
     */
    permissions?: pulumi.Input<pulumi.Input<inputs.DashboardPermissionPermission>[]>;
}

/**
 * The set of arguments for constructing a DashboardPermission resource.
 */
export interface DashboardPermissionArgs {
    /**
     * UID of the dashboard to apply permissions to.
     */
    dashboardUid?: pulumi.Input<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: pulumi.Input<string>;
    /**
     * The permission items to add/update. Items that are omitted from the list will be removed.
     */
    permissions?: pulumi.Input<pulumi.Input<inputs.DashboardPermissionPermission>[]>;
}
