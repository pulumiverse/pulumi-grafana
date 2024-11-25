// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/datasources/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/data_source/)
 *
 * The required arguments for this resource vary depending on the type of data
 * source selected (via the 'type' argument).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const arbitrary_data = new grafana.oss.DataSource("arbitrary-data", {
 *     type: "stackdriver",
 *     name: "sd-arbitrary-data",
 *     jsonDataEncoded: JSON.stringify({
 *         tokenUri: "https://oauth2.googleapis.com/token",
 *         authenticationType: "jwt",
 *         defaultProject: "default-project",
 *         clientEmail: "client-email@default-project.iam.gserviceaccount.com",
 *     }),
 *     secureJsonDataEncoded: JSON.stringify({
 *         privateKey: `-----BEGIN PRIVATE KEY-----
 * private-key
 * -----END PRIVATE KEY-----
 * `,
 *     }),
 * });
 * const influxdb = new grafana.oss.DataSource("influxdb", {
 *     type: "influxdb",
 *     name: "myapp-metrics",
 *     url: "http://influxdb.example.net:8086/",
 *     basicAuthEnabled: true,
 *     basicAuthUsername: "username",
 *     databaseName: "dbname",
 *     jsonDataEncoded: JSON.stringify({
 *         authType: "default",
 *         basicAuthPassword: "mypassword",
 *     }),
 * });
 * const cloudwatch = new grafana.oss.DataSource("cloudwatch", {
 *     type: "cloudwatch",
 *     name: "cw-example",
 *     jsonDataEncoded: JSON.stringify({
 *         defaultRegion: "us-east-1",
 *         authType: "keys",
 *     }),
 *     secureJsonDataEncoded: JSON.stringify({
 *         accessKey: "123",
 *         secretKey: "456",
 *     }),
 * });
 * const cloudwatchAssumeARN = new grafana.oss.DataSource("cloudwatch_assumeARN", {
 *     type: "cloudwatch",
 *     name: "cw-assumeARN-example",
 *     jsonDataEncoded: JSON.stringify({
 *         defaultRegion: "us-east-1",
 *         authType: "grafana_assume_role",
 *         assumeRoleArn: "arn:aws:iam::123456789012:root",
 *     }),
 * });
 * const prometheus = new grafana.oss.DataSource("prometheus", {
 *     type: "prometheus",
 *     name: "mimir",
 *     url: "https://my-instances.com",
 *     basicAuthEnabled: true,
 *     basicAuthUsername: "username",
 *     jsonDataEncoded: JSON.stringify({
 *         httpMethod: "POST",
 *         prometheusType: "Mimir",
 *         prometheusVersion: "2.4.0",
 *     }),
 *     secureJsonDataEncoded: JSON.stringify({
 *         basicAuthPassword: "password",
 *     }),
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:oss/dataSource:DataSource name "{{ uid }}"
 * ```
 *
 * ```sh
 * $ pulumi import grafana:oss/dataSource:DataSource name "{{ orgID }}:{{ uid }}"
 * ```
 */
export class DataSource extends pulumi.CustomResource {
    /**
     * Get an existing DataSource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataSourceState, opts?: pulumi.CustomResourceOptions): DataSource {
        return new DataSource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:oss/dataSource:DataSource';

    /**
     * Returns true if the given object is an instance of DataSource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataSource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataSource.__pulumiType;
    }

    /**
     * The method by which Grafana will access the data source: `proxy` or `direct`. Defaults to `proxy`.
     */
    public readonly accessMode!: pulumi.Output<string | undefined>;
    /**
     * Whether to enable basic auth for the data source. Defaults to `false`.
     */
    public readonly basicAuthEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Basic auth username. Defaults to ``.
     */
    public readonly basicAuthUsername!: pulumi.Output<string | undefined>;
    /**
     * (Required by some data source types) The name of the database to use on the selected data source server. Defaults to ``.
     */
    public readonly databaseName!: pulumi.Output<string | undefined>;
    /**
     * Custom HTTP headers
     */
    public readonly httpHeaders!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
     */
    public readonly isDefault!: pulumi.Output<boolean | undefined>;
    /**
     * Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
     */
    public readonly jsonDataEncoded!: pulumi.Output<string | undefined>;
    /**
     * A unique name for the data source.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
     */
    public readonly secureJsonDataEncoded!: pulumi.Output<string | undefined>;
    /**
     * The data source type. Must be one of the supported data source keywords.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Unique identifier. If unset, this will be automatically generated.
     */
    public readonly uid!: pulumi.Output<string>;
    /**
     * The URL for the data source. The type of URL required varies depending on the chosen data source type.
     */
    public readonly url!: pulumi.Output<string | undefined>;
    /**
     * (Required by some data source types) The username to use to authenticate to the data source. Defaults to ``.
     */
    public readonly username!: pulumi.Output<string | undefined>;

    /**
     * Create a DataSource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataSourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataSourceArgs | DataSourceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DataSourceState | undefined;
            resourceInputs["accessMode"] = state ? state.accessMode : undefined;
            resourceInputs["basicAuthEnabled"] = state ? state.basicAuthEnabled : undefined;
            resourceInputs["basicAuthUsername"] = state ? state.basicAuthUsername : undefined;
            resourceInputs["databaseName"] = state ? state.databaseName : undefined;
            resourceInputs["httpHeaders"] = state ? state.httpHeaders : undefined;
            resourceInputs["isDefault"] = state ? state.isDefault : undefined;
            resourceInputs["jsonDataEncoded"] = state ? state.jsonDataEncoded : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["secureJsonDataEncoded"] = state ? state.secureJsonDataEncoded : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as DataSourceArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["accessMode"] = args ? args.accessMode : undefined;
            resourceInputs["basicAuthEnabled"] = args ? args.basicAuthEnabled : undefined;
            resourceInputs["basicAuthUsername"] = args ? args.basicAuthUsername : undefined;
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["httpHeaders"] = args?.httpHeaders ? pulumi.secret(args.httpHeaders) : undefined;
            resourceInputs["isDefault"] = args ? args.isDefault : undefined;
            resourceInputs["jsonDataEncoded"] = args ? args.jsonDataEncoded : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["secureJsonDataEncoded"] = args?.secureJsonDataEncoded ? pulumi.secret(args.secureJsonDataEncoded) : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["uid"] = args ? args.uid : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "grafana:index/dataSource:DataSource" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        const secretOpts = { additionalSecretOutputs: ["httpHeaders", "secureJsonDataEncoded"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(DataSource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataSource resources.
 */
export interface DataSourceState {
    /**
     * The method by which Grafana will access the data source: `proxy` or `direct`. Defaults to `proxy`.
     */
    accessMode?: pulumi.Input<string>;
    /**
     * Whether to enable basic auth for the data source. Defaults to `false`.
     */
    basicAuthEnabled?: pulumi.Input<boolean>;
    /**
     * Basic auth username. Defaults to ``.
     */
    basicAuthUsername?: pulumi.Input<string>;
    /**
     * (Required by some data source types) The name of the database to use on the selected data source server. Defaults to ``.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * Custom HTTP headers
     */
    httpHeaders?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
     */
    jsonDataEncoded?: pulumi.Input<string>;
    /**
     * A unique name for the data source.
     */
    name?: pulumi.Input<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: pulumi.Input<string>;
    /**
     * Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
     */
    secureJsonDataEncoded?: pulumi.Input<string>;
    /**
     * The data source type. Must be one of the supported data source keywords.
     */
    type?: pulumi.Input<string>;
    /**
     * Unique identifier. If unset, this will be automatically generated.
     */
    uid?: pulumi.Input<string>;
    /**
     * The URL for the data source. The type of URL required varies depending on the chosen data source type.
     */
    url?: pulumi.Input<string>;
    /**
     * (Required by some data source types) The username to use to authenticate to the data source. Defaults to ``.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DataSource resource.
 */
export interface DataSourceArgs {
    /**
     * The method by which Grafana will access the data source: `proxy` or `direct`. Defaults to `proxy`.
     */
    accessMode?: pulumi.Input<string>;
    /**
     * Whether to enable basic auth for the data source. Defaults to `false`.
     */
    basicAuthEnabled?: pulumi.Input<boolean>;
    /**
     * Basic auth username. Defaults to ``.
     */
    basicAuthUsername?: pulumi.Input<string>;
    /**
     * (Required by some data source types) The name of the database to use on the selected data source server. Defaults to ``.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * Custom HTTP headers
     */
    httpHeaders?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
     */
    jsonDataEncoded?: pulumi.Input<string>;
    /**
     * A unique name for the data source.
     */
    name?: pulumi.Input<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: pulumi.Input<string>;
    /**
     * Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
     */
    secureJsonDataEncoded?: pulumi.Input<string>;
    /**
     * The data source type. Must be one of the supported data source keywords.
     */
    type: pulumi.Input<string>;
    /**
     * Unique identifier. If unset, this will be automatically generated.
     */
    uid?: pulumi.Input<string>;
    /**
     * The URL for the data source. The type of URL required varies depending on the chosen data source type.
     */
    url?: pulumi.Input<string>;
    /**
     * (Required by some data source types) The username to use to authenticate to the data source. Defaults to ``.
     */
    username?: pulumi.Input<string>;
}
