package kube2pulumi

import (
	"github.com/pulumi/kube2pulumi/pkg/pcl2pulumi"
	"github.com/pulumi/kube2pulumi/pkg/yaml2pcl"
	"path/filepath"
)

// Kube2PulumiFile generates an output file containing the converted YAML manifest
// file (filePath) into the specified language and returns the path of the generated
// code file
func Kube2PulumiFile(filePath string, language string) (string, error) {
	pcl, err := yaml2pcl.ConvertFile(filePath)
	if err != nil {
		return "", err
	}
	outPath := getOutputFile(filepath.Dir(filePath), language)
	outFile, err := pcl2pulumi.Pcl2Pulumi(pcl, outPath, language)
	if err != nil {
		return "", err
	}
	return outFile, nil
}

// Kube2PulumiDirectory generates an output file containing the converted directory
// containing YAML manifests (directoryPath) into the specified language and returns
// the path of the generated code file
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
