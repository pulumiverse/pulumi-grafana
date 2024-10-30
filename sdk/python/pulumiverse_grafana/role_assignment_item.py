# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RoleAssignmentItemArgs', 'RoleAssignmentItem']

@pulumi.input_type
class RoleAssignmentItemArgs:
    def __init__(__self__, *,
                 role_uid: pulumi.Input[str],
                 org_id: Optional[pulumi.Input[str]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RoleAssignmentItem resource.
        :param pulumi.Input[str] role_uid: the role UID onto which to assign an actor
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[str] service_account_id: the service account onto which the role is to be assigned
        :param pulumi.Input[str] team_id: the team onto which the role is to be assigned
        :param pulumi.Input[str] user_id: the user onto which the role is to be assigned
        """
        pulumi.set(__self__, "role_uid", role_uid)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if service_account_id is not None:
            pulumi.set(__self__, "service_account_id", service_account_id)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="roleUid")
    def role_uid(self) -> pulumi.Input[str]:
        """
        the role UID onto which to assign an actor
        """
        return pulumi.get(self, "role_uid")

    @role_uid.setter
    def role_uid(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_uid", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        the service account onto which the role is to be assigned
        """
        return pulumi.get(self, "service_account_id")

    @service_account_id.setter
    def service_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_id", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        the team onto which the role is to be assigned
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        the user onto which the role is to be assigned
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


@pulumi.input_type
class _RoleAssignmentItemState:
    def __init__(__self__, *,
                 org_id: Optional[pulumi.Input[str]] = None,
                 role_uid: Optional[pulumi.Input[str]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RoleAssignmentItem resources.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[str] role_uid: the role UID onto which to assign an actor
        :param pulumi.Input[str] service_account_id: the service account onto which the role is to be assigned
        :param pulumi.Input[str] team_id: the team onto which the role is to be assigned
        :param pulumi.Input[str] user_id: the user onto which the role is to be assigned
        """
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if role_uid is not None:
            pulumi.set(__self__, "role_uid", role_uid)
        if service_account_id is not None:
            pulumi.set(__self__, "service_account_id", service_account_id)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="roleUid")
    def role_uid(self) -> Optional[pulumi.Input[str]]:
        """
        the role UID onto which to assign an actor
        """
        return pulumi.get(self, "role_uid")

    @role_uid.setter
    def role_uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_uid", value)

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        the service account onto which the role is to be assigned
        """
        return pulumi.get(self, "service_account_id")

    @service_account_id.setter
    def service_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_id", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        the team onto which the role is to be assigned
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        the user onto which the role is to be assigned
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


warnings.warn("""grafana.index/roleassignmentitem.RoleAssignmentItem has been deprecated in favor of grafana.enterprise/roleassignmentitem.RoleAssignmentItem""", DeprecationWarning)


class RoleAssignmentItem(pulumi.CustomResource):
    warnings.warn("""grafana.index/roleassignmentitem.RoleAssignmentItem has been deprecated in favor of grafana.enterprise/roleassignmentitem.RoleAssignmentItem""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 role_uid: Optional[pulumi.Input[str]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a single assignment for a role. Conflicts with the "enterprise.RoleAssignment" resource which manages the entire set of assignments for a role.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        test_role = grafana.enterprise.Role("test_role",
            name="Test Role",
            uid="testrole",
            version=1,
            global_=True,
            permissions=[grafana.enterprise.RolePermissionArgs(
                action="org.users:add",
                scope="users:*",
            )])
        test_team = grafana.oss.Team("test_team", name="terraform_test_team")
        test_user = grafana.oss.User("test_user",
            email="terraform_user@test.com",
            login="terraform_user@test.com",
            password="password")
        test_sa = grafana.oss.ServiceAccount("test_sa",
            name="terraform_test_sa",
            role="Viewer")
        user = grafana.enterprise.RoleAssignmentItem("user",
            role_uid=test_role.uid,
            user_id=test_user.id)
        team = grafana.enterprise.RoleAssignmentItem("team",
            role_uid=test_role.uid,
            team_id=test_team.id)
        service_account = grafana.enterprise.RoleAssignmentItem("service_account",
            role_uid=test_role.uid,
            service_account_id=test_sa.id)
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/roleAssignmentItem:RoleAssignmentItem name "{{ roleUID }}:{{ type (user, team or service_account) }}:{{ identifier }}"
        ```

        ```sh
        $ pulumi import grafana:index/roleAssignmentItem:RoleAssignmentItem name "{{ orgID }}:{{ roleUID }}:{{ type (user, team or service_account) }}:{{ identifier }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[str] role_uid: the role UID onto which to assign an actor
        :param pulumi.Input[str] service_account_id: the service account onto which the role is to be assigned
        :param pulumi.Input[str] team_id: the team onto which the role is to be assigned
        :param pulumi.Input[str] user_id: the user onto which the role is to be assigned
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RoleAssignmentItemArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a single assignment for a role. Conflicts with the "enterprise.RoleAssignment" resource which manages the entire set of assignments for a role.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        test_role = grafana.enterprise.Role("test_role",
            name="Test Role",
            uid="testrole",
            version=1,
            global_=True,
            permissions=[grafana.enterprise.RolePermissionArgs(
                action="org.users:add",
                scope="users:*",
            )])
        test_team = grafana.oss.Team("test_team", name="terraform_test_team")
        test_user = grafana.oss.User("test_user",
            email="terraform_user@test.com",
            login="terraform_user@test.com",
            password="password")
        test_sa = grafana.oss.ServiceAccount("test_sa",
            name="terraform_test_sa",
            role="Viewer")
        user = grafana.enterprise.RoleAssignmentItem("user",
            role_uid=test_role.uid,
            user_id=test_user.id)
        team = grafana.enterprise.RoleAssignmentItem("team",
            role_uid=test_role.uid,
            team_id=test_team.id)
        service_account = grafana.enterprise.RoleAssignmentItem("service_account",
            role_uid=test_role.uid,
            service_account_id=test_sa.id)
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/roleAssignmentItem:RoleAssignmentItem name "{{ roleUID }}:{{ type (user, team or service_account) }}:{{ identifier }}"
        ```

        ```sh
        $ pulumi import grafana:index/roleAssignmentItem:RoleAssignmentItem name "{{ orgID }}:{{ roleUID }}:{{ type (user, team or service_account) }}:{{ identifier }}"
        ```

        :param str resource_name: The name of the resource.
        :param RoleAssignmentItemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RoleAssignmentItemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 role_uid: Optional[pulumi.Input[str]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""RoleAssignmentItem is deprecated: grafana.index/roleassignmentitem.RoleAssignmentItem has been deprecated in favor of grafana.enterprise/roleassignmentitem.RoleAssignmentItem""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RoleAssignmentItemArgs.__new__(RoleAssignmentItemArgs)

            __props__.__dict__["org_id"] = org_id
            if role_uid is None and not opts.urn:
                raise TypeError("Missing required property 'role_uid'")
            __props__.__dict__["role_uid"] = role_uid
            __props__.__dict__["service_account_id"] = service_account_id
            __props__.__dict__["team_id"] = team_id
            __props__.__dict__["user_id"] = user_id
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:index/roleAssignmentItem:RoleAssignmentItem")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(RoleAssignmentItem, __self__).__init__(
            'grafana:index/roleAssignmentItem:RoleAssignmentItem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            role_uid: Optional[pulumi.Input[str]] = None,
            service_account_id: Optional[pulumi.Input[str]] = None,
            team_id: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'RoleAssignmentItem':
        """
        Get an existing RoleAssignmentItem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[str] role_uid: the role UID onto which to assign an actor
        :param pulumi.Input[str] service_account_id: the service account onto which the role is to be assigned
        :param pulumi.Input[str] team_id: the team onto which the role is to be assigned
        :param pulumi.Input[str] user_id: the user onto which the role is to be assigned
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RoleAssignmentItemState.__new__(_RoleAssignmentItemState)

        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["role_uid"] = role_uid
        __props__.__dict__["service_account_id"] = service_account_id
        __props__.__dict__["team_id"] = team_id
        __props__.__dict__["user_id"] = user_id
        return RoleAssignmentItem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[str]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="roleUid")
    def role_uid(self) -> pulumi.Output[str]:
        """
        the role UID onto which to assign an actor
        """
        return pulumi.get(self, "role_uid")

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> pulumi.Output[Optional[str]]:
        """
        the service account onto which the role is to be assigned
        """
        return pulumi.get(self, "service_account_id")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Output[Optional[str]]:
        """
        the team onto which the role is to be assigned
        """
        return pulumi.get(self, "team_id")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[Optional[str]]:
        """
        the user onto which the role is to be assigned
        """
        return pulumi.get(self, "user_id")

