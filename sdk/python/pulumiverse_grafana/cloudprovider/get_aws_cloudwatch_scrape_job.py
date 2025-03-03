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
from . import outputs
from ._inputs import *

__all__ = [
    'GetAwsCloudwatchScrapeJobResult',
    'AwaitableGetAwsCloudwatchScrapeJobResult',
    'get_aws_cloudwatch_scrape_job',
    'get_aws_cloudwatch_scrape_job_output',
]

@pulumi.output_type
class GetAwsCloudwatchScrapeJobResult:
    """
    A collection of values returned by getAwsCloudwatchScrapeJob.
    """
    def __init__(__self__, aws_account_resource_id=None, custom_namespaces=None, disabled_reason=None, enabled=None, export_tags=None, id=None, name=None, regions=None, regions_subset_override_used=None, role_arn=None, services=None, stack_id=None, static_labels=None):
        if aws_account_resource_id and not isinstance(aws_account_resource_id, str):
            raise TypeError("Expected argument 'aws_account_resource_id' to be a str")
        pulumi.set(__self__, "aws_account_resource_id", aws_account_resource_id)
        if custom_namespaces and not isinstance(custom_namespaces, list):
            raise TypeError("Expected argument 'custom_namespaces' to be a list")
        pulumi.set(__self__, "custom_namespaces", custom_namespaces)
        if disabled_reason and not isinstance(disabled_reason, str):
            raise TypeError("Expected argument 'disabled_reason' to be a str")
        pulumi.set(__self__, "disabled_reason", disabled_reason)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if export_tags and not isinstance(export_tags, bool):
            raise TypeError("Expected argument 'export_tags' to be a bool")
        pulumi.set(__self__, "export_tags", export_tags)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if regions and not isinstance(regions, list):
            raise TypeError("Expected argument 'regions' to be a list")
        pulumi.set(__self__, "regions", regions)
        if regions_subset_override_used and not isinstance(regions_subset_override_used, bool):
            raise TypeError("Expected argument 'regions_subset_override_used' to be a bool")
        pulumi.set(__self__, "regions_subset_override_used", regions_subset_override_used)
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        pulumi.set(__self__, "role_arn", role_arn)
        if services and not isinstance(services, list):
            raise TypeError("Expected argument 'services' to be a list")
        pulumi.set(__self__, "services", services)
        if stack_id and not isinstance(stack_id, str):
            raise TypeError("Expected argument 'stack_id' to be a str")
        pulumi.set(__self__, "stack_id", stack_id)
        if static_labels and not isinstance(static_labels, dict):
            raise TypeError("Expected argument 'static_labels' to be a dict")
        pulumi.set(__self__, "static_labels", static_labels)

    @property
    @pulumi.getter(name="awsAccountResourceId")
    def aws_account_resource_id(self) -> str:
        """
        The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this CloudWatch Scrape Job. This can be provided by the `resource_id` attribute of the `cloudProvider.AwsAccount` resource.
        """
        return pulumi.get(self, "aws_account_resource_id")

    @property
    @pulumi.getter(name="customNamespaces")
    def custom_namespaces(self) -> Optional[Sequence['outputs.GetAwsCloudwatchScrapeJobCustomNamespaceResult']]:
        """
        Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        """
        return pulumi.get(self, "custom_namespaces")

    @property
    @pulumi.getter(name="disabledReason")
    def disabled_reason(self) -> str:
        """
        When the CloudWatch Scrape Job is disabled, this will show the reason that it is in that state.
        """
        return pulumi.get(self, "disabled_reason")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Whether the CloudWatch Scrape Job is enabled or not.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="exportTags")
    def export_tags(self) -> bool:
        """
        When enabled, AWS resource tags are exported as Prometheus labels to metrics formatted as `aws_<service_name>_info`.
        """
        return pulumi.get(self, "export_tags")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def regions(self) -> Sequence[str]:
        """
        The set of AWS region names that this CloudWatch Scrape Job is configured to scrape.
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter(name="regionsSubsetOverrideUsed")
    def regions_subset_override_used(self) -> bool:
        """
        When true, the `regions` attribute will be the set of regions configured in the override. When false, the `regions` attribute will be the set of regions belonging to the AWS Account resource that is associated with this CloudWatch Scrape Job.
        """
        return pulumi.get(self, "regions_subset_override_used")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> str:
        """
        The AWS ARN of the IAM role associated with the AWS Account resource that is being used by this CloudWatch Scrape Job.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter
    def services(self) -> Optional[Sequence['outputs.GetAwsCloudwatchScrapeJobServiceResult']]:
        """
        One or more configuration blocks to dictate what this CloudWatch Scrape Job should scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
        """
        return pulumi.get(self, "services")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> str:
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter(name="staticLabels")
    def static_labels(self) -> Mapping[str, str]:
        """
        A set of static labels to add to all metrics exported by this scrape job.
        """
        return pulumi.get(self, "static_labels")


