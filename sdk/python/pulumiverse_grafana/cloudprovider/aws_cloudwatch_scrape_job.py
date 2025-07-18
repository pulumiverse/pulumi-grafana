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
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AwsCloudwatchScrapeJobArgs', 'AwsCloudwatchScrapeJob']

@pulumi.input_type
class AwsCloudwatchScrapeJobArgs:
    def __init__(__self__, *,
                 aws_account_resource_id: pulumi.Input[builtins.str],
                 stack_id: pulumi.Input[builtins.str],
                 custom_namespaces: Optional[pulumi.Input[Sequence[pulumi.Input['AwsCloudwatchScrapeJobCustomNamespaceArgs']]]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 export_tags: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 regions_subset_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input['AwsCloudwatchScrapeJobServiceArgs']]]] = None,
                 static_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a AwsCloudwatchScrapeJob resource.
        :param pulumi.Input[builtins.str] aws_account_resource_id: The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resource_id` attribute of the `cloudProvider.AwsAccount` resource.
        :param pulumi.Input[Sequence[pulumi.Input['AwsCloudwatchScrapeJobCustomNamespaceArgs']]] custom_namespaces: Zero or more configuration blocks to configure custom namespaces for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        :param pulumi.Input[builtins.bool] enabled: Whether the AWS CloudWatch Scrape Job is enabled or not. Defaults to `true`.
        :param pulumi.Input[builtins.bool] export_tags: When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] regions_subset_overrides: A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
        :param pulumi.Input[Sequence[pulumi.Input['AwsCloudwatchScrapeJobServiceArgs']]] services: One or more configuration blocks to configure AWS services for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] static_labels: A set of static labels to add to all metrics exported by this scrape job.
        """
        pulumi.set(__self__, "aws_account_resource_id", aws_account_resource_id)
        pulumi.set(__self__, "stack_id", stack_id)
        if custom_namespaces is not None:
            pulumi.set(__self__, "custom_namespaces", custom_namespaces)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if export_tags is not None:
            pulumi.set(__self__, "export_tags", export_tags)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if regions_subset_overrides is not None:
            pulumi.set(__self__, "regions_subset_overrides", regions_subset_overrides)
        if services is not None:
            pulumi.set(__self__, "services", services)
        if static_labels is not None:
            pulumi.set(__self__, "static_labels", static_labels)

    @property
    @pulumi.getter(name="awsAccountResourceId")
    def aws_account_resource_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resource_id` attribute of the `cloudProvider.AwsAccount` resource.
        """
        return pulumi.get(self, "aws_account_resource_id")

    @aws_account_resource_id.setter
    def aws_account_resource_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "aws_account_resource_id", value)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "stack_id", value)

    @property
    @pulumi.getter(name="customNamespaces")
    def custom_namespaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AwsCloudwatchScrapeJobCustomNamespaceArgs']]]]:
        """
        Zero or more configuration blocks to configure custom namespaces for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        """
        return pulumi.get(self, "custom_namespaces")

    @custom_namespaces.setter
    def custom_namespaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AwsCloudwatchScrapeJobCustomNamespaceArgs']]]]):
        pulumi.set(self, "custom_namespaces", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether the AWS CloudWatch Scrape Job is enabled or not. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="exportTags")
    def export_tags(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`. Defaults to `true`.
        """
        return pulumi.get(self, "export_tags")

    @export_tags.setter
    def export_tags(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "export_tags", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="regionsSubsetOverrides")
    def regions_subset_overrides(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
        """
        return pulumi.get(self, "regions_subset_overrides")

    @regions_subset_overrides.setter
    def regions_subset_overrides(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "regions_subset_overrides", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AwsCloudwatchScrapeJobServiceArgs']]]]:
        """
        One or more configuration blocks to configure AWS services for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        """
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AwsCloudwatchScrapeJobServiceArgs']]]]):
        pulumi.set(self, "services", value)

    @property
    @pulumi.getter(name="staticLabels")
    def static_labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        A set of static labels to add to all metrics exported by this scrape job.
        """
        return pulumi.get(self, "static_labels")

    @static_labels.setter
    def static_labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "static_labels", value)


