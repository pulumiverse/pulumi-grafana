// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Inputs
{

    public sealed class SsoSettingsSamlSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether SAML IdP-initiated login is allowed.
        /// </summary>
        [Input("allowIdpInitiated")]
        public Input<bool>? AllowIdpInitiated { get; set; }

        /// <summary>
        /// Whether to allow new Grafana user creation through SAML login. If set to false, then only existing Grafana users can log in with SAML.
        /// </summary>
        [Input("allowSignUp")]
        public Input<bool>? AllowSignUp { get; set; }

        /// <summary>
        /// List of comma- or space-separated organizations. User should be a member of at least one organization to log in.
        /// </summary>
        [Input("allowedOrganizations")]
        public Input<string>? AllowedOrganizations { get; set; }

        /// <summary>
        /// Friendly name or name of the attribute within the SAML assertion to use as the user email.
        /// </summary>
        [Input("assertionAttributeEmail")]
        public Input<string>? AssertionAttributeEmail { get; set; }

        /// <summary>
        /// Friendly name or name of the attribute within the SAML assertion to use as the user groups.
        /// </summary>
        [Input("assertionAttributeGroups")]
        public Input<string>? AssertionAttributeGroups { get; set; }

        /// <summary>
        /// Friendly name or name of the attribute within the SAML assertion to use as the user login handle.
        /// </summary>
        [Input("assertionAttributeLogin")]
        public Input<string>? AssertionAttributeLogin { get; set; }

        /// <summary>
        /// Friendly name or name of the attribute within the SAML assertion to use as the user name. Alternatively, this can be a template with variables that match the names of attributes within the SAML assertion.
        /// </summary>
        [Input("assertionAttributeName")]
        public Input<string>? AssertionAttributeName { get; set; }

        /// <summary>
        /// Friendly name or name of the attribute within the SAML assertion to use as the user organization.
        /// </summary>
        [Input("assertionAttributeOrg")]
        public Input<string>? AssertionAttributeOrg { get; set; }

        /// <summary>
        /// Friendly name or name of the attribute within the SAML assertion to use as the user roles.
        /// </summary>
        [Input("assertionAttributeRole")]
        public Input<string>? AssertionAttributeRole { get; set; }

        /// <summary>
        /// Whether SAML auto login is enabled.
        /// </summary>
        [Input("autoLogin")]
        public Input<bool>? AutoLogin { get; set; }

        [Input("certificate")]
        private Input<string>? _certificate;

        /// <summary>
        /// Base64-encoded string for the SP X.509 certificate.
        /// </summary>
        public Input<string>? Certificate
        {
            get => _certificate;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _certificate = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Path for the SP X.509 certificate.
        /// </summary>
        [Input("certificatePath")]
        public Input<string>? CertificatePath { get; set; }

        /// <summary>
        /// The client Id of your OAuth2 app.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The client secret of your OAuth2 app.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// Define whether this configuration is enabled for SAML. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The entity ID is a globally unique identifier for the service provider. It is used to identify the service provider to the identity provider. Defaults to the URL of the Grafana instance if not set.
        /// </summary>
        [Input("entityId")]
        public Input<string>? EntityId { get; set; }

        /// <summary>
        /// If enabled, Grafana will fetch groups from Microsoft Graph API instead of using the groups claim from the ID token.
        /// </summary>
        [Input("forceUseGraphApi")]
        public Input<bool>? ForceUseGraphApi { get; set; }

        /// <summary>
        /// Base64-encoded string for the IdP SAML metadata XML.
        /// </summary>
        [Input("idpMetadata")]
        public Input<string>? IdpMetadata { get; set; }

        /// <summary>
        /// Path for the IdP SAML metadata XML.
        /// </summary>
        [Input("idpMetadataPath")]
        public Input<string>? IdpMetadataPath { get; set; }

        /// <summary>
        /// URL for the IdP SAML metadata XML.
        /// </summary>
        [Input("idpMetadataUrl")]
        public Input<string>? IdpMetadataUrl { get; set; }

        /// <summary>
        /// Duration, since the IdP issued a response and the SP is allowed to process it. For example: 90s, 1h.
        /// </summary>
        [Input("maxIssueDelay")]
        public Input<string>? MaxIssueDelay { get; set; }

        /// <summary>
        /// Duration, for how long the SP metadata is valid. For example: 48h, 5d.
        /// </summary>
        [Input("metadataValidDuration")]
        public Input<string>? MetadataValidDuration { get; set; }

        /// <summary>
        /// Name used to refer to the SAML authentication.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Name ID Format to request within the SAML assertion. Defaults to urn:oasis:names:tc:SAML:2.0:nameid-format:transient
        /// </summary>
        [Input("nameIdFormat")]
        public Input<string>? NameIdFormat { get; set; }

        /// <summary>
        /// List of comma- or space-separated Organization:OrgId:Role mappings. Organization can be * meaning “All users”. Role is optional and can have the following values: Viewer, Editor or Admin.
        /// </summary>
        [Input("orgMapping")]
        public Input<string>? OrgMapping { get; set; }

        [Input("privateKey")]
        private Input<string>? _privateKey;

        /// <summary>
        /// Base64-encoded string for the SP private key.
        /// </summary>
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Path for the SP private key.
        /// </summary>
        [Input("privateKeyPath")]
        public Input<string>? PrivateKeyPath { get; set; }

        /// <summary>
        /// Relay state for IdP-initiated login. Should match relay state configured in IdP.
        /// </summary>
        [Input("relayState")]
        public Input<string>? RelayState { get; set; }

        /// <summary>
        /// List of comma- or space-separated roles which will be mapped into the Admin role.
        /// </summary>
        [Input("roleValuesAdmin")]
        public Input<string>? RoleValuesAdmin { get; set; }

        /// <summary>
        /// List of comma- or space-separated roles which will be mapped into the Editor role.
        /// </summary>
        [Input("roleValuesEditor")]
        public Input<string>? RoleValuesEditor { get; set; }

        /// <summary>
        /// List of comma- or space-separated roles which will be mapped into the Grafana Admin (Super Admin) role.
        /// </summary>
        [Input("roleValuesGrafanaAdmin")]
        public Input<string>? RoleValuesGrafanaAdmin { get; set; }

        /// <summary>
        /// List of comma- or space-separated roles which will be mapped into the None role.
        /// </summary>
        [Input("roleValuesNone")]
        public Input<string>? RoleValuesNone { get; set; }

        /// <summary>
        /// List of comma- or space-separated roles which will be mapped into the Viewer role.
        /// </summary>
        [Input("roleValuesViewer")]
        public Input<string>? RoleValuesViewer { get; set; }

        /// <summary>
        /// Signature algorithm used for signing requests to the IdP. Supported values are rsa-sha1, rsa-sha256, rsa-sha512.
        /// </summary>
        [Input("signatureAlgorithm")]
        public Input<string>? SignatureAlgorithm { get; set; }

        /// <summary>
        /// Whether SAML Single Logout is enabled.
        /// </summary>
        [Input("singleLogout")]
        public Input<bool>? SingleLogout { get; set; }

        /// <summary>
        /// Prevent synchronizing users’ organization roles from your IdP.
        /// </summary>
        [Input("skipOrgRoleSync")]
        public Input<bool>? SkipOrgRoleSync { get; set; }

        /// <summary>
        /// The token endpoint of your OAuth2 provider. Required for Azure AD providers.
        /// </summary>
        [Input("tokenUrl")]
        public Input<string>? TokenUrl { get; set; }

        public SsoSettingsSamlSettingsGetArgs()
        {
        }
        public static new SsoSettingsSamlSettingsGetArgs Empty => new SsoSettingsSamlSettingsGetArgs();
    }
}
