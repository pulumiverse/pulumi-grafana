# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetRoleResult',
    'AwaitableGetRoleResult',
    'get_role',
    'get_role_output',
]

@pulumi.output_type
class GetRoleResult:
    """
    A collection of values returned by getRole.
    """
    def __init__(__self__, description=None, display_name=None, global_=None, group=None, hidden=None, id=None, name=None, org_id=None, permissions=None, uid=None, version=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if global_ and not isinstance(global_, bool):
            raise TypeError("Expected argument 'global_' to be a bool")
        pulumi.set(__self__, "global_", global_)
        if group and not isinstance(group, str):
            raise TypeError("Expected argument 'group' to be a str")
        pulumi.set(__self__, "group", group)
        if hidden and not isinstance(hidden, bool):
            raise TypeError("Expected argument 'hidden' to be a bool")
        pulumi.set(__self__, "hidden", hidden)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if permissions and not isinstance(permissions, list):
            raise TypeError("Expected argument 'permissions' to be a list")
        pulumi.set(__self__, "permissions", permissions)
        if uid and not isinstance(uid, str):
            raise TypeError("Expected argument 'uid' to be a str")
        pulumi.set(__self__, "uid", uid)
        if version and not isinstance(version, int):
            raise TypeError("Expected argument 'version' to be a int")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the role.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Display name of the role. Available with Grafana 8.5+.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="global")
    def global_(self) -> bool:
        """
        Boolean to state whether the role is available across all organizations or not.
        """
        return pulumi.get(self, "global_")

    @property
    @pulumi.getter
    def group(self) -> str:
        """
        Group of the role. Available with Grafana 8.5+.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def hidden(self) -> bool:
        """
        Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+.
        """
        return pulumi.get(self, "hidden")

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
        Name of the role
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> str:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def permissions(self) -> Sequence['outputs.GetRolePermissionResult']:
        """
        Specific set of actions granted by the role.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        Unique identifier of the role. Used for assignments.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter
    def version(self) -> int:
        """
        Version of the role. A role is updated only on version increase. This field or `auto_increment_version` should be set.
        """
        return pulumi.get(self, "version")


class AwaitableGetRoleResult(GetRoleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRoleResult(
            description=self.description,
            display_name=self.display_name,
            global_=self.global_,
            group=self.group,
            hidden=self.hidden,
            id=self.id,
            name=self.name,
            org_id=self.org_id,
            permissions=self.permissions,
            uid=self.uid,
            version=self.version)


def get_role(name: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRoleResult:
    """
    **Note:** This resource is available only with Grafana Enterprise 8.+.

    * [Official documentation](https://grafana.com/docs/grafana/latest/administration/roles-and-permissions/access-control/)
    * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/access_control/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana
    import pulumiverse_grafana as grafana

    test = grafana.enterprise.Role("test",
        name="test-role",
        description="test-role description",
        uid="test-ds-role-uid",
        version=1,
        global_=True,
        hidden=False,
        permissions=[
            grafana.enterprise.RolePermissionArgs(
                action="org.users:add",
                scope="users:*",
            ),
            grafana.enterprise.RolePermissionArgs(
                action="org.users:write",
                scope="users:*",
            ),
            grafana.enterprise.RolePermissionArgs(
                action="org.users:read",
                scope="users:*",
            ),
        ])
    from_name = grafana.enterprise.get_role_output(name=test.name)
    ```


    :param str name: Name of the role
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:enterprise/getRole:getRole', __args__, opts=opts, typ=GetRoleResult).value

    return AwaitableGetRoleResult(
        description=pulumi.get(__ret__, 'description'),
        display_name=pulumi.get(__ret__, 'display_name'),
        global_=pulumi.get(__ret__, 'global_'),
        group=pulumi.get(__ret__, 'group'),
        hidden=pulumi.get(__ret__, 'hidden'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        org_id=pulumi.get(__ret__, 'org_id'),
        permissions=pulumi.get(__ret__, 'permissions'),
        uid=pulumi.get(__ret__, 'uid'),
        version=pulumi.get(__ret__, 'version'))


@_utilities.lift_output_func(get_role)
def get_role_output(name: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRoleResult]:
    """
    **Note:** This resource is available only with Grafana Enterprise 8.+.

    * [Official documentation](https://grafana.com/docs/grafana/latest/administration/roles-and-permissions/access-control/)
    * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/access_control/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana
    import pulumiverse_grafana as grafana

    test = grafana.enterprise.Role("test",
        name="test-role",
        description="test-role description",
        uid="test-ds-role-uid",
        version=1,
        global_=True,
        hidden=False,
        permissions=[
            grafana.enterprise.RolePermissionArgs(
                action="org.users:add",
                scope="users:*",
            ),
            grafana.enterprise.RolePermissionArgs(
                action="org.users:write",
                scope="users:*",
            ),
            grafana.enterprise.RolePermissionArgs(
                action="org.users:read",
                scope="users:*",
            ),
        ])
    from_name = grafana.enterprise.get_role_output(name=test.name)
    ```


    :param str name: Name of the role
    """
    ...
