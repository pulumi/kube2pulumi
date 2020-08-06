package util

import (
	"fmt"
	"github.com/pulumi/kube2pulumi/pcl2pulumi"
	"github.com/pulumi/kube2pulumi/yaml2pcl"
)

func VerifyParams(dirPath string, filePath string, language string) error {
	if filePath == "" && dirPath == "" {
		return fmt.Errorf("must specify a path for a file or directory\n")
	}
	if filePath != "" && dirPath != "" {
		return fmt.Errorf("must specify EITHER a path for a file or directory, not both\n")
	}
	var result string
	var err error
	// filepath only
	if filePath != "" {
		result, err = yaml2pcl.ConvertFile(filePath)
	} else { // dir only
		result, err = yaml2pcl.ConvertDirectory(fmt.Sprintf("%smain", dirPath))
	}
	if err != nil {
		return err
	}
	err = pcl2pulumi.Pcl2Pulumi(result, filePath, language)
	if err != nil {
		return err
	}
	return nil
}
