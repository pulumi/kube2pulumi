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
	"sort"
	"strings"
)

// generates pulumi program for specified type given the input stream
func PclString2Pulumi(pcl string, yamlName string, output string) {
	pclFile, err := buildTempFile(pcl)
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		get original file name
	*/
	fileName := strings.Split(yamlName, ".")[0]
	ConvertPulumi(pclFile, fileName, output)
	err = os.Remove(pclFile.Name()) // delete temporary .pp file
	if err != nil {
		fmt.Println(err)
	}
}

func buildTempFile(pcl string) (*os.File, error) {
	var err error

	tempFile, err := ioutil.TempFile("", "pcl-*.pp")
	if err != nil {
		return nil, err
	}

	/*
		Write to the file
	*/
	text := []byte(pcl)
	if _, err = tempFile.Write(text); err != nil {
		log.Fatal("Failed to write to temporary file", err)
		return nil, err
	}

	/*
		Close the file
	*/
	if err := tempFile.Close(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return tempFile, err
}

// converts .pp file directly in the same directory as the input file
func ConvertPulumi(ppFile *os.File, newFileName string, outputLanguage string) {
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
		log.Fatalf("failed to parse %v: %v", ppFile.Name(), err)
	}
	if len(parser.Diagnostics) != 0 {
		writer := parser.NewDiagnosticWriter(os.Stderr, 0, true)
		writer.WriteDiagnostics(parser.Diagnostics)
		if parser.Diagnostics.HasErrors() {
			os.Exit(1)
		}
	}
	program, diags, err := hcl2.BindProgram(parser.Files)
	if err != nil {
		log.Fatalf("failed to bind program: %v", err)
	}
	if len(diags) != 0 {
		writer := program.NewDiagnosticWriter(os.Stderr, 0, true)
		writer.WriteDiagnostics(diags)
		if diags.HasErrors() {
			os.Exit(1)
		}
	}

	files, diags, err := generateProgram(program)
	if err != nil {
		log.Fatalf("failed to generate program: %v", err)
	}
	if len(diags) != 0 {
		writer := program.NewDiagnosticWriter(os.Stderr, 0, true)
		writer.WriteDiagnostics(diags)
		if diags.HasErrors() {
			os.Exit(1)
		}
	}

	fpaths := make([]string, 0, len(files))
	for fpath := range files {
		fpaths = append(fpaths, fpath)
	}
	sort.Strings(fpaths)

	outputFileName := fmt.Sprintf("%s%s", newFileName, fileExt)
	for _, p := range fpaths {
		if err := ioutil.WriteFile(outputFileName, files[p], 0600); err != nil {
			log.Fatalf("failed to write output file %v: %v", p, err)
		}
	}
}

func parseFile(parser *syntax.Parser, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	return parser.ParseFile(f, path.Base(filePath))
}
