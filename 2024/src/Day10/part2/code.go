package main

import (
	"fmt"
	"path/filepath"
	"strconv"

	"aoc/2024/pkg/file"
)

func findPath(grid [][]int, startY int, startX int, value int, paths int) int {
	numberOfPaths := paths

	//UP
	if startY < len(grid) && startY > 0 {
		nextValue := grid[startY-1][startX]

		if nextValue-value == 1 {
			if nextValue == 9 {
				numberOfPaths += 1
			}

			num := findPath(grid, startY-1, startX, nextValue, numberOfPaths)
			numberOfPaths = num
		}
	}

	//DOWN
	if startY >= 0 && startY < len(grid)-1 {
		nextValue := grid[startY+1][startX]

		if nextValue-value == 1 {
			if nextValue == 9 {
				numberOfPaths += 1
			}

			num := findPath(grid, startY+1, startX, nextValue, numberOfPaths)
			numberOfPaths = num
		}
	}

	//LEFT
	if startX < len(grid[0]) && startX > 0 {
		nextValue := grid[startY][startX-1]

		if nextValue-value == 1 {
			if nextValue == 9 {
				numberOfPaths += 1
			}

			num := findPath(grid, startY, startX-1, nextValue, numberOfPaths)
			numberOfPaths = num
		}
	}

	//RIGHT
	if startX >= 0 && startX < len(grid[0])-1 {
		nextValue := grid[startY][startX+1]

		if nextValue-value == 1 {
			if nextValue == 9 {
				numberOfPaths += 1
			}

			num := findPath(grid, startY, startX+1, nextValue, numberOfPaths)
			numberOfPaths = num
		}
	}

	return numberOfPaths
}

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	grid := [][]int{}
	trailHeads := make(map[string][2]int)
	result := 0

	for y, row := range output {
		colValues := []int{}
		for x, column := range row {
			value, _ := strconv.Atoi(string(column))
			colValues = append(colValues, value)

			if value == 0 {
				trailHeads[string(y)+string(x)] = [2]int{y, x}
			}
		}
		grid = append(grid, colValues)
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	for _, head := range trailHeads {
		paths := findPath(grid, head[0], head[1], 0, 0)
		result += paths
	}

	fmt.Println(result)
}
