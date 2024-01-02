resource "argocd_serverDeployment" "kubernetes:apps/v1:Deployment" {
apiVersion = "apps/v1"
kind = "Deployment"
metadata = {
labels = {
"app.kubernetes.io/component" = "server"
"app.kubernetes.io/instance" = "argocd"
"app.kubernetes.io/managed-by" = "pulumi"
"app.kubernetes.io/name" = "argocd-server"
"app.kubernetes.io/part-of" = "argocd"
"app.kubernetes.io/version" = "v1.6.1"
"helm.sh/chart" = "argo-cd-2.5.4"
}
name = "argocd-server"
}
spec = {
selector = {
matchLabels = {
app = "argocd"
}
}
template = {
spec = {
containers = [
{
name = "testcontainer"
readinessProbe = {
httpGet = {
port = 8080
}
}
}
]
}
}
}
}
