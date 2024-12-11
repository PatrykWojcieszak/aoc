package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"aoc/2024/pkg/file"
)

func arrangeStones(stones []int) []int {
	newStones := []int{}

	for _, number := range stones {
		strNumber := strconv.Itoa(number)

		if number == 0 {
			newStones = append(newStones, 1)
		} else if len(strNumber)%2 == 0 {
			half := len(strNumber) / 2
			firstHalf, _ := strconv.Atoi(strNumber[:half])
			secondHalf, _ := strconv.Atoi(strNumber[half:])

			newStones = append(newStones, firstHalf)
			newStones = append(newStones, secondHalf)
		} else {
			newStones = append(newStones, number*2024)
		}
	}

	return newStones
}

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	stones := []int{}

	for _, line := range output {
		values := strings.Fields(line)

		for _, number := range values {
			num, _ := strconv.Atoi(number)

			stones = append(stones, num)
		}
	}

	for i := 0; i < 25; i++ {
		stones = arrangeStones(stones)
	}

	fmt.Println(len(stones))
}
