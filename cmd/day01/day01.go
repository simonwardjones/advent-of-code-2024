// Day 1 advent of code 2020

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)


func main() {
	inputFileName := "input.txt"
	data := loadInput(inputFileName)

	part1Result := part1(data)
	fmt.Println(part1Result)

	part2Result := part2(data)
	fmt.Println(part2Result)
}

func part1(data []string) int {
	var listA, listB []int
	for _, line := range data {
		parts := strings.Split(line, "   ")
		// convert part[0] and part[1] to int
		a, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("failed to convert string to int: %v", err)
		}
		b, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("failed to convert string to int: %v", err)
		}
		// append a and b to listA and listB
		listA = append(listA, a)
		listB = append(listB, b)
	}
	// fmt.Println(listA)
	// fmt.Println(listB)

	// sort both lists
	sort.Ints(listA)
	sort.Ints(listB)
	var sum int
	var a, b int
	for i := 0; i < len(listA); i++ {
		a, b = listA[i], listB[i]
		if a > b {
			sum += a - b
		} else {
			sum += b - a
		}
	}
	return sum
}

func part2(data []string) int {

	var listA, listB []int
	for _, line := range data {
		parts := strings.Split(line, "   ")
		// convert part[0] and part[1] to int
		a, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("failed to convert string to int: %v", err)
		}
		b, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("failed to convert string to int: %v", err)
		}
		// append a and b to listA and listB
		listA = append(listA, a)
		listB = append(listB, b)
	}
	// fmt.Println(listA)
	// fmt.Println(listB)

	// condense b in to map of counts
	bCounts := make(map[int]int)
	for _, b := range listB {
		bCounts[b]++
	}

	// sort both lists
	aCountsInB := make(map[int]int)
	for _, a := range listA {
		aCountsInB[a] += bCounts[a]
	}
	// fmt.Println(aCountsInB)

	var sum int
	for a, count := range aCountsInB {
		sum += a * count
	}
	return sum
}

func loadInput(fileName string) []string {
	// Read the input file and return lines of strings
	_, currentFilePath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFilePath)

	// Construct the path to the input file relative to the current file
	inputFilePath := filepath.Join(currentDir, fileName)

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		data = append(data, line)
	}
	return data
}
