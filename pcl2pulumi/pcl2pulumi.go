package pcl2pulumi

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// will generate the pulumi code of specified type given the input stream
func GeneratePulumi(pcl string, yamlName string, ext string) {
	pclFile, err := buildTempFile(pcl)
	if err != nil {
		fmt.Println(err)
		return
	}
	//get original file name
	fileName := strings.Split(yamlName, ".")[0]
	convertPulumi(pclFile, fileName, ext)
	//os.Remove(pclFile.Name())
}

// open questions:
// will this work or will the os remove the file before contents are read?
// should os.Remove be called at the end of GeneratePulumi once the final file has been built?
//			refer to line 21
func buildTempFile(pcl string) (*os.File, error) {
	var err error

	tempFile, err := ioutil.TempFile("", "pcl-*.pp")
	if err != nil {
		return nil, err
	}
	// ensure to clean up the file
	defer os.Remove(tempFile.Name())
	// fmt.Println("created the file: ", tempFile.Name())

	// Write to the file
	text := []byte(pcl)
	if _, err = tempFile.Write(text); err != nil {
		log.Fatal("Failed to write to temporary file", err)
		return nil, err
	}

	// Close the file
	if err := tempFile.Close(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return tempFile, err
}

func convertPulumi(tempFile *os.File, newFileName string, fileExt string) {
	// invoke codegen here to convert the file
	// ask Evan/Lee about how to import the modules correctly
}
