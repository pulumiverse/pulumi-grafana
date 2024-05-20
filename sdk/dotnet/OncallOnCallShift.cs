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
    [GrafanaResourceType("grafana:index/oncallOnCallShift:OncallOnCallShift")]
    public partial class OncallOnCallShift : global::Pulumi.CustomResource
    {
        /// <summary>
        /// This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
        /// </summary>
        [Output("byDays")]
        public Output<ImmutableArray<string>> ByDays { get; private set; } = null!;

        /// <summary>
        /// This parameter takes a list of days of the month. Valid values are 1 to 31 or -31 to -1
        /// </summary>
        [Output("byMonthdays")]
        public Output<ImmutableArray<int>> ByMonthdays { get; private set; } = null!;

        /// <summary>
        /// This parameter takes a list of months. Valid values are 1 to 12
        /// </summary>
        [Output("byMonths")]
        public Output<ImmutableArray<int>> ByMonths { get; private set; } = null!;

        /// <summary>
        /// The duration of the event.
        /// </summary>
        [Output("duration")]
        public Output<int> Duration { get; private set; } = null!;

        /// <summary>
        /// The frequency of the event. Can be hourly, daily, weekly, monthly
        /// </summary>
        [Output("frequency")]
        public Output<string?> Frequency { get; private set; } = null!;

        /// <summary>
        /// The positive integer representing at which intervals the recurrence rule repeats.
        /// </summary>
        [Output("interval")]
        public Output<int?> Interval { get; private set; } = null!;

        /// <summary>
        /// The priority level. The higher the value, the higher the priority.
        /// </summary>
        [Output("level")]
        public Output<int?> Level { get; private set; } = null!;

        /// <summary>
        /// The shift's name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The list of lists with on-call users (for rolling_users event type)
        /// </summary>
        [Output("rollingUsers")]
        public Output<ImmutableArray<ImmutableArray<string>>> RollingUsers { get; private set; } = null!;

        /// <summary>
        /// The start time of the on-call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example
        /// "2020-09-05T08:00:00")
        /// </summary>
        [Output("start")]
        public Output<string> Start { get; private set; } = null!;

        /// <summary>
        /// The index of the list of users in rolling_users, from which on-call rotation starts.
        /// </summary>
        [Output("startRotationFromUserIndex")]
        public Output<int?> StartRotationFromUserIndex { get; private set; } = null!;

        /// <summary>
        /// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team
        /// with OnCall). You can then get the ID using the `grafana.getOncallTeam` datasource.
        /// </summary>
        [Output("teamId")]
        public Output<string?> TeamId { get; private set; } = null!;

        /// <summary>
        /// The shift's timezone. Overrides schedule's timezone.
        /// </summary>
        [Output("timeZone")]
        public Output<string?> TimeZone { get; private set; } = null!;

        /// <summary>
        /// The shift's type. Can be rolling_users, recurrent_event, single_event
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The list of on-call users (for single_event and recurrent_event event type).
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<string>> Users { get; private set; } = null!;

        /// <summary>
        /// Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
        /// </summary>
        [Output("weekStart")]
        public Output<string?> WeekStart { get; private set; } = null!;


        /// <summary>
        /// Create a OncallOnCallShift resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OncallOnCallShift(string name, OncallOnCallShiftArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/oncallOnCallShift:OncallOnCallShift", name, args ?? new OncallOnCallShiftArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OncallOnCallShift(string name, Input<string> id, OncallOnCallShiftState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/oncallOnCallShift:OncallOnCallShift", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OncallOnCallShift resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OncallOnCallShift Get(string name, Input<string> id, OncallOnCallShiftState? state = null, CustomResourceOptions? options = null)
        {
            return new OncallOnCallShift(name, id, state, options);
        }
    }

    public sealed class OncallOnCallShiftArgs : global::Pulumi.ResourceArgs
    {
        [Input("byDays")]
        private InputList<string>? _byDays;

        /// <summary>
        /// This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
        /// </summary>
        public InputList<string> ByDays
        {
            get => _byDays ?? (_byDays = new InputList<string>());
            set => _byDays = value;
        }

        [Input("byMonthdays")]
        private InputList<int>? _byMonthdays;

        /// <summary>
        /// This parameter takes a list of days of the month. Valid values are 1 to 31 or -31 to -1
        /// </summary>
        public InputList<int> ByMonthdays
        {
            get => _byMonthdays ?? (_byMonthdays = new InputList<int>());
            set => _byMonthdays = value;
        }

        [Input("byMonths")]
        private InputList<int>? _byMonths;

        /// <summary>
        /// This parameter takes a list of months. Valid values are 1 to 12
        /// </summary>
        public InputList<int> ByMonths
        {
            get => _byMonths ?? (_byMonths = new InputList<int>());
            set => _byMonths = value;
        }

        /// <summary>
        /// The duration of the event.
        /// </summary>
        [Input("duration", required: true)]
        public Input<int> Duration { get; set; } = null!;

        /// <summary>
        /// The frequency of the event. Can be hourly, daily, weekly, monthly
        /// </summary>
        [Input("frequency")]
        public Input<string>? Frequency { get; set; }

        /// <summary>
        /// The positive integer representing at which intervals the recurrence rule repeats.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// The priority level. The higher the value, the higher the priority.
        /// </summary>
        [Input("level")]
        public Input<int>? Level { get; set; }

        /// <summary>
        /// The shift's name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("rollingUsers")]
        private InputList<ImmutableArray<string>>? _rollingUsers;

        /// <summary>
        /// The list of lists with on-call users (for rolling_users event type)
        /// </summary>
        public InputList<ImmutableArray<string>> RollingUsers
        {
            get => _rollingUsers ?? (_rollingUsers = new InputList<ImmutableArray<string>>());
            set => _rollingUsers = value;
        }

        /// <summary>
        /// The start time of the on-call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example
        /// "2020-09-05T08:00:00")
        /// </summary>
        [Input("start", required: true)]
        public Input<string> Start { get; set; } = null!;

        /// <summary>
        /// The index of the list of users in rolling_users, from which on-call rotation starts.
        /// </summary>
        [Input("startRotationFromUserIndex")]
        public Input<int>? StartRotationFromUserIndex { get; set; }

        /// <summary>
        /// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team
        /// with OnCall). You can then get the ID using the `grafana.getOncallTeam` datasource.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        /// <summary>
        /// The shift's timezone. Overrides schedule's timezone.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        /// <summary>
        /// The shift's type. Can be rolling_users, recurrent_event, single_event
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("users")]
        private InputList<string>? _users;

        /// <summary>
        /// The list of on-call users (for single_event and recurrent_event event type).
        /// </summary>
        public InputList<string> Users
        {
            get => _users ?? (_users = new InputList<string>());
            set => _users = value;
        }

        /// <summary>
        /// Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
        /// </summary>
        [Input("weekStart")]
        public Input<string>? WeekStart { get; set; }

        public OncallOnCallShiftArgs()
        {
        }
        public static new OncallOnCallShiftArgs Empty => new OncallOnCallShiftArgs();
    }

    public sealed class OncallOnCallShiftState : global::Pulumi.ResourceArgs
    {
        [Input("byDays")]
        private InputList<string>? _byDays;

        /// <summary>
        /// This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
        /// </summary>
        public InputList<string> ByDays
        {
            get => _byDays ?? (_byDays = new InputList<string>());
            set => _byDays = value;
        }

        [Input("byMonthdays")]
        private InputList<int>? _byMonthdays;

        /// <summary>
        /// This parameter takes a list of days of the month. Valid values are 1 to 31 or -31 to -1
        /// </summary>
        public InputList<int> ByMonthdays
        {
            get => _byMonthdays ?? (_byMonthdays = new InputList<int>());
            set => _byMonthdays = value;
        }

        [Input("byMonths")]
        private InputList<int>? _byMonths;

        /// <summary>
        /// This parameter takes a list of months. Valid values are 1 to 12
        /// </summary>
        public InputList<int> ByMonths
        {
            get => _byMonths ?? (_byMonths = new InputList<int>());
            set => _byMonths = value;
        }

        /// <summary>
        /// The duration of the event.
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// The frequency of the event. Can be hourly, daily, weekly, monthly
        /// </summary>
        [Input("frequency")]
        public Input<string>? Frequency { get; set; }

        /// <summary>
        /// The positive integer representing at which intervals the recurrence rule repeats.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// The priority level. The higher the value, the higher the priority.
        /// </summary>
        [Input("level")]
        public Input<int>? Level { get; set; }

        /// <summary>
        /// The shift's name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("rollingUsers")]
        private InputList<ImmutableArray<string>>? _rollingUsers;

        /// <summary>
        /// The list of lists with on-call users (for rolling_users event type)
        /// </summary>
        public InputList<ImmutableArray<string>> RollingUsers
        {
            get => _rollingUsers ?? (_rollingUsers = new InputList<ImmutableArray<string>>());
            set => _rollingUsers = value;
        }

        /// <summary>
        /// The start time of the on-call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example
        /// "2020-09-05T08:00:00")
        /// </summary>
        [Input("start")]
        public Input<string>? Start { get; set; }

        /// <summary>
        /// The index of the list of users in rolling_users, from which on-call rotation starts.
        /// </summary>
        [Input("startRotationFromUserIndex")]
        public Input<int>? StartRotationFromUserIndex { get; set; }

        /// <summary>
        /// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team
        /// with OnCall). You can then get the ID using the `grafana.getOncallTeam` datasource.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        /// <summary>
        /// The shift's timezone. Overrides schedule's timezone.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        /// <summary>
        /// The shift's type. Can be rolling_users, recurrent_event, single_event
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("users")]
        private InputList<string>? _users;

        /// <summary>
        /// The list of on-call users (for single_event and recurrent_event event type).
        /// </summary>
        public InputList<string> Users
        {
            get => _users ?? (_users = new InputList<string>());
            set => _users = value;
        }

        /// <summary>
        /// Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
        /// </summary>
        [Input("weekStart")]
        public Input<string>? WeekStart { get; set; }

        public OncallOnCallShiftState()
        {
        }
        public static new OncallOnCallShiftState Empty => new OncallOnCallShiftState();
    }
}
