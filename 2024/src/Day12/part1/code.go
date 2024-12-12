package main

import (
	"fmt"
	"path/filepath"
	"slices"

	"aoc/2024/pkg/file"
)

func groupAdjacent(coords [][2]int) [][][2]int {
	visited := make(map[[2]int]bool)
	var groups [][][2]int

	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	var dfs func(coord [2]int, group *[][2]int)
	dfs = func(coord [2]int, group *[][2]int) {
		visited[coord] = true
		*group = append(*group, coord)

		for _, dir := range directions {
			adjacent := [2]int{coord[0] + dir[0], coord[1] + dir[1]}

			contains := slices.IndexFunc(coords, func(c [2]int) bool {
				return c == adjacent
			})

			if !visited[adjacent] && contains != -1 {
				dfs(adjacent, group)
			}
		}
	}

	for _, coord := range coords {
		if !visited[coord] {
			var group [][2]int
			dfs(coord, &group)
			groups = append(groups, group)
		}
	}

	return groups
}

func getFencePrice(regions [][2]int, height int, width int) int {
	perimeter := 0
	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	for _, plant := range regions {
		for _, dir := range directions {
			isAdjacent := slices.IndexFunc(regions, func(c [2]int) bool {
				if plant != c && (plant[1] > 0 || plant[1] < width || plant[0] > 0 || plant[0] < height) {
					return c[0] == plant[0]+dir[0] && c[1] == plant[1]+dir[1]
				}

				return false
			})

			if isAdjacent == -1 {
				perimeter++
			}
		}
	}

	return len(regions) * perimeter
}

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	regions := make(map[string][][2]int)
	result := 0

	for y, row := range output {
		for x, column := range row {
			regions[string(column)] = append(regions[string(column)], [2]int{y, x})
		}
	}

	for _, region := range regions {
		groups := groupAdjacent(region)
		for _, group := range groups {
			result += getFencePrice(group, len(output), len(output[0]))
		}
	}

	fmt.Println(result)
}
