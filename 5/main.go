package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	seedsRanges := [][]int{}
	var seedToSoil [][]int
	var soilToFertilizer [][]int
	var fertilizerToWater [][]int
	var waterToLight [][]int
	var lightToTemperature [][]int
	var temperatureToHumidity [][]int
	var humidityToLocation [][]int

	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error reading file: ", err)
	}

	input := string(fileContent)

	sections := strings.Split(input, "\n\n")

	seedsLine := strings.Split(sections[0], " ")[1:]
	for i := 0; i < len(seedsLine); i += 2 {
		start, err := strconv.Atoi(seedsLine[i])
		if err != nil {
			fmt.Println("error converting string to int: ", err)
		}
		length, err := strconv.Atoi(seedsLine[i+1])
		if err != nil {
			fmt.Println("error converting string to int: ", err)
		}
		seedRange := generateSeedRange(start, length)
		seedsRanges = append(seedsRanges, seedRange)
	}

	// for _, v := range strings.Split(sections[0], " ")[1:] {
	// 	n, err := strconv.Atoi(strings.TrimSpace(v))
	// 	if err != nil {
	// 		fmt.Println("error converting string to int: ", err)
	// 	}
	// 	seeds = append(seeds, n)
	// }

	for _, section := range sections[1:] {
		if strings.HasPrefix(section, "seed-to-soil map:") {
			seedToSoil = parseMapSection(section)
		} else if strings.HasPrefix(section, "soil-to-fertilizer map:") {
			soilToFertilizer = parseMapSection(section)
		} else if strings.HasPrefix(section, "fertilizer-to-water map:") {
			fertilizerToWater = parseMapSection(section)
		} else if strings.HasPrefix(section, "water-to-light map:") {
			waterToLight = parseMapSection(section)
		} else if strings.HasPrefix(section, "light-to-temperature map:") {
			lightToTemperature = parseMapSection(section)
		} else if strings.HasPrefix(section, "temperature-to-humidity map:") {
			temperatureToHumidity = parseMapSection(section)
		} else if strings.HasPrefix(section, "humidity-to-location map:") {
			humidityToLocation = parseMapSection(section)
		}
	}
	// fmt.Printf("\n%#v\n%v", sections, seed, )
	// fmt.Println(seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation)

	// lowest
	lowest := -1
	for _, seeds := range seedsRanges {
		for _, s := range seeds {
			soil := processMap(seedToSoil, s)
			// fmt.Printf("\n[%v] soil: %v", s, soil)

			fertilizer := processMap(soilToFertilizer, soil)
			// fmt.Printf("\n[%v] soilToFertilizer: %v", s, fertilizer)

			water := processMap(fertilizerToWater, fertilizer)
			// fmt.Printf("\n[%v] fertilizerToWater: %v", s, water)

			light := processMap(waterToLight, water)
			// fmt.Printf("\n[%v] waterToLight: %v", s, light)

			temperature := processMap(lightToTemperature, light)
			// fmt.Printf("\n[%v] lightToTemperature: %v", s, temperature)

			humidity := processMap(temperatureToHumidity, temperature)
			// fmt.Printf("\n[%v] temperatureToHumidity: %v\t %v", s, humidity, temperatureToHumidity)

			location := processMap(humidityToLocation, humidity)
			// fmt.Printf("\n[%v] humidityToLocation: %v", s, location)
			if lowest == -1 || location < lowest {
				lowest = location
			}
			// fmt.Println("\n---------------------------------------------")
		}
	}

	fmt.Println("Lowest Location: ", lowest)
}

func generateSeedRange(start, length int) []int {
	seedRange := make([]int, length)
	for i := 0; i < length; i++ {
		seedRange[i] = start + i
	}
	return seedRange
}

func diff(req, src, dst, r int) int {
	// fmt.Printf("diff:  src: %v, req: %v, src+r-1: %v", src, req, src+r-1)
	if (req >= src) && (req <= src+r-1) {
		return dst + (req - src)
	}
	return -1
}

func parseMapSection(section string) [][]int {
	lines := strings.Split(section, "\n")
	var result [][]int

	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		if len(fields) > 0 {
			var row []int
			for _, field := range fields {
				value, err := strconv.Atoi(field)
				if err != nil {
					fmt.Println("error converting to integer: ", err)
					return nil
				}
				row = append(row, value)
			}
			result = append(result, row)
		}
	}
	return result
}

func processMap(mappings [][]int, seed int) int {
	// result := make(map[int]int)
	// var res int
	// res = seed
	for _, mapping := range mappings {
		destStart, sourceStart, length := mapping[0], mapping[1], mapping[2]

		temp := diff(seed, sourceStart, destStart, length)
		if temp != -1 {
			return temp
		}
	}

	return seed
}
