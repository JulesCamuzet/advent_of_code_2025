package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
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

func dupLinesInt(lines []string) [][]int {
	height := len(lines)
	width := len(lines[0])
	result := make([][]int, height)
	for i := 0; i < height; i++ {
		result[i] = make([]int, width)
		for j := 0; j < width; j++ {
			result[i][j] = 0
		}
	}
	return result
}

func processInput(input string) (int, error) {
	lines := strings.Split(input, "\n")
	linesInt := dupLinesInt(lines)
	height := len(lines)
	width := len(lines[0])
	for i := 0; i < width; i++ {
		if (lines[0][i] == 'S') {
			linesInt[0][i] = 1
			break
		}
	}
	for y := 2; y < height; y += 2 {
		for x := 0; x < width; x++ {
			if (lines[y][x] == '^') {
				linesInt[y][x - 1] += linesInt[y - 2][x]
				linesInt[y][x + 1] += linesInt[y - 2][x]
			} else {
				linesInt[y][x] += linesInt[y - 2][x]
			}
		}
	}
	result := 0
	for x := 0; x < width; x++ {
		result += linesInt[height - 2][x]
	}
	return result, nil
}

func main() {
	input, err := getFileContent()
	if (err != nil) {
		panic(err)
	}
	start := time.Now().UnixMilli()
	answer, err := processInput(input)
	end := time.Now().UnixMilli()
	duration := end - start
	fmt.Printf("The answer is : %d. Found in %d ms.\n", answer, duration)
}
