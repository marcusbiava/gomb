package logs

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestElapsedTime(t *testing.T) {
	var actualFormat string
	var actualV interface{}

	logPrintf = func(format string, v ...any) {
		actualFormat = format
		actualV = v
	}

	start := time.Now()
	time.Sleep(2 * time.Second)

	ElapsedTime(start, "Elapsed time")

	assert.Equal(t, "%s %d hours, %d minutes, and %d seconds.\n", actualFormat)
	assert.ElementsMatch(t, []interface{}{"Elapsed time", 0, 0, 2}, actualV)
}
