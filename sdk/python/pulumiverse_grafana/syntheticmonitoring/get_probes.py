# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetProbesResult',
    'AwaitableGetProbesResult',
    'get_probes',
    'get_probes_output',
]

@pulumi.output_type
class GetProbesResult:
    """
    A collection of values returned by getProbes.
    """
    def __init__(__self__, filter_deprecated=None, id=None, probes=None):
        if filter_deprecated and not isinstance(filter_deprecated, bool):
            raise TypeError("Expected argument 'filter_deprecated' to be a bool")
        pulumi.set(__self__, "filter_deprecated", filter_deprecated)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if probes and not isinstance(probes, dict):
            raise TypeError("Expected argument 'probes' to be a dict")
        pulumi.set(__self__, "probes", probes)

    @property
    @pulumi.getter(name="filterDeprecated")
    def filter_deprecated(self) -> Optional[bool]:
        """
        If true, only probes that are not deprecated will be returned. Defaults to `true`.
        """
        return pulumi.get(self, "filter_deprecated")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def probes(self) -> Mapping[str, int]:
        """
        Map of probes with their names as keys and IDs as values.
        """
        return pulumi.get(self, "probes")


class AwaitableGetProbesResult(GetProbesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProbesResult(
            filter_deprecated=self.filter_deprecated,
            id=self.id,
            probes=self.probes)


def get_probes(filter_deprecated: Optional[bool] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProbesResult:
    """
    Data source for retrieving all probes.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    main = grafana.syntheticMonitoring.get_probes()
    ```


    :param bool filter_deprecated: If true, only probes that are not deprecated will be returned. Defaults to `true`.
    """
    __args__ = dict()
    __args__['filterDeprecated'] = filter_deprecated
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:syntheticMonitoring/getProbes:getProbes', __args__, opts=opts, typ=GetProbesResult).value

    return AwaitableGetProbesResult(
        filter_deprecated=pulumi.get(__ret__, 'filter_deprecated'),
        id=pulumi.get(__ret__, 'id'),
        probes=pulumi.get(__ret__, 'probes'))


@_utilities.lift_output_func(get_probes)
def get_probes_output(filter_deprecated: Optional[pulumi.Input[Optional[bool]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProbesResult]:
    """
    Data source for retrieving all probes.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    main = grafana.syntheticMonitoring.get_probes()
    ```


    :param bool filter_deprecated: If true, only probes that are not deprecated will be returned. Defaults to `true`.
    """
    ...