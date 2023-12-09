package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func parseInput(fileContents string) ([]string, map[string][]string) {

	steps := []string{}
	graph := map[string][]string{}

	lines := strings.Split(fileContents, "\n")

	stepsString, lines := lines[0], lines[2:]

	for _, char := range stepsString {
		steps = append(steps, string(char))
	}

	for _, line := range lines {
		node, left, right := line[0:3], line[7:10], line[12:15]
		graph[node] = []string{left, right}
	}

	return steps, graph
}

func main() {

	// Figure out where the current file is
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable info:", err)
	}
	exeDir := filepath.Dir(exePath)

	// Load the input file from the same dir
	// filePath := filepath.Join(exeDir, "example_input.txt")
	filePath := filepath.Join(exeDir, "input.txt")

	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fileContents := string(content)

	steps, graph := parseInput(fileContents)

	currNodes := []string{}
	pathLengthsToZ := []int{}

	for key := range graph {
		if string(key[len(key)-1]) == "A" {
			currNodes = append(currNodes, key)
			pathLengthsToZ = append(pathLengthsToZ, 0)
		}
	}

	numSteps := 0
	i := 0

	for {
		// Take a step
		nextStep := steps[i]
		numSteps += 1
		i += 1
		if i == len(steps) {
			i = 0
		}

		// Move curr nodes
		for nodeIdx, currNode := range currNodes {

			if string(currNode[len(currNode)-1]) == "Z" {
				// Skip paths that already found Z
				continue
			}

			left, right := graph[currNode][0], graph[currNode][1]

			if nextStep == "L" {
				currNode = left
			} else {
				currNode = right
			}

			currNodes[nodeIdx] = currNode

			if string(currNode[len(currNode)-1]) == "Z" {
				pathLengthsToZ[nodeIdx] = numSteps
			}
		}

		// Check if we're finished
		numZs := 0
		for _, node := range currNodes {
			if string(node[len(node)-1]) == "Z" {
				numZs += 1
			}
		}

		if numZs == len(currNodes) {
			break
		}

	}

	fmt.Println(pathLengthsToZ)

	// Pasted that into LCM calculator online

	// I cheated on this - I got stuck, and learned the trick was LCM on Reddit
	// From there, I just calculated each path's length to find 1st z, and did LCM for that -- externally

}
