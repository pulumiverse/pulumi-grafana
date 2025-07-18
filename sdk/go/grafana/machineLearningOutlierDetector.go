// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// An outlier detector monitors the results of a query and reports when its values are outside normal bands.
//
// The normal band is configured by choice of algorithm, its sensitivity and other configuration.
//
// Visit https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for more details.
//
// ## Example Usage
//
// ### DBSCAN Outlier Detector
//
// This outlier detector uses the DBSCAN algorithm to detect outliers.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/machinelearning"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := machinelearning.NewOutlierDetector(ctx, "my_dbscan_outlier_detector", &machinelearning.OutlierDetectorArgs{
//				Name:           pulumi.String("My DBSCAN outlier detector"),
//				Description:    pulumi.String("My DBSCAN Outlier Detector"),
//				Metric:         pulumi.String("tf_test_dbscan_job"),
//				DatasourceType: pulumi.String("prometheus"),
//				DatasourceUid:  pulumi.String("AbCd12345"),
//				QueryParams: pulumi.StringMap{
//					"expr": pulumi.String("grafanacloud_grafana_instance_active_user_count"),
//				},
//				Interval: pulumi.Int(300),
//				Algorithm: &machinelearning.OutlierDetectorAlgorithmArgs{
//					Name:        pulumi.String("dbscan"),
//					Sensitivity: pulumi.Float64(0.5),
//					Config: &machinelearning.OutlierDetectorAlgorithmConfigArgs{
//						Epsilon: pulumi.Float64(1),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ### MAD Outlier Detector
//
// This outlier detector uses the Median Absolute Deviation (MAD) algorithm to detect outliers.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/machinelearning"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := machinelearning.NewOutlierDetector(ctx, "my_mad_outlier_detector", &machinelearning.OutlierDetectorArgs{
//				Name:           pulumi.String("My MAD outlier detector"),
//				Description:    pulumi.String("My MAD Outlier Detector"),
//				Metric:         pulumi.String("tf_test_mad_job"),
//				DatasourceType: pulumi.String("prometheus"),
//				DatasourceUid:  pulumi.String("AbCd12345"),
//				QueryParams: pulumi.StringMap{
//					"expr": pulumi.String("grafanacloud_grafana_instance_active_user_count"),
//				},
//				Interval: pulumi.Int(300),
//				Algorithm: &machinelearning.OutlierDetectorAlgorithmArgs{
//					Name:        pulumi.String("mad"),
//					Sensitivity: pulumi.Float64(0.7),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
// $ pulumi import grafana:index/machineLearningOutlierDetector:MachineLearningOutlierDetector name "{{ id }}"
// ```
//
// Deprecated: grafana.index/machinelearningoutlierdetector.MachineLearningOutlierDetector has been deprecated in favor of grafana.machinelearning/outlierdetector.OutlierDetector
type MachineLearningOutlierDetector struct {
	pulumi.CustomResourceState

	// The algorithm to use and its configuration. See
	// https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
	Algorithm MachineLearningOutlierDetectorAlgorithmOutput `pulumi:"algorithm"`
	// The type of datasource being queried. Currently allowed values are prometheus, grafana-prometheus-datasource,
	// grafana-amazonprometheus-datasource, loki, grafana-loki-datasource, graphite, grafana-graphite-datasource,
	// grafana-datadog-datasource, postgres, grafana-postgresql-datasource, doitintl-bigquery-datasource,
	// grafana-bigquery-datasource, grafana-snowflake-datasource, influxdb, grafana-influxdb-datasource,
	// grafana-splunk-datasource, elasticsearch, grafana-elasticsearch-datasource, and grafana-mongodb-datasource.
	DatasourceType pulumi.StringOutput `pulumi:"datasourceType"`
	// The uid of the datasource to query.
	DatasourceUid pulumi.StringOutput `pulumi:"datasourceUid"`
	// A description of the outlier detector.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The data interval in seconds to monitor.
	Interval pulumi.IntPtrOutput `pulumi:"interval"`
	// The metric used to query the outlier detector results.
	Metric pulumi.StringOutput `pulumi:"metric"`
	// The name of the outlier detector.
	Name pulumi.StringOutput `pulumi:"name"`
	// An object representing the query params to query Grafana with.
	QueryParams pulumi.StringMapOutput `pulumi:"queryParams"`
}

