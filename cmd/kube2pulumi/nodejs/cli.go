package nodejs

import (
	"github.com/pulumi/kube2pulumi/cmd/kube2pulumi/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:  "typescript",
		Long: "convert k8s yaml to typescript",
		RunE: func(cmd *cobra.Command, args []string) error {
			dirPath := viper.GetString("directory")
			filePath := viper.GetString("manifest")

			return util.VerifyParams(dirPath, filePath, "python")
		}}

	return command
}
