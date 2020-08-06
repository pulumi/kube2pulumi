package dotnet

import (
	"github.com/pulumi/kube2pulumi/cmd/kube2pulumi/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:  "C#",
		Long: "convert k8s yaml to C#",
		RunE: func(cmd *cobra.Command, args []string) error {
			dirPath := viper.GetString("directory")
			filePath := viper.GetString("manifest")

			return util.RunConversion(dirPath, filePath, "dotnet")
		}}

	return command
}
