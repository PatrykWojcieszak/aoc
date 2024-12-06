package main

import (
	"aoc/2024/pkg/file"
	"fmt"
	"path/filepath"
	"strconv"
)

func copyArray(original [][]string) [][]string {
	copy := [][]string{}

	for i := range original {
		row := []string{}
		row = append(row, original[i]...)

		copy = append(copy, row)
	}

	return copy
}

func getNextIndex(direction int, guardX int, guardY int) (y int, x int) {
	if direction == 0 {
		return guardY - 1, guardX
	} else if direction == 90 {
		return guardY, guardX + 1
	} else if direction == 180 {
		return guardY + 1, guardX
	} else if direction == 270 {
		return guardY, guardX - 1
	}

	return guardY, guardX
}

func removeDuplicates(input []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, value := range input {
		if !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}

	return result
}

func hasExit(labMap [][]string, guardX int, guardY int) bool {
	direction := 0
	currPosY := guardY
	currPosX := guardX

	visitedPos := make(map[string]int)

	for {
		y, x := getNextIndex(direction, currPosX, currPosY)

		if y < 0 || x < 0 || currPosX < 0 || currPosX >= len(labMap)-1 || currPosY < 0 || currPosY >= len(labMap[0])-1 {
			return true
		}

		isFreePath := labMap[y][x] != "#"

		visitedPosition := strconv.Itoa(y) + "|" + strconv.Itoa(x)
		value, exists := visitedPos[visitedPosition]

		if exists && value == direction {
			return false
		}

		if isFreePath {
			currPosX = x
			currPosY = y
			visitedPos[visitedPosition] = direction
		} else {
			if direction == 270 {
				direction = 0
			} else {
				direction += 90
			}
		}

	}
}

func hasFoundExit(labMap [][]string, guardX int, guardY int) []string {
	direction := 0
	currPosY := guardY
	currPosX := guardX

	visitedPos := make(map[string]int)
	obstacles := []string{}
	for {
		y, x := getNextIndex(direction, currPosX, currPosY)

		if y < 0 || x < 0 || currPosX < 0 || currPosX >= len(labMap)-1 || currPosY < 0 || currPosY >= len(labMap[0])-1 {
			return obstacles
		}

		isFreePath := labMap[y][x] != "#"

		newLabMap := copyArray(labMap)
		newLabMap[y][x] = "#"
		foundExitWithObstacle := hasExit(newLabMap, guardX, guardY)

		if !foundExitWithObstacle {
			newLabMap[y][x] = "O"
			obstacles = append(obstacles, strconv.Itoa(y)+"|"+strconv.Itoa(x))
		}

		visitedPosition := strconv.Itoa(y) + "|" + strconv.Itoa(x)
		value, exists := visitedPos[visitedPosition]

		if exists && value == direction {
			return obstacles
		}

		if isFreePath {
			currPosX = x
			currPosY = y
			visitedPos[visitedPosition] = direction
		} else {
			if direction == 270 {
				direction = 0
			} else {
				direction += 90
			}
		}

	}
}

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	guardX, guardY := 0, 0
	labMap := [][]string{}

	for y, row := range output {
		colValues := []string{}
		for x, column := range row {
			if string(column) == "^" {
				guardX = x
				guardY = y
				colValues = append(colValues, "X")
			} else {
				colValues = append(colValues, string(column))
			}
		}
		labMap = append(labMap, colValues)
	}

	fmt.Println("Obstacles", len(removeDuplicates(hasFoundExit(labMap, guardX, guardY))))
}
