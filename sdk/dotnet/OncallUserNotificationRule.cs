// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana
{
    /// <summary>
    /// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/personal_notification_rules/)
    /// 
    /// **Note**: you must be running Grafana OnCall &gt;= v1.8.0 to use this resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Grafana = Pulumi.Grafana;
    /// using Grafana = Pulumiverse.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myUser = Grafana.OnCall.GetUser.Invoke(new()
    ///     {
    ///         Username = "my_username",
    ///     });
    /// 
    ///     var myUserStep1 = new Grafana.OncallUserNotificationRule("myUserStep1", new()
    ///     {
    ///         UserId = myUser.Apply(getUserResult =&gt; getUserResult.Id),
    ///         Position = 0,
    ///         Type = "notify_by_mobile_app",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = grafana.Oncall,
    ///     });
    /// 
    ///     var myUserStep2 = new Grafana.OncallUserNotificationRule("myUserStep2", new()
    ///     {
    ///         UserId = myUser.Apply(getUserResult =&gt; getUserResult.Id),
    ///         Position = 1,
    ///         Duration = 600,
    ///         Type = "wait",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = grafana.Oncall,
    ///     });
    /// 
    ///     var myUserStep3 = new Grafana.OncallUserNotificationRule("myUserStep3", new()
    ///     {
    ///         UserId = myUser.Apply(getUserResult =&gt; getUserResult.Id),
    ///         Position = 2,
    ///         Type = "notify_by_phone_call",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = grafana.Oncall,
    ///     });
    /// 
    ///     var myUserStep4 = new Grafana.OncallUserNotificationRule("myUserStep4", new()
    ///     {
    ///         UserId = myUser.Apply(getUserResult =&gt; getUserResult.Id),
    ///         Position = 3,
    ///         Duration = 300,
    ///         Type = "wait",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = grafana.Oncall,
    ///     });
    /// 
    ///     var myUserStep5 = new Grafana.OncallUserNotificationRule("myUserStep5", new()
    ///     {
    ///         UserId = myUser.Apply(getUserResult =&gt; getUserResult.Id),
    ///         Position = 4,
    ///         Type = "notify_by_slack",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = grafana.Oncall,
    ///     });
    /// 
    ///     var myUserImportantStep1 = new Grafana.OncallUserNotificationRule("myUserImportantStep1", new()
    ///     {
    ///         UserId = myUser.Apply(getUserResult =&gt; getUserResult.Id),
    ///         Important = true,
    ///         Position = 0,
    ///         Type = "notify_by_mobile_app_critical",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = grafana.Oncall,
    ///     });
    /// 
    ///     var myUserImportantStep2 = new Grafana.OncallUserNotificationRule("myUserImportantStep2", new()
    ///     {
    ///         UserId = myUser.Apply(getUserResult =&gt; getUserResult.Id),
    ///         Important = true,
    ///         Position = 1,
    ///         Duration = 300,
    ///         Type = "wait",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = grafana.Oncall,
    ///     });
    /// 
    ///     var myUserImportantStep3 = new Grafana.OncallUserNotificationRule("myUserImportantStep3", new()
    ///     {
    ///         UserId = myUser.Apply(getUserResult =&gt; getUserResult.Id),
    ///         Important = true,
    ///         Position = 2,
    ///         Type = "notify_by_mobile_app_critical",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = grafana.Oncall,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:index/oncallUserNotificationRule:OncallUserNotificationRule name "{{ id }}"
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:index/oncallUserNotificationRule:OncallUserNotificationRule")]
    public partial class OncallUserNotificationRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
        /// </summary>
        [Output("duration")]
        public Output<int?> Duration { get; private set; } = null!;

        /// <summary>
        /// Boolean value which indicates if a rule is “important”
        /// </summary>
        [Output("important")]
        public Output<bool> Important { get; private set; } = null!;

        /// <summary>
        /// Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
        /// </summary>
        [Output("position")]
        public Output<int?> Position { get; private set; } = null!;

        /// <summary>
        /// The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notify_by_msteams` is only available for Grafana Cloud customers.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// User ID
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a OncallUserNotificationRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OncallUserNotificationRule(string name, OncallUserNotificationRuleArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/oncallUserNotificationRule:OncallUserNotificationRule", name, args ?? new OncallUserNotificationRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OncallUserNotificationRule(string name, Input<string> id, OncallUserNotificationRuleState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/oncallUserNotificationRule:OncallUserNotificationRule", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OncallUserNotificationRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OncallUserNotificationRule Get(string name, Input<string> id, OncallUserNotificationRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new OncallUserNotificationRule(name, id, state, options);
        }
    }

    public sealed class OncallUserNotificationRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// Boolean value which indicates if a rule is “important”
        /// </summary>
        [Input("important")]
        public Input<bool>? Important { get; set; }

        /// <summary>
        /// Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        /// <summary>
        /// The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notify_by_msteams` is only available for Grafana Cloud customers.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// User ID
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public OncallUserNotificationRuleArgs()
        {
        }
        public static new OncallUserNotificationRuleArgs Empty => new OncallUserNotificationRuleArgs();
    }

    public sealed class OncallUserNotificationRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// Boolean value which indicates if a rule is “important”
        /// </summary>
        [Input("important")]
        public Input<bool>? Important { get; set; }

        /// <summary>
        /// Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        /// <summary>
        /// The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notify_by_msteams` is only available for Grafana Cloud customers.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// User ID
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public OncallUserNotificationRuleState()
        {
        }
        public static new OncallUserNotificationRuleState Empty => new OncallUserNotificationRuleState();
    }
}
