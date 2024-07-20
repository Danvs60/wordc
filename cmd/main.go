package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	// Define flags
	byteSize := flag.Bool("c", false, "Count bytes in file.")
	lineCount := flag.Bool("l", false, "Count lines in file.")
	wordCount := flag.Bool("w", false, "Count words in file.")
	characterCount := flag.Bool("m", false, "Count words in file.")

	var filePath string

	// Parse command-line flags
	flag.Parse()

	var file io.Reader
	if len(flag.Args()) > 0 {
		filePath := flag.Arg(0)
		f, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("Error opening file %s: %v", filePath, err)
		}
		defer f.Close()
		file = f
	} else {
		// Use stdin if no file argument is provided
		stat, err := os.Stdin.Stat()
		if err != nil {
			log.Fatalf("Error getting stdin stat: %v", err)
		}
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			file = os.Stdin
		} else {
			log.Fatal("No input file specified and stdin is a terminal")
		}
	}

	// Read content from the file or stdin into a buffer
	var buf strings.Builder
	_, err := io.Copy(&buf, file)
	if err != nil {
		log.Fatalf("Error reading from input: %v", err)
	}

	content := buf.String()
	// Process content as needed
	_ = content // Placeholder to prevent "declared and not used" error
	if !(*byteSize || *lineCount || *wordCount || *characterCount) {
		*wordCount = true
		*byteSize = true
		*lineCount = true
		*characterCount = true
	}

	var output string
	// process flags
	if *byteSize {
		output += fmt.Sprintf("%v ", len(content))
	}
	if *lineCount {
		output += fmt.Sprintf("%v ", strings.Count(content, "\n"))
	}
	if *wordCount {
		output += fmt.Sprintf("%v ", len(strings.Fields(content)))
	}
	if *characterCount {
		output += fmt.Sprintf("%v ", utf8.RuneCountInString(content))
	}

	fmt.Printf("\t %s %s\n", output, filePath)
}
