package main

import (
	"fmt"

	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := corev1.NewConfigMap(ctx, "myappConfigMap", &corev1.ConfigMapArgs{
			ApiVersion: pulumi.String("v1"),
			Kind:       pulumi.String("ConfigMap"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("myapp"),
			},
			Data: pulumi.StringMap{
				"key": pulumi.String(fmt.Sprintf("{\\\"uid\\\": \\\"$(datasource)\\\"}")),
			},
		})
		if err != nil {
			return err
		}
		_, err = corev1.NewConfigMap(ctx, "myapp_varConfigMap", &corev1.ConfigMapArgs{
			ApiVersion: pulumi.String("v1"),
			Kind:       pulumi.String("ConfigMap"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("myapp-var"),
			},
			Data: pulumi.StringMap{
				"key": pulumi.String(fmt.Sprintf("{\\\"uid\\\": \\\"${datasource}\\\"}")),
			},
		})
		if err != nil {
			return err
		}
		_, err = corev1.NewConfigMap(ctx, "myapp_no_end_bracketConfigMap", &corev1.ConfigMapArgs{
			ApiVersion: pulumi.String("v1"),
			Kind:       pulumi.String("ConfigMap"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("myapp-no-end-bracket"),
			},
			Data: pulumi.StringMap{
				"key": pulumi.String(fmt.Sprintf("{\\\"uid\\\": \\\"${datasource\\\"}")),
			},
		})
		if err != nil {
			return err
		}
		_, err = corev1.NewConfigMap(ctx, "myapp_no_bracketsConfigMap", &corev1.ConfigMapArgs{
			ApiVersion: pulumi.String("v1"),
			Kind:       pulumi.String("ConfigMap"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("myapp-no-brackets"),
			},
			Data: pulumi.StringMap{
				"key": pulumi.String(fmt.Sprintf("{\\\"uid\\\": \\\"${datasource\\\"")),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
