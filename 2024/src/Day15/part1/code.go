package main

import (
	"fmt"
	"path/filepath"

	"aoc/2024/pkg/file"
)

func getNextCoordinates(currentPos [2]int, move string) (int, int, bool) {
	if move == "<" {
		return currentPos[0], currentPos[1] - 1, true
	}

	if move == ">" {
		return currentPos[0], currentPos[1] + 1, true
	}

	if move == "^" {
		return currentPos[0] - 1, currentPos[1], true
	}

	if move == "v" {
		return currentPos[0] + 1, currentPos[1], true
	}

	return 0, 0, false
}

func canMoveBox(grid [][]string, move string, posY, posX int) (int, int, bool) {
	if move == ">" {
		for i := posX; i < len(grid[posY])-1; i++ {
			if grid[posY][i] == "." {
				return posY, i, true
			} else if grid[posY][i] == "#" {
				return 0, 0, false
			}
		}
	}

	if move == "<" {
		for i := posX; i > 0; i-- {
			if grid[posY][i] == "." {
				return posY, i, true
			} else if grid[posY][i] == "#" {
				return 0, 0, false
			}
		}
	}

	if move == "^" {
		for i := posY; i > 0; i-- {
			if grid[i][posX] == "." {
				return i, posX, true
			} else if grid[i][posX] == "#" {
				return 0, 0, false
			}
		}
	}

	if move == "v" {
		for i := posY; i < len(grid); i++ {
			if grid[i][posX] == "." {
				return i, posY, true
			} else if grid[i][posX] == "#" {
				return 0, 0, false
			}
		}
	}

	return 0, 0, false
}

func moveBox(grid [][]string, move string, posY, posX, emptyY, emptyX int) [][]string {
	grid[posY][posX] = "."

	if move == ">" {
		for i := posX + 1; i <= emptyX; i++ {
			grid[posY][i] = "O"
		}
	}

	if move == "<" {
		for i := posX - 1; i >= emptyX; i-- {
			grid[posY][i] = "O"
		}
	}

	if move == "^" {
		for i := posY - 1; i >= emptyY; i-- {
			grid[i][posX] = "O"
		}
	}

	if move == "v" {
		for i := posY + 1; i <= emptyY; i++ {
			grid[i][posX] = "O"
		}
	}

	return grid
}

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	grid := [][]string{}
	moves := []string{}
	robot := [2]int{}
	result := 0

	for y, row := range output {
		rowValues := []string{}
		for x, column := range row {
			if string(column) == "@" {
				robot = [2]int{y, x}
			}

			if string(column) == "<" || string(column) == ">" || string(column) == "^" || string(column) == "v" {
				moves = append(moves, string(column))
			} else if string(column) != "" {
				rowValues = append(rowValues, string(column))
			}
		}

		if len(rowValues) > 0 {
			grid = append(grid, rowValues)
		}
	}

	for _, move := range moves {
		nextY, nextX, _ := getNextCoordinates(robot, move)

		if grid[nextY][nextX] == "#" {
			continue
		}

		if grid[nextY][nextX] == "." {
			grid[robot[0]][robot[1]] = "."
			grid[nextY][nextX] = "@"
			robot = [2]int{nextY, nextX}

			continue
		}

		if grid[nextY][nextX] == "O" {
			emptyY, emptyX, canMove := canMoveBox(grid, move, nextY, nextX)

			if canMove {
				updatedGrid := moveBox(grid, move, nextY, nextX, emptyY, emptyX)
				grid = updatedGrid
				grid[robot[0]][robot[1]] = "."
				grid[nextY][nextX] = "@"
				robot = [2]int{nextY, nextX}
			}
		}
	}

	for y, row := range grid {
		for x, column := range row {
			if string(column) == "O" {
				result += 100*y + x
			}
		}
	}

	fmt.Println(result)
}
