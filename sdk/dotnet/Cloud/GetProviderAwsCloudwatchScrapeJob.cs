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
    public static class GetProviderAwsCloudwatchScrapeJob
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
        ///     var testGetProviderAwsCloudwatchScrapeJob = Grafana.Cloud.GetProviderAwsCloudwatchScrapeJob.Invoke(new()
        ///     {
        ///         StackId = test.Apply(getStackResult =&gt; getStackResult.Id),
        ///         Name = testProviderAwsCloudwatchScrapeJob.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetProviderAwsCloudwatchScrapeJobResult> InvokeAsync(GetProviderAwsCloudwatchScrapeJobArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProviderAwsCloudwatchScrapeJobResult>("grafana:cloud/getProviderAwsCloudwatchScrapeJob:getProviderAwsCloudwatchScrapeJob", args ?? new GetProviderAwsCloudwatchScrapeJobArgs(), options.WithDefaults());

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
        ///     var testGetProviderAwsCloudwatchScrapeJob = Grafana.Cloud.GetProviderAwsCloudwatchScrapeJob.Invoke(new()
        ///     {
        ///         StackId = test.Apply(getStackResult =&gt; getStackResult.Id),
        ///         Name = testProviderAwsCloudwatchScrapeJob.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetProviderAwsCloudwatchScrapeJobResult> Invoke(GetProviderAwsCloudwatchScrapeJobInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProviderAwsCloudwatchScrapeJobResult>("grafana:cloud/getProviderAwsCloudwatchScrapeJob:getProviderAwsCloudwatchScrapeJob", args ?? new GetProviderAwsCloudwatchScrapeJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProviderAwsCloudwatchScrapeJobArgs : global::Pulumi.InvokeArgs
    {
        [Input("customNamespaces")]
        private List<Inputs.GetProviderAwsCloudwatchScrapeJobCustomNamespaceArgs>? _customNamespaces;

        /// <summary>
        /// Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public List<Inputs.GetProviderAwsCloudwatchScrapeJobCustomNamespaceArgs> CustomNamespaces
        {
            get => _customNamespaces ?? (_customNamespaces = new List<Inputs.GetProviderAwsCloudwatchScrapeJobCustomNamespaceArgs>());
            set => _customNamespaces = value;
        }

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("services")]
        private List<Inputs.GetProviderAwsCloudwatchScrapeJobServiceArgs>? _services;

        /// <summary>
        /// One or more configuration blocks to dictate what this CloudWatch Scrape Job should scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public List<Inputs.GetProviderAwsCloudwatchScrapeJobServiceArgs> Services
        {
            get => _services ?? (_services = new List<Inputs.GetProviderAwsCloudwatchScrapeJobServiceArgs>());
            set => _services = value;
        }

        [Input("stackId", required: true)]
        public string StackId { get; set; } = null!;

        public GetProviderAwsCloudwatchScrapeJobArgs()
        {
        }
        public static new GetProviderAwsCloudwatchScrapeJobArgs Empty => new GetProviderAwsCloudwatchScrapeJobArgs();
    }

    public sealed class GetProviderAwsCloudwatchScrapeJobInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("customNamespaces")]
        private InputList<Inputs.GetProviderAwsCloudwatchScrapeJobCustomNamespaceInputArgs>? _customNamespaces;

        /// <summary>
        /// Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public InputList<Inputs.GetProviderAwsCloudwatchScrapeJobCustomNamespaceInputArgs> CustomNamespaces
        {
            get => _customNamespaces ?? (_customNamespaces = new InputList<Inputs.GetProviderAwsCloudwatchScrapeJobCustomNamespaceInputArgs>());
            set => _customNamespaces = value;
        }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("services")]
        private InputList<Inputs.GetProviderAwsCloudwatchScrapeJobServiceInputArgs>? _services;

        /// <summary>
        /// One or more configuration blocks to dictate what this CloudWatch Scrape Job should scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public InputList<Inputs.GetProviderAwsCloudwatchScrapeJobServiceInputArgs> Services
        {
            get => _services ?? (_services = new InputList<Inputs.GetProviderAwsCloudwatchScrapeJobServiceInputArgs>());
            set => _services = value;
        }

        [Input("stackId", required: true)]
        public Input<string> StackId { get; set; } = null!;

        public GetProviderAwsCloudwatchScrapeJobInvokeArgs()
        {
        }
        public static new GetProviderAwsCloudwatchScrapeJobInvokeArgs Empty => new GetProviderAwsCloudwatchScrapeJobInvokeArgs();
    }


    [OutputType]
    public sealed class GetProviderAwsCloudwatchScrapeJobResult
    {
        /// <summary>
        /// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resource_id` attribute of the `grafana.cloud.ProviderAwsAccount` resource.
        /// </summary>
        public readonly string AwsAccountResourceId;
        /// <summary>
        /// Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProviderAwsCloudwatchScrapeJobCustomNamespaceResult> CustomNamespaces;
        /// <summary>
        /// When the CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
        /// </summary>
        public readonly string DisabledReason;
        /// <summary>
        /// Whether the CloudWatch Scrape Job is enabled or not.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_&lt;service_name&gt;_info`.
        /// </summary>
        public readonly bool ExportTags;
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// The set of AWS region names that this CloudWatch Scrape Job is configured to scrape.
        /// </summary>
        public readonly ImmutableArray<string> Regions;
        /// <summary>
        /// When true, the `regions` attribute will be the set of regions configured in the override. When false, the `regions` attribute will be the set of regions belonging to the AWS Account resource that is associated with this CloudWatch Scrape Job.
        /// </summary>
        public readonly bool RegionsSubsetOverrideUsed;
        /// <summary>
        /// The AWS ARN of the IAM role associated with the AWS Account resource that is being used by this CloudWatch Scrape Job.
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// One or more configuration blocks to dictate what this CloudWatch Scrape Job should scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProviderAwsCloudwatchScrapeJobServiceResult> Services;
        public readonly string StackId;

        [OutputConstructor]
        private GetProviderAwsCloudwatchScrapeJobResult(
            string awsAccountResourceId,

            ImmutableArray<Outputs.GetProviderAwsCloudwatchScrapeJobCustomNamespaceResult> customNamespaces,

            string disabledReason,

            bool enabled,

            bool exportTags,

            string id,

            string name,

            ImmutableArray<string> regions,

            bool regionsSubsetOverrideUsed,

            string roleArn,

            ImmutableArray<Outputs.GetProviderAwsCloudwatchScrapeJobServiceResult> services,

            string stackId)
        {
            AwsAccountResourceId = awsAccountResourceId;
            CustomNamespaces = customNamespaces;
            DisabledReason = disabledReason;
            Enabled = enabled;
            ExportTags = exportTags;
            Id = id;
            Name = name;
            Regions = regions;
            RegionsSubsetOverrideUsed = regionsSubsetOverrideUsed;
            RoleArn = roleArn;
            Services = services;
            StackId = stackId;
        }
    }
}
