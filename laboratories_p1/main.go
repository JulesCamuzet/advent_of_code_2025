package main

import (
	"fmt"
	"os"
	"path/filepath"
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

func processInput(input string) (int, error) {
	lines := strings.Split(input, "\n")
	linesCount := len(lines)
	width := len(lines[0])
	tachyonMap := make([]bool, width)
	for i := 0; i < width; i++ {
		if (lines[0][i] == 'S') {
			tachyonMap[i] = true
		} else {
			tachyonMap[i] = false
		}
	}
	result := 0
	for i := 0; i < linesCount; i++ {
		for j := 0; j < width; j++ {
			if (lines[i][j] == '^' && tachyonMap[j] == true) {
				result++
				tachyonMap[j] = false
				if (j > 0) {
					tachyonMap[j - 1] = true
				}
				if (j < width - 1) {
					tachyonMap[j + 1] = true
				}
			}
		}
	}
	return result, nil
}

func main() {
	input, err := getFileContent()
	if (err != nil) {
		panic(err)
	}
	answer, err := processInput(input)
	fmt.Printf("The answer is : %d\n", answer)
}
