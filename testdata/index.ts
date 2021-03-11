import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const fooNamespace = new kubernetes.core.v1.Namespace("fooNamespace", {
    apiVersion: "v1",
    kind: "Namespace",
    metadata: {
        name: "foo",
    },
});
