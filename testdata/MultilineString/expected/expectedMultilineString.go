package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
				"Corefile": pulumi.String(`.:53 {
        errors
        health {
          lameduck 5s
        }
        ready
        kubernetes CLUSTER_DOMAIN REVERSE_CIDRS {
          fallthrough in-addr.arpa ip6.arpa
        }
        prometheus :9153
        forward . UPSTREAMNAMESERVER {
          max_concurrent 1000
        }
        cache 30
        loop
        reload
        loadbalance
    }STUBDOMAINS
`),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
