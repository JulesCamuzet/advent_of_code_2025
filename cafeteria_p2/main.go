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

//func (r Range) display() {
//	fmt.Printf("Min : %d\nMax : %d\n", r.min, r.max)
//}

//func displayRangeArray(arr []Range) {
//	len := len(arr)
//	for i := 0; i < len; i++ {
//		arr[i].display()
//		fmt.Printf("\n")
//	}
//}

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

func checkIfOverlap(range1 Range, range2 Range) bool {
	if (range1.max == -1 || range2.max == -1) {
		return false
	} else if (range2.max < range1.min || range1.max < range2.min) {
		return false
	} else {
		return true
	}
}

func checkIfOverlapArray(ranges []Range) bool {
	len := len(ranges)
	for i := 0; i < len; i++ {
		for j := 0; j < i; j++ {
			if (checkIfOverlap(ranges[i], ranges[j])) {
				return true
			}
		}
	}
	return false
}

// Sry for garbage function lol
func mergeRanges(ranges []Range) {
	len := len(ranges)
	for checkIfOverlapArray(ranges) {
		for i := 0; i < len; i++ {
			for j := 0; j < i; j++ {
				range1 := &ranges[i]
				range2 := &ranges[j]
				if (!checkIfOverlap(*range1, *range2)) {
					continue
				}
				if (range1.min <= range2.min && range1.max >= range2.min && range1.max <= range2.max) {
					range1.max = range2.max
					range2.max = -1
					range2.min = -1
				} else if (range1.min <= range2.max && range1.min >= range2.min && range1.max >= range2.max) {
					range1.min = range2.min
					range2.max = -1
					range2.min = -1
				} else if (range1.min <= range2.min && range1.max >= range2.max) {
					range2.max = -1
					range2.min = -1
				} else if (range1.min >= range2.min && range1.max <= range2.max) {
					range1.max = -1
					range1.min = -1
				}
			}
		}
	}
}

func filterDeadRanges(ranges []Range) []Range {
	len := len(ranges)
	count := 0
	for i := 0; i < len; i++ {
		if (ranges[i].max != -1 && ranges[i].min != -1) {
			count++
		}
	}
	newRanges := make([]Range, count)
	index := 0
	for i := 0; i < len; i++ {
		if (ranges[i].max != -1 && ranges[i].min != -1) {
			newRanges[index].min = ranges[i].min
			newRanges[index].max = ranges[i].max
			index++
		}
	}
	return newRanges
}

func processInput(input string) (int, error) {
	splitInput := strings.Split(input, "\n\n")
	if (len(splitInput) != 2) {
		return 0, errors.New("Error parsing.")
	}
	strRanges := strings.Split(splitInput[0], "\n")
	ranges, err := getRangesFromStr(strRanges)
	if (err != nil) {
		return 0, err
	}
	mergeRanges(ranges)
	filteredRanges := filterDeadRanges(ranges)
	len := len(filteredRanges)
	freshNumberCount := 0
	for i := 0; i < len; i++ {
		freshNumberCount += filteredRanges[i].max - filteredRanges[i].min + 1
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