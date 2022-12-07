package main

import (
	"container/list"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseStacks(rows []string) []*list.List {
	width := len(rows[0])
	numStacks := (width + 1) / 4
	stacks := make([]*list.List, numStacks)
	for i := 0; i < numStacks; i++ {
		stacks[i] = list.New()
	}
	for _, row := range rows {
		for s := 0; s < numStacks; s++ {
			val := row[(s*4)+1]
			if val != ' ' {
				stacks[s].PushBack(val)
			}
		}
	}
	return stacks
}

func parseInstruction(instruct string) []int {
	res := make([]int, 3)
	parts := strings.Split(instruct, " ")
	if numToMove, err := strconv.Atoi(parts[1]); err == nil {
		res[0] = numToMove
	}
	if fromStack, err := strconv.Atoi(parts[3]); err == nil {
		res[1] = fromStack
	}
	if toStack, err := strconv.Atoi(parts[5]); err == nil {
		res[2] = toStack
	}
	return res
}

func getTops(stacks []*list.List) string {
	res := make([]byte, len(stacks))
	for s := 0; s < len(stacks); s++ {
		res[s] = (stacks[s].Front().Value).(byte)
	}

	return string(res)
}

func findCratesOnTop(fileBytes []byte) string {
	lines := strings.Split(string(fileBytes), "\n")
	var i int
	for i = 0; len(lines[i]) > 0; i++ {
	}
	stacks := parseStacks(lines[:i-1])
	i++
	for ; i < len(lines); i++ {
		instruction := parseInstruction(lines[i])
		for k := 0; k < instruction[0]; k++ {
			fromStack := instruction[1] - 1
			toStack := instruction[2] - 1
			crate := stacks[fromStack].Front()
			stacks[fromStack].Remove(crate)
			stacks[toStack].PushFront(crate.Value)
		}
	}
	return getTops(stacks)
}

func findCratesOnTop2(fileBytes []byte) string {
	lines := strings.Split(string(fileBytes), "\n")
	var i int
	for i = 0; len(lines[i]) > 0; i++ {
	}
	stacks := parseStacks(lines[:i-1])
	i++
	for ; i < len(lines); i++ {
		instruction := parseInstruction(lines[i])
		var last *list.Element
		for k := 0; k < instruction[0]; k++ {
			fromStack := instruction[1] - 1
			toStack := instruction[2] - 1
			crate := stacks[fromStack].Front()
			stacks[fromStack].Remove(crate)
			if last != nil {
				last = stacks[toStack].InsertAfter(crate.Value, last)
			} else {
				last = stacks[toStack].PushFront(crate.Value)
			}
		}
	}
	return getTops(stacks)
}

func main() {
	const filename = "data/five.txt"
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	ans := findCratesOnTop(fileBytes)
	log.Printf("crates on top after moves: %s", ans)

	ans = findCratesOnTop2(fileBytes)
	log.Printf("crates on top after moves with CrateMover 9001: %s", ans)
}
