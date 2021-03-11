import pulumi
import pulumi_kubernetes as kubernetes

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",
    api_version="apps/v1",
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        labels={
            "app.kubernetes.io/component": "server",
            "aws:region": "us-west-2",
            "key%percent": "percent",
            "key...ellipse": "ellipse",
            "key{bracket": "bracket",
            "key}bracket": "bracket",
            "key*asterix": "asterix",
            "key?question": "question",
            "key,comma": "comma",
            "key&&and": "and",
            "key||or": "or",
            "key!not": "not",
            "key=>geq": "geq",
            "key==eq": "equal",
        },
        name="argocd-server",
    ))
