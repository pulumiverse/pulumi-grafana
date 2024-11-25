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

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 auth: Optional[pulumi.Input[str]] = None,
                 ca_cert: Optional[pulumi.Input[str]] = None,
                 cloud_access_policy_token: Optional[pulumi.Input[str]] = None,
                 cloud_api_url: Optional[pulumi.Input[str]] = None,
                 cloud_provider_access_token: Optional[pulumi.Input[str]] = None,
                 cloud_provider_url: Optional[pulumi.Input[str]] = None,
                 connections_api_access_token: Optional[pulumi.Input[str]] = None,
                 connections_api_url: Optional[pulumi.Input[str]] = None,
                 insecure_skip_verify: Optional[pulumi.Input[bool]] = None,
                 oncall_access_token: Optional[pulumi.Input[str]] = None,
                 oncall_url: Optional[pulumi.Input[str]] = None,
                 retries: Optional[pulumi.Input[int]] = None,
                 retry_status_codes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retry_wait: Optional[pulumi.Input[int]] = None,
                 sm_access_token: Optional[pulumi.Input[str]] = None,
                 sm_url: Optional[pulumi.Input[str]] = None,
                 store_dashboard_sha256: Optional[pulumi.Input[bool]] = None,
                 tls_cert: Optional[pulumi.Input[str]] = None,
                 tls_key: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] auth: API token, basic auth in the `username:password` format or `anonymous` (string literal). May alternatively be set via
               the `GRAFANA_AUTH` environment variable.
        :param pulumi.Input[str] ca_cert: Certificate CA bundle (file path or literal value) to use to verify the Grafana server's certificate. May alternatively
               be set via the `GRAFANA_CA_CERT` environment variable.
        :param pulumi.Input[str] cloud_access_policy_token: Access Policy Token for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_ACCESS_POLICY_TOKEN` environment
               variable.
        :param pulumi.Input[str] cloud_api_url: Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
        :param pulumi.Input[str] cloud_provider_access_token: A Grafana Cloud Provider access token. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_ACCESS_TOKEN`
               environment variable.
        :param pulumi.Input[str] cloud_provider_url: A Grafana Cloud Provider backend address. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_URL` environment
               variable.
        :param pulumi.Input[str] connections_api_access_token: A Grafana Connections API access token. May alternatively be set via the `GRAFANA_CONNECTIONS_API_ACCESS_TOKEN`
               environment variable.
        :param pulumi.Input[str] connections_api_url: A Grafana Connections API address. May alternatively be set via the `GRAFANA_CONNECTIONS_API_URL` environment variable.
        :param pulumi.Input[bool] insecure_skip_verify: Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
        :param pulumi.Input[str] oncall_access_token: A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
        :param pulumi.Input[str] oncall_url: An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
        :param pulumi.Input[int] retries: The amount of retries to use for Grafana API and Grafana Cloud API calls. May alternatively be set via the
               `GRAFANA_RETRIES` environment variable.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] retry_status_codes: The status codes to retry on for Grafana API and Grafana Cloud API calls. Use `x` as a digit wildcard. Defaults to 429
               and 5xx. May alternatively be set via the `GRAFANA_RETRY_STATUS_CODES` environment variable.
        :param pulumi.Input[int] retry_wait: The amount of time in seconds to wait between retries for Grafana API and Grafana Cloud API calls. May alternatively be
               set via the `GRAFANA_RETRY_WAIT` environment variable.
        :param pulumi.Input[str] sm_access_token: A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
        :param pulumi.Input[bool] store_dashboard_sha256: Set to true if you want to save only the sha256sum instead of complete dashboard model JSON in the tfstate.
        :param pulumi.Input[str] tls_cert: Client TLS certificate (file path or literal value) to use to authenticate to the Grafana server. May alternatively be
               set via the `GRAFANA_TLS_CERT` environment variable.
        :param pulumi.Input[str] tls_key: Client TLS key (file path or literal value) to use to authenticate to the Grafana server. May alternatively be set via
               the `GRAFANA_TLS_KEY` environment variable.
        :param pulumi.Input[str] url: The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
        """
        if auth is None:
            auth = _utilities.get_env('GRAFANA_AUTH')
        if auth is not None:
            pulumi.set(__self__, "auth", auth)
        if ca_cert is None:
            ca_cert = _utilities.get_env('GRAFANA_CA_CERT')
        if ca_cert is not None:
            pulumi.set(__self__, "ca_cert", ca_cert)
        if cloud_access_policy_token is None:
            cloud_access_policy_token = _utilities.get_env('GRAFANA_CLOUD_ACCESS_POLICY_TOKEN')
        if cloud_access_policy_token is not None:
            pulumi.set(__self__, "cloud_access_policy_token", cloud_access_policy_token)
        if cloud_api_url is None:
            cloud_api_url = _utilities.get_env('GRAFANA_CLOUD_API_URL')
        if cloud_api_url is not None:
            pulumi.set(__self__, "cloud_api_url", cloud_api_url)
        if cloud_provider_access_token is not None:
            pulumi.set(__self__, "cloud_provider_access_token", cloud_provider_access_token)
        if cloud_provider_url is not None:
            pulumi.set(__self__, "cloud_provider_url", cloud_provider_url)
        if connections_api_access_token is not None:
            pulumi.set(__self__, "connections_api_access_token", connections_api_access_token)
        if connections_api_url is not None:
            pulumi.set(__self__, "connections_api_url", connections_api_url)
        if insecure_skip_verify is None:
            insecure_skip_verify = _utilities.get_env_bool('GRAFANA_INSECURE_SKIP_VERIFY')
        if insecure_skip_verify is not None:
            pulumi.set(__self__, "insecure_skip_verify", insecure_skip_verify)
        if oncall_access_token is None:
            oncall_access_token = _utilities.get_env('GRAFANA_ONCALL_ACCESS_TOKEN')
        if oncall_access_token is not None:
            pulumi.set(__self__, "oncall_access_token", oncall_access_token)
        if oncall_url is None:
            oncall_url = _utilities.get_env('GRAFANA_ONCALL_URL')
        if oncall_url is not None:
            pulumi.set(__self__, "oncall_url", oncall_url)
        if retries is None:
            retries = _utilities.get_env_int('GRAFANA_RETRIES')
        if retries is not None:
            pulumi.set(__self__, "retries", retries)
        if retry_status_codes is not None:
            pulumi.set(__self__, "retry_status_codes", retry_status_codes)
        if retry_wait is None:
            retry_wait = _utilities.get_env_int('GRAFANA_RETRY_WAIT')
        if retry_wait is not None:
            pulumi.set(__self__, "retry_wait", retry_wait)
        if sm_access_token is None:
            sm_access_token = _utilities.get_env('GRAFANA_SM_ACCESS_TOKEN')
        if sm_access_token is not None:
            pulumi.set(__self__, "sm_access_token", sm_access_token)
        if sm_url is None:
            sm_url = _utilities.get_env('GRAFANA_SM_URL')
        if sm_url is not None:
            pulumi.set(__self__, "sm_url", sm_url)
        if store_dashboard_sha256 is None:
            store_dashboard_sha256 = _utilities.get_env_bool('GRAFANA_STORE_DASHBOARD_SHA256')
        if store_dashboard_sha256 is not None:
            pulumi.set(__self__, "store_dashboard_sha256", store_dashboard_sha256)
        if tls_cert is None:
            tls_cert = _utilities.get_env('GRAFANA_TLS_CERT')
        if tls_cert is not None:
            pulumi.set(__self__, "tls_cert", tls_cert)
        if tls_key is None:
            tls_key = _utilities.get_env('GRAFANA_TLS_KEY')
        if tls_key is not None:
            pulumi.set(__self__, "tls_key", tls_key)
        if url is None:
            url = _utilities.get_env('GRAFANA_URL')
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def auth(self) -> Optional[pulumi.Input[str]]:
        """
        API token, basic auth in the `username:password` format or `anonymous` (string literal). May alternatively be set via
        the `GRAFANA_AUTH` environment variable.
        """
        return pulumi.get(self, "auth")

    @auth.setter
    def auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth", value)

    @property
    @pulumi.getter(name="caCert")
    def ca_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate CA bundle (file path or literal value) to use to verify the Grafana server's certificate. May alternatively
        be set via the `GRAFANA_CA_CERT` environment variable.
        """
        return pulumi.get(self, "ca_cert")

    @ca_cert.setter
    def ca_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_cert", value)

    @property
    @pulumi.getter(name="cloudAccessPolicyToken")
    def cloud_access_policy_token(self) -> Optional[pulumi.Input[str]]:
        """
        Access Policy Token for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_ACCESS_POLICY_TOKEN` environment
        variable.
        """
        return pulumi.get(self, "cloud_access_policy_token")

    @cloud_access_policy_token.setter
    def cloud_access_policy_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_access_policy_token", value)

    @property
    @pulumi.getter(name="cloudApiUrl")
    def cloud_api_url(self) -> Optional[pulumi.Input[str]]:
        """
        Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
        """
        return pulumi.get(self, "cloud_api_url")

    @cloud_api_url.setter
    def cloud_api_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_api_url", value)

    @property
    @pulumi.getter(name="cloudProviderAccessToken")
    def cloud_provider_access_token(self) -> Optional[pulumi.Input[str]]:
        """
        A Grafana Cloud Provider access token. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_ACCESS_TOKEN`
        environment variable.
        """
        return pulumi.get(self, "cloud_provider_access_token")

    @cloud_provider_access_token.setter
    def cloud_provider_access_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_provider_access_token", value)

    @property
    @pulumi.getter(name="cloudProviderUrl")
    def cloud_provider_url(self) -> Optional[pulumi.Input[str]]:
        """
        A Grafana Cloud Provider backend address. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_URL` environment
        variable.
        """
        return pulumi.get(self, "cloud_provider_url")

    @cloud_provider_url.setter
    def cloud_provider_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_provider_url", value)

    @property
    @pulumi.getter(name="connectionsApiAccessToken")
    def connections_api_access_token(self) -> Optional[pulumi.Input[str]]:
        """
        A Grafana Connections API access token. May alternatively be set via the `GRAFANA_CONNECTIONS_API_ACCESS_TOKEN`
        environment variable.
        """
        return pulumi.get(self, "connections_api_access_token")

    @connections_api_access_token.setter
    def connections_api_access_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connections_api_access_token", value)

    @property
    @pulumi.getter(name="connectionsApiUrl")
    def connections_api_url(self) -> Optional[pulumi.Input[str]]:
        """
        A Grafana Connections API address. May alternatively be set via the `GRAFANA_CONNECTIONS_API_URL` environment variable.
        """
        return pulumi.get(self, "connections_api_url")

    @connections_api_url.setter
    def connections_api_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connections_api_url", value)

    @property
    @pulumi.getter(name="insecureSkipVerify")
    def insecure_skip_verify(self) -> Optional[pulumi.Input[bool]]:
        """
        Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
        """
        return pulumi.get(self, "insecure_skip_verify")

    @insecure_skip_verify.setter
    def insecure_skip_verify(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "insecure_skip_verify", value)

    @property
    @pulumi.getter(name="oncallAccessToken")
    def oncall_access_token(self) -> Optional[pulumi.Input[str]]:
        """
        A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
        """
        return pulumi.get(self, "oncall_access_token")

    @oncall_access_token.setter
    def oncall_access_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "oncall_access_token", value)

    @property
    @pulumi.getter(name="oncallUrl")
    def oncall_url(self) -> Optional[pulumi.Input[str]]:
        """
        An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
        """
        return pulumi.get(self, "oncall_url")

    @oncall_url.setter
    def oncall_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "oncall_url", value)

    @property
    @pulumi.getter
    def retries(self) -> Optional[pulumi.Input[int]]:
        """
        The amount of retries to use for Grafana API and Grafana Cloud API calls. May alternatively be set via the
        `GRAFANA_RETRIES` environment variable.
        """
        return pulumi.get(self, "retries")

    @retries.setter
    def retries(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retries", value)

    @property
    @pulumi.getter(name="retryStatusCodes")
    def retry_status_codes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The status codes to retry on for Grafana API and Grafana Cloud API calls. Use `x` as a digit wildcard. Defaults to 429
        and 5xx. May alternatively be set via the `GRAFANA_RETRY_STATUS_CODES` environment variable.
        """
        return pulumi.get(self, "retry_status_codes")

    @retry_status_codes.setter
    def retry_status_codes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "retry_status_codes", value)

    @property
    @pulumi.getter(name="retryWait")
    def retry_wait(self) -> Optional[pulumi.Input[int]]:
        """
        The amount of time in seconds to wait between retries for Grafana API and Grafana Cloud API calls. May alternatively be
        set via the `GRAFANA_RETRY_WAIT` environment variable.
        """
        return pulumi.get(self, "retry_wait")

    @retry_wait.setter
    def retry_wait(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retry_wait", value)

    @property
    @pulumi.getter(name="smAccessToken")
    def sm_access_token(self) -> Optional[pulumi.Input[str]]:
        """
        A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
        """
        return pulumi.get(self, "sm_access_token")

    @sm_access_token.setter
    def sm_access_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sm_access_token", value)

    @property
    @pulumi.getter(name="smUrl")
    def sm_url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sm_url")

    @sm_url.setter
    def sm_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sm_url", value)

    @property
    @pulumi.getter(name="storeDashboardSha256")
    def store_dashboard_sha256(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to true if you want to save only the sha256sum instead of complete dashboard model JSON in the tfstate.
        """
        return pulumi.get(self, "store_dashboard_sha256")

    @store_dashboard_sha256.setter
    def store_dashboard_sha256(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "store_dashboard_sha256", value)

    @property
    @pulumi.getter(name="tlsCert")
    def tls_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Client TLS certificate (file path or literal value) to use to authenticate to the Grafana server. May alternatively be
        set via the `GRAFANA_TLS_CERT` environment variable.
        """
        return pulumi.get(self, "tls_cert")

    @tls_cert.setter
    def tls_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_cert", value)

    @property
    @pulumi.getter(name="tlsKey")
    def tls_key(self) -> Optional[pulumi.Input[str]]:
        """
        Client TLS key (file path or literal value) to use to authenticate to the Grafana server. May alternatively be set via
        the `GRAFANA_TLS_KEY` environment variable.
        """
        return pulumi.get(self, "tls_key")

    @tls_key.setter
    def tls_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_key", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth: Optional[pulumi.Input[str]] = None,
                 ca_cert: Optional[pulumi.Input[str]] = None,
                 cloud_access_policy_token: Optional[pulumi.Input[str]] = None,
                 cloud_api_url: Optional[pulumi.Input[str]] = None,
                 cloud_provider_access_token: Optional[pulumi.Input[str]] = None,
                 cloud_provider_url: Optional[pulumi.Input[str]] = None,
                 connections_api_access_token: Optional[pulumi.Input[str]] = None,
                 connections_api_url: Optional[pulumi.Input[str]] = None,
                 insecure_skip_verify: Optional[pulumi.Input[bool]] = None,
                 oncall_access_token: Optional[pulumi.Input[str]] = None,
                 oncall_url: Optional[pulumi.Input[str]] = None,
                 retries: Optional[pulumi.Input[int]] = None,
                 retry_status_codes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retry_wait: Optional[pulumi.Input[int]] = None,
                 sm_access_token: Optional[pulumi.Input[str]] = None,
                 sm_url: Optional[pulumi.Input[str]] = None,
                 store_dashboard_sha256: Optional[pulumi.Input[bool]] = None,
                 tls_cert: Optional[pulumi.Input[str]] = None,
                 tls_key: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the grafana package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth: API token, basic auth in the `username:password` format or `anonymous` (string literal). May alternatively be set via
               the `GRAFANA_AUTH` environment variable.
        :param pulumi.Input[str] ca_cert: Certificate CA bundle (file path or literal value) to use to verify the Grafana server's certificate. May alternatively
               be set via the `GRAFANA_CA_CERT` environment variable.
        :param pulumi.Input[str] cloud_access_policy_token: Access Policy Token for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_ACCESS_POLICY_TOKEN` environment
               variable.
        :param pulumi.Input[str] cloud_api_url: Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
        :param pulumi.Input[str] cloud_provider_access_token: A Grafana Cloud Provider access token. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_ACCESS_TOKEN`
               environment variable.
        :param pulumi.Input[str] cloud_provider_url: A Grafana Cloud Provider backend address. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_URL` environment
               variable.
        :param pulumi.Input[str] connections_api_access_token: A Grafana Connections API access token. May alternatively be set via the `GRAFANA_CONNECTIONS_API_ACCESS_TOKEN`
               environment variable.
        :param pulumi.Input[str] connections_api_url: A Grafana Connections API address. May alternatively be set via the `GRAFANA_CONNECTIONS_API_URL` environment variable.
        :param pulumi.Input[bool] insecure_skip_verify: Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
        :param pulumi.Input[str] oncall_access_token: A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
        :param pulumi.Input[str] oncall_url: An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
        :param pulumi.Input[int] retries: The amount of retries to use for Grafana API and Grafana Cloud API calls. May alternatively be set via the
               `GRAFANA_RETRIES` environment variable.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] retry_status_codes: The status codes to retry on for Grafana API and Grafana Cloud API calls. Use `x` as a digit wildcard. Defaults to 429
               and 5xx. May alternatively be set via the `GRAFANA_RETRY_STATUS_CODES` environment variable.
        :param pulumi.Input[int] retry_wait: The amount of time in seconds to wait between retries for Grafana API and Grafana Cloud API calls. May alternatively be
               set via the `GRAFANA_RETRY_WAIT` environment variable.
        :param pulumi.Input[str] sm_access_token: A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
        :param pulumi.Input[bool] store_dashboard_sha256: Set to true if you want to save only the sha256sum instead of complete dashboard model JSON in the tfstate.
        :param pulumi.Input[str] tls_cert: Client TLS certificate (file path or literal value) to use to authenticate to the Grafana server. May alternatively be
               set via the `GRAFANA_TLS_CERT` environment variable.
        :param pulumi.Input[str] tls_key: Client TLS key (file path or literal value) to use to authenticate to the Grafana server. May alternatively be set via
               the `GRAFANA_TLS_KEY` environment variable.
        :param pulumi.Input[str] url: The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the grafana package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth: Optional[pulumi.Input[str]] = None,
                 ca_cert: Optional[pulumi.Input[str]] = None,
                 cloud_access_policy_token: Optional[pulumi.Input[str]] = None,
                 cloud_api_url: Optional[pulumi.Input[str]] = None,
                 cloud_provider_access_token: Optional[pulumi.Input[str]] = None,
                 cloud_provider_url: Optional[pulumi.Input[str]] = None,
                 connections_api_access_token: Optional[pulumi.Input[str]] = None,
                 connections_api_url: Optional[pulumi.Input[str]] = None,
                 insecure_skip_verify: Optional[pulumi.Input[bool]] = None,
                 oncall_access_token: Optional[pulumi.Input[str]] = None,
                 oncall_url: Optional[pulumi.Input[str]] = None,
                 retries: Optional[pulumi.Input[int]] = None,
                 retry_status_codes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retry_wait: Optional[pulumi.Input[int]] = None,
                 sm_access_token: Optional[pulumi.Input[str]] = None,
                 sm_url: Optional[pulumi.Input[str]] = None,
                 store_dashboard_sha256: Optional[pulumi.Input[bool]] = None,
                 tls_cert: Optional[pulumi.Input[str]] = None,
                 tls_key: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if auth is None:
                auth = _utilities.get_env('GRAFANA_AUTH')
            __props__.__dict__["auth"] = None if auth is None else pulumi.Output.secret(auth)
            if ca_cert is None:
                ca_cert = _utilities.get_env('GRAFANA_CA_CERT')
            __props__.__dict__["ca_cert"] = ca_cert
            if cloud_access_policy_token is None:
                cloud_access_policy_token = _utilities.get_env('GRAFANA_CLOUD_ACCESS_POLICY_TOKEN')
            __props__.__dict__["cloud_access_policy_token"] = None if cloud_access_policy_token is None else pulumi.Output.secret(cloud_access_policy_token)
            if cloud_api_url is None:
                cloud_api_url = _utilities.get_env('GRAFANA_CLOUD_API_URL')
            __props__.__dict__["cloud_api_url"] = cloud_api_url
            __props__.__dict__["cloud_provider_access_token"] = None if cloud_provider_access_token is None else pulumi.Output.secret(cloud_provider_access_token)
            __props__.__dict__["cloud_provider_url"] = cloud_provider_url
            __props__.__dict__["connections_api_access_token"] = None if connections_api_access_token is None else pulumi.Output.secret(connections_api_access_token)
            __props__.__dict__["connections_api_url"] = connections_api_url
            if insecure_skip_verify is None:
                insecure_skip_verify = _utilities.get_env_bool('GRAFANA_INSECURE_SKIP_VERIFY')
            __props__.__dict__["insecure_skip_verify"] = pulumi.Output.from_input(insecure_skip_verify).apply(pulumi.runtime.to_json) if insecure_skip_verify is not None else None
            if oncall_access_token is None:
                oncall_access_token = _utilities.get_env('GRAFANA_ONCALL_ACCESS_TOKEN')
            __props__.__dict__["oncall_access_token"] = None if oncall_access_token is None else pulumi.Output.secret(oncall_access_token)
            if oncall_url is None:
                oncall_url = _utilities.get_env('GRAFANA_ONCALL_URL')
            __props__.__dict__["oncall_url"] = oncall_url
            if retries is None:
                retries = _utilities.get_env_int('GRAFANA_RETRIES')
            __props__.__dict__["retries"] = pulumi.Output.from_input(retries).apply(pulumi.runtime.to_json) if retries is not None else None
            __props__.__dict__["retry_status_codes"] = pulumi.Output.from_input(retry_status_codes).apply(pulumi.runtime.to_json) if retry_status_codes is not None else None
            if retry_wait is None:
                retry_wait = _utilities.get_env_int('GRAFANA_RETRY_WAIT')
            __props__.__dict__["retry_wait"] = pulumi.Output.from_input(retry_wait).apply(pulumi.runtime.to_json) if retry_wait is not None else None
            if sm_access_token is None:
                sm_access_token = _utilities.get_env('GRAFANA_SM_ACCESS_TOKEN')
            __props__.__dict__["sm_access_token"] = None if sm_access_token is None else pulumi.Output.secret(sm_access_token)
            if sm_url is None:
                sm_url = _utilities.get_env('GRAFANA_SM_URL')
            __props__.__dict__["sm_url"] = sm_url
            if store_dashboard_sha256 is None:
                store_dashboard_sha256 = _utilities.get_env_bool('GRAFANA_STORE_DASHBOARD_SHA256')
            __props__.__dict__["store_dashboard_sha256"] = pulumi.Output.from_input(store_dashboard_sha256).apply(pulumi.runtime.to_json) if store_dashboard_sha256 is not None else None
            if tls_cert is None:
                tls_cert = _utilities.get_env('GRAFANA_TLS_CERT')
            __props__.__dict__["tls_cert"] = tls_cert
            if tls_key is None:
                tls_key = _utilities.get_env('GRAFANA_TLS_KEY')
            __props__.__dict__["tls_key"] = None if tls_key is None else pulumi.Output.secret(tls_key)
            if url is None:
                url = _utilities.get_env('GRAFANA_URL')
            __props__.__dict__["url"] = url
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["auth", "cloudAccessPolicyToken", "cloudProviderAccessToken", "connectionsApiAccessToken", "oncallAccessToken", "smAccessToken", "tlsKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Provider, __self__).__init__(
            'grafana',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter
    def auth(self) -> pulumi.Output[Optional[str]]:
        """
        API token, basic auth in the `username:password` format or `anonymous` (string literal). May alternatively be set via
        the `GRAFANA_AUTH` environment variable.
        """
        return pulumi.get(self, "auth")

    @property
    @pulumi.getter(name="caCert")
    def ca_cert(self) -> pulumi.Output[Optional[str]]:
        """
        Certificate CA bundle (file path or literal value) to use to verify the Grafana server's certificate. May alternatively
        be set via the `GRAFANA_CA_CERT` environment variable.
        """
        return pulumi.get(self, "ca_cert")

    @property
    @pulumi.getter(name="cloudAccessPolicyToken")
    def cloud_access_policy_token(self) -> pulumi.Output[Optional[str]]:
        """
        Access Policy Token for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_ACCESS_POLICY_TOKEN` environment
        variable.
        """
        return pulumi.get(self, "cloud_access_policy_token")

    @property
    @pulumi.getter(name="cloudApiUrl")
    def cloud_api_url(self) -> pulumi.Output[Optional[str]]:
        """
        Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
        """
        return pulumi.get(self, "cloud_api_url")

    @property
    @pulumi.getter(name="cloudProviderAccessToken")
    def cloud_provider_access_token(self) -> pulumi.Output[Optional[str]]:
        """
        A Grafana Cloud Provider access token. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_ACCESS_TOKEN`
        environment variable.
        """
        return pulumi.get(self, "cloud_provider_access_token")

    @property
    @pulumi.getter(name="cloudProviderUrl")
    def cloud_provider_url(self) -> pulumi.Output[Optional[str]]:
        """
        A Grafana Cloud Provider backend address. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_URL` environment
        variable.
        """
        return pulumi.get(self, "cloud_provider_url")

    @property
    @pulumi.getter(name="connectionsApiAccessToken")
    def connections_api_access_token(self) -> pulumi.Output[Optional[str]]:
        """
        A Grafana Connections API access token. May alternatively be set via the `GRAFANA_CONNECTIONS_API_ACCESS_TOKEN`
        environment variable.
        """
        return pulumi.get(self, "connections_api_access_token")

    @property
    @pulumi.getter(name="connectionsApiUrl")
    def connections_api_url(self) -> pulumi.Output[Optional[str]]:
        """
        A Grafana Connections API address. May alternatively be set via the `GRAFANA_CONNECTIONS_API_URL` environment variable.
        """
        return pulumi.get(self, "connections_api_url")

    @property
    @pulumi.getter(name="oncallAccessToken")
    def oncall_access_token(self) -> pulumi.Output[Optional[str]]:
        """
        A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
        """
        return pulumi.get(self, "oncall_access_token")

    @property
    @pulumi.getter(name="oncallUrl")
    def oncall_url(self) -> pulumi.Output[Optional[str]]:
        """
        An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
        """
        return pulumi.get(self, "oncall_url")

    @property
    @pulumi.getter(name="smAccessToken")
    def sm_access_token(self) -> pulumi.Output[Optional[str]]:
        """
        A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
        """
        return pulumi.get(self, "sm_access_token")

    @property
    @pulumi.getter(name="smUrl")
    def sm_url(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "sm_url")

    @property
    @pulumi.getter(name="tlsCert")
    def tls_cert(self) -> pulumi.Output[Optional[str]]:
        """
        Client TLS certificate (file path or literal value) to use to authenticate to the Grafana server. May alternatively be
        set via the `GRAFANA_TLS_CERT` environment variable.
        """
        return pulumi.get(self, "tls_cert")

    @property
    @pulumi.getter(name="tlsKey")
    def tls_key(self) -> pulumi.Output[Optional[str]]:
        """
        Client TLS key (file path or literal value) to use to authenticate to the Grafana server. May alternatively be set via
        the `GRAFANA_TLS_KEY` environment variable.
        """
        return pulumi.get(self, "tls_key")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[Optional[str]]:
        """
        The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
        """
        return pulumi.get(self, "url")

