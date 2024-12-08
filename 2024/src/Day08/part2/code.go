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

				antiNodeY1 := ant1[0]
				antiNodeX1 := ant1[1]

				for antiNodeY1 >= 0 && antiNodeY1 < len(grid) && antiNodeX1 >= 0 && antiNodeX1 < len(grid[0]) {
					antiNodes[[2]int{antiNodeY1, antiNodeX1}] = true

					antiNodeY1 += deltaY
					antiNodeX1 += deltaX
				}

				antiNodeY2 := ant2[0]
				antiNodeX2 := ant2[1]

				for antiNodeY2 >= 0 && antiNodeY2 < len(grid) && antiNodeX2 >= 0 && antiNodeX2 < len(grid[0]) {
					antiNodes[[2]int{antiNodeY2, antiNodeX2}] = true

					antiNodeY2 -= deltaY
					antiNodeX2 -= deltaX

				}
			}
		}
	}

	fmt.Println("Result -", len(antiNodes))
}
