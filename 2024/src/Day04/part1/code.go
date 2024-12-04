package main

import (
	"fmt"
	"path/filepath"

	"aoc/2024/pkg/file"
)

type Direction int

var (
	UP         = Direction(0)
	DOWN       = Direction(1)
	LEFT       = Direction(2)
	RIGHT      = Direction(3)
	LEFT_DOWN  = Direction(4)
	RIGHT_DOWN = Direction(5)
	LEFT_UP    = Direction(6)
	RIGHT_UP   = Direction(7)
)

func (d Direction) String() string {
	return [...]string{"UP", "DOWN", "LEFT", "RIGHT", "LEFT_DOWN", "RIGHT_DOWN", "LEFT_UP", "RIGHT_UP"}[d]
}

func getNextLetter(letter string) string {
	if letter == "X" {
		return "M"
	} else if letter == "M" {
		return "A"
	} else if letter == "A" {
		return "S"
	}

	return "END"
}

func checkDirection(letter string, input []string, yIndex int, xIndex int, direction *Direction) bool {

	if letter == "END" {
		return true
	}

	//UP
	if direction == &UP && yIndex > 0 && input[yIndex-1][xIndex] == letter[0] {
		return checkDirection(getNextLetter(letter), input, yIndex-1, xIndex, &UP)
	}

	//DOWN
	if direction == &DOWN && yIndex < len(input)-1 && input[yIndex+1][xIndex] == letter[0] {
		return checkDirection(getNextLetter(letter), input, yIndex+1, xIndex, &DOWN)
	}

	//LEFT
	if direction == &LEFT && xIndex > 0 && input[yIndex][xIndex-1] == letter[0] {
		return checkDirection(getNextLetter(letter), input, yIndex, xIndex-1, &LEFT)
	}

	//RIGHT
	if direction == &RIGHT && xIndex < len(input[yIndex])-1 && input[yIndex][xIndex+1] == letter[0] {
		return checkDirection(getNextLetter(letter), input, yIndex, xIndex+1, &RIGHT)
	}

	//LEFT_DOWN
	if direction == &LEFT_DOWN && xIndex > 0 && yIndex < len(input)-1 && input[yIndex+1][xIndex-1] == letter[0] {
		return checkDirection(getNextLetter(letter), input, yIndex+1, xIndex-1, &LEFT_DOWN)
	}

	//RIGHT_DOWN
	if direction == &RIGHT_DOWN && xIndex < len(input[yIndex])-1 && yIndex < len(input)-1 && input[yIndex+1][xIndex+1] == letter[0] {
		return checkDirection(getNextLetter(letter), input, yIndex+1, xIndex+1, &RIGHT_DOWN)
	}

	//LEFT_UP
	if direction == &LEFT_UP && xIndex > 0 && yIndex > 0 && input[yIndex-1][xIndex-1] == letter[0] {
		return checkDirection(getNextLetter(letter), input, yIndex-1, xIndex-1, &LEFT_UP)
	}

	//RIGHT_DOWN
	if direction == &RIGHT_UP && xIndex < len(input[yIndex])-1 && yIndex > 0 && input[yIndex-1][xIndex+1] == letter[0] {
		return checkDirection(getNextLetter(letter), input, yIndex-1, xIndex+1, &RIGHT_UP)
	}

	return false
}

func findPattern(input []string) int {
	result := 0

	for yIndex, line := range input {
		for xIndex, char := range line {
			if char == 'X' {

				up := checkDirection("M", input, yIndex, xIndex, &UP)
				down := checkDirection("M", input, yIndex, xIndex, &DOWN)
				left := checkDirection("M", input, yIndex, xIndex, &LEFT)
				right := checkDirection("M", input, yIndex, xIndex, &RIGHT)
				leftDown := checkDirection("M", input, yIndex, xIndex, &LEFT_DOWN)
				rightDown := checkDirection("M", input, yIndex, xIndex, &RIGHT_DOWN)
				leftUP := checkDirection("M", input, yIndex, xIndex, &LEFT_UP)
				rightUP := checkDirection("M", input, yIndex, xIndex, &RIGHT_UP)

				boolList := []bool{up, down, left, right, leftDown, rightDown, leftUP, rightUP}

				for _, value := range boolList {
					if value {
						result++
					}
				}
			}
		}
	}

	return result
}

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(findPattern(output))
}
