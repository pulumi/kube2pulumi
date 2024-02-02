package pcl2pulumi

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2"
	javagen "github.com/pulumi/pulumi-java/pkg/codegen/java"
	csgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2/syntax"
	tsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	"github.com/pulumi/pulumi/pkg/v3/codegen/pcl"
	pygen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
)

// generates pulumi program for specified language type given the input stream (pcl)
func Pcl2Pulumi(pcl string, outputFilePathAndName string, language string) (string, error) {
	pclFile, err := buildTempFile(pcl)
	if err != nil {
		return "", err
	}
	defer os.Remove(pclFile.Name())

	// get original file name
	dir, fileName := filepath.Split(outputFilePathAndName)
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	outPath := fmt.Sprintf("%s%s", dir, fileName)
	outPath, err = convertPulumi(pclFile, outPath, language)
	if err != nil {
		return "", err
	}
	return outPath, nil
}

func buildTempFile(pcl string) (*os.File, error) {
	tempFile, err := os.CreateTemp("", "pcl-*.pp")
	if err != nil {
		return nil, err
	}

	//Write to the file
	text := []byte(pcl)
	if _, err = tempFile.Write(text); err != nil {
		return nil, err
	}

	// Close the file
	if err := tempFile.Close(); err != nil {
		return nil, err
	}

	return tempFile, err
}

// converts .pp file directly in the same directory as the input file
func convertPulumi(ppFile *os.File, newFileName string, outputLanguage string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("unable to convert program:", r)
		}
	}()

	var generateProgram func(program *pcl.Program) (map[string][]byte, hcl.Diagnostics, error)
	var fileExt string
	switch outputLanguage {
	case "typescript", "javascript":
		generateProgram = tsgen.GenerateProgram
		fileExt = ".ts"
	case "python":
		generateProgram = pygen.GenerateProgram
		fileExt = ".py"
	case "csharp":
		generateProgram = csgen.GenerateProgram
		fileExt = ".cs"
	case "java":
		generateProgram = javagen.GenerateProgram
		fileExt = ".java"
	case "go":
		generateProgram = gogen.GenerateProgram
		fileExt = ".go"
	default:
		return "", fmt.Errorf("language input is invalid\n")
	}

	parser := syntax.NewParser()
	if err := parseFile(parser, ppFile.Name()); err != nil {
		log.Printf("failed to parse %v", ppFile.Name())
		return "", err
	}
	if len(parser.Diagnostics) != 0 {
		var diagOutput bytes.Buffer
		writer := parser.NewDiagnosticWriter(&diagOutput, 0, true)
		err := writer.WriteDiagnostics(parser.Diagnostics)
		if parser.Diagnostics.HasErrors() {
			return "", fmt.Errorf(diagOutput.String())
		}
		if err != nil {
			return "", err
		}
	}
	// Enable pcl.AllowMissingProperties to allow converting Kubernetes objects that don't
	// specify all required fields.
	program, diags, err := pcl.BindProgram(parser.Files, pcl.AllowMissingProperties)
	if err != nil {
		return "", err
	}
	if len(diags) != 0 {
		var diagOutput bytes.Buffer
		writer := program.NewDiagnosticWriter(&diagOutput, 0, true)
		err := writer.WriteDiagnostics(diags)
		if diags.HasErrors() {
			return "", fmt.Errorf(diagOutput.String())
		}
		if err != nil {
			return "", err
		}
	}

	files, diags, err := generateProgram(program)
	if err != nil {
		log.Print("failed to generate program: ")
		return "", err
	}
	if len(diags) != 0 {
		var diagOutput bytes.Buffer
		writer := program.NewDiagnosticWriter(&diagOutput, 0, true)
		err = writer.WriteDiagnostics(diags)
		if diags.HasErrors() {
			return "", fmt.Errorf(diagOutput.String())
		}
		if err != nil {
			return "", err
		}
	}

	var fpaths []string
	for fpath := range files {
		fpaths = append(fpaths, fpath)
	}
	sort.Strings(fpaths)

	outputFileName := fmt.Sprintf("%s%s", newFileName, fileExt)
	for _, p := range fpaths {
		if err := os.WriteFile(outputFileName, files[p], 0600); err != nil {
			log.Printf("failed to write output file %v: ", p)
			return "", err
		}
	}
	return outputFileName, nil
}

func parseFile(parser *syntax.Parser, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	return parser.ParseFile(f, path.Base(filePath))
}
