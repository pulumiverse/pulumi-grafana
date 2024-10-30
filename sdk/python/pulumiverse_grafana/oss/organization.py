# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['OrganizationArgs', 'Organization']

@pulumi.input_type
class OrganizationArgs:
    def __init__(__self__, *,
                 admin_user: Optional[pulumi.Input[str]] = None,
                 admins: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 create_users: Optional[pulumi.Input[bool]] = None,
                 editors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 users_without_accesses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 viewers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Organization resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] admins: A list of email addresses corresponding to users who should be given admin
               access to the organization. Note: users specified here must already exist in
               Grafana unless 'create_users' is set to true.
        :param pulumi.Input[bool] create_users: Whether or not to create Grafana users specified in the organization's
               membership if they don't already exist in Grafana. If unspecified, this
               parameter defaults to true, creating placeholder users with the name, login,
               and email set to the email of the user, and a random password. Setting this
               option to false will cause an error to be thrown for any users that do not
               already exist in Grafana.
               Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] editors: A list of email addresses corresponding to users who should be given editor
               access to the organization. Note: users specified here must already exist in
               Grafana unless 'create_users' is set to true.
        :param pulumi.Input[str] name: The display name for the Grafana organization created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] users_without_accesses: A list of email addresses corresponding to users who should be given none access to the organization.
               Note: users specified here must already exist in Grafana, unless 'create_users' is
               set to true. This feature is only available in Grafana 10.2+.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] viewers: A list of email addresses corresponding to users who should be given viewer
               access to the organization. Note: users specified here must already exist in
               Grafana unless 'create_users' is set to true.
        """
        if admin_user is not None:
            pulumi.set(__self__, "admin_user", admin_user)
        if admins is not None:
            pulumi.set(__self__, "admins", admins)
        if create_users is not None:
            pulumi.set(__self__, "create_users", create_users)
        if editors is not None:
            pulumi.set(__self__, "editors", editors)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if users_without_accesses is not None:
            pulumi.set(__self__, "users_without_accesses", users_without_accesses)
        if viewers is not None:
            pulumi.set(__self__, "viewers", viewers)

    @property
    @pulumi.getter(name="adminUser")
    def admin_user(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "admin_user")

    @admin_user.setter
    def admin_user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_user", value)

    @property
    @pulumi.getter
    def admins(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of email addresses corresponding to users who should be given admin
        access to the organization. Note: users specified here must already exist in
        Grafana unless 'create_users' is set to true.
        """
        return pulumi.get(self, "admins")

    @admins.setter
    def admins(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "admins", value)

    @property
    @pulumi.getter(name="createUsers")
    def create_users(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to create Grafana users specified in the organization's
        membership if they don't already exist in Grafana. If unspecified, this
        parameter defaults to true, creating placeholder users with the name, login,
        and email set to the email of the user, and a random password. Setting this
        option to false will cause an error to be thrown for any users that do not
        already exist in Grafana.
        Defaults to `true`.
        """
        return pulumi.get(self, "create_users")

    @create_users.setter
    def create_users(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "create_users", value)

    @property
    @pulumi.getter
    def editors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of email addresses corresponding to users who should be given editor
        access to the organization. Note: users specified here must already exist in
        Grafana unless 'create_users' is set to true.
        """
        return pulumi.get(self, "editors")

    @editors.setter
    def editors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "editors", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name for the Grafana organization created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="usersWithoutAccesses")
    def users_without_accesses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of email addresses corresponding to users who should be given none access to the organization.
        Note: users specified here must already exist in Grafana, unless 'create_users' is
        set to true. This feature is only available in Grafana 10.2+.
        """
        return pulumi.get(self, "users_without_accesses")

    @users_without_accesses.setter
    def users_without_accesses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "users_without_accesses", value)

    @property
    @pulumi.getter
    def viewers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of email addresses corresponding to users who should be given viewer
        access to the organization. Note: users specified here must already exist in
        Grafana unless 'create_users' is set to true.
        """
        return pulumi.get(self, "viewers")

    @viewers.setter
    def viewers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "viewers", value)


