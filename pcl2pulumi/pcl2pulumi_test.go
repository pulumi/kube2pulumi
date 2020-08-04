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
	pyExpected := `import pulumi
import pulumi_kubernetes as kubernetes

foo = kubernetes.core_v1.Namespace("foo",
    api_version="v1",
    kind="Namespace",
    metadata={
        "name": "foo",
    })
`
	pcl, _ := ioutil.ReadFile("testdata/Namespace.pp")
	err := Pcl2Pulumi(string(pcl), "testdata/Namespace", "python")
	if err != nil {
		return
	}
	py, _ := ioutil.ReadFile("testdata/Namespace.py")
	assert.Equal(t, pyExpected, string(py))
}

func TestPulumiTypeScript(t *testing.T) {
	namespaceTs(t)
}

func namespaceTs(t *testing.T) {
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
	pcl, _ := ioutil.ReadFile("testdata/Namespace.pp")
	err := Pcl2Pulumi(string(pcl), "testdata/Namespace", "nodejs")
	if err != nil {
		return
	}
	ts, _ := ioutil.ReadFile("testdata/Namespace.ts")
	assert.Equal(t, tsExpected, string(ts))
}

func TestPulumiDotNet(t *testing.T) {
	namespaceDotNet(t)
}

func namespaceDotNet(t *testing.T) {
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
	pcl, _ := ioutil.ReadFile("testdata/Namespace.pp")
	err := Pcl2Pulumi(string(pcl), "testdata/Namespace", "dotnet")
	if err != nil {
		return
	}
	cs, _ := ioutil.ReadFile("testdata/Namespace.cs")
	assert.Equal(t, csExpected, string(cs))
}

func TestPulumiGoLang(t *testing.T) {
	namespaceGo(t)
}

func namespaceGo(t *testing.T) {
	goExpected := `package main

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := core / v1.NewNamespace(ctx, "foo", &core/v1.NamespaceArgs{
			ApiVersion: pulumi.String("v1"),
			Kind:       pulumi.String("Namespace"),
			Metadata: &meta.ObjectMetaArgs{
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
	pcl, _ := ioutil.ReadFile("testdata/Namespace.pp")
	err := Pcl2Pulumi(string(pcl), "testdata/Namespace", "go")
	if err != nil {
		return
	}
	_go, _ := ioutil.ReadFile("testdata/Namespace.go")
	assert.Equal(t, goExpected, string(_go))
}
