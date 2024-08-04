package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "f", "", "File to check")

	byteCountFlag := flag.Bool("c", false, "File byte count")

	flag.Parse()

	if fileName == "" {
		fmt.Println("No file specified.")
		os.Exit(1)
	}
	absPath, err := filepath.Abs(fileName)
	if err != nil {
		fmt.Printf("Error getting absolut file path for file %s. Error: %s", fileName, err)
	}

	file, err := os.Open(absPath)
	if err != nil {
		fmt.Printf("File %s does exist: %s", fileName, err)
	}

	if *byteCountFlag {
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Printf("Error getting file info: %s", err)
		}

		byteCount := fileInfo.Size()
		fmt.Printf("Bytes: %d\n", byteCount)
	}

	if !*byteCountFlag {
		fmt.Println("No flags specified")
		flag.Usage()
		os.Exit(0)
	}
}
