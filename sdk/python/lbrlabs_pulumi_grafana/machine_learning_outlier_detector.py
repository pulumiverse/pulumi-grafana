# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['MachineLearningOutlierDetectorArgs', 'MachineLearningOutlierDetector']

@pulumi.input_type
class MachineLearningOutlierDetectorArgs:
    def __init__(__self__, *,
                 algorithm: pulumi.Input['MachineLearningOutlierDetectorAlgorithmArgs'],
                 datasource_type: pulumi.Input[str],
                 metric: pulumi.Input[str],
                 query_params: pulumi.Input[Mapping[str, Any]],
                 datasource_id: Optional[pulumi.Input[int]] = None,
                 datasource_uid: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MachineLearningOutlierDetector resource.
        :param pulumi.Input['MachineLearningOutlierDetectorAlgorithmArgs'] algorithm: The algorithm to use and its configuration. See https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
        :param pulumi.Input[str] datasource_type: The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        :param pulumi.Input[str] metric: The metric used to query the outlier detector results.
        :param pulumi.Input[Mapping[str, Any]] query_params: An object representing the query params to query Grafana with.
        :param pulumi.Input[int] datasource_id: The id of the datasource to query.
        :param pulumi.Input[str] datasource_uid: The uid of the datasource to query.
        :param pulumi.Input[str] description: A description of the outlier detector.
        :param pulumi.Input[int] interval: The data interval in seconds to monitor. Defaults to `300`.
        :param pulumi.Input[str] name: The name of the algorithm to use ('mad' or 'dbscan').
        """
        pulumi.set(__self__, "algorithm", algorithm)
        pulumi.set(__self__, "datasource_type", datasource_type)
        pulumi.set(__self__, "metric", metric)
        pulumi.set(__self__, "query_params", query_params)
        if datasource_id is not None:
            pulumi.set(__self__, "datasource_id", datasource_id)
        if datasource_uid is not None:
            pulumi.set(__self__, "datasource_uid", datasource_uid)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def algorithm(self) -> pulumi.Input['MachineLearningOutlierDetectorAlgorithmArgs']:
        """
        The algorithm to use and its configuration. See https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: pulumi.Input['MachineLearningOutlierDetectorAlgorithmArgs']):
        pulumi.set(self, "algorithm", value)

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
    @pulumi.getter
    def metric(self) -> pulumi.Input[str]:
        """
        The metric used to query the outlier detector results.
        """
        return pulumi.get(self, "metric")

    @metric.setter
    def metric(self, value: pulumi.Input[str]):
        pulumi.set(self, "metric", value)

    @property
    @pulumi.getter(name="queryParams")
    def query_params(self) -> pulumi.Input[Mapping[str, Any]]:
        """
        An object representing the query params to query Grafana with.
        """
        return pulumi.get(self, "query_params")

    @query_params.setter
    def query_params(self, value: pulumi.Input[Mapping[str, Any]]):
        pulumi.set(self, "query_params", value)

    @property
    @pulumi.getter(name="datasourceId")
    def datasource_id(self) -> Optional[pulumi.Input[int]]:
        """
        The id of the datasource to query.
        """
        return pulumi.get(self, "datasource_id")

    @datasource_id.setter
    def datasource_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "datasource_id", value)

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
        A description of the outlier detector.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        """
        The data interval in seconds to monitor. Defaults to `300`.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the algorithm to use ('mad' or 'dbscan').
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _MachineLearningOutlierDetectorState:
    def __init__(__self__, *,
                 algorithm: Optional[pulumi.Input['MachineLearningOutlierDetectorAlgorithmArgs']] = None,
                 datasource_id: Optional[pulumi.Input[int]] = None,
                 datasource_type: Optional[pulumi.Input[str]] = None,
                 datasource_uid: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 metric: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query_params: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering MachineLearningOutlierDetector resources.
        :param pulumi.Input['MachineLearningOutlierDetectorAlgorithmArgs'] algorithm: The algorithm to use and its configuration. See https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
        :param pulumi.Input[int] datasource_id: The id of the datasource to query.
        :param pulumi.Input[str] datasource_type: The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        :param pulumi.Input[str] datasource_uid: The uid of the datasource to query.
        :param pulumi.Input[str] description: A description of the outlier detector.
        :param pulumi.Input[int] interval: The data interval in seconds to monitor. Defaults to `300`.
        :param pulumi.Input[str] metric: The metric used to query the outlier detector results.
        :param pulumi.Input[str] name: The name of the algorithm to use ('mad' or 'dbscan').
        :param pulumi.Input[Mapping[str, Any]] query_params: An object representing the query params to query Grafana with.
        """
        if algorithm is not None:
            pulumi.set(__self__, "algorithm", algorithm)
        if datasource_id is not None:
            pulumi.set(__self__, "datasource_id", datasource_id)
        if datasource_type is not None:
            pulumi.set(__self__, "datasource_type", datasource_type)
        if datasource_uid is not None:
            pulumi.set(__self__, "datasource_uid", datasource_uid)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if metric is not None:
            pulumi.set(__self__, "metric", metric)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if query_params is not None:
            pulumi.set(__self__, "query_params", query_params)

    @property
    @pulumi.getter
    def algorithm(self) -> Optional[pulumi.Input['MachineLearningOutlierDetectorAlgorithmArgs']]:
        """
        The algorithm to use and its configuration. See https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input['MachineLearningOutlierDetectorAlgorithmArgs']]):
        pulumi.set(self, "algorithm", value)

    @property
    @pulumi.getter(name="datasourceId")
    def datasource_id(self) -> Optional[pulumi.Input[int]]:
        """
        The id of the datasource to query.
        """
        return pulumi.get(self, "datasource_id")

    @datasource_id.setter
    def datasource_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "datasource_id", value)

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
        A description of the outlier detector.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        """
        The data interval in seconds to monitor. Defaults to `300`.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def metric(self) -> Optional[pulumi.Input[str]]:
        """
        The metric used to query the outlier detector results.
        """
        return pulumi.get(self, "metric")

    @metric.setter
    def metric(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metric", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the algorithm to use ('mad' or 'dbscan').
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="queryParams")
    def query_params(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        An object representing the query params to query Grafana with.
        """
        return pulumi.get(self, "query_params")

    @query_params.setter
    def query_params(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "query_params", value)


class MachineLearningOutlierDetector(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input[pulumi.InputType['MachineLearningOutlierDetectorAlgorithmArgs']]] = None,
                 datasource_id: Optional[pulumi.Input[int]] = None,
                 datasource_type: Optional[pulumi.Input[str]] = None,
                 datasource_uid: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 metric: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query_params: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        An outlier detector monitors the results of a query and reports when its values are outside normal bands.

        The normal band is configured by choice of algorithm, its sensitivity and other configuration.

        Visit https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for more details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['MachineLearningOutlierDetectorAlgorithmArgs']] algorithm: The algorithm to use and its configuration. See https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
        :param pulumi.Input[int] datasource_id: The id of the datasource to query.
        :param pulumi.Input[str] datasource_type: The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        :param pulumi.Input[str] datasource_uid: The uid of the datasource to query.
        :param pulumi.Input[str] description: A description of the outlier detector.
        :param pulumi.Input[int] interval: The data interval in seconds to monitor. Defaults to `300`.
        :param pulumi.Input[str] metric: The metric used to query the outlier detector results.
        :param pulumi.Input[str] name: The name of the algorithm to use ('mad' or 'dbscan').
        :param pulumi.Input[Mapping[str, Any]] query_params: An object representing the query params to query Grafana with.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MachineLearningOutlierDetectorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An outlier detector monitors the results of a query and reports when its values are outside normal bands.

        The normal band is configured by choice of algorithm, its sensitivity and other configuration.

        Visit https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for more details.

        :param str resource_name: The name of the resource.
        :param MachineLearningOutlierDetectorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MachineLearningOutlierDetectorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input[pulumi.InputType['MachineLearningOutlierDetectorAlgorithmArgs']]] = None,
                 datasource_id: Optional[pulumi.Input[int]] = None,
                 datasource_type: Optional[pulumi.Input[str]] = None,
                 datasource_uid: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 metric: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query_params: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MachineLearningOutlierDetectorArgs.__new__(MachineLearningOutlierDetectorArgs)

            if algorithm is None and not opts.urn:
                raise TypeError("Missing required property 'algorithm'")
            __props__.__dict__["algorithm"] = algorithm
            __props__.__dict__["datasource_id"] = datasource_id
            if datasource_type is None and not opts.urn:
                raise TypeError("Missing required property 'datasource_type'")
            __props__.__dict__["datasource_type"] = datasource_type
            __props__.__dict__["datasource_uid"] = datasource_uid
            __props__.__dict__["description"] = description
            __props__.__dict__["interval"] = interval
            if metric is None and not opts.urn:
                raise TypeError("Missing required property 'metric'")
            __props__.__dict__["metric"] = metric
            __props__.__dict__["name"] = name
            if query_params is None and not opts.urn:
                raise TypeError("Missing required property 'query_params'")
            __props__.__dict__["query_params"] = query_params
        super(MachineLearningOutlierDetector, __self__).__init__(
            'grafana:index/machineLearningOutlierDetector:MachineLearningOutlierDetector',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            algorithm: Optional[pulumi.Input[pulumi.InputType['MachineLearningOutlierDetectorAlgorithmArgs']]] = None,
            datasource_id: Optional[pulumi.Input[int]] = None,
            datasource_type: Optional[pulumi.Input[str]] = None,
            datasource_uid: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            interval: Optional[pulumi.Input[int]] = None,
            metric: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            query_params: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'MachineLearningOutlierDetector':
        """
        Get an existing MachineLearningOutlierDetector resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['MachineLearningOutlierDetectorAlgorithmArgs']] algorithm: The algorithm to use and its configuration. See https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
        :param pulumi.Input[int] datasource_id: The id of the datasource to query.
        :param pulumi.Input[str] datasource_type: The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        :param pulumi.Input[str] datasource_uid: The uid of the datasource to query.
        :param pulumi.Input[str] description: A description of the outlier detector.
        :param pulumi.Input[int] interval: The data interval in seconds to monitor. Defaults to `300`.
        :param pulumi.Input[str] metric: The metric used to query the outlier detector results.
        :param pulumi.Input[str] name: The name of the algorithm to use ('mad' or 'dbscan').
        :param pulumi.Input[Mapping[str, Any]] query_params: An object representing the query params to query Grafana with.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MachineLearningOutlierDetectorState.__new__(_MachineLearningOutlierDetectorState)

        __props__.__dict__["algorithm"] = algorithm
        __props__.__dict__["datasource_id"] = datasource_id
        __props__.__dict__["datasource_type"] = datasource_type
        __props__.__dict__["datasource_uid"] = datasource_uid
        __props__.__dict__["description"] = description
        __props__.__dict__["interval"] = interval
        __props__.__dict__["metric"] = metric
        __props__.__dict__["name"] = name
        __props__.__dict__["query_params"] = query_params
        return MachineLearningOutlierDetector(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def algorithm(self) -> pulumi.Output['outputs.MachineLearningOutlierDetectorAlgorithm']:
        """
        The algorithm to use and its configuration. See https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter(name="datasourceId")
    def datasource_id(self) -> pulumi.Output[Optional[int]]:
        """
        The id of the datasource to query.
        """
        return pulumi.get(self, "datasource_id")

    @property
    @pulumi.getter(name="datasourceType")
    def datasource_type(self) -> pulumi.Output[str]:
        """
        The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        """
        return pulumi.get(self, "datasource_type")

    @property
    @pulumi.getter(name="datasourceUid")
    def datasource_uid(self) -> pulumi.Output[Optional[str]]:
        """
        The uid of the datasource to query.
        """
        return pulumi.get(self, "datasource_uid")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the outlier detector.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Output[Optional[int]]:
        """
        The data interval in seconds to monitor. Defaults to `300`.
        """
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def metric(self) -> pulumi.Output[str]:
        """
        The metric used to query the outlier detector results.
        """
        return pulumi.get(self, "metric")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the algorithm to use ('mad' or 'dbscan').
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="queryParams")
    def query_params(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        An object representing the query params to query Grafana with.
        """
        return pulumi.get(self, "query_params")

