package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type CategoryMap struct {
	source      int
	destination int
	length      int
}

type Category struct {
	categoryMappings []CategoryMap
}

type Categories struct {
	seedToSoil          Category
	soilToFertilizer    Category
	fertilizerToWater   Category
	waterToLight        Category
	lightToTemperature  Category
	tempatureToHumidity Category
	humidityToLocation  Category
}

// func (c *Category) getDestinationForSource(sourceId int) int {
// 	// todo
// }

// func (c *Categories) getLocationForSeed(seedId int) int {
// 	// todo
// }

func getCategoryForLines(lines []string) Category {
	// Pop Header
	_, lines = lines[0], lines[1:]
	mappings := []CategoryMap{}
	for _, line := range lines {
		if line == "" {
			break
		}
		map_values := strings.Fields(line)

		source, err := strconv.Atoi(map_values[0])
		if err != nil {
			panic(err)
		}

		destination, err := strconv.Atoi(map_values[1])
		if err != nil {
			panic(err)
		}

		length, err := strconv.Atoi(map_values[2])
		if err != nil {
			panic(err)
		}

		categoryMap := CategoryMap{
			source:      source,
			destination: destination,
			length:      length,
		}
		mappings = append(mappings, categoryMap)
	}
	return Category{categoryMappings: mappings}
}

func parseInput(fileContents string) ([]int, Categories, error) {
	lines := strings.Split(fileContents, "\n")

	first_line, lines := lines[0], lines[1:]

	seedsStrings := strings.Fields(strings.Split(first_line, ": ")[1])
	seeds := []int{}
	for _, seedString := range seedsStrings {
		seed, err := strconv.Atoi(seedString)
		if err != nil {
			return seeds, Categories{}, err
		}
		seeds = append(seeds, seed)
	}

	// Pop empty line
	_, lines = lines[0], lines[1:]

	seedsToSoilMappingStrings := []string{}
	for _, line := range lines {
		if line == "" {
			break
		}
		seedsToSoilMappingStrings = append(seedsToSoilMappingStrings, line)
	}
	soilToFertilizerMappingStrings := []string{}
	for _, line := range lines {
		if line == "" {
			break
		}
		soilToFertilizerMappingStrings = append(soilToFertilizerMappingStrings, line)
	}
	fertilizerToWaterMappingStrings := []string{}
	for _, line := range lines {
		if line == "" {
			break
		}
		fertilizerToWaterMappingStrings = append(fertilizerToWaterMappingStrings, line)
	}
	waterToLightMappingStrings := []string{}
	for _, line := range lines {
		if line == "" {
			break
		}
		waterToLightMappingStrings = append(waterToLightMappingStrings, line)
	}
	lightToTemperatureMappingStrings := []string{}
	for _, line := range lines {
		if line == "" {
			break
		}
		lightToTemperatureMappingStrings = append(lightToTemperatureMappingStrings, line)
	}
	temperatureToHumidityMappingStrings := []string{}
	for _, line := range lines {
		if line == "" {
			break
		}
		temperatureToHumidityMappingStrings = append(temperatureToHumidityMappingStrings, line)
	}
	humidityToLocationMappingStrings := []string{}
	for _, line := range lines {
		if line == "" {
			break
		}
		humidityToLocationMappingStrings = append(humidityToLocationMappingStrings, line)
	}

	categories := Categories{
		seedToSoil:          getCategoryForLines(seedsToSoilMappingStrings),
		soilToFertilizer:    getCategoryForLines(soilToFertilizerMappingStrings),
		fertilizerToWater:   getCategoryForLines(fertilizerToWaterMappingStrings),
		waterToLight:        getCategoryForLines(waterToLightMappingStrings),
		lightToTemperature:  getCategoryForLines(lightToTemperatureMappingStrings),
		tempatureToHumidity: getCategoryForLines(temperatureToHumidityMappingStrings),
		humidityToLocation:  getCategoryForLines(humidityToLocationMappingStrings),
	}

	return seeds, categories, nil
}

func main() {

	// Figure out where the current file is
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable info:", err)
	}
	exeDir := filepath.Dir(exePath)

	// Load the input file from the same dir
	filePath := filepath.Join(exeDir, "example_input.txt")
	// filePath := filepath.Join(exeDir, "input.txt")

	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	seeds, categories, err := parseInput(string(content))

	fmt.Println(seeds, categories)

}