// NewMachineLearningOutlierDetector registers a new resource with the given unique name, arguments, and options.
func NewMachineLearningOutlierDetector(ctx *pulumi.Context,
	name string, args *MachineLearningOutlierDetectorArgs, opts ...pulumi.ResourceOption) (*MachineLearningOutlierDetector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Algorithm == nil {
		return nil, errors.New("invalid value for required argument 'Algorithm'")
	}
	if args.DatasourceType == nil {
		return nil, errors.New("invalid value for required argument 'DatasourceType'")
	}
	if args.DatasourceUid == nil {
		return nil, errors.New("invalid value for required argument 'DatasourceUid'")
	}
	if args.Metric == nil {
		return nil, errors.New("invalid value for required argument 'Metric'")
	}
	if args.QueryParams == nil {
		return nil, errors.New("invalid value for required argument 'QueryParams'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MachineLearningOutlierDetector
	err := ctx.RegisterResource("grafana:index/machineLearningOutlierDetector:MachineLearningOutlierDetector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineLearningOutlierDetector gets an existing MachineLearningOutlierDetector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineLearningOutlierDetector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineLearningOutlierDetectorState, opts ...pulumi.ResourceOption) (*MachineLearningOutlierDetector, error) {
	var resource MachineLearningOutlierDetector
	err := ctx.ReadResource("grafana:index/machineLearningOutlierDetector:MachineLearningOutlierDetector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineLearningOutlierDetector resources.
type machineLearningOutlierDetectorState struct {
	// The algorithm to use and its configuration. See
	// https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
	Algorithm *MachineLearningOutlierDetectorAlgorithm `pulumi:"algorithm"`
	// The type of datasource being queried. Currently allowed values are prometheus, grafana-prometheus-datasource,
	// grafana-amazonprometheus-datasource, loki, grafana-loki-datasource, graphite, grafana-graphite-datasource,
	// grafana-datadog-datasource, postgres, grafana-postgresql-datasource, doitintl-bigquery-datasource,
	// grafana-bigquery-datasource, grafana-snowflake-datasource, influxdb, grafana-influxdb-datasource,
	// grafana-splunk-datasource, elasticsearch, grafana-elasticsearch-datasource, and grafana-mongodb-datasource.
	DatasourceType *string `pulumi:"datasourceType"`
	// The uid of the datasource to query.
	DatasourceUid *string `pulumi:"datasourceUid"`
	// A description of the outlier detector.
	Description *string `pulumi:"description"`
	// The data interval in seconds to monitor.
	Interval *int `pulumi:"interval"`
	// The metric used to query the outlier detector results.
	Metric *string `pulumi:"metric"`
	// The name of the outlier detector.
	Name *string `pulumi:"name"`
	// An object representing the query params to query Grafana with.
	QueryParams map[string]string `pulumi:"queryParams"`
}

type MachineLearningOutlierDetectorState struct {
	// The algorithm to use and its configuration. See
	// https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
	Algorithm MachineLearningOutlierDetectorAlgorithmPtrInput
	// The type of datasource being queried. Currently allowed values are prometheus, grafana-prometheus-datasource,
	// grafana-amazonprometheus-datasource, loki, grafana-loki-datasource, graphite, grafana-graphite-datasource,
	// grafana-datadog-datasource, postgres, grafana-postgresql-datasource, doitintl-bigquery-datasource,
	// grafana-bigquery-datasource, grafana-snowflake-datasource, influxdb, grafana-influxdb-datasource,
	// grafana-splunk-datasource, elasticsearch, grafana-elasticsearch-datasource, and grafana-mongodb-datasource.
	DatasourceType pulumi.StringPtrInput
	// The uid of the datasource to query.
	DatasourceUid pulumi.StringPtrInput
	// A description of the outlier detector.
	Description pulumi.StringPtrInput
	// The data interval in seconds to monitor.
	Interval pulumi.IntPtrInput
	// The metric used to query the outlier detector results.
	Metric pulumi.StringPtrInput
	// The name of the outlier detector.
	Name pulumi.StringPtrInput
	// An object representing the query params to query Grafana with.
	QueryParams pulumi.StringMapInput
}

func (MachineLearningOutlierDetectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineLearningOutlierDetectorState)(nil)).Elem()
}

type machineLearningOutlierDetectorArgs struct {
	// The algorithm to use and its configuration. See
	// https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
	Algorithm MachineLearningOutlierDetectorAlgorithm `pulumi:"algorithm"`
	// The type of datasource being queried. Currently allowed values are prometheus, grafana-prometheus-datasource,
	// grafana-amazonprometheus-datasource, loki, grafana-loki-datasource, graphite, grafana-graphite-datasource,
	// grafana-datadog-datasource, postgres, grafana-postgresql-datasource, doitintl-bigquery-datasource,
	// grafana-bigquery-datasource, grafana-snowflake-datasource, influxdb, grafana-influxdb-datasource,
	// grafana-splunk-datasource, elasticsearch, grafana-elasticsearch-datasource, and grafana-mongodb-datasource.
	DatasourceType string `pulumi:"datasourceType"`
	// The uid of the datasource to query.
	DatasourceUid string `pulumi:"datasourceUid"`
	// A description of the outlier detector.
	Description *string `pulumi:"description"`
	// The data interval in seconds to monitor.
	Interval *int `pulumi:"interval"`
	// The metric used to query the outlier detector results.
	Metric string `pulumi:"metric"`
	// The name of the outlier detector.
	Name *string `pulumi:"name"`
	// An object representing the query params to query Grafana with.
	QueryParams map[string]string `pulumi:"queryParams"`
}

// The set of arguments for constructing a MachineLearningOutlierDetector resource.
type MachineLearningOutlierDetectorArgs struct {
	// The algorithm to use and its configuration. See
	// https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
	Algorithm MachineLearningOutlierDetectorAlgorithmInput
	// The type of datasource being queried. Currently allowed values are prometheus, grafana-prometheus-datasource,
	// grafana-amazonprometheus-datasource, loki, grafana-loki-datasource, graphite, grafana-graphite-datasource,
	// grafana-datadog-datasource, postgres, grafana-postgresql-datasource, doitintl-bigquery-datasource,
	// grafana-bigquery-datasource, grafana-snowflake-datasource, influxdb, grafana-influxdb-datasource,
	// grafana-splunk-datasource, elasticsearch, grafana-elasticsearch-datasource, and grafana-mongodb-datasource.
	DatasourceType pulumi.StringInput
	// The uid of the datasource to query.
	DatasourceUid pulumi.StringInput
	// A description of the outlier detector.
	Description pulumi.StringPtrInput
	// The data interval in seconds to monitor.
	Interval pulumi.IntPtrInput
	// The metric used to query the outlier detector results.
	Metric pulumi.StringInput
	// The name of the outlier detector.
	Name pulumi.StringPtrInput
	// An object representing the query params to query Grafana with.
	QueryParams pulumi.StringMapInput
}

func (MachineLearningOutlierDetectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineLearningOutlierDetectorArgs)(nil)).Elem()
}

type MachineLearningOutlierDetectorInput interface {
	pulumi.Input

	ToMachineLearningOutlierDetectorOutput() MachineLearningOutlierDetectorOutput
	ToMachineLearningOutlierDetectorOutputWithContext(ctx context.Context) MachineLearningOutlierDetectorOutput
}

func (*MachineLearningOutlierDetector) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineLearningOutlierDetector)(nil)).Elem()
}

