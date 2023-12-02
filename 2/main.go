package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	max := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	contents := make([]int, 0)

	// Variable to store the sum of power of minimum sets
	sumPower := 0

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		contents = solution1(line, max, contents)
		sumPower += solution2(line)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	// for solution 1
	sum := 0
	for _, gameId := range contents {
		sum += gameId
	}
	//
	fmt.Println(sum, sumPower)
}

func solution1(line string, max map[string]int, contents []int) []int {
	gameID, err := strconv.Atoi(strings.Fields(strings.Split(line, ":")[0])[1])
	if err != nil {
		fmt.Println("Error converting game ID to integer:", err)
	}

	validLines := strings.Split(strings.Split(line, ":")[1], ";")
	valid := true
	for _, game := range validLines {
		dices := strings.Split(strings.TrimSpace(game), ",")
		for _, dice := range dices {
			parts := strings.Fields(strings.TrimSpace(dice))
			color := parts[1]
			count, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Error converting count to integer:", err)
			}
			if count > max[color] {
				valid = false
				break
			}
		}
	}

	if valid {
		contents = append(contents, gameID)
	}
	return contents
}

func solution2(line string) int {
	colorCount := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	validLines := strings.Split(strings.Split(line, ":")[1], ";")
	for _, game := range validLines {
		dices := strings.Split(strings.TrimSpace(game), ",")
		for _, dice := range dices {
			parts := strings.Fields(strings.TrimSpace(dice))
			color := parts[1]
			count, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Error converting count to integer:", err)
			}
			colorCount[color] = max(colorCount[color], count)
		}
	}

	power := colorCount["red"] * colorCount["green"] * colorCount["blue"]

	return power
}
