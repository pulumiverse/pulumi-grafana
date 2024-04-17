# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
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
        :param pulumi.Input['SsoSettingsOauth2SettingsArgs'] oauth2_settings: The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic_oauth providers.
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
        The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic_oauth providers.
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
        :param pulumi.Input['SsoSettingsOauth2SettingsArgs'] oauth2_settings: The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic_oauth providers.
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
        The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic_oauth providers.
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
                 oauth2_settings: Optional[pulumi.Input[pulumi.InputType['SsoSettingsOauth2SettingsArgs']]] = None,
                 provider_name: Optional[pulumi.Input[str]] = None,
                 saml_settings: Optional[pulumi.Input[pulumi.InputType['SsoSettingsSamlSettingsArgs']]] = None,
                 __props__=None):
        """
        Create a SsoSettings resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SsoSettingsOauth2SettingsArgs']] oauth2_settings: The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic_oauth providers.
        :param pulumi.Input[str] provider_name: The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml.
        :param pulumi.Input[pulumi.InputType['SsoSettingsSamlSettingsArgs']] saml_settings: The SAML settings set. Required for the saml provider.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SsoSettingsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SsoSettings resource with the given unique name, props, and options.
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
                 oauth2_settings: Optional[pulumi.Input[pulumi.InputType['SsoSettingsOauth2SettingsArgs']]] = None,
                 provider_name: Optional[pulumi.Input[str]] = None,
                 saml_settings: Optional[pulumi.Input[pulumi.InputType['SsoSettingsSamlSettingsArgs']]] = None,
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
        super(SsoSettings, __self__).__init__(
            'grafana:index/ssoSettings:SsoSettings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            oauth2_settings: Optional[pulumi.Input[pulumi.InputType['SsoSettingsOauth2SettingsArgs']]] = None,
            provider_name: Optional[pulumi.Input[str]] = None,
            saml_settings: Optional[pulumi.Input[pulumi.InputType['SsoSettingsSamlSettingsArgs']]] = None) -> 'SsoSettings':
        """
        Get an existing SsoSettings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SsoSettingsOauth2SettingsArgs']] oauth2_settings: The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic_oauth providers.
        :param pulumi.Input[str] provider_name: The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml.
        :param pulumi.Input[pulumi.InputType['SsoSettingsSamlSettingsArgs']] saml_settings: The SAML settings set. Required for the saml provider.
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
        The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic_oauth providers.
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

