package files

import (
	"github.com/marcusbiava/gomb/errors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	fileName := "file1"
	data := []string{"1", "2", "3", "4"}

	WriteFile(fileName, data)

	slice := ReadFile(fileName)

	assert.Equal(t, 4, len(slice))

	err := os.Remove(fileName)
	errors.IfAnErrorOccursCallsLogFatal(err, "Rremove")
}
