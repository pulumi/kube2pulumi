package main

import (
	"fmt"
	"github.com/pulumi/kube2pulumi/pcl2pulumi"
	"github.com/pulumi/kube2pulumi/yaml2pcl"
)

/**
converts YAML defined in the testData field to PCL and writes it to a temp .pp file
*/
func main() {
	fileName := "conversionTest.yaml"
	result, err := yaml2pcl.ConvertFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// open comments:
	// passing the correct extension should be taken care of by the CLI
	// should a pointer to the proper codegen tool also be passed.... e.g. something similar to
	/**
	var generateProgram func(program *hcl2.Program) (map[string][]byte, hcl.Diagnostics, error)
	switch target {
	case "dotnet":
		generateProgram = csgen.GenerateProgram
	case "go":
		generateProgram = gogen.GenerateProgram
	case "nodejs":
		generateProgram = tsgen.GenerateProgram
	case "python":
		generateProgram = pygen.GenerateProgram
	default:
		flag.Usage()
		os.Exit(1)
	}
	*/
	pcl2pulumi.GeneratePulumi(result, fileName, ".ts")
}
