resource operatorServiceAccount "kubernetes:core/v1:ServiceAccount" {
metadata = {
name = "pulumi-kubernetes-operator"
}
}

resource operatorRole "kubernetes:rbac.authorization.k8s.io/v1:Role" {
metadata = {
name = "pulumi-kubernetes-operator"
}
rules = [
{
apiGroups = [
""
]
resources = [
"pods",
"services",
"services/finalizers",
"endpoints",
"persistentvolumeclaims",
"events",
"configmaps",
"secrets"
]
verbs = [
"create",
"delete",
"get",
"list",
"patch",
"update",
"watch"
]
},
{
apiGroups = [
"apps"
]
resources = [
"deployments",
"daemonsets",
"replicasets",
"statefulsets",
]
verbs = [
"create",
"delete",
"get",
"list",
"patch",
"update",
"watch"
]
},
{
apiGroups = [
"monitoring.coreos.com"
]
resources = [
"servicemonitors",
]
verbs = [
"create",
"get",
]
},
{
apiGroups = [
"apps"
]
resourceNames = [
"pulumi-kubernetes-operator"
]
resources = [
"deployments/finalizers",
]
verbs = [
"update",
]
},
{
apiGroups = [
""
]
resources = [
"pods",
]
verbs = [
"get",
]
},
{
apiGroups = [
"apps"
]
resources = [
"replicasets",
"deployments",
]
verbs = [
"get",
]
},
{
apiGroups = [
"pulumi.com"
]
resources = [
"*",
]
verbs = [
"create",
"delete",
"get",
"list",
"patch",
"update",
"watch",
]
},
]
}

resource operatorRoleBinding "kubernetes:rbac.authorization.k8s.io/v1:RoleBinding" {
metadata = {
name = "pulumi-kubernetes-operator"
}
subjects = [
{
kind = "ServiceAccount"
name = "pulumi-kubernetes-operator"
}
]
roleRef = {
kind = "Role"
name = "pulumi-kubernetes-operator"
apiGroup = "rbac.authorization.k8s.io"
}
}

resource operatorDeployment "kubernetes:apps/v1:Deployment" {
metadata = {
name = "pulumi-kubernetes-operator"
}
spec = {
# Currently only 1 replica supported, until leader election: https://github.com/pulumi/pulumi-kubernetes-operator/issues/33
replicas = 1
selector = {
matchLabels = {
name = "pulumi-kubernetes-operator"
}
}
template = {
metadata = {
labels = {
name = "pulumi-kubernetes-operator"
}
}
spec = {
serviceAccountName = "pulumi-kubernetes-operator"
imagePullSecrets = [
{
name = "pulumi-kubernetes-operator"
}
]
containers = [
{
name = "pulumi-kubernetes-operator"
image = "pulumi/pulumi-kubernetes-operator:v0.0.1"
command = [
"pulumi-kubernetes-operator"
]
args = [
"--zap-level=debug"
]
imagePullPolicy = "Always"
env = [
{
name = "WATCH_NAMESPACE"
valueFrom = {
fieldRef = {
fieldPath = "metadata.namespace"
}
}
},
{
name = "POD_NAME"
valueFrom = {
fieldRef = {
fieldPath = "metadata.name"
}
}
},
{
name = "OPERATOR_NAME"
value = "pulumi-kubernetes-operator"
}
]
}
]
}
}
}
}
