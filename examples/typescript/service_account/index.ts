import * as pulumi from "@pulumi/pulumi";
import * as grafana from "@lbrlabs/pulumi-grafana";

const sa = new grafana.ServiceAccount("example", {
    role: "Viewer",
})