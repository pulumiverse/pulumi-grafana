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
    'GetOncallUserResult',
    'AwaitableGetOncallUserResult',
    'get_oncall_user',
    'get_oncall_user_output',
]

warnings.warn("""grafana.index/getoncalluser.getOncallUser has been deprecated in favor of grafana.oncall/getuser.getUser""", DeprecationWarning)

@pulumi.output_type
class GetOncallUserResult:
    """
    A collection of values returned by getOncallUser.
    """
    def __init__(__self__, email=None, id=None, role=None, username=None):
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        The email of the user.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the user.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        The role of the user.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        The username of the user.
        """
        return pulumi.get(self, "username")


class AwaitableGetOncallUserResult(GetOncallUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOncallUserResult(
            email=self.email,
            id=self.id,
            role=self.role,
            username=self.username)


def get_oncall_user(username: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOncallUserResult:
    """
    * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/users/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    alex = grafana.onCall.get_user(username="alex")
    ```


    :param str username: The username of the user.
    """
    pulumi.log.warn("""get_oncall_user is deprecated: grafana.index/getoncalluser.getOncallUser has been deprecated in favor of grafana.oncall/getuser.getUser""")
    __args__ = dict()
    __args__['username'] = username
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:index/getOncallUser:getOncallUser', __args__, opts=opts, typ=GetOncallUserResult).value

    return AwaitableGetOncallUserResult(
        email=pulumi.get(__ret__, 'email'),
        id=pulumi.get(__ret__, 'id'),
        role=pulumi.get(__ret__, 'role'),
        username=pulumi.get(__ret__, 'username'))
def get_oncall_user_output(username: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOncallUserResult]:
    """
    * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/users/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    alex = grafana.onCall.get_user(username="alex")
    ```


    :param str username: The username of the user.
    """
    pulumi.log.warn("""get_oncall_user is deprecated: grafana.index/getoncalluser.getOncallUser has been deprecated in favor of grafana.oncall/getuser.getUser""")
    __args__ = dict()
    __args__['username'] = username
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('grafana:index/getOncallUser:getOncallUser', __args__, opts=opts, typ=GetOncallUserResult)
    return __ret__.apply(lambda __response__: GetOncallUserResult(
        email=pulumi.get(__response__, 'email'),
        id=pulumi.get(__response__, 'id'),
        role=pulumi.get(__response__, 'role'),
        username=pulumi.get(__response__, 'username')))
