import pulumi
import pulumi_kubernetes as kubernetes

kube_system_coredns_config_map = kubernetes.core.v1.ConfigMap("kube_systemCorednsConfigMap",
    api_version="v1",
    kind="ConfigMap",
    metadata={
        "name": "coredns",
        "namespace": "kube-system",
    },
    data={
        "Corefile": """.:53 {
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
""",
    })
