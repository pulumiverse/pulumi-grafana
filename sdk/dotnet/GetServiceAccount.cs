// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana
{
    [Obsolete(@"grafana.index/getserviceaccount.getServiceAccount has been deprecated in favor of grafana.oss/getserviceaccount.getServiceAccount")]
    public static class GetServiceAccount
    {
        /// <summary>
        /// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/)
        /// 		* [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api)
        /// </summary>
        public static Task<GetServiceAccountResult> InvokeAsync(GetServiceAccountArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceAccountResult>("grafana:index/getServiceAccount:getServiceAccount", args ?? new GetServiceAccountArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/)
        /// 		* [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api)
        /// </summary>
        public static Output<GetServiceAccountResult> Invoke(GetServiceAccountInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceAccountResult>("grafana:index/getServiceAccount:getServiceAccount", args ?? new GetServiceAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceAccountArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Service Account.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public string? OrgId { get; set; }

        public GetServiceAccountArgs()
        {
        }
        public static new GetServiceAccountArgs Empty => new GetServiceAccountArgs();
    }

    public sealed class GetServiceAccountInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Service Account.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        public GetServiceAccountInvokeArgs()
        {
        }
        public static new GetServiceAccountInvokeArgs Empty => new GetServiceAccountInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceAccountResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The disabled status for the service account.
        /// </summary>
        public readonly bool IsDisabled;
        /// <summary>
        /// The name of the Service Account.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        public readonly string? OrgId;
        /// <summary>
        /// The basic role of the service account in the organization.
        /// </summary>
        public readonly string Role;

        [OutputConstructor]
        private GetServiceAccountResult(
            string id,

            bool isDisabled,

            string name,

            string? orgId,

            string role)
        {
            Id = id;
            IsDisabled = isDisabled;
            Name = name;
            OrgId = orgId;
            Role = role;
        }
    }
}
