// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearning

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// A job defines the queries and model parameters for a machine learning task.
//
// See [the Grafana Cloud docs](https://grafana.com/docs/grafana-cloud/alerting-and-irm/machine-learning/dynamic-alerting/forecasting/config/) for more information
// on available hyperparameters for use in the `hyperParams` field.
//
// ## Example Usage
//
// ### Basic Forecast
//
// This forecast uses a Prometheus datasource, where the source query is defined in the `expr` field of the `queryParams` attribute.
//
// Other datasources are supported, but the structure `queryParams` may differ.
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/machinelearning"
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
//			foo, err := oss.NewDataSource(ctx, "foo", &oss.DataSourceArgs{
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
//			_, err = machinelearning.NewJob(ctx, "test_job", &machinelearning.JobArgs{
//				Name:           pulumi.String("Test Job"),
//				Metric:         pulumi.String("tf_test_job"),
//				DatasourceType: pulumi.String("prometheus"),
//				DatasourceUid:  foo.Uid,
//				QueryParams: pulumi.StringMap{
//					"expr": pulumi.String("grafanacloud_grafana_instance_active_user_count"),
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
// ### Tuned Forecast
//
// This forecast has tuned hyperparameters to improve the accuracy of the model.
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/machinelearning"
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
//			foo, err := oss.NewDataSource(ctx, "foo", &oss.DataSourceArgs{
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
//			_, err = machinelearning.NewJob(ctx, "test_job", &machinelearning.JobArgs{
//				Name:           pulumi.String("Test Job"),
//				Metric:         pulumi.String("tf_test_job"),
//				DatasourceType: pulumi.String("prometheus"),
//				DatasourceUid:  foo.Uid,
//				QueryParams: pulumi.StringMap{
//					"expr": pulumi.String("grafanacloud_grafana_instance_active_user_count"),
//				},
//				HyperParams: pulumi.StringMap{
//					"daily_seasonality":  pulumi.String("15"),
//					"weekly_seasonality": pulumi.String("10"),
//				},
//				CustomLabels: pulumi.StringMap{
//					"example_label": pulumi.String("example_value"),
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
// ### Rescaled Forecast
//
// This forecast has had the data transformed using a power transformation in order to avoid negative lower predictions.
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/machinelearning"
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
//			foo, err := oss.NewDataSource(ctx, "foo", &oss.DataSourceArgs{
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
//			_, err = machinelearning.NewJob(ctx, "test_job", &machinelearning.JobArgs{
//				Name:           pulumi.String("Test Job"),
//				Metric:         pulumi.String("tf_test_job"),
//				DatasourceType: pulumi.String("prometheus"),
//				DatasourceUid:  foo.Uid,
//				QueryParams: pulumi.StringMap{
//					"expr": pulumi.String("grafanacloud_grafana_instance_active_user_count"),
//				},
//				HyperParams: pulumi.StringMap{
//					"transformation_id": pulumi.String("power"),
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
// ### Forecast with Holidays
//
// This forecast has holidays which will be taken into account when training the model.
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/machinelearning"
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
//			foo, err := oss.NewDataSource(ctx, "foo", &oss.DataSourceArgs{
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
//			testHoliday, err := machinelearning.NewHoliday(ctx, "test_holiday", &machinelearning.HolidayArgs{
//				Name: pulumi.String("Test Holiday"),
//				CustomPeriods: machinelearning.HolidayCustomPeriodArray{
//					&machinelearning.HolidayCustomPeriodArgs{
//						Name:      pulumi.String("First of January"),
//						StartTime: pulumi.String("2023-01-01T00:00:00Z"),
//						EndTime:   pulumi.String("2023-01-02T00:00:00Z"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = machinelearning.NewJob(ctx, "test_job", &machinelearning.JobArgs{
//				Name:           pulumi.String("Test Job"),
//				Metric:         pulumi.String("tf_test_job"),
//				DatasourceType: pulumi.String("prometheus"),
//				DatasourceUid:  foo.Uid,
//				QueryParams: pulumi.StringMap{
//					"expr": pulumi.String("grafanacloud_grafana_instance_active_user_count"),
//				},
//				Holidays: pulumi.StringArray{
//					testHoliday.ID(),
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
// $ pulumi import grafana:machineLearning/job:Job name "{{ id }}"
// ```
type Job struct {
	pulumi.CustomResourceState

	// An object representing the custom labels added on the forecast.
	CustomLabels pulumi.StringMapOutput `pulumi:"customLabels"`
	// The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
	DatasourceType pulumi.StringOutput `pulumi:"datasourceType"`
	// The uid of the datasource to query.
	DatasourceUid pulumi.StringOutput `pulumi:"datasourceUid"`
	// A description of the job.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A list of holiday IDs or names to take into account when training the model.
	Holidays pulumi.StringArrayOutput `pulumi:"holidays"`
	// The hyperparameters used to fine tune the algorithm. See
	// https://grafana.com/docs/grafana-cloud/alerting-and-irm/machine-learning/forecasts/models/ for the full list of
	// available hyperparameters.
	HyperParams pulumi.StringMapOutput `pulumi:"hyperParams"`
	// The data interval in seconds to train the data on.
	Interval pulumi.IntPtrOutput `pulumi:"interval"`
	// The metric used to query the job results.
	Metric pulumi.StringOutput `pulumi:"metric"`
	// The name of the job.
	Name pulumi.StringOutput `pulumi:"name"`
	// An object representing the query params to query Grafana with.
	QueryParams pulumi.StringMapOutput `pulumi:"queryParams"`
	// The data interval in seconds to train the data on.
	TrainingWindow pulumi.IntPtrOutput `pulumi:"trainingWindow"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/machineLearningJob:MachineLearningJob"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Job
	err := ctx.RegisterResource("grafana:machineLearning/job:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("grafana:machineLearning/job:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// An object representing the custom labels added on the forecast.
	CustomLabels map[string]string `pulumi:"customLabels"`
	// The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
	DatasourceType *string `pulumi:"datasourceType"`
	// The uid of the datasource to query.
	DatasourceUid *string `pulumi:"datasourceUid"`
	// A description of the job.
	Description *string `pulumi:"description"`
	// A list of holiday IDs or names to take into account when training the model.
	Holidays []string `pulumi:"holidays"`
	// The hyperparameters used to fine tune the algorithm. See
	// https://grafana.com/docs/grafana-cloud/alerting-and-irm/machine-learning/forecasts/models/ for the full list of
	// available hyperparameters.
	HyperParams map[string]string `pulumi:"hyperParams"`
	// The data interval in seconds to train the data on.
	Interval *int `pulumi:"interval"`
	// The metric used to query the job results.
	Metric *string `pulumi:"metric"`
	// The name of the job.
	Name *string `pulumi:"name"`
	// An object representing the query params to query Grafana with.
	QueryParams map[string]string `pulumi:"queryParams"`
	// The data interval in seconds to train the data on.
	TrainingWindow *int `pulumi:"trainingWindow"`
}

type JobState struct {
	// An object representing the custom labels added on the forecast.
	CustomLabels pulumi.StringMapInput
	// The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
	DatasourceType pulumi.StringPtrInput
	// The uid of the datasource to query.
	DatasourceUid pulumi.StringPtrInput
	// A description of the job.
	Description pulumi.StringPtrInput
	// A list of holiday IDs or names to take into account when training the model.
	Holidays pulumi.StringArrayInput
	// The hyperparameters used to fine tune the algorithm. See
	// https://grafana.com/docs/grafana-cloud/alerting-and-irm/machine-learning/forecasts/models/ for the full list of
	// available hyperparameters.
	HyperParams pulumi.StringMapInput
	// The data interval in seconds to train the data on.
	Interval pulumi.IntPtrInput
	// The metric used to query the job results.
	Metric pulumi.StringPtrInput
	// The name of the job.
	Name pulumi.StringPtrInput
	// An object representing the query params to query Grafana with.
	QueryParams pulumi.StringMapInput
	// The data interval in seconds to train the data on.
	TrainingWindow pulumi.IntPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// An object representing the custom labels added on the forecast.
	CustomLabels map[string]string `pulumi:"customLabels"`
	// The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
	DatasourceType string `pulumi:"datasourceType"`
	// The uid of the datasource to query.
	DatasourceUid string `pulumi:"datasourceUid"`
	// A description of the job.
	Description *string `pulumi:"description"`
	// A list of holiday IDs or names to take into account when training the model.
	Holidays []string `pulumi:"holidays"`
	// The hyperparameters used to fine tune the algorithm. See
	// https://grafana.com/docs/grafana-cloud/alerting-and-irm/machine-learning/forecasts/models/ for the full list of
	// available hyperparameters.
	HyperParams map[string]string `pulumi:"hyperParams"`
	// The data interval in seconds to train the data on.
	Interval *int `pulumi:"interval"`
	// The metric used to query the job results.
	Metric string `pulumi:"metric"`
	// The name of the job.
	Name *string `pulumi:"name"`
	// An object representing the query params to query Grafana with.
	QueryParams map[string]string `pulumi:"queryParams"`
	// The data interval in seconds to train the data on.
	TrainingWindow *int `pulumi:"trainingWindow"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// An object representing the custom labels added on the forecast.
	CustomLabels pulumi.StringMapInput
	// The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
	DatasourceType pulumi.StringInput
	// The uid of the datasource to query.
	DatasourceUid pulumi.StringInput
	// A description of the job.
	Description pulumi.StringPtrInput
	// A list of holiday IDs or names to take into account when training the model.
	Holidays pulumi.StringArrayInput
	// The hyperparameters used to fine tune the algorithm. See
	// https://grafana.com/docs/grafana-cloud/alerting-and-irm/machine-learning/forecasts/models/ for the full list of
	// available hyperparameters.
	HyperParams pulumi.StringMapInput
	// The data interval in seconds to train the data on.
	Interval pulumi.IntPtrInput
	// The metric used to query the job results.
	Metric pulumi.StringInput
	// The name of the job.
	Name pulumi.StringPtrInput
	// An object representing the query params to query Grafana with.
	QueryParams pulumi.StringMapInput
	// The data interval in seconds to train the data on.
	TrainingWindow pulumi.IntPtrInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

// JobArrayInput is an input type that accepts JobArray and JobArrayOutput values.
// You can construct a concrete instance of `JobArrayInput` via:
//
//	JobArray{ JobArgs{...} }
type JobArrayInput interface {
	pulumi.Input

	ToJobArrayOutput() JobArrayOutput
	ToJobArrayOutputWithContext(context.Context) JobArrayOutput
}

type JobArray []JobInput

func (JobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Job)(nil)).Elem()
}

