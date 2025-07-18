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
    'GetTeamResult',
    'AwaitableGetTeamResult',
    'get_team',
    'get_team_output',
]

@pulumi.output_type
class GetTeamResult:
    """
    A collection of values returned by getTeam.
    """
    def __init__(__self__, email=None, id=None, members=None, name=None, org_id=None, preferences=None, read_team_sync=None, team_id=None, team_syncs=None, team_uid=None):
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if members and not isinstance(members, list):
            raise TypeError("Expected argument 'members' to be a list")
        pulumi.set(__self__, "members", members)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if preferences and not isinstance(preferences, list):
            raise TypeError("Expected argument 'preferences' to be a list")
        pulumi.set(__self__, "preferences", preferences)
        if read_team_sync and not isinstance(read_team_sync, bool):
            raise TypeError("Expected argument 'read_team_sync' to be a bool")
        pulumi.set(__self__, "read_team_sync", read_team_sync)
        if team_id and not isinstance(team_id, int):
            raise TypeError("Expected argument 'team_id' to be a int")
        pulumi.set(__self__, "team_id", team_id)
        if team_syncs and not isinstance(team_syncs, list):
            raise TypeError("Expected argument 'team_syncs' to be a list")
        pulumi.set(__self__, "team_syncs", team_syncs)
        if team_uid and not isinstance(team_uid, str):
            raise TypeError("Expected argument 'team_uid' to be a str")
        pulumi.set(__self__, "team_uid", team_uid)

    @property
    @pulumi.getter
    def email(self) -> builtins.str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def members(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def preferences(self) -> Sequence['outputs.GetTeamPreferenceResult']:
        return pulumi.get(self, "preferences")

    @property
    @pulumi.getter(name="readTeamSync")
    def read_team_sync(self) -> Optional[builtins.bool]:
        return pulumi.get(self, "read_team_sync")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> builtins.int:
        return pulumi.get(self, "team_id")

    @property
    @pulumi.getter(name="teamSyncs")
    def team_syncs(self) -> Sequence['outputs.GetTeamTeamSyncResult']:
        return pulumi.get(self, "team_syncs")

    @property
    @pulumi.getter(name="teamUid")
    def team_uid(self) -> builtins.str:
        return pulumi.get(self, "team_uid")


class AwaitableGetTeamResult(GetTeamResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTeamResult(
            email=self.email,
            id=self.id,
            members=self.members,
            name=self.name,
            org_id=self.org_id,
            preferences=self.preferences,
            read_team_sync=self.read_team_sync,
            team_id=self.team_id,
            team_syncs=self.team_syncs,
            team_uid=self.team_uid)


def get_team(name: Optional[builtins.str] = None,
             org_id: Optional[builtins.str] = None,
             read_team_sync: Optional[builtins.bool] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTeamResult:
    """
    * [Official documentation](https://grafana.com/docs/grafana/latest/administration/team-management/)
    * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/team/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana
    import pulumiverse_grafana as grafana

    test = grafana.oss.Team("test",
        name="test-team",
        email="test-team-email@test.com",
        preferences={
            "theme": "dark",
            "timezone": "utc",
        })
    from_name = grafana.oss.get_team_output(name=test.name)
    ```
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['orgId'] = org_id
    __args__['readTeamSync'] = read_team_sync
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:oss/getTeam:getTeam', __args__, opts=opts, typ=GetTeamResult).value

    return AwaitableGetTeamResult(
        email=pulumi.get(__ret__, 'email'),
        id=pulumi.get(__ret__, 'id'),
        members=pulumi.get(__ret__, 'members'),
        name=pulumi.get(__ret__, 'name'),
        org_id=pulumi.get(__ret__, 'org_id'),
        preferences=pulumi.get(__ret__, 'preferences'),
        read_team_sync=pulumi.get(__ret__, 'read_team_sync'),
        team_id=pulumi.get(__ret__, 'team_id'),
        team_syncs=pulumi.get(__ret__, 'team_syncs'),
        team_uid=pulumi.get(__ret__, 'team_uid'))
def get_team_output(name: Optional[pulumi.Input[builtins.str]] = None,
                    org_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                    read_team_sync: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetTeamResult]:
    """
    * [Official documentation](https://grafana.com/docs/grafana/latest/administration/team-management/)
    * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/team/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana
    import pulumiverse_grafana as grafana

    test = grafana.oss.Team("test",
        name="test-team",
        email="test-team-email@test.com",
        preferences={
            "theme": "dark",
            "timezone": "utc",
        })
    from_name = grafana.oss.get_team_output(name=test.name)
    ```
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['orgId'] = org_id
    __args__['readTeamSync'] = read_team_sync
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('grafana:oss/getTeam:getTeam', __args__, opts=opts, typ=GetTeamResult)
    return __ret__.apply(lambda __response__: GetTeamResult(
        email=pulumi.get(__response__, 'email'),
        id=pulumi.get(__response__, 'id'),
        members=pulumi.get(__response__, 'members'),
        name=pulumi.get(__response__, 'name'),
        org_id=pulumi.get(__response__, 'org_id'),
        preferences=pulumi.get(__response__, 'preferences'),
        read_team_sync=pulumi.get(__response__, 'read_team_sync'),
        team_id=pulumi.get(__response__, 'team_id'),
        team_syncs=pulumi.get(__response__, 'team_syncs'),
        team_uid=pulumi.get(__response__, 'team_uid')))
