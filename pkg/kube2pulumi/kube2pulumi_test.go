package kube2pulumi

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

// PYTHON CODEGEN TESTS

func TestNamespacePy(t *testing.T) {
	assertion := assert.New(t)

	pyExpected := `import pulumi
import pulumi_kubernetes as kubernetes

foo_namespace = kubernetes.core.v1.Namespace("fooNamespace",
    api_version="v1",
    kind="Namespace",
    metadata={
        "name": "foo",
    })
`
	path, err := Kube2PulumiFile("../../testdata/Namespace.yaml", "python")
	assertion.NoError(err)

	py, err := ioutil.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(pyExpected, string(py), "python codegen is incorrect")
}

func TestOperatorPy(t *testing.T) {
	assertion := assert.New(t)

	pyExpected, err := ioutil.ReadFile("../../testdata/k8sOperator/expectedMain.py")
	assertion.NoError(err)

	path, err := Kube2PulumiDirectory("../../testdata/k8sOperator/", "python")
	assertion.NoError(err)

	py, err := ioutil.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(string(pyExpected), string(py), "python operator codegen is incorrect")
}

func TestDoubleQuotesPy(t *testing.T) {
	assertion := assert.New(t)

	_, err := Kube2PulumiFile("../../testdata/doubleQuotes.yaml", "python")
	assertion.NoError(err)
}

func TestSpecialCharPy(t *testing.T) {
	assertion := assert.New(t)

	_, err := Kube2PulumiFile("../../testdata/specialChar.yaml", "python")
	assertion.NoError(err)
}

func TestAnnotationsPy(t *testing.T) {
	assertion := assert.New(t)

	_, err := Kube2PulumiFile("../../testdata/testDep.yaml", "python")
	assertion.NoError(err)
}

func TestMultiLineStringPy(t *testing.T) {
	assertion := assert.New(t)

	pyExpected, err := ioutil.ReadFile("../../testdata/expectedMultilineString.py")
	assertion.NoError(err)

	path, err := Kube2PulumiFile("../../testdata/MultilineString.yaml", "python")
	assertion.NoError(err)

	py, err := ioutil.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(string(pyExpected), string(py), "multiline gen is incorrect")
}

// TYPESCRIPT CODEGEN TESTS

func TestNamespaceTs(t *testing.T) {
	assertion := assert.New(t)

	tsExpected := `import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const fooNamespace = new kubernetes.core.v1.Namespace("fooNamespace", {
    apiVersion: "v1",
    kind: "Namespace",
    metadata: {
        name: "foo",
    },
});
`
	path, err := Kube2PulumiFile("../../testdata/Namespace.yaml", "typescript")
	assertion.NoError(err)

	ts, err := ioutil.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(tsExpected, string(ts), "typescript codegen is incorrect")
}

func TestOperatorTs(t *testing.T) {
	assertion := assert.New(t)

	tsExpected, err := ioutil.ReadFile("../../testdata/k8sOperator/expectedMain.ts")
	assertion.NoError(err)

	path, err := Kube2PulumiDirectory("../../testdata/k8sOperator/", "typescript")
	assertion.NoError(err)

	ts, err := ioutil.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(string(tsExpected), string(ts), "typescript operator codegen is incorrect")
}

func TestDoubleQuotesTs(t *testing.T) {
	assertion := assert.New(t)

	_, err := Kube2PulumiFile("../../testdata/doubleQuotes.yaml", "typescript")
	assertion.NoError(err)
}

func TestSpecialCharTs(t *testing.T) {
	assertion := assert.New(t)

	_, err := Kube2PulumiFile("../../testdata/specialChar.yaml", "typescript")
	assertion.NoError(err)
}

func TestAnnotationsTs(t *testing.T) {
	assertion := assert.New(t)

	_, err := Kube2PulumiFile("../../testdata/testDep.yaml", "typescript")
	assertion.NoError(err)
}

