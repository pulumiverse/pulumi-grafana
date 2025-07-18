# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
from . import _utilities

__all__ = ['DataSourceConfigArgs', 'DataSourceConfig']

@pulumi.input_type
class DataSourceConfigArgs:
    def __init__(__self__, *,
                 http_headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 json_data_encoded: Optional[pulumi.Input[builtins.str]] = None,
                 org_id: Optional[pulumi.Input[builtins.str]] = None,
                 secure_json_data_encoded: Optional[pulumi.Input[builtins.str]] = None,
                 uid: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a DataSourceConfig resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] http_headers: Custom HTTP headers
        :param pulumi.Input[builtins.str] json_data_encoded: Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        :param pulumi.Input[builtins.str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[builtins.str] secure_json_data_encoded: Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        :param pulumi.Input[builtins.str] uid: Unique identifier. If unset, this will be automatically generated.
        """
        if http_headers is not None:
            pulumi.set(__self__, "http_headers", http_headers)
        if json_data_encoded is not None:
            pulumi.set(__self__, "json_data_encoded", json_data_encoded)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if secure_json_data_encoded is not None:
            pulumi.set(__self__, "secure_json_data_encoded", secure_json_data_encoded)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)

    @property
    @pulumi.getter(name="httpHeaders")
    def http_headers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Custom HTTP headers
        """
        return pulumi.get(self, "http_headers")

    @http_headers.setter
    def http_headers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "http_headers", value)

    @property
    @pulumi.getter(name="jsonDataEncoded")
    def json_data_encoded(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        """
        return pulumi.get(self, "json_data_encoded")

    @json_data_encoded.setter
    def json_data_encoded(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "json_data_encoded", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="secureJsonDataEncoded")
    def secure_json_data_encoded(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        """
        return pulumi.get(self, "secure_json_data_encoded")

    @secure_json_data_encoded.setter
    def secure_json_data_encoded(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "secure_json_data_encoded", value)

    @property
    @pulumi.getter
    def uid(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Unique identifier. If unset, this will be automatically generated.
        """
        return pulumi.get(self, "uid")

    @uid.setter
    def uid(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "uid", value)


@pulumi.input_type
class _DataSourceConfigState:
    def __init__(__self__, *,
                 http_headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 json_data_encoded: Optional[pulumi.Input[builtins.str]] = None,
                 org_id: Optional[pulumi.Input[builtins.str]] = None,
                 secure_json_data_encoded: Optional[pulumi.Input[builtins.str]] = None,
                 uid: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering DataSourceConfig resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] http_headers: Custom HTTP headers
        :param pulumi.Input[builtins.str] json_data_encoded: Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        :param pulumi.Input[builtins.str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[builtins.str] secure_json_data_encoded: Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        :param pulumi.Input[builtins.str] uid: Unique identifier. If unset, this will be automatically generated.
        """
        if http_headers is not None:
            pulumi.set(__self__, "http_headers", http_headers)
        if json_data_encoded is not None:
            pulumi.set(__self__, "json_data_encoded", json_data_encoded)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if secure_json_data_encoded is not None:
            pulumi.set(__self__, "secure_json_data_encoded", secure_json_data_encoded)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)

    @property
    @pulumi.getter(name="httpHeaders")
    def http_headers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Custom HTTP headers
        """
        return pulumi.get(self, "http_headers")

    @http_headers.setter
    def http_headers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "http_headers", value)

    @property
    @pulumi.getter(name="jsonDataEncoded")
    def json_data_encoded(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        """
        return pulumi.get(self, "json_data_encoded")

    @json_data_encoded.setter
    def json_data_encoded(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "json_data_encoded", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="secureJsonDataEncoded")
    def secure_json_data_encoded(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        """
        return pulumi.get(self, "secure_json_data_encoded")

    @secure_json_data_encoded.setter
    def secure_json_data_encoded(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "secure_json_data_encoded", value)

    @property
    @pulumi.getter
    def uid(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Unique identifier. If unset, this will be automatically generated.
        """
        return pulumi.get(self, "uid")

    @uid.setter
    def uid(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "uid", value)


warnings.warn("""grafana.index/datasourceconfig.DataSourceConfig has been deprecated in favor of grafana.oss/datasourceconfig.DataSourceConfig""", DeprecationWarning)


@pulumi.type_token("grafana:index/dataSourceConfig:DataSourceConfig")
class DataSourceConfig(pulumi.CustomResource):
    warnings.warn("""grafana.index/datasourceconfig.DataSourceConfig has been deprecated in favor of grafana.oss/datasourceconfig.DataSourceConfig""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 http_headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 json_data_encoded: Optional[pulumi.Input[builtins.str]] = None,
                 org_id: Optional[pulumi.Input[builtins.str]] = None,
                 secure_json_data_encoded: Optional[pulumi.Input[builtins.str]] = None,
                 uid: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        * [Official documentation](https://grafana.com/docs/grafana/latest/datasources/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/data_source/)

        The required arguments for this resource vary depending on the type of data
        source selected (via the 'type' argument).

        Use this resource for configuring multiple datasources, when that configuration (`json_data_encoded` field) requires circular references like in the example below.

        > When using the `oss.DataSourceConfig` resource, the corresponding `oss.DataSource` resources must have the `json_data_encoded` and `http_headers` fields ignored. Otherwise, an infinite update loop will occur. See the example below.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumiverse_grafana as grafana

        loki = grafana.oss.DataSource("loki",
            type="loki",
            name="loki",
            url="http://localhost:3100")
        tempo = grafana.oss.DataSource("tempo",
            type="tempo",
            name="tempo",
            url="http://localhost:3200")
        loki_data_source_config = grafana.oss.DataSourceConfig("loki",
            uid=loki.uid,
            json_data_encoded=pulumi.Output.json_dumps({
                "derivedFields": [{
                    "datasourceUid": tempo.uid,
                    "matcherRegex": "[tT]race_?[iI][dD]\\"?[:=]\\"?(\\\\w+)",
                    "matcherType": "regex",
                    "name": "traceID",
                    "url": "${__value.raw}",
                }],
            }))
        tempo_data_source_config = grafana.oss.DataSourceConfig("tempo",
            uid=tempo.uid,
            json_data_encoded=pulumi.Output.json_dumps({
                "tracesToLogsV2": {
                    "customQuery": True,
                    "datasourceUid": loki.uid,
                    "filterBySpanID": False,
                    "filterByTraceID": False,
                    "query": "|=\\"${__trace.traceId}\\" | json",
                },
            }))
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/dataSourceConfig:DataSourceConfig name "{{ uid }}"
        ```

        ```sh
        $ pulumi import grafana:index/dataSourceConfig:DataSourceConfig name "{{ orgID }}:{{ uid }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] http_headers: Custom HTTP headers
        :param pulumi.Input[builtins.str] json_data_encoded: Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        :param pulumi.Input[builtins.str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[builtins.str] secure_json_data_encoded: Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        :param pulumi.Input[builtins.str] uid: Unique identifier. If unset, this will be automatically generated.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DataSourceConfigArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [Official documentation](https://grafana.com/docs/grafana/latest/datasources/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/data_source/)

        The required arguments for this resource vary depending on the type of data
        source selected (via the 'type' argument).

        Use this resource for configuring multiple datasources, when that configuration (`json_data_encoded` field) requires circular references like in the example below.

        > When using the `oss.DataSourceConfig` resource, the corresponding `oss.DataSource` resources must have the `json_data_encoded` and `http_headers` fields ignored. Otherwise, an infinite update loop will occur. See the example below.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumiverse_grafana as grafana

        loki = grafana.oss.DataSource("loki",
            type="loki",
            name="loki",
            url="http://localhost:3100")
        tempo = grafana.oss.DataSource("tempo",
            type="tempo",
            name="tempo",
            url="http://localhost:3200")
        loki_data_source_config = grafana.oss.DataSourceConfig("loki",
            uid=loki.uid,
            json_data_encoded=pulumi.Output.json_dumps({
                "derivedFields": [{
                    "datasourceUid": tempo.uid,
                    "matcherRegex": "[tT]race_?[iI][dD]\\"?[:=]\\"?(\\\\w+)",
                    "matcherType": "regex",
                    "name": "traceID",
                    "url": "${__value.raw}",
                }],
            }))
        tempo_data_source_config = grafana.oss.DataSourceConfig("tempo",
            uid=tempo.uid,
            json_data_encoded=pulumi.Output.json_dumps({
                "tracesToLogsV2": {
                    "customQuery": True,
                    "datasourceUid": loki.uid,
                    "filterBySpanID": False,
                    "filterByTraceID": False,
                    "query": "|=\\"${__trace.traceId}\\" | json",
                },
            }))
        ```

        ## Import

        ```sh
        $ pulumi import grafana:index/dataSourceConfig:DataSourceConfig name "{{ uid }}"
        ```

        ```sh
        $ pulumi import grafana:index/dataSourceConfig:DataSourceConfig name "{{ orgID }}:{{ uid }}"
        ```

        :param str resource_name: The name of the resource.
        :param DataSourceConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataSourceConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 http_headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 json_data_encoded: Optional[pulumi.Input[builtins.str]] = None,
                 org_id: Optional[pulumi.Input[builtins.str]] = None,
                 secure_json_data_encoded: Optional[pulumi.Input[builtins.str]] = None,
                 uid: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        pulumi.log.warn("""DataSourceConfig is deprecated: grafana.index/datasourceconfig.DataSourceConfig has been deprecated in favor of grafana.oss/datasourceconfig.DataSourceConfig""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DataSourceConfigArgs.__new__(DataSourceConfigArgs)

            __props__.__dict__["http_headers"] = None if http_headers is None else pulumi.Output.secret(http_headers)
            __props__.__dict__["json_data_encoded"] = json_data_encoded
            __props__.__dict__["org_id"] = org_id
            __props__.__dict__["secure_json_data_encoded"] = None if secure_json_data_encoded is None else pulumi.Output.secret(secure_json_data_encoded)
            __props__.__dict__["uid"] = uid
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:index/dataSourceConfig:DataSourceConfig")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["httpHeaders", "secureJsonDataEncoded"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(DataSourceConfig, __self__).__init__(
            'grafana:index/dataSourceConfig:DataSourceConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            http_headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            json_data_encoded: Optional[pulumi.Input[builtins.str]] = None,
            org_id: Optional[pulumi.Input[builtins.str]] = None,
            secure_json_data_encoded: Optional[pulumi.Input[builtins.str]] = None,
            uid: Optional[pulumi.Input[builtins.str]] = None) -> 'DataSourceConfig':
        """
        Get an existing DataSourceConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] http_headers: Custom HTTP headers
        :param pulumi.Input[builtins.str] json_data_encoded: Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        :param pulumi.Input[builtins.str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        :param pulumi.Input[builtins.str] secure_json_data_encoded: Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        :param pulumi.Input[builtins.str] uid: Unique identifier. If unset, this will be automatically generated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DataSourceConfigState.__new__(_DataSourceConfigState)

        __props__.__dict__["http_headers"] = http_headers
        __props__.__dict__["json_data_encoded"] = json_data_encoded
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["secure_json_data_encoded"] = secure_json_data_encoded
        __props__.__dict__["uid"] = uid
        return DataSourceConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="httpHeaders")
    def http_headers(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Custom HTTP headers
        """
        return pulumi.get(self, "http_headers")

    @property
    @pulumi.getter(name="jsonDataEncoded")
    def json_data_encoded(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        """
        return pulumi.get(self, "json_data_encoded")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="secureJsonDataEncoded")
    def secure_json_data_encoded(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        """
        return pulumi.get(self, "secure_json_data_encoded")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[builtins.str]:
        """
        Unique identifier. If unset, this will be automatically generated.
        """
        return pulumi.get(self, "uid")

