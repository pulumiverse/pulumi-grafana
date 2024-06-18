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

__all__ = ['DataSourcePermissionArgs', 'DataSourcePermission']

@pulumi.input_type
class DataSourcePermissionArgs:
    def __init__(__self__, *,
                 datasource_uid: pulumi.Input[str],
                 org_id: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourcePermissionPermissionArgs']]]] = None):
        """
        The set of arguments for constructing a DataSourcePermission resource.
        :param pulumi.Input[str] datasource_uid: UID of the datasource to apply permissions to.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[Sequence[pulumi.Input['DataSourcePermissionPermissionArgs']]] permissions: The permission items to add/update. Items that are omitted from the list will be removed.
        """
        pulumi.set(__self__, "datasource_uid", datasource_uid)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)

    @property
    @pulumi.getter(name="datasourceUid")
    def datasource_uid(self) -> pulumi.Input[str]:
        """
        UID of the datasource to apply permissions to.
        """
        return pulumi.get(self, "datasource_uid")

    @datasource_uid.setter
    def datasource_uid(self, value: pulumi.Input[str]):
        pulumi.set(self, "datasource_uid", value)

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
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataSourcePermissionPermissionArgs']]]]:
        """
        The permission items to add/update. Items that are omitted from the list will be removed.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourcePermissionPermissionArgs']]]]):
        pulumi.set(self, "permissions", value)


@pulumi.input_type
class _DataSourcePermissionState:
    def __init__(__self__, *,
                 datasource_uid: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourcePermissionPermissionArgs']]]] = None):
        """
        Input properties used for looking up and filtering DataSourcePermission resources.
        :param pulumi.Input[str] datasource_uid: UID of the datasource to apply permissions to.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[Sequence[pulumi.Input['DataSourcePermissionPermissionArgs']]] permissions: The permission items to add/update. Items that are omitted from the list will be removed.
        """
        if datasource_uid is not None:
            pulumi.set(__self__, "datasource_uid", datasource_uid)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)

    @property
    @pulumi.getter(name="datasourceUid")
    def datasource_uid(self) -> Optional[pulumi.Input[str]]:
        """
        UID of the datasource to apply permissions to.
        """
        return pulumi.get(self, "datasource_uid")

    @datasource_uid.setter
    def datasource_uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "datasource_uid", value)

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
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataSourcePermissionPermissionArgs']]]]:
        """
        The permission items to add/update. Items that are omitted from the list will be removed.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourcePermissionPermissionArgs']]]]):
        pulumi.set(self, "permissions", value)


