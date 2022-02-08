package kube2pulumi

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/kube2pulumi/pkg/pcl2pulumi"
	"github.com/pulumi/kube2pulumi/pkg/yaml2pcl"
	"path/filepath"
)

// Kube2PulumiFile generates an output file containing the converted YAML manifest
// file (filePath) into the specified language and returns the path of the generated
// code file
func Kube2PulumiFile(filePath string, outputFile string, language string) (string, hcl.Diagnostics, error) {
	pcl, diags, err := yaml2pcl.ConvertFile(filePath)
	if err != nil {
		return "", diags, err
	}
	outPath := getOutputFile(filepath.Dir(filePath), outputFile, language)
	outFile, err := pcl2pulumi.Pcl2Pulumi(pcl, outPath, language)
	if err != nil {
		return "", diags, err
	}
	return outFile, diags, nil
}

// Kube2PulumiDirectory generates an output file containing the converted directory
// containing YAML manifests (directoryPath) into the specified language and returns
// the path of the generated code file
func Kube2PulumiDirectory(directoryPath string, outputFile string, language string) (string, hcl.Diagnostics, error) {
	pcl, diags, err := yaml2pcl.ConvertDirectory(directoryPath)
	if err != nil {
		return "", diags, err
	}
	outPath := getOutputFile(filepath.Dir(directoryPath), outputFile, language)
	outFile, err := pcl2pulumi.Pcl2Pulumi(pcl, outPath, language)
	if err != nil {
		return "", diags, err
	}
	return outFile, diags, nil
}

func getOutputFile(dir, outputFile, language string) string {
	if outputFile != "" {
		return outputFile
	}
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
