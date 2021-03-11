import pulumi
import pulumi_kubernetes as kubernetes

foo_namespace = kubernetes.core.v1.Namespace("fooNamespace",
    api_version="v1",
    kind="Namespace",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="foo",
    ))
