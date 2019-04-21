package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	GitVersion string
)

func isFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	cmdLine := ReadCommandLine()

	if len(cmdLine.FilePath) == 0 || !isFileExists(cmdLine.FilePath) ||
		len(cmdLine.TemplatePath) == 0 || !isFileExists(cmdLine.TemplatePath) {

		flag.Usage()
		os.Exit(-1)
	}

	if err := Process(cmdLine); err != nil {
		flag.Usage()
		fmt.Printf("Failed to process file: %+v\n", err)
		os.Exit(-1)
	}
	fmt.Println("Cart parsing succeeded!")
}
