using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Kubernetes.Core.v1.Namespace("foo", new Kubernetes.Core.v1.NamespaceArgs
        {
            ApiVersion = "v1",
            Kind = "Namespace",
            Metadata = new Kubernetes.Meta.Inputs.ObjectMetaArgs
            {
                Name = "foo",
            },
        });
    }

}
