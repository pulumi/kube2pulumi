package main

import (
	"fmt"

	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appsv1.NewDeployment(ctx, "defaultArgocd_serverDeployment", &appsv1.DeploymentArgs{
			ApiVersion: pulumi.String("apps/v1"),
			Kind:       pulumi.String("Deployment"),
			Metadata: &metav1.ObjectMetaArgs{
				Annotations: pulumi.StringMap{
					"deployment.kubernetes.io/revision":                pulumi.String("1"),
					"kubectl.kubernetes.io/last-applied-configuration": pulumi.String("{\"apiVersion\":\"apps/v1\",\"kind\":\"Deployment\",\"metadata\":{\"labels\":{\"app.kubernetes.io/component\":\"server\",\"app.kubernetes.io/instance\":\"argocd\",\"app.kubernetes.io/managed-by\":\"pulumi\",\"app.kubernetes.io/name\":\"argocd-server\",\"app.kubernetes.io/part-of\":\"argocd\",\"app.kubernetes.io/version\":\"v1.6.1\",\"helm.sh/chart\":\"argo-cd-2.5.4\"},\"name\":\"argocd-server\",\"namespace\":\"default\"},\"spec\":{\"replicas\":1,\"revisionHistoryLimit\":5,\"selector\":{\"matchLabels\":{\"app.kubernetes.io/instance\":\"argocd\",\"app.kubernetes.io/name\":\"argocd-server\"}},\"template\":{\"metadata\":{\"labels\":{\"app.kubernetes.io/component\":\"server\",\"app.kubernetes.io/instance\":\"argocd\",\"app.kubernetes.io/managed-by\":\"Helm\",\"app.kubernetes.io/name\":\"argocd-server\",\"app.kubernetes.io/part-of\":\"argocd\",\"app.kubernetes.io/version\":\"v1.6.1\",\"helm.sh/chart\":\"argo-cd-2.5.4\"}},\"spec\":{\"containers\":[{\"command\":[\"argocd-server\",\"--staticassets\",\"/shared/app\",\"--repo-server\",\"argocd-repo-server:8081\",\"--dex-server\",\"http://argocd-dex-server:5556\",\"--loglevel\",\"info\",\"--redis\",\"argocd-redis:6379\"],\"image\":\"argoproj/argocd:v1.6.1\",\"imagePullPolicy\":\"IfNotPresent\",\"livenessProbe\":{\"failureThreshold\":3,\"httpGet\":{\"path\":\"/healthz\",\"port\":8080},\"initialDelaySeconds\":10,\"periodSeconds\":10,\"successThreshold\":1,\"timeoutSeconds\":1},\"name\":\"server\",\"ports\":[{\"containerPort\":8080,\"name\":\"server\",\"protocol\":\"TCP\"}],\"readinessProbe\":{\"failureThreshold\":3,\"httpGet\":{\"path\":\"/healthz\",\"port\":8080},\"initialDelaySeconds\":10,\"periodSeconds\":10,\"successThreshold\":1,\"timeoutSeconds\":1},\"resources\":{},\"volumeMounts\":[{\"mountPath\":\"/app/config/ssh\",\"name\":\"ssh-known-hosts\"}]}],\"serviceAccountName\":\"argocd-server\",\"volumes\":[{\"emptyDir\":{},\"name\":\"static-files\"},{\"configMap\":{\"name\":\"argocd-ssh-known-hosts-cm\"},\"name\":\"ssh-known-hosts\"}]}}}}\n"),
				},
				CreationTimestamp: pulumi.String("2020-08-04T18:50:43Z"),
				Generation:        pulumi.Int(1),
				Labels: pulumi.StringMap{
					"app.kubernetes.io/component":  pulumi.String("server"),
					"app.kubernetes.io/instance":   pulumi.String("argocd"),
					"app.kubernetes.io/managed-by": pulumi.String("pulumi"),
					"app.kubernetes.io/name":       pulumi.String("argocd-server"),
					"app.kubernetes.io/part-of":    pulumi.String("argocd"),
					"app.kubernetes.io/version":    pulumi.String("v1.6.1"),
					"helm.sh/chart":                pulumi.String("argo-cd-2.5.4"),
				},
				Name:            pulumi.String("argocd-server"),
				Namespace:       pulumi.String("default"),
				ResourceVersion: pulumi.String("1406"),
				SelfLink:        pulumi.String("/apis/apps/v1/namespaces/default/deployments/argocd-server"),
				Uid:             pulumi.String("4b806e77-b035-41a3-bdf9-9781b76445f9"),
			},
			Spec: &appsv1.DeploymentSpecArgs{
				ProgressDeadlineSeconds: pulumi.Int(600),
				Replicas:                pulumi.Int(1),
				RevisionHistoryLimit:    pulumi.Int(5),
				Selector: &metav1.LabelSelectorArgs{
					MatchLabels: pulumi.StringMap{
						"app.kubernetes.io/instance": pulumi.String("argocd"),
						"app.kubernetes.io/name":     pulumi.String("argocd-server"),
					},
				},
				Strategy: &appsv1.DeploymentStrategyArgs{
					RollingUpdate: &appsv1.RollingUpdateDeploymentArgs{
						MaxSurge:       pulumi.Any(fmt.Sprintf("25%v", "%")),
						MaxUnavailable: pulumi.Any(fmt.Sprintf("25%v", "%")),
					},
					Type: pulumi.String("RollingUpdate"),
				},
				Template: &corev1.PodTemplateSpecArgs{
					Metadata: &metav1.ObjectMetaArgs{
						CreationTimestamp: nil,
						Labels: pulumi.StringMap{
							"app.kubernetes.io/component":  pulumi.String("server"),
							"app.kubernetes.io/instance":   pulumi.String("argocd"),
							"app.kubernetes.io/managed-by": pulumi.String("Helm"),
							"app.kubernetes.io/name":       pulumi.String("argocd-server"),
							"app.kubernetes.io/part-of":    pulumi.String("argocd"),
							"app.kubernetes.io/version":    pulumi.String("v1.6.1"),
							"helm.sh/chart":                pulumi.String("argo-cd-2.5.4"),
						},
					},
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							&corev1.ContainerArgs{
								Command: pulumi.StringArray{
									pulumi.String("argocd-server"),
									pulumi.String("--staticassets"),
									pulumi.String("/shared/app"),
									pulumi.String("--repo-server"),
									pulumi.String("argocd-repo-server:8081"),
									pulumi.String("--dex-server"),
									pulumi.String("http://argocd-dex-server:5556"),
									pulumi.String("--loglevel"),
									pulumi.String("info"),
									pulumi.String("--redis"),
									pulumi.String("argocd-redis:6379"),
								},
								Image:           pulumi.String("argoproj/argocd:v1.6.1"),
								ImagePullPolicy: pulumi.String("IfNotPresent"),
								LivenessProbe: &corev1.ProbeArgs{
									FailureThreshold: pulumi.Int(3),
									HttpGet: &corev1.HTTPGetActionArgs{
										Path:   pulumi.String("/healthz"),
										Port:   pulumi.Any(8080),
										Scheme: pulumi.String("HTTP"),
									},
									InitialDelaySeconds: pulumi.Int(10),
									PeriodSeconds:       pulumi.Int(10),
									SuccessThreshold:    pulumi.Int(1),
									TimeoutSeconds:      pulumi.Int(1),
								},
								Name: pulumi.String("server"),
								Ports: corev1.ContainerPortArray{
									&corev1.ContainerPortArgs{
										ContainerPort: pulumi.Int(8080),
										Name:          pulumi.String("server"),
										Protocol:      pulumi.String("TCP"),
									},
								},
								ReadinessProbe: &corev1.ProbeArgs{
									FailureThreshold: pulumi.Int(3),
									HttpGet: &corev1.HTTPGetActionArgs{
										Path:   pulumi.String("/healthz"),
										Port:   pulumi.Any(8080),
										Scheme: pulumi.String("HTTP"),
									},
									InitialDelaySeconds: pulumi.Int(10),
									PeriodSeconds:       pulumi.Int(10),
									SuccessThreshold:    pulumi.Int(1),
									TimeoutSeconds:      pulumi.Int(1),
								},
								Resources:                nil,
								TerminationMessagePath:   pulumi.String("/dev/termination-log"),
								TerminationMessagePolicy: pulumi.String("File"),
								VolumeMounts: corev1.VolumeMountArray{
									&corev1.VolumeMountArgs{
										MountPath: pulumi.String("/app/config/ssh"),
										Name:      pulumi.String("ssh-known-hosts"),
									},
								},
							},
						},
						DnsPolicy:                     pulumi.String("ClusterFirst"),
						RestartPolicy:                 pulumi.String("Always"),
						SchedulerName:                 pulumi.String("default-scheduler"),
						SecurityContext:               nil,
						ServiceAccount:                pulumi.String("argocd-server"),
						ServiceAccountName:            pulumi.String("argocd-server"),
						TerminationGracePeriodSeconds: pulumi.Int(30),
						Volumes: corev1.VolumeArray{
							&corev1.VolumeArgs{
								EmptyDir: nil,
								Name:     pulumi.String("static-files"),
							},
							&corev1.VolumeArgs{
								ConfigMap: &corev1.ConfigMapVolumeSourceArgs{
									DefaultMode: pulumi.Int(420),
									Name:        pulumi.String("argocd-ssh-known-hosts-cm"),
								},
								Name: pulumi.String("ssh-known-hosts"),
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
