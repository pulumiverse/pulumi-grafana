// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
 * import * as grafana from "@pulumi/grafana";
 *
 * const test = grafana.cloud.getStack({
 *     slug: "gcloudstacktest",
 * });
 * const testGetProviderAwsCloudwatchScrapeJobs = test.then(test => grafana.cloud.getProviderAwsCloudwatchScrapeJobs({
 *     stackId: test.id,
 * }));
 * ```
 */
export function getProviderAwsCloudwatchScrapeJobs(args: GetProviderAwsCloudwatchScrapeJobsArgs, opts?: pulumi.InvokeOptions): Promise<GetProviderAwsCloudwatchScrapeJobsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:cloud/getProviderAwsCloudwatchScrapeJobs:getProviderAwsCloudwatchScrapeJobs", {
        "scrapeJobs": args.scrapeJobs,
        "stackId": args.stackId,
    }, opts);
}

/**
 * A collection of arguments for invoking getProviderAwsCloudwatchScrapeJobs.
 */
export interface GetProviderAwsCloudwatchScrapeJobsArgs {
    /**
     * A list of AWS CloudWatch Scrape Job objects associated with the given StackID.
     */
    scrapeJobs?: inputs.cloud.GetProviderAwsCloudwatchScrapeJobsScrapeJob[];
    stackId: string;
}

/**
 * A collection of values returned by getProviderAwsCloudwatchScrapeJobs.
 */
export interface GetProviderAwsCloudwatchScrapeJobsResult {
    readonly id: string;
    /**
     * A list of AWS CloudWatch Scrape Job objects associated with the given StackID.
     */
    readonly scrapeJobs?: outputs.cloud.GetProviderAwsCloudwatchScrapeJobsScrapeJob[];
    readonly stackId: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 *
 * const test = grafana.cloud.getStack({
 *     slug: "gcloudstacktest",
 * });
 * const testGetProviderAwsCloudwatchScrapeJobs = test.then(test => grafana.cloud.getProviderAwsCloudwatchScrapeJobs({
 *     stackId: test.id,
 * }));
 * ```
 */
export function getProviderAwsCloudwatchScrapeJobsOutput(args: GetProviderAwsCloudwatchScrapeJobsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetProviderAwsCloudwatchScrapeJobsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("grafana:cloud/getProviderAwsCloudwatchScrapeJobs:getProviderAwsCloudwatchScrapeJobs", {
        "scrapeJobs": args.scrapeJobs,
        "stackId": args.stackId,
    }, opts);
}

/**
 * A collection of arguments for invoking getProviderAwsCloudwatchScrapeJobs.
 */
export interface GetProviderAwsCloudwatchScrapeJobsOutputArgs {
    /**
     * A list of AWS CloudWatch Scrape Job objects associated with the given StackID.
     */
    scrapeJobs?: pulumi.Input<pulumi.Input<inputs.cloud.GetProviderAwsCloudwatchScrapeJobsScrapeJobArgs>[]>;
    stackId: pulumi.Input<string>;
}
