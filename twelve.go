package main

import (
	"container/list"
	"log"
	"os"
)

func createHeightMap(fileBytes []byte) [][]byte {
	res := make([][]byte, 0)
	row := make([]byte, 0)
	for _, b := range fileBytes {
		if b == '\n' {
			res = append(res, row)
			row = make([]byte, 0)
		} else {
			row = append(row, b)
		}
	}
	res = append(res, row)
	return res
}

func findStart(heightmap [][]byte, startByte byte) []int {
	for r := range heightmap {
		for c := range heightmap[r] {
			if heightmap[r][c] == startByte {
				return []int{r, c, 0}
			}
		}
	}
	log.Fatal("did not find start in heightmap")
	return make([]int, 0)
}

func getHeight(b byte) int {
	if b == 'S' {
		return int('a')
	} else if b == 'E' {
		return int('z')
	} else {
		return int(b)
	}
}

func new2D[T int | bool](n, m int) [][]T {
	grid := make([][]T, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]T, m)
	}
	return grid
}

type StepperFunc func(int, int) bool

func stepsToBestSignal(fileBytes []byte, startByte, stopByte byte, canStep StepperFunc) int {
	dirs := [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, 1}, []int{0, -1}}
	queue := list.New()
	heightmap := createHeightMap(fileBytes)
	numRows := len(heightmap)
	numCols := len(heightmap[0])
	visited := new2D[bool](numRows, numCols)
	start := findStart(heightmap, startByte)
	queue.PushBack(start)
	visited[start[0]][start[1]] = true
	for curr := queue.Front(); curr != nil; curr = queue.Front() {
		queue.Remove(curr)
		locAndSteps := curr.Value.([]int)
		row := locAndSteps[0]
		col := locAndSteps[1]
		steps := locAndSteps[2]
		height := getHeight(heightmap[row][col])
		if heightmap[row][col] == stopByte {
			return steps
		}
		for _, newDir := range dirs {
			newRow := row + newDir[0]
			newCol := col + newDir[1]
			isOutOfBounds := newRow < 0 || newRow >= numRows || newCol < 0 || newCol >= numCols
			if !isOutOfBounds && !visited[newRow][newCol] && canStep(getHeight(heightmap[newRow][newCol]), height) {
				visited[newRow][newCol] = true
				queue.PushBack([]int{newRow, newCol, steps + 1})
			}
		}
	}

	return -1
}

func main() {
	const filename = "data/twelve.txt"
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	ans := stepsToBestSignal(fileBytes, 'S', 'E', func(nextHeight, currHeight int) bool {
		return nextHeight-1 <= currHeight
	})
	log.Printf("steps to best signal: %d", ans)

	ans = stepsToBestSignal(fileBytes, 'E', 'a', func(nextHeight, currHeight int) bool {
		return nextHeight+1 >= currHeight
	})
	log.Printf("steps from lowest elevation: %d", ans)
}
