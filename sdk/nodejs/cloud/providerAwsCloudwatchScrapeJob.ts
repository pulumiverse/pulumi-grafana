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
 * const testProviderAwsAccount = new grafana.cloud.ProviderAwsAccount("test", {
 *     stackId: test.then(test => test.id),
 *     roleArn: testGetRole.then(testGetRole => testGetRole.arn),
 *     regions: [
 *         "us-east-1",
 *         "us-east-2",
 *         "us-west-1",
 *     ],
 * });
 * const testProviderAwsCloudwatchScrapeJob = new grafana.cloud.ProviderAwsCloudwatchScrapeJob("test", {
 *     stackId: test.then(test => test.id),
 *     name: "my-cloudwatch-scrape-job",
 *     awsAccountResourceId: testProviderAwsAccount.resourceId,
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
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:cloud/providerAwsCloudwatchScrapeJob:ProviderAwsCloudwatchScrapeJob name "{{ stack_id }}:{{ name }}"
 * ```
 */
export class ProviderAwsCloudwatchScrapeJob extends pulumi.CustomResource {
    /**
     * Get an existing ProviderAwsCloudwatchScrapeJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProviderAwsCloudwatchScrapeJobState, opts?: pulumi.CustomResourceOptions): ProviderAwsCloudwatchScrapeJob {
        return new ProviderAwsCloudwatchScrapeJob(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:cloud/providerAwsCloudwatchScrapeJob:ProviderAwsCloudwatchScrapeJob';

    /**
     * Returns true if the given object is an instance of ProviderAwsCloudwatchScrapeJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProviderAwsCloudwatchScrapeJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProviderAwsCloudwatchScrapeJob.__pulumiType;
    }

    /**
     * The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resourceId` attribute of the `grafana.cloud.ProviderAwsAccount` resource.
     */
    public readonly awsAccountResourceId!: pulumi.Output<string>;
    /**
     * Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
     */
    public readonly customNamespaces!: pulumi.Output<outputs.cloud.ProviderAwsCloudwatchScrapeJobCustomNamespace[] | undefined>;
    /**
     * When the CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
     */
    public /*out*/ readonly disabledReason!: pulumi.Output<string>;
    /**
     * Whether the CloudWatch Scrape Job is enabled or not.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`.
     */
    public readonly exportTags!: pulumi.Output<boolean>;
    public readonly name!: pulumi.Output<string>;
    /**
     * A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
     */
    public readonly regionsSubsetOverrides!: pulumi.Output<string[]>;
    /**
     * One or more configuration blocks to configure AWS services for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
     */
    public readonly services!: pulumi.Output<outputs.cloud.ProviderAwsCloudwatchScrapeJobService[] | undefined>;
    public readonly stackId!: pulumi.Output<string>;

    /**
     * Create a ProviderAwsCloudwatchScrapeJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderAwsCloudwatchScrapeJobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProviderAwsCloudwatchScrapeJobArgs | ProviderAwsCloudwatchScrapeJobState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProviderAwsCloudwatchScrapeJobState | undefined;
            resourceInputs["awsAccountResourceId"] = state ? state.awsAccountResourceId : undefined;
            resourceInputs["customNamespaces"] = state ? state.customNamespaces : undefined;
            resourceInputs["disabledReason"] = state ? state.disabledReason : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["exportTags"] = state ? state.exportTags : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["regionsSubsetOverrides"] = state ? state.regionsSubsetOverrides : undefined;
            resourceInputs["services"] = state ? state.services : undefined;
            resourceInputs["stackId"] = state ? state.stackId : undefined;
        } else {
            const args = argsOrState as ProviderAwsCloudwatchScrapeJobArgs | undefined;
            if ((!args || args.awsAccountResourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'awsAccountResourceId'");
            }
            if ((!args || args.stackId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stackId'");
            }
            resourceInputs["awsAccountResourceId"] = args ? args.awsAccountResourceId : undefined;
            resourceInputs["customNamespaces"] = args ? args.customNamespaces : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["exportTags"] = args ? args.exportTags : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["regionsSubsetOverrides"] = args ? args.regionsSubsetOverrides : undefined;
            resourceInputs["services"] = args ? args.services : undefined;
            resourceInputs["stackId"] = args ? args.stackId : undefined;
            resourceInputs["disabledReason"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProviderAwsCloudwatchScrapeJob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProviderAwsCloudwatchScrapeJob resources.
 */
export interface ProviderAwsCloudwatchScrapeJobState {
    /**
     * The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resourceId` attribute of the `grafana.cloud.ProviderAwsAccount` resource.
     */
    awsAccountResourceId?: pulumi.Input<string>;
    /**
     * Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
     */
    customNamespaces?: pulumi.Input<pulumi.Input<inputs.cloud.ProviderAwsCloudwatchScrapeJobCustomNamespace>[]>;
    /**
     * When the CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
     */
    disabledReason?: pulumi.Input<string>;
    /**
     * Whether the CloudWatch Scrape Job is enabled or not.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`.
     */
    exportTags?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    /**
     * A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
     */
    regionsSubsetOverrides?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * One or more configuration blocks to configure AWS services for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
     */
    services?: pulumi.Input<pulumi.Input<inputs.cloud.ProviderAwsCloudwatchScrapeJobService>[]>;
    stackId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProviderAwsCloudwatchScrapeJob resource.
 */
export interface ProviderAwsCloudwatchScrapeJobArgs {
    /**
     * The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resourceId` attribute of the `grafana.cloud.ProviderAwsAccount` resource.
     */
    awsAccountResourceId: pulumi.Input<string>;
    /**
     * Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
     */
    customNamespaces?: pulumi.Input<pulumi.Input<inputs.cloud.ProviderAwsCloudwatchScrapeJobCustomNamespace>[]>;
    /**
     * Whether the CloudWatch Scrape Job is enabled or not.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`.
     */
    exportTags?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    /**
     * A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
     */
    regionsSubsetOverrides?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * One or more configuration blocks to configure AWS services for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
     */
    services?: pulumi.Input<pulumi.Input<inputs.cloud.ProviderAwsCloudwatchScrapeJobService>[]>;
    stackId: pulumi.Input<string>;
}