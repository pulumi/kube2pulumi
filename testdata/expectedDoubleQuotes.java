package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.kubernetes.apps_v1.Deployment;
import com.pulumi.kubernetes.apps_v1.DeploymentArgs;
import com.pulumi.kubernetes.meta_v1.inputs.ObjectMetaArgs;
import com.pulumi.kubernetes.apps_v1.inputs.DeploymentSpecArgs;
import com.pulumi.kubernetes.core_v1.inputs.PodTemplateSpecArgs;
import com.pulumi.kubernetes.core_v1.inputs.PodSpecArgs;
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
        var argocd_serverDeployment = new Deployment("argocd_serverDeployment", DeploymentArgs.builder()        
            .apiVersion("apps/v1")
            .kind("Deployment")
            .metadata(ObjectMetaArgs.builder()
                .labels(Map.ofEntries(
                    Map.entry("app.kubernetes.io/component", "server"),
                    Map.entry("app.kubernetes.io/instance", "argocd"),
                    Map.entry("app.kubernetes.io/managed-by", "pulumi"),
                    Map.entry("app.kubernetes.io/name", "argocd-server"),
                    Map.entry("app.kubernetes.io/part-of", "argocd"),
                    Map.entry("app.kubernetes.io/version", "v1.6.1"),
                    Map.entry("helm.sh/chart", "argo-cd-2.5.4")
                ))
                .name("argocd-server")
                .build())
            .spec(DeploymentSpecArgs.builder()
                .template(PodTemplateSpecArgs.builder()
                    .spec(PodSpecArgs.builder()
                        .containers(ContainerArgs.builder()
                            .readinessProbe(ProbeArgs.builder()
                                .httpGet(HTTPGetActionArgs.builder()
                                    .port(8080)
                                    .build())
                                .build())
                            .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
