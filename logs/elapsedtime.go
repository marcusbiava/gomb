package logs

import (
	"log"
	"time"
)

var logPrintf = log.Printf

func ElapsedTime(start time.Time, startMessage string) {
	end := time.Now()
	duration := end.Sub(start)
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	logPrintf("%s %d hours, %d minutes, and %d seconds.\n", startMessage, hours, minutes, seconds)
}
