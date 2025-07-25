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

__all__ = [
    'GetAzureCredentialResult',
    'AwaitableGetAzureCredentialResult',
    'get_azure_credential',
    'get_azure_credential_output',
]

@pulumi.output_type
class GetAzureCredentialResult:
    """
    A collection of values returned by getAzureCredential.
    """
    def __init__(__self__, auto_discovery_configurations=None, client_id=None, client_secret=None, id=None, name=None, resource_discovery_tag_filters=None, resource_id=None, resource_tags_to_add_to_metrics=None, stack_id=None, tenant_id=None):
        if auto_discovery_configurations and not isinstance(auto_discovery_configurations, list):
            raise TypeError("Expected argument 'auto_discovery_configurations' to be a list")
        pulumi.set(__self__, "auto_discovery_configurations", auto_discovery_configurations)
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if client_secret and not isinstance(client_secret, str):
            raise TypeError("Expected argument 'client_secret' to be a str")
        pulumi.set(__self__, "client_secret", client_secret)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_discovery_tag_filters and not isinstance(resource_discovery_tag_filters, list):
            raise TypeError("Expected argument 'resource_discovery_tag_filters' to be a list")
        pulumi.set(__self__, "resource_discovery_tag_filters", resource_discovery_tag_filters)
        if resource_id and not isinstance(resource_id, str):
            raise TypeError("Expected argument 'resource_id' to be a str")
        pulumi.set(__self__, "resource_id", resource_id)
        if resource_tags_to_add_to_metrics and not isinstance(resource_tags_to_add_to_metrics, list):
            raise TypeError("Expected argument 'resource_tags_to_add_to_metrics' to be a list")
        pulumi.set(__self__, "resource_tags_to_add_to_metrics", resource_tags_to_add_to_metrics)
        if stack_id and not isinstance(stack_id, str):
            raise TypeError("Expected argument 'stack_id' to be a str")
        pulumi.set(__self__, "stack_id", stack_id)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="autoDiscoveryConfigurations")
    def auto_discovery_configurations(self) -> Optional[Sequence['outputs.GetAzureCredentialAutoDiscoveryConfigurationResult']]:
        """
        The list of auto discovery configurations.
        """
        return pulumi.get(self, "auto_discovery_configurations")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> builtins.str:
        """
        The client ID of the Azure Credential.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> builtins.str:
        """
        The client secret of the Azure Credential.
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the Azure Credential.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceDiscoveryTagFilters")
    def resource_discovery_tag_filters(self) -> Optional[Sequence['outputs.GetAzureCredentialResourceDiscoveryTagFilterResult']]:
        """
        The list of tag filters to apply to resources.
        """
        return pulumi.get(self, "resource_discovery_tag_filters")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> builtins.str:
        """
        The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceTagsToAddToMetrics")
    def resource_tags_to_add_to_metrics(self) -> Sequence[builtins.str]:
        """
        The list of resource tags to add to metrics.
        """
        return pulumi.get(self, "resource_tags_to_add_to_metrics")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> builtins.str:
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> builtins.str:
        """
        The tenant ID of the Azure Credential.
        """
        return pulumi.get(self, "tenant_id")


class AwaitableGetAzureCredentialResult(GetAzureCredentialResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAzureCredentialResult(
            auto_discovery_configurations=self.auto_discovery_configurations,
            client_id=self.client_id,
            client_secret=self.client_secret,
            id=self.id,
            name=self.name,
            resource_discovery_tag_filters=self.resource_discovery_tag_filters,
            resource_id=self.resource_id,
            resource_tags_to_add_to_metrics=self.resource_tags_to_add_to_metrics,
            stack_id=self.stack_id,
            tenant_id=self.tenant_id)


def get_azure_credential(auto_discovery_configurations: Optional[Sequence[Union['GetAzureCredentialAutoDiscoveryConfigurationArgs', 'GetAzureCredentialAutoDiscoveryConfigurationArgsDict']]] = None,
                         resource_discovery_tag_filters: Optional[Sequence[Union['GetAzureCredentialResourceDiscoveryTagFilterArgs', 'GetAzureCredentialResourceDiscoveryTagFilterArgsDict']]] = None,
                         resource_id: Optional[builtins.str] = None,
                         stack_id: Optional[builtins.str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAzureCredentialResult:
    """
    ## Example Usage


    :param Sequence[Union['GetAzureCredentialAutoDiscoveryConfigurationArgs', 'GetAzureCredentialAutoDiscoveryConfigurationArgsDict']] auto_discovery_configurations: The list of auto discovery configurations.
    :param Sequence[Union['GetAzureCredentialResourceDiscoveryTagFilterArgs', 'GetAzureCredentialResourceDiscoveryTagFilterArgsDict']] resource_discovery_tag_filters: The list of tag filters to apply to resources.
    :param builtins.str resource_id: The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
    """
    __args__ = dict()
    __args__['autoDiscoveryConfigurations'] = auto_discovery_configurations
    __args__['resourceDiscoveryTagFilters'] = resource_discovery_tag_filters
    __args__['resourceId'] = resource_id
    __args__['stackId'] = stack_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:cloudProvider/getAzureCredential:getAzureCredential', __args__, opts=opts, typ=GetAzureCredentialResult).value

    return AwaitableGetAzureCredentialResult(
        auto_discovery_configurations=pulumi.get(__ret__, 'auto_discovery_configurations'),
        client_id=pulumi.get(__ret__, 'client_id'),
        client_secret=pulumi.get(__ret__, 'client_secret'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        resource_discovery_tag_filters=pulumi.get(__ret__, 'resource_discovery_tag_filters'),
        resource_id=pulumi.get(__ret__, 'resource_id'),
        resource_tags_to_add_to_metrics=pulumi.get(__ret__, 'resource_tags_to_add_to_metrics'),
        stack_id=pulumi.get(__ret__, 'stack_id'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))
def get_azure_credential_output(auto_discovery_configurations: Optional[pulumi.Input[Optional[Sequence[Union['GetAzureCredentialAutoDiscoveryConfigurationArgs', 'GetAzureCredentialAutoDiscoveryConfigurationArgsDict']]]]] = None,
                                resource_discovery_tag_filters: Optional[pulumi.Input[Optional[Sequence[Union['GetAzureCredentialResourceDiscoveryTagFilterArgs', 'GetAzureCredentialResourceDiscoveryTagFilterArgsDict']]]]] = None,
                                resource_id: Optional[pulumi.Input[builtins.str]] = None,
                                stack_id: Optional[pulumi.Input[builtins.str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAzureCredentialResult]:
    """
    ## Example Usage


    :param Sequence[Union['GetAzureCredentialAutoDiscoveryConfigurationArgs', 'GetAzureCredentialAutoDiscoveryConfigurationArgsDict']] auto_discovery_configurations: The list of auto discovery configurations.
    :param Sequence[Union['GetAzureCredentialResourceDiscoveryTagFilterArgs', 'GetAzureCredentialResourceDiscoveryTagFilterArgsDict']] resource_discovery_tag_filters: The list of tag filters to apply to resources.
    :param builtins.str resource_id: The ID given by the Grafana Cloud Provider API to this Azure Credential resource.
    """
    __args__ = dict()
    __args__['autoDiscoveryConfigurations'] = auto_discovery_configurations
    __args__['resourceDiscoveryTagFilters'] = resource_discovery_tag_filters
    __args__['resourceId'] = resource_id
    __args__['stackId'] = stack_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('grafana:cloudProvider/getAzureCredential:getAzureCredential', __args__, opts=opts, typ=GetAzureCredentialResult)
    return __ret__.apply(lambda __response__: GetAzureCredentialResult(
        auto_discovery_configurations=pulumi.get(__response__, 'auto_discovery_configurations'),
        client_id=pulumi.get(__response__, 'client_id'),
        client_secret=pulumi.get(__response__, 'client_secret'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        resource_discovery_tag_filters=pulumi.get(__response__, 'resource_discovery_tag_filters'),
        resource_id=pulumi.get(__response__, 'resource_id'),
        resource_tags_to_add_to_metrics=pulumi.get(__response__, 'resource_tags_to_add_to_metrics'),
        stack_id=pulumi.get(__response__, 'stack_id'),
        tenant_id=pulumi.get(__response__, 'tenant_id')))
