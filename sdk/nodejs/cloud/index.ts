// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AccessPolicyArgs, AccessPolicyState } from "./accessPolicy";
export type AccessPolicy = import("./accessPolicy").AccessPolicy;
export const AccessPolicy: typeof import("./accessPolicy").AccessPolicy = null as any;
utilities.lazyLoad(exports, ["AccessPolicy"], () => require("./accessPolicy"));

export { AccessPolicyTokenArgs, AccessPolicyTokenState } from "./accessPolicyToken";
export type AccessPolicyToken = import("./accessPolicyToken").AccessPolicyToken;
export const AccessPolicyToken: typeof import("./accessPolicyToken").AccessPolicyToken = null as any;
utilities.lazyLoad(exports, ["AccessPolicyToken"], () => require("./accessPolicyToken"));

export { GetIpsResult } from "./getIps";
export const getIps: typeof import("./getIps").getIps = null as any;
export const getIpsOutput: typeof import("./getIps").getIpsOutput = null as any;
utilities.lazyLoad(exports, ["getIps","getIpsOutput"], () => require("./getIps"));

export { GetOrganizationArgs, GetOrganizationResult, GetOrganizationOutputArgs } from "./getOrganization";
export const getOrganization: typeof import("./getOrganization").getOrganization = null as any;
export const getOrganizationOutput: typeof import("./getOrganization").getOrganizationOutput = null as any;
utilities.lazyLoad(exports, ["getOrganization","getOrganizationOutput"], () => require("./getOrganization"));

export { GetStackArgs, GetStackResult, GetStackOutputArgs } from "./getStack";
export const getStack: typeof import("./getStack").getStack = null as any;
export const getStackOutput: typeof import("./getStack").getStackOutput = null as any;
utilities.lazyLoad(exports, ["getStack","getStackOutput"], () => require("./getStack"));

export { OrgMemberArgs, OrgMemberState } from "./orgMember";
export type OrgMember = import("./orgMember").OrgMember;
export const OrgMember: typeof import("./orgMember").OrgMember = null as any;
utilities.lazyLoad(exports, ["OrgMember"], () => require("./orgMember"));

export { PluginInstallationArgs, PluginInstallationState } from "./pluginInstallation";
export type PluginInstallation = import("./pluginInstallation").PluginInstallation;
export const PluginInstallation: typeof import("./pluginInstallation").PluginInstallation = null as any;
utilities.lazyLoad(exports, ["PluginInstallation"], () => require("./pluginInstallation"));

export { StackArgs, StackState } from "./stack";
export type Stack = import("./stack").Stack;
export const Stack: typeof import("./stack").Stack = null as any;
utilities.lazyLoad(exports, ["Stack"], () => require("./stack"));

export { StackServiceAccountArgs, StackServiceAccountState } from "./stackServiceAccount";
export type StackServiceAccount = import("./stackServiceAccount").StackServiceAccount;
export const StackServiceAccount: typeof import("./stackServiceAccount").StackServiceAccount = null as any;
utilities.lazyLoad(exports, ["StackServiceAccount"], () => require("./stackServiceAccount"));

export { StackServiceAccountTokenArgs, StackServiceAccountTokenState } from "./stackServiceAccountToken";
export type StackServiceAccountToken = import("./stackServiceAccountToken").StackServiceAccountToken;
export const StackServiceAccountToken: typeof import("./stackServiceAccountToken").StackServiceAccountToken = null as any;
utilities.lazyLoad(exports, ["StackServiceAccountToken"], () => require("./stackServiceAccountToken"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "grafana:cloud/accessPolicy:AccessPolicy":
                return new AccessPolicy(name, <any>undefined, { urn })
            case "grafana:cloud/accessPolicyToken:AccessPolicyToken":
                return new AccessPolicyToken(name, <any>undefined, { urn })
            case "grafana:cloud/orgMember:OrgMember":
                return new OrgMember(name, <any>undefined, { urn })
            case "grafana:cloud/pluginInstallation:PluginInstallation":
                return new PluginInstallation(name, <any>undefined, { urn })
            case "grafana:cloud/stack:Stack":
                return new Stack(name, <any>undefined, { urn })
            case "grafana:cloud/stackServiceAccount:StackServiceAccount":
                return new StackServiceAccount(name, <any>undefined, { urn })
            case "grafana:cloud/stackServiceAccountToken:StackServiceAccountToken":
                return new StackServiceAccountToken(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("grafana", "cloud/accessPolicy", _module)
pulumi.runtime.registerResourceModule("grafana", "cloud/accessPolicyToken", _module)
pulumi.runtime.registerResourceModule("grafana", "cloud/orgMember", _module)
pulumi.runtime.registerResourceModule("grafana", "cloud/pluginInstallation", _module)
pulumi.runtime.registerResourceModule("grafana", "cloud/stack", _module)
pulumi.runtime.registerResourceModule("grafana", "cloud/stackServiceAccount", _module)
pulumi.runtime.registerResourceModule("grafana", "cloud/stackServiceAccountToken", _module)
