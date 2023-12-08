package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getPoints(data string) int {
	values := strings.Split(strings.Split(data, ":")[1], "|")
	winnerData := strings.Fields(values[0])
	numberData := strings.Fields(values[1])

	var winners []int
	for _, w := range winnerData {
		winners = append(winners, parseInt(w))
	}

	wDict := toDict(winners)

	var numbers []int
	for _, n := range numberData {
		numbers = append(numbers, parseInt(n))
	}

	var points int
	for _, x := range numbers {
		if wDict[x] {
			points++
		}
	}
	if points > 0 {
		points = 1 << (points - 1)
	}
	return points
}

func toDict(numbers []int) map[int]bool {
	dict := make(map[int]bool)
	for _, x := range numbers {
		dict[x] = true
	}
	return dict
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	fileContent, err := os.ReadFile("d4.txt")
	if err != nil {
		panic(err)
	}

	cards := strings.Split(string(fileContent), "\n")
	var points []int
	for _, v := range cards {
		points = append(points, getPoints(v))
	}
	fmt.Println(points)
	sum := 0
	for _, p := range points {
		sum += p
	}

	fmt.Println(sum)
}
