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
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = grafana.cloud.getStack({
 *     slug: "gcloudstacktest",
 * });
 * const testAwsAccount = new grafana.cloudprovider.AwsAccount("test", {
 *     stackId: test.then(test => test.id),
 *     roleArn: testAwsIamRole.arn,
 *     regions: [
 *         "us-east-2",
 *         "eu-west-3",
 *     ],
 * });
 * const testGetAwsAccount = pulumi.all([test, testAwsAccount.resourceId]).apply(([test, resourceId]) => grafana.cloudProvider.getAwsAccountOutput({
 *     stackId: test.id,
 *     resourceId: resourceId,
 * }));
 * ```
 */
export function getAwsAccount(args: GetAwsAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetAwsAccountResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:cloudProvider/getAwsAccount:getAwsAccount", {
        "resourceId": args.resourceId,
        "stackId": args.stackId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAwsAccount.
 */
export interface GetAwsAccountArgs {
    /**
     * The ID given by the Grafana Cloud Provider API to this AWS Account resource.
     */
    resourceId: string;
    stackId: string;
}

/**
 * A collection of values returned by getAwsAccount.
 */
export interface GetAwsAccountResult {
    readonly id: string;
    /**
     * An optional human-readable name for this AWS Account resource.
     */
    readonly name: string;
    /**
     * A set of regions that this AWS Account resource applies to.
     */
    readonly regions: string[];
    /**
     * The ID given by the Grafana Cloud Provider API to this AWS Account resource.
     */
    readonly resourceId: string;
    /**
     * An IAM Role ARN string to represent with this AWS Account resource.
     */
    readonly roleArn: string;
    readonly stackId: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = grafana.cloud.getStack({
 *     slug: "gcloudstacktest",
 * });
 * const testAwsAccount = new grafana.cloudprovider.AwsAccount("test", {
 *     stackId: test.then(test => test.id),
 *     roleArn: testAwsIamRole.arn,
 *     regions: [
 *         "us-east-2",
 *         "eu-west-3",
 *     ],
 * });
 * const testGetAwsAccount = pulumi.all([test, testAwsAccount.resourceId]).apply(([test, resourceId]) => grafana.cloudProvider.getAwsAccountOutput({
 *     stackId: test.id,
 *     resourceId: resourceId,
 * }));
 * ```
 */
export function getAwsAccountOutput(args: GetAwsAccountOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAwsAccountResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("grafana:cloudProvider/getAwsAccount:getAwsAccount", {
        "resourceId": args.resourceId,
        "stackId": args.stackId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAwsAccount.
 */
export interface GetAwsAccountOutputArgs {
    /**
     * The ID given by the Grafana Cloud Provider API to this AWS Account resource.
     */
    resourceId: pulumi.Input<string>;
    stackId: pulumi.Input<string>;
}
