// *** WARNING: this file was generated by pulumi-language-nodejs. ***
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
 * const test = new grafana.oss.LibraryPanel("test", {
 *     name: "panelname",
 *     modelJson: JSON.stringify({
 *         title: "test name",
 *         type: "text",
 *         version: 0,
 *         description: "test description",
 *     }),
 * });
 * const testFolder = new grafana.oss.Folder("test", {
 *     title: "Panel Folder",
 *     uid: "panelname-folder",
 * });
 * const folder = new grafana.oss.LibraryPanel("folder", {
 *     name: "panelname In Folder",
 *     folderUid: testFolder.uid,
 *     modelJson: JSON.stringify({
 *         gridPos: {
 *             x: 0,
 *             y: 0,
 *             h: 10,
 *             w: 10,
 *         },
 *         title: "panel",
 *         type: "text",
 *         version: 0,
 *     }),
 * });
 * const all = grafana.oss.getLibraryPanels({});
 * ```
 */
export function getLibraryPanels(args?: GetLibraryPanelsArgs, opts?: pulumi.InvokeOptions): Promise<GetLibraryPanelsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:oss/getLibraryPanels:getLibraryPanels", {
        "orgId": args.orgId,
    }, opts);
}

/**
 * A collection of arguments for invoking getLibraryPanels.
 */
export interface GetLibraryPanelsArgs {
    /**
     * The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
     */
    orgId?: string;
}

/**
 * A collection of values returned by getLibraryPanels.
 */
export interface GetLibraryPanelsResult {
    /**
     * The ID of this resource.
     */
    readonly id: string;
    /**
     * The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
     */
    readonly orgId: string;
    readonly panels: outputs.oss.GetLibraryPanelsPanel[];
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = new grafana.oss.LibraryPanel("test", {
 *     name: "panelname",
 *     modelJson: JSON.stringify({
 *         title: "test name",
 *         type: "text",
 *         version: 0,
 *         description: "test description",
 *     }),
 * });
 * const testFolder = new grafana.oss.Folder("test", {
 *     title: "Panel Folder",
 *     uid: "panelname-folder",
 * });
 * const folder = new grafana.oss.LibraryPanel("folder", {
 *     name: "panelname In Folder",
 *     folderUid: testFolder.uid,
 *     modelJson: JSON.stringify({
 *         gridPos: {
 *             x: 0,
 *             y: 0,
 *             h: 10,
 *             w: 10,
 *         },
 *         title: "panel",
 *         type: "text",
 *         version: 0,
 *     }),
 * });
 * const all = grafana.oss.getLibraryPanels({});
 * ```
 */
export function getLibraryPanelsOutput(args?: GetLibraryPanelsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLibraryPanelsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("grafana:oss/getLibraryPanels:getLibraryPanels", {
        "orgId": args.orgId,
    }, opts);
}

/**
 * A collection of arguments for invoking getLibraryPanels.
 */
export interface GetLibraryPanelsOutputArgs {
    /**
     * The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
     */
    orgId?: pulumi.Input<string>;
}
