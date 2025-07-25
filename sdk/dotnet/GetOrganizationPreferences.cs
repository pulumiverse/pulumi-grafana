// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana
{
    [Obsolete(@"grafana.index/getorganizationpreferences.getOrganizationPreferences has been deprecated in favor of grafana.oss/getorganizationpreferences.getOrganizationPreferences")]
    public static class GetOrganizationPreferences
    {
        /// <summary>
        /// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/organization-management/)
        /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/preferences/#get-current-org-prefs)
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
        ///     var test = Grafana.Oss.GetOrganizationPreferences.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetOrganizationPreferencesResult> InvokeAsync(GetOrganizationPreferencesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationPreferencesResult>("grafana:index/getOrganizationPreferences:getOrganizationPreferences", args ?? new GetOrganizationPreferencesArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/organization-management/)
        /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/preferences/#get-current-org-prefs)
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
        ///     var test = Grafana.Oss.GetOrganizationPreferences.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetOrganizationPreferencesResult> Invoke(GetOrganizationPreferencesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationPreferencesResult>("grafana:index/getOrganizationPreferences:getOrganizationPreferences", args ?? new GetOrganizationPreferencesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/organization-management/)
        /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/preferences/#get-current-org-prefs)
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
        ///     var test = Grafana.Oss.GetOrganizationPreferences.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetOrganizationPreferencesResult> Invoke(GetOrganizationPreferencesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationPreferencesResult>("grafana:index/getOrganizationPreferences:getOrganizationPreferences", args ?? new GetOrganizationPreferencesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrganizationPreferencesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public string? OrgId { get; set; }

        public GetOrganizationPreferencesArgs()
        {
        }
        public static new GetOrganizationPreferencesArgs Empty => new GetOrganizationPreferencesArgs();
    }

    public sealed class GetOrganizationPreferencesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        public GetOrganizationPreferencesInvokeArgs()
        {
        }
        public static new GetOrganizationPreferencesInvokeArgs Empty => new GetOrganizationPreferencesInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrganizationPreferencesResult
    {
        /// <summary>
        /// The Organization home dashboard UID. This is only available in Grafana 9.0+.
        /// </summary>
        public readonly string HomeDashboardUid;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        public readonly string? OrgId;
        /// <summary>
        /// The Organization theme. Available values are `light`, `dark`, `system`, or an empty string for the default.
        /// </summary>
        public readonly string Theme;
        /// <summary>
        /// The Organization timezone. Available values are `utc`, `browser`, or an empty string for the default.
        /// </summary>
        public readonly string Timezone;
        /// <summary>
        /// The Organization week start day. Available values are `sunday`, `monday`, `saturday`, or an empty string for the default.
        /// </summary>
        public readonly string WeekStart;

        [OutputConstructor]
        private GetOrganizationPreferencesResult(
            string homeDashboardUid,

            string id,

            string? orgId,

            string theme,

            string timezone,

            string weekStart)
        {
            HomeDashboardUid = homeDashboardUid;
            Id = id;
            OrgId = orgId;
            Theme = theme;
            Timezone = timezone;
            WeekStart = weekStart;
        }
    }
}
