// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Connections
{
    public static class GetMetricsEndpointScrapeJob
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
        ///     var dsTest = Grafana.Connections.GetMetricsEndpointScrapeJob.Invoke(new()
        ///     {
        ///         StackId = "1",
        ///         Name = "my-scrape-job",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetMetricsEndpointScrapeJobResult> InvokeAsync(GetMetricsEndpointScrapeJobArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMetricsEndpointScrapeJobResult>("grafana:connections/getMetricsEndpointScrapeJob:getMetricsEndpointScrapeJob", args ?? new GetMetricsEndpointScrapeJobArgs(), options.WithDefaults());

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
        ///     var dsTest = Grafana.Connections.GetMetricsEndpointScrapeJob.Invoke(new()
        ///     {
        ///         StackId = "1",
        ///         Name = "my-scrape-job",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetMetricsEndpointScrapeJobResult> Invoke(GetMetricsEndpointScrapeJobInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMetricsEndpointScrapeJobResult>("grafana:connections/getMetricsEndpointScrapeJob:getMetricsEndpointScrapeJob", args ?? new GetMetricsEndpointScrapeJobInvokeArgs(), options.WithDefaults());

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
        ///     var dsTest = Grafana.Connections.GetMetricsEndpointScrapeJob.Invoke(new()
        ///     {
        ///         StackId = "1",
        ///         Name = "my-scrape-job",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetMetricsEndpointScrapeJobResult> Invoke(GetMetricsEndpointScrapeJobInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMetricsEndpointScrapeJobResult>("grafana:connections/getMetricsEndpointScrapeJob:getMetricsEndpointScrapeJob", args ?? new GetMetricsEndpointScrapeJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMetricsEndpointScrapeJobArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("stackId", required: true)]
        public string StackId { get; set; } = null!;

        public GetMetricsEndpointScrapeJobArgs()
        {
        }
        public static new GetMetricsEndpointScrapeJobArgs Empty => new GetMetricsEndpointScrapeJobArgs();
    }

    public sealed class GetMetricsEndpointScrapeJobInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("stackId", required: true)]
        public Input<string> StackId { get; set; } = null!;

        public GetMetricsEndpointScrapeJobInvokeArgs()
        {
        }
        public static new GetMetricsEndpointScrapeJobInvokeArgs Empty => new GetMetricsEndpointScrapeJobInvokeArgs();
    }


    [OutputType]
    public sealed class GetMetricsEndpointScrapeJobResult
    {
        /// <summary>
        /// Password for basic authentication.
        /// </summary>
        public readonly string AuthenticationBasicPassword;
        /// <summary>
        /// Username for basic authentication.
        /// </summary>
        public readonly string AuthenticationBasicUsername;
        /// <summary>
        /// Token for authentication bearer.
        /// </summary>
        public readonly string AuthenticationBearerToken;
        /// <summary>
        /// Method to pass authentication credentials: basic or bearer.
        /// </summary>
        public readonly string AuthenticationMethod;
        /// <summary>
        /// Whether the metrics endpoint scrape job is enabled or not.
        /// </summary>
        public readonly bool Enabled;
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// Frequency for scraping the metrics endpoint: 30, 60, or 120 seconds.
        /// </summary>
        public readonly int ScrapeIntervalSeconds;
        public readonly string StackId;
        /// <summary>
        /// The url to scrape metrics.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetMetricsEndpointScrapeJobResult(
            string authenticationBasicPassword,

            string authenticationBasicUsername,

            string authenticationBearerToken,

            string authenticationMethod,

            bool enabled,

            string id,

            string name,

            int scrapeIntervalSeconds,

            string stackId,

            string url)
        {
            AuthenticationBasicPassword = authenticationBasicPassword;
            AuthenticationBasicUsername = authenticationBasicUsername;
            AuthenticationBearerToken = authenticationBearerToken;
            AuthenticationMethod = authenticationMethod;
            Enabled = enabled;
            Id = id;
            Name = name;
            ScrapeIntervalSeconds = scrapeIntervalSeconds;
            StackId = stackId;
            Url = url;
        }
    }
}
