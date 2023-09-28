package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	part1(scanner)
	part2(scanner)
}

func getCratePositions(lines []string) [][]byte {
	var crates [][]byte
	lastLine := lines[len(lines)-1]
	numCrates := len(strings.Split(lastLine, "   "))
	lines = lines[:len(lines)-1]

	for i := 0; i < numCrates; i++ {
		crates = append(crates, []byte{})
	}

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.ReplaceAll(lines[i], "     ", " [X] ")
		lines[i] = strings.ReplaceAll(lines[i], "    ", "[X] ")
		lines[i] = strings.ReplaceAll(lines[i], " ", "")
		lines[i] = strings.ReplaceAll(lines[i], "[", "")
		lines[i] = strings.ReplaceAll(lines[i], "]", "")
		// fmt.Println(lines[i])
	}

	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		for j := 0; j < numCrates; j++ {
			if line[j] == 'X' {
				continue
			}
			crates[j] = append(crates[j], line[j])
		}
	}

	return crates
}

type Instruction struct {
	num  int
	from int
	to   int
}

func getInstructions(lines []string) []Instruction {
	var instructions []Instruction

	for _, line := range lines {
		parts := strings.Split(line, " ")
		num, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])
		instructions = append(instructions, Instruction{num, from, to})
	}

	return instructions
}

func part1(scanner *bufio.Scanner) {
	fmt.Println("Part 1")
	var crates [][]byte
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			crates = getCratePositions(lines)
			lines = nil
		} else {
			lines = append(lines, line)
		}
	}

	instructions := getInstructions(lines)

	for _, c := range crates {
		fmt.Println(string(c))
	}
	fmt.Println()

	for _, inst := range instructions {
		toAdd := crates[inst.from-1][len(crates[inst.from-1])-inst.num:]
		// slices.Reverse(toAdd)
		crates[inst.to-1] = append(crates[inst.to-1], toAdd...)
		crates[inst.from-1] = crates[inst.from-1][:len(crates[inst.from-1])-inst.num]

		// for _, c := range crates {
		// 	fmt.Println(string(c))
		// }
	}

	fmt.Println()

	for _, c := range crates {
		fmt.Println(string(c))
	}
	fmt.Println()

	for _, crate := range crates {
		fmt.Print(string(crate[len(crate)-1]))
	}
	fmt.Println()
}

func part2(scanner *bufio.Scanner) {
	fmt.Println("Part 2")
}
