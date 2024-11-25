# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SsoSettingsArgs', 'SsoSettings']

@pulumi.input_type
class SsoSettingsArgs:
    def __init__(__self__, *,
                 provider_name: pulumi.Input[str],
                 oauth2_settings: Optional[pulumi.Input['SsoSettingsOauth2SettingsArgs']] = None,
                 saml_settings: Optional[pulumi.Input['SsoSettingsSamlSettingsArgs']] = None):
        """
        The set of arguments for constructing a SsoSettings resource.
        :param pulumi.Input[str] provider_name: The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml.
        :param pulumi.Input['SsoSettingsOauth2SettingsArgs'] oauth2_settings: The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic*oauth providers.
        :param pulumi.Input['SsoSettingsSamlSettingsArgs'] saml_settings: The SAML settings set. Required for the saml provider.
        """
        pulumi.set(__self__, "provider_name", provider_name)
        if oauth2_settings is not None:
            pulumi.set(__self__, "oauth2_settings", oauth2_settings)
        if saml_settings is not None:
            pulumi.set(__self__, "saml_settings", saml_settings)

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> pulumi.Input[str]:
        """
        The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml.
        """
        return pulumi.get(self, "provider_name")

    @provider_name.setter
    def provider_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "provider_name", value)

    @property
    @pulumi.getter(name="oauth2Settings")
    def oauth2_settings(self) -> Optional[pulumi.Input['SsoSettingsOauth2SettingsArgs']]:
        """
        The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic*oauth providers.
        """
        return pulumi.get(self, "oauth2_settings")

    @oauth2_settings.setter
    def oauth2_settings(self, value: Optional[pulumi.Input['SsoSettingsOauth2SettingsArgs']]):
        pulumi.set(self, "oauth2_settings", value)

    @property
    @pulumi.getter(name="samlSettings")
    def saml_settings(self) -> Optional[pulumi.Input['SsoSettingsSamlSettingsArgs']]:
        """
        The SAML settings set. Required for the saml provider.
        """
        return pulumi.get(self, "saml_settings")

    @saml_settings.setter
    def saml_settings(self, value: Optional[pulumi.Input['SsoSettingsSamlSettingsArgs']]):
        pulumi.set(self, "saml_settings", value)


