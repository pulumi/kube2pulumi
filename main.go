package main

import (
	"fmt"
	"github.com/pulumi/kube2pulumi/yaml2pcl"
)

// converts a single YAML file (@filepath) to a various pulumi files
// ("nodejs", "python", "dotnet", "go") in the same directory
func main() {
	filePath := "testdata/k8sOperator/operator.yaml"
	result, err := yaml2pcl.ConvertFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	//expected, err := ioutil.ReadFile("testdata/expK8sOperator.pp")
	//
	////output format options are "nodejs", "python", "dotnet", "go"
	//err = pcl2pulumi.Pcl2Pulumi(string(expected), filePath, "python")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
}
