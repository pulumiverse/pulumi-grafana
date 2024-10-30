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

__all__ = ['DashboardPublicArgs', 'DashboardPublic']

@pulumi.input_type
class DashboardPublicArgs:
    def __init__(__self__, *,
                 dashboard_uid: pulumi.Input[str],
                 access_token: Optional[pulumi.Input[str]] = None,
                 annotations_enabled: Optional[pulumi.Input[bool]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 share: Optional[pulumi.Input[str]] = None,
                 time_selection_enabled: Optional[pulumi.Input[bool]] = None,
                 uid: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DashboardPublic resource.
        :param pulumi.Input[str] dashboard_uid: The unique identifier of the original dashboard.
        :param pulumi.Input[str] access_token: A public unique identifier of a public dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a public dashboard.
        :param pulumi.Input[bool] annotations_enabled: Set to `true` to show annotations. The default value is `false`.
        :param pulumi.Input[bool] is_enabled: Set to `true` to enable the public dashboard. The default value is `false`.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[str] share: Set the share mode. The default value is `public`.
        :param pulumi.Input[bool] time_selection_enabled: Set to `true` to enable the time picker in the public dashboard. The default value is `false`.
        :param pulumi.Input[str] uid: The unique identifier of a public dashboard. It's automatically generated if not provided when creating a public dashboard.
        """
        pulumi.set(__self__, "dashboard_uid", dashboard_uid)
        if access_token is not None:
            pulumi.set(__self__, "access_token", access_token)
        if annotations_enabled is not None:
            pulumi.set(__self__, "annotations_enabled", annotations_enabled)
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if share is not None:
            pulumi.set(__self__, "share", share)
        if time_selection_enabled is not None:
            pulumi.set(__self__, "time_selection_enabled", time_selection_enabled)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)

    @property
    @pulumi.getter(name="dashboardUid")
    def dashboard_uid(self) -> pulumi.Input[str]:
        """
        The unique identifier of the original dashboard.
        """
        return pulumi.get(self, "dashboard_uid")

    @dashboard_uid.setter
    def dashboard_uid(self, value: pulumi.Input[str]):
        pulumi.set(self, "dashboard_uid", value)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> Optional[pulumi.Input[str]]:
        """
        A public unique identifier of a public dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a public dashboard.
        """
        return pulumi.get(self, "access_token")

    @access_token.setter
    def access_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_token", value)

    @property
    @pulumi.getter(name="annotationsEnabled")
    def annotations_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to show annotations. The default value is `false`.
        """
        return pulumi.get(self, "annotations_enabled")

    @annotations_enabled.setter
    def annotations_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "annotations_enabled", value)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to enable the public dashboard. The default value is `false`.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_enabled", value)

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
    def share(self) -> Optional[pulumi.Input[str]]:
        """
        Set the share mode. The default value is `public`.
        """
        return pulumi.get(self, "share")

    @share.setter
    def share(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "share", value)

    @property
    @pulumi.getter(name="timeSelectionEnabled")
    def time_selection_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to enable the time picker in the public dashboard. The default value is `false`.
        """
        return pulumi.get(self, "time_selection_enabled")

    @time_selection_enabled.setter
    def time_selection_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "time_selection_enabled", value)

    @property
    @pulumi.getter
    def uid(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of a public dashboard. It's automatically generated if not provided when creating a public dashboard.
        """
        return pulumi.get(self, "uid")

    @uid.setter
    def uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uid", value)


@pulumi.input_type
class _DashboardPublicState:
    def __init__(__self__, *,
                 access_token: Optional[pulumi.Input[str]] = None,
                 annotations_enabled: Optional[pulumi.Input[bool]] = None,
                 dashboard_uid: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 share: Optional[pulumi.Input[str]] = None,
                 time_selection_enabled: Optional[pulumi.Input[bool]] = None,
                 uid: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DashboardPublic resources.
        :param pulumi.Input[str] access_token: A public unique identifier of a public dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a public dashboard.
        :param pulumi.Input[bool] annotations_enabled: Set to `true` to show annotations. The default value is `false`.
        :param pulumi.Input[str] dashboard_uid: The unique identifier of the original dashboard.
        :param pulumi.Input[bool] is_enabled: Set to `true` to enable the public dashboard. The default value is `false`.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[str] share: Set the share mode. The default value is `public`.
        :param pulumi.Input[bool] time_selection_enabled: Set to `true` to enable the time picker in the public dashboard. The default value is `false`.
        :param pulumi.Input[str] uid: The unique identifier of a public dashboard. It's automatically generated if not provided when creating a public dashboard.
        """
        if access_token is not None:
            pulumi.set(__self__, "access_token", access_token)
        if annotations_enabled is not None:
            pulumi.set(__self__, "annotations_enabled", annotations_enabled)
        if dashboard_uid is not None:
            pulumi.set(__self__, "dashboard_uid", dashboard_uid)
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if share is not None:
            pulumi.set(__self__, "share", share)
        if time_selection_enabled is not None:
            pulumi.set(__self__, "time_selection_enabled", time_selection_enabled)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> Optional[pulumi.Input[str]]:
        """
        A public unique identifier of a public dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a public dashboard.
        """
        return pulumi.get(self, "access_token")

    @access_token.setter
    def access_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_token", value)

    @property
    @pulumi.getter(name="annotationsEnabled")
    def annotations_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to show annotations. The default value is `false`.
        """
        return pulumi.get(self, "annotations_enabled")

    @annotations_enabled.setter
    def annotations_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "annotations_enabled", value)

    @property
    @pulumi.getter(name="dashboardUid")
    def dashboard_uid(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of the original dashboard.
        """
        return pulumi.get(self, "dashboard_uid")

    @dashboard_uid.setter
    def dashboard_uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dashboard_uid", value)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to enable the public dashboard. The default value is `false`.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_enabled", value)

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
    def share(self) -> Optional[pulumi.Input[str]]:
        """
        Set the share mode. The default value is `public`.
        """
        return pulumi.get(self, "share")

    @share.setter
    def share(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "share", value)

    @property
    @pulumi.getter(name="timeSelectionEnabled")
    def time_selection_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to enable the time picker in the public dashboard. The default value is `false`.
        """
        return pulumi.get(self, "time_selection_enabled")

    @time_selection_enabled.setter
    def time_selection_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "time_selection_enabled", value)

    @property
    @pulumi.getter
    def uid(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of a public dashboard. It's automatically generated if not provided when creating a public dashboard.
        """
        return pulumi.get(self, "uid")

    @uid.setter
    def uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uid", value)


warnings.warn("""grafana.index/dashboardpublic.DashboardPublic has been deprecated in favor of grafana.oss/dashboardpublic.DashboardPublic""", DeprecationWarning)


class DashboardPublic(pulumi.CustomResource):
    warnings.warn("""grafana.index/dashboardpublic.DashboardPublic has been deprecated in favor of grafana.oss/dashboardpublic.DashboardPublic""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_token: Optional[pulumi.Input[str]] = None,
                 annotations_enabled: Optional[pulumi.Input[bool]] = None,
                 dashboard_uid: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 share: Optional[pulumi.Input[str]] = None,
                 time_selection_enabled: Optional[pulumi.Input[bool]] = None,
                 uid: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages Grafana public dashboards.

        **Note:** This resource is available only with Grafana 10.2+.

        * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/dashboard-public/)
        * [HTTP API](https://grafana.com/docs/grafana/next/developers/http_api/dashboard_public/)

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumiverse_grafana as grafana

        # Optional (On-premise, not supported in Grafana Cloud): Create an organization
        my_org = grafana.oss.Organization("my_org", name="test 1")
        # Create resources (optional: within the organization)
        my_folder = grafana.oss.Folder("my_folder",
            org_id=my_org.org_id,
            title="test Folder")
        test_dash = grafana.oss.Dashboard("test_dash",
            org_id=my_org.org_id,
            folder=my_folder.id,
            config_json=json.dumps({
                "title": "My Terraform Dashboard",
                "uid": "my-dashboard-uid",
            }))
        my_public_dashboard = grafana.oss.DashboardPublic("my_public_dashboard",
            org_id=my_org.org_id,
            dashboard_uid=test_dash.uid,
            uid="my-custom-public-uid",
            access_token="e99e4275da6f410d83760eefa934d8d2",
            time_selection_enabled=True,
            is_enabled=True,
            annotations_enabled=True,
            share="public")
        # Optional (On-premise, not supported in Grafana Cloud): Create an organization
        my_org2 = grafana.oss.Organization("my_org2", name="test 2")
        test_dash2 = grafana.oss.Dashboard("test_dash2",
            org_id=my_org2.org_id,
            config_json=json.dumps({
                "title": "My Terraform Dashboard2",
                "uid": "my-dashboard-uid2",
            }))
        my_public_dashboard2 = grafana.oss.DashboardPublic("my_public_dashboard2",
            org_id=my_org2.org_id,
            dashboard_uid=test_dash2.uid,
            share="public")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/dashboardPublic:DashboardPublic name "{{ dashboardUID }}:{{ publicDashboardUID }}"
        ```

        ```sh
        $ pulumi import grafana:index/dashboardPublic:DashboardPublic name "{{ orgID }}:{{ dashboardUID }}:{{ publicDashboardUID }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_token: A public unique identifier of a public dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a public dashboard.
        :param pulumi.Input[bool] annotations_enabled: Set to `true` to show annotations. The default value is `false`.
        :param pulumi.Input[str] dashboard_uid: The unique identifier of the original dashboard.
        :param pulumi.Input[bool] is_enabled: Set to `true` to enable the public dashboard. The default value is `false`.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[str] share: Set the share mode. The default value is `public`.
        :param pulumi.Input[bool] time_selection_enabled: Set to `true` to enable the time picker in the public dashboard. The default value is `false`.
        :param pulumi.Input[str] uid: The unique identifier of a public dashboard. It's automatically generated if not provided when creating a public dashboard.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DashboardPublicArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Grafana public dashboards.

        **Note:** This resource is available only with Grafana 10.2+.

        * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/dashboard-public/)
        * [HTTP API](https://grafana.com/docs/grafana/next/developers/http_api/dashboard_public/)

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumiverse_grafana as grafana

        # Optional (On-premise, not supported in Grafana Cloud): Create an organization
        my_org = grafana.oss.Organization("my_org", name="test 1")
        # Create resources (optional: within the organization)
        my_folder = grafana.oss.Folder("my_folder",
            org_id=my_org.org_id,
            title="test Folder")
        test_dash = grafana.oss.Dashboard("test_dash",
            org_id=my_org.org_id,
            folder=my_folder.id,
            config_json=json.dumps({
                "title": "My Terraform Dashboard",
                "uid": "my-dashboard-uid",
            }))
        my_public_dashboard = grafana.oss.DashboardPublic("my_public_dashboard",
            org_id=my_org.org_id,
            dashboard_uid=test_dash.uid,
            uid="my-custom-public-uid",
            access_token="e99e4275da6f410d83760eefa934d8d2",
            time_selection_enabled=True,
            is_enabled=True,
            annotations_enabled=True,
            share="public")
        # Optional (On-premise, not supported in Grafana Cloud): Create an organization
        my_org2 = grafana.oss.Organization("my_org2", name="test 2")
        test_dash2 = grafana.oss.Dashboard("test_dash2",
            org_id=my_org2.org_id,
            config_json=json.dumps({
                "title": "My Terraform Dashboard2",
                "uid": "my-dashboard-uid2",
            }))
        my_public_dashboard2 = grafana.oss.DashboardPublic("my_public_dashboard2",
            org_id=my_org2.org_id,
            dashboard_uid=test_dash2.uid,
            share="public")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/dashboardPublic:DashboardPublic name "{{ dashboardUID }}:{{ publicDashboardUID }}"
        ```

        ```sh
        $ pulumi import grafana:index/dashboardPublic:DashboardPublic name "{{ orgID }}:{{ dashboardUID }}:{{ publicDashboardUID }}"
        ```

        :param str resource_name: The name of the resource.
        :param DashboardPublicArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DashboardPublicArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_token: Optional[pulumi.Input[str]] = None,
                 annotations_enabled: Optional[pulumi.Input[bool]] = None,
                 dashboard_uid: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 share: Optional[pulumi.Input[str]] = None,
                 time_selection_enabled: Optional[pulumi.Input[bool]] = None,
                 uid: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""DashboardPublic is deprecated: grafana.index/dashboardpublic.DashboardPublic has been deprecated in favor of grafana.oss/dashboardpublic.DashboardPublic""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DashboardPublicArgs.__new__(DashboardPublicArgs)

            __props__.__dict__["access_token"] = access_token
            __props__.__dict__["annotations_enabled"] = annotations_enabled
            if dashboard_uid is None and not opts.urn:
                raise TypeError("Missing required property 'dashboard_uid'")
            __props__.__dict__["dashboard_uid"] = dashboard_uid
            __props__.__dict__["is_enabled"] = is_enabled
            __props__.__dict__["org_id"] = org_id
            __props__.__dict__["share"] = share
            __props__.__dict__["time_selection_enabled"] = time_selection_enabled
            __props__.__dict__["uid"] = uid
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:index/dashboardPublic:DashboardPublic")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DashboardPublic, __self__).__init__(
            'grafana:index/dashboardPublic:DashboardPublic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_token: Optional[pulumi.Input[str]] = None,
            annotations_enabled: Optional[pulumi.Input[bool]] = None,
            dashboard_uid: Optional[pulumi.Input[str]] = None,
            is_enabled: Optional[pulumi.Input[bool]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            share: Optional[pulumi.Input[str]] = None,
            time_selection_enabled: Optional[pulumi.Input[bool]] = None,
            uid: Optional[pulumi.Input[str]] = None) -> 'DashboardPublic':
        """
        Get an existing DashboardPublic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_token: A public unique identifier of a public dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a public dashboard.
        :param pulumi.Input[bool] annotations_enabled: Set to `true` to show annotations. The default value is `false`.
        :param pulumi.Input[str] dashboard_uid: The unique identifier of the original dashboard.
        :param pulumi.Input[bool] is_enabled: Set to `true` to enable the public dashboard. The default value is `false`.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[str] share: Set the share mode. The default value is `public`.
        :param pulumi.Input[bool] time_selection_enabled: Set to `true` to enable the time picker in the public dashboard. The default value is `false`.
        :param pulumi.Input[str] uid: The unique identifier of a public dashboard. It's automatically generated if not provided when creating a public dashboard.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DashboardPublicState.__new__(_DashboardPublicState)

        __props__.__dict__["access_token"] = access_token
        __props__.__dict__["annotations_enabled"] = annotations_enabled
        __props__.__dict__["dashboard_uid"] = dashboard_uid
        __props__.__dict__["is_enabled"] = is_enabled
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["share"] = share
        __props__.__dict__["time_selection_enabled"] = time_selection_enabled
        __props__.__dict__["uid"] = uid
        return DashboardPublic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> pulumi.Output[str]:
        """
        A public unique identifier of a public dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a public dashboard.
        """
        return pulumi.get(self, "access_token")

    @property
    @pulumi.getter(name="annotationsEnabled")
    def annotations_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `true` to show annotations. The default value is `false`.
        """
        return pulumi.get(self, "annotations_enabled")

    @property
    @pulumi.getter(name="dashboardUid")
    def dashboard_uid(self) -> pulumi.Output[str]:
        """
        The unique identifier of the original dashboard.
        """
        return pulumi.get(self, "dashboard_uid")

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `true` to enable the public dashboard. The default value is `false`.
        """
        return pulumi.get(self, "is_enabled")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def share(self) -> pulumi.Output[Optional[str]]:
        """
        Set the share mode. The default value is `public`.
        """
        return pulumi.get(self, "share")

    @property
    @pulumi.getter(name="timeSelectionEnabled")
    def time_selection_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `true` to enable the time picker in the public dashboard. The default value is `false`.
        """
        return pulumi.get(self, "time_selection_enabled")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[str]:
        """
        The unique identifier of a public dashboard. It's automatically generated if not provided when creating a public dashboard.
        """
        return pulumi.get(self, "uid")

