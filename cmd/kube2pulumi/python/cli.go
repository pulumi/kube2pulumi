package python

import (
	"fmt"
	"github.com/pulumi/kube2pulumi/yaml2pcl"

	"github.com/pulumi/kube2pulumi/pcl2pulumi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:  "python",
		Long: "convert k8s yaml to python",
		RunE: func(cmd *cobra.Command, args []string) error {
			dirPath := viper.GetString("directory")
			filePath := viper.GetString("manifest")
			if filePath == "" && dirPath == "" {
				return fmt.Errorf("must specify a path for a file or directory\n")
			}
			if filePath != "" && dirPath != "" {
				return fmt.Errorf("must specify EITHER a path for a file or directory, not both\n")
			}
			var result string
			var err error
			// filepath only
			if filePath != "" {
				result, err = yaml2pcl.ConvertFile(filePath)
			} else { // dir only
				result, err = yaml2pcl.ConvertDirectory(fmt.Sprintf("%smain", dirPath))
			}
			if err != nil {
				return err
			}
			err = pcl2pulumi.Pcl2Pulumi(result, filePath, "python")
			if err != nil {
				return err
			}
			return nil
		}}
	return command
}
