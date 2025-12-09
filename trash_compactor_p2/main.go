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

func processOperation(prev int, num int, op byte) int {
	switch (op) {
		case '+':
			return prev + num
		case '*':
			return prev * num
		default:
			return prev
	}
}

func processInput(input string) (int, error) {
	lines := strings.Split(input, "\n")
	height := len(lines) - 1
	width := len(lines[0]) - 1
	result := 0
	signIndex := width
	for lines[height][signIndex] == ' ' {
		signIndex--
	}
	sign := lines[height][signIndex]
	currResult := 0
	if (sign == '*') {
		currResult = 1
	}
	for x := width; x >= 0; x-- {
		if lines[0][x] == ' ' && lines[1][x] == ' ' && lines[2][x] == ' ' && lines[3][x] == ' ' {
			result += currResult
			currResult = 0
			signIndex--
			for signIndex > 0 && lines[height][signIndex] == ' ' {
				signIndex--
			}
			sign = lines[height][signIndex]
			currResult = 0
			if (sign == '*') {
				currResult = 1
			}
			continue
		}
		currStrNum := ""
		for y := 0; y < height; y++ {
			if (lines[y][x] != ' ') {
				currStrNum += string(lines[y][x])
			}
		}
		num, err := strconv.Atoi(currStrNum)
		if (err != nil) {
			return result, err
		}
		currResult = processOperation(currResult, num, sign)
	}
	result += currResult
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
