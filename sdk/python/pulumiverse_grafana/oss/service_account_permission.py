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
from ._inputs import *

__all__ = ['ServiceAccountPermissionArgs', 'ServiceAccountPermission']

@pulumi.input_type
class ServiceAccountPermissionArgs:
    def __init__(__self__, *,
                 service_account_id: pulumi.Input[str],
                 org_id: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceAccountPermissionPermissionArgs']]]] = None):
        """
        The set of arguments for constructing a ServiceAccountPermission resource.
        :param pulumi.Input[str] service_account_id: The id of the service account.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceAccountPermissionPermissionArgs']]] permissions: The permission items to add/update. Items that are omitted from the list will be removed.
        """
        pulumi.set(__self__, "service_account_id", service_account_id)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> pulumi.Input[str]:
        """
        The id of the service account.
        """
        return pulumi.get(self, "service_account_id")

    @service_account_id.setter
    def service_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_account_id", value)

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
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceAccountPermissionPermissionArgs']]]]:
        """
        The permission items to add/update. Items that are omitted from the list will be removed.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceAccountPermissionPermissionArgs']]]]):
        pulumi.set(self, "permissions", value)


@pulumi.input_type
class _ServiceAccountPermissionState:
    def __init__(__self__, *,
                 org_id: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceAccountPermissionPermissionArgs']]]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceAccountPermission resources.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceAccountPermissionPermissionArgs']]] permissions: The permission items to add/update. Items that are omitted from the list will be removed.
        :param pulumi.Input[str] service_account_id: The id of the service account.
        """
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if service_account_id is not None:
            pulumi.set(__self__, "service_account_id", service_account_id)

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
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceAccountPermissionPermissionArgs']]]]:
        """
        The permission items to add/update. Items that are omitted from the list will be removed.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceAccountPermissionPermissionArgs']]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the service account.
        """
        return pulumi.get(self, "service_account_id")

    @service_account_id.setter
    def service_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_id", value)


class ServiceAccountPermission(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceAccountPermissionPermissionArgs']]]]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages the entire set of permissions for a service account. Permissions that aren't specified when applying this resource will be removed.

        **Note:** This resource is available from Grafana 9.2.4 onwards.

        * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/#manage-users-and-teams-permissions-for-a-service-account-in-grafana)

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        test = grafana.oss.ServiceAccount("test",
            name="sa-terraform-test",
            role="Editor",
            is_disabled=False)
        test_team = grafana.oss.Team("test_team", name="tf_test_team")
        test_user = grafana.oss.User("test_user",
            email="tf_user@test.com",
            login="tf_user@test.com",
            password="password")
        test_permissions = grafana.oss.ServiceAccountPermission("test_permissions",
            service_account_id=test.id,
            permissions=[
                grafana.oss.ServiceAccountPermissionPermissionArgs(
                    user_id=test_user.id,
                    permission="Edit",
                ),
                grafana.oss.ServiceAccountPermissionPermissionArgs(
                    team_id=test_team.id,
                    permission="Admin",
                ),
            ])
        ```

        ## Import

        ```sh
        $ pulumi import grafana:oss/serviceAccountPermission:ServiceAccountPermission name "{{ serviceAccountID }}"
        ```

        ```sh
        $ pulumi import grafana:oss/serviceAccountPermission:ServiceAccountPermission name "{{ orgID }}:{{ serviceAccountID }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceAccountPermissionPermissionArgs']]]] permissions: The permission items to add/update. Items that are omitted from the list will be removed.
        :param pulumi.Input[str] service_account_id: The id of the service account.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceAccountPermissionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages the entire set of permissions for a service account. Permissions that aren't specified when applying this resource will be removed.

        **Note:** This resource is available from Grafana 9.2.4 onwards.

        * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/#manage-users-and-teams-permissions-for-a-service-account-in-grafana)

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        test = grafana.oss.ServiceAccount("test",
            name="sa-terraform-test",
            role="Editor",
            is_disabled=False)
        test_team = grafana.oss.Team("test_team", name="tf_test_team")
        test_user = grafana.oss.User("test_user",
            email="tf_user@test.com",
            login="tf_user@test.com",
            password="password")
        test_permissions = grafana.oss.ServiceAccountPermission("test_permissions",
            service_account_id=test.id,
            permissions=[
                grafana.oss.ServiceAccountPermissionPermissionArgs(
                    user_id=test_user.id,
                    permission="Edit",
                ),
                grafana.oss.ServiceAccountPermissionPermissionArgs(
                    team_id=test_team.id,
                    permission="Admin",
                ),
            ])
        ```

        ## Import

        ```sh
        $ pulumi import grafana:oss/serviceAccountPermission:ServiceAccountPermission name "{{ serviceAccountID }}"
        ```

        ```sh
        $ pulumi import grafana:oss/serviceAccountPermission:ServiceAccountPermission name "{{ orgID }}:{{ serviceAccountID }}"
        ```

        :param str resource_name: The name of the resource.
        :param ServiceAccountPermissionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceAccountPermissionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceAccountPermissionPermissionArgs']]]]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceAccountPermissionArgs.__new__(ServiceAccountPermissionArgs)

            __props__.__dict__["org_id"] = org_id
            __props__.__dict__["permissions"] = permissions
            if service_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_account_id'")
            __props__.__dict__["service_account_id"] = service_account_id
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:index/serviceAccountPermission:ServiceAccountPermission")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ServiceAccountPermission, __self__).__init__(
            'grafana:oss/serviceAccountPermission:ServiceAccountPermission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceAccountPermissionPermissionArgs']]]]] = None,
            service_account_id: Optional[pulumi.Input[str]] = None) -> 'ServiceAccountPermission':
        """
        Get an existing ServiceAccountPermission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceAccountPermissionPermissionArgs']]]] permissions: The permission items to add/update. Items that are omitted from the list will be removed.
        :param pulumi.Input[str] service_account_id: The id of the service account.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceAccountPermissionState.__new__(_ServiceAccountPermissionState)

        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["permissions"] = permissions
        __props__.__dict__["service_account_id"] = service_account_id
        return ServiceAccountPermission(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Optional[Sequence['outputs.ServiceAccountPermissionPermission']]]:
        """
        The permission items to add/update. Items that are omitted from the list will be removed.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> pulumi.Output[str]:
        """
        The id of the service account.
        """
        return pulumi.get(self, "service_account_id")