@pulumi.input_type
class _OrganizationState:
    def __init__(__self__, *,
                 admin_user: Optional[pulumi.Input[str]] = None,
                 admins: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 create_users: Optional[pulumi.Input[bool]] = None,
                 editors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[int]] = None,
                 users_without_accesses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 viewers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Organization resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] admins: A list of email addresses corresponding to users who should be given admin
               access to the organization. Note: users specified here must already exist in
               Grafana unless 'create_users' is set to true.
        :param pulumi.Input[bool] create_users: Whether or not to create Grafana users specified in the organization's
               membership if they don't already exist in Grafana. If unspecified, this
               parameter defaults to true, creating placeholder users with the name, login,
               and email set to the email of the user, and a random password. Setting this
               option to false will cause an error to be thrown for any users that do not
               already exist in Grafana.
               Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] editors: A list of email addresses corresponding to users who should be given editor
               access to the organization. Note: users specified here must already exist in
               Grafana unless 'create_users' is set to true.
        :param pulumi.Input[str] name: The display name for the Grafana organization created.
        :param pulumi.Input[int] org_id: The organization id assigned to this organization by Grafana.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] users_without_accesses: A list of email addresses corresponding to users who should be given none access to the organization.
               Note: users specified here must already exist in Grafana, unless 'create_users' is
               set to true. This feature is only available in Grafana 10.2+.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] viewers: A list of email addresses corresponding to users who should be given viewer
               access to the organization. Note: users specified here must already exist in
               Grafana unless 'create_users' is set to true.
        """
        if admin_user is not None:
            pulumi.set(__self__, "admin_user", admin_user)
        if admins is not None:
            pulumi.set(__self__, "admins", admins)
        if create_users is not None:
            pulumi.set(__self__, "create_users", create_users)
        if editors is not None:
            pulumi.set(__self__, "editors", editors)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if users_without_accesses is not None:
            pulumi.set(__self__, "users_without_accesses", users_without_accesses)
        if viewers is not None:
            pulumi.set(__self__, "viewers", viewers)

    @property
    @pulumi.getter(name="adminUser")
    def admin_user(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "admin_user")

    @admin_user.setter
    def admin_user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_user", value)

    @property
    @pulumi.getter
    def admins(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of email addresses corresponding to users who should be given admin
        access to the organization. Note: users specified here must already exist in
        Grafana unless 'create_users' is set to true.
        """
        return pulumi.get(self, "admins")

    @admins.setter
    def admins(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "admins", value)

    @property
    @pulumi.getter(name="createUsers")
    def create_users(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to create Grafana users specified in the organization's
        membership if they don't already exist in Grafana. If unspecified, this
        parameter defaults to true, creating placeholder users with the name, login,
        and email set to the email of the user, and a random password. Setting this
        option to false will cause an error to be thrown for any users that do not
        already exist in Grafana.
        Defaults to `true`.
        """
        return pulumi.get(self, "create_users")

    @create_users.setter
    def create_users(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "create_users", value)

    @property
    @pulumi.getter
    def editors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of email addresses corresponding to users who should be given editor
        access to the organization. Note: users specified here must already exist in
        Grafana unless 'create_users' is set to true.
        """
        return pulumi.get(self, "editors")

    @editors.setter
    def editors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "editors", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name for the Grafana organization created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[int]]:
        """
        The organization id assigned to this organization by Grafana.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="usersWithoutAccesses")
    def users_without_accesses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of email addresses corresponding to users who should be given none access to the organization.
        Note: users specified here must already exist in Grafana, unless 'create_users' is
        set to true. This feature is only available in Grafana 10.2+.
        """
        return pulumi.get(self, "users_without_accesses")

    @users_without_accesses.setter
    def users_without_accesses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "users_without_accesses", value)

    @property
    @pulumi.getter
    def viewers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of email addresses corresponding to users who should be given viewer
        access to the organization. Note: users specified here must already exist in
        Grafana unless 'create_users' is set to true.
        """
        return pulumi.get(self, "viewers")

    @viewers.setter
    def viewers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "viewers", value)


class Organization(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_user: Optional[pulumi.Input[str]] = None,
                 admins: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 create_users: Optional[pulumi.Input[bool]] = None,
                 editors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 users_without_accesses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 viewers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        * [Official documentation](https://grafana.com/docs/grafana/latest/administration/organization-management/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/org/)

        This resource represents an instance-scoped resource and uses Grafana's admin APIs.
        It does not work with API tokens or service accounts which are org-scoped.
        You must use basic auth.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        test = grafana.oss.Organization("test",
            name="Test Organization",
            admin_user="admin",
            create_users=True,
            admins=["admin@example.com"],
            editors=[
                "editor-01@example.com",
                "editor-02@example.com",
            ],
            viewers=[
                "viewer-01@example.com",
                "viewer-02@example.com",
            ])
        ```

        ## Import

        ```sh
        $ pulumi import grafana:oss/organization:Organization name "{{ id }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] admins: A list of email addresses corresponding to users who should be given admin
               access to the organization. Note: users specified here must already exist in
               Grafana unless 'create_users' is set to true.
        :param pulumi.Input[bool] create_users: Whether or not to create Grafana users specified in the organization's
               membership if they don't already exist in Grafana. If unspecified, this
               parameter defaults to true, creating placeholder users with the name, login,
               and email set to the email of the user, and a random password. Setting this
               option to false will cause an error to be thrown for any users that do not
               already exist in Grafana.
               Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] editors: A list of email addresses corresponding to users who should be given editor
               access to the organization. Note: users specified here must already exist in
               Grafana unless 'create_users' is set to true.
        :param pulumi.Input[str] name: The display name for the Grafana organization created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] users_without_accesses: A list of email addresses corresponding to users who should be given none access to the organization.
               Note: users specified here must already exist in Grafana, unless 'create_users' is
               set to true. This feature is only available in Grafana 10.2+.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] viewers: A list of email addresses corresponding to users who should be given viewer
               access to the organization. Note: users specified here must already exist in
               Grafana unless 'create_users' is set to true.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[OrganizationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [Official documentation](https://grafana.com/docs/grafana/latest/administration/organization-management/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/org/)

        This resource represents an instance-scoped resource and uses Grafana's admin APIs.
        It does not work with API tokens or service accounts which are org-scoped.
        You must use basic auth.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        test = grafana.oss.Organization("test",
            name="Test Organization",
            admin_user="admin",
            create_users=True,
            admins=["admin@example.com"],
            editors=[
                "editor-01@example.com",
                "editor-02@example.com",
            ],
            viewers=[
                "viewer-01@example.com",
                "viewer-02@example.com",
            ])
        ```

        ## Import

        ```sh
        $ pulumi import grafana:oss/organization:Organization name "{{ id }}"
        ```

        :param str resource_name: The name of the resource.
        :param OrganizationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_user: Optional[pulumi.Input[str]] = None,
                 admins: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 create_users: Optional[pulumi.Input[bool]] = None,
                 editors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 users_without_accesses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 viewers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrganizationArgs.__new__(OrganizationArgs)

            __props__.__dict__["admin_user"] = admin_user
            __props__.__dict__["admins"] = admins
            __props__.__dict__["create_users"] = create_users
            __props__.__dict__["editors"] = editors
            __props__.__dict__["name"] = name
            __props__.__dict__["users_without_accesses"] = users_without_accesses
            __props__.__dict__["viewers"] = viewers
            __props__.__dict__["org_id"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:index/organization:Organization")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Organization, __self__).__init__(
            'grafana:oss/organization:Organization',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_user: Optional[pulumi.Input[str]] = None,
            admins: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            create_users: Optional[pulumi.Input[bool]] = None,
            editors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[int]] = None,
            users_without_accesses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            viewers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Organization':
        """
        Get an existing Organization resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] admins: A list of email addresses corresponding to users who should be given admin
               access to the organization. Note: users specified here must already exist in
               Grafana unless 'create_users' is set to true.
        :param pulumi.Input[bool] create_users: Whether or not to create Grafana users specified in the organization's
               membership if they don't already exist in Grafana. If unspecified, this
               parameter defaults to true, creating placeholder users with the name, login,
               and email set to the email of the user, and a random password. Setting this
               option to false will cause an error to be thrown for any users that do not
               already exist in Grafana.
               Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] editors: A list of email addresses corresponding to users who should be given editor
               access to the organization. Note: users specified here must already exist in
               Grafana unless 'create_users' is set to true.
        :param pulumi.Input[str] name: The display name for the Grafana organization created.
        :param pulumi.Input[int] org_id: The organization id assigned to this organization by Grafana.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] users_without_accesses: A list of email addresses corresponding to users who should be given none access to the organization.
               Note: users specified here must already exist in Grafana, unless 'create_users' is
               set to true. This feature is only available in Grafana 10.2+.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] viewers: A list of email addresses corresponding to users who should be given viewer
               access to the organization. Note: users specified here must already exist in
               Grafana unless 'create_users' is set to true.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrganizationState.__new__(_OrganizationState)

        __props__.__dict__["admin_user"] = admin_user
        __props__.__dict__["admins"] = admins
        __props__.__dict__["create_users"] = create_users
        __props__.__dict__["editors"] = editors
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["users_without_accesses"] = users_without_accesses
        __props__.__dict__["viewers"] = viewers
        return Organization(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminUser")
    def admin_user(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "admin_user")

    @property
    @pulumi.getter
    def admins(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of email addresses corresponding to users who should be given admin
        access to the organization. Note: users specified here must already exist in
        Grafana unless 'create_users' is set to true.
        """
        return pulumi.get(self, "admins")

    @property
    @pulumi.getter(name="createUsers")
    def create_users(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not to create Grafana users specified in the organization's
        membership if they don't already exist in Grafana. If unspecified, this
        parameter defaults to true, creating placeholder users with the name, login,
        and email set to the email of the user, and a random password. Setting this
        option to false will cause an error to be thrown for any users that do not
        already exist in Grafana.
        Defaults to `true`.
        """
        return pulumi.get(self, "create_users")

    @property
    @pulumi.getter
    def editors(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of email addresses corresponding to users who should be given editor
        access to the organization. Note: users specified here must already exist in
        Grafana unless 'create_users' is set to true.
        """
        return pulumi.get(self, "editors")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The display name for the Grafana organization created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[int]:
        """
        The organization id assigned to this organization by Grafana.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="usersWithoutAccesses")
    def users_without_accesses(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of email addresses corresponding to users who should be given none access to the organization.
        Note: users specified here must already exist in Grafana, unless 'create_users' is
        set to true. This feature is only available in Grafana 10.2+.
        """
        return pulumi.get(self, "users_without_accesses")

    @property
    @pulumi.getter
    def viewers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of email addresses corresponding to users who should be given viewer
        access to the organization. Note: users specified here must already exist in
        Grafana unless 'create_users' is set to true.
        """
        return pulumi.get(self, "viewers")

