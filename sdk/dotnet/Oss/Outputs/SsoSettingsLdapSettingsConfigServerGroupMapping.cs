// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Oss.Outputs
{

    [OutputType]
    public sealed class SsoSettingsLdapSettingsConfigServerGroupMapping
    {
        /// <summary>
        /// If set to true, it makes the user of group_dn Grafana server admin.
        /// </summary>
        public readonly bool? GrafanaAdmin;
        /// <summary>
        /// LDAP distinguished name (DN) of LDAP group. If you want to match all (or no LDAP groups) then you can use wildcard ("*").
        /// </summary>
        public readonly string GroupDn;
        /// <summary>
        /// The Grafana organization database id.
        /// </summary>
        public readonly int? OrgId;
        /// <summary>
        /// Assign users of group_dn the organization role Admin, Editor, or Viewer.
        /// </summary>
        public readonly string OrgRole;

        [OutputConstructor]
        private SsoSettingsLdapSettingsConfigServerGroupMapping(
            bool? grafanaAdmin,

            string groupDn,

            int? orgId,

            string orgRole)
        {
            GrafanaAdmin = grafanaAdmin;
            GroupDn = groupDn;
            OrgId = orgId;
            OrgRole = orgRole;
        }
    }
}
