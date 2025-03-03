// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * * [Official documentation](https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#stacks/)
 *
 * Required access policy scopes:
 *
 * * stacks:read
 * * stacks:write
 * * stacks:delete
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = new grafana.cloud.Stack("test", {
 *     name: "gcloudstacktest",
 *     slug: "gcloudstacktest",
 *     regionSlug: "eu",
 *     description: "Test Grafana Cloud Stack",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:cloud/stack:Stack name "{{ stackSlugOrID }}"
 * ```
 */
export class Stack extends pulumi.CustomResource {
    /**
     * Get an existing Stack resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StackState, opts?: pulumi.CustomResourceOptions): Stack {
        return new Stack(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:cloud/stack:Stack';

    /**
     * Returns true if the given object is an instance of Stack.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Stack {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Stack.__pulumiType;
    }

    /**
     * Name of the Alertmanager instance configured for this stack.
     */
    public /*out*/ readonly alertmanagerName!: pulumi.Output<string>;
    /**
     * Status of the Alertmanager instance configured for this stack.
     */
    public /*out*/ readonly alertmanagerStatus!: pulumi.Output<string>;
    /**
     * Base URL of the Alertmanager instance configured for this stack.
     */
    public /*out*/ readonly alertmanagerUrl!: pulumi.Output<string>;
    /**
     * User ID of the Alertmanager instance configured for this stack.
     */
    public /*out*/ readonly alertmanagerUserId!: pulumi.Output<number>;
    /**
     * Slug of the cluster where this stack resides.
     */
    public /*out*/ readonly clusterSlug!: pulumi.Output<string>;
    /**
     * Description of stack.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the Fleet Management instance configured for this stack.
     */
    public /*out*/ readonly fleetManagementName!: pulumi.Output<string>;
    /**
     * Status of the Fleet Management instance configured for this stack.
     */
    public /*out*/ readonly fleetManagementStatus!: pulumi.Output<string>;
    /**
     * Base URL of the Fleet Management instance configured for this stack.
     */
    public /*out*/ readonly fleetManagementUrl!: pulumi.Output<string>;
    /**
     * User ID of the Fleet Management instance configured for this stack.
     */
    public /*out*/ readonly fleetManagementUserId!: pulumi.Output<number>;
    public /*out*/ readonly graphiteName!: pulumi.Output<string>;
    public /*out*/ readonly graphiteStatus!: pulumi.Output<string>;
    public /*out*/ readonly graphiteUrl!: pulumi.Output<string>;
    public /*out*/ readonly graphiteUserId!: pulumi.Output<number>;
    /**
     * Base URL of the InfluxDB instance configured for this stack. The username is the same as the metrics' (`prometheusUserId` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/metrics/metrics-influxdb/push-from-telegraf/ for docs on how to use this.
     */
    public /*out*/ readonly influxUrl!: pulumi.Output<string>;
    /**
     * A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\-.]+$" and stacks cannot have more than 10 labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly logsName!: pulumi.Output<string>;
    public /*out*/ readonly logsStatus!: pulumi.Output<string>;
    public /*out*/ readonly logsUrl!: pulumi.Output<string>;
    public /*out*/ readonly logsUserId!: pulumi.Output<number>;
    /**
     * Name of stack. Conventionally matches the url of the instance (e.g. `<stack_slug>.grafana.net`).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Organization id to assign to this stack.
     */
    public /*out*/ readonly orgId!: pulumi.Output<number>;
    /**
     * Organization name to assign to this stack.
     */
    public /*out*/ readonly orgName!: pulumi.Output<string>;
    /**
     * Organization slug to assign to this stack.
     */
    public /*out*/ readonly orgSlug!: pulumi.Output<string>;
    /**
     * Base URL of the OTLP instance configured for this stack. The username is the stack's ID (`id` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/ for docs on how to use this.
     */
    public /*out*/ readonly otlpUrl!: pulumi.Output<string>;
    public /*out*/ readonly profilesName!: pulumi.Output<string>;
    public /*out*/ readonly profilesStatus!: pulumi.Output<string>;
    public /*out*/ readonly profilesUrl!: pulumi.Output<string>;
    public /*out*/ readonly profilesUserId!: pulumi.Output<number>;
    /**
     * Prometheus name for this instance.
     */
    public /*out*/ readonly prometheusName!: pulumi.Output<string>;
    /**
     * Use this URL to query hosted metrics data e.g. Prometheus data source in Grafana
     */
    public /*out*/ readonly prometheusRemoteEndpoint!: pulumi.Output<string>;
    /**
     * Use this URL to send prometheus metrics to Grafana cloud
     */
    public /*out*/ readonly prometheusRemoteWriteEndpoint!: pulumi.Output<string>;
    /**
     * Prometheus status for this instance.
     */
    public /*out*/ readonly prometheusStatus!: pulumi.Output<string>;
    /**
     * Prometheus url for this instance.
     */
    public /*out*/ readonly prometheusUrl!: pulumi.Output<string>;
    /**
     * Prometheus user ID. Used for e.g. remote_write.
     */
    public /*out*/ readonly prometheusUserId!: pulumi.Output<number>;
    /**
     * Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired region. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
     */
    public readonly regionSlug!: pulumi.Output<string | undefined>;
    /**
     * Subdomain that the Grafana instance will be available at. Setting slug to `<stack_slug>` will make the instance available at `https://<stack_slug>.grafana.net`.
     */
    public readonly slug!: pulumi.Output<string>;
    /**
     * Status of the stack.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    public /*out*/ readonly tracesName!: pulumi.Output<string>;
    public /*out*/ readonly tracesStatus!: pulumi.Output<string>;
    /**
     * Base URL of the Traces instance configured for this stack. To use this in the Tempo data source in Grafana, append `/tempo` to the URL.
     */
    public /*out*/ readonly tracesUrl!: pulumi.Output<string>;
    public /*out*/ readonly tracesUserId!: pulumi.Output<number>;
    /**
     * Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
     */
    public readonly url!: pulumi.Output<string | undefined>;
    /**
     * Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana instance). Defaults to `true`.
     */
    public readonly waitForReadiness!: pulumi.Output<boolean | undefined>;
    /**
     * How long to wait for readiness (if enabled). Defaults to `5m0s`.
     */
    public readonly waitForReadinessTimeout!: pulumi.Output<string | undefined>;

    /**
     * Create a Stack resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StackArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StackArgs | StackState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StackState | undefined;
            resourceInputs["alertmanagerName"] = state ? state.alertmanagerName : undefined;
            resourceInputs["alertmanagerStatus"] = state ? state.alertmanagerStatus : undefined;
            resourceInputs["alertmanagerUrl"] = state ? state.alertmanagerUrl : undefined;
            resourceInputs["alertmanagerUserId"] = state ? state.alertmanagerUserId : undefined;
            resourceInputs["clusterSlug"] = state ? state.clusterSlug : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["fleetManagementName"] = state ? state.fleetManagementName : undefined;
            resourceInputs["fleetManagementStatus"] = state ? state.fleetManagementStatus : undefined;
            resourceInputs["fleetManagementUrl"] = state ? state.fleetManagementUrl : undefined;
            resourceInputs["fleetManagementUserId"] = state ? state.fleetManagementUserId : undefined;
            resourceInputs["graphiteName"] = state ? state.graphiteName : undefined;
            resourceInputs["graphiteStatus"] = state ? state.graphiteStatus : undefined;
            resourceInputs["graphiteUrl"] = state ? state.graphiteUrl : undefined;
            resourceInputs["graphiteUserId"] = state ? state.graphiteUserId : undefined;
            resourceInputs["influxUrl"] = state ? state.influxUrl : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["logsName"] = state ? state.logsName : undefined;
            resourceInputs["logsStatus"] = state ? state.logsStatus : undefined;
            resourceInputs["logsUrl"] = state ? state.logsUrl : undefined;
            resourceInputs["logsUserId"] = state ? state.logsUserId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["orgName"] = state ? state.orgName : undefined;
            resourceInputs["orgSlug"] = state ? state.orgSlug : undefined;
            resourceInputs["otlpUrl"] = state ? state.otlpUrl : undefined;
            resourceInputs["profilesName"] = state ? state.profilesName : undefined;
            resourceInputs["profilesStatus"] = state ? state.profilesStatus : undefined;
            resourceInputs["profilesUrl"] = state ? state.profilesUrl : undefined;
            resourceInputs["profilesUserId"] = state ? state.profilesUserId : undefined;
            resourceInputs["prometheusName"] = state ? state.prometheusName : undefined;
            resourceInputs["prometheusRemoteEndpoint"] = state ? state.prometheusRemoteEndpoint : undefined;
            resourceInputs["prometheusRemoteWriteEndpoint"] = state ? state.prometheusRemoteWriteEndpoint : undefined;
            resourceInputs["prometheusStatus"] = state ? state.prometheusStatus : undefined;
            resourceInputs["prometheusUrl"] = state ? state.prometheusUrl : undefined;
            resourceInputs["prometheusUserId"] = state ? state.prometheusUserId : undefined;
            resourceInputs["regionSlug"] = state ? state.regionSlug : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tracesName"] = state ? state.tracesName : undefined;
            resourceInputs["tracesStatus"] = state ? state.tracesStatus : undefined;
            resourceInputs["tracesUrl"] = state ? state.tracesUrl : undefined;
            resourceInputs["tracesUserId"] = state ? state.tracesUserId : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["waitForReadiness"] = state ? state.waitForReadiness : undefined;
            resourceInputs["waitForReadinessTimeout"] = state ? state.waitForReadinessTimeout : undefined;
        } else {
            const args = argsOrState as StackArgs | undefined;
            if ((!args || args.slug === undefined) && !opts.urn) {
                throw new Error("Missing required property 'slug'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["regionSlug"] = args ? args.regionSlug : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["waitForReadiness"] = args ? args.waitForReadiness : undefined;
            resourceInputs["waitForReadinessTimeout"] = args ? args.waitForReadinessTimeout : undefined;
            resourceInputs["alertmanagerName"] = undefined /*out*/;
            resourceInputs["alertmanagerStatus"] = undefined /*out*/;
            resourceInputs["alertmanagerUrl"] = undefined /*out*/;
            resourceInputs["alertmanagerUserId"] = undefined /*out*/;
            resourceInputs["clusterSlug"] = undefined /*out*/;
            resourceInputs["fleetManagementName"] = undefined /*out*/;
            resourceInputs["fleetManagementStatus"] = undefined /*out*/;
            resourceInputs["fleetManagementUrl"] = undefined /*out*/;
            resourceInputs["fleetManagementUserId"] = undefined /*out*/;
            resourceInputs["graphiteName"] = undefined /*out*/;
            resourceInputs["graphiteStatus"] = undefined /*out*/;
            resourceInputs["graphiteUrl"] = undefined /*out*/;
            resourceInputs["graphiteUserId"] = undefined /*out*/;
            resourceInputs["influxUrl"] = undefined /*out*/;
            resourceInputs["logsName"] = undefined /*out*/;
            resourceInputs["logsStatus"] = undefined /*out*/;
            resourceInputs["logsUrl"] = undefined /*out*/;
            resourceInputs["logsUserId"] = undefined /*out*/;
            resourceInputs["orgId"] = undefined /*out*/;
            resourceInputs["orgName"] = undefined /*out*/;
            resourceInputs["orgSlug"] = undefined /*out*/;
            resourceInputs["otlpUrl"] = undefined /*out*/;
            resourceInputs["profilesName"] = undefined /*out*/;
            resourceInputs["profilesStatus"] = undefined /*out*/;
            resourceInputs["profilesUrl"] = undefined /*out*/;
            resourceInputs["profilesUserId"] = undefined /*out*/;
            resourceInputs["prometheusName"] = undefined /*out*/;
            resourceInputs["prometheusRemoteEndpoint"] = undefined /*out*/;
            resourceInputs["prometheusRemoteWriteEndpoint"] = undefined /*out*/;
            resourceInputs["prometheusStatus"] = undefined /*out*/;
            resourceInputs["prometheusUrl"] = undefined /*out*/;
            resourceInputs["prometheusUserId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tracesName"] = undefined /*out*/;
            resourceInputs["tracesStatus"] = undefined /*out*/;
            resourceInputs["tracesUrl"] = undefined /*out*/;
            resourceInputs["tracesUserId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Stack.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Stack resources.
 */
export interface StackState {
    /**
     * Name of the Alertmanager instance configured for this stack.
     */
    alertmanagerName?: pulumi.Input<string>;
    /**
     * Status of the Alertmanager instance configured for this stack.
     */
    alertmanagerStatus?: pulumi.Input<string>;
    /**
     * Base URL of the Alertmanager instance configured for this stack.
     */
    alertmanagerUrl?: pulumi.Input<string>;
    /**
     * User ID of the Alertmanager instance configured for this stack.
     */
    alertmanagerUserId?: pulumi.Input<number>;
    /**
     * Slug of the cluster where this stack resides.
     */
    clusterSlug?: pulumi.Input<string>;
    /**
     * Description of stack.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the Fleet Management instance configured for this stack.
     */
    fleetManagementName?: pulumi.Input<string>;
    /**
     * Status of the Fleet Management instance configured for this stack.
     */
    fleetManagementStatus?: pulumi.Input<string>;
    /**
     * Base URL of the Fleet Management instance configured for this stack.
     */
    fleetManagementUrl?: pulumi.Input<string>;
    /**
     * User ID of the Fleet Management instance configured for this stack.
     */
    fleetManagementUserId?: pulumi.Input<number>;
    graphiteName?: pulumi.Input<string>;
    graphiteStatus?: pulumi.Input<string>;
    graphiteUrl?: pulumi.Input<string>;
    graphiteUserId?: pulumi.Input<number>;
    /**
     * Base URL of the InfluxDB instance configured for this stack. The username is the same as the metrics' (`prometheusUserId` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/metrics/metrics-influxdb/push-from-telegraf/ for docs on how to use this.
     */
    influxUrl?: pulumi.Input<string>;
    /**
     * A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\-.]+$" and stacks cannot have more than 10 labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    logsName?: pulumi.Input<string>;
    logsStatus?: pulumi.Input<string>;
    logsUrl?: pulumi.Input<string>;
    logsUserId?: pulumi.Input<number>;
    /**
     * Name of stack. Conventionally matches the url of the instance (e.g. `<stack_slug>.grafana.net`).
     */
    name?: pulumi.Input<string>;
    /**
     * Organization id to assign to this stack.
     */
    orgId?: pulumi.Input<number>;
    /**
     * Organization name to assign to this stack.
     */
    orgName?: pulumi.Input<string>;
    /**
     * Organization slug to assign to this stack.
     */
    orgSlug?: pulumi.Input<string>;
    /**
     * Base URL of the OTLP instance configured for this stack. The username is the stack's ID (`id` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/ for docs on how to use this.
     */
    otlpUrl?: pulumi.Input<string>;
    profilesName?: pulumi.Input<string>;
    profilesStatus?: pulumi.Input<string>;
    profilesUrl?: pulumi.Input<string>;
    profilesUserId?: pulumi.Input<number>;
    /**
     * Prometheus name for this instance.
     */
    prometheusName?: pulumi.Input<string>;
    /**
     * Use this URL to query hosted metrics data e.g. Prometheus data source in Grafana
     */
    prometheusRemoteEndpoint?: pulumi.Input<string>;
    /**
     * Use this URL to send prometheus metrics to Grafana cloud
     */
    prometheusRemoteWriteEndpoint?: pulumi.Input<string>;
    /**
     * Prometheus status for this instance.
     */
    prometheusStatus?: pulumi.Input<string>;
    /**
     * Prometheus url for this instance.
     */
    prometheusUrl?: pulumi.Input<string>;
    /**
     * Prometheus user ID. Used for e.g. remote_write.
     */
    prometheusUserId?: pulumi.Input<number>;
    /**
     * Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired region. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
     */
    regionSlug?: pulumi.Input<string>;
    /**
     * Subdomain that the Grafana instance will be available at. Setting slug to `<stack_slug>` will make the instance available at `https://<stack_slug>.grafana.net`.
     */
    slug?: pulumi.Input<string>;
    /**
     * Status of the stack.
     */
    status?: pulumi.Input<string>;
    tracesName?: pulumi.Input<string>;
    tracesStatus?: pulumi.Input<string>;
    /**
     * Base URL of the Traces instance configured for this stack. To use this in the Tempo data source in Grafana, append `/tempo` to the URL.
     */
    tracesUrl?: pulumi.Input<string>;
    tracesUserId?: pulumi.Input<number>;
    /**
     * Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
     */
    url?: pulumi.Input<string>;
    /**
     * Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana instance). Defaults to `true`.
     */
    waitForReadiness?: pulumi.Input<boolean>;
    /**
     * How long to wait for readiness (if enabled). Defaults to `5m0s`.
     */
    waitForReadinessTimeout?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Stack resource.
 */
export interface StackArgs {
    /**
     * Description of stack.
     */
    description?: pulumi.Input<string>;
    /**
     * A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\-.]+$" and stacks cannot have more than 10 labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of stack. Conventionally matches the url of the instance (e.g. `<stack_slug>.grafana.net`).
     */
    name?: pulumi.Input<string>;
    /**
     * Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired region. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
     */
    regionSlug?: pulumi.Input<string>;
    /**
     * Subdomain that the Grafana instance will be available at. Setting slug to `<stack_slug>` will make the instance available at `https://<stack_slug>.grafana.net`.
     */
    slug: pulumi.Input<string>;
    /**
     * Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
     */
    url?: pulumi.Input<string>;
    /**
     * Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana instance). Defaults to `true`.
     */
    waitForReadiness?: pulumi.Input<boolean>;
    /**
     * How long to wait for readiness (if enabled). Defaults to `5m0s`.
     */
    waitForReadinessTimeout?: pulumi.Input<string>;
}
