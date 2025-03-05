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

__all__ = [
    'GetLibraryPanelsResult',
    'AwaitableGetLibraryPanelsResult',
    'get_library_panels',
    'get_library_panels_output',
]

@pulumi.output_type
class GetLibraryPanelsResult:
    """
    A collection of values returned by getLibraryPanels.
    """
    def __init__(__self__, id=None, org_id=None, panels=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if panels and not isinstance(panels, list):
            raise TypeError("Expected argument 'panels' to be a list")
        pulumi.set(__self__, "panels", panels)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> str:
        """
        The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def panels(self) -> Sequence['outputs.GetLibraryPanelsPanelResult']:
        return pulumi.get(self, "panels")


class AwaitableGetLibraryPanelsResult(GetLibraryPanelsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLibraryPanelsResult(
            id=self.id,
            org_id=self.org_id,
            panels=self.panels)


def get_library_panels(org_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLibraryPanelsResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import json
    import pulumi_grafana as grafana
    import pulumiverse_grafana as grafana

    test = grafana.oss.LibraryPanel("test",
        name="panelname",
        model_json=json.dumps({
            "title": "test name",
            "type": "text",
            "version": 0,
            "description": "test description",
        }))
    test_folder = grafana.oss.Folder("test",
        title="Panel Folder",
        uid="panelname-folder")
    folder = grafana.oss.LibraryPanel("folder",
        name="panelname In Folder",
        folder_uid=test_folder.uid,
        model_json=json.dumps({
            "gridPos": {
                "x": 0,
                "y": 0,
                "h": 10,
                "w": 10,
            },
            "title": "panel",
            "type": "text",
            "version": 0,
        }))
    all = grafana.oss.get_library_panels()
    ```


    :param str org_id: The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
    """
    __args__ = dict()
    __args__['orgId'] = org_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:oss/getLibraryPanels:getLibraryPanels', __args__, opts=opts, typ=GetLibraryPanelsResult).value

    return AwaitableGetLibraryPanelsResult(
        id=pulumi.get(__ret__, 'id'),
        org_id=pulumi.get(__ret__, 'org_id'),
        panels=pulumi.get(__ret__, 'panels'))
def get_library_panels_output(org_id: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetLibraryPanelsResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import json
    import pulumi_grafana as grafana
    import pulumiverse_grafana as grafana

    test = grafana.oss.LibraryPanel("test",
        name="panelname",
        model_json=json.dumps({
            "title": "test name",
            "type": "text",
            "version": 0,
            "description": "test description",
        }))
    test_folder = grafana.oss.Folder("test",
        title="Panel Folder",
        uid="panelname-folder")
    folder = grafana.oss.LibraryPanel("folder",
        name="panelname In Folder",
        folder_uid=test_folder.uid,
        model_json=json.dumps({
            "gridPos": {
                "x": 0,
                "y": 0,
                "h": 10,
                "w": 10,
            },
            "title": "panel",
            "type": "text",
            "version": 0,
        }))
    all = grafana.oss.get_library_panels()
    ```


    :param str org_id: The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
    """
    __args__ = dict()
    __args__['orgId'] = org_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('grafana:oss/getLibraryPanels:getLibraryPanels', __args__, opts=opts, typ=GetLibraryPanelsResult)
    return __ret__.apply(lambda __response__: GetLibraryPanelsResult(
        id=pulumi.get(__response__, 'id'),
        org_id=pulumi.get(__response__, 'org_id'),
        panels=pulumi.get(__response__, 'panels')))
