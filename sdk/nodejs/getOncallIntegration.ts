// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/** @deprecated grafana.index/getoncallintegration.getOncallIntegration has been deprecated in favor of grafana.oncall/getintegration.getIntegration */
export function getOncallIntegration(args: GetOncallIntegrationArgs, opts?: pulumi.InvokeOptions): Promise<GetOncallIntegrationResult> {
    pulumi.log.warn("getOncallIntegration is deprecated: grafana.index/getoncallintegration.getOncallIntegration has been deprecated in favor of grafana.oncall/getintegration.getIntegration")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:index/getOncallIntegration:getOncallIntegration", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getOncallIntegration.
 */
export interface GetOncallIntegrationArgs {
    id: string;
}

/**
 * A collection of values returned by getOncallIntegration.
 */
export interface GetOncallIntegrationResult {
    readonly id: string;
    readonly name: string;
}
/** @deprecated grafana.index/getoncallintegration.getOncallIntegration has been deprecated in favor of grafana.oncall/getintegration.getIntegration */
export function getOncallIntegrationOutput(args: GetOncallIntegrationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOncallIntegrationResult> {
    return pulumi.output(args).apply((a: any) => getOncallIntegration(a, opts))
}

/**
 * A collection of arguments for invoking getOncallIntegration.
 */
export interface GetOncallIntegrationOutputArgs {
    id: pulumi.Input<string>;
}
