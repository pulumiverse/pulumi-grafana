# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SyntheticMonitoringInstallationArgs', 'SyntheticMonitoringInstallation']

@pulumi.input_type
class SyntheticMonitoringInstallationArgs:
    def __init__(__self__, *,
                 metrics_publisher_key: pulumi.Input[str],
                 stack_id: pulumi.Input[str],
                 stack_sm_api_url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SyntheticMonitoringInstallation resource.
        :param pulumi.Input[str] metrics_publisher_key: The [Grafana Cloud access
               policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with
               the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and
               logs to Grafana Cloud stack.
        :param pulumi.Input[str] stack_id: The ID or slug of the stack to install SM on.
        :param pulumi.Input[str] stack_sm_api_url: The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here:
               https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/set-up/set-up-private-probes/#probe-api-server-url. A
               static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region,
               this field is computed automatically and readable.
        """
        pulumi.set(__self__, "metrics_publisher_key", metrics_publisher_key)
        pulumi.set(__self__, "stack_id", stack_id)
        if stack_sm_api_url is not None:
            pulumi.set(__self__, "stack_sm_api_url", stack_sm_api_url)

    @property
    @pulumi.getter(name="metricsPublisherKey")
    def metrics_publisher_key(self) -> pulumi.Input[str]:
        """
        The [Grafana Cloud access
        policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with
        the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and
        logs to Grafana Cloud stack.
        """
        return pulumi.get(self, "metrics_publisher_key")

    @metrics_publisher_key.setter
    def metrics_publisher_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "metrics_publisher_key", value)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Input[str]:
        """
        The ID or slug of the stack to install SM on.
        """
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "stack_id", value)

    @property
    @pulumi.getter(name="stackSmApiUrl")
    def stack_sm_api_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here:
        https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/set-up/set-up-private-probes/#probe-api-server-url. A
        static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region,
        this field is computed automatically and readable.
        """
        return pulumi.get(self, "stack_sm_api_url")

    @stack_sm_api_url.setter
    def stack_sm_api_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_sm_api_url", value)


@pulumi.input_type
class _SyntheticMonitoringInstallationState:
    def __init__(__self__, *,
                 metrics_publisher_key: Optional[pulumi.Input[str]] = None,
                 sm_access_token: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 stack_sm_api_url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SyntheticMonitoringInstallation resources.
        :param pulumi.Input[str] metrics_publisher_key: The [Grafana Cloud access
               policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with
               the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and
               logs to Grafana Cloud stack.
        :param pulumi.Input[str] sm_access_token: Generated token to access the SM API.
        :param pulumi.Input[str] stack_id: The ID or slug of the stack to install SM on.
        :param pulumi.Input[str] stack_sm_api_url: The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here:
               https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/set-up/set-up-private-probes/#probe-api-server-url. A
               static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region,
               this field is computed automatically and readable.
        """
        if metrics_publisher_key is not None:
            pulumi.set(__self__, "metrics_publisher_key", metrics_publisher_key)
        if sm_access_token is not None:
            pulumi.set(__self__, "sm_access_token", sm_access_token)
        if stack_id is not None:
            pulumi.set(__self__, "stack_id", stack_id)
        if stack_sm_api_url is not None:
            pulumi.set(__self__, "stack_sm_api_url", stack_sm_api_url)

    @property
    @pulumi.getter(name="metricsPublisherKey")
    def metrics_publisher_key(self) -> Optional[pulumi.Input[str]]:
        """
        The [Grafana Cloud access
        policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with
        the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and
        logs to Grafana Cloud stack.
        """
        return pulumi.get(self, "metrics_publisher_key")

    @metrics_publisher_key.setter
    def metrics_publisher_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metrics_publisher_key", value)

    @property
    @pulumi.getter(name="smAccessToken")
    def sm_access_token(self) -> Optional[pulumi.Input[str]]:
        """
        Generated token to access the SM API.
        """
        return pulumi.get(self, "sm_access_token")

    @sm_access_token.setter
    def sm_access_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sm_access_token", value)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or slug of the stack to install SM on.
        """
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_id", value)

    @property
    @pulumi.getter(name="stackSmApiUrl")
    def stack_sm_api_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here:
        https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/set-up/set-up-private-probes/#probe-api-server-url. A
        static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region,
        this field is computed automatically and readable.
        """
        return pulumi.get(self, "stack_sm_api_url")

    @stack_sm_api_url.setter
    def stack_sm_api_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_sm_api_url", value)


class SyntheticMonitoringInstallation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 metrics_publisher_key: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 stack_sm_api_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SyntheticMonitoringInstallation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] metrics_publisher_key: The [Grafana Cloud access
               policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with
               the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and
               logs to Grafana Cloud stack.
        :param pulumi.Input[str] stack_id: The ID or slug of the stack to install SM on.
        :param pulumi.Input[str] stack_sm_api_url: The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here:
               https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/set-up/set-up-private-probes/#probe-api-server-url. A
               static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region,
               this field is computed automatically and readable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SyntheticMonitoringInstallationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SyntheticMonitoringInstallation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SyntheticMonitoringInstallationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SyntheticMonitoringInstallationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 metrics_publisher_key: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 stack_sm_api_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SyntheticMonitoringInstallationArgs.__new__(SyntheticMonitoringInstallationArgs)

            if metrics_publisher_key is None and not opts.urn:
                raise TypeError("Missing required property 'metrics_publisher_key'")
            __props__.__dict__["metrics_publisher_key"] = None if metrics_publisher_key is None else pulumi.Output.secret(metrics_publisher_key)
            if stack_id is None and not opts.urn:
                raise TypeError("Missing required property 'stack_id'")
            __props__.__dict__["stack_id"] = stack_id
            __props__.__dict__["stack_sm_api_url"] = stack_sm_api_url
            __props__.__dict__["sm_access_token"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["metricsPublisherKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SyntheticMonitoringInstallation, __self__).__init__(
            'grafana:index/syntheticMonitoringInstallation:SyntheticMonitoringInstallation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            metrics_publisher_key: Optional[pulumi.Input[str]] = None,
            sm_access_token: Optional[pulumi.Input[str]] = None,
            stack_id: Optional[pulumi.Input[str]] = None,
            stack_sm_api_url: Optional[pulumi.Input[str]] = None) -> 'SyntheticMonitoringInstallation':
        """
        Get an existing SyntheticMonitoringInstallation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] metrics_publisher_key: The [Grafana Cloud access
               policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with
               the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and
               logs to Grafana Cloud stack.
        :param pulumi.Input[str] sm_access_token: Generated token to access the SM API.
        :param pulumi.Input[str] stack_id: The ID or slug of the stack to install SM on.
        :param pulumi.Input[str] stack_sm_api_url: The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here:
               https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/set-up/set-up-private-probes/#probe-api-server-url. A
               static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region,
               this field is computed automatically and readable.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SyntheticMonitoringInstallationState.__new__(_SyntheticMonitoringInstallationState)

        __props__.__dict__["metrics_publisher_key"] = metrics_publisher_key
        __props__.__dict__["sm_access_token"] = sm_access_token
        __props__.__dict__["stack_id"] = stack_id
        __props__.__dict__["stack_sm_api_url"] = stack_sm_api_url
        return SyntheticMonitoringInstallation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="metricsPublisherKey")
    def metrics_publisher_key(self) -> pulumi.Output[str]:
        """
        The [Grafana Cloud access
        policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with
        the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and
        logs to Grafana Cloud stack.
        """
        return pulumi.get(self, "metrics_publisher_key")

    @property
    @pulumi.getter(name="smAccessToken")
    def sm_access_token(self) -> pulumi.Output[str]:
        """
        Generated token to access the SM API.
        """
        return pulumi.get(self, "sm_access_token")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Output[str]:
        """
        The ID or slug of the stack to install SM on.
        """
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter(name="stackSmApiUrl")
    def stack_sm_api_url(self) -> pulumi.Output[str]:
        """
        The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here:
        https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/set-up/set-up-private-probes/#probe-api-server-url. A
        static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region,
        this field is computed automatically and readable.
        """
        return pulumi.get(self, "stack_sm_api_url")

