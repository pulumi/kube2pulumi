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
        var myappConfigMap = new ConfigMap("myappConfigMap", ConfigMapArgs.builder()        
            .apiVersion("v1")
            .kind("ConfigMap")
            .metadata(ObjectMetaArgs.builder()
                .name("myapp")
                .build())
            .data(Map.of("key", "{\\\"uid\\\": \\\"$(datasource)\\\"}"))
            .build());

        var myapp_varConfigMap = new ConfigMap("myapp_varConfigMap", ConfigMapArgs.builder()        
            .apiVersion("v1")
            .kind("ConfigMap")
            .metadata(ObjectMetaArgs.builder()
                .name("myapp-var")
                .build())
            .data(Map.of("key", "{\\\"uid\\\": \\\"${datasource}\\\"}"))
            .build());

        var myapp_no_end_bracketConfigMap = new ConfigMap("myapp_no_end_bracketConfigMap", ConfigMapArgs.builder()        
            .apiVersion("v1")
            .kind("ConfigMap")
            .metadata(ObjectMetaArgs.builder()
                .name("myapp-no-end-bracket")
                .build())
            .data(Map.of("key", "{\\\"uid\\\": \\\"${datasource\\\"}"))
            .build());

        var myapp_no_bracketsConfigMap = new ConfigMap("myapp_no_bracketsConfigMap", ConfigMapArgs.builder()        
            .apiVersion("v1")
            .kind("ConfigMap")
            .metadata(ObjectMetaArgs.builder()
                .name("myapp-no-brackets")
                .build())
            .data(Map.of("key", "{\\\"uid\\\": \\\"$datasource\\\""))
            .build());

    }
}
