// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a single permission item for a datasource. Conflicts with the "grafana.enterprise.DataSourcePermission" resource which manages the entire set of permissions for a datasource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const team = new grafana.oss.Team("team", {name: "Team Name"});
 * const foo = new grafana.oss.DataSource("foo", {
 *     type: "cloudwatch",
 *     name: "cw-example",
 *     jsonDataEncoded: JSON.stringify({
 *         defaultRegion: "us-east-1",
 *         authType: "keys",
 *     }),
 *     secureJsonDataEncoded: JSON.stringify({
 *         accessKey: "123",
 *         secretKey: "456",
 *     }),
 * });
 * const user = new grafana.oss.User("user", {
 *     name: "test-ds-permissions",
 *     email: "test-ds-permissions@example.com",
 *     login: "test-ds-permissions",
 *     password: "hunter2",
 * });
 * const sa = new grafana.oss.ServiceAccount("sa", {
 *     name: "test-ds-permissions",
 *     role: "Viewer",
 * });
 * const teamDataSourcePermissionItem = new grafana.enterprise.DataSourcePermissionItem("team", {
 *     datasourceUid: foo.uid,
 *     team: team.id,
 *     permission: "Edit",
 * });
 * const userDataSourcePermissionItem = new grafana.enterprise.DataSourcePermissionItem("user", {
 *     datasourceUid: foo.uid,
 *     user: user.id,
 *     permission: "Edit",
 * });
 * const role = new grafana.enterprise.DataSourcePermissionItem("role", {
 *     datasourceUid: foo.uid,
 *     role: "Viewer",
 *     permission: "Query",
 * });
 * const serviceAccount = new grafana.enterprise.DataSourcePermissionItem("service_account", {
 *     datasourceUid: foo.uid,
 *     user: sa.id,
 *     permission: "Query",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:index/dataSourcePermissionItem:DataSourcePermissionItem name "{{ datasourceUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
 * ```
 *
 * ```sh
 * $ pulumi import grafana:index/dataSourcePermissionItem:DataSourcePermissionItem name "{{ orgID }}:{{ datasourceUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
 * ```
 *
 * @deprecated grafana.index/datasourcepermissionitem.DataSourcePermissionItem has been deprecated in favor of grafana.enterprise/datasourcepermissionitem.DataSourcePermissionItem
 */
export class DataSourcePermissionItem extends pulumi.CustomResource {
    /**
     * Get an existing DataSourcePermissionItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataSourcePermissionItemState, opts?: pulumi.CustomResourceOptions): DataSourcePermissionItem {
        pulumi.log.warn("DataSourcePermissionItem is deprecated: grafana.index/datasourcepermissionitem.DataSourcePermissionItem has been deprecated in favor of grafana.enterprise/datasourcepermissionitem.DataSourcePermissionItem")
        return new DataSourcePermissionItem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/dataSourcePermissionItem:DataSourcePermissionItem';

    /**
     * Returns true if the given object is an instance of DataSourcePermissionItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataSourcePermissionItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataSourcePermissionItem.__pulumiType;
    }

    /**
     * The UID of the datasource.
     */
    public readonly datasourceUid!: pulumi.Output<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * the permission to be assigned
     */
    public readonly permission!: pulumi.Output<string>;
    /**
     * the role onto which the permission is to be assigned
     */
    public readonly role!: pulumi.Output<string | undefined>;
    /**
     * the team onto which the permission is to be assigned
     */
    public readonly team!: pulumi.Output<string | undefined>;
    /**
     * the user or service account onto which the permission is to be assigned
     */
    public readonly user!: pulumi.Output<string | undefined>;

    /**
     * Create a DataSourcePermissionItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated grafana.index/datasourcepermissionitem.DataSourcePermissionItem has been deprecated in favor of grafana.enterprise/datasourcepermissionitem.DataSourcePermissionItem */
    constructor(name: string, args: DataSourcePermissionItemArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated grafana.index/datasourcepermissionitem.DataSourcePermissionItem has been deprecated in favor of grafana.enterprise/datasourcepermissionitem.DataSourcePermissionItem */
    constructor(name: string, argsOrState?: DataSourcePermissionItemArgs | DataSourcePermissionItemState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DataSourcePermissionItem is deprecated: grafana.index/datasourcepermissionitem.DataSourcePermissionItem has been deprecated in favor of grafana.enterprise/datasourcepermissionitem.DataSourcePermissionItem")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DataSourcePermissionItemState | undefined;
            resourceInputs["datasourceUid"] = state ? state.datasourceUid : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["permission"] = state ? state.permission : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["team"] = state ? state.team : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as DataSourcePermissionItemArgs | undefined;
            if ((!args || args.datasourceUid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasourceUid'");
            }
            if ((!args || args.permission === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permission'");
            }
            resourceInputs["datasourceUid"] = args ? args.datasourceUid : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["permission"] = args ? args.permission : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["team"] = args ? args.team : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "grafana:index/dataSourcePermissionItem:DataSourcePermissionItem" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(DataSourcePermissionItem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataSourcePermissionItem resources.
 */
export interface DataSourcePermissionItemState {
    /**
     * The UID of the datasource.
     */
    datasourceUid?: pulumi.Input<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: pulumi.Input<string>;
    /**
     * the permission to be assigned
     */
    permission?: pulumi.Input<string>;
    /**
     * the role onto which the permission is to be assigned
     */
    role?: pulumi.Input<string>;
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
 * The set of arguments for constructing a DataSourcePermissionItem resource.
 */
export interface DataSourcePermissionItemArgs {
    /**
     * The UID of the datasource.
     */
    datasourceUid: pulumi.Input<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: pulumi.Input<string>;
    /**
     * the permission to be assigned
     */
    permission: pulumi.Input<string>;
    /**
     * the role onto which the permission is to be assigned
     */
    role?: pulumi.Input<string>;
    /**
     * the team onto which the permission is to be assigned
     */
    team?: pulumi.Input<string>;
    /**
     * the user or service account onto which the permission is to be assigned
     */
    user?: pulumi.Input<string>;
}
