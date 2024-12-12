// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Data source for Grafana Stack
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const testStack = new grafana.cloud.Stack("test", {
 *     name: "gcloudstacktest",
 *     slug: "gcloudstacktest",
 *     regionSlug: "eu",
 *     description: "Test Grafana Cloud Stack",
 * });
 * const test = grafana.cloud.getStackOutput({
 *     slug: testStack.slug,
 * });
 * ```
 */
export function getStack(args: GetStackArgs, opts?: pulumi.InvokeOptions): Promise<GetStackResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:cloud/getStack:getStack", {
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getStack.
 */
export interface GetStackArgs {
    /**
     * Subdomain that the Grafana instance will be available at (i.e. setting slug to “\n\n” will make the instance
     * available at “https://\n\n.grafana.net".
     */
    slug: string;
}

/**
 * A collection of values returned by getStack.
 */
export interface GetStackResult {
    /**
     * Name of the Alertmanager instance configured for this stack.
     */
    readonly alertmanagerName: string;
    /**
     * Status of the Alertmanager instance configured for this stack.
     */
    readonly alertmanagerStatus: string;
    /**
     * Base URL of the Alertmanager instance configured for this stack.
     */
    readonly alertmanagerUrl: string;
    /**
     * User ID of the Alertmanager instance configured for this stack.
     */
    readonly alertmanagerUserId: number;
    /**
     * Description of stack.
     */
    readonly description: string;
    readonly graphiteName: string;
    readonly graphiteStatus: string;
    readonly graphiteUrl: string;
    readonly graphiteUserId: number;
    /**
     * The stack id assigned to this stack by Grafana.
     */
    readonly id: string;
    /**
     * Base URL of the InfluxDB instance configured for this stack. The username is the same as the metrics' (`prometheusUserId` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/metrics/metrics-influxdb/push-from-telegraf/ for docs on how to use this.
     */
    readonly influxUrl: string;
    /**
     * A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\-.]+$" and stacks cannot have more than 10 labels.
     */
    readonly labels: {[key: string]: string};
    readonly logsName: string;
    readonly logsStatus: string;
    readonly logsUrl: string;
    readonly logsUserId: number;
    /**
     * Name of stack. Conventionally matches the url of the instance (e.g. `<stack_slug>.grafana.net`).
     */
    readonly name: string;
    /**
     * Organization id to assign to this stack.
     */
    readonly orgId: number;
    /**
     * Organization name to assign to this stack.
     */
    readonly orgName: string;
    /**
     * Organization slug to assign to this stack.
     */
    readonly orgSlug: string;
    /**
     * Base URL of the OTLP instance configured for this stack. The username is the stack's ID (`id` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/ for docs on how to use this.
     */
    readonly otlpUrl: string;
    readonly profilesName: string;
    readonly profilesStatus: string;
    readonly profilesUrl: string;
    readonly profilesUserId: number;
    /**
     * Prometheus name for this instance.
     */
    readonly prometheusName: string;
    /**
     * Use this URL to query hosted metrics data e.g. Prometheus data source in Grafana
     */
    readonly prometheusRemoteEndpoint: string;
    /**
     * Use this URL to send prometheus metrics to Grafana cloud
     */
    readonly prometheusRemoteWriteEndpoint: string;
    /**
     * Prometheus status for this instance.
     */
    readonly prometheusStatus: string;
    /**
     * Prometheus url for this instance.
     */
    readonly prometheusUrl: string;
    /**
     * Prometheus user ID. Used for e.g. remote_write.
     */
    readonly prometheusUserId: number;
    /**
     * The region this stack is deployed to.
     */
    readonly regionSlug: string;
    /**
     * Subdomain that the Grafana instance will be available at (i.e. setting slug to “\n\n” will make the instance
     * available at “https://\n\n.grafana.net".
     */
    readonly slug: string;
    /**
     * Status of the stack.
     */
    readonly status: string;
    readonly tracesName: string;
    readonly tracesStatus: string;
    /**
     * Base URL of the Traces instance configured for this stack. To use this in the Tempo data source in Grafana, append `/tempo` to the URL.
     */
    readonly tracesUrl: string;
    readonly tracesUserId: number;
    /**
     * Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
     */
    readonly url: string;
}
/**
 * Data source for Grafana Stack
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const testStack = new grafana.cloud.Stack("test", {
 *     name: "gcloudstacktest",
 *     slug: "gcloudstacktest",
 *     regionSlug: "eu",
 *     description: "Test Grafana Cloud Stack",
 * });
 * const test = grafana.cloud.getStackOutput({
 *     slug: testStack.slug,
 * });
 * ```
 */
export function getStackOutput(args: GetStackOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetStackResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("grafana:cloud/getStack:getStack", {
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getStack.
 */
export interface GetStackOutputArgs {
    /**
     * Subdomain that the Grafana instance will be available at (i.e. setting slug to “\n\n” will make the instance
     * available at “https://\n\n.grafana.net".
     */
    slug: pulumi.Input<string>;
}
