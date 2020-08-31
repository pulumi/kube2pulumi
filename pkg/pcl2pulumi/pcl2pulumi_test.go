package pcl2pulumi

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

foo = kubernetes.core.v1.Namespace("foo",
    api_version="v1",
    kind="Namespace",
    metadata={
        "name": "foo",
    })
`
	pcl, err := ioutil.ReadFile("../../testdata/Namespace.pp")
	assertion.NoError(err)

	outPath, err := Pcl2Pulumi(string(pcl), "../../testdata/Namespace", "python")
	assertion.NoError(err)

	py, err := ioutil.ReadFile(outPath)
	assertion.NoError(err)

	assertion.Equal(pyExpected, string(py), "python codegen is incorrect")
}

func TestOperatorPy(t *testing.T) {
	assertion := assert.New(t)

	pyExpected, err := ioutil.ReadFile("../../testdata/k8sOperator/expectedMain.py")
	assertion.NoError(err)

	pcl, err := ioutil.ReadFile("../../testdata/expK8sOperator.pp")
	assertion.NoError(err)

	outPath, err := Pcl2Pulumi(string(pcl), "../../testdata/k8sOperator/main", "python")
	assertion.NoError(err)

	py, err := ioutil.ReadFile(outPath)
	assertion.NoError(err)

	assertion.Equal(string(pyExpected), string(py), "python operator codegen is incorrect")
}

func TestDoubleQuotesPy(t *testing.T) {
	assertion := assert.New(t)

	pcl, err := ioutil.ReadFile("../../testdata/doubleQuotes.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/doubleQuotes", "python")
	assertion.NoError(err)
}

func TestSpecialCharPy(t *testing.T) {
	assertion := assert.New(t)

	pcl, err := ioutil.ReadFile("../../testdata/specialChar.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/specialChar", "python")
	assertion.NoError(err)
}

func TestAnnotationsPy(t *testing.T) {
	assertion := assert.New(t)

	pcl, err := ioutil.ReadFile("../../testdata/testDep.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/testDep", "python")
	assertion.NoError(err)
}

func TestMultiLineStringPy(t *testing.T) {
	assertion := assert.New(t)

	pyExpected, err := ioutil.ReadFile("../../testdata/expectedMultilineString.py")
	assertion.NoError(err)

	pcl, err := ioutil.ReadFile("../../testdata/MultilineString.pp")
	assertion.NoError(err)

	outPath, err := Pcl2Pulumi(string(pcl), "../../testdata/MultilineString", "python")
	assertion.NoError(err)

	py, err := ioutil.ReadFile(outPath)
	assertion.NoError(err)

	assertion.Equal(string(pyExpected), string(py), "multiline gen is incorrect")
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

	outPath, err := Pcl2Pulumi(string(pcl), "../../testdata/Namespace", "typescript")
	assertion.NoError(err)

	ts, err := ioutil.ReadFile(outPath)
	assertion.NoError(err)

	assertion.Equal(tsExpected, string(ts), "typescript codegen is incorrect")
}

func TestOperatorTs(t *testing.T) {
	assertion := assert.New(t)

	tsExpected, err := ioutil.ReadFile("../../testdata/k8sOperator/expectedMain.ts")
	assertion.NoError(err)

	pcl, err := ioutil.ReadFile("../../testdata/expK8sOperator.pp")
	assertion.NoError(err)

	outPath, err := Pcl2Pulumi(string(pcl), "../../testdata/k8sOperator/main", "typescript")
	assertion.NoError(err)

	ts, err := ioutil.ReadFile(outPath)
	assertion.NoError(err)

	assertion.Equal(string(tsExpected), string(ts), "typescript operator codegen is incorrect")
}

func TestDoubleQuotesTs(t *testing.T) {
	assertion := assert.New(t)

	pcl, err := ioutil.ReadFile("../../testdata/doubleQuotes.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/doubleQuotes", "typescript")
	assertion.NoError(err)
}

func TestSpecialCharTs(t *testing.T) {
	assertion := assert.New(t)

	pcl, err := ioutil.ReadFile("../../testdata/specialChar.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/specialChar", "typescript")
	assertion.NoError(err)
}

func TestAnnotationsTs(t *testing.T) {
	assertion := assert.New(t)

	pcl, err := ioutil.ReadFile("../../testdata/testDep.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/testDep", "typescript")
	assertion.NoError(err)
}

func TestMultiLineStringTs(t *testing.T) {
	assertion := assert.New(t)

	tsExpected, err := ioutil.ReadFile("../../testdata/expectedMultilineString.ts")
	assertion.NoError(err)

	pcl, err := ioutil.ReadFile("../../testdata/MultilineString.pp")
	assertion.NoError(err)

	outPath, err := Pcl2Pulumi(string(pcl), "../../testdata/MultilineString", "typescript")
	assertion.NoError(err)

	ts, err := ioutil.ReadFile(outPath)
	assertion.NoError(err)

	assertion.Equal(string(tsExpected), string(ts), "multiline gen is incorrect")
}

// C# CODEGEN TESTS

func TestNamespaceCS(t *testing.T) {
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

	outPath, err := Pcl2Pulumi(string(pcl), "../../testdata/Namespace", "csharp")
	assertion.NoError(err)

	cs, err := ioutil.ReadFile(outPath)
	assertion.NoError(err)

	assertion.Equal(csExpected, string(cs), "C# codegen is incorrect")
}

func TestOperatorCs(t *testing.T) {
	assertion := assert.New(t)

	csExpected, err := ioutil.ReadFile("../../testdata/k8sOperator/expectedMain.cs")
	assertion.NoError(err)

	pcl, err := ioutil.ReadFile("../../testdata/expK8sOperator.pp")
	assertion.NoError(err)

	outPath, err := Pcl2Pulumi(string(pcl), "../../testdata/k8sOperator/main", "csharp")
	assertion.NoError(err)

	cs, err := ioutil.ReadFile(outPath)
	assertion.NoError(err)

	assertion.Equal(string(csExpected), string(cs), "csharp operator codegen is incorrect")
}

func TestDoubleQuotesCs(t *testing.T) {
	assertion := assert.New(t)

	pcl, err := ioutil.ReadFile("../../testdata/doubleQuotes.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/doubleQuotes", "csharp")
	assertion.NoError(err)
}

func TestSpecialCharCs(t *testing.T) {
	assertion := assert.New(t)

	pcl, err := ioutil.ReadFile("../../testdata/specialChar.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/specialChar", "csharp")
	assertion.NoError(err)
}

func TestAnnotationsCs(t *testing.T) {
	assertion := assert.New(t)

	pcl, err := ioutil.ReadFile("../../testdata/testDep.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/testDep", "csharp")
	assertion.NoError(err)
}

func TestMultiLineStringCs(t *testing.T) {
	assertion := assert.New(t)

	csExpected, err := ioutil.ReadFile("../../testdata/expectedMultilineString.cs")
	assertion.NoError(err)

	pcl, err := ioutil.ReadFile("../../testdata/MultilineString.pp")
	assertion.NoError(err)

	outPath, err := Pcl2Pulumi(string(pcl), "../../testdata/MultilineString", "csharp")
	assertion.NoError(err)

	cs, err := ioutil.ReadFile(outPath)
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

	outPath, err := Pcl2Pulumi(string(pcl), "../../testdata/Namespace", "go")
	assertion.NoError(err)

	_go, err := ioutil.ReadFile(outPath)
	assertion.NoError(err)

	assertion.Equal(goExpected, string(_go), "golang codegen is incorrect")
}

func TestOperatorGo(t *testing.T) {
	assertion := assert.New(t)

	goExpected, err := ioutil.ReadFile("../../testdata/k8sOperator/expectedMain.go")
	assertion.NoError(err)

	pcl, err := ioutil.ReadFile("../../testdata/expK8sOperator.pp")
	assertion.NoError(err)

	outPath, err := Pcl2Pulumi(string(pcl), "../../testdata/k8sOperator/main", "go")
	assertion.NoError(err)

	_go, err := ioutil.ReadFile(outPath)
	assertion.NoError(err)

	assertion.Equal(string(goExpected), string(_go), "golang operator codegen is incorrect")
}

func TestDoubleQuotesGo(t *testing.T) {
	assertion := assert.New(t)

	pcl, err := ioutil.ReadFile("../../testdata/doubleQuotes.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/doubleQuotes", "go")
	assertion.NoError(err)
}

func TestSpecialCharGo(t *testing.T) {
	assertion := assert.New(t)

	pcl, err := ioutil.ReadFile("../../testdata/specialChar.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/specialChar", "go")
	assertion.NoError(err)
}

func TestAnnotationsGo(t *testing.T) {
	assertion := assert.New(t)

	pcl, err := ioutil.ReadFile("../../testdata/testDep.pp")
	assertion.NoError(err)

	_, err = Pcl2Pulumi(string(pcl), "../../testdata/testDep", "go")
	assertion.NoError(err)
}

func TestMultiLineStringGo(t *testing.T) {
	assertion := assert.New(t)

	goExpected, err := ioutil.ReadFile("../../testdata/expectedMultilineString.go")
	assertion.NoError(err)

	pcl, err := ioutil.ReadFile("../../testdata/MultilineString.pp")
	assertion.NoError(err)

	outFile, err := Pcl2Pulumi(string(pcl), "../../testdata/MultilineString", "go")
	assertion.NoError(err)

	_go, err := ioutil.ReadFile(outFile)
	assertion.NoError(err)

	assertion.Equal(string(goExpected), string(_go), "multiline gen is incorrect")
}
