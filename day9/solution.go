package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	instructions := readInstructions(lines)

	fmt.Println("Part 1:", part1(instructions))
	fmt.Println("Part 2:", part2(instructions))
}

func readInstructions(lines []string) []Instruction {
	instructions := make([]Instruction, 0)

	for _, line := range lines {
		var op rune
		var steps int

		fmt.Sscanf(line, "%c %d", &op, &steps)

		instructions = append(instructions, Instruction{op: op, steps: steps})
	}

	return instructions
}

type Instruction struct {
	op    rune
	steps int
}

type Position struct {
	x int
	y int
}

type History struct {
	visited []Position
}

func NewHistory() *History {
	return &History{visited: make([]Position, 0)}
}

func (h *History) add(pos Position) bool {
	for _, p := range h.visited {
		if p == pos {
			return false
		}
	}

	h.visited = append(h.visited, pos)
	return true
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(instructions []Instruction) int {
	history := NewHistory()
	head := &Position{x: 0, y: 0}
	tail := &Position{x: 0, y: 0}
	history.add(*tail)

	for _, instr := range instructions {

		for k := 0; k < instr.steps; k++ {
			switch instr.op {
			case 'U':
				head.y++
			case 'D':
				head.y--
			case 'L':
				head.x--
			case 'R':
				head.x++
			}

			if absInt(head.x-tail.x) > 1 || absInt(head.y-tail.y) > 1 {
				if head.x-tail.x > 0 {
					tail.x++
				} else if head.x-tail.x < 0 {
					tail.x--
				}

				if head.y-tail.y > 0 {
					tail.y++
				} else if head.y-tail.y < 0 {
					tail.y--
				}
			}
			history.add(*tail)
		}
	}

	return len(history.visited)
}

func part2(instructions []Instruction) int {
	history := NewHistory()
	head := &Position{x: 0, y: 0}
	tail := &Position{x: 0, y: 0}

	body := make([]*Position, 0)
	body = append(body, head)
	for i := 0; i < 8; i++ {
		body = append(body, &Position{x: 0, y: 0})
	}
	body = append(body, tail)

	for _, instr := range instructions {
		for k := 0; k < instr.steps; k++ {
			switch instr.op {
			case 'U':
				head.y++
			case 'D':
				head.y--
			case 'L':
				head.x--
			case 'R':
				head.x++
			}

			for i, knot := range body[1:] {
				prev := body[i]
				if absInt(prev.x-knot.x) > 1 || absInt(prev.y-knot.y) > 1 {
					if prev.x-knot.x > 0 {
						knot.x++
					} else if prev.x-knot.x < 0 {
						knot.x--
					}

					if prev.y-knot.y > 0 {
						knot.y++
					} else if prev.y-knot.y < 0 {
						knot.y--
					}
				}
			}

			history.add(*tail)
		}
	}

	return len(history.visited)
}