func (i *MachineLearningOutlierDetector) ToMachineLearningOutlierDetectorOutput() MachineLearningOutlierDetectorOutput {
	return i.ToMachineLearningOutlierDetectorOutputWithContext(context.Background())
}

func (i *MachineLearningOutlierDetector) ToMachineLearningOutlierDetectorOutputWithContext(ctx context.Context) MachineLearningOutlierDetectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningOutlierDetectorOutput)
}

// MachineLearningOutlierDetectorArrayInput is an input type that accepts MachineLearningOutlierDetectorArray and MachineLearningOutlierDetectorArrayOutput values.
// You can construct a concrete instance of `MachineLearningOutlierDetectorArrayInput` via:
//
//	MachineLearningOutlierDetectorArray{ MachineLearningOutlierDetectorArgs{...} }
type MachineLearningOutlierDetectorArrayInput interface {
	pulumi.Input

	ToMachineLearningOutlierDetectorArrayOutput() MachineLearningOutlierDetectorArrayOutput
	ToMachineLearningOutlierDetectorArrayOutputWithContext(context.Context) MachineLearningOutlierDetectorArrayOutput
}

type MachineLearningOutlierDetectorArray []MachineLearningOutlierDetectorInput

func (MachineLearningOutlierDetectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MachineLearningOutlierDetector)(nil)).Elem()
}

