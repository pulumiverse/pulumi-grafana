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
from .. import _utilities

__all__ = ['MessageTemplateArgs', 'MessageTemplate']

@pulumi.input_type
class MessageTemplateArgs:
    def __init__(__self__, *,
                 template: pulumi.Input[str],
                 disable_provenance: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MessageTemplate resource.
        :param pulumi.Input[str] template: The content of the notification template group.
        :param pulumi.Input[str] name: The name of the notification template group.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        pulumi.set(__self__, "template", template)
        if disable_provenance is not None:
            pulumi.set(__self__, "disable_provenance", disable_provenance)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)

    @property
    @pulumi.getter
    def template(self) -> pulumi.Input[str]:
        """
        The content of the notification template group.
        """
        return pulumi.get(self, "template")

    @template.setter
    def template(self, value: pulumi.Input[str]):
        pulumi.set(self, "template", value)

    @property
    @pulumi.getter(name="disableProvenance")
    def disable_provenance(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "disable_provenance")

    @disable_provenance.setter
    def disable_provenance(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_provenance", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the notification template group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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


@pulumi.input_type
class _MessageTemplateState:
    def __init__(__self__, *,
                 disable_provenance: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MessageTemplate resources.
        :param pulumi.Input[str] name: The name of the notification template group.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[str] template: The content of the notification template group.
        """
        if disable_provenance is not None:
            pulumi.set(__self__, "disable_provenance", disable_provenance)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if template is not None:
            pulumi.set(__self__, "template", template)

    @property
    @pulumi.getter(name="disableProvenance")
    def disable_provenance(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "disable_provenance")

    @disable_provenance.setter
    def disable_provenance(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_provenance", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the notification template group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
    def template(self) -> Optional[pulumi.Input[str]]:
        """
        The content of the notification template group.
        """
        return pulumi.get(self, "template")

    @template.setter
    def template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template", value)


class MessageTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_provenance: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages Grafana Alerting notification template groups, including notification templates.

        * Official documentation
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/alerting_provisioning/#notification-template-groups)

        This resource requires Grafana 9.1.0 or later.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        my_template = grafana.alerting.MessageTemplate("my_template",
            name="My Notification Template Group",
            template=\"\"\"{{define "custom.message" }}
         template content
        {{ end }}\"\"\")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:alerting/messageTemplate:MessageTemplate name "{{ name }}"
        ```

        ```sh
        $ pulumi import grafana:alerting/messageTemplate:MessageTemplate name "{{ orgID }}:{{ name }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the notification template group.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[str] template: The content of the notification template group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MessageTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Grafana Alerting notification template groups, including notification templates.

        * Official documentation
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/alerting_provisioning/#notification-template-groups)

        This resource requires Grafana 9.1.0 or later.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        my_template = grafana.alerting.MessageTemplate("my_template",
            name="My Notification Template Group",
            template=\"\"\"{{define "custom.message" }}
         template content
        {{ end }}\"\"\")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:alerting/messageTemplate:MessageTemplate name "{{ name }}"
        ```

        ```sh
        $ pulumi import grafana:alerting/messageTemplate:MessageTemplate name "{{ orgID }}:{{ name }}"
        ```

        :param str resource_name: The name of the resource.
        :param MessageTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MessageTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_provenance: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MessageTemplateArgs.__new__(MessageTemplateArgs)

            __props__.__dict__["disable_provenance"] = disable_provenance
            __props__.__dict__["name"] = name
            __props__.__dict__["org_id"] = org_id
            if template is None and not opts.urn:
                raise TypeError("Missing required property 'template'")
            __props__.__dict__["template"] = template
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:index/messageTemplate:MessageTemplate")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(MessageTemplate, __self__).__init__(
            'grafana:alerting/messageTemplate:MessageTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disable_provenance: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            template: Optional[pulumi.Input[str]] = None) -> 'MessageTemplate':
        """
        Get an existing MessageTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the notification template group.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[str] template: The content of the notification template group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MessageTemplateState.__new__(_MessageTemplateState)

        __props__.__dict__["disable_provenance"] = disable_provenance
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["template"] = template
        return MessageTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="disableProvenance")
    def disable_provenance(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "disable_provenance")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the notification template group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def template(self) -> pulumi.Output[str]:
        """
        The content of the notification template group.
        """
        return pulumi.get(self, "template")

