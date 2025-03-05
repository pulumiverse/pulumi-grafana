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
from . import outputs

__all__ = [
    'GetAccessPoliciesResult',
    'AwaitableGetAccessPoliciesResult',
    'get_access_policies',
    'get_access_policies_output',
]

@pulumi.output_type
class GetAccessPoliciesResult:
    """
    A collection of values returned by getAccessPolicies.
    """
    def __init__(__self__, access_policies=None, id=None, name_filter=None, region_filter=None):
        if access_policies and not isinstance(access_policies, list):
            raise TypeError("Expected argument 'access_policies' to be a list")
        pulumi.set(__self__, "access_policies", access_policies)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name_filter and not isinstance(name_filter, str):
            raise TypeError("Expected argument 'name_filter' to be a str")
        pulumi.set(__self__, "name_filter", name_filter)
        if region_filter and not isinstance(region_filter, str):
            raise TypeError("Expected argument 'region_filter' to be a str")
        pulumi.set(__self__, "region_filter", region_filter)

    @property
    @pulumi.getter(name="accessPolicies")
    def access_policies(self) -> Sequence['outputs.GetAccessPoliciesAccessPolicyResult']:
        return pulumi.get(self, "access_policies")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of this datasource. This is an internal identifier used by the provider to track this datasource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="nameFilter")
    def name_filter(self) -> Optional[str]:
        return pulumi.get(self, "name_filter")

    @property
    @pulumi.getter(name="regionFilter")
    def region_filter(self) -> Optional[str]:
        return pulumi.get(self, "region_filter")


class AwaitableGetAccessPoliciesResult(GetAccessPoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccessPoliciesResult(
            access_policies=self.access_policies,
            id=self.id,
            name_filter=self.name_filter,
            region_filter=self.region_filter)


def get_access_policies(name_filter: Optional[str] = None,
                        region_filter: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccessPoliciesResult:
    """
    Fetches access policies from Grafana Cloud.

    * [Official documentation](https://grafana.com/docs/grafana-cloud/security-and-account-management/authentication-and-permissions/access-policies/)
    * [API documentation](https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-access-policies)

    Required access policy scopes:

    * accesspolicies:read
    """
    __args__ = dict()
    __args__['nameFilter'] = name_filter
    __args__['regionFilter'] = region_filter
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:cloud/getAccessPolicies:getAccessPolicies', __args__, opts=opts, typ=GetAccessPoliciesResult).value

    return AwaitableGetAccessPoliciesResult(
        access_policies=pulumi.get(__ret__, 'access_policies'),
        id=pulumi.get(__ret__, 'id'),
        name_filter=pulumi.get(__ret__, 'name_filter'),
        region_filter=pulumi.get(__ret__, 'region_filter'))
def get_access_policies_output(name_filter: Optional[pulumi.Input[Optional[str]]] = None,
                               region_filter: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAccessPoliciesResult]:
    """
    Fetches access policies from Grafana Cloud.

    * [Official documentation](https://grafana.com/docs/grafana-cloud/security-and-account-management/authentication-and-permissions/access-policies/)
    * [API documentation](https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-access-policies)

    Required access policy scopes:

    * accesspolicies:read
    """
    __args__ = dict()
    __args__['nameFilter'] = name_filter
    __args__['regionFilter'] = region_filter
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('grafana:cloud/getAccessPolicies:getAccessPolicies', __args__, opts=opts, typ=GetAccessPoliciesResult)
    return __ret__.apply(lambda __response__: GetAccessPoliciesResult(
        access_policies=pulumi.get(__response__, 'access_policies'),
        id=pulumi.get(__response__, 'id'),
        name_filter=pulumi.get(__response__, 'name_filter'),
        region_filter=pulumi.get(__response__, 'region_filter')))
