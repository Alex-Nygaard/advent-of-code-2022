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

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		elf1 := strings.Split(parts[0], "-")
		elf2 := strings.Split(parts[1], "-")

		elf11, err := strconv.Atoi(elf1[0])
		elf12, err := strconv.Atoi(elf1[1])

		elf21, err := strconv.Atoi(elf2[0])
		elf22, err := strconv.Atoi(elf2[1])

		if err != nil {
			panic(err)
		}

		if (elf11 >= elf21 && elf11 <= elf22) || (elf12 >= elf21 && elf12 <= elf22) || (elf21 >= elf11 && elf21 <= elf12) || (elf22 >= elf11 && elf22 <= elf12) {
			count++
		}
	}

	fmt.Println(count)
}
