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

type Object struct {
	posX   int
	posY   int
	length int
}

type Number struct {
	Object
	value int
}

type Symbol struct {
	Object
	value string
}

func (a *Object) isAdjacent(b Object) bool {
	yDistance := a.posY - b.posY
	if math.Abs(float64(yDistance)) > 1 {
		return false
	}

	xMin := a.posX
	xMax := a.posX + a.length

	return b.posX <= xMax && (b.posX+b.length) >= xMin
}

func parseNumbersAndSybmolsFromLine(pos_y int, line string) ([]Number, []Symbol, error) {
	numbers := []Number{}
	symbols := []Symbol{}
	numBuffer := ""
	posXStart := 0
	numFound := false
	for posX, char := range line + "." {
		isDigit := unicode.IsDigit(char)
		isSymbol := !isDigit && string(char) != "."

		if isSymbol {
			symbols = append(symbols, Symbol{
				value: string(char),
				Object: Object{
					posX:   posX,
					posY:   pos_y,
					length: 1,
				},
			})
		}

		if isDigit {
			if !numFound {
				numFound = true
				posXStart = posX
			}
			numBuffer += string(char)
		}

		if !isDigit && numFound {
			value, err := strconv.Atoi(numBuffer)
			numBuffer = ""
			if err != nil {
				return numbers, symbols, err
			}
			numbers = append(numbers, Number{
				value: value,
				Object: Object{
					posX:   posXStart,
					posY:   pos_y,
					length: posX - posXStart,
				},
			})

			// Reset
			numBuffer = ""
			numFound = false
		}
	}

	return numbers, symbols, nil
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

	numbers := []Number{}
	symbols := []Symbol{}

	for pos_y, line := range lines {
		lineNumbers, lineSymbols, err := parseNumbersAndSybmolsFromLine(pos_y, line)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, lineNumbers...)
		symbols = append(symbols, lineSymbols...)
	}

	sumAccum := 0

	for _, symbol := range symbols {
		if symbol.value != "*" {
			continue
		}

		gearRatio := 1
		adjacentPartCount := 0
		for _, number := range numbers {
			if symbol.isAdjacent(number.Object) {
				adjacentPartCount += 1
				gearRatio *= number.value
			}
		}

		if adjacentPartCount == 2 {
			sumAccum += gearRatio
		}

	}

	fmt.Println(sumAccum)

}
