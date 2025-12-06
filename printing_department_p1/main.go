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

func countCloseRolls(lines []string, x, y, linesCount, columnsCount int) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			isOutside :=  x + dx < 0 || y + dy < 0 || x + dx >= columnsCount || y + dy >= linesCount
			if (!isOutside) {
				isAdjRoll := (dx != 0 || dy != 0) && lines[y + dy][x + dx] == '@'
				if (isAdjRoll) {
					count++
				}
			}
		}
	}
	return count
}

func processInput(input string) int {
	lines := strings.Split(input, "\n")
	linesCount := len(lines)
	columnsCount := len(lines[0])
	result := 0
	for y := 0; y < linesCount; y++ {
		for x := 0; x < columnsCount; x++ {
			if (lines[y][x] == '@') {
				if (countCloseRolls(lines, x, y, linesCount, columnsCount) < 4) {
					result++
				}
			}
		}
	}
	return result
}

func main() {
	input, err := getFileContent()
	if (err != nil) {
		panic(err)
	}
	answer := processInput(input)
	fmt.Printf("The answer is : %d", answer)
}
