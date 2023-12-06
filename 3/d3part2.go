package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Pair struct {
	R int
	C int
}

func main() {
	fileContent, err := os.ReadFile("d3.txt")
	if err != nil {
		panic(err)
	}
	input := string(fileContent)

	grid := strings.Split(input, "\n")
	total := 0

	dirs := [][]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	// extracting digit adjacent to a symbol
	for r, row := range grid {
		for c, ch := range row {
			if ch != '*' {
				continue
			}
			cs := make(map[Pair]bool)
			// processing symbols all direction adjacency
			for _, dir := range dirs {
				dr, dc := r+dir[0], c+dir[1]
				if dr < 0 || dr >= len(grid) || dc < 0 || dc >= len(grid[dr]) || !unicode.IsDigit(rune(grid[dr][dc])) {
					continue
				}
				for dc > 0 && unicode.IsDigit(rune(grid[dr][dc-1])) {
					dc--
				}
				cs[Pair{dr, dc}] = true
			}

			if len(cs) != 2 {
				continue
			}

			var ns []int
			// from digit location extracting numbers
			for r := range cs {
				s := ""
				for r.C < len(grid[r.R]) && unicode.IsDigit(rune(grid[r.R][r.C])) {
					s += string(grid[r.R][r.C])
					r.C++
				}
				num, err := strconv.Atoi(s)
				if err == nil {
					ns = append(ns, num)
				}
			}

			total += ns[0] * ns[1]
		}
	}

	fmt.Println(total)
}
