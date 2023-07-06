package files

import (
	"bufio"
	"fmt"
	"github.com/marcusbiava/gomb/errors"
	"github.com/marcusbiava/gomb/slices"
	"os"
)

func WriteStringToFile(path, fileContent string) {
	err := os.WriteFile(path, []byte(fileContent), os.ModePerm)
	errors.IfAnErrorOccursCallsLogFatal(err, "WriteStringToFile: "+path)
}

func WriteSliceToFile[T any](path string, slice []T) {
	fileContent := slices.SliceToStringWithoutBracket(slice)
	WriteStringToFile(path, fileContent)
}

func ReadFile(path string) string {
	b, err := os.ReadFile(path)
	errors.IfAnErrorOccursCallsLogFatal(err, "ReadFile: "+path)
	return string(b)
}

func OpenFileOrCreate(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	errors.IfAnErrorOccursCallsLogFatal(err, "OpenFileOrCreate")
	return file
}

func RemoveFile(fileName string) {
	err := os.Remove(fileName)
	errors.IfAnErrorOccursCallsLogFatal(err, "Remove")
}

func FileToSlice(filename string) []string {
	// Open the file
	file, err := os.Open(filename)
	errors.IfAnErrorOccursCallsLogFatal(err, fmt.Sprintf("os.Open(%s)", filename))
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var lines []string
	// Read each line of the file
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Check for any errors encountered during scanning
	err = scanner.Err()
	errors.IfAnErrorOccursCallsLogFatal(err, "scanner.Err()")

	return lines
}
