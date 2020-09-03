using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {
        var argocd_serverDeployment = new Kubernetes.Apps.V1.Deployment("argocd_serverDeployment", new Kubernetes.Types.Inputs.Apps.V1.DeploymentArgs
        {
            ApiVersion = "apps/v1",
            Kind = "Deployment",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Labels = 
                {
                    { "app.kubernetes.io/component", "server" },
                    { "aws:region", "us-west-2" },
                    { "key%percent", "percent" },
                    { "key...ellipse", "ellipse" },
                    { "key{bracket", "bracket" },
                    { "key}bracket", "bracket" },
                    { "key*asterix", "asterix" },
                    { "key?question", "question" },
                    { "key,comma", "comma" },
                    { "key&&and", "and" },
                    { "key||or", "or" },
                    { "key!not", "not" },
                    { "key=>geq", "geq" },
                    { "key==eq", "equal" },
                },
                Name = "argocd-server",
            },
        });
    }

}
