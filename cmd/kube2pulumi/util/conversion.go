package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pulumi/kube2pulumi/pcl2pulumi"
	"github.com/pulumi/kube2pulumi/yaml2pcl"
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
	if filePath != "" {
		result, err = yaml2pcl.ConvertFile(filePath)
		outPath, err = pcl2pulumi.Pcl2Pulumi(result, filePath, language)
		if err != nil {
			return "", err
		}
	} else { // dir only
		result, err = yaml2pcl.ConvertDirectory(dirPath)
		if err != nil {
			return "", err
		}
		outPath, err = pcl2pulumi.Pcl2Pulumi(result, filepath.Join(dirPath, "main"), language)
		if err != nil {
			return "", err
		}
	}
	return outPath, nil
}
