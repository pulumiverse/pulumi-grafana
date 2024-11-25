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

__all__ = ['JobArgs', 'Job']

@pulumi.input_type
class JobArgs:
    def __init__(__self__, *,
                 datasource_type: pulumi.Input[str],
                 datasource_uid: pulumi.Input[str],
                 metric: pulumi.Input[str],
                 query_params: pulumi.Input[Mapping[str, pulumi.Input[str]]],
                 custom_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 holidays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 hyper_params: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 training_window: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Job resource.
        :param pulumi.Input[str] datasource_type: The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        :param pulumi.Input[str] datasource_uid: The uid of the datasource to query.
        :param pulumi.Input[str] metric: The metric used to query the job results.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] query_params: An object representing the query params to query Grafana with.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_labels: An object representing the custom labels added on the forecast.
        :param pulumi.Input[str] description: A description of the job.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] holidays: A list of holiday IDs or names to take into account when training the model.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] hyper_params: The hyperparameters used to fine tune the algorithm. See https://grafana.com/docs/grafana-cloud/machine-learning/models/
               for the full list of available hyperparameters.
        :param pulumi.Input[int] interval: The data interval in seconds to train the data on.
        :param pulumi.Input[str] name: The name of the job.
        :param pulumi.Input[int] training_window: The data interval in seconds to train the data on.
        """
        pulumi.set(__self__, "datasource_type", datasource_type)
        pulumi.set(__self__, "datasource_uid", datasource_uid)
        pulumi.set(__self__, "metric", metric)
        pulumi.set(__self__, "query_params", query_params)
        if custom_labels is not None:
            pulumi.set(__self__, "custom_labels", custom_labels)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if holidays is not None:
            pulumi.set(__self__, "holidays", holidays)
        if hyper_params is not None:
            pulumi.set(__self__, "hyper_params", hyper_params)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if training_window is not None:
            pulumi.set(__self__, "training_window", training_window)

    @property
    @pulumi.getter(name="datasourceType")
    def datasource_type(self) -> pulumi.Input[str]:
        """
        The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        """
        return pulumi.get(self, "datasource_type")

    @datasource_type.setter
    def datasource_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "datasource_type", value)

    @property
    @pulumi.getter(name="datasourceUid")
    def datasource_uid(self) -> pulumi.Input[str]:
        """
        The uid of the datasource to query.
        """
        return pulumi.get(self, "datasource_uid")

    @datasource_uid.setter
    def datasource_uid(self, value: pulumi.Input[str]):
        pulumi.set(self, "datasource_uid", value)

    @property
    @pulumi.getter
    def metric(self) -> pulumi.Input[str]:
        """
        The metric used to query the job results.
        """
        return pulumi.get(self, "metric")

    @metric.setter
    def metric(self, value: pulumi.Input[str]):
        pulumi.set(self, "metric", value)

    @property
    @pulumi.getter(name="queryParams")
    def query_params(self) -> pulumi.Input[Mapping[str, pulumi.Input[str]]]:
        """
        An object representing the query params to query Grafana with.
        """
        return pulumi.get(self, "query_params")

    @query_params.setter
    def query_params(self, value: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        pulumi.set(self, "query_params", value)

    @property
    @pulumi.getter(name="customLabels")
    def custom_labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        An object representing the custom labels added on the forecast.
        """
        return pulumi.get(self, "custom_labels")

    @custom_labels.setter
    def custom_labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "custom_labels", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the job.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def holidays(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of holiday IDs or names to take into account when training the model.
        """
        return pulumi.get(self, "holidays")

    @holidays.setter
    def holidays(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "holidays", value)

    @property
    @pulumi.getter(name="hyperParams")
    def hyper_params(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The hyperparameters used to fine tune the algorithm. See https://grafana.com/docs/grafana-cloud/machine-learning/models/
        for the full list of available hyperparameters.
        """
        return pulumi.get(self, "hyper_params")

    @hyper_params.setter
    def hyper_params(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "hyper_params", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        """
        The data interval in seconds to train the data on.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the job.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="trainingWindow")
    def training_window(self) -> Optional[pulumi.Input[int]]:
        """
        The data interval in seconds to train the data on.
        """
        return pulumi.get(self, "training_window")

    @training_window.setter
    def training_window(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "training_window", value)


@pulumi.input_type
class _JobState:
    def __init__(__self__, *,
                 custom_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 datasource_type: Optional[pulumi.Input[str]] = None,
                 datasource_uid: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 holidays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 hyper_params: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 metric: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query_params: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 training_window: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Job resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_labels: An object representing the custom labels added on the forecast.
        :param pulumi.Input[str] datasource_type: The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        :param pulumi.Input[str] datasource_uid: The uid of the datasource to query.
        :param pulumi.Input[str] description: A description of the job.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] holidays: A list of holiday IDs or names to take into account when training the model.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] hyper_params: The hyperparameters used to fine tune the algorithm. See https://grafana.com/docs/grafana-cloud/machine-learning/models/
               for the full list of available hyperparameters.
        :param pulumi.Input[int] interval: The data interval in seconds to train the data on.
        :param pulumi.Input[str] metric: The metric used to query the job results.
        :param pulumi.Input[str] name: The name of the job.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] query_params: An object representing the query params to query Grafana with.
        :param pulumi.Input[int] training_window: The data interval in seconds to train the data on.
        """
        if custom_labels is not None:
            pulumi.set(__self__, "custom_labels", custom_labels)
        if datasource_type is not None:
            pulumi.set(__self__, "datasource_type", datasource_type)
        if datasource_uid is not None:
            pulumi.set(__self__, "datasource_uid", datasource_uid)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if holidays is not None:
            pulumi.set(__self__, "holidays", holidays)
        if hyper_params is not None:
            pulumi.set(__self__, "hyper_params", hyper_params)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if metric is not None:
            pulumi.set(__self__, "metric", metric)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if query_params is not None:
            pulumi.set(__self__, "query_params", query_params)
        if training_window is not None:
            pulumi.set(__self__, "training_window", training_window)

    @property
    @pulumi.getter(name="customLabels")
    def custom_labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        An object representing the custom labels added on the forecast.
        """
        return pulumi.get(self, "custom_labels")

    @custom_labels.setter
    def custom_labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "custom_labels", value)

    @property
    @pulumi.getter(name="datasourceType")
    def datasource_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        """
        return pulumi.get(self, "datasource_type")

    @datasource_type.setter
    def datasource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "datasource_type", value)

    @property
    @pulumi.getter(name="datasourceUid")
    def datasource_uid(self) -> Optional[pulumi.Input[str]]:
        """
        The uid of the datasource to query.
        """
        return pulumi.get(self, "datasource_uid")

    @datasource_uid.setter
    def datasource_uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "datasource_uid", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the job.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def holidays(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of holiday IDs or names to take into account when training the model.
        """
        return pulumi.get(self, "holidays")

    @holidays.setter
    def holidays(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "holidays", value)

    @property
    @pulumi.getter(name="hyperParams")
    def hyper_params(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The hyperparameters used to fine tune the algorithm. See https://grafana.com/docs/grafana-cloud/machine-learning/models/
        for the full list of available hyperparameters.
        """
        return pulumi.get(self, "hyper_params")

    @hyper_params.setter
    def hyper_params(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "hyper_params", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        """
        The data interval in seconds to train the data on.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def metric(self) -> Optional[pulumi.Input[str]]:
        """
        The metric used to query the job results.
        """
        return pulumi.get(self, "metric")

    @metric.setter
    def metric(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metric", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the job.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="queryParams")
    def query_params(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        An object representing the query params to query Grafana with.
        """
        return pulumi.get(self, "query_params")

    @query_params.setter
    def query_params(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "query_params", value)

    @property
    @pulumi.getter(name="trainingWindow")
    def training_window(self) -> Optional[pulumi.Input[int]]:
        """
        The data interval in seconds to train the data on.
        """
        return pulumi.get(self, "training_window")

    @training_window.setter
    def training_window(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "training_window", value)


class Job(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 datasource_type: Optional[pulumi.Input[str]] = None,
                 datasource_uid: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 holidays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 hyper_params: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 metric: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query_params: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 training_window: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        A job defines the queries and model parameters for a machine learning task.

        ## Example Usage

        ### Basic Forecast

        This forecast uses a Prometheus datasource, where the source query is defined in the `expr` field of the `query_params` attribute.

        Other datasources are supported, but the structure `query_params` may differ.

        ```python
        import pulumi
        import json
        import pulumiverse_grafana as grafana

        foo = grafana.oss.DataSource("foo",
            type="prometheus",
            name="prometheus-ds-test",
            uid="prometheus-ds-test-uid",
            url="https://my-instance.com",
            basic_auth_enabled=True,
            basic_auth_username="username",
            json_data_encoded=json.dumps({
                "httpMethod": "POST",
                "prometheusType": "Mimir",
                "prometheusVersion": "2.4.0",
            }),
            secure_json_data_encoded=json.dumps({
                "basicAuthPassword": "password",
            }))
        test_job = grafana.machine_learning.Job("test_job",
            name="Test Job",
            metric="tf_test_job",
            datasource_type="prometheus",
            datasource_uid=foo.uid,
            query_params={
                "expr": "grafanacloud_grafana_instance_active_user_count",
            })
        ```

        ### Tuned Forecast

        This forecast has tuned hyperparameters to improve the accuracy of the model.

        ```python
        import pulumi
        import json
        import pulumiverse_grafana as grafana

        foo = grafana.oss.DataSource("foo",
            type="prometheus",
            name="prometheus-ds-test",
            uid="prometheus-ds-test-uid",
            url="https://my-instance.com",
            basic_auth_enabled=True,
            basic_auth_username="username",
            json_data_encoded=json.dumps({
                "httpMethod": "POST",
                "prometheusType": "Mimir",
                "prometheusVersion": "2.4.0",
            }),
            secure_json_data_encoded=json.dumps({
                "basicAuthPassword": "password",
            }))
        test_job = grafana.machine_learning.Job("test_job",
            name="Test Job",
            metric="tf_test_job",
            datasource_type="prometheus",
            datasource_uid=foo.uid,
            query_params={
                "expr": "grafanacloud_grafana_instance_active_user_count",
            },
            hyper_params={
                "daily_seasonality": "15",
                "weekly_seasonality": "10",
            },
            custom_labels={
                "example_label": "example_value",
            })
        ```

        ### Forecast with Holidays

        This forecast has holidays which will be taken into account when training the model.

        ```python
        import pulumi
        import json
        import pulumiverse_grafana as grafana

        foo = grafana.oss.DataSource("foo",
            type="prometheus",
            name="prometheus-ds-test",
            uid="prometheus-ds-test-uid",
            url="https://my-instance.com",
            basic_auth_enabled=True,
            basic_auth_username="username",
            json_data_encoded=json.dumps({
                "httpMethod": "POST",
                "prometheusType": "Mimir",
                "prometheusVersion": "2.4.0",
            }),
            secure_json_data_encoded=json.dumps({
                "basicAuthPassword": "password",
            }))
        test_holiday = grafana.machine_learning.Holiday("test_holiday",
            name="Test Holiday",
            custom_periods=[{
                "name": "First of January",
                "start_time": "2023-01-01T00:00:00Z",
                "end_time": "2023-01-02T00:00:00Z",
            }])
        test_job = grafana.machine_learning.Job("test_job",
            name="Test Job",
            metric="tf_test_job",
            datasource_type="prometheus",
            datasource_uid=foo.uid,
            query_params={
                "expr": "grafanacloud_grafana_instance_active_user_count",
            },
            holidays=[test_holiday.id])
        ```

        ## Import

        ```sh
        $ pulumi import grafana:machineLearning/job:Job name "{{ id }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_labels: An object representing the custom labels added on the forecast.
        :param pulumi.Input[str] datasource_type: The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        :param pulumi.Input[str] datasource_uid: The uid of the datasource to query.
        :param pulumi.Input[str] description: A description of the job.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] holidays: A list of holiday IDs or names to take into account when training the model.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] hyper_params: The hyperparameters used to fine tune the algorithm. See https://grafana.com/docs/grafana-cloud/machine-learning/models/
               for the full list of available hyperparameters.
        :param pulumi.Input[int] interval: The data interval in seconds to train the data on.
        :param pulumi.Input[str] metric: The metric used to query the job results.
        :param pulumi.Input[str] name: The name of the job.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] query_params: An object representing the query params to query Grafana with.
        :param pulumi.Input[int] training_window: The data interval in seconds to train the data on.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: JobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A job defines the queries and model parameters for a machine learning task.

        ## Example Usage

        ### Basic Forecast

        This forecast uses a Prometheus datasource, where the source query is defined in the `expr` field of the `query_params` attribute.

        Other datasources are supported, but the structure `query_params` may differ.

        ```python
        import pulumi
        import json
        import pulumiverse_grafana as grafana

        foo = grafana.oss.DataSource("foo",
            type="prometheus",
            name="prometheus-ds-test",
            uid="prometheus-ds-test-uid",
            url="https://my-instance.com",
            basic_auth_enabled=True,
            basic_auth_username="username",
            json_data_encoded=json.dumps({
                "httpMethod": "POST",
                "prometheusType": "Mimir",
                "prometheusVersion": "2.4.0",
            }),
            secure_json_data_encoded=json.dumps({
                "basicAuthPassword": "password",
            }))
        test_job = grafana.machine_learning.Job("test_job",
            name="Test Job",
            metric="tf_test_job",
            datasource_type="prometheus",
            datasource_uid=foo.uid,
            query_params={
                "expr": "grafanacloud_grafana_instance_active_user_count",
            })
        ```

        ### Tuned Forecast

        This forecast has tuned hyperparameters to improve the accuracy of the model.

        ```python
        import pulumi
        import json
        import pulumiverse_grafana as grafana

        foo = grafana.oss.DataSource("foo",
            type="prometheus",
            name="prometheus-ds-test",
            uid="prometheus-ds-test-uid",
            url="https://my-instance.com",
            basic_auth_enabled=True,
            basic_auth_username="username",
            json_data_encoded=json.dumps({
                "httpMethod": "POST",
                "prometheusType": "Mimir",
                "prometheusVersion": "2.4.0",
            }),
            secure_json_data_encoded=json.dumps({
                "basicAuthPassword": "password",
            }))
        test_job = grafana.machine_learning.Job("test_job",
            name="Test Job",
            metric="tf_test_job",
            datasource_type="prometheus",
            datasource_uid=foo.uid,
            query_params={
                "expr": "grafanacloud_grafana_instance_active_user_count",
            },
            hyper_params={
                "daily_seasonality": "15",
                "weekly_seasonality": "10",
            },
            custom_labels={
                "example_label": "example_value",
            })
        ```

        ### Forecast with Holidays

        This forecast has holidays which will be taken into account when training the model.

        ```python
        import pulumi
        import json
        import pulumiverse_grafana as grafana

        foo = grafana.oss.DataSource("foo",
            type="prometheus",
            name="prometheus-ds-test",
            uid="prometheus-ds-test-uid",
            url="https://my-instance.com",
            basic_auth_enabled=True,
            basic_auth_username="username",
            json_data_encoded=json.dumps({
                "httpMethod": "POST",
                "prometheusType": "Mimir",
                "prometheusVersion": "2.4.0",
            }),
            secure_json_data_encoded=json.dumps({
                "basicAuthPassword": "password",
            }))
        test_holiday = grafana.machine_learning.Holiday("test_holiday",
            name="Test Holiday",
            custom_periods=[{
                "name": "First of January",
                "start_time": "2023-01-01T00:00:00Z",
                "end_time": "2023-01-02T00:00:00Z",
            }])
        test_job = grafana.machine_learning.Job("test_job",
            name="Test Job",
            metric="tf_test_job",
            datasource_type="prometheus",
            datasource_uid=foo.uid,
            query_params={
                "expr": "grafanacloud_grafana_instance_active_user_count",
            },
            holidays=[test_holiday.id])
        ```

        ## Import

        ```sh
        $ pulumi import grafana:machineLearning/job:Job name "{{ id }}"
        ```

        :param str resource_name: The name of the resource.
        :param JobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(JobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 datasource_type: Optional[pulumi.Input[str]] = None,
                 datasource_uid: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 holidays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 hyper_params: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 metric: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query_params: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 training_window: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = JobArgs.__new__(JobArgs)

            __props__.__dict__["custom_labels"] = custom_labels
            if datasource_type is None and not opts.urn:
                raise TypeError("Missing required property 'datasource_type'")
            __props__.__dict__["datasource_type"] = datasource_type
            if datasource_uid is None and not opts.urn:
                raise TypeError("Missing required property 'datasource_uid'")
            __props__.__dict__["datasource_uid"] = datasource_uid
            __props__.__dict__["description"] = description
            __props__.__dict__["holidays"] = holidays
            __props__.__dict__["hyper_params"] = hyper_params
            __props__.__dict__["interval"] = interval
            if metric is None and not opts.urn:
                raise TypeError("Missing required property 'metric'")
            __props__.__dict__["metric"] = metric
            __props__.__dict__["name"] = name
            if query_params is None and not opts.urn:
                raise TypeError("Missing required property 'query_params'")
            __props__.__dict__["query_params"] = query_params
            __props__.__dict__["training_window"] = training_window
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:index/machineLearningJob:MachineLearningJob")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Job, __self__).__init__(
            'grafana:machineLearning/job:Job',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            custom_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            datasource_type: Optional[pulumi.Input[str]] = None,
            datasource_uid: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            holidays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            hyper_params: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            interval: Optional[pulumi.Input[int]] = None,
            metric: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            query_params: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            training_window: Optional[pulumi.Input[int]] = None) -> 'Job':
        """
        Get an existing Job resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_labels: An object representing the custom labels added on the forecast.
        :param pulumi.Input[str] datasource_type: The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        :param pulumi.Input[str] datasource_uid: The uid of the datasource to query.
        :param pulumi.Input[str] description: A description of the job.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] holidays: A list of holiday IDs or names to take into account when training the model.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] hyper_params: The hyperparameters used to fine tune the algorithm. See https://grafana.com/docs/grafana-cloud/machine-learning/models/
               for the full list of available hyperparameters.
        :param pulumi.Input[int] interval: The data interval in seconds to train the data on.
        :param pulumi.Input[str] metric: The metric used to query the job results.
        :param pulumi.Input[str] name: The name of the job.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] query_params: An object representing the query params to query Grafana with.
        :param pulumi.Input[int] training_window: The data interval in seconds to train the data on.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _JobState.__new__(_JobState)

        __props__.__dict__["custom_labels"] = custom_labels
        __props__.__dict__["datasource_type"] = datasource_type
        __props__.__dict__["datasource_uid"] = datasource_uid
        __props__.__dict__["description"] = description
        __props__.__dict__["holidays"] = holidays
        __props__.__dict__["hyper_params"] = hyper_params
        __props__.__dict__["interval"] = interval
        __props__.__dict__["metric"] = metric
        __props__.__dict__["name"] = name
        __props__.__dict__["query_params"] = query_params
        __props__.__dict__["training_window"] = training_window
        return Job(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="customLabels")
    def custom_labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        An object representing the custom labels added on the forecast.
        """
        return pulumi.get(self, "custom_labels")

    @property
    @pulumi.getter(name="datasourceType")
    def datasource_type(self) -> pulumi.Output[str]:
        """
        The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        """
        return pulumi.get(self, "datasource_type")

    @property
    @pulumi.getter(name="datasourceUid")
    def datasource_uid(self) -> pulumi.Output[str]:
        """
        The uid of the datasource to query.
        """
        return pulumi.get(self, "datasource_uid")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the job.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def holidays(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of holiday IDs or names to take into account when training the model.
        """
        return pulumi.get(self, "holidays")

    @property
    @pulumi.getter(name="hyperParams")
    def hyper_params(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The hyperparameters used to fine tune the algorithm. See https://grafana.com/docs/grafana-cloud/machine-learning/models/
        for the full list of available hyperparameters.
        """
        return pulumi.get(self, "hyper_params")

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Output[Optional[int]]:
        """
        The data interval in seconds to train the data on.
        """
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def metric(self) -> pulumi.Output[str]:
        """
        The metric used to query the job results.
        """
        return pulumi.get(self, "metric")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the job.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="queryParams")
    def query_params(self) -> pulumi.Output[Mapping[str, str]]:
        """
        An object representing the query params to query Grafana with.
        """
        return pulumi.get(self, "query_params")

    @property
    @pulumi.getter(name="trainingWindow")
    def training_window(self) -> pulumi.Output[Optional[int]]:
        """
        The data interval in seconds to train the data on.
        """
        return pulumi.get(self, "training_window")

