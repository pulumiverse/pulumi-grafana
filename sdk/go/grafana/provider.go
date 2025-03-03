// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// The provider type for the grafana package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// API token, basic auth in the `username:password` format or `anonymous` (string literal). May alternatively be set via
	// the `GRAFANA_AUTH` environment variable.
	Auth pulumi.StringPtrOutput `pulumi:"auth"`
	// Certificate CA bundle (file path or literal value) to use to verify the Grafana server's certificate. May alternatively
	// be set via the `GRAFANA_CA_CERT` environment variable.
	CaCert pulumi.StringPtrOutput `pulumi:"caCert"`
	// Access Policy Token for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_ACCESS_POLICY_TOKEN` environment
	// variable.
	CloudAccessPolicyToken pulumi.StringPtrOutput `pulumi:"cloudAccessPolicyToken"`
	// Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
	CloudApiUrl pulumi.StringPtrOutput `pulumi:"cloudApiUrl"`
	// A Grafana Cloud Provider access token. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_ACCESS_TOKEN`
	// environment variable.
	CloudProviderAccessToken pulumi.StringPtrOutput `pulumi:"cloudProviderAccessToken"`
	// A Grafana Cloud Provider backend address. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_URL` environment
	// variable.
	CloudProviderUrl pulumi.StringPtrOutput `pulumi:"cloudProviderUrl"`
	// A Grafana Connections API access token. May alternatively be set via the `GRAFANA_CONNECTIONS_API_ACCESS_TOKEN`
	// environment variable.
	ConnectionsApiAccessToken pulumi.StringPtrOutput `pulumi:"connectionsApiAccessToken"`
	// A Grafana Connections API address. May alternatively be set via the `GRAFANA_CONNECTIONS_API_URL` environment variable.
	ConnectionsApiUrl pulumi.StringPtrOutput `pulumi:"connectionsApiUrl"`
	// A Grafana Fleet Management basic auth in the `username:password` format. May alternatively be set via the
	// `GRAFANA_FLEET_MANAGEMENT_AUTH` environment variable.
	FleetManagementAuth pulumi.StringPtrOutput `pulumi:"fleetManagementAuth"`
	// A Grafana Fleet Management API address. May alternatively be set via the `GRAFANA_FLEET_MANAGEMENT_URL` environment
	// variable.
	FleetManagementUrl pulumi.StringPtrOutput `pulumi:"fleetManagementUrl"`
	// A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
	OncallAccessToken pulumi.StringPtrOutput `pulumi:"oncallAccessToken"`
	// An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
	OncallUrl pulumi.StringPtrOutput `pulumi:"oncallUrl"`
	// A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
	SmAccessToken pulumi.StringPtrOutput `pulumi:"smAccessToken"`
	SmUrl         pulumi.StringPtrOutput `pulumi:"smUrl"`
	// Client TLS certificate (file path or literal value) to use to authenticate to the Grafana server. May alternatively be
	// set via the `GRAFANA_TLS_CERT` environment variable.
	TlsCert pulumi.StringPtrOutput `pulumi:"tlsCert"`
	// Client TLS key (file path or literal value) to use to authenticate to the Grafana server. May alternatively be set via
	// the `GRAFANA_TLS_KEY` environment variable.
	TlsKey pulumi.StringPtrOutput `pulumi:"tlsKey"`
	// The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
	Url pulumi.StringPtrOutput `pulumi:"url"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.Auth == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_AUTH"); d != nil {
			args.Auth = pulumi.StringPtr(d.(string))
		}
	}
	if args.CaCert == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_CA_CERT"); d != nil {
			args.CaCert = pulumi.StringPtr(d.(string))
		}
	}
	if args.CloudAccessPolicyToken == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_CLOUD_ACCESS_POLICY_TOKEN"); d != nil {
			args.CloudAccessPolicyToken = pulumi.StringPtr(d.(string))
		}
	}
	if args.CloudApiUrl == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_CLOUD_API_URL"); d != nil {
			args.CloudApiUrl = pulumi.StringPtr(d.(string))
		}
	}
	if args.InsecureSkipVerify == nil {
		if d := internal.GetEnvOrDefault(nil, internal.ParseEnvBool, "GRAFANA_INSECURE_SKIP_VERIFY"); d != nil {
			args.InsecureSkipVerify = pulumi.BoolPtr(d.(bool))
		}
	}
	if args.OncallAccessToken == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_ONCALL_ACCESS_TOKEN"); d != nil {
			args.OncallAccessToken = pulumi.StringPtr(d.(string))
		}
	}
	if args.OncallUrl == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_ONCALL_URL"); d != nil {
			args.OncallUrl = pulumi.StringPtr(d.(string))
		}
	}
	if args.Retries == nil {
		if d := internal.GetEnvOrDefault(nil, internal.ParseEnvInt, "GRAFANA_RETRIES"); d != nil {
			args.Retries = pulumi.IntPtr(d.(int))
		}
	}
	if args.RetryWait == nil {
		if d := internal.GetEnvOrDefault(nil, internal.ParseEnvInt, "GRAFANA_RETRY_WAIT"); d != nil {
			args.RetryWait = pulumi.IntPtr(d.(int))
		}
	}
	if args.SmAccessToken == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_SM_ACCESS_TOKEN"); d != nil {
			args.SmAccessToken = pulumi.StringPtr(d.(string))
		}
	}
	if args.SmUrl == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_SM_URL"); d != nil {
			args.SmUrl = pulumi.StringPtr(d.(string))
		}
	}
	if args.StoreDashboardSha256 == nil {
		if d := internal.GetEnvOrDefault(nil, internal.ParseEnvBool, "GRAFANA_STORE_DASHBOARD_SHA256"); d != nil {
			args.StoreDashboardSha256 = pulumi.BoolPtr(d.(bool))
		}
	}
	if args.TlsCert == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_TLS_CERT"); d != nil {
			args.TlsCert = pulumi.StringPtr(d.(string))
		}
	}
	if args.TlsKey == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_TLS_KEY"); d != nil {
			args.TlsKey = pulumi.StringPtr(d.(string))
		}
	}
	if args.Url == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "GRAFANA_URL"); d != nil {
			args.Url = pulumi.StringPtr(d.(string))
		}
	}
	if args.Auth != nil {
		args.Auth = pulumi.ToSecret(args.Auth).(pulumi.StringPtrInput)
	}
	if args.CloudAccessPolicyToken != nil {
		args.CloudAccessPolicyToken = pulumi.ToSecret(args.CloudAccessPolicyToken).(pulumi.StringPtrInput)
	}
	if args.CloudProviderAccessToken != nil {
		args.CloudProviderAccessToken = pulumi.ToSecret(args.CloudProviderAccessToken).(pulumi.StringPtrInput)
	}
	if args.ConnectionsApiAccessToken != nil {
		args.ConnectionsApiAccessToken = pulumi.ToSecret(args.ConnectionsApiAccessToken).(pulumi.StringPtrInput)
	}
	if args.FleetManagementAuth != nil {
		args.FleetManagementAuth = pulumi.ToSecret(args.FleetManagementAuth).(pulumi.StringPtrInput)
	}
	if args.OncallAccessToken != nil {
		args.OncallAccessToken = pulumi.ToSecret(args.OncallAccessToken).(pulumi.StringPtrInput)
	}
	if args.SmAccessToken != nil {
		args.SmAccessToken = pulumi.ToSecret(args.SmAccessToken).(pulumi.StringPtrInput)
	}
	if args.TlsKey != nil {
		args.TlsKey = pulumi.ToSecret(args.TlsKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"auth",
		"cloudAccessPolicyToken",
		"cloudProviderAccessToken",
		"connectionsApiAccessToken",
		"fleetManagementAuth",
		"oncallAccessToken",
		"smAccessToken",
		"tlsKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:grafana", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// API token, basic auth in the `username:password` format or `anonymous` (string literal). May alternatively be set via
	// the `GRAFANA_AUTH` environment variable.
	Auth *string `pulumi:"auth"`
	// Certificate CA bundle (file path or literal value) to use to verify the Grafana server's certificate. May alternatively
	// be set via the `GRAFANA_CA_CERT` environment variable.
	CaCert *string `pulumi:"caCert"`
	// Access Policy Token for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_ACCESS_POLICY_TOKEN` environment
	// variable.
	CloudAccessPolicyToken *string `pulumi:"cloudAccessPolicyToken"`
	// Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
	CloudApiUrl *string `pulumi:"cloudApiUrl"`
	// A Grafana Cloud Provider access token. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_ACCESS_TOKEN`
	// environment variable.
	CloudProviderAccessToken *string `pulumi:"cloudProviderAccessToken"`
	// A Grafana Cloud Provider backend address. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_URL` environment
	// variable.
	CloudProviderUrl *string `pulumi:"cloudProviderUrl"`
	// A Grafana Connections API access token. May alternatively be set via the `GRAFANA_CONNECTIONS_API_ACCESS_TOKEN`
	// environment variable.
	ConnectionsApiAccessToken *string `pulumi:"connectionsApiAccessToken"`
	// A Grafana Connections API address. May alternatively be set via the `GRAFANA_CONNECTIONS_API_URL` environment variable.
	ConnectionsApiUrl *string `pulumi:"connectionsApiUrl"`
	// A Grafana Fleet Management basic auth in the `username:password` format. May alternatively be set via the
	// `GRAFANA_FLEET_MANAGEMENT_AUTH` environment variable.
	FleetManagementAuth *string `pulumi:"fleetManagementAuth"`
	// A Grafana Fleet Management API address. May alternatively be set via the `GRAFANA_FLEET_MANAGEMENT_URL` environment
	// variable.
	FleetManagementUrl *string `pulumi:"fleetManagementUrl"`
	// Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
	InsecureSkipVerify *bool `pulumi:"insecureSkipVerify"`
	// A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
	OncallAccessToken *string `pulumi:"oncallAccessToken"`
	// An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
	OncallUrl *string `pulumi:"oncallUrl"`
	// The amount of retries to use for Grafana API and Grafana Cloud API calls. May alternatively be set via the
	// `GRAFANA_RETRIES` environment variable.
	Retries *int `pulumi:"retries"`
	// The status codes to retry on for Grafana API and Grafana Cloud API calls. Use `x` as a digit wildcard. Defaults to 429
	// and 5xx. May alternatively be set via the `GRAFANA_RETRY_STATUS_CODES` environment variable.
	RetryStatusCodes []string `pulumi:"retryStatusCodes"`
	// The amount of time in seconds to wait between retries for Grafana API and Grafana Cloud API calls. May alternatively be
	// set via the `GRAFANA_RETRY_WAIT` environment variable.
	RetryWait *int `pulumi:"retryWait"`
	// A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
	SmAccessToken *string `pulumi:"smAccessToken"`
	SmUrl         *string `pulumi:"smUrl"`
	// Set to true if you want to save only the sha256sum instead of complete dashboard model JSON in the tfstate.
	StoreDashboardSha256 *bool `pulumi:"storeDashboardSha256"`
	// Client TLS certificate (file path or literal value) to use to authenticate to the Grafana server. May alternatively be
	// set via the `GRAFANA_TLS_CERT` environment variable.
	TlsCert *string `pulumi:"tlsCert"`
	// Client TLS key (file path or literal value) to use to authenticate to the Grafana server. May alternatively be set via
	// the `GRAFANA_TLS_KEY` environment variable.
	TlsKey *string `pulumi:"tlsKey"`
	// The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
	Url *string `pulumi:"url"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// API token, basic auth in the `username:password` format or `anonymous` (string literal). May alternatively be set via
	// the `GRAFANA_AUTH` environment variable.
	Auth pulumi.StringPtrInput
	// Certificate CA bundle (file path or literal value) to use to verify the Grafana server's certificate. May alternatively
	// be set via the `GRAFANA_CA_CERT` environment variable.
	CaCert pulumi.StringPtrInput
	// Access Policy Token for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_ACCESS_POLICY_TOKEN` environment
	// variable.
	CloudAccessPolicyToken pulumi.StringPtrInput
	// Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
	CloudApiUrl pulumi.StringPtrInput
	// A Grafana Cloud Provider access token. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_ACCESS_TOKEN`
	// environment variable.
	CloudProviderAccessToken pulumi.StringPtrInput
	// A Grafana Cloud Provider backend address. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_URL` environment
	// variable.
	CloudProviderUrl pulumi.StringPtrInput
	// A Grafana Connections API access token. May alternatively be set via the `GRAFANA_CONNECTIONS_API_ACCESS_TOKEN`
	// environment variable.
	ConnectionsApiAccessToken pulumi.StringPtrInput
	// A Grafana Connections API address. May alternatively be set via the `GRAFANA_CONNECTIONS_API_URL` environment variable.
	ConnectionsApiUrl pulumi.StringPtrInput
	// A Grafana Fleet Management basic auth in the `username:password` format. May alternatively be set via the
	// `GRAFANA_FLEET_MANAGEMENT_AUTH` environment variable.
	FleetManagementAuth pulumi.StringPtrInput
	// A Grafana Fleet Management API address. May alternatively be set via the `GRAFANA_FLEET_MANAGEMENT_URL` environment
	// variable.
	FleetManagementUrl pulumi.StringPtrInput
	// Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
	InsecureSkipVerify pulumi.BoolPtrInput
	// A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
	OncallAccessToken pulumi.StringPtrInput
	// An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
	OncallUrl pulumi.StringPtrInput
	// The amount of retries to use for Grafana API and Grafana Cloud API calls. May alternatively be set via the
	// `GRAFANA_RETRIES` environment variable.
	Retries pulumi.IntPtrInput
	// The status codes to retry on for Grafana API and Grafana Cloud API calls. Use `x` as a digit wildcard. Defaults to 429
	// and 5xx. May alternatively be set via the `GRAFANA_RETRY_STATUS_CODES` environment variable.
	RetryStatusCodes pulumi.StringArrayInput
	// The amount of time in seconds to wait between retries for Grafana API and Grafana Cloud API calls. May alternatively be
	// set via the `GRAFANA_RETRY_WAIT` environment variable.
	RetryWait pulumi.IntPtrInput
	// A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
	SmAccessToken pulumi.StringPtrInput
	SmUrl         pulumi.StringPtrInput
	// Set to true if you want to save only the sha256sum instead of complete dashboard model JSON in the tfstate.
	StoreDashboardSha256 pulumi.BoolPtrInput
	// Client TLS certificate (file path or literal value) to use to authenticate to the Grafana server. May alternatively be
	// set via the `GRAFANA_TLS_CERT` environment variable.
	TlsCert pulumi.StringPtrInput
	// Client TLS key (file path or literal value) to use to authenticate to the Grafana server. May alternatively be set via
	// the `GRAFANA_TLS_KEY` environment variable.
	TlsKey pulumi.StringPtrInput
	// The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
	Url pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// API token, basic auth in the `username:password` format or `anonymous` (string literal). May alternatively be set via
