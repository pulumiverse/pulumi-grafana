// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.OnCall.Outputs
{

    [OutputType]
    public sealed class GetUsersUserResult
    {
        public readonly string Email;
        public readonly string Id;
        public readonly string Role;
        public readonly string Username;

        [OutputConstructor]
        private GetUsersUserResult(
            string email,

            string id,

            string role,

            string username)
        {
            Email = email;
            Id = id;
            Role = role;
            Username = username;
        }
    }
}