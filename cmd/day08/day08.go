// Day 1 advent of code 2020

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	// fmt.Println(data)

	part1Result := part1(data)
	fmt.Println("Part 1:", part1Result)

	part2Result := part2(data)
	fmt.Println("Part 2:", part2Result)

}

func part1(data []string) int {
	antennasByFrequency := passInput(data)
	var antinodes map[[2]int]bool = make(map[[2]int]bool)
	for _, locations := range antennasByFrequency {
		for i, pos1 := range locations {
			for _, pos2 := range locations[i+1:] {
				deltaX := pos2[0] - pos1[0]
				deltaY := pos2[1] - pos1[1]
				candidate1 := [2]int{pos1[0] - deltaX, pos1[1] - deltaY}
				candidate2 := [2]int{pos2[0] + deltaX, pos2[1] + deltaY}
				antinodes[candidate1] = true
				antinodes[candidate2] = true
			}

		}
	}
	total := 0
	// sum antinodes in map
	n, m := len(data), len(data[0])
	for node := range antinodes {
		if node[0] >= 0 && node[0] < n && node[1] >= 0 && node[1] < m {
			total++
		}
	}
	return total
}

func part2(data []string) int {
	antennasByFrequency := passInput(data)
	var antinodes map[[2]int]bool = make(map[[2]int]bool)
	n, m := len(data), len(data[0])

	for _, locations := range antennasByFrequency {
		for i, pos1 := range locations {
			for _, pos2 := range locations[i+1:] {
				deltaX := pos2[0] - pos1[0]
				deltaY := pos2[1] - pos1[1]

				antinodes[pos1] = true
				antinodes[pos2] = true
				pos := pos1
				for {
					pos = [2]int{pos[0] - deltaX, pos[1] - deltaY}
					if pos[0] < 0 || pos[0] >= n || pos[1] < 0 || pos[1] >= m {
						break
					}
					antinodes[pos] = true
				}
				pos = pos1
				for {
					pos = [2]int{pos[0] + deltaX, pos[1] + deltaY}
					if pos[0] < 0 || pos[0] >= n || pos[1] < 0 || pos[1] >= m {
						break
					}
					antinodes[pos] = true
				}
			}

		}
	}
	total := 0
	for range antinodes {
		total++
	}
	return total
}

func passInput(data []string) map[rune][][2]int {
	antennasByFrequency := make(map[rune][][2]int)
	for i, line := range data {
		for j, char := range line {
			if char != '.' {
				antennasByFrequency[char] = append(antennasByFrequency[char], [2]int{i, j})
			}
		}
	}
	return antennasByFrequency
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