func (i JobArray) ToJobArrayOutput() JobArrayOutput {
	return i.ToJobArrayOutputWithContext(context.Background())
}

func (i JobArray) ToJobArrayOutputWithContext(ctx context.Context) JobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobArrayOutput)
}

// JobMapInput is an input type that accepts JobMap and JobMapOutput values.
// You can construct a concrete instance of `JobMapInput` via:
//
//	JobMap{ "key": JobArgs{...} }
type JobMapInput interface {
	pulumi.Input

	ToJobMapOutput() JobMapOutput
	ToJobMapOutputWithContext(context.Context) JobMapOutput
}

type JobMap map[string]JobInput

func (JobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Job)(nil)).Elem()
}

func (i JobMap) ToJobMapOutput() JobMapOutput {
	return i.ToJobMapOutputWithContext(context.Background())
}

func (i JobMap) ToJobMapOutputWithContext(ctx context.Context) JobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobMapOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

// An object representing the custom labels added on the forecast.
func (o JobOutput) CustomLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Job) pulumi.StringMapOutput { return v.CustomLabels }).(pulumi.StringMapOutput)
}

// The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
func (o JobOutput) DatasourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.DatasourceType }).(pulumi.StringOutput)
}

// The uid of the datasource to query.
func (o JobOutput) DatasourceUid() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.DatasourceUid }).(pulumi.StringOutput)
}

