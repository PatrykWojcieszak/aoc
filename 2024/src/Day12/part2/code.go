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

	for _, plant := range regions {
		left := slices.IndexFunc(regions, func(c [2]int) bool {
			if plant != c && plant[1] > 0 {
				return c[0] == plant[0] && c[1] == plant[1]-1
			}

			return false
		})

		right := slices.IndexFunc(regions, func(c [2]int) bool {
			if plant != c && plant[1] < width {
				return c[0] == plant[0] && c[1] == plant[1]+1
			}

			return false
		})

		top := slices.IndexFunc(regions, func(c [2]int) bool {
			if plant != c && plant[0] > 0 {
				return c[0] == plant[0]-1 && c[1] == plant[1]
			}

			return false
		})

		bottom := slices.IndexFunc(regions, func(c [2]int) bool {
			if plant != c && plant[0] < height {
				return c[0] == plant[0]+1 && c[1] == plant[1]
			}

			return false
		})

		innerBottomRight := slices.IndexFunc(regions, func(c [2]int) bool {
			if plant != c && (plant[1] > 0 || plant[1] < width || plant[0] > 0 || plant[0] < height) {
				return c[0] == plant[0]+1 && c[1] == plant[1]+1
			}

			return false
		})

		innerBottomLeft := slices.IndexFunc(regions, func(c [2]int) bool {
			if plant != c && (plant[1] > 0 || plant[1] < width || plant[0] > 0 || plant[0] < height) {
				return c[0] == plant[0]+1 && c[1] == plant[1]-1
			}

			return false
		})

		innerTopLeft := slices.IndexFunc(regions, func(c [2]int) bool {
			if plant != c && (plant[1] > 0 || plant[1] < width || plant[0] > 0 || plant[0] < height) {
				return c[0] == plant[0]-1 && c[1] == plant[1]-1
			}

			return false
		})

		innerTopRight := slices.IndexFunc(regions, func(c [2]int) bool {
			if plant != c && (plant[1] > 0 || plant[1] < width || plant[0] > 0 || plant[0] < height) {
				return c[0] == plant[0]-1 && c[1] == plant[1]+1
			}

			return false
		})

		//OUTER TOP-LEFT CORNER
		if top == -1 && left == -1 {
			perimeter++
		}

		//OUTER BOTTOM-LEFT CORNER
		if bottom == -1 && left == -1 {
			perimeter++
		}

		//OUTER TOP-RIGHT CORNER
		if top == -1 && right == -1 {
			perimeter++
		}

		//OUTER BOTTOM-RIGHT CORNER
		if bottom == -1 && right == -1 {
			perimeter++
		}

		//INNER BOTTOM-LEFT CORNER
		if bottom == -1 && innerBottomLeft != -1 && left != -1 {
			perimeter++
		}

		//INNER BOTTOM-RIGHT CORNER
		if bottom == -1 && innerBottomRight != -1 && right != -1 {
			perimeter++
		}

		//INNER TOP-LEFT CORNER
		if top == -1 && innerTopLeft != -1 && left != -1 {
			perimeter++
		}
		//INNER TOP-RIGHT CORNER
		if top == -1 && innerTopRight != -1 && right != -1 {
			perimeter++
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

	for letter, region := range regions {
		groups := groupAdjacent(region)
		for _, group := range groups {
			sides := getFencePrice(group, len(output), len(output[0]))
			fmt.Println(letter, ":", sides)
			result += sides
		}
	}

	fmt.Println("result:", result)
}
