// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.CloudProvider
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
    ///     var testAwsAccount = new Grafana.CloudProvider.AwsAccount("test", new()
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
    ///     var testAwsCloudwatchScrapeJob = new Grafana.CloudProvider.AwsCloudwatchScrapeJob("test", new()
    ///     {
    ///         StackId = test.Apply(getStackResult =&gt; getStackResult.Id),
    ///         Name = "my-cloudwatch-scrape-job",
    ///         AwsAccountResourceId = testAwsAccount.ResourceId,
    ///         ExportTags = true,
    ///         Services = new[]
    ///         {
    ///             new Grafana.CloudProvider.Inputs.AwsCloudwatchScrapeJobServiceArgs
    ///             {
    ///                 Name = "AWS/EC2",
    ///                 Metrics = new[]
    ///                 {
    ///                     new Grafana.CloudProvider.Inputs.AwsCloudwatchScrapeJobServiceMetricArgs
    ///                     {
    ///                         Name = "CPUUtilization",
    ///                         Statistics = new[]
    ///                         {
    ///                             "Average",
    ///                         },
    ///                     },
    ///                     new Grafana.CloudProvider.Inputs.AwsCloudwatchScrapeJobServiceMetricArgs
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
    ///                     new Grafana.CloudProvider.Inputs.AwsCloudwatchScrapeJobServiceResourceDiscoveryTagFilterArgs
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
    ///             new Grafana.CloudProvider.Inputs.AwsCloudwatchScrapeJobCustomNamespaceArgs
    ///             {
    ///                 Name = "CoolApp",
    ///                 Metrics = new[]
    ///                 {
    ///                     new Grafana.CloudProvider.Inputs.AwsCloudwatchScrapeJobCustomNamespaceMetricArgs
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
    ///         StaticLabels = 
    ///         {
    ///             { "label1", "value1" },
    ///             { "label2", "value2" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:cloudProvider/awsCloudwatchScrapeJob:AwsCloudwatchScrapeJob name "{{ stack_id }}:{{ name }}"
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:cloudProvider/awsCloudwatchScrapeJob:AwsCloudwatchScrapeJob")]
    public partial class AwsCloudwatchScrapeJob : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resource_id` attribute of the `grafana.cloudProvider.AwsAccount` resource.
        /// </summary>
        [Output("awsAccountResourceId")]
        public Output<string> AwsAccountResourceId { get; private set; } = null!;

        /// <summary>
        /// Zero or more configuration blocks to configure custom namespaces for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        [Output("customNamespaces")]
        public Output<ImmutableArray<Outputs.AwsCloudwatchScrapeJobCustomNamespace>> CustomNamespaces { get; private set; } = null!;

        /// <summary>
        /// When the AWS CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
        /// </summary>
        [Output("disabledReason")]
        public Output<string> DisabledReason { get; private set; } = null!;

        /// <summary>
        /// Whether the AWS CloudWatch Scrape Job is enabled or not. Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_&lt;service_name&gt;_info`. Defaults to `true`.
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
        /// One or more configuration blocks to configure AWS services for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        [Output("services")]
        public Output<ImmutableArray<Outputs.AwsCloudwatchScrapeJobService>> Services { get; private set; } = null!;

        [Output("stackId")]
        public Output<string> StackId { get; private set; } = null!;

        /// <summary>
        /// A set of static labels to add to all metrics exported by this scrape job.
        /// </summary>
        [Output("staticLabels")]
        public Output<ImmutableDictionary<string, string>> StaticLabels { get; private set; } = null!;


        /// <summary>
        /// Create a AwsCloudwatchScrapeJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AwsCloudwatchScrapeJob(string name, AwsCloudwatchScrapeJobArgs args, CustomResourceOptions? options = null)
            : base("grafana:cloudProvider/awsCloudwatchScrapeJob:AwsCloudwatchScrapeJob", name, args ?? new AwsCloudwatchScrapeJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AwsCloudwatchScrapeJob(string name, Input<string> id, AwsCloudwatchScrapeJobState? state = null, CustomResourceOptions? options = null)
            : base("grafana:cloudProvider/awsCloudwatchScrapeJob:AwsCloudwatchScrapeJob", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "grafana:cloud/providerAwsCloudwatchScrapeJob:ProviderAwsCloudwatchScrapeJob" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AwsCloudwatchScrapeJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AwsCloudwatchScrapeJob Get(string name, Input<string> id, AwsCloudwatchScrapeJobState? state = null, CustomResourceOptions? options = null)
        {
            return new AwsCloudwatchScrapeJob(name, id, state, options);
        }
    }

    public sealed class AwsCloudwatchScrapeJobArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resource_id` attribute of the `grafana.cloudProvider.AwsAccount` resource.
        /// </summary>
        [Input("awsAccountResourceId", required: true)]
        public Input<string> AwsAccountResourceId { get; set; } = null!;

        [Input("customNamespaces")]
        private InputList<Inputs.AwsCloudwatchScrapeJobCustomNamespaceArgs>? _customNamespaces;

        /// <summary>
        /// Zero or more configuration blocks to configure custom namespaces for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public InputList<Inputs.AwsCloudwatchScrapeJobCustomNamespaceArgs> CustomNamespaces
        {
            get => _customNamespaces ?? (_customNamespaces = new InputList<Inputs.AwsCloudwatchScrapeJobCustomNamespaceArgs>());
            set => _customNamespaces = value;
        }

        /// <summary>
        /// Whether the AWS CloudWatch Scrape Job is enabled or not. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_&lt;service_name&gt;_info`. Defaults to `true`.
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
        private InputList<Inputs.AwsCloudwatchScrapeJobServiceArgs>? _services;

        /// <summary>
        /// One or more configuration blocks to configure AWS services for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public InputList<Inputs.AwsCloudwatchScrapeJobServiceArgs> Services
        {
            get => _services ?? (_services = new InputList<Inputs.AwsCloudwatchScrapeJobServiceArgs>());
            set => _services = value;
        }

        [Input("stackId", required: true)]
        public Input<string> StackId { get; set; } = null!;

        [Input("staticLabels")]
        private InputMap<string>? _staticLabels;

        /// <summary>
        /// A set of static labels to add to all metrics exported by this scrape job.
        /// </summary>
        public InputMap<string> StaticLabels
        {
            get => _staticLabels ?? (_staticLabels = new InputMap<string>());
            set => _staticLabels = value;
        }

        public AwsCloudwatchScrapeJobArgs()
        {
        }
        public static new AwsCloudwatchScrapeJobArgs Empty => new AwsCloudwatchScrapeJobArgs();
    }

    public sealed class AwsCloudwatchScrapeJobState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resource_id` attribute of the `grafana.cloudProvider.AwsAccount` resource.
        /// </summary>
        [Input("awsAccountResourceId")]
        public Input<string>? AwsAccountResourceId { get; set; }

        [Input("customNamespaces")]
        private InputList<Inputs.AwsCloudwatchScrapeJobCustomNamespaceGetArgs>? _customNamespaces;

        /// <summary>
        /// Zero or more configuration blocks to configure custom namespaces for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public InputList<Inputs.AwsCloudwatchScrapeJobCustomNamespaceGetArgs> CustomNamespaces
        {
            get => _customNamespaces ?? (_customNamespaces = new InputList<Inputs.AwsCloudwatchScrapeJobCustomNamespaceGetArgs>());
            set => _customNamespaces = value;
        }

        /// <summary>
        /// When the AWS CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
        /// </summary>
        [Input("disabledReason")]
        public Input<string>? DisabledReason { get; set; }

        /// <summary>
        /// Whether the AWS CloudWatch Scrape Job is enabled or not. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_&lt;service_name&gt;_info`. Defaults to `true`.
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
        private InputList<Inputs.AwsCloudwatchScrapeJobServiceGetArgs>? _services;

        /// <summary>
        /// One or more configuration blocks to configure AWS services for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public InputList<Inputs.AwsCloudwatchScrapeJobServiceGetArgs> Services
        {
            get => _services ?? (_services = new InputList<Inputs.AwsCloudwatchScrapeJobServiceGetArgs>());
            set => _services = value;
        }

        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        [Input("staticLabels")]
        private InputMap<string>? _staticLabels;

        /// <summary>
        /// A set of static labels to add to all metrics exported by this scrape job.
        /// </summary>
        public InputMap<string> StaticLabels
        {
            get => _staticLabels ?? (_staticLabels = new InputMap<string>());
            set => _staticLabels = value;
        }

        public AwsCloudwatchScrapeJobState()
        {
        }
        public static new AwsCloudwatchScrapeJobState Empty => new AwsCloudwatchScrapeJobState();
    }
}
