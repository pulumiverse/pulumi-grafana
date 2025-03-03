// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetMetricsEndpointScrapeJobArgs, GetMetricsEndpointScrapeJobResult, GetMetricsEndpointScrapeJobOutputArgs } from "./getMetricsEndpointScrapeJob";
export const getMetricsEndpointScrapeJob: typeof import("./getMetricsEndpointScrapeJob").getMetricsEndpointScrapeJob = null as any;
export const getMetricsEndpointScrapeJobOutput: typeof import("./getMetricsEndpointScrapeJob").getMetricsEndpointScrapeJobOutput = null as any;
utilities.lazyLoad(exports, ["getMetricsEndpointScrapeJob","getMetricsEndpointScrapeJobOutput"], () => require("./getMetricsEndpointScrapeJob"));

export { MetricsEndpointScrapeJobArgs, MetricsEndpointScrapeJobState } from "./metricsEndpointScrapeJob";
export type MetricsEndpointScrapeJob = import("./metricsEndpointScrapeJob").MetricsEndpointScrapeJob;
export const MetricsEndpointScrapeJob: typeof import("./metricsEndpointScrapeJob").MetricsEndpointScrapeJob = null as any;
utilities.lazyLoad(exports, ["MetricsEndpointScrapeJob"], () => require("./metricsEndpointScrapeJob"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "grafana:connections/metricsEndpointScrapeJob:MetricsEndpointScrapeJob":
                return new MetricsEndpointScrapeJob(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("grafana", "connections/metricsEndpointScrapeJob", _module)
