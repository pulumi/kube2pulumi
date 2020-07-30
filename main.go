package main

import (
	"fmt"
	"github.com/pulumi/kube2pulumi/yaml2pcl"
)

var testData = `

apiVersion: v1
kind: Namespace
metadata:
	name: foo

`

/**
converts YAML defined in the testData field to PCL and prints it out
*/
func main() {
	result, err := yaml2pcl.Convert([]byte(testData))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
