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
	cardNum        int
	winningNumbers map[int]bool
	playedNumbers  map[int]bool
}

func parseCardFromLine(line string) (Card, error) {
	tokens := strings.Split(line, ": ")
	cardIdString := strings.Fields(tokens[0])[1]
	cardNum, err := strconv.Atoi(cardIdString)
	if err != nil {
		return Card{}, err
	}
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
		cardNum:        cardNum,
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

	cardLookup := map[int]Card{}
	cardsToPlay := []Card{}

	for _, line := range lines {
		card, err := parseCardFromLine(line)

		if err != nil {
			panic(err)
		}

		cardLookup[card.cardNum] = card
		cardsToPlay = append(cardsToPlay, card)
	}

	numCards := len(cardsToPlay)

	for len(cardsToPlay) > 0 {
		cardToPlay := cardsToPlay[0]
		cardsToPlay = cardsToPlay[1:]

		winningNumbers := 0
		for playedNumber := range cardToPlay.playedNumbers {
			value, has_key := cardToPlay.winningNumbers[playedNumber]
			if has_key && value {
				winningNumbers += 1
			}
		}

		if winningNumbers > 0 {
			for i := cardToPlay.cardNum + 1; i <= cardToPlay.cardNum+winningNumbers; i++ {
				cardsToPlay = append(cardsToPlay, cardLookup[i])
				numCards += 1
			}
		}
	}
	fmt.Println(numCards)
}
