package main

import (
	"fmt"
	"path/filepath"
	"strconv"

	"aoc/2024/pkg/file"
)

func findPath(grid [][]int, startY int, startX int, value int, visited map[[2]int]bool) map[[2]int]bool {
	visitedPaths := visited

	//UP
	if startY < len(grid) && startY > 0 {
		nextValue := grid[startY-1][startX]

		if nextValue-value == 1 {
			if nextValue == 9 {
				visitedPaths[[2]int{startY - 1, startX}] = true
			}

			num := findPath(grid, startY-1, startX, nextValue, visitedPaths)
			visitedPaths = num
		}
	}

	//DOWN
	if startY >= 0 && startY < len(grid)-1 {
		nextValue := grid[startY+1][startX]

		if nextValue-value == 1 {
			if nextValue == 9 {
				visitedPaths[[2]int{startY + 1, startX}] = true
			}

			num := findPath(grid, startY+1, startX, nextValue, visitedPaths)
			visitedPaths = num
		}
	}

	//LEFT
	if startX < len(grid[0]) && startX > 0 {
		nextValue := grid[startY][startX-1]

		if nextValue-value == 1 {
			if nextValue == 9 {
				visitedPaths[[2]int{startY, startX - 1}] = true
			}

			num := findPath(grid, startY, startX-1, nextValue, visitedPaths)
			visitedPaths = num
		}
	}

	//RIGHT
	if startX >= 0 && startX < len(grid[0])-1 {
		nextValue := grid[startY][startX+1]

		if nextValue-value == 1 {
			if nextValue == 9 {
				visitedPaths[[2]int{startY, startX + 1}] = true
			}

			num := findPath(grid, startY, startX+1, nextValue, visitedPaths)
			visitedPaths = num
		}
	}

	return visitedPaths
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

	for _, head := range trailHeads {
		visited := map[[2]int]bool{}
		paths := findPath(grid, head[0], head[1], 0, visited)
		result += len(paths)
	}

	fmt.Println(result)
}
