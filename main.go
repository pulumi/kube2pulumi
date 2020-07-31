package main

import (
	"fmt"
	"github.com/pulumi/kube2pulumi/yaml2pcl"
	"io/ioutil"
	"log"
	"os"
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
	}

	tempFile, err := ioutil.TempFile("", "test-*.pp")
	if err != nil {
		panic(err)
	}

	// ensure to clean up the file
	defer os.Remove(tempFile.Name())
	fmt.Println("created the file: ", tempFile.Name())

	// Write to the file
	text := []byte(result)
	if _, err = tempFile.Write(text); err != nil {
		log.Fatal("Failed to write to temporary file", err)
	}

	// test printing out file contents
	out, err := ioutil.ReadFile(tempFile.Name())
	fmt.Println(string(out))

	// Close the file
	if err := tempFile.Close(); err != nil {
		log.Fatal(err)
	}
}
