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
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    """
    A collection of values returned by getUser.
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
    def email(self) -> builtins.str:
        """
        The email of the user.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The ID of the user.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def role(self) -> builtins.str:
        """
        The role of the user.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def username(self) -> builtins.str:
        """
        The username of the user.
        """
        return pulumi.get(self, "username")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            email=self.email,
            id=self.id,
            role=self.role,
            username=self.username)


def get_user(username: Optional[builtins.str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/users/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    alex = grafana.onCall.get_user(username="alex")
    ```


    :param builtins.str username: The username of the user.
    """
    __args__ = dict()
    __args__['username'] = username
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:onCall/getUser:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        email=pulumi.get(__ret__, 'email'),
        id=pulumi.get(__ret__, 'id'),
        role=pulumi.get(__ret__, 'role'),
        username=pulumi.get(__ret__, 'username'))
def get_user_output(username: Optional[pulumi.Input[builtins.str]] = None,
                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetUserResult]:
    """
    * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/users/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    alex = grafana.onCall.get_user(username="alex")
    ```


    :param builtins.str username: The username of the user.
    """
    __args__ = dict()
    __args__['username'] = username
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('grafana:onCall/getUser:getUser', __args__, opts=opts, typ=GetUserResult)
    return __ret__.apply(lambda __response__: GetUserResult(
        email=pulumi.get(__response__, 'email'),
        id=pulumi.get(__response__, 'id'),
        role=pulumi.get(__response__, 'role'),
        username=pulumi.get(__response__, 'username')))
