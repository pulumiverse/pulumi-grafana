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

__all__ = ['FolderPermissionItemArgs', 'FolderPermissionItem']

@pulumi.input_type
class FolderPermissionItemArgs:
    def __init__(__self__, *,
                 folder_uid: pulumi.Input[str],
                 permission: pulumi.Input[str],
                 org_id: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 team: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FolderPermissionItem resource.
        :param pulumi.Input[str] folder_uid: The UID of the folder.
        :param pulumi.Input[str] permission: the permission to be assigned
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
        :param pulumi.Input[str] role: the role onto which the permission is to be assigned
        :param pulumi.Input[str] team: the team onto which the permission is to be assigned
        :param pulumi.Input[str] user: the user or service account onto which the permission is to be assigned
        """
        pulumi.set(__self__, "folder_uid", folder_uid)
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
    @pulumi.getter(name="folderUid")
    def folder_uid(self) -> pulumi.Input[str]:
        """
        The UID of the folder.
        """
        return pulumi.get(self, "folder_uid")

    @folder_uid.setter
    def folder_uid(self, value: pulumi.Input[str]):
        pulumi.set(self, "folder_uid", value)

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
class _FolderPermissionItemState:
    def __init__(__self__, *,
                 folder_uid: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 permission: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 team: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FolderPermissionItem resources.
        :param pulumi.Input[str] folder_uid: The UID of the folder.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
        :param pulumi.Input[str] permission: the permission to be assigned
        :param pulumi.Input[str] role: the role onto which the permission is to be assigned
        :param pulumi.Input[str] team: the team onto which the permission is to be assigned
        :param pulumi.Input[str] user: the user or service account onto which the permission is to be assigned
        """
        if folder_uid is not None:
            pulumi.set(__self__, "folder_uid", folder_uid)
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
    @pulumi.getter(name="folderUid")
    def folder_uid(self) -> Optional[pulumi.Input[str]]:
        """
        The UID of the folder.
        """
        return pulumi.get(self, "folder_uid")

    @folder_uid.setter
    def folder_uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_uid", value)

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


warnings.warn("""grafana.index/folderpermissionitem.FolderPermissionItem has been deprecated in favor of grafana.oss/folderpermissionitem.FolderPermissionItem""", DeprecationWarning)


class FolderPermissionItem(pulumi.CustomResource):
    warnings.warn("""grafana.index/folderpermissionitem.FolderPermissionItem has been deprecated in favor of grafana.oss/folderpermissionitem.FolderPermissionItem""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 folder_uid: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 permission: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 team: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a single permission item for a folder. Conflicts with the "oss.FolderPermission" resource which manages the entire set of permissions for a folder.
        		* [Official documentation](https://grafana.com/docs/grafana/latest/administration/roles-and-permissions/access-control/)
        		* [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/folder_permissions/)

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        team = grafana.oss.Team("team", name="Team Name")
        user = grafana.oss.User("user",
            email="user.name@example.com",
            login="user.name",
            password="my-password")
        collection = grafana.oss.Folder("collection", title="Folder Title")
        on_role = grafana.oss.FolderPermissionItem("on_role",
            folder_uid=collection.uid,
            role="Viewer",
            permission="Edit")
        on_team = grafana.oss.FolderPermissionItem("on_team",
            folder_uid=collection.uid,
            team=team.id,
            permission="View")
        on_user = grafana.oss.FolderPermissionItem("on_user",
            folder_uid=collection.uid,
            user=user.id,
            permission="Admin")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/folderPermissionItem:FolderPermissionItem name "{{ folderUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
        ```

        ```sh
        $ pulumi import grafana:index/folderPermissionItem:FolderPermissionItem name "{{ orgID }}:{{ folderUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] folder_uid: The UID of the folder.
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
                 args: FolderPermissionItemArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a single permission item for a folder. Conflicts with the "oss.FolderPermission" resource which manages the entire set of permissions for a folder.
        		* [Official documentation](https://grafana.com/docs/grafana/latest/administration/roles-and-permissions/access-control/)
        		* [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/folder_permissions/)

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        team = grafana.oss.Team("team", name="Team Name")
        user = grafana.oss.User("user",
            email="user.name@example.com",
            login="user.name",
            password="my-password")
        collection = grafana.oss.Folder("collection", title="Folder Title")
        on_role = grafana.oss.FolderPermissionItem("on_role",
            folder_uid=collection.uid,
            role="Viewer",
            permission="Edit")
        on_team = grafana.oss.FolderPermissionItem("on_team",
            folder_uid=collection.uid,
            team=team.id,
            permission="View")
        on_user = grafana.oss.FolderPermissionItem("on_user",
            folder_uid=collection.uid,
            user=user.id,
            permission="Admin")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/folderPermissionItem:FolderPermissionItem name "{{ folderUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
        ```

        ```sh
        $ pulumi import grafana:index/folderPermissionItem:FolderPermissionItem name "{{ orgID }}:{{ folderUID }}:{{ type (role, team, or user) }}:{{ identifier }}"
        ```

        :param str resource_name: The name of the resource.
        :param FolderPermissionItemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FolderPermissionItemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 folder_uid: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 permission: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 team: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""FolderPermissionItem is deprecated: grafana.index/folderpermissionitem.FolderPermissionItem has been deprecated in favor of grafana.oss/folderpermissionitem.FolderPermissionItem""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FolderPermissionItemArgs.__new__(FolderPermissionItemArgs)

            if folder_uid is None and not opts.urn:
                raise TypeError("Missing required property 'folder_uid'")
            __props__.__dict__["folder_uid"] = folder_uid
            __props__.__dict__["org_id"] = org_id
            if permission is None and not opts.urn:
                raise TypeError("Missing required property 'permission'")
            __props__.__dict__["permission"] = permission
            __props__.__dict__["role"] = role
            __props__.__dict__["team"] = team
            __props__.__dict__["user"] = user
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:index/folderPermissionItem:FolderPermissionItem")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(FolderPermissionItem, __self__).__init__(
            'grafana:index/folderPermissionItem:FolderPermissionItem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            folder_uid: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            permission: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            team: Optional[pulumi.Input[str]] = None,
            user: Optional[pulumi.Input[str]] = None) -> 'FolderPermissionItem':
        """
        Get an existing FolderPermissionItem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] folder_uid: The UID of the folder.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the default organization is used for basic authentication, or the one that owns your service account for token authentication.
        :param pulumi.Input[str] permission: the permission to be assigned
        :param pulumi.Input[str] role: the role onto which the permission is to be assigned
        :param pulumi.Input[str] team: the team onto which the permission is to be assigned
        :param pulumi.Input[str] user: the user or service account onto which the permission is to be assigned
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FolderPermissionItemState.__new__(_FolderPermissionItemState)

        __props__.__dict__["folder_uid"] = folder_uid
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["permission"] = permission
        __props__.__dict__["role"] = role
        __props__.__dict__["team"] = team
        __props__.__dict__["user"] = user
        return FolderPermissionItem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="folderUid")
    def folder_uid(self) -> pulumi.Output[str]:
        """
        The UID of the folder.
        """
        return pulumi.get(self, "folder_uid")

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