@pulumi.input_type
class _SsoSettingsState:
    def __init__(__self__, *,
                 oauth2_settings: Optional[pulumi.Input['SsoSettingsOauth2SettingsArgs']] = None,
                 provider_name: Optional[pulumi.Input[str]] = None,
                 saml_settings: Optional[pulumi.Input['SsoSettingsSamlSettingsArgs']] = None):
        """
        Input properties used for looking up and filtering SsoSettings resources.
        :param pulumi.Input['SsoSettingsOauth2SettingsArgs'] oauth2_settings: The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic*oauth providers.
        :param pulumi.Input[str] provider_name: The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml.
        :param pulumi.Input['SsoSettingsSamlSettingsArgs'] saml_settings: The SAML settings set. Required for the saml provider.
        """
        if oauth2_settings is not None:
            pulumi.set(__self__, "oauth2_settings", oauth2_settings)
        if provider_name is not None:
            pulumi.set(__self__, "provider_name", provider_name)
        if saml_settings is not None:
            pulumi.set(__self__, "saml_settings", saml_settings)

    @property
    @pulumi.getter(name="oauth2Settings")
    def oauth2_settings(self) -> Optional[pulumi.Input['SsoSettingsOauth2SettingsArgs']]:
        """
        The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic*oauth providers.
        """
        return pulumi.get(self, "oauth2_settings")

    @oauth2_settings.setter
    def oauth2_settings(self, value: Optional[pulumi.Input['SsoSettingsOauth2SettingsArgs']]):
        pulumi.set(self, "oauth2_settings", value)

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml.
        """
        return pulumi.get(self, "provider_name")

    @provider_name.setter
    def provider_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_name", value)

    @property
    @pulumi.getter(name="samlSettings")
    def saml_settings(self) -> Optional[pulumi.Input['SsoSettingsSamlSettingsArgs']]:
        """
        The SAML settings set. Required for the saml provider.
        """
        return pulumi.get(self, "saml_settings")

    @saml_settings.setter
    def saml_settings(self, value: Optional[pulumi.Input['SsoSettingsSamlSettingsArgs']]):
        pulumi.set(self, "saml_settings", value)


class SsoSettings(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 oauth2_settings: Optional[pulumi.Input[Union['SsoSettingsOauth2SettingsArgs', 'SsoSettingsOauth2SettingsArgsDict']]] = None,
                 provider_name: Optional[pulumi.Input[str]] = None,
                 saml_settings: Optional[pulumi.Input[Union['SsoSettingsSamlSettingsArgs', 'SsoSettingsSamlSettingsArgsDict']]] = None,
                 __props__=None):
        """
        Manages Grafana SSO Settings for OAuth2 and SAML. Support for SAML is currently in preview, it will be available in Grafana Enterprise starting with v11.1.

        * [Official documentation](https://grafana.com/docs/grafana/latest/setup-grafana/configure-security/configure-authentication/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/sso-settings/)

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        # Configure SSO for GitHub using OAuth2
        github_sso_settings = grafana.oss.SsoSettings("github_sso_settings",
            provider_name="github",
            oauth2_settings={
                "name": "Github",
                "client_id": "<your GitHub app client id>",
                "client_secret": "<your GitHub app client secret>",
                "allow_sign_up": True,
                "auto_login": False,
                "scopes": "user:email,read:org",
                "team_ids": "150,300",
                "allowed_organizations": "[\\"My Organization\\", \\"Octocats\\"]",
                "allowed_domains": "mycompany.com mycompany.org",
            })
        # Configure SSO using generic OAuth2
        generic_sso_settings = grafana.oss.SsoSettings("generic_sso_settings",
            provider_name="generic_oauth",
            oauth2_settings={
                "name": "Auth0",
                "auth_url": "https://<domain>/authorize",
                "token_url": "https://<domain>/oauth/token",
                "api_url": "https://<domain>/userinfo",
                "client_id": "<client id>",
                "client_secret": "<client secret>",
                "allow_sign_up": True,
                "auto_login": False,
                "scopes": "openid profile email offline_access",
                "use_pkce": True,
                "use_refresh_token": True,
            })
        # Configure SSO using SAML
        saml_sso_settings = grafana.oss.SsoSettings("saml_sso_settings",
            provider_name="saml",
            saml_settings={
                "allow_sign_up": True,
                "certificate_path": "devenv/docker/blocks/auth/saml-enterprise/cert.crt",
                "private_key_path": "devenv/docker/blocks/auth/saml-enterprise/key.pem",
                "idp_metadata_url": "https://nexus.microsoftonline-p.com/federationmetadata/saml20/federationmetadata.xml",
                "signature_algorithm": "rsa-sha256",
                "assertion_attribute_login": "login",
                "assertion_attribute_email": "email",
                "name_id_format": "urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress",
            })
        ```

        ## Import

        ```sh
        $ pulumi import grafana:oss/ssoSettings:SsoSettings name "{{ provider }}"
        ```

        ```sh
        $ pulumi import grafana:oss/ssoSettings:SsoSettings name "{{ orgID }}:{{ provider }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['SsoSettingsOauth2SettingsArgs', 'SsoSettingsOauth2SettingsArgsDict']] oauth2_settings: The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic*oauth providers.
        :param pulumi.Input[str] provider_name: The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml.
        :param pulumi.Input[Union['SsoSettingsSamlSettingsArgs', 'SsoSettingsSamlSettingsArgsDict']] saml_settings: The SAML settings set. Required for the saml provider.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SsoSettingsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Grafana SSO Settings for OAuth2 and SAML. Support for SAML is currently in preview, it will be available in Grafana Enterprise starting with v11.1.

        * [Official documentation](https://grafana.com/docs/grafana/latest/setup-grafana/configure-security/configure-authentication/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/sso-settings/)

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        # Configure SSO for GitHub using OAuth2
        github_sso_settings = grafana.oss.SsoSettings("github_sso_settings",
            provider_name="github",
            oauth2_settings={
                "name": "Github",
                "client_id": "<your GitHub app client id>",
                "client_secret": "<your GitHub app client secret>",
                "allow_sign_up": True,
                "auto_login": False,
                "scopes": "user:email,read:org",
                "team_ids": "150,300",
                "allowed_organizations": "[\\"My Organization\\", \\"Octocats\\"]",
                "allowed_domains": "mycompany.com mycompany.org",
            })
        # Configure SSO using generic OAuth2
        generic_sso_settings = grafana.oss.SsoSettings("generic_sso_settings",
            provider_name="generic_oauth",
            oauth2_settings={
                "name": "Auth0",
                "auth_url": "https://<domain>/authorize",
                "token_url": "https://<domain>/oauth/token",
                "api_url": "https://<domain>/userinfo",
                "client_id": "<client id>",
                "client_secret": "<client secret>",
                "allow_sign_up": True,
                "auto_login": False,
                "scopes": "openid profile email offline_access",
                "use_pkce": True,
                "use_refresh_token": True,
            })
        # Configure SSO using SAML
        saml_sso_settings = grafana.oss.SsoSettings("saml_sso_settings",
            provider_name="saml",
            saml_settings={
                "allow_sign_up": True,
                "certificate_path": "devenv/docker/blocks/auth/saml-enterprise/cert.crt",
                "private_key_path": "devenv/docker/blocks/auth/saml-enterprise/key.pem",
                "idp_metadata_url": "https://nexus.microsoftonline-p.com/federationmetadata/saml20/federationmetadata.xml",
                "signature_algorithm": "rsa-sha256",
                "assertion_attribute_login": "login",
                "assertion_attribute_email": "email",
                "name_id_format": "urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress",
            })
        ```

        ## Import

        ```sh
        $ pulumi import grafana:oss/ssoSettings:SsoSettings name "{{ provider }}"
        ```

        ```sh
        $ pulumi import grafana:oss/ssoSettings:SsoSettings name "{{ orgID }}:{{ provider }}"
        ```

        :param str resource_name: The name of the resource.
        :param SsoSettingsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SsoSettingsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 oauth2_settings: Optional[pulumi.Input[Union['SsoSettingsOauth2SettingsArgs', 'SsoSettingsOauth2SettingsArgsDict']]] = None,
                 provider_name: Optional[pulumi.Input[str]] = None,
                 saml_settings: Optional[pulumi.Input[Union['SsoSettingsSamlSettingsArgs', 'SsoSettingsSamlSettingsArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SsoSettingsArgs.__new__(SsoSettingsArgs)

            __props__.__dict__["oauth2_settings"] = oauth2_settings
            if provider_name is None and not opts.urn:
                raise TypeError("Missing required property 'provider_name'")
            __props__.__dict__["provider_name"] = provider_name
            __props__.__dict__["saml_settings"] = saml_settings
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:index/ssoSettings:SsoSettings")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(SsoSettings, __self__).__init__(
            'grafana:oss/ssoSettings:SsoSettings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            oauth2_settings: Optional[pulumi.Input[Union['SsoSettingsOauth2SettingsArgs', 'SsoSettingsOauth2SettingsArgsDict']]] = None,
            provider_name: Optional[pulumi.Input[str]] = None,
            saml_settings: Optional[pulumi.Input[Union['SsoSettingsSamlSettingsArgs', 'SsoSettingsSamlSettingsArgsDict']]] = None) -> 'SsoSettings':
        """
        Get an existing SsoSettings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['SsoSettingsOauth2SettingsArgs', 'SsoSettingsOauth2SettingsArgsDict']] oauth2_settings: The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic*oauth providers.
        :param pulumi.Input[str] provider_name: The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml.
        :param pulumi.Input[Union['SsoSettingsSamlSettingsArgs', 'SsoSettingsSamlSettingsArgsDict']] saml_settings: The SAML settings set. Required for the saml provider.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SsoSettingsState.__new__(_SsoSettingsState)

        __props__.__dict__["oauth2_settings"] = oauth2_settings
        __props__.__dict__["provider_name"] = provider_name
        __props__.__dict__["saml_settings"] = saml_settings
        return SsoSettings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="oauth2Settings")
    def oauth2_settings(self) -> pulumi.Output[Optional['outputs.SsoSettingsOauth2Settings']]:
        """
        The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic*oauth providers.
        """
        return pulumi.get(self, "oauth2_settings")

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> pulumi.Output[str]:
        """
        The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml.
        """
        return pulumi.get(self, "provider_name")

    @property
    @pulumi.getter(name="samlSettings")
    def saml_settings(self) -> pulumi.Output[Optional['outputs.SsoSettingsSamlSettings']]:
        """
        The SAML settings set. Required for the saml provider.
        """
        return pulumi.get(self, "saml_settings")
