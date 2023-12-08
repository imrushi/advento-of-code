package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getPoints(data string, idx int) []int {
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

	var points int = 0
	for _, x := range numbers {
		if wDict[x] {
			points++
		}
	}
	// if points > 0 {
	// 	points = 1 << (points - 1)
	// }
	var newArray []int
	for i := 0; i < points; i++ {
		newArray = append(newArray, idx+1+i)
	}
	return newArray
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

	toProcess := make([]int, len(cards))
	for i := range toProcess {
		toProcess[i] = i + 1 // 1-based index
	}

	// Create maps to store processed data
	seen := make(map[int][]int)
	count := make(map[int]int)

	// for _, v := range cards {
	// 	points = append(points, getPoints(v))
	// }

	for len(toProcess) > 0 {
		idx := toProcess[len(toProcess)-1]
		toProcess = toProcess[:len(toProcess)-1]

		if c, ok := count[idx]; ok {
			count[idx] = c + 1
		} else {
			count[idx] = 1
		}

		var points []int

		// Check if the card has been processed before
		if p, ok := seen[idx]; ok {
			points = p
		} else {
			points = getPoints(cards[idx-1], idx) // Adjust for 0-based index
		}
		seen[idx] = points

		// fmt.Println("processing", idx, points)
		toProcess = append(toProcess, points...)
	}

	sum := 0
	for _, p := range count {
		sum += p
	}

	fmt.Println(sum)
}
