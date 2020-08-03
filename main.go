package main

import (
	"fmt"
	"github.com/pulumi/kube2pulumi/pcl2pulumi"
	"github.com/pulumi/kube2pulumi/yaml2pcl"
)

// converts a single YAML file (@filepath) to a various pulumi files
// ("nodejs", "python", "dotnet", "go") in the same directory
func main() {
	filePath := "conversionTest.yaml"
	result, err := yaml2pcl.ConvertFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// output format options are "nodejs", "python", "dotnet", "go"
	pcl2pulumi.Pcl2Pulumi(result, filePath, "nodejs")
	pcl2pulumi.Pcl2Pulumi(result, filePath, "python")
	pcl2pulumi.Pcl2Pulumi(result, filePath, "dotnet")
	pcl2pulumi.Pcl2Pulumi(result, filePath, "go")
}
