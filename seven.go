package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Name     string
	IsFile   bool
	Size     uint64
	Parent   *Node
	Children map[string]*Node
}

func (n *Node) UpdateSize() uint64 {
	total := uint64(0)
	for _, child := range n.Children {
		total += child.Size
	}
	n.Size = total
	return total
}

func createFileTree(fileBytes []byte) *Node {
	root := Node{"/", false, 0, nil, make(map[string]*Node, 0)}
	currentNode := &root

	for _, line := range strings.Split(string(fileBytes), "\n")[1:] {
		args := strings.Split(line, " ")
		if args[0] == "$" {
			if args[1] == "ls" {
				// do nothing?
			} else if args[1] == "cd" {
				args := strings.Split(line, " ")
				if args[2] == ".." {
					currentNode.UpdateSize()
					currentNode = currentNode.Parent
				} else {
					currentNode = currentNode.Children[args[2]]
				}
			} else {
				// panic: unknown arg
				log.Fatalf("parse error: unknown cmd: %s", line)
			}
		} else if strings.Contains(line, "dir") {
			// directory listing
			currentNode.Children[args[1]] = &Node{
				Name:     args[1],
				IsFile:   false,
				Size:     0,
				Parent:   currentNode,
				Children: make(map[string]*Node),
			}
		} else {
			// regular file
			size, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				log.Fatalf("parse error: bad size: %s", line)
			}
			currentNode.Children[args[1]] = &Node{
				Name:     args[1],
				IsFile:   true,
				Size:     size,
				Parent:   currentNode,
				Children: nil,
			}
		}
	}
	for currentNode != nil {
		currentNode.UpdateSize()
		currentNode = currentNode.Parent
	}
	return &root
}

func dfsSum(node *Node, maxSize uint64) uint64 {
	var total uint64
	if !node.IsFile && node.Size <= maxSize {
		total += node.Size
	}
	for _, child := range node.Children {
		total += dfsSum(child, maxSize)
	}
	return total
}

func getSumOfDirectoriesWithMaxSize(fileBytes []byte, maxSize uint64) uint64 {
	root := createFileTree(fileBytes)
	return dfsSum(root, maxSize)
}

func dfsMinFind(node *Node, minSize uint64) uint64 {
	var runningMin uint64
	runningMin = math.MaxUint64
	for _, child := range node.Children {
		if !child.IsFile {
			thisMin := dfsMinFind(child, minSize)
			if thisMin >= minSize && thisMin <= runningMin {
				runningMin = thisMin
			}
		}
	}
	if node.Size > minSize && node.Size < runningMin {
		runningMin = node.Size
	}
	return runningMin
}

func findDirectoryToDelete(fileBytes []byte) uint64 {
	const diskSize = uint64(70000000)
	const updateSize = uint64(30000000)
	root := createFileTree(fileBytes)
	freeSpace := diskSize - root.Size
	deleteSizeMin := updateSize - freeSpace
	return dfsMinFind(root, deleteSizeMin)
}

func main() {
	const filename = "data/seven.txt"
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	maxSize := uint64(100000)
	ans := getSumOfDirectoriesWithMaxSize(fileBytes, maxSize)
	log.Printf("sum of directories with max size %d: %d", maxSize, ans)

	ans = findDirectoryToDelete(fileBytes)
	log.Printf("size of directory to delete: %d", ans)
}
