package files

import (
	"github.com/marcusbiava/gomb/errors"
	"github.com/marcusbiava/gomb/slices"
	"os"
	"strings"
)

func WriteFile[T any](path string, slice []T) {
	s := slices.SliceToStringWithoutBracket(slice)
	err := os.WriteFile(path, []byte(s), 0666)
	errors.IfAnErrorOccursCallsLogFatal(err, "WriteFile: "+path)
}

func ReadFile(path string) []string {
	b, err := os.ReadFile(path)
	errors.IfAnErrorOccursCallsLogFatal(err, "ReadFile: "+path)
	s := string(b)
	split := strings.Split(s, ",")
	return split
}
