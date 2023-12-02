package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// content, err := os.ReadFile("inputs/example_input.txt")
	content, err := os.ReadFile("inputs/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	file_contents := string(content)
	lines := strings.Split(file_contents, "\n")

	var sum_accum = 0

	for _, line := range lines {
		runes := []rune(line)
		var lineNumber = ""

		// Get 1st number
		for _, rune := range runes {
			if unicode.IsNumber(rune) {
				lineNumber += string(rune)
				break
			}
		}

		// Get 2nd Number
		for i := len(runes) - 1; i >= 0; i-- {
			rune := runes[i]
			if unicode.IsNumber(rune) {
				lineNumber += string(rune)
				break
			}
		}

		numericVal, err := strconv.Atoi(lineNumber)
		if err != nil {
			log.Fatal(err)
		}

		sum_accum += numericVal

	}

	fmt.Println(sum_accum)
}
