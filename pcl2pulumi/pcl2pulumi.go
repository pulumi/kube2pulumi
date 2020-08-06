package pcl2pulumi

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	csgen "github.com/pulumi/pulumi/pkg/v2/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v2/codegen/go"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	tsgen "github.com/pulumi/pulumi/pkg/v2/codegen/nodejs"
	pygen "github.com/pulumi/pulumi/pkg/v2/codegen/python"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

// generates pulumi program for specified type given the input stream
func Pcl2Pulumi(pcl string, outputFilePathAndName string, output string) error {
	pclFile, err := buildTempFile(pcl)
	if err != nil {
		return err
	}
	defer os.Remove(pclFile.Name())

	// get original file name
	dir, fileName := filepath.Split(outputFilePathAndName)
	fileName = strings.Split(fileName, ".")[0]
	err = convertPulumi(pclFile, fmt.Sprintf("%s%s", dir, fileName), output)
	if err != nil {
		return err
	}
	return nil
}

func buildTempFile(pcl string) (*os.File, error) {
	tempFile, err := ioutil.TempFile("", "pcl-*.pp")
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
func convertPulumi(ppFile *os.File, newFileName string, outputLanguage string) error {
	var generateProgram func(program *hcl2.Program) (map[string][]byte, hcl.Diagnostics, error)
	var fileExt string
	switch outputLanguage {
	case "nodejs":
		generateProgram = tsgen.GenerateProgram
		fileExt = ".ts"
	case "python":
		generateProgram = pygen.GenerateProgram
		fileExt = ".py"
	case "dotnet":
		generateProgram = csgen.GenerateProgram
		fileExt = ".cs"
	case "go":
		generateProgram = gogen.GenerateProgram
		fileExt = ".go"
	}

	parser := syntax.NewParser()
	if err := parseFile(parser, ppFile.Name()); err != nil {
		log.Printf("failed to parse %v", ppFile.Name())
		return err
	}
	if len(parser.Diagnostics) != 0 {
		writer := parser.NewDiagnosticWriter(os.Stderr, 0, true)
		err := writer.WriteDiagnostics(parser.Diagnostics)
		if parser.Diagnostics.HasErrors() {
			return err
		}
	}
	program, diags, err := hcl2.BindProgram(parser.Files)
	if err != nil {
		log.Print("failed to bind program: ")
		return err
	}
	if len(diags) != 0 {
		writer := program.NewDiagnosticWriter(os.Stderr, 0, true)
		err := writer.WriteDiagnostics(diags)
		if diags.HasErrors() {
			return err
		}
	}

	files, diags, err := generateProgram(program)
	if err != nil {
		log.Print("failed to generate program: ")
		return err
	}
	if len(diags) != 0 {
		writer := program.NewDiagnosticWriter(os.Stderr, 0, true)
		err = writer.WriteDiagnostics(diags)
		if diags.HasErrors() {
			return err
		}
	}

	var fpaths []string
	for fpath := range files {
		fpaths = append(fpaths, fpath)
	}
	sort.Strings(fpaths)

	outputFileName := fmt.Sprintf("%s%s", newFileName, fileExt)
	for _, p := range fpaths {
		if err := ioutil.WriteFile(outputFileName, files[p], 0600); err != nil {
			log.Printf("failed to write output file %v: ", p)
			return err
		}
	}
	return nil
}

func parseFile(parser *syntax.Parser, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	return parser.ParseFile(f, path.Base(filePath))
}
