package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"

	"aoc/2024/pkg/file"
)

func regexFilter(line []string) []string {
	pattern := `mul\(\d{1,3},\d{1,3}\)`

	var validMul []string

	for _, value := range line {
		regex := regexp.MustCompile(pattern)
		matches := regex.FindAllString(value, -1)

		validMul = append(validMul, matches...)
	}

	return validMul
}

func multiply(instr []string) int {
	pattern := `\d+`

	result := 0

	for _, value := range instr {
		regex := regexp.MustCompile(pattern)
		matches := regex.FindAllString(value, -1)

		num1, _ := strconv.Atoi((matches[0]))
		num2, _ := strconv.Atoi((matches[1]))

		result += num1 * num2
	}

	return result
}

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	validInstr := regexFilter(output)
	fmt.Println(multiply(validInstr))
}
