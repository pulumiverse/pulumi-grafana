// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.CloudProvider.Inputs
{

    public sealed class GetAwsCloudwatchScrapeJobsScrapeJobArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resource_id` attribute of the `grafana.cloudProvider.AwsAccount` resource.
        /// </summary>
        [Input("awsAccountResourceId", required: true)]
        public string AwsAccountResourceId { get; set; } = null!;

        [Input("customNamespaces")]
        private List<Inputs.GetAwsCloudwatchScrapeJobsScrapeJobCustomNamespaceArgs>? _customNamespaces;

        /// <summary>
        /// Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public List<Inputs.GetAwsCloudwatchScrapeJobsScrapeJobCustomNamespaceArgs> CustomNamespaces
        {
            get => _customNamespaces ?? (_customNamespaces = new List<Inputs.GetAwsCloudwatchScrapeJobsScrapeJobCustomNamespaceArgs>());
            set => _customNamespaces = value;
        }

        /// <summary>
        /// When the CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
        /// </summary>
        [Input("disabledReason", required: true)]
        public string DisabledReason { get; set; } = null!;

        /// <summary>
        /// Whether the CloudWatch Scrape Job is enabled or not.
        /// </summary>
        [Input("enabled", required: true)]
        public bool Enabled { get; set; }

        /// <summary>
        /// When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_&lt;service_name&gt;_info`.
        /// </summary>
        [Input("exportTags", required: true)]
        public bool ExportTags { get; set; }

        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("regions", required: true)]
        private List<string>? _regions;

        /// <summary>
        /// The set of AWS region names that this CloudWatch Scrape Job is configured to scrape.
        /// </summary>
        public List<string> Regions
        {
            get => _regions ?? (_regions = new List<string>());
            set => _regions = value;
        }

        /// <summary>
        /// When true, the `regions` attribute will be the set of regions configured in the override. When false, the `regions` attribute will be the set of regions belonging to the AWS Account resource that is associated with this CloudWatch Scrape Job.
        /// </summary>
        [Input("regionsSubsetOverrideUsed", required: true)]
        public bool RegionsSubsetOverrideUsed { get; set; }

        /// <summary>
        /// The AWS ARN of the IAM role associated with the AWS Account resource that is being used by this CloudWatch Scrape Job.
        /// </summary>
        [Input("roleArn", required: true)]
        public string RoleArn { get; set; } = null!;

        [Input("services")]
        private List<Inputs.GetAwsCloudwatchScrapeJobsScrapeJobServiceArgs>? _services;

        /// <summary>
        /// One or more configuration blocks to dictate what this CloudWatch Scrape Job should scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public List<Inputs.GetAwsCloudwatchScrapeJobsScrapeJobServiceArgs> Services
        {
            get => _services ?? (_services = new List<Inputs.GetAwsCloudwatchScrapeJobsScrapeJobServiceArgs>());
            set => _services = value;
        }

        [Input("stackId", required: true)]
        public string StackId { get; set; } = null!;

        [Input("staticLabels", required: true)]
        private Dictionary<string, string>? _staticLabels;

        /// <summary>
        /// A set of static labels to add to all metrics exported by this scrape job.
        /// </summary>
        public Dictionary<string, string> StaticLabels
        {
            get => _staticLabels ?? (_staticLabels = new Dictionary<string, string>());
            set => _staticLabels = value;
        }

        public GetAwsCloudwatchScrapeJobsScrapeJobArgs()
        {
        }
        public static new GetAwsCloudwatchScrapeJobsScrapeJobArgs Empty => new GetAwsCloudwatchScrapeJobsScrapeJobArgs();
    }
}
