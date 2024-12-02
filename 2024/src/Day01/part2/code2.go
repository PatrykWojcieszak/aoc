package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func filterArr(arr []int, valueToCheck int) []int {
	var result []int

	for _, value := range arr {
		if value == valueToCheck {
			result = append(result, value)
		}
	}

	return result
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var arr1, arr2 []int

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)

		intOneVal, err1 := strconv.Atoi(values[0])
		intTwoVal, err2 := strconv.Atoi(values[1])

		if err1 != nil || err2 != nil {
			fmt.Printf("Error parsing integers in line: %s\n", line)
			return
		}

		arr1 = append(arr1, intOneVal)
		arr2 = append(arr2, intTwoVal)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	var result int

	for _, value := range arr1 {
		filtered := filterArr(arr2, value)

		result = result + (value * len(filtered))
	}

	fmt.Println(result)
}
