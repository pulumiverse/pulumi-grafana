// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.SyntheticMonitoring.Inputs
{

    public sealed class CheckSettingsMultihttpEntryGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("assertions")]
        private InputList<Inputs.CheckSettingsMultihttpEntryAssertionGetArgs>? _assertions;

        /// <summary>
        /// Assertions to make on the request response
        /// </summary>
        public InputList<Inputs.CheckSettingsMultihttpEntryAssertionGetArgs> Assertions
        {
            get => _assertions ?? (_assertions = new InputList<Inputs.CheckSettingsMultihttpEntryAssertionGetArgs>());
            set => _assertions = value;
        }

        /// <summary>
        /// An individual MultiHTTP request
        /// </summary>
        [Input("request")]
        public Input<Inputs.CheckSettingsMultihttpEntryRequestGetArgs>? Request { get; set; }

        [Input("variables")]
        private InputList<Inputs.CheckSettingsMultihttpEntryVariableGetArgs>? _variables;

        /// <summary>
        /// Variables to extract from the request response
        /// </summary>
        public InputList<Inputs.CheckSettingsMultihttpEntryVariableGetArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.CheckSettingsMultihttpEntryVariableGetArgs>());
            set => _variables = value;
        }

        public CheckSettingsMultihttpEntryGetArgs()
        {
        }
        public static new CheckSettingsMultihttpEntryGetArgs Empty => new CheckSettingsMultihttpEntryGetArgs();
    }
}
