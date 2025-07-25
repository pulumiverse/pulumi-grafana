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
    'DataSourcePermissionPermission',
    'ReportDashboard',
    'ReportDashboardTimeRange',
    'ReportSchedule',
    'RolePermission',
    'GetRolePermissionResult',
]

@pulumi.output_type
class DataSourcePermissionPermission(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "builtInRole":
            suggest = "built_in_role"
        elif key == "teamId":
            suggest = "team_id"
        elif key == "userId":
            suggest = "user_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DataSourcePermissionPermission. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DataSourcePermissionPermission.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DataSourcePermissionPermission.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 permission: builtins.str,
                 built_in_role: Optional[builtins.str] = None,
                 team_id: Optional[builtins.str] = None,
                 user_id: Optional[builtins.str] = None):
        """
        :param builtins.str permission: Permission to associate with item. Options: `Query`, `Edit` or `Admin` (`Admin` can only be used with Grafana v10.3.0+).
        :param builtins.str built_in_role: Name of the basic role to manage permissions for. Options: `Viewer`, `Editor` or `Admin`.
        :param builtins.str team_id: ID of the team to manage permissions for. Defaults to `0`.
        :param builtins.str user_id: ID of the user or service account to manage permissions for. Defaults to `0`.
        """
        pulumi.set(__self__, "permission", permission)
        if built_in_role is not None:
            pulumi.set(__self__, "built_in_role", built_in_role)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter
    def permission(self) -> builtins.str:
        """
        Permission to associate with item. Options: `Query`, `Edit` or `Admin` (`Admin` can only be used with Grafana v10.3.0+).
        """
        return pulumi.get(self, "permission")

    @property
    @pulumi.getter(name="builtInRole")
    def built_in_role(self) -> Optional[builtins.str]:
        """
        Name of the basic role to manage permissions for. Options: `Viewer`, `Editor` or `Admin`.
        """
        return pulumi.get(self, "built_in_role")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[builtins.str]:
        """
        ID of the team to manage permissions for. Defaults to `0`.
        """
        return pulumi.get(self, "team_id")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[builtins.str]:
        """
        ID of the user or service account to manage permissions for. Defaults to `0`.
        """
        return pulumi.get(self, "user_id")


@pulumi.output_type
class ReportDashboard(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "reportVariables":
            suggest = "report_variables"
        elif key == "timeRange":
            suggest = "time_range"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ReportDashboard. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ReportDashboard.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ReportDashboard.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 uid: builtins.str,
                 report_variables: Optional[Mapping[str, builtins.str]] = None,
                 time_range: Optional['outputs.ReportDashboardTimeRange'] = None):
        """
        :param builtins.str uid: Dashboard uid.
        :param Mapping[str, builtins.str] report_variables: Add report variables to the dashboard. Values should be separated by commas.
        :param 'ReportDashboardTimeRangeArgs' time_range: Time range of the report.
        """
        pulumi.set(__self__, "uid", uid)
        if report_variables is not None:
            pulumi.set(__self__, "report_variables", report_variables)
        if time_range is not None:
            pulumi.set(__self__, "time_range", time_range)

    @property
    @pulumi.getter
    def uid(self) -> builtins.str:
        """
        Dashboard uid.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="reportVariables")
    def report_variables(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Add report variables to the dashboard. Values should be separated by commas.
        """
        return pulumi.get(self, "report_variables")

    @property
    @pulumi.getter(name="timeRange")
    def time_range(self) -> Optional['outputs.ReportDashboardTimeRange']:
        """
        Time range of the report.
        """
        return pulumi.get(self, "time_range")


@pulumi.output_type
class ReportDashboardTimeRange(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "from":
            suggest = "from_"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ReportDashboardTimeRange. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ReportDashboardTimeRange.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ReportDashboardTimeRange.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 from_: Optional[builtins.str] = None,
                 to: Optional[builtins.str] = None):
        """
        :param builtins.str from_: Start of the time range.
        :param builtins.str to: End of the time range.
        """
        if from_ is not None:
            pulumi.set(__self__, "from_", from_)
        if to is not None:
            pulumi.set(__self__, "to", to)

    @property
    @pulumi.getter(name="from")
    def from_(self) -> Optional[builtins.str]:
        """
        Start of the time range.
        """
        return pulumi.get(self, "from_")

    @property
    @pulumi.getter
    def to(self) -> Optional[builtins.str]:
        """
        End of the time range.
        """
        return pulumi.get(self, "to")


