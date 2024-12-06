// Day 1 advent of code 2020

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
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
	// fmt.Println(data)

	part1Result := part1(data)
	fmt.Println("Part 1:", part1Result)

	part2Result := part2(data)
	fmt.Println("Part 2:", part2Result)

}

var (
	Up    = [2]int{-1, 0}
	Right = [2]int{0, 1}
	Down  = [2]int{1, 0}
	Left  = [2]int{0, -1}
)

func turnRight(direction [2]int) [2]int {
	switch direction {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	default:
		panic("invalid direction")
	}
}

func findInitialPosAndDirection(data []string) (row, col int, direction [2]int) {
	for i, line := range data {
		for j, char := range line {
			switch char {
			case '^':
				return i, j, Up
			case '>':
				return i, j, Right
			case 'v':
				return i, j, Down
			case '<':
				return i, j, Left
			}
		}
	}
	log.Panicf("initial position not found")
	return
}

func getChar(direction [2]int) rune {
	switch direction {
	case Up:
		return '^'
	case Right:
		return '>'
	case Down:
		return 'v'
	case Left:
		return '<'
	default:
		return 'o'
	}
}

func printBoard(data []string, row, col int, direction [2]int) {
    fmt.Print("\033[H\033[2J") // Clear the terminal
    for i, line := range data {
        for j, char := range line {
            if i == row && j == col {
                fmt.Print(string(getChar(direction)))
            } else {
                fmt.Print(string(char))
            }
        }
        fmt.Println()
    }
    fmt.Println()
    fmt.Println()
    time.Sleep(100 * time.Millisecond) // Sleep for 100 milliseconds
}

func findVisited(data []string) map[[2]int]bool {
	row, col, initialDirection := findInitialPosAndDirection(data)
	direction := initialDirection
	n, m := len(data), len(data[0]) // rows, cols
	visited := make(map[[2]int]bool)
	visited[[2]int{row, col}] = true
	for {
        // printBoard(data, row, col, direction)
		step_row, step_col := row+direction[0], col+direction[1]
		if step_row < 0 || step_row >= n || step_col < 0 || step_col >= m {
			// fmt.Println("Stepped off the board")
			break
		}
		step_found := rune(data[step_row][step_col])
		if step_found == '.' || step_found == getChar(initialDirection) {
			row, col = step_row, step_col
			visited[[2]int{row, col}] = true
		} else if step_found == '#' {
			direction = turnRight(direction)
		} else {
			log.Panicf("invalid character")
		}
	}
	return visited
}

func part1(data []string) int {
	return len(findVisited(data))
}

func part2(data []string) int {
	row, col, initialDirection := findInitialPosAndDirection(data)
	visited := findVisited(data)

	var loops int
	for swap := range visited {
		if isLoop(data, row, col, initialDirection, swap[0], swap[1]) {
			loops++
		}
	}

	return loops
}

func isLoop(data []string, row, col int, initialDirection [2]int, modRow, modCol int) bool {
	direction := initialDirection
	n, m := len(data), len(data[0]) // rows, cols
	visited := make(map[[4]int]bool)
	visited[[4]int{row, col, direction[0], direction[1]}] = true

	for {
		step_row, step_col := row+direction[0], col+direction[1]
		if step_row < 0 || step_row >= n || step_col < 0 || step_col >= m {
			return false
		}
		step_found := rune(data[step_row][step_col])
		if step_row == modRow && step_col == modCol {
			step_found = '#'
		}
		if step_found == '.' || step_found == getChar(initialDirection) {
			row, col = step_row, step_col
		} else if step_found == '#' {
			direction = turnRight(direction)
		} else {
			log.Panicf("invalid character")
		}
		if visited[[4]int{row, col, direction[0], direction[1]}] {
			return true
		}
		visited[[4]int{row, col, direction[0], direction[1]}] = true
	}
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
