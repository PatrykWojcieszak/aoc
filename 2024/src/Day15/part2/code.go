package main

import (
	"fmt"
	"os"
	"path/filepath"

	"aoc/2024/pkg/file"
)

type Point struct {
	x int
	y int
}

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	file, err := os.OpenFile("result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	grid := [][]string{}
	moves := []string{}
	robot := Point{}
	result := 0

	directions := map[string]Point{
		"^": Point{0, -1},
		">": Point{1, 0},
		"v": Point{0, 1},
		"<": Point{-1, 0},
	}

	for y, row := range output {
		rowValues := []string{}
		for _, column := range row {

			if string(column) == "<" || string(column) == ">" || string(column) == "^" || string(column) == "v" {
				moves = append(moves, string(column))
			} else if string(column) != "" {
				if string(column) == "#" {
					rowValues = append(rowValues, "#")
					rowValues = append(rowValues, "#")
				} else if string(column) == "O" {
					rowValues = append(rowValues, "[")
					rowValues = append(rowValues, "]")
				} else if string(column) == "." {
					rowValues = append(rowValues, ".")
					rowValues = append(rowValues, ".")
				} else if string(column) == "@" {
					rowValues = append(rowValues, "@")
					robot = Point{len(rowValues) - 1, y}
					rowValues = append(rowValues, ".")
				}
			}
		}

		if len(rowValues) > 0 {
			grid = append(grid, rowValues)
		}
	}

	for _, move := range moves {
		dir := directions[move]
		nexPos := Point{robot.x + dir.x, robot.y + dir.y}

		if grid[nexPos.y][nexPos.x] == "#" {
			continue
		} else if grid[nexPos.y][nexPos.x] == "." {
			grid[robot.y][robot.x] = "."
			grid[nexPos.y][nexPos.x] = "@"
			robot = Point{nexPos.x, nexPos.y}
		} else if grid[nexPos.y][nexPos.x] == "[" || grid[nexPos.y][nexPos.x] == "]" {
			boxes := [][][2]Point{}

			if grid[nexPos.y][nexPos.x] == "[" {
				boxes = append(boxes, [][2]Point{{
					Point{nexPos.x, nexPos.y},
					Point{nexPos.x + 1, nexPos.y},
				}})
			} else {
				boxes = append(boxes, [][2]Point{{
					Point{nexPos.x - 1, nexPos.y},
					Point{nexPos.x, nexPos.y},
				}})
			}

			currentBoxes := boxes[0]

			wallFound := false
			for len(currentBoxes) > 0 && !wallFound {
				nextBoxes := [][2]Point{}
				for _, box := range currentBoxes {
					if grid[box[0].y+dir.y][box[0].x+dir.x] == "#" || grid[box[1].y+dir.y][box[1].x+dir.x] == "#" {
						wallFound = true
						break
					}

					if move == "^" || move == "v" {
						if grid[box[0].y+dir.y][box[0].x+dir.x] == "[" {
							nextBoxes = append(nextBoxes, [2]Point{
								Point{box[0].x, box[0].y + dir.y},
								Point{box[0].x + 1, box[0].y + dir.y},
							})
						} else if grid[box[0].y+dir.y][box[0].x+dir.x] == "]" {
							nextBoxes = append(nextBoxes, [2]Point{
								Point{box[0].x - 1, box[0].y + dir.y},
								Point{box[0].x, box[0].y + dir.y},
							})
						}
						if grid[box[1].y+dir.y][box[1].x+dir.x] == "[" {
							nextBoxes = append(nextBoxes, [2]Point{
								Point{box[1].x, box[1].y + dir.y},
								Point{box[1].x + 1, box[1].y + dir.y},
							})
						}

					} else if move == "<" || move == ">" {
						if move == "<" {
							if grid[box[0].y+dir.y][box[0].x+dir.x] == "]" {
								nextBoxes = append(nextBoxes, [2]Point{
									Point{box[0].x + dir.x*2, box[0].y + dir.y},
									Point{box[0].x + dir.x, box[0].y + dir.y},
								})
							}
						} else {
							if grid[box[1].y+dir.y][box[1].x+dir.x] == "[" {
								nextBoxes = append(nextBoxes, [2]Point{
									Point{box[1].x + dir.x, box[1].y + dir.y},
									Point{box[1].x + dir.x*2, box[1].y + dir.y},
								})
							}
						}
					}
				}

				if len(boxes) > 0 {
					boxes = append(boxes, nextBoxes)
				}
				currentBoxes = nextBoxes
			}

			if wallFound {
				continue
			}

			if len(boxes) > 0 {
				for i := len(boxes) - 1; i >= 0; i-- {
					for _, box := range boxes[i] {
						grid[box[0].y][box[0].x] = "."
						grid[box[1].y][box[1].x] = "."
						grid[box[0].y+dir.y][box[0].x+dir.x] = "["
						grid[box[1].y+dir.y][box[1].x+dir.x] = "]"
					}
				}
			}

			grid[robot.y][robot.x] = "."
			robot = Point{nexPos.x, nexPos.y}
			grid[robot.y][robot.x] = "@"
		}
	}

	for y := 1; y < len(grid); y++ {
		for x := 1; x < len(grid[y]); x++ {
			if grid[y][x] == "[" {
				result += 100*y + x
				x++
			}
		}
	}

	fmt.Println(result)
}
