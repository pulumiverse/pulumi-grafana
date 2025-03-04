// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CollectorArgs, CollectorState } from "./collector";
export type Collector = import("./collector").Collector;
export const Collector: typeof import("./collector").Collector = null as any;
utilities.lazyLoad(exports, ["Collector"], () => require("./collector"));

export { PipelineArgs, PipelineState } from "./pipeline";
export type Pipeline = import("./pipeline").Pipeline;
export const Pipeline: typeof import("./pipeline").Pipeline = null as any;
utilities.lazyLoad(exports, ["Pipeline"], () => require("./pipeline"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "grafana:fleetManagement/collector:Collector":
                return new Collector(name, <any>undefined, { urn })
            case "grafana:fleetManagement/pipeline:Pipeline":
                return new Pipeline(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("grafana", "fleetManagement/collector", _module)
pulumi.runtime.registerResourceModule("grafana", "fleetManagement/pipeline", _module)
