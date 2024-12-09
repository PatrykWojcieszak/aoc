package main

import (
	"fmt"
	"path/filepath"
	"strconv"

	"aoc/2024/pkg/file"
)

func getFirstEmptyBlock(blocks []string, blockLength int) (int, int, bool) {
	startIndex := 0
	endIndex := 0

	for i := 0; i < len(blocks); i++ {
		if blocks[i] == "." && startIndex == 0 {
			startIndex = i
		} else if blocks[i] != "." {
			if startIndex != 0 {
				endIndex = i - 1
			}
		}

		if startIndex != 0 && endIndex != 0 {
			emptyBlockLength := (endIndex - startIndex) + 1

			if emptyBlockLength >= blockLength {
				return startIndex, endIndex, true
			} else {
				startIndex = 0
				endIndex = 0

			}
		}
	}

	return 0, 0, false
}

func getLastFileBlock(blocks []string, lastIndex int) (int, int, bool) {
	startIndex := 0
	endIndex := 0
	blockId := ""

	for i := lastIndex; i > 0; i-- {
		if blocks[i] != "." && endIndex == 0 {
			endIndex = i
			blockId = blocks[i]
		} else if blocks[i] != blockId {
			if endIndex != 0 {
				startIndex = i + 1
			}
		}

		if startIndex != 0 && endIndex != 0 {
			return startIndex, endIndex, true
		}
	}

	return 0, 0, false
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

	lastFileIndex := len(blocks) - 1
	for {
		fileBlockStart, fileBlockEnd, isFileBlock := getLastFileBlock(blocks, lastFileIndex)
		fileBlockLength := (fileBlockEnd - fileBlockStart) + 1
		emptyBlockStart, emptyBlockEnd, isEmptyBlock := getFirstEmptyBlock(blocks, fileBlockLength)
		emptyBlockLength := (emptyBlockEnd - emptyBlockStart) + 1

		if !isFileBlock {
			break
		}

		if !isEmptyBlock || emptyBlockStart > fileBlockEnd {
			lastFileIndex = fileBlockStart - 1
			continue
		}

		if fileBlockStart > emptyBlockEnd {
			if emptyBlockLength >= fileBlockLength {
				for i := emptyBlockStart; i <= (emptyBlockStart+fileBlockLength)-1; i++ {
					blocks[i] = blocks[fileBlockStart]
				}

				for i := fileBlockStart; i <= fileBlockEnd; i++ {
					blocks[i] = "."
				}

				lastFileIndex = fileBlockStart - 1
			}
		}
	}

	for index, block := range blocks {
		if block != "." {
			blockNum, _ := strconv.Atoi(string(block))

			result += index * blockNum

		}
	}

	fmt.Println("Result -", result)
}