func (i MachineLearningOutlierDetectorArray) ToMachineLearningOutlierDetectorArrayOutput() MachineLearningOutlierDetectorArrayOutput {
	return i.ToMachineLearningOutlierDetectorArrayOutputWithContext(context.Background())
}

func (i MachineLearningOutlierDetectorArray) ToMachineLearningOutlierDetectorArrayOutputWithContext(ctx context.Context) MachineLearningOutlierDetectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningOutlierDetectorArrayOutput)
}

// MachineLearningOutlierDetectorMapInput is an input type that accepts MachineLearningOutlierDetectorMap and MachineLearningOutlierDetectorMapOutput values.
// You can construct a concrete instance of `MachineLearningOutlierDetectorMapInput` via:
//
//	MachineLearningOutlierDetectorMap{ "key": MachineLearningOutlierDetectorArgs{...} }
type MachineLearningOutlierDetectorMapInput interface {
	pulumi.Input

	ToMachineLearningOutlierDetectorMapOutput() MachineLearningOutlierDetectorMapOutput
	ToMachineLearningOutlierDetectorMapOutputWithContext(context.Context) MachineLearningOutlierDetectorMapOutput
}

type MachineLearningOutlierDetectorMap map[string]MachineLearningOutlierDetectorInput

func (MachineLearningOutlierDetectorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MachineLearningOutlierDetector)(nil)).Elem()
}

func (i MachineLearningOutlierDetectorMap) ToMachineLearningOutlierDetectorMapOutput() MachineLearningOutlierDetectorMapOutput {
	return i.ToMachineLearningOutlierDetectorMapOutputWithContext(context.Background())
}

func (i MachineLearningOutlierDetectorMap) ToMachineLearningOutlierDetectorMapOutputWithContext(ctx context.Context) MachineLearningOutlierDetectorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningOutlierDetectorMapOutput)
}

type MachineLearningOutlierDetectorOutput struct{ *pulumi.OutputState }

func (MachineLearningOutlierDetectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineLearningOutlierDetector)(nil)).Elem()
}

func (o MachineLearningOutlierDetectorOutput) ToMachineLearningOutlierDetectorOutput() MachineLearningOutlierDetectorOutput {
	return o
}

func (o MachineLearningOutlierDetectorOutput) ToMachineLearningOutlierDetectorOutputWithContext(ctx context.Context) MachineLearningOutlierDetectorOutput {
	return o
}

// The algorithm to use and its configuration. See
// https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
func (o MachineLearningOutlierDetectorOutput) Algorithm() MachineLearningOutlierDetectorAlgorithmOutput {
	return o.ApplyT(func(v *MachineLearningOutlierDetector) MachineLearningOutlierDetectorAlgorithmOutput {
		return v.Algorithm
	}).(MachineLearningOutlierDetectorAlgorithmOutput)
}

