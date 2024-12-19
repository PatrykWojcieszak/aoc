package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"aoc/2024/pkg/file"
)

func canConstructDesign(design string, towels []string, memo map[string]bool) bool {
	if result, exists := memo[design]; exists {
		return result
	}

	if design == "" {
		return true
	}

	for _, pattern := range towels {
		if strings.HasPrefix(design, pattern) {
			if canConstructDesign(design[len(pattern):], towels, memo) {
				memo[design] = true
				return true
			}
		}
	}

	memo[design] = false
	return false
}

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	towels := []string{}
	designs := []string{}
	memo := make(map[string]bool)
	possibleCount := 0

	avlTowels := strings.Split(output[0], ",")
	for _, towel := range avlTowels {
		towels = append(towels, strings.TrimSpace(towel))
	}

	for i := 2; i < len(output); i++ {
		designs = append(designs, output[i])
	}

	for _, design := range designs {
		if canConstructDesign(design, towels, memo) {
			possibleCount++
		}
	}

	fmt.Println(possibleCount)
}
