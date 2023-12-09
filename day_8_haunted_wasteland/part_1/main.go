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
	// filePath := filepath.Join(exeDir, "example_input_2.txt")
	filePath := filepath.Join(exeDir, "input.txt")

	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fileContents := string(content)

	steps, graph := parseInput(fileContents)

	currNode := "AAA"
	numSteps := 0
	i := 0
	for currNode != "ZZZ" {
		left, right := graph[currNode][0], graph[currNode][1]

		nextStep := steps[i]

		if nextStep == "L" {
			currNode = left
		} else {
			currNode = right
		}

		i += 1
		if i == len(steps) {
			i = 0
		}

		numSteps += 1
	}

	fmt.Println(numSteps)

}
