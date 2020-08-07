package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pulumi/kube2pulumi/pkg/pcl2pulumi"
	"github.com/pulumi/kube2pulumi/pkg/yaml2pcl"
)

func RunConversion(dirPath string, filePath string, language string) (string, error) {
	if filePath == "" && dirPath == "" {
		path, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("unable to detect working directory, must specify a path for a file or directory\n")
		}
		dirPath = path
	}
	if filePath != "" && dirPath != "" {
		return "", fmt.Errorf("must specify EITHER a path for a file or directory, not both\n")
	}
	var result string
	var outPath string
	var err error
	// filepath only
	var fileName string
	if filePath != "" {
		fileName = getOutputFile(filepath.Dir(filePath), language)
		result, err = yaml2pcl.ConvertFile(filePath)
		outPath, err = pcl2pulumi.Pcl2Pulumi(result, fileName, language)
		if err != nil {
			return "", err
		}
	} else { // dir only
		result, err = yaml2pcl.ConvertDirectory(dirPath)
		if err != nil {
			return "", err
		}
		fileName = getOutputFile(dirPath, language)
		outPath, err = pcl2pulumi.Pcl2Pulumi(result, fileName, language)
		if err != nil {
			return "", err
		}
	}
	return outPath, nil
}

func getOutputFile(dir, language string) string {
	var fName string
	switch language {
	case "nodejs":
		fName = "index"
	case "python":
		fName = "__main__"
	case "dotnet":
		fName = "Program"
	case "go":
		fName = "main"
	}

	return filepath.Join(dir, fName)
}