@pulumi.input_type
class _AwsCloudwatchScrapeJobState:
    def __init__(__self__, *,
                 aws_account_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 custom_namespaces: Optional[pulumi.Input[Sequence[pulumi.Input['AwsCloudwatchScrapeJobCustomNamespaceArgs']]]] = None,
                 disabled_reason: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 export_tags: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 regions_subset_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input['AwsCloudwatchScrapeJobServiceArgs']]]] = None,
                 stack_id: Optional[pulumi.Input[builtins.str]] = None,
                 static_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering AwsCloudwatchScrapeJob resources.
        :param pulumi.Input[builtins.str] aws_account_resource_id: The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resource_id` attribute of the `cloudProvider.AwsAccount` resource.
        :param pulumi.Input[Sequence[pulumi.Input['AwsCloudwatchScrapeJobCustomNamespaceArgs']]] custom_namespaces: Zero or more configuration blocks to configure custom namespaces for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        :param pulumi.Input[builtins.str] disabled_reason: When the AWS CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
        :param pulumi.Input[builtins.bool] enabled: Whether the AWS CloudWatch Scrape Job is enabled or not. Defaults to `true`.
        :param pulumi.Input[builtins.bool] export_tags: When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] regions_subset_overrides: A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
        :param pulumi.Input[Sequence[pulumi.Input['AwsCloudwatchScrapeJobServiceArgs']]] services: One or more configuration blocks to configure AWS services for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] static_labels: A set of static labels to add to all metrics exported by this scrape job.
        """
        if aws_account_resource_id is not None:
            pulumi.set(__self__, "aws_account_resource_id", aws_account_resource_id)
        if custom_namespaces is not None:
            pulumi.set(__self__, "custom_namespaces", custom_namespaces)
        if disabled_reason is not None:
            pulumi.set(__self__, "disabled_reason", disabled_reason)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if export_tags is not None:
            pulumi.set(__self__, "export_tags", export_tags)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if regions_subset_overrides is not None:
            pulumi.set(__self__, "regions_subset_overrides", regions_subset_overrides)
        if services is not None:
            pulumi.set(__self__, "services", services)
        if stack_id is not None:
            pulumi.set(__self__, "stack_id", stack_id)
        if static_labels is not None:
            pulumi.set(__self__, "static_labels", static_labels)

    @property
    @pulumi.getter(name="awsAccountResourceId")
    def aws_account_resource_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resource_id` attribute of the `cloudProvider.AwsAccount` resource.
        """
        return pulumi.get(self, "aws_account_resource_id")

    @aws_account_resource_id.setter
    def aws_account_resource_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "aws_account_resource_id", value)

    @property
    @pulumi.getter(name="customNamespaces")
    def custom_namespaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AwsCloudwatchScrapeJobCustomNamespaceArgs']]]]:
        """
        Zero or more configuration blocks to configure custom namespaces for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        """
        return pulumi.get(self, "custom_namespaces")

    @custom_namespaces.setter
    def custom_namespaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AwsCloudwatchScrapeJobCustomNamespaceArgs']]]]):
        pulumi.set(self, "custom_namespaces", value)

    @property
    @pulumi.getter(name="disabledReason")
    def disabled_reason(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        When the AWS CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
        """
        return pulumi.get(self, "disabled_reason")

    @disabled_reason.setter
    def disabled_reason(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "disabled_reason", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether the AWS CloudWatch Scrape Job is enabled or not. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="exportTags")
    def export_tags(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`. Defaults to `true`.
        """
        return pulumi.get(self, "export_tags")

    @export_tags.setter
    def export_tags(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "export_tags", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="regionsSubsetOverrides")
    def regions_subset_overrides(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
        """
        return pulumi.get(self, "regions_subset_overrides")

    @regions_subset_overrides.setter
    def regions_subset_overrides(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "regions_subset_overrides", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AwsCloudwatchScrapeJobServiceArgs']]]]:
        """
        One or more configuration blocks to configure AWS services for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        """
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AwsCloudwatchScrapeJobServiceArgs']]]]):
        pulumi.set(self, "services", value)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "stack_id", value)

    @property
    @pulumi.getter(name="staticLabels")
    def static_labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        A set of static labels to add to all metrics exported by this scrape job.
        """
        return pulumi.get(self, "static_labels")

    @static_labels.setter
    def static_labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "static_labels", value)


@pulumi.type_token("grafana:cloudProvider/awsCloudwatchScrapeJob:AwsCloudwatchScrapeJob")
class AwsCloudwatchScrapeJob(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_account_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 custom_namespaces: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AwsCloudwatchScrapeJobCustomNamespaceArgs', 'AwsCloudwatchScrapeJobCustomNamespaceArgsDict']]]]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 export_tags: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 regions_subset_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AwsCloudwatchScrapeJobServiceArgs', 'AwsCloudwatchScrapeJobServiceArgsDict']]]]] = None,
                 stack_id: Optional[pulumi.Input[builtins.str]] = None,
                 static_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_grafana as grafana
        import pulumiverse_grafana as grafana

        test = grafana.cloud.get_stack(slug="gcloudstacktest")
        test_get_role = aws.iam.get_role(name="my-role")
        test_aws_account = grafana.cloud_provider.AwsAccount("test",
            stack_id=test.id,
            role_arn=test_get_role.arn,
            regions=[
                "us-east-1",
                "us-east-2",
                "us-west-1",
            ])
        test_aws_cloudwatch_scrape_job = grafana.cloud_provider.AwsCloudwatchScrapeJob("test",
            stack_id=test.id,
            name="my-cloudwatch-scrape-job",
            aws_account_resource_id=test_aws_account.resource_id,
            export_tags=True,
            services=[{
                "name": "AWS/EC2",
                "metrics": [
                    {
                        "name": "CPUUtilization",
                        "statistics": ["Average"],
                    },
                    {
                        "name": "StatusCheckFailed",
                        "statistics": ["Maximum"],
                    },
                ],
                "scrape_interval_seconds": 300,
                "resource_discovery_tag_filters": [{
                    "key": "k8s.io/cluster-autoscaler/enabled",
                    "value": "true",
                }],
                "tags_to_add_to_metrics": ["eks:cluster-name"],
            }],
            custom_namespaces=[{
                "name": "CoolApp",
                "metrics": [{
                    "name": "CoolMetric",
                    "statistics": [
                        "Maximum",
                        "Sum",
                    ],
                }],
                "scrape_interval_seconds": 300,
            }],
            static_labels={
                "label1": "value1",
                "label2": "value2",
            })
        ```

        ## Import

        ```sh
        $ pulumi import grafana:cloudProvider/awsCloudwatchScrapeJob:AwsCloudwatchScrapeJob name "{{ stack_id }}:{{ name }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] aws_account_resource_id: The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resource_id` attribute of the `cloudProvider.AwsAccount` resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['AwsCloudwatchScrapeJobCustomNamespaceArgs', 'AwsCloudwatchScrapeJobCustomNamespaceArgsDict']]]] custom_namespaces: Zero or more configuration blocks to configure custom namespaces for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        :param pulumi.Input[builtins.bool] enabled: Whether the AWS CloudWatch Scrape Job is enabled or not. Defaults to `true`.
        :param pulumi.Input[builtins.bool] export_tags: When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] regions_subset_overrides: A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
        :param pulumi.Input[Sequence[pulumi.Input[Union['AwsCloudwatchScrapeJobServiceArgs', 'AwsCloudwatchScrapeJobServiceArgsDict']]]] services: One or more configuration blocks to configure AWS services for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] static_labels: A set of static labels to add to all metrics exported by this scrape job.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AwsCloudwatchScrapeJobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_grafana as grafana
        import pulumiverse_grafana as grafana

        test = grafana.cloud.get_stack(slug="gcloudstacktest")
        test_get_role = aws.iam.get_role(name="my-role")
        test_aws_account = grafana.cloud_provider.AwsAccount("test",
            stack_id=test.id,
            role_arn=test_get_role.arn,
            regions=[
                "us-east-1",
                "us-east-2",
                "us-west-1",
            ])
        test_aws_cloudwatch_scrape_job = grafana.cloud_provider.AwsCloudwatchScrapeJob("test",
            stack_id=test.id,
            name="my-cloudwatch-scrape-job",
            aws_account_resource_id=test_aws_account.resource_id,
            export_tags=True,
            services=[{
                "name": "AWS/EC2",
                "metrics": [
                    {
                        "name": "CPUUtilization",
                        "statistics": ["Average"],
                    },
                    {
                        "name": "StatusCheckFailed",
                        "statistics": ["Maximum"],
                    },
                ],
                "scrape_interval_seconds": 300,
                "resource_discovery_tag_filters": [{
                    "key": "k8s.io/cluster-autoscaler/enabled",
                    "value": "true",
                }],
                "tags_to_add_to_metrics": ["eks:cluster-name"],
            }],
            custom_namespaces=[{
                "name": "CoolApp",
                "metrics": [{
                    "name": "CoolMetric",
                    "statistics": [
                        "Maximum",
                        "Sum",
                    ],
                }],
                "scrape_interval_seconds": 300,
            }],
            static_labels={
                "label1": "value1",
                "label2": "value2",
            })
        ```

        ## Import

        ```sh
        $ pulumi import grafana:cloudProvider/awsCloudwatchScrapeJob:AwsCloudwatchScrapeJob name "{{ stack_id }}:{{ name }}"
        ```

        :param str resource_name: The name of the resource.
        :param AwsCloudwatchScrapeJobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AwsCloudwatchScrapeJobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_account_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 custom_namespaces: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AwsCloudwatchScrapeJobCustomNamespaceArgs', 'AwsCloudwatchScrapeJobCustomNamespaceArgsDict']]]]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 export_tags: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 regions_subset_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AwsCloudwatchScrapeJobServiceArgs', 'AwsCloudwatchScrapeJobServiceArgsDict']]]]] = None,
                 stack_id: Optional[pulumi.Input[builtins.str]] = None,
                 static_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AwsCloudwatchScrapeJobArgs.__new__(AwsCloudwatchScrapeJobArgs)

            if aws_account_resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'aws_account_resource_id'")
            __props__.__dict__["aws_account_resource_id"] = aws_account_resource_id
            __props__.__dict__["custom_namespaces"] = custom_namespaces
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["export_tags"] = export_tags
            __props__.__dict__["name"] = name
            __props__.__dict__["regions_subset_overrides"] = regions_subset_overrides
            __props__.__dict__["services"] = services
            if stack_id is None and not opts.urn:
                raise TypeError("Missing required property 'stack_id'")
            __props__.__dict__["stack_id"] = stack_id
            __props__.__dict__["static_labels"] = static_labels
            __props__.__dict__["disabled_reason"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:cloud/providerAwsCloudwatchScrapeJob:ProviderAwsCloudwatchScrapeJob")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AwsCloudwatchScrapeJob, __self__).__init__(
            'grafana:cloudProvider/awsCloudwatchScrapeJob:AwsCloudwatchScrapeJob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            aws_account_resource_id: Optional[pulumi.Input[builtins.str]] = None,
            custom_namespaces: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AwsCloudwatchScrapeJobCustomNamespaceArgs', 'AwsCloudwatchScrapeJobCustomNamespaceArgsDict']]]]] = None,
            disabled_reason: Optional[pulumi.Input[builtins.str]] = None,
            enabled: Optional[pulumi.Input[builtins.bool]] = None,
            export_tags: Optional[pulumi.Input[builtins.bool]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            regions_subset_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            services: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AwsCloudwatchScrapeJobServiceArgs', 'AwsCloudwatchScrapeJobServiceArgsDict']]]]] = None,
            stack_id: Optional[pulumi.Input[builtins.str]] = None,
            static_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None) -> 'AwsCloudwatchScrapeJob':
        """
        Get an existing AwsCloudwatchScrapeJob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] aws_account_resource_id: The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resource_id` attribute of the `cloudProvider.AwsAccount` resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['AwsCloudwatchScrapeJobCustomNamespaceArgs', 'AwsCloudwatchScrapeJobCustomNamespaceArgsDict']]]] custom_namespaces: Zero or more configuration blocks to configure custom namespaces for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        :param pulumi.Input[builtins.str] disabled_reason: When the AWS CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
        :param pulumi.Input[builtins.bool] enabled: Whether the AWS CloudWatch Scrape Job is enabled or not. Defaults to `true`.
        :param pulumi.Input[builtins.bool] export_tags: When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] regions_subset_overrides: A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
        :param pulumi.Input[Sequence[pulumi.Input[Union['AwsCloudwatchScrapeJobServiceArgs', 'AwsCloudwatchScrapeJobServiceArgsDict']]]] services: One or more configuration blocks to configure AWS services for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] static_labels: A set of static labels to add to all metrics exported by this scrape job.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AwsCloudwatchScrapeJobState.__new__(_AwsCloudwatchScrapeJobState)

        __props__.__dict__["aws_account_resource_id"] = aws_account_resource_id
        __props__.__dict__["custom_namespaces"] = custom_namespaces
        __props__.__dict__["disabled_reason"] = disabled_reason
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["export_tags"] = export_tags
        __props__.__dict__["name"] = name
        __props__.__dict__["regions_subset_overrides"] = regions_subset_overrides
        __props__.__dict__["services"] = services
        __props__.__dict__["stack_id"] = stack_id
        __props__.__dict__["static_labels"] = static_labels
        return AwsCloudwatchScrapeJob(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsAccountResourceId")
    def aws_account_resource_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resource_id` attribute of the `cloudProvider.AwsAccount` resource.
        """
        return pulumi.get(self, "aws_account_resource_id")

    @property
    @pulumi.getter(name="customNamespaces")
    def custom_namespaces(self) -> pulumi.Output[Optional[Sequence['outputs.AwsCloudwatchScrapeJobCustomNamespace']]]:
        """
        Zero or more configuration blocks to configure custom namespaces for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        """
        return pulumi.get(self, "custom_namespaces")

    @property
    @pulumi.getter(name="disabledReason")
    def disabled_reason(self) -> pulumi.Output[builtins.str]:
        """
        When the AWS CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
        """
        return pulumi.get(self, "disabled_reason")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[builtins.bool]:
        """
        Whether the AWS CloudWatch Scrape Job is enabled or not. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="exportTags")
    def export_tags(self) -> pulumi.Output[builtins.bool]:
        """
        When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`. Defaults to `true`.
        """
        return pulumi.get(self, "export_tags")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="regionsSubsetOverrides")
    def regions_subset_overrides(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
        """
        return pulumi.get(self, "regions_subset_overrides")

    @property
    @pulumi.getter
    def services(self) -> pulumi.Output[Optional[Sequence['outputs.AwsCloudwatchScrapeJobService']]]:
        """
        One or more configuration blocks to configure AWS services for the AWS CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        """
        return pulumi.get(self, "services")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter(name="staticLabels")
    def static_labels(self) -> pulumi.Output[Mapping[str, builtins.str]]:
        """
        A set of static labels to add to all metrics exported by this scrape job.
        """
        return pulumi.get(self, "static_labels")

