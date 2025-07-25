// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { DataSourceConfigLbacRulesArgs, DataSourceConfigLbacRulesState } from "./dataSourceConfigLbacRules";
export type DataSourceConfigLbacRules = import("./dataSourceConfigLbacRules").DataSourceConfigLbacRules;
export const DataSourceConfigLbacRules: typeof import("./dataSourceConfigLbacRules").DataSourceConfigLbacRules = null as any;
utilities.lazyLoad(exports, ["DataSourceConfigLbacRules"], () => require("./dataSourceConfigLbacRules"));

export { DataSourcePermissionArgs, DataSourcePermissionState } from "./dataSourcePermission";
export type DataSourcePermission = import("./dataSourcePermission").DataSourcePermission;
export const DataSourcePermission: typeof import("./dataSourcePermission").DataSourcePermission = null as any;
utilities.lazyLoad(exports, ["DataSourcePermission"], () => require("./dataSourcePermission"));

export { DataSourcePermissionItemArgs, DataSourcePermissionItemState } from "./dataSourcePermissionItem";
export type DataSourcePermissionItem = import("./dataSourcePermissionItem").DataSourcePermissionItem;
export const DataSourcePermissionItem: typeof import("./dataSourcePermissionItem").DataSourcePermissionItem = null as any;
utilities.lazyLoad(exports, ["DataSourcePermissionItem"], () => require("./dataSourcePermissionItem"));

export { GetRoleArgs, GetRoleResult, GetRoleOutputArgs } from "./getRole";
export const getRole: typeof import("./getRole").getRole = null as any;
export const getRoleOutput: typeof import("./getRole").getRoleOutput = null as any;
utilities.lazyLoad(exports, ["getRole","getRoleOutput"], () => require("./getRole"));

export { ReportArgs, ReportState } from "./report";
export type Report = import("./report").Report;
export const Report: typeof import("./report").Report = null as any;
utilities.lazyLoad(exports, ["Report"], () => require("./report"));

export { RoleArgs, RoleState } from "./role";
export type Role = import("./role").Role;
export const Role: typeof import("./role").Role = null as any;
utilities.lazyLoad(exports, ["Role"], () => require("./role"));

export { RoleAssignmentArgs, RoleAssignmentState } from "./roleAssignment";
export type RoleAssignment = import("./roleAssignment").RoleAssignment;
export const RoleAssignment: typeof import("./roleAssignment").RoleAssignment = null as any;
utilities.lazyLoad(exports, ["RoleAssignment"], () => require("./roleAssignment"));

export { RoleAssignmentItemArgs, RoleAssignmentItemState } from "./roleAssignmentItem";
export type RoleAssignmentItem = import("./roleAssignmentItem").RoleAssignmentItem;
export const RoleAssignmentItem: typeof import("./roleAssignmentItem").RoleAssignmentItem = null as any;
utilities.lazyLoad(exports, ["RoleAssignmentItem"], () => require("./roleAssignmentItem"));

export { ScimConfigArgs, ScimConfigState } from "./scimConfig";
export type ScimConfig = import("./scimConfig").ScimConfig;
export const ScimConfig: typeof import("./scimConfig").ScimConfig = null as any;
utilities.lazyLoad(exports, ["ScimConfig"], () => require("./scimConfig"));

export { TeamExternalGroupArgs, TeamExternalGroupState } from "./teamExternalGroup";
export type TeamExternalGroup = import("./teamExternalGroup").TeamExternalGroup;
export const TeamExternalGroup: typeof import("./teamExternalGroup").TeamExternalGroup = null as any;
utilities.lazyLoad(exports, ["TeamExternalGroup"], () => require("./teamExternalGroup"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "grafana:enterprise/dataSourceConfigLbacRules:DataSourceConfigLbacRules":
                return new DataSourceConfigLbacRules(name, <any>undefined, { urn })
            case "grafana:enterprise/dataSourcePermission:DataSourcePermission":
                return new DataSourcePermission(name, <any>undefined, { urn })
            case "grafana:enterprise/dataSourcePermissionItem:DataSourcePermissionItem":
                return new DataSourcePermissionItem(name, <any>undefined, { urn })
            case "grafana:enterprise/report:Report":
                return new Report(name, <any>undefined, { urn })
            case "grafana:enterprise/role:Role":
                return new Role(name, <any>undefined, { urn })
            case "grafana:enterprise/roleAssignment:RoleAssignment":
                return new RoleAssignment(name, <any>undefined, { urn })
            case "grafana:enterprise/roleAssignmentItem:RoleAssignmentItem":
                return new RoleAssignmentItem(name, <any>undefined, { urn })
            case "grafana:enterprise/scimConfig:ScimConfig":
                return new ScimConfig(name, <any>undefined, { urn })
            case "grafana:enterprise/teamExternalGroup:TeamExternalGroup":
                return new TeamExternalGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("grafana", "enterprise/dataSourceConfigLbacRules", _module)
pulumi.runtime.registerResourceModule("grafana", "enterprise/dataSourcePermission", _module)
pulumi.runtime.registerResourceModule("grafana", "enterprise/dataSourcePermissionItem", _module)
pulumi.runtime.registerResourceModule("grafana", "enterprise/report", _module)
pulumi.runtime.registerResourceModule("grafana", "enterprise/role", _module)
pulumi.runtime.registerResourceModule("grafana", "enterprise/roleAssignment", _module)
pulumi.runtime.registerResourceModule("grafana", "enterprise/roleAssignmentItem", _module)
pulumi.runtime.registerResourceModule("grafana", "enterprise/scimConfig", _module)
pulumi.runtime.registerResourceModule("grafana", "enterprise/teamExternalGroup", _module)
