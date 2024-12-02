package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func isIncreasingOrDecreasing(line string) bool {
	values := strings.Fields(line)
	var isDecreasing = false

	for index, value := range values {
		if index < len(values) && index > 0 {
			currentValue, err1 := strconv.Atoi(value)
			previousValue, err2 := strconv.Atoi(values[index-1])

			if err1 != nil || err2 != nil {
				fmt.Printf("Error parsing integers in line: %s\n", line)
				return false
			}

			if index == 1 {
				isDecreasing = previousValue > currentValue
			}

			if currentValue == previousValue {
				return false
			}

			if (!isDecreasing && currentValue < previousValue) || (isDecreasing && currentValue > previousValue) {
				return false
			}
		}
	}

	return true
}

func isDifferInRange(line string) bool {
	values := strings.Fields(line)
	var isInRange = true

	for index, value := range values {
		if index < len(values) && index > 0 {
			int1, err1 := strconv.Atoi(value)
			int2, err2 := strconv.Atoi(values[index-1])

			if err1 != nil || err2 != nil {
				fmt.Printf("Error parsing integers in line: %s\n", line)
				return false
			}

			var difference = absInt(int2 - int1)

			if difference < 1 || difference > 3 {
				isInRange = false
			}
		}
	}

	return isInRange
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result = 0

	for scanner.Scan() {
		line := scanner.Text()

		firstCondition := isIncreasingOrDecreasing((line))
		secondCondition := isDifferInRange(line)

		if firstCondition && secondCondition {
			result = result + 1
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	fmt.Println(result)
}
