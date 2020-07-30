package main

import (
	"fmt"
	"kube2pulumi/yaml2pcl"
)

var testData = `

apiVersion: v1
kind: Pod
# this is a test comment
metadata:
  namespace: foo
  name: bar
spec:
  containers:
    - name: nginx
      image: nginx:1.14-alpine
      resources:
        limits:
          memory: 20Mi
          cpu: 0.2

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
