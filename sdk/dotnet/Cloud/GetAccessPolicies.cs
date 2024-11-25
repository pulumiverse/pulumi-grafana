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
    public static class GetAccessPolicies
    {
        /// <summary>
        /// Fetches access policies from Grafana Cloud.
        /// 
        /// * [Official documentation](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/)
        /// * [API documentation](https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-access-policies)
        /// 
        /// Required access policy scopes:
        /// 
        /// * accesspolicies:read
        /// </summary>
        public static Task<GetAccessPoliciesResult> InvokeAsync(GetAccessPoliciesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccessPoliciesResult>("grafana:cloud/getAccessPolicies:getAccessPolicies", args ?? new GetAccessPoliciesArgs(), options.WithDefaults());

        /// <summary>
        /// Fetches access policies from Grafana Cloud.
        /// 
        /// * [Official documentation](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/)
        /// * [API documentation](https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-access-policies)
        /// 
        /// Required access policy scopes:
        /// 
        /// * accesspolicies:read
        /// </summary>
        public static Output<GetAccessPoliciesResult> Invoke(GetAccessPoliciesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccessPoliciesResult>("grafana:cloud/getAccessPolicies:getAccessPolicies", args ?? new GetAccessPoliciesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccessPoliciesArgs : global::Pulumi.InvokeArgs
    {
        [Input("nameFilter")]
        public string? NameFilter { get; set; }

        [Input("regionFilter")]
        public string? RegionFilter { get; set; }

        public GetAccessPoliciesArgs()
        {
        }
        public static new GetAccessPoliciesArgs Empty => new GetAccessPoliciesArgs();
    }

    public sealed class GetAccessPoliciesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("nameFilter")]
        public Input<string>? NameFilter { get; set; }

        [Input("regionFilter")]
        public Input<string>? RegionFilter { get; set; }

        public GetAccessPoliciesInvokeArgs()
        {
        }
        public static new GetAccessPoliciesInvokeArgs Empty => new GetAccessPoliciesInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccessPoliciesResult
    {
        public readonly ImmutableArray<Outputs.GetAccessPoliciesAccessPolicyResult> AccessPolicies;
        /// <summary>
        /// The ID of this datasource. This is an internal identifier used by the provider to track this datasource.
        /// </summary>
        public readonly string Id;
        public readonly string? NameFilter;
        public readonly string? RegionFilter;

        [OutputConstructor]
        private GetAccessPoliciesResult(
            ImmutableArray<Outputs.GetAccessPoliciesAccessPolicyResult> accessPolicies,

            string id,

            string? nameFilter,

            string? regionFilter)
        {
            AccessPolicies = accessPolicies;
            Id = id;
            NameFilter = nameFilter;
            RegionFilter = regionFilter;
        }
    }
}
