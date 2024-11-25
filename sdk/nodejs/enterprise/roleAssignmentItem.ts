// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a single assignment for a role. Conflicts with the "grafana.enterprise.RoleAssignment" resource which manages the entire set of assignments for a role.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const testRole = new grafana.enterprise.Role("test_role", {
 *     name: "Test Role",
 *     uid: "testrole",
 *     version: 1,
 *     global: true,
 *     permissions: [{
 *         action: "org.users:add",
 *         scope: "users:*",
 *     }],
 * });
 * const testTeam = new grafana.oss.Team("test_team", {name: "terraform_test_team"});
 * const testUser = new grafana.oss.User("test_user", {
 *     email: "terraform_user@test.com",
 *     login: "terraform_user@test.com",
 *     password: "password",
 * });
 * const testSa = new grafana.oss.ServiceAccount("test_sa", {
 *     name: "terraform_test_sa",
 *     role: "Viewer",
 * });
 * const user = new grafana.enterprise.RoleAssignmentItem("user", {
 *     roleUid: testRole.uid,
 *     userId: testUser.id,
 * });
 * const team = new grafana.enterprise.RoleAssignmentItem("team", {
 *     roleUid: testRole.uid,
 *     teamId: testTeam.id,
 * });
 * const serviceAccount = new grafana.enterprise.RoleAssignmentItem("service_account", {
 *     roleUid: testRole.uid,
 *     serviceAccountId: testSa.id,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:enterprise/roleAssignmentItem:RoleAssignmentItem name "{{ roleUID }}:{{ type (user, team or service_account) }}:{{ identifier }}"
 * ```
 *
 * ```sh
 * $ pulumi import grafana:enterprise/roleAssignmentItem:RoleAssignmentItem name "{{ orgID }}:{{ roleUID }}:{{ type (user, team or service_account) }}:{{ identifier }}"
 * ```
 */
export class RoleAssignmentItem extends pulumi.CustomResource {
    /**
     * Get an existing RoleAssignmentItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RoleAssignmentItemState, opts?: pulumi.CustomResourceOptions): RoleAssignmentItem {
        return new RoleAssignmentItem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:enterprise/roleAssignmentItem:RoleAssignmentItem';

    /**
     * Returns true if the given object is an instance of RoleAssignmentItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RoleAssignmentItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RoleAssignmentItem.__pulumiType;
    }

    /**
     * The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * the role UID onto which to assign an actor
     */
    public readonly roleUid!: pulumi.Output<string>;
    /**
     * the service account onto which the role is to be assigned
     */
    public readonly serviceAccountId!: pulumi.Output<string | undefined>;
    /**
     * the team onto which the role is to be assigned
     */
    public readonly teamId!: pulumi.Output<string | undefined>;
    /**
     * the user onto which the role is to be assigned
     */
    public readonly userId!: pulumi.Output<string | undefined>;

    /**
     * Create a RoleAssignmentItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoleAssignmentItemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RoleAssignmentItemArgs | RoleAssignmentItemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RoleAssignmentItemState | undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["roleUid"] = state ? state.roleUid : undefined;
            resourceInputs["serviceAccountId"] = state ? state.serviceAccountId : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as RoleAssignmentItemArgs | undefined;
            if ((!args || args.roleUid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleUid'");
            }
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["roleUid"] = args ? args.roleUid : undefined;
            resourceInputs["serviceAccountId"] = args ? args.serviceAccountId : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "grafana:index/roleAssignmentItem:RoleAssignmentItem" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(RoleAssignmentItem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RoleAssignmentItem resources.
 */
export interface RoleAssignmentItemState {
    /**
     * The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
     */
    orgId?: pulumi.Input<string>;
    /**
     * the role UID onto which to assign an actor
     */
    roleUid?: pulumi.Input<string>;
    /**
     * the service account onto which the role is to be assigned
     */
    serviceAccountId?: pulumi.Input<string>;
    /**
     * the team onto which the role is to be assigned
     */
    teamId?: pulumi.Input<string>;
    /**
     * the user onto which the role is to be assigned
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RoleAssignmentItem resource.
 */
export interface RoleAssignmentItemArgs {
    /**
     * The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
     */
    orgId?: pulumi.Input<string>;
    /**
     * the role UID onto which to assign an actor
     */
    roleUid: pulumi.Input<string>;
    /**
     * the service account onto which the role is to be assigned
     */
    serviceAccountId?: pulumi.Input<string>;
    /**
     * the team onto which the role is to be assigned
     */
    teamId?: pulumi.Input<string>;
    /**
     * the user onto which the role is to be assigned
     */
    userId?: pulumi.Input<string>;
}
