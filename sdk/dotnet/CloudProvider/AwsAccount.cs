// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.CloudProvider
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// using Grafana = Pulumi.Grafana;
    /// using Grafana = Pulumiverse.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = Grafana.Cloud.GetStack.Invoke(new()
    ///     {
    ///         Slug = "gcloudstacktest",
    ///     });
    /// 
    ///     var testGetRole = Aws.Iam.GetRole.Invoke(new()
    ///     {
    ///         Name = "my-role",
    ///     });
    /// 
    ///     var testAwsAccount = new Grafana.CloudProvider.AwsAccount("test", new()
    ///     {
    ///         StackId = test.Apply(getStackResult =&gt; getStackResult.Id),
    ///         RoleArn = testGetRole.Apply(getRoleResult =&gt; getRoleResult.Arn),
    ///         Regions = new[]
    ///         {
    ///             "us-east-1",
    ///             "us-east-2",
    ///             "us-west-1",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:cloudProvider/awsAccount:AwsAccount name "{{ stack_id }}:{{ resource_id }}"
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:cloudProvider/awsAccount:AwsAccount")]
    public partial class AwsAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An optional human-readable name for this AWS Account resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A set of regions that this AWS Account resource applies to.
        /// </summary>
        [Output("regions")]
        public Output<ImmutableArray<string>> Regions { get; private set; } = null!;

        /// <summary>
        /// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// An IAM Role ARN string to represent with this AWS Account resource.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        [Output("stackId")]
        public Output<string> StackId { get; private set; } = null!;


        /// <summary>
        /// Create a AwsAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AwsAccount(string name, AwsAccountArgs args, CustomResourceOptions? options = null)
            : base("grafana:cloudProvider/awsAccount:AwsAccount", name, args ?? new AwsAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AwsAccount(string name, Input<string> id, AwsAccountState? state = null, CustomResourceOptions? options = null)
            : base("grafana:cloudProvider/awsAccount:AwsAccount", name, state, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "grafana:cloud/providerAwsAccount:ProviderAwsAccount" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AwsAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AwsAccount Get(string name, Input<string> id, AwsAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new AwsAccount(name, id, state, options);
        }
    }

    public sealed class AwsAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional human-readable name for this AWS Account resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("regions", required: true)]
        private InputList<string>? _regions;

        /// <summary>
        /// A set of regions that this AWS Account resource applies to.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        /// <summary>
        /// An IAM Role ARN string to represent with this AWS Account resource.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("stackId", required: true)]
        public Input<string> StackId { get; set; } = null!;

        public AwsAccountArgs()
        {
        }
        public static new AwsAccountArgs Empty => new AwsAccountArgs();
    }

    public sealed class AwsAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional human-readable name for this AWS Account resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("regions")]
        private InputList<string>? _regions;

        /// <summary>
        /// A set of regions that this AWS Account resource applies to.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        /// <summary>
        /// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// An IAM Role ARN string to represent with this AWS Account resource.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        public AwsAccountState()
        {
        }
        public static new AwsAccountState Empty => new AwsAccountState();
    }
}
