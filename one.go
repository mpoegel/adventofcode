package main

import (
	"container/heap"
	"log"
	"os"
	"strconv"
	"strings"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findElfWithMostCalories(calorieList []byte) int {
	runningMax := 0
	elfTotal := 0
	calories := strings.Split(string(calorieList), "\n")
	for _, cal := range calories {
		if len(cal) == 0 {
			if elfTotal > runningMax {
				runningMax = elfTotal
			}
			elfTotal = 0
		} else if calInt, err := strconv.Atoi(cal); err != nil {
			log.Fatal(err)
		} else {
			elfTotal += calInt
		}
	}
	if elfTotal > runningMax {
		runningMax = elfTotal
	}
	return runningMax
}

func findTopKElvesWithMostCalories(calorieList []byte, k int) int {
	topElves := &IntHeap{}
	heap.Init(topElves)
	elfTotal := 0
	calories := strings.Split(string(calorieList), "\n")
	for _, cal := range calories {
		if len(cal) == 0 {
			heap.Push(topElves, elfTotal)
			if topElves.Len() > k {
				heap.Pop(topElves)
			}
			elfTotal = 0
		} else if calInt, err := strconv.Atoi(cal); err != nil {
			log.Fatal(err)
		} else {
			elfTotal += calInt
		}
	}
	heap.Push(topElves, elfTotal)
	if topElves.Len() > k {
		heap.Pop(topElves)
	}

	total := 0
	for topElves.Len() > 0 {
		cal := heap.Pop(topElves).(int)
		log.Printf("top elf: %d", cal)
		total += cal
	}
	return total
}

func main() {
	const filename = "data/one.txt"
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	ans := findElfWithMostCalories(fileBytes)
	log.Printf("elf with most calories has %d calories\n", ans)

	ans = findTopKElvesWithMostCalories(fileBytes, 3)
	log.Printf("total calories of top 3 elves: %d", ans)
}
