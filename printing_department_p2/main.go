package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Position struct {
	x int
	y int
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

func cloneMap(src []string) [][]byte {
	srcLines := len(src)
	srcCol := len(src[0])
	dst := make([][]byte, srcLines)
	for y := 0; y < srcLines; y++ {
		dst[y] = make([]byte, srcCol)
		for x := 0; x < srcCol; x++ {
			dst[y][x] = src[y][x]
		}
	}
	return dst
}

func countCloseRolls(lines [][]byte, x, y, linesCount, columnsCount int) int {
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

func getPositionsToDelete(bMap [][]byte) []Position {
	linesCount := len(bMap)
	columnsCount := len(bMap[0])
	positionsToDelete := []Position{}
	y := 0
	for y < linesCount {
		x := 0
		for x < columnsCount {
			if (bMap[y][x] == '@' && countCloseRolls(bMap, x, y, linesCount, columnsCount) < 4) {
				positionsToDelete = append(positionsToDelete, Position{ x, y })
			}
			x++
		}
		y++
	}
	return positionsToDelete
}

func processInput(input string) int {
	lines := strings.Split(input, "\n")
	bMap := cloneMap(lines)
	result := 0
	run := true
	for run == true {
		positionsToDelete := getPositionsToDelete(bMap)
		positionsToDeleteLen := len(positionsToDelete)
		if (positionsToDeleteLen == 0) {
			run = false
		} else {
			result += positionsToDeleteLen
			for i := 0; i < positionsToDeleteLen; i++ {
				bMap[positionsToDelete[i].y][positionsToDelete[i].x] = '.'
			}
		}
	}
	return result
}

func main() {
	start := time.Now().UnixMilli()
	input, err := getFileContent()
	if (err != nil) {
		panic(err)
	}
	answer := processInput(input)
	end := time.Now().UnixMilli()
	duration := end - start
	fmt.Printf("The answer is : %d. Found in %d ms", answer, duration)
}
