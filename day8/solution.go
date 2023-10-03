package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	forest := createForest(file)

	fmt.Println("Part 1 answer: ", part1(forest))
	fmt.Println("Part 2 UNFINISHED")
	fmt.Println("Part 2 answer (WRONG): ", part2(forest))
}

func createForest(file *os.File) *Forest {
	scanner := bufio.NewScanner(file)

	forest := NewForest()

	for scanner.Scan() {
		line := scanner.Text()

		var heights []int
		var visible []bool

		for _, r := range line {
			heights = append(heights, int(r-'0'))
			visible = append(visible, false)
		}

		forest.heights = append(forest.heights, heights)
		forest.visible = append(forest.visible, visible)
	}

	return forest
}

type Forest struct {
	heights [][]int
	visible [][]bool
}

func NewForest() *Forest {
	return &Forest{
		heights: make([][]int, 0),
		visible: make([][]bool, 0),
	}
}

func part1(forest *Forest) int {
	// rows
	for i := 0; i < len(forest.heights); i++ {
		left, right := -1, -1
		for j := 0; j < len(forest.heights[i]); j++ {
			var tree int

			tree = forest.heights[i][j]
			if tree > left {
				forest.visible[i][j] = true
				left = tree
			}

			tree = forest.heights[i][len(forest.heights[i])-1-j]
			if tree > right {
				forest.visible[i][len(forest.heights[i])-1-j] = true
				right = tree
			}
		}
	}

	// cols
	for j := 0; j < len(forest.heights[0]); j++ {
		top, bottom := -1, -1
		for i := 0; i < len(forest.heights); i++ {
			var tree int

			tree = forest.heights[i][j]
			if tree > top {
				forest.visible[i][j] = true
				top = tree
			}

			tree = forest.heights[len(forest.heights[j])-1-i][j]
			if tree > bottom {
				forest.visible[len(forest.heights[j])-1-i][j] = true
				bottom = tree
			}
		}
	}

	// count true
	count := 0
	for _, row := range forest.visible {
		for _, cell := range row {
			if cell {
				count++
			}
		}
	}

	return count
}

func part2(forest *Forest) int {
	maxScore := 0

	for i := 0; i < len(forest.heights); i++ {
		// fmt.Println("Col: ", i)
		for j := 0; j < len(forest.heights[i]); j++ {
			// fmt.Println("Row: ", j)
			up, down, left, right := 0, 0, 0, 0

			tallest := -1
			// up
			for cursor := i - 1; cursor >= 0 && tallest < forest.heights[i][j]; cursor-- {
				if forest.heights[cursor][j] >= tallest {
					up++
					tallest = forest.heights[cursor][j]
				}
			}

			tallest = -1
			// down
			for cursor := i + 1; cursor < len(forest.heights) && tallest < forest.heights[i][j]; cursor++ {
				if forest.heights[cursor][j] >= tallest {
					down++
					tallest = forest.heights[cursor][j]
				}
			}

			tallest = -1
			// left
			for cursor := j - 1; cursor >= 0 && tallest < forest.heights[i][j]; cursor-- {
				if forest.heights[i][cursor] >= tallest {
					left++
					tallest = forest.heights[i][cursor]
				}
			}

			tallest = -1
			// right
			for cursor := j + 1; cursor < len(forest.heights[i]) && tallest < forest.heights[i][j]; cursor++ {
				if forest.heights[i][cursor] >= tallest {
					right++
					tallest = forest.heights[i][cursor]
				}
			}

			if mult := up * down * left * right; mult > maxScore {
				maxScore = mult
			}
		}
	}

	return maxScore
}
