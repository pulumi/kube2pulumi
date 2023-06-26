import pulumi
import pulumi_kubernetes as kubernetes

foo = kubernetes.core.v1.Namespace("foo",
    api_version="v1",
    kind="Namespace",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="foo",
    ))
