package main

import (
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appsv1.NewDaemonSet(ctx, "defaultCiliumDaemonSet", &appsv1.DaemonSetArgs{
			ApiVersion: pulumi.String("apps/v1"),
			Kind:       pulumi.String("DaemonSet"),
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String("cilium"),
				Namespace: pulumi.String("default"),
				Labels: pulumi.StringMap{
					"k8s-app":                   pulumi.String("cilium"),
					"app.kubernetes.io/part-of": pulumi.String("cilium"),
					"app.kubernetes.io/name":    pulumi.String("cilium-agent"),
				},
			},
			Spec: &appsv1.DaemonSetSpecArgs{
				Selector: &metav1.LabelSelectorArgs{
					MatchLabels: pulumi.StringMap{
						"k8s-app": pulumi.String("cilium"),
					},
				},
				UpdateStrategy: &appsv1.DaemonSetUpdateStrategyArgs{
					RollingUpdate: &appsv1.RollingUpdateDaemonSetArgs{
						MaxUnavailable: pulumi.Any(2),
					},
					Type: pulumi.String("RollingUpdate"),
				},
				Template: &corev1.PodTemplateSpecArgs{
					Metadata: &metav1.ObjectMetaArgs{
						Annotations: pulumi.StringMap{
							"container.apparmor.security.beta.kubernetes.io/cilium-agent":            pulumi.String("unconfined"),
							"container.apparmor.security.beta.kubernetes.io/clean-cilium-state":      pulumi.String("unconfined"),
							"container.apparmor.security.beta.kubernetes.io/mount-cgroup":            pulumi.String("unconfined"),
							"container.apparmor.security.beta.kubernetes.io/apply-sysctl-overwrites": pulumi.String("unconfined"),
						},
						Labels: pulumi.StringMap{
							"k8s-app":                   pulumi.String("cilium"),
							"app.kubernetes.io/name":    pulumi.String("cilium-agent"),
							"app.kubernetes.io/part-of": pulumi.String("cilium"),
						},
					},
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							&corev1.ContainerArgs{
								Name:            pulumi.String("cilium-agent"),
								Image:           pulumi.String("quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35"),
								ImagePullPolicy: pulumi.String("IfNotPresent"),
								Command: pulumi.StringArray{
									pulumi.String("cilium-agent"),
								},
								Args: pulumi.StringArray{
									pulumi.String("--config-dir=/tmp/cilium/config-map"),
								},
								StartupProbe: &corev1.ProbeArgs{
									HttpGet: &corev1.HTTPGetActionArgs{
										Host:   pulumi.String("127.0.0.1"),
										Path:   pulumi.String("/healthz"),
										Port:   pulumi.Any(9879),
										Scheme: pulumi.String("HTTP"),
										HttpHeaders: corev1.HTTPHeaderArray{
											&corev1.HTTPHeaderArgs{
												Name:  pulumi.String("brief"),
												Value: pulumi.String("true"),
											},
										},
									},
									FailureThreshold: pulumi.Int(105),
									PeriodSeconds:    pulumi.Int(2),
									SuccessThreshold: pulumi.Int(1),
								},
								LivenessProbe: &corev1.ProbeArgs{
									HttpGet: &corev1.HTTPGetActionArgs{
										Host:   pulumi.String("127.0.0.1"),
										Path:   pulumi.String("/healthz"),
										Port:   pulumi.Any(9879),
										Scheme: pulumi.String("HTTP"),
										HttpHeaders: corev1.HTTPHeaderArray{
											&corev1.HTTPHeaderArgs{
												Name:  pulumi.String("brief"),
												Value: pulumi.String("true"),
											},
										},
									},
									PeriodSeconds:    pulumi.Int(30),
									SuccessThreshold: pulumi.Int(1),
									FailureThreshold: pulumi.Int(10),
									TimeoutSeconds:   pulumi.Int(5),
								},
								ReadinessProbe: &corev1.ProbeArgs{
									HttpGet: &corev1.HTTPGetActionArgs{
										Host:   pulumi.String("127.0.0.1"),
										Path:   pulumi.String("/healthz"),
										Port:   pulumi.Any(9879),
										Scheme: pulumi.String("HTTP"),
										HttpHeaders: corev1.HTTPHeaderArray{
											&corev1.HTTPHeaderArgs{
												Name:  pulumi.String("brief"),
												Value: pulumi.String("true"),
											},
										},
									},
									PeriodSeconds:    pulumi.Int(30),
									SuccessThreshold: pulumi.Int(1),
									FailureThreshold: pulumi.Int(3),
									TimeoutSeconds:   pulumi.Int(5),
								},
								Env: corev1.EnvVarArray{
									&corev1.EnvVarArgs{
										Name: pulumi.String("K8S_NODE_NAME"),
										ValueFrom: &corev1.EnvVarSourceArgs{
											FieldRef: &corev1.ObjectFieldSelectorArgs{
												ApiVersion: pulumi.String("v1"),
												FieldPath:  pulumi.String("spec.nodeName"),
											},
										},
									},
									&corev1.EnvVarArgs{
										Name: pulumi.String("CILIUM_K8S_NAMESPACE"),
										ValueFrom: &corev1.EnvVarSourceArgs{
											FieldRef: &corev1.ObjectFieldSelectorArgs{
												ApiVersion: pulumi.String("v1"),
												FieldPath:  pulumi.String("metadata.namespace"),
											},
										},
									},
									&corev1.EnvVarArgs{
										Name:  pulumi.String("CILIUM_CLUSTERMESH_CONFIG"),
										Value: pulumi.String("/var/lib/cilium/clustermesh/"),
									},
								},
								Lifecycle: &corev1.LifecycleArgs{
									PreStop: &corev1.LifecycleHandlerArgs{
										Exec: &corev1.ExecActionArgs{
											Command: pulumi.StringArray{
												pulumi.String("/cni-uninstall.sh"),
											},
										},
									},
								},
								SecurityContext: &corev1.SecurityContextArgs{
									SeLinuxOptions: &corev1.SELinuxOptionsArgs{
										Level: pulumi.String("s0"),
										Type:  pulumi.String("spc_t"),
									},
									Capabilities: &corev1.CapabilitiesArgs{
										Add: pulumi.StringArray{
											pulumi.String("CHOWN"),
											pulumi.String("KILL"),
											pulumi.String("NET_ADMIN"),
											pulumi.String("NET_RAW"),
											pulumi.String("IPC_LOCK"),
											pulumi.String("SYS_MODULE"),
											pulumi.String("SYS_ADMIN"),
											pulumi.String("SYS_RESOURCE"),
											pulumi.String("DAC_OVERRIDE"),
											pulumi.String("FOWNER"),
											pulumi.String("SETGID"),
											pulumi.String("SETUID"),
										},
										Drop: pulumi.StringArray{
											pulumi.String("ALL"),
										},
									},
								},
								TerminationMessagePolicy: pulumi.String("FallbackToLogsOnError"),
								VolumeMounts: corev1.VolumeMountArray{
									&corev1.VolumeMountArgs{
										MountPath: pulumi.String("/host/proc/sys/net"),
										Name:      pulumi.String("host-proc-sys-net"),
									},
									&corev1.VolumeMountArgs{
										MountPath: pulumi.String("/host/proc/sys/kernel"),
										Name:      pulumi.String("host-proc-sys-kernel"),
									},
									&corev1.VolumeMountArgs{
										Name:             pulumi.String("bpf-maps"),
										MountPath:        pulumi.String("/sys/fs/bpf"),
										MountPropagation: pulumi.String("HostToContainer"),
									},
									&corev1.VolumeMountArgs{
										Name:      pulumi.String("cilium-run"),
										MountPath: pulumi.String("/var/run/cilium"),
									},
									&corev1.VolumeMountArgs{
										Name:      pulumi.String("etc-cni-netd"),
										MountPath: pulumi.String("/host/etc/cni/net.d"),
									},
									&corev1.VolumeMountArgs{
										Name:      pulumi.String("clustermesh-secrets"),
										MountPath: pulumi.String("/var/lib/cilium/clustermesh"),
										ReadOnly:  pulumi.Bool(true),
									},
									&corev1.VolumeMountArgs{
										Name:      pulumi.String("lib-modules"),
										MountPath: pulumi.String("/lib/modules"),
										ReadOnly:  pulumi.Bool(true),
									},
									&corev1.VolumeMountArgs{
										Name:      pulumi.String("xtables-lock"),
										MountPath: pulumi.String("/run/xtables.lock"),
									},
									&corev1.VolumeMountArgs{
										Name:      pulumi.String("hubble-tls"),
										MountPath: pulumi.String("/var/lib/cilium/tls/hubble"),
										ReadOnly:  pulumi.Bool(true),
									},
									&corev1.VolumeMountArgs{
										Name:      pulumi.String("tmp"),
										MountPath: pulumi.String("/tmp"),
									},
								},
							},
						},
						InitContainers: corev1.ContainerArray{
							&corev1.ContainerArgs{
								Name:            pulumi.String("config"),
								Image:           pulumi.String("quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35"),
								ImagePullPolicy: pulumi.String("IfNotPresent"),
								Command: pulumi.StringArray{
									pulumi.String("cilium"),
									pulumi.String("build-config"),
								},
								Env: corev1.EnvVarArray{
									&corev1.EnvVarArgs{
										Name: pulumi.String("K8S_NODE_NAME"),
										ValueFrom: &corev1.EnvVarSourceArgs{
											FieldRef: &corev1.ObjectFieldSelectorArgs{
												ApiVersion: pulumi.String("v1"),
												FieldPath:  pulumi.String("spec.nodeName"),
											},
										},
									},
									&corev1.EnvVarArgs{
										Name: pulumi.String("CILIUM_K8S_NAMESPACE"),
										ValueFrom: &corev1.EnvVarSourceArgs{
											FieldRef: &corev1.ObjectFieldSelectorArgs{
												ApiVersion: pulumi.String("v1"),
												FieldPath:  pulumi.String("metadata.namespace"),
											},
										},
									},
								},
								VolumeMounts: corev1.VolumeMountArray{
									&corev1.VolumeMountArgs{
										Name:      pulumi.String("tmp"),
										MountPath: pulumi.String("/tmp"),
									},
								},
								TerminationMessagePolicy: pulumi.String("FallbackToLogsOnError"),
							},
							&corev1.ContainerArgs{
								Name:            pulumi.String("mount-cgroup"),
								Image:           pulumi.String("quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35"),
								ImagePullPolicy: pulumi.String("IfNotPresent"),
								Env: corev1.EnvVarArray{
									&corev1.EnvVarArgs{
										Name:  pulumi.String("CGROUP_ROOT"),
										Value: pulumi.String("/run/cilium/cgroupv2"),
									},
									&corev1.EnvVarArgs{
										Name:  pulumi.String("BIN_PATH"),
										Value: pulumi.String("/opt/cni/bin"),
									},
								},
								Command: pulumi.StringArray{
									pulumi.String("sh"),
									pulumi.String("-ec"),
									pulumi.String("cp /usr/bin/cilium-mount /hostbin/cilium-mount;\nnsenter --cgroup=/hostproc/1/ns/cgroup --mount=/hostproc/1/ns/mnt \"${BIN_PATH}/cilium-mount\" $CGROUP_ROOT;\nrm /hostbin/cilium-mount\n"),
								},
								VolumeMounts: corev1.VolumeMountArray{
									&corev1.VolumeMountArgs{
										Name:      pulumi.String("hostproc"),
										MountPath: pulumi.String("/hostproc"),
									},
									&corev1.VolumeMountArgs{
										Name:      pulumi.String("cni-path"),
										MountPath: pulumi.String("/hostbin"),
									},
								},
								TerminationMessagePolicy: pulumi.String("FallbackToLogsOnError"),
								SecurityContext: &corev1.SecurityContextArgs{
									SeLinuxOptions: &corev1.SELinuxOptionsArgs{
										Level: pulumi.String("s0"),
										Type:  pulumi.String("spc_t"),
									},
									Capabilities: &corev1.CapabilitiesArgs{
										Add: pulumi.StringArray{
											pulumi.String("SYS_ADMIN"),
											pulumi.String("SYS_CHROOT"),
											pulumi.String("SYS_PTRACE"),
										},
										Drop: pulumi.StringArray{
											pulumi.String("ALL"),
										},
									},
								},
							},
							&corev1.ContainerArgs{
								Name:            pulumi.String("apply-sysctl-overwrites"),
								Image:           pulumi.String("quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35"),
								ImagePullPolicy: pulumi.String("IfNotPresent"),
								Env: corev1.EnvVarArray{
									&corev1.EnvVarArgs{
										Name:  pulumi.String("BIN_PATH"),
										Value: pulumi.String("/opt/cni/bin"),
									},
								},
								Command: pulumi.StringArray{
									pulumi.String("sh"),
									pulumi.String("-ec"),
									pulumi.String("cp /usr/bin/cilium-sysctlfix /hostbin/cilium-sysctlfix;\nnsenter --mount=/hostproc/1/ns/mnt \"${BIN_PATH}/cilium-sysctlfix\";\nrm /hostbin/cilium-sysctlfix\n"),
								},
								VolumeMounts: corev1.VolumeMountArray{
									&corev1.VolumeMountArgs{
										Name:      pulumi.String("hostproc"),
										MountPath: pulumi.String("/hostproc"),
									},
									&corev1.VolumeMountArgs{
										Name:      pulumi.String("cni-path"),
										MountPath: pulumi.String("/hostbin"),
									},
								},
								TerminationMessagePolicy: pulumi.String("FallbackToLogsOnError"),
								SecurityContext: &corev1.SecurityContextArgs{
									SeLinuxOptions: &corev1.SELinuxOptionsArgs{
										Level: pulumi.String("s0"),
										Type:  pulumi.String("spc_t"),
									},
									Capabilities: &corev1.CapabilitiesArgs{
										Add: pulumi.StringArray{
											pulumi.String("SYS_ADMIN"),
											pulumi.String("SYS_CHROOT"),
											pulumi.String("SYS_PTRACE"),
										},
										Drop: pulumi.StringArray{
											pulumi.String("ALL"),
										},
									},
								},
							},
							&corev1.ContainerArgs{
								Name:            pulumi.String("mount-bpf-fs"),
								Image:           pulumi.String("quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35"),
								ImagePullPolicy: pulumi.String("IfNotPresent"),
								Args: pulumi.StringArray{
									pulumi.String("mount | grep \"/sys/fs/bpf type bpf\" || mount -t bpf bpf /sys/fs/bpf"),
								},
								Command: pulumi.StringArray{
									pulumi.String("/bin/bash"),
									pulumi.String("-c"),
									pulumi.String("--"),
								},
								TerminationMessagePolicy: pulumi.String("FallbackToLogsOnError"),
								SecurityContext: &corev1.SecurityContextArgs{
									Privileged: pulumi.Bool(true),
								},
								VolumeMounts: corev1.VolumeMountArray{
									&corev1.VolumeMountArgs{
										Name:             pulumi.String("bpf-maps"),
										MountPath:        pulumi.String("/sys/fs/bpf"),
										MountPropagation: pulumi.String("Bidirectional"),
									},
								},
							},
							&corev1.ContainerArgs{
								Name:            pulumi.String("clean-cilium-state"),
								Image:           pulumi.String("quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35"),
								ImagePullPolicy: pulumi.String("IfNotPresent"),
								Command: pulumi.StringArray{
									pulumi.String("/init-container.sh"),
								},
								Env: corev1.EnvVarArray{
									&corev1.EnvVarArgs{
										Name: pulumi.String("CILIUM_ALL_STATE"),
										ValueFrom: &corev1.EnvVarSourceArgs{
											ConfigMapKeyRef: &corev1.ConfigMapKeySelectorArgs{
												Name:     pulumi.String("cilium-config"),
												Key:      pulumi.String("clean-cilium-state"),
												Optional: pulumi.Bool(true),
											},
										},
									},
									&corev1.EnvVarArgs{
										Name: pulumi.String("CILIUM_BPF_STATE"),
										ValueFrom: &corev1.EnvVarSourceArgs{
											ConfigMapKeyRef: &corev1.ConfigMapKeySelectorArgs{
												Name:     pulumi.String("cilium-config"),
												Key:      pulumi.String("clean-cilium-bpf-state"),
												Optional: pulumi.Bool(true),
											},
										},
									},
								},
								TerminationMessagePolicy: pulumi.String("FallbackToLogsOnError"),
								SecurityContext: &corev1.SecurityContextArgs{
									SeLinuxOptions: &corev1.SELinuxOptionsArgs{
										Level: pulumi.String("s0"),
										Type:  pulumi.String("spc_t"),
									},
									Capabilities: &corev1.CapabilitiesArgs{
										Add: pulumi.StringArray{
											pulumi.String("NET_ADMIN"),
											pulumi.String("SYS_MODULE"),
											pulumi.String("SYS_ADMIN"),
											pulumi.String("SYS_RESOURCE"),
										},
										Drop: pulumi.StringArray{
											pulumi.String("ALL"),
										},
									},
								},
								VolumeMounts: corev1.VolumeMountArray{
									&corev1.VolumeMountArgs{
										Name:      pulumi.String("bpf-maps"),
										MountPath: pulumi.String("/sys/fs/bpf"),
									},
									&corev1.VolumeMountArgs{
										Name:             pulumi.String("cilium-cgroup"),
										MountPath:        pulumi.String("/run/cilium/cgroupv2"),
										MountPropagation: pulumi.String("HostToContainer"),
									},
									&corev1.VolumeMountArgs{
										Name:      pulumi.String("cilium-run"),
										MountPath: pulumi.String("/var/run/cilium"),
									},
								},
								Resources: &corev1.ResourceRequirementsArgs{
									Requests: pulumi.StringMap{
										"cpu":    pulumi.String("100m"),
										"memory": pulumi.String("100Mi"),
									},
								},
							},
							&corev1.ContainerArgs{
								Name:            pulumi.String("install-cni-binaries"),
								Image:           pulumi.String("quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35"),
								ImagePullPolicy: pulumi.String("IfNotPresent"),
								Command: pulumi.StringArray{
									pulumi.String("/install-plugin.sh"),
								},
								Resources: &corev1.ResourceRequirementsArgs{
									Requests: pulumi.StringMap{
										"cpu":    pulumi.String("100m"),
										"memory": pulumi.String("10Mi"),
									},
								},
								SecurityContext: &corev1.SecurityContextArgs{
									SeLinuxOptions: &corev1.SELinuxOptionsArgs{
										Level: pulumi.String("s0"),
										Type:  pulumi.String("spc_t"),
									},
									Capabilities: &corev1.CapabilitiesArgs{
										Drop: pulumi.StringArray{
											pulumi.String("ALL"),
										},
									},
								},
								TerminationMessagePolicy: pulumi.String("FallbackToLogsOnError"),
								VolumeMounts: corev1.VolumeMountArray{
									&corev1.VolumeMountArgs{
										Name:      pulumi.String("cni-path"),
										MountPath: pulumi.String("/host/opt/cni/bin"),
									},
								},
							},
						},
						RestartPolicy:                 pulumi.String("Always"),
						PriorityClassName:             pulumi.String("system-node-critical"),
						ServiceAccount:                pulumi.String("cilium"),
						ServiceAccountName:            pulumi.String("cilium"),
						AutomountServiceAccountToken:  pulumi.Bool(true),
						TerminationGracePeriodSeconds: pulumi.Int(1),
						HostNetwork:                   pulumi.Bool(true),
						Affinity: &corev1.AffinityArgs{
							PodAntiAffinity: &corev1.PodAntiAffinityArgs{
								RequiredDuringSchedulingIgnoredDuringExecution: corev1.PodAffinityTermArray{
									&corev1.PodAffinityTermArgs{
										LabelSelector: &metav1.LabelSelectorArgs{
											MatchLabels: pulumi.StringMap{
												"k8s-app": pulumi.String("cilium"),
											},
										},
										TopologyKey: pulumi.String("kubernetes.io/hostname"),
									},
								},
							},
						},
						NodeSelector: pulumi.StringMap{
							"kubernetes.io/os": pulumi.String("linux"),
						},
						Tolerations: corev1.TolerationArray{
							&corev1.TolerationArgs{
								Operator: pulumi.String("Exists"),
							},
						},
						Volumes: corev1.VolumeArray{
							&corev1.VolumeArgs{
								Name:     pulumi.String("tmp"),
								EmptyDir: nil,
							},
							&corev1.VolumeArgs{
								Name: pulumi.String("cilium-run"),
								HostPath: &corev1.HostPathVolumeSourceArgs{
									Path: pulumi.String("/var/run/cilium"),
									Type: pulumi.String("DirectoryOrCreate"),
								},
							},
							&corev1.VolumeArgs{
								Name: pulumi.String("bpf-maps"),
								HostPath: &corev1.HostPathVolumeSourceArgs{
									Path: pulumi.String("/sys/fs/bpf"),
									Type: pulumi.String("DirectoryOrCreate"),
								},
							},
							&corev1.VolumeArgs{
								Name: pulumi.String("hostproc"),
								HostPath: &corev1.HostPathVolumeSourceArgs{
									Path: pulumi.String("/proc"),
									Type: pulumi.String("Directory"),
								},
							},
							&corev1.VolumeArgs{
								Name: pulumi.String("cilium-cgroup"),
								HostPath: &corev1.HostPathVolumeSourceArgs{
									Path: pulumi.String("/run/cilium/cgroupv2"),
									Type: pulumi.String("DirectoryOrCreate"),
								},
							},
							&corev1.VolumeArgs{
								Name: pulumi.String("cni-path"),
								HostPath: &corev1.HostPathVolumeSourceArgs{
									Path: pulumi.String("/opt/cni/bin"),
									Type: pulumi.String("DirectoryOrCreate"),
								},
							},
							&corev1.VolumeArgs{
								Name: pulumi.String("etc-cni-netd"),
								HostPath: &corev1.HostPathVolumeSourceArgs{
									Path: pulumi.String("/etc/cni/net.d"),
									Type: pulumi.String("DirectoryOrCreate"),
								},
							},
							&corev1.VolumeArgs{
								Name: pulumi.String("lib-modules"),
								HostPath: &corev1.HostPathVolumeSourceArgs{
									Path: pulumi.String("/lib/modules"),
								},
							},
							&corev1.VolumeArgs{
								Name: pulumi.String("xtables-lock"),
								HostPath: &corev1.HostPathVolumeSourceArgs{
									Path: pulumi.String("/run/xtables.lock"),
									Type: pulumi.String("FileOrCreate"),
								},
							},
							&corev1.VolumeArgs{
								Name: pulumi.String("clustermesh-secrets"),
								Projected: &corev1.ProjectedVolumeSourceArgs{
									DefaultMode: pulumi.Int(256),
									Sources: corev1.VolumeProjectionArray{
										&corev1.VolumeProjectionArgs{
											Secret: &corev1.SecretProjectionArgs{
												Name:     pulumi.String("cilium-clustermesh"),
												Optional: pulumi.Bool(true),
											},
										},
										&corev1.VolumeProjectionArgs{
											Secret: &corev1.SecretProjectionArgs{
												Name:     pulumi.String("clustermesh-apiserver-remote-cert"),
												Optional: pulumi.Bool(true),
												Items: corev1.KeyToPathArray{
													&corev1.KeyToPathArgs{
														Key:  pulumi.String("tls.key"),
														Path: pulumi.String("common-etcd-client.key"),
													},
													&corev1.KeyToPathArgs{
														Key:  pulumi.String("tls.crt"),
														Path: pulumi.String("common-etcd-client.crt"),
													},
													&corev1.KeyToPathArgs{
														Key:  pulumi.String("ca.crt"),
														Path: pulumi.String("common-etcd-client-ca.crt"),
													},
												},
											},
										},
									},
								},
							},
							&corev1.VolumeArgs{
								Name: pulumi.String("host-proc-sys-net"),
								HostPath: &corev1.HostPathVolumeSourceArgs{
									Path: pulumi.String("/proc/sys/net"),
									Type: pulumi.String("Directory"),
								},
							},
							&corev1.VolumeArgs{
								Name: pulumi.String("host-proc-sys-kernel"),
								HostPath: &corev1.HostPathVolumeSourceArgs{
									Path: pulumi.String("/proc/sys/kernel"),
									Type: pulumi.String("Directory"),
								},
							},
							&corev1.VolumeArgs{
								Name: pulumi.String("hubble-tls"),
								Projected: &corev1.ProjectedVolumeSourceArgs{
									DefaultMode: pulumi.Int(256),
									Sources: corev1.VolumeProjectionArray{
										&corev1.VolumeProjectionArgs{
											Secret: &corev1.SecretProjectionArgs{
												Name:     pulumi.String("hubble-server-certs"),
												Optional: pulumi.Bool(true),
												Items: corev1.KeyToPathArray{
													&corev1.KeyToPathArgs{
														Key:  pulumi.String("tls.crt"),
														Path: pulumi.String("server.crt"),
													},
													&corev1.KeyToPathArgs{
														Key:  pulumi.String("tls.key"),
														Path: pulumi.String("server.key"),
													},
													&corev1.KeyToPathArgs{
														Key:  pulumi.String("ca.crt"),
														Path: pulumi.String("client-ca.crt"),
													},
												},
											},
										},
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
