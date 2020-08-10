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

			python, err := util.RunConversion(dirPath, filePath, "python")
			if err != nil {
				return err
			}
			fmt.Printf("Conversion successful! Generated File: %s.py\n", python)

			nodejs, err := util.RunConversion(dirPath, filePath, "nodejs")
			if err != nil {
				return err
			}
			fmt.Printf("Conversion successful! Generated File: %s.ts\n", nodejs)

			dotnet, err := util.RunConversion(dirPath, filePath, "dotnet")
			if err != nil {
				return err
			}
			fmt.Printf("Conversion successful! Generated File: %s.cs\n", dotnet)

			golang, err := util.RunConversion(dirPath, filePath, "go")
			if err != nil {
				return err
			}
			fmt.Printf("Conversion successful! Generated File: %s.go\n", golang)
			return nil
		}}

	return command
}
