package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	indexes := map[string]int{
		"A": 0, // rock
		"X": 0, //lose
		"B": 1, // paper
		"Y": 1, //draw
		"C": 2, // scissors
		"Z": 2, //win
	}

	movePoints := []int{1, 2, 3}

	// 0 = rock
	// 1 = paper
	// 2 = scissors

	mat := [][]int{
		{2, 0, 1},
		{0, 1, 2},
		{1, 2, 0}}

	score := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		if len(parts) != 2 {
			log.Fatal("Invalid input")
			continue
		}

		oppMove := indexes[parts[0]]
		myMove := indexes[parts[1]]

		score += movePoints[mat[oppMove][myMove]]

		switch parts[1] {
		case "Y":
			score += 3
		case "Z":
			score += 6
		default:
			score += 0
		}
	}

	fmt.Println(score)
}
