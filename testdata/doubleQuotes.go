package main

import (
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appsv1.NewDeployment(ctx, "argocd_serverDeployment", &appsv1.DeploymentArgs{
			ApiVersion: pulumi.String("apps/v1"),
			Kind:       pulumi.String("Deployment"),
			Metadata: &metav1.ObjectMetaArgs{
				Labels: pulumi.StringMap{
					"app.kubernetes.io/component":  pulumi.String("server"),
					"app.kubernetes.io/instance":   pulumi.String("argocd"),
					"app.kubernetes.io/managed-by": pulumi.String("pulumi"),
					"app.kubernetes.io/name":       pulumi.String("argocd-server"),
					"app.kubernetes.io/part-of":    pulumi.String("argocd"),
					"app.kubernetes.io/version":    pulumi.String("v1.6.1"),
					"helm.sh/chart":                pulumi.String("argo-cd-2.5.4"),
				},
				Name: pulumi.String("argocd-server"),
			},
			Spec: &appsv1.DeploymentSpecArgs{
				Template: &corev1.PodTemplateSpecArgs{
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							&corev1.ContainerArgs{
								ReadinessProbe: &corev1.ProbeArgs{
									HttpGet: &corev1.HTTPGetActionArgs{
										Port: pulumi.Any(8080),
									},
								},
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
