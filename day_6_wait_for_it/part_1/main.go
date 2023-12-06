package main

import (
	"errors"
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

func (rd *RaceData) getLowestNumberThatBeatsBestTime(time int) (int, error) {
	timeRange := 2
	for i := time - timeRange; i <= time+timeRange; i++ {
		result := rd.getDistance(i)
		if result > rd.bestDistance {
			return i, nil
		}
	}
	return 0, errors.New("could not find lowest number")
}

func (rd *RaceData) getHighestNumberThatBeatsBestTime(time int) (int, error) {
	timeRange := 2
	for i := time + timeRange; i > time-timeRange; i-- {
		result := rd.getDistance(i)
		if result > rd.bestDistance {
			return i, nil
		}
	}
	return 0, errors.New("could not find highest number")
}

func (rd *RaceData) getDistance(chargeTime int) int {
	return chargeTime * (rd.racetime - chargeTime)
}

func (rd *RaceData) getNumWaysToWin() int {
	// Quadratic formula
	a := -float64(1)
	b := float64(rd.racetime)
	c := -float64(rd.bestDistance)

	xMin := int(math.Floor((-b + math.Sqrt((math.Pow(b, float64(2)))-(4*a*c))) / (2 * a)))
	xMax := int(math.Ceil((-b - math.Sqrt((math.Pow(b, float64(2)))-(4*a*c))) / (2 * a)))

	x0, err := rd.getLowestNumberThatBeatsBestTime(xMin)
	if err != nil {
		panic(err)
	}
	x1, err := rd.getHighestNumberThatBeatsBestTime(xMax)
	if err != nil {
		panic(err)
	}

	return (x1 - x0 + 1)
}

func parseInput(fileContents string) ([]RaceData, error) {
	racesData := []RaceData{}
	lines := strings.Split(fileContents, "\n")
	times := strings.Fields(lines[0])
	distances := strings.Fields(lines[1])

	for i := 1; i < len(times); i++ {
		raceTime, err := strconv.Atoi(times[i])
		if err != nil {
			return []RaceData{}, err
		}

		raceBestDistance, err := strconv.Atoi(distances[i])
		if err != nil {
			return []RaceData{}, err
		}

		raceData := RaceData{
			racetime:     raceTime,
			bestDistance: raceBestDistance,
		}
		racesData = append(racesData, raceData)
	}

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
