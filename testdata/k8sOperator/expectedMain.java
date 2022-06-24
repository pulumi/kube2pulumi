package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.kubernetes.apps_v1.Deployment;
import com.pulumi.kubernetes.apps_v1.DeploymentArgs;
import com.pulumi.kubernetes.meta_v1.inputs.ObjectMetaArgs;
import com.pulumi.kubernetes.apps_v1.inputs.DeploymentSpecArgs;
import com.pulumi.kubernetes.meta_v1.inputs.LabelSelectorArgs;
import com.pulumi.kubernetes.core_v1.inputs.PodTemplateSpecArgs;
import com.pulumi.kubernetes.core_v1.inputs.PodSpecArgs;
import com.pulumi.kubernetes.rbac.authorization.k8s.io_v1.Role;
import com.pulumi.kubernetes.rbac.authorization.k8s.io_v1.RoleArgs;
import com.pulumi.kubernetes.rbac.authorization.k8s.io_v1.inputs.PolicyRuleArgs;
import com.pulumi.kubernetes.rbac.authorization.k8s.io_v1.RoleBinding;
import com.pulumi.kubernetes.rbac.authorization.k8s.io_v1.RoleBindingArgs;
import com.pulumi.kubernetes.rbac.authorization.k8s.io_v1.inputs.SubjectArgs;
import com.pulumi.kubernetes.rbac.authorization.k8s.io_v1.inputs.RoleRefArgs;
import com.pulumi.kubernetes.core_v1.ServiceAccount;
import com.pulumi.kubernetes.core_v1.ServiceAccountArgs;
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
        var pulumi_kubernetes_operatorDeployment = new Deployment("pulumi_kubernetes_operatorDeployment", DeploymentArgs.builder()        
            .apiVersion("apps/v1")
            .kind("Deployment")
            .metadata(ObjectMetaArgs.builder()
                .name("pulumi-kubernetes-operator")
                .build())
            .spec(DeploymentSpecArgs.builder()
                .replicas(1)
                .selector(LabelSelectorArgs.builder()
                    .matchLabels(Map.of("name", "pulumi-kubernetes-operator"))
                    .build())
                .template(PodTemplateSpecArgs.builder()
                    .metadata(ObjectMetaArgs.builder()
                        .labels(Map.of("name", "pulumi-kubernetes-operator"))
                        .build())
                    .spec(PodSpecArgs.builder()
                        .serviceAccountName("pulumi-kubernetes-operator")
                        .imagePullSecrets(LocalObjectReferenceArgs.builder()
                            .name("pulumi-kubernetes-operator")
                            .build())
                        .containers(ContainerArgs.builder()
                            .name("pulumi-kubernetes-operator")
                            .image("pulumi/pulumi-kubernetes-operator:v0.0.2")
                            .command("pulumi-kubernetes-operator")
                            .args("--zap-level=debug")
                            .imagePullPolicy("Always")
                            .env(                            
                                EnvVarArgs.builder()
                                    .name("WATCH_NAMESPACE")
                                    .valueFrom(EnvVarSourceArgs.builder()
                                        .fieldRef(ObjectFieldSelectorArgs.builder()
                                            .fieldPath("metadata.namespace")
                                            .build())
                                        .build())
                                    .build(),
                                EnvVarArgs.builder()
                                    .name("POD_NAME")
                                    .valueFrom(EnvVarSourceArgs.builder()
                                        .fieldRef(ObjectFieldSelectorArgs.builder()
                                            .fieldPath("metadata.name")
                                            .build())
                                        .build())
                                    .build(),
                                EnvVarArgs.builder()
                                    .name("OPERATOR_NAME")
                                    .value("pulumi-kubernetes-operator")
                                    .build())
                            .build())
                        .build())
                    .build())
                .build())
            .build());

        var pulumi_kubernetes_operatorRole = new Role("pulumi_kubernetes_operatorRole", RoleArgs.builder()        
            .apiVersion("rbac.authorization.k8s.io/v1")
            .kind("Role")
            .metadata(ObjectMetaArgs.builder()
                .creationTimestamp(null)
                .name("pulumi-kubernetes-operator")
                .build())
            .rules(            
                PolicyRuleArgs.builder()
                    .apiGroups("")
                    .resources(                    
                        "pods",
                        "services",
                        "services/finalizers",
                        "endpoints",
                        "persistentvolumeclaims",
                        "events",
                        "configmaps",
                        "secrets")
                    .verbs(                    
                        "create",
                        "delete",
                        "get",
                        "list",
                        "patch",
                        "update",
                        "watch")
                    .build(),
                PolicyRuleArgs.builder()
                    .apiGroups("apps")
                    .resources(                    
                        "deployments",
                        "daemonsets",
                        "replicasets",
                        "statefulsets")
                    .verbs(                    
                        "create",
                        "delete",
                        "get",
                        "list",
                        "patch",
                        "update",
                        "watch")
                    .build(),
                PolicyRuleArgs.builder()
                    .apiGroups("monitoring.coreos.com")
                    .resources("servicemonitors")
                    .verbs(                    
                        "get",
                        "create")
                    .build(),
                PolicyRuleArgs.builder()
                    .apiGroups("apps")
                    .resourceNames("pulumi-kubernetes-operator")
                    .resources("deployments/finalizers")
                    .verbs("update")
                    .build(),
                PolicyRuleArgs.builder()
                    .apiGroups("")
                    .resources("pods")
                    .verbs("get")
                    .build(),
                PolicyRuleArgs.builder()
                    .apiGroups("apps")
                    .resources(                    
                        "replicasets",
                        "deployments")
                    .verbs("get")
                    .build(),
                PolicyRuleArgs.builder()
                    .apiGroups("pulumi.com")
                    .resources("*")
                    .verbs(                    
                        "create",
                        "delete",
                        "get",
                        "list",
                        "patch",
                        "update",
                        "watch")
                    .build())
            .build());

        var pulumi_kubernetes_operatorRoleBinding = new RoleBinding("pulumi_kubernetes_operatorRoleBinding", RoleBindingArgs.builder()        
            .kind("RoleBinding")
            .apiVersion("rbac.authorization.k8s.io/v1")
            .metadata(ObjectMetaArgs.builder()
                .name("pulumi-kubernetes-operator")
                .build())
            .subjects(SubjectArgs.builder()
                .kind("ServiceAccount")
                .name("pulumi-kubernetes-operator")
                .build())
            .roleRef(RoleRefArgs.builder()
                .kind("Role")
                .name("pulumi-kubernetes-operator")
                .apiGroup("rbac.authorization.k8s.io")
                .build())
            .build());

        var pulumi_kubernetes_operatorServiceAccount = new ServiceAccount("pulumi_kubernetes_operatorServiceAccount", ServiceAccountArgs.builder()        
            .apiVersion("v1")
            .kind("ServiceAccount")
            .metadata(ObjectMetaArgs.builder()
                .name("pulumi-kubernetes-operator")
                .build())
            .build());

    }
}