// the `GRAFANA_AUTH` environment variable.
func (o ProviderOutput) Auth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Auth }).(pulumi.StringPtrOutput)
}

// Certificate CA bundle (file path or literal value) to use to verify the Grafana server's certificate. May alternatively
// be set via the `GRAFANA_CA_CERT` environment variable.
func (o ProviderOutput) CaCert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CaCert }).(pulumi.StringPtrOutput)
}

// Access Policy Token for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_ACCESS_POLICY_TOKEN` environment
// variable.
func (o ProviderOutput) CloudAccessPolicyToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CloudAccessPolicyToken }).(pulumi.StringPtrOutput)
}

// Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
func (o ProviderOutput) CloudApiUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CloudApiUrl }).(pulumi.StringPtrOutput)
}

// A Grafana Cloud Provider access token. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_ACCESS_TOKEN`
// environment variable.
func (o ProviderOutput) CloudProviderAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CloudProviderAccessToken }).(pulumi.StringPtrOutput)
}

// A Grafana Cloud Provider backend address. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_URL` environment
// variable.
func (o ProviderOutput) CloudProviderUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CloudProviderUrl }).(pulumi.StringPtrOutput)
}

// A Grafana Connections API access token. May alternatively be set via the `GRAFANA_CONNECTIONS_API_ACCESS_TOKEN`
// environment variable.
func (o ProviderOutput) ConnectionsApiAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ConnectionsApiAccessToken }).(pulumi.StringPtrOutput)
}

