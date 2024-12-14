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

		minCost := -1
		found := false

		for aPress := 0; aPress < 100; aPress++ {
			for bPress := 0; bPress < 100; bPress++ {
				if ax*aPress+bx*bPress == px && ay*aPress+by*bPress == py {
					cost := aPress*3 + bPress*1
					if !found || cost < minCost {
						minCost = cost
						found = true
					}
				}
			}
		}

		if found {
			result += minCost
		}
	}

	fmt.Println("Result:", result)
}
