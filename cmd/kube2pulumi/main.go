package main

import (
	"fmt"
	"github.com/pulumi/kube2pulumi/cmd/kube2pulumi/all"
	"os"

	"github.com/pulumi/kube2pulumi/cmd/kube2pulumi/csharp"
	_go "github.com/pulumi/kube2pulumi/cmd/kube2pulumi/go"
	"github.com/pulumi/kube2pulumi/cmd/kube2pulumi/python"
	"github.com/pulumi/kube2pulumi/cmd/kube2pulumi/typescript"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	manifestFile  string
	directoryPath string
)

func configureCLI() *cobra.Command {
	rootCmd := &cobra.Command{Use: "kube2pulumi", Long: "converts input files to desired output language"}

	// 4 commands for the distinct languages
	rootCmd.AddCommand(python.Command())
	rootCmd.AddCommand(_go.Command())
	rootCmd.AddCommand(typescript.Command())
	rootCmd.AddCommand(csharp.Command())
	rootCmd.AddCommand(all.Command())

	rootCmd.PersistentFlags().StringVarP(&manifestFile, "file", "f", "", "YAML file to convert")
	viper.BindPFlag("file", rootCmd.PersistentFlags().Lookup("file"))

	rootCmd.PersistentFlags().StringVarP(&directoryPath, "directory", "d", "", "file path for directory to convert")
	viper.BindPFlag("directory", rootCmd.PersistentFlags().Lookup("directory"))

	return rootCmd
}

func main() {
	rootCmd := configureCLI()
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("unable to run program: %v\n", err)
		os.Exit(1)
	}
}
