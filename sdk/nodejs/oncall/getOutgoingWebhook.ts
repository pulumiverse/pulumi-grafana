// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/outgoing_webhooks/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const exampleOutgoingWebhook = grafana.onCall.getOutgoingWebhook({
 *     name: "example_outgoing_webhook",
 * });
 * ```
 */
export function getOutgoingWebhook(args: GetOutgoingWebhookArgs, opts?: pulumi.InvokeOptions): Promise<GetOutgoingWebhookResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:onCall/getOutgoingWebhook:getOutgoingWebhook", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getOutgoingWebhook.
 */
export interface GetOutgoingWebhookArgs {
    /**
     * The outgoing webhook name.
     */
    name: string;
}

/**
 * A collection of values returned by getOutgoingWebhook.
 */
export interface GetOutgoingWebhookResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The outgoing webhook name.
     */
    readonly name: string;
}
/**
 * * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/outgoing_webhooks/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const exampleOutgoingWebhook = grafana.onCall.getOutgoingWebhook({
 *     name: "example_outgoing_webhook",
 * });
 * ```
 */
export function getOutgoingWebhookOutput(args: GetOutgoingWebhookOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetOutgoingWebhookResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("grafana:onCall/getOutgoingWebhook:getOutgoingWebhook", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getOutgoingWebhook.
 */
export interface GetOutgoingWebhookOutputArgs {
    /**
     * The outgoing webhook name.
     */
    name: pulumi.Input<string>;
}
