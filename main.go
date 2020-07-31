package main

import (
	"fmt"
	"github.com/pulumi/kube2pulumi/yaml2pcl"
)

/**
converts YAML defined in the testData field to PCL and writes it to a temp .pp file
*/
func main() {
	result, err := yaml2pcl.ConvertFile("conversionTest.yaml")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
