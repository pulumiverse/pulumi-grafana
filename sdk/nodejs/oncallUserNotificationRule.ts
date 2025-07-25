// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/personal_notification_rules/)
 *
 * **Note**: you must be running Grafana OnCall >= v1.8.0 to use this resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const myUser = grafana.onCall.getUser({
 *     username: "my_username",
 * });
 * const myUserStep1 = new grafana.oncall.UserNotificationRule("my_user_step_1", {
 *     userId: myUser.then(myUser => myUser.id),
 *     position: 0,
 *     type: "notify_by_mobile_app",
 * });
 * const myUserStep2 = new grafana.oncall.UserNotificationRule("my_user_step_2", {
 *     userId: myUser.then(myUser => myUser.id),
 *     position: 1,
 *     duration: 600,
 *     type: "wait",
 * });
 * const myUserStep3 = new grafana.oncall.UserNotificationRule("my_user_step_3", {
 *     userId: myUser.then(myUser => myUser.id),
 *     position: 2,
 *     type: "notify_by_phone_call",
 * });
 * const myUserStep4 = new grafana.oncall.UserNotificationRule("my_user_step_4", {
 *     userId: myUser.then(myUser => myUser.id),
 *     position: 3,
 *     duration: 300,
 *     type: "wait",
 * });
 * const myUserStep5 = new grafana.oncall.UserNotificationRule("my_user_step_5", {
 *     userId: myUser.then(myUser => myUser.id),
 *     position: 4,
 *     type: "notify_by_slack",
 * });
 * const myUserImportantStep1 = new grafana.oncall.UserNotificationRule("my_user_important_step_1", {
 *     userId: myUser.then(myUser => myUser.id),
 *     important: true,
 *     position: 0,
 *     type: "notify_by_mobile_app_critical",
 * });
 * const myUserImportantStep2 = new grafana.oncall.UserNotificationRule("my_user_important_step_2", {
 *     userId: myUser.then(myUser => myUser.id),
 *     important: true,
 *     position: 1,
 *     duration: 300,
 *     type: "wait",
 * });
 * const myUserImportantStep3 = new grafana.oncall.UserNotificationRule("my_user_important_step_3", {
 *     userId: myUser.then(myUser => myUser.id),
 *     important: true,
 *     position: 2,
 *     type: "notify_by_mobile_app_critical",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:index/oncallUserNotificationRule:OncallUserNotificationRule name "{{ id }}"
 * ```
 *
 * @deprecated grafana.index/oncallusernotificationrule.OncallUserNotificationRule has been deprecated in favor of grafana.oncall/usernotificationrule.UserNotificationRule
 */
export class OncallUserNotificationRule extends pulumi.CustomResource {
    /**
     * Get an existing OncallUserNotificationRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OncallUserNotificationRuleState, opts?: pulumi.CustomResourceOptions): OncallUserNotificationRule {
        pulumi.log.warn("OncallUserNotificationRule is deprecated: grafana.index/oncallusernotificationrule.OncallUserNotificationRule has been deprecated in favor of grafana.oncall/usernotificationrule.UserNotificationRule")
        return new OncallUserNotificationRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/oncallUserNotificationRule:OncallUserNotificationRule';

    /**
     * Returns true if the given object is an instance of OncallUserNotificationRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OncallUserNotificationRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OncallUserNotificationRule.__pulumiType;
    }

    /**
     * A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
     */
    public readonly duration!: pulumi.Output<number | undefined>;
    /**
     * Boolean value which indicates if a rule is “important”
     */
    public readonly important!: pulumi.Output<boolean>;
    /**
     * Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
     */
    public readonly position!: pulumi.Output<number | undefined>;
    /**
     * The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notifyByMsteams` is only available for Grafana Cloud customers.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * User ID
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a OncallUserNotificationRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated grafana.index/oncallusernotificationrule.OncallUserNotificationRule has been deprecated in favor of grafana.oncall/usernotificationrule.UserNotificationRule */
    constructor(name: string, args: OncallUserNotificationRuleArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated grafana.index/oncallusernotificationrule.OncallUserNotificationRule has been deprecated in favor of grafana.oncall/usernotificationrule.UserNotificationRule */
    constructor(name: string, argsOrState?: OncallUserNotificationRuleArgs | OncallUserNotificationRuleState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("OncallUserNotificationRule is deprecated: grafana.index/oncallusernotificationrule.OncallUserNotificationRule has been deprecated in favor of grafana.oncall/usernotificationrule.UserNotificationRule")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OncallUserNotificationRuleState | undefined;
            resourceInputs["duration"] = state ? state.duration : undefined;
            resourceInputs["important"] = state ? state.important : undefined;
            resourceInputs["position"] = state ? state.position : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as OncallUserNotificationRuleArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["duration"] = args ? args.duration : undefined;
            resourceInputs["important"] = args ? args.important : undefined;
            resourceInputs["position"] = args ? args.position : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OncallUserNotificationRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OncallUserNotificationRule resources.
 */
export interface OncallUserNotificationRuleState {
    /**
     * A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
     */
    duration?: pulumi.Input<number>;
    /**
     * Boolean value which indicates if a rule is “important”
     */
    important?: pulumi.Input<boolean>;
    /**
     * Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
     */
    position?: pulumi.Input<number>;
    /**
     * The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notifyByMsteams` is only available for Grafana Cloud customers.
     */
    type?: pulumi.Input<string>;
    /**
     * User ID
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OncallUserNotificationRule resource.
 */
export interface OncallUserNotificationRuleArgs {
    /**
     * A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
     */
    duration?: pulumi.Input<number>;
    /**
     * Boolean value which indicates if a rule is “important”
     */
    important?: pulumi.Input<boolean>;
    /**
     * Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
     */
    position?: pulumi.Input<number>;
    /**
     * The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notifyByMsteams` is only available for Grafana Cloud customers.
     */
    type: pulumi.Input<string>;
    /**
     * User ID
     */
    userId: pulumi.Input<string>;
}
