package pcl2pulumi

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: https://github.com/pulumi/pulumi/issues/5101

// PYTHON CODEGEN TESTS

func TestNamespacePy(t *testing.T) {
	assertion := assert.New(t)

	pyExpected := `import pulumi
import pulumi_kubernetes as kubernetes

foo = kubernetes.core.v1.Namespace("foo",
    api_version="v1",
    kind="Namespace",
    metadata={
        "name": "foo",
    })
`
	pcl, err := ioutil.ReadFile("../../testdata/Namespace.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/Namespace", "python")
	assertion.NoError(err)

	py, err := ioutil.ReadFile("../../testdata/Namespace.py")
	assertion.NoError(err)

	assertion.Equal(pyExpected, string(py), "python codegen is incorrect")
}

func TestOperatorPy(t *testing.T) {
	assertion := assert.New(t)

	pyExpected, err := ioutil.ReadFile("../../testdata/k8sOperator/expectedMain.py")
	assertion.NoError(err)

	pcl, err := ioutil.ReadFile("../../testdata/expK8sOperator.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/k8sOperator/main", "python")
	assertion.NoError(err)

	py, err := ioutil.ReadFile("../../testdata/k8sOperator/main.py")
	assertion.NoError(err)

	assertion.Equal(string(pyExpected), string(py), "python operator codegen is incorrect")
}

// TYPESCRIPT CODEGEN TESTS

func TestNamespaceTs(t *testing.T) {
	assertion := assert.New(t)

	tsExpected := `import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const foo = new kubernetes.core.v1.Namespace("foo", {
    apiVersion: "v1",
    kind: "Namespace",
    metadata: {
        name: "foo",
    },
});
`
	pcl, err := ioutil.ReadFile("../../testdata/Namespace.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/Namespace", "nodejs")
	assertion.NoError(err)

	ts, err := ioutil.ReadFile("../../testdata/Namespace.ts")
	assertion.NoError(err)

	assertion.Equal(tsExpected, string(ts), "typescript codegen is incorrect")
}

func TestOperatorTs(t *testing.T) {
	assertion := assert.New(t)

	tsExpected, err := ioutil.ReadFile("../../testdata/k8sOperator/expectedMain.ts")
	assertion.NoError(err)

	pcl, err := ioutil.ReadFile("../../testdata/expK8sOperator.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/k8sOperator/main", "nodejs")
	assertion.NoError(err)

	ts, err := ioutil.ReadFile("../../testdata/k8sOperator/main.ts")
	assertion.NoError(err)

	assertion.Equal(string(tsExpected), string(ts), "typescript operator codegen is incorrect")
}

// C# CODEGEN TESTS

func TestNamespaceDotNet(t *testing.T) {
	assertion := assert.New(t)

	csExpected := `using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Kubernetes.Core.V1.Namespace("foo", new Kubernetes.Types.Inputs.Core.V1.NamespaceArgs
        {
            ApiVersion = "v1",
            Kind = "Namespace",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Name = "foo",
            },
        });
    }

}
`
	pcl, err := ioutil.ReadFile("../../testdata/Namespace.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/Namespace", "dotnet")
	assertion.NoError(err)

	cs, err := ioutil.ReadFile("../../testdata/Namespace.cs")
	assertion.NoError(err)

	assertion.Equal(csExpected, string(cs), "C# codegen is incorrect")
}

func TestOperatorCs(t *testing.T) {
	assertion := assert.New(t)

	csExpected, err := ioutil.ReadFile("../../testdata/k8sOperator/expectedMain.cs")
	assertion.NoError(err)

	pcl, err := ioutil.ReadFile("../../testdata/expK8sOperator.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/k8sOperator/main", "dotnet")
	assertion.NoError(err)

	cs, err := ioutil.ReadFile("../../testdata/k8sOperator/main.cs")
	assertion.NoError(err)

	assertion.Equal(string(csExpected), string(cs), "dotnet operator codegen is incorrect")
}

// GOLANG CODEGEN TESTS

func TestNamespaceGo(t *testing.T) {
	assertion := assert.New(t)

	goExpected := `package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := corev1.NewNamespace(ctx, "foo", &corev1.NamespaceArgs{
			ApiVersion: pulumi.String("v1"),
			Kind:       pulumi.String("Namespace"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("foo"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
`
	pcl, err := ioutil.ReadFile("../../testdata/Namespace.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/Namespace", "go")
	assertion.NoError(err)

	_go, err := ioutil.ReadFile("../../testdata/Namespace.go")
	assertion.NoError(err)

	assertion.Equal(goExpected, string(_go), "golang codegen is incorrect")
}

func TestOperatorGo(t *testing.T) {
	assertion := assert.New(t)

	goExpected, err := ioutil.ReadFile("../../testdata/k8sOperator/expectedMain.go")
	assertion.NoError(err)

	pcl, err := ioutil.ReadFile("../../testdata/expK8sOperator.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/k8sOperator/main", "go")
	assertion.NoError(err)

	_go, err := ioutil.ReadFile("../../testdata/k8sOperator/main.go")
	assertion.NoError(err)

	assertion.Equal(string(goExpected), string(_go), "golang operator codegen is incorrect")
}
