package yaml2pcl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertDirectory(t *testing.T) {
	testK8sOperator(t)
}

func testK8sOperator(t *testing.T) {
	assertion := assert.New(t)

	expected := `resource operatorDeployment "kubernetes:apps/v1:Deployment" {
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

resource operatorServiceAccount "kubernetes:core/v1:ServiceAccount" {
metadata = {
name = "pulumi-kubernetes-operator"
}
}
`
	result, err := ConvertDirectory("../testdata/k8sOperator/")
	if err != nil {
		assertion.Error(err)
	} else {
		assertion.Equal(expected, result, "Directory output does not match")
	}
}

func TestNamespace(t *testing.T) {
	assertion := assert.New(t)

	expected := `resource foo "kubernetes:core/v1:Namespace" {
apiVersion = "v1"
kind = "Namespace"
metadata = {
name = "foo"
}
}
`
	result, err := ConvertFile("testdata/Namespace.yaml")
	if err != nil {
		assertion.Error(err)
	} else {
		assertion.Equal(expected, result, "Single resource conversion was incorrect")
	}
}

func TestNamespaceComments(t *testing.T) {
	assertion := assert.New(t)

	expected := `resource foo "kubernetes:core/v1:Namespace" {
apiVersion = "v1"
kind = "Namespace"
# this is a codegentest comment
metadata = {
name = "foo"
}
}
`
	result, err := ConvertFile("testdata/NamespaceWithComments.yaml")
	if err != nil {
		assertion.Error(err)
	} else {
		assertion.Equal(expected, result, "Comments are converted incorrectly")
	}
}

func Test1PodArray(t *testing.T) {
	assertion := assert.New(t)

	expected := `resource bar "kubernetes:core/v1:Pod" {
apiVersion = "v1"
kind = "Pod"
metadata = {
namespace = "foo"
name = "bar"
}
spec = {
containers = [
{
name = "nginx"
image = "nginx:1.14-alpine"
resources = {
limits = {
memory = "20Mi"
cpu = 0.2
}
}
}
]
}
}
`
	result, err := ConvertFile("testdata/OnePodArray.yaml")
	if err != nil {
		assertion.Error(err)
	} else {
		assertion.Equal(expected, result, "Nested array is converted incorrectly")
	}
}
