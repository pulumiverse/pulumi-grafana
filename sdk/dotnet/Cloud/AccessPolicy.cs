// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Cloud
{
    /// <summary>
    /// * [Official documentation](https://grafana.com/docs/grafana-cloud/security-and-account-management/authentication-and-permissions/access-policies/)
    /// * [API documentation](https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#create-an-access-policy)
    /// 
    /// Required access policy scopes:
    /// 
    /// * accesspolicies:read
    /// * accesspolicies:write
    /// * accesspolicies:delete
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
    ///     var current = Grafana.Cloud.GetOrganization.Invoke(new()
    ///     {
    ///         Slug = "&lt;your org slug&gt;",
    ///     });
    /// 
    ///     var test = new Grafana.Cloud.AccessPolicy("test", new()
    ///     {
    ///         Region = "prod-us-east-0",
    ///         Name = "my-policy",
    ///         DisplayName = "My Policy",
    ///         Scopes = new[]
    ///         {
    ///             "metrics:read",
    ///             "logs:read",
    ///         },
    ///         Realms = new[]
    ///         {
    ///             new Grafana.Cloud.Inputs.AccessPolicyRealmArgs
    ///             {
    ///                 Type = "org",
    ///                 Identifier = current.Apply(getOrganizationResult =&gt; getOrganizationResult.Id),
    ///                 LabelPolicies = new[]
    ///                 {
    ///                     new Grafana.Cloud.Inputs.AccessPolicyRealmLabelPolicyArgs
    ///                     {
    ///                         Selector = "{namespace=\"default\"}",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var testAccessPolicyToken = new Grafana.Cloud.AccessPolicyToken("test", new()
    ///     {
    ///         Region = "prod-us-east-0",
    ///         AccessPolicyId = test.PolicyId,
    ///         Name = "my-policy-token",
    ///         DisplayName = "My Policy Token",
    ///         ExpiresAt = "2023-01-01T00:00:00Z",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:cloud/accessPolicy:AccessPolicy name "{{ region }}:{{ policyId }}"
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:cloud/accessPolicy:AccessPolicy")]
    public partial class AccessPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Conditions for the access policy.
        /// </summary>
        [Output("conditions")]
        public Output<ImmutableArray<Outputs.AccessPolicyCondition>> Conditions { get; private set; } = null!;

        /// <summary>
        /// Creation date of the access policy.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Display name of the access policy. Defaults to the name.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Name of the access policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the access policy.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        [Output("realms")]
        public Output<ImmutableArray<Outputs.AccessPolicyRealm>> Realms { get; private set; } = null!;

        /// <summary>
        /// Region where the API is deployed. Generally where the stack is deployed. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Scopes of the access policy. See https://grafana.com/docs/grafana-cloud/security-and-account-management/authentication-and-permissions/access-policies/#scopes for possible values.
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        /// <summary>
        /// Last update date of the access policy.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a AccessPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessPolicy(string name, AccessPolicyArgs args, CustomResourceOptions? options = null)
            : base("grafana:cloud/accessPolicy:AccessPolicy", name, args ?? new AccessPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessPolicy(string name, Input<string> id, AccessPolicyState? state = null, CustomResourceOptions? options = null)
            : base("grafana:cloud/accessPolicy:AccessPolicy", name, state, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "grafana:index/cloudAccessPolicy:CloudAccessPolicy" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AccessPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessPolicy Get(string name, Input<string> id, AccessPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new AccessPolicy(name, id, state, options);
        }
    }

    public sealed class AccessPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("conditions")]
        private InputList<Inputs.AccessPolicyConditionArgs>? _conditions;

        /// <summary>
        /// Conditions for the access policy.
        /// </summary>
        public InputList<Inputs.AccessPolicyConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.AccessPolicyConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// Display name of the access policy. Defaults to the name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Name of the access policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("realms", required: true)]
        private InputList<Inputs.AccessPolicyRealmArgs>? _realms;
        public InputList<Inputs.AccessPolicyRealmArgs> Realms
        {
            get => _realms ?? (_realms = new InputList<Inputs.AccessPolicyRealmArgs>());
            set => _realms = value;
        }

        /// <summary>
        /// Region where the API is deployed. Generally where the stack is deployed. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("scopes", required: true)]
        private InputList<string>? _scopes;

        /// <summary>
        /// Scopes of the access policy. See https://grafana.com/docs/grafana-cloud/security-and-account-management/authentication-and-permissions/access-policies/#scopes for possible values.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        public AccessPolicyArgs()
        {
        }
        public static new AccessPolicyArgs Empty => new AccessPolicyArgs();
    }

    public sealed class AccessPolicyState : global::Pulumi.ResourceArgs
    {
        [Input("conditions")]
        private InputList<Inputs.AccessPolicyConditionGetArgs>? _conditions;

        /// <summary>
        /// Conditions for the access policy.
        /// </summary>
        public InputList<Inputs.AccessPolicyConditionGetArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.AccessPolicyConditionGetArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// Creation date of the access policy.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Display name of the access policy. Defaults to the name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Name of the access policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the access policy.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        [Input("realms")]
        private InputList<Inputs.AccessPolicyRealmGetArgs>? _realms;
        public InputList<Inputs.AccessPolicyRealmGetArgs> Realms
        {
            get => _realms ?? (_realms = new InputList<Inputs.AccessPolicyRealmGetArgs>());
            set => _realms = value;
        }

        /// <summary>
        /// Region where the API is deployed. Generally where the stack is deployed. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// Scopes of the access policy. See https://grafana.com/docs/grafana-cloud/security-and-account-management/authentication-and-permissions/access-policies/#scopes for possible values.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// Last update date of the access policy.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public AccessPolicyState()
        {
        }
        public static new AccessPolicyState Empty => new AccessPolicyState();
    }
}
