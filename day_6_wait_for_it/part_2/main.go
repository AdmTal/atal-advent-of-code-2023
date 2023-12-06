package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type RaceData struct {
	racetime     int
	bestDistance int
}

func (rd *RaceData) getNumWaysToWin() int {
	// Quadratic formula
	a := -float64(1)
	b := float64(rd.racetime)
	c := -float64(rd.bestDistance)

	xMin := int(math.Floor((-b+math.Sqrt((math.Pow(b, float64(2)))-(4*a*c)))/(2*a))) + 1
	xMax := int(math.Ceil((-b-math.Sqrt((math.Pow(b, float64(2)))-(4*a*c)))/(2*a))) - 1

	return (xMax - xMin + 1)
}

func parseInput(fileContents string) ([]RaceData, error) {
	racesData := []RaceData{}
	lines := strings.Split(fileContents, "\n")
	time := strings.Join(strings.Fields(lines[0])[1:], "")
	distance := strings.Join(strings.Fields(lines[1])[1:], "")

	raceTime, err := strconv.Atoi(time)
	if err != nil {
		return []RaceData{}, err
	}

	raceBestDistance, err := strconv.Atoi(distance)
	if err != nil {
		return []RaceData{}, err
	}

	raceData := RaceData{
		racetime:     raceTime,
		bestDistance: raceBestDistance,
	}
	racesData = append(racesData, raceData)

	return racesData, nil
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

	races, err := parseInput(fileContents)
	if err != nil {
		panic(err)
	}

	waysToWin := 1
	for _, race := range races {
		waysToWin *= race.getNumWaysToWin()
	}

	fmt.Println(waysToWin)
}
