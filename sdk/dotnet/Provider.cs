// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana
{
    /// <summary>
    /// The provider type for the grafana package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [GrafanaResourceType("pulumi:providers:grafana")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// API token, basic auth in the `username:password` format or `anonymous` (string literal). May alternatively be set via
        /// the `GRAFANA_AUTH` environment variable.
        /// </summary>
        [Output("auth")]
        public Output<string?> Auth { get; private set; } = null!;

        /// <summary>
        /// Certificate CA bundle (file path or literal value) to use to verify the Grafana server's certificate. May alternatively
        /// be set via the `GRAFANA_CA_CERT` environment variable.
        /// </summary>
        [Output("caCert")]
        public Output<string?> CaCert { get; private set; } = null!;

        /// <summary>
        /// Access Policy Token for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_ACCESS_POLICY_TOKEN` environment
        /// variable.
        /// </summary>
        [Output("cloudAccessPolicyToken")]
        public Output<string?> CloudAccessPolicyToken { get; private set; } = null!;

        /// <summary>
        /// Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
        /// </summary>
        [Output("cloudApiUrl")]
        public Output<string?> CloudApiUrl { get; private set; } = null!;

        /// <summary>
        /// A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
        /// </summary>
        [Output("oncallAccessToken")]
        public Output<string?> OncallAccessToken { get; private set; } = null!;

        /// <summary>
        /// An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
        /// </summary>
        [Output("oncallUrl")]
        public Output<string?> OncallUrl { get; private set; } = null!;

        /// <summary>
        /// A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
        /// </summary>
        [Output("smAccessToken")]
        public Output<string?> SmAccessToken { get; private set; } = null!;

        [Output("smUrl")]
        public Output<string?> SmUrl { get; private set; } = null!;

        /// <summary>
        /// Client TLS certificate (file path or literal value) to use to authenticate to the Grafana server. May alternatively be
        /// set via the `GRAFANA_TLS_CERT` environment variable.
        /// </summary>
        [Output("tlsCert")]
        public Output<string?> TlsCert { get; private set; } = null!;

        /// <summary>
        /// Client TLS key (file path or literal value) to use to authenticate to the Grafana server. May alternatively be set via
        /// the `GRAFANA_TLS_KEY` environment variable.
        /// </summary>
        [Output("tlsKey")]
        public Output<string?> TlsKey { get; private set; } = null!;

        /// <summary>
        /// The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
        /// </summary>
        [Output("url")]
        public Output<string?> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("grafana", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                AdditionalSecretOutputs =
                {
                    "auth",
                    "cloudAccessPolicyToken",
                    "oncallAccessToken",
                    "smAccessToken",
                    "tlsKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        [Input("auth")]
        private Input<string>? _auth;

        /// <summary>
        /// API token, basic auth in the `username:password` format or `anonymous` (string literal). May alternatively be set via
        /// the `GRAFANA_AUTH` environment variable.
        /// </summary>
        public Input<string>? Auth
        {
            get => _auth;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _auth = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Certificate CA bundle (file path or literal value) to use to verify the Grafana server's certificate. May alternatively
        /// be set via the `GRAFANA_CA_CERT` environment variable.
        /// </summary>
        [Input("caCert")]
        public Input<string>? CaCert { get; set; }

        [Input("cloudAccessPolicyToken")]
        private Input<string>? _cloudAccessPolicyToken;

        /// <summary>
        /// Access Policy Token for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_ACCESS_POLICY_TOKEN` environment
        /// variable.
        /// </summary>
        public Input<string>? CloudAccessPolicyToken
        {
            get => _cloudAccessPolicyToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _cloudAccessPolicyToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
        /// </summary>
        [Input("cloudApiUrl")]
        public Input<string>? CloudApiUrl { get; set; }

        /// <summary>
        /// Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
        /// </summary>
        [Input("insecureSkipVerify", json: true)]
        public Input<bool>? InsecureSkipVerify { get; set; }

        [Input("oncallAccessToken")]
        private Input<string>? _oncallAccessToken;

        /// <summary>
        /// A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
        /// </summary>
        public Input<string>? OncallAccessToken
        {
            get => _oncallAccessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _oncallAccessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
        /// </summary>
        [Input("oncallUrl")]
        public Input<string>? OncallUrl { get; set; }

        /// <summary>
        /// The amount of retries to use for Grafana API and Grafana Cloud API calls. May alternatively be set via the
        /// `GRAFANA_RETRIES` environment variable.
        /// </summary>
        [Input("retries", json: true)]
        public Input<int>? Retries { get; set; }

        [Input("retryStatusCodes", json: true)]
        private InputList<string>? _retryStatusCodes;

        /// <summary>
        /// The status codes to retry on for Grafana API and Grafana Cloud API calls. Use `x` as a digit wildcard. Defaults to 429
        /// and 5xx. May alternatively be set via the `GRAFANA_RETRY_STATUS_CODES` environment variable.
        /// </summary>
        public InputList<string> RetryStatusCodes
        {
            get => _retryStatusCodes ?? (_retryStatusCodes = new InputList<string>());
            set => _retryStatusCodes = value;
        }

        /// <summary>
        /// The amount of time in seconds to wait between retries for Grafana API and Grafana Cloud API calls. May alternatively be
        /// set via the `GRAFANA_RETRY_WAIT` environment variable.
        /// </summary>
        [Input("retryWait", json: true)]
        public Input<int>? RetryWait { get; set; }

        [Input("smAccessToken")]
        private Input<string>? _smAccessToken;

        /// <summary>
        /// A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
        /// </summary>
        public Input<string>? SmAccessToken
        {
            get => _smAccessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _smAccessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("smUrl")]
        public Input<string>? SmUrl { get; set; }

        /// <summary>
        /// Set to true if you want to save only the sha256sum instead of complete dashboard model JSON in the tfstate.
        /// </summary>
        [Input("storeDashboardSha256", json: true)]
        public Input<bool>? StoreDashboardSha256 { get; set; }

        /// <summary>
        /// Client TLS certificate (file path or literal value) to use to authenticate to the Grafana server. May alternatively be
        /// set via the `GRAFANA_TLS_CERT` environment variable.
        /// </summary>
        [Input("tlsCert")]
        public Input<string>? TlsCert { get; set; }

        [Input("tlsKey")]
        private Input<string>? _tlsKey;

        /// <summary>
        /// Client TLS key (file path or literal value) to use to authenticate to the Grafana server. May alternatively be set via
        /// the `GRAFANA_TLS_KEY` environment variable.
        /// </summary>
        public Input<string>? TlsKey
        {
            get => _tlsKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _tlsKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ProviderArgs()
        {
            Auth = Utilities.GetEnv("GRAFANA_AUTH");
            CaCert = Utilities.GetEnv("GRAFANA_CA_CERT");
            CloudApiUrl = Utilities.GetEnv("GRAFANA_CLOUD_API_URL");
            InsecureSkipVerify = Utilities.GetEnvBoolean("GRAFANA_INSECURE_SKIP_VERIFY");
            OncallAccessToken = Utilities.GetEnv("GRAFANA_ONCALL_ACCESS_TOKEN");
            OncallUrl = Utilities.GetEnv("GRAFANA_ONCALL_URL");
            Retries = Utilities.GetEnvInt32("GRAFANA_RETRIES");
            SmAccessToken = Utilities.GetEnv("GRAFANA_SM_ACCESS_TOKEN");
            SmUrl = Utilities.GetEnv("GRAFANA_SM_URL");
            StoreDashboardSha256 = Utilities.GetEnvBoolean("GRAFANA_STORE_DASHBOARD_SHA256");
            TlsCert = Utilities.GetEnv("GRAFANA_TLS_CERT");
            TlsKey = Utilities.GetEnv("GRAFANA_TLS_KEY");
            Url = Utilities.GetEnv("GRAFANA_URL");
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}
