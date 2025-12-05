package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getFileContent() (string, error) {
	wd, err := os.Getwd()
	if (err != nil) {
		return "", err
	}
	path := filepath.Join(wd, "input")
	inputBytes, err := os.ReadFile(path)
	if (err != nil) {
		return "", err
	}
	input := string(inputBytes)
	return input, nil
}

func processLine(line string) (int, error) {
	bResult := make([]byte, 12)
	lineLen := len(line)
	previousMaxIndex := -1
	for i := 0; i < 12; i++ {
		maxIndex := previousMaxIndex + 1
		for j := maxIndex; j <= lineLen - (12 - i); j++ {
			if (line[j] > line[maxIndex]) {
				maxIndex = j
			}
		}
		bResult[i] = line[maxIndex]
		previousMaxIndex = maxIndex
	}
	result, err := strconv.Atoi(string(bResult))
	if (err != nil) {
		return 0, err
	}
	return result, nil
}

func processInput(input string) (uint64, error) {
	lines := strings.Split(input, "\n")
	linesCount := len(lines)
	result := (uint64)(0)
	for i := 0; i < linesCount; i++ {
		lineResult, err := processLine(lines[i])
		if (err != nil) {
			return 0, err
		}
		result += (uint64)(lineResult)
	}
	return result, nil
}

func main() {
	input, err := getFileContent()
	if (err != nil) {
		panic(err)
	}
	answer, err := processInput(input)
	if (err != nil) {
		panic(err)
	}
	fmt.Printf("The answer is : %d", answer)
}
