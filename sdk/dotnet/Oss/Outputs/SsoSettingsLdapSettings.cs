// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
    public sealed class SsoSettingsLdapSettings
    {
        /// <summary>
        /// Whether to allow new Grafana user creation through LDAP login. If set to false, then only existing Grafana users can log in with LDAP.
        /// </summary>
        public readonly bool? AllowSignUp;
        /// <summary>
        /// The LDAP configuration.
        /// </summary>
        public readonly Outputs.SsoSettingsLdapSettingsConfig Config;
        /// <summary>
        /// Define whether this configuration is enabled for LDAP. Defaults to `true`.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Prevent synchronizing users’ organization roles from LDAP.
        /// </summary>
        public readonly bool? SkipOrgRoleSync;

        [OutputConstructor]
        private SsoSettingsLdapSettings(
            bool? allowSignUp,

            Outputs.SsoSettingsLdapSettingsConfig config,

            bool? enabled,

            bool? skipOrgRoleSync)
        {
            AllowSignUp = allowSignUp;
            Config = config;
            Enabled = enabled;
            SkipOrgRoleSync = skipOrgRoleSync;
        }
    }
}
