// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * **Note:** This resource is available only with Grafana Enterprise 7.+.
 *
 * * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/create-reports/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/reporting/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = new grafana.oss.Dashboard("test", {
 *     configJson: `{
 *   "uid": "report-dashboard",
 *   "title": "report-dashboard"
 * }
 * `,
 *     message: "inital commit.",
 * });
 * const testReport = new grafana.enterprise.Report("test", {
 *     name: "my report",
 *     recipients: ["some@email.com"],
 *     dashboards: [{
 *         uid: test.uid,
 *     }],
 *     schedule: {
 *         frequency: "hourly",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:index/report:Report name "{{ id }}"
 * ```
 *
 * ```sh
 * $ pulumi import grafana:index/report:Report name "{{ orgID }}:{{ id }}"
 * ```
 *
 * @deprecated grafana.index/report.Report has been deprecated in favor of grafana.enterprise/report.Report
 */
export class Report extends pulumi.CustomResource {
    /**
     * Get an existing Report resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReportState, opts?: pulumi.CustomResourceOptions): Report {
        pulumi.log.warn("Report is deprecated: grafana.index/report.Report has been deprecated in favor of grafana.enterprise/report.Report")
        return new Report(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/report:Report';

    /**
     * Returns true if the given object is an instance of Report.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Report {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Report.__pulumiType;
    }

    /**
     * List of dashboards to render into the report
     */
    public readonly dashboards!: pulumi.Output<outputs.ReportDashboard[] | undefined>;
    /**
     * Specifies what kind of attachment to generate for the report. Allowed values: `pdf`, `csv`, `image`.
     */
    public readonly formats!: pulumi.Output<string[] | undefined>;
    /**
     * Whether to include a link to the dashboard in the report. Defaults to `true`.
     */
    public readonly includeDashboardLink!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to include a CSV file of table panel data. Defaults to `false`.
     */
    public readonly includeTableCsv!: pulumi.Output<boolean | undefined>;
    /**
     * Layout of the report. Allowed values: `simple`, `grid`. Defaults to `grid`.
     */
    public readonly layout!: pulumi.Output<string | undefined>;
    /**
     * Message to be sent in the report.
     */
    public readonly message!: pulumi.Output<string | undefined>;
    /**
     * Name of the report.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * Orientation of the report. Allowed values: `landscape`, `portrait`. Defaults to `landscape`.
     */
    public readonly orientation!: pulumi.Output<string | undefined>;
    /**
     * List of recipients of the report.
     */
    public readonly recipients!: pulumi.Output<string[]>;
    /**
     * Reply-to email address of the report.
     */
    public readonly replyTo!: pulumi.Output<string | undefined>;
    /**
     * Schedule of the report.
     */
    public readonly schedule!: pulumi.Output<outputs.ReportSchedule>;

    /**
     * Create a Report resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated grafana.index/report.Report has been deprecated in favor of grafana.enterprise/report.Report */
    constructor(name: string, args: ReportArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated grafana.index/report.Report has been deprecated in favor of grafana.enterprise/report.Report */
    constructor(name: string, argsOrState?: ReportArgs | ReportState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Report is deprecated: grafana.index/report.Report has been deprecated in favor of grafana.enterprise/report.Report")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReportState | undefined;
            resourceInputs["dashboards"] = state ? state.dashboards : undefined;
            resourceInputs["formats"] = state ? state.formats : undefined;
            resourceInputs["includeDashboardLink"] = state ? state.includeDashboardLink : undefined;
            resourceInputs["includeTableCsv"] = state ? state.includeTableCsv : undefined;
            resourceInputs["layout"] = state ? state.layout : undefined;
            resourceInputs["message"] = state ? state.message : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["orientation"] = state ? state.orientation : undefined;
            resourceInputs["recipients"] = state ? state.recipients : undefined;
            resourceInputs["replyTo"] = state ? state.replyTo : undefined;
            resourceInputs["schedule"] = state ? state.schedule : undefined;
        } else {
            const args = argsOrState as ReportArgs | undefined;
            if ((!args || args.recipients === undefined) && !opts.urn) {
                throw new Error("Missing required property 'recipients'");
            }
            if ((!args || args.schedule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schedule'");
            }
            resourceInputs["dashboards"] = args ? args.dashboards : undefined;
            resourceInputs["formats"] = args ? args.formats : undefined;
            resourceInputs["includeDashboardLink"] = args ? args.includeDashboardLink : undefined;
            resourceInputs["includeTableCsv"] = args ? args.includeTableCsv : undefined;
            resourceInputs["layout"] = args ? args.layout : undefined;
            resourceInputs["message"] = args ? args.message : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["orientation"] = args ? args.orientation : undefined;
            resourceInputs["recipients"] = args ? args.recipients : undefined;
            resourceInputs["replyTo"] = args ? args.replyTo : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "grafana:index/report:Report" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Report.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Report resources.
 */
export interface ReportState {
    /**
     * List of dashboards to render into the report
     */
    dashboards?: pulumi.Input<pulumi.Input<inputs.ReportDashboard>[]>;
    /**
     * Specifies what kind of attachment to generate for the report. Allowed values: `pdf`, `csv`, `image`.
     */
    formats?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to include a link to the dashboard in the report. Defaults to `true`.
     */
    includeDashboardLink?: pulumi.Input<boolean>;
    /**
     * Whether to include a CSV file of table panel data. Defaults to `false`.
     */
    includeTableCsv?: pulumi.Input<boolean>;
    /**
     * Layout of the report. Allowed values: `simple`, `grid`. Defaults to `grid`.
     */
    layout?: pulumi.Input<string>;
    /**
     * Message to be sent in the report.
     */
    message?: pulumi.Input<string>;
    /**
     * Name of the report.
     */
    name?: pulumi.Input<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: pulumi.Input<string>;
    /**
     * Orientation of the report. Allowed values: `landscape`, `portrait`. Defaults to `landscape`.
     */
    orientation?: pulumi.Input<string>;
    /**
     * List of recipients of the report.
     */
    recipients?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Reply-to email address of the report.
     */
    replyTo?: pulumi.Input<string>;
    /**
     * Schedule of the report.
     */
    schedule?: pulumi.Input<inputs.ReportSchedule>;
}

/**
 * The set of arguments for constructing a Report resource.
 */
export interface ReportArgs {
    /**
     * List of dashboards to render into the report
     */
    dashboards?: pulumi.Input<pulumi.Input<inputs.ReportDashboard>[]>;
    /**
     * Specifies what kind of attachment to generate for the report. Allowed values: `pdf`, `csv`, `image`.
     */
    formats?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to include a link to the dashboard in the report. Defaults to `true`.
     */
    includeDashboardLink?: pulumi.Input<boolean>;
    /**
     * Whether to include a CSV file of table panel data. Defaults to `false`.
     */
    includeTableCsv?: pulumi.Input<boolean>;
    /**
     * Layout of the report. Allowed values: `simple`, `grid`. Defaults to `grid`.
     */
    layout?: pulumi.Input<string>;
    /**
     * Message to be sent in the report.
     */
    message?: pulumi.Input<string>;
    /**
     * Name of the report.
     */
    name?: pulumi.Input<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: pulumi.Input<string>;
    /**
     * Orientation of the report. Allowed values: `landscape`, `portrait`. Defaults to `landscape`.
     */
    orientation?: pulumi.Input<string>;
    /**
     * List of recipients of the report.
     */
    recipients: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Reply-to email address of the report.
     */
    replyTo?: pulumi.Input<string>;
    /**
     * Schedule of the report.
     */
    schedule: pulumi.Input<inputs.ReportSchedule>;
}
