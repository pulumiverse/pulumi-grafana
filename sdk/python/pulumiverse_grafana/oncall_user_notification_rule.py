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

__all__ = ['OncallUserNotificationRuleArgs', 'OncallUserNotificationRule']

@pulumi.input_type
class OncallUserNotificationRuleArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 user_id: pulumi.Input[str],
                 duration: Optional[pulumi.Input[int]] = None,
                 important: Optional[pulumi.Input[bool]] = None,
                 position: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a OncallUserNotificationRule resource.
        :param pulumi.Input[str] type: The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notify_by_msteams` is only available for Grafana Cloud customers.
        :param pulumi.Input[str] user_id: User ID
        :param pulumi.Input[int] duration: A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
        :param pulumi.Input[bool] important: Boolean value which indicates if a rule is “important”
        :param pulumi.Input[int] position: Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "user_id", user_id)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if important is not None:
            pulumi.set(__self__, "important", important)
        if position is not None:
            pulumi.set(__self__, "position", position)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notify_by_msteams` is only available for Grafana Cloud customers.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[str]:
        """
        User ID
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[int]]:
        """
        A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter
    def important(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean value which indicates if a rule is “important”
        """
        return pulumi.get(self, "important")

    @important.setter
    def important(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "important", value)

    @property
    @pulumi.getter
    def position(self) -> Optional[pulumi.Input[int]]:
        """
        Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "position", value)


