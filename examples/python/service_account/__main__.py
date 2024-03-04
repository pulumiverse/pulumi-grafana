"""A Python Pulumi program"""

import pulumi
import pulumiverse_grafana as grafana

service_account = grafana.ServiceAccount(
    "example",
    role="Viewer"
)
