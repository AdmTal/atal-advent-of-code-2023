package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

func getNumber(runes []rune) (int, error) {

	var numbers []string = []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	index_of_earliest_number := math.MaxInt
	earliest_number := ""

	index_of_latest_number := -1
	latest_number := ""

	// For every target number to match
	for i := 0; i < len(numbers); i++ {
		number := numbers[i]

		// for every starting position in runes
		for j := 0; j < len(runes); j++ {

			// If this Char is a digit, keep track of that
			if unicode.IsDigit(runes[j]) {

				// Keep track of earliest number
				if j < index_of_earliest_number {
					index_of_earliest_number = j
					_, err := strconv.Atoi(string(runes[j]))
					if err != nil {
						return -1, err
					}
					earliest_number = string(string(runes[j]))
				}

				// Keep track of latest number
				if j > index_of_latest_number {
					index_of_latest_number = j
					_, err := strconv.Atoi(string(runes[j]))
					if err != nil {
						return -1, err
					}
					latest_number = string(runes[j])
				}

			} else if len(runes)-j >= len(number) {
				// Otherwise - if there are enough letters to spell out the target number

				// Take slice of runes
				slice := runes[j : j+len(number)]

				// If it matches the target number
				if string(slice) == number {

					// Keep track of earliest number
					if j < index_of_earliest_number {
						index_of_earliest_number = j
						earliest_number = strconv.Itoa(i + 1)
					}

					// Keep track of latest number
					if j > index_of_latest_number {
						index_of_latest_number = j
						latest_number = strconv.Itoa(i + 1)
					}
				}
			}
		}
	}

	response, err := strconv.Atoi(earliest_number + latest_number)
	if err != nil {
		return -1, err
	}

	return response, nil
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

	// Figure out the numbers on each line, and add them together
	accum_sum := 0
	for _, line := range lines {
		runes := []rune(line)

		number, err := getNumber(runes)
		if err != nil {
			fmt.Println("Failed to find first number")
		}

		accum_sum += number
	}

	// Print the response
	fmt.Println(accum_sum)
}
