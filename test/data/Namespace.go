package main

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