// The type of datasource being queried. Currently allowed values are prometheus, grafana-prometheus-datasource,
// grafana-amazonprometheus-datasource, loki, grafana-loki-datasource, graphite, grafana-graphite-datasource,
// grafana-datadog-datasource, postgres, grafana-postgresql-datasource, doitintl-bigquery-datasource,
// grafana-bigquery-datasource, grafana-snowflake-datasource, influxdb, grafana-influxdb-datasource,
// grafana-splunk-datasource, elasticsearch, grafana-elasticsearch-datasource, and grafana-mongodb-datasource.
func (o MachineLearningOutlierDetectorOutput) DatasourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineLearningOutlierDetector) pulumi.StringOutput { return v.DatasourceType }).(pulumi.StringOutput)
}

// The uid of the datasource to query.
func (o MachineLearningOutlierDetectorOutput) DatasourceUid() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineLearningOutlierDetector) pulumi.StringOutput { return v.DatasourceUid }).(pulumi.StringOutput)
}

// A description of the outlier detector.
func (o MachineLearningOutlierDetectorOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineLearningOutlierDetector) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The data interval in seconds to monitor.
func (o MachineLearningOutlierDetectorOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MachineLearningOutlierDetector) pulumi.IntPtrOutput { return v.Interval }).(pulumi.IntPtrOutput)
}

// The metric used to query the outlier detector results.
func (o MachineLearningOutlierDetectorOutput) Metric() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineLearningOutlierDetector) pulumi.StringOutput { return v.Metric }).(pulumi.StringOutput)
}

// The name of the outlier detector.
func (o MachineLearningOutlierDetectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineLearningOutlierDetector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// An object representing the query params to query Grafana with.
func (o MachineLearningOutlierDetectorOutput) QueryParams() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MachineLearningOutlierDetector) pulumi.StringMapOutput { return v.QueryParams }).(pulumi.StringMapOutput)
}

type MachineLearningOutlierDetectorArrayOutput struct{ *pulumi.OutputState }

func (MachineLearningOutlierDetectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MachineLearningOutlierDetector)(nil)).Elem()
}

func (o MachineLearningOutlierDetectorArrayOutput) ToMachineLearningOutlierDetectorArrayOutput() MachineLearningOutlierDetectorArrayOutput {
	return o
}

func (o MachineLearningOutlierDetectorArrayOutput) ToMachineLearningOutlierDetectorArrayOutputWithContext(ctx context.Context) MachineLearningOutlierDetectorArrayOutput {
	return o
}

func (o MachineLearningOutlierDetectorArrayOutput) Index(i pulumi.IntInput) MachineLearningOutlierDetectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MachineLearningOutlierDetector {
		return vs[0].([]*MachineLearningOutlierDetector)[vs[1].(int)]
	}).(MachineLearningOutlierDetectorOutput)
}

type MachineLearningOutlierDetectorMapOutput struct{ *pulumi.OutputState }

func (MachineLearningOutlierDetectorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MachineLearningOutlierDetector)(nil)).Elem()
}

func (o MachineLearningOutlierDetectorMapOutput) ToMachineLearningOutlierDetectorMapOutput() MachineLearningOutlierDetectorMapOutput {
	return o
}

func (o MachineLearningOutlierDetectorMapOutput) ToMachineLearningOutlierDetectorMapOutputWithContext(ctx context.Context) MachineLearningOutlierDetectorMapOutput {
	return o
}

func (o MachineLearningOutlierDetectorMapOutput) MapIndex(k pulumi.StringInput) MachineLearningOutlierDetectorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MachineLearningOutlierDetector {
		return vs[0].(map[string]*MachineLearningOutlierDetector)[vs[1].(string)]
	}).(MachineLearningOutlierDetectorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MachineLearningOutlierDetectorInput)(nil)).Elem(), &MachineLearningOutlierDetector{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineLearningOutlierDetectorArrayInput)(nil)).Elem(), MachineLearningOutlierDetectorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineLearningOutlierDetectorMapInput)(nil)).Elem(), MachineLearningOutlierDetectorMap{})
	pulumi.RegisterOutputType(MachineLearningOutlierDetectorOutput{})
	pulumi.RegisterOutputType(MachineLearningOutlierDetectorArrayOutput{})
	pulumi.RegisterOutputType(MachineLearningOutlierDetectorMapOutput{})
}
