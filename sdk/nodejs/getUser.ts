// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/user-management/server-user-management/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/user/)
 *
 * This data source uses Grafana's admin APIs for reading users which
 * does not currently work with API Tokens. You must use basic auth.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = new grafana.oss.User("test", {
 *     email: "test.datasource@example.com",
 *     name: "Testing Datasource",
 *     login: "test-datasource",
 *     password: "my-password",
 *     isAdmin: true,
 * });
 * const fromId = grafana.oss.getUserOutput({
 *     userId: test.userId,
 * });
 * const fromEmail = grafana.oss.getUserOutput({
 *     email: test.email,
 * });
 * const fromLogin = test.login.apply(login => grafana.oss.getUserOutput({
 *     login: login,
 * }));
 * ```
 */
/** @deprecated grafana.index/getuser.getUser has been deprecated in favor of grafana.oss/getuser.getUser */
export function getUser(args?: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> {
    pulumi.log.warn("getUser is deprecated: grafana.index/getuser.getUser has been deprecated in favor of grafana.oss/getuser.getUser")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:index/getUser:getUser", {
        "email": args.email,
        "login": args.login,
        "userId": args.userId,
    }, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserArgs {
    /**
     * The email address of the Grafana user. Defaults to ``.
     */
    email?: string;
    /**
     * The username for the Grafana user. Defaults to ``.
     */
    login?: string;
    /**
     * The numerical ID of the Grafana user. Defaults to `-1`.
     */
    userId?: number;
}

/**
 * A collection of values returned by getUser.
 */
export interface GetUserResult {
    /**
     * The email address of the Grafana user. Defaults to ``.
     */
    readonly email?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Whether the user is an admin.
     */
    readonly isAdmin: boolean;
    /**
     * The username for the Grafana user. Defaults to ``.
     */
    readonly login?: string;
    /**
     * The display name for the Grafana user.
     */
    readonly name: string;
    /**
     * The numerical ID of the Grafana user. Defaults to `-1`.
     */
    readonly userId?: number;
}
/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/user-management/server-user-management/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/user/)
 *
 * This data source uses Grafana's admin APIs for reading users which
 * does not currently work with API Tokens. You must use basic auth.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = new grafana.oss.User("test", {
 *     email: "test.datasource@example.com",
 *     name: "Testing Datasource",
 *     login: "test-datasource",
 *     password: "my-password",
 *     isAdmin: true,
 * });
 * const fromId = grafana.oss.getUserOutput({
 *     userId: test.userId,
 * });
 * const fromEmail = grafana.oss.getUserOutput({
 *     email: test.email,
 * });
 * const fromLogin = test.login.apply(login => grafana.oss.getUserOutput({
 *     login: login,
 * }));
 * ```
 */
/** @deprecated grafana.index/getuser.getUser has been deprecated in favor of grafana.oss/getuser.getUser */
export function getUserOutput(args?: GetUserOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserResult> {
    pulumi.log.warn("getUser is deprecated: grafana.index/getuser.getUser has been deprecated in favor of grafana.oss/getuser.getUser")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("grafana:index/getUser:getUser", {
        "email": args.email,
        "login": args.login,
        "userId": args.userId,
    }, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserOutputArgs {
    /**
     * The email address of the Grafana user. Defaults to ``.
     */
    email?: pulumi.Input<string>;
    /**
     * The username for the Grafana user. Defaults to ``.
     */
    login?: pulumi.Input<string>;
    /**
     * The numerical ID of the Grafana user. Defaults to `-1`.
     */
    userId?: pulumi.Input<number>;
}
