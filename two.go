package main

import (
	"log"
	"os"
	"strings"
)

/*
A/X - rock -> 1
B/Y - paper -> 2
C/Z - scissors -> 3
win/draw/loss - 6/3/0
*/
func scoreGame(opp, you byte) int {
	switch opp {
	case 'A': // rock
		switch you {
		case 'X':
			return 3 + 1
		case 'Y':
			return 6 + 2
		case 'Z':
			return 0 + 3
		}
	case 'B': // paper
		switch you {
		case 'X':
			return 0 + 1
		case 'Y':
			return 3 + 2
		case 'Z':
			return 6 + 3
		}
	case 'C': // scissors
		switch you {
		case 'X':
			return 6 + 1
		case 'Y':
			return 0 + 2
		case 'Z':
			return 3 + 3
		}
	default:
		log.Fatalf("invalid move: %s", string(opp))
	}
	return -1
}

/*
A - rock -> 1
B - paper -> 2
C - scissors -> 3
win/draw/loss - 6/3/0
X - lose
Y - draw
Z - win
*/
func scoreGame2(opp, desiredOutcome byte) int {
	switch opp {
	case 'A': // rock
		switch desiredOutcome {
		case 'X':
			return 0 + 3
		case 'Y':
			return 3 + 1
		case 'Z':
			return 6 + 2
		}
	case 'B': // paper
		switch desiredOutcome {
		case 'X':
			return 0 + 1
		case 'Y':
			return 3 + 2
		case 'Z':
			return 6 + 3
		}
	case 'C': // scissors
		switch desiredOutcome {
		case 'X':
			return 0 + 2
		case 'Y':
			return 3 + 3
		case 'Z':
			return 6 + 1
		}
	default:
		log.Fatalf("invalid move: %s", string(opp))
	}
	return -1
}

func calculateTotalScore(fileBytes []byte) int {
	total := 0
	for _, match := range strings.Split(string(fileBytes), "\n") {
		total += scoreGame(match[0], match[2])
	}
	return total
}

func calculateTotalScore2(fileBytes []byte) int {
	total := 0
	for _, match := range strings.Split(string(fileBytes), "\n") {
		total += scoreGame2(match[0], match[2])
	}
	return total
}

func main() {
	const filename = "data/two.txt"
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	ans := calculateTotalScore(fileBytes)
	log.Printf("total score: %d", ans)

	ans = calculateTotalScore2(fileBytes)
	log.Printf("total score again: %d", ans)
}
