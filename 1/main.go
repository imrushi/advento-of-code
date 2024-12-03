package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Part 1: 2756096
// Part 2: 23117829

func main() {
	// 	testSlice := strings.Split(testInput, "\n")
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	dataString := string(data)

	var arr1, arr2 []int

	for i, v := range strings.Fields(dataString) {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			arr1 = append(arr1, val)
		} else {
			arr2 = append(arr2, val)
		}
	}
	part1(arr1, arr2)

	part2(arr1, arr2)
}

func part1(arr1 []int, arr2 []int) {
	sort.Ints(arr1)
	sort.Ints(arr2)

	sum := 0

	for i := range len(arr1) {
		temp := arr1[i] - arr2[i]
		sum += int(math.Abs(float64(temp)))
	}

	fmt.Println(sum)
}

type valOccurrs struct {
	val     int
	occurrs int
}

func part2(arr1, arr2 []int) {
	m := make(map[int]valOccurrs)

	for _, val := range arr1 {
		if d, ok := m[val]; ok {
			m[val] = valOccurrs{val: 0, occurrs: d.occurrs + 1}
		} else {
			m[val] = valOccurrs{0, 0}
		}
	}

	for _, v := range arr2 {
		if d, ok := m[v]; ok {
			m[v] = valOccurrs{d.val + 1, d.occurrs}
		}
	}

	prod := 0
	for key, v := range m {
		for range v.occurrs + 1 {
			prod += key * v.val
		}
	}
	fmt.Println(prod)
}
