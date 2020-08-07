package python

import (
	"fmt"

	"github.com/pulumi/kube2pulumi/cmd/kube2pulumi/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:  "python",
		Long: "convert k8s yaml to python",
		RunE: func(cmd *cobra.Command, args []string) error {
			dirPath := viper.GetString("directory")
			filePath := viper.GetString("file")
			result, err := util.RunConversion(dirPath, filePath, "python")
			if err != nil {
				return err
			}
			fmt.Printf("Conversion successful! Generated File: %s.py\n", result)
			return nil
		}}
	return command
}
