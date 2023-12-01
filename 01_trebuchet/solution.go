package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var NUMBERS = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {

	file, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(file), "\n")

	Part1(lines)
	Part2(lines)

}

func Part1(lines []string) {
	start_time := time.Now()
	total := 0
	for _, line := range lines {
		total += FindValue(string(line))
	}

	elapsed := time.Since(start_time)
	fmt.Printf("Part 1 answer: %v\n", total)
	fmt.Printf("Part 1 Time: %s\n\n", elapsed)
}

func Part2(lines []string) {
	start_time := time.Now()
	total := 0

	for _, line := range lines {
		value, err := strconv.Atoi(FindFirstDigit(line, false) + FindFirstDigit(line, true))
		if err != nil {
			log.Fatal(err)
		}
		total += value
	}

	elapsed := time.Since(start_time)
	fmt.Printf("Part 2 answer: %v\n", total)
	fmt.Printf("Part 2 Time: %s\n\n", elapsed)
}

func FindValue(s string) int {
	digit := '0'
	firstDigit := '0'

	for _, c := range s {
		if unicode.IsDigit(c) {
			digit = c
			if firstDigit == '0' {
				firstDigit = digit
			}
		}
	}
	ret, err := strconv.Atoi(string(firstDigit) + string(digit))
	if err != nil {
		log.Fatal(err)
	}
	return ret
}

func FindFirstDigit(s string, reverse bool) string {
	if reverse {
		s = Reverse(s)
	}
	for i, c := range s {
		if unicode.IsDigit(c) {
			return string(c)
		}
		for j, word := range NUMBERS {
			if reverse {
				word = Reverse(word)
			}

			if i+len(word) > len(s) {
				continue
			}
			if s[i:i+len(word)] == word {
				return fmt.Sprint(j + 1)
			}
		}
	}
	return ""
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
