// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Data source for retrieving all probes.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 *
 * const main = grafana.getSyntheticMonitoringProbes({});
 * ```
 */
export function getSyntheticMonitoringProbes(args?: GetSyntheticMonitoringProbesArgs, opts?: pulumi.InvokeOptions): Promise<GetSyntheticMonitoringProbesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:index/getSyntheticMonitoringProbes:getSyntheticMonitoringProbes", {
        "filterDeprecated": args.filterDeprecated,
    }, opts);
}

/**
 * A collection of arguments for invoking getSyntheticMonitoringProbes.
 */
export interface GetSyntheticMonitoringProbesArgs {
    /**
     * If true, only probes that are not deprecated will be returned. Defaults to `true`.
     */
    filterDeprecated?: boolean;
}

/**
 * A collection of values returned by getSyntheticMonitoringProbes.
 */
export interface GetSyntheticMonitoringProbesResult {
    /**
     * If true, only probes that are not deprecated will be returned. Defaults to `true`.
     */
    readonly filterDeprecated?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Map of probes with their names as keys and IDs as values.
     */
    readonly probes: {[key: string]: number};
}
/**
 * Data source for retrieving all probes.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 *
 * const main = grafana.getSyntheticMonitoringProbes({});
 * ```
 */
export function getSyntheticMonitoringProbesOutput(args?: GetSyntheticMonitoringProbesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSyntheticMonitoringProbesResult> {
    return pulumi.output(args).apply((a: any) => getSyntheticMonitoringProbes(a, opts))
}

/**
 * A collection of arguments for invoking getSyntheticMonitoringProbes.
 */
export interface GetSyntheticMonitoringProbesOutputArgs {
    /**
     * If true, only probes that are not deprecated will be returned. Defaults to `true`.
     */
    filterDeprecated?: pulumi.Input<boolean>;
}
