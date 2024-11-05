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
from . import _utilities

__all__ = ['CloudStackServiceAccountTokenArgs', 'CloudStackServiceAccountToken']

@pulumi.input_type
class CloudStackServiceAccountTokenArgs:
    def __init__(__self__, *,
                 service_account_id: pulumi.Input[str],
                 stack_slug: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 seconds_to_live: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a CloudStackServiceAccountToken resource.
        """
        pulumi.set(__self__, "service_account_id", service_account_id)
        pulumi.set(__self__, "stack_slug", stack_slug)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if seconds_to_live is not None:
            pulumi.set(__self__, "seconds_to_live", seconds_to_live)

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "service_account_id")

    @service_account_id.setter
    def service_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_account_id", value)

    @property
    @pulumi.getter(name="stackSlug")
    def stack_slug(self) -> pulumi.Input[str]:
        return pulumi.get(self, "stack_slug")

    @stack_slug.setter
    def stack_slug(self, value: pulumi.Input[str]):
        pulumi.set(self, "stack_slug", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="secondsToLive")
    def seconds_to_live(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "seconds_to_live")

    @seconds_to_live.setter
    def seconds_to_live(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "seconds_to_live", value)


@pulumi.input_type
class _CloudStackServiceAccountTokenState:
    def __init__(__self__, *,
                 expiration: Optional[pulumi.Input[str]] = None,
                 has_expired: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 seconds_to_live: Optional[pulumi.Input[int]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None,
                 stack_slug: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CloudStackServiceAccountToken resources.
        """
        if expiration is not None:
            pulumi.set(__self__, "expiration", expiration)
        if has_expired is not None:
            pulumi.set(__self__, "has_expired", has_expired)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if seconds_to_live is not None:
            pulumi.set(__self__, "seconds_to_live", seconds_to_live)
        if service_account_id is not None:
            pulumi.set(__self__, "service_account_id", service_account_id)
        if stack_slug is not None:
            pulumi.set(__self__, "stack_slug", stack_slug)

    @property
    @pulumi.getter
    def expiration(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expiration")

    @expiration.setter
    def expiration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration", value)

    @property
    @pulumi.getter(name="hasExpired")
    def has_expired(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "has_expired")

    @has_expired.setter
    def has_expired(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "has_expired", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="secondsToLive")
    def seconds_to_live(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "seconds_to_live")

    @seconds_to_live.setter
    def seconds_to_live(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "seconds_to_live", value)

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "service_account_id")

    @service_account_id.setter
    def service_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_id", value)

    @property
    @pulumi.getter(name="stackSlug")
    def stack_slug(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "stack_slug")

    @stack_slug.setter
    def stack_slug(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_slug", value)


warnings.warn("""grafana.index/cloudstackserviceaccounttoken.CloudStackServiceAccountToken has been deprecated in favor of grafana.cloud/stackserviceaccounttoken.StackServiceAccountToken""", DeprecationWarning)


class CloudStackServiceAccountToken(pulumi.CustomResource):
    warnings.warn("""grafana.index/cloudstackserviceaccounttoken.CloudStackServiceAccountToken has been deprecated in favor of grafana.cloud/stackserviceaccounttoken.StackServiceAccountToken""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 seconds_to_live: Optional[pulumi.Input[int]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None,
                 stack_slug: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages service account tokens of a Grafana Cloud stack using the Cloud API
        This can be used to bootstrap a management service account token for a new stack

        * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api)

        Required access policy scopes:

        * stack-service-accounts:write

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        cloud_sa = grafana.cloud.StackServiceAccount("cloud_sa",
            stack_slug="<your stack slug>",
            name="cloud service account",
            role="Admin",
            is_disabled=False)
        foo = grafana.cloud.StackServiceAccountToken("foo",
            name="key_foo",
            service_account_id=cloud_sa.id)
        pulumi.export("serviceAccountTokenFooKey", foo.key)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CloudStackServiceAccountTokenArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages service account tokens of a Grafana Cloud stack using the Cloud API
        This can be used to bootstrap a management service account token for a new stack

        * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api)

        Required access policy scopes:

        * stack-service-accounts:write

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        cloud_sa = grafana.cloud.StackServiceAccount("cloud_sa",
            stack_slug="<your stack slug>",
            name="cloud service account",
            role="Admin",
            is_disabled=False)
        foo = grafana.cloud.StackServiceAccountToken("foo",
            name="key_foo",
            service_account_id=cloud_sa.id)
        pulumi.export("serviceAccountTokenFooKey", foo.key)
        ```

        :param str resource_name: The name of the resource.
        :param CloudStackServiceAccountTokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CloudStackServiceAccountTokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 seconds_to_live: Optional[pulumi.Input[int]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None,
                 stack_slug: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""CloudStackServiceAccountToken is deprecated: grafana.index/cloudstackserviceaccounttoken.CloudStackServiceAccountToken has been deprecated in favor of grafana.cloud/stackserviceaccounttoken.StackServiceAccountToken""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CloudStackServiceAccountTokenArgs.__new__(CloudStackServiceAccountTokenArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["seconds_to_live"] = seconds_to_live
            if service_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_account_id'")
            __props__.__dict__["service_account_id"] = service_account_id
            if stack_slug is None and not opts.urn:
                raise TypeError("Missing required property 'stack_slug'")
            __props__.__dict__["stack_slug"] = stack_slug
            __props__.__dict__["expiration"] = None
            __props__.__dict__["has_expired"] = None
            __props__.__dict__["key"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["key"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(CloudStackServiceAccountToken, __self__).__init__(
            'grafana:index/cloudStackServiceAccountToken:CloudStackServiceAccountToken',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            expiration: Optional[pulumi.Input[str]] = None,
            has_expired: Optional[pulumi.Input[bool]] = None,
            key: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            seconds_to_live: Optional[pulumi.Input[int]] = None,
            service_account_id: Optional[pulumi.Input[str]] = None,
            stack_slug: Optional[pulumi.Input[str]] = None) -> 'CloudStackServiceAccountToken':
        """
        Get an existing CloudStackServiceAccountToken resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CloudStackServiceAccountTokenState.__new__(_CloudStackServiceAccountTokenState)

        __props__.__dict__["expiration"] = expiration
        __props__.__dict__["has_expired"] = has_expired
        __props__.__dict__["key"] = key
        __props__.__dict__["name"] = name
        __props__.__dict__["seconds_to_live"] = seconds_to_live
        __props__.__dict__["service_account_id"] = service_account_id
        __props__.__dict__["stack_slug"] = stack_slug
        return CloudStackServiceAccountToken(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def expiration(self) -> pulumi.Output[str]:
        return pulumi.get(self, "expiration")

    @property
    @pulumi.getter(name="hasExpired")
    def has_expired(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "has_expired")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="secondsToLive")
    def seconds_to_live(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "seconds_to_live")

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "service_account_id")

    @property
    @pulumi.getter(name="stackSlug")
    def stack_slug(self) -> pulumi.Output[str]:
        return pulumi.get(self, "stack_slug")

