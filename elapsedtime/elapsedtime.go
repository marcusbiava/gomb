package elapsedtime

import "time"

func MeasureExecutionTime(fn func()) time.Duration {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	return elapsed
}
