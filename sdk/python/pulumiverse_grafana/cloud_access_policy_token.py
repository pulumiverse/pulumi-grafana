# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['CloudAccessPolicyTokenArgs', 'CloudAccessPolicyToken']

@pulumi.input_type
class CloudAccessPolicyTokenArgs:
    def __init__(__self__, *,
                 access_policy_id: pulumi.Input[str],
                 region: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CloudAccessPolicyToken resource.
        :param pulumi.Input[str] access_policy_id: ID of the access policy for which to create a token.
        :param pulumi.Input[str] region: Region of the access policy. Should be set to the same region as the access policy. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
        :param pulumi.Input[str] display_name: Display name of the access policy token. Defaults to the name.
        :param pulumi.Input[str] expires_at: Expiration date of the access policy token. Does not expire by default.
        :param pulumi.Input[str] name: Name of the access policy token.
        """
        pulumi.set(__self__, "access_policy_id", access_policy_id)
        pulumi.set(__self__, "region", region)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="accessPolicyId")
    def access_policy_id(self) -> pulumi.Input[str]:
        """
        ID of the access policy for which to create a token.
        """
        return pulumi.get(self, "access_policy_id")

    @access_policy_id.setter
    def access_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_policy_id", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        Region of the access policy. Should be set to the same region as the access policy. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name of the access policy token. Defaults to the name.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        Expiration date of the access policy token. Does not expire by default.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the access policy token.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _CloudAccessPolicyTokenState:
    def __init__(__self__, *,
                 access_policy_id: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CloudAccessPolicyToken resources.
        :param pulumi.Input[str] access_policy_id: ID of the access policy for which to create a token.
        :param pulumi.Input[str] created_at: Creation date of the access policy token.
        :param pulumi.Input[str] display_name: Display name of the access policy token. Defaults to the name.
        :param pulumi.Input[str] expires_at: Expiration date of the access policy token. Does not expire by default.
        :param pulumi.Input[str] name: Name of the access policy token.
        :param pulumi.Input[str] region: Region of the access policy. Should be set to the same region as the access policy. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
        :param pulumi.Input[str] updated_at: Last update date of the access policy token.
        """
        if access_policy_id is not None:
            pulumi.set(__self__, "access_policy_id", access_policy_id)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="accessPolicyId")
    def access_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the access policy for which to create a token.
        """
        return pulumi.get(self, "access_policy_id")

    @access_policy_id.setter
    def access_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_policy_id", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Creation date of the access policy token.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name of the access policy token. Defaults to the name.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        Expiration date of the access policy token. Does not expire by default.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the access policy token.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Region of the access policy. Should be set to the same region as the access policy. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        Last update date of the access policy token.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


warnings.warn("""grafana.index/cloudaccesspolicytoken.CloudAccessPolicyToken has been deprecated in favor of grafana.cloud/accesspolicytoken.AccessPolicyToken""", DeprecationWarning)


class CloudAccessPolicyToken(pulumi.CustomResource):
    warnings.warn("""grafana.index/cloudaccesspolicytoken.CloudAccessPolicyToken has been deprecated in favor of grafana.cloud/accesspolicytoken.AccessPolicyToken""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_policy_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        * [Official documentation](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/)
        * [API documentation](https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#create-a-token)

        Required access policy scopes:

        * accesspolicies:read
        * accesspolicies:write
        * accesspolicies:delete

        ## Example Usage

        ```python
        import pulumi
        import pulumi_grafana as grafana
        import pulumiverse_grafana as grafana

        current = grafana.cloud.get_organization(slug="<your org slug>")
        test = grafana.cloud.AccessPolicy("test",
            region="us",
            name="my-policy",
            display_name="My Policy",
            scopes=[
                "metrics:read",
                "logs:read",
            ],
            realms=[grafana.cloud.AccessPolicyRealmArgs(
                type="org",
                identifier=current.id,
                label_policies=[grafana.cloud.AccessPolicyRealmLabelPolicyArgs(
                    selector="{namespace=\\"default\\"}",
                )],
            )])
        test_access_policy_token = grafana.cloud.AccessPolicyToken("test",
            region="us",
            access_policy_id=test.policy_id,
            name="my-policy-token",
            display_name="My Policy Token",
            expires_at="2023-01-01T00:00:00Z")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/cloudAccessPolicyToken:CloudAccessPolicyToken name "{{ region }}:{{ tokenId }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_policy_id: ID of the access policy for which to create a token.
        :param pulumi.Input[str] display_name: Display name of the access policy token. Defaults to the name.
        :param pulumi.Input[str] expires_at: Expiration date of the access policy token. Does not expire by default.
        :param pulumi.Input[str] name: Name of the access policy token.
        :param pulumi.Input[str] region: Region of the access policy. Should be set to the same region as the access policy. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CloudAccessPolicyTokenArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [Official documentation](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/)
        * [API documentation](https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#create-a-token)

        Required access policy scopes:

        * accesspolicies:read
        * accesspolicies:write
        * accesspolicies:delete

        ## Example Usage

        ```python
        import pulumi
        import pulumi_grafana as grafana
        import pulumiverse_grafana as grafana

        current = grafana.cloud.get_organization(slug="<your org slug>")
        test = grafana.cloud.AccessPolicy("test",
            region="us",
            name="my-policy",
            display_name="My Policy",
            scopes=[
                "metrics:read",
                "logs:read",
            ],
            realms=[grafana.cloud.AccessPolicyRealmArgs(
                type="org",
                identifier=current.id,
                label_policies=[grafana.cloud.AccessPolicyRealmLabelPolicyArgs(
                    selector="{namespace=\\"default\\"}",
                )],
            )])
        test_access_policy_token = grafana.cloud.AccessPolicyToken("test",
            region="us",
            access_policy_id=test.policy_id,
            name="my-policy-token",
            display_name="My Policy Token",
            expires_at="2023-01-01T00:00:00Z")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/cloudAccessPolicyToken:CloudAccessPolicyToken name "{{ region }}:{{ tokenId }}"
        ```

        :param str resource_name: The name of the resource.
        :param CloudAccessPolicyTokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CloudAccessPolicyTokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_policy_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""CloudAccessPolicyToken is deprecated: grafana.index/cloudaccesspolicytoken.CloudAccessPolicyToken has been deprecated in favor of grafana.cloud/accesspolicytoken.AccessPolicyToken""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CloudAccessPolicyTokenArgs.__new__(CloudAccessPolicyTokenArgs)

            if access_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'access_policy_id'")
            __props__.__dict__["access_policy_id"] = access_policy_id
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["expires_at"] = expires_at
            __props__.__dict__["name"] = name
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["created_at"] = None
            __props__.__dict__["token"] = None
            __props__.__dict__["updated_at"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:index/cloudAccessPolicyToken:CloudAccessPolicyToken")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["token"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(CloudAccessPolicyToken, __self__).__init__(
            'grafana:index/cloudAccessPolicyToken:CloudAccessPolicyToken',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_policy_id: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            expires_at: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'CloudAccessPolicyToken':
        """
        Get an existing CloudAccessPolicyToken resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_policy_id: ID of the access policy for which to create a token.
        :param pulumi.Input[str] created_at: Creation date of the access policy token.
        :param pulumi.Input[str] display_name: Display name of the access policy token. Defaults to the name.
        :param pulumi.Input[str] expires_at: Expiration date of the access policy token. Does not expire by default.
        :param pulumi.Input[str] name: Name of the access policy token.
        :param pulumi.Input[str] region: Region of the access policy. Should be set to the same region as the access policy. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
        :param pulumi.Input[str] updated_at: Last update date of the access policy token.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CloudAccessPolicyTokenState.__new__(_CloudAccessPolicyTokenState)

        __props__.__dict__["access_policy_id"] = access_policy_id
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["expires_at"] = expires_at
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["token"] = token
        __props__.__dict__["updated_at"] = updated_at
        return CloudAccessPolicyToken(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessPolicyId")
    def access_policy_id(self) -> pulumi.Output[str]:
        """
        ID of the access policy for which to create a token.
        """
        return pulumi.get(self, "access_policy_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Creation date of the access policy token.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        Display name of the access policy token. Defaults to the name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> pulumi.Output[Optional[str]]:
        """
        Expiration date of the access policy token. Does not expire by default.
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the access policy token.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Region of the access policy. Should be set to the same region as the access policy. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        Last update date of the access policy token.
        """
        return pulumi.get(self, "updated_at")

