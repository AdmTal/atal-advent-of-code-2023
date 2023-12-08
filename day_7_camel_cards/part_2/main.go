package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type CamelCardHand struct {
	cards []string
	bid   int
}

type CamelCardHands []CamelCardHand

func (cch CamelCardHands) Len() int      { return len(cch) }
func (cch CamelCardHands) Swap(i, j int) { cch[i], cch[j] = cch[j], cch[i] }
func (cch CamelCardHands) Less(i, j int) bool {
	handTypeRanks := map[string]int{
		FIVE_OF_A_KIND:  7,
		FOUR_OF_A_KIND:  6,
		FULL_HOUSE:      5,
		THREE_OF_A_KIND: 4,
		TWO_PAIR:        3,
		ONE_PAIR:        2,
		HIGH_CARD:       1,
	}

	iRank := handTypeRanks[cch[i].getType()]
	jRank := handTypeRanks[cch[j].getType()]

	if iRank != jRank {
		return iRank < jRank
	}

	cardRanks := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
		"J": 1,
	}

	for idx := 0; idx < len(cch[i].cards); idx++ {
		if cch[i].cards[idx] != cch[j].cards[idx] {
			return cardRanks[cch[i].cards[idx]] < cardRanks[cch[j].cards[idx]]
		}
	}

	return false
}

const FIVE_OF_A_KIND = "FIVE_OF_A_KIND"
const FOUR_OF_A_KIND = "FOUR_OF_A_KIND"
const FULL_HOUSE = "FULL_HOUSE"
const THREE_OF_A_KIND = "THREE_OF_A_KIND"
const TWO_PAIR = "TWO_PAIR"
const ONE_PAIR = "ONE_PAIR"
const HIGH_CARD = "HIGH_CARD"

func sliceMatches(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (cch *CamelCardHand) getType() string {
	cardCounts := map[string]int{}
	jCount := 0
	for _, card := range cch.cards {
		if card == "J" {
			jCount += 1
			continue
		}
		cardCounts[card] += 1
	}

	// Inject J as best card
	maxKey := ""
	maxValue := 0
	for key, value := range cardCounts {
		if value > maxValue {
			maxKey = key
			maxValue = value
			continue
		}
	}

	cardCounts[maxKey] += jCount

	cardCountValues := []int{}
	for _, value := range cardCounts {
		cardCountValues = append(cardCountValues, value)
	}

	sort.Ints(cardCountValues)

	if sliceMatches(cardCountValues, []int{5}) {
		return FIVE_OF_A_KIND
	}

	if sliceMatches(cardCountValues, []int{1, 4}) {
		return FOUR_OF_A_KIND
	}

	if sliceMatches(cardCountValues, []int{2, 3}) {
		return FULL_HOUSE
	}

	if sliceMatches(cardCountValues, []int{1, 1, 3}) {
		return THREE_OF_A_KIND
	}

	if sliceMatches(cardCountValues, []int{1, 2, 2}) {
		return TWO_PAIR
	}

	if sliceMatches(cardCountValues, []int{1, 1, 1, 2}) {
		return ONE_PAIR
	}

	return HIGH_CARD
}

func parseInput(fileContents string) CamelCardHands {
	lines := strings.Split(fileContents, "\n")

	hands := CamelCardHands{}

	for _, line := range lines {
		tokens := strings.Fields(line)
		bid, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}
		hands = append(
			hands,
			CamelCardHand{
				cards: strings.Split(tokens[0], ""),
				bid:   bid,
			},
		)
	}
	return hands
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

	hands := parseInput(fileContents)

	sort.Sort(hands)

	totalWinnings := 0

	for idx, hand := range hands {
		rank := idx + 1
		totalWinnings += rank * hand.bid
	}

	fmt.Println(totalWinnings)

}
