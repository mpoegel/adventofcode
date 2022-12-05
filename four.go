package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func parsePair(strPair string) []int {
	parts := strings.Split(strPair, "-")
	res := make([]int, 2)
	if start, err := strconv.Atoi(parts[0]); err == nil {
		res[0] = start
	}
	if end, err := strconv.Atoi(parts[1]); err == nil {
		res[1] = end
	}
	return res
}

func findNumberofCompleteOverlaps(fileBytes []byte) int {
	total := 0
	for _, group := range strings.Split(string(fileBytes), "\n") {
		groupParts := strings.Split(group, ",")
		first := parsePair(groupParts[0])
		second := parsePair(groupParts[1])
		if (first[0] <= second[0] && second[1] <= first[1]) || (second[0] <= first[0] && first[1] <= second[1]) {
			total++
		}
	}
	return total
}

func findNumberofPartialOverlaps(fileBytes []byte) int {
	total := 0
	for _, group := range strings.Split(string(fileBytes), "\n") {
		groupParts := strings.Split(group, ",")
		first := parsePair(groupParts[0])
		second := parsePair(groupParts[1])
		if (second[0] <= first[1] && second[0] >= first[0]) || (first[0] <= second[1] && first[0] >= second[0]) {
			total++
		}
	}
	return total
}

func main() {
	const filename = "data/four.txt"
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	ans := findNumberofCompleteOverlaps(fileBytes)
	log.Printf("number of completely overlapping assignments: %d", ans)

	ans = findNumberofPartialOverlaps(fileBytes)
	log.Printf("number of partial overlapping assignments: %d", ans)
}
