package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"aoc/2024/pkg/file"
)

func blink(stones map[int]int) map[int]int {
	newStones := make(map[int]int)

	for stone := range stones {
		strNumber := strconv.Itoa(stone)

		if stone == 0 {
			newStones[1] += stones[stone]
		} else if len(strNumber)%2 == 0 {
			half := len(strNumber) / 2
			firstHalf, _ := strconv.Atoi(strNumber[:half])
			secondHalf, _ := strconv.Atoi(strNumber[half:])

			newStones[firstHalf] += stones[stone]
			newStones[secondHalf] += stones[stone]
		} else {
			newNumber := stone * 2024
			newStones[newNumber] += stones[stone]
		}
	}

	return newStones
}

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	stones := make(map[int]int)
	numberOfStones := 0

	for _, line := range output {
		values := strings.Fields(line)

		for _, number := range values {
			num, _ := strconv.Atoi(number)

			stones[num] = stones[num] + 1

		}
	}

	for i := 0; i < 75; i++ {
		stones = blink(stones)
	}

	for _, stone := range stones {
		numberOfStones += stone
	}

	fmt.Println(numberOfStones)
}
