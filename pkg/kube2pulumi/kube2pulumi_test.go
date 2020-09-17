package kube2pulumi

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getLangs() map[string]string {
	return map[string]string{
		"python":     ".py",
		"typescript": ".ts",
		"csharp":     ".cs",
		"go":         ".go",
	}
}

// GENERAL TESTS

func TestOperator(t *testing.T) {
	assertion := assert.New(t)
	langs := getLangs()

	for language, ext := range langs {
		expected, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata",
			"k8sOperator", fmt.Sprintf("expectedMain%s", ext)))
		assertion.NoError(err)

		path, diags, err := Kube2PulumiDirectory(filepath.Join("..", "..", "testdata", "k8sOperator"), language)
		assertion.NoError(err)
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")

		generated, err := ioutil.ReadFile(path)
		assertion.NoError(err)

		assertion.Equal(string(expected), string(generated), fmt.Sprintf("%s codegen is incorrect", language))
	}
}

func TestDoubleQuotes(t *testing.T) {
	assertion := assert.New(t)
	langs := getLangs()

	for language := range langs {
		_, diags, err := Kube2PulumiFile(filepath.Join("..", "..", "testdata", "doubleQuotes.yaml"), language)
		if diags != nil {
			assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
		}
		assertion.NoError(err)
	}
}

func TestSpecialChar(t *testing.T) {
	assertion := assert.New(t)
	langs := getLangs()

	for language := range langs {
		_, diags, err := Kube2PulumiFile(filepath.Join("..", "..", "testdata", "specialChar.yaml"), language)
		if diags != nil {
			assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
		}
		assertion.NoError(err)
	}
}

func TestAnnotations(t *testing.T) {
	assertion := assert.New(t)
	langs := getLangs()

	for language := range langs {
		_, diags, err := Kube2PulumiFile(filepath.Join("..", "..", "testdata", "testDep.yaml"), language)
		if diags != nil {
			assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
		}
		assertion.NoError(err)
	}
}

func TestMultiLineString(t *testing.T) {
	assertion := assert.New(t)
	langs := getLangs()

	for language, ext := range langs {
		expected, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata",
			fmt.Sprintf("MultilineString%s", ext)))
		assertion.NoError(err)

		path, diags, err := Kube2PulumiFile(filepath.Join("..", "..", "testdata", "MultilineString.yaml"), language)
		assertion.NoError(err)
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")

		generated, err := ioutil.ReadFile(path)
		assertion.NoError(err)

		assertion.Equal(string(expected), string(generated), fmt.Sprintf("%s codegen is incorrect", language))
	}
}

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
	path, diags, err := Kube2PulumiFile(filepath.Join("..", "..", "testdata", "Namespace.yaml"), "python")
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)

	py, err := ioutil.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(pyExpected, string(py), "python codegen is incorrect")
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
	path, diags, err := Kube2PulumiFile(filepath.Join("..", "..", "testdata", "Namespace.yaml"), "typescript")
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)

	ts, err := ioutil.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(tsExpected, string(ts), "typescript codegen is incorrect")
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
	path, diags, err := Kube2PulumiFile(filepath.Join("..", "..", "testdata", "Namespace.yaml"), "csharp")
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)

	cs, err := ioutil.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(csExpected, string(cs), "C# codegen is incorrect")
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
	path, diags, err := Kube2PulumiFile(filepath.Join("..", "..", "testdata", "Namespace.yaml"), "go")
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)

	_go, err := ioutil.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(goExpected, string(_go), "golang codegen is incorrect")
}
