package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func abs(n int) int {
	if (n < 0) {
		return n * -1
	} else {
		return n
	}
}

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


func processLine(line string, currentNum int) (int, int,  error) {
	var newCurrentNum int

	lineLen := len(line)
	if (len(line) < 2) {
		return currentNum, 0, errors.New("Line too short.")
	}
	direction := line[0]
	strCount := line[1:lineLen]
	count, err := strconv.Atoi(strCount)
	if (err != nil) {
		return currentNum, 0, errors.New("Error while converting count.")
	}
	zeroMetCount := 0
	newCurrentNum = currentNum
	switch direction {
		case 'R':
			for i := 0; i < count; i++ {
				newCurrentNum++
				if (newCurrentNum == 0 || newCurrentNum == 100) {
					newCurrentNum = 0
					zeroMetCount++
				}
			}
		case 'L':
			for i := 0; i < count; i++ {
				newCurrentNum--
				if (newCurrentNum == 0 || newCurrentNum == -100) {
					newCurrentNum = 0
					zeroMetCount++
				}
			}
		default:
			return currentNum, 0, errors.New("Wrong direction.")
	}
	return newCurrentNum, zeroMetCount, nil
}

func processInput(input string) (int, error) {
	var err error
	var zeroMetCount int
	lines := strings.Split(input, "\n")
	linesLen := len(lines)
	answer := 0
	currentNum := 50

	for i := 0; i < linesLen; i++ {
		currentNum, zeroMetCount, err = processLine(lines[i], currentNum)
		if (err != nil) {
			return answer, err
		}
		answer += zeroMetCount
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
	fmt.Printf("The answer is : %d\n", answer)
}
