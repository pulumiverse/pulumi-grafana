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

__all__ = ['AwsAccountArgs', 'AwsAccount']

@pulumi.input_type
class AwsAccountArgs:
    def __init__(__self__, *,
                 regions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 role_arn: pulumi.Input[str],
                 stack_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AwsAccount resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] regions: A set of regions that this AWS Account resource applies to.
        :param pulumi.Input[str] role_arn: An IAM Role ARN string to represent with this AWS Account resource.
        :param pulumi.Input[str] name: An optional human-readable name for this AWS Account resource.
        """
        pulumi.set(__self__, "regions", regions)
        pulumi.set(__self__, "role_arn", role_arn)
        pulumi.set(__self__, "stack_id", stack_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def regions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A set of regions that this AWS Account resource applies to.
        """
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        """
        An IAM Role ARN string to represent with this AWS Account resource.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "stack_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        An optional human-readable name for this AWS Account resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _AwsAccountState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AwsAccount resources.
        :param pulumi.Input[str] name: An optional human-readable name for this AWS Account resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] regions: A set of regions that this AWS Account resource applies to.
        :param pulumi.Input[str] resource_id: The ID given by the Grafana Cloud Provider API to this AWS Account resource.
        :param pulumi.Input[str] role_arn: An IAM Role ARN string to represent with this AWS Account resource.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if regions is not None:
            pulumi.set(__self__, "regions", regions)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if stack_id is not None:
            pulumi.set(__self__, "stack_id", stack_id)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        An optional human-readable name for this AWS Account resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def regions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of regions that this AWS Account resource applies to.
        """
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID given by the Grafana Cloud Provider API to this AWS Account resource.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        An IAM Role ARN string to represent with this AWS Account resource.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_id", value)


class AwsAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
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
        ```

        ## Import

        ```sh
        $ pulumi import grafana:cloudProvider/awsAccount:AwsAccount name "{{ stack_id }}:{{ resource_id }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: An optional human-readable name for this AWS Account resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] regions: A set of regions that this AWS Account resource applies to.
        :param pulumi.Input[str] role_arn: An IAM Role ARN string to represent with this AWS Account resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AwsAccountArgs,
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
        ```

        ## Import

        ```sh
        $ pulumi import grafana:cloudProvider/awsAccount:AwsAccount name "{{ stack_id }}:{{ resource_id }}"
        ```

        :param str resource_name: The name of the resource.
        :param AwsAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AwsAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AwsAccountArgs.__new__(AwsAccountArgs)

            __props__.__dict__["name"] = name
            if regions is None and not opts.urn:
                raise TypeError("Missing required property 'regions'")
            __props__.__dict__["regions"] = regions
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__.__dict__["role_arn"] = role_arn
            if stack_id is None and not opts.urn:
                raise TypeError("Missing required property 'stack_id'")
            __props__.__dict__["stack_id"] = stack_id
            __props__.__dict__["resource_id"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:cloud/providerAwsAccount:ProviderAwsAccount")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AwsAccount, __self__).__init__(
            'grafana:cloudProvider/awsAccount:AwsAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            role_arn: Optional[pulumi.Input[str]] = None,
            stack_id: Optional[pulumi.Input[str]] = None) -> 'AwsAccount':
        """
        Get an existing AwsAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: An optional human-readable name for this AWS Account resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] regions: A set of regions that this AWS Account resource applies to.
        :param pulumi.Input[str] resource_id: The ID given by the Grafana Cloud Provider API to this AWS Account resource.
        :param pulumi.Input[str] role_arn: An IAM Role ARN string to represent with this AWS Account resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AwsAccountState.__new__(_AwsAccountState)

        __props__.__dict__["name"] = name
        __props__.__dict__["regions"] = regions
        __props__.__dict__["resource_id"] = resource_id
        __props__.__dict__["role_arn"] = role_arn
        __props__.__dict__["stack_id"] = stack_id
        return AwsAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        An optional human-readable name for this AWS Account resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def regions(self) -> pulumi.Output[Sequence[str]]:
        """
        A set of regions that this AWS Account resource applies to.
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The ID given by the Grafana Cloud Provider API to this AWS Account resource.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        An IAM Role ARN string to represent with this AWS Account resource.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "stack_id")

