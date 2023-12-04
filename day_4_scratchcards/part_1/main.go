package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Card struct {
	cardId         string
	winningNumbers map[int]bool
	playedNumbers  map[int]bool
}

func parseCardFromLine(line string) (Card, error) {
	tokens := strings.Split(line, ": ")
	cardId := tokens[0]
	numbers := strings.Split(tokens[1], " | ")
	winningNumbers := map[int]bool{}

	for _, numberString := range strings.Fields(numbers[0]) {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			return Card{}, err
		}
		winningNumbers[number] = true
	}

	playedNumbers := map[int]bool{}
	for _, numberString := range strings.Fields(numbers[1]) {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			return Card{}, err
		}
		playedNumbers[number] = true
	}

	return Card{
		cardId:         cardId,
		winningNumbers: winningNumbers,
		playedNumbers:  playedNumbers,
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

	cardScoreSums := 0

	for _, line := range lines {
		card, err := parseCardFromLine(line)

		if err != nil {
			panic(err)
		}

		cardScore := 0
		for playedNumber := range card.playedNumbers {
			value, has_key := card.winningNumbers[playedNumber]
			if has_key && value {
				if cardScore == 0 {
					cardScore = 1
				} else {
					cardScore *= 2
				}
			}
		}

		cardScoreSums += cardScore
	}

	fmt.Println(cardScoreSums)

}
