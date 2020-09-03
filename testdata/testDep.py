import pulumi
import pulumi_kubernetes as kubernetes

default_argocd_server_deployment = kubernetes.apps.v1.Deployment("defaultArgocd_serverDeployment",
    api_version="apps/v1",
    kind="Deployment",
    metadata={
        "annotations": {
            "deployment.kubernetes.io/revision": "1",
            "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"apps/v1\",\"kind\":\"Deployment\",\"metadata\":{\"labels\":{\"app.kubernetes.io/component\":\"server\",\"app.kubernetes.io/instance\":\"argocd\",\"app.kubernetes.io/managed-by\":\"pulumi\",\"app.kubernetes.io/name\":\"argocd-server\",\"app.kubernetes.io/part-of\":\"argocd\",\"app.kubernetes.io/version\":\"v1.6.1\",\"helm.sh/chart\":\"argo-cd-2.5.4\"},\"name\":\"argocd-server\",\"namespace\":\"default\"},\"spec\":{\"replicas\":1,\"revisionHistoryLimit\":5,\"selector\":{\"matchLabels\":{\"app.kubernetes.io/instance\":\"argocd\",\"app.kubernetes.io/name\":\"argocd-server\"}},\"template\":{\"metadata\":{\"labels\":{\"app.kubernetes.io/component\":\"server\",\"app.kubernetes.io/instance\":\"argocd\",\"app.kubernetes.io/managed-by\":\"Helm\",\"app.kubernetes.io/name\":\"argocd-server\",\"app.kubernetes.io/part-of\":\"argocd\",\"app.kubernetes.io/version\":\"v1.6.1\",\"helm.sh/chart\":\"argo-cd-2.5.4\"}},\"spec\":{\"containers\":[{\"command\":[\"argocd-server\",\"--staticassets\",\"/shared/app\",\"--repo-server\",\"argocd-repo-server:8081\",\"--dex-server\",\"http://argocd-dex-server:5556\",\"--loglevel\",\"info\",\"--redis\",\"argocd-redis:6379\"],\"image\":\"argoproj/argocd:v1.6.1\",\"imagePullPolicy\":\"IfNotPresent\",\"livenessProbe\":{\"failureThreshold\":3,\"httpGet\":{\"path\":\"/healthz\",\"port\":8080},\"initialDelaySeconds\":10,\"periodSeconds\":10,\"successThreshold\":1,\"timeoutSeconds\":1},\"name\":\"server\",\"ports\":[{\"containerPort\":8080,\"name\":\"server\",\"protocol\":\"TCP\"}],\"readinessProbe\":{\"failureThreshold\":3,\"httpGet\":{\"path\":\"/healthz\",\"port\":8080},\"initialDelaySeconds\":10,\"periodSeconds\":10,\"successThreshold\":1,\"timeoutSeconds\":1},\"resources\":{},\"volumeMounts\":[{\"mountPath\":\"/app/config/ssh\",\"name\":\"ssh-known-hosts\"}]}],\"serviceAccountName\":\"argocd-server\",\"volumes\":[{\"emptyDir\":{},\"name\":\"static-files\"},{\"configMap\":{\"name\":\"argocd-ssh-known-hosts-cm\"},\"name\":\"ssh-known-hosts\"}]}}}}\n",
        },
        "creation_timestamp": "2020-08-04T18:50:43Z",
        "generation": 1,
        "labels": {
            "app.kubernetes.io/component": "server",
            "app.kubernetes.io/instance": "argocd",
            "app.kubernetes.io/managed-by": "pulumi",
            "app.kubernetes.io/name": "argocd-server",
            "app.kubernetes.io/part-of": "argocd",
            "app.kubernetes.io/version": "v1.6.1",
            "helm.sh/chart": "argo-cd-2.5.4",
        },
        "name": "argocd-server",
        "namespace": "default",
        "resource_version": "1406",
        "self_link": "/apis/apps/v1/namespaces/default/deployments/argocd-server",
        "uid": "4b806e77-b035-41a3-bdf9-9781b76445f9",
    },
    spec={
        "progress_deadline_seconds": 600,
        "replicas": 1,
        "revision_history_limit": 5,
        "selector": {
            "match_labels": {
                "app.kubernetes.io/instance": "argocd",
                "app.kubernetes.io/name": "argocd-server",
            },
        },
        "strategy": {
            "rolling_update": {
                "max_surge": "25%",
                "max_unavailable": "25%",
            },
            "type": "RollingUpdate",
        },
        "template": {
            "metadata": {
                "creation_timestamp": None,
                "labels": {
                    "app.kubernetes.io/component": "server",
                    "app.kubernetes.io/instance": "argocd",
                    "app.kubernetes.io/managed-by": "Helm",
                    "app.kubernetes.io/name": "argocd-server",
                    "app.kubernetes.io/part-of": "argocd",
                    "app.kubernetes.io/version": "v1.6.1",
                    "helm.sh/chart": "argo-cd-2.5.4",
                },
            },
            "spec": {
                "containers": [{
                    "command": [
                        "argocd-server",
                        "--staticassets",
                        "/shared/app",
                        "--repo-server",
                        "argocd-repo-server:8081",
                        "--dex-server",
                        "http://argocd-dex-server:5556",
                        "--loglevel",
                        "info",
                        "--redis",
                        "argocd-redis:6379",
                    ],
                    "image": "argoproj/argocd:v1.6.1",
                    "image_pull_policy": "IfNotPresent",
                    "liveness_probe": {
                        "failure_threshold": 3,
                        "http_get": {
                            "path": "/healthz",
                            "port": 8080,
                            "scheme": "HTTP",
                        },
                        "initial_delay_seconds": 10,
                        "period_seconds": 10,
                        "success_threshold": 1,
                        "timeout_seconds": 1,
                    },
                    "name": "server",
                    "ports": [{
                        "container_port": 8080,
                        "name": "server",
                        "protocol": "TCP",
                    }],
                    "readiness_probe": {
                        "failure_threshold": 3,
                        "http_get": {
                            "path": "/healthz",
                            "port": 8080,
                            "scheme": "HTTP",
                        },
                        "initial_delay_seconds": 10,
                        "period_seconds": 10,
                        "success_threshold": 1,
                        "timeout_seconds": 1,
                    },
                    "resources": {},
                    "termination_message_path": "/dev/termination-log",
                    "termination_message_policy": "File",
                    "volume_mounts": [{
                        "mount_path": "/app/config/ssh",
                        "name": "ssh-known-hosts",
                    }],
                }],
                "dns_policy": "ClusterFirst",
                "restart_policy": "Always",
                "scheduler_name": "default-scheduler",
                "security_context": {},
                "service_account": "argocd-server",
                "service_account_name": "argocd-server",
                "termination_grace_period_seconds": 30,
                "volumes": [
                    {
                        "empty_dir": {},
                        "name": "static-files",
                    },
                    {
                        "config_map": {
                            "default_mode": 420,
                            "name": "argocd-ssh-known-hosts-cm",
                        },
                        "name": "ssh-known-hosts",
                    },
                ],
            },
        },
    })
