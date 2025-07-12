package srt

import (
	"fmt"
	"time"
)

var srtTimeFormat = "15:04:05,000"

// expects a time string in the format "HH:MM:SS,mmm", was already check by the main function
func ParseSRTTime(s string) time.Time {
	t, _ := time.Parse(srtTimeFormat, s)
	return t
}

func ApplyOffset(t time.Time, offset time.Duration) time.Time {
	base := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
	newTime := t.Add(offset)
	if newTime.Before(base) {
		return base
	}
	return newTime
}

func FormatSRTTime(t time.Time) string {
	return fmt.Sprintf("%02d:%02d:%02d,%03d",
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1e6) // the /1e6 is to converts nanoseconds to milliseconds
}
