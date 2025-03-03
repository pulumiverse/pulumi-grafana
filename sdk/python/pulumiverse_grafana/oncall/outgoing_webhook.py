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

__all__ = ['OutgoingWebhookArgs', 'OutgoingWebhook']

@pulumi.input_type
class OutgoingWebhookArgs:
    def __init__(__self__, *,
                 url: pulumi.Input[str],
                 authorization_header: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 forward_whole_payload: Optional[pulumi.Input[bool]] = None,
                 headers: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 integration_filters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 is_webhook_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 trigger_template: Optional[pulumi.Input[str]] = None,
                 trigger_type: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OutgoingWebhook resource.
        :param pulumi.Input[str] url: The webhook URL.
        :param pulumi.Input[str] authorization_header: The auth data of the webhook. Used in Authorization header instead of user/password auth.
        :param pulumi.Input[str] data: The data of the webhook.
        :param pulumi.Input[bool] forward_whole_payload: Toggle to send the entire webhook payload instead of using the values in the Data field.
        :param pulumi.Input[str] headers: Headers to add to the outgoing webhook request.
        :param pulumi.Input[str] http_method: The HTTP method used in the request made by the outgoing webhook. Defaults to `POST`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] integration_filters: Restricts the outgoing webhook to only trigger if the event came from a selected integration. If no integrations are selected the outgoing webhook will trigger for any integration.
        :param pulumi.Input[bool] is_webhook_enabled: Controls whether the outgoing webhook will trigger or is ignored. Defaults to `true`.
        :param pulumi.Input[str] name: The name of the outgoing webhook.
        :param pulumi.Input[str] password: The auth data of the webhook. Used for Basic authentication
        :param pulumi.Input[str] team_id: The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `on_call_get_team` datasource.
        :param pulumi.Input[str] trigger_template: A template used to dynamically determine whether the webhook should execute based on the content of the payload.
        :param pulumi.Input[str] trigger_type: The type of event that will cause this outgoing webhook to execute. The types of triggers are: `escalation`, `alert group created`, `acknowledge`, `resolve`, `silence`, `unsilence`, `unresolve`, `unacknowledge`. Defaults to `escalation`.
        :param pulumi.Input[str] user: Username to use when making the outgoing webhook request.
        """
        pulumi.set(__self__, "url", url)
        if authorization_header is not None:
            pulumi.set(__self__, "authorization_header", authorization_header)
        if data is not None:
            pulumi.set(__self__, "data", data)
        if forward_whole_payload is not None:
            pulumi.set(__self__, "forward_whole_payload", forward_whole_payload)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if http_method is not None:
            pulumi.set(__self__, "http_method", http_method)
        if integration_filters is not None:
            pulumi.set(__self__, "integration_filters", integration_filters)
        if is_webhook_enabled is not None:
            pulumi.set(__self__, "is_webhook_enabled", is_webhook_enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)
        if trigger_template is not None:
            pulumi.set(__self__, "trigger_template", trigger_template)
        if trigger_type is not None:
            pulumi.set(__self__, "trigger_type", trigger_type)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        """
        The webhook URL.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="authorizationHeader")
    def authorization_header(self) -> Optional[pulumi.Input[str]]:
        """
        The auth data of the webhook. Used in Authorization header instead of user/password auth.
        """
        return pulumi.get(self, "authorization_header")

    @authorization_header.setter
    def authorization_header(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization_header", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        """
        The data of the webhook.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="forwardWholePayload")
    def forward_whole_payload(self) -> Optional[pulumi.Input[bool]]:
        """
        Toggle to send the entire webhook payload instead of using the values in the Data field.
        """
        return pulumi.get(self, "forward_whole_payload")

    @forward_whole_payload.setter
    def forward_whole_payload(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "forward_whole_payload", value)

    @property
    @pulumi.getter
    def headers(self) -> Optional[pulumi.Input[str]]:
        """
        Headers to add to the outgoing webhook request.
        """
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "headers", value)

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> Optional[pulumi.Input[str]]:
        """
        The HTTP method used in the request made by the outgoing webhook. Defaults to `POST`.
        """
        return pulumi.get(self, "http_method")

    @http_method.setter
    def http_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_method", value)

    @property
    @pulumi.getter(name="integrationFilters")
    def integration_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Restricts the outgoing webhook to only trigger if the event came from a selected integration. If no integrations are selected the outgoing webhook will trigger for any integration.
        """
        return pulumi.get(self, "integration_filters")

    @integration_filters.setter
    def integration_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "integration_filters", value)

    @property
    @pulumi.getter(name="isWebhookEnabled")
    def is_webhook_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Controls whether the outgoing webhook will trigger or is ignored. Defaults to `true`.
        """
        return pulumi.get(self, "is_webhook_enabled")

    @is_webhook_enabled.setter
    def is_webhook_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_webhook_enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the outgoing webhook.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The auth data of the webhook. Used for Basic authentication
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `on_call_get_team` datasource.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)

    @property
    @pulumi.getter(name="triggerTemplate")
    def trigger_template(self) -> Optional[pulumi.Input[str]]:
        """
        A template used to dynamically determine whether the webhook should execute based on the content of the payload.
        """
        return pulumi.get(self, "trigger_template")

    @trigger_template.setter
    def trigger_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trigger_template", value)

    @property
    @pulumi.getter(name="triggerType")
    def trigger_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of event that will cause this outgoing webhook to execute. The types of triggers are: `escalation`, `alert group created`, `acknowledge`, `resolve`, `silence`, `unsilence`, `unresolve`, `unacknowledge`. Defaults to `escalation`.
        """
        return pulumi.get(self, "trigger_type")

    @trigger_type.setter
    def trigger_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trigger_type", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        Username to use when making the outgoing webhook request.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


@pulumi.input_type
class _OutgoingWebhookState:
    def __init__(__self__, *,
                 authorization_header: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 forward_whole_payload: Optional[pulumi.Input[bool]] = None,
                 headers: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 integration_filters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 is_webhook_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 trigger_template: Optional[pulumi.Input[str]] = None,
                 trigger_type: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OutgoingWebhook resources.
        :param pulumi.Input[str] authorization_header: The auth data of the webhook. Used in Authorization header instead of user/password auth.
        :param pulumi.Input[str] data: The data of the webhook.
        :param pulumi.Input[bool] forward_whole_payload: Toggle to send the entire webhook payload instead of using the values in the Data field.
        :param pulumi.Input[str] headers: Headers to add to the outgoing webhook request.
        :param pulumi.Input[str] http_method: The HTTP method used in the request made by the outgoing webhook. Defaults to `POST`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] integration_filters: Restricts the outgoing webhook to only trigger if the event came from a selected integration. If no integrations are selected the outgoing webhook will trigger for any integration.
        :param pulumi.Input[bool] is_webhook_enabled: Controls whether the outgoing webhook will trigger or is ignored. Defaults to `true`.
        :param pulumi.Input[str] name: The name of the outgoing webhook.
        :param pulumi.Input[str] password: The auth data of the webhook. Used for Basic authentication
        :param pulumi.Input[str] team_id: The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `on_call_get_team` datasource.
        :param pulumi.Input[str] trigger_template: A template used to dynamically determine whether the webhook should execute based on the content of the payload.
        :param pulumi.Input[str] trigger_type: The type of event that will cause this outgoing webhook to execute. The types of triggers are: `escalation`, `alert group created`, `acknowledge`, `resolve`, `silence`, `unsilence`, `unresolve`, `unacknowledge`. Defaults to `escalation`.
        :param pulumi.Input[str] url: The webhook URL.
        :param pulumi.Input[str] user: Username to use when making the outgoing webhook request.
        """
        if authorization_header is not None:
            pulumi.set(__self__, "authorization_header", authorization_header)
        if data is not None:
            pulumi.set(__self__, "data", data)
        if forward_whole_payload is not None:
            pulumi.set(__self__, "forward_whole_payload", forward_whole_payload)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if http_method is not None:
            pulumi.set(__self__, "http_method", http_method)
        if integration_filters is not None:
            pulumi.set(__self__, "integration_filters", integration_filters)
        if is_webhook_enabled is not None:
            pulumi.set(__self__, "is_webhook_enabled", is_webhook_enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)
        if trigger_template is not None:
            pulumi.set(__self__, "trigger_template", trigger_template)
        if trigger_type is not None:
            pulumi.set(__self__, "trigger_type", trigger_type)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter(name="authorizationHeader")
    def authorization_header(self) -> Optional[pulumi.Input[str]]:
        """
        The auth data of the webhook. Used in Authorization header instead of user/password auth.
        """
        return pulumi.get(self, "authorization_header")

    @authorization_header.setter
    def authorization_header(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization_header", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        """
        The data of the webhook.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="forwardWholePayload")
    def forward_whole_payload(self) -> Optional[pulumi.Input[bool]]:
        """
        Toggle to send the entire webhook payload instead of using the values in the Data field.
        """
        return pulumi.get(self, "forward_whole_payload")

    @forward_whole_payload.setter
    def forward_whole_payload(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "forward_whole_payload", value)

    @property
    @pulumi.getter
    def headers(self) -> Optional[pulumi.Input[str]]:
        """
        Headers to add to the outgoing webhook request.
        """
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "headers", value)

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> Optional[pulumi.Input[str]]:
        """
        The HTTP method used in the request made by the outgoing webhook. Defaults to `POST`.
        """
        return pulumi.get(self, "http_method")

    @http_method.setter
    def http_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_method", value)

    @property
    @pulumi.getter(name="integrationFilters")
    def integration_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Restricts the outgoing webhook to only trigger if the event came from a selected integration. If no integrations are selected the outgoing webhook will trigger for any integration.
        """
        return pulumi.get(self, "integration_filters")

    @integration_filters.setter
    def integration_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "integration_filters", value)

    @property
    @pulumi.getter(name="isWebhookEnabled")
    def is_webhook_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Controls whether the outgoing webhook will trigger or is ignored. Defaults to `true`.
        """
        return pulumi.get(self, "is_webhook_enabled")

    @is_webhook_enabled.setter
    def is_webhook_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_webhook_enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the outgoing webhook.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The auth data of the webhook. Used for Basic authentication
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `on_call_get_team` datasource.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)

    @property
    @pulumi.getter(name="triggerTemplate")
    def trigger_template(self) -> Optional[pulumi.Input[str]]:
        """
        A template used to dynamically determine whether the webhook should execute based on the content of the payload.
        """
        return pulumi.get(self, "trigger_template")

    @trigger_template.setter
    def trigger_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trigger_template", value)

    @property
    @pulumi.getter(name="triggerType")
    def trigger_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of event that will cause this outgoing webhook to execute. The types of triggers are: `escalation`, `alert group created`, `acknowledge`, `resolve`, `silence`, `unsilence`, `unresolve`, `unacknowledge`. Defaults to `escalation`.
        """
        return pulumi.get(self, "trigger_type")

    @trigger_type.setter
    def trigger_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trigger_type", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The webhook URL.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        Username to use when making the outgoing webhook request.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


class OutgoingWebhook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization_header: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 forward_whole_payload: Optional[pulumi.Input[bool]] = None,
                 headers: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 integration_filters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 is_webhook_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 trigger_template: Optional[pulumi.Input[str]] = None,
                 trigger_type: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/outgoing_webhooks/)

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        test_acc_outgoing_webhook = grafana.on_call.OutgoingWebhook("test-acc-outgoing_webhook",
            name="my outgoing webhook",
            url="https://example.com/")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:onCall/outgoingWebhook:OutgoingWebhook name "{{ id }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorization_header: The auth data of the webhook. Used in Authorization header instead of user/password auth.
        :param pulumi.Input[str] data: The data of the webhook.
        :param pulumi.Input[bool] forward_whole_payload: Toggle to send the entire webhook payload instead of using the values in the Data field.
        :param pulumi.Input[str] headers: Headers to add to the outgoing webhook request.
        :param pulumi.Input[str] http_method: The HTTP method used in the request made by the outgoing webhook. Defaults to `POST`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] integration_filters: Restricts the outgoing webhook to only trigger if the event came from a selected integration. If no integrations are selected the outgoing webhook will trigger for any integration.
        :param pulumi.Input[bool] is_webhook_enabled: Controls whether the outgoing webhook will trigger or is ignored. Defaults to `true`.
        :param pulumi.Input[str] name: The name of the outgoing webhook.
        :param pulumi.Input[str] password: The auth data of the webhook. Used for Basic authentication
        :param pulumi.Input[str] team_id: The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `on_call_get_team` datasource.
        :param pulumi.Input[str] trigger_template: A template used to dynamically determine whether the webhook should execute based on the content of the payload.
        :param pulumi.Input[str] trigger_type: The type of event that will cause this outgoing webhook to execute. The types of triggers are: `escalation`, `alert group created`, `acknowledge`, `resolve`, `silence`, `unsilence`, `unresolve`, `unacknowledge`. Defaults to `escalation`.
        :param pulumi.Input[str] url: The webhook URL.
        :param pulumi.Input[str] user: Username to use when making the outgoing webhook request.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OutgoingWebhookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/outgoing_webhooks/)

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        test_acc_outgoing_webhook = grafana.on_call.OutgoingWebhook("test-acc-outgoing_webhook",
            name="my outgoing webhook",
            url="https://example.com/")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:onCall/outgoingWebhook:OutgoingWebhook name "{{ id }}"
        ```

        :param str resource_name: The name of the resource.
        :param OutgoingWebhookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OutgoingWebhookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization_header: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 forward_whole_payload: Optional[pulumi.Input[bool]] = None,
                 headers: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 integration_filters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 is_webhook_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 trigger_template: Optional[pulumi.Input[str]] = None,
                 trigger_type: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OutgoingWebhookArgs.__new__(OutgoingWebhookArgs)

            __props__.__dict__["authorization_header"] = None if authorization_header is None else pulumi.Output.secret(authorization_header)
            __props__.__dict__["data"] = data
            __props__.__dict__["forward_whole_payload"] = forward_whole_payload
            __props__.__dict__["headers"] = headers
            __props__.__dict__["http_method"] = http_method
            __props__.__dict__["integration_filters"] = integration_filters
            __props__.__dict__["is_webhook_enabled"] = is_webhook_enabled
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["team_id"] = team_id
            __props__.__dict__["trigger_template"] = trigger_template
            __props__.__dict__["trigger_type"] = trigger_type
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            __props__.__dict__["user"] = user
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["authorizationHeader", "password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(OutgoingWebhook, __self__).__init__(
            'grafana:onCall/outgoingWebhook:OutgoingWebhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization_header: Optional[pulumi.Input[str]] = None,
            data: Optional[pulumi.Input[str]] = None,
            forward_whole_payload: Optional[pulumi.Input[bool]] = None,
            headers: Optional[pulumi.Input[str]] = None,
            http_method: Optional[pulumi.Input[str]] = None,
            integration_filters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            is_webhook_enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            team_id: Optional[pulumi.Input[str]] = None,
            trigger_template: Optional[pulumi.Input[str]] = None,
            trigger_type: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            user: Optional[pulumi.Input[str]] = None) -> 'OutgoingWebhook':
        """
        Get an existing OutgoingWebhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorization_header: The auth data of the webhook. Used in Authorization header instead of user/password auth.
        :param pulumi.Input[str] data: The data of the webhook.
        :param pulumi.Input[bool] forward_whole_payload: Toggle to send the entire webhook payload instead of using the values in the Data field.
        :param pulumi.Input[str] headers: Headers to add to the outgoing webhook request.
        :param pulumi.Input[str] http_method: The HTTP method used in the request made by the outgoing webhook. Defaults to `POST`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] integration_filters: Restricts the outgoing webhook to only trigger if the event came from a selected integration. If no integrations are selected the outgoing webhook will trigger for any integration.
        :param pulumi.Input[bool] is_webhook_enabled: Controls whether the outgoing webhook will trigger or is ignored. Defaults to `true`.
        :param pulumi.Input[str] name: The name of the outgoing webhook.
        :param pulumi.Input[str] password: The auth data of the webhook. Used for Basic authentication
        :param pulumi.Input[str] team_id: The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `on_call_get_team` datasource.
        :param pulumi.Input[str] trigger_template: A template used to dynamically determine whether the webhook should execute based on the content of the payload.
        :param pulumi.Input[str] trigger_type: The type of event that will cause this outgoing webhook to execute. The types of triggers are: `escalation`, `alert group created`, `acknowledge`, `resolve`, `silence`, `unsilence`, `unresolve`, `unacknowledge`. Defaults to `escalation`.
        :param pulumi.Input[str] url: The webhook URL.
        :param pulumi.Input[str] user: Username to use when making the outgoing webhook request.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OutgoingWebhookState.__new__(_OutgoingWebhookState)

        __props__.__dict__["authorization_header"] = authorization_header
        __props__.__dict__["data"] = data
        __props__.__dict__["forward_whole_payload"] = forward_whole_payload
        __props__.__dict__["headers"] = headers
        __props__.__dict__["http_method"] = http_method
        __props__.__dict__["integration_filters"] = integration_filters
        __props__.__dict__["is_webhook_enabled"] = is_webhook_enabled
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["team_id"] = team_id
        __props__.__dict__["trigger_template"] = trigger_template
        __props__.__dict__["trigger_type"] = trigger_type
        __props__.__dict__["url"] = url
        __props__.__dict__["user"] = user
        return OutgoingWebhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authorizationHeader")
    def authorization_header(self) -> pulumi.Output[Optional[str]]:
        """
        The auth data of the webhook. Used in Authorization header instead of user/password auth.
        """
        return pulumi.get(self, "authorization_header")

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[Optional[str]]:
        """
        The data of the webhook.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="forwardWholePayload")
    def forward_whole_payload(self) -> pulumi.Output[Optional[bool]]:
        """
        Toggle to send the entire webhook payload instead of using the values in the Data field.
        """
        return pulumi.get(self, "forward_whole_payload")

    @property
    @pulumi.getter
    def headers(self) -> pulumi.Output[Optional[str]]:
        """
        Headers to add to the outgoing webhook request.
        """
        return pulumi.get(self, "headers")

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> pulumi.Output[Optional[str]]:
        """
        The HTTP method used in the request made by the outgoing webhook. Defaults to `POST`.
        """
        return pulumi.get(self, "http_method")

    @property
    @pulumi.getter(name="integrationFilters")
    def integration_filters(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Restricts the outgoing webhook to only trigger if the event came from a selected integration. If no integrations are selected the outgoing webhook will trigger for any integration.
        """
        return pulumi.get(self, "integration_filters")

    @property
    @pulumi.getter(name="isWebhookEnabled")
    def is_webhook_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Controls whether the outgoing webhook will trigger or is ignored. Defaults to `true`.
        """
        return pulumi.get(self, "is_webhook_enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the outgoing webhook.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The auth data of the webhook. Used for Basic authentication
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `on_call_get_team` datasource.
        """
        return pulumi.get(self, "team_id")

    @property
    @pulumi.getter(name="triggerTemplate")
    def trigger_template(self) -> pulumi.Output[Optional[str]]:
        """
        A template used to dynamically determine whether the webhook should execute based on the content of the payload.
        """
        return pulumi.get(self, "trigger_template")

    @property
    @pulumi.getter(name="triggerType")
    def trigger_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of event that will cause this outgoing webhook to execute. The types of triggers are: `escalation`, `alert group created`, `acknowledge`, `resolve`, `silence`, `unsilence`, `unresolve`, `unacknowledge`. Defaults to `escalation`.
        """
        return pulumi.get(self, "trigger_type")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The webhook URL.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[Optional[str]]:
        """
        Username to use when making the outgoing webhook request.
        """
        return pulumi.get(self, "user")

