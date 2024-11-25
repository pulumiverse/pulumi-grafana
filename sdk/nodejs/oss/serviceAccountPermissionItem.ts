// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a single permission item for a service account. Conflicts with the "grafana.oss.ServiceAccountPermission" resource which manages the entire set of permissions for a service account.
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/#manage-users-and-teams-permissions-for-a-service-account-in-grafana)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = new grafana.oss.ServiceAccount("test", {
 *     name: "terraform-sa",
 *     role: "Editor",
 *     isDisabled: false,
 * });
 * const team = new grafana.oss.Team("team", {name: "Team Name"});
 * const user = new grafana.oss.User("user", {
 *     email: "user.name@example.com",
 *     login: "user.name",
 *     password: "my-password",
 * });
 * const onTeam = new grafana.oss.ServiceAccountPermissionItem("on_team", {
 *     serviceAccountId: test.id,
 *     team: team.id,
 *     permission: "Admin",
 * });
 * const onUser = new grafana.oss.ServiceAccountPermissionItem("on_user", {
 *     serviceAccountId: test.id,
 *     user: user.id,
 *     permission: "Admin",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:oss/serviceAccountPermissionItem:ServiceAccountPermissionItem name "{{ serviceAccountID }}:{{ type (role, team, or user) }}:{{ identifier }}"
 * ```
 *
 * ```sh
 * $ pulumi import grafana:oss/serviceAccountPermissionItem:ServiceAccountPermissionItem name "{{ orgID }}:{{ serviceAccountID }}:{{ type (role, team, or user) }}:{{ identifier }}"
 * ```
 */
export class ServiceAccountPermissionItem extends pulumi.CustomResource {
    /**
     * Get an existing ServiceAccountPermissionItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceAccountPermissionItemState, opts?: pulumi.CustomResourceOptions): ServiceAccountPermissionItem {
        return new ServiceAccountPermissionItem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:oss/serviceAccountPermissionItem:ServiceAccountPermissionItem';

    /**
     * Returns true if the given object is an instance of ServiceAccountPermissionItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceAccountPermissionItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceAccountPermissionItem.__pulumiType;
    }

    /**
     * The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * the permission to be assigned
     */
    public readonly permission!: pulumi.Output<string>;
    /**
     * The ID of the service account.
     */
    public readonly serviceAccountId!: pulumi.Output<string>;
    /**
     * the team onto which the permission is to be assigned
     */
    public readonly team!: pulumi.Output<string | undefined>;
    /**
     * the user or service account onto which the permission is to be assigned
     */
    public readonly user!: pulumi.Output<string | undefined>;

    /**
     * Create a ServiceAccountPermissionItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceAccountPermissionItemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceAccountPermissionItemArgs | ServiceAccountPermissionItemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceAccountPermissionItemState | undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["permission"] = state ? state.permission : undefined;
            resourceInputs["serviceAccountId"] = state ? state.serviceAccountId : undefined;
            resourceInputs["team"] = state ? state.team : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as ServiceAccountPermissionItemArgs | undefined;
            if ((!args || args.permission === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permission'");
            }
            if ((!args || args.serviceAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceAccountId'");
            }
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["permission"] = args ? args.permission : undefined;
            resourceInputs["serviceAccountId"] = args ? args.serviceAccountId : undefined;
            resourceInputs["team"] = args ? args.team : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "grafana:index/serviceAccountPermissionItem:ServiceAccountPermissionItem" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ServiceAccountPermissionItem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceAccountPermissionItem resources.
 */
export interface ServiceAccountPermissionItemState {
    /**
     * The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
     */
    orgId?: pulumi.Input<string>;
    /**
     * the permission to be assigned
     */
    permission?: pulumi.Input<string>;
    /**
     * The ID of the service account.
     */
    serviceAccountId?: pulumi.Input<string>;
    /**
     * the team onto which the permission is to be assigned
     */
    team?: pulumi.Input<string>;
    /**
     * the user or service account onto which the permission is to be assigned
     */
    user?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceAccountPermissionItem resource.
 */
export interface ServiceAccountPermissionItemArgs {
    /**
     * The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
     */
    orgId?: pulumi.Input<string>;
    /**
     * the permission to be assigned
     */
    permission: pulumi.Input<string>;
    /**
     * The ID of the service account.
     */
    serviceAccountId: pulumi.Input<string>;
    /**
     * the team onto which the permission is to be assigned
     */
    team?: pulumi.Input<string>;
    /**
     * the user or service account onto which the permission is to be assigned
     */
    user?: pulumi.Input<string>;
}
