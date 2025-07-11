// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.SyntheticMonitoring.Outputs
{

    [OutputType]
    public sealed class CheckSettingsMultihttpEntryAssertion
    {
        /// <summary>
        /// The condition of the assertion: NOT*CONTAINS, EQUALS, STARTS*WITH, ENDS*WITH, TYPE*OF, CONTAINS
        /// </summary>
        public readonly string? Condition;
        /// <summary>
        /// The expression of the assertion. Should start with $.
        /// </summary>
        public readonly string? Expression;
        /// <summary>
        /// The subject of the assertion: RESPONSE*HEADERS, HTTP*STATUS*CODE, RESPONSE*BODY
        /// </summary>
        public readonly string? Subject;
        /// <summary>
        /// The type of assertion to make: TEXT, JSON*PATH*VALUE, JSON*PATH*ASSERTION, REGEX_ASSERTION
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The value of the assertion
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private CheckSettingsMultihttpEntryAssertion(
            string? condition,

            string? expression,

            string? subject,

            string type,

            string? value)
        {
            Condition = condition;
            Expression = expression;
            Subject = subject;
            Type = type;
            Value = value;
        }
    }
}
