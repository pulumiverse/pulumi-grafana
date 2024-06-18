# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

auth: Optional[str]
"""
API token, basic auth in the `username:password` format or `anonymous` (string literal). May alternatively be set via
the `GRAFANA_AUTH` environment variable.
"""

caCert: Optional[str]
"""
Certificate CA bundle (file path or literal value) to use to verify the Grafana server's certificate. May alternatively
be set via the `GRAFANA_CA_CERT` environment variable.
"""

cloudAccessPolicyToken: Optional[str]
"""
Access Policy Token for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_ACCESS_POLICY_TOKEN` environment
variable.
"""

cloudApiUrl: Optional[str]
"""
Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
"""

insecureSkipVerify: Optional[bool]
"""
Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
"""

oncallAccessToken: Optional[str]
"""
A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
"""

oncallUrl: Optional[str]
"""
An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
"""

retries: Optional[int]
"""
The amount of retries to use for Grafana API and Grafana Cloud API calls. May alternatively be set via the
`GRAFANA_RETRIES` environment variable.
"""

retryStatusCodes: Optional[str]
"""
The status codes to retry on for Grafana API and Grafana Cloud API calls. Use `x` as a digit wildcard. Defaults to 429
and 5xx. May alternatively be set via the `GRAFANA_RETRY_STATUS_CODES` environment variable.
"""

retryWait: Optional[int]
"""
The amount of time in seconds to wait between retries for Grafana API and Grafana Cloud API calls. May alternatively be
set via the `GRAFANA_RETRY_WAIT` environment variable.
"""

smAccessToken: Optional[str]
"""
A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
"""

smUrl: Optional[str]

storeDashboardSha256: Optional[bool]
"""
Set to true if you want to save only the sha256sum instead of complete dashboard model JSON in the tfstate.
"""

tlsCert: Optional[str]
"""
Client TLS certificate (file path or literal value) to use to authenticate to the Grafana server. May alternatively be
set via the `GRAFANA_TLS_CERT` environment variable.
"""

tlsKey: Optional[str]
"""
Client TLS key (file path or literal value) to use to authenticate to the Grafana server. May alternatively be set via
the `GRAFANA_TLS_KEY` environment variable.
"""

url: Optional[str]
"""
The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
"""

