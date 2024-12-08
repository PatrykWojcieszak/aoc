package main

import (
	"fmt"
	"path/filepath"

	"aoc/2024/pkg/file"
)

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)
	grid := [][]string{}
	antennas := make(map[string][][2]int)
	antiNodes := make(map[[2]int]bool)

	for _, row := range output {
		values := []string{}
		for _, column := range row {
			values = append(values, string(column))
		}

		grid = append(grid, values)
	}

	for y, row := range output {
		for x, column := range row {
			if string(column) != "." {
				antennas[string(column)] = append(antennas[string(column)], [2]int{y, x})
			}
		}
	}

	for _, antenna := range antennas {
		for i := 0; i < len(antenna); i++ {
			for j := i + 1; j < len(antenna); j++ {
				ant1 := antenna[i]
				ant2 := antenna[j]

				deltaY := ant1[0] - ant2[0]
				deltaX := ant1[1] - ant2[1]

				antiNode1 := [2]int{ant1[0] + deltaY, ant1[1] + deltaX}
				antiNode2 := [2]int{ant2[0] - deltaY, ant2[1] - deltaX}

				if antiNode1[0] >= 0 && antiNode1[0] < len(grid) && antiNode1[1] >= 0 && antiNode1[1] < len(grid[0]) {
					antiNodes[antiNode1] = true
				}

				if antiNode2[0] >= 0 && antiNode2[0] < len(grid) && antiNode2[1] >= 0 && antiNode2[1] < len(grid[0]) {
					antiNodes[antiNode2] = true
				}
			}
		}
	}

	fmt.Println("Result -", len(antiNodes))
}
