// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages Grafana SSO Settings for OAuth2, SAML and LDAP. Support for LDAP is currently in preview, it will be available in Grafana starting with v11.3.
 *
 * * [Official documentation](https://grafana.com/docs/grafana/latest/setup-grafana/configure-security/configure-authentication/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/sso-settings/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * // Configure SSO for GitHub using OAuth2
 * const githubSsoSettings = new grafana.oss.SsoSettings("github_sso_settings", {
 *     providerName: "github",
 *     oauth2Settings: {
 *         name: "Github",
 *         clientId: "<your GitHub app client id>",
 *         clientSecret: "<your GitHub app client secret>",
 *         allowSignUp: true,
 *         autoLogin: false,
 *         scopes: "user:email,read:org",
 *         teamIds: "150,300",
 *         allowedOrganizations: "[\"My Organization\", \"Octocats\"]",
 *         allowedDomains: "mycompany.com mycompany.org",
 *     },
 * });
 * // Configure SSO using generic OAuth2
 * const genericSsoSettings = new grafana.oss.SsoSettings("generic_sso_settings", {
 *     providerName: "generic_oauth",
 *     oauth2Settings: {
 *         name: "Auth0",
 *         authUrl: "https://<domain>/authorize",
 *         tokenUrl: "https://<domain>/oauth/token",
 *         apiUrl: "https://<domain>/userinfo",
 *         clientId: "<client id>",
 *         clientSecret: "<client secret>",
 *         allowSignUp: true,
 *         autoLogin: false,
 *         scopes: "openid profile email offline_access",
 *         usePkce: true,
 *         useRefreshToken: true,
 *     },
 * });
 * // Configure SSO using SAML
 * const samlSsoSettings = new grafana.oss.SsoSettings("saml_sso_settings", {
 *     providerName: "saml",
 *     samlSettings: {
 *         allowSignUp: true,
 *         certificatePath: "devenv/docker/blocks/auth/saml-enterprise/cert.crt",
 *         privateKeyPath: "devenv/docker/blocks/auth/saml-enterprise/key.pem",
 *         idpMetadataUrl: "https://nexus.microsoftonline-p.com/federationmetadata/saml20/federationmetadata.xml",
 *         signatureAlgorithm: "rsa-sha256",
 *         assertionAttributeLogin: "login",
 *         assertionAttributeEmail: "email",
 *         nameIdFormat: "urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress",
 *     },
 * });
 * // Configure SSO using LDAP
 * const ldapSsoSettings = new grafana.oss.SsoSettings("ldap_sso_settings", {
 *     providerName: "ldap",
 *     ldapSettings: {
 *         enabled: true,
 *         config: {
 *             servers: [{
 *                 host: "127.0.0.1",
 *                 port: 389,
 *                 searchFilter: "(cn=%s)",
 *                 bindDn: "cn=admin,dc=grafana,dc=org",
 *                 bindPassword: "grafana",
 *                 searchBaseDns: ["dc=grafana,dc=org"],
 *                 attributes: {
 *                     name: "givenName",
 *                     surname: "sn",
 *                     username: "cn",
 *                     member_of: "memberOf",
 *                     email: "email",
 *                 },
 *                 groupMappings: [
 *                     {
 *                         groupDn: "cn=superadmins,dc=grafana,dc=org",
 *                         orgRole: "Admin",
 *                         orgId: 1,
 *                         grafanaAdmin: true,
 *                     },
 *                     {
 *                         groupDn: "cn=users,dc=grafana,dc=org",
 *                         orgRole: "Editor",
 *                     },
 *                     {
 *                         groupDn: "*",
 *                         orgRole: "Viewer",
 *                     },
 *                 ],
 *             }],
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:oss/ssoSettings:SsoSettings name "{{ provider }}"
 * ```
 *
 * ```sh
 * $ pulumi import grafana:oss/ssoSettings:SsoSettings name "{{ orgID }}:{{ provider }}"
 * ```
 */
export class SsoSettings extends pulumi.CustomResource {
    /**
     * Get an existing SsoSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SsoSettingsState, opts?: pulumi.CustomResourceOptions): SsoSettings {
        return new SsoSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:oss/ssoSettings:SsoSettings';

    /**
     * Returns true if the given object is an instance of SsoSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SsoSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SsoSettings.__pulumiType;
    }

    /**
     * The LDAP settings set. Required for the ldap provider.
     */
    public readonly ldapSettings!: pulumi.Output<outputs.oss.SsoSettingsLdapSettings | undefined>;
    /**
     * The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic*oauth providers.
     */
    public readonly oauth2Settings!: pulumi.Output<outputs.oss.SsoSettingsOauth2Settings | undefined>;
    /**
     * The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml, ldap.
     */
    public readonly providerName!: pulumi.Output<string>;
    /**
     * The SAML settings set. Required for the saml provider.
     */
    public readonly samlSettings!: pulumi.Output<outputs.oss.SsoSettingsSamlSettings | undefined>;

    /**
     * Create a SsoSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SsoSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SsoSettingsArgs | SsoSettingsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SsoSettingsState | undefined;
            resourceInputs["ldapSettings"] = state ? state.ldapSettings : undefined;
            resourceInputs["oauth2Settings"] = state ? state.oauth2Settings : undefined;
            resourceInputs["providerName"] = state ? state.providerName : undefined;
            resourceInputs["samlSettings"] = state ? state.samlSettings : undefined;
        } else {
            const args = argsOrState as SsoSettingsArgs | undefined;
            if ((!args || args.providerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'providerName'");
            }
            resourceInputs["ldapSettings"] = args ? args.ldapSettings : undefined;
            resourceInputs["oauth2Settings"] = args ? args.oauth2Settings : undefined;
            resourceInputs["providerName"] = args ? args.providerName : undefined;
            resourceInputs["samlSettings"] = args ? args.samlSettings : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "grafana:index/ssoSettings:SsoSettings" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(SsoSettings.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SsoSettings resources.
 */
export interface SsoSettingsState {
    /**
     * The LDAP settings set. Required for the ldap provider.
     */
    ldapSettings?: pulumi.Input<inputs.oss.SsoSettingsLdapSettings>;
    /**
     * The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic*oauth providers.
     */
    oauth2Settings?: pulumi.Input<inputs.oss.SsoSettingsOauth2Settings>;
    /**
     * The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml, ldap.
     */
    providerName?: pulumi.Input<string>;
    /**
     * The SAML settings set. Required for the saml provider.
     */
    samlSettings?: pulumi.Input<inputs.oss.SsoSettingsSamlSettings>;
}

/**
 * The set of arguments for constructing a SsoSettings resource.
 */
export interface SsoSettingsArgs {
    /**
     * The LDAP settings set. Required for the ldap provider.
     */
    ldapSettings?: pulumi.Input<inputs.oss.SsoSettingsLdapSettings>;
    /**
     * The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic*oauth providers.
     */
    oauth2Settings?: pulumi.Input<inputs.oss.SsoSettingsOauth2Settings>;
    /**
     * The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml, ldap.
     */
    providerName: pulumi.Input<string>;
    /**
     * The SAML settings set. Required for the saml provider.
     */
    samlSettings?: pulumi.Input<inputs.oss.SsoSettingsSamlSettings>;
}