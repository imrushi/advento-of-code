package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(fileContent)
	s := strings.Split(input, "\n")
	t, d := parseInt(strings.Join(strings.Fields(s[0])[1:], "")), parseInt(strings.Join(strings.Fields(s[1])[1:], ""))
	times := make([]int, 0)
	distances := make([]int, 0)
	times = append(times, t)
	distances = append(distances, d)
	// for i := 0; i < len(t); i++ {
	// 	times = append(times, parseInt(t[i]))
	// 	distances = append(distances, parseInt(d[i]))
	// }

	n := 1

	for i := 0; i < len(times); i++ {
		margin := 0
		for j := 0; j < times[i]; j++ {
			if j*(times[i]-j) > distances[i] {
				margin += 1
			}
		}
		n *= margin
	}

	fmt.Println(n)
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
