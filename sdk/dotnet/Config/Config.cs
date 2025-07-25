// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumiverse.Grafana
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("grafana");

        private static readonly __Value<string?> _auth = new __Value<string?>(() => __config.Get("auth") ?? Utilities.GetEnv("GRAFANA_AUTH"));
        /// <summary>
        /// API token, basic auth in the `username:password` format or `anonymous` (string literal). May alternatively be set via
        /// the `GRAFANA_AUTH` environment variable.
        /// </summary>
        public static string? Auth
        {
            get => _auth.Get();
            set => _auth.Set(value);
        }

        private static readonly __Value<string?> _caCert = new __Value<string?>(() => __config.Get("caCert") ?? Utilities.GetEnv("GRAFANA_CA_CERT"));
        /// <summary>
        /// Certificate CA bundle (file path or literal value) to use to verify the Grafana server's certificate. May alternatively
        /// be set via the `GRAFANA_CA_CERT` environment variable.
        /// </summary>
        public static string? CaCert
        {
            get => _caCert.Get();
            set => _caCert.Set(value);
        }

        private static readonly __Value<string?> _cloudAccessPolicyToken = new __Value<string?>(() => __config.Get("cloudAccessPolicyToken") ?? Utilities.GetEnv("GRAFANA_CLOUD_ACCESS_POLICY_TOKEN"));
        /// <summary>
        /// Access Policy Token for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_ACCESS_POLICY_TOKEN` environment
        /// variable.
        /// </summary>
        public static string? CloudAccessPolicyToken
        {
            get => _cloudAccessPolicyToken.Get();
            set => _cloudAccessPolicyToken.Set(value);
        }

        private static readonly __Value<string?> _cloudApiUrl = new __Value<string?>(() => __config.Get("cloudApiUrl") ?? Utilities.GetEnv("GRAFANA_CLOUD_API_URL"));
        /// <summary>
        /// Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
        /// </summary>
        public static string? CloudApiUrl
        {
            get => _cloudApiUrl.Get();
            set => _cloudApiUrl.Set(value);
        }

        private static readonly __Value<string?> _cloudProviderAccessToken = new __Value<string?>(() => __config.Get("cloudProviderAccessToken"));
        /// <summary>
        /// A Grafana Cloud Provider access token. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_ACCESS_TOKEN`
        /// environment variable.
        /// </summary>
        public static string? CloudProviderAccessToken
        {
            get => _cloudProviderAccessToken.Get();
            set => _cloudProviderAccessToken.Set(value);
        }

        private static readonly __Value<string?> _cloudProviderUrl = new __Value<string?>(() => __config.Get("cloudProviderUrl"));
        /// <summary>
        /// A Grafana Cloud Provider backend address. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_URL` environment
        /// variable.
        /// </summary>
        public static string? CloudProviderUrl
        {
            get => _cloudProviderUrl.Get();
            set => _cloudProviderUrl.Set(value);
        }

        private static readonly __Value<string?> _connectionsApiAccessToken = new __Value<string?>(() => __config.Get("connectionsApiAccessToken"));
        /// <summary>
        /// A Grafana Connections API access token. May alternatively be set via the `GRAFANA_CONNECTIONS_API_ACCESS_TOKEN`
        /// environment variable.
        /// </summary>
        public static string? ConnectionsApiAccessToken
        {
            get => _connectionsApiAccessToken.Get();
            set => _connectionsApiAccessToken.Set(value);
        }

        private static readonly __Value<string?> _connectionsApiUrl = new __Value<string?>(() => __config.Get("connectionsApiUrl"));
        /// <summary>
        /// A Grafana Connections API address. May alternatively be set via the `GRAFANA_CONNECTIONS_API_URL` environment variable.
        /// </summary>
        public static string? ConnectionsApiUrl
        {
            get => _connectionsApiUrl.Get();
            set => _connectionsApiUrl.Set(value);
        }

        private static readonly __Value<string?> _fleetManagementAuth = new __Value<string?>(() => __config.Get("fleetManagementAuth"));
        /// <summary>
        /// A Grafana Fleet Management basic auth in the `username:password` format. May alternatively be set via the
        /// `GRAFANA_FLEET_MANAGEMENT_AUTH` environment variable.
        /// </summary>
        public static string? FleetManagementAuth
        {
            get => _fleetManagementAuth.Get();
            set => _fleetManagementAuth.Set(value);
        }

        private static readonly __Value<string?> _fleetManagementUrl = new __Value<string?>(() => __config.Get("fleetManagementUrl"));
        /// <summary>
        /// A Grafana Fleet Management API address. May alternatively be set via the `GRAFANA_FLEET_MANAGEMENT_URL` environment
        /// variable.
        /// </summary>
        public static string? FleetManagementUrl
        {
            get => _fleetManagementUrl.Get();
            set => _fleetManagementUrl.Set(value);
        }

        private static readonly __Value<string?> _frontendO11yApiAccessToken = new __Value<string?>(() => __config.Get("frontendO11yApiAccessToken"));
        /// <summary>
        /// A Grafana Frontend Observability API access token. May alternatively be set via the
        /// `GRAFANA_FRONTEND_O11Y_API_ACCESS_TOKEN` environment variable.
        /// </summary>
        public static string? FrontendO11yApiAccessToken
        {
            get => _frontendO11yApiAccessToken.Get();
            set => _frontendO11yApiAccessToken.Set(value);
        }

        private static readonly __Value<ImmutableDictionary<string, string>?> _httpHeaders = new __Value<ImmutableDictionary<string, string>?>(() => __config.GetObject<ImmutableDictionary<string, string>>("httpHeaders"));
        /// <summary>
        /// Optional. HTTP headers mapping keys to values used for accessing the Grafana and Grafana Cloud APIs. May alternatively
        /// be set via the `GRAFANA_HTTP_HEADERS` environment variable in JSON format.
        /// </summary>
        public static ImmutableDictionary<string, string>? HttpHeaders
        {
            get => _httpHeaders.Get();
            set => _httpHeaders.Set(value);
        }

        private static readonly __Value<bool?> _insecureSkipVerify = new __Value<bool?>(() => __config.GetBoolean("insecureSkipVerify") ?? Utilities.GetEnvBoolean("GRAFANA_INSECURE_SKIP_VERIFY"));
        /// <summary>
        /// Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
        /// </summary>
        public static bool? InsecureSkipVerify
        {
            get => _insecureSkipVerify.Get();
            set => _insecureSkipVerify.Set(value);
        }

        private static readonly __Value<string?> _k6AccessToken = new __Value<string?>(() => __config.Get("k6AccessToken"));
        /// <summary>
        /// The k6 Cloud API token. May alternatively be set via the `GRAFANA_K6_ACCESS_TOKEN` environment variable.
        /// </summary>
        public static string? K6AccessToken
        {
            get => _k6AccessToken.Get();
            set => _k6AccessToken.Set(value);
        }

        private static readonly __Value<string?> _k6Url = new __Value<string?>(() => __config.Get("k6Url"));
        /// <summary>
        /// The k6 Cloud API url. May alternatively be set via the `GRAFANA_K6_URL` environment variable.
        /// </summary>
        public static string? K6Url
        {
            get => _k6Url.Get();
            set => _k6Url.Set(value);
        }

        private static readonly __Value<string?> _oncallAccessToken = new __Value<string?>(() => __config.Get("oncallAccessToken") ?? Utilities.GetEnv("GRAFANA_ONCALL_ACCESS_TOKEN"));
        /// <summary>
        /// A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
        /// </summary>
        public static string? OncallAccessToken
        {
            get => _oncallAccessToken.Get();
            set => _oncallAccessToken.Set(value);
        }

        private static readonly __Value<string?> _oncallUrl = new __Value<string?>(() => __config.Get("oncallUrl") ?? Utilities.GetEnv("GRAFANA_ONCALL_URL"));
        /// <summary>
        /// An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
        /// </summary>
        public static string? OncallUrl
        {
            get => _oncallUrl.Get();
            set => _oncallUrl.Set(value);
        }

        private static readonly __Value<int?> _orgId = new __Value<int?>(() => __config.GetInt32("orgId"));
        /// <summary>
        /// The Grafana org ID, if you are using a self-hosted OSS or enterprise Grafana instance. May alternatively be set via the
        /// `GRAFANA_ORG_ID` environment variable.
        /// </summary>
        public static int? OrgId
        {
            get => _orgId.Get();
            set => _orgId.Set(value);
        }

        private static readonly __Value<int?> _retries = new __Value<int?>(() => __config.GetInt32("retries") ?? Utilities.GetEnvInt32("GRAFANA_RETRIES"));
        /// <summary>
        /// The amount of retries to use for Grafana API and Grafana Cloud API calls. May alternatively be set via the
        /// `GRAFANA_RETRIES` environment variable.
        /// </summary>
        public static int? Retries
        {
            get => _retries.Get();
            set => _retries.Set(value);
        }

        private static readonly __Value<ImmutableArray<string>> _retryStatusCodes = new __Value<ImmutableArray<string>>(() => __config.GetObject<ImmutableArray<string>>("retryStatusCodes"));
        /// <summary>
        /// The status codes to retry on for Grafana API and Grafana Cloud API calls. Use `x` as a digit wildcard. Defaults to 429
        /// and 5xx. May alternatively be set via the `GRAFANA_RETRY_STATUS_CODES` environment variable.
        /// </summary>
        public static ImmutableArray<string> RetryStatusCodes
        {
            get => _retryStatusCodes.Get();
            set => _retryStatusCodes.Set(value);
        }

        private static readonly __Value<int?> _retryWait = new __Value<int?>(() => __config.GetInt32("retryWait") ?? Utilities.GetEnvInt32("GRAFANA_RETRY_WAIT"));
        /// <summary>
        /// The amount of time in seconds to wait between retries for Grafana API and Grafana Cloud API calls. May alternatively be
        /// set via the `GRAFANA_RETRY_WAIT` environment variable.
        /// </summary>
        public static int? RetryWait
        {
            get => _retryWait.Get();
            set => _retryWait.Set(value);
        }

        private static readonly __Value<string?> _smAccessToken = new __Value<string?>(() => __config.Get("smAccessToken") ?? Utilities.GetEnv("GRAFANA_SM_ACCESS_TOKEN"));
        /// <summary>
        /// A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
        /// </summary>
        public static string? SmAccessToken
        {
            get => _smAccessToken.Get();
            set => _smAccessToken.Set(value);
        }

        private static readonly __Value<string?> _smUrl = new __Value<string?>(() => __config.Get("smUrl") ?? Utilities.GetEnv("GRAFANA_SM_URL"));
        public static string? SmUrl
        {
            get => _smUrl.Get();
            set => _smUrl.Set(value);
        }

        private static readonly __Value<int?> _stackId = new __Value<int?>(() => __config.GetInt32("stackId"));
        /// <summary>
        /// The Grafana stack ID, if you are using a Grafana Cloud stack. May alternatively be set via the `GRAFANA_STACK_ID`
        /// environment variable.
        /// </summary>
        public static int? StackId
        {
            get => _stackId.Get();
            set => _stackId.Set(value);
        }

        private static readonly __Value<bool?> _storeDashboardSha256 = new __Value<bool?>(() => __config.GetBoolean("storeDashboardSha256") ?? Utilities.GetEnvBoolean("GRAFANA_STORE_DASHBOARD_SHA256"));
        /// <summary>
        /// Set to true if you want to save only the sha256sum instead of complete dashboard model JSON in the tfstate.
        /// </summary>
        public static bool? StoreDashboardSha256
        {
            get => _storeDashboardSha256.Get();
            set => _storeDashboardSha256.Set(value);
        }

        private static readonly __Value<string?> _tlsCert = new __Value<string?>(() => __config.Get("tlsCert") ?? Utilities.GetEnv("GRAFANA_TLS_CERT"));
        /// <summary>
        /// Client TLS certificate (file path or literal value) to use to authenticate to the Grafana server. May alternatively be
        /// set via the `GRAFANA_TLS_CERT` environment variable.
        /// </summary>
        public static string? TlsCert
        {
            get => _tlsCert.Get();
            set => _tlsCert.Set(value);
        }

        private static readonly __Value<string?> _tlsKey = new __Value<string?>(() => __config.Get("tlsKey") ?? Utilities.GetEnv("GRAFANA_TLS_KEY"));
        /// <summary>
        /// Client TLS key (file path or literal value) to use to authenticate to the Grafana server. May alternatively be set via
        /// the `GRAFANA_TLS_KEY` environment variable.
        /// </summary>
        public static string? TlsKey
        {
            get => _tlsKey.Get();
            set => _tlsKey.Set(value);
        }

        private static readonly __Value<string?> _url = new __Value<string?>(() => __config.Get("url") ?? Utilities.GetEnv("GRAFANA_URL"));
        /// <summary>
        /// The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
        /// </summary>
        public static string? Url
        {
            get => _url.Get();
            set => _url.Set(value);
        }

    }
}
