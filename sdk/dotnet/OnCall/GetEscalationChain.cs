// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.OnCall
{
    public static class GetEscalationChain
    {
        /// <summary>
        /// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/escalation_chains/)
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
        ///     var @default = Grafana.OnCall.GetEscalationChain.Invoke(new()
        ///     {
        ///         Name = "default",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetEscalationChainResult> InvokeAsync(GetEscalationChainArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEscalationChainResult>("grafana:onCall/getEscalationChain:getEscalationChain", args ?? new GetEscalationChainArgs(), options.WithDefaults());

        /// <summary>
        /// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/escalation_chains/)
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
        ///     var @default = Grafana.OnCall.GetEscalationChain.Invoke(new()
        ///     {
        ///         Name = "default",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetEscalationChainResult> Invoke(GetEscalationChainInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEscalationChainResult>("grafana:onCall/getEscalationChain:getEscalationChain", args ?? new GetEscalationChainInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/escalation_chains/)
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
        ///     var @default = Grafana.OnCall.GetEscalationChain.Invoke(new()
        ///     {
        ///         Name = "default",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetEscalationChainResult> Invoke(GetEscalationChainInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetEscalationChainResult>("grafana:onCall/getEscalationChain:getEscalationChain", args ?? new GetEscalationChainInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEscalationChainArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The escalation chain name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetEscalationChainArgs()
        {
        }
        public static new GetEscalationChainArgs Empty => new GetEscalationChainArgs();
    }

    public sealed class GetEscalationChainInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The escalation chain name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetEscalationChainInvokeArgs()
        {
        }
        public static new GetEscalationChainInvokeArgs Empty => new GetEscalationChainInvokeArgs();
    }


    [OutputType]
    public sealed class GetEscalationChainResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The escalation chain name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetEscalationChainResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
