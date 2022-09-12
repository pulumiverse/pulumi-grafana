"""A Python Pulumi program"""

import pulumi
import lbrlabs_pulumi_grafana as grafana

service_account = grafana.ServiceAccount(
    "example",
    role="Viewer"
)
