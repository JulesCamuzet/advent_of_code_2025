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

func isNumberValid(number int) bool {
	strNum := strconv.Itoa(number)
	if (strNum[0] == '0') {
		return false
	}
	numLen := len(strNum)
	if (numLen % 2 == 1) {
		return true
	}
	middle := numLen / 2
	for i := 0; i < middle; i++ {
		if (strNum[i] != strNum[i + middle]) {
			return true
		}
	}
	return false
}

func processRange(currRange string) (int, error) {
	bounds := strings.Split(currRange, "-")
	if (len(bounds) != 2) {
		return 0, errors.New("Error parsing range.")
	}
	min, err := strconv.Atoi(bounds[0])
	if (err != nil) {
		return 0, err
	}
	max, err := strconv.Atoi(bounds[1])
	if (err != nil) {
		return 0, err
	}
	if (max < min) {
		return 0, errors.New("Min and max in wrong order.")
	}
	result := 0
	for i := min; i <= max; i++ {
		if (isNumberValid(i) == false) {
			result += i
		}
	}
	return result, nil
}

func processInput(input string) (int, error) {
	var result int
	var ranges []string
	var rangesLen int

	result = 0
	ranges = strings.Split(input, ",")
	rangesLen = len(ranges)
	for i := 0; i < rangesLen; i++ {
		rangeResult, err := processRange(ranges[i])
		if (err != nil) {
			return 0, err
		}
		result += rangeResult
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
