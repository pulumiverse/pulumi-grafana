// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the grafana package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'grafana';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * API token, basic auth in the `username:password` format or `anonymous` (string literal). May alternatively be set via
     * the `GRAFANA_AUTH` environment variable.
     */
    public readonly auth!: pulumi.Output<string | undefined>;
    /**
     * Certificate CA bundle (file path or literal value) to use to verify the Grafana server's certificate. May alternatively
     * be set via the `GRAFANA_CA_CERT` environment variable.
     */
    public readonly caCert!: pulumi.Output<string | undefined>;
    /**
     * Access Policy Token for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_ACCESS_POLICY_TOKEN` environment
     * variable.
     */
    public readonly cloudAccessPolicyToken!: pulumi.Output<string | undefined>;
    /**
     * Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
     */
    public readonly cloudApiUrl!: pulumi.Output<string | undefined>;
    /**
     * A Grafana Cloud Provider access token. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_ACCESS_TOKEN`
     * environment variable.
     */
    public readonly cloudProviderAccessToken!: pulumi.Output<string | undefined>;
    /**
     * A Grafana Cloud Provider backend address. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_URL` environment
     * variable.
     */
    public readonly cloudProviderUrl!: pulumi.Output<string | undefined>;
    /**
     * A Grafana Connections API access token. May alternatively be set via the `GRAFANA_CONNECTIONS_API_ACCESS_TOKEN`
     * environment variable.
     */
    public readonly connectionsApiAccessToken!: pulumi.Output<string | undefined>;
    /**
     * A Grafana Connections API address. May alternatively be set via the `GRAFANA_CONNECTIONS_API_URL` environment variable.
     */
    public readonly connectionsApiUrl!: pulumi.Output<string | undefined>;
    /**
     * A Grafana Fleet Management basic auth in the `username:password` format. May alternatively be set via the
     * `GRAFANA_FLEET_MANAGEMENT_AUTH` environment variable.
     */
    public readonly fleetManagementAuth!: pulumi.Output<string | undefined>;
    /**
     * A Grafana Fleet Management API address. May alternatively be set via the `GRAFANA_FLEET_MANAGEMENT_URL` environment
     * variable.
     */
    public readonly fleetManagementUrl!: pulumi.Output<string | undefined>;
    /**
     * A Grafana Frontend Observability API access token. May alternatively be set via the
     * `GRAFANA_FRONTEND_O11Y_API_ACCESS_TOKEN` environment variable.
     */
    public readonly frontendO11yApiAccessToken!: pulumi.Output<string | undefined>;
    /**
     * A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
     */
    public readonly oncallAccessToken!: pulumi.Output<string | undefined>;
    /**
     * An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
     */
    public readonly oncallUrl!: pulumi.Output<string | undefined>;
    /**
     * A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
     */
    public readonly smAccessToken!: pulumi.Output<string | undefined>;
    public readonly smUrl!: pulumi.Output<string | undefined>;
    /**
     * Client TLS certificate (file path or literal value) to use to authenticate to the Grafana server. May alternatively be
     * set via the `GRAFANA_TLS_CERT` environment variable.
     */
    public readonly tlsCert!: pulumi.Output<string | undefined>;
    /**
     * Client TLS key (file path or literal value) to use to authenticate to the Grafana server. May alternatively be set via
     * the `GRAFANA_TLS_KEY` environment variable.
     */
    public readonly tlsKey!: pulumi.Output<string | undefined>;
    /**
     * The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
     */
    public readonly url!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["auth"] = (args?.auth ? pulumi.secret(args.auth) : undefined) ?? utilities.getEnv("GRAFANA_AUTH");
            resourceInputs["caCert"] = (args ? args.caCert : undefined) ?? utilities.getEnv("GRAFANA_CA_CERT");
            resourceInputs["cloudAccessPolicyToken"] = (args?.cloudAccessPolicyToken ? pulumi.secret(args.cloudAccessPolicyToken) : undefined) ?? utilities.getEnv("GRAFANA_CLOUD_ACCESS_POLICY_TOKEN");
            resourceInputs["cloudApiUrl"] = (args ? args.cloudApiUrl : undefined) ?? utilities.getEnv("GRAFANA_CLOUD_API_URL");
            resourceInputs["cloudProviderAccessToken"] = args?.cloudProviderAccessToken ? pulumi.secret(args.cloudProviderAccessToken) : undefined;
            resourceInputs["cloudProviderUrl"] = args ? args.cloudProviderUrl : undefined;
            resourceInputs["connectionsApiAccessToken"] = args?.connectionsApiAccessToken ? pulumi.secret(args.connectionsApiAccessToken) : undefined;
            resourceInputs["connectionsApiUrl"] = args ? args.connectionsApiUrl : undefined;
            resourceInputs["fleetManagementAuth"] = args?.fleetManagementAuth ? pulumi.secret(args.fleetManagementAuth) : undefined;
            resourceInputs["fleetManagementUrl"] = args ? args.fleetManagementUrl : undefined;
            resourceInputs["frontendO11yApiAccessToken"] = args?.frontendO11yApiAccessToken ? pulumi.secret(args.frontendO11yApiAccessToken) : undefined;
            resourceInputs["httpHeaders"] = pulumi.output(args?.httpHeaders ? pulumi.secret(args.httpHeaders) : undefined).apply(JSON.stringify);
            resourceInputs["insecureSkipVerify"] = pulumi.output((args ? args.insecureSkipVerify : undefined) ?? utilities.getEnvBoolean("GRAFANA_INSECURE_SKIP_VERIFY")).apply(JSON.stringify);
            resourceInputs["oncallAccessToken"] = (args?.oncallAccessToken ? pulumi.secret(args.oncallAccessToken) : undefined) ?? utilities.getEnv("GRAFANA_ONCALL_ACCESS_TOKEN");
            resourceInputs["oncallUrl"] = (args ? args.oncallUrl : undefined) ?? utilities.getEnv("GRAFANA_ONCALL_URL");
            resourceInputs["retries"] = pulumi.output((args ? args.retries : undefined) ?? utilities.getEnvNumber("GRAFANA_RETRIES")).apply(JSON.stringify);
            resourceInputs["retryStatusCodes"] = pulumi.output(args ? args.retryStatusCodes : undefined).apply(JSON.stringify);
            resourceInputs["retryWait"] = pulumi.output((args ? args.retryWait : undefined) ?? utilities.getEnvNumber("GRAFANA_RETRY_WAIT")).apply(JSON.stringify);
            resourceInputs["smAccessToken"] = (args?.smAccessToken ? pulumi.secret(args.smAccessToken) : undefined) ?? utilities.getEnv("GRAFANA_SM_ACCESS_TOKEN");
            resourceInputs["smUrl"] = (args ? args.smUrl : undefined) ?? utilities.getEnv("GRAFANA_SM_URL");
            resourceInputs["storeDashboardSha256"] = pulumi.output((args ? args.storeDashboardSha256 : undefined) ?? utilities.getEnvBoolean("GRAFANA_STORE_DASHBOARD_SHA256")).apply(JSON.stringify);
            resourceInputs["tlsCert"] = (args ? args.tlsCert : undefined) ?? utilities.getEnv("GRAFANA_TLS_CERT");
            resourceInputs["tlsKey"] = (args?.tlsKey ? pulumi.secret(args.tlsKey) : undefined) ?? utilities.getEnv("GRAFANA_TLS_KEY");
            resourceInputs["url"] = (args ? args.url : undefined) ?? utilities.getEnv("GRAFANA_URL");
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["auth", "cloudAccessPolicyToken", "cloudProviderAccessToken", "connectionsApiAccessToken", "fleetManagementAuth", "frontendO11yApiAccessToken", "oncallAccessToken", "smAccessToken", "tlsKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * API token, basic auth in the `username:password` format or `anonymous` (string literal). May alternatively be set via
     * the `GRAFANA_AUTH` environment variable.
     */
    auth?: pulumi.Input<string>;
    /**
     * Certificate CA bundle (file path or literal value) to use to verify the Grafana server's certificate. May alternatively
     * be set via the `GRAFANA_CA_CERT` environment variable.
     */
    caCert?: pulumi.Input<string>;
    /**
     * Access Policy Token for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_ACCESS_POLICY_TOKEN` environment
     * variable.
     */
    cloudAccessPolicyToken?: pulumi.Input<string>;
    /**
     * Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
     */
    cloudApiUrl?: pulumi.Input<string>;
    /**
     * A Grafana Cloud Provider access token. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_ACCESS_TOKEN`
     * environment variable.
     */
    cloudProviderAccessToken?: pulumi.Input<string>;
    /**
     * A Grafana Cloud Provider backend address. May alternatively be set via the `GRAFANA_CLOUD_PROVIDER_URL` environment
     * variable.
     */
    cloudProviderUrl?: pulumi.Input<string>;
    /**
     * A Grafana Connections API access token. May alternatively be set via the `GRAFANA_CONNECTIONS_API_ACCESS_TOKEN`
     * environment variable.
     */
    connectionsApiAccessToken?: pulumi.Input<string>;
    /**
     * A Grafana Connections API address. May alternatively be set via the `GRAFANA_CONNECTIONS_API_URL` environment variable.
     */
    connectionsApiUrl?: pulumi.Input<string>;
    /**
     * A Grafana Fleet Management basic auth in the `username:password` format. May alternatively be set via the
     * `GRAFANA_FLEET_MANAGEMENT_AUTH` environment variable.
     */
    fleetManagementAuth?: pulumi.Input<string>;
    /**
     * A Grafana Fleet Management API address. May alternatively be set via the `GRAFANA_FLEET_MANAGEMENT_URL` environment
     * variable.
     */
    fleetManagementUrl?: pulumi.Input<string>;
    /**
     * A Grafana Frontend Observability API access token. May alternatively be set via the
     * `GRAFANA_FRONTEND_O11Y_API_ACCESS_TOKEN` environment variable.
     */
    frontendO11yApiAccessToken?: pulumi.Input<string>;
    /**
     * Optional. HTTP headers mapping keys to values used for accessing the Grafana and Grafana Cloud APIs. May alternatively
     * be set via the `GRAFANA_HTTP_HEADERS` environment variable in JSON format.
     */
    httpHeaders?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
     */
    insecureSkipVerify?: pulumi.Input<boolean>;
    /**
     * A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
     */
    oncallAccessToken?: pulumi.Input<string>;
    /**
     * An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
     */
    oncallUrl?: pulumi.Input<string>;
    /**
     * The amount of retries to use for Grafana API and Grafana Cloud API calls. May alternatively be set via the
     * `GRAFANA_RETRIES` environment variable.
     */
    retries?: pulumi.Input<number>;
    /**
     * The status codes to retry on for Grafana API and Grafana Cloud API calls. Use `x` as a digit wildcard. Defaults to 429
     * and 5xx. May alternatively be set via the `GRAFANA_RETRY_STATUS_CODES` environment variable.
     */
    retryStatusCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The amount of time in seconds to wait between retries for Grafana API and Grafana Cloud API calls. May alternatively be
     * set via the `GRAFANA_RETRY_WAIT` environment variable.
     */
    retryWait?: pulumi.Input<number>;
    /**
     * A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
     */
    smAccessToken?: pulumi.Input<string>;
    smUrl?: pulumi.Input<string>;
    /**
     * Set to true if you want to save only the sha256sum instead of complete dashboard model JSON in the tfstate.
     */
    storeDashboardSha256?: pulumi.Input<boolean>;
    /**
     * Client TLS certificate (file path or literal value) to use to authenticate to the Grafana server. May alternatively be
     * set via the `GRAFANA_TLS_CERT` environment variable.
     */
    tlsCert?: pulumi.Input<string>;
    /**
     * Client TLS key (file path or literal value) to use to authenticate to the Grafana server. May alternatively be set via
     * the `GRAFANA_TLS_KEY` environment variable.
     */
    tlsKey?: pulumi.Input<string>;
    /**
     * The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
     */
    url?: pulumi.Input<string>;
}
