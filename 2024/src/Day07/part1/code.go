package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"aoc/2024/pkg/file"
)

func verifyNumber(expected int, values []int, index int, currentResult int) bool {

	if index >= len(values) {
		return expected == currentResult
	}

	newResult := currentResult + values[index]
	if newResult <= expected {
		if verifyNumber(expected, values, index+1, newResult) {
			return true
		}
	}

	newResult = currentResult * values[index]
	if newResult <= expected {
		return verifyNumber(expected, values, index+1, newResult)
	}

	return false
}

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)
	result := 0
	input := make(map[int][]int)

	for _, row := range output {
		valuesNum := []int{}
		splitted := strings.Split(row, ":")

		expectedNum, _ := strconv.Atoi(splitted[0])

		splittedValues := strings.Fields(splitted[1])

		for _, value := range splittedValues {
			if value != "" {
				valueNum, _ := strconv.Atoi(value)
				valuesNum = append(valuesNum, valueNum)
			}
		}
		input[expectedNum] = valuesNum

	}

	for expected, values := range input {
		if verifyNumber(expected, values, 0, 0) {
			result += expected
		}
	}

	fmt.Println("Result -", result)
}
