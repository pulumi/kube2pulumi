package csharp

import (
	"fmt"

	"github.com/pulumi/kube2pulumi/cmd/kube2pulumi/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:  "csharp",
		Long: "convert k8s yaml to C#",
		RunE: func(cmd *cobra.Command, args []string) error {
			dirPath := viper.GetString("directory")
			filePath := viper.GetString("file")
			result, err := util.RunConversion(dirPath, filePath, "csharp")
			if err != nil {
				return err
			}
			fmt.Printf("Conversion successful! Generated File: %s\n", result)
			return nil
		}}

	return command
}
