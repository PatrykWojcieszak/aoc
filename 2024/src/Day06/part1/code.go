package main

import (
	"fmt"
	"path/filepath"

	"aoc/2024/pkg/file"
)

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

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	guardX, guardY := 0, 0
	labMap := [][]string{}
	direction := 0
	result := 0

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

	for {
		if guardX < 0 || guardX >= len(labMap)-1 || guardY < 0 || guardY >= len(labMap[0])-1 {
			break
		}

		y, x := getNextIndex(direction, guardX, guardY)
		isFreePath := labMap[y][x] != "#"

		if isFreePath {
			guardX = x
			guardY = y
			labMap[y][x] = "X"
		} else {
			if direction == 270 {
				direction = 0
			} else {
				direction += 90
			}
		}
	}

	for _, row := range labMap {
		for _, column := range row {
			if string(column) == "X" {
				result++
			}
		}
	}

	fmt.Println("result", result)
}
