package main

import (
	"errors"
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

func processLine(line string, currentNum int) (int, error) {
	var newCurrentNum int

	lineLen := len(line)
	if (len(line) < 2) {
		return currentNum, errors.New("Line too short.")
	}
	direction := line[0]
	strCount := line[1:lineLen]
	count, err := strconv.Atoi(strCount)
	if (err != nil) {
		return currentNum, errors.New("Error while converting count.")
	}
	switch direction {
		case 'R':
			newCurrentNum = currentNum + count
		case 'L':
			newCurrentNum = currentNum - count
		default:
			return currentNum, errors.New("Wrong direction.")
	}
	newCurrentNum %= 100
	return newCurrentNum, nil
}

func processInput(input string) (int, error) {
	var currentNum int
	var err error
	lines := strings.Split(input, "\n")
	linesLen := len(lines)
	answer := 0
	currentNum = 50

	for i := 0; i < linesLen; i++ {
		currentNum, err = processLine(lines[i], currentNum)
		if (err != nil) {
			return answer, err
		}
		if (currentNum == 0) {
			answer++
		}
	}
	return answer, nil
}

func main()  {
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
