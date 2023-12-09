package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func parseInput(fileContents string) ([][]int, error) {
	lines := strings.Split(fileContents, "\n")

	response := [][]int{}
	for _, line := range lines {
		tokens := strings.Fields(line)
		numberLine := []int{}
		for _, token := range tokens {
			number, err := strconv.Atoi(token)
			if err != nil {
				return response, err
			}
			numberLine = append(numberLine, number)
		}
		response = append(response, numberLine)
	}

	return response, nil
}

func predictNextNumber(numbers []int) int {

	stacks := [][]int{numbers}

	for {
		lastStackIdx := len(stacks) - 1
		lastStack := stacks[lastStackIdx]
		lastStackAllZeros := true
		for _, i := range lastStack {
			if i != 0 {
				lastStackAllZeros = false
			}
		}
		if lastStackAllZeros {
			break
		}

		newStack := []int{}
		for i := 1; i < len(lastStack); i++ {
			a := lastStack[i]
			b := lastStack[i-1]
			newStack = append(newStack, a-b)
		}
		stacks = append(stacks, newStack)

	}

	// Add a ZERO to the last stack
	numStacks := len(stacks)
	stacks[numStacks-1] = append(stacks[numStacks-1], 0)

	// Work up the stacks, adding 2nd to last to last of the next row
	for i := numStacks - 2; i >= 0; i-- {
		currStack := stacks[i]
		stackLen := len(currStack)
		nextStack := stacks[i+1]
		nextStackLen := len(nextStack)
		a := currStack[stackLen-1]
		b := nextStack[nextStackLen-1]
		stacks[i] = append(stacks[i], a+b)
	}

	return stacks[0][len(stacks[0])-1]
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

	inputLines, err := parseInput(string(content))

	if err != nil {
		fmt.Println(err)
		return
	}

	predictionsSum := 0

	for _, inputLine := range inputLines {
		predictionsSum += predictNextNumber(inputLine)
	}

	fmt.Println(predictionsSum)
}
