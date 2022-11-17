package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.kubernetes.apps_v1.Deployment;
import com.pulumi.kubernetes.apps_v1.DeploymentArgs;
import com.pulumi.kubernetes.meta_v1.inputs.ObjectMetaArgs;
import com.pulumi.kubernetes.apps_v1.inputs.DeploymentSpecArgs;
import com.pulumi.kubernetes.meta_v1.inputs.LabelSelectorArgs;
import com.pulumi.kubernetes.apps_v1.inputs.DeploymentStrategyArgs;
import com.pulumi.kubernetes.apps_v1.inputs.RollingUpdateDeploymentArgs;
import com.pulumi.kubernetes.core_v1.inputs.PodTemplateSpecArgs;
import com.pulumi.kubernetes.core_v1.inputs.PodSpecArgs;
import com.pulumi.kubernetes.core_v1.inputs.PodSecurityContextArgs;
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
        var defaultArgocd_serverDeployment = new Deployment("defaultArgocd_serverDeployment", DeploymentArgs.builder()        
            .apiVersion("apps/v1")
            .kind("Deployment")
            .metadata(ObjectMetaArgs.builder()
                .annotations(Map.ofEntries(
                    Map.entry("deployment.kubernetes.io/revision", "1"),
                    Map.entry("kubectl.kubernetes.io/last-applied-configuration", """
{"apiVersion":"apps/v1","kind":"Deployment","metadata":{"labels":{"app.kubernetes.io/component":"server","app.kubernetes.io/instance":"argocd","app.kubernetes.io/managed-by":"pulumi","app.kubernetes.io/name":"argocd-server","app.kubernetes.io/part-of":"argocd","app.kubernetes.io/version":"v1.6.1","helm.sh/chart":"argo-cd-2.5.4"},"name":"argocd-server","namespace":"default"},"spec":{"replicas":1,"revisionHistoryLimit":5,"selector":{"matchLabels":{"app.kubernetes.io/instance":"argocd","app.kubernetes.io/name":"argocd-server"}},"template":{"metadata":{"labels":{"app.kubernetes.io/component":"server","app.kubernetes.io/instance":"argocd","app.kubernetes.io/managed-by":"Helm","app.kubernetes.io/name":"argocd-server","app.kubernetes.io/part-of":"argocd","app.kubernetes.io/version":"v1.6.1","helm.sh/chart":"argo-cd-2.5.4"}},"spec":{"containers":[{"command":["argocd-server","--staticassets","/shared/app","--repo-server","argocd-repo-server:8081","--dex-server","http://argocd-dex-server:5556","--loglevel","info","--redis","argocd-redis:6379"],"image":"argoproj/argocd:v1.6.1","imagePullPolicy":"IfNotPresent","livenessProbe":{"failureThreshold":3,"httpGet":{"path":"/healthz","port":8080},"initialDelaySeconds":10,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1},"name":"server","ports":[{"containerPort":8080,"name":"server","protocol":"TCP"}],"readinessProbe":{"failureThreshold":3,"httpGet":{"path":"/healthz","port":8080},"initialDelaySeconds":10,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1},"resources":{},"volumeMounts":[{"mountPath":"/app/config/ssh","name":"ssh-known-hosts"}]}],"serviceAccountName":"argocd-server","volumes":[{"emptyDir":{},"name":"static-files"},{"configMap":{"name":"argocd-ssh-known-hosts-cm"},"name":"ssh-known-hosts"}]}}}}
                    """)
                ))
                .creationTimestamp("2020-08-04T18:50:43Z")
                .generation(1)
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
                .namespace("default")
                .resourceVersion("1406")
                .selfLink("/apis/apps/v1/namespaces/default/deployments/argocd-server")
                .uid("4b806e77-b035-41a3-bdf9-9781b76445f9")
                .build())
            .spec(DeploymentSpecArgs.builder()
                .progressDeadlineSeconds(600)
                .replicas(1)
                .revisionHistoryLimit(5)
                .selector(LabelSelectorArgs.builder()
                    .matchLabels(Map.ofEntries(
                        Map.entry("app.kubernetes.io/instance", "argocd"),
                        Map.entry("app.kubernetes.io/name", "argocd-server")
                    ))
                    .build())
                .strategy(DeploymentStrategyArgs.builder()
                    .rollingUpdate(RollingUpdateDeploymentArgs.builder()
                        .maxSurge("25%")
                        .maxUnavailable("25%")
                        .build())
                    .type("RollingUpdate")
                    .build())
                .template(PodTemplateSpecArgs.builder()
                    .metadata(ObjectMetaArgs.builder()
                        .creationTimestamp(null)
                        .labels(Map.ofEntries(
                            Map.entry("app.kubernetes.io/component", "server"),
                            Map.entry("app.kubernetes.io/instance", "argocd"),
                            Map.entry("app.kubernetes.io/managed-by", "Helm"),
                            Map.entry("app.kubernetes.io/name", "argocd-server"),
                            Map.entry("app.kubernetes.io/part-of", "argocd"),
                            Map.entry("app.kubernetes.io/version", "v1.6.1"),
                            Map.entry("helm.sh/chart", "argo-cd-2.5.4")
                        ))
                        .build())
                    .spec(PodSpecArgs.builder()
                        .containers(ContainerArgs.builder()
                            .command(                            
                                "argocd-server",
                                "--staticassets",
                                "/shared/app",
                                "--repo-server",
                                "argocd-repo-server:8081",
                                "--dex-server",
                                "http://argocd-dex-server:5556",
                                "--loglevel",
                                "info",
                                "--redis",
                                "argocd-redis:6379")
                            .image("argoproj/argocd:v1.6.1")
                            .imagePullPolicy("IfNotPresent")
                            .livenessProbe(ProbeArgs.builder()
                                .failureThreshold(3)
                                .httpGet(HTTPGetActionArgs.builder()
                                    .path("/healthz")
                                    .port(8080)
                                    .scheme("HTTP")
                                    .build())
                                .initialDelaySeconds(10)
                                .periodSeconds(10)
                                .successThreshold(1)
                                .timeoutSeconds(1)
                                .build())
                            .name("server")
                            .ports(ContainerPortArgs.builder()
                                .containerPort(8080)
                                .name("server")
                                .protocol("TCP")
                                .build())
                            .readinessProbe(ProbeArgs.builder()
                                .failureThreshold(3)
                                .httpGet(HTTPGetActionArgs.builder()
                                    .path("/healthz")
                                    .port(8080)
                                    .scheme("HTTP")
                                    .build())
                                .initialDelaySeconds(10)
                                .periodSeconds(10)
                                .successThreshold(1)
                                .timeoutSeconds(1)
                                .build())
                            .resources()
                            .terminationMessagePath("/dev/termination-log")
                            .terminationMessagePolicy("File")
                            .volumeMounts(VolumeMountArgs.builder()
                                .mountPath("/app/config/ssh")
                                .name("ssh-known-hosts")
                                .build())
                            .build())
                        .dnsPolicy("ClusterFirst")
                        .restartPolicy("Always")
                        .schedulerName("default-scheduler")
                        .securityContext()
                        .serviceAccount("argocd-server")
                        .serviceAccountName("argocd-server")
                        .terminationGracePeriodSeconds(30)
                        .volumes(                        
                            VolumeArgs.builder()
                                .emptyDir()
                                .name("static-files")
                                .build(),
                            VolumeArgs.builder()
                                .configMap(ConfigMapVolumeSourceArgs.builder()
                                    .defaultMode(420)
                                    .name("argocd-ssh-known-hosts-cm")
                                    .build())
                                .name("ssh-known-hosts")
                                .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
