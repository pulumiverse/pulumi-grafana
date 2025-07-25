// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * A holiday describes time periods where a time series is expected to behave differently to normal.
 *
 * To use a holiday in a job, use its id in the `holidays` attribute of a `grafana.machineLearning.Job`:
 *
 * ### iCal Holiday
 *
 * This holiday uses an iCal file to define the holidays.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const ical = new grafana.machinelearning.Holiday("ical", {
 *     name: "My iCal holiday",
 *     description: "My Holiday",
 *     icalUrl: "https://calendar.google.com/calendar/ical/en.uk%23holiday%40group.v.calendar.google.com/public/basic.ics",
 *     icalTimezone: "Europe/London",
 * });
 * ```
 *
 * ### Custom Periods Holiday
 *
 * This holiday uses custom periods to define the holidays.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const customPeriods = new grafana.machinelearning.Holiday("custom_periods", {
 *     name: "My custom periods holiday",
 *     description: "My Holiday",
 *     customPeriods: [
 *         {
 *             name: "First of January",
 *             startTime: "2023-01-01T00:00:00Z",
 *             endTime: "2023-01-02T00:00:00Z",
 *         },
 *         {
 *             name: "First of Feburary",
 *             startTime: "2023-02-01T00:00:00Z",
 *             endTime: "2023-02-02T00:00:00Z",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:index/machineLearningHoliday:MachineLearningHoliday name "{{ id }}"
 * ```
 *
 * @deprecated grafana.index/machinelearningholiday.MachineLearningHoliday has been deprecated in favor of grafana.machinelearning/holiday.Holiday
 */
export class MachineLearningHoliday extends pulumi.CustomResource {
    /**
     * Get an existing MachineLearningHoliday resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MachineLearningHolidayState, opts?: pulumi.CustomResourceOptions): MachineLearningHoliday {
        pulumi.log.warn("MachineLearningHoliday is deprecated: grafana.index/machinelearningholiday.MachineLearningHoliday has been deprecated in favor of grafana.machinelearning/holiday.Holiday")
        return new MachineLearningHoliday(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/machineLearningHoliday:MachineLearningHoliday';

    /**
     * Returns true if the given object is an instance of MachineLearningHoliday.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MachineLearningHoliday {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MachineLearningHoliday.__pulumiType;
    }

    /**
     * A list of custom periods for the holiday.
     */
    public readonly customPeriods!: pulumi.Output<outputs.MachineLearningHolidayCustomPeriod[] | undefined>;
    /**
     * A description of the holiday.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The timezone to use for events in the iCal file pointed to by ical_url.
     */
    public readonly icalTimezone!: pulumi.Output<string | undefined>;
    /**
     * A URL to an iCal file containing all occurrences of the holiday.
     */
    public readonly icalUrl!: pulumi.Output<string | undefined>;
    /**
     * The name of the holiday.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a MachineLearningHoliday resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated grafana.index/machinelearningholiday.MachineLearningHoliday has been deprecated in favor of grafana.machinelearning/holiday.Holiday */
    constructor(name: string, args?: MachineLearningHolidayArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated grafana.index/machinelearningholiday.MachineLearningHoliday has been deprecated in favor of grafana.machinelearning/holiday.Holiday */
    constructor(name: string, argsOrState?: MachineLearningHolidayArgs | MachineLearningHolidayState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("MachineLearningHoliday is deprecated: grafana.index/machinelearningholiday.MachineLearningHoliday has been deprecated in favor of grafana.machinelearning/holiday.Holiday")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MachineLearningHolidayState | undefined;
            resourceInputs["customPeriods"] = state ? state.customPeriods : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["icalTimezone"] = state ? state.icalTimezone : undefined;
            resourceInputs["icalUrl"] = state ? state.icalUrl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as MachineLearningHolidayArgs | undefined;
            resourceInputs["customPeriods"] = args ? args.customPeriods : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["icalTimezone"] = args ? args.icalTimezone : undefined;
            resourceInputs["icalUrl"] = args ? args.icalUrl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MachineLearningHoliday.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MachineLearningHoliday resources.
 */
export interface MachineLearningHolidayState {
    /**
     * A list of custom periods for the holiday.
     */
    customPeriods?: pulumi.Input<pulumi.Input<inputs.MachineLearningHolidayCustomPeriod>[]>;
    /**
     * A description of the holiday.
     */
    description?: pulumi.Input<string>;
    /**
     * The timezone to use for events in the iCal file pointed to by ical_url.
     */
    icalTimezone?: pulumi.Input<string>;
    /**
     * A URL to an iCal file containing all occurrences of the holiday.
     */
    icalUrl?: pulumi.Input<string>;
    /**
     * The name of the holiday.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MachineLearningHoliday resource.
 */
export interface MachineLearningHolidayArgs {
    /**
     * A list of custom periods for the holiday.
     */
    customPeriods?: pulumi.Input<pulumi.Input<inputs.MachineLearningHolidayCustomPeriod>[]>;
    /**
     * A description of the holiday.
     */
    description?: pulumi.Input<string>;
    /**
     * The timezone to use for events in the iCal file pointed to by ical_url.
     */
    icalTimezone?: pulumi.Input<string>;
    /**
     * A URL to an iCal file containing all occurrences of the holiday.
     */
    icalUrl?: pulumi.Input<string>;
    /**
     * The name of the holiday.
     */
    name?: pulumi.Input<string>;
}