// A description of the job.
func (o JobOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A list of holiday IDs or names to take into account when training the model.
func (o JobOutput) Holidays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Job) pulumi.StringArrayOutput { return v.Holidays }).(pulumi.StringArrayOutput)
}

// The hyperparameters used to fine tune the algorithm. See
// https://grafana.com/docs/grafana-cloud/alerting-and-irm/machine-learning/forecasts/models/ for the full list of
// available hyperparameters.
func (o JobOutput) HyperParams() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Job) pulumi.StringMapOutput { return v.HyperParams }).(pulumi.StringMapOutput)
}

// The data interval in seconds to train the data on.
func (o JobOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.IntPtrOutput { return v.Interval }).(pulumi.IntPtrOutput)
}

// The metric used to query the job results.
func (o JobOutput) Metric() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Metric }).(pulumi.StringOutput)
}

// The name of the job.
func (o JobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// An object representing the query params to query Grafana with.
func (o JobOutput) QueryParams() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Job) pulumi.StringMapOutput { return v.QueryParams }).(pulumi.StringMapOutput)
}

// The data interval in seconds to train the data on.
func (o JobOutput) TrainingWindow() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.IntPtrOutput { return v.TrainingWindow }).(pulumi.IntPtrOutput)
}

type JobArrayOutput struct{ *pulumi.OutputState }

func (JobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Job)(nil)).Elem()
}

func (o JobArrayOutput) ToJobArrayOutput() JobArrayOutput {
	return o
}

func (o JobArrayOutput) ToJobArrayOutputWithContext(ctx context.Context) JobArrayOutput {
	return o
}

func (o JobArrayOutput) Index(i pulumi.IntInput) JobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Job {
		return vs[0].([]*Job)[vs[1].(int)]
	}).(JobOutput)
}

type JobMapOutput struct{ *pulumi.OutputState }

func (JobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Job)(nil)).Elem()
}

func (o JobMapOutput) ToJobMapOutput() JobMapOutput {
	return o
}

func (o JobMapOutput) ToJobMapOutputWithContext(ctx context.Context) JobMapOutput {
	return o
}

func (o JobMapOutput) MapIndex(k pulumi.StringInput) JobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Job {
		return vs[0].(map[string]*Job)[vs[1].(string)]
	}).(JobOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobInput)(nil)).Elem(), &Job{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobArrayInput)(nil)).Elem(), JobArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobMapInput)(nil)).Elem(), JobMap{})
	pulumi.RegisterOutputType(JobOutput{})
	pulumi.RegisterOutputType(JobArrayOutput{})
	pulumi.RegisterOutputType(JobMapOutput{})
}
