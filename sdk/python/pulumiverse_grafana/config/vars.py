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

import types

__config__ = pulumi.Config('grafana')


class _ExportableConfig(types.ModuleType):
    @property
    def auth(self) -> Optional[str]:
        """
        API token, basic auth in the `username:password` format or `anonymous` (string literal). May alternatively be set via
        the `GRAFANA_AUTH` environment variable.
        """
        return __config__.get('auth') or _utilities.get_env('GRAFANA_AUTH')

    @property
    def ca_cert(self) -> Optional[str]:
        """
        Certificate CA bundle (file path or literal value) to use to verify the Grafana server's certificate. May alternatively
        be set via the `GRAFANA_CA_CERT` environment variable.
        """
        return __config__.get('caCert') or _utilities.get_env('GRAFANA_CA_CERT')

    @property
    def cloud_access_policy_token(self) -> Optional[str]:
        """
        Access Policy Token for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_ACCESS_POLICY_TOKEN` environment
        variable.
        """
        return __config__.get('cloudAccessPolicyToken') or _utilities.get_env('GRAFANA_CLOUD_ACCESS_POLICY_TOKEN')

    @property
    def cloud_api_url(self) -> Optional[str]:
        """
        Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
        """
        return __config__.get('cloudApiUrl') or _utilities.get_env('GRAFANA_CLOUD_API_URL')

    @property
    def cloud_provider_access_token(self) -> Optional[str]:
        """
        A Grafana Cloud Provider access token. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_ACCESS_TOKEN`
        environment variable.
        """
        return __config__.get('cloudProviderAccessToken')

    @property
    def cloud_provider_url(self) -> Optional[str]:
        """
        A Grafana Cloud Provider backend address. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_URL` environment
        variable.
        """
        return __config__.get('cloudProviderUrl')

    @property
    def connections_api_access_token(self) -> Optional[str]:
        """
        A Grafana Connections API access token. May alternatively be set via the `GRAFANA_CONNECTIONS_API_ACCESS_TOKEN`
        environment variable.
        """
        return __config__.get('connectionsApiAccessToken')

    @property
    def connections_api_url(self) -> Optional[str]:
        """
        A Grafana Connections API address. May alternatively be set via the `GRAFANA_CONNECTIONS_API_URL` environment variable.
        """
        return __config__.get('connectionsApiUrl')

    @property
    def fleet_management_auth(self) -> Optional[str]:
        """
        A Grafana Fleet Management basic auth in the `username:password` format. May alternatively be set via the
        `GRAFANA_FLEET_MANAGEMENT_AUTH` environment variable.
        """
        return __config__.get('fleetManagementAuth')

    @property
    def fleet_management_url(self) -> Optional[str]:
        """
        A Grafana Fleet Management API address. May alternatively be set via the `GRAFANA_FLEET_MANAGEMENT_URL` environment
        variable.
        """
        return __config__.get('fleetManagementUrl')

    @property
    def frontend_o11y_api_access_token(self) -> Optional[str]:
        """
        A Grafana Frontend Observability API access token. May alternatively be set via the
        `GRAFANA_FRONTEND_O11Y_API_ACCESS_TOKEN` environment variable.
        """
        return __config__.get('frontendO11yApiAccessToken')

    @property
    def insecure_skip_verify(self) -> Optional[bool]:
        """
        Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
        """
        return __config__.get_bool('insecureSkipVerify') or _utilities.get_env_bool('GRAFANA_INSECURE_SKIP_VERIFY')

    @property
    def oncall_access_token(self) -> Optional[str]:
        """
        A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
        """
        return __config__.get('oncallAccessToken') or _utilities.get_env('GRAFANA_ONCALL_ACCESS_TOKEN')

    @property
    def oncall_url(self) -> Optional[str]:
        """
        An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
        """
        return __config__.get('oncallUrl') or _utilities.get_env('GRAFANA_ONCALL_URL')

    @property
    def retries(self) -> Optional[int]:
        """
        The amount of retries to use for Grafana API and Grafana Cloud API calls. May alternatively be set via the
        `GRAFANA_RETRIES` environment variable.
        """
        return __config__.get_int('retries') or _utilities.get_env_int('GRAFANA_RETRIES')

    @property
    def retry_status_codes(self) -> Optional[str]:
        """
        The status codes to retry on for Grafana API and Grafana Cloud API calls. Use `x` as a digit wildcard. Defaults to 429
        and 5xx. May alternatively be set via the `GRAFANA_RETRY_STATUS_CODES` environment variable.
        """
        return __config__.get('retryStatusCodes')

    @property
    def retry_wait(self) -> Optional[int]:
        """
        The amount of time in seconds to wait between retries for Grafana API and Grafana Cloud API calls. May alternatively be
        set via the `GRAFANA_RETRY_WAIT` environment variable.
        """
        return __config__.get_int('retryWait') or _utilities.get_env_int('GRAFANA_RETRY_WAIT')

    @property
    def sm_access_token(self) -> Optional[str]:
        """
        A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
        """
        return __config__.get('smAccessToken') or _utilities.get_env('GRAFANA_SM_ACCESS_TOKEN')

    @property
    def sm_url(self) -> Optional[str]:
        return __config__.get('smUrl') or _utilities.get_env('GRAFANA_SM_URL')

    @property
    def store_dashboard_sha256(self) -> Optional[bool]:
        """
        Set to true if you want to save only the sha256sum instead of complete dashboard model JSON in the tfstate.
        """
        return __config__.get_bool('storeDashboardSha256') or _utilities.get_env_bool('GRAFANA_STORE_DASHBOARD_SHA256')

    @property
    def tls_cert(self) -> Optional[str]:
        """
        Client TLS certificate (file path or literal value) to use to authenticate to the Grafana server. May alternatively be
        set via the `GRAFANA_TLS_CERT` environment variable.
        """
        return __config__.get('tlsCert') or _utilities.get_env('GRAFANA_TLS_CERT')

    @property
    def tls_key(self) -> Optional[str]:
        """
        Client TLS key (file path or literal value) to use to authenticate to the Grafana server. May alternatively be set via
        the `GRAFANA_TLS_KEY` environment variable.
        """
        return __config__.get('tlsKey') or _utilities.get_env('GRAFANA_TLS_KEY')

    @property
    def url(self) -> Optional[str]:
        """
        The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
        """
        return __config__.get('url') or _utilities.get_env('GRAFANA_URL')

