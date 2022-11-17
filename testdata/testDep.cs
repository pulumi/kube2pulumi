using System.Collections.Generic;
using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

return await Deployment.RunAsync(() => 
{
    var defaultArgocd_serverDeployment = new Kubernetes.Apps.V1.Deployment("defaultArgocd_serverDeployment", new()
    {
        ApiVersion = "apps/v1",
        Kind = "Deployment",
        Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
        {
            Annotations = 
            {
                { "deployment.kubernetes.io/revision", "1" },
                { "kubectl.kubernetes.io/last-applied-configuration", @"{""apiVersion"":""apps/v1"",""kind"":""Deployment"",""metadata"":{""labels"":{""app.kubernetes.io/component"":""server"",""app.kubernetes.io/instance"":""argocd"",""app.kubernetes.io/managed-by"":""pulumi"",""app.kubernetes.io/name"":""argocd-server"",""app.kubernetes.io/part-of"":""argocd"",""app.kubernetes.io/version"":""v1.6.1"",""helm.sh/chart"":""argo-cd-2.5.4""},""name"":""argocd-server"",""namespace"":""default""},""spec"":{""replicas"":1,""revisionHistoryLimit"":5,""selector"":{""matchLabels"":{""app.kubernetes.io/instance"":""argocd"",""app.kubernetes.io/name"":""argocd-server""}},""template"":{""metadata"":{""labels"":{""app.kubernetes.io/component"":""server"",""app.kubernetes.io/instance"":""argocd"",""app.kubernetes.io/managed-by"":""Helm"",""app.kubernetes.io/name"":""argocd-server"",""app.kubernetes.io/part-of"":""argocd"",""app.kubernetes.io/version"":""v1.6.1"",""helm.sh/chart"":""argo-cd-2.5.4""}},""spec"":{""containers"":[{""command"":[""argocd-server"",""--staticassets"",""/shared/app"",""--repo-server"",""argocd-repo-server:8081"",""--dex-server"",""http://argocd-dex-server:5556"",""--loglevel"",""info"",""--redis"",""argocd-redis:6379""],""image"":""argoproj/argocd:v1.6.1"",""imagePullPolicy"":""IfNotPresent"",""livenessProbe"":{""failureThreshold"":3,""httpGet"":{""path"":""/healthz"",""port"":8080},""initialDelaySeconds"":10,""periodSeconds"":10,""successThreshold"":1,""timeoutSeconds"":1},""name"":""server"",""ports"":[{""containerPort"":8080,""name"":""server"",""protocol"":""TCP""}],""readinessProbe"":{""failureThreshold"":3,""httpGet"":{""path"":""/healthz"",""port"":8080},""initialDelaySeconds"":10,""periodSeconds"":10,""successThreshold"":1,""timeoutSeconds"":1},""resources"":{},""volumeMounts"":[{""mountPath"":""/app/config/ssh"",""name"":""ssh-known-hosts""}]}],""serviceAccountName"":""argocd-server"",""volumes"":[{""emptyDir"":{},""name"":""static-files""},{""configMap"":{""name"":""argocd-ssh-known-hosts-cm""},""name"":""ssh-known-hosts""}]}}}}
" },
            },
            CreationTimestamp = "2020-08-04T18:50:43Z",
            Generation = 1,
            Labels = 
            {
                { "app.kubernetes.io/component", "server" },
                { "app.kubernetes.io/instance", "argocd" },
                { "app.kubernetes.io/managed-by", "pulumi" },
                { "app.kubernetes.io/name", "argocd-server" },
                { "app.kubernetes.io/part-of", "argocd" },
                { "app.kubernetes.io/version", "v1.6.1" },
                { "helm.sh/chart", "argo-cd-2.5.4" },
            },
            Name = "argocd-server",
            Namespace = "default",
            ResourceVersion = "1406",
            SelfLink = "/apis/apps/v1/namespaces/default/deployments/argocd-server",
            Uid = "4b806e77-b035-41a3-bdf9-9781b76445f9",
        },
        Spec = new Kubernetes.Types.Inputs.Apps.V1.DeploymentSpecArgs
        {
            ProgressDeadlineSeconds = 600,
            Replicas = 1,
            RevisionHistoryLimit = 5,
            Selector = new Kubernetes.Types.Inputs.Meta.V1.LabelSelectorArgs
            {
                MatchLabels = 
                {
                    { "app.kubernetes.io/instance", "argocd" },
                    { "app.kubernetes.io/name", "argocd-server" },
                },
            },
            Strategy = new Kubernetes.Types.Inputs.Apps.V1.DeploymentStrategyArgs
            {
                RollingUpdate = new Kubernetes.Types.Inputs.Apps.V1.RollingUpdateDeploymentArgs
                {
                    MaxSurge = "25%",
                    MaxUnavailable = "25%",
                },
                Type = "RollingUpdate",
            },
            Template = new Kubernetes.Types.Inputs.Core.V1.PodTemplateSpecArgs
            {
                Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
                {
                    CreationTimestamp = null,
                    Labels = 
                    {
                        { "app.kubernetes.io/component", "server" },
                        { "app.kubernetes.io/instance", "argocd" },
                        { "app.kubernetes.io/managed-by", "Helm" },
                        { "app.kubernetes.io/name", "argocd-server" },
                        { "app.kubernetes.io/part-of", "argocd" },
                        { "app.kubernetes.io/version", "v1.6.1" },
                        { "helm.sh/chart", "argo-cd-2.5.4" },
                    },
                },
                Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
                {
                    Containers = new[]
                    {
                        new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                        {
                            Command = new[]
                            {
                                "argocd-server",
                                "--staticassets",
                                "/shared/app",
                                "--repo-server",
                                "argocd-repo-server:8081",
                                "--dex-server",
                                "http://argocd-dex-server:5556",
                                "--loglevel",
                                "info",
                                "--redis",
                                "argocd-redis:6379",
                            },
                            Image = "argoproj/argocd:v1.6.1",
                            ImagePullPolicy = "IfNotPresent",
                            LivenessProbe = new Kubernetes.Types.Inputs.Core.V1.ProbeArgs
                            {
                                FailureThreshold = 3,
                                HttpGet = new Kubernetes.Types.Inputs.Core.V1.HTTPGetActionArgs
                                {
                                    Path = "/healthz",
                                    Port = 8080,
                                    Scheme = "HTTP",
                                },
                                InitialDelaySeconds = 10,
                                PeriodSeconds = 10,
                                SuccessThreshold = 1,
                                TimeoutSeconds = 1,
                            },
                            Name = "server",
                            Ports = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.ContainerPortArgs
                                {
                                    ContainerPortValue = 8080,
                                    Name = "server",
                                    Protocol = "TCP",
                                },
                            },
                            ReadinessProbe = new Kubernetes.Types.Inputs.Core.V1.ProbeArgs
                            {
                                FailureThreshold = 3,
                                HttpGet = new Kubernetes.Types.Inputs.Core.V1.HTTPGetActionArgs
                                {
                                    Path = "/healthz",
                                    Port = 8080,
                                    Scheme = "HTTP",
                                },
                                InitialDelaySeconds = 10,
                                PeriodSeconds = 10,
                                SuccessThreshold = 1,
                                TimeoutSeconds = 1,
                            },
                            Resources = null,
                            TerminationMessagePath = "/dev/termination-log",
                            TerminationMessagePolicy = "File",
                            VolumeMounts = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs
                                {
                                    MountPath = "/app/config/ssh",
                                    Name = "ssh-known-hosts",
                                },
                            },
                        },
                    },
                    DnsPolicy = "ClusterFirst",
                    RestartPolicy = "Always",
                    SchedulerName = "default-scheduler",
                    SecurityContext = null,
                    ServiceAccount = "argocd-server",
                    ServiceAccountName = "argocd-server",
                    TerminationGracePeriodSeconds = 30,
                    Volumes = new[]
                    {
                        new Kubernetes.Types.Inputs.Core.V1.VolumeArgs
                        {
                            EmptyDir = null,
                            Name = "static-files",
                        },
                        new Kubernetes.Types.Inputs.Core.V1.VolumeArgs
                        {
                            ConfigMap = new Kubernetes.Types.Inputs.Core.V1.ConfigMapVolumeSourceArgs
                            {
                                DefaultMode = 420,
                                Name = "argocd-ssh-known-hosts-cm",
                            },
                            Name = "ssh-known-hosts",
                        },
                    },
                },
            },
        },
    });

});