class DataSourcePermission(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 datasource_uid: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourcePermissionPermissionArgs']]]]] = None,
                 __props__=None):
        """
        Manages the entire set of permissions for a datasource. Permissions that aren't specified when applying this resource will be removed.
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/datasource_permissions/)

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumiverse_grafana as grafana

        team = grafana.oss.Team("team")
        foo = grafana.oss.DataSource("foo",
            type="cloudwatch",
            json_data_encoded=json.dumps({
                "defaultRegion": "us-east-1",
                "authType": "keys",
            }),
            secure_json_data_encoded=json.dumps({
                "accessKey": "123",
                "secretKey": "456",
            }))
        user = grafana.oss.User("user",
            email="test-ds-permissions@example.com",
            login="test-ds-permissions",
            password="hunter2")
        sa = grafana.oss.ServiceAccount("sa", role="Viewer")
        foo_permissions = grafana.enterprise.DataSourcePermission("fooPermissions",
            datasource_uid=foo.uid,
            permissions=[
                grafana.enterprise.DataSourcePermissionPermissionArgs(
                    team_id=team.id,
                    permission="Edit",
                ),
                grafana.enterprise.DataSourcePermissionPermissionArgs(
                    user_id=user.id,
                    permission="Edit",
                ),
                grafana.enterprise.DataSourcePermissionPermissionArgs(
                    built_in_role="Viewer",
                    permission="Query",
                ),
                grafana.enterprise.DataSourcePermissionPermissionArgs(
                    user_id=sa.id,
                    permission="Query",
                ),
            ])
        ```

        ## Import

        ```sh
        $ pulumi import grafana:enterprise/dataSourcePermission:DataSourcePermission name "{{ datasourceID }}"
        ```

        ```sh
        $ pulumi import grafana:enterprise/dataSourcePermission:DataSourcePermission name "{{ orgID }}:{{ datasourceID }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] datasource_uid: UID of the datasource to apply permissions to.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourcePermissionPermissionArgs']]]] permissions: The permission items to add/update. Items that are omitted from the list will be removed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataSourcePermissionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages the entire set of permissions for a datasource. Permissions that aren't specified when applying this resource will be removed.
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/datasource_permissions/)

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumiverse_grafana as grafana

        team = grafana.oss.Team("team")
        foo = grafana.oss.DataSource("foo",
            type="cloudwatch",
            json_data_encoded=json.dumps({
                "defaultRegion": "us-east-1",
                "authType": "keys",
            }),
            secure_json_data_encoded=json.dumps({
                "accessKey": "123",
                "secretKey": "456",
            }))
        user = grafana.oss.User("user",
            email="test-ds-permissions@example.com",
            login="test-ds-permissions",
            password="hunter2")
        sa = grafana.oss.ServiceAccount("sa", role="Viewer")
        foo_permissions = grafana.enterprise.DataSourcePermission("fooPermissions",
            datasource_uid=foo.uid,
            permissions=[
                grafana.enterprise.DataSourcePermissionPermissionArgs(
                    team_id=team.id,
                    permission="Edit",
                ),
                grafana.enterprise.DataSourcePermissionPermissionArgs(
                    user_id=user.id,
                    permission="Edit",
                ),
                grafana.enterprise.DataSourcePermissionPermissionArgs(
                    built_in_role="Viewer",
                    permission="Query",
                ),
                grafana.enterprise.DataSourcePermissionPermissionArgs(
                    user_id=sa.id,
                    permission="Query",
                ),
            ])
        ```

        ## Import

        ```sh
        $ pulumi import grafana:enterprise/dataSourcePermission:DataSourcePermission name "{{ datasourceID }}"
        ```

        ```sh
        $ pulumi import grafana:enterprise/dataSourcePermission:DataSourcePermission name "{{ orgID }}:{{ datasourceID }}"
        ```

        :param str resource_name: The name of the resource.
        :param DataSourcePermissionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataSourcePermissionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 datasource_uid: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourcePermissionPermissionArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DataSourcePermissionArgs.__new__(DataSourcePermissionArgs)

            if datasource_uid is None and not opts.urn:
                raise TypeError("Missing required property 'datasource_uid'")
            __props__.__dict__["datasource_uid"] = datasource_uid
            __props__.__dict__["org_id"] = org_id
            __props__.__dict__["permissions"] = permissions
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:index/dataSourcePermission:DataSourcePermission")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DataSourcePermission, __self__).__init__(
            'grafana:enterprise/dataSourcePermission:DataSourcePermission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            datasource_uid: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourcePermissionPermissionArgs']]]]] = None) -> 'DataSourcePermission':
        """
        Get an existing DataSourcePermission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] datasource_uid: UID of the datasource to apply permissions to.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourcePermissionPermissionArgs']]]] permissions: The permission items to add/update. Items that are omitted from the list will be removed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DataSourcePermissionState.__new__(_DataSourcePermissionState)

        __props__.__dict__["datasource_uid"] = datasource_uid
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["permissions"] = permissions
        return DataSourcePermission(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="datasourceUid")
    def datasource_uid(self) -> pulumi.Output[str]:
        """
        UID of the datasource to apply permissions to.
        """
        return pulumi.get(self, "datasource_uid")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Optional[Sequence['outputs.DataSourcePermissionPermission']]]:
        """
        The permission items to add/update. Items that are omitted from the list will be removed.
        """
        return pulumi.get(self, "permissions")

