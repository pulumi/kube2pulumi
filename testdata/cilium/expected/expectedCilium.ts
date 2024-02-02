import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const defaultCiliumDaemonSet = new kubernetes.apps.v1.DaemonSet("defaultCiliumDaemonSet", {
    apiVersion: "apps/v1",
    kind: "DaemonSet",
    metadata: {
        name: "cilium",
        namespace: "default",
        labels: {
            "k8s-app": "cilium",
            "app.kubernetes.io/part-of": "cilium",
            "app.kubernetes.io/name": "cilium-agent",
        },
    },
    spec: {
        selector: {
            matchLabels: {
                "k8s-app": "cilium",
            },
        },
        updateStrategy: {
            rollingUpdate: {
                maxUnavailable: 2,
            },
            type: "RollingUpdate",
        },
        template: {
            metadata: {
                annotations: {
                    "container.apparmor.security.beta.kubernetes.io/cilium-agent": "unconfined",
                    "container.apparmor.security.beta.kubernetes.io/clean-cilium-state": "unconfined",
                    "container.apparmor.security.beta.kubernetes.io/mount-cgroup": "unconfined",
                    "container.apparmor.security.beta.kubernetes.io/apply-sysctl-overwrites": "unconfined",
                },
                labels: {
                    "k8s-app": "cilium",
                    "app.kubernetes.io/name": "cilium-agent",
                    "app.kubernetes.io/part-of": "cilium",
                },
            },
            spec: {
                containers: [{
                    name: "cilium-agent",
                    image: "quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35",
                    imagePullPolicy: "IfNotPresent",
                    command: ["cilium-agent"],
                    args: ["--config-dir=/tmp/cilium/config-map"],
                    startupProbe: {
                        httpGet: {
                            host: "127.0.0.1",
                            path: "/healthz",
                            port: 9879,
                            scheme: "HTTP",
                            httpHeaders: [{
                                name: "brief",
                                value: "true",
                            }],
                        },
                        failureThreshold: 105,
                        periodSeconds: 2,
                        successThreshold: 1,
                    },
                    livenessProbe: {
                        httpGet: {
                            host: "127.0.0.1",
                            path: "/healthz",
                            port: 9879,
                            scheme: "HTTP",
                            httpHeaders: [{
                                name: "brief",
                                value: "true",
                            }],
                        },
                        periodSeconds: 30,
                        successThreshold: 1,
                        failureThreshold: 10,
                        timeoutSeconds: 5,
                    },
                    readinessProbe: {
                        httpGet: {
                            host: "127.0.0.1",
                            path: "/healthz",
                            port: 9879,
                            scheme: "HTTP",
                            httpHeaders: [{
                                name: "brief",
                                value: "true",
                            }],
                        },
                        periodSeconds: 30,
                        successThreshold: 1,
                        failureThreshold: 3,
                        timeoutSeconds: 5,
                    },
                    env: [
                        {
                            name: "K8S_NODE_NAME",
                            valueFrom: {
                                fieldRef: {
                                    apiVersion: "v1",
                                    fieldPath: "spec.nodeName",
                                },
                            },
                        },
                        {
                            name: "CILIUM_K8S_NAMESPACE",
                            valueFrom: {
                                fieldRef: {
                                    apiVersion: "v1",
                                    fieldPath: "metadata.namespace",
                                },
                            },
                        },
                        {
                            name: "CILIUM_CLUSTERMESH_CONFIG",
                            value: "/var/lib/cilium/clustermesh/",
                        },
                    ],
                    lifecycle: {
                        preStop: {
                            exec: {
                                command: ["/cni-uninstall.sh"],
                            },
                        },
                    },
                    securityContext: {
                        seLinuxOptions: {
                            level: "s0",
                            type: "spc_t",
                        },
                        capabilities: {
                            add: [
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
                            ],
                            drop: ["ALL"],
                        },
                    },
                    terminationMessagePolicy: "FallbackToLogsOnError",
                    volumeMounts: [
                        {
                            mountPath: "/host/proc/sys/net",
                            name: "host-proc-sys-net",
                        },
                        {
                            mountPath: "/host/proc/sys/kernel",
                            name: "host-proc-sys-kernel",
                        },
                        {
                            name: "bpf-maps",
                            mountPath: "/sys/fs/bpf",
                            mountPropagation: "HostToContainer",
                        },
                        {
                            name: "cilium-run",
                            mountPath: "/var/run/cilium",
                        },
                        {
                            name: "etc-cni-netd",
                            mountPath: "/host/etc/cni/net.d",
                        },
                        {
                            name: "clustermesh-secrets",
                            mountPath: "/var/lib/cilium/clustermesh",
                            readOnly: true,
                        },
                        {
                            name: "lib-modules",
                            mountPath: "/lib/modules",
                            readOnly: true,
                        },
                        {
                            name: "xtables-lock",
                            mountPath: "/run/xtables.lock",
                        },
                        {
                            name: "hubble-tls",
                            mountPath: "/var/lib/cilium/tls/hubble",
                            readOnly: true,
                        },
                        {
                            name: "tmp",
                            mountPath: "/tmp",
                        },
                    ],
                }],
                initContainers: [
                    {
                        name: "config",
                        image: "quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35",
                        imagePullPolicy: "IfNotPresent",
                        command: [
                            "cilium",
                            "build-config",
                        ],
                        env: [
                            {
                                name: "K8S_NODE_NAME",
                                valueFrom: {
                                    fieldRef: {
                                        apiVersion: "v1",
                                        fieldPath: "spec.nodeName",
                                    },
                                },
                            },
                            {
                                name: "CILIUM_K8S_NAMESPACE",
                                valueFrom: {
                                    fieldRef: {
                                        apiVersion: "v1",
                                        fieldPath: "metadata.namespace",
                                    },
                                },
                            },
                        ],
                        volumeMounts: [{
                            name: "tmp",
                            mountPath: "/tmp",
                        }],
                        terminationMessagePolicy: "FallbackToLogsOnError",
                    },
                    {
                        name: "mount-cgroup",
                        image: "quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35",
                        imagePullPolicy: "IfNotPresent",
                        env: [
                            {
                                name: "CGROUP_ROOT",
                                value: "/run/cilium/cgroupv2",
                            },
                            {
                                name: "BIN_PATH",
                                value: "/opt/cni/bin",
                            },
                        ],
                        command: [
                            "sh",
                            "-ec",
                            `cp /usr/bin/cilium-mount /hostbin/cilium-mount;
nsenter --cgroup=/hostproc/1/ns/cgroup --mount=/hostproc/1/ns/mnt "\${BIN_PATH}/cilium-mount" $CGROUP_ROOT;
rm /hostbin/cilium-mount
`,
                        ],
                        volumeMounts: [
                            {
                                name: "hostproc",
                                mountPath: "/hostproc",
                            },
                            {
                                name: "cni-path",
                                mountPath: "/hostbin",
                            },
                        ],
                        terminationMessagePolicy: "FallbackToLogsOnError",
                        securityContext: {
                            seLinuxOptions: {
                                level: "s0",
                                type: "spc_t",
                            },
                            capabilities: {
                                add: [
                                    "SYS_ADMIN",
                                    "SYS_CHROOT",
                                    "SYS_PTRACE",
                                ],
                                drop: ["ALL"],
                            },
                        },
                    },
                    {
                        name: "apply-sysctl-overwrites",
                        image: "quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35",
                        imagePullPolicy: "IfNotPresent",
                        env: [{
                            name: "BIN_PATH",
                            value: "/opt/cni/bin",
                        }],
                        command: [
                            "sh",
                            "-ec",
                            `cp /usr/bin/cilium-sysctlfix /hostbin/cilium-sysctlfix;
nsenter --mount=/hostproc/1/ns/mnt "\${BIN_PATH}/cilium-sysctlfix";
rm /hostbin/cilium-sysctlfix
`,
                        ],
                        volumeMounts: [
                            {
                                name: "hostproc",
                                mountPath: "/hostproc",
                            },
                            {
                                name: "cni-path",
                                mountPath: "/hostbin",
                            },
                        ],
                        terminationMessagePolicy: "FallbackToLogsOnError",
                        securityContext: {
                            seLinuxOptions: {
                                level: "s0",
                                type: "spc_t",
                            },
                            capabilities: {
                                add: [
                                    "SYS_ADMIN",
                                    "SYS_CHROOT",
                                    "SYS_PTRACE",
                                ],
                                drop: ["ALL"],
                            },
                        },
                    },
                    {
                        name: "mount-bpf-fs",
                        image: "quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35",
                        imagePullPolicy: "IfNotPresent",
                        args: ["mount | grep \"/sys/fs/bpf type bpf\" || mount -t bpf bpf /sys/fs/bpf"],
                        command: [
                            "/bin/bash",
                            "-c",
                            "--",
                        ],
                        terminationMessagePolicy: "FallbackToLogsOnError",
                        securityContext: {
                            privileged: true,
                        },
                        volumeMounts: [{
                            name: "bpf-maps",
                            mountPath: "/sys/fs/bpf",
                            mountPropagation: "Bidirectional",
                        }],
                    },
                    {
                        name: "clean-cilium-state",
                        image: "quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35",
                        imagePullPolicy: "IfNotPresent",
                        command: ["/init-container.sh"],
                        env: [
                            {
                                name: "CILIUM_ALL_STATE",
                                valueFrom: {
                                    configMapKeyRef: {
                                        name: "cilium-config",
                                        key: "clean-cilium-state",
                                        optional: true,
                                    },
                                },
                            },
                            {
                                name: "CILIUM_BPF_STATE",
                                valueFrom: {
                                    configMapKeyRef: {
                                        name: "cilium-config",
                                        key: "clean-cilium-bpf-state",
                                        optional: true,
                                    },
                                },
                            },
                        ],
                        terminationMessagePolicy: "FallbackToLogsOnError",
                        securityContext: {
                            seLinuxOptions: {
                                level: "s0",
                                type: "spc_t",
                            },
                            capabilities: {
                                add: [
                                    "NET_ADMIN",
                                    "SYS_MODULE",
                                    "SYS_ADMIN",
                                    "SYS_RESOURCE",
                                ],
                                drop: ["ALL"],
                            },
                        },
                        volumeMounts: [
                            {
                                name: "bpf-maps",
                                mountPath: "/sys/fs/bpf",
                            },
                            {
                                name: "cilium-cgroup",
                                mountPath: "/run/cilium/cgroupv2",
                                mountPropagation: "HostToContainer",
                            },
                            {
                                name: "cilium-run",
                                mountPath: "/var/run/cilium",
                            },
                        ],
                        resources: {
                            requests: {
                                cpu: "100m",
                                memory: "100Mi",
                            },
                        },
                    },
                    {
                        name: "install-cni-binaries",
                        image: "quay.io/cilium/cilium:v1.14.2@sha256:6263f3a3d5d63b267b538298dbeb5ae87da3efacf09a2c620446c873ba807d35",
                        imagePullPolicy: "IfNotPresent",
                        command: ["/install-plugin.sh"],
                        resources: {
                            requests: {
                                cpu: "100m",
                                memory: "10Mi",
                            },
                        },
                        securityContext: {
                            seLinuxOptions: {
                                level: "s0",
                                type: "spc_t",
                            },
                            capabilities: {
                                drop: ["ALL"],
                            },
                        },
                        terminationMessagePolicy: "FallbackToLogsOnError",
                        volumeMounts: [{
                            name: "cni-path",
                            mountPath: "/host/opt/cni/bin",
                        }],
                    },
                ],
                restartPolicy: "Always",
                priorityClassName: "system-node-critical",
                serviceAccount: "cilium",
                serviceAccountName: "cilium",
                automountServiceAccountToken: true,
                terminationGracePeriodSeconds: 1,
                hostNetwork: true,
                affinity: {
                    podAntiAffinity: {
                        requiredDuringSchedulingIgnoredDuringExecution: [{
                            labelSelector: {
                                matchLabels: {
                                    "k8s-app": "cilium",
                                },
                            },
                            topologyKey: "kubernetes.io/hostname",
                        }],
                    },
                },
                nodeSelector: {
                    "kubernetes.io/os": "linux",
                },
                tolerations: [{
                    operator: "Exists",
                }],
                volumes: [
                    {
                        name: "tmp",
                        emptyDir: {},
                    },
                    {
                        name: "cilium-run",
                        hostPath: {
                            path: "/var/run/cilium",
                            type: "DirectoryOrCreate",
                        },
                    },
                    {
                        name: "bpf-maps",
                        hostPath: {
                            path: "/sys/fs/bpf",
                            type: "DirectoryOrCreate",
                        },
                    },
                    {
                        name: "hostproc",
                        hostPath: {
                            path: "/proc",
                            type: "Directory",
                        },
                    },
                    {
                        name: "cilium-cgroup",
                        hostPath: {
                            path: "/run/cilium/cgroupv2",
                            type: "DirectoryOrCreate",
                        },
                    },
                    {
                        name: "cni-path",
                        hostPath: {
                            path: "/opt/cni/bin",
                            type: "DirectoryOrCreate",
                        },
                    },
                    {
                        name: "etc-cni-netd",
                        hostPath: {
                            path: "/etc/cni/net.d",
                            type: "DirectoryOrCreate",
                        },
                    },
                    {
                        name: "lib-modules",
                        hostPath: {
                            path: "/lib/modules",
                        },
                    },
                    {
                        name: "xtables-lock",
                        hostPath: {
                            path: "/run/xtables.lock",
                            type: "FileOrCreate",
                        },
                    },
                    {
                        name: "clustermesh-secrets",
                        projected: {
                            defaultMode: 256,
                            sources: [
                                {
                                    secret: {
                                        name: "cilium-clustermesh",
                                        optional: true,
                                    },
                                },
                                {
                                    secret: {
                                        name: "clustermesh-apiserver-remote-cert",
                                        optional: true,
                                        items: [
                                            {
                                                key: "tls.key",
                                                path: "common-etcd-client.key",
                                            },
                                            {
                                                key: "tls.crt",
                                                path: "common-etcd-client.crt",
                                            },
                                            {
                                                key: "ca.crt",
                                                path: "common-etcd-client-ca.crt",
                                            },
                                        ],
                                    },
                                },
                            ],
                        },
                    },
                    {
                        name: "host-proc-sys-net",
                        hostPath: {
                            path: "/proc/sys/net",
                            type: "Directory",
                        },
                    },
                    {
                        name: "host-proc-sys-kernel",
                        hostPath: {
                            path: "/proc/sys/kernel",
                            type: "Directory",
                        },
                    },
                    {
                        name: "hubble-tls",
                        projected: {
                            defaultMode: 256,
                            sources: [{
                                secret: {
                                    name: "hubble-server-certs",
                                    optional: true,
                                    items: [
                                        {
                                            key: "tls.crt",
                                            path: "server.crt",
                                        },
                                        {
                                            key: "tls.key",
                                            path: "server.key",
                                        },
                                        {
                                            key: "ca.crt",
                                            path: "client-ca.crt",
                                        },
                                    ],
                                },
                            }],
                        },
                    },
                ],
            },
        },
    },
});
