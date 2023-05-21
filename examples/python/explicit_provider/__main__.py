"""A Python Pulumi program"""

import pulumi
import pulumi_random as random
import lbrlabs_pulumi_grafana as grafana


service_account = grafana.ServiceAccount(
    "ci-testing",
    role="Admin"
)

service_account_token = grafana.ServiceAccountToken(
    "ci-testing",
    service_account_id = service_account.id,
    seconds_to_live = 600,
)

provider = grafana.Provider(
    "ci-testing",
    auth=service_account_token.key,
)

folder_name = random.RandomString(
    "example",
    length=10,
)

folder = grafana.Folder(
    "example",
    title=folder_name.result,
)

