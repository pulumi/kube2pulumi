using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Kubernetes.Core.V1.Namespace("foo", new Kubernetes.Types.Inputs.Core.V1.NamespaceArgs
        {
            ApiVersion = "v1",
            Kind = "Namespace",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Name = "foo",
            },
        });
    }

}
