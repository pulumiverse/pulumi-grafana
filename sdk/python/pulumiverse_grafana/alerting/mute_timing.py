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

__all__ = ['MuteTimingArgs', 'MuteTiming']

@pulumi.input_type
class MuteTimingArgs:
    def __init__(__self__, *,
                 disable_provenance: Optional[pulumi.Input[bool]] = None,
                 intervals: Optional[pulumi.Input[Sequence[pulumi.Input['MuteTimingIntervalArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MuteTiming resource.
        :param pulumi.Input[Sequence[pulumi.Input['MuteTimingIntervalArgs']]] intervals: The time intervals at which to mute notifications. Use an empty block to mute all the time.
        :param pulumi.Input[str] name: The name of the mute timing.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        if disable_provenance is not None:
            pulumi.set(__self__, "disable_provenance", disable_provenance)
        if intervals is not None:
            pulumi.set(__self__, "intervals", intervals)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)

    @property
    @pulumi.getter(name="disableProvenance")
    def disable_provenance(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "disable_provenance")

    @disable_provenance.setter
    def disable_provenance(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_provenance", value)

    @property
    @pulumi.getter
    def intervals(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MuteTimingIntervalArgs']]]]:
        """
        The time intervals at which to mute notifications. Use an empty block to mute all the time.
        """
        return pulumi.get(self, "intervals")

    @intervals.setter
    def intervals(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MuteTimingIntervalArgs']]]]):
        pulumi.set(self, "intervals", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the mute timing.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)


@pulumi.input_type
class _MuteTimingState:
    def __init__(__self__, *,
                 disable_provenance: Optional[pulumi.Input[bool]] = None,
                 intervals: Optional[pulumi.Input[Sequence[pulumi.Input['MuteTimingIntervalArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MuteTiming resources.
        :param pulumi.Input[Sequence[pulumi.Input['MuteTimingIntervalArgs']]] intervals: The time intervals at which to mute notifications. Use an empty block to mute all the time.
        :param pulumi.Input[str] name: The name of the mute timing.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        if disable_provenance is not None:
            pulumi.set(__self__, "disable_provenance", disable_provenance)
        if intervals is not None:
            pulumi.set(__self__, "intervals", intervals)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)

    @property
    @pulumi.getter(name="disableProvenance")
    def disable_provenance(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "disable_provenance")

    @disable_provenance.setter
    def disable_provenance(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_provenance", value)

    @property
    @pulumi.getter
    def intervals(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MuteTimingIntervalArgs']]]]:
        """
        The time intervals at which to mute notifications. Use an empty block to mute all the time.
        """
        return pulumi.get(self, "intervals")

    @intervals.setter
    def intervals(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MuteTimingIntervalArgs']]]]):
        pulumi.set(self, "intervals", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the mute timing.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)


class MuteTiming(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_provenance: Optional[pulumi.Input[bool]] = None,
                 intervals: Optional[pulumi.Input[Sequence[pulumi.Input[Union['MuteTimingIntervalArgs', 'MuteTimingIntervalArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages Grafana Alerting mute timings.

        * Official documentation
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/alerting_provisioning/#mute-timings)

        This resource requires Grafana 9.1.0 or later.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        my_mute_timing = grafana.alerting.MuteTiming("my_mute_timing",
            name="My Mute Timing",
            intervals=[{
                "times": [{
                    "start": "04:56",
                    "end": "14:17",
                }],
                "weekdays": [
                    "monday",
                    "tuesday:thursday",
                ],
                "days_of_months": [
                    "1:7",
                    "-1",
                ],
                "months": [
                    "1:3",
                    "december",
                ],
                "years": [
                    "2030",
                    "2025:2026",
                ],
                "location": "America/New_York",
            }])
        ```

        ## Import

        ```sh
        $ pulumi import grafana:alerting/muteTiming:MuteTiming name "{{ name }}"
        ```

        ```sh
        $ pulumi import grafana:alerting/muteTiming:MuteTiming name "{{ orgID }}:{{ name }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['MuteTimingIntervalArgs', 'MuteTimingIntervalArgsDict']]]] intervals: The time intervals at which to mute notifications. Use an empty block to mute all the time.
        :param pulumi.Input[str] name: The name of the mute timing.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[MuteTimingArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Grafana Alerting mute timings.

        * Official documentation
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/alerting_provisioning/#mute-timings)

        This resource requires Grafana 9.1.0 or later.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        my_mute_timing = grafana.alerting.MuteTiming("my_mute_timing",
            name="My Mute Timing",
            intervals=[{
                "times": [{
                    "start": "04:56",
                    "end": "14:17",
                }],
                "weekdays": [
                    "monday",
                    "tuesday:thursday",
                ],
                "days_of_months": [
                    "1:7",
                    "-1",
                ],
                "months": [
                    "1:3",
                    "december",
                ],
                "years": [
                    "2030",
                    "2025:2026",
                ],
                "location": "America/New_York",
            }])
        ```

        ## Import

        ```sh
        $ pulumi import grafana:alerting/muteTiming:MuteTiming name "{{ name }}"
        ```

        ```sh
        $ pulumi import grafana:alerting/muteTiming:MuteTiming name "{{ orgID }}:{{ name }}"
        ```

        :param str resource_name: The name of the resource.
        :param MuteTimingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MuteTimingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_provenance: Optional[pulumi.Input[bool]] = None,
                 intervals: Optional[pulumi.Input[Sequence[pulumi.Input[Union['MuteTimingIntervalArgs', 'MuteTimingIntervalArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MuteTimingArgs.__new__(MuteTimingArgs)

            __props__.__dict__["disable_provenance"] = disable_provenance
            __props__.__dict__["intervals"] = intervals
            __props__.__dict__["name"] = name
            __props__.__dict__["org_id"] = org_id
        super(MuteTiming, __self__).__init__(
            'grafana:alerting/muteTiming:MuteTiming',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disable_provenance: Optional[pulumi.Input[bool]] = None,
            intervals: Optional[pulumi.Input[Sequence[pulumi.Input[Union['MuteTimingIntervalArgs', 'MuteTimingIntervalArgsDict']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None) -> 'MuteTiming':
        """
        Get an existing MuteTiming resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['MuteTimingIntervalArgs', 'MuteTimingIntervalArgsDict']]]] intervals: The time intervals at which to mute notifications. Use an empty block to mute all the time.
        :param pulumi.Input[str] name: The name of the mute timing.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MuteTimingState.__new__(_MuteTimingState)

        __props__.__dict__["disable_provenance"] = disable_provenance
        __props__.__dict__["intervals"] = intervals
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        return MuteTiming(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="disableProvenance")
    def disable_provenance(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "disable_provenance")

    @property
    @pulumi.getter
    def intervals(self) -> pulumi.Output[Optional[Sequence['outputs.MuteTimingInterval']]]:
        """
        The time intervals at which to mute notifications. Use an empty block to mute all the time.
        """
        return pulumi.get(self, "intervals")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the mute timing.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

