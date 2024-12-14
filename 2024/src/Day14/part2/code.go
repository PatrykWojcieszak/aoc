package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"aoc/2024/pkg/file"
)

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	file, err := os.OpenFile("result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	grid := [103][101]int{}
	robots := make(map[int][2][2]int)
	height := 103
	width := 101

	for i := 0; i < len(output); i += 1 {
		var px, py, vx, vy int
		fmt.Sscanf(output[i], "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)

		position := [2]int{px, py}
		velocity := [2]int{vx, vy}
		robots[i] = [2][2]int{position, velocity}
		grid[py][px] += 1
	}

	for i := 1; i <= 10000; i++ {
		for robotIndex, robot := range robots {
			pos := robot[0]
			vel := robot[1]

			nextPosX := pos[0] + vel[0]
			nextPosY := pos[1] + vel[1]

			if nextPosX < 0 {
				nextPosX += width
			}

			if nextPosX >= width {
				nextPosX -= width
			}

			if nextPosY < 0 {
				nextPosY += height
			}

			if nextPosY >= height {
				nextPosY -= height
			}

			grid[nextPosY][nextPosX] += 1
			grid[pos[1]][pos[0]] -= 1
			robots[robotIndex] = [2][2]int{{nextPosX, nextPosY}, vel}
		}

		for _, row := range grid {
			strs := make([]string, len(row))
			for i, num := range row {
				if num == 0 {
					strs[i] = " "
				} else {
					strs[i] = "."
				}
			}

			content := strings.Join(strs, "") + "\n"
			_, err = file.WriteString(content)
			if err != nil {
				fmt.Println("Error appending to file:", err)
				return
			}
		}
		content := "Seconds " + strconv.Itoa(i) + "\n" + "\n"
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println("Error appending to file:", err)
			return
		}
	}

}
