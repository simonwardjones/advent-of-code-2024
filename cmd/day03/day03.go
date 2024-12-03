// Day 1 advent of code 2020

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	inputFileName := "input.txt"
	for _, arg := range os.Args {
		if arg == "--test" || arg == "test" {
			inputFileName = "input_test_1.txt"
			break
		} else if arg == "--test2" || arg == "test2" {
			inputFileName = "input_test_2.txt"
		}
	}

	data := loadInput(inputFileName)

	part1Result := part1(data)
	fmt.Println("Part 1:", part1Result)

	part2Result := part2(data)
	fmt.Println("Part 2:", part2Result)
}

func findProductsSum(data []string, allPairs bool) int {
	allData := strings.Join(data, "")
	// use regex
	var mulRegexp = regexp.MustCompile(`(?:mul\((\d+),(\d+)\))|don\'t\(\)|do`)
	matches := mulRegexp.FindAllStringSubmatch(allData, -1)
	var on bool = true
	var pairs [][]int
	for _, match := range matches {
		// fmt.Println(match)
		switch {
		case strings.HasPrefix(match[0], "don"):
			on = false
		case strings.HasPrefix(match[0], "do"):
			on = true
		default:
			if allPairs || on {
				var pair []int
				for _, matchVal := range match[1:] {
					value, _ := strconv.Atoi(matchVal)
					pair = append(pair, value)
				}
				pairs = append(pairs, pair)
			}
		}
	}
	// fmt.Println(pairs)
	var sum int
	for _, pair := range pairs {
		sum += pair[0] * pair[1]
	}
	return sum
}

func part1(data []string) int {
	return findProductsSum(data, true)
}
func part2(data []string) int {
	return findProductsSum(data, false)
}

func loadInput(fileName string) []string {
	_, currentFilePath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFilePath)
	inputFilePath := filepath.Join(currentDir, fileName)

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}
