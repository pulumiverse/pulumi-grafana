// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Inputs
{

    public sealed class SsoSettingsLdapSettingsConfigServerGroupMappingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to true, it makes the user of group_dn Grafana server admin.
        /// </summary>
        [Input("grafanaAdmin")]
        public Input<bool>? GrafanaAdmin { get; set; }

        /// <summary>
        /// LDAP distinguished name (DN) of LDAP group. If you want to match all (or no LDAP groups) then you can use wildcard ("*").
        /// </summary>
        [Input("groupDn", required: true)]
        public Input<string> GroupDn { get; set; } = null!;

        /// <summary>
        /// The Grafana organization database id.
        /// </summary>
        [Input("orgId")]
        public Input<int>? OrgId { get; set; }

        /// <summary>
        /// Assign users of group_dn the organization role Admin, Editor, or Viewer.
        /// </summary>
        [Input("orgRole", required: true)]
        public Input<string> OrgRole { get; set; } = null!;

        public SsoSettingsLdapSettingsConfigServerGroupMappingArgs()
        {
        }
        public static new SsoSettingsLdapSettingsConfigServerGroupMappingArgs Empty => new SsoSettingsLdapSettingsConfigServerGroupMappingArgs();
    }
}
