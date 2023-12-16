import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const myappConfigMap = new kubernetes.core.v1.ConfigMap("myappConfigMap", {
    apiVersion: "v1",
    kind: "ConfigMap",
    metadata: {
        name: "myapp",
    },
    data: {
        key: "{\\\"uid\\\": \\\"$(datasource)\\\"}",
    },
});
const myapp_varConfigMap = new kubernetes.core.v1.ConfigMap("myapp_varConfigMap", {
    apiVersion: "v1",
    kind: "ConfigMap",
    metadata: {
        name: "myapp-var",
    },
    data: {
        key: "{\\\"uid\\\": \\\"${datasource}\\\"}",
    },
});
const myapp_no_end_bracketConfigMap = new kubernetes.core.v1.ConfigMap("myapp_no_end_bracketConfigMap", {
    apiVersion: "v1",
    kind: "ConfigMap",
    metadata: {
        name: "myapp-no-end-bracket",
    },
    data: {
        key: "{\\\"uid\\\": \\\"${datasource\\\"}",
    },
});
const myapp_no_bracketsConfigMap = new kubernetes.core.v1.ConfigMap("myapp_no_bracketsConfigMap", {
    apiVersion: "v1",
    kind: "ConfigMap",
    metadata: {
        name: "myapp-no-brackets",
    },
    data: {
        key: "{\\\"uid\\\": \\\"$datasource\\\"",
    },
});
