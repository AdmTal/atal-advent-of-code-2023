package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Node struct {
	north   *Node
	south   *Node
	east    *Node
	west    *Node
	steps   int
	visited bool
	posX    int
	posY    int
	char    string
}

func parseInput(fileContents string) (int, int, [][]Node) {
	SposX := 0
	SposY := 0

	lines := strings.Split(fileContents, "\n")

	numRows := len(lines)
	numCols := len(lines[0])
	graph := make([][]Node, numRows)
	for row := range graph {
		graph[row] = make([]Node, numCols)
	}

	for posY, line := range lines {
		for posX, char := range line {
			graph[posY][posX].posX = posX
			graph[posY][posX].posY = posY
			c := string(char)
			graph[posY][posX].char = c
			if c == "." {
				continue
			}

			if c == "S" {
				SposX = posX
				SposY = posY
			}

			currNode := &graph[posY][posX]

			if c == "S" || c == "|" || c == "L" || c == "J" {
				// Connect to North
				if posY >= 1 {
					currNode.north = &graph[posY-1][posX]
					// graph[posY-1][posX].south = currNode
				}
			}

			if c == "S" || c == "|" || c == "7" || c == "F" {
				// Connect to South
				if posY < numRows-1 {
					currNode.south = &graph[posY+1][posX]
					// graph[posY+1][posX].north = currNode
				}
			}

			if c == "S" || c == "-" || c == "L" || c == "F" {
				// Connect to East
				if posX < numCols-1 {
					currNode.east = &graph[posY][posX+1]
					// graph[posY][posX+1].west = currNode
				}
			}

			if c == "S" || c == "-" || c == "J" || c == "7" {
				// Connect to West
				if posX >= 1 {
					currNode.west = &graph[posY][posX-1]
					// graph[posY][posX-1].east = currNode
				}
			}
		}
	}

	return SposX, SposY, graph
}

func main() {

	// Figure out where the current file is
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable info:", err)
	}
	exeDir := filepath.Dir(exePath)

	// Load the input file from the same dir
	// filePath := filepath.Join(exeDir, "example_input_2.txt")
	filePath := filepath.Join(exeDir, "input.txt")

	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fileContents := string(content)
	SposX, SposY, graph := parseInput(fileContents)

	startNode := &graph[SposY][SposX]
	startNode.visited = true

	toVisit := []*Node{}

	if startNode.north != nil && startNode.north.south == startNode {
		toVisit = append(toVisit, startNode.north)
		startNode.north.steps += 1
	}

	fmt.Printf("startNode: %p, startNode.east: %p, startNode.east.west: %p\n", &startNode, startNode.east, startNode.east.west)
	if startNode.east != nil && startNode.east.west == startNode {
		toVisit = append(toVisit, startNode.east)
		startNode.east.steps += 1
	}

	if startNode.south != nil && startNode.south.north == startNode {
		toVisit = append(toVisit, startNode.south)
		startNode.south.steps += 1
	}

	if startNode.west != nil && startNode.west.east == startNode {
		toVisit = append(toVisit, startNode.west)
		startNode.west.steps += 1
	}

	for len(toVisit) > 0 {
		currNode, nextToVisit := toVisit[0], toVisit[1:]

		if currNode.visited {
			toVisit = nextToVisit
			continue
		}

		currNode.visited = true

		allEdgesVisited := true
		if currNode.north != nil && !currNode.north.visited {
			nextToVisit = append(nextToVisit, currNode.north)
			currNode.north.steps = currNode.steps + 1
			allEdgesVisited = false
		}

		if currNode.east != nil && !currNode.east.visited {
			nextToVisit = append(nextToVisit, currNode.east)
			currNode.east.steps = currNode.steps + 1
			allEdgesVisited = false
		}

		if currNode.south != nil && !currNode.south.visited {
			nextToVisit = append(nextToVisit, currNode.south)
			currNode.south.steps = currNode.steps + 1
			allEdgesVisited = false
		}

		if currNode.west != nil && !currNode.west.visited {
			nextToVisit = append(nextToVisit, currNode.west)
			currNode.west.steps = currNode.steps + 1
			allEdgesVisited = false
		}

		if allEdgesVisited {
			fmt.Println("steps=", currNode.steps)
		}

		toVisit = nextToVisit
	}
	// for _, row := range graph {
	// 	fmt.Println()
	// 	for _, node := range row {
	// 		if node.char == "." {
	// 			fmt.Print(".")
	// 			continue
	// 		}
	// 		if node.char == "S" {
	// 			fmt.Print("S")
	// 			continue
	// 		}
	// 		if node.steps < 10 {
	// 			fmt.Print(node.steps)
	// 		} else {
	// 			fmt.Print("*")
	// 		}
	// 	}
	// }
}
