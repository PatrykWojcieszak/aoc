package main

import (
	"fmt"
	"path/filepath"

	"aoc/2024/pkg/file"
)

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	grid := [103][101]int{}
	height := 103
	width := 101
	seconds := 100

	for i := 0; i < len(output); i += 1 {
		var px, py, vx, vy int
		fmt.Sscanf(output[i], "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)

		newX := ((px+seconds*vx)%width + width) % width
		newY := ((py+seconds*vy)%height + height) % height
		grid[newY][newX] += 1
	}

	q1, q2, q3, q4 := 0, 0, 0, 0
	halfX := (width - 1) / 2
	halfY := (height - 1) / 2

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x < halfX && y < halfY {
				q1 += grid[y][x]
			}

			if x > halfX && y < halfY {
				q2 += grid[y][x]
			}

			if x < halfX && y > halfY {
				q3 += grid[y][x]
			}

			if x > halfX && y > halfY {
				q4 += grid[y][x]
			}
		}

	}

	fmt.Println("Result:", q1*q2*q3*q4)
}
