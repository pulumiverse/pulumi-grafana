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

__all__ = ['PipelineArgs', 'Pipeline']

@pulumi.input_type
class PipelineArgs:
    def __init__(__self__, *,
                 contents: pulumi.Input[str],
                 enabled: Optional[pulumi.Input[bool]] = None,
                 matchers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Pipeline resource.
        :param pulumi.Input[str] contents: Configuration contents of the pipeline to be used by collectors
        :param pulumi.Input[bool] enabled: Whether the pipeline is enabled for collectors
        :param pulumi.Input[Sequence[pulumi.Input[str]]] matchers: Used to match against collectors and assign pipelines to them; follows the syntax of Prometheus Alertmanager matchers
        :param pulumi.Input[str] name: Name of the pipeline which is the unique identifier for the pipeline
        """
        pulumi.set(__self__, "contents", contents)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if matchers is not None:
            pulumi.set(__self__, "matchers", matchers)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def contents(self) -> pulumi.Input[str]:
        """
        Configuration contents of the pipeline to be used by collectors
        """
        return pulumi.get(self, "contents")

    @contents.setter
    def contents(self, value: pulumi.Input[str]):
        pulumi.set(self, "contents", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the pipeline is enabled for collectors
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def matchers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Used to match against collectors and assign pipelines to them; follows the syntax of Prometheus Alertmanager matchers
        """
        return pulumi.get(self, "matchers")

    @matchers.setter
    def matchers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "matchers", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the pipeline which is the unique identifier for the pipeline
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _PipelineState:
    def __init__(__self__, *,
                 contents: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 matchers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Pipeline resources.
        :param pulumi.Input[str] contents: Configuration contents of the pipeline to be used by collectors
        :param pulumi.Input[bool] enabled: Whether the pipeline is enabled for collectors
        :param pulumi.Input[Sequence[pulumi.Input[str]]] matchers: Used to match against collectors and assign pipelines to them; follows the syntax of Prometheus Alertmanager matchers
        :param pulumi.Input[str] name: Name of the pipeline which is the unique identifier for the pipeline
        """
        if contents is not None:
            pulumi.set(__self__, "contents", contents)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if matchers is not None:
            pulumi.set(__self__, "matchers", matchers)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def contents(self) -> Optional[pulumi.Input[str]]:
        """
        Configuration contents of the pipeline to be used by collectors
        """
        return pulumi.get(self, "contents")

    @contents.setter
    def contents(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "contents", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the pipeline is enabled for collectors
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def matchers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Used to match against collectors and assign pipelines to them; follows the syntax of Prometheus Alertmanager matchers
        """
        return pulumi.get(self, "matchers")

    @matchers.setter
    def matchers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "matchers", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the pipeline which is the unique identifier for the pipeline
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class Pipeline(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 contents: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 matchers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages Grafana Fleet Management pipelines.

        * [Official documentation](https://grafana.com/docs/grafana-cloud/send-data/fleet-management/)
        * [API documentation](https://grafana.com/docs/grafana-cloud/send-data/fleet-management/api-reference/pipeline-api/)

        Required access policy scopes:

        * fleet-management:read
        * fleet-management:write

        ## Import

        ```sh
        $ pulumi import grafana:fleetManagement/pipeline:Pipeline name "{{ name }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] contents: Configuration contents of the pipeline to be used by collectors
        :param pulumi.Input[bool] enabled: Whether the pipeline is enabled for collectors
        :param pulumi.Input[Sequence[pulumi.Input[str]]] matchers: Used to match against collectors and assign pipelines to them; follows the syntax of Prometheus Alertmanager matchers
        :param pulumi.Input[str] name: Name of the pipeline which is the unique identifier for the pipeline
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PipelineArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Grafana Fleet Management pipelines.

        * [Official documentation](https://grafana.com/docs/grafana-cloud/send-data/fleet-management/)
        * [API documentation](https://grafana.com/docs/grafana-cloud/send-data/fleet-management/api-reference/pipeline-api/)

        Required access policy scopes:

        * fleet-management:read
        * fleet-management:write

        ## Import

        ```sh
        $ pulumi import grafana:fleetManagement/pipeline:Pipeline name "{{ name }}"
        ```

        :param str resource_name: The name of the resource.
        :param PipelineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PipelineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 contents: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 matchers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PipelineArgs.__new__(PipelineArgs)

            if contents is None and not opts.urn:
                raise TypeError("Missing required property 'contents'")
            __props__.__dict__["contents"] = contents
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["matchers"] = matchers
            __props__.__dict__["name"] = name
        super(Pipeline, __self__).__init__(
            'grafana:fleetManagement/pipeline:Pipeline',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            contents: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            matchers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'Pipeline':
        """
        Get an existing Pipeline resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] contents: Configuration contents of the pipeline to be used by collectors
        :param pulumi.Input[bool] enabled: Whether the pipeline is enabled for collectors
        :param pulumi.Input[Sequence[pulumi.Input[str]]] matchers: Used to match against collectors and assign pipelines to them; follows the syntax of Prometheus Alertmanager matchers
        :param pulumi.Input[str] name: Name of the pipeline which is the unique identifier for the pipeline
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PipelineState.__new__(_PipelineState)

        __props__.__dict__["contents"] = contents
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["matchers"] = matchers
        __props__.__dict__["name"] = name
        return Pipeline(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def contents(self) -> pulumi.Output[str]:
        """
        Configuration contents of the pipeline to be used by collectors
        """
        return pulumi.get(self, "contents")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Whether the pipeline is enabled for collectors
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def matchers(self) -> pulumi.Output[Sequence[str]]:
        """
        Used to match against collectors and assign pipelines to them; follows the syntax of Prometheus Alertmanager matchers
        """
        return pulumi.get(self, "matchers")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the pipeline which is the unique identifier for the pipeline
        """
        return pulumi.get(self, "name")

