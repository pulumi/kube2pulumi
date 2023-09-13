import pulumi
import pulumi_kubernetes as kubernetes

myapp_config_map = kubernetes.core.v1.ConfigMap("myappConfigMap",
    api_version="v1",
    kind="ConfigMap",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="myapp",
    ),
    data={
        "key": "{\\\"uid\\\": \\\"$(datasource)\\\"}",
    })
myapp_var_config_map = kubernetes.core.v1.ConfigMap("myapp_varConfigMap",
    api_version="v1",
    kind="ConfigMap",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="myapp-var",
    ),
    data={
        "key": "{\\\"uid\\\": \\\"${datasource}\\\"}",
    })
myapp_no_end_bracket_config_map = kubernetes.core.v1.ConfigMap("myapp_no_end_bracketConfigMap",
    api_version="v1",
    kind="ConfigMap",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="myapp-no-end-bracket",
    ),
    data={
        "key": "{\\\"uid\\\": \\\"${datasource\\\"}",
    })
myapp_no_brackets_config_map = kubernetes.core.v1.ConfigMap("myapp_no_bracketsConfigMap",
    api_version="v1",
    kind="ConfigMap",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="myapp-no-brackets",
    ),
    data={
        "key": "{\\\"uid\\\": \\\"$datasource\\\"",
    })