func TestMultiLineStringTs(t *testing.T) {
	assertion := assert.New(t)

	tsExpected, err := ioutil.ReadFile("../../testdata/expectedMultilineString.ts")
	assertion.NoError(err)

	path, err := Kube2PulumiFile("../../testdata/MultilineString.yaml", "typescript")
	assertion.NoError(err)

	ts, err := ioutil.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(string(tsExpected), string(ts), "multiline gen is incorrect")
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
        var fooNamespace = new Kubernetes.Core.V1.Namespace("fooNamespace", new Kubernetes.Types.Inputs.Core.V1.NamespaceArgs
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
	path, err := Kube2PulumiFile("../../testdata/Namespace.yaml", "csharp")
	assertion.NoError(err)

	cs, err := ioutil.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(csExpected, string(cs), "C# codegen is incorrect")
}

func TestOperatorCs(t *testing.T) {
	assertion := assert.New(t)

	csExpected, err := ioutil.ReadFile("../../testdata/k8sOperator/expectedMain.cs")
	assertion.NoError(err)

	path, err := Kube2PulumiDirectory("../../testdata/k8sOperator", "csharp")
	assertion.NoError(err)

	cs, err := ioutil.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(string(csExpected), string(cs), "csharp operator codegen is incorrect")
}

func TestDoubleQuotesCs(t *testing.T) {
	assertion := assert.New(t)

	_, err := Kube2PulumiFile("../../testdata/doubleQuotes.yaml", "csharp")
	assertion.NoError(err)
}

func TestSpecialCharCs(t *testing.T) {
	assertion := assert.New(t)

	_, err := Kube2PulumiFile("../../testdata/specialChar.yaml", "csharp")
	assertion.NoError(err)
}

func TestAnnotationsCs(t *testing.T) {
	assertion := assert.New(t)

	_, err := Kube2PulumiFile("../../testdata/testDep.yaml", "csharp")
	assertion.NoError(err)
}

func TestMultiLineStringCs(t *testing.T) {
	assertion := assert.New(t)

	csExpected, err := ioutil.ReadFile("../../testdata/expectedMultilineString.cs")
	assertion.NoError(err)

	path, err := Kube2PulumiFile("../../testdata/MultilineString.yaml", "csharp")
	assertion.NoError(err)

	cs, err := ioutil.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(string(csExpected), string(cs), "multiline gen is incorrect")
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
		_, err := corev1.NewNamespace(ctx, "fooNamespace", &corev1.NamespaceArgs{
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
	path, err := Kube2PulumiFile("../../testdata/Namespace.yaml", "go")
	assertion.NoError(err)

	_go, err := ioutil.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(goExpected, string(_go), "golang codegen is incorrect")
}

func TestOperatorGo(t *testing.T) {
	assertion := assert.New(t)

	goExpected, err := ioutil.ReadFile("../../testdata/k8sOperator/expectedMain.go")
	assertion.NoError(err)

	path, err := Kube2PulumiDirectory("../../testdata/k8sOperator/", "go")
	assertion.NoError(err)

	_go, err := ioutil.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(string(goExpected), string(_go), "golang operator codegen is incorrect")
}

func TestDoubleQuotesGo(t *testing.T) {
	assertion := assert.New(t)

	_, err := Kube2PulumiFile("../../testdata/doubleQuotes.yaml", "go")
	assertion.NoError(err)
}

func TestSpecialCharGo(t *testing.T) {
	assertion := assert.New(t)

	_, err := Kube2PulumiFile("../../testdata/specialChar.yaml", "go")
	assertion.NoError(err)
}

func TestAnnotationsGo(t *testing.T) {
	assertion := assert.New(t)

	_, err := Kube2PulumiFile("../../testdata/testDep.yaml", "go")
	assertion.NoError(err)
}

func TestMultiLineStringGo(t *testing.T) {
	assertion := assert.New(t)

	goExpected, err := ioutil.ReadFile("../../testdata/expectedMultilineString.go")
	assertion.NoError(err)

	path, err := Kube2PulumiFile("../../testdata/MultilineString.yaml", "go")
	assertion.NoError(err)

	_go, err := ioutil.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(string(goExpected), string(_go), "multiline gen is incorrect")
}
