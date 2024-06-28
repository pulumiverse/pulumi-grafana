# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SyntheticMonitoringProbeArgs', 'SyntheticMonitoringProbe']

@pulumi.input_type
class SyntheticMonitoringProbeArgs:
    def __init__(__self__, *,
                 latitude: pulumi.Input[float],
                 longitude: pulumi.Input[float],
                 region: pulumi.Input[str],
                 disable_scripted_checks: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a SyntheticMonitoringProbe resource.
        :param pulumi.Input[float] latitude: Latitude coordinates.
        :param pulumi.Input[float] longitude: Longitude coordinates.
        :param pulumi.Input[str] region: Region of the probe.
        :param pulumi.Input[bool] disable_scripted_checks: Disables scripted checks for this probe. Defaults to `false`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Custom labels to be included with collected metrics and logs.
        :param pulumi.Input[str] name: Name of the probe.
        :param pulumi.Input[bool] public: Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
        """
        pulumi.set(__self__, "latitude", latitude)
        pulumi.set(__self__, "longitude", longitude)
        pulumi.set(__self__, "region", region)
        if disable_scripted_checks is not None:
            pulumi.set(__self__, "disable_scripted_checks", disable_scripted_checks)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if public is not None:
            pulumi.set(__self__, "public", public)

    @property
    @pulumi.getter
    def latitude(self) -> pulumi.Input[float]:
        """
        Latitude coordinates.
        """
        return pulumi.get(self, "latitude")

    @latitude.setter
    def latitude(self, value: pulumi.Input[float]):
        pulumi.set(self, "latitude", value)

    @property
    @pulumi.getter
    def longitude(self) -> pulumi.Input[float]:
        """
        Longitude coordinates.
        """
        return pulumi.get(self, "longitude")

    @longitude.setter
    def longitude(self, value: pulumi.Input[float]):
        pulumi.set(self, "longitude", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        Region of the probe.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="disableScriptedChecks")
    def disable_scripted_checks(self) -> Optional[pulumi.Input[bool]]:
        """
        Disables scripted checks for this probe. Defaults to `false`.
        """
        return pulumi.get(self, "disable_scripted_checks")

    @disable_scripted_checks.setter
    def disable_scripted_checks(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_scripted_checks", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Custom labels to be included with collected metrics and logs.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the probe.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def public(self) -> Optional[pulumi.Input[bool]]:
        """
        Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
        """
        return pulumi.get(self, "public")

    @public.setter
    def public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "public", value)


@pulumi.input_type
class _SyntheticMonitoringProbeState:
    def __init__(__self__, *,
                 auth_token: Optional[pulumi.Input[str]] = None,
                 disable_scripted_checks: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 latitude: Optional[pulumi.Input[float]] = None,
                 longitude: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering SyntheticMonitoringProbe resources.
        :param pulumi.Input[str] auth_token: The probe authentication token. Your probe must use this to authenticate with Grafana Cloud.
        :param pulumi.Input[bool] disable_scripted_checks: Disables scripted checks for this probe. Defaults to `false`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Custom labels to be included with collected metrics and logs.
        :param pulumi.Input[float] latitude: Latitude coordinates.
        :param pulumi.Input[float] longitude: Longitude coordinates.
        :param pulumi.Input[str] name: Name of the probe.
        :param pulumi.Input[bool] public: Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
        :param pulumi.Input[str] region: Region of the probe.
        :param pulumi.Input[int] tenant_id: The tenant ID of the probe.
        """
        if auth_token is not None:
            pulumi.set(__self__, "auth_token", auth_token)
        if disable_scripted_checks is not None:
            pulumi.set(__self__, "disable_scripted_checks", disable_scripted_checks)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if latitude is not None:
            pulumi.set(__self__, "latitude", latitude)
        if longitude is not None:
            pulumi.set(__self__, "longitude", longitude)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if public is not None:
            pulumi.set(__self__, "public", public)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="authToken")
    def auth_token(self) -> Optional[pulumi.Input[str]]:
        """
        The probe authentication token. Your probe must use this to authenticate with Grafana Cloud.
        """
        return pulumi.get(self, "auth_token")

    @auth_token.setter
    def auth_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_token", value)

    @property
    @pulumi.getter(name="disableScriptedChecks")
    def disable_scripted_checks(self) -> Optional[pulumi.Input[bool]]:
        """
        Disables scripted checks for this probe. Defaults to `false`.
        """
        return pulumi.get(self, "disable_scripted_checks")

    @disable_scripted_checks.setter
    def disable_scripted_checks(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_scripted_checks", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Custom labels to be included with collected metrics and logs.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def latitude(self) -> Optional[pulumi.Input[float]]:
        """
        Latitude coordinates.
        """
        return pulumi.get(self, "latitude")

    @latitude.setter
    def latitude(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "latitude", value)

    @property
    @pulumi.getter
    def longitude(self) -> Optional[pulumi.Input[float]]:
        """
        Longitude coordinates.
        """
        return pulumi.get(self, "longitude")

    @longitude.setter
    def longitude(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "longitude", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the probe.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def public(self) -> Optional[pulumi.Input[bool]]:
        """
        Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
        """
        return pulumi.get(self, "public")

    @public.setter
    def public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "public", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Region of the probe.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[int]]:
        """
        The tenant ID of the probe.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "tenant_id", value)


warnings.warn("""grafana.index/syntheticmonitoringprobe.SyntheticMonitoringProbe has been deprecated in favor of grafana.syntheticmonitoring/probe.Probe""", DeprecationWarning)


class SyntheticMonitoringProbe(pulumi.CustomResource):
    warnings.warn("""grafana.index/syntheticmonitoringprobe.SyntheticMonitoringProbe has been deprecated in favor of grafana.syntheticmonitoring/probe.Probe""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_scripted_checks: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 latitude: Optional[pulumi.Input[float]] = None,
                 longitude: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Besides the public probes run by Grafana Labs, you can also install your
        own private probes. These are only accessible to you and only write data to
        your Grafana Cloud account. Private probes are instances of the open source
        Grafana Synthetic Monitoring Agent.

        * [Official documentation](https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/set-up-private-probes/)

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        main = grafana.synthetic_monitoring.Probe("main",
            labels={
                "type": "mountain",
            },
            latitude=27.98606,
            longitude=86.92262,
            region="APAC")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/syntheticMonitoringProbe:SyntheticMonitoringProbe name "{{ id }}"
        ```

        ```sh
        $ pulumi import grafana:index/syntheticMonitoringProbe:SyntheticMonitoringProbe name "{{ id }}:{{ authToken }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_scripted_checks: Disables scripted checks for this probe. Defaults to `false`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Custom labels to be included with collected metrics and logs.
        :param pulumi.Input[float] latitude: Latitude coordinates.
        :param pulumi.Input[float] longitude: Longitude coordinates.
        :param pulumi.Input[str] name: Name of the probe.
        :param pulumi.Input[bool] public: Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
        :param pulumi.Input[str] region: Region of the probe.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SyntheticMonitoringProbeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Besides the public probes run by Grafana Labs, you can also install your
        own private probes. These are only accessible to you and only write data to
        your Grafana Cloud account. Private probes are instances of the open source
        Grafana Synthetic Monitoring Agent.

        * [Official documentation](https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/set-up-private-probes/)

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        main = grafana.synthetic_monitoring.Probe("main",
            labels={
                "type": "mountain",
            },
            latitude=27.98606,
            longitude=86.92262,
            region="APAC")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/syntheticMonitoringProbe:SyntheticMonitoringProbe name "{{ id }}"
        ```

        ```sh
        $ pulumi import grafana:index/syntheticMonitoringProbe:SyntheticMonitoringProbe name "{{ id }}:{{ authToken }}"
        ```

        :param str resource_name: The name of the resource.
        :param SyntheticMonitoringProbeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SyntheticMonitoringProbeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_scripted_checks: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 latitude: Optional[pulumi.Input[float]] = None,
                 longitude: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""SyntheticMonitoringProbe is deprecated: grafana.index/syntheticmonitoringprobe.SyntheticMonitoringProbe has been deprecated in favor of grafana.syntheticmonitoring/probe.Probe""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SyntheticMonitoringProbeArgs.__new__(SyntheticMonitoringProbeArgs)

            __props__.__dict__["disable_scripted_checks"] = disable_scripted_checks
            __props__.__dict__["labels"] = labels
            if latitude is None and not opts.urn:
                raise TypeError("Missing required property 'latitude'")
            __props__.__dict__["latitude"] = latitude
            if longitude is None and not opts.urn:
                raise TypeError("Missing required property 'longitude'")
            __props__.__dict__["longitude"] = longitude
            __props__.__dict__["name"] = name
            __props__.__dict__["public"] = public
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["auth_token"] = None
            __props__.__dict__["tenant_id"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:index/syntheticMonitoringProbe:SyntheticMonitoringProbe")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["authToken"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SyntheticMonitoringProbe, __self__).__init__(
            'grafana:index/syntheticMonitoringProbe:SyntheticMonitoringProbe',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_token: Optional[pulumi.Input[str]] = None,
            disable_scripted_checks: Optional[pulumi.Input[bool]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            latitude: Optional[pulumi.Input[float]] = None,
            longitude: Optional[pulumi.Input[float]] = None,
            name: Optional[pulumi.Input[str]] = None,
            public: Optional[pulumi.Input[bool]] = None,
            region: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[int]] = None) -> 'SyntheticMonitoringProbe':
        """
        Get an existing SyntheticMonitoringProbe resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_token: The probe authentication token. Your probe must use this to authenticate with Grafana Cloud.
        :param pulumi.Input[bool] disable_scripted_checks: Disables scripted checks for this probe. Defaults to `false`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Custom labels to be included with collected metrics and logs.
        :param pulumi.Input[float] latitude: Latitude coordinates.
        :param pulumi.Input[float] longitude: Longitude coordinates.
        :param pulumi.Input[str] name: Name of the probe.
        :param pulumi.Input[bool] public: Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
        :param pulumi.Input[str] region: Region of the probe.
        :param pulumi.Input[int] tenant_id: The tenant ID of the probe.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SyntheticMonitoringProbeState.__new__(_SyntheticMonitoringProbeState)

        __props__.__dict__["auth_token"] = auth_token
        __props__.__dict__["disable_scripted_checks"] = disable_scripted_checks
        __props__.__dict__["labels"] = labels
        __props__.__dict__["latitude"] = latitude
        __props__.__dict__["longitude"] = longitude
        __props__.__dict__["name"] = name
        __props__.__dict__["public"] = public
        __props__.__dict__["region"] = region
        __props__.__dict__["tenant_id"] = tenant_id
        return SyntheticMonitoringProbe(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authToken")
    def auth_token(self) -> pulumi.Output[str]:
        """
        The probe authentication token. Your probe must use this to authenticate with Grafana Cloud.
        """
        return pulumi.get(self, "auth_token")

    @property
    @pulumi.getter(name="disableScriptedChecks")
    def disable_scripted_checks(self) -> pulumi.Output[Optional[bool]]:
        """
        Disables scripted checks for this probe. Defaults to `false`.
        """
        return pulumi.get(self, "disable_scripted_checks")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Custom labels to be included with collected metrics and logs.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def latitude(self) -> pulumi.Output[float]:
        """
        Latitude coordinates.
        """
        return pulumi.get(self, "latitude")

    @property
    @pulumi.getter
    def longitude(self) -> pulumi.Output[float]:
        """
        Longitude coordinates.
        """
        return pulumi.get(self, "longitude")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the probe.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def public(self) -> pulumi.Output[Optional[bool]]:
        """
        Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
        """
        return pulumi.get(self, "public")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Region of the probe.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[int]:
        """
        The tenant ID of the probe.
        """
        return pulumi.get(self, "tenant_id")

