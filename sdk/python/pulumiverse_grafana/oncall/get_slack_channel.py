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

__all__ = [
    'GetSlackChannelResult',
    'AwaitableGetSlackChannelResult',
    'get_slack_channel',
    'get_slack_channel_output',
]

@pulumi.output_type
class GetSlackChannelResult:
    """
    A collection of values returned by getSlackChannel.
    """
    def __init__(__self__, id=None, name=None, slack_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if slack_id and not isinstance(slack_id, str):
            raise TypeError("Expected argument 'slack_id' to be a str")
        pulumi.set(__self__, "slack_id", slack_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The Slack channel name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="slackId")
    def slack_id(self) -> str:
        """
        The Slack ID of the channel.
        """
        return pulumi.get(self, "slack_id")


class AwaitableGetSlackChannelResult(GetSlackChannelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSlackChannelResult(
            id=self.id,
            name=self.name,
            slack_id=self.slack_id)


def get_slack_channel(name: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSlackChannelResult:
    """
    * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/slack_channels/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    example_slack_channel = grafana.onCall.get_slack_channel(name="example_slack_channel")
    ```


    :param str name: The Slack channel name.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:onCall/getSlackChannel:getSlackChannel', __args__, opts=opts, typ=GetSlackChannelResult).value

    return AwaitableGetSlackChannelResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        slack_id=pulumi.get(__ret__, 'slack_id'))
def get_slack_channel_output(name: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSlackChannelResult]:
    """
    * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/slack_channels/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    example_slack_channel = grafana.onCall.get_slack_channel(name="example_slack_channel")
    ```


    :param str name: The Slack channel name.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('grafana:onCall/getSlackChannel:getSlackChannel', __args__, opts=opts, typ=GetSlackChannelResult)
    return __ret__.apply(lambda __response__: GetSlackChannelResult(
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        slack_id=pulumi.get(__response__, 'slack_id')))
