// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Data source for retrieving a single probe by name.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 *
 * const atlanta = grafana.getSyntheticMonitoringProbe({
 *     name: "Atlanta",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSyntheticMonitoringProbe(args: GetSyntheticMonitoringProbeArgs, opts?: pulumi.InvokeOptions): Promise<GetSyntheticMonitoringProbeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:index/getSyntheticMonitoringProbe:getSyntheticMonitoringProbe", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getSyntheticMonitoringProbe.
 */
export interface GetSyntheticMonitoringProbeArgs {
    /**
     * Name of the probe.
     */
    name: string;
}

/**
 * A collection of values returned by getSyntheticMonitoringProbe.
 */
export interface GetSyntheticMonitoringProbeResult {
    /**
     * The ID of the probe.
     */
    readonly id: string;
    /**
     * Custom labels to be included with collected metrics and logs.
     */
    readonly labels: {[key: string]: string};
    /**
     * Latitude coordinates.
     */
    readonly latitude: number;
    /**
     * Longitude coordinates.
     */
    readonly longitude: number;
    /**
     * Name of the probe.
     */
    readonly name: string;
    /**
     * Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`.
     */
    readonly public: boolean;
    /**
     * Region of the probe.
     */
    readonly region: string;
    /**
     * The tenant ID of the probe.
     */
    readonly tenantId: number;
}
/**
 * Data source for retrieving a single probe by name.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 *
 * const atlanta = grafana.getSyntheticMonitoringProbe({
 *     name: "Atlanta",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSyntheticMonitoringProbeOutput(args: GetSyntheticMonitoringProbeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSyntheticMonitoringProbeResult> {
    return pulumi.output(args).apply((a: any) => getSyntheticMonitoringProbe(a, opts))
}

/**
 * A collection of arguments for invoking getSyntheticMonitoringProbe.
 */
export interface GetSyntheticMonitoringProbeOutputArgs {
    /**
     * Name of the probe.
     */
    name: pulumi.Input<string>;
}
