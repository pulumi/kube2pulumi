resource test "kubernetes:rbac.authorization.k8s.io/v1:Role" {
metadata = {
name = "test"
}
rules = [
{
apiGroups = [
""
]
resources = [
"pods"
]
verbs = [
"create"
]
},
]
}
