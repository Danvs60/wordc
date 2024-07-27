package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

const maxFileSize = 100 * 1024 * 1024 // 100 MB

func main() {
	// Define flags
	byteSize := flag.Bool("c", false, "Count bytes in file.")
	lineCount := flag.Bool("l", false, "Count lines in file.")
	wordCount := flag.Bool("w", false, "Count words in file.")
	characterCount := flag.Bool("m", false, "Count words in file.")

	// Parse command-line flags
	flag.Parse()
	// Use all if user does not set any
	if !(*byteSize || *lineCount || *wordCount || *characterCount) {
		*wordCount = true
		*byteSize = true
		*lineCount = true
		*characterCount = true
	}

	var filePath string
	var file io.Reader
	if len(flag.Args()) > 0 {
		filePath := flag.Arg(0)

		// ensure input is of reasonable size
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			log.Fatal(err)
		}
		if fileInfo.Size() > maxFileSize {
			log.Fatalf("The file size limit is %v MB", maxFileSize/(1024*1024))
		}

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

			// ensure input is of reasonable size
			if stat.Size() > maxFileSize {
				log.Fatalf("The file size limit is %v MB", maxFileSize/(1024*1024))
			}
		} else {
			log.Fatal("No input file specified and stdin is a terminal")
		}
	}

	reader := bufio.NewReader(file)

	var byteSizeOut, lineCountOut, wordCountOut, characterCountOut int
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				if line != "" {
					lineCountOut++
					bc, wc, cc := getLineCounters(line)
					byteSizeOut += bc
					wordCountOut += wc
					characterCountOut += cc
				}
				break
			}
			log.Fatalf("Error reading line %v", err)
		}
		lineCountOut++
		bc, wc, cc := getLineCounters(line)
		byteSizeOut += bc
		wordCountOut += wc
		characterCountOut += cc
	}

	var output string
	if *byteSize {
		output += fmt.Sprintf("%v ", byteSizeOut)
	}
	if *lineCount {
		output += fmt.Sprintf("%v ", lineCountOut)
	}
	if *wordCount {
		output += fmt.Sprintf("%v ", wordCountOut)
	}
	if *characterCount {
		output += fmt.Sprintf("%v ", characterCountOut)
	}

	fmt.Printf("\t %s %s\n", output, filePath)
}

func getLineCounters(line string) (byteSize, wordCount, characterCount int) {
	byteSize = len(line)
	wordCount = len(strings.Fields(line))
	characterCount = utf8.RuneCountInString(line)
	return
}
