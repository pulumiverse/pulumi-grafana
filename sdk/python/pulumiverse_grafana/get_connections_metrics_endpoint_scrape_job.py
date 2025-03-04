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
from . import _utilities

__all__ = [
    'GetConnectionsMetricsEndpointScrapeJobResult',
    'AwaitableGetConnectionsMetricsEndpointScrapeJobResult',
    'get_connections_metrics_endpoint_scrape_job',
    'get_connections_metrics_endpoint_scrape_job_output',
]

warnings.warn("""grafana.index/getconnectionsmetricsendpointscrapejob.getConnectionsMetricsEndpointScrapeJob has been deprecated in favor of grafana.connections/getmetricsendpointscrapejob.getMetricsEndpointScrapeJob""", DeprecationWarning)

@pulumi.output_type
class GetConnectionsMetricsEndpointScrapeJobResult:
    """
    A collection of values returned by getConnectionsMetricsEndpointScrapeJob.
    """
    def __init__(__self__, authentication_basic_password=None, authentication_basic_username=None, authentication_bearer_token=None, authentication_method=None, enabled=None, id=None, name=None, scrape_interval_seconds=None, stack_id=None, url=None):
        if authentication_basic_password and not isinstance(authentication_basic_password, str):
            raise TypeError("Expected argument 'authentication_basic_password' to be a str")
        pulumi.set(__self__, "authentication_basic_password", authentication_basic_password)
        if authentication_basic_username and not isinstance(authentication_basic_username, str):
            raise TypeError("Expected argument 'authentication_basic_username' to be a str")
        pulumi.set(__self__, "authentication_basic_username", authentication_basic_username)
        if authentication_bearer_token and not isinstance(authentication_bearer_token, str):
            raise TypeError("Expected argument 'authentication_bearer_token' to be a str")
        pulumi.set(__self__, "authentication_bearer_token", authentication_bearer_token)
        if authentication_method and not isinstance(authentication_method, str):
            raise TypeError("Expected argument 'authentication_method' to be a str")
        pulumi.set(__self__, "authentication_method", authentication_method)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if scrape_interval_seconds and not isinstance(scrape_interval_seconds, int):
            raise TypeError("Expected argument 'scrape_interval_seconds' to be a int")
        pulumi.set(__self__, "scrape_interval_seconds", scrape_interval_seconds)
        if stack_id and not isinstance(stack_id, str):
            raise TypeError("Expected argument 'stack_id' to be a str")
        pulumi.set(__self__, "stack_id", stack_id)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="authenticationBasicPassword")
    def authentication_basic_password(self) -> str:
        """
        Password for basic authentication.
        """
        return pulumi.get(self, "authentication_basic_password")

    @property
    @pulumi.getter(name="authenticationBasicUsername")
    def authentication_basic_username(self) -> str:
        """
        Username for basic authentication.
        """
        return pulumi.get(self, "authentication_basic_username")

    @property
    @pulumi.getter(name="authenticationBearerToken")
    def authentication_bearer_token(self) -> str:
        """
        Token for authentication bearer.
        """
        return pulumi.get(self, "authentication_bearer_token")

    @property
    @pulumi.getter(name="authenticationMethod")
    def authentication_method(self) -> str:
        """
        Method to pass authentication credentials: basic or bearer.
        """
        return pulumi.get(self, "authentication_method")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Whether the metrics endpoint scrape job is enabled or not.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="scrapeIntervalSeconds")
    def scrape_interval_seconds(self) -> int:
        """
        Frequency for scraping the metrics endpoint: 30, 60, or 120 seconds.
        """
        return pulumi.get(self, "scrape_interval_seconds")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> str:
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        The url to scrape metrics.
        """
        return pulumi.get(self, "url")


class AwaitableGetConnectionsMetricsEndpointScrapeJobResult(GetConnectionsMetricsEndpointScrapeJobResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConnectionsMetricsEndpointScrapeJobResult(
            authentication_basic_password=self.authentication_basic_password,
            authentication_basic_username=self.authentication_basic_username,
            authentication_bearer_token=self.authentication_bearer_token,
            authentication_method=self.authentication_method,
            enabled=self.enabled,
            id=self.id,
            name=self.name,
            scrape_interval_seconds=self.scrape_interval_seconds,
            stack_id=self.stack_id,
            url=self.url)


def get_connections_metrics_endpoint_scrape_job(name: Optional[str] = None,
                                                stack_id: Optional[str] = None,
                                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConnectionsMetricsEndpointScrapeJobResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    ds_test = grafana.connections.get_metrics_endpoint_scrape_job(stack_id="1",
        name="my-scrape-job")
    ```
    """
    pulumi.log.warn("""get_connections_metrics_endpoint_scrape_job is deprecated: grafana.index/getconnectionsmetricsendpointscrapejob.getConnectionsMetricsEndpointScrapeJob has been deprecated in favor of grafana.connections/getmetricsendpointscrapejob.getMetricsEndpointScrapeJob""")
    __args__ = dict()
    __args__['name'] = name
    __args__['stackId'] = stack_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:index/getConnectionsMetricsEndpointScrapeJob:getConnectionsMetricsEndpointScrapeJob', __args__, opts=opts, typ=GetConnectionsMetricsEndpointScrapeJobResult).value

    return AwaitableGetConnectionsMetricsEndpointScrapeJobResult(
        authentication_basic_password=pulumi.get(__ret__, 'authentication_basic_password'),
        authentication_basic_username=pulumi.get(__ret__, 'authentication_basic_username'),
        authentication_bearer_token=pulumi.get(__ret__, 'authentication_bearer_token'),
        authentication_method=pulumi.get(__ret__, 'authentication_method'),
        enabled=pulumi.get(__ret__, 'enabled'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        scrape_interval_seconds=pulumi.get(__ret__, 'scrape_interval_seconds'),
        stack_id=pulumi.get(__ret__, 'stack_id'),
        url=pulumi.get(__ret__, 'url'))
def get_connections_metrics_endpoint_scrape_job_output(name: Optional[pulumi.Input[str]] = None,
                                                       stack_id: Optional[pulumi.Input[str]] = None,
                                                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetConnectionsMetricsEndpointScrapeJobResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    ds_test = grafana.connections.get_metrics_endpoint_scrape_job(stack_id="1",
        name="my-scrape-job")
    ```
    """
    pulumi.log.warn("""get_connections_metrics_endpoint_scrape_job is deprecated: grafana.index/getconnectionsmetricsendpointscrapejob.getConnectionsMetricsEndpointScrapeJob has been deprecated in favor of grafana.connections/getmetricsendpointscrapejob.getMetricsEndpointScrapeJob""")
    __args__ = dict()
    __args__['name'] = name
    __args__['stackId'] = stack_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('grafana:index/getConnectionsMetricsEndpointScrapeJob:getConnectionsMetricsEndpointScrapeJob', __args__, opts=opts, typ=GetConnectionsMetricsEndpointScrapeJobResult)
    return __ret__.apply(lambda __response__: GetConnectionsMetricsEndpointScrapeJobResult(
        authentication_basic_password=pulumi.get(__response__, 'authentication_basic_password'),
        authentication_basic_username=pulumi.get(__response__, 'authentication_basic_username'),
        authentication_bearer_token=pulumi.get(__response__, 'authentication_bearer_token'),
        authentication_method=pulumi.get(__response__, 'authentication_method'),
        enabled=pulumi.get(__response__, 'enabled'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        scrape_interval_seconds=pulumi.get(__response__, 'scrape_interval_seconds'),
        stack_id=pulumi.get(__response__, 'stack_id'),
        url=pulumi.get(__response__, 'url')))
