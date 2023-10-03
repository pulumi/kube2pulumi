using System.Collections.Generic;
using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

return await Deployment.RunAsync(() => 
{
    var defaultCiliumDaemonSet = new Kubernetes.Apps.V1.DaemonSet("defaultCiliumDaemonSet", new()
    {
        ApiVersion = "apps/v1",
        Kind = "DaemonSet",
        Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
        {
            Name = "cilium",
            Namespace = "default",
            Labels = 
            {
                { "k8s-app", "cilium" },
                { "app.kubernetes.io/part-of", "cilium" },
                { "app.kubernetes.io/name", "cilium-agent" },
            },
        },
        Spec = new Kubernetes.Types.Inputs.Apps.V1.DaemonSetSpecArgs
        {
            Selector = new Kubernetes.Types.Inputs.Meta.V1.LabelSelectorArgs
            {
                MatchLabels = 
                {
                    { "k8s-app", "cilium" },
                },
            },
            UpdateStrategy = new Kubernetes.Types.Inputs.Apps.V1.DaemonSetUpdateStrategyArgs
            {
                RollingUpdate = new Kubernetes.Types.Inputs.Apps.V1.RollingUpdateDaemonSetArgs
                {
                    MaxUnavailable = 2,
                },
                Type = "RollingUpdate",
            },
            Template = new Kubernetes.Types.Inputs.Core.V1.PodTemplateSpecArgs
            {
                Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
                {
                    Annotations = 
                    {
                        { "container.apparmor.security.beta.kubernetes.io/cilium-agent", "unconfined" },
                        { "container.apparmor.security.beta.kubernetes.io/clean-cilium-state", "unconfined" },
                        { "container.apparmor.security.beta.kubernetes.io/mount-cgroup", "unconfined" },
                        { "container.apparmor.security.beta.kubernetes.io/apply-sysctl-overwrites", "unconfined" },
                    },
                    Labels = 
                    {
                        { "k8s-app", "cilium" },
                        { "app.kubernetes.io/name", "cilium-agent" },
                        { "app.kubernetes.io/part-of", "cilium" },
                    },
                },
                Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
                {
                    Containers = new[]
                    {
                        new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                        {
                            Name = "cilium-agent",
                            Image = "quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35",
                            ImagePullPolicy = "IfNotPresent",
                            Command = new[]
                            {
                                "cilium-agent",
                            },
                            Args = new[]
                            {
                                "--config-dir=/tmp/cilium/config-map",
                            },
                            StartupProbe = new Kubernetes.Types.Inputs.Core.V1.ProbeArgs
                            {
                                HttpGet = new Kubernetes.Types.Inputs.Core.V1.HTTPGetActionArgs
                                {
                                    Host = "127.0.0.1",
                                    Path = "/healthz",
                                    Port = 9879,
                                    Scheme = "HTTP",
                                    HttpHeaders = new[]
                                    {
                                        new Kubernetes.Types.Inputs.Core.V1.HTTPHeaderArgs
                                        {
                                            Name = "brief",
                                            Value = "true",
                                        },
                                    },
                                },
                                FailureThreshold = 105,
                                PeriodSeconds = 2,
                                SuccessThreshold = 1,
                            },
                            LivenessProbe = new Kubernetes.Types.Inputs.Core.V1.ProbeArgs
                            {
                                HttpGet = new Kubernetes.Types.Inputs.Core.V1.HTTPGetActionArgs
                                {
                                    Host = "127.0.0.1",
                                    Path = "/healthz",
                                    Port = 9879,
                                    Scheme = "HTTP",
                                    HttpHeaders = new[]
                                    {
                                        new Kubernetes.Types.Inputs.Core.V1.HTTPHeaderArgs
                                        {
                                            Name = "brief",
                                            Value = "true",
                                        },
                                    },
                                },
                                PeriodSeconds = 30,
                                SuccessThreshold = 1,
                                FailureThreshold = 10,
                                TimeoutSeconds = 5,
                            },
                            ReadinessProbe = new Kubernetes.Types.Inputs.Core.V1.ProbeArgs
                            {
                                HttpGet = new Kubernetes.Types.Inputs.Core.V1.HTTPGetActionArgs
                                {
                                    Host = "127.0.0.1",
                                    Path = "/healthz",
                                    Port = 9879,
                                    Scheme = "HTTP",
                                    HttpHeaders = new[]
                                    {
                                        new Kubernetes.Types.Inputs.Core.V1.HTTPHeaderArgs
                                        {
                                            Name = "brief",
                                            Value = "true",
                                        },
                                    },
                                },
                                PeriodSeconds = 30,
                                SuccessThreshold = 1,
                                FailureThreshold = 3,
                                TimeoutSeconds = 5,
                            },
                            Env = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                {
                                    Name = "K8S_NODE_NAME",
                                    ValueFrom = new Kubernetes.Types.Inputs.Core.V1.EnvVarSourceArgs
                                    {
                                        FieldRef = new Kubernetes.Types.Inputs.Core.V1.ObjectFieldSelectorArgs
                                        {
                                            ApiVersion = "v1",
                                            FieldPath = "spec.nodeName",
                                        },
                                    },
                                },
                                new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                {
                                    Name = "CILIUM_K8S_NAMESPACE",
                                    ValueFrom = new Kubernetes.Types.Inputs.Core.V1.EnvVarSourceArgs
                                    {
                                        FieldRef = new Kubernetes.Types.Inputs.Core.V1.ObjectFieldSelectorArgs
                                        {
                                            ApiVersion = "v1",
                                            FieldPath = "metadata.namespace",
                                        },
                                    },
                                },
                                new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                {
                                    Name = "CILIUM_CLUSTERMESH_CONFIG",
                                    Value = "/var/lib/cilium/clustermesh/",
                                },
                            },
                            Lifecycle = new Kubernetes.Types.Inputs.Core.V1.LifecycleArgs
                            {
                                PreStop = new Kubernetes.Types.Inputs.Core.V1.LifecycleHandlerArgs
                                {
                                    Exec = new Kubernetes.Types.Inputs.Core.V1.ExecActionArgs
                                    {
                                        Command = new[]
                                        {
                                            "/cni-uninstall.sh",
                                        },
                                    },
                                },
                            },
                            SecurityContext = new Kubernetes.Types.Inputs.Core.V1.SecurityContextArgs
                            {
                                SeLinuxOptions = new Kubernetes.Types.Inputs.Core.V1.SELinuxOptionsArgs
                                {
                                    Level = "s0",
                                    Type = "spc_t",
                                },
                                Capabilities = new Kubernetes.Types.Inputs.Core.V1.CapabilitiesArgs
                                {
                                    Add = new[]
                                    {
                                        "CHOWN",
                                        "KILL",
                                        "NET_ADMIN",
                                        "NET_RAW",
                                        "IPC_LOCK",
                                        "SYS_MODULE",
                                        "SYS_ADMIN",
                                        "SYS_RESOURCE",
                                        "DAC_OVERRIDE",
                                        "FOWNER",
                                        "SETGID",
                                        "SETUID",
                                    },
                                    Drop = new[]
                                    {
                                        "ALL",
                                    },
                                },
                            },
                            TerminationMessagePolicy = "FallbackToLogsOnError",
                            VolumeMounts = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    MountPath = "/host/proc/sys/net",
                                    Name = "host-proc-sys-net",
                                },
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    MountPath = "/host/proc/sys/kernel",
                                    Name = "host-proc-sys-kernel",
                                },
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "bpf-maps",
                                    MountPath = "/sys/fs/bpf",
                                    MountPropagation = "HostToContainer",
                                },
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "cilium-run",
                                    MountPath = "/var/run/cilium",
                                },
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "etc-cni-netd",
                                    MountPath = "/host/etc/cni/net.d",
                                },
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "clustermesh-secrets",
                                    MountPath = "/var/lib/cilium/clustermesh",
                                    ReadOnly = true,
                                },
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "lib-modules",
                                    MountPath = "/lib/modules",
                                    ReadOnly = true,
                                },
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "xtables-lock",
                                    MountPath = "/run/xtables.lock",
                                },
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "hubble-tls",
                                    MountPath = "/var/lib/cilium/tls/hubble",
                                    ReadOnly = true,
                                },
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "tmp",
                                    MountPath = "/tmp",
                                },
                            },
                        },
                    },
                    InitContainers = new[]
                    {
                        new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                        {
                            Name = "config",
                            Image = "quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35",
                            ImagePullPolicy = "IfNotPresent",
                            Command = new[]
                            {
                                "cilium",
                                "build-config",
                            },
                            Env = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                {
                                    Name = "K8S_NODE_NAME",
                                    ValueFrom = new Kubernetes.Types.Inputs.Core.V1.EnvVarSourceArgs
                                    {
                                        FieldRef = new Kubernetes.Types.Inputs.Core.V1.ObjectFieldSelectorArgs
                                        {
                                            ApiVersion = "v1",
                                            FieldPath = "spec.nodeName",
                                        },
                                    },
                                },
                                new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                {
                                    Name = "CILIUM_K8S_NAMESPACE",
                                    ValueFrom = new Kubernetes.Types.Inputs.Core.V1.EnvVarSourceArgs
                                    {
                                        FieldRef = new Kubernetes.Types.Inputs.Core.V1.ObjectFieldSelectorArgs
                                        {
                                            ApiVersion = "v1",
                                            FieldPath = "metadata.namespace",
                                        },
                                    },
                                },
                            },
                            VolumeMounts = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "tmp",
                                    MountPath = "/tmp",
                                },
                            },
                            TerminationMessagePolicy = "FallbackToLogsOnError",
                        },
                        new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                        {
                            Name = "mount-cgroup",
                            Image = "quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35",
                            ImagePullPolicy = "IfNotPresent",
                            Env = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                {
                                    Name = "CGROUP_ROOT",
                                    Value = "/run/cilium/cgroupv2",
                                },
                                new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                {
                                    Name = "BIN_PATH",
                                    Value = "/opt/cni/bin",
                                },
                            },
                            Command = new[]
                            {
                                "sh",
                                "-ec",
                                @"cp /usr/bin/cilium-mount /hostbin/cilium-mount;
              nsenter --cgroup=/hostproc/1/ns/cgroup --mount=/hostproc/1/ns/mnt ""${BIN_PATH}/cilium-mount"" $CGROUP_ROOT;
              rm /hostbin/cilium-mount
",
                            },
                            VolumeMounts = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "hostproc",
                                    MountPath = "/hostproc",
                                },
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "cni-path",
                                    MountPath = "/hostbin",
                                },
                            },
                            TerminationMessagePolicy = "FallbackToLogsOnError",
                            SecurityContext = new Kubernetes.Types.Inputs.Core.V1.SecurityContextArgs
                            {
                                SeLinuxOptions = new Kubernetes.Types.Inputs.Core.V1.SELinuxOptionsArgs
                                {
                                    Level = "s0",
                                    Type = "spc_t",
                                },
                                Capabilities = new Kubernetes.Types.Inputs.Core.V1.CapabilitiesArgs
                                {
                                    Add = new[]
                                    {
                                        "SYS_ADMIN",
                                        "SYS_CHROOT",
                                        "SYS_PTRACE",
                                    },
                                    Drop = new[]
                                    {
                                        "ALL",
                                    },
                                },
                            },
                        },
                        new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                        {
                            Name = "apply-sysctl-overwrites",
                            Image = "quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35",
                            ImagePullPolicy = "IfNotPresent",
                            Env = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                {
                                    Name = "BIN_PATH",
                                    Value = "/opt/cni/bin",
                                },
                            },
                            Command = new[]
                            {
                                "sh",
                                "-ec",
                                @"cp /usr/bin/cilium-sysctlfix /hostbin/cilium-sysctlfix;
              nsenter --mount=/hostproc/1/ns/mnt ""${BIN_PATH}/cilium-sysctlfix"";
              rm /hostbin/cilium-sysctlfix
",
                            },
                            VolumeMounts = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "hostproc",
                                    MountPath = "/hostproc",
                                },
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "cni-path",
                                    MountPath = "/hostbin",
                                },
                            },
                            TerminationMessagePolicy = "FallbackToLogsOnError",
                            SecurityContext = new Kubernetes.Types.Inputs.Core.V1.SecurityContextArgs
                            {
                                SeLinuxOptions = new Kubernetes.Types.Inputs.Core.V1.SELinuxOptionsArgs
                                {
                                    Level = "s0",
                                    Type = "spc_t",
                                },
                                Capabilities = new Kubernetes.Types.Inputs.Core.V1.CapabilitiesArgs
                                {
                                    Add = new[]
                                    {
                                        "SYS_ADMIN",
                                        "SYS_CHROOT",
                                        "SYS_PTRACE",
                                    },
                                    Drop = new[]
                                    {
                                        "ALL",
                                    },
                                },
                            },
                        },
                        new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                        {
                            Name = "mount-bpf-fs",
                            Image = "quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35",
                            ImagePullPolicy = "IfNotPresent",
                            Args = new[]
                            {
                                "mount | grep \"/sys/fs/bpf type bpf\" || mount -t bpf bpf /sys/fs/bpf",
                            },
                            Command = new[]
                            {
                                "/bin/bash",
                                "-c",
                                "--",
                            },
                            TerminationMessagePolicy = "FallbackToLogsOnError",
                            SecurityContext = new Kubernetes.Types.Inputs.Core.V1.SecurityContextArgs
                            {
                                Privileged = true,
                            },
                            VolumeMounts = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "bpf-maps",
                                    MountPath = "/sys/fs/bpf",
                                    MountPropagation = "Bidirectional",
                                },
                            },
                        },
                        new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                        {
                            Name = "clean-cilium-state",
                            Image = "quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35",
                            ImagePullPolicy = "IfNotPresent",
                            Command = new[]
                            {
                                "/init-container.sh",
                            },
                            Env = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                {
                                    Name = "CILIUM_ALL_STATE",
                                    ValueFrom = new Kubernetes.Types.Inputs.Core.V1.EnvVarSourceArgs
                                    {
                                        ConfigMapKeyRef = new Kubernetes.Types.Inputs.Core.V1.ConfigMapKeySelectorArgs
                                        {
                                            Name = "cilium-config",
                                            Key = "clean-cilium-state",
                                            Optional = true,
                                        },
                                    },
                                },
                                new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                {
                                    Name = "CILIUM_BPF_STATE",
                                    ValueFrom = new Kubernetes.Types.Inputs.Core.V1.EnvVarSourceArgs
                                    {
                                        ConfigMapKeyRef = new Kubernetes.Types.Inputs.Core.V1.ConfigMapKeySelectorArgs
                                        {
                                            Name = "cilium-config",
                                            Key = "clean-cilium-bpf-state",
                                            Optional = true,
                                        },
                                    },
                                },
                            },
                            TerminationMessagePolicy = "FallbackToLogsOnError",
                            SecurityContext = new Kubernetes.Types.Inputs.Core.V1.SecurityContextArgs
                            {
                                SeLinuxOptions = new Kubernetes.Types.Inputs.Core.V1.SELinuxOptionsArgs
                                {
                                    Level = "s0",
                                    Type = "spc_t",
                                },
                                Capabilities = new Kubernetes.Types.Inputs.Core.V1.CapabilitiesArgs
                                {
                                    Add = new[]
                                    {
                                        "NET_ADMIN",
                                        "SYS_MODULE",
                                        "SYS_ADMIN",
                                        "SYS_RESOURCE",
                                    },
                                    Drop = new[]
                                    {
                                        "ALL",
                                    },
                                },
                            },
                            VolumeMounts = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "bpf-maps",
                                    MountPath = "/sys/fs/bpf",
                                },
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "cilium-cgroup",
                                    MountPath = "/run/cilium/cgroupv2",
                                    MountPropagation = "HostToContainer",
                                },
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "cilium-run",
                                    MountPath = "/var/run/cilium",
                                },
                            },
                            Resources = new Kubernetes.Types.Inputs.Core.V1.ResourceRequirementsArgs
                            {
                                Requests = 
                                {
                                    { "cpu", "100m" },
                                    { "memory", "100Mi" },
                                },
                            },
                        },
                        new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                        {
                            Name = "install-cni-binaries",
                            Image = "quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35",
                            ImagePullPolicy = "IfNotPresent",
                            Command = new[]
                            {
                                "/install-plugin.sh",
                            },
                            Resources = new Kubernetes.Types.Inputs.Core.V1.ResourceRequirementsArgs
                            {
                                Requests = 
                                {
                                    { "cpu", "100m" },
                                    { "memory", "10Mi" },
                                },
                            },
                            SecurityContext = new Kubernetes.Types.Inputs.Core.V1.SecurityContextArgs
                            {
                                SeLinuxOptions = new Kubernetes.Types.Inputs.Core.V1.SELinuxOptionsArgs
                                {
                                    Level = "s0",
                                    Type = "spc_t",
                                },
                                Capabilities = new Kubernetes.Types.Inputs.Core.V1.CapabilitiesArgs
                                {
                                    Drop = new[]
                                    {
                                        "ALL",
                                    },
                                },
                            },
                            TerminationMessagePolicy = "FallbackToLogsOnError",
                            VolumeMounts = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    Name = "cni-path",
                                    MountPath = "/host/opt/cni/bin",
                                },
                            },
                        },
                    },
                    RestartPolicy = "Always",
                    PriorityClassName = "system-node-critical",
                    ServiceAccount = "cilium",
                    ServiceAccountName = "cilium",
                    AutomountServiceAccountToken = true,
                    TerminationGracePeriodSeconds = 1,
                    HostNetwork = true,
                    Affinity = new Kubernetes.Types.Inputs.Core.V1.AffinityArgs
                    {
                        PodAntiAffinity = new Kubernetes.Types.Inputs.Core.V1.PodAntiAffinityArgs
                        {
                            RequiredDuringSchedulingIgnoredDuringExecution = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.PodAffinityTermArgs
                                {
                                    LabelSelector = new Kubernetes.Types.Inputs.Meta.V1.LabelSelectorArgs
                                    {
                                        MatchLabels = 
                                        {
                                            { "k8s-app", "cilium" },
                                        },
                                    },
                                    TopologyKey = "kubernetes.io/hostname",
                                },
                            },
                        },
                    },
                    NodeSelector = 
                    {
                        { "kubernetes.io/os", "linux" },
                    },
                    Tolerations = new[]
                    {
                        new Kubernetes.Types.Inputs.Core.V1.TolerationArgs
                        {
                            Operator = "Exists",
                        },
                    },
                    Volumes = new[]
                    {
                        new Kubernetes.Types.Inputs.Core.V1.VolumeArgs
                        {
                            Name = "tmp",
                            EmptyDir = null,
                        },
                        new Kubernetes.Types.Inputs.Core.V1.VolumeArgs
                        {
                            Name = "cilium-run",
                            HostPath = new Kubernetes.Types.Inputs.Core.V1.HostPathVolumeSourceArgs
                            {
                                Path = "/var/run/cilium",
                                Type = "DirectoryOrCreate",
                            },
                        },
                        new Kubernetes.Types.Inputs.Core.V1.VolumeArgs
                        {
                            Name = "bpf-maps",
                            HostPath = new Kubernetes.Types.Inputs.Core.V1.HostPathVolumeSourceArgs
                            {
                                Path = "/sys/fs/bpf",
                                Type = "DirectoryOrCreate",
                            },
                        },
                        new Kubernetes.Types.Inputs.Core.V1.VolumeArgs
                        {
                            Name = "hostproc",
                            HostPath = new Kubernetes.Types.Inputs.Core.V1.HostPathVolumeSourceArgs
                            {
                                Path = "/proc",
                                Type = "Directory",
                            },
                        },
                        new Kubernetes.Types.Inputs.Core.V1.VolumeArgs
                        {
                            Name = "cilium-cgroup",
                            HostPath = new Kubernetes.Types.Inputs.Core.V1.HostPathVolumeSourceArgs
                            {
                                Path = "/run/cilium/cgroupv2",
                                Type = "DirectoryOrCreate",
                            },
                        },
                        new Kubernetes.Types.Inputs.Core.V1.VolumeArgs
                        {
                            Name = "cni-path",
                            HostPath = new Kubernetes.Types.Inputs.Core.V1.HostPathVolumeSourceArgs
                            {
                                Path = "/opt/cni/bin",
                                Type = "DirectoryOrCreate",
                            },
                        },
                        new Kubernetes.Types.Inputs.Core.V1.VolumeArgs
                        {
                            Name = "etc-cni-netd",
                            HostPath = new Kubernetes.Types.Inputs.Core.V1.HostPathVolumeSourceArgs
                            {
                                Path = "/etc/cni/net.d",
                                Type = "DirectoryOrCreate",
                            },
                        },
                        new Kubernetes.Types.Inputs.Core.V1.VolumeArgs
                        {
                            Name = "lib-modules",
                            HostPath = new Kubernetes.Types.Inputs.Core.V1.HostPathVolumeSourceArgs
                            {
                                Path = "/lib/modules",
                            },
                        },
                        new Kubernetes.Types.Inputs.Core.V1.VolumeArgs
                        {
                            Name = "xtables-lock",
                            HostPath = new Kubernetes.Types.Inputs.Core.V1.HostPathVolumeSourceArgs
                            {
                                Path = "/run/xtables.lock",
                                Type = "FileOrCreate",
                            },
                        },
                        new Kubernetes.Types.Inputs.Core.V1.VolumeArgs
                        {
                            Name = "clustermesh-secrets",
                            Projected = new Kubernetes.Types.Inputs.Core.V1.ProjectedVolumeSourceArgs
                            {
                                DefaultMode = 400,
                                Sources = new[]
                                {
                                    new Kubernetes.Types.Inputs.Core.V1.VolumeProjectionArgs
                                    {
                                        Secret = new Kubernetes.Types.Inputs.Core.V1.SecretProjectionArgs
                                        {
                                            Name = "cilium-clustermesh",
                                            Optional = true,
                                        },
                                    },
                                    new Kubernetes.Types.Inputs.Core.V1.VolumeProjectionArgs
                                    {
                                        Secret = new Kubernetes.Types.Inputs.Core.V1.SecretProjectionArgs
                                        {
                                            Name = "clustermesh-apiserver-remote-cert",
                                            Optional = true,
                                            Items = new[]
                                            {
                                                new Kubernetes.Types.Inputs.Core.V1.KeyToPathArgs
                                                {
                                                    Key = "tls.key",
                                                    Path = "common-etcd-client.key",
                                                },
                                                new Kubernetes.Types.Inputs.Core.V1.KeyToPathArgs
                                                {
                                                    Key = "tls.crt",
                                                    Path = "common-etcd-client.crt",
                                                },
                                                new Kubernetes.Types.Inputs.Core.V1.KeyToPathArgs
                                                {
                                                    Key = "ca.crt",
                                                    Path = "common-etcd-client-ca.crt",
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        new Kubernetes.Types.Inputs.Core.V1.VolumeArgs
                        {
                            Name = "host-proc-sys-net",
                            HostPath = new Kubernetes.Types.Inputs.Core.V1.HostPathVolumeSourceArgs
                            {
                                Path = "/proc/sys/net",
                                Type = "Directory",
                            },
                        },
                        new Kubernetes.Types.Inputs.Core.V1.VolumeArgs
                        {
                            Name = "host-proc-sys-kernel",
                            HostPath = new Kubernetes.Types.Inputs.Core.V1.HostPathVolumeSourceArgs
                            {
                                Path = "/proc/sys/kernel",
                                Type = "Directory",
                            },
                        },
                        new Kubernetes.Types.Inputs.Core.V1.VolumeArgs
                        {
                            Name = "hubble-tls",
                            Projected = new Kubernetes.Types.Inputs.Core.V1.ProjectedVolumeSourceArgs
                            {
                                DefaultMode = 400,
                                Sources = new[]
                                {
                                    new Kubernetes.Types.Inputs.Core.V1.VolumeProjectionArgs
                                    {
                                        Secret = new Kubernetes.Types.Inputs.Core.V1.SecretProjectionArgs
                                        {
                                            Name = "hubble-server-certs",
                                            Optional = true,
                                            Items = new[]
                                            {
                                                new Kubernetes.Types.Inputs.Core.V1.KeyToPathArgs
                                                {
                                                    Key = "tls.crt",
                                                    Path = "server.crt",
                                                },
                                                new Kubernetes.Types.Inputs.Core.V1.KeyToPathArgs
                                                {
                                                    Key = "tls.key",
                                                    Path = "server.key",
                                                },
                                                new Kubernetes.Types.Inputs.Core.V1.KeyToPathArgs
                                                {
                                                    Key = "ca.crt",
                                                    Path = "client-ca.crt",
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
    });

});

