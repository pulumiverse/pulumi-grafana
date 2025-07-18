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

__all__ = [
    'GetCollectorsResult',
    'AwaitableGetCollectorsResult',
    'get_collectors',
    'get_collectors_output',
]

@pulumi.output_type
class GetCollectorsResult:
    """
    A collection of values returned by getCollectors.
    """
    def __init__(__self__, collectors=None, id=None):
        if collectors and not isinstance(collectors, list):
            raise TypeError("Expected argument 'collectors' to be a list")
        pulumi.set(__self__, "collectors", collectors)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def collectors(self) -> Sequence['outputs.GetCollectorsCollectorResult']:
        """
        List of collectors
        """
        return pulumi.get(self, "collectors")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetCollectorsResult(GetCollectorsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCollectorsResult(
            collectors=self.collectors,
            id=self.id)


def get_collectors(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCollectorsResult:
    """
    Represents a list of Grafana Fleet Management collectors.

    * [Official documentation](https://grafana.com/docs/grafana-cloud/send-data/fleet-management/)
    * [API documentation](https://grafana.com/docs/grafana-cloud/send-data/fleet-management/api-reference/collector-api/)

    Required access policy scopes:

    * fleet-management:read

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    test = grafana.fleetManagement.get_collectors()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:fleetManagement/getCollectors:getCollectors', __args__, opts=opts, typ=GetCollectorsResult).value

    return AwaitableGetCollectorsResult(
        collectors=pulumi.get(__ret__, 'collectors'),
        id=pulumi.get(__ret__, 'id'))
def get_collectors_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCollectorsResult]:
    """
    Represents a list of Grafana Fleet Management collectors.

    * [Official documentation](https://grafana.com/docs/grafana-cloud/send-data/fleet-management/)
    * [API documentation](https://grafana.com/docs/grafana-cloud/send-data/fleet-management/api-reference/collector-api/)

    Required access policy scopes:

    * fleet-management:read

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    test = grafana.fleetManagement.get_collectors()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('grafana:fleetManagement/getCollectors:getCollectors', __args__, opts=opts, typ=GetCollectorsResult)
    return __ret__.apply(lambda __response__: GetCollectorsResult(
        collectors=pulumi.get(__response__, 'collectors'),
        id=pulumi.get(__response__, 'id')))
