import * as pulumi from "@pulumi/pulumi";
import * as grafana from "@pulumiverse/grafana";

const sa = new grafana.ServiceAccount("example", {
    role: "Viewer",
})