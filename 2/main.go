package main

import (
	"advent/common"
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	inputFileName = flag.String("input-file", "input.txt", "this will be used as input")
)

func main() {
	// 	testSlice := strings.Split(testInput, "\n")
	flag.Parse()
	inputData := common.ReadFile(*inputFileName)

	lines := strings.Split(string(inputData), "\n")

	var lineArr [][]int
	for _, line := range lines {
		n := strings.Split(line, " ")
		var l []int
		for _, no := range n {
			data, err := strconv.Atoi(no)
			if err != nil {
				panic(data)
			}
			l = append(l, data)
		}
		lineArr = append(lineArr, l)
	}

	// fmt.Println(lineArr)
	part1(lineArr)

	part2(lineArr)
}

func part1(lineArr [][]int) {
	var count int
	for _, i := range lineArr {
		if validate(i) {
			count++
		}
	}

	fmt.Println(count)
}

func part2(lineArr [][]int) {
	var count int
	for _, s := range lineArr {
		if validate(s) || validateByRemoving(s) {
			count++
		}
	}

	fmt.Println(count)
}

func validate(data []int) bool {
	return (isIncreasing(data) || isDecreasing(data)) && maxDiff(data, 3) && minDiff(data, 1)
}

func validateByRemoving(data []int) bool {
	for i := 0; i < len(data); i++ {
		data1 := append([]int{}, data[:i]...)
		data1 = append(data1, data[i+1:]...)
		if (isIncreasing(data1) || isDecreasing(data1)) && maxDiff(data1, 3) && minDiff(data1, 1) {
			return true
		}
	}
	return false
}

func isIncreasing(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func isDecreasing(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			return false
		}
	}
	return true
}

func maxDiff(levels []int, v int) bool {
	for i := 1; i < len(levels); i++ {
		diff := int(math.Abs(float64(levels[i] - levels[i-1])))
		if diff > v {
			return false
		}
	}
	return true
}

func minDiff(levels []int, v int) bool {
	for i := 1; i < len(levels); i++ {
		diff := int(math.Abs(float64(levels[i] - levels[i-1])))
		if diff < v {
			return false
		}
	}
	return true
}

// func part2(arr1, arr2 []int) {
// 	m := make(map[int]valOccurrs)

// 	for _, val := range arr1 {
// 		if d, ok := m[val]; ok {
// 			m[val] = valOccurrs{val: 0, occurrs: d.occurrs + 1}
// 		} else {
// 			m[val] = valOccurrs{0, 0}
// 		}
// 	}

// 	for _, v := range arr2 {
// 		if d, ok := m[v]; ok {
// 			m[v] = valOccurrs{d.val + 1, d.occurrs}
// 		}
// 	}

// 	prod := 0
// 	for key, v := range m {
// 		for range v.occurrs + 1 {
// 			prod += key * v.val
// 		}
// 	}
// 	fmt.Println(prod)
// }
