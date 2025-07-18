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

__all__ = ['ProviderAzureCredentialArgs', 'ProviderAzureCredential']

@pulumi.input_type
class ProviderAzureCredentialArgs:
    def __init__(__self__, *,
                 client_id: pulumi.Input[builtins.str],
                 client_secret: pulumi.Input[builtins.str],
                 stack_id: pulumi.Input[builtins.str],
                 tenant_id: pulumi.Input[builtins.str],
                 auto_discovery_configurations: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderAzureCredentialAutoDiscoveryConfigurationArgs']]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_discovery_tag_filters: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderAzureCredentialResourceDiscoveryTagFilterArgs']]]] = None,
                 resource_tags_to_add_to_metrics: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a ProviderAzureCredential resource.
        :param pulumi.Input[builtins.str] client_id: The client ID of the Azure Credential.
        :param pulumi.Input[builtins.str] client_secret: The client secret of the Azure Credential.
        :param pulumi.Input[builtins.str] tenant_id: The tenant ID of the Azure Credential.
        :param pulumi.Input[Sequence[pulumi.Input['ProviderAzureCredentialAutoDiscoveryConfigurationArgs']]] auto_discovery_configurations: The list of auto discovery configurations.
        :param pulumi.Input[builtins.str] name: The name of the Azure Credential.
        :param pulumi.Input[Sequence[pulumi.Input['ProviderAzureCredentialResourceDiscoveryTagFilterArgs']]] resource_discovery_tag_filters: The list of tag filters to apply to resources.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] resource_tags_to_add_to_metrics: The list of resource tags to add to metrics.
        """
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "client_secret", client_secret)
        pulumi.set(__self__, "stack_id", stack_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        if auto_discovery_configurations is not None:
            pulumi.set(__self__, "auto_discovery_configurations", auto_discovery_configurations)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_discovery_tag_filters is not None:
            pulumi.set(__self__, "resource_discovery_tag_filters", resource_discovery_tag_filters)
        if resource_tags_to_add_to_metrics is not None:
            pulumi.set(__self__, "resource_tags_to_add_to_metrics", resource_tags_to_add_to_metrics)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Input[builtins.str]:
        """
        The client ID of the Azure Credential.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Input[builtins.str]:
        """
        The client secret of the Azure Credential.
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "stack_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[builtins.str]:
        """
        The tenant ID of the Azure Credential.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="autoDiscoveryConfigurations")
    def auto_discovery_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProviderAzureCredentialAutoDiscoveryConfigurationArgs']]]]:
        """
        The list of auto discovery configurations.
        """
        return pulumi.get(self, "auto_discovery_configurations")

    @auto_discovery_configurations.setter
    def auto_discovery_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderAzureCredentialAutoDiscoveryConfigurationArgs']]]]):
        pulumi.set(self, "auto_discovery_configurations", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the Azure Credential.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceDiscoveryTagFilters")
    def resource_discovery_tag_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProviderAzureCredentialResourceDiscoveryTagFilterArgs']]]]:
        """
        The list of tag filters to apply to resources.
        """
        return pulumi.get(self, "resource_discovery_tag_filters")

    @resource_discovery_tag_filters.setter
    def resource_discovery_tag_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderAzureCredentialResourceDiscoveryTagFilterArgs']]]]):
        pulumi.set(self, "resource_discovery_tag_filters", value)

    @property
    @pulumi.getter(name="resourceTagsToAddToMetrics")
    def resource_tags_to_add_to_metrics(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The list of resource tags to add to metrics.
        """
        return pulumi.get(self, "resource_tags_to_add_to_metrics")

    @resource_tags_to_add_to_metrics.setter
    def resource_tags_to_add_to_metrics(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "resource_tags_to_add_to_metrics", value)


@pulumi.input_type
class _ProviderAzureCredentialState:
    def __init__(__self__, *,
                 auto_discovery_configurations: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderAzureCredentialAutoDiscoveryConfigurationArgs']]]] = None,
                 client_id: Optional[pulumi.Input[builtins.str]] = None,
                 client_secret: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_discovery_tag_filters: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderAzureCredentialResourceDiscoveryTagFilterArgs']]]] = None,
                 resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_tags_to_add_to_metrics: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 stack_id: Optional[pulumi.Input[builtins.str]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering ProviderAzureCredential resources.
        :param pulumi.Input[Sequence[pulumi.Input['ProviderAzureCredentialAutoDiscoveryConfigurationArgs']]] auto_discovery_configurations: The list of auto discovery configurations.
        :param pulumi.Input[builtins.str] client_id: The client ID of the Azure Credential.
        :param pulumi.Input[builtins.str] client_secret: The client secret of the Azure Credential.
        :param pulumi.Input[builtins.str] name: The name of the Azure Credential.
        :param pulumi.Input[Sequence[pulumi.Input['ProviderAzureCredentialResourceDiscoveryTagFilterArgs']]] resource_discovery_tag_filters: The list of tag filters to apply to resources.
        :param pulumi.Input[builtins.str] resource_id: The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] resource_tags_to_add_to_metrics: The list of resource tags to add to metrics.
        :param pulumi.Input[builtins.str] tenant_id: The tenant ID of the Azure Credential.
        """
        if auto_discovery_configurations is not None:
            pulumi.set(__self__, "auto_discovery_configurations", auto_discovery_configurations)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_discovery_tag_filters is not None:
            pulumi.set(__self__, "resource_discovery_tag_filters", resource_discovery_tag_filters)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_tags_to_add_to_metrics is not None:
            pulumi.set(__self__, "resource_tags_to_add_to_metrics", resource_tags_to_add_to_metrics)
        if stack_id is not None:
            pulumi.set(__self__, "stack_id", stack_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="autoDiscoveryConfigurations")
    def auto_discovery_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProviderAzureCredentialAutoDiscoveryConfigurationArgs']]]]:
        """
        The list of auto discovery configurations.
        """
        return pulumi.get(self, "auto_discovery_configurations")

    @auto_discovery_configurations.setter
    def auto_discovery_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderAzureCredentialAutoDiscoveryConfigurationArgs']]]]):
        pulumi.set(self, "auto_discovery_configurations", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The client ID of the Azure Credential.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The client secret of the Azure Credential.
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the Azure Credential.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceDiscoveryTagFilters")
    def resource_discovery_tag_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProviderAzureCredentialResourceDiscoveryTagFilterArgs']]]]:
        """
        The list of tag filters to apply to resources.
        """
        return pulumi.get(self, "resource_discovery_tag_filters")

    @resource_discovery_tag_filters.setter
    def resource_discovery_tag_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderAzureCredentialResourceDiscoveryTagFilterArgs']]]]):
        pulumi.set(self, "resource_discovery_tag_filters", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceTagsToAddToMetrics")
    def resource_tags_to_add_to_metrics(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The list of resource tags to add to metrics.
        """
        return pulumi.get(self, "resource_tags_to_add_to_metrics")

    @resource_tags_to_add_to_metrics.setter
    def resource_tags_to_add_to_metrics(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "resource_tags_to_add_to_metrics", value)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "stack_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The tenant ID of the Azure Credential.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "tenant_id", value)


warnings.warn("""grafana.cloud/providerazurecredential.ProviderAzureCredential has been deprecated in favor of grafana.cloudprovider/azurecredential.AzureCredential""", DeprecationWarning)


@pulumi.type_token("grafana:cloud/providerAzureCredential:ProviderAzureCredential")
class ProviderAzureCredential(pulumi.CustomResource):
    warnings.warn("""grafana.cloud/providerazurecredential.ProviderAzureCredential has been deprecated in favor of grafana.cloudprovider/azurecredential.AzureCredential""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_discovery_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProviderAzureCredentialAutoDiscoveryConfigurationArgs', 'ProviderAzureCredentialAutoDiscoveryConfigurationArgsDict']]]]] = None,
                 client_id: Optional[pulumi.Input[builtins.str]] = None,
                 client_secret: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_discovery_tag_filters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProviderAzureCredentialResourceDiscoveryTagFilterArgs', 'ProviderAzureCredentialResourceDiscoveryTagFilterArgsDict']]]]] = None,
                 resource_tags_to_add_to_metrics: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 stack_id: Optional[pulumi.Input[builtins.str]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import grafana:cloud/providerAzureCredential:ProviderAzureCredential name "{{ stack_id }}:{{ resource_id }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ProviderAzureCredentialAutoDiscoveryConfigurationArgs', 'ProviderAzureCredentialAutoDiscoveryConfigurationArgsDict']]]] auto_discovery_configurations: The list of auto discovery configurations.
        :param pulumi.Input[builtins.str] client_id: The client ID of the Azure Credential.
        :param pulumi.Input[builtins.str] client_secret: The client secret of the Azure Credential.
        :param pulumi.Input[builtins.str] name: The name of the Azure Credential.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ProviderAzureCredentialResourceDiscoveryTagFilterArgs', 'ProviderAzureCredentialResourceDiscoveryTagFilterArgsDict']]]] resource_discovery_tag_filters: The list of tag filters to apply to resources.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] resource_tags_to_add_to_metrics: The list of resource tags to add to metrics.
        :param pulumi.Input[builtins.str] tenant_id: The tenant ID of the Azure Credential.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderAzureCredentialArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import grafana:cloud/providerAzureCredential:ProviderAzureCredential name "{{ stack_id }}:{{ resource_id }}"
        ```

        :param str resource_name: The name of the resource.
        :param ProviderAzureCredentialArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderAzureCredentialArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_discovery_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProviderAzureCredentialAutoDiscoveryConfigurationArgs', 'ProviderAzureCredentialAutoDiscoveryConfigurationArgsDict']]]]] = None,
                 client_id: Optional[pulumi.Input[builtins.str]] = None,
                 client_secret: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_discovery_tag_filters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProviderAzureCredentialResourceDiscoveryTagFilterArgs', 'ProviderAzureCredentialResourceDiscoveryTagFilterArgsDict']]]]] = None,
                 resource_tags_to_add_to_metrics: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 stack_id: Optional[pulumi.Input[builtins.str]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        pulumi.log.warn("""ProviderAzureCredential is deprecated: grafana.cloud/providerazurecredential.ProviderAzureCredential has been deprecated in favor of grafana.cloudprovider/azurecredential.AzureCredential""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderAzureCredentialArgs.__new__(ProviderAzureCredentialArgs)

            __props__.__dict__["auto_discovery_configurations"] = auto_discovery_configurations
            if client_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_id'")
            __props__.__dict__["client_id"] = client_id
            if client_secret is None and not opts.urn:
                raise TypeError("Missing required property 'client_secret'")
            __props__.__dict__["client_secret"] = None if client_secret is None else pulumi.Output.secret(client_secret)
            __props__.__dict__["name"] = name
            __props__.__dict__["resource_discovery_tag_filters"] = resource_discovery_tag_filters
            __props__.__dict__["resource_tags_to_add_to_metrics"] = resource_tags_to_add_to_metrics
            if stack_id is None and not opts.urn:
                raise TypeError("Missing required property 'stack_id'")
            __props__.__dict__["stack_id"] = stack_id
            if tenant_id is None and not opts.urn:
                raise TypeError("Missing required property 'tenant_id'")
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["resource_id"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["clientSecret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ProviderAzureCredential, __self__).__init__(
            'grafana:cloud/providerAzureCredential:ProviderAzureCredential',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_discovery_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProviderAzureCredentialAutoDiscoveryConfigurationArgs', 'ProviderAzureCredentialAutoDiscoveryConfigurationArgsDict']]]]] = None,
            client_id: Optional[pulumi.Input[builtins.str]] = None,
            client_secret: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            resource_discovery_tag_filters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProviderAzureCredentialResourceDiscoveryTagFilterArgs', 'ProviderAzureCredentialResourceDiscoveryTagFilterArgsDict']]]]] = None,
            resource_id: Optional[pulumi.Input[builtins.str]] = None,
            resource_tags_to_add_to_metrics: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            stack_id: Optional[pulumi.Input[builtins.str]] = None,
            tenant_id: Optional[pulumi.Input[builtins.str]] = None) -> 'ProviderAzureCredential':
        """
        Get an existing ProviderAzureCredential resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ProviderAzureCredentialAutoDiscoveryConfigurationArgs', 'ProviderAzureCredentialAutoDiscoveryConfigurationArgsDict']]]] auto_discovery_configurations: The list of auto discovery configurations.
        :param pulumi.Input[builtins.str] client_id: The client ID of the Azure Credential.
        :param pulumi.Input[builtins.str] client_secret: The client secret of the Azure Credential.
        :param pulumi.Input[builtins.str] name: The name of the Azure Credential.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ProviderAzureCredentialResourceDiscoveryTagFilterArgs', 'ProviderAzureCredentialResourceDiscoveryTagFilterArgsDict']]]] resource_discovery_tag_filters: The list of tag filters to apply to resources.
        :param pulumi.Input[builtins.str] resource_id: The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] resource_tags_to_add_to_metrics: The list of resource tags to add to metrics.
        :param pulumi.Input[builtins.str] tenant_id: The tenant ID of the Azure Credential.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProviderAzureCredentialState.__new__(_ProviderAzureCredentialState)

        __props__.__dict__["auto_discovery_configurations"] = auto_discovery_configurations
        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["client_secret"] = client_secret
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_discovery_tag_filters"] = resource_discovery_tag_filters
        __props__.__dict__["resource_id"] = resource_id
        __props__.__dict__["resource_tags_to_add_to_metrics"] = resource_tags_to_add_to_metrics
        __props__.__dict__["stack_id"] = stack_id
        __props__.__dict__["tenant_id"] = tenant_id
        return ProviderAzureCredential(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoDiscoveryConfigurations")
    def auto_discovery_configurations(self) -> pulumi.Output[Optional[Sequence['outputs.ProviderAzureCredentialAutoDiscoveryConfiguration']]]:
        """
        The list of auto discovery configurations.
        """
        return pulumi.get(self, "auto_discovery_configurations")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[builtins.str]:
        """
        The client ID of the Azure Credential.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[builtins.str]:
        """
        The client secret of the Azure Credential.
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the Azure Credential.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceDiscoveryTagFilters")
    def resource_discovery_tag_filters(self) -> pulumi.Output[Optional[Sequence['outputs.ProviderAzureCredentialResourceDiscoveryTagFilter']]]:
        """
        The list of tag filters to apply to resources.
        """
        return pulumi.get(self, "resource_discovery_tag_filters")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceTagsToAddToMetrics")
    def resource_tags_to_add_to_metrics(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        The list of resource tags to add to metrics.
        """
        return pulumi.get(self, "resource_tags_to_add_to_metrics")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[builtins.str]:
        """
        The tenant ID of the Azure Credential.
        """
        return pulumi.get(self, "tenant_id")

