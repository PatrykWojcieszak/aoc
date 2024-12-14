package main

import (
	"fmt"
	"path/filepath"

	"aoc/2024/pkg/file"
)

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	height := 103
	width := 101
	seconds := 100
	halfX := (width - 1) / 2
	halfY := (height - 1) / 2
	q1, q2, q3, q4 := 0, 0, 0, 0

	for i := 0; i < len(output); i += 1 {
		var px, py, vx, vy int
		fmt.Sscanf(output[i], "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)

		newX := ((px+seconds*vx)%width + width) % width
		newY := ((py+seconds*vy)%height + height) % height

		if newX < halfX && newY < halfY {
			q1++
		}

		if newX > halfX && newY < halfY {
			q2++
		}

		if newX < halfX && newY > halfY {
			q3++
		}

		if newX > halfX && newY > halfY {
			q4++
		}
	}

	fmt.Println("Result:", q1*q2*q3*q4)
}
