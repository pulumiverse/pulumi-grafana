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
    public static class GetProviderAwsCloudwatchScrapeJobs
    {
        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Grafana = Pulumi.Grafana;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test = Grafana.Cloud.GetStack.Invoke(new()
        ///     {
        ///         Slug = "gcloudstacktest",
        ///     });
        /// 
        ///     var testGetProviderAwsCloudwatchScrapeJobs = Grafana.Cloud.GetProviderAwsCloudwatchScrapeJobs.Invoke(new()
        ///     {
        ///         StackId = test.Apply(getStackResult =&gt; getStackResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetProviderAwsCloudwatchScrapeJobsResult> InvokeAsync(GetProviderAwsCloudwatchScrapeJobsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProviderAwsCloudwatchScrapeJobsResult>("grafana:cloud/getProviderAwsCloudwatchScrapeJobs:getProviderAwsCloudwatchScrapeJobs", args ?? new GetProviderAwsCloudwatchScrapeJobsArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Grafana = Pulumi.Grafana;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test = Grafana.Cloud.GetStack.Invoke(new()
        ///     {
        ///         Slug = "gcloudstacktest",
        ///     });
        /// 
        ///     var testGetProviderAwsCloudwatchScrapeJobs = Grafana.Cloud.GetProviderAwsCloudwatchScrapeJobs.Invoke(new()
        ///     {
        ///         StackId = test.Apply(getStackResult =&gt; getStackResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetProviderAwsCloudwatchScrapeJobsResult> Invoke(GetProviderAwsCloudwatchScrapeJobsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProviderAwsCloudwatchScrapeJobsResult>("grafana:cloud/getProviderAwsCloudwatchScrapeJobs:getProviderAwsCloudwatchScrapeJobs", args ?? new GetProviderAwsCloudwatchScrapeJobsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProviderAwsCloudwatchScrapeJobsArgs : global::Pulumi.InvokeArgs
    {
        [Input("scrapeJobs")]
        private List<Inputs.GetProviderAwsCloudwatchScrapeJobsScrapeJobArgs>? _scrapeJobs;

        /// <summary>
        /// A list of AWS CloudWatch Scrape Job objects associated with the given StackID.
        /// </summary>
        public List<Inputs.GetProviderAwsCloudwatchScrapeJobsScrapeJobArgs> ScrapeJobs
        {
            get => _scrapeJobs ?? (_scrapeJobs = new List<Inputs.GetProviderAwsCloudwatchScrapeJobsScrapeJobArgs>());
            set => _scrapeJobs = value;
        }

        [Input("stackId", required: true)]
        public string StackId { get; set; } = null!;

        public GetProviderAwsCloudwatchScrapeJobsArgs()
        {
        }
        public static new GetProviderAwsCloudwatchScrapeJobsArgs Empty => new GetProviderAwsCloudwatchScrapeJobsArgs();
    }

    public sealed class GetProviderAwsCloudwatchScrapeJobsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("scrapeJobs")]
        private InputList<Inputs.GetProviderAwsCloudwatchScrapeJobsScrapeJobInputArgs>? _scrapeJobs;

        /// <summary>
        /// A list of AWS CloudWatch Scrape Job objects associated with the given StackID.
        /// </summary>
        public InputList<Inputs.GetProviderAwsCloudwatchScrapeJobsScrapeJobInputArgs> ScrapeJobs
        {
            get => _scrapeJobs ?? (_scrapeJobs = new InputList<Inputs.GetProviderAwsCloudwatchScrapeJobsScrapeJobInputArgs>());
            set => _scrapeJobs = value;
        }

        [Input("stackId", required: true)]
        public Input<string> StackId { get; set; } = null!;

        public GetProviderAwsCloudwatchScrapeJobsInvokeArgs()
        {
        }
        public static new GetProviderAwsCloudwatchScrapeJobsInvokeArgs Empty => new GetProviderAwsCloudwatchScrapeJobsInvokeArgs();
    }


    [OutputType]
    public sealed class GetProviderAwsCloudwatchScrapeJobsResult
    {
        public readonly string Id;
        /// <summary>
        /// A list of AWS CloudWatch Scrape Job objects associated with the given StackID.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProviderAwsCloudwatchScrapeJobsScrapeJobResult> ScrapeJobs;
        public readonly string StackId;

        [OutputConstructor]
        private GetProviderAwsCloudwatchScrapeJobsResult(
            string id,

            ImmutableArray<Outputs.GetProviderAwsCloudwatchScrapeJobsScrapeJobResult> scrapeJobs,

            string stackId)
        {
            Id = id;
            ScrapeJobs = scrapeJobs;
            StackId = stackId;
        }
    }
}