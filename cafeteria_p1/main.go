package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Range struct {
	min int
	max int 
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

func getRangesFromStr(strRanges []string) ([]Range, error) {
	ranges := []Range{}
	strRangesLen := len(strRanges)
	for i := 0; i < strRangesLen; i++ {
		splitRange := strings.Split(strRanges[i], "-")
		if (len(splitRange) != 2) {
			return ranges, errors.New("Error parsing range.")
		}
		min, err := strconv.Atoi(splitRange[0])
		if (err != nil) {
			return ranges, err
		}
		max, err := strconv.Atoi(splitRange[1])
		if (err != nil) {
			return ranges, err
		}
		ranges = append(ranges, Range{ min, max })
	}
	return ranges, nil
}

func getNumbersFromStr(strNumbers []string) ([]int, error) {
	strNumbersLen := len(strNumbers)
	numbers := []int{}
	for i := 0; i < strNumbersLen; i++ {
		num, err := strconv.Atoi(strNumbers[i])
		if (err != nil) {
			return numbers, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func checkIfNumberFresh(number int, ranges []Range) bool {
	rangesLen := len(ranges)
	for i := 0; i < rangesLen; i++ {
		if (number >= ranges[i].min && number <= ranges[i].max) {
			return true
		}
	}
	return false
}

func processInput(input string) (int, error) {
	splitInput := strings.Split(input, "\n\n")
	if (len(splitInput) != 2) {
		return 0, errors.New("Error parsing.")
	}
	strRanges := strings.Split(splitInput[0], "\n")
	strNumbers := strings.Split(splitInput[1], "\n")
	ranges, err := getRangesFromStr(strRanges)
	if (err != nil) {
		return 0, err
	}
	numbers, err := getNumbersFromStr(strNumbers)
	if (err != nil) {
		return 0, err
	}
	numbersLen := len(numbers)
	freshNumberCount := 0
	for i := 0; i < numbersLen; i++ {
		if (checkIfNumberFresh(numbers[i], ranges)) {
			freshNumberCount++
		}
	}
	return freshNumberCount, nil
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