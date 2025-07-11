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

__all__ = [
    'GetIntegrationResult',
    'AwaitableGetIntegrationResult',
    'get_integration',
    'get_integration_output',
]

@pulumi.output_type
class GetIntegrationResult:
    """
    A collection of values returned by getIntegration.
    """
    def __init__(__self__, id=None, link=None, name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if link and not isinstance(link, str):
            raise TypeError("Expected argument 'link' to be a str")
        pulumi.set(__self__, "link", link)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The integration ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def link(self) -> builtins.str:
        """
        The link for the integration.
        """
        return pulumi.get(self, "link")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The integration name.
        """
        return pulumi.get(self, "name")


class AwaitableGetIntegrationResult(GetIntegrationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIntegrationResult(
            id=self.id,
            link=self.link,
            name=self.name)


def get_integration(id: Optional[builtins.str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIntegrationResult:
    """
    * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/integrations/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    example_integration = grafana.onCall.get_integration(id="CEXAMPLEID123")
    ```


    :param builtins.str id: The integration ID.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:onCall/getIntegration:getIntegration', __args__, opts=opts, typ=GetIntegrationResult).value

    return AwaitableGetIntegrationResult(
        id=pulumi.get(__ret__, 'id'),
        link=pulumi.get(__ret__, 'link'),
        name=pulumi.get(__ret__, 'name'))
def get_integration_output(id: Optional[pulumi.Input[builtins.str]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetIntegrationResult]:
    """
    * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/integrations/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    example_integration = grafana.onCall.get_integration(id="CEXAMPLEID123")
    ```


    :param builtins.str id: The integration ID.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('grafana:onCall/getIntegration:getIntegration', __args__, opts=opts, typ=GetIntegrationResult)
    return __ret__.apply(lambda __response__: GetIntegrationResult(
        id=pulumi.get(__response__, 'id'),
        link=pulumi.get(__response__, 'link'),
        name=pulumi.get(__response__, 'name')))
