package all

import (
	"fmt"
	"github.com/pulumi/kube2pulumi/cmd/kube2pulumi/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:  "all",
		Long: "convert k8s yaml to all supported languages",
		RunE: func(cmd *cobra.Command, args []string) error {
			dirPath := viper.GetString("directory")
			filePath := viper.GetString("file")
			outputFile := viper.GetString("outputFile")

			python, err := util.RunConversion(dirPath, filePath, outputFile, "python")
			if err != nil {
				return err
			}
			fmt.Printf("Conversion successful! Generated File: %s\n", python)

			typescript, err := util.RunConversion(dirPath, filePath, outputFile, "typescript")
			if err != nil {
				return err
			}
			fmt.Printf("Conversion successful! Generated File: %s\n", typescript)

			csharp, err := util.RunConversion(dirPath, filePath, outputFile, "csharp")
			if err != nil {
				return err
			}
			fmt.Printf("Conversion successful! Generated File: %s\n", csharp)

			java, err := util.RunConversion(dirPath, filePath, outputFile, "java")
			if err != nil {
				return err
			}
			fmt.Printf("Conversion successful! Generated File: %s\n", java)

			golang, err := util.RunConversion(dirPath, filePath, outputFile, "go")
			if err != nil {
				return err
			}
			fmt.Printf("Conversion successful! Generated File: %s\n", golang)
			return nil
		}}

	return command
}
