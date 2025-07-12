package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/zambrinf/srt-offset/srt"
)

func main() {
	inputPath := flag.String("i", "", "Input .srt file")
	outputPath := flag.String("o", "", "Output .srt file (optional)")
	offsetSec := flag.Float64("offset", 0.0, "Offset in seconds (e.g., -1.25 or 2.75)")

	flag.Parse()

	if *inputPath == "" {
		fmt.Println("Usage: srt-offset -i <file.srt> -offset <seconds> [-o <output.srt>]")
		os.Exit(1)
	}

	offset := time.Duration(*offsetSec * float64(time.Second))

	inputFile, err := os.Open(*inputPath)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	timeLineRegex := regexp.MustCompile(`^(\d{2}:\d{2}:\d{2},\d{3})\s-->\s(\d{2}:\d{2}:\d{2},\d{3})$`)

	var outputLines []string

	for scanner.Scan() {
		line := scanner.Text()
		matches := timeLineRegex.FindStringSubmatch(line)

		if matches != nil {
			startTime := srt.ParseSRTTime(matches[1])
			endTime := srt.ParseSRTTime(matches[2])

			newStart := srt.ApplyOffset(startTime, offset)
			newEnd := srt.ApplyOffset(endTime, offset)

			outputLines = append(outputLines, fmt.Sprintf("%s --> %s", srt.FormatSRTTime(newStart), srt.FormatSRTTime(newEnd)))
		} else {
			outputLines = append(outputLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	if *outputPath == "" {
		writeToFile(*inputPath, outputLines)
	} else {
		writeToFile(*outputPath, outputLines)
	}
}

func writeToFile(path string, lines []string) {
	outputFile, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			os.Exit(1)
		}
	}
	writer.Flush()
}
