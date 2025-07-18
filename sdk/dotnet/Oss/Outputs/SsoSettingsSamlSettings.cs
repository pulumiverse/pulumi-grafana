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
    public sealed class SsoSettingsSamlSettings
    {
        /// <summary>
        /// Whether SAML IdP-initiated login is allowed.
        /// </summary>
        public readonly bool? AllowIdpInitiated;
        /// <summary>
        /// Whether to allow new Grafana user creation through SAML login. If set to false, then only existing Grafana users can log in with SAML.
        /// </summary>
        public readonly bool? AllowSignUp;
        /// <summary>
        /// List of comma- or space-separated organizations. User should be a member of at least one organization to log in.
        /// </summary>
        public readonly string? AllowedOrganizations;
        /// <summary>
        /// Friendly name or name of the attribute within the SAML assertion to use as the user email.
        /// </summary>
        public readonly string? AssertionAttributeEmail;
        /// <summary>
        /// Friendly name or name of the attribute within the SAML assertion to use as the user groups.
        /// </summary>
        public readonly string? AssertionAttributeGroups;
        /// <summary>
        /// Friendly name or name of the attribute within the SAML assertion to use as the user login handle.
        /// </summary>
        public readonly string? AssertionAttributeLogin;
        /// <summary>
        /// Friendly name or name of the attribute within the SAML assertion to use as the user name. Alternatively, this can be a template with variables that match the names of attributes within the SAML assertion.
        /// </summary>
        public readonly string? AssertionAttributeName;
        /// <summary>
        /// Friendly name or name of the attribute within the SAML assertion to use as the user organization.
        /// </summary>
        public readonly string? AssertionAttributeOrg;
        /// <summary>
        /// Friendly name or name of the attribute within the SAML assertion to use as the user roles.
        /// </summary>
        public readonly string? AssertionAttributeRole;
        /// <summary>
        /// Whether SAML auto login is enabled.
        /// </summary>
        public readonly bool? AutoLogin;
        /// <summary>
        /// Base64-encoded string for the SP X.509 certificate.
        /// </summary>
        public readonly string? Certificate;
        /// <summary>
        /// Path for the SP X.509 certificate.
        /// </summary>
        public readonly string? CertificatePath;
        /// <summary>
        /// The client Id of your OAuth2 app.
        /// </summary>
        public readonly string? ClientId;
        /// <summary>
        /// The client secret of your OAuth2 app.
        /// </summary>
        public readonly string? ClientSecret;
        /// <summary>
        /// Define whether this configuration is enabled for SAML. Defaults to `true`.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The entity ID is a globally unique identifier for the service provider. It is used to identify the service provider to the identity provider. Defaults to the URL of the Grafana instance if not set.
        /// </summary>
        public readonly string? EntityId;
        /// <summary>
        /// If enabled, Grafana will fetch groups from Microsoft Graph API instead of using the groups claim from the ID token.
        /// </summary>
        public readonly bool? ForceUseGraphApi;
        /// <summary>
        /// Base64-encoded string for the IdP SAML metadata XML.
        /// </summary>
        public readonly string? IdpMetadata;
        /// <summary>
        /// Path for the IdP SAML metadata XML.
        /// </summary>
        public readonly string? IdpMetadataPath;
        /// <summary>
        /// URL for the IdP SAML metadata XML.
        /// </summary>
        public readonly string? IdpMetadataUrl;
        /// <summary>
        /// Duration, since the IdP issued a response and the SP is allowed to process it. For example: 90s, 1h.
        /// </summary>
        public readonly string? MaxIssueDelay;
        /// <summary>
        /// Duration, for how long the SP metadata is valid. For example: 48h, 5d.
        /// </summary>
        public readonly string? MetadataValidDuration;
        /// <summary>
        /// Name used to refer to the SAML authentication.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The Name ID Format to request within the SAML assertion. Defaults to urn:oasis:names:tc:SAML:2.0:nameid-format:transient
        /// </summary>
        public readonly string? NameIdFormat;
        /// <summary>
        /// List of comma- or space-separated Organization:OrgId:Role mappings. Organization can be * meaning “All users”. Role is optional and can have the following values: Viewer, Editor or Admin.
        /// </summary>
        public readonly string? OrgMapping;
        /// <summary>
        /// Base64-encoded string for the SP private key.
        /// </summary>
        public readonly string? PrivateKey;
        /// <summary>
        /// Path for the SP private key.
        /// </summary>
        public readonly string? PrivateKeyPath;
        /// <summary>
        /// Relay state for IdP-initiated login. Should match relay state configured in IdP.
        /// </summary>
        public readonly string? RelayState;
        /// <summary>
        /// List of comma- or space-separated roles which will be mapped into the Admin role.
        /// </summary>
        public readonly string? RoleValuesAdmin;
        /// <summary>
        /// List of comma- or space-separated roles which will be mapped into the Editor role.
        /// </summary>
        public readonly string? RoleValuesEditor;
        /// <summary>
        /// List of comma- or space-separated roles which will be mapped into the Grafana Admin (Super Admin) role.
        /// </summary>
        public readonly string? RoleValuesGrafanaAdmin;
        /// <summary>
        /// List of comma- or space-separated roles which will be mapped into the None role.
        /// </summary>
        public readonly string? RoleValuesNone;
        /// <summary>
        /// List of comma- or space-separated roles which will be mapped into the Viewer role.
        /// </summary>
        public readonly string? RoleValuesViewer;
        /// <summary>
        /// Signature algorithm used for signing requests to the IdP. Supported values are rsa-sha1, rsa-sha256, rsa-sha512.
        /// </summary>
        public readonly string? SignatureAlgorithm;
        /// <summary>
        /// Whether SAML Single Logout is enabled.
        /// </summary>
        public readonly bool? SingleLogout;
        /// <summary>
        /// Prevent synchronizing users’ organization roles from your IdP.
        /// </summary>
        public readonly bool? SkipOrgRoleSync;
        /// <summary>
        /// The token endpoint of your OAuth2 provider. Required for Azure AD providers.
        /// </summary>
        public readonly string? TokenUrl;

        [OutputConstructor]
        private SsoSettingsSamlSettings(
            bool? allowIdpInitiated,

            bool? allowSignUp,

            string? allowedOrganizations,

            string? assertionAttributeEmail,

            string? assertionAttributeGroups,

            string? assertionAttributeLogin,

            string? assertionAttributeName,

            string? assertionAttributeOrg,

            string? assertionAttributeRole,

            bool? autoLogin,

            string? certificate,

            string? certificatePath,

            string? clientId,

            string? clientSecret,

            bool? enabled,

            string? entityId,

            bool? forceUseGraphApi,

            string? idpMetadata,

            string? idpMetadataPath,

            string? idpMetadataUrl,

            string? maxIssueDelay,

            string? metadataValidDuration,

            string? name,

            string? nameIdFormat,

            string? orgMapping,

            string? privateKey,

            string? privateKeyPath,

            string? relayState,

            string? roleValuesAdmin,

            string? roleValuesEditor,

            string? roleValuesGrafanaAdmin,

            string? roleValuesNone,

            string? roleValuesViewer,

            string? signatureAlgorithm,

            bool? singleLogout,

            bool? skipOrgRoleSync,

            string? tokenUrl)
        {
            AllowIdpInitiated = allowIdpInitiated;
            AllowSignUp = allowSignUp;
            AllowedOrganizations = allowedOrganizations;
            AssertionAttributeEmail = assertionAttributeEmail;
            AssertionAttributeGroups = assertionAttributeGroups;
            AssertionAttributeLogin = assertionAttributeLogin;
            AssertionAttributeName = assertionAttributeName;
            AssertionAttributeOrg = assertionAttributeOrg;
            AssertionAttributeRole = assertionAttributeRole;
            AutoLogin = autoLogin;
            Certificate = certificate;
            CertificatePath = certificatePath;
            ClientId = clientId;
            ClientSecret = clientSecret;
            Enabled = enabled;
            EntityId = entityId;
            ForceUseGraphApi = forceUseGraphApi;
            IdpMetadata = idpMetadata;
            IdpMetadataPath = idpMetadataPath;
            IdpMetadataUrl = idpMetadataUrl;
            MaxIssueDelay = maxIssueDelay;
            MetadataValidDuration = metadataValidDuration;
            Name = name;
            NameIdFormat = nameIdFormat;
            OrgMapping = orgMapping;
            PrivateKey = privateKey;
            PrivateKeyPath = privateKeyPath;
            RelayState = relayState;
            RoleValuesAdmin = roleValuesAdmin;
            RoleValuesEditor = roleValuesEditor;
            RoleValuesGrafanaAdmin = roleValuesGrafanaAdmin;
            RoleValuesNone = roleValuesNone;
            RoleValuesViewer = roleValuesViewer;
            SignatureAlgorithm = signatureAlgorithm;
            SingleLogout = singleLogout;
            SkipOrgRoleSync = skipOrgRoleSync;
            TokenUrl = tokenUrl;
        }
    }
}
