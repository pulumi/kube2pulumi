package main

import (
	"fmt"

	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := corev1.NewConfigMap(ctx, "kube_systemCorednsConfigMap", &corev1.ConfigMapArgs{
			ApiVersion: pulumi.String("v1"),
			Kind:       pulumi.String("ConfigMap"),
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String("coredns"),
				Namespace: pulumi.String("kube-system"),
			},
			Data: pulumi.StringMap{
				"Corefile": pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", ".:53 {\n", "        errors\n", "        health {\n", "          lameduck 5s\n", "        }\n", "        ready\n", "        kubernetes CLUSTER_DOMAIN REVERSE_CIDRS {\n", "          fallthrough in-addr.arpa ip6.arpa\n", "        }\n", "        prometheus :9153\n", "        forward . UPSTREAMNAMESERVER {\n", "          max_concurrent 1000\n", "        }\n", "        cache 30\n", "        loop\n", "        reload\n", "        loadbalance\n", "    }STUBDOMAINS\n")),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
