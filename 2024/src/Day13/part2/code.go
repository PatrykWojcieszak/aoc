package main

import (
	"fmt"
	"path/filepath"

	"aoc/2024/pkg/file"
)

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	result := 0

	for i := 0; i < len(output); i += 4 {
		var ax, ay, bx, by, px, py int
		fmt.Sscanf(output[i], "Button A: X+%d, Y+%d", &ax, &ay)
		fmt.Sscanf(output[i+1], "Button B: X+%d, Y+%d", &bx, &by)
		fmt.Sscanf(output[i+2], "Prize: X=%d, Y=%d", &px, &py)

		px += 10000000000000
		py += 10000000000000

		det := ax*by - ay*bx

		if det == 0 {
			continue
		}

		aPress := (px*by - py*bx) / det
		bPress := (ax*py - ay*px) / det

		if ax*aPress+bx*bPress == px && ay*aPress+by*bPress == py && aPress >= 0 && bPress >= 0 {
			cost := aPress*3 + bPress
			result += cost
		}

	}

	fmt.Println("Result:", result)
}
