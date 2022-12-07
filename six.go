package main

import (
	"container/list"
	"log"
	"os"
)

func getCharsBeforeStartOfPacket(fileBytes []byte, startOfMark int) int {
	queue := list.New()
	counts := make(map[byte]int)

	for i, b := range fileBytes {
		counts[b]++
		queue.PushFront(b)
		if queue.Len() > startOfMark {
			val := queue.Back()
			popByte := val.Value.(byte)
			counts[popByte]--
			if counts[popByte] == 0 {
				delete(counts, popByte)
			}
			queue.Remove(val)
			if len(counts) == startOfMark {
				return i + 1
			}
		}
	}

	return -1
}

func main() {
	const filename = "data/six.txt"
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	ans := getCharsBeforeStartOfPacket(fileBytes, 4)
	log.Printf("characters before start of packet: %d", ans)

	ans = getCharsBeforeStartOfPacket(fileBytes, 14)
	log.Printf("characters before start of message: %d", ans)

}
