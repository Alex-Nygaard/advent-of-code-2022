package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	instructions := readInstructions(lines)

	fmt.Println("Part 1:", part1(instructions))
	fmt.Println("Part 2:", part2(instructions))
}

type Instruction struct {
	op  string
	num int
}

func readInstructions(lines []string) []Instruction {
	instructions := make([]Instruction, 0)

	for _, line := range lines {
		var op string
		var num int

		fmt.Sscanf(line, "%s %d", &op, &num)
		instructions = append(instructions, Instruction{op: op, num: num})
	}

	return instructions
}

func part1(instructions []Instruction) int {
	cycle := 1
	x := 1
	sum := 0

	for _, instr := range instructions {
		if instr.op == "noop" {
			sum += addSignalStrength(cycle, x)
			cycle++
			// done
		} else if instr.op == "addx" {
			sum += addSignalStrength(cycle, x)
			cycle++
			sum += addSignalStrength(cycle, x)
			cycle++
			x += instr.num
		}
	}

	return sum
}

func addSignalStrength(cycle int, x int) int {
	if (cycle-20)%40 == 0 {
		return cycle * x
	}

	return 0
}

func part2(instructions []Instruction) int {
	cycle := 1
	x := 1

	for _, instr := range instructions {
		printPixel(cycle, x)

		if instr.op == "noop" {
			cycle++
			// done
		} else if instr.op == "addx" {
			cycle++
			printPixel(cycle, x)

			cycle++
			x += instr.num
		}
	}
	return 0
}

func printPixel(cycle int, sprite int) {
	pixel := (cycle - 1) % 40
	if absInt(pixel-sprite) <= 1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if cycle%40 == 0 {
		fmt.Println()
	}
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
