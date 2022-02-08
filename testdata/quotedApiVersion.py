import pulumi
import pulumi_kubernetes as kubernetes

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",
    api_version="apps/v1",
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        labels={
            "app.kubernetes.io/component": "server",
            "aws:region": "us-west-2",
        },
        name="argocd-server",
    ))
