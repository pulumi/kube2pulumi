package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
