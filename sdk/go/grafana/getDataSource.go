// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// Get details about a Grafana Datasource querying by either name, uid or ID
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"httpMethod":        "POST",
//				"prometheusType":    "Mimir",
//				"prometheusVersion": "2.4.0",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			tmpJSON1, err := json.Marshal(map[string]interface{}{
//				"basicAuthPassword": "password",
//			})
//			if err != nil {
//				return err
//			}
//			json1 := string(tmpJSON1)
//			prometheus, err := oss.NewDataSource(ctx, "prometheus", &oss.DataSourceArgs{
//				Type:                  pulumi.String("prometheus"),
//				Name:                  pulumi.String("prometheus-ds-test"),
//				Uid:                   pulumi.String("prometheus-ds-test-uid"),
//				Url:                   pulumi.String("https://my-instance.com"),
//				BasicAuthEnabled:      pulumi.Bool(true),
//				BasicAuthUsername:     pulumi.String("username"),
//				JsonDataEncoded:       pulumi.String(json0),
//				SecureJsonDataEncoded: pulumi.String(json1),
//			})
//			if err != nil {
//				return err
//			}
//			_ = oss.LookupDataSourceOutput(ctx, oss.GetDataSourceOutputArgs{
//				Name: prometheus.Name,
//			}, nil)
//			_ = oss.LookupDataSourceOutput(ctx, oss.GetDataSourceOutputArgs{
//				Uid: prometheus.Uid,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: grafana.index/getdatasource.getDataSource has been deprecated in favor of grafana.oss/getdatasource.getDataSource
func LookupDataSource(ctx *pulumi.Context, args *LookupDataSourceArgs, opts ...pulumi.InvokeOption) (*LookupDataSourceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDataSourceResult
	err := ctx.Invoke("grafana:index/getDataSource:getDataSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDataSource.
type LookupDataSourceArgs struct {
	Name *string `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	Uid   *string `pulumi:"uid"`
}

// A collection of values returned by getDataSource.
type LookupDataSourceResult struct {
	// The method by which Grafana will access the data source: `proxy` or `direct`.
	AccessMode string `pulumi:"accessMode"`
	// Whether to enable basic auth for the data source.
	BasicAuthEnabled bool `pulumi:"basicAuthEnabled"`
	// Basic auth username.
	BasicAuthUsername string `pulumi:"basicAuthUsername"`
	// (Required by some data source types) The name of the database to use on the selected data source server.
	DatabaseName string `pulumi:"databaseName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Whether to set the data source as default. This should only be `true` to a single data source.
	IsDefault bool `pulumi:"isDefault"`
	// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	JsonDataEncoded string `pulumi:"jsonDataEncoded"`
	Name            string `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// (Can only be used with data sources in Grafana Cloud) The ID of the Private Data source Connect network to use with this data source.
	PrivateDataSourceConnectNetworkId string `pulumi:"privateDataSourceConnectNetworkId"`
	// The data source type. Must be one of the supported data source keywords.
	Type string `pulumi:"type"`
	Uid  string `pulumi:"uid"`
	// The URL for the data source. The type of URL required varies depending on the chosen data source type.
	Url string `pulumi:"url"`
	// (Required by some data source types) The username to use to authenticate to the data source.
	Username string `pulumi:"username"`
}

func LookupDataSourceOutput(ctx *pulumi.Context, args LookupDataSourceOutputArgs, opts ...pulumi.InvokeOption) LookupDataSourceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDataSourceResultOutput, error) {
			args := v.(LookupDataSourceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("grafana:index/getDataSource:getDataSource", args, LookupDataSourceResultOutput{}, options).(LookupDataSourceResultOutput), nil
		}).(LookupDataSourceResultOutput)
}

// A collection of arguments for invoking getDataSource.
type LookupDataSourceOutputArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
	Uid   pulumi.StringPtrInput `pulumi:"uid"`
}

func (LookupDataSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataSourceArgs)(nil)).Elem()
}

// A collection of values returned by getDataSource.
type LookupDataSourceResultOutput struct{ *pulumi.OutputState }

func (LookupDataSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataSourceResult)(nil)).Elem()
}

func (o LookupDataSourceResultOutput) ToLookupDataSourceResultOutput() LookupDataSourceResultOutput {
	return o
}

func (o LookupDataSourceResultOutput) ToLookupDataSourceResultOutputWithContext(ctx context.Context) LookupDataSourceResultOutput {
	return o
}

// The method by which Grafana will access the data source: `proxy` or `direct`.
func (o LookupDataSourceResultOutput) AccessMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.AccessMode }).(pulumi.StringOutput)
}

// Whether to enable basic auth for the data source.
func (o LookupDataSourceResultOutput) BasicAuthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDataSourceResult) bool { return v.BasicAuthEnabled }).(pulumi.BoolOutput)
}

// Basic auth username.
func (o LookupDataSourceResultOutput) BasicAuthUsername() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.BasicAuthUsername }).(pulumi.StringOutput)
}

// (Required by some data source types) The name of the database to use on the selected data source server.
func (o LookupDataSourceResultOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.DatabaseName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDataSourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether to set the data source as default. This should only be `true` to a single data source.
func (o LookupDataSourceResultOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDataSourceResult) bool { return v.IsDefault }).(pulumi.BoolOutput)
}

// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
func (o LookupDataSourceResultOutput) JsonDataEncoded() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.JsonDataEncoded }).(pulumi.StringOutput)
}

func (o LookupDataSourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o LookupDataSourceResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// (Can only be used with data sources in Grafana Cloud) The ID of the Private Data source Connect network to use with this data source.
func (o LookupDataSourceResultOutput) PrivateDataSourceConnectNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.PrivateDataSourceConnectNetworkId }).(pulumi.StringOutput)
}

// The data source type. Must be one of the supported data source keywords.
func (o LookupDataSourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDataSourceResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.Uid }).(pulumi.StringOutput)
}

// The URL for the data source. The type of URL required varies depending on the chosen data source type.
func (o LookupDataSourceResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.Url }).(pulumi.StringOutput)
}

// (Required by some data source types) The username to use to authenticate to the data source.
func (o LookupDataSourceResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataSourceResultOutput{})
}
