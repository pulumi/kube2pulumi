package pcl2pulumi

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

// TODO: https://github.com/pulumi/pulumi/issues/5101

func TestPulumiPython(t *testing.T) {
	namespacePy(t)
}

func namespacePy(t *testing.T) {
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
	pcl, err := ioutil.ReadFile("../testdata/Namespace.pp")
	assertion.NoError(err)

	err = Pcl2Pulumi(string(pcl), "../testdata/Namespace", "python")
	assertion.NoError(err)

	py, err := ioutil.ReadFile("../testdata/Namespace.py")
	assertion.NoError(err)

	assertion.Equal(pyExpected, string(py), "python codegen is incorrect")
}

func TestPulumiTypeScript(t *testing.T) {
	namespaceTs(t)
}

func namespaceTs(t *testing.T) {
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
	pcl, err := ioutil.ReadFile("../testdata/Namespace.pp")
	assertion.NoError(err)

	err = Pcl2Pulumi(string(pcl), "../testdata/Namespace", "nodejs")
	assertion.NoError(err)

	ts, err := ioutil.ReadFile("../testdata/Namespace.ts")
	assertion.NoError(err)

	assertion.Equal(tsExpected, string(ts), "typescript codegen is incorrect")
}

func TestPulumiDotNet(t *testing.T) {
	namespaceDotNet(t)
}

func namespaceDotNet(t *testing.T) {
	assertion := assert.New(t)

	csExpected := `using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Kubernetes.Core.v1.Namespace("foo", new Kubernetes.Core.v1.NamespaceArgs
        {
            ApiVersion = "v1",
            Kind = "Namespace",
            Metadata = new Kubernetes.Meta.Inputs.ObjectMetaArgs
            {
                Name = "foo",
            },
        });
    }

}
`
	pcl, err := ioutil.ReadFile("../testdata/Namespace.pp")
	assertion.NoError(err)

	err = Pcl2Pulumi(string(pcl), "../testdata/Namespace", "dotnet")
	assertion.NoError(err)

	cs, err := ioutil.ReadFile("../testdata/Namespace.cs")
	assertion.NoError(err)

	assertion.Equal(csExpected, string(cs), "C# codegen is incorrect")
}

func TestPulumiGoLang(t *testing.T) {
	namespaceGo(t)
}

func namespaceGo(t *testing.T) {
	assertion := assert.New(t)

	goExpected := `package main

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := core.v1.NewNamespace(ctx, "foo", &core.v1.NamespaceArgs{
			ApiVersion: pulumi.String("v1"),
			Kind:       pulumi.String("Namespace"),
			Metadata: &meta.ObjectMetaArgs{
				Name: pulumi.String("foo"),
			},
		})
		
			return err
		}
		return nil
	})
}
`
	pcl, err := ioutil.ReadFile("../testdata/Namespace.pp")
	assertion.NoError(err)

	err = Pcl2Pulumi(string(pcl), "../testdata/Namespace", "go")
	assertion.NoError(err)

	_go, err := ioutil.ReadFile("../testdata/Namespace.go")
	assertion.NoError(err)

	assertion.Equal(goExpected, string(_go), "golang codegen is incorrect")
}
