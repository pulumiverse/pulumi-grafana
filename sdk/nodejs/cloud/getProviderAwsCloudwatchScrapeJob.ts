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
 * import * as aws from "@pulumi/aws";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = grafana.cloud.getStack({
 *     slug: "gcloudstacktest",
 * });
 * const testGetRole = aws.iam.getRole({
 *     name: "my-role",
 * });
 * const testAwsAccount = new grafana.cloudprovider.AwsAccount("test", {
 *     stackId: test.then(test => test.id),
 *     roleArn: testGetRole.then(testGetRole => testGetRole.arn),
 *     regions: [
 *         "us-east-1",
 *         "us-east-2",
 *         "us-west-1",
 *     ],
 * });
 * const testAwsCloudwatchScrapeJob = new grafana.cloudprovider.AwsCloudwatchScrapeJob("test", {
 *     stackId: test.then(test => test.id),
 *     name: "my-cloudwatch-scrape-job",
 *     awsAccountResourceId: testAwsAccount.resourceId,
 *     exportTags: true,
 *     services: [{
 *         name: "AWS/EC2",
 *         metrics: [
 *             {
 *                 name: "CPUUtilization",
 *                 statistics: ["Average"],
 *             },
 *             {
 *                 name: "StatusCheckFailed",
 *                 statistics: ["Maximum"],
 *             },
 *         ],
 *         scrapeIntervalSeconds: 300,
 *         resourceDiscoveryTagFilters: [{
 *             key: "k8s.io/cluster-autoscaler/enabled",
 *             value: "true",
 *         }],
 *         tagsToAddToMetrics: ["eks:cluster-name"],
 *     }],
 *     customNamespaces: [{
 *         name: "CoolApp",
 *         metrics: [{
 *             name: "CoolMetric",
 *             statistics: [
 *                 "Maximum",
 *                 "Sum",
 *             ],
 *         }],
 *         scrapeIntervalSeconds: 300,
 *     }],
 *     staticLabels: {
 *         label1: "value1",
 *         label2: "value2",
 *     },
 * });
 * const testGetAwsCloudwatchScrapeJob = pulumi.all([test, testAwsCloudwatchScrapeJob.name]).apply(([test, name]) => grafana.cloudProvider.getAwsCloudwatchScrapeJobOutput({
 *     stackId: test.id,
 *     name: name,
 * }));
 * ```
 */
/** @deprecated grafana.cloud/getproviderawscloudwatchscrapejob.getProviderAwsCloudwatchScrapeJob has been deprecated in favor of grafana.cloudprovider/getawscloudwatchscrapejob.getAwsCloudwatchScrapeJob */
export function getProviderAwsCloudwatchScrapeJob(args: GetProviderAwsCloudwatchScrapeJobArgs, opts?: pulumi.InvokeOptions): Promise<GetProviderAwsCloudwatchScrapeJobResult> {
    pulumi.log.warn("getProviderAwsCloudwatchScrapeJob is deprecated: grafana.cloud/getproviderawscloudwatchscrapejob.getProviderAwsCloudwatchScrapeJob has been deprecated in favor of grafana.cloudprovider/getawscloudwatchscrapejob.getAwsCloudwatchScrapeJob")
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:cloud/getProviderAwsCloudwatchScrapeJob:getProviderAwsCloudwatchScrapeJob", {
        "customNamespaces": args.customNamespaces,
        "name": args.name,
        "services": args.services,
        "stackId": args.stackId,
    }, opts);
}

/**
 * A collection of arguments for invoking getProviderAwsCloudwatchScrapeJob.
 */
export interface GetProviderAwsCloudwatchScrapeJobArgs {
    /**
     * Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
     */
    customNamespaces?: inputs.cloud.GetProviderAwsCloudwatchScrapeJobCustomNamespace[];
    name: string;
    /**
     * One or more configuration blocks to dictate what this CloudWatch Scrape Job should scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
     */
    services?: inputs.cloud.GetProviderAwsCloudwatchScrapeJobService[];
    stackId: string;
}

/**
 * A collection of values returned by getProviderAwsCloudwatchScrapeJob.
 */
export interface GetProviderAwsCloudwatchScrapeJobResult {
    /**
     * The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resourceId` attribute of the `grafana.cloudProvider.AwsAccount` resource.
     */
    readonly awsAccountResourceId: string;
    /**
     * Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
     */
    readonly customNamespaces?: outputs.cloud.GetProviderAwsCloudwatchScrapeJobCustomNamespace[];
    /**
     * When the CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
     */
    readonly disabledReason: string;
    /**
     * Whether the CloudWatch Scrape Job is enabled or not.
     */
    readonly enabled: boolean;
    /**
     * When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`.
     */
    readonly exportTags: boolean;
    readonly id: string;
    readonly name: string;
    /**
     * The set of AWS region names that this CloudWatch Scrape Job is configured to scrape.
     */
    readonly regions: string[];
    /**
     * When true, the `regions` attribute will be the set of regions configured in the override. When false, the `regions` attribute will be the set of regions belonging to the AWS Account resource that is associated with this CloudWatch Scrape Job.
     */
    readonly regionsSubsetOverrideUsed: boolean;
    /**
     * The AWS ARN of the IAM role associated with the AWS Account resource that is being used by this CloudWatch Scrape Job.
     */
    readonly roleArn: string;
    /**
     * One or more configuration blocks to dictate what this CloudWatch Scrape Job should scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
     */
    readonly services?: outputs.cloud.GetProviderAwsCloudwatchScrapeJobService[];
    readonly stackId: string;
    /**
     * A set of static labels to add to all metrics exported by this scrape job.
     */
    readonly staticLabels: {[key: string]: string};
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = grafana.cloud.getStack({
 *     slug: "gcloudstacktest",
 * });
 * const testGetRole = aws.iam.getRole({
 *     name: "my-role",
 * });
 * const testAwsAccount = new grafana.cloudprovider.AwsAccount("test", {
 *     stackId: test.then(test => test.id),
 *     roleArn: testGetRole.then(testGetRole => testGetRole.arn),
 *     regions: [
 *         "us-east-1",
 *         "us-east-2",
 *         "us-west-1",
 *     ],
 * });
 * const testAwsCloudwatchScrapeJob = new grafana.cloudprovider.AwsCloudwatchScrapeJob("test", {
 *     stackId: test.then(test => test.id),
 *     name: "my-cloudwatch-scrape-job",
 *     awsAccountResourceId: testAwsAccount.resourceId,
 *     exportTags: true,
 *     services: [{
 *         name: "AWS/EC2",
 *         metrics: [
 *             {
 *                 name: "CPUUtilization",
 *                 statistics: ["Average"],
 *             },
 *             {
 *                 name: "StatusCheckFailed",
 *                 statistics: ["Maximum"],
 *             },
 *         ],
 *         scrapeIntervalSeconds: 300,
 *         resourceDiscoveryTagFilters: [{
 *             key: "k8s.io/cluster-autoscaler/enabled",
 *             value: "true",
 *         }],
 *         tagsToAddToMetrics: ["eks:cluster-name"],
 *     }],
 *     customNamespaces: [{
 *         name: "CoolApp",
 *         metrics: [{
 *             name: "CoolMetric",
 *             statistics: [
 *                 "Maximum",
 *                 "Sum",
 *             ],
 *         }],
 *         scrapeIntervalSeconds: 300,
 *     }],
 *     staticLabels: {
 *         label1: "value1",
 *         label2: "value2",
 *     },
 * });
 * const testGetAwsCloudwatchScrapeJob = pulumi.all([test, testAwsCloudwatchScrapeJob.name]).apply(([test, name]) => grafana.cloudProvider.getAwsCloudwatchScrapeJobOutput({
 *     stackId: test.id,
 *     name: name,
 * }));
 * ```
 */
/** @deprecated grafana.cloud/getproviderawscloudwatchscrapejob.getProviderAwsCloudwatchScrapeJob has been deprecated in favor of grafana.cloudprovider/getawscloudwatchscrapejob.getAwsCloudwatchScrapeJob */
export function getProviderAwsCloudwatchScrapeJobOutput(args: GetProviderAwsCloudwatchScrapeJobOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetProviderAwsCloudwatchScrapeJobResult> {
    pulumi.log.warn("getProviderAwsCloudwatchScrapeJob is deprecated: grafana.cloud/getproviderawscloudwatchscrapejob.getProviderAwsCloudwatchScrapeJob has been deprecated in favor of grafana.cloudprovider/getawscloudwatchscrapejob.getAwsCloudwatchScrapeJob")
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("grafana:cloud/getProviderAwsCloudwatchScrapeJob:getProviderAwsCloudwatchScrapeJob", {
        "customNamespaces": args.customNamespaces,
        "name": args.name,
        "services": args.services,
        "stackId": args.stackId,
    }, opts);
}

/**
 * A collection of arguments for invoking getProviderAwsCloudwatchScrapeJob.
 */
export interface GetProviderAwsCloudwatchScrapeJobOutputArgs {
    /**
     * Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
     */
    customNamespaces?: pulumi.Input<pulumi.Input<inputs.cloud.GetProviderAwsCloudwatchScrapeJobCustomNamespaceArgs>[]>;
    name: pulumi.Input<string>;
    /**
     * One or more configuration blocks to dictate what this CloudWatch Scrape Job should scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
     */
    services?: pulumi.Input<pulumi.Input<inputs.cloud.GetProviderAwsCloudwatchScrapeJobServiceArgs>[]>;
    stackId: pulumi.Input<string>;
}
