# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = ['AppsDashboardArgs', 'AppsDashboard']

@pulumi.input_type
class AppsDashboardArgs:
    def __init__(__self__, *,
                 metadata: Optional[pulumi.Input['AppsDashboardMetadataArgs']] = None,
                 options: Optional[pulumi.Input['AppsDashboardOptionsArgs']] = None,
                 spec: Optional[pulumi.Input['AppsDashboardSpecArgs']] = None):
        """
        The set of arguments for constructing a AppsDashboard resource.
        :param pulumi.Input['AppsDashboardMetadataArgs'] metadata: The metadata of the resource.
        :param pulumi.Input['AppsDashboardOptionsArgs'] options: Options for applying the resource.
        :param pulumi.Input['AppsDashboardSpecArgs'] spec: The spec of the resource.
        """
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if options is not None:
            pulumi.set(__self__, "options", options)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['AppsDashboardMetadataArgs']]:
        """
        The metadata of the resource.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['AppsDashboardMetadataArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input['AppsDashboardOptionsArgs']]:
        """
        Options for applying the resource.
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input['AppsDashboardOptionsArgs']]):
        pulumi.set(self, "options", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['AppsDashboardSpecArgs']]:
        """
        The spec of the resource.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['AppsDashboardSpecArgs']]):
        pulumi.set(self, "spec", value)


@pulumi.input_type
class _AppsDashboardState:
    def __init__(__self__, *,
                 metadata: Optional[pulumi.Input['AppsDashboardMetadataArgs']] = None,
                 options: Optional[pulumi.Input['AppsDashboardOptionsArgs']] = None,
                 spec: Optional[pulumi.Input['AppsDashboardSpecArgs']] = None):
        """
        Input properties used for looking up and filtering AppsDashboard resources.
        :param pulumi.Input['AppsDashboardMetadataArgs'] metadata: The metadata of the resource.
        :param pulumi.Input['AppsDashboardOptionsArgs'] options: Options for applying the resource.
        :param pulumi.Input['AppsDashboardSpecArgs'] spec: The spec of the resource.
        """
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if options is not None:
            pulumi.set(__self__, "options", options)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['AppsDashboardMetadataArgs']]:
        """
        The metadata of the resource.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['AppsDashboardMetadataArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input['AppsDashboardOptionsArgs']]:
        """
        Options for applying the resource.
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input['AppsDashboardOptionsArgs']]):
        pulumi.set(self, "options", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['AppsDashboardSpecArgs']]:
        """
        The spec of the resource.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['AppsDashboardSpecArgs']]):
        pulumi.set(self, "spec", value)


@pulumi.type_token("grafana:experimental/appsDashboard:AppsDashboard")
class AppsDashboard(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 metadata: Optional[pulumi.Input[Union['AppsDashboardMetadataArgs', 'AppsDashboardMetadataArgsDict']]] = None,
                 options: Optional[pulumi.Input[Union['AppsDashboardOptionsArgs', 'AppsDashboardOptionsArgsDict']]] = None,
                 spec: Optional[pulumi.Input[Union['AppsDashboardSpecArgs', 'AppsDashboardSpecArgsDict']]] = None,
                 __props__=None):
        """
        Manages Grafana dashboards using the new Grafana APIs.

        * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/dashboard/#new-dashboard-apis)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['AppsDashboardMetadataArgs', 'AppsDashboardMetadataArgsDict']] metadata: The metadata of the resource.
        :param pulumi.Input[Union['AppsDashboardOptionsArgs', 'AppsDashboardOptionsArgsDict']] options: Options for applying the resource.
        :param pulumi.Input[Union['AppsDashboardSpecArgs', 'AppsDashboardSpecArgsDict']] spec: The spec of the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AppsDashboardArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Grafana dashboards using the new Grafana APIs.

        * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/dashboard/#new-dashboard-apis)

        :param str resource_name: The name of the resource.
        :param AppsDashboardArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppsDashboardArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 metadata: Optional[pulumi.Input[Union['AppsDashboardMetadataArgs', 'AppsDashboardMetadataArgsDict']]] = None,
                 options: Optional[pulumi.Input[Union['AppsDashboardOptionsArgs', 'AppsDashboardOptionsArgsDict']]] = None,
                 spec: Optional[pulumi.Input[Union['AppsDashboardSpecArgs', 'AppsDashboardSpecArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppsDashboardArgs.__new__(AppsDashboardArgs)

            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["options"] = options
            __props__.__dict__["spec"] = spec
        super(AppsDashboard, __self__).__init__(
            'grafana:experimental/appsDashboard:AppsDashboard',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            metadata: Optional[pulumi.Input[Union['AppsDashboardMetadataArgs', 'AppsDashboardMetadataArgsDict']]] = None,
            options: Optional[pulumi.Input[Union['AppsDashboardOptionsArgs', 'AppsDashboardOptionsArgsDict']]] = None,
            spec: Optional[pulumi.Input[Union['AppsDashboardSpecArgs', 'AppsDashboardSpecArgsDict']]] = None) -> 'AppsDashboard':
        """
        Get an existing AppsDashboard resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['AppsDashboardMetadataArgs', 'AppsDashboardMetadataArgsDict']] metadata: The metadata of the resource.
        :param pulumi.Input[Union['AppsDashboardOptionsArgs', 'AppsDashboardOptionsArgsDict']] options: Options for applying the resource.
        :param pulumi.Input[Union['AppsDashboardSpecArgs', 'AppsDashboardSpecArgsDict']] spec: The spec of the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppsDashboardState.__new__(_AppsDashboardState)

        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["options"] = options
        __props__.__dict__["spec"] = spec
        return AppsDashboard(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional['outputs.AppsDashboardMetadata']]:
        """
        The metadata of the resource.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def options(self) -> pulumi.Output[Optional['outputs.AppsDashboardOptions']]:
        """
        Options for applying the resource.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Output[Optional['outputs.AppsDashboardSpec']]:
        """
        The spec of the resource.
        """
        return pulumi.get(self, "spec")

