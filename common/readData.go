package common

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile(inputFileName string) []byte {
	// read input.txt
	input, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	return input
}

func ReadString(in *bufio.Scanner) string {
	nStr := in.Text()
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	nStr = strings.TrimSpace(nStr)
	nStr = strings.Trim(nStr, "\t \n")
	return nStr
}

func ReadStrArr(in *bufio.Scanner) []string {
	line := ReadString(in)
	numbs := strings.Split(line, " ")
	return numbs
}

func ReadArrInt(in *bufio.Scanner) []int {
	numbs := ReadStrArr(in)
	arr := make([]int, 0)
	for _, n := range numbs {
		val, _ := strconv.Atoi(n)
		arr = append(arr, val)
	}
	return arr
}
