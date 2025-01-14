// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.SyntheticMonitoring.Inputs
{

    public sealed class CheckSettingsHttpGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Basic auth settings.
        /// </summary>
        [Input("basicAuth")]
        public Input<Inputs.CheckSettingsHttpBasicAuthGetArgs>? BasicAuth { get; set; }

        [Input("bearerToken")]
        private Input<string>? _bearerToken;

        /// <summary>
        /// Token for use with bearer authorization header.
        /// </summary>
        public Input<string>? BearerToken
        {
            get => _bearerToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bearerToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The body of the HTTP request used in probe.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        /// <summary>
        /// The name of the query parameter used to prevent the server from using a cached response. Each probe will assign a random value to this parameter each time a request is made.
        /// </summary>
        [Input("cacheBustingQueryParamName")]
        public Input<string>? CacheBustingQueryParamName { get; set; }

        /// <summary>
        /// Check fails if the response body is not compressed using this compression algorithm. One of `none`, `identity`, `br`, `gzip`, `deflate`.
        /// </summary>
        [Input("compression")]
        public Input<string>? Compression { get; set; }

        [Input("failIfBodyMatchesRegexps")]
        private InputList<string>? _failIfBodyMatchesRegexps;

        /// <summary>
        /// List of regexes. If any match the response body, the check will fail.
        /// </summary>
        public InputList<string> FailIfBodyMatchesRegexps
        {
            get => _failIfBodyMatchesRegexps ?? (_failIfBodyMatchesRegexps = new InputList<string>());
            set => _failIfBodyMatchesRegexps = value;
        }

        [Input("failIfBodyNotMatchesRegexps")]
        private InputList<string>? _failIfBodyNotMatchesRegexps;

        /// <summary>
        /// List of regexes. If any do not match the response body, the check will fail.
        /// </summary>
        public InputList<string> FailIfBodyNotMatchesRegexps
        {
            get => _failIfBodyNotMatchesRegexps ?? (_failIfBodyNotMatchesRegexps = new InputList<string>());
            set => _failIfBodyNotMatchesRegexps = value;
        }

        [Input("failIfHeaderMatchesRegexps")]
        private InputList<Inputs.CheckSettingsHttpFailIfHeaderMatchesRegexpGetArgs>? _failIfHeaderMatchesRegexps;

        /// <summary>
        /// Check fails if headers match.
        /// </summary>
        public InputList<Inputs.CheckSettingsHttpFailIfHeaderMatchesRegexpGetArgs> FailIfHeaderMatchesRegexps
        {
            get => _failIfHeaderMatchesRegexps ?? (_failIfHeaderMatchesRegexps = new InputList<Inputs.CheckSettingsHttpFailIfHeaderMatchesRegexpGetArgs>());
            set => _failIfHeaderMatchesRegexps = value;
        }

        [Input("failIfHeaderNotMatchesRegexps")]
        private InputList<Inputs.CheckSettingsHttpFailIfHeaderNotMatchesRegexpGetArgs>? _failIfHeaderNotMatchesRegexps;

        /// <summary>
        /// Check fails if headers do not match.
        /// </summary>
        public InputList<Inputs.CheckSettingsHttpFailIfHeaderNotMatchesRegexpGetArgs> FailIfHeaderNotMatchesRegexps
        {
            get => _failIfHeaderNotMatchesRegexps ?? (_failIfHeaderNotMatchesRegexps = new InputList<Inputs.CheckSettingsHttpFailIfHeaderNotMatchesRegexpGetArgs>());
            set => _failIfHeaderNotMatchesRegexps = value;
        }

        /// <summary>
        /// Fail if SSL is not present. Defaults to `false`.
        /// </summary>
        [Input("failIfNotSsl")]
        public Input<bool>? FailIfNotSsl { get; set; }

        /// <summary>
        /// Fail if SSL is present. Defaults to `false`.
        /// </summary>
        [Input("failIfSsl")]
        public Input<bool>? FailIfSsl { get; set; }

        [Input("headers")]
        private InputList<string>? _headers;

        /// <summary>
        /// The HTTP headers set for the probe.
        /// </summary>
        public InputList<string> Headers
        {
            get => _headers ?? (_headers = new InputList<string>());
            set => _headers = value;
        }

        /// <summary>
        /// Options are `V4`, `V6`, `Any`. Specifies whether the corresponding check will be performed using IPv4 or IPv6. The `Any` value indicates that IPv6 should be used, falling back to IPv4 if that's not available. Defaults to `V4`.
        /// </summary>
        [Input("ipVersion")]
        public Input<string>? IpVersion { get; set; }

        /// <summary>
        /// Request method. One of `GET`, `CONNECT`, `DELETE`, `HEAD`, `OPTIONS`, `POST`, `PUT`, `TRACE` Defaults to `GET`.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// Do not follow redirects. Defaults to `false`.
        /// </summary>
        [Input("noFollowRedirects")]
        public Input<bool>? NoFollowRedirects { get; set; }

        [Input("proxyConnectHeaders")]
        private InputList<string>? _proxyConnectHeaders;

        /// <summary>
        /// The HTTP headers sent to the proxy URL
        /// </summary>
        public InputList<string> ProxyConnectHeaders
        {
            get => _proxyConnectHeaders ?? (_proxyConnectHeaders = new InputList<string>());
            set => _proxyConnectHeaders = value;
        }

        /// <summary>
        /// Proxy URL.
        /// </summary>
        [Input("proxyUrl")]
        public Input<string>? ProxyUrl { get; set; }

        /// <summary>
        /// TLS config.
        /// </summary>
        [Input("tlsConfig")]
        public Input<Inputs.CheckSettingsHttpTlsConfigGetArgs>? TlsConfig { get; set; }

        [Input("validHttpVersions")]
        private InputList<string>? _validHttpVersions;

        /// <summary>
        /// List of valid HTTP versions. Options include `HTTP/1.0`, `HTTP/1.1`, `HTTP/2.0`
        /// </summary>
        public InputList<string> ValidHttpVersions
        {
            get => _validHttpVersions ?? (_validHttpVersions = new InputList<string>());
            set => _validHttpVersions = value;
        }

        [Input("validStatusCodes")]
        private InputList<int>? _validStatusCodes;

        /// <summary>
        /// Accepted status codes. If unset, defaults to 2xx.
        /// </summary>
        public InputList<int> ValidStatusCodes
        {
            get => _validStatusCodes ?? (_validStatusCodes = new InputList<int>());
            set => _validStatusCodes = value;
        }

        public CheckSettingsHttpGetArgs()
        {
        }
        public static new CheckSettingsHttpGetArgs Empty => new CheckSettingsHttpGetArgs();
    }
}