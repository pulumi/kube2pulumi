using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

return await Deployment.RunAsync(() => 
{
    var foo = new Kubernetes.Core.V1.Namespace("foo", new()
    {
        ApiVersion = "v1",
        Kind = "Namespace",
        Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
        {
            Name = "foo",
        },
    });

});

