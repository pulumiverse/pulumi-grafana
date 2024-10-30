# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InstallationArgs', 'Installation']

@pulumi.input_type
class InstallationArgs:
    def __init__(__self__, *,
                 metrics_publisher_key: pulumi.Input[str],
                 stack_id: pulumi.Input[str],
                 stack_sm_api_url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Installation resource.
        :param pulumi.Input[str] metrics_publisher_key: The [Grafana Cloud access policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and logs to Grafana Cloud stack.
        :param pulumi.Input[str] stack_id: The ID or slug of the stack to install SM on.
        :param pulumi.Input[str] stack_sm_api_url: The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/set-up-private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
        """
        pulumi.set(__self__, "metrics_publisher_key", metrics_publisher_key)
        pulumi.set(__self__, "stack_id", stack_id)
        if stack_sm_api_url is not None:
            pulumi.set(__self__, "stack_sm_api_url", stack_sm_api_url)

    @property
    @pulumi.getter(name="metricsPublisherKey")
    def metrics_publisher_key(self) -> pulumi.Input[str]:
        """
        The [Grafana Cloud access policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and logs to Grafana Cloud stack.
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
        The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/set-up-private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
        """
        return pulumi.get(self, "stack_sm_api_url")

    @stack_sm_api_url.setter
    def stack_sm_api_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_sm_api_url", value)


@pulumi.input_type
class _InstallationState:
    def __init__(__self__, *,
                 metrics_publisher_key: Optional[pulumi.Input[str]] = None,
                 sm_access_token: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 stack_sm_api_url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Installation resources.
        :param pulumi.Input[str] metrics_publisher_key: The [Grafana Cloud access policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and logs to Grafana Cloud stack.
        :param pulumi.Input[str] sm_access_token: Generated token to access the SM API.
        :param pulumi.Input[str] stack_id: The ID or slug of the stack to install SM on.
        :param pulumi.Input[str] stack_sm_api_url: The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/set-up-private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
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
        The [Grafana Cloud access policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and logs to Grafana Cloud stack.
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
        The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/set-up-private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
        """
        return pulumi.get(self, "stack_sm_api_url")

    @stack_sm_api_url.setter
    def stack_sm_api_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_sm_api_url", value)


class Installation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 metrics_publisher_key: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 stack_sm_api_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Sets up Synthetic Monitoring on a Grafana cloud stack and generates a token.
        Once a Grafana Cloud stack is created, a user can either use this resource or go into the UI to install synthetic monitoring.
        This resource cannot be imported but it can be used on an existing Synthetic Monitoring installation without issues.

        **Note that this resource must be used on a provider configured with Grafana Cloud credentials.**

        * [Official documentation](https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/)
        * [API documentation](https://github.com/grafana/synthetic-monitoring-api-go-client/blob/main/docs/API.md#apiv1registerinstall)

        Required access policy scopes:

        * stacks:read

        ## Example Usage

        ```python
        import pulumi
        import pulumi_grafana as grafana
        import pulumiverse_grafana as grafana

        config = pulumi.Config()
        # Cloud Access Policy token for Grafana Cloud with the following scopes: accesspolicies:read|write|delete, stacks:read|write|delete
        cloud_access_policy_token = config.require_object("cloudAccessPolicyToken")
        stack_slug = config.require_object("stackSlug")
        cloud_region = config.get("cloudRegion")
        if cloud_region is None:
            cloud_region = "us"
        sm_stack = grafana.cloud.Stack("sm_stack",
            name=stack_slug,
            slug=stack_slug,
            region_slug=cloud_region)
        # Step 2: Install Synthetic Monitoring on the stack
        sm_metrics_publish = grafana.cloud.AccessPolicy("sm_metrics_publish",
            region=cloud_region,
            name="metric-publisher-for-sm",
            scopes=[
                "metrics:write",
                "stacks:read",
                "logs:write",
                "traces:write",
            ],
            realms=[grafana.cloud.AccessPolicyRealmArgs(
                type="stack",
                identifier=sm_stack.id,
            )])
        sm_metrics_publish_access_policy_token = grafana.cloud.AccessPolicyToken("sm_metrics_publish",
            region=cloud_region,
            access_policy_id=sm_metrics_publish.policy_id,
            name="metric-publisher-for-sm")
        sm_stack_installation = grafana.synthetic_monitoring.Installation("sm_stack",
            stack_id=sm_stack.id,
            metrics_publisher_key=sm_metrics_publish_access_policy_token.token)
        main = grafana.syntheticMonitoring.get_probes()
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] metrics_publisher_key: The [Grafana Cloud access policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and logs to Grafana Cloud stack.
        :param pulumi.Input[str] stack_id: The ID or slug of the stack to install SM on.
        :param pulumi.Input[str] stack_sm_api_url: The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/set-up-private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstallationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Sets up Synthetic Monitoring on a Grafana cloud stack and generates a token.
        Once a Grafana Cloud stack is created, a user can either use this resource or go into the UI to install synthetic monitoring.
        This resource cannot be imported but it can be used on an existing Synthetic Monitoring installation without issues.

        **Note that this resource must be used on a provider configured with Grafana Cloud credentials.**

        * [Official documentation](https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/)
        * [API documentation](https://github.com/grafana/synthetic-monitoring-api-go-client/blob/main/docs/API.md#apiv1registerinstall)

        Required access policy scopes:

        * stacks:read

        ## Example Usage

        ```python
        import pulumi
        import pulumi_grafana as grafana
        import pulumiverse_grafana as grafana

        config = pulumi.Config()
        # Cloud Access Policy token for Grafana Cloud with the following scopes: accesspolicies:read|write|delete, stacks:read|write|delete
        cloud_access_policy_token = config.require_object("cloudAccessPolicyToken")
        stack_slug = config.require_object("stackSlug")
        cloud_region = config.get("cloudRegion")
        if cloud_region is None:
            cloud_region = "us"
        sm_stack = grafana.cloud.Stack("sm_stack",
            name=stack_slug,
            slug=stack_slug,
            region_slug=cloud_region)
        # Step 2: Install Synthetic Monitoring on the stack
        sm_metrics_publish = grafana.cloud.AccessPolicy("sm_metrics_publish",
            region=cloud_region,
            name="metric-publisher-for-sm",
            scopes=[
                "metrics:write",
                "stacks:read",
                "logs:write",
                "traces:write",
            ],
            realms=[grafana.cloud.AccessPolicyRealmArgs(
                type="stack",
                identifier=sm_stack.id,
            )])
        sm_metrics_publish_access_policy_token = grafana.cloud.AccessPolicyToken("sm_metrics_publish",
            region=cloud_region,
            access_policy_id=sm_metrics_publish.policy_id,
            name="metric-publisher-for-sm")
        sm_stack_installation = grafana.synthetic_monitoring.Installation("sm_stack",
            stack_id=sm_stack.id,
            metrics_publisher_key=sm_metrics_publish_access_policy_token.token)
        main = grafana.syntheticMonitoring.get_probes()
        ```

        :param str resource_name: The name of the resource.
        :param InstallationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstallationArgs, pulumi.ResourceOptions, *args, **kwargs)
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
            __props__ = InstallationArgs.__new__(InstallationArgs)

            if metrics_publisher_key is None and not opts.urn:
                raise TypeError("Missing required property 'metrics_publisher_key'")
            __props__.__dict__["metrics_publisher_key"] = None if metrics_publisher_key is None else pulumi.Output.secret(metrics_publisher_key)
            if stack_id is None and not opts.urn:
                raise TypeError("Missing required property 'stack_id'")
            __props__.__dict__["stack_id"] = stack_id
            __props__.__dict__["stack_sm_api_url"] = stack_sm_api_url
            __props__.__dict__["sm_access_token"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:index/syntheticMonitoringInstallation:SyntheticMonitoringInstallation")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["metricsPublisherKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Installation, __self__).__init__(
            'grafana:syntheticMonitoring/installation:Installation',
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
            stack_sm_api_url: Optional[pulumi.Input[str]] = None) -> 'Installation':
        """
        Get an existing Installation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] metrics_publisher_key: The [Grafana Cloud access policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and logs to Grafana Cloud stack.
        :param pulumi.Input[str] sm_access_token: Generated token to access the SM API.
        :param pulumi.Input[str] stack_id: The ID or slug of the stack to install SM on.
        :param pulumi.Input[str] stack_sm_api_url: The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/set-up-private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstallationState.__new__(_InstallationState)

        __props__.__dict__["metrics_publisher_key"] = metrics_publisher_key
        __props__.__dict__["sm_access_token"] = sm_access_token
        __props__.__dict__["stack_id"] = stack_id
        __props__.__dict__["stack_sm_api_url"] = stack_sm_api_url
        return Installation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="metricsPublisherKey")
    def metrics_publisher_key(self) -> pulumi.Output[str]:
        """
        The [Grafana Cloud access policy](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/) with the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and logs to Grafana Cloud stack.
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
        The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/set-up-private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
        """
        return pulumi.get(self, "stack_sm_api_url")

