package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// todo

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	score := 0

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		if len(lines) == 3 {
			score += processLines(lines)
			lines = nil
		}
	}

	fmt.Println(score)
}

func processLines(lines []string) int {
	ruck1 := map[rune]bool{}
	ruck2 := map[rune]bool{}

	for _, c := range lines[0] {
		ruck1[c] = true
	}

	for _, c := range lines[1] {
		ruck2[c] = true
	}

	var common rune

	for _, c := range lines[2] {
		if ruck1[c] && ruck2[c] {
			common = c
			break
		}
	}

	if unicode.IsUpper(common) {
		return int(common - rune('A') + 27)
	} else {
		return int(common - rune('a') + 1)
	}
}
