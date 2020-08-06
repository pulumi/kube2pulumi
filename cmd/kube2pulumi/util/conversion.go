package util

import (
	"fmt"
	"github.com/pulumi/kube2pulumi/pcl2pulumi"
	"github.com/pulumi/kube2pulumi/yaml2pcl"
)

func RunConversion(dirPath string, filePath string, language string) error {
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
		err = pcl2pulumi.Pcl2Pulumi(result, filePath, language)

		fmt.Println(result)

		if err != nil {
			return err
		}
	} else { // dir only
		result, err = yaml2pcl.ConvertDirectory(dirPath)
		err = pcl2pulumi.Pcl2Pulumi(result, fmt.Sprintf("%smain", dirPath), language)
		if err != nil {
			return err
		}
	}
	return nil
}