@pulumi.output_type
class ReportSchedule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "customInterval":
            suggest = "custom_interval"
        elif key == "endTime":
            suggest = "end_time"
        elif key == "lastDayOfMonth":
            suggest = "last_day_of_month"
        elif key == "startTime":
            suggest = "start_time"
        elif key == "workdaysOnly":
            suggest = "workdays_only"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ReportSchedule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ReportSchedule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ReportSchedule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 frequency: builtins.str,
                 custom_interval: Optional[builtins.str] = None,
                 end_time: Optional[builtins.str] = None,
                 last_day_of_month: Optional[builtins.bool] = None,
                 start_time: Optional[builtins.str] = None,
                 timezone: Optional[builtins.str] = None,
                 workdays_only: Optional[builtins.bool] = None):
        """
        :param builtins.str frequency: Frequency of the report. Allowed values: `never`, `once`, `hourly`, `daily`, `weekly`, `monthly`, `custom`.
        :param builtins.str custom_interval: Custom interval of the report.
               **Note:** This field is only available when frequency is set to `custom`.
        :param builtins.str end_time: End time of the report. If empty, the report will be sent indefinitely (according to frequency). Note that times will be saved as UTC in Grafana. Use 2006-01-02T15:04:05 format if you want to set a custom timezone
        :param builtins.bool last_day_of_month: Send the report on the last day of the month Defaults to `false`.
        :param builtins.str start_time: Start time of the report. If empty, the start date will be set to the creation time. Note that times will be saved as UTC in Grafana. Use 2006-01-02T15:04:05 format if you want to set a custom timezone
        :param builtins.str timezone: Set the report time zone. Defaults to `GMT`.
        :param builtins.bool workdays_only: Whether to send the report only on work days. Defaults to `false`.
        """
        pulumi.set(__self__, "frequency", frequency)
        if custom_interval is not None:
            pulumi.set(__self__, "custom_interval", custom_interval)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if last_day_of_month is not None:
            pulumi.set(__self__, "last_day_of_month", last_day_of_month)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if timezone is not None:
            pulumi.set(__self__, "timezone", timezone)
        if workdays_only is not None:
            pulumi.set(__self__, "workdays_only", workdays_only)

    @property
    @pulumi.getter
    def frequency(self) -> builtins.str:
        """
        Frequency of the report. Allowed values: `never`, `once`, `hourly`, `daily`, `weekly`, `monthly`, `custom`.
        """
        return pulumi.get(self, "frequency")

    @property
    @pulumi.getter(name="customInterval")
    def custom_interval(self) -> Optional[builtins.str]:
        """
        Custom interval of the report.
        **Note:** This field is only available when frequency is set to `custom`.
        """
        return pulumi.get(self, "custom_interval")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[builtins.str]:
        """
        End time of the report. If empty, the report will be sent indefinitely (according to frequency). Note that times will be saved as UTC in Grafana. Use 2006-01-02T15:04:05 format if you want to set a custom timezone
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="lastDayOfMonth")
    def last_day_of_month(self) -> Optional[builtins.bool]:
        """
        Send the report on the last day of the month Defaults to `false`.
        """
        return pulumi.get(self, "last_day_of_month")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[builtins.str]:
        """
        Start time of the report. If empty, the start date will be set to the creation time. Note that times will be saved as UTC in Grafana. Use 2006-01-02T15:04:05 format if you want to set a custom timezone
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def timezone(self) -> Optional[builtins.str]:
        """
        Set the report time zone. Defaults to `GMT`.
        """
        return pulumi.get(self, "timezone")

    @property
    @pulumi.getter(name="workdaysOnly")
    def workdays_only(self) -> Optional[builtins.bool]:
        """
        Whether to send the report only on work days. Defaults to `false`.
        """
        return pulumi.get(self, "workdays_only")


@pulumi.output_type
class RolePermission(dict):
    def __init__(__self__, *,
                 action: builtins.str,
                 scope: Optional[builtins.str] = None):
        """
        :param builtins.str action: Specific action users granted with the role will be allowed to perform (for example: `users:read`)
        :param builtins.str scope: Scope to restrict the action to a set of resources (for example: `users:*` or `roles:customrole1`) Defaults to ``.
        """
        pulumi.set(__self__, "action", action)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter
    def action(self) -> builtins.str:
        """
        Specific action users granted with the role will be allowed to perform (for example: `users:read`)
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def scope(self) -> Optional[builtins.str]:
        """
        Scope to restrict the action to a set of resources (for example: `users:*` or `roles:customrole1`) Defaults to ``.
        """
        return pulumi.get(self, "scope")


@pulumi.output_type
class GetRolePermissionResult(dict):
    def __init__(__self__, *,
                 action: builtins.str,
                 scope: Optional[builtins.str] = None):
        """
        :param builtins.str action: Specific action users granted with the role will be allowed to perform (for example: `users:read`)
        :param builtins.str scope: Scope to restrict the action to a set of resources (for example: `users:*` or `roles:customrole1`)
        """
        pulumi.set(__self__, "action", action)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter
    def action(self) -> builtins.str:
        """
        Specific action users granted with the role will be allowed to perform (for example: `users:read`)
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def scope(self) -> Optional[builtins.str]:
        """
        Scope to restrict the action to a set of resources (for example: `users:*` or `roles:customrole1`)
        """
        return pulumi.get(self, "scope")


