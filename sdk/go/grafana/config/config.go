// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

var _ = internal.GetEnvOrDefault

// API token, basic auth in the `username:password` format or `anonymous` (string literal). May alternatively be set via
// the `GRAFANA_AUTH` environment variable.
func GetAuth(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "grafana:auth")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_AUTH"); d != nil {
		value = d.(string)
	}
	return value
}

// Certificate CA bundle (file path or literal value) to use to verify the Grafana server's certificate. May alternatively
// be set via the `GRAFANA_CA_CERT` environment variable.
func GetCaCert(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "grafana:caCert")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_CA_CERT"); d != nil {
		value = d.(string)
	}
	return value
}

// Access Policy Token for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_ACCESS_POLICY_TOKEN` environment
// variable.
func GetCloudAccessPolicyToken(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "grafana:cloudAccessPolicyToken")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_CLOUD_ACCESS_POLICY_TOKEN"); d != nil {
		value = d.(string)
	}
	return value
}

// Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
func GetCloudApiUrl(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "grafana:cloudApiUrl")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_CLOUD_API_URL"); d != nil {
		value = d.(string)
	}
	return value
}

// A Grafana Cloud Provider access token. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_ACCESS_TOKEN`
// environment variable.
func GetCloudProviderAccessToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "grafana:cloudProviderAccessToken")
}

// A Grafana Cloud Provider backend address. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_URL` environment
// variable.
func GetCloudProviderUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "grafana:cloudProviderUrl")
}

// A Grafana Connections API access token. May alternatively be set via the `GRAFANA_CONNECTIONS_API_ACCESS_TOKEN`
// environment variable.
func GetConnectionsApiAccessToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "grafana:connectionsApiAccessToken")
}

// A Grafana Connections API address. May alternatively be set via the `GRAFANA_CONNECTIONS_API_URL` environment variable.
func GetConnectionsApiUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "grafana:connectionsApiUrl")
}

// A Grafana Fleet Management basic auth in the `username:password` format. May alternatively be set via the
// `GRAFANA_FLEET_MANAGEMENT_AUTH` environment variable.
func GetFleetManagementAuth(ctx *pulumi.Context) string {
	return config.Get(ctx, "grafana:fleetManagementAuth")
}

// A Grafana Fleet Management API address. May alternatively be set via the `GRAFANA_FLEET_MANAGEMENT_URL` environment
// variable.
func GetFleetManagementUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "grafana:fleetManagementUrl")
}

// Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
func GetInsecureSkipVerify(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "grafana:insecureSkipVerify")
	if err == nil {
		return v
	}
	var value bool
	if d := internal.GetEnvOrDefault(nil, internal.ParseEnvBool, "GRAFANA_INSECURE_SKIP_VERIFY"); d != nil {
		value = d.(bool)
	}
	return value
}

// A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
func GetOncallAccessToken(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "grafana:oncallAccessToken")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_ONCALL_ACCESS_TOKEN"); d != nil {
		value = d.(string)
	}
	return value
}

// An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
func GetOncallUrl(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "grafana:oncallUrl")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_ONCALL_URL"); d != nil {
		value = d.(string)
	}
	return value
}

// The amount of retries to use for Grafana API and Grafana Cloud API calls. May alternatively be set via the
// `GRAFANA_RETRIES` environment variable.
func GetRetries(ctx *pulumi.Context) int {
	v, err := config.TryInt(ctx, "grafana:retries")
	if err == nil {
		return v
	}
	var value int
	if d := internal.GetEnvOrDefault(nil, internal.ParseEnvInt, "GRAFANA_RETRIES"); d != nil {
		value = d.(int)
	}
	return value
}

// The status codes to retry on for Grafana API and Grafana Cloud API calls. Use `x` as a digit wildcard. Defaults to 429
// and 5xx. May alternatively be set via the `GRAFANA_RETRY_STATUS_CODES` environment variable.
func GetRetryStatusCodes(ctx *pulumi.Context) string {
	return config.Get(ctx, "grafana:retryStatusCodes")
}

// The amount of time in seconds to wait between retries for Grafana API and Grafana Cloud API calls. May alternatively be
// set via the `GRAFANA_RETRY_WAIT` environment variable.
func GetRetryWait(ctx *pulumi.Context) int {
	v, err := config.TryInt(ctx, "grafana:retryWait")
	if err == nil {
		return v
	}
	var value int
	if d := internal.GetEnvOrDefault(nil, internal.ParseEnvInt, "GRAFANA_RETRY_WAIT"); d != nil {
		value = d.(int)
	}
	return value
}

// A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
func GetSmAccessToken(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "grafana:smAccessToken")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_SM_ACCESS_TOKEN"); d != nil {
		value = d.(string)
	}
	return value
}
func GetSmUrl(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "grafana:smUrl")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_SM_URL"); d != nil {
		value = d.(string)
	}
	return value
}

// Set to true if you want to save only the sha256sum instead of complete dashboard model JSON in the tfstate.
func GetStoreDashboardSha256(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "grafana:storeDashboardSha256")
	if err == nil {
		return v
	}
	var value bool
	if d := internal.GetEnvOrDefault(nil, internal.ParseEnvBool, "GRAFANA_STORE_DASHBOARD_SHA256"); d != nil {
		value = d.(bool)
	}
	return value
}

// Client TLS certificate (file path or literal value) to use to authenticate to the Grafana server. May alternatively be
// set via the `GRAFANA_TLS_CERT` environment variable.
func GetTlsCert(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "grafana:tlsCert")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_TLS_CERT"); d != nil {
		value = d.(string)
	}
	return value
}

// Client TLS key (file path or literal value) to use to authenticate to the Grafana server. May alternatively be set via
// the `GRAFANA_TLS_KEY` environment variable.
func GetTlsKey(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "grafana:tlsKey")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_TLS_KEY"); d != nil {
		value = d.(string)
	}
	return value
}

// The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
func GetUrl(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "grafana:url")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_URL"); d != nil {
		value = d.(string)
	}
	return value
}
