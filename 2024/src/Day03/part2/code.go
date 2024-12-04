package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"aoc/2024/pkg/file"
)

func regexFilter(line []string) []string {
	pattern := `(?:mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\))`

	var validMul []string

	for _, value := range line {
		regex := regexp.MustCompile(pattern)
		matches := regex.FindAllString(value, -1)

		validMul = append(validMul, matches...)
	}

	return validMul
}

func multiply(instr []string) int {
	numbersPattern := `\d+`
	dontPattern := `don't`
	doPattern := `do`
	shouldMultiply := true
	result := 0

	for _, value := range instr {
		regex := regexp.MustCompile(numbersPattern)
		matches := regex.FindAllString(value, -1)

		if len(matches) == 0 && strings.Contains(value, dontPattern) {
			shouldMultiply = false
		} else if len(matches) == 0 && strings.Contains(value, doPattern) {
			shouldMultiply = true
		}

		if shouldMultiply && len(matches) > 1 {
			num1, _ := strconv.Atoi((matches[0]))
			num2, _ := strconv.Atoi((matches[1]))

			result += num1 * num2
		}
	}

	return result
}

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	validInstr := regexFilter(output)
	fmt.Println(multiply(validInstr))
}
