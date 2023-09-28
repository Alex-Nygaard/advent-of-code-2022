package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	fmt.Println("Result part 1: ", part1(input, 4))
	fmt.Println("Result part 2: ", part1(input, 14))
}

func part1(line string, uniqueCharsNeeded int) int {
	chars := []rune{}

	for i, nextChar := range line {
		isUnique := true
		for j, prevChar := range chars {
			if nextChar == prevChar {
				chars = chars[j+1:]
				isUnique = false
				break
			}
		}

		if isUnique && len(chars) == uniqueCharsNeeded-1 {
			return i + 1
		}

		chars = append(chars, nextChar)
	}

	return -1
}
