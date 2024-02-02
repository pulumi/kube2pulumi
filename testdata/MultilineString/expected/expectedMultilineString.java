package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.kubernetes.core_v1.ConfigMap;
import com.pulumi.kubernetes.core_v1.ConfigMapArgs;
import com.pulumi.kubernetes.meta_v1.inputs.ObjectMetaArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var kube_systemCorednsConfigMap = new ConfigMap("kube_systemCorednsConfigMap", ConfigMapArgs.builder()        
            .apiVersion("v1")
            .kind("ConfigMap")
            .metadata(ObjectMetaArgs.builder()
                .name("coredns")
                .namespace("kube-system")
                .build())
            .data(Map.of("Corefile", """
.:53 {
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
            """))
            .build());

    }
}
