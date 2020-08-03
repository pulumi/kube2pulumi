package main

import (
	"fmt"
	"os"

	"github.com/pulumi/kube2pulumi/cmd/kube2pulumi/python"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	manifestFile string
)

func configureCLI() *cobra.Command {
	rootCmd := &cobra.Command{Use: "kube2pulumi", Long: "foo"}
	rootCmd.AddCommand(python.Command())
	rootCmd.PersistentFlags().StringVarP(&manifestFile, "manifest", "m", "", "manifest file to convert")
	viper.BindPFlag("manifest", rootCmd.PersistentFlags().Lookup("manifest"))

	return rootCmd
}

func main() {
	rootCmd := configureCLI()
	if err := rootCmd.Execute(); err != nil {
		fmt.Errorf("unable to run program")
		os.Exit(1)
	}
}
