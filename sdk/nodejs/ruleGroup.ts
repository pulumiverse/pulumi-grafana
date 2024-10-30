// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manages Grafana Alerting rule groups.
 *
 * * [Official documentation](https://grafana.com/docs/grafana/latest/alerting/alerting-rules/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/alerting_provisioning/#alert-rules)
 *
 * This resource requires Grafana 9.1.0 or later.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const ruleFolder = new grafana.oss.Folder("rule_folder", {title: "My Alert Rule Folder"});
 * const myAlertRule = new grafana.alerting.RuleGroup("my_alert_rule", {
 *     name: "My Rule Group",
 *     folderUid: ruleFolder.uid,
 *     intervalSeconds: 240,
 *     orgId: "1",
 *     rules: [{
 *         name: "My Alert Rule 1",
 *         "for": "2m",
 *         condition: "B",
 *         noDataState: "NoData",
 *         execErrState: "Alerting",
 *         annotations: {
 *             a: "b",
 *             c: "d",
 *         },
 *         labels: {
 *             e: "f",
 *             g: "h",
 *         },
 *         isPaused: false,
 *         datas: [
 *             {
 *                 refId: "A",
 *                 queryType: "",
 *                 relativeTimeRange: {
 *                     from: 600,
 *                     to: 0,
 *                 },
 *                 datasourceUid: "PD8C576611E62080A",
 *                 model: JSON.stringify({
 *                     hide: false,
 *                     intervalMs: 1000,
 *                     maxDataPoints: 43200,
 *                     refId: "A",
 *                 }),
 *             },
 *             {
 *                 refId: "B",
 *                 queryType: "",
 *                 relativeTimeRange: {
 *                     from: 0,
 *                     to: 0,
 *                 },
 *                 datasourceUid: "-100",
 *                 model: `{
 *     "conditions": [
 *         {
 *         "evaluator": {
 *             "params": [
 *             3
 *             ],
 *             "type": "gt"
 *         },
 *         "operator": {
 *             "type": "and"
 *         },
 *         "query": {
 *             "params": [
 *             "A"
 *             ]
 *         },
 *         "reducer": {
 *             "params": [],
 *             "type": "last"
 *         },
 *         "type": "query"
 *         }
 *     ],
 *     "datasource": {
 *         "type": "__expr__",
 *         "uid": "-100"
 *     },
 *     "hide": false,
 *     "intervalMs": 1000,
 *     "maxDataPoints": 43200,
 *     "refId": "B",
 *     "type": "classic_conditions"
 * }
 * `,
 *             },
 *         ],
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:index/ruleGroup:RuleGroup name "{{ folderUID }}:{{ title }}"
 * ```
 *
 * ```sh
 * $ pulumi import grafana:index/ruleGroup:RuleGroup name "{{ orgID }}:{{ folderUID }}:{{ title }}"
 * ```
 *
 * @deprecated grafana.index/rulegroup.RuleGroup has been deprecated in favor of grafana.alerting/rulegroup.RuleGroup
 */
export class RuleGroup extends pulumi.CustomResource {
    /**
     * Get an existing RuleGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleGroupState, opts?: pulumi.CustomResourceOptions): RuleGroup {
        pulumi.log.warn("RuleGroup is deprecated: grafana.index/rulegroup.RuleGroup has been deprecated in favor of grafana.alerting/rulegroup.RuleGroup")
        return new RuleGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/ruleGroup:RuleGroup';

    /**
     * Returns true if the given object is an instance of RuleGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuleGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuleGroup.__pulumiType;
    }

    public readonly disableProvenance!: pulumi.Output<boolean | undefined>;
    /**
     * The UID of the folder that the group belongs to.
     */
    public readonly folderUid!: pulumi.Output<string>;
    /**
     * The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are evaluated sequentially.
     */
    public readonly intervalSeconds!: pulumi.Output<number>;
    /**
     * The name of the rule group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * The rules within the group.
     */
    public readonly rules!: pulumi.Output<outputs.RuleGroupRule[]>;

    /**
     * Create a RuleGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated grafana.index/rulegroup.RuleGroup has been deprecated in favor of grafana.alerting/rulegroup.RuleGroup */
    constructor(name: string, args: RuleGroupArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated grafana.index/rulegroup.RuleGroup has been deprecated in favor of grafana.alerting/rulegroup.RuleGroup */
    constructor(name: string, argsOrState?: RuleGroupArgs | RuleGroupState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("RuleGroup is deprecated: grafana.index/rulegroup.RuleGroup has been deprecated in favor of grafana.alerting/rulegroup.RuleGroup")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleGroupState | undefined;
            resourceInputs["disableProvenance"] = state ? state.disableProvenance : undefined;
            resourceInputs["folderUid"] = state ? state.folderUid : undefined;
            resourceInputs["intervalSeconds"] = state ? state.intervalSeconds : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
        } else {
            const args = argsOrState as RuleGroupArgs | undefined;
            if ((!args || args.folderUid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'folderUid'");
            }
            if ((!args || args.intervalSeconds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'intervalSeconds'");
            }
            if ((!args || args.rules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rules'");
            }
            resourceInputs["disableProvenance"] = args ? args.disableProvenance : undefined;
            resourceInputs["folderUid"] = args ? args.folderUid : undefined;
            resourceInputs["intervalSeconds"] = args ? args.intervalSeconds : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "grafana:index/ruleGroup:RuleGroup" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(RuleGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RuleGroup resources.
 */
export interface RuleGroupState {
    disableProvenance?: pulumi.Input<boolean>;
    /**
     * The UID of the folder that the group belongs to.
     */
    folderUid?: pulumi.Input<string>;
    /**
     * The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are evaluated sequentially.
     */
    intervalSeconds?: pulumi.Input<number>;
    /**
     * The name of the rule group.
     */
    name?: pulumi.Input<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: pulumi.Input<string>;
    /**
     * The rules within the group.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.RuleGroupRule>[]>;
}

/**
 * The set of arguments for constructing a RuleGroup resource.
 */
export interface RuleGroupArgs {
    disableProvenance?: pulumi.Input<boolean>;
    /**
     * The UID of the folder that the group belongs to.
     */
    folderUid: pulumi.Input<string>;
    /**
     * The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are evaluated sequentially.
     */
    intervalSeconds: pulumi.Input<number>;
    /**
     * The name of the rule group.
     */
    name?: pulumi.Input<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: pulumi.Input<string>;
    /**
     * The rules within the group.
     */
    rules: pulumi.Input<pulumi.Input<inputs.RuleGroupRule>[]>;
}
