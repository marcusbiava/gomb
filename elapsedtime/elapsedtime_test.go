package elapsedtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func exampleFunction() {
	time.Sleep(2 * time.Second)
}

func TestMeasureExecutionTime(t *testing.T) {
	expectedTime := 2 * time.Second
	maxDelta := 100 * time.Millisecond

	elapsedTime := MeasureExecutionTime(exampleFunction)

	assert.InDelta(t, expectedTime, elapsedTime, float64(maxDelta), "Execution time must be within the margin of error")
}
