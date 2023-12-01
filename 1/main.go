package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var letters map[string]string

// Part 1: 55621
// Part 2: 53592

func main() {
	letters = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"eno":   "1",
		"owt":   "2",
		"eerht": "3",
		"ruof":  "4",
		"evif":  "5",
		"xis":   "6",
		"neves": "7",
		"thgie": "8",
		"enin":  "9",
	}
	//	testInput := `1abc2
	//
	// pqr3stu8vwx
	// a1b2c3d4e5f
	// treb7uchet`
	// 	testInput := `two1nine
	// eightwothree
	// abcone2threexyz
	// xtwone3four
	// 4nineeightseven2
	// zoneight234
	// 7pqrstsixteen`

	// 	testSlice := strings.Split(testInput, "\n")
	testSlice := make([]string, 0)
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		testSlice = append(testSlice, scanner.Text())
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	// sliceOfTwoDigitNum := make([]int, 0)
	sum := 0
	for _, d := range testSlice {
		// sum += extractFirstLastDigit(d)
		sum += extractDigitFromLetter(d)
	}
	fmt.Println(sum)
}

func extractDigitFromLetter(s string) int {
	re_front := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
	re_back := regexp.MustCompile(`(\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)`)
	front_matches := re_front.FindAllString(s, 1)
	back_mathches := re_back.FindAllString(reverse(s), 1)
	if len(front_matches) != 0 {
		twoDigits := checkInMap(front_matches[0]) + checkInMap(back_mathches[0])
		digits, err := strconv.ParseInt(twoDigits, 10, 0)
		if err != nil {
			fmt.Printf("Error parsing into int: %v", err)
		}
		return int(digits)
	}
	return 0
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func checkInMap(s string) string {
	if val, ok := letters[s]; ok {
		return val
	}
	return s
}

func extractFirstLastDigit(s string) int {
	re := regexp.MustCompile(`\d`)
	match := re.FindAllString(s, -1)
	if len(match) != 0 {
		digits, err := strconv.ParseInt(match[0]+match[len(match)-1], 10, 0)
		if err != nil {
			fmt.Printf("Error parsing into int: %v", err)
		}
		return int(digits)
	}
	return 0
}
