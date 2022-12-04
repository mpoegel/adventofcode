package main

import (
	"log"
	"os"
	"strings"
)

func getPriority(item byte) int {
	if int(item) > 97 {
		return int(item) - 97
	}
	return int(item) - 65 + 26
}

func findItemsInBothCompartments(filesBytes []byte) int {
	total := 0
	for _, rucksack := range strings.Split(string(filesBytes), "\n") {
		counts := make([]int, 52)
		halfLen := len(rucksack) / 2
		for i, item := range rucksack {
			p := getPriority(byte(item))
			if i < halfLen {
				counts[p]++
			} else {
				if counts[p] > 0 {
					total += p + 1
					break
				}
			}
		}
	}
	return total
}

func findGroupBadges(filesBytes []byte) int {
	total := 0
	counts := make([]int, 52)
	for g, rucksack := range strings.Split(string(filesBytes), "\n") {
		for _, item := range rucksack {
			p := getPriority(byte(item))
			if g%3 == 0 {
				counts[p] |= 1
			} else if g%3 == 1 {
				counts[p] |= 2
			} else {
				if counts[p] == 3 {
					total += p + 1
					counts = make([]int, 52)
					break
				}
			}
		}
	}
	return total
}

func main() {
	const filename = "data/three.txt"
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	ans := findItemsInBothCompartments(fileBytes)
	log.Printf("total priority: %d", ans)

	ans = findGroupBadges(fileBytes)
	log.Printf("total badge priority: %d", ans)
}
