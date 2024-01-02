using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

return await Deployment.RunAsync(() => 
{
    var myappConfigMap = new Kubernetes.Core.V1.ConfigMap("myappConfigMap", new()
    {
        ApiVersion = "v1",
        Kind = "ConfigMap",
        Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
        {
            Name = "myapp",
        },
        Data = 
        {
            { "key", "{\\\"uid\\\": \\\"$(datasource)\\\"}" },
        },
    });

    var myapp_varConfigMap = new Kubernetes.Core.V1.ConfigMap("myapp_varConfigMap", new()
    {
        ApiVersion = "v1",
        Kind = "ConfigMap",
        Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
        {
            Name = "myapp-var",
        },
        Data = 
        {
            { "key", "{\\\"uid\\\": \\\"${datasource}\\\"}" },
        },
    });

    var myapp_no_end_bracketConfigMap = new Kubernetes.Core.V1.ConfigMap("myapp_no_end_bracketConfigMap", new()
    {
        ApiVersion = "v1",
        Kind = "ConfigMap",
        Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
        {
            Name = "myapp-no-end-bracket",
        },
        Data = 
        {
            { "key", "{\\\"uid\\\": \\\"${datasource\\\"}" },
        },
    });

    var myapp_no_bracketsConfigMap = new Kubernetes.Core.V1.ConfigMap("myapp_no_bracketsConfigMap", new()
    {
        ApiVersion = "v1",
        Kind = "ConfigMap",
        Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
        {
            Name = "myapp-no-brackets",
        },
        Data = 
        {
            { "key", "{\\\"uid\\\": \\\"$datasource\\\"" },
        },
    });

});

