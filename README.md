# kube2pulumi

Convert Kubernetes YAML to Pulumi programs in Go, TypeScript, Python, and C#. Improve your Kubernetes development experience by taking advantage of strong types, compilation errors, full IDE support for features like autocomplete. Declare and manage the infrastructure in any cloud in the same program that manages your Kubernetes resources.

## Prerequisites
1. [Pulumi CLI](https://pulumi.io/quickstart/install.html)
2. Install the Pulumi Kubernetes plugin:
```console
$ pulumi plugin install resource kubernetes v2.4.2
```

## Building and Installation

If you wish to use `kube2pulumi` without developing the tool itself, you can use one of the [binary
releases](https://github.com/pulumi/kube2pulumi/releases) hosted on GitHub.


kube2pulumi uses [Go modules](https://github.com/golang/go/wiki/Modules) to manage dependencies. If you want to develop `kube2pulumi` itself, you'll need to have [Go](https://golang.org/)  installed in order to build.
Once this prerequisite is installed, run the following to build the `kube2pulumi` binary and install it into `$GOPATH/bin`:

```console
$ go build -o $GOPATH/bin/kube2pulumi cmd/kube2pulumi/main.go
```

Go should automatically handle pulling the dependencies for you.

If `$GOPATH/bin` is not on your path, you may want to move the `kube2pulumi` binary from `$GOPATH/bin`
into a directory that is on your path.

## Usage

In order to use `kube2pulumi` to convert Kubernetes YAML to Pulumi and then deploy it,
you'll first need to install the [Pulumi CLI](https://pulumi.io/quickstart/install.html). 

Once the
Pulumi CLI has been installed, you'll need to install the Kubernetes plugin:

```console
$ pulumi plugin install resource kubernetes v2.4.2
```

Now, navigate to the same directory as the the YAML you'd like to
convert and create a new Pulumi stack in your favorite language:

```console
// For a Go project
$ pulumi new kubernetes-go -f

// For a TypeScript project
$ pulumi new kubernetes-typescript -f

// For a Python project
$ pulumi new kubernetes-python -f

// For a C# project
$ pulumi new kubernetes-csharp -f
```

Then run `kube2pulumi` which will write a file in the directory that
contains the Pulumi project you just created:

```console
// For a Go project
$ kube2pulumi go -d ./

// For a TypeScript project
$ kube2pulumi typescript -d ./

// For a Python project
$ kube2pulumi python -d ./

// For a C# project
$ kube2pulumi C# -d ./
```

This will generate a Pulumi  program that when run with `pulumi update` will deploy the
Kubernetes resources originally described by your YAML. Note that before deployment you will need to [configure Kubernetes](https://www.pulumi.com/docs/intro/cloud-providers/kubernetes/setup/) so the Pulumi CLI can connect to a Kubernetes cluster. If you have previously configured the [kubectl CLI](https://kubernetes.io/docs/reference/kubectl/overview/), `kubectl`, Pulumi will respect and use your configuration settings.

## Example

Let's convert a simple YAML file describing a pod with a single container running nginx:

```yaml
apiVersion: v1
kind: Pod
metadata:
  namespace: foo
  name: bar
spec:
  containers:
    - name: nginx
      image: nginx:1.14-alpine
      resources:
        limits:
          memory: 20Mi
          cpu: 0.2

```

### Go

```console
kube2pulumi go -m ./pod.yaml
```

```go
package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := corev1.NewPod(ctx, "fooBarPod", &corev1.PodArgs{
			ApiVersion: pulumi.String("v1"),
			Kind:       pulumi.String("Pod"),
			Metadata: &metav1.ObjectMetaArgs{
				Namespace: pulumi.String("foo"),
				Name:      pulumi.String("bar"),
			},
			Spec: &corev1.PodSpecArgs{
				Containers: corev1.ContainerArray{
					&corev1.ContainerArgs{
						Name:  pulumi.String("nginx"),
						Image: pulumi.String("nginx:1.14-alpine"),
						Resources: &corev1.ResourceRequirementsArgs{
							Limits: pulumi.StringMap{
								"memory": pulumi.String("20Mi"),
								"cpu":    pulumi.String("0.2"),
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```

### TypeScript

```console
kube2pulumi typescript -m ./pod.yaml
```

```ts
import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const fooBarPod = new kubernetes.core.v1.Pod("fooBarPod", {
    apiVersion: "v1",
    kind: "Pod",
    metadata: {
        namespace: "foo",
        name: "bar",
    },
    spec: {
        containers: [{
            name: "nginx",
            image: "nginx:1.14-alpine",
            resources: {
                limits: {
                    memory: "20Mi",
                    cpu: 0.2,
                },
            },
        }],
    },
});
```

### Python

```console
kube2pulumi python -m ./pod.yaml
```

```py
import pulumi
import pulumi_kubernetes as kubernetes

foo_bar_pod = kubernetes.core.v1.Pod("fooBarPod",
    api_version="v1",
    kind="Pod",
    metadata={
        "namespace": "foo",
        "name": "bar",
    },
    spec={
        "containers": [{
            "name": "nginx",
            "image": "nginx:1.14-alpine",
            "resources": {
                "limits": {
                    "memory": "20Mi",
                    "cpu": "0.2",
                },
            },
        }],
    })
```

### C#

```console
kube2pulumi C# -m ./pod.yaml
```

```cs
using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {
        var fooBarPod = new Kubernetes.Core.v1.Pod("fooBarPod", new Kubernetes.Core.v1.PodArgs
        {
            ApiVersion = "v1",
            Kind = "Pod",
            Metadata = new Kubernetes.Meta.Inputs.ObjectMetaArgs
            {
                Namespace = "foo",
                Name = "bar",
            },
            Spec = new Kubernetes.Core.Inputs.PodSpecArgs
            {
                Containers = 
                {
                    new Kubernetes.Core.Inputs.ContainerArgs
                    {
                        Name = "nginx",
                        Image = "nginx:1.14-alpine",
                        Resources = new Kubernetes.Core.Inputs.ResourceRequirementsArgs
                        {
                            Limits = 
                            {
                                { "memory", "20Mi" },
                                { "cpu", "0.2" },
                            },
                        },
                    },
                },
            },
        });
    }

}
```