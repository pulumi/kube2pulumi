package main

import (
	"fmt"
	"os"

	"github.com/pulumi/kube2pulumi/cmd/kube2pulumi/dotnet"
	_go "github.com/pulumi/kube2pulumi/cmd/kube2pulumi/go"
	"github.com/pulumi/kube2pulumi/cmd/kube2pulumi/nodejs"
	"github.com/pulumi/kube2pulumi/cmd/kube2pulumi/python"
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
	rootCmd.AddCommand(nodejs.Command())
	rootCmd.AddCommand(dotnet.Command())

	rootCmd.PersistentFlags().StringVarP(&manifestFile, "file", "f", "", "manifest file to convert")
	viper.BindPFlag("manifest", rootCmd.PersistentFlags().Lookup("manifest"))

	rootCmd.PersistentFlags().StringVarP(&directoryPath, "directory", "d", "", "file path for directory to convert")
	viper.BindPFlag("directory", rootCmd.PersistentFlags().Lookup("directory"))

	return rootCmd
}

func main() {
	rootCmd := configureCLI()
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("unable to run program %v\n", err)
		os.Exit(1)
	}
}
