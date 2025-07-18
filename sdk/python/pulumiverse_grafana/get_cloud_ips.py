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

__all__ = [
    'GetCloudIpsResult',
    'AwaitableGetCloudIpsResult',
    'get_cloud_ips',
    'get_cloud_ips_output',
]

warnings.warn("""grafana.index/getcloudips.getCloudIps has been deprecated in favor of grafana.cloud/getips.getIps""", DeprecationWarning)

@pulumi.output_type
class GetCloudIpsResult:
    """
    A collection of values returned by getCloudIps.
    """
    def __init__(__self__, hosted_alerts=None, hosted_grafanas=None, hosted_logs=None, hosted_metrics=None, hosted_traces=None, id=None):
        if hosted_alerts and not isinstance(hosted_alerts, list):
            raise TypeError("Expected argument 'hosted_alerts' to be a list")
        pulumi.set(__self__, "hosted_alerts", hosted_alerts)
        if hosted_grafanas and not isinstance(hosted_grafanas, list):
            raise TypeError("Expected argument 'hosted_grafanas' to be a list")
        pulumi.set(__self__, "hosted_grafanas", hosted_grafanas)
        if hosted_logs and not isinstance(hosted_logs, list):
            raise TypeError("Expected argument 'hosted_logs' to be a list")
        pulumi.set(__self__, "hosted_logs", hosted_logs)
        if hosted_metrics and not isinstance(hosted_metrics, list):
            raise TypeError("Expected argument 'hosted_metrics' to be a list")
        pulumi.set(__self__, "hosted_metrics", hosted_metrics)
        if hosted_traces and not isinstance(hosted_traces, list):
            raise TypeError("Expected argument 'hosted_traces' to be a list")
        pulumi.set(__self__, "hosted_traces", hosted_traces)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="hostedAlerts")
    def hosted_alerts(self) -> Sequence[builtins.str]:
        """
        Set of IP addresses that are used for hosted alerts.
        """
        return pulumi.get(self, "hosted_alerts")

    @property
    @pulumi.getter(name="hostedGrafanas")
    def hosted_grafanas(self) -> Sequence[builtins.str]:
        """
        Set of IP addresses that are used for hosted Grafana.
        """
        return pulumi.get(self, "hosted_grafanas")

    @property
    @pulumi.getter(name="hostedLogs")
    def hosted_logs(self) -> Sequence[builtins.str]:
        """
        Set of IP addresses that are used for hosted logs.
        """
        return pulumi.get(self, "hosted_logs")

    @property
    @pulumi.getter(name="hostedMetrics")
    def hosted_metrics(self) -> Sequence[builtins.str]:
        """
        Set of IP addresses that are used for hosted metrics.
        """
        return pulumi.get(self, "hosted_metrics")

    @property
    @pulumi.getter(name="hostedTraces")
    def hosted_traces(self) -> Sequence[builtins.str]:
        """
        Set of IP addresses that are used for hosted traces.
        """
        return pulumi.get(self, "hosted_traces")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetCloudIpsResult(GetCloudIpsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudIpsResult(
            hosted_alerts=self.hosted_alerts,
            hosted_grafanas=self.hosted_grafanas,
            hosted_logs=self.hosted_logs,
            hosted_metrics=self.hosted_metrics,
            hosted_traces=self.hosted_traces,
            id=self.id)


def get_cloud_ips(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudIpsResult:
    """
    Data source for retrieving sets of cloud IPs. See https://grafana.com/docs/grafana-cloud/reference/allow-list/ for more info

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    test = grafana.cloud.get_ips()
    ```
    """
    pulumi.log.warn("""get_cloud_ips is deprecated: grafana.index/getcloudips.getCloudIps has been deprecated in favor of grafana.cloud/getips.getIps""")
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:index/getCloudIps:getCloudIps', __args__, opts=opts, typ=GetCloudIpsResult).value

    return AwaitableGetCloudIpsResult(
        hosted_alerts=pulumi.get(__ret__, 'hosted_alerts'),
        hosted_grafanas=pulumi.get(__ret__, 'hosted_grafanas'),
        hosted_logs=pulumi.get(__ret__, 'hosted_logs'),
        hosted_metrics=pulumi.get(__ret__, 'hosted_metrics'),
        hosted_traces=pulumi.get(__ret__, 'hosted_traces'),
        id=pulumi.get(__ret__, 'id'))
def get_cloud_ips_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCloudIpsResult]:
    """
    Data source for retrieving sets of cloud IPs. See https://grafana.com/docs/grafana-cloud/reference/allow-list/ for more info

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    test = grafana.cloud.get_ips()
    ```
    """
    pulumi.log.warn("""get_cloud_ips is deprecated: grafana.index/getcloudips.getCloudIps has been deprecated in favor of grafana.cloud/getips.getIps""")
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('grafana:index/getCloudIps:getCloudIps', __args__, opts=opts, typ=GetCloudIpsResult)
    return __ret__.apply(lambda __response__: GetCloudIpsResult(
        hosted_alerts=pulumi.get(__response__, 'hosted_alerts'),
        hosted_grafanas=pulumi.get(__response__, 'hosted_grafanas'),
        hosted_logs=pulumi.get(__response__, 'hosted_logs'),
        hosted_metrics=pulumi.get(__response__, 'hosted_metrics'),
        hosted_traces=pulumi.get(__response__, 'hosted_traces'),
        id=pulumi.get(__response__, 'id')))
