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
from ._inputs import *

__all__ = ['RuleGroupArgs', 'RuleGroup']

@pulumi.input_type
class RuleGroupArgs:
    def __init__(__self__, *,
                 folder_uid: pulumi.Input[builtins.str],
                 interval_seconds: pulumi.Input[builtins.int],
                 rules: pulumi.Input[Sequence[pulumi.Input['RuleGroupRuleArgs']]],
                 disable_provenance: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 org_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a RuleGroup resource.
        :param pulumi.Input[builtins.str] folder_uid: The UID of the folder that the group belongs to.
        :param pulumi.Input[builtins.int] interval_seconds: The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are evaluated sequentially.
        :param pulumi.Input[Sequence[pulumi.Input['RuleGroupRuleArgs']]] rules: The rules within the group.
        :param pulumi.Input[builtins.str] name: The name of the rule group.
        :param pulumi.Input[builtins.str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        pulumi.set(__self__, "folder_uid", folder_uid)
        pulumi.set(__self__, "interval_seconds", interval_seconds)
        pulumi.set(__self__, "rules", rules)
        if disable_provenance is not None:
            pulumi.set(__self__, "disable_provenance", disable_provenance)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)

    @property
    @pulumi.getter(name="folderUid")
    def folder_uid(self) -> pulumi.Input[builtins.str]:
        """
        The UID of the folder that the group belongs to.
        """
        return pulumi.get(self, "folder_uid")

    @folder_uid.setter
    def folder_uid(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "folder_uid", value)

    @property
    @pulumi.getter(name="intervalSeconds")
    def interval_seconds(self) -> pulumi.Input[builtins.int]:
        """
        The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are evaluated sequentially.
        """
        return pulumi.get(self, "interval_seconds")

    @interval_seconds.setter
    def interval_seconds(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "interval_seconds", value)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Input[Sequence[pulumi.Input['RuleGroupRuleArgs']]]:
        """
        The rules within the group.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: pulumi.Input[Sequence[pulumi.Input['RuleGroupRuleArgs']]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter(name="disableProvenance")
    def disable_provenance(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "disable_provenance")

    @disable_provenance.setter
    def disable_provenance(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "disable_provenance", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the rule group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "org_id", value)


@pulumi.input_type
class _RuleGroupState:
    def __init__(__self__, *,
                 disable_provenance: Optional[pulumi.Input[builtins.bool]] = None,
                 folder_uid: Optional[pulumi.Input[builtins.str]] = None,
                 interval_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 org_id: Optional[pulumi.Input[builtins.str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['RuleGroupRuleArgs']]]] = None):
        """
        Input properties used for looking up and filtering RuleGroup resources.
        :param pulumi.Input[builtins.str] folder_uid: The UID of the folder that the group belongs to.
        :param pulumi.Input[builtins.int] interval_seconds: The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are evaluated sequentially.
        :param pulumi.Input[builtins.str] name: The name of the rule group.
        :param pulumi.Input[builtins.str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[Sequence[pulumi.Input['RuleGroupRuleArgs']]] rules: The rules within the group.
        """
        if disable_provenance is not None:
            pulumi.set(__self__, "disable_provenance", disable_provenance)
        if folder_uid is not None:
            pulumi.set(__self__, "folder_uid", folder_uid)
        if interval_seconds is not None:
            pulumi.set(__self__, "interval_seconds", interval_seconds)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter(name="disableProvenance")
    def disable_provenance(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "disable_provenance")

    @disable_provenance.setter
    def disable_provenance(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "disable_provenance", value)

    @property
    @pulumi.getter(name="folderUid")
    def folder_uid(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The UID of the folder that the group belongs to.
        """
        return pulumi.get(self, "folder_uid")

    @folder_uid.setter
    def folder_uid(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "folder_uid", value)

    @property
    @pulumi.getter(name="intervalSeconds")
    def interval_seconds(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are evaluated sequentially.
        """
        return pulumi.get(self, "interval_seconds")

    @interval_seconds.setter
    def interval_seconds(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "interval_seconds", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the rule group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RuleGroupRuleArgs']]]]:
        """
        The rules within the group.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RuleGroupRuleArgs']]]]):
        pulumi.set(self, "rules", value)


@pulumi.type_token("grafana:alerting/ruleGroup:RuleGroup")
class RuleGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_provenance: Optional[pulumi.Input[builtins.bool]] = None,
                 folder_uid: Optional[pulumi.Input[builtins.str]] = None,
                 interval_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 org_id: Optional[pulumi.Input[builtins.str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RuleGroupRuleArgs', 'RuleGroupRuleArgsDict']]]]] = None,
                 __props__=None):
        """
        Manages Grafana Alerting rule groups.

        * Official documentation
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/alerting_provisioning/#alert-rules)

        This resource requires Grafana 9.1.0 or later.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumiverse_grafana as grafana

        rule_folder = grafana.oss.Folder("rule_folder", title="My Alert Rule Folder")
        my_alert_rule = grafana.alerting.RuleGroup("my_alert_rule",
            name="My Rule Group",
            folder_uid=rule_folder.uid,
            interval_seconds=240,
            org_id="1",
            rules=[{
                "name": "My Alert Rule 1",
                "for_": "2m",
                "condition": "B",
                "no_data_state": "NoData",
                "exec_err_state": "Alerting",
                "annotations": {
                    "a": "b",
                    "c": "d",
                },
                "labels": {
                    "e": "f",
                    "g": "h",
                },
                "is_paused": False,
                "datas": [
                    {
                        "ref_id": "A",
                        "query_type": "",
                        "relative_time_range": {
                            "from_": 600,
                            "to": 0,
                        },
                        "datasource_uid": "PD8C576611E62080A",
                        "model": json.dumps({
                            "hide": False,
                            "intervalMs": 1000,
                            "maxDataPoints": 43200,
                            "refId": "A",
                        }),
                    },
                    {
                        "ref_id": "B",
                        "query_type": "",
                        "relative_time_range": {
                            "from_": 0,
                            "to": 0,
                        },
                        "datasource_uid": "-100",
                        "model": \"\"\"{
            "conditions": [
                {
                "evaluator": {
                    "params": [
                    3
                    ],
                    "type": "gt"
                },
                "operator": {
                    "type": "and"
                },
                "query": {
                    "params": [
                    "A"
                    ]
                },
                "reducer": {
                    "params": [],
                    "type": "last"
                },
                "type": "query"
                }
            ],
            "datasource": {
                "type": "__expr__",
                "uid": "-100"
            },
            "hide": false,
            "intervalMs": 1000,
            "maxDataPoints": 43200,
            "refId": "B",
            "type": "classic_conditions"
        }
        \"\"\",
                    },
                ],
            }])
        ```

        ## Import

        ```sh
        $ pulumi import grafana:alerting/ruleGroup:RuleGroup name "{{ folderUID }}:{{ title }}"
        ```

        ```sh
        $ pulumi import grafana:alerting/ruleGroup:RuleGroup name "{{ orgID }}:{{ folderUID }}:{{ title }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] folder_uid: The UID of the folder that the group belongs to.
        :param pulumi.Input[builtins.int] interval_seconds: The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are evaluated sequentially.
        :param pulumi.Input[builtins.str] name: The name of the rule group.
        :param pulumi.Input[builtins.str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RuleGroupRuleArgs', 'RuleGroupRuleArgsDict']]]] rules: The rules within the group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RuleGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Grafana Alerting rule groups.

        * Official documentation
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/alerting_provisioning/#alert-rules)

        This resource requires Grafana 9.1.0 or later.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumiverse_grafana as grafana

        rule_folder = grafana.oss.Folder("rule_folder", title="My Alert Rule Folder")
        my_alert_rule = grafana.alerting.RuleGroup("my_alert_rule",
            name="My Rule Group",
            folder_uid=rule_folder.uid,
            interval_seconds=240,
            org_id="1",
            rules=[{
                "name": "My Alert Rule 1",
                "for_": "2m",
                "condition": "B",
                "no_data_state": "NoData",
                "exec_err_state": "Alerting",
                "annotations": {
                    "a": "b",
                    "c": "d",
                },
                "labels": {
                    "e": "f",
                    "g": "h",
                },
                "is_paused": False,
                "datas": [
                    {
                        "ref_id": "A",
                        "query_type": "",
                        "relative_time_range": {
                            "from_": 600,
                            "to": 0,
                        },
                        "datasource_uid": "PD8C576611E62080A",
                        "model": json.dumps({
                            "hide": False,
                            "intervalMs": 1000,
                            "maxDataPoints": 43200,
                            "refId": "A",
                        }),
                    },
                    {
                        "ref_id": "B",
                        "query_type": "",
                        "relative_time_range": {
                            "from_": 0,
                            "to": 0,
                        },
                        "datasource_uid": "-100",
                        "model": \"\"\"{
            "conditions": [
                {
                "evaluator": {
                    "params": [
                    3
                    ],
                    "type": "gt"
                },
                "operator": {
                    "type": "and"
                },
                "query": {
                    "params": [
                    "A"
                    ]
                },
                "reducer": {
                    "params": [],
                    "type": "last"
                },
                "type": "query"
                }
            ],
            "datasource": {
                "type": "__expr__",
                "uid": "-100"
            },
            "hide": false,
            "intervalMs": 1000,
            "maxDataPoints": 43200,
            "refId": "B",
            "type": "classic_conditions"
        }
        \"\"\",
                    },
                ],
            }])
        ```

        ## Import

        ```sh
        $ pulumi import grafana:alerting/ruleGroup:RuleGroup name "{{ folderUID }}:{{ title }}"
        ```

        ```sh
        $ pulumi import grafana:alerting/ruleGroup:RuleGroup name "{{ orgID }}:{{ folderUID }}:{{ title }}"
        ```

        :param str resource_name: The name of the resource.
        :param RuleGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RuleGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_provenance: Optional[pulumi.Input[builtins.bool]] = None,
                 folder_uid: Optional[pulumi.Input[builtins.str]] = None,
                 interval_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 org_id: Optional[pulumi.Input[builtins.str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RuleGroupRuleArgs', 'RuleGroupRuleArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RuleGroupArgs.__new__(RuleGroupArgs)

            __props__.__dict__["disable_provenance"] = disable_provenance
            if folder_uid is None and not opts.urn:
                raise TypeError("Missing required property 'folder_uid'")
            __props__.__dict__["folder_uid"] = folder_uid
            if interval_seconds is None and not opts.urn:
                raise TypeError("Missing required property 'interval_seconds'")
            __props__.__dict__["interval_seconds"] = interval_seconds
            __props__.__dict__["name"] = name
            __props__.__dict__["org_id"] = org_id
            if rules is None and not opts.urn:
                raise TypeError("Missing required property 'rules'")
            __props__.__dict__["rules"] = rules
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:index/ruleGroup:RuleGroup")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(RuleGroup, __self__).__init__(
            'grafana:alerting/ruleGroup:RuleGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disable_provenance: Optional[pulumi.Input[builtins.bool]] = None,
            folder_uid: Optional[pulumi.Input[builtins.str]] = None,
            interval_seconds: Optional[pulumi.Input[builtins.int]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            org_id: Optional[pulumi.Input[builtins.str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RuleGroupRuleArgs', 'RuleGroupRuleArgsDict']]]]] = None) -> 'RuleGroup':
        """
        Get an existing RuleGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] folder_uid: The UID of the folder that the group belongs to.
        :param pulumi.Input[builtins.int] interval_seconds: The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are evaluated sequentially.
        :param pulumi.Input[builtins.str] name: The name of the rule group.
        :param pulumi.Input[builtins.str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RuleGroupRuleArgs', 'RuleGroupRuleArgsDict']]]] rules: The rules within the group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RuleGroupState.__new__(_RuleGroupState)

        __props__.__dict__["disable_provenance"] = disable_provenance
        __props__.__dict__["folder_uid"] = folder_uid
        __props__.__dict__["interval_seconds"] = interval_seconds
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["rules"] = rules
        return RuleGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="disableProvenance")
    def disable_provenance(self) -> pulumi.Output[Optional[builtins.bool]]:
        return pulumi.get(self, "disable_provenance")

    @property
    @pulumi.getter(name="folderUid")
    def folder_uid(self) -> pulumi.Output[builtins.str]:
        """
        The UID of the folder that the group belongs to.
        """
        return pulumi.get(self, "folder_uid")

    @property
    @pulumi.getter(name="intervalSeconds")
    def interval_seconds(self) -> pulumi.Output[builtins.int]:
        """
        The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are evaluated sequentially.
        """
        return pulumi.get(self, "interval_seconds")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the rule group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.RuleGroupRule']]:
        """
        The rules within the group.
        """
        return pulumi.get(self, "rules")

