// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Cloud
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// using Grafana = Pulumi.Grafana;
    /// using Grafana = Pulumiverse.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = Grafana.Cloud.GetStack.Invoke(new()
    ///     {
    ///         Slug = "gcloudstacktest",
    ///     });
    /// 
    ///     var testGetRole = Aws.Iam.GetRole.Invoke(new()
    ///     {
    ///         Name = "my-role",
    ///     });
    /// 
    ///     var testProviderAwsAccount = new Grafana.Cloud.ProviderAwsAccount("test", new()
    ///     {
    ///         StackId = test.Apply(getStackResult =&gt; getStackResult.Id),
    ///         RoleArn = testGetRole.Apply(getRoleResult =&gt; getRoleResult.Arn),
    ///         Regions = new[]
    ///         {
    ///             "us-east-1",
    ///             "us-east-2",
    ///             "us-west-1",
    ///         },
    ///     });
    /// 
    ///     var testProviderAwsCloudwatchScrapeJob = new Grafana.Cloud.ProviderAwsCloudwatchScrapeJob("test", new()
    ///     {
    ///         StackId = test.Apply(getStackResult =&gt; getStackResult.Id),
    ///         Name = "my-cloudwatch-scrape-job",
    ///         AwsAccountResourceId = testProviderAwsAccount.ResourceId,
    ///         ExportTags = true,
    ///         Services = new[]
    ///         {
    ///             new Grafana.Cloud.Inputs.ProviderAwsCloudwatchScrapeJobServiceArgs
    ///             {
    ///                 Name = "AWS/EC2",
    ///                 Metrics = new[]
    ///                 {
    ///                     new Grafana.Cloud.Inputs.ProviderAwsCloudwatchScrapeJobServiceMetricArgs
    ///                     {
    ///                         Name = "CPUUtilization",
    ///                         Statistics = new[]
    ///                         {
    ///                             "Average",
    ///                         },
    ///                     },
    ///                     new Grafana.Cloud.Inputs.ProviderAwsCloudwatchScrapeJobServiceMetricArgs
    ///                     {
    ///                         Name = "StatusCheckFailed",
    ///                         Statistics = new[]
    ///                         {
    ///                             "Maximum",
    ///                         },
    ///                     },
    ///                 },
    ///                 ScrapeIntervalSeconds = 300,
    ///                 ResourceDiscoveryTagFilters = new[]
    ///                 {
    ///                     new Grafana.Cloud.Inputs.ProviderAwsCloudwatchScrapeJobServiceResourceDiscoveryTagFilterArgs
    ///                     {
    ///                         Key = "k8s.io/cluster-autoscaler/enabled",
    ///                         Value = "true",
    ///                     },
    ///                 },
    ///                 TagsToAddToMetrics = new[]
    ///                 {
    ///                     "eks:cluster-name",
    ///                 },
    ///             },
    ///         },
    ///         CustomNamespaces = new[]
    ///         {
    ///             new Grafana.Cloud.Inputs.ProviderAwsCloudwatchScrapeJobCustomNamespaceArgs
    ///             {
    ///                 Name = "CoolApp",
    ///                 Metrics = new[]
    ///                 {
    ///                     new Grafana.Cloud.Inputs.ProviderAwsCloudwatchScrapeJobCustomNamespaceMetricArgs
    ///                     {
    ///                         Name = "CoolMetric",
    ///                         Statistics = new[]
    ///                         {
    ///                             "Maximum",
    ///                             "Sum",
    ///                         },
    ///                     },
    ///                 },
    ///                 ScrapeIntervalSeconds = 300,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:cloud/providerAwsCloudwatchScrapeJob:ProviderAwsCloudwatchScrapeJob name "{{ stack_id }}:{{ name }}"
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:cloud/providerAwsCloudwatchScrapeJob:ProviderAwsCloudwatchScrapeJob")]
    public partial class ProviderAwsCloudwatchScrapeJob : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resource_id` attribute of the `grafana.cloud.ProviderAwsAccount` resource.
        /// </summary>
        [Output("awsAccountResourceId")]
        public Output<string> AwsAccountResourceId { get; private set; } = null!;

        /// <summary>
        /// Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        [Output("customNamespaces")]
        public Output<ImmutableArray<Outputs.ProviderAwsCloudwatchScrapeJobCustomNamespace>> CustomNamespaces { get; private set; } = null!;

        /// <summary>
        /// When the CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
        /// </summary>
        [Output("disabledReason")]
        public Output<string> DisabledReason { get; private set; } = null!;

        /// <summary>
        /// Whether the CloudWatch Scrape Job is enabled or not.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_&lt;service_name&gt;_info`.
        /// </summary>
        [Output("exportTags")]
        public Output<bool> ExportTags { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
        /// </summary>
        [Output("regionsSubsetOverrides")]
        public Output<ImmutableArray<string>> RegionsSubsetOverrides { get; private set; } = null!;

        /// <summary>
        /// One or more configuration blocks to configure AWS services for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        [Output("services")]
        public Output<ImmutableArray<Outputs.ProviderAwsCloudwatchScrapeJobService>> Services { get; private set; } = null!;

        [Output("stackId")]
        public Output<string> StackId { get; private set; } = null!;


        /// <summary>
        /// Create a ProviderAwsCloudwatchScrapeJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProviderAwsCloudwatchScrapeJob(string name, ProviderAwsCloudwatchScrapeJobArgs args, CustomResourceOptions? options = null)
            : base("grafana:cloud/providerAwsCloudwatchScrapeJob:ProviderAwsCloudwatchScrapeJob", name, args ?? new ProviderAwsCloudwatchScrapeJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProviderAwsCloudwatchScrapeJob(string name, Input<string> id, ProviderAwsCloudwatchScrapeJobState? state = null, CustomResourceOptions? options = null)
            : base("grafana:cloud/providerAwsCloudwatchScrapeJob:ProviderAwsCloudwatchScrapeJob", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProviderAwsCloudwatchScrapeJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProviderAwsCloudwatchScrapeJob Get(string name, Input<string> id, ProviderAwsCloudwatchScrapeJobState? state = null, CustomResourceOptions? options = null)
        {
            return new ProviderAwsCloudwatchScrapeJob(name, id, state, options);
        }
    }

    public sealed class ProviderAwsCloudwatchScrapeJobArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resource_id` attribute of the `grafana.cloud.ProviderAwsAccount` resource.
        /// </summary>
        [Input("awsAccountResourceId", required: true)]
        public Input<string> AwsAccountResourceId { get; set; } = null!;

        [Input("customNamespaces")]
        private InputList<Inputs.ProviderAwsCloudwatchScrapeJobCustomNamespaceArgs>? _customNamespaces;

        /// <summary>
        /// Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public InputList<Inputs.ProviderAwsCloudwatchScrapeJobCustomNamespaceArgs> CustomNamespaces
        {
            get => _customNamespaces ?? (_customNamespaces = new InputList<Inputs.ProviderAwsCloudwatchScrapeJobCustomNamespaceArgs>());
            set => _customNamespaces = value;
        }

        /// <summary>
        /// Whether the CloudWatch Scrape Job is enabled or not.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_&lt;service_name&gt;_info`.
        /// </summary>
        [Input("exportTags")]
        public Input<bool>? ExportTags { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("regionsSubsetOverrides")]
        private InputList<string>? _regionsSubsetOverrides;

        /// <summary>
        /// A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
        /// </summary>
        public InputList<string> RegionsSubsetOverrides
        {
            get => _regionsSubsetOverrides ?? (_regionsSubsetOverrides = new InputList<string>());
            set => _regionsSubsetOverrides = value;
        }

        [Input("services")]
        private InputList<Inputs.ProviderAwsCloudwatchScrapeJobServiceArgs>? _services;

        /// <summary>
        /// One or more configuration blocks to configure AWS services for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public InputList<Inputs.ProviderAwsCloudwatchScrapeJobServiceArgs> Services
        {
            get => _services ?? (_services = new InputList<Inputs.ProviderAwsCloudwatchScrapeJobServiceArgs>());
            set => _services = value;
        }

        [Input("stackId", required: true)]
        public Input<string> StackId { get; set; } = null!;

        public ProviderAwsCloudwatchScrapeJobArgs()
        {
        }
        public static new ProviderAwsCloudwatchScrapeJobArgs Empty => new ProviderAwsCloudwatchScrapeJobArgs();
    }

    public sealed class ProviderAwsCloudwatchScrapeJobState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resource_id` attribute of the `grafana.cloud.ProviderAwsAccount` resource.
        /// </summary>
        [Input("awsAccountResourceId")]
        public Input<string>? AwsAccountResourceId { get; set; }

        [Input("customNamespaces")]
        private InputList<Inputs.ProviderAwsCloudwatchScrapeJobCustomNamespaceGetArgs>? _customNamespaces;

        /// <summary>
        /// Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public InputList<Inputs.ProviderAwsCloudwatchScrapeJobCustomNamespaceGetArgs> CustomNamespaces
        {
            get => _customNamespaces ?? (_customNamespaces = new InputList<Inputs.ProviderAwsCloudwatchScrapeJobCustomNamespaceGetArgs>());
            set => _customNamespaces = value;
        }

        /// <summary>
        /// When the CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
        /// </summary>
        [Input("disabledReason")]
        public Input<string>? DisabledReason { get; set; }

        /// <summary>
        /// Whether the CloudWatch Scrape Job is enabled or not.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_&lt;service_name&gt;_info`.
        /// </summary>
        [Input("exportTags")]
        public Input<bool>? ExportTags { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("regionsSubsetOverrides")]
        private InputList<string>? _regionsSubsetOverrides;

        /// <summary>
        /// A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
        /// </summary>
        public InputList<string> RegionsSubsetOverrides
        {
            get => _regionsSubsetOverrides ?? (_regionsSubsetOverrides = new InputList<string>());
            set => _regionsSubsetOverrides = value;
        }

        [Input("services")]
        private InputList<Inputs.ProviderAwsCloudwatchScrapeJobServiceGetArgs>? _services;

        /// <summary>
        /// One or more configuration blocks to configure AWS services for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public InputList<Inputs.ProviderAwsCloudwatchScrapeJobServiceGetArgs> Services
        {
            get => _services ?? (_services = new InputList<Inputs.ProviderAwsCloudwatchScrapeJobServiceGetArgs>());
            set => _services = value;
        }

        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        public ProviderAwsCloudwatchScrapeJobState()
        {
        }
        public static new ProviderAwsCloudwatchScrapeJobState Empty => new ProviderAwsCloudwatchScrapeJobState();
    }
}
