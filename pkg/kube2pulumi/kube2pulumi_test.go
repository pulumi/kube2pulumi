package kube2pulumi

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/kube2pulumi/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

// GENERAL TESTS

func TestOperator(t *testing.T) {
	for language, ext := range testutil.Languages() {
		language, ext := language, ext
		t.Run(language, func(t *testing.T) {
			t.Parallel()
			assertion := assert.New(t)
			testdataDir := filepath.Join("..", "..", "testdata", "k8sOperator")
			testDir := testutil.MakeTestDir(t, testdataDir)

			path, diags, err := Kube2PulumiDirectory(testDir, "", language)
			assertion.NoError(err)
			assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")

			testutil.AssertFilesEqual(t,
				filepath.Join(testdataDir, "expected", fmt.Sprintf("expectedMain%s", ext)),
				path)
		})
	}
}

func TestDoubleQuotes(t *testing.T) {
	for language := range testutil.Languages() {
		language := language
		t.Run(language, func(t *testing.T) {
			t.Parallel()
			assertion := assert.New(t)
			testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "doubleQuotes"))
			_, diags, err := Kube2PulumiFile(filepath.Join(testDir, "doubleQuotes.yaml"), "", language)
			if diags != nil {
				assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
			}
			assertion.NoError(err)
		})
	}
}

func TestSpecialChar(t *testing.T) {
	for language := range testutil.Languages() {
		language := language
		t.Run(language, func(t *testing.T) {
			t.Parallel()
			assertion := assert.New(t)
			testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "specialChar"))
			_, diags, err := Kube2PulumiFile(filepath.Join(testDir, "specialChar.yaml"), "", language)
			if diags != nil {
				assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
			}
			assertion.NoError(err)
		})
	}
}

func TestQuotedApiVersion(t *testing.T) {
	for language := range testutil.Languages() {
		language := language
		t.Run(language, func(t *testing.T) {
			t.Parallel()
			assertion := assert.New(t)
			testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "quotedApiVersion"))
			_, diags, err := Kube2PulumiFile(filepath.Join(testDir, "quotedApiVersion.yaml"), "", language)
			if diags != nil {
				assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
			}
			assertion.NoError(err)
		})
	}
}

func TestAnnotations(t *testing.T) {
	for language := range testutil.Languages() {
		language := language
		t.Run(language, func(t *testing.T) {
			t.Parallel()
			assertion := assert.New(t)
			testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "testDep"))
			_, diags, err := Kube2PulumiFile(filepath.Join(testDir, "testDep.yaml"), "", language)
			if diags != nil {
				assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
			}
			assertion.NoError(err)
		})
	}
}

func TestMultiLineString(t *testing.T) {
	for language, ext := range testutil.Languages() {
		language, ext := language, ext
		t.Run(language, func(t *testing.T) {
			t.Parallel()
			assertion := assert.New(t)
			testdataDir := filepath.Join("..", "..", "testdata", "MultilineString")
			testDir := testutil.MakeTestDir(t, testdataDir)
			expected := filepath.Join(testdataDir, "expected", fmt.Sprintf("expectedMultilineString%s", ext))

			path, diags, err := Kube2PulumiFile(filepath.Join(testDir, "MultilineString.yaml"), "", language)
			assertion.NoError(err)
			assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")

			testutil.AssertFilesEqual(t, expected, path)
		})
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
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="foo",
    ))
`
	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "Namespace"))
	path, diags, err := Kube2PulumiFile(filepath.Join(testDir, "Namespace.yaml"), "", "python")
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)

	py, err := os.ReadFile(path)
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
	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "Namespace"))
	path, diags, err := Kube2PulumiFile(filepath.Join(testDir, "Namespace.yaml"), "", "typescript")
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)

	ts, err := os.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(tsExpected, string(ts), "typescript codegen is incorrect")
}

// C# CODEGEN TESTS

func TestNamespaceDotNet(t *testing.T) {
	assertion := assert.New(t)

	csExpected := `using System.Collections.Generic;
using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

return await Deployment.RunAsync(() => 
{
    var fooNamespace = new Kubernetes.Core.V1.Namespace("fooNamespace", new()
    {
        ApiVersion = "v1",
        Kind = "Namespace",
        Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
        {
            Name = "foo",
        },
    });

});

`
	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "Namespace"))
	path, diags, err := Kube2PulumiFile(filepath.Join(testDir, "Namespace.yaml"), "", "csharp")
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)

	cs, err := os.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(csExpected, string(cs), "C# codegen is incorrect")
}

// GOLANG CODEGEN TESTS

func TestNamespaceGo(t *testing.T) {
	assertion := assert.New(t)

	goExpected := `package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "Namespace"))
	path, diags, err := Kube2PulumiFile(filepath.Join(testDir, "Namespace.yaml"), "", "go")
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)

	_go, err := os.ReadFile(path)
	assertion.NoError(err)

	assertion.Equal(goExpected, string(_go), "golang codegen is incorrect")
}

func TestStringLiteral(t *testing.T) {
	for language := range testutil.Languages() {
		language := language
		t.Run(language, func(t *testing.T) {
			t.Parallel()
			assertion := assert.New(t)
			testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "stringLiteral"))
			kubeManifest := filepath.Join(testDir, "cm.yaml")
			outFile, diags, err := Kube2PulumiFile(kubeManifest, "", language)
			if diags != nil {
				assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
			}
			assertion.NoError(err)

			// Ensure that the generated file does not contain $$. This is a sign that the string literal was not
			// properly escaped.
			generated, err := os.ReadFile(outFile)
			assertion.NoError(err)
			assertion.NotContains(string(generated), "$$")
		})
	}
}
