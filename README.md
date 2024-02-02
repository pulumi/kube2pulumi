# ⚠️ This repository is deprecated.

kube2pulumi has been deprecated and is no longer maintained. Please note that this repository is no longer actively developed or supported.

Development for this project has moved to [pulumi/pulumi-converter-kubernetes](https://github.com/pulumi/pulumi-converter-kubernetes). If you are looking to convert Kubernetes manifests to Pulumi programs, we **strongly recommend** you utilize the new converter plugin. You can find detailed instructions on using the new converter plugin [here](https://www.pulumi.com/docs/using-pulumi/adopting-pulumi/migrating-to-pulumi/from-kubernetes/#converting-kubernetes-yaml).

Many of the long-requested features and bug fixes have been integrated into the newer converter plugin, making it more versatile and reliable.

---

# kube2pulumi

Convert Kubernetes YAML to Pulumi programs in Go, TypeScript, Python, C# and Java. Improve your Kubernetes development experience by taking advantage of strong types, compilation errors, full IDE support for features like autocomplete. Declare and manage the infrastructure in any cloud in the same program that manages your Kubernetes resources.

## Prerequisites
1. [Pulumi CLI](https://pulumi.io/quickstart/install.html)
2. Install the Pulumi Kubernetes plugin:
```console
$ pulumi plugin install resource kubernetes v3.0.0
```

## Building and Installation

If you wish to use `kube2pulumi` without developing the tool itself, you can use one of the [binary
releases](https://github.com/pulumi/kube2pulumi/releases) hosted on GitHub.

### Homebrew
`kube2pulumi` can be installed on Mac from the Pulumi Homebrew tap.
```console
brew install pulumi/tap/kube2pulumi
```


kube2pulumi uses [Go modules](https://github.com/golang/go/wiki/Modules) to manage dependencies. If you want to develop `kube2pulumi` itself, you'll need to have [Go](https://golang.org/)  installed in order to build.
Once this prerequisite is installed, run the following to build the `kube2pulumi` binary and install it into `$GOPATH/bin`:

```console
$ go build -o $GOPATH/bin/kube2pulumi -ldflags="-X github.com/pulumi/kube2pulumi/pkg/version.Version=1.0.0" cmd/kube2pulumi/main.go
```

The `ldflags` argument is necessary to dynamically set the `kube2pulumi` version at build time. However, the version 
itself can be anything, so you don't have to set it to `dev`.

Go should automatically handle pulling the dependencies for you. If `$GOPATH/bin` is not on your path, you may want to 
move the `kube2pulumi` binary from `$GOPATH/bin` into a directory that is on your path.

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

// For a Java project
$ pulumi new kubernetes-java -f
```

Then run `kube2pulumi` which will write a file in the directory that
contains the Pulumi project you just created:

```console
// For a Go project
$ kube2pulumi go 

// For a TypeScript project
$ kube2pulumi typescript

// For a Python project
$ kube2pulumi python

// For a C# project
$ kube2pulumi csharp

// For a Java project
$ kube2pulumi java
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
kube2pulumi go -f ./pod.yaml
```

```go
package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
kube2pulumi typescript -f ./pod.yaml
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
kube2pulumi python -f ./pod.yaml
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
kube2pulumi csharp -f ./pod.yaml
```

```cs
using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {
        var fooBarPod = new Kubernetes.Core.V1.Pod("fooBarPod", new Kubernetes.Types.Inputs.Core.V1.PodArgs
        {
            ApiVersion = "v1",
            Kind = "Pod",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Namespace = "foo",
                Name = "bar",
            },
            Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
            {
                Containers = 
                {
                    new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                    {
                        Name = "nginx",
                        Image = "nginx:1.14-alpine",
                        Resources = new Kubernetes.Types.Inputs.Core.V1.ResourceRequirementsArgs
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

### Java

```console
kube2pulumi java -f ./pod.yaml
```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.kubernetes.core_v1.Pod;
import com.pulumi.kubernetes.core_v1.PodArgs;
import com.pulumi.kubernetes.meta_v1.inputs.ObjectMetaArgs;
import com.pulumi.kubernetes.core_v1.inputs.PodSpecArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var fooBarPod = new Pod("fooBarPod", PodArgs.builder()
                .apiVersion("v1")
                .kind("Pod")
                .metadata(ObjectMetaArgs.builder()
                        .namespace("foo")
                        .name("bar")
                        .build())
                .spec(PodSpecArgs.builder()
                        .containers(ContainerArgs.builder()
                                .name("nginx")
                                .image("nginx:1.14-alpine")
                                .resources(ResourceRequirementsArgs.builder()
                                        .limits(Map.ofEntries(
                                                Map.entry("memory", "20Mi"),
                                                Map.entry("cpu", 0.2)
                                        ))
                                        .build())
                                .build())
                        .build())
                .build());

    }
}
```

# Limitations

`kube2pulumi` currently does not handle the conversion of CustomResourceDefinitions or CustomResources. However, our 
new tool `crd2pulumi`, creates strongly-typed args for a Resource based on your CRD! If using CRD/CR's make sure to check out the following tool!

1. [crd2pulumi README](https://github.com/pulumi/crd2pulumi)

