package _go

import (
	"fmt"

	"github.com/pulumi/kube2pulumi/cmd/kube2pulumi/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:  "go",
		Long: "convert k8s yaml to golang",
		RunE: func(cmd *cobra.Command, args []string) error {
			dirPath := viper.GetString("directory")
			filePath := viper.GetString("file")
			outputFile := viper.GetString("outputFile")
			result, err := util.RunConversion(dirPath, filePath, outputFile, "go")
			if err != nil {
				return err
			}
			fmt.Printf("Conversion successful! Generated File: %s\n", result)
			return nil
		}}

	return command
}
