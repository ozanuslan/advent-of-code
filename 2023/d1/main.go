package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func main() {
	lines := readInput()
	p1 := part1(lines)
	p2 := part2(lines)

	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
}

func part1(lines []string) int {
	lineValues := make([]int, 0)

	for _, line := range lines {
		str := ""
		for _, char := range line {
			if unicode.IsDigit(char) {
				str += string(char)
			}
		}

		var firstDigit byte
		var lastDigit byte
		if len(str) == 1 {
			firstDigit = str[0]
			lastDigit = str[0]
		} else {
			firstDigit = str[0]
			lastDigit = str[len(str)-1]
		}
		numStr := string(firstDigit) + string(lastDigit)
		num, _ := strconv.Atoi(numStr)

		lineValues = append(lineValues, num)
	}

	sum := 0
	for _, lineValue := range lineValues {
		sum += lineValue
	}

	return sum
}

func part2(lines []string) int {
	lineValues := make([]int, 0)

	pattern := `\d|(one|two|three|four|five|six|seven|eight|nine)`
	for _, line := range lines {
		matches := findAllOverlappingMatches(line, pattern)

		str := ""
		for _, match := range matches {
			digitChr, err := strToDigitChar(match)
			if err == nil {
				str += string(digitChr)
			}
		}

		var firstDigit byte
		var lastDigit byte
		if len(str) == 1 {
			firstDigit = str[0]
			lastDigit = str[0]
		} else {
			firstDigit = str[0]
			lastDigit = str[len(str)-1]
		}
		numStr := string(firstDigit) + string(lastDigit)
		num, _ := strconv.Atoi(numStr)

		lineValues = append(lineValues, num)
	}

	sum := 0
	for _, lineValue := range lineValues {
		sum += lineValue
	}

	return sum
}

func findAllOverlappingMatches(input string, pattern string) []string {
	var matches []string
	regex := regexp.MustCompile(pattern)

	for pos := 0; pos < len(input); {
		loc := regex.FindStringIndex(input[pos:])
		if loc == nil {
			break
		}

		start, end := loc[0]+pos, loc[1]+pos
		matches = append(matches, input[start:end])

		pos = start + 1
	}

	return matches
}

func strToDigitChar(str string) (byte, error) {
	if len(str) == 1 && unicode.IsDigit(rune(str[0])) {
		return str[0], nil
	}

	switch str {
	case "one":
		return '1', nil
	case "two":
		return '2', nil
	case "three":
		return '3', nil
	case "four":
		return '4', nil
	case "five":
		return '5', nil
	case "six":
		return '6', nil
	case "seven":
		return '7', nil
	case "eight":
		return '8', nil
	case "nine":
		return '9', nil
	}

	return ' ', errors.New("error: digit not recognized")
}

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
