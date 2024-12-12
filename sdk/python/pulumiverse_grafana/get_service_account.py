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

__all__ = [
    'GetServiceAccountResult',
    'AwaitableGetServiceAccountResult',
    'get_service_account',
    'get_service_account_output',
]

warnings.warn("""grafana.index/getserviceaccount.getServiceAccount has been deprecated in favor of grafana.oss/getserviceaccount.getServiceAccount""", DeprecationWarning)

@pulumi.output_type
class GetServiceAccountResult:
    """
    A collection of values returned by getServiceAccount.
    """
    def __init__(__self__, id=None, is_disabled=None, name=None, org_id=None, role=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_disabled and not isinstance(is_disabled, bool):
            raise TypeError("Expected argument 'is_disabled' to be a bool")
        pulumi.set(__self__, "is_disabled", is_disabled)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isDisabled")
    def is_disabled(self) -> bool:
        """
        The disabled status for the service account.
        """
        return pulumi.get(self, "is_disabled")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Service Account.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[str]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        The basic role of the service account in the organization.
        """
        return pulumi.get(self, "role")


class AwaitableGetServiceAccountResult(GetServiceAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceAccountResult(
            id=self.id,
            is_disabled=self.is_disabled,
            name=self.name,
            org_id=self.org_id,
            role=self.role)


def get_service_account(name: Optional[str] = None,
                        org_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceAccountResult:
    """
    * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/)
            * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana
    import pulumiverse_grafana as grafana

    admin_service_account = grafana.oss.ServiceAccount("admin",
        name="admin sa",
        role="Admin",
        is_disabled=False)
    admin = grafana.oss.get_service_account_output(name=admin_service_account.name)
    ```


    :param str name: The name of the Service Account.
    :param str org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
    """
    pulumi.log.warn("""get_service_account is deprecated: grafana.index/getserviceaccount.getServiceAccount has been deprecated in favor of grafana.oss/getserviceaccount.getServiceAccount""")
    __args__ = dict()
    __args__['name'] = name
    __args__['orgId'] = org_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:index/getServiceAccount:getServiceAccount', __args__, opts=opts, typ=GetServiceAccountResult).value

    return AwaitableGetServiceAccountResult(
        id=pulumi.get(__ret__, 'id'),
        is_disabled=pulumi.get(__ret__, 'is_disabled'),
        name=pulumi.get(__ret__, 'name'),
        org_id=pulumi.get(__ret__, 'org_id'),
        role=pulumi.get(__ret__, 'role'))
def get_service_account_output(name: Optional[pulumi.Input[str]] = None,
                               org_id: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetServiceAccountResult]:
    """
    * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/)
            * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana
    import pulumiverse_grafana as grafana

    admin_service_account = grafana.oss.ServiceAccount("admin",
        name="admin sa",
        role="Admin",
        is_disabled=False)
    admin = grafana.oss.get_service_account_output(name=admin_service_account.name)
    ```


    :param str name: The name of the Service Account.
    :param str org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
    """
    pulumi.log.warn("""get_service_account is deprecated: grafana.index/getserviceaccount.getServiceAccount has been deprecated in favor of grafana.oss/getserviceaccount.getServiceAccount""")
    __args__ = dict()
    __args__['name'] = name
    __args__['orgId'] = org_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('grafana:index/getServiceAccount:getServiceAccount', __args__, opts=opts, typ=GetServiceAccountResult)
    return __ret__.apply(lambda __response__: GetServiceAccountResult(
        id=pulumi.get(__response__, 'id'),
        is_disabled=pulumi.get(__response__, 'is_disabled'),
        name=pulumi.get(__response__, 'name'),
        org_id=pulumi.get(__response__, 'org_id'),
        role=pulumi.get(__response__, 'role')))
