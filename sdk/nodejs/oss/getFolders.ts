// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/manage-dashboards/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/folder/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const testA = new grafana.oss.Folder("test_a", {
 *     title: "test-folder-a",
 *     uid: "test-ds-folder-uid-a",
 * });
 * const testB = new grafana.oss.Folder("test_b", {
 *     title: "test-folder-b",
 *     uid: "test-ds-folder-uid-b",
 * });
 * const test = grafana.oss.getFolders({});
 * ```
 */
export function getFolders(args?: GetFoldersArgs, opts?: pulumi.InvokeOptions): Promise<GetFoldersResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:oss/getFolders:getFolders", {
        "orgId": args.orgId,
    }, opts);
}

/**
 * A collection of arguments for invoking getFolders.
 */
export interface GetFoldersArgs {
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: string;
}

/**
 * A collection of values returned by getFolders.
 */
export interface GetFoldersResult {
    /**
     * The Grafana instance's folders.
     */
    readonly folders: outputs.oss.GetFoldersFolder[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    readonly orgId?: string;
}
/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/manage-dashboards/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/folder/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const testA = new grafana.oss.Folder("test_a", {
 *     title: "test-folder-a",
 *     uid: "test-ds-folder-uid-a",
 * });
 * const testB = new grafana.oss.Folder("test_b", {
 *     title: "test-folder-b",
 *     uid: "test-ds-folder-uid-b",
 * });
 * const test = grafana.oss.getFolders({});
 * ```
 */
export function getFoldersOutput(args?: GetFoldersOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetFoldersResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("grafana:oss/getFolders:getFolders", {
        "orgId": args.orgId,
    }, opts);
}

/**
 * A collection of arguments for invoking getFolders.
 */
export interface GetFoldersOutputArgs {
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: pulumi.Input<string>;
}
