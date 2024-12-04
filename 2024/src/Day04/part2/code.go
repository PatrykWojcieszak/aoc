package main

import (
	"fmt"
	"path/filepath"

	"aoc/2024/pkg/file"
)

func findPattern(input []string) int {
	result := 0

	for yIndex := 1; yIndex < len(input)-1; yIndex++ {
		for xIndex := 1; xIndex < len(input[yIndex])-1; xIndex++ {
			if input[yIndex][xIndex] == "A"[0] {
				leftDiagonal := (input[yIndex-1][xIndex-1] == "M"[0] && input[yIndex+1][xIndex+1] == "S"[0]) || (input[yIndex-1][xIndex-1] == "S"[0] && input[yIndex+1][xIndex+1] == "M"[0])
				rightDiagonal := (input[yIndex+1][xIndex-1] == "M"[0] && input[yIndex-1][xIndex+1] == "S"[0]) || (input[yIndex+1][xIndex-1] == "S"[0] && input[yIndex-1][xIndex+1] == "M"[0])

				if leftDiagonal && rightDiagonal {
					result++
				}
			}
		}
	}

	return result
}

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(findPattern(output))
}
