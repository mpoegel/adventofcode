package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func signalStrenth(cycle, regisiterVal, start, end, interval int) int {
	if cycle == start || (cycle <= end && (cycle-start)%interval == 0) {
		strength := regisiterVal * cycle
		log.Printf("signal strength at cycle %d, register value %d: %d", cycle, regisiterVal, strength)
		return strength
	}
	return 0
}

func sumSignalStrengths(fileBytes []byte, start, end, interval int) int {
	totalSum := 0
	cycle := 1
	regisiterVal := 1

	for _, line := range strings.Split(string(fileBytes), "\n") {
		args := strings.Split(line, " ")
		if args[0] == "noop" {
			cycle++
			totalSum += signalStrenth(cycle, regisiterVal, start, end, interval)
		} else if args[0] == "addx" {
			cycle++
			totalSum += signalStrenth(cycle, regisiterVal, start, end, interval)
			cycle++
			val, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatalf("bad addx instruction: %s", line)
			}
			regisiterVal += val
			totalSum += signalStrenth(cycle, regisiterVal, start, end, interval)
		} else {
			log.Fatalf("bad instruction: %s", line)
		}
	}

	return totalSum
}

type Simluation struct {
	crt       [][]byte
	cycle     int
	spritePos int
	drawIndex int
}

func NewSimulation() Simluation {
	s := Simluation{
		crt:       make([][]byte, 6),
		cycle:     1,
		spritePos: 1,
		drawIndex: 0,
	}
	for i := 0; i < len(s.crt); i++ {
		s.crt[i] = make([]byte, 40)
	}
	return s
}

func (s *Simluation) Noop() {
	s.Cycle()
}

func (s *Simluation) Add(x int) {
	s.Cycle()
	s.Cycle()
	s.spritePos += x
	s.drawIndex = 0
}

func (s *Simluation) Cycle() {
	row := int(math.Floor(float64(s.cycle-1) / float64(len(s.crt[0]))))
	col := (s.cycle - 1) % len(s.crt[0])
	if math.Abs(float64(col-s.spritePos)) <= 1 {
		s.crt[row][col] = '#'
	} else {
		s.crt[row][col] = '.'
	}
	s.cycle++
}

func (s *Simluation) Draw() {
	for _, row := range s.crt {
		for _, val := range row {
			fmt.Printf("%s", string(val))
		}
		fmt.Printf("\n")
	}
}

func processInstructionsAndDraw(fileBytes []byte) {
	sim := NewSimulation()
	for _, line := range strings.Split(string(fileBytes), "\n") {
		args := strings.Split(line, " ")
		if args[0] == "noop" {
			sim.Noop()
		} else if args[0] == "addx" {
			val, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatalf("bad addx instruction: %s", line)
			}
			sim.Add(val)
		} else {
			log.Fatalf("bad instruction: %s", line)
		}
	}
	sim.Draw()
}

func main() {
	const filename = "data/ten.txt"
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	ans := sumSignalStrengths(fileBytes, 20, 220, 40)
	log.Printf("sum of signal strengths: %d", ans)

	processInstructionsAndDraw(fileBytes)
}
