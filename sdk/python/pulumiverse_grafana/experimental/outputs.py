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

__all__ = [
    'AppsDashboardMetadata',
    'AppsDashboardOptions',
    'AppsDashboardSpec',
    'AppsPlaylistV0Alpha1Metadata',
    'AppsPlaylistV0Alpha1Options',
    'AppsPlaylistV0Alpha1Spec',
    'AppsPlaylistV0Alpha1SpecItem',
]

@pulumi.output_type
class AppsDashboardMetadata(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "folderUid":
            suggest = "folder_uid"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AppsDashboardMetadata. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AppsDashboardMetadata.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AppsDashboardMetadata.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 uid: builtins.str,
                 folder_uid: Optional[builtins.str] = None,
                 url: Optional[builtins.str] = None,
                 uuid: Optional[builtins.str] = None,
                 version: Optional[builtins.str] = None):
        """
        :param builtins.str uid: The unique identifier of the resource.
        :param builtins.str folder_uid: The UID of the folder to save the resource in.
        :param builtins.str url: The full URL of the resource.
        :param builtins.str uuid: The globally unique identifier of a resource, used by the API for tracking.
        :param builtins.str version: The version of the resource.
        """
        pulumi.set(__self__, "uid", uid)
        if folder_uid is not None:
            pulumi.set(__self__, "folder_uid", folder_uid)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def uid(self) -> builtins.str:
        """
        The unique identifier of the resource.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="folderUid")
    def folder_uid(self) -> Optional[builtins.str]:
        """
        The UID of the folder to save the resource in.
        """
        return pulumi.get(self, "folder_uid")

    @property
    @pulumi.getter
    def url(self) -> Optional[builtins.str]:
        """
        The full URL of the resource.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def uuid(self) -> Optional[builtins.str]:
        """
        The globally unique identifier of a resource, used by the API for tracking.
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def version(self) -> Optional[builtins.str]:
        """
        The version of the resource.
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class AppsDashboardOptions(dict):
    def __init__(__self__, *,
                 overwrite: Optional[builtins.bool] = None):
        """
        :param builtins.bool overwrite: Set to true if you want to overwrite existing resource with newer version, same resource title in folder or same resource uid.
        """
        if overwrite is not None:
            pulumi.set(__self__, "overwrite", overwrite)

    @property
    @pulumi.getter
    def overwrite(self) -> Optional[builtins.bool]:
        """
        Set to true if you want to overwrite existing resource with newer version, same resource title in folder or same resource uid.
        """
        return pulumi.get(self, "overwrite")


@pulumi.output_type
class AppsDashboardSpec(dict):
    def __init__(__self__, *,
                 json: builtins.str,
                 tags: Optional[Sequence[builtins.str]] = None,
                 title: Optional[builtins.str] = None):
        """
        :param builtins.str json: The JSON representation of the dashboard spec.
        :param Sequence[builtins.str] tags: The tags of the dashboard. If not set, the tags will be derived from the JSON spec.
        :param builtins.str title: The title of the dashboard. If not set, the title will be derived from the JSON spec.
        """
        pulumi.set(__self__, "json", json)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if title is not None:
            pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter
    def json(self) -> builtins.str:
        """
        The JSON representation of the dashboard spec.
        """
        return pulumi.get(self, "json")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[builtins.str]]:
        """
        The tags of the dashboard. If not set, the tags will be derived from the JSON spec.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def title(self) -> Optional[builtins.str]:
        """
        The title of the dashboard. If not set, the title will be derived from the JSON spec.
        """
        return pulumi.get(self, "title")


@pulumi.output_type
class AppsPlaylistV0Alpha1Metadata(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "folderUid":
            suggest = "folder_uid"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AppsPlaylistV0Alpha1Metadata. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AppsPlaylistV0Alpha1Metadata.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AppsPlaylistV0Alpha1Metadata.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 uid: builtins.str,
                 folder_uid: Optional[builtins.str] = None,
                 url: Optional[builtins.str] = None,
                 uuid: Optional[builtins.str] = None,
                 version: Optional[builtins.str] = None):
        """
        :param builtins.str uid: The unique identifier of the resource.
        :param builtins.str folder_uid: The UID of the folder to save the resource in.
        :param builtins.str url: The full URL of the resource.
        :param builtins.str uuid: The globally unique identifier of a resource, used by the API for tracking.
        :param builtins.str version: The version of the resource.
        """
        pulumi.set(__self__, "uid", uid)
        if folder_uid is not None:
            pulumi.set(__self__, "folder_uid", folder_uid)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def uid(self) -> builtins.str:
        """
        The unique identifier of the resource.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="folderUid")
    def folder_uid(self) -> Optional[builtins.str]:
        """
        The UID of the folder to save the resource in.
        """
        return pulumi.get(self, "folder_uid")

    @property
    @pulumi.getter
    def url(self) -> Optional[builtins.str]:
        """
        The full URL of the resource.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def uuid(self) -> Optional[builtins.str]:
        """
        The globally unique identifier of a resource, used by the API for tracking.
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def version(self) -> Optional[builtins.str]:
        """
        The version of the resource.
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class AppsPlaylistV0Alpha1Options(dict):
    def __init__(__self__, *,
                 overwrite: Optional[builtins.bool] = None):
        """
        :param builtins.bool overwrite: Set to true if you want to overwrite existing resource with newer version, same resource title in folder or same resource uid.
        """
        if overwrite is not None:
            pulumi.set(__self__, "overwrite", overwrite)

    @property
    @pulumi.getter
    def overwrite(self) -> Optional[builtins.bool]:
        """
        Set to true if you want to overwrite existing resource with newer version, same resource title in folder or same resource uid.
        """
        return pulumi.get(self, "overwrite")


@pulumi.output_type
class AppsPlaylistV0Alpha1Spec(dict):
    def __init__(__self__, *,
                 items: Sequence['outputs.AppsPlaylistV0Alpha1SpecItem'],
                 title: builtins.str,
                 interval: Optional[builtins.str] = None):
        """
        :param Sequence['AppsPlaylistV0Alpha1SpecItemArgs'] items: The items of the playlist.
        :param builtins.str title: The title of the playlist.
        :param builtins.str interval: The interval of the playlist.
        """
        pulumi.set(__self__, "items", items)
        pulumi.set(__self__, "title", title)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)

    @property
    @pulumi.getter
    def items(self) -> Sequence['outputs.AppsPlaylistV0Alpha1SpecItem']:
        """
        The items of the playlist.
        """
        return pulumi.get(self, "items")

    @property
    @pulumi.getter
    def title(self) -> builtins.str:
        """
        The title of the playlist.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def interval(self) -> Optional[builtins.str]:
        """
        The interval of the playlist.
        """
        return pulumi.get(self, "interval")


@pulumi.output_type
class AppsPlaylistV0Alpha1SpecItem(dict):
    def __init__(__self__, *,
                 type: builtins.str,
                 value: builtins.str):
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> builtins.str:
        return pulumi.get(self, "value")


