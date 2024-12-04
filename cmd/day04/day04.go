// Day 1 advent of code 2020

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
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

var Directions = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

var NextLetter = map[rune]rune{'X': 'M', 'M': 'A', 'A': 'S'}

func nextLetter(char rune) rune {
	if next, ok := NextLetter[char]; ok {
		return next
	}
	return char
}

func inGrid(x, y, N, M int) bool {
	return x >= 0 && x < N && y >= 0 && y < M
}

func Walk(data []string, x, y int, direction [2]int, M, N int) func(func(rune) bool) {
	return func(yield func(rune) bool) {
		for {
			x, y = x+direction[0], y+direction[1]
			if !inGrid(x, y, N, M) {
				return
			}
			if !yield(rune(data[x][y])) {
				return
			}
		}
	}

}

func part1(data []string) int {
	found := 0
	N, M := len(data), len(data[0]) // N rows, M columns
	for row, line := range data {
		for col, char := range line {
			if char == 'X' {
				for _, dir := range Directions {
					currentChar := char
					for nextRune := range Walk(data, row, col, dir, M, N) {
						if nextRune != nextLetter(currentChar) {
							break
						}
						if nextRune == 'S' {
							found++
							break
						}
						currentChar = nextRune
					}
				}
			}
		}
	}
	return found
}

func checkForXMask(data []string, row, col int, M, N int) bool {
	if data[row][col] != 'A' {
		return false
	}
	if row < 1 || row > N-2 || col < 1 || col > M-2 {
		return false
	}
	TL := []int{row - 1, col - 1}
	TR := []int{row - 1, col + 1}
	BL := []int{row + 1, col - 1}
	BR := []int{row + 1, col + 1}
	// create a set with data from tl and br
	var expected = map[rune]bool{'M': true, 'S': true}
	var tl_br = map[rune]bool{
		rune(data[TL[0]][TL[1]]): true,
		rune(data[BR[0]][BR[1]]): true,
	}
	var tr_bl = map[rune]bool{
		rune(data[TR[0]][TR[1]]): true,
		rune(data[BL[0]][BL[1]]): true,
	}
	if reflect.DeepEqual(expected, tl_br) && reflect.DeepEqual(expected, tr_bl) {
		return true
	}
	return false
}

func part2(data []string) int {
	found := 0
	N, M := len(data), len(data[0]) // N rows, M columns
	for row, line := range data {
		for col := range line {
			if checkForXMask(data, row, col, M, N) {
				found++
			}
		}
	}
	return found
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
