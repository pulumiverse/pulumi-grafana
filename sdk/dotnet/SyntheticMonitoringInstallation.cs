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
    /// Sets up Synthetic Monitoring on a Grafana cloud stack and generates a token.
    /// Once a Grafana Cloud stack is created, a user can either use this resource or go into the UI to install synthetic monitoring.
    /// This resource cannot be imported but it can be used on an existing Synthetic Monitoring installation without issues.
    /// 
    /// **Note that this resource must be used on a provider configured with Grafana Cloud credentials.**
    /// 
    /// * [Official documentation](https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/)
    /// * [API documentation](https://github.com/grafana/synthetic-monitoring-api-go-client/blob/main/docs/API.md#apiv1registerinstall)
    /// 
    /// Required access policy scopes:
    /// 
    /// * stacks:read
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
    ///     var config = new Config();
    ///     // Cloud Access Policy token for Grafana Cloud with the following scopes: accesspolicies:read|write|delete, stacks:read|write|delete
    ///     var cloudAccessPolicyToken = config.RequireObject&lt;dynamic&gt;("cloudAccessPolicyToken");
    ///     var stackSlug = config.RequireObject&lt;dynamic&gt;("stackSlug");
    ///     var cloudRegion = config.Get("cloudRegion") ?? "prod-us-east-0";
    ///     var smStack = new Grafana.Cloud.Stack("sm_stack", new()
    ///     {
    ///         Name = stackSlug,
    ///         Slug = stackSlug,
    ///         RegionSlug = cloudRegion,
    ///     });
    /// 
    ///     // Step 2: Install Synthetic Monitoring on the stack
    ///     var smMetricsPublish = new Grafana.Cloud.AccessPolicy("sm_metrics_publish", new()
    ///     {
    ///         Region = cloudRegion,
    ///         Name = "metric-publisher-for-sm",
    ///         Scopes = new[]
    ///         {
    ///             "metrics:write",
    ///             "stacks:read",
    ///             "logs:write",
    ///             "traces:write",
    ///         },
    ///         Realms = new[]
    ///         {
    ///             new Grafana.Cloud.Inputs.AccessPolicyRealmArgs
    ///             {
    ///                 Type = "stack",
    ///                 Identifier = smStack.Id,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var smMetricsPublishAccessPolicyToken = new Grafana.Cloud.AccessPolicyToken("sm_metrics_publish", new()
    ///     {
    ///         Region = cloudRegion,
    ///         AccessPolicyId = smMetricsPublish.PolicyId,
    ///         Name = "metric-publisher-for-sm",
    ///     });
    /// 
    ///     var smStackInstallation = new Grafana.SyntheticMonitoring.Installation("sm_stack", new()
    ///     {
    ///         StackId = smStack.Id,
    ///         MetricsPublisherKey = smMetricsPublishAccessPolicyToken.Token,
    ///     });
    /// 
    ///     var main = Grafana.SyntheticMonitoring.GetProbes.Invoke();
    /// 
    /// });
    /// ```
    /// </summary>
    [Obsolete(@"grafana.index/syntheticmonitoringinstallation.SyntheticMonitoringInstallation has been deprecated in favor of grafana.syntheticmonitoring/installation.Installation")]
    [GrafanaResourceType("grafana:index/syntheticMonitoringInstallation:SyntheticMonitoringInstallation")]
    public partial class SyntheticMonitoringInstallation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The [Grafana Cloud access policy](https://grafana.com/docs/grafana-cloud/security-and-account-management/authentication-and-permissions/access-policies/) with the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and logs to Grafana Cloud stack.
        /// </summary>
        [Output("metricsPublisherKey")]
        public Output<string> MetricsPublisherKey { get; private set; } = null!;

        /// <summary>
        /// Generated token to access the SM API.
        /// </summary>
        [Output("smAccessToken")]
        public Output<string> SmAccessToken { get; private set; } = null!;

        /// <summary>
        /// The ID or slug of the stack to install SM on.
        /// </summary>
        [Output("stackId")]
        public Output<string> StackId { get; private set; } = null!;

        /// <summary>
        /// The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/set-up-private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
        /// </summary>
        [Output("stackSmApiUrl")]
        public Output<string> StackSmApiUrl { get; private set; } = null!;


        /// <summary>
        /// Create a SyntheticMonitoringInstallation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SyntheticMonitoringInstallation(string name, SyntheticMonitoringInstallationArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/syntheticMonitoringInstallation:SyntheticMonitoringInstallation", name, args ?? new SyntheticMonitoringInstallationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SyntheticMonitoringInstallation(string name, Input<string> id, SyntheticMonitoringInstallationState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/syntheticMonitoringInstallation:SyntheticMonitoringInstallation", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                AdditionalSecretOutputs =
                {
                    "metricsPublisherKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SyntheticMonitoringInstallation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SyntheticMonitoringInstallation Get(string name, Input<string> id, SyntheticMonitoringInstallationState? state = null, CustomResourceOptions? options = null)
        {
            return new SyntheticMonitoringInstallation(name, id, state, options);
        }
    }

    public sealed class SyntheticMonitoringInstallationArgs : global::Pulumi.ResourceArgs
    {
        [Input("metricsPublisherKey", required: true)]
        private Input<string>? _metricsPublisherKey;

        /// <summary>
        /// The [Grafana Cloud access policy](https://grafana.com/docs/grafana-cloud/security-and-account-management/authentication-and-permissions/access-policies/) with the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and logs to Grafana Cloud stack.
        /// </summary>
        public Input<string>? MetricsPublisherKey
        {
            get => _metricsPublisherKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _metricsPublisherKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The ID or slug of the stack to install SM on.
        /// </summary>
        [Input("stackId", required: true)]
        public Input<string> StackId { get; set; } = null!;

        /// <summary>
        /// The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/set-up-private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
        /// </summary>
        [Input("stackSmApiUrl")]
        public Input<string>? StackSmApiUrl { get; set; }

        public SyntheticMonitoringInstallationArgs()
        {
        }
        public static new SyntheticMonitoringInstallationArgs Empty => new SyntheticMonitoringInstallationArgs();
    }

    public sealed class SyntheticMonitoringInstallationState : global::Pulumi.ResourceArgs
    {
        [Input("metricsPublisherKey")]
        private Input<string>? _metricsPublisherKey;

        /// <summary>
        /// The [Grafana Cloud access policy](https://grafana.com/docs/grafana-cloud/security-and-account-management/authentication-and-permissions/access-policies/) with the following scopes: `stacks:read`, `metrics:write`, `logs:write`, `traces:write`. This is used to publish metrics and logs to Grafana Cloud stack.
        /// </summary>
        public Input<string>? MetricsPublisherKey
        {
            get => _metricsPublisherKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _metricsPublisherKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Generated token to access the SM API.
        /// </summary>
        [Input("smAccessToken")]
        public Input<string>? SmAccessToken { get; set; }

        /// <summary>
        /// The ID or slug of the stack to install SM on.
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        /// <summary>
        /// The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/set-up/set-up-private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
        /// </summary>
        [Input("stackSmApiUrl")]
        public Input<string>? StackSmApiUrl { get; set; }

        public SyntheticMonitoringInstallationState()
        {
        }
        public static new SyntheticMonitoringInstallationState Empty => new SyntheticMonitoringInstallationState();
    }
}
