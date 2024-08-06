package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "f", "", "File to check")

	byteCountFlag := flag.Bool("c", false, "File byte count")
	lineCountFlag := flag.Bool("l", false, "File line count")
	wordCountFlag := flag.Bool("w", false, "File word count")
	charCountFlag := flag.Bool("m", false, "File character count")
	helpFlag := flag.Bool("h", false, "Help")

	flag.Parse()

	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	if fileName == "" {
		fmt.Println("No file specified.")
		os.Exit(1)
	}

	if !*byteCountFlag && !*lineCountFlag && !*wordCountFlag && !*charCountFlag {
		fmt.Println("No flags specified")
		flag.Usage()
		os.Exit(0)
	}

	absPath, err := filepath.Abs(fileName)
	if err != nil {
		fmt.Printf("Error getting absolut file path for file %s. Error: %s", fileName, err)
	}

	file, err := os.Open(absPath)
	if err != nil {
		fmt.Printf("File %s does exist: %s", fileName, err)
	}
	defer file.Close()

	if !*byteCountFlag && !*lineCountFlag && !*wordCountFlag && !*charCountFlag {
		*byteCountFlag = true
		*lineCountFlag = true
		*wordCountFlag = true
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Error getting file info: %s", err)
	}

	if *byteCountFlag {
		byteCount := fileInfo.Size()
		fmt.Printf("%d ", byteCount)
	}

	if *lineCountFlag {
		scanner := bufio.NewScanner(file)
		lineCount := 0
		for scanner.Scan() {
			lineCount++
		}

		fmt.Printf("%d ", lineCount)
	}

	if *wordCountFlag {
		_, err := file.Seek(0, 0)
		if err != nil {
			fmt.Printf("Error seeking start of file: %s\n", err)
		}
		scanner := bufio.NewScanner(file)
		wordCount := 0

		for scanner.Scan() {
			words := strings.Fields(scanner.Text())
			wordCount += len(words)
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Encountered error while reading file: %s", err)
		}

		fmt.Printf("%d ", wordCount)
	}

	if *charCountFlag {
		_, err := file.Seek(0, 0)
		if err != nil {
			fmt.Printf("Error seeking start of file: %s\n", err)
		}
		readFile, err := os.ReadFile(absPath)
		if err != nil {
			fmt.Printf("Error reading file to count chars: %s\n", err)
		}

		fileString := string(readFile)

		charCount := utf8.RuneCountInString(fileString)

		fmt.Printf("%d ", charCount)
	}

	fmt.Printf("%s\n", fileName)
}
