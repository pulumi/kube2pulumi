resource "argocd_serverDeployment" "kubernetes:apps/v1:Deployment" {
apiVersion = "apps/v1"
kind = "Deployment"
metadata = {
labels = {
"app.kubernetes.io/component" = "server"
"aws:region" = "us-west-2"
}
name = "argocd-server"
}
}
