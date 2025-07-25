// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Data source for Grafana Stack
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
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
/** @deprecated grafana.index/getcloudstack.getCloudStack has been deprecated in favor of grafana.cloud/getstack.getStack */
export function getCloudStack(args: GetCloudStackArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudStackResult> {
    pulumi.log.warn("getCloudStack is deprecated: grafana.index/getcloudstack.getCloudStack has been deprecated in favor of grafana.cloud/getstack.getStack")
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:index/getCloudStack:getCloudStack", {
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudStack.
 */
export interface GetCloudStackArgs {
    /**
     * Subdomain that the Grafana instance will be available at (i.e. setting slug to “\n\n” will make the instance
     * available at “https://\n\n.grafana.net".
     */
    slug: string;
}

/**
 * A collection of values returned by getCloudStack.
 */
export interface GetCloudStackResult {
    /**
     * Comma-separated list of CNAMEs that can be whitelisted to access the Alertmanager instances (Optional)
     */
    readonly alertmanagerIpAllowListCname: string;
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
     * Slug of the cluster where this stack resides.
     */
    readonly clusterSlug: string;
    /**
     * Description of stack.
     */
    readonly description: string;
    /**
     * Name of the Fleet Management instance configured for this stack.
     */
    readonly fleetManagementName: string;
    /**
     * Private DNS for Fleet Management when using AWS PrivateLink (only for AWS stacks)
     */
    readonly fleetManagementPrivateConnectivityInfoPrivateDns: string;
    /**
     * Service Name for Fleet Management when using AWS PrivateLink (only for AWS stacks)
     */
    readonly fleetManagementPrivateConnectivityInfoServiceName: string;
    /**
     * Status of the Fleet Management instance configured for this stack.
     */
    readonly fleetManagementStatus: string;
    /**
     * Base URL of the Fleet Management instance configured for this stack.
     */
    readonly fleetManagementUrl: string;
    /**
     * User ID of the Fleet Management instance configured for this stack.
     */
    readonly fleetManagementUserId: number;
    /**
     * Comma-separated list of CNAMEs that can be whitelisted to access the grafana instance (Optional)
     */
    readonly grafanasIpAllowListCname: string;
    /**
     * Comma-separated list of CNAMEs that can be whitelisted to access the Graphite instance (Optional)
     */
    readonly graphiteIpAllowListCname: string;
    readonly graphiteName: string;
    /**
     * Private DNS for Graphite when using AWS PrivateLink (only for AWS stacks)
     */
    readonly graphitePrivateConnectivityInfoPrivateDns: string;
    /**
     * Service Name for Graphite when using AWS PrivateLink (only for AWS stacks)
     */
    readonly graphitePrivateConnectivityInfoServiceName: string;
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
    /**
     * Comma-separated list of CNAMEs that can be whitelisted to access the Logs instance (Optional)
     */
    readonly logsIpAllowListCname: string;
    readonly logsName: string;
    /**
     * Private DNS for Logs when using AWS PrivateLink (only for AWS stacks)
     */
    readonly logsPrivateConnectivityInfoPrivateDns: string;
    /**
     * Service Name for Logs when using AWS PrivateLink (only for AWS stacks)
     */
    readonly logsPrivateConnectivityInfoServiceName: string;
    readonly logsStatus: string;
    readonly logsUrl: string;
    readonly logsUserId: number;
    /**
     * Name of stack. Conventionally matches the url of the instance (e.g. `<stack_slug>.grafana.net`).
     */
    readonly name: string;
    /**
     * Base URL of the OnCall API instance configured for this stack.
     */
    readonly oncallApiUrl: string;
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
     * Private DNS for OTLP when using AWS PrivateLink (only for AWS stacks)
     */
    readonly otlpPrivateConnectivityInfoPrivateDns: string;
    /**
     * Service Name for OTLP when using AWS PrivateLink (only for AWS stacks)
     */
    readonly otlpPrivateConnectivityInfoServiceName: string;
    /**
     * Base URL of the OTLP instance configured for this stack. The username is the stack's ID (`id` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/ for docs on how to use this.
     */
    readonly otlpUrl: string;
    /**
     * Private DNS for PDC's API when using AWS PrivateLink (only for AWS stacks)
     */
    readonly pdcApiPrivateConnectivityInfoPrivateDns: string;
    /**
     * Service Name for PDC's API when using AWS PrivateLink (only for AWS stacks)
     */
    readonly pdcApiPrivateConnectivityInfoServiceName: string;
    /**
     * Private DNS for PDC's Gateway when using AWS PrivateLink (only for AWS stacks)
     */
    readonly pdcGatewayPrivateConnectivityInfoPrivateDns: string;
    /**
     * Service Name for PDC's Gateway when using AWS PrivateLink (only for AWS stacks)
     */
    readonly pdcGatewayPrivateConnectivityInfoServiceName: string;
    /**
     * Comma-separated list of CNAMEs that can be whitelisted to access the Profiles instance (Optional)
     */
    readonly profilesIpAllowListCname: string;
    readonly profilesName: string;
    /**
     * Private DNS for Profiles when using AWS PrivateLink (only for AWS stacks)
     */
    readonly profilesPrivateConnectivityInfoPrivateDns: string;
    /**
     * Service Name for Profiles when using AWS PrivateLink (only for AWS stacks)
     */
    readonly profilesPrivateConnectivityInfoServiceName: string;
    readonly profilesStatus: string;
    readonly profilesUrl: string;
    readonly profilesUserId: number;
    /**
     * Comma-separated list of CNAMEs that can be whitelisted to access the Prometheus instance (Optional)
     */
    readonly prometheusIpAllowListCname: string;
    /**
     * Prometheus name for this instance.
     */
    readonly prometheusName: string;
    /**
     * Private DNS for Prometheus when using AWS PrivateLink (only for AWS stacks)
     */
    readonly prometheusPrivateConnectivityInfoPrivateDns: string;
    /**
     * Service Name for Prometheus when using AWS PrivateLink (only for AWS stacks)
     */
    readonly prometheusPrivateConnectivityInfoServiceName: string;
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
    /**
     * Comma-separated list of CNAMEs that can be whitelisted to access the Traces instance (Optional)
     */
    readonly tracesIpAllowListCname: string;
    readonly tracesName: string;
    /**
     * Private DNS for Traces when using AWS PrivateLink (only for AWS stacks)
     */
    readonly tracesPrivateConnectivityInfoPrivateDns: string;
    /**
     * Service Name for Traces when using AWS PrivateLink (only for AWS stacks)
     */
    readonly tracesPrivateConnectivityInfoServiceName: string;
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
/** @deprecated grafana.index/getcloudstack.getCloudStack has been deprecated in favor of grafana.cloud/getstack.getStack */
export function getCloudStackOutput(args: GetCloudStackOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCloudStackResult> {
    pulumi.log.warn("getCloudStack is deprecated: grafana.index/getcloudstack.getCloudStack has been deprecated in favor of grafana.cloud/getstack.getStack")
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("grafana:index/getCloudStack:getCloudStack", {
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudStack.
 */
export interface GetCloudStackOutputArgs {
    /**
     * Subdomain that the Grafana instance will be available at (i.e. setting slug to “\n\n” will make the instance
     * available at “https://\n\n.grafana.net".
     */
    slug: pulumi.Input<string>;
}
