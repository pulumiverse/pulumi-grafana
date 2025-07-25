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
    public static class GetTeam
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
        ///     var myTeam = Grafana.Oss.GetTeam.Invoke(new()
        ///     {
        ///         Name = "my team",
        ///     });
        /// 
        ///     var myTeamGetTeam = Grafana.OnCall.GetTeam.Invoke(new()
        ///     {
        ///         Name = myTeam.Apply(getTeamResult =&gt; getTeamResult.Name),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetTeamResult> InvokeAsync(GetTeamArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTeamResult>("grafana:onCall/getTeam:getTeam", args ?? new GetTeamArgs(), options.WithDefaults());

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
        ///     var myTeam = Grafana.Oss.GetTeam.Invoke(new()
        ///     {
        ///         Name = "my team",
        ///     });
        /// 
        ///     var myTeamGetTeam = Grafana.OnCall.GetTeam.Invoke(new()
        ///     {
        ///         Name = myTeam.Apply(getTeamResult =&gt; getTeamResult.Name),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetTeamResult> Invoke(GetTeamInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTeamResult>("grafana:onCall/getTeam:getTeam", args ?? new GetTeamInvokeArgs(), options.WithDefaults());

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
        ///     var myTeam = Grafana.Oss.GetTeam.Invoke(new()
        ///     {
        ///         Name = "my team",
        ///     });
        /// 
        ///     var myTeamGetTeam = Grafana.OnCall.GetTeam.Invoke(new()
        ///     {
        ///         Name = myTeam.Apply(getTeamResult =&gt; getTeamResult.Name),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetTeamResult> Invoke(GetTeamInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTeamResult>("grafana:onCall/getTeam:getTeam", args ?? new GetTeamInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTeamArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The team name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetTeamArgs()
        {
        }
        public static new GetTeamArgs Empty => new GetTeamArgs();
    }

    public sealed class GetTeamInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The team name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetTeamInvokeArgs()
        {
        }
        public static new GetTeamInvokeArgs Empty => new GetTeamInvokeArgs();
    }


    [OutputType]
    public sealed class GetTeamResult
    {
        public readonly string AvatarUrl;
        public readonly string Email;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The team name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetTeamResult(
            string avatarUrl,

            string email,

            string id,

            string name)
        {
            AvatarUrl = avatarUrl;
            Email = email;
            Id = id;
            Name = name;
        }
    }
}
