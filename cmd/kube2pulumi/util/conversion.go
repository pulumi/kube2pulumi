package util

import (
	"fmt"
	"github.com/pulumi/kube2pulumi/pkg/kube2pulumi"
	"os"
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
	var outPath string
	var err error
	// filepath only
	if filePath != "" {
		outPath, err = kube2pulumi.Kube2PulumiFile(filePath, language)
		if err != nil {
			return "", err
		}
	} else { // dir only
		outPath, err = kube2pulumi.Kube2PulumiDirectory(dirPath, language)
		if err != nil {
			return "", err
		}
	}
	return outPath, nil
}
