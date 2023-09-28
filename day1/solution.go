package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func day1() {
	// read data

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	elves := []int{0, 0, 0}
	currentCals := 0

	for scanner.Scan() {
		line := scanner.Text()
		val, err := strconv.Atoi(line)
		if err != nil {
			minCalInd := argMin(elves)
			if elves[minCalInd] < currentCals {
				elves[minCalInd] = currentCals
			}
			currentCals = 0
		} else {
			currentCals += val
		}
	}

	fmt.Println(elves[0] + elves[1] + elves[2])

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
