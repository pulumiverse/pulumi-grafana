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

    public sealed class SsoSettingsLdapSettingsConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("servers", required: true)]
        private InputList<Inputs.SsoSettingsLdapSettingsConfigServerArgs>? _servers;

        /// <summary>
        /// The LDAP servers configuration.
        /// </summary>
        public InputList<Inputs.SsoSettingsLdapSettingsConfigServerArgs> Servers
        {
            get => _servers ?? (_servers = new InputList<Inputs.SsoSettingsLdapSettingsConfigServerArgs>());
            set => _servers = value;
        }

        public SsoSettingsLdapSettingsConfigArgs()
        {
        }
        public static new SsoSettingsLdapSettingsConfigArgs Empty => new SsoSettingsLdapSettingsConfigArgs();
    }
}
