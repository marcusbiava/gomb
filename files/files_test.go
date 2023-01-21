package files

import (
	"github.com/marcusbiava/gomb/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteFile(t *testing.T) {
	fileName := "file1"
	data := []string{"1", "2", "3", "4"}

	WriteFile(fileName, data)

	slice := ReadFile(fileName)

	assert.Equal(t, 4, len(slice))

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
