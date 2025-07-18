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
    'GetOrganizationResult',
    'AwaitableGetOrganizationResult',
    'get_organization',
    'get_organization_output',
]

@pulumi.output_type
class GetOrganizationResult:
    """
    A collection of values returned by getOrganization.
    """
    def __init__(__self__, admins=None, editors=None, id=None, name=None, viewers=None):
        if admins and not isinstance(admins, list):
            raise TypeError("Expected argument 'admins' to be a list")
        pulumi.set(__self__, "admins", admins)
        if editors and not isinstance(editors, list):
            raise TypeError("Expected argument 'editors' to be a list")
        pulumi.set(__self__, "editors", editors)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if viewers and not isinstance(viewers, list):
            raise TypeError("Expected argument 'viewers' to be a list")
        pulumi.set(__self__, "viewers", viewers)

    @property
    @pulumi.getter
    def admins(self) -> Sequence[builtins.str]:
        """
        A list of email addresses corresponding to users given admin access to the organization.
        """
        return pulumi.get(self, "admins")

    @property
    @pulumi.getter
    def editors(self) -> Sequence[builtins.str]:
        """
        A list of email addresses corresponding to users given editor access to the organization.
        """
        return pulumi.get(self, "editors")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the Organization.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def viewers(self) -> Sequence[builtins.str]:
        """
        A list of email addresses corresponding to users given viewer access to the organization.
        """
        return pulumi.get(self, "viewers")


class AwaitableGetOrganizationResult(GetOrganizationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationResult(
            admins=self.admins,
            editors=self.editors,
            id=self.id,
            name=self.name,
            viewers=self.viewers)


def get_organization(name: Optional[builtins.str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrganizationResult:
    """
    * [Official documentation](https://grafana.com/docs/grafana/latest/administration/organization-management/)
    * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/org/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana
    import pulumiverse_grafana as grafana

    test = grafana.oss.Organization("test",
        name="test-org",
        admin_user="admin",
        create_users=True,
        viewers=[
            "viewer-01@example.com",
            "viewer-02@example.com",
        ])
    from_name = grafana.oss.get_organization_output(name=test.name)
    ```


    :param builtins.str name: The name of the Organization.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:oss/getOrganization:getOrganization', __args__, opts=opts, typ=GetOrganizationResult).value

    return AwaitableGetOrganizationResult(
        admins=pulumi.get(__ret__, 'admins'),
        editors=pulumi.get(__ret__, 'editors'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        viewers=pulumi.get(__ret__, 'viewers'))
def get_organization_output(name: Optional[pulumi.Input[builtins.str]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetOrganizationResult]:
    """
    * [Official documentation](https://grafana.com/docs/grafana/latest/administration/organization-management/)
    * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/org/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana
    import pulumiverse_grafana as grafana

    test = grafana.oss.Organization("test",
        name="test-org",
        admin_user="admin",
        create_users=True,
        viewers=[
            "viewer-01@example.com",
            "viewer-02@example.com",
        ])
    from_name = grafana.oss.get_organization_output(name=test.name)
    ```


    :param builtins.str name: The name of the Organization.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('grafana:oss/getOrganization:getOrganization', __args__, opts=opts, typ=GetOrganizationResult)
    return __ret__.apply(lambda __response__: GetOrganizationResult(
        admins=pulumi.get(__response__, 'admins'),
        editors=pulumi.get(__response__, 'editors'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        viewers=pulumi.get(__response__, 'viewers')))
