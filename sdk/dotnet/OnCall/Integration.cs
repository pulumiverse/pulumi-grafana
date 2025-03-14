// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.OnCall
{
    /// <summary>
    /// * [Official documentation](https://grafana.com/docs/oncall/latest/configure/integrations/)
    /// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Grafana = Pulumiverse.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test_acc_integration = new Grafana.OnCall.Integration("test-acc-integration", new()
    ///     {
    ///         Name = "my integration",
    ///         Type = "grafana",
    ///         DefaultRoute = null,
    ///     });
    /// 
    ///     // Also it's possible to manage integration templates.
    ///     // Check docs to see all available templates.
    ///     var integrationWithTemplates = new Grafana.OnCall.Integration("integration_with_templates", new()
    ///     {
    ///         Name = "integration_with_templates",
    ///         Type = "webhook",
    ///         DefaultRoute = null,
    ///         Templates = new Grafana.OnCall.Inputs.IntegrationTemplatesArgs
    ///         {
    ///             GroupingKey = "{{ payload.group_id }}",
    ///             Slack = new Grafana.OnCall.Inputs.IntegrationTemplatesSlackArgs
    ///             {
    ///                 Title = "Slack title",
    ///                 Message = @"This is example of multiline template
    /// {{ payload.message }}
    /// ",
    ///                 ImageUrl = "{{ payload.image_url }}",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:onCall/integration:Integration name "{{ id }}"
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:onCall/integration:Integration")]
    public partial class Integration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Default route for all alerts from the given integration
        /// </summary>
        [Output("defaultRoute")]
        public Output<Outputs.IntegrationDefaultRoute> DefaultRoute { get; private set; } = null!;

        /// <summary>
        /// The link for using in an integrated tool.
        /// </summary>
        [Output("link")]
        public Output<string> Link { get; private set; } = null!;

        /// <summary>
        /// The name of the service integration.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana.onCall.getTeam` datasource.
        /// </summary>
        [Output("teamId")]
        public Output<string?> TeamId { get; private set; } = null!;

        /// <summary>
        /// Jinja2 templates for Alert payload. An empty templates block will be ignored.
        /// </summary>
        [Output("templates")]
        public Output<Outputs.IntegrationTemplates?> Templates { get; private set; } = null!;

        /// <summary>
        /// The type of integration. Can be grafana, grafana*alerting, webhook, alertmanager, kapacitor, fabric, newrelic, datadog, pagerduty, pingdom, elastalert, amazon*sns, curler, sentry, formatted*webhook, heartbeat, demo, manual, stackdriver, uptimerobot, sentry*platform, zabbix, prtg, slack*channel, inbound*email, direct_paging, jira, zendesk.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Integration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Integration(string name, IntegrationArgs args, CustomResourceOptions? options = null)
            : base("grafana:onCall/integration:Integration", name, args ?? new IntegrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Integration(string name, Input<string> id, IntegrationState? state = null, CustomResourceOptions? options = null)
            : base("grafana:onCall/integration:Integration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "grafana:index/oncallIntegration:OncallIntegration" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Integration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Integration Get(string name, Input<string> id, IntegrationState? state = null, CustomResourceOptions? options = null)
        {
            return new Integration(name, id, state, options);
        }
    }

    public sealed class IntegrationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Default route for all alerts from the given integration
        /// </summary>
        [Input("defaultRoute", required: true)]
        public Input<Inputs.IntegrationDefaultRouteArgs> DefaultRoute { get; set; } = null!;

        /// <summary>
        /// The name of the service integration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana.onCall.getTeam` datasource.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        /// <summary>
        /// Jinja2 templates for Alert payload. An empty templates block will be ignored.
        /// </summary>
        [Input("templates")]
        public Input<Inputs.IntegrationTemplatesArgs>? Templates { get; set; }

        /// <summary>
        /// The type of integration. Can be grafana, grafana*alerting, webhook, alertmanager, kapacitor, fabric, newrelic, datadog, pagerduty, pingdom, elastalert, amazon*sns, curler, sentry, formatted*webhook, heartbeat, demo, manual, stackdriver, uptimerobot, sentry*platform, zabbix, prtg, slack*channel, inbound*email, direct_paging, jira, zendesk.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public IntegrationArgs()
        {
        }
        public static new IntegrationArgs Empty => new IntegrationArgs();
    }

    public sealed class IntegrationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Default route for all alerts from the given integration
        /// </summary>
        [Input("defaultRoute")]
        public Input<Inputs.IntegrationDefaultRouteGetArgs>? DefaultRoute { get; set; }

        /// <summary>
        /// The link for using in an integrated tool.
        /// </summary>
        [Input("link")]
        public Input<string>? Link { get; set; }

        /// <summary>
        /// The name of the service integration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana.onCall.getTeam` datasource.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        /// <summary>
        /// Jinja2 templates for Alert payload. An empty templates block will be ignored.
        /// </summary>
        [Input("templates")]
        public Input<Inputs.IntegrationTemplatesGetArgs>? Templates { get; set; }

        /// <summary>
        /// The type of integration. Can be grafana, grafana*alerting, webhook, alertmanager, kapacitor, fabric, newrelic, datadog, pagerduty, pingdom, elastalert, amazon*sns, curler, sentry, formatted*webhook, heartbeat, demo, manual, stackdriver, uptimerobot, sentry*platform, zabbix, prtg, slack*channel, inbound*email, direct_paging, jira, zendesk.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public IntegrationState()
        {
        }
        public static new IntegrationState Empty => new IntegrationState();
    }
}