// A Grafana Connections API address. May alternatively be set via the `GRAFANA_CONNECTIONS_API_URL` environment variable.
func (o ProviderOutput) ConnectionsApiUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ConnectionsApiUrl }).(pulumi.StringPtrOutput)
}

// A Grafana Fleet Management basic auth in the `username:password` format. May alternatively be set via the
// `GRAFANA_FLEET_MANAGEMENT_AUTH` environment variable.
func (o ProviderOutput) FleetManagementAuth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.FleetManagementAuth }).(pulumi.StringPtrOutput)
}

// A Grafana Fleet Management API address. May alternatively be set via the `GRAFANA_FLEET_MANAGEMENT_URL` environment
// variable.
func (o ProviderOutput) FleetManagementUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.FleetManagementUrl }).(pulumi.StringPtrOutput)
}

// A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
func (o ProviderOutput) OncallAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.OncallAccessToken }).(pulumi.StringPtrOutput)
}

// An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
func (o ProviderOutput) OncallUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.OncallUrl }).(pulumi.StringPtrOutput)
}

// A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
func (o ProviderOutput) SmAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.SmAccessToken }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) SmUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.SmUrl }).(pulumi.StringPtrOutput)
}

// Client TLS certificate (file path or literal value) to use to authenticate to the Grafana server. May alternatively be
// set via the `GRAFANA_TLS_CERT` environment variable.
func (o ProviderOutput) TlsCert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.TlsCert }).(pulumi.StringPtrOutput)
}

// Client TLS key (file path or literal value) to use to authenticate to the Grafana server. May alternatively be set via
// the `GRAFANA_TLS_KEY` environment variable.
func (o ProviderOutput) TlsKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.TlsKey }).(pulumi.StringPtrOutput)
}

// The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
func (o ProviderOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
