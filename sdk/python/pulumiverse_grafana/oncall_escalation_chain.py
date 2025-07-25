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
from . import _utilities

__all__ = ['OncallEscalationChainArgs', 'OncallEscalationChain']

@pulumi.input_type
class OncallEscalationChainArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 team_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a OncallEscalationChain resource.
        :param pulumi.Input[builtins.str] name: The name of the escalation chain.
        :param pulumi.Input[builtins.str] team_id: The ID of the OnCall team (using the `on_call_get_team` datasource).
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the escalation chain.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the OnCall team (using the `on_call_get_team` datasource).
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "team_id", value)


@pulumi.input_type
class _OncallEscalationChainState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 team_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering OncallEscalationChain resources.
        :param pulumi.Input[builtins.str] name: The name of the escalation chain.
        :param pulumi.Input[builtins.str] team_id: The ID of the OnCall team (using the `on_call_get_team` datasource).
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the escalation chain.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the OnCall team (using the `on_call_get_team` datasource).
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "team_id", value)


warnings.warn("""grafana.index/oncallescalationchain.OncallEscalationChain has been deprecated in favor of grafana.oncall/escalationchain.EscalationChain""", DeprecationWarning)


@pulumi.type_token("grafana:index/oncallEscalationChain:OncallEscalationChain")
class OncallEscalationChain(pulumi.CustomResource):
    warnings.warn("""grafana.index/oncallescalationchain.OncallEscalationChain has been deprecated in favor of grafana.oncall/escalationchain.EscalationChain""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 team_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/escalation_chains/)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_grafana as grafana
        import pulumiverse_grafana as grafana

        my_team = grafana.oss.get_team(name="my team")
        my_team_get_team = grafana.onCall.get_team(name=my_team.name)
        default = grafana.on_call.EscalationChain("default",
            name="default",
            team_id=my_team_get_team.id)
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/oncallEscalationChain:OncallEscalationChain name "{{ id }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] name: The name of the escalation chain.
        :param pulumi.Input[builtins.str] team_id: The ID of the OnCall team (using the `on_call_get_team` datasource).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[OncallEscalationChainArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/escalation_chains/)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_grafana as grafana
        import pulumiverse_grafana as grafana

        my_team = grafana.oss.get_team(name="my team")
        my_team_get_team = grafana.onCall.get_team(name=my_team.name)
        default = grafana.on_call.EscalationChain("default",
            name="default",
            team_id=my_team_get_team.id)
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/oncallEscalationChain:OncallEscalationChain name "{{ id }}"
        ```

        :param str resource_name: The name of the resource.
        :param OncallEscalationChainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OncallEscalationChainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 team_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        pulumi.log.warn("""OncallEscalationChain is deprecated: grafana.index/oncallescalationchain.OncallEscalationChain has been deprecated in favor of grafana.oncall/escalationchain.EscalationChain""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OncallEscalationChainArgs.__new__(OncallEscalationChainArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["team_id"] = team_id
        super(OncallEscalationChain, __self__).__init__(
            'grafana:index/oncallEscalationChain:OncallEscalationChain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            team_id: Optional[pulumi.Input[builtins.str]] = None) -> 'OncallEscalationChain':
        """
        Get an existing OncallEscalationChain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] name: The name of the escalation chain.
        :param pulumi.Input[builtins.str] team_id: The ID of the OnCall team (using the `on_call_get_team` datasource).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OncallEscalationChainState.__new__(_OncallEscalationChainState)

        __props__.__dict__["name"] = name
        __props__.__dict__["team_id"] = team_id
        return OncallEscalationChain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the escalation chain.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The ID of the OnCall team (using the `on_call_get_team` datasource).
        """
        return pulumi.get(self, "team_id")

