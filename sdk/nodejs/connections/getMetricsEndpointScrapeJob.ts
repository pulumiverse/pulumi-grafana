// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 *
 * const dsTest = grafana.connections.getMetricsEndpointScrapeJob({
 *     stackId: "1",
 *     name: "my-scrape-job",
 * });
 * ```
 */
export function getMetricsEndpointScrapeJob(args: GetMetricsEndpointScrapeJobArgs, opts?: pulumi.InvokeOptions): Promise<GetMetricsEndpointScrapeJobResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:connections/getMetricsEndpointScrapeJob:getMetricsEndpointScrapeJob", {
        "name": args.name,
        "stackId": args.stackId,
    }, opts);
}

/**
 * A collection of arguments for invoking getMetricsEndpointScrapeJob.
 */
export interface GetMetricsEndpointScrapeJobArgs {
    name: string;
    stackId: string;
}

/**
 * A collection of values returned by getMetricsEndpointScrapeJob.
 */
export interface GetMetricsEndpointScrapeJobResult {
    /**
     * Password for basic authentication.
     */
    readonly authenticationBasicPassword: string;
    /**
     * Username for basic authentication.
     */
    readonly authenticationBasicUsername: string;
    /**
     * Token for authentication bearer.
     */
    readonly authenticationBearerToken: string;
    /**
     * Method to pass authentication credentials: basic or bearer.
     */
    readonly authenticationMethod: string;
    /**
     * Whether the metrics endpoint scrape job is enabled or not.
     */
    readonly enabled: boolean;
    readonly id: string;
    readonly name: string;
    /**
     * Frequency for scraping the metrics endpoint: 30, 60, or 120 seconds.
     */
    readonly scrapeIntervalSeconds: number;
    readonly stackId: string;
    /**
     * The url to scrape metrics.
     */
    readonly url: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 *
 * const dsTest = grafana.connections.getMetricsEndpointScrapeJob({
 *     stackId: "1",
 *     name: "my-scrape-job",
 * });
 * ```
 */
export function getMetricsEndpointScrapeJobOutput(args: GetMetricsEndpointScrapeJobOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetMetricsEndpointScrapeJobResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("grafana:connections/getMetricsEndpointScrapeJob:getMetricsEndpointScrapeJob", {
        "name": args.name,
        "stackId": args.stackId,
    }, opts);
}

/**
 * A collection of arguments for invoking getMetricsEndpointScrapeJob.
 */
export interface GetMetricsEndpointScrapeJobOutputArgs {
    name: pulumi.Input<string>;
    stackId: pulumi.Input<string>;
}
