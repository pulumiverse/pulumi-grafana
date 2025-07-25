// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Data source for retrieving sets of cloud IPs. See https://grafana.com/docs/grafana-cloud/reference/allow-list/ for more info
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = grafana.cloud.getIps({});
 * ```
 */
export function getIps(opts?: pulumi.InvokeOptions): Promise<GetIpsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:cloud/getIps:getIps", {
    }, opts);
}

/**
 * A collection of values returned by getIps.
 */
export interface GetIpsResult {
    /**
     * Set of IP addresses that are used for hosted alerts.
     */
    readonly hostedAlerts: string[];
    /**
     * Set of IP addresses that are used for hosted Grafana.
     */
    readonly hostedGrafanas: string[];
    /**
     * Set of IP addresses that are used for hosted logs.
     */
    readonly hostedLogs: string[];
    /**
     * Set of IP addresses that are used for hosted metrics.
     */
    readonly hostedMetrics: string[];
    /**
     * Set of IP addresses that are used for hosted traces.
     */
    readonly hostedTraces: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
/**
 * Data source for retrieving sets of cloud IPs. See https://grafana.com/docs/grafana-cloud/reference/allow-list/ for more info
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = grafana.cloud.getIps({});
 * ```
 */
export function getIpsOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetIpsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("grafana:cloud/getIps:getIps", {
    }, opts);
}
