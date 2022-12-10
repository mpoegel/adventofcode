package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func posToStr(pos []int) string {
	return fmt.Sprintf("%d,%d", pos[0], pos[1])
}

func updateTail(head, tail []int) []int {
	distX := head[0] - tail[0]
	distY := head[1] - tail[1]
	if distX > 1 || (math.Abs(float64(distY)) > 1 && distX > 0) {
		tail[0] += 1
	} else if distX < -1 || (math.Abs(float64(distY)) > 1 && distX < 0) {
		tail[0] -= 1
	}
	if distY > 1 || (math.Abs(float64(distX)) > 1 && distY > 0) {
		tail[1] += 1
	} else if distY < -1 || (math.Abs(float64(distX)) > 1 && distY < 0) {
		tail[1] -= 1
	}
	return tail
}

func numberOfTailPositions(fileBytes []byte, ropeLen int) int {
	total := 0
	rope := make([][]int, ropeLen)
	for i := 0; i < ropeLen; i++ {
		rope[i] = make([]int, 2)
	}
	allTailPositions := make(map[string]bool)
	allTailPositions[posToStr(rope[ropeLen-1])] = true
	total++

	for _, line := range strings.Split(string(fileBytes), "\n") {
		args := strings.Split(line, " ")
		direction := args[0]
		steps, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal("bad input line: %s", line)
		}
		dx := 0
		dy := 0
		if direction == "U" {
			dy = 1
		} else if direction == "R" {
			dx = 1
		} else if direction == "L" {
			dx = -1
		} else if direction == "D" {
			dy = -1
		}
		for i := 0; i < steps; i++ {
			rope[0][0] += dx
			rope[0][1] += dy
			for r := 1; r < ropeLen; r++ {
				rope[r] = updateTail(rope[r-1], rope[r])
			}
			tailPosStr := posToStr(rope[ropeLen-1])
			if !allTailPositions[tailPosStr] {
				allTailPositions[tailPosStr] = true
				total++
			}
		}
	}

	return total
}

func main() {
	const filename = "data/nine.txt"
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	ans := numberOfTailPositions(fileBytes, 2)
	log.Printf("number of tail positions: %d", ans)

	ans = numberOfTailPositions(fileBytes, 10)
	log.Printf("number of tail positions with long rope: %d", ans)

}
