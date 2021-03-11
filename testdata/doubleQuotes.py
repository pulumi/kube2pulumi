import pulumi
import pulumi_kubernetes as kubernetes

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",
    api_version="apps/v1",
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        labels={
            "app.kubernetes.io/component": "server",
            "app.kubernetes.io/instance": "argocd",
            "app.kubernetes.io/managed-by": "pulumi",
            "app.kubernetes.io/name": "argocd-server",
            "app.kubernetes.io/part-of": "argocd",
            "app.kubernetes.io/version": "v1.6.1",
            "helm.sh/chart": "argo-cd-2.5.4",
        },
        name="argocd-server",
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(
        template=kubernetes.core.v1.PodTemplateSpecArgs(
            spec=kubernetes.core.v1.PodSpecArgs(
                containers=[kubernetes.core.v1.ContainerArgs(
                    readiness_probe={
                        "http_get": {
                            "port": 8080,
                        },
                    },
                )],
            ),
        ),
    ))
