package _go

import (
	"fmt"
	"github.com/pulumi/kube2pulumi/pcl2pulumi"
	"github.com/pulumi/kube2pulumi/yaml2pcl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:  "go",
		Long: "convert k8s yaml to golang",
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath := viper.GetString("manifest")
			if filePath == "" {
				return fmt.Errorf("must specify manifest file")
			}
			result, err := yaml2pcl.ConvertFile(filePath)
			if err != nil {
				return err
			}
			pcl2pulumi.Pcl2Pulumi(result, filePath, "go")
			return nil
		}}

	return command
}