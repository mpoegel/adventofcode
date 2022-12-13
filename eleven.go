package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items         *list.List
	operation     func(uint64) uint64
	test          func(uint64) bool
	monkeyOnTrue  int
	monkeyOnFalse int
	inspections   int
}

func NewMonkey() *Monkey {
	return &Monkey{
		items:         list.New(),
		operation:     nil,
		test:          nil,
		monkeyOnTrue:  -1,
		monkeyOnFalse: -1,
		inspections:   0,
	}
}

func (m *Monkey) DoOperation(worry uint64) uint64 {
	m.inspections++
	newWorry := m.operation(worry)
	if newWorry < worry {
		log.Fatal("int overflow detected")
	}
	return newWorry
}

func PrintMonkeys(monkeys []*Monkey) {
	for i, m := range monkeys {
		fmt.Printf("Monkey %d: ", i)
		elem := m.items.Front()
		for elem != nil {
			fmt.Printf("%d ", elem.Value.(uint64))
			elem = elem.Next()
		}
		fmt.Printf("\n")
	}
}

func parseMonkeys(fileBytes []byte) ([]*Monkey, uint64) {
	monkeys := make([]*Monkey, 0)
	lcm := uint64(1)
	var currMonkey *Monkey
	for i, line := range strings.Split(string(fileBytes), "\n") {
		switch i % 7 {
		case 0:
			// monkey ID
			currMonkey = NewMonkey()
		case 1:
			// starting items
			args := strings.Split(line, ": ")
			for _, itemStr := range strings.Split(args[1], ", ") {
				item, err := strconv.ParseUint(itemStr, 10, 64)
				if err != nil {
					log.Fatalf("bad monkey starting items: %s", line)
				}
				currMonkey.items.PushBack(item)
			}
		case 2:
			// operation
			args := strings.Split(line, ": ")
			equation := strings.Split(args[1], " = ")
			if strings.Contains(equation[1], "+") {
				lhs := strings.Split(equation[1], " + ")
				if lhs[1] == "old" {
					currMonkey.operation = func(old uint64) uint64 {
						return old + old
					}
				} else {
					val, err := strconv.ParseUint(lhs[1], 10, 64)
					if err != nil {
						log.Fatalf("bad operation equation: %s", line)
					}
					currMonkey.operation = func(old uint64) uint64 {
						return old + val
					}
				}
			} else if strings.Contains(equation[1], "*") {
				lhs := strings.Split(equation[1], " * ")
				if lhs[1] == "old" {
					currMonkey.operation = func(old uint64) uint64 {
						return old * old
					}
				} else {
					val, err := strconv.ParseUint(lhs[1], 10, 64)
					if err != nil {
						log.Fatalf("bad operation equation: %s", line)
					}
					currMonkey.operation = func(old uint64) uint64 {
						return old * val
					}
				}
			} else {
				log.Fatalf("unrecognized operation: %s", line)
			}
		case 3:
			// test
			args := strings.Split(line, " by ")
			testVal, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				log.Fatalf("bad test value: %s", line)
			}
			currMonkey.test = func(item uint64) bool {
				return item%testVal == 0
			}
			lcm *= testVal
		case 4:
			// if true
			args := strings.Split(line, " monkey ")
			val, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatalf("bad if true: %s", line)
			}
			currMonkey.monkeyOnTrue = val
		case 5:
			// if false
			args := strings.Split(line, " monkey ")
			val, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatalf("bad if false: %s", line)
			}
			currMonkey.monkeyOnFalse = val
		case 6:
			// empty line
			monkeys = append(monkeys, currMonkey)
		}
	}
	monkeys = append(monkeys, currMonkey)

	return monkeys, lcm
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func monkeyBusiness(fileBytes []byte, rounds int, reliefFactor float64) uint64 {
	monkeys, lcm := parseMonkeys(fileBytes)
	for i := 1; i <= rounds; i++ {
		for m := range monkeys {
			numItems := monkeys[m].items.Len()
			for w := 0; w < numItems; w++ {
				item := monkeys[m].items.Front()
				monkeys[m].items.Remove(item)
				worry := monkeys[m].DoOperation(item.Value.(uint64))
				worry %= lcm
				worry = uint64(math.Floor(float64(worry) / reliefFactor))
				toMonkey := monkeys[m].monkeyOnFalse
				if monkeys[m].test(worry) {
					toMonkey = monkeys[m].monkeyOnTrue
				}
				monkeys[toMonkey].items.PushBack(worry)
			}
		}
	}

	h := &IntHeap{}
	heap.Init(h)
	for _, m := range monkeys {
		heap.Push(h, m.inspections)
	}

	first := heap.Pop(h).(int)
	second := heap.Pop(h).(int)

	return uint64(first) * uint64(second)
}

func main() {
	const filename = "data/eleven.txt"
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	rounds := 20
	reliefFactor := 3.0
	ans := monkeyBusiness(fileBytes, rounds, reliefFactor)
	log.Printf("monkey business level after %d rounds with relief %0.1f: %d", rounds, reliefFactor, ans)

	rounds = 10000
	reliefFactor = 1.0
	ans = monkeyBusiness(fileBytes, rounds, reliefFactor)
	log.Printf("monkey business level after %d rounds with relief %0.1f: %d", rounds, reliefFactor, ans)
}
