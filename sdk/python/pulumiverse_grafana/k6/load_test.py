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

__all__ = ['LoadTestArgs', 'LoadTest']

@pulumi.input_type
class LoadTestArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[builtins.str],
                 script: pulumi.Input[builtins.str],
                 baseline_test_run_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a LoadTest resource.
        :param pulumi.Input[builtins.str] project_id: The identifier of the project this load test belongs to.
        :param pulumi.Input[builtins.str] script: The k6 test script content. Can be provided inline or via the `file()` function.
        :param pulumi.Input[builtins.str] baseline_test_run_id: Identifier of a baseline test run used for results comparison.
        :param pulumi.Input[builtins.str] name: Human-friendly identifier of the load test.
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "script", script)
        if baseline_test_run_id is not None:
            pulumi.set(__self__, "baseline_test_run_id", baseline_test_run_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[builtins.str]:
        """
        The identifier of the project this load test belongs to.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def script(self) -> pulumi.Input[builtins.str]:
        """
        The k6 test script content. Can be provided inline or via the `file()` function.
        """
        return pulumi.get(self, "script")

    @script.setter
    def script(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "script", value)

    @property
    @pulumi.getter(name="baselineTestRunId")
    def baseline_test_run_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Identifier of a baseline test run used for results comparison.
        """
        return pulumi.get(self, "baseline_test_run_id")

    @baseline_test_run_id.setter
    def baseline_test_run_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "baseline_test_run_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Human-friendly identifier of the load test.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _LoadTestState:
    def __init__(__self__, *,
                 baseline_test_run_id: Optional[pulumi.Input[builtins.str]] = None,
                 created: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 script: Optional[pulumi.Input[builtins.str]] = None,
                 updated: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering LoadTest resources.
        :param pulumi.Input[builtins.str] baseline_test_run_id: Identifier of a baseline test run used for results comparison.
        :param pulumi.Input[builtins.str] created: The date when the load test was created.
        :param pulumi.Input[builtins.str] name: Human-friendly identifier of the load test.
        :param pulumi.Input[builtins.str] project_id: The identifier of the project this load test belongs to.
        :param pulumi.Input[builtins.str] script: The k6 test script content. Can be provided inline or via the `file()` function.
        :param pulumi.Input[builtins.str] updated: The date when the load test was last updated.
        """
        if baseline_test_run_id is not None:
            pulumi.set(__self__, "baseline_test_run_id", baseline_test_run_id)
        if created is not None:
            pulumi.set(__self__, "created", created)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if script is not None:
            pulumi.set(__self__, "script", script)
        if updated is not None:
            pulumi.set(__self__, "updated", updated)

    @property
    @pulumi.getter(name="baselineTestRunId")
    def baseline_test_run_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Identifier of a baseline test run used for results comparison.
        """
        return pulumi.get(self, "baseline_test_run_id")

    @baseline_test_run_id.setter
    def baseline_test_run_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "baseline_test_run_id", value)

    @property
    @pulumi.getter
    def created(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The date when the load test was created.
        """
        return pulumi.get(self, "created")

    @created.setter
    def created(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Human-friendly identifier of the load test.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The identifier of the project this load test belongs to.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def script(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The k6 test script content. Can be provided inline or via the `file()` function.
        """
        return pulumi.get(self, "script")

    @script.setter
    def script(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "script", value)

    @property
    @pulumi.getter
    def updated(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The date when the load test was last updated.
        """
        return pulumi.get(self, "updated")

    @updated.setter
    def updated(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "updated", value)


@pulumi.type_token("grafana:k6/loadTest:LoadTest")
class LoadTest(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 baseline_test_run_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 script: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Manages a k6 load test.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        load_test_project = grafana.k6.Project("load_test_project", name="Terraform Load Test Project")
        test_load_test = grafana.k6.LoadTest("test_load_test",
            project_id=load_test_project.id,
            name="Terraform Test Load Test",
            script=\"\"\"export default function() {
          console.log('Hello from k6!');
        }
        \"\"\")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:k6/loadTest:LoadTest name "{{ id }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] baseline_test_run_id: Identifier of a baseline test run used for results comparison.
        :param pulumi.Input[builtins.str] name: Human-friendly identifier of the load test.
        :param pulumi.Input[builtins.str] project_id: The identifier of the project this load test belongs to.
        :param pulumi.Input[builtins.str] script: The k6 test script content. Can be provided inline or via the `file()` function.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LoadTestArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a k6 load test.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        load_test_project = grafana.k6.Project("load_test_project", name="Terraform Load Test Project")
        test_load_test = grafana.k6.LoadTest("test_load_test",
            project_id=load_test_project.id,
            name="Terraform Test Load Test",
            script=\"\"\"export default function() {
          console.log('Hello from k6!');
        }
        \"\"\")
        ```

        ## Import

        ```sh
        $ pulumi import grafana:k6/loadTest:LoadTest name "{{ id }}"
        ```

        :param str resource_name: The name of the resource.
        :param LoadTestArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LoadTestArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 baseline_test_run_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 script: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LoadTestArgs.__new__(LoadTestArgs)

            __props__.__dict__["baseline_test_run_id"] = baseline_test_run_id
            __props__.__dict__["name"] = name
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if script is None and not opts.urn:
                raise TypeError("Missing required property 'script'")
            __props__.__dict__["script"] = script
            __props__.__dict__["created"] = None
            __props__.__dict__["updated"] = None
        super(LoadTest, __self__).__init__(
            'grafana:k6/loadTest:LoadTest',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            baseline_test_run_id: Optional[pulumi.Input[builtins.str]] = None,
            created: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None,
            script: Optional[pulumi.Input[builtins.str]] = None,
            updated: Optional[pulumi.Input[builtins.str]] = None) -> 'LoadTest':
        """
        Get an existing LoadTest resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] baseline_test_run_id: Identifier of a baseline test run used for results comparison.
        :param pulumi.Input[builtins.str] created: The date when the load test was created.
        :param pulumi.Input[builtins.str] name: Human-friendly identifier of the load test.
        :param pulumi.Input[builtins.str] project_id: The identifier of the project this load test belongs to.
        :param pulumi.Input[builtins.str] script: The k6 test script content. Can be provided inline or via the `file()` function.
        :param pulumi.Input[builtins.str] updated: The date when the load test was last updated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LoadTestState.__new__(_LoadTestState)

        __props__.__dict__["baseline_test_run_id"] = baseline_test_run_id
        __props__.__dict__["created"] = created
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["script"] = script
        __props__.__dict__["updated"] = updated
        return LoadTest(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="baselineTestRunId")
    def baseline_test_run_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Identifier of a baseline test run used for results comparison.
        """
        return pulumi.get(self, "baseline_test_run_id")

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[builtins.str]:
        """
        The date when the load test was created.
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Human-friendly identifier of the load test.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        """
        The identifier of the project this load test belongs to.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def script(self) -> pulumi.Output[builtins.str]:
        """
        The k6 test script content. Can be provided inline or via the `file()` function.
        """
        return pulumi.get(self, "script")

    @property
    @pulumi.getter
    def updated(self) -> pulumi.Output[builtins.str]:
        """
        The date when the load test was last updated.
        """
        return pulumi.get(self, "updated")

