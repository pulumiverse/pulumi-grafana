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

__all__ = ['DataSourcePermissionItemArgs', 'DataSourcePermissionItem']

@pulumi.input_type
class DataSourcePermissionItemArgs:
    def __init__(__self__, *,
                 datasource_uid: pulumi.Input[str],
                 permission: pulumi.Input[str],
                 org_id: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 team: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DataSourcePermissionItem resource.
        :param pulumi.Input[str] datasource_uid: The UID of the datasource.
        :param pulumi.Input[str] permission: the permission to be assigned
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
        :param pulumi.Input[str] role: the role onto which the permission is to be assigned
        :param pulumi.Input[str] team: the team onto which the permission is to be assigned
        :param pulumi.Input[str] user: the user or service account onto which the permission is to be assigned
        """
        pulumi.set(__self__, "datasource_uid", datasource_uid)
        pulumi.set(__self__, "permission", permission)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if team is not None:
            pulumi.set(__self__, "team", team)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter(name="datasourceUid")
    def datasource_uid(self) -> pulumi.Input[str]:
        """
        The UID of the datasource.
        """
        return pulumi.get(self, "datasource_uid")

    @datasource_uid.setter
    def datasource_uid(self, value: pulumi.Input[str]):
        pulumi.set(self, "datasource_uid", value)

    @property
    @pulumi.getter
    def permission(self) -> pulumi.Input[str]:
        """
        the permission to be assigned
        """
        return pulumi.get(self, "permission")

    @permission.setter
    def permission(self, value: pulumi.Input[str]):
        pulumi.set(self, "permission", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        the role onto which the permission is to be assigned
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def team(self) -> Optional[pulumi.Input[str]]:
        """
        the team onto which the permission is to be assigned
        """
        return pulumi.get(self, "team")

    @team.setter
    def team(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        the user or service account onto which the permission is to be assigned
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


@pulumi.input_type
class _DataSourcePermissionItemState:
    def __init__(__self__, *,
                 datasource_uid: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 permission: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 team: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DataSourcePermissionItem resources.
        :param pulumi.Input[str] datasource_uid: The UID of the datasource.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
        :param pulumi.Input[str] permission: the permission to be assigned
        :param pulumi.Input[str] role: the role onto which the permission is to be assigned
        :param pulumi.Input[str] team: the team onto which the permission is to be assigned
        :param pulumi.Input[str] user: the user or service account onto which the permission is to be assigned
        """
        if datasource_uid is not None:
            pulumi.set(__self__, "datasource_uid", datasource_uid)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if permission is not None:
            pulumi.set(__self__, "permission", permission)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if team is not None:
            pulumi.set(__self__, "team", team)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter(name="datasourceUid")
    def datasource_uid(self) -> Optional[pulumi.Input[str]]:
        """
        The UID of the datasource.
        """
        return pulumi.get(self, "datasource_uid")

    @datasource_uid.setter
    def datasource_uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "datasource_uid", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter
    def permission(self) -> Optional[pulumi.Input[str]]:
        """
        the permission to be assigned
        """
        return pulumi.get(self, "permission")

    @permission.setter
    def permission(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "permission", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        the role onto which the permission is to be assigned
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def team(self) -> Optional[pulumi.Input[str]]:
        """
        the team onto which the permission is to be assigned
        """
        return pulumi.get(self, "team")

    @team.setter
    def team(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        the user or service account onto which the permission is to be assigned
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


warnings.warn("""grafana.index/datasourcepermissionitem.DataSourcePermissionItem has been deprecated in favor of grafana.enterprise/datasourcepermissionitem.DataSourcePermissionItem""", DeprecationWarning)


class DataSourcePermissionItem(pulumi.CustomResource):
    warnings.warn("""grafana.index/datasourcepermissionitem.DataSourcePermissionItem has been deprecated in favor of grafana.enterprise/datasourcepermissionitem.DataSourcePermissionItem""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 datasource_uid: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 permission: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 team: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a single permission item for a datasource. Conflicts with the "enterprise.DataSourcePermission" resource which manages the entire set of permissions for a datasource.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumiverse_grafana as grafana

        team = grafana.oss.Team("team", name="Team Name")
        foo = grafana.oss.DataSource("foo",
            type="cloudwatch",
            name="cw-example",
            json_data_encoded=json.dumps({
                "defaultRegion": "us-east-1",
                "authType": "keys",
            }),
            secure_json_data_encoded=json.dumps({
                "accessKey": "123",
                "secretKey": "456",
            }))
        user = grafana.oss.User("user",
            name="test-ds-permissions",
            email="test-ds-permissions@example.com",
            login="test-ds-permissions",
            password="hunter2")
        sa = grafana.oss.ServiceAccount("sa",
            name="test-ds-permissions",
            role="Viewer")
        team_data_source_permission_item = grafana.enterprise.DataSourcePermissionItem("team",
            datasource_uid=foo.uid,
            team=team.id,
            permission="Edit")
        user_data_source_permission_item = grafana.enterprise.DataSourcePermissionItem("user",
            datasource_uid=foo.uid,
            user=user.id,
            permission="Edit")
        role = grafana.enterprise.DataSourcePermissionItem("role",
            datasource_uid=foo.uid,
            role="Viewer",
            permission="Query")
        service_account = grafana.enterprise.DataSourcePermissionItem("service_account",
            datasource_uid=foo.uid,
            user=sa.id,
            permission="Query")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/dataSourcePermissionItem:DataSourcePermissionItem name "{{ datasourceUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
        ```

        ```sh
        $ pulumi import grafana:index/dataSourcePermissionItem:DataSourcePermissionItem name "{{ orgID }}:{{ datasourceUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] datasource_uid: The UID of the datasource.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
        :param pulumi.Input[str] permission: the permission to be assigned
        :param pulumi.Input[str] role: the role onto which the permission is to be assigned
        :param pulumi.Input[str] team: the team onto which the permission is to be assigned
        :param pulumi.Input[str] user: the user or service account onto which the permission is to be assigned
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataSourcePermissionItemArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a single permission item for a datasource. Conflicts with the "enterprise.DataSourcePermission" resource which manages the entire set of permissions for a datasource.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumiverse_grafana as grafana

        team = grafana.oss.Team("team", name="Team Name")
        foo = grafana.oss.DataSource("foo",
            type="cloudwatch",
            name="cw-example",
            json_data_encoded=json.dumps({
                "defaultRegion": "us-east-1",
                "authType": "keys",
            }),
            secure_json_data_encoded=json.dumps({
                "accessKey": "123",
                "secretKey": "456",
            }))
        user = grafana.oss.User("user",
            name="test-ds-permissions",
            email="test-ds-permissions@example.com",
            login="test-ds-permissions",
            password="hunter2")
        sa = grafana.oss.ServiceAccount("sa",
            name="test-ds-permissions",
            role="Viewer")
        team_data_source_permission_item = grafana.enterprise.DataSourcePermissionItem("team",
            datasource_uid=foo.uid,
            team=team.id,
            permission="Edit")
        user_data_source_permission_item = grafana.enterprise.DataSourcePermissionItem("user",
            datasource_uid=foo.uid,
            user=user.id,
            permission="Edit")
        role = grafana.enterprise.DataSourcePermissionItem("role",
            datasource_uid=foo.uid,
            role="Viewer",
            permission="Query")
        service_account = grafana.enterprise.DataSourcePermissionItem("service_account",
            datasource_uid=foo.uid,
            user=sa.id,
            permission="Query")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/dataSourcePermissionItem:DataSourcePermissionItem name "{{ datasourceUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
        ```

        ```sh
        $ pulumi import grafana:index/dataSourcePermissionItem:DataSourcePermissionItem name "{{ orgID }}:{{ datasourceUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
        ```

        :param str resource_name: The name of the resource.
        :param DataSourcePermissionItemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataSourcePermissionItemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 datasource_uid: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 permission: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 team: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""DataSourcePermissionItem is deprecated: grafana.index/datasourcepermissionitem.DataSourcePermissionItem has been deprecated in favor of grafana.enterprise/datasourcepermissionitem.DataSourcePermissionItem""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DataSourcePermissionItemArgs.__new__(DataSourcePermissionItemArgs)

            if datasource_uid is None and not opts.urn:
                raise TypeError("Missing required property 'datasource_uid'")
            __props__.__dict__["datasource_uid"] = datasource_uid
            __props__.__dict__["org_id"] = org_id
            if permission is None and not opts.urn:
                raise TypeError("Missing required property 'permission'")
            __props__.__dict__["permission"] = permission
            __props__.__dict__["role"] = role
            __props__.__dict__["team"] = team
            __props__.__dict__["user"] = user
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:index/dataSourcePermissionItem:DataSourcePermissionItem")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DataSourcePermissionItem, __self__).__init__(
            'grafana:index/dataSourcePermissionItem:DataSourcePermissionItem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            datasource_uid: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            permission: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            team: Optional[pulumi.Input[str]] = None,
            user: Optional[pulumi.Input[str]] = None) -> 'DataSourcePermissionItem':
        """
        Get an existing DataSourcePermissionItem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] datasource_uid: The UID of the datasource.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
        :param pulumi.Input[str] permission: the permission to be assigned
        :param pulumi.Input[str] role: the role onto which the permission is to be assigned
        :param pulumi.Input[str] team: the team onto which the permission is to be assigned
        :param pulumi.Input[str] user: the user or service account onto which the permission is to be assigned
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DataSourcePermissionItemState.__new__(_DataSourcePermissionItemState)

        __props__.__dict__["datasource_uid"] = datasource_uid
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["permission"] = permission
        __props__.__dict__["role"] = role
        __props__.__dict__["team"] = team
        __props__.__dict__["user"] = user
        return DataSourcePermissionItem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="datasourceUid")
    def datasource_uid(self) -> pulumi.Output[str]:
        """
        The UID of the datasource.
        """
        return pulumi.get(self, "datasource_uid")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[str]:
        """
        The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def permission(self) -> pulumi.Output[str]:
        """
        the permission to be assigned
        """
        return pulumi.get(self, "permission")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[Optional[str]]:
        """
        the role onto which the permission is to be assigned
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def team(self) -> pulumi.Output[Optional[str]]:
        """
        the team onto which the permission is to be assigned
        """
        return pulumi.get(self, "team")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[Optional[str]]:
        """
        the user or service account onto which the permission is to be assigned
        """
        return pulumi.get(self, "user")

