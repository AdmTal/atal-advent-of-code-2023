package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type GameGubeSet struct {
	numBlue  int
	numRed   int
	numGreen int
}

type GameRecord struct {
	gameId   int
	cubeSets []GameGubeSet
}

func (gr *GameRecord) isPossible(numRed int, numBlue int, numGreen int) bool {
	for _, cubeSet := range gr.cubeSets {
		if cubeSet.numBlue > numBlue || cubeSet.numGreen > numGreen || cubeSet.numRed > numRed {
			return false
		}
	}
	return true
}

func parseLine(line string) (GameRecord, error) {
	// Example Line: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"

	// First, split into "Game 3" and "8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	gameNumberClause := strings.Split(line, ": ")
	cubeSetsClause := gameNumberClause[1]

	// Then, split "Game 3" into "Game" and "3"
	gameNumberClause = strings.Split(gameNumberClause[0], " ")
	gameNumber, err := strconv.Atoi(gameNumberClause[1])
	if err != nil {
		panic(err)
	}

	// Then split the rest of the line into "8 green, 6 blue, 20 red", "5 blue, 4 red, 13 green", and "5 green, 1 red"
	cubeSetClauses := strings.Split(cubeSetsClause, "; ")
	cubeSets := []GameGubeSet{}

	// yadda yadda yadda...
	for _, cubeSet := range cubeSetClauses {
		cubeClauses := strings.Split(cubeSet, ", ")
		numRed := 0
		numBlue := 0
		numGreen := 0
		for _, cubeClause := range cubeClauses {
			cubeData := strings.Split(cubeClause, " ")
			num, err := strconv.Atoi(cubeData[0])
			if err != nil {
				return GameRecord{}, err
			}
			color := cubeData[1]

			switch color {
			case "red":
				numRed = num
			case "blue":
				numBlue = num
			case "green":
				numGreen = num
			default:
				panic("Color not recognized")
			}

		}

		cubeSets = append(cubeSets, GameGubeSet{
			numBlue:  numBlue,
			numRed:   numRed,
			numGreen: numGreen,
		})

	}

	return GameRecord{
		gameId:   gameNumber,
		cubeSets: cubeSets,
	}, nil
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
	file_contents := string(content)
	lines := strings.Split(file_contents, "\n")

	accum_possible := 0

	numRed := 12
	numGreen := 13
	numBlue := 14

	for _, line := range lines {
		gameRecord, err := parseLine(line)
		if err != nil {
			panic(err)
		}

		if gameRecord.isPossible(numRed, numBlue, numGreen) {
			accum_possible += gameRecord.gameId
		}
	}

	fmt.Println(accum_possible)
}
