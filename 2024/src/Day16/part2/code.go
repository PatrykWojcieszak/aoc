package main

import (
	"container/heap"
	"fmt"
	"path/filepath"

	"aoc/2024/pkg/file"
)

type Position struct {
	x, y, orientation int
}

type PriorityQueueItem struct {
	position Position
	cost     int
	index    int
	path     []Position
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

func shortestPath(grid []string) (int, [][]Position) {
	var startX, startY, endX, endY int

	for y, row := range grid {
		for x, cell := range row {
			if cell == 'S' {
				startX, startY = x, y
			} else if cell == 'E' {
				endX, endY = x, y
			}
		}
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &PriorityQueueItem{
		position: Position{startX, startY, 1},
		cost:     0,
		path:     []Position{{startX, startY, 1}},
	})

	visited := make(map[Position]int)
	minCost := -1
	var allPaths [][]Position

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*PriorityQueueItem)
		pos := current.position
		cost := current.cost
		path := current.path

		if pos.x == endX && pos.y == endY {
			if minCost == -1 || cost < minCost {
				minCost = cost
				allPaths = [][]Position{path}
			} else if cost == minCost {
				allPaths = append(allPaths, path)
			}
			continue
		}

		if v, ok := visited[pos]; ok && v < cost {
			continue
		}
		visited[pos] = cost

		for i, dir := range directions {
			newX, newY := pos.x+dir.dx, pos.y+dir.dy

			if isValid(grid, newX, newY) && pos.orientation == i {
				heap.Push(pq, &PriorityQueueItem{
					position: Position{newX, newY, i},
					cost:     cost + 1,
					path:     append(append([]Position{}, path...), Position{newX, newY, i}),
				})
			}

			if pos.orientation != i {
				heap.Push(pq, &PriorityQueueItem{
					position: Position{pos.x, pos.y, i},
					cost:     cost + 1000,
					path:     append(append([]Position{}, path...), Position{pos.x, pos.y, i}),
				})
			}
		}
	}

	return minCost, allPaths
}

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	grid := []string{}
	grid = append(grid, output...)
	result := 0

	_, seats := shortestPath(grid)

	newGrid := [][]string{}

	for _, row := range output {
		rowValues := []string{}
		for _, column := range row {
			rowValues = append(rowValues, string(column))
		}

		newGrid = append(newGrid, rowValues)
	}

	for _, seat := range seats {
		for _, s := range seat {
			newGrid[s.y][s.x] = "O"
		}
	}

	for _, row := range newGrid {
		fmt.Println(row)
	}

	for _, row := range newGrid {
		for _, column := range row {
			if string(column) == "O" {
				result++
			}
		}
	}

	fmt.Println("Result:", result)
}
