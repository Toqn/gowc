package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var byteCountFlag string
	var fileName string

	flag.StringVar(&byteCountFlag, "c", "", "File byte count")
	flag.StringVar(&fileName, "f2", "", "File to check")

	flag.Parse()
	
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("File %s does exist", &fileName)
	}

	byteCount, err = os.Read 

	fmt.Printf("%s %s\n", byteCountFlag, fileName)
}
