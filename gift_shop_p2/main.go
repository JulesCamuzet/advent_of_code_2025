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

func splitByN(str string, n int) []string {
	strLen := len(str)
	segCount := strLen / n
	var result []string
	for i := 0; i < segCount; i++ {
		var currSeg []byte
		for j := 0; j < n; j++ {
			currSeg = append(currSeg, str[i * n + j])
		}
		result = append(result, string(currSeg))
	}
	return result
}

func isNumberValid(number int) bool {
	strNum := strconv.Itoa(number)
	numLen := len(strNum)
	if (numLen < 2) {
		return true
	}
	for segLen := 1; segLen <= numLen / 2; segLen++ {
		if (numLen % segLen == 0 && numLen / segLen > 1) {
			segments := splitByN(strNum, segLen)
			for j := 0; j < numLen / segLen; j++ {
				if (segments[0] != segments[j]) {
					break
				} else if (j == numLen / segLen - 1) {
					return false
				}
			}
		}
	}
	return true
}

func processRange(currRange string) (uint64, error) {
	bounds := strings.Split(currRange, "-")
	var result uint64 = 0
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
	for i := min; i <= max; i++ {
		if (isNumberValid(i) == false) {
			fmt.Printf("invalid : %d\n", i)
			result += (uint64)(i)
		}
	}
	return result, nil
}

func processInput(input string) (uint64, error) {
	var result uint64
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
