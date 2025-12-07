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

func processInput(input string) (int, error) {
	lines := strings.Split(input, "\n")
	linesCount := len(lines)
	splitLines := make([][]string, linesCount)
	for i := 0; i < linesCount; i++ {
		splitLines[i] = strings.FieldsFunc(lines[i], func(c rune) bool {
			return c == ' '
		})
	}
	columnsCount := len(splitLines[0])
	result := 0
	for col := 0; col < columnsCount; col++ {
		sign := splitLines[linesCount - 1][col]
		switch (sign) {
			case "*":
				currResult := 1
				for line := 0; line < linesCount - 1; line++ {
					num, err := strconv.Atoi(splitLines[line][col])
					if (err != nil) {
						return 0, err
					}
					currResult *= num
				}
				result += currResult
			case "+":
				currResult := 0
				for line := 0; line < linesCount - 1; line++ {
					num, err := strconv.Atoi(splitLines[line][col])
					if (err != nil) {
						return 0, err
					}
					currResult += num
				}
				result += currResult
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
	if (err != nil) {
		panic(err)
	}
	fmt.Printf("The answer is : %d\n", answer)
}
