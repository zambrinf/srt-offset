package srt

import (
	"testing"
	"time"
)

func TestParseSRTTime(t *testing.T) {
	input := "00:01:30,500"
	expected := time.Date(0, 1, 1, 0, 1, 30, 500*1e6, time.UTC) // 500 milliseconds converted to nanoseconds
	parsed := ParseSRTTime(input)

	if !parsed.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, parsed)
	}
}

func TestFormatSRTTime(t *testing.T) {
	tm := time.Date(0, 1, 1, 0, 2, 10, 120*1e6, time.UTC) // 120 milliseconds converted to nanoseconds
	formatted := FormatSRTTime(tm)
	expected := "00:02:10,120"

	if formatted != expected {
		t.Errorf("Expected %s, got %s", expected, formatted)
	}
}

func TestApplyOffset_Positive(t *testing.T) {
	tm := time.Date(0, 1, 1, 0, 0, 30, 0, time.UTC)
	offset := 2*time.Second + 500*time.Millisecond
	result := ApplyOffset(tm, offset)
	expected := time.Date(0, 1, 1, 0, 0, 32, 500*1e6, time.UTC)

	if !result.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestApplyOffset_Negative(t *testing.T) {
	tm := time.Date(0, 1, 1, 0, 0, 5, 0, time.UTC)
	offset := -10 * time.Second
	result := ApplyOffset(tm, offset)
	expected := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)

	if !result.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
