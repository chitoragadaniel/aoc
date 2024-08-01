package main

import (
	"bufio"
	"fmt"
	"os"
)

func printP1[T any](v T) {
	fmt.Printf("Result of part 1: %v\n", v)
}

func printP2[T any](v T) {
	fmt.Printf("Result of part 2: %v\n", v)
}

func scanLines(path string) []string {
	readFile, _ := os.Open(path)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines
}
