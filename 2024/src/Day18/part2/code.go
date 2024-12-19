package main

import (
	"container/heap"
	"fmt"
	"path/filepath"
	"strings"

	"aoc/2024/pkg/file"
)

type Position struct {
	x, y, orientation int
}

type PriorityQueueItem struct {
	position Position
	cost     int
	index    int
}

type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PriorityQueueItem)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

var directions = []struct{ dx, dy int }{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

func isValid(grid []string, x, y int) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid) && grid[y][x] != '#'
}

func shortestPath(grid []string) int {
	startX, startY, endX, endY := 0, 0, 70, 70

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &PriorityQueueItem{
		position: Position{startX, startY, 1},
		cost:     0,
	})

	visited := make(map[Position]int)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*PriorityQueueItem)
		pos := current.position
		cost := current.cost

		if pos.x == endX && pos.y == endY {
			return cost
		}

		if v, ok := visited[pos]; ok && v <= cost {
			continue
		}
		visited[pos] = cost

		for i, dir := range directions {
			newX, newY := pos.x+dir.dx, pos.y+dir.dy

			if isValid(grid, newX, newY) && pos.orientation == i {
				heap.Push(pq, &PriorityQueueItem{
					position: Position{newX, newY, i},
					cost:     cost + 1,
				})
			}

			if pos.orientation != i {
				heap.Push(pq, &PriorityQueueItem{
					position: Position{pos.x, pos.y, i},
					cost:     cost + 0,
				})
			}
		}
	}

	return -1
}

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	size := 71
	maxSize := 1024

	corrMem := [][2]int{}
	corrMemRest := [][2]int{}

	for index, row := range output {
		if index < maxSize {
			var x, y int
			fmt.Sscanf(row, "%d,%d", &x, &y)

			corrMem = append(corrMem, [2]int{x, y})
		} else {
			var x, y int
			fmt.Sscanf(row, "%d,%d", &x, &y)

			corrMemRest = append(corrMemRest, [2]int{x, y})
		}
	}

	grid := [][]string{}

	for y := 0; y < size; y++ {
		row := []string{}
		for x := 0; x < size; x++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}

	for _, byte := range corrMem {
		grid[byte[1]][byte[0]] = "#"
	}

	for _, rBytes := range corrMemRest {
		grid[rBytes[1]][rBytes[0]] = "#"

		pGrid := []string{}

		for _, row := range grid {
			pGrid = append(pGrid, strings.Join(row, ""))
		}

		result := shortestPath(pGrid)
		if result == -1 {
			fmt.Println("Coordinates:", rBytes[0], rBytes[1])
			break
		}
	}
}