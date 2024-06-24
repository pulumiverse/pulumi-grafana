// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.OnCall
{
    public static class GetUser
    {
        /// <summary>
        /// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/users/)
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
        ///     var alex = Grafana.OnCall.GetUser.Invoke(new()
        ///     {
        ///         Username = "alex",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetUserResult> InvokeAsync(GetUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("grafana:onCall/getUser:getUser", args ?? new GetUserArgs(), options.WithDefaults());

        /// <summary>
        /// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/users/)
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
        ///     var alex = Grafana.OnCall.GetUser.Invoke(new()
        ///     {
        ///         Username = "alex",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("grafana:onCall/getUser:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The username of the user.
        /// </summary>
        [Input("username", required: true)]
        public string Username { get; set; } = null!;

        public GetUserArgs()
        {
        }
        public static new GetUserArgs Empty => new GetUserArgs();
    }

    public sealed class GetUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The username of the user.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public GetUserInvokeArgs()
        {
        }
        public static new GetUserInvokeArgs Empty => new GetUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// The email of the user.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The role of the user.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// The username of the user.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GetUserResult(
            string email,

            string id,

            string role,

            string username)
        {
            Email = email;
            Id = id;
            Role = role;
            Username = username;
        }
    }
}