@pulumi.input_type
class _OncallUserNotificationRuleState:
    def __init__(__self__, *,
                 duration: Optional[pulumi.Input[int]] = None,
                 important: Optional[pulumi.Input[bool]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OncallUserNotificationRule resources.
        :param pulumi.Input[int] duration: A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
        :param pulumi.Input[bool] important: Boolean value which indicates if a rule is “important”
        :param pulumi.Input[int] position: Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
        :param pulumi.Input[str] type: The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notify_by_msteams` is only available for Grafana Cloud customers.
        :param pulumi.Input[str] user_id: User ID
        """
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if important is not None:
            pulumi.set(__self__, "important", important)
        if position is not None:
            pulumi.set(__self__, "position", position)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[int]]:
        """
        A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter
    def important(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean value which indicates if a rule is “important”
        """
        return pulumi.get(self, "important")

    @important.setter
    def important(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "important", value)

    @property
    @pulumi.getter
    def position(self) -> Optional[pulumi.Input[int]]:
        """
        Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notify_by_msteams` is only available for Grafana Cloud customers.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        User ID
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


warnings.warn("""grafana.index/oncallusernotificationrule.OncallUserNotificationRule has been deprecated in favor of grafana.oncall/usernotificationrule.UserNotificationRule""", DeprecationWarning)


class OncallUserNotificationRule(pulumi.CustomResource):
    warnings.warn("""grafana.index/oncallusernotificationrule.OncallUserNotificationRule has been deprecated in favor of grafana.oncall/usernotificationrule.UserNotificationRule""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 duration: Optional[pulumi.Input[int]] = None,
                 important: Optional[pulumi.Input[bool]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/personal_notification_rules/)

        **Note**: you must be running Grafana OnCall >= v1.8.0 to use this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_grafana as grafana
        import pulumiverse_grafana as grafana

        my_user = grafana.onCall.get_user(username="my_username")
        my_user_step1 = grafana.on_call.UserNotificationRule("my_user_step_1",
            user_id=my_user.id,
            position=0,
            type="notify_by_mobile_app")
        my_user_step2 = grafana.on_call.UserNotificationRule("my_user_step_2",
            user_id=my_user.id,
            position=1,
            duration=600,
            type="wait")
        my_user_step3 = grafana.on_call.UserNotificationRule("my_user_step_3",
            user_id=my_user.id,
            position=2,
            type="notify_by_phone_call")
        my_user_step4 = grafana.on_call.UserNotificationRule("my_user_step_4",
            user_id=my_user.id,
            position=3,
            duration=300,
            type="wait")
        my_user_step5 = grafana.on_call.UserNotificationRule("my_user_step_5",
            user_id=my_user.id,
            position=4,
            type="notify_by_slack")
        my_user_important_step1 = grafana.on_call.UserNotificationRule("my_user_important_step_1",
            user_id=my_user.id,
            important=True,
            position=0,
            type="notify_by_mobile_app_critical")
        my_user_important_step2 = grafana.on_call.UserNotificationRule("my_user_important_step_2",
            user_id=my_user.id,
            important=True,
            position=1,
            duration=300,
            type="wait")
        my_user_important_step3 = grafana.on_call.UserNotificationRule("my_user_important_step_3",
            user_id=my_user.id,
            important=True,
            position=2,
            type="notify_by_mobile_app_critical")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/oncallUserNotificationRule:OncallUserNotificationRule name "{{ id }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] duration: A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
        :param pulumi.Input[bool] important: Boolean value which indicates if a rule is “important”
        :param pulumi.Input[int] position: Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
        :param pulumi.Input[str] type: The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notify_by_msteams` is only available for Grafana Cloud customers.
        :param pulumi.Input[str] user_id: User ID
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OncallUserNotificationRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/personal_notification_rules/)

        **Note**: you must be running Grafana OnCall >= v1.8.0 to use this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_grafana as grafana
        import pulumiverse_grafana as grafana

        my_user = grafana.onCall.get_user(username="my_username")
        my_user_step1 = grafana.on_call.UserNotificationRule("my_user_step_1",
            user_id=my_user.id,
            position=0,
            type="notify_by_mobile_app")
        my_user_step2 = grafana.on_call.UserNotificationRule("my_user_step_2",
            user_id=my_user.id,
            position=1,
            duration=600,
            type="wait")
        my_user_step3 = grafana.on_call.UserNotificationRule("my_user_step_3",
            user_id=my_user.id,
            position=2,
            type="notify_by_phone_call")
        my_user_step4 = grafana.on_call.UserNotificationRule("my_user_step_4",
            user_id=my_user.id,
            position=3,
            duration=300,
            type="wait")
        my_user_step5 = grafana.on_call.UserNotificationRule("my_user_step_5",
            user_id=my_user.id,
            position=4,
            type="notify_by_slack")
        my_user_important_step1 = grafana.on_call.UserNotificationRule("my_user_important_step_1",
            user_id=my_user.id,
            important=True,
            position=0,
            type="notify_by_mobile_app_critical")
        my_user_important_step2 = grafana.on_call.UserNotificationRule("my_user_important_step_2",
            user_id=my_user.id,
            important=True,
            position=1,
            duration=300,
            type="wait")
        my_user_important_step3 = grafana.on_call.UserNotificationRule("my_user_important_step_3",
            user_id=my_user.id,
            important=True,
            position=2,
            type="notify_by_mobile_app_critical")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/oncallUserNotificationRule:OncallUserNotificationRule name "{{ id }}"
        ```

        :param str resource_name: The name of the resource.
        :param OncallUserNotificationRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OncallUserNotificationRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 duration: Optional[pulumi.Input[int]] = None,
                 important: Optional[pulumi.Input[bool]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""OncallUserNotificationRule is deprecated: grafana.index/oncallusernotificationrule.OncallUserNotificationRule has been deprecated in favor of grafana.oncall/usernotificationrule.UserNotificationRule""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OncallUserNotificationRuleArgs.__new__(OncallUserNotificationRuleArgs)

            __props__.__dict__["duration"] = duration
            __props__.__dict__["important"] = important
            __props__.__dict__["position"] = position
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
        super(OncallUserNotificationRule, __self__).__init__(
            'grafana:index/oncallUserNotificationRule:OncallUserNotificationRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            duration: Optional[pulumi.Input[int]] = None,
            important: Optional[pulumi.Input[bool]] = None,
            position: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'OncallUserNotificationRule':
        """
        Get an existing OncallUserNotificationRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] duration: A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
        :param pulumi.Input[bool] important: Boolean value which indicates if a rule is “important”
        :param pulumi.Input[int] position: Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
        :param pulumi.Input[str] type: The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notify_by_msteams` is only available for Grafana Cloud customers.
        :param pulumi.Input[str] user_id: User ID
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OncallUserNotificationRuleState.__new__(_OncallUserNotificationRuleState)

        __props__.__dict__["duration"] = duration
        __props__.__dict__["important"] = important
        __props__.__dict__["position"] = position
        __props__.__dict__["type"] = type
        __props__.__dict__["user_id"] = user_id
        return OncallUserNotificationRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Output[Optional[int]]:
        """
        A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter
    def important(self) -> pulumi.Output[bool]:
        """
        Boolean value which indicates if a rule is “important”
        """
        return pulumi.get(self, "important")

    @property
    @pulumi.getter
    def position(self) -> pulumi.Output[Optional[int]]:
        """
        Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
        """
        return pulumi.get(self, "position")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notify_by_msteams` is only available for Grafana Cloud customers.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        User ID
        """
        return pulumi.get(self, "user_id")

