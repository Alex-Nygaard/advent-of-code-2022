package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput()
	dir := createDirectoryStructure(input)
	log.Printf("Part 1: %d", part1(dir))
	log.Printf("Part 2: %d", part2(dir))
}

func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

type Directory struct {
	parent  *Directory
	files   map[string]int
	subdirs map[string]*Directory
	size    int
}

func NewDirectory(parent *Directory) *Directory {
	return &Directory{
		parent:  parent,
		files:   make(map[string]int),
		subdirs: make(map[string]*Directory),
		size:    0,
	}
}

func (d *Directory) addSubdir(name string) *Directory {
	if _, ok := d.subdirs[name]; !ok {
		d.subdirs[name] = NewDirectory(d)
	}
	return d.subdirs[name]
}

func (d *Directory) calcSizeRec() int {
	size := 0

	for _, fileSize := range d.files {
		size += fileSize
	}

	for _, subDir := range d.subdirs {
		size += subDir.calcSizeRec()
	}

	d.size = size
	return size
}

func (d *Directory) sumDirsUnder100k() int {
	sum := 0

	for _, subdir := range d.subdirs {
		sum += subdir.sumDirsUnder100k()
	}

	if d.size <= 100_000 {
		sum += d.size
	}

	return sum
}

func (d *Directory) findDirsAboveThreshold(threshold int) []int {
	var matchingDirs []int

	for _, subdir := range d.subdirs {
		matchingDirs = append(matchingDirs, subdir.findDirsAboveThreshold(threshold)...)
	}

	if d.size >= threshold {
		matchingDirs = append(matchingDirs, d.size)
	}

	return matchingDirs
}

func createDirectoryStructure(input []string) *Directory {
	dir := NewDirectory(nil)
	currentDir := dir

	for _, line := range input {
		parts := strings.Split(line, " ")
		if strings.HasPrefix(parts[0], "$") {
			// command
			if parts[1] == "ls" {
				continue
			}
			switch parts[2] {
			case "/":
				currentDir = dir
			case "..":
				currentDir = currentDir.parent
			default:
				currentDir = currentDir.addSubdir(parts[2])
			}
		} else if strings.HasPrefix(parts[0], "dir") {
			currentDir.addSubdir(parts[1])
		} else {
			// file
			fileSize, _ := strconv.Atoi(parts[0])
			fileName := parts[1]
			currentDir.files[fileName] = fileSize
		}
	}

	dir.calcSizeRec()

	return dir
}

func part1(dir *Directory) int {
	return dir.sumDirsUnder100k()
}

func part2(dir *Directory) int {
	totalSpace := 70_000_000
	neededSpace := 30_000_000
	freeSpace := totalSpace - dir.size
	missingSpace := neededSpace - freeSpace

	suitableDirs := dir.findDirsAboveThreshold(missingSpace)
	minValue := suitableDirs[0]

	for _, size := range suitableDirs {
		if size < minValue {
			minValue = size
		}
	}

	return minValue
}
