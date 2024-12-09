package main

import (
	"fmt"
	"path/filepath"
	"strconv"

	"aoc/2024/pkg/file"
)

func getLastBlock(blocks []string) (int, bool) {
	for i := len(blocks) - 1; i > 0; i-- {
		if blocks[i] != "." {
			return i, true
		}
	}

	return 0, false
}

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	blocks := []string{}
	fileId := 0

	result := 0

	for _, line := range output {
		for index, block := range line {
			isFile := index%2 == 0

			blockNum, _ := strconv.Atoi(string(block))

			for i := 0; i < blockNum; i++ {
				if isFile {
					blocks = append(blocks, strconv.Itoa(fileId))
				} else {
					blocks = append(blocks, ".")
				}
			}

			if isFile {
				fileId++
			}
		}
	}

	for i := 0; i < len(blocks); i++ {
		if blocks[i] == "." {
			index, isFile := getLastBlock(blocks)

			if isFile && i < index {
				blocks[i] = blocks[index]
				blocks[index] = "."
			}
		}
	}

	for index, block := range blocks {
		if block != "." {
			blockNum, _ := strconv.Atoi(string(block))

			result += index * blockNum

		} else {
			break
		}
	}

	fmt.Println("Result -", result)
}
