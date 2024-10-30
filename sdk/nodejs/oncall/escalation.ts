// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * * [Official documentation](https://grafana.com/docs/oncall/latest/configure/escalation-chains-and-routes/)
 * * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/escalation_policies/)
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:onCall/escalation:Escalation name "{{ id }}"
 * ```
 */
export class Escalation extends pulumi.CustomResource {
    /**
     * Get an existing Escalation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EscalationState, opts?: pulumi.CustomResourceOptions): Escalation {
        return new Escalation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:onCall/escalation:Escalation';

    /**
     * Returns true if the given object is an instance of Escalation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Escalation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Escalation.__pulumiType;
    }

    /**
     * The ID of an Action for triggerWebhook type step.
     */
    public readonly actionToTrigger!: pulumi.Output<string | undefined>;
    /**
     * The duration of delay for wait type step.
     */
    public readonly duration!: pulumi.Output<number | undefined>;
    /**
     * The ID of the escalation chain.
     */
    public readonly escalationChainId!: pulumi.Output<string>;
    /**
     * The ID of a User Group for notify*user*group type step.
     */
    public readonly groupToNotify!: pulumi.Output<string | undefined>;
    /**
     * Will activate "important" personal notification rules. Actual for steps: notify*persons, notify*on*call*from*schedule and notify*user*group,notify*team_members
     */
    public readonly important!: pulumi.Output<boolean | undefined>;
    /**
     * The beginning of the time interval for notify*if*time*from*to type step in UTC (for example 08:00:00Z).
     */
    public readonly notifyIfTimeFrom!: pulumi.Output<string | undefined>;
    /**
     * The end of the time interval for notify*if*time*from*to type step in UTC (for example 18:00:00Z).
     */
    public readonly notifyIfTimeTo!: pulumi.Output<string | undefined>;
    /**
     * ID of a Schedule for notify*on*call*from*schedule type step.
     */
    public readonly notifyOnCallFromSchedule!: pulumi.Output<string | undefined>;
    /**
     * The ID of a Team for a notify*team*members type step.
     */
    public readonly notifyToTeamMembers!: pulumi.Output<string | undefined>;
    /**
     * The list of ID's of users for notifyPersons type step.
     */
    public readonly personsToNotifies!: pulumi.Output<string[] | undefined>;
    /**
     * The list of ID's of users for notify*person*next*each*time type step.
     */
    public readonly personsToNotifyNextEachTimes!: pulumi.Output<string[] | undefined>;
    /**
     * The position of the escalation step (starts from 0).
     */
    public readonly position!: pulumi.Output<number>;
    /**
     * The type of escalation policy. Can be wait, notify*persons, notify*person*next*each*time, notify*on*call*from*schedule, trigger*webhook, notify*user*group, resolve, notify*whole*channel, notify*if*time*from*to, repeat*escalation, notify*team_members
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Escalation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EscalationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EscalationArgs | EscalationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EscalationState | undefined;
            resourceInputs["actionToTrigger"] = state ? state.actionToTrigger : undefined;
            resourceInputs["duration"] = state ? state.duration : undefined;
            resourceInputs["escalationChainId"] = state ? state.escalationChainId : undefined;
            resourceInputs["groupToNotify"] = state ? state.groupToNotify : undefined;
            resourceInputs["important"] = state ? state.important : undefined;
            resourceInputs["notifyIfTimeFrom"] = state ? state.notifyIfTimeFrom : undefined;
            resourceInputs["notifyIfTimeTo"] = state ? state.notifyIfTimeTo : undefined;
            resourceInputs["notifyOnCallFromSchedule"] = state ? state.notifyOnCallFromSchedule : undefined;
            resourceInputs["notifyToTeamMembers"] = state ? state.notifyToTeamMembers : undefined;
            resourceInputs["personsToNotifies"] = state ? state.personsToNotifies : undefined;
            resourceInputs["personsToNotifyNextEachTimes"] = state ? state.personsToNotifyNextEachTimes : undefined;
            resourceInputs["position"] = state ? state.position : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as EscalationArgs | undefined;
            if ((!args || args.escalationChainId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'escalationChainId'");
            }
            if ((!args || args.position === undefined) && !opts.urn) {
                throw new Error("Missing required property 'position'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["actionToTrigger"] = args ? args.actionToTrigger : undefined;
            resourceInputs["duration"] = args ? args.duration : undefined;
            resourceInputs["escalationChainId"] = args ? args.escalationChainId : undefined;
            resourceInputs["groupToNotify"] = args ? args.groupToNotify : undefined;
            resourceInputs["important"] = args ? args.important : undefined;
            resourceInputs["notifyIfTimeFrom"] = args ? args.notifyIfTimeFrom : undefined;
            resourceInputs["notifyIfTimeTo"] = args ? args.notifyIfTimeTo : undefined;
            resourceInputs["notifyOnCallFromSchedule"] = args ? args.notifyOnCallFromSchedule : undefined;
            resourceInputs["notifyToTeamMembers"] = args ? args.notifyToTeamMembers : undefined;
            resourceInputs["personsToNotifies"] = args ? args.personsToNotifies : undefined;
            resourceInputs["personsToNotifyNextEachTimes"] = args ? args.personsToNotifyNextEachTimes : undefined;
            resourceInputs["position"] = args ? args.position : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "grafana:index/oncallEscalation:OncallEscalation" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Escalation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Escalation resources.
 */
export interface EscalationState {
    /**
     * The ID of an Action for triggerWebhook type step.
     */
    actionToTrigger?: pulumi.Input<string>;
    /**
     * The duration of delay for wait type step.
     */
    duration?: pulumi.Input<number>;
    /**
     * The ID of the escalation chain.
     */
    escalationChainId?: pulumi.Input<string>;
    /**
     * The ID of a User Group for notify*user*group type step.
     */
    groupToNotify?: pulumi.Input<string>;
    /**
     * Will activate "important" personal notification rules. Actual for steps: notify*persons, notify*on*call*from*schedule and notify*user*group,notify*team_members
     */
    important?: pulumi.Input<boolean>;
    /**
     * The beginning of the time interval for notify*if*time*from*to type step in UTC (for example 08:00:00Z).
     */
    notifyIfTimeFrom?: pulumi.Input<string>;
    /**
     * The end of the time interval for notify*if*time*from*to type step in UTC (for example 18:00:00Z).
     */
    notifyIfTimeTo?: pulumi.Input<string>;
    /**
     * ID of a Schedule for notify*on*call*from*schedule type step.
     */
    notifyOnCallFromSchedule?: pulumi.Input<string>;
    /**
     * The ID of a Team for a notify*team*members type step.
     */
    notifyToTeamMembers?: pulumi.Input<string>;
    /**
     * The list of ID's of users for notifyPersons type step.
     */
    personsToNotifies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of ID's of users for notify*person*next*each*time type step.
     */
    personsToNotifyNextEachTimes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The position of the escalation step (starts from 0).
     */
    position?: pulumi.Input<number>;
    /**
     * The type of escalation policy. Can be wait, notify*persons, notify*person*next*each*time, notify*on*call*from*schedule, trigger*webhook, notify*user*group, resolve, notify*whole*channel, notify*if*time*from*to, repeat*escalation, notify*team_members
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Escalation resource.
 */
export interface EscalationArgs {
    /**
     * The ID of an Action for triggerWebhook type step.
     */
    actionToTrigger?: pulumi.Input<string>;
    /**
     * The duration of delay for wait type step.
     */
    duration?: pulumi.Input<number>;
    /**
     * The ID of the escalation chain.
     */
    escalationChainId: pulumi.Input<string>;
    /**
     * The ID of a User Group for notify*user*group type step.
     */
    groupToNotify?: pulumi.Input<string>;
    /**
     * Will activate "important" personal notification rules. Actual for steps: notify*persons, notify*on*call*from*schedule and notify*user*group,notify*team_members
     */
    important?: pulumi.Input<boolean>;
    /**
     * The beginning of the time interval for notify*if*time*from*to type step in UTC (for example 08:00:00Z).
     */
    notifyIfTimeFrom?: pulumi.Input<string>;
    /**
     * The end of the time interval for notify*if*time*from*to type step in UTC (for example 18:00:00Z).
     */
    notifyIfTimeTo?: pulumi.Input<string>;
    /**
     * ID of a Schedule for notify*on*call*from*schedule type step.
     */
    notifyOnCallFromSchedule?: pulumi.Input<string>;
    /**
     * The ID of a Team for a notify*team*members type step.
     */
    notifyToTeamMembers?: pulumi.Input<string>;
    /**
     * The list of ID's of users for notifyPersons type step.
     */
    personsToNotifies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of ID's of users for notify*person*next*each*time type step.
     */
    personsToNotifyNextEachTimes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The position of the escalation step (starts from 0).
     */
    position: pulumi.Input<number>;
    /**
     * The type of escalation policy. Can be wait, notify*persons, notify*person*next*each*time, notify*on*call*from*schedule, trigger*webhook, notify*user*group, resolve, notify*whole*channel, notify*if*time*from*to, repeat*escalation, notify*team_members
     */
    type: pulumi.Input<string>;
}