class AwaitableGetAwsCloudwatchScrapeJobResult(GetAwsCloudwatchScrapeJobResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAwsCloudwatchScrapeJobResult(
            aws_account_resource_id=self.aws_account_resource_id,
            custom_namespaces=self.custom_namespaces,
            disabled_reason=self.disabled_reason,
            enabled=self.enabled,
            export_tags=self.export_tags,
            id=self.id,
            name=self.name,
            regions=self.regions,
            regions_subset_override_used=self.regions_subset_override_used,
            role_arn=self.role_arn,
            services=self.services,
            stack_id=self.stack_id,
            static_labels=self.static_labels)


def get_aws_cloudwatch_scrape_job(custom_namespaces: Optional[Sequence[Union['GetAwsCloudwatchScrapeJobCustomNamespaceArgs', 'GetAwsCloudwatchScrapeJobCustomNamespaceArgsDict']]] = None,
                                  name: Optional[str] = None,
                                  services: Optional[Sequence[Union['GetAwsCloudwatchScrapeJobServiceArgs', 'GetAwsCloudwatchScrapeJobServiceArgsDict']]] = None,
                                  stack_id: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAwsCloudwatchScrapeJobResult:
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
    test_get_aws_cloudwatch_scrape_job = test_aws_cloudwatch_scrape_job.name.apply(lambda name: grafana.cloudProvider.get_aws_cloudwatch_scrape_job_output(stack_id=test.id,
        name=name))
    ```


    :param Sequence[Union['GetAwsCloudwatchScrapeJobCustomNamespaceArgs', 'GetAwsCloudwatchScrapeJobCustomNamespaceArgsDict']] custom_namespaces: Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
    :param Sequence[Union['GetAwsCloudwatchScrapeJobServiceArgs', 'GetAwsCloudwatchScrapeJobServiceArgsDict']] services: One or more configuration blocks to dictate what this CloudWatch Scrape Job should scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
    """
    __args__ = dict()
    __args__['customNamespaces'] = custom_namespaces
    __args__['name'] = name
    __args__['services'] = services
    __args__['stackId'] = stack_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:cloudProvider/getAwsCloudwatchScrapeJob:getAwsCloudwatchScrapeJob', __args__, opts=opts, typ=GetAwsCloudwatchScrapeJobResult).value

    return AwaitableGetAwsCloudwatchScrapeJobResult(
        aws_account_resource_id=pulumi.get(__ret__, 'aws_account_resource_id'),
        custom_namespaces=pulumi.get(__ret__, 'custom_namespaces'),
        disabled_reason=pulumi.get(__ret__, 'disabled_reason'),
        enabled=pulumi.get(__ret__, 'enabled'),
        export_tags=pulumi.get(__ret__, 'export_tags'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        regions=pulumi.get(__ret__, 'regions'),
        regions_subset_override_used=pulumi.get(__ret__, 'regions_subset_override_used'),
        role_arn=pulumi.get(__ret__, 'role_arn'),
        services=pulumi.get(__ret__, 'services'),
        stack_id=pulumi.get(__ret__, 'stack_id'),
        static_labels=pulumi.get(__ret__, 'static_labels'))
def get_aws_cloudwatch_scrape_job_output(custom_namespaces: Optional[pulumi.Input[Optional[Sequence[Union['GetAwsCloudwatchScrapeJobCustomNamespaceArgs', 'GetAwsCloudwatchScrapeJobCustomNamespaceArgsDict']]]]] = None,
                                         name: Optional[pulumi.Input[str]] = None,
                                         services: Optional[pulumi.Input[Optional[Sequence[Union['GetAwsCloudwatchScrapeJobServiceArgs', 'GetAwsCloudwatchScrapeJobServiceArgsDict']]]]] = None,
                                         stack_id: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAwsCloudwatchScrapeJobResult]:
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
    test_get_aws_cloudwatch_scrape_job = test_aws_cloudwatch_scrape_job.name.apply(lambda name: grafana.cloudProvider.get_aws_cloudwatch_scrape_job_output(stack_id=test.id,
        name=name))
    ```


    :param Sequence[Union['GetAwsCloudwatchScrapeJobCustomNamespaceArgs', 'GetAwsCloudwatchScrapeJobCustomNamespaceArgsDict']] custom_namespaces: Zero or more configuration blocks to configure custom namespaces for the CloudWatch Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
    :param Sequence[Union['GetAwsCloudwatchScrapeJobServiceArgs', 'GetAwsCloudwatchScrapeJobServiceArgsDict']] services: One or more configuration blocks to dictate what this CloudWatch Scrape Job should scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
    """
    __args__ = dict()
    __args__['customNamespaces'] = custom_namespaces
    __args__['name'] = name
    __args__['services'] = services
    __args__['stackId'] = stack_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('grafana:cloudProvider/getAwsCloudwatchScrapeJob:getAwsCloudwatchScrapeJob', __args__, opts=opts, typ=GetAwsCloudwatchScrapeJobResult)
    return __ret__.apply(lambda __response__: GetAwsCloudwatchScrapeJobResult(
        aws_account_resource_id=pulumi.get(__response__, 'aws_account_resource_id'),
        custom_namespaces=pulumi.get(__response__, 'custom_namespaces'),
        disabled_reason=pulumi.get(__response__, 'disabled_reason'),
        enabled=pulumi.get(__response__, 'enabled'),
        export_tags=pulumi.get(__response__, 'export_tags'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        regions=pulumi.get(__response__, 'regions'),
        regions_subset_override_used=pulumi.get(__response__, 'regions_subset_override_used'),
        role_arn=pulumi.get(__response__, 'role_arn'),
        services=pulumi.get(__response__, 'services'),
        stack_id=pulumi.get(__response__, 'stack_id'),
        static_labels=pulumi.get(__response__, 'static_labels')))
