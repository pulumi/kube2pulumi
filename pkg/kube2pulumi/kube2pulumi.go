package kube2pulumi

import (
	"github.com/pulumi/kube2pulumi/pkg/pcl2pulumi"
	"github.com/pulumi/kube2pulumi/pkg/yaml2pcl"
	"path/filepath"
)

func Kube2PulumiFile(filePathAndName string, language string) (string, error) {
	pcl, err := yaml2pcl.ConvertFile(filePathAndName)
	if err != nil {
		return "", err
	}
	outFile, err := pcl2pulumi.Pcl2Pulumi(pcl, filePathAndName, language)
	if err != nil {
		return "", err
	}
	return outFile, nil
}

func Kube2PulumiDirectory(directoryPath string, language string) (string, error) {
	pcl, err := yaml2pcl.ConvertDirectory(directoryPath)
	if err != nil {
		return "", err
	}
	outPath := getOutputFile(directoryPath, language)
	outFile, err := pcl2pulumi.Pcl2Pulumi(pcl, outPath, language)
	if err != nil {
		return "", err
	}
	return outFile, nil
}

func getOutputFile(dir, language string) string {
	var fName string
	switch language {
	case "typescript":
		fName = "index"
	case "python":
		fName = "__main__"
	case "csharp":
		fName = "Program"
	case "go":
		fName = "main"
	}

	return filepath.Join(dir, fName)
}
