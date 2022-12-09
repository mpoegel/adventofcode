package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func parseTreeGrid(fileBytes []byte) [][]int {
	grid := make([][]int, 0)
	row := make([]int, 0)
	for _, b := range fileBytes {
		if b == '\n' {
			grid = append(grid, row)
			row = make([]int, 0)
		} else {
			row = append(row, int(b)-int('0'))
		}
	}
	grid = append(grid, row)
	return grid
}

func initEmptyGrid[T int | bool](n, m int) [][]T {
	grid := make([][]T, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]T, m)
	}
	return grid
}

func printGrid[T int | bool](grid [][]T) {
	for _, row := range grid {
		for _, val := range row {
			fmt.Printf("%v ", val)
		}
		fmt.Printf("\n")
	}
}

func numberOfVisibleTrees(fileBytes []byte) int {
	total := 0
	grid := parseTreeGrid(fileBytes)
	numRows := len(grid)
	numCols := len(grid[0])
	tallestUp := initEmptyGrid[int](numRows, numCols)
	tallestLeft := initEmptyGrid[int](numRows, numCols)
	counted := initEmptyGrid[bool](numRows, numCols)

	for r, _ := range grid {
		for c, _ := range grid[r] {
			if r > 0 {
				if tallestUp[r-1][c] < grid[r][c] {
					total++
					tallestUp[r][c] = grid[r][c]
					counted[r][c] = true
				} else {
					tallestUp[r][c] = tallestUp[r-1][c]
				}
			} else {
				total++
				counted[r][c] = true
				tallestUp[r][c] = grid[r][c]
			}
			if c > 0 {
				if tallestLeft[r][c-1] < grid[r][c] {
					if !counted[r][c] {
						total++
					}
					tallestLeft[r][c] = grid[r][c]
					counted[r][c] = true
				} else {
					tallestLeft[r][c] = tallestLeft[r][c-1]
				}
			} else {
				if !counted[r][c] {
					total++
					counted[r][c] = true
				}
				tallestLeft[r][c] = grid[r][c]
			}
		}
	}

	tallestRight := initEmptyGrid[int](numRows, numCols)
	tallestDown := initEmptyGrid[int](numRows, numCols)
	for r := numRows - 1; r >= 0; r-- {
		for c := numCols - 1; c >= 0; c-- {
			if r < numRows-1 {
				if tallestDown[r+1][c] < grid[r][c] {
					if !counted[r][c] {
						total++
					}
					tallestDown[r][c] = grid[r][c]
					counted[r][c] = true
				} else {
					tallestDown[r][c] = tallestDown[r+1][c]
				}
			} else {
				if !counted[r][c] {
					total++
				}
				counted[r][c] = true
				tallestDown[r][c] = grid[r][c]
			}
			if c < numCols-1 {
				if tallestRight[r][c+1] < grid[r][c] {
					if !counted[r][c] {
						total++
					}
					tallestRight[r][c] = grid[r][c]
					counted[r][c] = true
				} else {
					tallestRight[r][c] = tallestRight[r][c+1]
				}
			} else {
				if !counted[r][c] {
					total++
					counted[r][c] = true
				}
				tallestRight[r][c] = grid[r][c]
			}
		}
	}

	return total
}

func minNonNegativeFromIndex(arr []int, start int) int {
	var runningMin int
	runningMin = math.MaxInt
	for _, val := range arr[start:] {
		if val > 0 && val < runningMin {
			runningMin = val
		}
	}
	if runningMin == math.MaxInt {
		return 0
	}
	return runningMin
}

func maxFromIndex(arr []int, start int) int {
	runningMax := 0
	for _, val := range arr[start:] {
		if val > runningMax {
			runningMax = val
		}
	}
	return runningMax
}

func highestScenicScore(fileBytes []byte) int {
	maxScore := 0
	grid := parseTreeGrid(fileBytes)
	numRows := len(grid)
	numCols := len(grid[0])
	scenicScore := initEmptyGrid[int](numRows, numCols)
	lastTreeOfSizeLeft := initEmptyGrid[int](numRows, 10)
	lastTreeOfSizeUp := initEmptyGrid[int](numCols, 10)
	lastTreeOfSizeRight := initEmptyGrid[int](numRows, 10)
	lastTreeOfSizeDown := initEmptyGrid[int](numCols, 10)

	for r := range grid {
		for c := range grid[r] {
			if r > 0 {
				scenicScore[r][c] = r - maxFromIndex(lastTreeOfSizeUp[c], grid[r][c])
			}
			lastTreeOfSizeUp[c][grid[r][c]] = r
			if c > 0 {
				scenicScore[r][c] *= c - maxFromIndex(lastTreeOfSizeLeft[r], grid[r][c])
			} else {
				scenicScore[r][c] = 0
			}
			lastTreeOfSizeLeft[r][grid[r][c]] = c
		}
	}

	for r := numRows - 1; r >= 0; r-- {
		for c := numCols - 1; c >= 0; c-- {
			if r < numRows-1 {
				viewDown := minNonNegativeFromIndex(lastTreeOfSizeDown[c], grid[r][c])
				if viewDown > 0 {
					scenicScore[r][c] *= viewDown - r
				} else {
					scenicScore[r][c] *= numRows - r - 1
				}
			} else {
				scenicScore[r][c] = 0
			}
			lastTreeOfSizeDown[c][grid[r][c]] = r
			if c < numCols-1 {
				viewRight := minNonNegativeFromIndex(lastTreeOfSizeRight[r], grid[r][c])
				if viewRight > 0 {
					scenicScore[r][c] *= viewRight - c
				} else {
					scenicScore[r][c] *= numCols - c - 1
				}
				if scenicScore[r][c] > maxScore {
					maxScore = scenicScore[r][c]
				}
			} else {
				scenicScore[r][c] = 0
			}
			lastTreeOfSizeRight[r][grid[r][c]] = c
		}
	}

	return maxScore
}

func main() {
	const filename = "data/eight.txt"
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	ans := numberOfVisibleTrees(fileBytes)
	log.Printf("number of trees visible from the outside: %d", ans)

	ans = highestScenicScore(fileBytes)
	log.Printf("highest scenic score: %d", ans)
}
