import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {
        labels: {
            "app.kubernetes.io/component": "server",
            "aws:region": "us-west-2",
            `key%percent`: "percent",
        },
        name: "argocd-server",
    },
});
