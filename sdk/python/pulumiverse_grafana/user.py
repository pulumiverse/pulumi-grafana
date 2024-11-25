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

__all__ = ['UserArgs', 'User']

@pulumi.input_type
class UserArgs:
    def __init__(__self__, *,
                 email: pulumi.Input[str],
                 password: pulumi.Input[str],
                 is_admin: Optional[pulumi.Input[bool]] = None,
                 login: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a User resource.
        :param pulumi.Input[str] email: The email address of the Grafana user.
        :param pulumi.Input[str] password: The password for the Grafana user.
        :param pulumi.Input[bool] is_admin: Whether to make user an admin. Defaults to `false`.
        :param pulumi.Input[str] login: The username for the Grafana user.
        :param pulumi.Input[str] name: The display name for the Grafana user.
        """
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "password", password)
        if is_admin is not None:
            pulumi.set(__self__, "is_admin", is_admin)
        if login is not None:
            pulumi.set(__self__, "login", login)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        """
        The email address of the Grafana user.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        The password for the Grafana user.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="isAdmin")
    def is_admin(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to make user an admin. Defaults to `false`.
        """
        return pulumi.get(self, "is_admin")

    @is_admin.setter
    def is_admin(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_admin", value)

    @property
    @pulumi.getter
    def login(self) -> Optional[pulumi.Input[str]]:
        """
        The username for the Grafana user.
        """
        return pulumi.get(self, "login")

    @login.setter
    def login(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "login", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name for the Grafana user.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _UserState:
    def __init__(__self__, *,
                 email: Optional[pulumi.Input[str]] = None,
                 is_admin: Optional[pulumi.Input[bool]] = None,
                 login: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering User resources.
        :param pulumi.Input[str] email: The email address of the Grafana user.
        :param pulumi.Input[bool] is_admin: Whether to make user an admin. Defaults to `false`.
        :param pulumi.Input[str] login: The username for the Grafana user.
        :param pulumi.Input[str] name: The display name for the Grafana user.
        :param pulumi.Input[str] password: The password for the Grafana user.
        :param pulumi.Input[int] user_id: The numerical ID of the Grafana user.
        """
        if email is not None:
            pulumi.set(__self__, "email", email)
        if is_admin is not None:
            pulumi.set(__self__, "is_admin", is_admin)
        if login is not None:
            pulumi.set(__self__, "login", login)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        The email address of the Grafana user.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="isAdmin")
    def is_admin(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to make user an admin. Defaults to `false`.
        """
        return pulumi.get(self, "is_admin")

    @is_admin.setter
    def is_admin(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_admin", value)

    @property
    @pulumi.getter
    def login(self) -> Optional[pulumi.Input[str]]:
        """
        The username for the Grafana user.
        """
        return pulumi.get(self, "login")

    @login.setter
    def login(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "login", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name for the Grafana user.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password for the Grafana user.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[int]]:
        """
        The numerical ID of the Grafana user.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "user_id", value)


warnings.warn("""grafana.index/user.User has been deprecated in favor of grafana.oss/user.User""", DeprecationWarning)


class User(pulumi.CustomResource):
    warnings.warn("""grafana.index/user.User has been deprecated in favor of grafana.oss/user.User""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 is_admin: Optional[pulumi.Input[bool]] = None,
                 login: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        * [Official documentation](https://grafana.com/docs/grafana/latest/administration/user-management/server-user-management/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/user/)

        This resource represents an instance-scoped resource and uses Grafana's admin APIs.
        It does not work with API tokens or service accounts which are org-scoped.
        You must use basic auth.
        This resource is also not compatible with Grafana Cloud, as it does not allow basic auth.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        staff = grafana.oss.User("staff",
            email="staff.name@example.com",
            name="Staff Name",
            login="staff",
            password="my-password",
            is_admin=False)
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/user:User name "{{ id }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] email: The email address of the Grafana user.
        :param pulumi.Input[bool] is_admin: Whether to make user an admin. Defaults to `false`.
        :param pulumi.Input[str] login: The username for the Grafana user.
        :param pulumi.Input[str] name: The display name for the Grafana user.
        :param pulumi.Input[str] password: The password for the Grafana user.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [Official documentation](https://grafana.com/docs/grafana/latest/administration/user-management/server-user-management/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/user/)

        This resource represents an instance-scoped resource and uses Grafana's admin APIs.
        It does not work with API tokens or service accounts which are org-scoped.
        You must use basic auth.
        This resource is also not compatible with Grafana Cloud, as it does not allow basic auth.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        staff = grafana.oss.User("staff",
            email="staff.name@example.com",
            name="Staff Name",
            login="staff",
            password="my-password",
            is_admin=False)
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/user:User name "{{ id }}"
        ```

        :param str resource_name: The name of the resource.
        :param UserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 is_admin: Optional[pulumi.Input[bool]] = None,
                 login: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""User is deprecated: grafana.index/user.User has been deprecated in favor of grafana.oss/user.User""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserArgs.__new__(UserArgs)

            if email is None and not opts.urn:
                raise TypeError("Missing required property 'email'")
            __props__.__dict__["email"] = email
            __props__.__dict__["is_admin"] = is_admin
            __props__.__dict__["login"] = login
            __props__.__dict__["name"] = name
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["user_id"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:index/user:User")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(User, __self__).__init__(
            'grafana:index/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            email: Optional[pulumi.Input[str]] = None,
            is_admin: Optional[pulumi.Input[bool]] = None,
            login: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[int]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] email: The email address of the Grafana user.
        :param pulumi.Input[bool] is_admin: Whether to make user an admin. Defaults to `false`.
        :param pulumi.Input[str] login: The username for the Grafana user.
        :param pulumi.Input[str] name: The display name for the Grafana user.
        :param pulumi.Input[str] password: The password for the Grafana user.
        :param pulumi.Input[int] user_id: The numerical ID of the Grafana user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserState.__new__(_UserState)

        __props__.__dict__["email"] = email
        __props__.__dict__["is_admin"] = is_admin
        __props__.__dict__["login"] = login
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["user_id"] = user_id
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        The email address of the Grafana user.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="isAdmin")
    def is_admin(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to make user an admin. Defaults to `false`.
        """
        return pulumi.get(self, "is_admin")

    @property
    @pulumi.getter
    def login(self) -> pulumi.Output[Optional[str]]:
        """
        The username for the Grafana user.
        """
        return pulumi.get(self, "login")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The display name for the Grafana user.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        The password for the Grafana user.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[int]:
        """
        The numerical ID of the Grafana user.
        """
        return pulumi.get(self, "user_id")

