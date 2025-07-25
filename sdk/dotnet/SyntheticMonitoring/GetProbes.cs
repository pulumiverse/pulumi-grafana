// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.SyntheticMonitoring
{
    public static class GetProbes
    {
        /// <summary>
        /// Data source for retrieving all probes.
        /// 
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
        ///     var main = Grafana.SyntheticMonitoring.GetProbes.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetProbesResult> InvokeAsync(GetProbesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProbesResult>("grafana:syntheticMonitoring/getProbes:getProbes", args ?? new GetProbesArgs(), options.WithDefaults());

        /// <summary>
        /// Data source for retrieving all probes.
        /// 
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
        ///     var main = Grafana.SyntheticMonitoring.GetProbes.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetProbesResult> Invoke(GetProbesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProbesResult>("grafana:syntheticMonitoring/getProbes:getProbes", args ?? new GetProbesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Data source for retrieving all probes.
        /// 
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
        ///     var main = Grafana.SyntheticMonitoring.GetProbes.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetProbesResult> Invoke(GetProbesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetProbesResult>("grafana:syntheticMonitoring/getProbes:getProbes", args ?? new GetProbesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProbesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// If true, only probes that are not deprecated will be returned. Defaults to `true`.
        /// </summary>
        [Input("filterDeprecated")]
        public bool? FilterDeprecated { get; set; }

        public GetProbesArgs()
        {
        }
        public static new GetProbesArgs Empty => new GetProbesArgs();
    }

    public sealed class GetProbesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// If true, only probes that are not deprecated will be returned. Defaults to `true`.
        /// </summary>
        [Input("filterDeprecated")]
        public Input<bool>? FilterDeprecated { get; set; }

        public GetProbesInvokeArgs()
        {
        }
        public static new GetProbesInvokeArgs Empty => new GetProbesInvokeArgs();
    }


    [OutputType]
    public sealed class GetProbesResult
    {
        /// <summary>
        /// If true, only probes that are not deprecated will be returned. Defaults to `true`.
        /// </summary>
        public readonly bool? FilterDeprecated;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Map of probes with their names as keys and IDs as values.
        /// </summary>
        public readonly ImmutableDictionary<string, int> Probes;

        [OutputConstructor]
        private GetProbesResult(
            bool? filterDeprecated,

            string id,

            ImmutableDictionary<string, int> probes)
        {
            FilterDeprecated = filterDeprecated;
            Id = id;
            Probes = probes;
        }
    }
}
