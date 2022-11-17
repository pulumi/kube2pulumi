package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.kubernetes.apps_v1.Deployment;
import com.pulumi.kubernetes.apps_v1.DeploymentArgs;
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
        var argocd_serverDeployment = new Deployment("argocd_serverDeployment", DeploymentArgs.builder()        
            .apiVersion("apps/v1")
            .kind("Deployment")
            .metadata(ObjectMetaArgs.builder()
                .labels(Map.ofEntries(
                    Map.entry("app.kubernetes.io/component", "server"),
                    Map.entry("aws:region", "us-west-2"),
                    Map.entry("key%percent", "percent"),
                    Map.entry("key...ellipse", "ellipse"),
                    Map.entry("key{bracket", "bracket"),
                    Map.entry("key}bracket", "bracket"),
                    Map.entry("key*asterix", "asterix"),
                    Map.entry("key?question", "question"),
                    Map.entry("key,comma", "comma"),
                    Map.entry("key&&and", "and"),
                    Map.entry("key||or", "or"),
                    Map.entry("key!not", "not"),
                    Map.entry("key=>geq", "geq"),
                    Map.entry("key==eq", "equal")
                ))
                .name("argocd-server")
                .build())
            .build());

    }
}
