resource test "kubernetes:rbac.authorization.k8s.io/v1:Role" {
apiVersion = "rbac.authorization.k8s.io/v1"
kind = "Role"
metadata = {
name = "test"
}
rules = [
{
apiGroups = [
""
]
resources = [
"pods\\\""
]
verbs = [
"create"
]
},
{
apiGroups = [
"pulumi.com"
]
resources = [
"*"
]
verbs = [
"create"
]
}
]
}
