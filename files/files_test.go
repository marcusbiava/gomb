package files

import (
	"github.com/marcusbiava/gomb/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteFile(t *testing.T) {
	fileName := "file1"
	data := []string{"1", "2", "3", "4"}

	WriteSliceToFile(fileName, data)

	slice := ReadFile(fileName)

	assert.Equal(t, "1,2,3,4", slice)

	RemoveFile(fileName)
}

func TestOpenFileOrCreateFile(t *testing.T) {
	fileName := "file1"
	file := OpenFileOrCreate(fileName)
	stat, err := file.Stat()
	errors.IfAnErrorOccursCallsLogFatal(err, "Stat")

	assert.Equal(t, fileName, stat.Name())

	err = file.Close()
	errors.IfAnErrorOccursCallsLogFatal(err, "Close file")

	RemoveFile(fileName)
}

func TestFileToSlice(t *testing.T) {
	fileName := "file1"
	WriteStringToFile(fileName, "line 1\nline 2\nline 3")
	r := FileToSlice(fileName)

	assert.Equal(t, r[0], "line 1")
	assert.Equal(t, r[1], "line 2")
	assert.Equal(t, r[2], "line 3")

	RemoveFile(fileName)
}
